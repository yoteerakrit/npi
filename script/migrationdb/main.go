package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// https://github.com/golang-migrate/migrate/tree/master/database/mysql
	cmd := os.Args[1]
	if cmd != "migrate" && cmd != "drop" {
		panic("invalid cmd")
	}

	dir, _ := os.Getwd()
	db, _ := sql.Open("mysql",
		fmt.Sprintf("%s:%s@%s(%s:%d)/%s?multiStatements=true",
			"root",
			"keep1234",
			"tcp",
			"0.0.0.0",
			3306,
			"keeplearning",
		),
	)

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", path.Join(dir, "script", "migrationdb", "migrations")), // migration file path
		"mysql",
		driver,
	)

	if cmd == "migrate" {
		if err := m.Up(); err != nil {
			fmt.Println(err.Error())
			panic(err)
		}

		log.Printf("up success")
	}

	if cmd == "drop" {
		if err := m.Drop(); err != nil {
			panic(err)
		}

		log.Printf("drop success")
	}

}
