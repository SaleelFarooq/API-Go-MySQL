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

    // perform a db.Query insert
    results, err := db.Query("SELECT username,password FROM credentials")
  	if err != nil {
  		panic(err.Error()) // proper error handling instead of panic in your app
  	}

  	for results.Next() {
  		var tag person
  		// for each row, scan the result into our tag composite object
  		err = results.Scan(&tag.username, &tag.password)
  		if err != nil {
  			panic(err.Error()) // proper error handling instead of panic in your app
  		}
                  // and then print out the tag's Name attribute
  		fmt.Printf("User ==> %s --- Password ==> %s\n ",tag.username,tag.password)
  	}


}
