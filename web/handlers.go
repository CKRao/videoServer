package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

type HomePage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	hp := &HomePage{Name: "clarkrao"}
	t, e := template.ParseFiles("./templates/home.html")

	if e != nil {
		log.Printf("Parse templates home.html error:%s", e)
	}

	t.Execute(w, hp)
}

func userHomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func apiHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}
