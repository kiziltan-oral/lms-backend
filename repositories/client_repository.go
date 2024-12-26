package repositories

import (
	"errors"

	datamodels "lms-web-services-main/models/data"

	"github.com/LGYtech/lgo"
	"gorm.io/gorm"
)

type ClientRepository interface {
	Create(client *datamodels.Client) *lgo.OperationResult
	Update(client *datamodels.Client) *lgo.OperationResult
	Delete(id int) *lgo.OperationResult
	GetById(id int) *lgo.OperationResult
	GetAll() *lgo.OperationResult
}

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{db: db}
}

// #region Create Client
func (r *clientRepository) Create(client *datamodels.Client) *lgo.OperationResult {
	if err := client.Validate(); err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}

	result := r.db.Create(&client)
	if result.Error != nil {
		return lgo.NewLogicError(result.Error.Error(), nil)
	}
	return lgo.NewSuccess(client)
}

// #endregion Create Client

// #region Update Client
func (r *clientRepository) Update(client *datamodels.Client) *lgo.OperationResult {
	if err := client.ValidateForUpdate(); err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}

	existingClient := &datamodels.Client{}
	if err := r.db.First(&existingClient, client.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lgo.NewLogicError("Müşteri bulunamadı.", nil)
		}
		return lgo.NewLogicError(err.Error(), nil)
	}

	existingClient.ShortTitle = client.ShortTitle
	existingClient.Title = client.Title
	existingClient.Notes = client.Notes
	existingClient.IsActive = client.IsActive

	if err := r.db.Save(&existingClient).Error; err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}
	return lgo.NewSuccess(existingClient)
}

// #endregion Update Client

// #region Delete Client
func (r *clientRepository) Delete(id int) *lgo.OperationResult {
	client := &datamodels.Client{}
	if err := r.db.First(&client, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lgo.NewLogicError("Müşteri bulunamadı.", nil)
		}
		return lgo.NewLogicError(err.Error(), nil)
	}

	if err := r.db.Delete(&client).Error; err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}
	return lgo.NewSuccess(nil)
}

// #endregion Delete Client

// #region Get Client By Id
func (r *clientRepository) GetById(id int) *lgo.OperationResult {
	client := &datamodels.Client{}
	if err := r.db.First(&client, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lgo.NewLogicError("Müşteri bulunamadı.", nil)
		}
		return lgo.NewLogicError(err.Error(), nil)
	}
	return lgo.NewSuccess(client)
}

// #endregion Get Client By Id

// #region Get All Clients
func (r *clientRepository) GetAll() *lgo.OperationResult {
	var clients []*datamodels.Client
	result := r.db.Find(&clients)
	if result.Error != nil {
		return lgo.NewLogicError(result.Error.Error(), nil)
	}
	return lgo.NewSuccess(clients)
}

// #endregion Get All Clients
