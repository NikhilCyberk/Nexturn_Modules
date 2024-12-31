package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDatabase() error {
    var err error
    DB, err = sql.Open("sqlite3", "./inventory.db")
    if err != nil {
        return fmt.Errorf("failed to open database: %v", err)
    }

    if err = DB.Ping(); err != nil {
        return fmt.Errorf("failed to ping database: %v", err)
    }

    // Create products table
    _, err = DB.Exec(`CREATE TABLE IF NOT EXISTS products (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT,
        price REAL NOT NULL,
        stock INTEGER NOT NULL,
        category_id INTEGER
    )`)
    if err != nil {
        return fmt.Errorf("failed to create products table: %v", err)
    }

    log.Println("Database initialized successfully")
    return nil
}

func GetDB() *sql.DB {
    return DB
}