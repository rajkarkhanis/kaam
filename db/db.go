package db

import (
    "database/sql"
    "os"
    "fmt"
    "errors"
    "path/filepath"
    "modernc.org/sqlite"
    "modernc.org/sqlite/lib"
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
        title TEXT NOT NULL UNIQUE,
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

func AddTask(task models.Task) error {
    statement, err := database.Prepare("INSERT INTO tasks (title, time_spent, last_started_at, status) VALUES (?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer statement.Close()
    
    _, err = statement.Exec(task.Title, task.TimeSpent, task.LastStartedAt, task.Status)
    if err != nil {
        var sqliteErr *sqlite.Error
        if errors.As(err, &sqliteErr) && sqliteErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
            return fmt.Errorf("\nTask with title \"%s\" already exists", task.Title)
        }
        return fmt.Errorf("Failed to execute statement: %w", err)
    }

    return nil
}

func CloseDB() error {
    if database == nil {
        return nil
    }

    return database.Close()
}
