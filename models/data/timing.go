package data

import (
	"errors"
	"time"

	"lms-web-services-main/models/enum"

	"github.com/google/uuid"
)

type Timing struct {
	Id              int             `gorm:"column:Id;type:serial;primary_key" json:"id"`
	ClientProjectId int             `gorm:"column:ClientProjectId;type:integer;not null" json:"cpid"`
	SystemUserId    uuid.UUID       `gorm:"column:SystemUserId;type:uuid;not null" json:"suid"`
	Title           string          `gorm:"column:Title;type:varchar(100);not null" json:"t"`
	Description     string          `gorm:"column:Description;type:text" json:"desc"`
	StartDateTime   time.Time       `gorm:"column:StartDateTime;type:timestamptz;not null" json:"sdt"`
	EndDateTime     time.Time       `gorm:"column:EndDateTime;type:timestamptz;not null" json:"edt"`
	Status          enum.StatusEnum `gorm:"column:Status;type:integer;not null" json:"st"`
}

func (Timing) TableName() string {
	return "Timings"
}

func (model *Timing) Validate() error {
	if model.ClientProjectId <= 0 {
		return errors.New("client_project_id alanı zorunludur")
	}
	if model.SystemUserId == uuid.Nil {
		return errors.New("system_user_id alanı zorunludur")
	}
	if model.Title == "" {
		return errors.New("başlık alanı zorunludur")
	}
	if len(model.Title) > 100 {
		return errors.New("başlık 100 karakterden uzun olamaz")
	}
	if model.StartDateTime.IsZero() {
		return errors.New("başlangıç zamanı alanı zorunludur")
	}
	if model.EndDateTime.IsZero() {
		return errors.New("bitiş zamanı alanı zorunludur")
	}
	if model.EndDateTime.Before(model.StartDateTime) {
		return errors.New("bitiş zamanı, başlangıç zamanından önce olamaz")
	}
	if !model.Status.IsValid() {
		return errors.New("geçersiz durum değeri")
	}
	return nil
}

func (model *Timing) ValidateForUpdate() error {
	if model.Title == "" {
		return errors.New("başlık alanı zorunludur")
	}
	if len(model.Title) > 100 {
		return errors.New("başlık 100 karakterden uzun olamaz")
	}
	if model.StartDateTime.IsZero() {
		return errors.New("başlangıç zamanı alanı zorunludur")
	}
	if model.EndDateTime.IsZero() {
		return errors.New("bitiş zamanı alanı zorunludur")
	}
	if model.EndDateTime.Before(model.StartDateTime) {
		return errors.New("bitiş zamanı, başlangıç zamanından önce olamaz")
	}
	if !model.Status.IsValid() {
		return errors.New("geçersiz durum değeri")
	}
	return nil
}
