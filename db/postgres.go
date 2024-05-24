package db

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

func ConnectPostgresDB(connStr string) (*sql.DB, error) {
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    err = db.Ping()
    if err != nil {
        return nil, err
    }
    log.Println("Connected to PostgreSQL database")
    return db, nil
}

func ClosePostgresDB(db *sql.DB) {
    if db != nil {
        db.Close()
        log.Println("Disconnected from PostgreSQL database")
    }
}
