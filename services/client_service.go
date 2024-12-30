package services

import (
	"lms-web-services-main/models"
	datamodels "lms-web-services-main/models/data"
	"lms-web-services-main/models/mvc"
	repositories "lms-web-services-main/repositories"

	"github.com/LGYtech/lgo"
)

// #region Client Service Interface
type ClientService interface {
	Create(client *datamodels.Client, c *models.Context) *lgo.OperationResult
	Update(client *datamodels.Client, c *models.Context) *lgo.OperationResult
	Delete(id int, c *models.Context) *lgo.OperationResult
	GetById(id int, c *models.Context) *lgo.OperationResult
	GetAll(query *mvc.QueryModel, c *models.Context) *lgo.OperationResult
}

//#endregion Client Service Interface

// #region Client Service Implementation
type clientService struct {
	repo        repositories.ClientRepository
	saveRules   ClientRuleHandler
	updateRules ClientRuleHandler
	deleteRules ClientRuleHandler
	readRules   ClientRuleHandler
}

func NewClientService(repo repositories.ClientRepository) ClientService {
	return &clientService{
		repo: repo,
		saveRules: (&ClientRuleHandlerValidation{}).
			SetNext(&ClientRuleHandlerCheckAlterAuthorization{}),
		updateRules: (&ClientRuleHandlerUpdateValidation{}).
			SetNext(&ClientRuleHandlerCheckAlterAuthorization{}),
		deleteRules: (&ClientRuleHandlerCheckDeleteAuthorization{}),
		readRules:   &ClientRuleHandlerCheckReadAuthorization{},
	}
}

//#endregion Client Service Implementation

// #region Create Client
func (s *clientService) Create(client *datamodels.Client, c *models.Context) *lgo.OperationResult {
	if result := s.saveRules.Handle(client, c); !result.IsSuccess() {
		return result
	}
	return s.repo.Create(client)
}

//#endregion Create Client

// #region Update Client
func (s *clientService) Update(client *datamodels.Client, c *models.Context) *lgo.OperationResult {
	if result := s.updateRules.Handle(client, c); !result.IsSuccess() {
		return result
	}
	return s.repo.Update(client)
}

//#endregion Update Client

// #region Delete Client
func (s *clientService) Delete(id int, c *models.Context) *lgo.OperationResult {
	if id <= 0 {
		return lgo.NewLogicError("Geçersiz ID.", nil)
	}

	client := &datamodels.Client{Id: id}
	if result := s.deleteRules.Handle(client, c); !result.IsSuccess() {
		return result
	}

	return s.repo.Delete(id)
}

//#endregion Delete Client

// #region Get Client By Id
func (s *clientService) GetById(id int, c *models.Context) *lgo.OperationResult {
	if id <= 0 {
		return lgo.NewLogicError("Geçersiz ID.", nil)
	}

	client := &datamodels.Client{Id: id}
	if result := s.readRules.Handle(client, c); !result.IsSuccess() {
		return result
	}

	return s.repo.GetById(id)
}

//#endregion Get Client By Id

// #region Get All Clients
func (s *clientService) GetAll(query *mvc.QueryModel, c *models.Context) *lgo.OperationResult {
	client := &datamodels.Client{}
	if result := s.readRules.Handle(client, c); !result.IsSuccess() {
		return result
	}

	if result := query.Validate(); !result.IsSuccess() {
		return lgo.NewLogicError("Geçersiz sorgu parametreleri: "+result.ErrorMessage, nil)
	}

	return s.repo.GetAll(query)
}

//#endregion Get All Clients
