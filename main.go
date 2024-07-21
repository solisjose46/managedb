package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	username := flag.String("u", "", "Username for the new user")
	password := flag.String("p", "", "Password for the new user")
	dbPath := flag.String("d", "", "Path to the SQLite database")
	flag.Parse()

	if *username == "" || *password == "" || *dbPath == "" {
		fmt.Println("Usage: -u USERNAME -p PASSWORD -d DATABASE_PATH")
		os.Exit(1)
	}

	// Hash the password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the SQLite database
	db, err := sql.Open("sqlite3", *dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert the new user
	insertUserQuery := `INSERT INTO users (username, password) VALUES (?, ?)`
	_, err = db.Exec(insertUserQuery, *username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User registered successfully")
}
