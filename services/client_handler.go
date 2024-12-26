package services

import (
	"lms-web-services-main/models"
	"lms-web-services-main/models/data"

	"github.com/LGYtech/lgo"
)

type ClientRuleHandler interface {
	Handle(model *data.Client, c *models.Context) *lgo.OperationResult
	SetNext(handler ClientRuleHandler) ClientRuleHandler
}

type BaseClientRuleHandler struct {
	next ClientRuleHandler
}

func (h *BaseClientRuleHandler) SetNext(next ClientRuleHandler) ClientRuleHandler {
	h.next = next
	return h
}

func (h *BaseClientRuleHandler) Handle(model *data.Client, c *models.Context) *lgo.OperationResult {
	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

// #region Validation Handler

type ClientRuleHandlerValidation struct {
	BaseClientRuleHandler
}

func (h *ClientRuleHandlerValidation) Handle(model *data.Client, c *models.Context) *lgo.OperationResult {
	// Model doğrulaması
	err := model.Validate()
	if err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

// #endregion Validation Handler

// #region Update Validation Handler

type ClientRuleHandlerUpdateValidation struct {
	BaseClientRuleHandler
}

func (h *ClientRuleHandlerUpdateValidation) Handle(model *data.Client, c *models.Context) *lgo.OperationResult {
	// Güncelleme doğrulaması
	err := model.ValidateForUpdate()
	if err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

// #endregion Update Validation Handler

// #region Alter Authorization Handler

type ClientRuleHandlerCheckAlterAuthorization struct {
	BaseClientRuleHandler
}

func (h *ClientRuleHandlerCheckAlterAuthorization) Handle(model *data.Client, c *models.Context) *lgo.OperationResult {
	// #region Yetki Kontrolü
	var permissionKey string
	if model.Id == 0 {
		permissionKey = data.CLIENTS_ADD
	} else {
		permissionKey = data.CLIENTS_UPDATE
	}

	result := CacheService.GetSystemUserSetting(c, permissionKey)
	if !result.IsSuccess() {
		return result
	}
	if result.ReturnObject.(string) != "1" {
		result = lgo.NewAutoError()
		result.ErrorMessage = permissionKey
		return result
	}
	// #endregion Yetki Kontrolü

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

// #endregion Alter Authorization Handler

// #region Read Authorization Handler

type ClientRuleHandlerCheckReadAuthorization struct {
	BaseClientRuleHandler
}

func (h *ClientRuleHandlerCheckReadAuthorization) Handle(model *data.Client, c *models.Context) *lgo.OperationResult {
	// #region Yetki Kontrolü
	result := CacheService.GetSystemUserSetting(c, data.CLIENTS_VIEW)
	if !result.IsSuccess() {
		return result
	}
	if result.ReturnObject.(string) != "1" {
		result = lgo.NewAutoError()
		result.ErrorMessage = data.CLIENTS_VIEW
		return result
	}
	// #endregion Yetki Kontrolü

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

// #endregion Read Authorization Handler

// #region Delete Authorization Handler

type ClientRuleHandlerCheckDeleteAuthorization struct {
	BaseClientRuleHandler
}

func (h *ClientRuleHandlerCheckDeleteAuthorization) Handle(model *data.Client, c *models.Context) *lgo.OperationResult {
	// #region Yetki Kontrolü
	result := CacheService.GetSystemUserSetting(c, data.CLIENTS_DELETE)
	if !result.IsSuccess() {
		return result
	}
	if result.ReturnObject.(string) != "1" {
		result = lgo.NewAutoError()
		result.ErrorMessage = data.CLIENTS_DELETE
		return result
	}
	// #endregion Yetki Kontrolü

	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

// #endregion Delete Authorization Handler
