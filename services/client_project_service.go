package services

import (
	"lms-web-services-main/models"
	datamodels "lms-web-services-main/models/data"
	repositories "lms-web-services-main/repositories"

	"github.com/LGYtech/lgo"
)

type ClientProjectService interface {
	Create(clientProject *datamodels.ClientProject, c *models.Context) *lgo.OperationResult
	Update(clientProject *datamodels.ClientProject, c *models.Context) *lgo.OperationResult
	Delete(id int, c *models.Context) *lgo.OperationResult
	GetById(id int, c *models.Context) *lgo.OperationResult
	GetAll(c *models.Context) *lgo.OperationResult
	GetByClientId(clientId int, c *models.Context) *lgo.OperationResult
}

type clientProjectService struct {
	repo        repositories.ClientProjectRepository
	saveRules   ClientProjectRuleHandler
	updateRules ClientProjectRuleHandler
	deleteRules ClientProjectRuleHandler
	readRules   ClientProjectRuleHandler
}

func NewClientProjectService(repo repositories.ClientProjectRepository) ClientProjectService {
	return &clientProjectService{
		repo: repo,
		saveRules: (&ClientProjectRuleHandlerValidation{}).
			SetNext(&ClientProjectRuleHandlerCheckAlterAuthorization{}),
		updateRules: (&ClientProjectRuleHandlerUpdateValidation{}).
			SetNext(&ClientProjectRuleHandlerCheckAlterAuthorization{}),
		deleteRules: (&ClientProjectRuleHandlerCheckDeleteAuthorization{}),
		readRules:   &ClientProjectRuleHandlerCheckReadAuthorization{},
	}
}

// #region Create ClientProject
func (s *clientProjectService) Create(clientProject *datamodels.ClientProject, c *models.Context) *lgo.OperationResult {
	if result := s.saveRules.Handle(clientProject, c); !result.IsSuccess() {
		return result
	}
	return s.repo.Create(clientProject)
}

//#endregion Create ClientProject

// #region Update ClientProject
func (s *clientProjectService) Update(clientProject *datamodels.ClientProject, c *models.Context) *lgo.OperationResult {
	if result := s.updateRules.Handle(clientProject, c); !result.IsSuccess() {
		return result
	}
	return s.repo.Update(clientProject)
}

//#endregion Update ClientProject

// #region Delete ClientProject
func (s *clientProjectService) Delete(id int, c *models.Context) *lgo.OperationResult {
	if id <= 0 {
		return lgo.NewLogicError("Geçersiz ID.", nil)
	}

	clientProject := &datamodels.ClientProject{Id: id}
	if result := s.deleteRules.Handle(clientProject, c); !result.IsSuccess() {
		return result
	}
	return s.repo.Delete(id)
}

//#endregion Delete ClientProject

// #region Get ClientProject By Id
func (s *clientProjectService) GetById(id int, c *models.Context) *lgo.OperationResult {
	if id <= 0 {
		return lgo.NewLogicError("Geçersiz ID.", nil)
	}

	clientProject := &datamodels.ClientProject{Id: id}
	if result := s.readRules.Handle(clientProject, c); !result.IsSuccess() {
		return result
	}
	return s.repo.GetById(id)
}

//#endregion Get ClientProject By Id

// #region Get All ClientProjects
func (s *clientProjectService) GetAll(c *models.Context) *lgo.OperationResult {
	clientProject := &datamodels.ClientProject{}
	if result := s.readRules.Handle(clientProject, c); !result.IsSuccess() {
		return result
	}
	return s.repo.GetAll()
}

//#endregion Get All ClientProjects

// #region Get ClientProjects By ClientId
func (s *clientProjectService) GetByClientId(clientId int, c *models.Context) *lgo.OperationResult {
	if clientId <= 0 {
		return lgo.NewLogicError("Geçersiz müşteri ID.", nil)
	}

	clientProject := &datamodels.ClientProject{ClientId: clientId}
	if result := s.readRules.Handle(clientProject, c); !result.IsSuccess() {
		return result
	}
	return s.repo.GetByClientId(clientId)
}

//#endregion Get ClientProjects By ClientId
