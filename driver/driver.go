package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var db *sql.DB

func getConnectionString() string {
	host := os.Getenv("POSTGRES_HOST")
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, db)
}

func connectToDB() *sql.DB {
	psqlInfo := getConnectionString()
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Print(err)
		log.Print("Could not connect to Postgres.")
		return &sql.DB{}
	}
	return conn
}

// Database will attempt to connect to postgres db and return a db object on a successfull connection.
// If the connection to the db is lost it will continually retry to connect until a successful connection can be made.
func Database() *sql.DB {
	if db == nil { // during startup - if it does not exist, create it
		db = connectToDB()
	}
	err := db.Ping()
	for err != nil { // reconnect if we lost connection
		log.Print("Connection to Postgres was lost. Waiting for 5s...")
		db.Close()
		time.Sleep(5 * time.Second)
		log.Print("Reconnecting...")
		db = connectToDB()
		err = db.Ping()
		log.Print("Connection to Postgres successful.")
	}
	return db
}
