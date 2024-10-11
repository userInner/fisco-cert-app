package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConf)

type AppConf struct {
	Name             string `mapstructure:"name"`
	Version          string `mapstructure:"version"`
	StartTime        string `mapstructure:"start_time"`
	MachineID        int64  `mapstructure:"machine_id"`
	Port             int    `mapstructure:"port"`
	Mode             string `mapstructure:"mode"`
	Local            string `mapstructure:"local"`
	*AuthConfig      `mapstructure:"auth"`
	*LogConfig       `mapstructure:"log"`
	*MySQLConfig     `mapstructure:"mysql"`
	*RedisConfig     `mapstructure:"redis"`
	*Tencent         `mapstructure:"tencent"`
	*FiscoBcosConfig `mapstructure:"fisco-bcos-config"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type AuthConfig struct {
	PwdSecret   string `mapstructure:"pwd_secret"`
	TokenSecret string `mapstructure:"token_secret"`
	TokenExpire int64  `mapstructure:"token_expire"`
}

type Tencent struct {
	CosRegion  string `mapstructure:"cos_region"`
	GetService string `mapstructure:"get_service"`
	SecretId   string `mapstructure:"secret_id"`
	SecretKey  string `mapstructure:"secret_key"`
}

type FiscoBcosConfig struct {
	TLSCa           string
	TLSKeyFile      string
	TLSCertFile     string
	TLSSMEnKeyFile  string
	TLSSMEnCertFile string
	IsSMCrypto      bool
	PrivateKey      string
	GroupID         string
	Host            string
	Port            int
	DisableSSL      bool
}

func Init(filePath string) (err error) {
	// 1. 加载配置文件
	viper.SetConfigFile(filePath)
	err = viper.ReadInConfig()

	// 读取配置失败
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	// 2.序列化到结构体中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\v", err)
	}
	// 开启实时检测配置文件是否更新
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("%v file already change", filePath)
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\v", err)
		}
	})
	return
}
