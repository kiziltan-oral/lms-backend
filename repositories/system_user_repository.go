package repositories

import (
	"errors"

	datamodels "lms-web-services-main/models/data"

	"github.com/LGYtech/lgo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SystemUserRepository interface {
	Create(systemUser *datamodels.SystemUser) *lgo.OperationResult
	Update(systemUser *datamodels.SystemUser) *lgo.OperationResult
	Delete(id uuid.UUID) *lgo.OperationResult
	GetById(id uuid.UUID) *lgo.OperationResult
	GetByEmail(email string) *lgo.OperationResult
	GetAll() *lgo.OperationResult
	CheckForeignReferences(systemUser *datamodels.SystemUser) *lgo.OperationResult
	CheckExistingSystemUser(systemUser *datamodels.SystemUser) *lgo.OperationResult
}

type systemUserRepository struct {
	db *gorm.DB
}

func NewSystemUserRepository(db *gorm.DB) SystemUserRepository {
	return &systemUserRepository{db: db}
}

// #region Create SystemUser
func (r *systemUserRepository) Create(systemUser *datamodels.SystemUser) *lgo.OperationResult {
	result := r.db.Create(&systemUser)
	if result.Error != nil {
		return lgo.NewLogicError(result.Error.Error(), nil)
	}
	return lgo.NewSuccess(systemUser)
}

// #endregion Create SystemUser

// #region Update SystemUser
func (r *systemUserRepository) Update(systemUser *datamodels.SystemUser) *lgo.OperationResult {
	existingUser := &datamodels.SystemUser{}
	if err := r.db.First(&existingUser, systemUser.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lgo.NewLogicError("Kullanıcı bulunamadı.", nil)
		}
		return lgo.NewLogicError(err.Error(), nil)
	}

	existingUser.Name = systemUser.Name
	existingUser.Surname = systemUser.Surname
	existingUser.Email = systemUser.Email
	existingUser.Password = systemUser.Password
	existingUser.PasswordSalt = systemUser.PasswordSalt
	existingUser.IsActive = systemUser.IsActive

	if err := r.db.Save(&existingUser).Error; err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}

	return lgo.NewSuccess(existingUser)
}

// #endregion Update SystemUser

// #region Delete SystemUser
func (r *systemUserRepository) Delete(id uuid.UUID) *lgo.OperationResult {
	existingUser := &datamodels.SystemUser{}
	if err := r.db.First(&existingUser, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lgo.NewLogicError("Kullanıcı bulunamadı.", nil)
		}
		return lgo.NewLogicError(err.Error(), nil)
	}

	if err := r.db.Delete(&existingUser).Error; err != nil {
		return lgo.NewLogicError(err.Error(), nil)
	}

	return lgo.NewSuccess(nil)
}

// #endregion Delete SystemUser

// #region Get SystemUser By Id
func (r *systemUserRepository) GetById(id uuid.UUID) *lgo.OperationResult {
	systemUser := &datamodels.SystemUser{}
	if err := r.db.First(&systemUser, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lgo.NewLogicError("Kullanıcı bulunamadı.", nil)
		}
		return lgo.NewLogicError(err.Error(), nil)
	}
	return lgo.NewSuccess(systemUser)
}

// #endregion Get SystemUser By Id

// #region GetByEmail
func (r *systemUserRepository) GetByEmail(email string) *lgo.OperationResult {
	var systemUser *datamodels.SystemUser

	result := r.db.Where("\"Email\" = ?", email).First(&systemUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return lgo.NewSuccess(nil)
		}
		return lgo.NewLogicError(result.Error.Error(), nil)
	}

	return lgo.NewSuccess(systemUser)
}

// #endregion GetByEmail

// #region GetAll
func (r *systemUserRepository) GetAll() *lgo.OperationResult {
	var systemUsers []*datamodels.SystemUser
	result := r.db.Select("\"Id\",\"Name\", \"Surname\",\"Email\"").Find(&systemUsers)
	if result.Error != nil {
		return lgo.NewLogicError(result.Error.Error(), nil)
	}
	return lgo.NewSuccess(systemUsers)
}

// #endregion GetAll

// #region Check Foreign References
func (r *systemUserRepository) CheckForeignReferences(systemUser *datamodels.SystemUser) *lgo.OperationResult {
	var referenceCount int64

	// Check references in SystemUserSetting
	if err := r.db.Model(&datamodels.SystemUserSetting{}).Where("\"SystemUserId\"=?", systemUser.Id).Count(&referenceCount).Error; err != nil {
		return lgo.NewLogicError("Error checking references in SystemUserSetting: "+err.Error(), nil)
	}
	if referenceCount > 0 {
		return lgo.NewLogicError("SystemUser is referenced in SystemUserSetting.", nil)
	}

	// Check references in Timing
	if err := r.db.Model(&datamodels.Timing{}).Where("\"SystemUserId\"=?", systemUser.Id).Count(&referenceCount).Error; err != nil {
		return lgo.NewLogicError("Error checking references in Timing: "+err.Error(), nil)
	}
	if referenceCount > 0 {
		return lgo.NewLogicError("SystemUser is referenced in Timing.", nil)
	}

	return lgo.NewSuccess(nil)
}

// #endregion Check Foreign References

// #region Check Existing SystemUser
func (r *systemUserRepository) CheckExistingSystemUser(systemUser *datamodels.SystemUser) *lgo.OperationResult {
	existingUser := &datamodels.SystemUser{}
	if err := r.db.Where("\"Email\" = ?", systemUser.Email).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lgo.NewSuccess(nil)
		}
		return lgo.NewLogicError(err.Error(), nil)
	}
	return lgo.NewSuccess(existingUser)
}

// #endregion Check Existing SystemUser
