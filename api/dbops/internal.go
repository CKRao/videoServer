package dbops

import (
	"database/sql"
	"log"
	"server/api/defs"
	"strconv"
	"sync"
)

//插入Session
func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	insertSql := "INSERT INTO sessions (session_id,TTL,login_name) VALUES (?,?,?) "

	stmtIns, err := dbConn.Prepare(insertSql)
	if err != nil {
		log.Fatal("InsertSession err ：", err)
		return err
	}

	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		log.Fatal("InsertSession err ：", err)
		return err
	}

	defer stmtIns.Close()
	return nil
}

func RetriveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}

	selectSql := "SELECT TTL,login_name FROM sessions WHERE session_id = ?"
	stmtOut, err := dbConn.Prepare(selectSql)
	if err != nil {
		log.Fatal("RetriveSession err ：", err)
		return nil, err
	}
	var ttl string
	var uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)

	if err != nil && err != sql.ErrNoRows {
		log.Fatal("RetriveSession err ：", err)
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.UserName = uname
	} else {
		return nil, err
	}

	defer stmtOut.Close()

	return ss, nil
}

//获取所有的会话信息
func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	selectSql := "SELECT session_id, TTL,login_name FROM sessions"
	stmtOut, err := dbConn.Prepare(selectSql)
	if err != nil {
		log.Fatal("RetrieveAllSessions err ：", err)
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		log.Fatal("RetrieveAllSessions err ：", err)
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlstr string
		var uname string
		if err := rows.Scan(&id, &ttlstr, &uname); err != nil {
			log.Fatal("RetriveSession err ：", err)
			break
		}

		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err == nil {
			ss := &defs.SimpleSession{UserName: uname, TTL: ttl}
			m.Store(id, ss)
			log.Printf("session id:&s, ttl:%d", id, ss.TTL)
		}
	}
	defer stmtOut.Close()
	return m, nil
}

//删除会话
func DeleteSession(sid string) error {
	delSql := "DELETE FROM sessions WHERE session_id = ?"
	stmtDel, err := dbConn.Prepare(delSql)
	if err != nil {
		log.Fatal("DeleteSession err ：", err)
		return err
	}
	_, err = stmtDel.Exec(sid)

	if err != nil {
		return err
	}

	defer stmtDel.Close()

	return nil
}
