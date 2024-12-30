package services

import (
	"lms-web-services-main/database/datasources"
	"lms-web-services-main/models"
	datamodels "lms-web-services-main/models/data"
	mvc "lms-web-services-main/models/mvc"
	repositories "lms-web-services-main/repositories"
	utils "lms-web-services-main/utils"

	"github.com/LGYtech/lgo"
	"github.com/google/uuid"
)

type SystemUserService interface {
	Create(systemUser *datamodels.SystemUser, c *models.Context) *lgo.OperationResult
	Update(systemUser *datamodels.SystemUser, c *models.Context) *lgo.OperationResult
	Delete(id uuid.UUID, c *models.Context) *lgo.OperationResult
	GetById(id uuid.UUID, c *models.Context) *lgo.OperationResult
	GetByEmail(email string) *lgo.OperationResult
	GetAll(query *mvc.QueryModel, c *models.Context) *lgo.OperationResult
	CheckForeignReferences(systemUser *datamodels.SystemUser) *lgo.OperationResult
	CheckExistingSystemUser(systemUser *datamodels.SystemUser) *lgo.OperationResult
	Login(c *models.Context, request *mvc.SystemUserLoginRequest) *lgo.OperationResult
	Logout(token string) *lgo.OperationResult
}

type systemUserService struct {
	repo            repositories.SystemUserRepository
	saveRules       SystemUserRuleHandler
	deleteRules     SystemUserRuleHandler
	updateRules     SystemUserRuleHandler
	readRules       SystemUserRuleHandler
	postUpdateRules SystemUserRuleHandler
}

func NewSystemUserService(repo repositories.SystemUserRepository) SystemUserService {
	service := &systemUserService{
		repo: repo,
	}
	service.saveRules = (&SystemUserRuleHandlerValidation{}).
		SetNext(&SystemUserRuleHandlerCheckAlterAuthorization{}).
		SetNext(&SystemUserRuleHandlerDataIntegrity{
			SystemUserService: service,
		})

	service.deleteRules = (&SystemUserRuleHandlerCheckDeleteAuthorization{}).
		SetNext(&SystemUserRuleHandlerCheckForeignReferences{SystemUserService: service})
	service.updateRules = (&SystemUserRuleHandlerValidation{}).
		SetNext(&SystemUserRuleHandlerCheckAlterAuthorization{}).
		SetNext(&SystemUserRuleHandlerDataIntegrity{
			SystemUserService: service,
		})

	service.postUpdateRules = &SystemUserRuleHandlerIsActiveChange{
		SystemUserService: service,
	}
	service.readRules = &SystemUserRuleHandlerCheckReadAuthorization{}

	return service
}

// #region Create
func (s *systemUserService) Create(systemUser *datamodels.SystemUser, c *models.Context) *lgo.OperationResult {
	if systemUser.Email == "" {
		return lgo.NewLogicError("E-posta adresi zorunludur.", nil)
	}

	// Şifre tuzu oluştur ve şifreyi hash'le
	salt := utils.GenerateRandomNumeric(15)
	systemUser.PasswordSalt = salt
	hashedPassword := utils.ComputeSHA256(systemUser.Password, salt)
	systemUser.Password = hashedPassword

	// Kuralları çalıştır
	if result := s.saveRules.Handle(systemUser, c); !result.IsSuccess() {
		return result
	}

	// Kullanıcıyı veritabanına ekle
	createResult := s.repo.Create(systemUser)
	if !createResult.IsSuccess() {
		return createResult
	}

	// Varsayılan izinleri ata
	permissionResult := s.assignDefaultPermissions(systemUser.Id, c)
	if !permissionResult.IsSuccess() {
		return permissionResult
	}

	return createResult
}

func (s *systemUserService) assignDefaultPermissions(systemUserId uuid.UUID, c *models.Context) *lgo.OperationResult {
	// Varsayılan izinler
	defaultPermissions := []datamodels.SystemUserSetting{
		{SystemUserId: systemUserId, Key: datamodels.SYSTEM_USERS_VIEW, Value: "1", Description: "Sistem kullanıcılarını görüntüleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.SYSTEM_USERS_ADD, Value: "1", Description: "Sistem kullanıcılarını ekleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.SYSTEM_USERS_UPDATE, Value: "1", Description: "Sistem kullanıcılarını güncelleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.SYSTEM_USERS_DELETE, Value: "1", Description: "Sistem kullanıcılarını silme yetkisi"},

		{SystemUserId: systemUserId, Key: datamodels.SYSTEM_SETTINGS_VIEW, Value: "1", Description: "Sistem ayarlarını görüntüleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.SYSTEM_SETTINGS_ADD, Value: "1", Description: "Sistem ayarlarını ekleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.SYSTEM_SETTINGS_UPDATE, Value: "1", Description: "Sistem ayarlarını güncelleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.SYSTEM_SETTINGS_DELETE, Value: "1", Description: "Sistem ayarlarını silme yetkisi"},

		{SystemUserId: systemUserId, Key: datamodels.CLIENTS_VIEW, Value: "1", Description: "Müşterileri görüntüleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.CLIENTS_ADD, Value: "1", Description: "Müşterileri ekleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.CLIENTS_UPDATE, Value: "1", Description: "Müşterileri güncelleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.CLIENTS_DELETE, Value: "1", Description: "Müşterileri silme yetkisi"},

		{SystemUserId: systemUserId, Key: datamodels.CLIENTPROJECTS_VIEW, Value: "1", Description: "Müşteri projelerini görüntüleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.CLIENTPROJECTS_ADD, Value: "1", Description: "Müşteri projelerini ekleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.CLIENTPROJECTS_UPDATE, Value: "1", Description: "Müşteri projelerini güncelleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.CLIENTPROJECTS_DELETE, Value: "1", Description: "Müşteri projelerini silme yetkisi"},

		{SystemUserId: systemUserId, Key: datamodels.TIMINGS_VIEW, Value: "1", Description: "Zamanlamaları görüntüleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.TIMINGS_ADD, Value: "1", Description: "Zamanlamaları ekleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.TIMINGS_UPDATE, Value: "1", Description: "Zamanlamaları güncelleme yetkisi"},
		{SystemUserId: systemUserId, Key: datamodels.TIMINGS_DELETE, Value: "1", Description: "Zamanlamaları silme yetkisi"},
	}

	// Her bir izin için `Set` metodunu çağır
	for _, permission := range defaultPermissions {
		var systemUserSettingService = NewSystemUserSettingService(repositories.NewSystemUserSettingRepository(datasources.Database))
		result := systemUserSettingService.Set(&permission, c)
		if !result.IsSuccess() {
			return result // Hata varsa işlemi durdur
		}
	}

	return lgo.NewSuccess(nil) // Tüm izinler başarıyla eklendi
}

//#endregion Create

// #region Update
func (s *systemUserService) Update(systemUser *datamodels.SystemUser, c *models.Context) *lgo.OperationResult {
	if systemUser.Id == uuid.Nil {
		return lgo.NewLogicError("Geçersiz kullanıcı ID.", nil)
	}

	if systemUser.Password != "" {
		salt := utils.GenerateRandomAlphaNumeric(15)
		systemUser.PasswordSalt = salt
		hashedPassword := utils.ComputeSHA256(systemUser.Password, salt)
		systemUser.Password = hashedPassword
	}

	if result := s.updateRules.Handle(systemUser, c); !result.IsSuccess() {
		return result
	}

	updateResult := s.repo.Update(systemUser)
	if !updateResult.IsSuccess() {
		return updateResult
	}

	return s.postUpdateRules.Handle(systemUser, c)
}

//#endregion Update

// #region Delete
func (s *systemUserService) Delete(id uuid.UUID, c *models.Context) *lgo.OperationResult {
	if id == uuid.Nil {
		return lgo.NewLogicError("Geçersiz kullanıcı ID.", nil)
	}

	systemUser := &datamodels.SystemUser{Id: id}
	result := s.deleteRules.Handle(systemUser, c)
	if !result.IsSuccess() {
		return result
	}

	deleteResult := s.repo.Delete(id)
	if !deleteResult.IsSuccess() {
		return deleteResult
	}

	return CacheService.DeleteSystemUserCredentialById(id)
}

//#endregion Delete

// #region GetById
func (s *systemUserService) GetById(id uuid.UUID, c *models.Context) *lgo.OperationResult {
	if id == uuid.Nil {
		return lgo.NewLogicError("Geçersiz kullanıcı ID.", nil)
	}

	if result := s.readRules.Handle(&datamodels.SystemUser{Id: id}, c); !result.IsSuccess() {
		return result
	}

	return s.repo.GetById(id)
}

//#endregion GetById

// #region GetByEmail
func (s *systemUserService) GetByEmail(email string) *lgo.OperationResult {
	if email == "" {
		return lgo.NewLogicError("E-posta adresi zorunludur.", nil)
	}

	return s.repo.GetByEmail(email)
}

//#endregion GetByEmail

// #region GetAll
func (s *systemUserService) GetAll(query *mvc.QueryModel, c *models.Context) *lgo.OperationResult {
	if result := query.Validate(); !result.IsSuccess() {
		return lgo.NewLogicError("Geçersiz sorgu parametreleri: "+result.ErrorMessage, nil)
	}
	return s.repo.GetAll(query)
}

//#endregion GetAll

// #region Check Foreign References
func (s *systemUserService) CheckForeignReferences(systemUser *datamodels.SystemUser) *lgo.OperationResult {
	return s.repo.CheckForeignReferences(systemUser)
}

//#endregion Check Foreign References

// #region Check Existing SystemUser
func (s *systemUserService) CheckExistingSystemUser(systemUser *datamodels.SystemUser) *lgo.OperationResult {
	return s.repo.CheckExistingSystemUser(systemUser)
}

//#endregion Check Existing SystemUser

// #region Login
func (s *systemUserService) Login(c *models.Context, request *mvc.SystemUserLoginRequest) *lgo.OperationResult {
	if result := request.Validate(); !result.IsSuccess() {
		return result
	}

	systemUserResult := s.GetByEmail(request.Email)
	if !systemUserResult.IsSuccess() {
		return systemUserResult
	}

	systemUser, ok := systemUserResult.ReturnObject.(*datamodels.SystemUser)
	if !ok {
		return lgo.NewFailure()
	}

	hashedRequestPassword := utils.ComputeSHA256(request.Password, systemUser.PasswordSalt)
	if hashedRequestPassword != systemUser.Password {
		return lgo.NewLogicError("Email veya Şifre hatalı.", nil)
	}

	uuidV4, err := uuid.NewRandom()
	if err != nil {
		return lgo.NewFailure()
	}
	systemUserToken := uuidV4.String()

	systemUserTokenResult := CacheService.RegisterSystemUserCredential(systemUserToken, systemUser)
	if !systemUserTokenResult.IsSuccess() {
		return systemUserTokenResult
	}

	c.Token = systemUserToken

	userData := map[string]string{
		"e": systemUser.Email,
		"n": systemUser.Name,
		"s": systemUser.Surname,
		"t": systemUserToken,
	}

	return lgo.NewSuccess(userData)
}

//#endregion Login

// #region Logout
func (s *systemUserService) Logout(token string) *lgo.OperationResult {
	return CacheService.DeleteSystemUserCredential(token)
}

//#endregion Logout
