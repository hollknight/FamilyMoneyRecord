package model

// User 用户表
type User struct {
	ID                 uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Username           string    `gorm:"type:varchar(100);not null;unique" json:"username"`
	Password           string    `gorm:"type:char(32);not null" json:"password"`
	Name               string    `gorm:"type:varchar(20);default:''" json:"name"`
	ReceiptSum         float64   `gorm:"type:double;default:0" json:"receipt_sum"`
	DisbursementSum    float64   `gorm:"type:double;default:0" json:"disbursement_sum"`
	AdvanceConsumption float64   `gorm:"type:double;default:0" json:"advance_consumption"`
	Bills              []Bill    `json:"bills,omitempty"`
	Accounts           []Account `json:"accounts,omitempty"`
}
