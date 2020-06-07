package driver

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "user_sor"
)

// var db *sql.DB

// // ConnectDB Creates a postgres db connection
// func ConnectDB() *sql.DB {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	var err error
// 	db, err = sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		log.Print(err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		log.Print(err)
// 	}

// 	return db
// }

var db *sql.DB // this is "internal", as in: we should NOT use this directly

func connectToDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Print(err)
		log.Print("Could not connect to Postgres.")
		return &sql.DB{}
	}
	return conn
}

// Database returns a connection to the db and will rety to connect if it is unable to
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
	}
	return db
}
