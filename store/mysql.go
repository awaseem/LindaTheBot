package store

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/awaseem/LindaTheBot/helpers"
	// internal sql driver
	_ "github.com/go-sql-driver/mysql"
)

const sqlURL = "SQL_DATABASE_URL"
const insertSmt string = "INSERT INTO chats (firstName, lastName, username, message, date) VALUES (?, ?, ?, ?, ?)"

var dbConn *sql.DB

var databasePath = helpers.GetEnvOrElse(sqlURL, "root:test@/telegram?charset=utf8mb4")

// Init create initial connection
func Init() error {
	db, dbErr := sql.Open("mysql", databasePath)
	if dbErr != nil {
		return dbErr
	}
	dbConn = db
	tableErr := createTelegramTable()
	if tableErr != nil {
		fmt.Println(tableErr)
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
    (id INT PRIMARY KEY AUTO_INCREMENT, firstName TEXT, lastName TEXT, username TEXT, message TEXT, date DATETIME)
		CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
  `
	_, err := dbConn.Exec(sqlSmt)
	if err != nil {
		return err
	}
	return nil
}
