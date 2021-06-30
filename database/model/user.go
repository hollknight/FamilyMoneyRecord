package model

// User 用户表
type User struct {
	ID                 uint64  `gorm:"primary_key;auto_increment"`
	Username           string  `gorm:"type:varchar(100);not null;unique"`
	Password           string  `gorm:"type:char(32);not null"`
	Name               string  `gorm:"type:varchar(20);default:''"`
	ReceiptSum         float64 `gorm:"type:double;default:0"`
	DisbursementSum    float64 `gorm:"type:double;default:0"`
	AdvanceConsumption float64 `gorm:"type:double;default:0"`
	Bills              []Bill
	Accounts           []Account
}
