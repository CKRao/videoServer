package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"server/scheduler/utils"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	sqlDriver := "mysql"

	log.Printf("%s Database connection start init.", sqlDriver)

	dataSourceConfig, err := utils.GetDataSourceConfig(sqlDriver)

	if err != nil {
		log.Printf("Get DataSource Config Failed.")
		panic(err)
	}

	//获取数据库参数
	userName := dataSourceConfig.UserName
	password := dataSourceConfig.Password
	url := dataSourceConfig.Url
	//拼接url
	dataSourceName := fmt.Sprintf("%s:%s@tcp%s", userName, password, url)
	dbConn, err = sql.Open(sqlDriver, dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	log.Printf("Database connection success.")
}
