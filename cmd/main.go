package main

import (
	"fmt"
	"log"
	"net/http"
	"real-time-voting/internal/auth"
	"real-time-voting/internal/websocket"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("voting system")

	router := mux.NewRouter()

	router.HandleFunc("/register", auth.Register)
	router.HandleFunc("/login", auth.Login)

	router.HandleFunc("/ws", websocket.HandleConnections)

	log.Println("server started on :8000")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("listenAnd Serve:", err)
	}
}
