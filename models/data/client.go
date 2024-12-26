package data

import "errors"

type Client struct {
	Id         int    `gorm:"column:Id;type:serial;primary_key" json:"id"`
	ShortTitle string `gorm:"column:ShortTitle;type:varchar(50);not null" json:"st"`
	Title      string `gorm:"column:Title;type:varchar(200);not null" json:"t"`
	Notes      string `gorm:"column:Notes;type:text" json:"nt"`
	IsActive   bool   `gorm:"column:IsActive;type:boolean;not null;default:true" json:"ia"`
}

func (Client) TableName() string {
	return "Clients"
}

func (model *Client) Validate() error {
	if model.ShortTitle == "" {
		return errors.New("kısa başlık alanı zorunludur")
	}
	if len(model.ShortTitle) > 50 {
		return errors.New("kısa başlık 50 karakterden uzun olamaz")
	}
	if model.Title == "" {
		return errors.New("başlık alanı zorunludur")
	}
	if len(model.Title) > 200 {
		return errors.New("başlık 200 karakterden uzun olamaz")
	}
	return nil
}

func (model *Client) ValidateForUpdate() error {
	if model.ShortTitle == "" {
		return errors.New("kısa başlık alanı zorunludur")
	}
	if len(model.ShortTitle) > 50 {
		return errors.New("kısa başlık 50 karakterden uzun olamaz")
	}
	if model.Title == "" {
		return errors.New("başlık alanı zorunludur")
	}
	if len(model.Title) > 200 {
		return errors.New("başlık 200 karakterden uzun olamaz")
	}
	return nil
}
