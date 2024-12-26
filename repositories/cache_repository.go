package repositories

import (
	"log"
	"sync"
	"time"

	"lms-web-services-main/database/datasources"
	"lms-web-services-main/models"
	datamodels "lms-web-services-main/models/data"
	mvcmodels "lms-web-services-main/models/mvc"

	"github.com/LGYtech/lgo"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

var (
	CacheRepository           cacheRepositoryInterface = &cacheRepository{}
	getSystemUserSettingMutex sync.Mutex
)

type cacheRepositoryInterface interface {
	AuthenticateSystemUser(token string) *lgo.OperationResult
	GetSystemUserCredential(token string) *lgo.OperationResult
	RegisterSystemUserCredential(token string, systemUser *datamodels.SystemUser) *lgo.OperationResult
	DeleteSystemUserCredential(token string) *lgo.OperationResult
	DeleteSystemUserCredentialById(id uuid.UUID) *lgo.OperationResult
	GetSystemUserSetting(c *models.Context, setting string) *lgo.OperationResult
	RemoveSystemUserSetting(systemUserId uuid.UUID, key string) *lgo.OperationResult
}

type cacheRepository struct{}

// #region Authenticate System User
func (r *cacheRepository) AuthenticateSystemUser(token string) *lgo.OperationResult {
	exists, err := datasources.Cache.Exists("su:" + token).Result()
	if err != nil {
		return lgo.NewFailureWithError(err)
	}
	return lgo.NewSuccess(exists == 1)
}

// #endregion Authenticate System User

// #region Get System User Credential
func (r *cacheRepository) GetSystemUserCredential(token string) *lgo.OperationResult {
	fields, err := datasources.Cache.HGetAll("su:" + token).Result()
	if err != nil {
		return lgo.NewFailureWithError(err)
	}

	if len(fields) == 0 {
		return lgo.NewFailure()
	}

	systemUserCredential := &mvcmodels.SystemUserCredential{
		Id:      fields["id"],
		Name:    fields["n"],
		Surname: fields["sn"],
		Email:   fields["e"],
	}

	return lgo.NewSuccess(systemUserCredential)
}

// #endregion Get System User Credential

// #region Register System User Credential
func (r *cacheRepository) RegisterSystemUserCredential(token string, systemUser *datamodels.SystemUser) *lgo.OperationResult {
	pipe := datasources.Cache.Pipeline()

	result := pipe.HSet("su:"+token, "id", systemUser.Id.String())
	if result.Err() != nil {
		return lgo.NewFailureWithError(result.Err())
	}

	result = pipe.HSet("su:"+token, "n", systemUser.Name)
	if result.Err() != nil {
		return lgo.NewFailureWithError(result.Err())
	}

	result = pipe.HSet("su:"+token, "sn", systemUser.Surname)
	if result.Err() != nil {
		return lgo.NewFailureWithError(result.Err())
	}

	result = pipe.HSet("su:"+token, "e", systemUser.Email)
	if result.Err() != nil {
		return lgo.NewFailureWithError(result.Err())
	}

	result = pipe.Expire("su:"+token, 5*time.Hour)
	if result.Err() != nil {
		return lgo.NewFailureWithError(result.Err())
	}

	listResult := pipe.LPush("su:rev:"+systemUser.Id.String(), token)
	if listResult.Err() != nil {
		return lgo.NewFailureWithError(listResult.Err())
	}

	result = pipe.Expire("su:rev:"+systemUser.Id.String(), 5*time.Hour)
	if result.Err() != nil {
		return lgo.NewFailureWithError(result.Err())
	}

	_, err := pipe.Exec()
	if err != nil {
		return lgo.NewFailureWithError(err)
	}

	return lgo.NewSuccess(nil)
}

// #endregion Register System User Credential

// #region Delete System User Credential
func (r *cacheRepository) DeleteSystemUserCredential(token string) *lgo.OperationResult {
	_, err := datasources.Cache.Del("su:" + token).Result()
	if err != nil {
		return lgo.NewFailureWithError(err)
	}

	return lgo.NewSuccess(nil)
}

// #endregion Delete System User Credential

// #region Delete System User Credential By Id
func (r *cacheRepository) DeleteSystemUserCredentialById(id uuid.UUID) *lgo.OperationResult {
	tokens, err := datasources.Cache.LRange("su:rev:"+id.String(), 0, -1).Result()
	if err != nil {
		return lgo.NewFailureWithError(err)
	}

	pipe := datasources.Cache.Pipeline()
	for _, token := range tokens {
		pipe.Del("su:" + token)
	}
	pipe.Del("su:rev:" + id.String())

	_, err = pipe.Exec()
	if err != nil {
		return lgo.NewFailureWithError(err)
	}

	return lgo.NewSuccess(nil)
}

// #endregion Delete System User Credential By Id

// #region Get System User Setting
func (r *cacheRepository) GetSystemUserSetting(c *models.Context, setting string) *lgo.OperationResult {
	// #region Get SystemUserCredential
	systemUserCredentialResult := r.GetSystemUserCredential(c.Token)
	if !systemUserCredentialResult.IsSuccess() {
		return systemUserCredentialResult
	}
	systemUserCredential := systemUserCredentialResult.ReturnObject.(*mvcmodels.SystemUserCredential)
	systemUserIdParsed, err := uuid.Parse(systemUserCredential.Id)
	if err != nil {
		return lgo.NewFailureWithError(err)
	}
	// #endregion Get SystemUserCredential

	// #region Try Get From Cache
	settingValueExists := true
	redisKey := "sus:" + systemUserCredential.Id                            // Redis Hash anahtarı
	settingValue, err := datasources.Cache.HGet(redisKey, setting).Result() // HGet ile Hash alt anahtarı okuyoruz
	if err == redis.Nil {
		settingValueExists = false
	} else if err != nil {
		return lgo.NewFailureWithError(err)
	}
	// #endregion Try Get From Cache

	// #region Lock and Try Cache Again
	if !settingValueExists {
		getSystemUserSettingMutex.Lock()
		defer getSystemUserSettingMutex.Unlock()

		// Tekrar Cache'ten Deniyoruz
		settingValue, err = datasources.Cache.HGet(redisKey, setting).Result()
		if err == redis.Nil {
			settingValueExists = false
		} else if err != nil {
			return lgo.NewFailureWithError(err)
		}

		if !settingValueExists {
			// #region Get From SystemUserSetting Repository
			var systemUserSettingRepository = NewSystemUserSettingRepository(datasources.Database)
			systemUserSettingResult := systemUserSettingRepository.GetValue(c, systemUserIdParsed, setting)
			if !systemUserSettingResult.IsSuccess() {
				return systemUserSettingResult
			}

			// Veritabanından Gelen Değeri İşliyoruz
			settingValue, ok := systemUserSettingResult.ReturnObject.(string)
			if !ok {
				return lgo.NewLogicError("Hatalı veri türü", nil)
			}
			// #endregion Get From SystemUserSetting Repository

			// #region Populate Cache
			err = datasources.Cache.HSet(redisKey, setting, settingValue).Err()
			if err != nil {
				log.Println("Redis Cache Güncellenemedi:", err)
			}
			// #endregion Populate Cache
		}
	}
	// #endregion Lock and Try Cache Again

	return lgo.NewSuccess(settingValue)
}

// #endregion Get System User Setting

// #region Remove System User Setting
func (r *cacheRepository) RemoveSystemUserSetting(systemUserId uuid.UUID, key string) *lgo.OperationResult {
	cacheKey := "sus:" + systemUserId.String() + ":" + key
	_, err := datasources.Cache.Del(cacheKey).Result()
	if err != nil {
		return lgo.NewFailureWithError(err)
	}

	return lgo.NewSuccess(nil)
}

// #endregion Remove System User Setting
