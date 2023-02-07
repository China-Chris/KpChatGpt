package configs

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	config  GlobalConfig
	rConfig sync.RWMutex
)

// MysqlConfig mysql配置参数
type MysqlConfig struct {
	User     string
	Password string
	Ip       string
	Port     string
	DbName   string
}

// GlobalConfig 全局配置
type GlobalConfig struct {
	Port  string
	Mysql MysqlConfig
	Debug bool
	Proxy string
	Eth   []string
	Bsc   []string
	Fibo  []string
}

// Config 返回配置文件
func Config() GlobalConfig {
	rConfig.RLock()
	configCopy := config
	rConfig.RUnlock()
	return configCopy
}

//加载配置文件
func ParseConfig(cfg string) {
	viper.SetConfigFile(cfg)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}
