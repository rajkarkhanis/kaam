package cmd

import (
    "log"
    "fmt"
    "time"
    "strings"

    "kaam/db"
    "kaam/models"

    "github.com/spf13/cobra"
)

func formatTimeSpent(seconds int) string {
    h := seconds / 3600
    m := (seconds % 3600) / 60

    return fmt.Sprintf("%dh%dm", h, m)
}

func truncateTitle(title string, maxLength int) string {
    if len(title) > maxLength {
        return title[:maxLength-3] + "..."
    }

    return title
}

func showTasks(tasks []models.Task) {
    fmt.Printf("%-5s %-30s %-10s %-12s\n",
        "ID",
        "TITLE",
        "TIME",
        "STATUS",
    )
    fmt.Println(strings.Repeat("-", 65))

    for _, task := range tasks {
        elapsedTime := task.TimeSpent
        if task.Status == "IN PROGRESS" && task.LastStartedAt > 0 {
            elapsedTime += int(time.Now().Unix() - task.LastStartedAt)
        }

        fmt.Printf("%-5d %-30s %-10s %-12s\n",
            task.ID,
            truncateTitle(task.Title, 30),
            formatTimeSpent(task.TimeSpent),
            task.Status,
        )
    }
}
var rootCmd = &cobra.Command {
    Use: "kaam",
    Short: "Kaam is a simple CLI tool to track time on your tasks",
    Run: func (cmd *cobra.Command, args []string) {
        tasks, err := db.GetAllTasks()
        if err != nil {
            log.Fatal("Failed to fetch tasks: ", err)
        }

        showTasks(tasks)
    },
}

func Execute() {
    if err := db.InitDB(); err != nil {
        log.Fatal("Failed to initialise database: ", err)
    }
    defer db.CloseDB()

    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
