package app

import (
	"fmt"
	"log"
	"os"

	"bacurin.de/recipeDB/backend/middlewares"
	"bacurin.de/recipeDB/backend/models"
	"golang.org/x/crypto/bcrypt"
)

// Start initialize and start everything
func Start() {
	len := len(os.Args)
	if len == 2 && os.Args[1] == "hashPassword" {
		fmt.Println("Enter password")
		var password string
		_, err := fmt.Scan(&password)
		if err != nil {
			log.Println(err)
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			log.Print(err)
		}

		fmt.Println("Salted Hash: ", string(hash))
		os.Exit(0)
	}

	if len == 2 && os.Args[1] == "noAuth" {
		fmt.Println("Turning authentication off")
		middlewares.NoAuth = true
	}

	_, err := models.Model.Initialize()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	route()

	r.Run()
}
