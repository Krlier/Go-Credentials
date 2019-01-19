package main

import (
	"fmt"
	"os"
)

func main() {
	databaseName := "53CR3TD4T4B453"
	secretKey := "5UP3R53CR3T"
	secretPhrase := "Always know where your towel is."

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
