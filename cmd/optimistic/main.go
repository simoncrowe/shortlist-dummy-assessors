package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	profileFilePath := os.Getenv("PROFILE_PATH")
	notifyUrl := os.Getenv("NOTIFIER_URL")

	profileFile, err := os.Open(profileFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer profileFile.Close()

	if _, err := http.Post(notifyUrl, "application/json", profileFile); err != nil {
		log.Fatal(err.Error())
	}

	log.Print("Successful result written")
}
