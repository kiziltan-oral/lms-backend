package services

import (
	"lms-web-services-main/models"
	"lms-web-services-main/models/data"

	"github.com/LGYtech/lgo"
	"github.com/google/uuid"
)

// #region System User Setting Rule Handler Interface
type SystemUserSettingRuleHandler interface {
	Handle(model *data.SystemUserSetting, c *models.Context) *lgo.OperationResult
	SetNext(handler SystemUserSettingRuleHandler) SystemUserSettingRuleHandler
}

//#endregion System User Setting Rule Handler Interface

// #region Base System User Setting Rule Handler
type BaseSystemUserSettingRuleHandler struct {
	next SystemUserSettingRuleHandler
}

func (h *BaseSystemUserSettingRuleHandler) SetNext(next SystemUserSettingRuleHandler) SystemUserSettingRuleHandler {
	h.next = next
	return next
}

func (h *BaseSystemUserSettingRuleHandler) Handle(model *data.SystemUserSetting, c *models.Context) *lgo.OperationResult {
	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Base System User Setting Rule Handler

// #region Data Integrity Handler
type SystemUserSettingRuleHandlerDataIntegrity struct {
	BaseSystemUserSettingRuleHandler
}

func (h *SystemUserSettingRuleHandlerDataIntegrity) Handle(model *data.SystemUserSetting, c *models.Context) *lgo.OperationResult {
	if model.Key == "" || model.Value == "" {
		return lgo.NewLogicError("Key ve Value alanları zorunludur.", nil)
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Data Integrity Handler

// #region Validation Handler
type SystemUserSettingRuleHandlerValidation struct {
	BaseSystemUserSettingRuleHandler
}

func (h *SystemUserSettingRuleHandlerValidation) Handle(model *data.SystemUserSetting, c *models.Context) *lgo.OperationResult {
	if model.SystemUserId == uuid.Nil {
		return lgo.NewLogicError("SystemUserId boş olamaz.", nil)
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Validation Handler

// #region Alter Authorization Handler
type SystemUserSettingRuleHandlerCheckAlterAuthorization struct {
	BaseSystemUserSettingRuleHandler
}

func (h *SystemUserSettingRuleHandlerCheckAlterAuthorization) Handle(model *data.SystemUserSetting, c *models.Context) *lgo.OperationResult {
	var permissionKey string
	if model.Id == 0 {
		permissionKey = data.SYSTEM_SETTINGS_ADD
	} else {
		permissionKey = data.SYSTEM_SETTINGS_UPDATE
	}

	result := CacheService.GetSystemUserSetting(c, permissionKey)
	if !result.IsSuccess() {
		return result
	}

	if result.ReturnObject.(string) != "1" {
		return lgo.NewLogicError("Bu işlem için yetkiniz yok.", nil)
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Alter Authorization Handler

// #region Read Authorization Handler
type SystemUserSettingRuleHandlerCheckReadAuthorization struct {
	BaseSystemUserSettingRuleHandler
}

func (h *SystemUserSettingRuleHandlerCheckReadAuthorization) Handle(model *data.SystemUserSetting, c *models.Context) *lgo.OperationResult {
	result := CacheService.GetSystemUserSetting(c, data.SYSTEM_SETTINGS_VIEW)
	if !result.IsSuccess() {
		return result
	}

	if result.ReturnObject.(string) != "1" {
		return lgo.NewLogicError("Bu kaydı görüntülemek için yetkiniz yok.", nil)
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Read Authorization Handler

// #region Delete Authorization Handler
type SystemUserSettingRuleHandlerCheckDeleteAuthorization struct {
	BaseSystemUserSettingRuleHandler
}

func (h *SystemUserSettingRuleHandlerCheckDeleteAuthorization) Handle(model *data.SystemUserSetting, c *models.Context) *lgo.OperationResult {
	result := CacheService.GetSystemUserSetting(c, data.SYSTEM_SETTINGS_DELETE)
	if !result.IsSuccess() {
		return result
	}

	if result.ReturnObject.(string) != "1" {
		return lgo.NewLogicError("Bu kaydı silmek için yetkiniz yok.", nil)
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Delete Authorization Handler
