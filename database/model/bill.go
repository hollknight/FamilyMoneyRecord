package model

import "time"

// Bill 账单表
type Bill struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID       uint64    `json:"user_id"`
	Receipt      float64   `gorm:"type:double;default:0" json:"receipt"`
	Disbursement float64   `gorm:"type:double;default:0" json:"disbursement"`
	Type         string    `gorm:"type:varchar(25);default:''" json:"type"`
	Note         string    `gorm:"type:varchar(200);default:''" json:"note"`
	Time         time.Time `gorm:"type:datetime;default:'1000-01-01 00:00:00'" json:"time"`
}
