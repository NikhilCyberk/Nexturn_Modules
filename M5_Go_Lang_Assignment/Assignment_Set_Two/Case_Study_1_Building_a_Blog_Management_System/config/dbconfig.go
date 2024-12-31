package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDatabase() error {
    var err error
    DB, err = sql.Open("sqlite3", "./Blog_Management_System.db")
    if err != nil {
        return fmt.Errorf("failed to open database: %v", err)
    }

    if err = DB.Ping(); err != nil {
        return fmt.Errorf("failed to ping database: %v", err)
    }

    fmt.Println("Connected to the database")

    // Create blogs table
    _, err = DB.Exec(`CREATE TABLE IF NOT EXISTS blogs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        author TEXT NOT NULL,
        timestamp TEXT NOT NULL
    )`)
    if err != nil {
        return fmt.Errorf("failed to create blogs table: %v", err)
    }

    // Create users table
    _, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL
    )`)
    if err != nil {
        return fmt.Errorf("failed to create users table: %v", err)
    }

    log.Println("Database tables created successfully")
    return nil
}

func GetDB() *sql.DB {
    if DB == nil {
        log.Fatal("Database not initialized")
    }
    return DB
}