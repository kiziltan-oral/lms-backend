package services

import (
	"lms-web-services-main/models"
	datamodels "lms-web-services-main/models/data"
	"lms-web-services-main/repositories"

	"github.com/LGYtech/lgo"
	"github.com/google/uuid"
)

var (
	CacheService cacheServiceInterface = &cacheService{}
)

type cacheServiceInterface interface {
	AuthenticateSystemUser(token string) *lgo.OperationResult
	GetSystemUserCredential(token string) *lgo.OperationResult
	RegisterSystemUserCredential(token string, systemUser *datamodels.SystemUser) *lgo.OperationResult
	DeleteSystemUserCredential(token string) *lgo.OperationResult
	DeleteSystemUserCredentialById(id uuid.UUID) *lgo.OperationResult
	GetSystemUserSetting(c *models.Context, setting string) *lgo.OperationResult
	RemoveSystemUserSetting(systemUserId uuid.UUID, setting string) *lgo.OperationResult
}

type cacheService struct {
}

func (*cacheService) AuthenticateSystemUser(token string) *lgo.OperationResult {
	return repositories.CacheRepository.AuthenticateSystemUser(token)
}

func (*cacheService) GetSystemUserCredential(token string) *lgo.OperationResult {
	return repositories.CacheRepository.GetSystemUserCredential(token)
}

func (*cacheService) RegisterSystemUserCredential(token string, systemUser *datamodels.SystemUser) *lgo.OperationResult {
	return repositories.CacheRepository.RegisterSystemUserCredential(token, systemUser)
}

func (*cacheService) DeleteSystemUserCredential(token string) *lgo.OperationResult {
	return repositories.CacheRepository.DeleteSystemUserCredential(token)
}

func (*cacheService) GetSystemUserSetting(c *models.Context, setting string) *lgo.OperationResult {
	return repositories.CacheRepository.GetSystemUserSetting(c, setting)
}

func (*cacheService) RemoveSystemUserSetting(systemUserId uuid.UUID, setting string) *lgo.OperationResult {
	return repositories.CacheRepository.RemoveSystemUserSetting(systemUserId, setting)
}

func (*cacheService) DeleteSystemUserCredentialById(id uuid.UUID) *lgo.OperationResult {
	return repositories.CacheRepository.DeleteSystemUserCredentialById(id)
}
