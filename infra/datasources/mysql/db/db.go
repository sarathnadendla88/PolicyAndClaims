package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_username = "root"
	mysql_password = "root"
	mysql_host     = "localhost"
	//mysql_host     = "127.0.0.1:3306"
	mysql_schema = "policy"
)

var (
	Client *sql.DB
	//userName=os.Getenv(mysql_username)
	//password=os.Getenv(mysql_password)
	//host=os.Getenv(mysql_host)
	//schema=os.Getenv(mysql_schema)

	userName = mysql_username
	password = mysql_password
	host     = mysql_host
	schema   = mysql_schema
)

func init() {
	// init db connection
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		userName, password, host, schema)

	log.Println(datasourceName)
	var err error

	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Print("DB connection successful")
}
