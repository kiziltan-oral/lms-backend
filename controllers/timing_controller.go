package controllers

import (
	"net/http"
	"strconv"
	"time"

	"lms-web-services-main/models"
	datamodels "lms-web-services-main/models/data"
	"lms-web-services-main/models/mvc"
	"lms-web-services-main/services"

	"github.com/LGYtech/lgo"
	"github.com/gin-gonic/gin"
)

type TimingController struct {
	service services.TimingService
}

func NewTimingController(service services.TimingService) *TimingController {
	return &TimingController{service: service}
}

// #region Create Timing
func (ctrl *TimingController) Create(c *gin.Context) {
	var timing datamodels.Timing
	if err := c.ShouldBindJSON(&timing); err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Veri doğrulama hatası: "+err.Error(), nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Create(&timing, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Create Timing

// #region Update Timing
func (ctrl *TimingController) Update(c *gin.Context) {
	var timing datamodels.Timing
	if err := c.ShouldBindJSON(&timing); err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Veri doğrulama hatası: "+err.Error(), nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Update(&timing, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Update Timing

// #region Delete Timing
func (ctrl *TimingController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz ID formatı.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.Delete(id, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Delete Timing

// #region Get Timing By Id
func (ctrl *TimingController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz ID formatı.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.GetById(id, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Get Timing By Id

// #region Get All Timings
func (ctrl *TimingController) GetAll(c *gin.Context) {
	// İstekten QueryModel'i oluştur
	var query mvc.QueryModel
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Veri doğrulama hatası: "+err.Error(), nil))
		return
	}

	// Servis çağrısı
	context := models.NewContext(c)
	result := ctrl.service.GetAll(&query, context)

	// Sonuç döndür
	c.JSON(http.StatusOK, result)
}

//#endregion Get All Timings

// #region Get Timings By ClientProjectId
func (ctrl *TimingController) GetByClientProjectId(c *gin.Context) {
	clientProjectId, err := strconv.Atoi(c.Param("clientProjectId"))
	if err != nil || clientProjectId <= 0 {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz ClientProject ID formatı.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.GetByClientProjectId(clientProjectId, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Get Timings By ClientProjectId

// #region Get Timings By Date Range
func (ctrl *TimingController) GetByDateRange(c *gin.Context) {
	startDateStr := c.Query("startDate")
	endDateStr := c.Query("endDate")

	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz başlangıç tarihi formatı.", nil))
		return
	}

	endDate, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, lgo.NewLogicError("Geçersiz bitiş tarihi formatı.", nil))
		return
	}

	context := models.NewContext(c)
	result := ctrl.service.GetByDateRange(startDate, endDate, context)
	c.JSON(http.StatusOK, result)
}

//#endregion Get Timings By Date Range
