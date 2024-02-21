package server

import (
	"log"
	"net"
	"net/http"
	"path"
	"time"
)

var (
	addr string
)

type ServerHandler struct {
	signupHandler *signupHandler
	loginHandler  *loginHandler
}

func (handler ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	basePath := path.Clean(url)[1:]

	switch basePath {
	case "login":
		handler.loginHandler.ServeHTTP(w, r)
	case "signup":
		handler.signupHandler.ServeHTTP(w, r)
	case "":
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte("hello world"))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte("http method not supported"))
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("404 NOT FOUND"))
	}
}

func SetupServer(ip string, port string) (*http.Server, error) {
	if len(ip) > 0 {
		addr += ip
	}
	addr += ":"
	if len(port) > 0 {
		addr += port
	} else {
		log.Println("using port 8000 by default")
		addr += "8000"
	}
	if len(addr) > 21 {
		return nil, &net.AddrError{Err: "there is some issue with the address", Addr: addr}
	}

	log.Printf("setting up TCP Listener Server on http://%s\n", addr)

	serverHandler := &ServerHandler{}
	server := http.Server{
		Handler:      serverHandler,
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return &server, nil
}
