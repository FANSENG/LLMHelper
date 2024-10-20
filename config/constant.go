package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	APIKEY  string `json:"APIKEY"`
	BaseURI string `json:"BaseURI"`
	URI     string `json:"URI"`
}

const (
	configPath = "config/config.json"

	Mode_GLM4Flash = "glm-4-flash"
)

var (
	ModeMap = map[string]Config{}
)

func InitConfig() error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("read config failed: %v", err)
	}
	var configMap map[string]Config
	err = json.Unmarshal(data, &configMap)
	if err != nil {
		return fmt.Errorf("unmarshal config failed: %v", err)
	}
	for mode, config := range configMap {
		ModeMap[mode] = config
	}
	return nil
}
