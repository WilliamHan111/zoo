package conf

import (
	"os"

	"gopkg.in/yaml.v2"
)

// 配置
var AllConfig Config

// 全部配置
type Config struct {
	Mongo MongoConfig `yaml:"mongo"`
	Web   WebConfig   `yaml:"web"`
}

// MongoDB配置
type MongoConfig struct {
	Uri        string `json:"uri" yaml:"uri"`
	User       string `json:"user" yaml:"user"`
	Password   string `json:"password" yaml:"password"`
	DB         string `json:"db" yaml:"db"`
	Collection string `json:"collection" yaml:"collection"`
}

// web配置
type WebConfig struct {
	Host     string `json:"host" yaml:"host"`
	HttpPort string `json:"http_port" yaml:"http_port"`
}

// 初始化配置
func InitConfigs(configPath string) error {
	cfg, err := NewConfig(configPath)
	if err != nil {
		return err
	}
	AllConfig = *cfg
	return nil
}

// 读取并解析配置文件
func NewConfig(file string) (*Config, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
