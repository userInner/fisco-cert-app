package models

import "gorm.io/gorm"

/**

 */

type User struct {
	ID           uint64 `gorm:"autoIncrement"`          // 用户主键
	UserID       uint64 `gorm:"not null; uniqueIndex"`  // 用户唯一标志
	UserName     string `gorm:"not null; uniqueIndex"`  // 用户名
	Password     string `gorm:"not null;"`              // 用户密码
	Email        string `gorm:"default:null"`           // 用户邮箱
	Gender       uint8  `gorm:"default:0; not null"`    // 用户性别
	Address      string `gorm:"not null"`               // 用户链上地址（用于盲水印嵌入）
	Status       uint8  `gorm:"default:0; not null"`    // 用户状态（0-活跃，1-禁用，2-未验证）
	AccessToken  string `gorm:"-" json:"access_token"`  // 短期token
	RefreshToken string `gorm:"-" json:"refresh_token"` // 持久化token
	CertSum      uint64 `gorm:"-" json:"cert_sum"`      // 存证数量
	*gorm.Model
}

func (*User) TableName() string {
	return "t_user"
}
