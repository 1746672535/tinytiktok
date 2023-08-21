package config

import (
	"github.com/spf13/viper"
	"tinytiktok/utils/msg"
)

type Config struct {
	Viper *viper.Viper
}

// NewConfig 初始化配置
func NewConfig(path, name, ty string) *Config {
	v := viper.New()
	// 设置读取的文件路径
	v.AddConfigPath(path)
	// 设置读取的文件名
	v.SetConfigName(name)
	// 设置文件的类型
	v.SetConfigType(ty)
	// 尝试进行配置读取
	err := v.ReadInConfig()
	if err != nil {
		panic(msg.UnableReadConfig)
	}
	return &Config{Viper: v}
}

// ReadInt 读取配置
func (config *Config) ReadInt(name string) int {
	return config.Viper.GetInt(name)
}

// ReadString 读取配置
func (config *Config) ReadString(name string) string {
	return config.Viper.GetString(name)
}

// ReadStringSlice 读取配置
func (config *Config) ReadStringSlice(name string) []string {
	return config.Viper.GetStringSlice(name)
}

// ReadAny 读取配置
func (config *Config) ReadAny(name string) any {
	return config.Viper.Get(name)
}
