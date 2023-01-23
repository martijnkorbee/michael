package cmd

import (
	"fmt"
	"log"

	"github.com/upper/db/v4/adapter/sqlite"
)

func mustConnectToDB() {
	// connect to database
	sess, err := sqlite.Open(sqlite.ConnectionURL{
		Database: fmt.Sprintf("%s/mortgage.db", app.rootpath),
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	// test connection
	if err = sess.Ping(); err != nil {
		log.Fatalln(err.Error())

	}

	log.Println("connected to database")

	// assign db to app
	app.database = sess
}

func mustPrepareDB(name string) {
	// remove current table if exists
	_, err := app.database.SQL().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s_mortgage;", name))
	if err != nil {
		log.Fatalln(err.Error())
	}

	// create new table
	file := fmt.Sprintf("files/migrations/%s_mortgage_tables.sql", name)

	q, err := binFS.ReadFile(file)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = app.database.SQL().Exec(string(q))
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("prepared database for %s mortgage\n", name)
}
