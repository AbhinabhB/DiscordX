package server

import "net/http"

type loginHandler struct{}

func (handler *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("welcome to login page"))
}

type signupHandler struct{}

func (handler *signupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("welcome to signup page"))
}
