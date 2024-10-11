package models

import "gorm.io/gorm"

/**
验证表(t_auth)
*/

type Auth struct {
	ID      uint64 `gorm:"autoIncrement"` // 版权主键
	AuthID  uint64 `gorm:"not null;"`     // 交易ID
	UserID  uint64 `gorm:"not null;"`     // 所属人
	ImageID uint64 `gorm:"not null;"`     // 文件ID
	Address string `gorm:"not null;"`     // 水印内容
	gorm.Model
}

func (*Auth) TableName() string {
	return "t_auth"
}
