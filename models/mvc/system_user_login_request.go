package mvc

import (
	"net/mail"

	"github.com/LGYtech/lgo"
)

type SystemUserLoginRequest struct {
	Email    string `json:"e"`
	Password string `json:"p"`
}

func (model *SystemUserLoginRequest) Validate() *lgo.OperationResult {

	if len(model.Password) == 0 {
		return lgo.NewLogicError("Giriş yapabilmeniz için Parolanızı girmeniz gereklidir", nil)
	}
	if len(model.Email) == 0 {
		return lgo.NewLogicError("Giriş yapabilmeniz için Email girmeniz gereklidir", nil)
	}

	// #region Check If Email Valid if Exists
	_, err := mail.ParseAddress(model.Email)
	if err != nil {
		return lgo.NewLogicError("Email adresi doğrulanamadı", nil)
	}
	// #endregion Check If Email Valid if Exists

	return lgo.NewSuccess(nil)
}
