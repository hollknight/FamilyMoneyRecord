package user

// User 用户表
type User struct {
	ID                 uint64 `gorm:"primary_key;auto_increment"`
	Username           string `gorm:"type:varchar(100);not null;unique"`
	Password           string `gorm:"type:varchar(150);not null"`
	Name               string `gorm:"type:varchar(20);default:''"`
	ReceiptSum         int    `gorm:"type:int;default:0"`
	DisbursementSum    int    `gorm:"type:int;default:0"`
	AdvanceConsumption int    `gorm:"type:int;default:0"`
}
