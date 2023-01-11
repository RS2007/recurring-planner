package db

import (
	"better_calendar_backend/utils"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type DB struct {
	Db          *sql.DB
	Initialized bool
}

var DbStructInstance DB

func dbInitialize() {

	err := godotenv.Load()
	utils.ErrorHandler(err)
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB_NAME")
	const (
		host = "localhost"
	)
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	utils.ErrorHandler(err)

	sqlInit := `
    CREATE TABLE IF NOT EXISTS users(
        userId SERIAL PRIMARY KEY,
        name TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        email TEXT NOT NULL,
        accentColor TEXT,
        apiKey TEXT
    ); 
    CREATE TABLE IF NOT EXISTS templates(
      templateId INTEGER PRIMARY KEY,
      name TEXT NOT NULL,
      startDateTime TIMESTAMP NOT NULL,
      startTimeZone TEXT NOT NULL,
      endDateTime TIMESTAMP NOT NULL,
      endTimeZone TEXT NOT NULL,
      location TEXT,
      summary TEXT,
      description TEXT
    );
  `
	result, err := db.Exec(sqlInit)
	fmt.Println(result)
	utils.ErrorHandler(err)
	DbStructInstance.Db = db
	DbStructInstance.Initialized = true

}

func GetDB() *sql.DB {
	if !DbStructInstance.Initialized {
		dbInitialize()
	}
	return DbStructInstance.Db
}
