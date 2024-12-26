package controllers

import (
	"net/http"
	"strconv"

	"lms-web-services-main/models"
	datamodels "lms-web-services-main/models/data"
	"lms-web-services-main/services"

	"github.com/LGYtech/lgo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SystemUserSettingController struct {
	service services.SystemUserSettingService
}

func NewSystemUserSettingController(service services.SystemUserSettingService) *SystemUserSettingController {
	return &SystemUserSettingController{service: service}
}

// #region GetByUserId
func (ctrl *SystemUserSettingController) GetByUserId(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz kullanıcı ID formatı.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.GetByUserId(userId, context)
	c.JSON(http.StatusOK, result)
}

//#endregion GetByUserId

// #region GetById
func (ctrl *SystemUserSettingController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz ID formatı.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.GetById(id, context)
	c.JSON(http.StatusOK, result)
}

//#endregion GetById

// #region Set
func (ctrl *SystemUserSettingController) Set(c *gin.Context) {
	var setting datamodels.SystemUserSetting
	if err := c.ShouldBindJSON(&setting); err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Veri doğrulama hatası: "+err.Error(), nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Set(&setting, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Set

// #region Delete
func (ctrl *SystemUserSettingController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz ID formatı.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Delete(id, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Delete

// #region GetValue
func (ctrl *SystemUserSettingController) GetValue(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz kullanıcı ID formatı.", nil))
		return
	}

	key := c.Query("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Anahtar (key) zorunludur.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.GetValue(userId, key, context)
	c.JSON(http.StatusOK, result)
}

//#endregion GetValue
