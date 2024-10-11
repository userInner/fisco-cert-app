package models

import "gorm.io/gorm"

/**
版权表(t_cert)
*/

type Cert struct {
	ID       uint64 `gorm:"autoIncrement"`         // 版权主键
	TxID     string `gorm:"not null; uniqueIndex"` // 交易ID
	UserID   uint64 `gorm:"not null;"`             // 所属人
	ImageID  uint64 `gorm:"not null;"`             // 文件ID
	Evidence string `gorm:"not null;"`             // 版权合约地址
	gorm.Model
}

func (*Cert) TableName() string {
	return "t_cert"
}
