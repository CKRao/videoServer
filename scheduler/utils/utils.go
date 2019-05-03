package utils

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"strings"
)

var fileName = "config.json"
var relativePath = "C:/Users/clarkrao/go/src/server/scheduler/config.json"

type Config struct {
	ServerPort string     `json:"server_port"`
	DataSource DataSource `json:"data_source"`
}

type DataSource struct {
	DBName   string `json:"db_name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

//新建UUID
func NewUUID() (string, error) {
	uuid, err := uuid.NewV4()

	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}

//获取配置JSON数据
func getConfigJsonData() (*Config, error) {
	data, _ := ioutil.ReadFile(relativePath)
	cfg := &Config{}
	err := json.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

//获取数据库连接
func GetDataSourceConfig(dataBase string) (*DataSource, error) {
	cfg, err := getConfigJsonData()
	if err != nil {
		return nil, err
	}

	if !strings.EqualFold(dataBase, cfg.DataSource.DBName) {
		log.Fatal("ERROR DB")
		panic("ERROR DB")
	}

	return &cfg.DataSource, nil
}

//获取端口号
func GetSeverPort() (string, error) {
	cfg, err := getConfigJsonData()
	if err != nil {
		return "", err
	}
	portString := strings.Join([]string{":", cfg.ServerPort}, "")

	return portString, nil
}

//判断字符串是否为空
func IsStringEmpty(str string) bool {
	if strings.EqualFold("", str) || len(str) == 0 {
		return true
	}
	return false
}

//判断字符串是否为空
func IsStringsEmpty(str ...string) bool {

	for _, s := range str {
		if IsStringEmpty(s) {
			return true
		}
	}
	return false
}
