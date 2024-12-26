package data

import (
	"errors"

	"github.com/google/uuid"
)

type SystemUserSetting struct {
	Id           int       `gorm:"column:Id;type:integer;primary_key" json:"id"`
	SystemUserId uuid.UUID `gorm:"column:SystemUserId;type:uuid;not null" json:"suid"`
	Key          string    `gorm:"column:Key;type:varchar(50);not null" json:"key"`
	Value        string    `gorm:"column:Value;type:varchar(200);not null" json:"val"`
	Description  string    `gorm:"column:Description;type:varchar(200);not null" json:"desc"`
}

func (SystemUserSetting) TableName() string {
	return "SystemUserSettings"
}

func (model *SystemUserSetting) Validate() error {
	if model.SystemUserId == uuid.Nil {
		return errors.New("system_user_id alanı zorunludur")
	}
	if len(model.Key) == 0 {
		return errors.New("key alanı zorunludur")
	}
	if len(model.Key) > 50 {
		return errors.New("key alanı 50 karakterden uzun olamaz")
	}
	if len(model.Value) == 0 {
		return errors.New("value alanı zorunludur")
	}
	if len(model.Value) > 200 {
		return errors.New("value alanı 200 karakterden uzun olamaz")
	}
	return nil
}

const (
	// System Users
	SYSTEM_USERS_VIEW   = "system.users.view"
	SYSTEM_USERS_ADD    = "system.users.add"
	SYSTEM_USERS_UPDATE = "system.users.update"
	SYSTEM_USERS_DELETE = "system.users.delete"

	// System Settings
	SYSTEM_SETTINGS_VIEW   = "system.settings.view"
	SYSTEM_SETTINGS_ADD    = "system.settings.add"
	SYSTEM_SETTINGS_UPDATE = "system.settings.update"
	SYSTEM_SETTINGS_DELETE = "system.settings.delete"

	// Clients
	CLIENTS_VIEW   = "clients.view"
	CLIENTS_ADD    = "clients.add"
	CLIENTS_UPDATE = "clients.update"
	CLIENTS_DELETE = "clients.delete"

	// Client Projects
	CLIENTPROJECTS_VIEW   = "clientprojects.view"
	CLIENTPROJECTS_ADD    = "clientprojects.add"
	CLIENTPROJECTS_UPDATE = "clientprojects.update"
	CLIENTPROJECTS_DELETE = "clientprojects.delete"

	// Timings
	TIMINGS_VIEW   = "timings.view"
	TIMINGS_ADD    = "timings.add"
	TIMINGS_UPDATE = "timings.update"
	TIMINGS_DELETE = "timings.delete"
)
