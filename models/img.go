package models

import "gorm.io/gorm"

/**
图片表(t_img)
用于存储经过盲水印处理之后的图片存储
*/

type Img struct {
	ID       uint64 `gorm:"autoIncrement"`         // 图片主键
	ImgID    uint64 `gorm:"not null; uniqueIndex"` // 图片ID
	UserID   uint64 `gorm:"not null;"`             // 上传的用户ID
	FileName string `gorm:"not null;"`             // 存在腾讯云万象桶内的名字
	FileSize uint64 `gorm:"not null;"`             // 图片大小
	Suffix   string `gorm:"not null;"`             // 图片后缀
	gorm.Model
}

func (*Img) TableName() string {
	return "t_img"
}
