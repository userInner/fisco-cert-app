package mysql

import (
	"database/sql"
	"fisco-cert-app/models"
	"fisco-cert-app/setting"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var client *sql.DB

// Init 初始化mysql
func Init(cfg *setting.MySQLConfig) (err error) {
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB), // DSN data source name
		DefaultStringSize:         256,                                                                                                                          // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                                                         // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                                                         // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                                                         // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                                                        // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		CreateBatchSize: 1000, //
		PrepareStmt:     true, // 开启全局预编译
	})

	if err != nil {
		return err
	}
	client, _ = db.DB()
	// SetMaxIdleConns: 设置空闲连接池中链接的最大数量
	client.SetMaxIdleConns(10)
	// SetMaxOpenConns: 设置打开数据库链接的最大数量
	client.SetMaxOpenConns(50)
	// SetConnMaxLifetime: 设置链接可复用的最大时间
	client.SetConnMaxLifetime(time.Hour)
	zap.L().Info("mysql connected success!")
	err = initMySQL()
	return
}

// gorm自动迁移初始化所有表
func initMySQL() (err error) {
	migrat := db.Migrator()

	if migrat.HasTable(&models.User{}) {
		user := &models.User{}
		// 提醒有些表已存在
		zap.L().Info("mysql already exist", zap.String("table", user.TableName()))
	}

	if migrat.HasTable(&models.Cert{}) {
		cert := &models.Cert{}
		zap.L().Info("mysql already exist", zap.String("table", cert.TableName()))
	}

	if migrat.HasTable(&models.Img{}) {
		img := &models.Img{}
		zap.L().Info("mysql already exist", zap.String("table", img.TableName()))
	}

	if migrat.HasTable(&models.Auth{}) {
		auth := &models.Auth{}
		zap.L().Info("mysql already exist", zap.String("table", auth.TableName()))
	}

	if migrat.HasTable(&models.TransactionMonitor{}) {
		monitor := &models.TransactionMonitor{}
		zap.L().Info("mysql already exist", zap.String("table", monitor.TableName()))
	}

	err = migrat.AutoMigrate(&models.User{}, &models.Cert{}, &models.Img{}, &models.Auth{}, &models.TransactionMonitor{})
	if err != nil {
		return
	}
	return
}

// Close 手动关闭连接
func Close() {
	_ = client.Close()
}
