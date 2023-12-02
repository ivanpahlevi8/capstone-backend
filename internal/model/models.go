package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type DbModel struct {
	conn *sql.DB
}

// create function to init db model
func InitDbModel(conn *sql.DB) *DbModel {
	// create model
	var dbModel DbModel

	// assign db model
	dbModel.conn = conn

	return &dbModel
}

// create model for object at database
type ModelObj struct {
	ID        int    `json:"id"`
	Orang     int    `json:"orang"`
	Timestamp int    `json:"timestamp"`
	Waktu     string `json:"waktu"`
}

// create function to get single data by id
func (db *DbModel) GetDataById(id int) (ModelObj, error) {
	// create context
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)

	// defer cancle
	defer cancle()

	// create object to be return
	var rtrnObj ModelObj

	// create variable to hold data
	var getId int
	var getOrang int
	var getTimeStamp int
	var getWaktu string

	// crete query
	querytxt := `SELECT * FROM Halte WHERE id=?`

	// execute query
	query := db.conn.QueryRowContext(ctx, querytxt, id)

	// scan data
	err := query.Scan(
		&getId,
		&getOrang,
		&getTimeStamp,
		&getWaktu,
	)

	// check for an error
	if err != nil {
		log.Println("error when querying to get single data")
		return rtrnObj, err
	}

	// assign column to objecy
	rtrnObj = ModelObj{
		ID:        getId,
		Orang:     getOrang,
		Timestamp: getTimeStamp,
		Waktu:     getWaktu,
	}

	// check another error from query
	err = query.Err()
	if err != nil {
		log.Println("error when querying to get single data in query")
		return rtrnObj, err
	}

	// if okay
	return rtrnObj, nil
}

// create function to get all data
func (db *DbModel) GetAllData() ([]ModelObj, error) {
	// create context
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)

	// defer cancle
	defer cancle()

	// create variable to hold object
	var allData []ModelObj
	var singleData ModelObj

	// create variable to hold data from column
	var getId int
	var getOrang int
	var getTimeStamp int
	var getWaktu string

	// create query
	querytxt := `SELECT * FROM Halte WHERE id=?`

	// do query
	query, err := db.conn.QueryContext(ctx, querytxt)

	// check for an error
	if err != nil {
		log.Println("error when querying obejct from database")
		return allData, err
	}

	// do query
	for query.Next() {
		err = query.Scan(
			&getId,
			&getOrang,
			&getTimeStamp,
			&getWaktu,
		)

		// check for an error
		if err != nil {
			log.Println("error when querying to get single data")
			return allData, err
		}

		// assign column to objecy
		singleData = ModelObj{
			ID:        getId,
			Orang:     getOrang,
			Timestamp: getTimeStamp,
			Waktu:     getWaktu,
		}

		allData = append(allData, singleData)
	}

	// check for query error
	if err = query.Err(); err != nil {
		log.Println("error when querying to get single data")
		return allData, err
	}

	// if all okay
	return allData, nil
}

// create function to delete data
func (db *DbModel) DeletDataById(id int) error {
	// create context
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)

	// defer cancle
	defer cancle()

	// create query
	query := `DELETE FROM Halte WHERE id=?`

	// execute query
	_, err := db.conn.ExecContext(ctx, query, id)

	// check for an error
	if err != nil {
		log.Println("error when deleting obejct in database")
		return err
	}

	// if okay
	return nil
}
