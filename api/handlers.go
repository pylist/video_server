package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	io.WriteString(w, "Create User Handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	uname := p.ByName("user_name")
	io.WriteString(w, "你的用户名是"+uname)
}