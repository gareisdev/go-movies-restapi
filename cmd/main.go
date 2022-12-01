package main

import (
	"fmt"
	"net/http"

	"github.com/gareisdev/go-movies-restapi/pkg/controllers"
	"github.com/gorilla/mux"
)

const PORT string = ":8000"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/movies", controllers.GetMovies).Methods("GET")
	r.HandleFunc("/api/movies/{id}", controllers.GetMovie).Methods("GET")
	r.HandleFunc("/api/movies", controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", controllers.UpdateMovies).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", controllers.DeleteMovies).Methods("DELETE")

	fmt.Printf("Server running at port %v\n", PORT)
	http.ListenAndServe(PORT, r)

}
