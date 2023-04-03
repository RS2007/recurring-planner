package db

import (
	"better_calendar_backend/utils"
	"database/sql"
	"fmt"
	"os"

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
	connectionString := os.Getenv("POSTGRES_CONNECTION_URI")
	db, err := sql.Open("postgres", connectionString)
	utils.ErrorHandler(err)

	sqlInit := `
    CREATE TABLE IF NOT EXISTS users(
        userId SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE,
        accentColor TEXT
    ); 

    CREATE TABLE IF NOT EXISTS tags(
      tagId SERIAL PRIMARY KEY,
      name  TEXT NOT NULL UNIQUE
    );

    CREATE TABLE IF NOT EXISTS events(
      eventId SERIAL PRIMARY KEY,
      startTime TIME WITH TIME ZONE NOT NULL,
      endTime TIME WITH TIME ZONE NOT NULL,
      location TEXT,
      summary TEXT UNIQUE,
      description TEXT
    );


    CREATE TABLE IF NOT EXISTS templates(
      templateId SERIAL PRIMARY KEY,
      name TEXT NOT NULL UNIQUE,
      userId INTEGER NOT NULL,
      timeZone TEXT NOT NULL,
      FOREIGN KEY(userId)
        REFERENCES users (userId)
    );

    CREATE TABLE IF NOT EXISTS tag_templates(
      tagId INTEGER NOT NULL,
      FOREIGN KEY(tagId)
        REFERENCES tags (tagId),
      templateId INTEGER NOT NULL,
      FOREIGN KEY(templateId)
        REFERENCES templates (templateId)
    );

    CREATE TABLE IF NOT EXISTS event_templates(
      eventId INTEGER NOT NULL,
      FOREIGN KEY(eventId)
        REFERENCES events (eventId),
      templateId INTEGER NOT NULL,
      FOREIGN KEY(templateId)
        REFERENCES templates (templateId)
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
