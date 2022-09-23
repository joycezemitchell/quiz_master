package dataq

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	/*username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)*/

	username = "quipper"
	password = "qwer1234"
	host     = "localhost"
	schema   = "quizmaster"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println(err)
	}
	if err = Client.Ping(); err != nil {
		log.Println(err)
	}

	log.Println("database successfully configured")
}
