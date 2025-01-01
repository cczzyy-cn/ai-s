package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// 定义配置结构体
type Config struct {
	DeepSeekAPIURL string `json:"deepseek_api_url"`
	APIKey         string `json:"api_key"`
}

// 加载配置文件
func LoadConfig(filename string) (*Config, error) {
	// 读取文件
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	// 解析 JSON
	var cfg Config
	if err := json.Unmarshal(file, &cfg); err != nil {
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}

	return &cfg, nil
}
