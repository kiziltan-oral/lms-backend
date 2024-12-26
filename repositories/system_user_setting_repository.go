package repositories

import (
	"errors"

	"github.com/LGYtech/lgo"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"lms-web-services-main/models"
	datamodels "lms-web-services-main/models/data"
)

type SystemUserSettingRepository interface {
	GetByUserId(systemUserId uuid.UUID) *lgo.OperationResult
	GetById(id int) *lgo.OperationResult
	Set(setting *datamodels.SystemUserSetting) *lgo.OperationResult
	Delete(id int) *lgo.OperationResult
	GetValue(c *models.Context, systemUserId uuid.UUID, key string) *lgo.OperationResult
}

type systemUserSettingRepository struct {
	db *gorm.DB
}

func NewSystemUserSettingRepository(db *gorm.DB) SystemUserSettingRepository {
	return &systemUserSettingRepository{db: db}
}

// #region GetByUserId
func (r *systemUserSettingRepository) GetByUserId(systemUserId uuid.UUID) *lgo.OperationResult {
	var settings []*datamodels.SystemUserSetting
	result := r.db.Where("\"SystemUserId\" = ?", systemUserId).Find(&settings)
	if result.Error != nil {
		return lgo.NewLogicError(result.Error.Error(), nil)
	}
	return lgo.NewSuccess(settings)
}

// #endregion GetByUserId

// #region GetById
func (r *systemUserSettingRepository) GetById(id int) *lgo.OperationResult {
	var setting datamodels.SystemUserSetting
	result := r.db.First(&setting, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return lgo.NewLogicError("Kayıt bulunamadı.", nil)
		}
		return lgo.NewLogicError(result.Error.Error(), nil)
	}
	return lgo.NewSuccess(setting)
}

// #endregion GetById

// #region Set
func (r *systemUserSettingRepository) Set(setting *datamodels.SystemUserSetting) *lgo.OperationResult {
	var existingSetting datamodels.SystemUserSetting
	result := r.db.Where("\"SystemUserId\" = ? AND \"Key\" = ?", setting.SystemUserId, setting.Key).First(&existingSetting)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			if err := r.db.Create(setting).Error; err != nil {
				return lgo.NewFailureWithError(err)
			}
			return lgo.NewSuccess(setting)
		}
		return lgo.NewFailureWithError(result.Error)
	}

	existingSetting.Value = setting.Value
	saveResult := r.db.Save(&existingSetting)
	if saveResult.Error != nil {
		return lgo.NewFailureWithError(saveResult.Error)
	}
	if saveResult.RowsAffected == 0 {
		return lgo.NewLogicError("Kayıt güncellenemedi.", nil)
	}
	return lgo.NewSuccess(existingSetting)
}

// #endregion Set

// #region Delete
func (r *systemUserSettingRepository) Delete(id int) *lgo.OperationResult {
	result := r.db.Delete(&datamodels.SystemUserSetting{}, id)
	if result.Error != nil {
		return lgo.NewFailureWithError(result.Error)
	}
	if result.RowsAffected == 0 {
		return lgo.NewLogicError("Kayıt bulunamadı", nil)
	}
	return lgo.NewSuccess(nil)
}

// #endregion Delete

// #region GetValue
func (r *systemUserSettingRepository) GetValue(c *models.Context, systemUserId uuid.UUID, key string) *lgo.OperationResult {
	var setting datamodels.SystemUserSetting
	result := r.db.Where("\"SystemUserId\" = ? AND \"Key\" = ?", systemUserId, key).First(&setting)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return lgo.NewSuccess(nil)
		}
		return lgo.NewFailureWithError(result.Error)
	}
	return lgo.NewSuccess(setting)
}

// #endregion GetValue
