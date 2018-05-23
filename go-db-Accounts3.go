package main

import (
    "fmt"
    "database/sql"
    _ "mysql-master"
)
type person struct{
  username string
  password string
}
func main() {
    fmt.Println("Dealing with accounts database")

    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Accounts")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }else {
      fmt.Println("Connection Established Successfully..")
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()
    s1:=person{"saleel","saleel123"}
    insert, err := db.Query("INSERT INTO credentials VALUES (? , ?)",s1.username,s1.password)

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }else {
      fmt.Println("Data added successfully..")
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()


}
