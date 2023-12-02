package driver

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// create variable to hold data of database
const maximumOpenConnection = 10
const maximumLifetime = 5 * time.Minute
const maxIdleDbConn = 5

var loopIterate = 0

// cretate function to init connection
func InitConnection(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("mysql", dsn)

	// check for an error
	if err != nil {
		log.Println("error when starting connection database : ", err)
		return nil, err
	}

	// loop if not yet connected
	for {
		if err != nil {
			log.Println("trying to connecting to database ...")
			loopIterate++
		} else {
			log.Println("success connected to datavase ...")
			// set db characteristic
			conn.SetMaxOpenConns(maximumOpenConnection)
			conn.SetConnMaxLifetime(maximumLifetime)
			conn.SetMaxIdleConns(maxIdleDbConn)

			// test for db connection
			err = conn.Ping()

			// check for an error
			if err != nil {
				log.Println("error when trying to ping database")
				return nil, err
			}

			return conn, nil
		}

		if loopIterate > 10 {
			log.Println("cannot connect to database")
			return conn, nil
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(1 * time.Second)
		continue
	}
}
