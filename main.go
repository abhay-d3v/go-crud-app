package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var database *sql.DB

// db connection
func init() {
	var err error

	database, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS games (id INTEGER PRIMARY KEY, name TEXT, genre TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

func AddGame() {
	fmt.Print("Enter Name & Genre >> ")

	var name, genre string

	_, err := fmt.Scan(&name, &genre)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	statement, err := database.Prepare("INSERT INTO games (name, genre) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(name, genre)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var menuCode int

	for {
		fmt.Println("Select an option:")
		fmt.Println("1.) Add Game")
		fmt.Println("2.) View Games")
		fmt.Println("3.) Update Game")
		fmt.Println("4.) Delete Game")
		fmt.Println("5.) Quit")

		// Read user input
		fmt.Print(">> ")
		_, err := fmt.Scan(&menuCode)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch menuCode {
		case 1:
			fmt.Println("Adding a game...")
			AddGame()
			fmt.Println("")
		case 2:
			fmt.Println("2")
		case 3:
			fmt.Println("3")
		case 4:
			fmt.Println("4")
		case 5:
			fmt.Println("Exiting...")
			goto exit

		default:
			fmt.Println("\nError: Only values from the given menu are accepted")
		}

	}
exit:
	fmt.Println("\n------------------")
}
