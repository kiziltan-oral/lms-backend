package controllers

import (
	"net/http"
	"strconv"

	"lms-web-services-main/models"
	"lms-web-services-main/models/data"
	"lms-web-services-main/models/mvc"
	"lms-web-services-main/services"

	"github.com/LGYtech/lgo"
	"github.com/gin-gonic/gin"
)

// #region Client Controller Definition
type ClientController struct {
	service services.ClientService
}

func NewClientController(service services.ClientService) *ClientController {
	return &ClientController{service: service}
}

//#endregion Client Controller Definition

// #region Create Client
func (ctrl *ClientController) Create(c *gin.Context) {

	var client data.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Veri doğrulama hatası: "+err.Error(), nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Create(&client, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Create Client

// #region Update Client
func (ctrl *ClientController) Update(c *gin.Context) {

	var client data.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Veri doğrulama hatası: "+err.Error(), nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Update(&client, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Update Client

// #region Delete Client
func (ctrl *ClientController) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz ID formatı.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Delete(id, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Delete Client

// #region Get Client By Id
func (ctrl *ClientController) GetById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz ID formatı.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.GetById(id, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Get Client By Id

// #region Get All Clients
func (ctrl *ClientController) GetAll(c *gin.Context) {
	// İstekten QueryModel'i oluştur
	var query mvc.QueryModel
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Veri doğrulama hatası: "+err.Error(), nil))
		return
	}

	// Servis çağrısı
	context := models.NewContext(c)
	result := ctrl.service.GetAll(&query, context)

	// Sonuç döndür
	c.JSON(http.StatusOK, result)
}

//#endregion Get All Clients
