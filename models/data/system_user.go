package data

import (
	"errors"
	"net/mail"

	"github.com/google/uuid"
)

type SystemUser struct {
	Id           uuid.UUID `gorm:"column:Id;type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name         string    `gorm:"column:Name;type:varchar(50);not null" json:"n"`
	Surname      string    `gorm:"column:Surname;type:varchar(50);not null" json:"sn"`
	Email        string    `gorm:"column:Email;type:varchar(255);not null" json:"e"`
	Password     string    `gorm:"column:Password;type:varchar(64);not null" json:"p"`
	PasswordSalt string    `gorm:"column:PasswordSalt;type:varchar(15);not null" json:"ps"`
	IsActive     bool      `gorm:"column:IsActive;type:boolean;not null;default:true" json:"ia"`
}

func (SystemUser) TableName() string {
	return "SystemUsers"
}

func (model *SystemUser) Validate() error {
	if model.Name == "" {
		return errors.New("ad alanı zorunludur")
	}
	if len(model.Name) > 50 {
		return errors.New("ad 50 karakterden uzun olamaz")
	}
	if model.Surname == "" {
		return errors.New("soyad alanı zorunludur")
	}
	if len(model.Surname) > 50 {
		return errors.New("soyad 50 karakterden uzun olamaz")
	}
	if model.Email == "" {
		return errors.New("e-posta alanı zorunludur")
	}
	if len(model.Email) > 255 {
		return errors.New("e-posta 255 karakterden uzun olamaz")
	}
	if _, err := mail.ParseAddress(model.Email); err != nil {
		return errors.New("geçerli bir e-posta adresi giriniz")
	}
	if model.Password == "" {
		return errors.New("şifre alanı zorunludur")
	}
	if len(model.Password) > 64 {
		return errors.New("şifre 64 karakterden uzun olamaz")
	}
	if model.PasswordSalt == "" {
		return errors.New("şifre tuzu (salt) alanı zorunludur")
	}
	if len(model.PasswordSalt) > 15 {
		return errors.New("şifre tuzu (salt) 15 karakterden uzun olamaz")
	}
	return nil
}

func (model *SystemUser) ValidateForUpdate() error {
	if model.Name == "" {
		return errors.New("ad alanı zorunludur")
	}
	if len(model.Name) > 50 {
		return errors.New("ad 50 karakterden uzun olamaz")
	}
	if model.Surname == "" {
		return errors.New("soyad alanı zorunludur")
	}
	if len(model.Surname) > 50 {
		return errors.New("soyad 50 karakterden uzun olamaz")
	}
	if model.Email == "" {
		return errors.New("e-posta alanı zorunludur")
	}
	if len(model.Email) > 255 {
		return errors.New("e-posta 255 karakterden uzun olamaz")
	}
	if _, err := mail.ParseAddress(model.Email); err != nil {
		return errors.New("geçerli bir e-posta adresi giriniz")
	}
	return nil
}
