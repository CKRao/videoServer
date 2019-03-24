package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Creater User Handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}

func GetRouterJson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	router = "./api/config/routers.json"
	data, _ := ioutil.ReadFile(router)
	io.WriteString(w, string(data))
}
