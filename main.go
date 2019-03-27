package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

//注册路由
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	handles, err := GetRouters()
	if err != nil {
		panic(err)
	}
	for path, info := range handles {
		switch info.Method {
		case GET:
			router.GET(path, info.Handler)
		case POST:
			router.POST(path, info.Handler)
		case PUT:
			router.PUT(path, info.Handler)
		case DELETE:
			router.DELETE(path, info.Handler)
		default:
			log.Printf("not allow this method!")
		}
	}
	return router
}
func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)

}

/*
handler -> validation{1.request,2.user} -> business logic -> response
1. data model
2. error handling
*/
