package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gareisdev/go-movies-restapi/pkg/models"
	"github.com/gorilla/mux"
)

func GetMovies(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	var movies = models.GetData()
	json.NewEncoder(writter).Encode(movies)
}

func GetMovie(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(request)
	var movies = models.GetData()
	for idx, movie := range movies {
		if movie.ID == parameters["id"] {
			json.NewEncoder(writter).Encode(movies[idx])
			return
		}
	}
}

func CreateMovie(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	var movies = models.GetData()

	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000000))
	movies = append(movies, movie)

	models.SaveData(movies)

	json.NewEncoder(writter).Encode(movie)
}

func UpdateMovies(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(request)
	var movies = models.GetData()

	for idx, movie := range movies {
		if movie.ID == parameters["id"] {
			var updatedMovie models.Movie
			_ = json.NewDecoder(request.Body).Decode(&updatedMovie)

			movieToReplace := movies[idx]

			updatedMovie.ID = movieToReplace.ID
			movies[idx] = updatedMovie

			models.SaveData(movies)

			json.NewEncoder(writter).Encode(updatedMovie)
			return
		}
	}
}

func DeleteMovies(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(request)
	var movieDeleted models.Movie
	var movies = models.GetData()

	for idx, movie := range movies {
		if movie.ID == parameters["id"] {
			movieDeleted = movies[idx]
			movies = append(movies[:idx], movies[idx+1:]...)
			models.SaveData(movies)
			json.NewEncoder(writter).Encode(movieDeleted)
			break
		}
	}
}
