package models

// Struct to represent a task
type Task struct {
	ID            int
	Title         string
	TimeSpent     int
	LastStartedAt int64
	Status        string
}
