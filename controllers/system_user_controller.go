package controllers

import (
	"net/http"

	"lms-web-services-main/models"
	datamodels "lms-web-services-main/models/data"
	mvc "lms-web-services-main/models/mvc"
	"lms-web-services-main/services"

	"github.com/LGYtech/lgo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SystemUserController struct {
	service services.SystemUserService
}

func NewSystemUserController(service services.SystemUserService) *SystemUserController {
	return &SystemUserController{service: service}
}

// #region Create System User
func (ctrl *SystemUserController) Create(c *gin.Context) {
	var systemUser datamodels.SystemUser
	if err := c.ShouldBindJSON(&systemUser); err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Veri doğrulama hatası: "+err.Error(), nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Create(&systemUser, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Create System User

// #region Update System User
func (ctrl *SystemUserController) Update(c *gin.Context) {
	var systemUser datamodels.SystemUser
	if err := c.ShouldBindJSON(&systemUser); err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Veri doğrulama hatası: "+err.Error(), nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Update(&systemUser, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Update System User

// #region Delete System User
func (ctrl *SystemUserController) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil || id == uuid.Nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz ID formatı.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Delete(id, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Delete System User

// #region Get System User By Id
func (ctrl *SystemUserController) GetById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil || id == uuid.Nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz ID formatı.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.GetById(id, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Get System User By Id

// #region Get All System Users
func (ctrl *SystemUserController) GetAll(c *gin.Context) {
	var query mvc.QueryModel
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Veri doğrulama hatası: "+err.Error(), nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.GetAll(&query, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Get All System Users

// #region Get System User By Email
func (ctrl *SystemUserController) GetByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("E-posta adresi zorunludur.", nil))
		return
	}

	result := ctrl.service.GetByEmail(email)
	c.JSON(http.StatusOK, result)
}

//#endregion Get System User By Email

// #region Login
func (ctrl *SystemUserController) Login(c *gin.Context) {
	var loginRequest mvc.SystemUserLoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Giriş verisi doğrulama hatası: "+err.Error(), nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Login(context, &loginRequest)
	c.JSON(http.StatusOK, result)
}

//#endregion Login

// #region Logout
func (ctrl *SystemUserController) Logout(c *gin.Context) {
	token := c.GetHeader("X-Token")
	if token == "" {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Token eksik.", nil))
		return
	}

	result := ctrl.service.Logout(token)
	c.JSON(http.StatusOK, result)
}

//#endregion Logout
