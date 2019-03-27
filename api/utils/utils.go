package utils

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"strings"
)

var fileName = "config.json"
var relativePath = "config.json"

type Config struct {
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

//获取数据库连接
func GetDataSourceConfig(dataBase string) (*DataSource, error) {
	data, _ := ioutil.ReadFile(relativePath)
	cfg := &Config{}
	err := json.Unmarshal(data, cfg)
	if err != nil {
		panic(err)
	}

	if !strings.EqualFold(dataBase, cfg.DataSource.DBName) {
		log.Fatal("ERROR DB")
		panic("ERROR DB")
	}

	return &cfg.DataSource, nil
}
