package services

import (
	"lms-web-services-main/models"
	"lms-web-services-main/models/data"

	"github.com/LGYtech/lgo"
)

type TimingRuleHandler interface {
	Handle(model *data.Timing, c *models.Context) *lgo.OperationResult
	SetNext(handler TimingRuleHandler) TimingRuleHandler
}

type BaseTimingRuleHandler struct {
	next TimingRuleHandler
}

func (h *BaseTimingRuleHandler) SetNext(next TimingRuleHandler) TimingRuleHandler {
	h.next = next
	return h
}

func (h *BaseTimingRuleHandler) Handle(model *data.Timing, c *models.Context) *lgo.OperationResult {
	if h.next != nil {
		return h.next.Handle(model, c)
	}
	return lgo.NewSuccess(nil)
}

//#region Validation Handler

type TimingRuleHandlerValidation struct {
	BaseTimingRuleHandler
}

func (h *TimingRuleHandlerValidation) Handle(model *data.Timing, c *models.Context) *lgo.OperationResult {
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

//#region Update Validation Handler

type TimingRuleHandlerUpdateValidation struct {
	BaseTimingRuleHandler
}

func (h *TimingRuleHandlerUpdateValidation) Handle(model *data.Timing, c *models.Context) *lgo.OperationResult {
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

//#region Alter Authorization Handler

type TimingRuleHandlerCheckAlterAuthorization struct {
	BaseTimingRuleHandler
}

func (h *TimingRuleHandlerCheckAlterAuthorization) Handle(model *data.Timing, c *models.Context) *lgo.OperationResult {
	var permissionKey string
	if model.Id == 0 {
		permissionKey = data.TIMINGS_ADD
	} else {
		permissionKey = data.TIMINGS_UPDATE
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

//#region Read Authorization Handler

type TimingRuleHandlerCheckReadAuthorization struct {
	BaseTimingRuleHandler
}

func (h *TimingRuleHandlerCheckReadAuthorization) Handle(model *data.Timing, c *models.Context) *lgo.OperationResult {
	result := CacheService.GetSystemUserSetting(c, data.TIMINGS_VIEW)
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

type TimingRuleHandlerCheckDeleteAuthorization struct {
	BaseTimingRuleHandler
}

func (h *TimingRuleHandlerCheckDeleteAuthorization) Handle(model *data.Timing, c *models.Context) *lgo.OperationResult {
	result := CacheService.GetSystemUserSetting(c, data.TIMINGS_DELETE)
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
