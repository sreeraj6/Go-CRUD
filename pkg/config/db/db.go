package db

import (
    "fmt"
    "database/sql"
     _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "root"
    dbname   = "go-test"
  )


func CreateDBConnection() {

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)

    if err != nil {
        fmt.Println("Something went wrong")
    }
    
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Database connected Successfully")
}