package db

import (
    "database/sql"
    "os"
    "path/filepath"
    _ "modernc.org/sqlite"
    "kaam/models"
)

var database *sql.DB

func InitDB() error {
    home, err := os.UserHomeDir()
    if err != nil {
        return err
    }

    dbPath := filepath.Join(home, ".local", "share", "kaam", "kaam.db")
    err = os.MkdirAll(filepath.Dir(dbPath), 0755)
    if err != nil {
        return err
    }

    database, err = sql.Open("sqlite", dbPath)
    if err != nil {
        return err
    }

    _, err = database.Exec(`CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        time_spent INTEGER NOT NULL DEFAULT 0,
        last_started_at INTEGER NOT NULL DEFAULT 0,
        status TEXT NOT NULL DEFAULT 'TODO'
    )`)
    if err != nil {
        return err
    }

    return nil
}

func GetAllTasks() ([]models.Task, error) {
    rows, err := database.Query("SELECT id, title, time_spent, last_started_at, status FROM tasks")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var task models.Task
        if err := rows.Scan(&task.ID, &task.Title, &task.TimeSpent, &task.LastStartedAt, &task.Status); err != nil {
            return nil, err
        }

        tasks = append(tasks, task)
    }

    return tasks, nil
}
func CloseDB() error {
    if database == nil {
        return nil
    }

    return database.Close()
}
