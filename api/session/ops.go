package session

import (
	"log"
	"server/api/dbops"
	"server/api/defs"
	"server/api/utils"
	"sync"
	"time"
)

//定义一个Map来做Session缓存,sync.Map是线程安全的
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func GetSessionMap() *sync.Map {
	return sessionMap
}

//加载Session
func LoadSessionFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		log.Printf("LoadSessionFromDB err: %s", err)
		panic(err)
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

//创建Session
func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := noeInMilli()
	ttl := ct + 30*60*1000

	ss := &defs.SimpleSession{UserName: un, TTL: ttl}
	sessionMap.Store(id, ss)

	dbops.InsertSession(id, ttl, un)

	return id
}

//判断Session是否过期
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := noeInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).UserName, false
	}
	return "", true
}

func noeInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

//删除会话
func deleteExpiredSession(sid string) {
	//先删缓存，再删数据库
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}
