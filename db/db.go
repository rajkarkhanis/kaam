package db

import (
    "database/sql"
    "os"
    "path/filepath"
    _ "modernc.org/sqlite"
)

func InitDB() (*sql.DB, error) {
    home, err := os.UserHomeDir()
    if err != nil {
        return nil, err
    }

    dbPath := filepath.Join(home, ".local", "share", "kaam", "kaam.db")
    err = os.MkdirAll(filepath.Dir(dbPath), 0755)
    if err != nil {
        return nil, err
    }

    db, err := sql.Open("sqlite", dbPath)
    if err != nil {
        return nil, err
    }

    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        time_spent INTEGER NOT NULL DEFAULT 0,
        last_started_at INTEGER NOT NULL DEFAULT 0,
        status TEXT NOT NULL DEFAULT 'TODO'
    )`)
    if err != nil {
        return nil, err
    }

    return db, nil
}
