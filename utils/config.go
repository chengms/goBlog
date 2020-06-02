package utils

import (
	"encoding/json"
	"os"
)

type WebConfig struct {
	AppName		string `json:"app_name"`
	Port		string `json:"port"`
	StaticPath	string `json:"static_path"`
	Mode 		string `json:"mode"`
}

// 初始化配置
func InitConfig() *WebConfig {
	file, err := os.Open("./configs/config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := WebConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return &conf
}
