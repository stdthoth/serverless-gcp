package main

import (
	"fmt"

	"github.com/ichthoth/prod-api/internal/db"
)

// Run handles the instatiation of our application
func Run() error {
	fmt.Println("Starting our application")

	//Creates a New database
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to our database")
		return err
	}

	if err := db.Migrate(); err != nil {
		fmt.Println("failed to migrate db")
		return err
	}

	fmt.Println("successfully connected and pinged database")
	return nil

}

func main() {
	fmt.Println("GO REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
