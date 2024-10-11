package mysql

import (
	"fisco-cert-app/models"
	"time"
)

type TxSum struct {
	Date       string `json:"date"`
	TransCount int64  `json:"trans_count" gorm:"column:sum"`
}

func InsertTxMonitor(monitor *models.TransactionMonitor) (err error) {
	result := db.Model(&models.TransactionMonitor{}).Create(monitor)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateTxMonitor(monitor *models.TransactionMonitor) (err error) {
	result := db.Model(&models.TransactionMonitor{}).
		Where("user_id", monitor.UserID).
		Update("trans_count", monitor.TransCount).
		Update("trans_hash_lastest", monitor.TransHashLastest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 获取用户交易数量
func GetUserTxMonitor(userID uint64) (monitor *models.TransactionMonitor, err error) {
	err = db.Model(&models.TransactionMonitor{}).
		Select("trans_count, trans_hash_lastest").
		Where("user_id", userID).
		First(&monitor).Error
	return
}

// 获取近十日的交易总量
func GetNowTenTxCount() (sums []*TxSum, err error) {
	now := time.Now()
	startDate := now.AddDate(0, 0, -9)

	err = db.Model(&models.TransactionMonitor{}).
		Select("DATE(updated_at) as date, sum(trans_count) as sum").
		Where("updated_at >= ?", startDate).
		Group("date").
		Find(&sums).
		Error
	return
}
