package dbops

import "log"

//添加用户
func AddUserCredential(loginName string, pwd string) error {
	sql := "INSERT INTO users (login_name,pwd) VALUES (?,?)"
	stmtIns, err := dbConn.Prepare(sql)

	if err != nil {
		log.Printf("AddUserCredential error: %s", err)
		return err
	}

	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()

	return nil
}

//获取用户
func GetUserCredential(loginName string) (string, error) {
	sql := "SELECT pwd FROM users WHERE login_name = ?"
	stmtOut, err := dbConn.Prepare(sql)

	if err != nil {
		log.Printf("GetUserCredential error: %s", err)
		return "", err
	}

	var pwd string

	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()

	return pwd, nil
}

//删除用户
func DeleteUser(login_name string,pwd string) error {
	sql := "DELETE FROM users WHERE login_name = ? AND pwd = ?"
	stmtDel, err := dbConn.Prepare(sql)

	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}

	stmtDel.Exec(login_name,pwd)
	stmtDel.Close()

	return nil
}
