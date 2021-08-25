package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName  string      `json:"app_name"`
	AppMode  string      `json:"app_mode"`
	AppHost  string      `json:"app_host"`
	AppPort  string      `json:"app_port"`
	Sms      SmsConfig   `json:"sms"`
	Database Database    `json:"database"`
	Redis    RedisConfig `json:"redis_config"`
}

var _cfg *Config

type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
	RegionId     string `json:"region_id"`
}
type Database struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	Showsql  bool   `json:"showsql"`
}
type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

func GetCofig() *Config {
	return _cfg
}
func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&_cfg)
	if err != nil {
		return nil, err
	}
	//fmt.Println(_cfg)
	return _cfg, nil
}
