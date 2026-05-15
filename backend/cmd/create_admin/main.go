package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	_ "github.com/lib/pq"
	"database/sql"

)

func main() {
	password := "bobo19751020"
	email := "boboidvtw@gmail.com"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", "host=localhost port=5432 user=liyungchih dbname=sub2api sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Delete existing user if any
	_, _ = db.Exec("DELETE FROM users WHERE email = $1", email)

	_, err = db.Exec(`INSERT INTO users (email, password_hash, role, status, username, notes, balance, concurrency, total_recharged, signup_source) 
		VALUES ($1, $2, 'admin', 'active', 'admin', 'Manually created admin via Go', 0, 100, 0, 'email')`,
		email, string(hash))
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Admin user created successfully with Go-generated bcrypt hash.")
}
