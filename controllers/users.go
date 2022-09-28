package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Email: ", r.FormValue("email"), "\n")
	fmt.Fprint(w, "Passowrd: ", r.FormValue("password"))
}
