package utils

import (
	"log"
	"os"
)

func CreateFileIfDoesntExists(path string) {
	f, err := os.OpenFile(path, os.O_RDWR, 0664)
	if err != nil {
		log.Printf("The file doesn't exists.")
		log.Printf("Creating file.")

		myfile, e := os.Create(path)
		if e != nil {
			log.Fatal("Error we can't create the file: ", e)
		}

		myfile.WriteString("[]")
		if err := myfile.Close(); err != nil {
			log.Fatal("Error we can't close the file: ", err)
		}
		return
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
