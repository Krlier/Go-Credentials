package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnvVars() {
	err := godotenv.Load("credentials.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	getEnvVars()
	databaseName := os.Getenv("DATABASE_NAME")
	secretKey := os.Getenv("SECRET_KEY")
	secretPhrase := os.Getenv("SECRET_PHRASE")

	var dbName string
	var dbPass string

	fmt.Println("Please enter database name:")
	fmt.Scanf("%s", &dbName)

	fmt.Println("Please enter database password:")
	fmt.Scanf("%s", &dbPass)

	if dbName == databaseName && dbPass == secretKey {
		fmt.Println("Welcome to the database!")
		fmt.Println("Your secret phrase is: ", secretPhrase)
		os.Exit(0)
	}
	fmt.Println("Sorry, wrong database name or password")
}
