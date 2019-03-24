package utils

import (
	"github.com/Unknwon/goconfig"
	"github.com/satori/go.uuid"
	"log"
)

var conf_path = "../config/conf.ini"

//新建UUID
func NewUUID() (string, error) {
	uuids, err := uuid.NewV4()

	if err != nil {
		return "", err
	}

	return uuids.String(), nil
}

//获取数据库连接
func GetDataSourceConfig(dataBase string) (map[string]string, error) {
	cfg, err := goconfig.LoadConfigFile(conf_path)
	if err != nil {
		return nil, err
	}
	//获取配置文件信息
	section, err := cfg.GetSection(dataBase)
	if err != nil {
		return nil, err
	}
	for k, v := range section {
		log.Printf("key : %s    val : %s ", k, v)
	}
	return section, nil

}
