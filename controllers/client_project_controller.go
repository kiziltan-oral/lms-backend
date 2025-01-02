package controllers

import (
	"net/http"
	"strconv"

	"lms-web-services-main/models"
	datamodels "lms-web-services-main/models/data"
	"lms-web-services-main/models/mvc"
	"lms-web-services-main/services"

	"github.com/LGYtech/lgo"
	"github.com/gin-gonic/gin"
)

type ClientProjectController struct {
	service services.ClientProjectService
}

func NewClientProjectController(service services.ClientProjectService) *ClientProjectController {
	return &ClientProjectController{service: service}
}

// #region Create ClientProject
func (ctrl *ClientProjectController) Create(c *gin.Context) {
	var clientProject datamodels.ClientProject
	if err := c.ShouldBindJSON(&clientProject); err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Veri doğrulama hatası: "+err.Error(), nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Create(&clientProject, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Create ClientProject

// #region Update ClientProject
func (ctrl *ClientProjectController) Update(c *gin.Context) {
	var clientProject datamodels.ClientProject
	if err := c.ShouldBindJSON(&clientProject); err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Veri doğrulama hatası: "+err.Error(), nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Update(&clientProject, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Update ClientProject

// #region Delete ClientProject
func (ctrl *ClientProjectController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz ID formatı.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Delete(id, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Delete ClientProject

// #region Get ClientProject By Id
func (ctrl *ClientProjectController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz ID formatı.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.GetById(id, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Get ClientProject By Id

// #region Get All ClientProjects
func (ctrl *ClientProjectController) GetAll(c *gin.Context) {
	// İstekten QueryModel'i oluştur
	var query mvc.QueryModel
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Veri doğrulama hatası: "+err.Error(), nil))
		return
	}
	context := models.NewContext(c)
	result := ctrl.service.GetAll(&query, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Get All ClientProjects

// #region Get ClientProjects By ClientId
func (ctrl *ClientProjectController) GetByClientId(c *gin.Context) {
	clientId, err := strconv.Atoi(c.Param("clientId"))
	if err != nil || clientId <= 0 {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz Client ID formatı.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.GetByClientId(clientId, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Get ClientProjects By ClientId
