package main
 
import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
 
func main() {
    db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:port)/database_name")
    if err != nil {
        panic(err)
    }
    defer db.Close()
 
    err = db.Ping()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
 
    fmt.Printf("Database connected!")
}