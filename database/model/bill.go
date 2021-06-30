package model

import "time"

// Bill 账单表
type Bill struct {
	ID           uint64 `gorm:"primary_key;auto_increment"`
	UserID       uint64
	Receipt      float64   `gorm:"type:double;default:0"`
	Disbursement float64   `gorm:"type:double;default:0"`
	Type         string    `gorm:"type:varchar(25);default:''"`
	Note         string    `gorm:"type:varchar(200);default:''"`
	Time         time.Time `gorm:"type:datetime;default:'1000-01-01 00:00:00'"`
}
