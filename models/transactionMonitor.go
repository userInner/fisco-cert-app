package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type TransactionMonitor struct {
	ID               uint64         `gorm:"autoIncrement"`
	UserID           uint64         `gorm:"not null; uniqueIndex"`                                                                  // 用户唯一标志
	TransCount       int            `gorm:"column:trans_count;type:int(11);comment:交易量;NOT NULL" json:"trans_count"`                // 用户交易总数
	TransHashLastest sql.NullString `gorm:"column:trans_hash_lastest;type:varchar(128);comment:最新交易hash" json:"trans_hash_lastest"` // 用户最新交易hash
	*gorm.Model
}

func (m *TransactionMonitor) TableName() string {
	return "t_transaction_monitor"
}
