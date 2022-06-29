package users

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	Client *sql.DB
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "123456", "localhost", "users")
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
