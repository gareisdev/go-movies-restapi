package models

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gareisdev/go-movies-restapi/pkg/utils"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

const pathFile string = "./data/movies.json"

func SaveData(movies []Movie) {

	utils.CreateFileIfDoesntExists(pathFile)
	content, err := json.Marshal(movies)
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile(pathFile, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func GetData() []Movie {
	utils.CreateFileIfDoesntExists(pathFile)

	var movies []Movie
	content, err := os.ReadFile(pathFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &movies)
	if err != nil {
		log.Fatal(err)
	}

	return movies
}
