package cmd

import (
    "fmt"
    "log"

    "kaam/models"
    "kaam/db"

    "github.com/spf13/cobra"
)

var addCmd = &cobra.Command {
    Use: "add",
    Short: "Add a new task",
    Run: func (cmd *cobra.Command, args []string) {
        title, err := cmd.Flags().GetString("title")
        if err != nil {
            log.Fatal("Error getting title: ", err)
        }

        if title == "" {
            log.Fatal("Title cannot be empty")
        }

        task := models.Task {
            Title: title,
            TimeSpent: 0,
            LastStartedAt: 0,
            Status: "TODO",
        }

        err = db.AddTask(task)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("Task added: ", title)
    },
}
func init() {
    addCmd.Flags().String("title", "", "Title of the task")
    addCmd.MarkFlagRequired("title")
    rootCmd.AddCommand(addCmd)
}
