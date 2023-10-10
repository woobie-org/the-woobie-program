package main

import (
    "database/sql"
    "fmt"
    "log"
)

func main() {
    tmp := "EcwyZC8TQj8R"
    db, err := sql.Open("postgres", fmt.Sprintf("user=myuser password=%v dbname=mydb sslmode=disable", tmp))
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
}
