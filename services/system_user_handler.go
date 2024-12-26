package services

import (
	"lms-web-services-main/models"
	datamodels "lms-web-services-main/models/data"

	"github.com/LGYtech/lgo"
	"github.com/google/uuid"
)

type SystemUserRuleHandler interface {
	Handle(model *datamodels.SystemUser, c *models.Context) *lgo.OperationResult
	SetNext(handler SystemUserRuleHandler) SystemUserRuleHandler
}

type BaseSystemUserRuleHandler struct {
	next SystemUserRuleHandler
}

func (h *BaseSystemUserRuleHandler) SetNext(next SystemUserRuleHandler) SystemUserRuleHandler {
	h.next = next
	return h
}

func (h *BaseSystemUserRuleHandler) Handle(model *datamodels.SystemUser, c *models.Context) *lgo.OperationResult {
	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

// #region Data Integrity
type SystemUserRuleHandlerDataIntegrity struct {
	BaseSystemUserRuleHandler
	SystemUserService SystemUserService
}

func (h *SystemUserRuleHandlerDataIntegrity) Handle(model *datamodels.SystemUser, c *models.Context) *lgo.OperationResult {
	result := h.SystemUserService.CheckExistingSystemUser(model)
	if !result.IsSuccess() {
		return result
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Data Integrity

// #region Validation
type SystemUserRuleHandlerValidation struct {
	BaseSystemUserRuleHandler
}

func (h *SystemUserRuleHandlerValidation) Handle(model *datamodels.SystemUser, c *models.Context) *lgo.OperationResult {
	var err error
	if model.Id == uuid.Nil {
		err = model.Validate()
	} else {
		err = model.ValidateForUpdate()
	}
	if err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Validation

// #region Check Foreign References
type SystemUserRuleHandlerCheckForeignReferences struct {
	BaseSystemUserRuleHandler
	SystemUserService SystemUserService
}

func (h *SystemUserRuleHandlerCheckForeignReferences) Handle(model *datamodels.SystemUser, c *models.Context) *lgo.OperationResult {
	result := h.SystemUserService.CheckForeignReferences(model)
	if !result.IsSuccess() {
		return result
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Check Foreign References

// #region Alter Authorization
type SystemUserRuleHandlerCheckAlterAuthorization struct {
	BaseSystemUserRuleHandler
}

func (h *SystemUserRuleHandlerCheckAlterAuthorization) Handle(model *datamodels.SystemUser, c *models.Context) *lgo.OperationResult {
	var actionKey string

	if model.Id == uuid.Nil {
		actionKey = datamodels.SYSTEM_USERS_ADD
	} else {
		actionKey = datamodels.SYSTEM_SETTINGS_UPDATE
	}

	result := CacheService.GetSystemUserSetting(c, actionKey)
	if !result.IsSuccess() {
		return result
	}
	if result.ReturnObject.(string) != "1" {
		return lgo.NewLogicError("Unauthorized", nil)
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Alter Authorization

// #region IsActive Change
type SystemUserRuleHandlerIsActiveChange struct {
	BaseSystemUserRuleHandler
	SystemUserService SystemUserService
}

func (h *SystemUserRuleHandlerIsActiveChange) Handle(model *datamodels.SystemUser, c *models.Context) *lgo.OperationResult {
	existingResult := h.SystemUserService.GetById(model.Id, c)
	if !existingResult.IsSuccess() {
		return existingResult
	}

	existingUser := existingResult.ReturnObject.(*datamodels.SystemUser)

	if existingUser.IsActive != model.IsActive {
		if result := CacheService.DeleteSystemUserCredentialById(model.Id); !result.IsSuccess() {
			return result
		}
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion IsActive Change

// #region Read Authorization
type SystemUserRuleHandlerCheckReadAuthorization struct {
	BaseSystemUserRuleHandler
}

func (h *SystemUserRuleHandlerCheckReadAuthorization) Handle(model *datamodels.SystemUser, c *models.Context) *lgo.OperationResult {
	result := CacheService.GetSystemUserSetting(c, datamodels.SYSTEM_USERS_VIEW)
	if !result.IsSuccess() {
		return result
	}
	if result.ReturnObject.(string) != "1" {
		return lgo.NewLogicError("Unauthorized", nil)
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Read Authorization

// #region Delete Authorization
type SystemUserRuleHandlerCheckDeleteAuthorization struct {
	BaseSystemUserRuleHandler
}

func (h *SystemUserRuleHandlerCheckDeleteAuthorization) Handle(model *datamodels.SystemUser, c *models.Context) *lgo.OperationResult {
	result := CacheService.GetSystemUserSetting(c, datamodels.SYSTEM_USERS_DELETE)
	if !result.IsSuccess() {
		return result
	}
	if result.ReturnObject.(string) != "1" {
		return lgo.NewLogicError("Unauthorized", nil)
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Delete Authorization
