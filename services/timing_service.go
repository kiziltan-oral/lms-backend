package services

import (
	"lms-web-services-main/models"
	datamodels "lms-web-services-main/models/data"
	"lms-web-services-main/models/mvc"
	repositories "lms-web-services-main/repositories"

	"time"

	"github.com/LGYtech/lgo"
)

type TimingService interface {
	Create(timing *datamodels.Timing, c *models.Context) *lgo.OperationResult
	Update(timing *datamodels.Timing, c *models.Context) *lgo.OperationResult
	Delete(id int, c *models.Context) *lgo.OperationResult
	GetById(id int, c *models.Context) *lgo.OperationResult
	GetAll(query *mvc.QueryModel, c *models.Context) *lgo.OperationResult
	GetByClientProjectId(clientProjectId int, c *models.Context) *lgo.OperationResult
	GetByDateRange(startDate time.Time, endDate time.Time, c *models.Context) *lgo.OperationResult
}

type timingService struct {
	repo        repositories.TimingRepository
	saveRules   TimingRuleHandler
	updateRules TimingRuleHandler
	deleteRules TimingRuleHandler
	readRules   TimingRuleHandler
}

func NewTimingService(repo repositories.TimingRepository) TimingService {
	return &timingService{
		repo: repo,
		saveRules: (&TimingRuleHandlerValidation{}).
			SetNext(&TimingRuleHandlerCheckAlterAuthorization{}),
		updateRules: (&TimingRuleHandlerUpdateValidation{}).
			SetNext(&TimingRuleHandlerCheckAlterAuthorization{}),
		deleteRules: (&TimingRuleHandlerCheckDeleteAuthorization{}),
		readRules:   &TimingRuleHandlerCheckReadAuthorization{},
	}
}

// #region Create Timing
func (s *timingService) Create(timing *datamodels.Timing, c *models.Context) *lgo.OperationResult {
	if result := s.saveRules.Handle(timing, c); !result.IsSuccess() {
		return result
	}
	return s.repo.Create(timing)
}

//#endregion Create Timing

// #region Update Timing
func (s *timingService) Update(timing *datamodels.Timing, c *models.Context) *lgo.OperationResult {
	if result := s.updateRules.Handle(timing, c); !result.IsSuccess() {
		return result
	}
	return s.repo.Update(timing)
}

//#endregion Update Timing

// #region Delete Timing
func (s *timingService) Delete(id int, c *models.Context) *lgo.OperationResult {
	if id <= 0 {
		return lgo.NewLogicError("Geçersiz ID.", nil)
	}

	timing := &datamodels.Timing{Id: id}
	if result := s.deleteRules.Handle(timing, c); !result.IsSuccess() {
		return result
	}

	return s.repo.Delete(id)
}

//#endregion Delete Timing

// #region Get Timing By Id
func (s *timingService) GetById(id int, c *models.Context) *lgo.OperationResult {
	if id <= 0 {
		return lgo.NewLogicError("Geçersiz ID.", nil)
	}

	timing := &datamodels.Timing{Id: id}
	if result := s.readRules.Handle(timing, c); !result.IsSuccess() {
		return result
	}

	return s.repo.GetById(id)
}

//#endregion Get Timing By Id

// #region Get All Timings
func (s *timingService) GetAll(query *mvc.QueryModel, c *models.Context) *lgo.OperationResult {
	timing := &datamodels.Timing{}
	if result := s.readRules.Handle(timing, c); !result.IsSuccess() {
		return result
	}

	if result := query.Validate(); !result.IsSuccess() {
		return lgo.NewLogicError("Geçersiz sorgu parametreleri: "+result.ErrorMessage, nil)
	}

	return s.repo.GetAll(query)
}

//#endregion Get All Timings

// #region Get Timings By ClientProjectId
func (s *timingService) GetByClientProjectId(clientProjectId int, c *models.Context) *lgo.OperationResult {
	if clientProjectId <= 0 {
		return lgo.NewLogicError("Geçersiz ClientProject ID.", nil)
	}

	timing := &datamodels.Timing{ClientProjectId: clientProjectId}
	if result := s.readRules.Handle(timing, c); !result.IsSuccess() {
		return result
	}

	return s.repo.GetByClientProjectId(clientProjectId)
}

//#endregion Get Timings By ClientProjectId

// #region Get Timings By Date Range
func (s *timingService) GetByDateRange(startDate time.Time, endDate time.Time, c *models.Context) *lgo.OperationResult {
	if startDate.IsZero() || endDate.IsZero() {
		return lgo.NewLogicError("Başlangıç ve bitiş tarihleri zorunludur.", nil)
	}
	if endDate.Before(startDate) {
		return lgo.NewLogicError("Bitiş tarihi, başlangıç tarihinden önce olamaz.", nil)
	}

	// Handle read rules (optional)
	timing := &datamodels.Timing{
		StartDateTime: startDate,
		EndDateTime:   endDate,
	}
	if result := s.readRules.Handle(timing, c); !result.IsSuccess() {
		return result
	}

	return s.repo.GetByDateRange(startDate, endDate)
}

//#endregion Get Timings By Date Range
