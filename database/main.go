package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "himanshusingh"
	DB_PASSWORD = "postgres"
	DB_NAME     = "test"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("# Inserting values")

	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username, department, created) VALUES($1, $2, $3) returning uid;", "hsnice16", "CS", "2024-08-13").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("last inserted id=", lastInsertId)

	fmt.Println("# Updating")
	stmt, err := db.Prepare("UPDATE userinfo SET username=$1 WHERE uid=$2")
	checkErr(err)

	res, err := stmt.Exec("hsnice16-update", lastInsertId)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time

		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)

		fmt.Println("uid | username | department | created")
		fmt.Printf("%3v | %8v | %10v | %7v\n", uid, username, department, created)
	}

	fmt.Println("# Deleting")
	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=$1")
	checkErr(err)

	res, err = stmt.Exec(lastInsertId)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
