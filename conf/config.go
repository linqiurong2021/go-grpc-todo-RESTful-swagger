package conf

import "gopkg.in/ini.v1"

// AppConfig AppConfig
type AppConfig struct {
	Port        string `ini:"port"`
	GRPCPort    string `ini:"gRPCPort"`
	SwaggerPort string `ini:"swaggerPort"`
}

// DBConfig DBConfig
type DBConfig struct {
	Host     string `ini:"host"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Sechma   string `ini:"sechma"`
}

// Config 配置文件
type Config struct {
	AppConfig `ini:"app"`
	DBConfig  `ini:"db"`
}

// Conf 配置文件
var Conf = new(Config)

// InitConfig 初始化
func InitConfig(file string) error {
	return ini.MapTo(Conf, file)
}
