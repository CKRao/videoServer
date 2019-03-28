package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"server/api/dbops"
	"server/api/defs"
	"server/api/session"
	"server/api/utils"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.UserName, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.UserName)
	su := &defs.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalError)
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	//判断参数是否为空，如果为空返回参数为空的错误响应
	if utils.IsStringsEmpty(ubody.UserName, ubody.Pwd) {
		sendErrorResponse(w, defs.ErrorParamsNullError)
		return
	}
	//查找用户，如果找不到返回未找到用户错误
	password, err := dbops.GetUserCredential(ubody.UserName)
	if err != nil {
		sendErrorResponse(w, defs.ErrorUserNotFoundError)
		return
	}
	//比较密码，如果密码不匹配，返回密码错误
	if !strings.EqualFold(ubody.Pwd, password) {
		sendErrorResponse(w, defs.ErrorPasswordWrongError)
		return
	}
	su := &defs.SignedUp{}
	//从会话里查找SessionId
	sid, err := dbops.RetriveSessionByLoginName(ubody.UserName)
	if err != nil {
		sid = session.GenerateNewSessionId(ubody.UserName)
	}
	su.SessionId = sid
	su.Success = true
	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalError)
	} else {
		sendNormalResponse(w, string(resp), 201)
	}

}
