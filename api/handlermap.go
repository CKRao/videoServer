package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
)

type jsonData struct {
	Routers []Routers `json:"routers"`
}

type Routers struct {
	Path    string `json:"path"`
	Method  string `json:"method"`
	Handler string `json:"handler"`
}

type RouterInfo struct {
	Path    string
	Method  string
	Handler httprouter.Handle
}

var (
	HandlerMap = make(map[string]httprouter.Handle)
	router     = "routers.json"
)

func init() {
	HandlerMap["CreateUser"] = CreateUser
	HandlerMap["Login"] = Login
}

//获取路由信息
func GetRouters() (map[string]RouterInfo, error) {
	data, err := ioutil.ReadFile(router)
	if err != nil {
		return nil, err
	}
	m := &jsonData{}
	json.Unmarshal(data, &m)
	routers := m.Routers

	mapResult := make(map[string]RouterInfo)

	for _, e := range routers {
		info := RouterInfo{Path: e.Path, Method: e.Method, Handler: HandlerMap[e.Handler]}
		mapResult[e.Path] = info
	}
	return mapResult, nil
}
