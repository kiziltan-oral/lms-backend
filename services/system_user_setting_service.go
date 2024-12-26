package services

import (
	"lms-web-services-main/models"
	datamodels "lms-web-services-main/models/data"
	repositories "lms-web-services-main/repositories"

	"github.com/LGYtech/lgo"
	"github.com/google/uuid"
)

// #region System User Setting Service Interface
type SystemUserSettingService interface {
	GetByUserId(systemUserId uuid.UUID, c *models.Context) *lgo.OperationResult
	GetById(id int, c *models.Context) *lgo.OperationResult
	Set(setting *datamodels.SystemUserSetting, c *models.Context) *lgo.OperationResult
	Delete(id int, c *models.Context) *lgo.OperationResult
	GetValue(systemUserId uuid.UUID, key string, c *models.Context) *lgo.OperationResult
}

//#endregion System User Setting Service Interface

// #region System User Setting Service Implementation
type systemUserSettingService struct {
	repo        repositories.SystemUserSettingRepository
	saveRules   SystemUserSettingRuleHandler
	deleteRules SystemUserSettingRuleHandler
	readRules   SystemUserSettingRuleHandler
}

func NewSystemUserSettingService(repo repositories.SystemUserSettingRepository) SystemUserSettingService {
	return &systemUserSettingService{
		repo: repo,
		saveRules: (&SystemUserSettingRuleHandlerValidation{}).
			SetNext(&SystemUserSettingRuleHandlerCheckAlterAuthorization{}).
			SetNext(&SystemUserSettingRuleHandlerDataIntegrity{}),
		deleteRules: &SystemUserSettingRuleHandlerCheckDeleteAuthorization{},
		readRules:   &SystemUserSettingRuleHandlerCheckReadAuthorization{},
	}
}

//#endregion System User Setting Service Implementation

// #region GetByUserId
func (s *systemUserSettingService) GetByUserId(systemUserId uuid.UUID, c *models.Context) *lgo.OperationResult {
	if systemUserId == uuid.Nil {
		return lgo.NewLogicError("Geçersiz kullanıcı ID.", nil)
	}

	if result := s.readRules.Handle(&datamodels.SystemUserSetting{SystemUserId: systemUserId}, c); !result.IsSuccess() {
		return result
	}

	return s.repo.GetByUserId(systemUserId)
}

//#endregion GetByUserId

// #region GetById
func (s *systemUserSettingService) GetById(id int, c *models.Context) *lgo.OperationResult {
	if id <= 0 {
		return lgo.NewLogicError("Geçersiz ID.", nil)
	}

	settingResult := s.repo.GetById(id)
	if !settingResult.IsSuccess() {
		return settingResult
	}

	setting := settingResult.ReturnObject.(*datamodels.SystemUserSetting)
	if result := s.readRules.Handle(setting, c); !result.IsSuccess() {
		return result
	}

	return settingResult
}

//#endregion GetById

// #region Set
func (s *systemUserSettingService) Set(setting *datamodels.SystemUserSetting, c *models.Context) *lgo.OperationResult {
	if setting.SystemUserId == uuid.Nil {
		return lgo.NewLogicError("Geçersiz kullanıcı ID.", nil)
	}
	if setting.Key == "" {
		return lgo.NewLogicError("Anahtar (key) alanı zorunludur.", nil)
	}
	if setting.Value == "" {
		return lgo.NewLogicError("Değer (value) alanı zorunludur.", nil)
	}

	if result := s.saveRules.Handle(setting, c); !result.IsSuccess() {
		return result
	}

	return s.repo.Set(setting)
}

//#endregion Set

// #region Delete
func (s *systemUserSettingService) Delete(id int, c *models.Context) *lgo.OperationResult {
	if id <= 0 {
		return lgo.NewLogicError("Geçersiz ID.", nil)
	}

	settingResult := s.repo.GetById(id)
	if !settingResult.IsSuccess() {
		return settingResult
	}

	setting := settingResult.ReturnObject.(*datamodels.SystemUserSetting)
	if result := s.deleteRules.Handle(setting, c); !result.IsSuccess() {
		return result
	}

	return s.repo.Delete(id)
}

//#endregion Delete

// #region GetValue
func (s *systemUserSettingService) GetValue(systemUserId uuid.UUID, key string, c *models.Context) *lgo.OperationResult {
	if systemUserId == uuid.Nil {
		return lgo.NewLogicError("Geçersiz kullanıcı ID.", nil)
	}
	if key == "" {
		return lgo.NewLogicError("Anahtar (key) alanı zorunludur.", nil)
	}

	if result := s.readRules.Handle(&datamodels.SystemUserSetting{SystemUserId: systemUserId, Key: key}, c); !result.IsSuccess() {
		return result
	}

	return s.repo.GetValue(c, systemUserId, key)
}

//#endregion GetValue
