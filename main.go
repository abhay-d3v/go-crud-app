package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
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

func CheckIDExists(id int) bool {
	var exists bool
	_ = database.QueryRow("SELECT EXISTS(SELECT 1 FROM games WHERE id = ?)", id).Scan(&exists)
	return exists
}

func AddGame() {
	var name, genre string

	fmt.Print("Enter Name & Genre >> ")

	_, err := fmt.Scan(&name, &genre)
	if err != nil {
		fmt.Println("[Error]:", err)
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

	if err := statement.Close(); err != nil {
		log.Fatal(err)
	}
}

func ViewAll() {
	rows, err := database.Query("SELECT id, name, genre FROM games")
	if err != nil {
		log.Fatal(err)
	}

	var id int
	var name string
	var genre string

	for rows.Next() {
		err := rows.Scan(&id, &name, &genre)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(strconv.Itoa(id) + ": " + name + " " + genre)
	}

	if err := rows.Close(); err != nil {
		log.Fatal(err)
	}
}

func Update() {
	var id int
	var name, genre string

	fmt.Print("Enter ID, Name & Genre >> ")

	_, err := fmt.Scan(&id, &name, &genre)
	if err != nil {
		fmt.Println("[Error]:", err)
		return
	}

	if !CheckIDExists(id) {
		fmt.Println("[Error]: ID does not exist!")
		return
	}

	statement, err := database.Prepare("UPDATE games SET name = ?, genre = ? WHERE id = ?;")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(name, genre, id)
	if err != nil {
		log.Fatal(err)
	}

	if err := statement.Close(); err != nil {
		log.Fatal(err)
	}
}

func Delete() {
	var id int

	fmt.Print("Enter ID >> ")

	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("[Error]:", err)
		return
	}

	if !CheckIDExists(id) {
		fmt.Println("[Error]: ID does not exist!")
		return
	}

	statement, err := database.Prepare("DELETE FROM games WHERE id = ?;")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	if err := statement.Close(); err != nil {
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
			AddGame()
			fmt.Println("")
		case 2:
			ViewAll()
			fmt.Println("")
		case 3:
			Update()
			fmt.Println("")
		case 4:
			Delete()
			fmt.Println("")
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
