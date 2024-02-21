package main

import (
	"backend/internal/server"
	"log"
)

func main() {
	log.Println("Starting Server")

	ser, err := server.SetupServer("127.0.0.1", "8080")
	if err != nil {
		log.Fatal("error occured during the configuration of serving, ", err)
	}
	err = ser.ListenAndServe()
	if err != nil {
		log.Fatal("error occured when starting the server ", err)
	}
}
