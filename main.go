package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nebtown/gmancloud/pkg/gameserver"
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("beep boop"))
}

func bootServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	server := gameserver.NewGameServer(vars["game"], vars["instance"])
	server.Start()

	w.Write([]byte("Trying to start server"))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/servers/{game}/{instance}", bootServer).Methods("PUT")

	router.Use(loggingMiddleware)

	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":80", router))
}
