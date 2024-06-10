package models

import "gorm.io/gorm" // simplified sql

type UserBasic struct {
	gorm.Model
	Name       string
	Password   string
	Phone      string
	Email      string
	Identity   string
	ClientIp   string
	ClientPort string
	LoginTime  uint64
	LogoutTime uint64
	IsLogout   bool
	DeviceInfo string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
