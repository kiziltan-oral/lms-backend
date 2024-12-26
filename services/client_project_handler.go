package services

import (
	"lms-web-services-main/models"
	"lms-web-services-main/models/data"

	"github.com/LGYtech/lgo"
)

type ClientProjectRuleHandler interface {
	Handle(model *data.ClientProject, c *models.Context) *lgo.OperationResult
	SetNext(handler ClientProjectRuleHandler) ClientProjectRuleHandler
}

type BaseClientProjectRuleHandler struct {
	next ClientProjectRuleHandler
}

func (h *BaseClientProjectRuleHandler) SetNext(handler ClientProjectRuleHandler) ClientProjectRuleHandler {
	h.next = handler
	return handler
}

func (h *BaseClientProjectRuleHandler) Handle(model *data.ClientProject, c *models.Context) *lgo.OperationResult {
	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

// #region Validation Handler
type ClientProjectRuleHandlerValidation struct {
	BaseClientProjectRuleHandler
}

func (h *ClientProjectRuleHandlerValidation) Handle(model *data.ClientProject, c *models.Context) *lgo.OperationResult {
	err := model.Validate()
	if err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Validation Handler

// #region Update Validation Handler
type ClientProjectRuleHandlerUpdateValidation struct {
	BaseClientProjectRuleHandler
}

func (h *ClientProjectRuleHandlerUpdateValidation) Handle(model *data.ClientProject, c *models.Context) *lgo.OperationResult {
	err := model.ValidateForUpdate()
	if err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Update Validation Handler

// #region Alter Authorization Handler
type ClientProjectRuleHandlerCheckAlterAuthorization struct {
	BaseClientProjectRuleHandler
}

func (h *ClientProjectRuleHandlerCheckAlterAuthorization) Handle(model *data.ClientProject, c *models.Context) *lgo.OperationResult {
	var permissionKey string
	if model.Id == 0 {
		permissionKey = data.CLIENTPROJECTS_ADD
	} else {
		permissionKey = data.CLIENTPROJECTS_UPDATE
	}

	result := CacheService.GetSystemUserSetting(c, permissionKey)
	if !result.IsSuccess() {
		return result
	}
	if result.ReturnObject.(string) != "1" {
		result = lgo.NewLogicError("Bu işlem için yetkiniz yok.", nil)
		return result
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#endregion Alter Authorization Handler

//#region Read Authorization Handler

type ClientProjectRuleHandlerCheckReadAuthorization struct {
	BaseClientProjectRuleHandler
}

func (h *ClientProjectRuleHandlerCheckReadAuthorization) Handle(model *data.ClientProject, c *models.Context) *lgo.OperationResult {
	result := CacheService.GetSystemUserSetting(c, data.CLIENTPROJECTS_VIEW)
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

//#region Delete Authorization Handler

type ClientProjectRuleHandlerCheckDeleteAuthorization struct {
	BaseClientProjectRuleHandler
}

func (h *ClientProjectRuleHandlerCheckDeleteAuthorization) Handle(model *data.ClientProject, c *models.Context) *lgo.OperationResult {
	result := CacheService.GetSystemUserSetting(c, data.CLIENTPROJECTS_DELETE)
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
