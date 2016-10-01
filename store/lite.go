package store

import (
	"database/sql"
	"log"
	"time"
	// internal sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

const databasePath string = "./LindaDataStore"
const insertSmt string = "INSERT INTO chats (firstName, lastName, username, message, date) VALUES (?, ?, ?, ?, ?)"

var dbConn *sql.DB

// Init create initial connection
func Init() error {
	db, dbErr := sql.Open("sqlite3", databasePath)
	if dbErr != nil {
		return dbErr
	}
	dbConn = db
	tableErr := createTelegramTable()
	if tableErr != nil {
		log.Fatal("Error: failed to create database table for telegram chat!")
	}
	return nil
}

// Save message props to datastore
func Save(firstName string, lastName string, username string, message string, date time.Time) error {
	_, err := dbConn.Exec(insertSmt, firstName, lastName, username, message, date)
	if err != nil {
		return err
	}
	return nil
}

func createTelegramTable() error {
	sqlSmt := `
    create table if not exists chats 
    (id INTEGER PRIMARY KEY AUTOINCREMENT, firstName TEXT, lastName TEXT, username TEXT, message TEXT, date DATETIME);
  `
	_, err := dbConn.Exec(sqlSmt)
	if err != nil {
		return err
	}
	return nil
}
