package utils

import (
	"github.com/shiyongxi/go-common/server"
	"gopkg.in/natefinch/lumberjack.v2"
)

type (
	Config struct {
		System *server.Config     `yaml:"system"`
		Logger *lumberjack.Logger `yaml:"logger"`
		//Mysql      *mysql.MysqlConf             `yaml:"mysql"`
		//Grpc       map[string]string            `yaml:"grpc"`
		//Remotes    *client.Config               `yaml:"remote"`
		//RocketMQ   *rocketmq.Config             `yaml:"rocketMQ"`
	}
)

var (
	config *Config
)

func Set(cfg *Config) {
	config = cfg
}

func GetConfig() *Config {
	return config
}
