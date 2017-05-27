package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Init() {
	port := 9000

	router := mux.NewRouter()
	router.HandleFunc("/", createAppHandler).Methods("POST")

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		panic(err)
	}
}

func createAppHandler(w http.ResponseWriter, r *http.Request) {

}
