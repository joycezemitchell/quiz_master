package schema

import (
	"github.com/joycezemitchell/quiz_master/datasource/mysql/dataq"
	"log"
)

func CreateSchema() {
	sql := make(map[int]string)

	sql[0] = `CREATE TABLE IF NOT EXISTS  questions (
			id INT NOT NULL,
			question TEXT,
			answer TEXT
		);
	`

	log.Println("Creating database tables")

	for _, v := range sql {
		stmt, err := dataq.Client.Prepare(v)
		if err != nil {
			log.Println("error when trying to prepare save user statement", err)
		}

		defer stmt.Close()
		_, saveErr := stmt.Exec()
		if saveErr != nil {
			log.Println("error when trying to save user", saveErr)
		}
	}

}