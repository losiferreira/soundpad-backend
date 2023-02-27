package handlers

import (
	"html/template"
	"net/http"
)

type RootHandler struct {
}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (r *RootHandler) HandleRoot(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("test_templates/google_sign_test.html")
	t.Execute(w, false)
}
