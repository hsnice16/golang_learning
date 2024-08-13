package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/astaxie/beedb"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "himanshusingh"
	DB_PASSWORD = "postgres"
	DB_NAME     = "test"
)

type Userinfo struct {
	Uid        int `PK` // if the primary key is not id, you need to add tag `PK` for your customized primary key.
	Username   string
	Departname string
	Created    time.Time
}

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	orm := beedb.New(db, "pg")
	beedb.OnDebug = true

	// saveone := Userinfo{Username: "Test Add User", Departname: "Test Add Departname", Created: time.Now()}
	// orm.Save(&saveone)

	var saveone Userinfo
	saveone.Username = "Test Add User"
	saveone.Departname = "Test Add Departname"
	saveone.Created = time.Now()
	orm.Save(&saveone)
}
