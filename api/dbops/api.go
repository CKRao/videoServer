package dbops

import (
	"database/sql"
	"log"
	"server/api/defs"
	"server/api/utils"
	"time"
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

//新建视频信息
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	//	create uuid
	vid, err := utils.NewUUID()
	log.Print("uuid:", vid)
	if err != nil {
		return nil, err
	}

	t := time.Now()
	//M D y, HH:MM:SS 这个格式化日期时间不能更改，必须是这个时间点
	ctime := t.Format("Jan 02 2006,15:04:05")
	insSql := `INSERT INTO video_info 
			(id,author_id,name,display_ctime) VALUES(?,?,?,?)`

	stmtIns, err := dbConn.Prepare(insSql)

	if err != nil {
		log.Printf("AddNewVideo error: %s", err)
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	defer stmtIns.Close()

	return res, nil
}

//获取视频信息
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	getSql := "SELECT author_id, name, display_ctime FROM video_info WHERE id = ?"
	stmtOut, err := dbConn.Prepare(getSql)

	var aid int
	var dct string
	var name string

	if err != nil {
		log.Printf("GetVideoInfo error: %s", err)
		return nil, err
	}

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer stmtOut.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

//删除视频信息
func DeleteVideoInfo(vid string) error {
	delSql := "DELETE FROM video_info WHERE id = ?"

	stmtDel, err := dbConn.Prepare(delSql)
	if err != nil {
		log.Printf("DeleteVideoInfo error: %s", err)
		return err
	}

	_, err = stmtDel.Exec(vid)

	if err != nil {
		return err
	}

	defer stmtDel.Close()

	return nil
}
