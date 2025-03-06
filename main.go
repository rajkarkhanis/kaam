package main

import (
    "log"
    "fmt"
    "kaam/db"
) 


func main() {
    _, err := db.InitDB()
    if err != nil {
        log.Fatal("Failed to initialise database: ", err)
    }
    fmt.Println("Connected to database successfully!")
}
