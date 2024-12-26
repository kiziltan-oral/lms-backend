package repositories

import (
	"errors"
	"time"

	datamodels "lms-web-services-main/models/data"

	"github.com/LGYtech/lgo"
	"gorm.io/gorm"
)

type TimingRepository interface {
	Create(timing *datamodels.Timing) *lgo.OperationResult
	Update(timing *datamodels.Timing) *lgo.OperationResult
	Delete(id int) *lgo.OperationResult
	GetById(id int) *lgo.OperationResult
	GetAll() *lgo.OperationResult
	GetByClientProjectId(clientProjectId int) *lgo.OperationResult
	GetByDateRange(startDate time.Time, endDate time.Time) *lgo.OperationResult
}

type timingRepository struct {
	db *gorm.DB
}

func NewTimingRepository(db *gorm.DB) TimingRepository {
	return &timingRepository{db: db}
}

// #region Create Timing
func (r *timingRepository) Create(timing *datamodels.Timing) *lgo.OperationResult {
	if err := timing.Validate(); err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}

	result := r.db.Create(&timing)
	if result.Error != nil {
		return lgo.NewLogicError(result.Error.Error(), nil)
	}
	return lgo.NewSuccess(timing)
}

// #endregion Create Timing

// #region Update Timing
func (r *timingRepository) Update(timing *datamodels.Timing) *lgo.OperationResult {
	if err := timing.ValidateForUpdate(); err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}

	existingTiming := &datamodels.Timing{}
	if err := r.db.First(&existingTiming, timing.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lgo.NewLogicError("Timing not found.", nil)
		}
		return lgo.NewLogicError(err.Error(), nil)
	}

	existingTiming.Title = timing.Title
	existingTiming.Description = timing.Description
	existingTiming.StartDateTime = timing.StartDateTime
	existingTiming.EndDateTime = timing.EndDateTime
	existingTiming.Status = timing.Status

	if err := r.db.Save(&existingTiming).Error; err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}
	return lgo.NewSuccess(existingTiming)
}

// #endregion Update Timing

// #region Delete Timing
func (r *timingRepository) Delete(id int) *lgo.OperationResult {
	timing := &datamodels.Timing{}
	if err := r.db.First(&timing, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lgo.NewLogicError("Timing not found.", nil)
		}
		return lgo.NewLogicError(err.Error(), nil)
	}

	if err := r.db.Delete(&timing).Error; err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}
	return lgo.NewSuccess(nil)
}

// #endregion Delete Timing

// #region Get Timing By Id
func (r *timingRepository) GetById(id int) *lgo.OperationResult {
	timing := &datamodels.Timing{}
	if err := r.db.First(&timing, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lgo.NewLogicError("Timing not found.", nil)
		}
		return lgo.NewLogicError(err.Error(), nil)
	}
	return lgo.NewSuccess(timing)
}

// #endregion Get Timing By Id

// #region Get All Timings
func (r *timingRepository) GetAll() *lgo.OperationResult {
	var timings []*datamodels.Timing
	result := r.db.Find(&timings)
	if result.Error != nil {
		return lgo.NewLogicError(result.Error.Error(), nil)
	}
	return lgo.NewSuccess(timings)
}

// #endregion Get All Timings

// #region Get Timings By ClientProjectId
func (r *timingRepository) GetByClientProjectId(clientProjectId int) *lgo.OperationResult {
	var timings []*datamodels.Timing
	result := r.db.Where("client_project_id = ?", clientProjectId).Find(&timings)
	if result.Error != nil {
		return lgo.NewLogicError(result.Error.Error(), nil)
	}
	return lgo.NewSuccess(timings)
}

// #endregion Get Timings By ClientProjectId

// #region Get Timings By Date Range
func (r *timingRepository) GetByDateRange(startDate time.Time, endDate time.Time) *lgo.OperationResult {
	var timings []*datamodels.Timing
	result := r.db.Where("start_date_time >= ? AND end_date_time <= ?", startDate, endDate).Find(&timings)
	if result.Error != nil {
		return lgo.NewLogicError(result.Error.Error(), nil)
	}
	return lgo.NewSuccess(timings)
}

// #endregion Get Timings By Date Range
