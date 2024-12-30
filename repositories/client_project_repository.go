package repositories

import (
	"errors"

	datamodels "lms-web-services-main/models/data"
	"lms-web-services-main/models/mvc"

	"github.com/LGYtech/lgo"
	"gorm.io/gorm"
)

type ClientProjectRepository interface {
	Create(clientProject *datamodels.ClientProject) *lgo.OperationResult
	Update(clientProject *datamodels.ClientProject) *lgo.OperationResult
	Delete(id int) *lgo.OperationResult
	GetById(id int) *lgo.OperationResult
	GetAll(query *mvc.QueryModel) *lgo.OperationResult
	GetByClientId(clientId int) *lgo.OperationResult
}

type clientProjectRepository struct {
	db *gorm.DB
}

func NewClientProjectRepository(db *gorm.DB) ClientProjectRepository {
	return &clientProjectRepository{db: db}
}

// #region Create ClientProject
func (r *clientProjectRepository) Create(clientProject *datamodels.ClientProject) *lgo.OperationResult {
	if err := clientProject.Validate(); err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}

	result := r.db.Create(&clientProject)
	if result.Error != nil {
		return lgo.NewLogicError(result.Error.Error(), nil)
	}
	return lgo.NewSuccess(clientProject)
}

// #endregion Create ClientProject

// #region Update ClientProject
func (r *clientProjectRepository) Update(clientProject *datamodels.ClientProject) *lgo.OperationResult {
	if err := clientProject.ValidateForUpdate(); err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}

	existingProject := &datamodels.ClientProject{}
	if err := r.db.First(&existingProject, clientProject.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lgo.NewLogicError("Proje bulunamadı.", nil)
		}
		return lgo.NewLogicError(err.Error(), nil)
	}

	existingProject.Name = clientProject.Name
	existingProject.IsActive = clientProject.IsActive

	if err := r.db.Save(&existingProject).Error; err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}
	return lgo.NewSuccess(existingProject)
}

// #endregion Update ClientProject

// #region Delete ClientProject
func (r *clientProjectRepository) Delete(id int) *lgo.OperationResult {
	clientProject := &datamodels.ClientProject{}
	if err := r.db.First(&clientProject, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lgo.NewLogicError("Proje bulunamadı.", nil)
		}
		return lgo.NewLogicError(err.Error(), nil)
	}

	if err := r.db.Delete(&clientProject).Error; err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}
	return lgo.NewSuccess(nil)
}

// #endregion Delete ClientProject

// #region Get ClientProject By Id
func (r *clientProjectRepository) GetById(id int) *lgo.OperationResult {
	clientProject := &datamodels.ClientProject{}
	if err := r.db.First(&clientProject, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lgo.NewLogicError("Proje bulunamadı.", nil)
		}
		return lgo.NewLogicError(err.Error(), nil)
	}
	return lgo.NewSuccess(clientProject)
}

// #endregion Get ClientProject By Id

// #region Get All ClientProjects
func (r *clientProjectRepository) GetAll(query *mvc.QueryModel) *lgo.OperationResult {
	var clientProjects []*datamodels.ClientProject

	defaultSorting := &mvc.DataSortingOptionItem{
		ColumnName: "Name",
		Sorting:    0, // 0: ASC, 1: DESC
	}

	searchableColumns := []string{"Name"}

	// QueryModel'i uygula
	db, result := ApplyQueryModel(r.db, query, searchableColumns, defaultSorting)
	if !result.IsSuccess() {
		return lgo.NewLogicError("Sorgu modeli uygulanırken bir hata oluştu: "+result.ErrorMessage, nil)
	}

	queryResult := db.Find(&clientProjects)
	if queryResult.Error != nil {
		return lgo.NewLogicError("Veritabanı sorgusu başarısız: "+queryResult.Error.Error(), nil)
	}

	return lgo.NewSuccess(clientProjects)
}

// #endregion Get All ClientProjects

// #region Get ClientProjects By ClientId
func (r *clientProjectRepository) GetByClientId(clientId int) *lgo.OperationResult {
	var clientProjects []*datamodels.ClientProject
	result := r.db.Where("\"ClientId\" = ?", clientId).Find(&clientProjects)
	if result.Error != nil {
		return lgo.NewLogicError(result.Error.Error(), nil)
	}
	return lgo.NewSuccess(clientProjects)
}

// #endregion Get ClientProjects By ClientId
