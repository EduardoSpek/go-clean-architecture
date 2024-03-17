package mysql

import (
	"database/sql"
	"fmt"
	"os"
)

type Mysql struct {}

var conn Mysql

func (repo *Mysql) Connect() (*sql.DB, error) {    
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USERNAME") +":"+ os.Getenv("MYSQL_PASSWORD") +"@tcp("+ os.Getenv("MYSQL_HOST") +":"+ os.Getenv("MYSQL_PORT") +")/"+ os.Getenv("MYSQL_DB") +"")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}