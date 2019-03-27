package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"server/api/dbops"
	"server/api/defs"
	"server/api/session"
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
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}

func GetRouterJson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	router = "./routers.json"
	data, _ := ioutil.ReadFile(router)
	io.WriteString(w, string(data))
}
