package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	MySQL  MySQLConfig  `yaml:"mysql"`
	JWT    JWTConfig    `yaml:"jwt"`
}
type ServerConfig struct {
	Addr string `yaml:"addr"`
}
type MySQLConfig struct {
	DSN string `yaml:"dsn"`
}

type JWTConfig struct {
	Secret string `yaml:"secret"`
}

func LoadConfig() (*Config, error) {

	path := os.Getenv("TODO_CONFIG")
	if path == "" {
		path = "configs/config.yaml"
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}
	var cfg Config
	err = yaml.Unmarshal(raw, &cfg)
	if err != nil {
		return nil, fmt.Errorf("配置文件解析失败：%w", err)
	}

	if dsn := os.Getenv("TODO_MYSQL_DSN"); dsn != "" {
		cfg.MySQL.DSN = dsn
	}
	if cfg.MySQL.DSN == "" || cfg.Server.Addr == "" || cfg.JWT.Secret == "" {
		return nil, fmt.Errorf("缺少必要的配置项:cfg.MySQL.DSN, cfg.Server.Addr, cfg.JWT.Secret")
	}
	return &cfg, nil

}
