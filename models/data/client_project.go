package data

import "errors"

type ClientProject struct {
	Id       int    `gorm:"column:Id;type:serial;primary_key" json:"id"`
	ClientId int    `gorm:"column:ClientId;type:integer;not null" json:"cid"`
	Name     string `gorm:"column:Name;type:varchar(100);not null" json:"n"`
	IsActive bool   `gorm:"column:IsActive;type:boolean;not null;default:true" json:"ia"`
}

func (ClientProject) TableName() string {
	return "ClientProjects"
}

func (model *ClientProject) Validate() error {
	if model.ClientId <= 0 {
		return errors.New("client_id alanı zorunludur")
	}
	if model.Name == "" {
		return errors.New("name alanı zorunludur")
	}
	if len(model.Name) > 100 {
		return errors.New("name alanı 100 karakterden uzun olamaz")
	}
	return nil
}

func (model *ClientProject) ValidateForUpdate() error {
	if model.Name == "" {
		return errors.New("name alanı zorunludur")
	}
	if len(model.Name) > 100 {
		return errors.New("name alanı 100 karakterden uzun olamaz")
	}
	return nil
}
