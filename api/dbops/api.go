package dbops

import (
	"database/sql"
	"log"
)

//添加用户
func AddUserCredential(loginName string, pwd string) error {
	sql := "INSERT INTO users (login_name,pwd) VALUES (?,?)"
	stmtIns, err := dbConn.Prepare(sql)

	if err != nil {
		log.Printf("AddUserCredential error: %s", err)
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtIns.Close()

	return nil
}

//获取用户
func GetUserCredential(loginName string) (string, error) {
	get_sql := "SELECT pwd FROM users WHERE login_name = ?"
	stmtOut, err := dbConn.Prepare(get_sql)

	if err != nil {
		log.Printf("GetUserCredential error: %s", err)
		return "", err
	}

	var pwd string

	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()

	return pwd, nil
}

//删除用户
func DeleteUser(login_name string, pwd string) error {
	sql := "DELETE FROM users WHERE login_name = ? AND pwd = ?"
	stmtDel, err := dbConn.Prepare(sql)

	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}

	_, err = stmtDel.Exec(login_name, pwd)
	if err != nil {
		return err
	}

	defer stmtDel.Close()

	return nil
}
