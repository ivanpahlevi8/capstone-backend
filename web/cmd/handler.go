package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// create model as response
type ResponseStatus struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// cretae handler function to get data by id
func (app *Config) GetDataById(w http.ResponseWriter, r *http.Request) {
	// set header as json
	w.Header().Set("Content-Type", "application/json")

	// get params from request
	urlQuery := r.URL.Query()

	// get id from params
	getIdStr := urlQuery.Get("id")

	// convert id string into integer
	getId, err := strconv.Atoi(getIdStr)

	// check for an error
	if err != nil {
		log.Println("error when converting id string into integer")
		app.ErrorJsonResponse(w, http.StatusInternalServerError, err)
		return
	}

	// get data from id
	getData, err := app.DB.GetDataById(getId)

	// check for an error
	if err != nil {
		log.Println("error when get data from database")
		app.ErrorJsonResponse(w, http.StatusInternalServerError, err)
		return
	}

	// create response
	resp := ResponseStatus{
		Status: http.StatusAccepted,
		Data:   getData,
	}

	// if all okay write response to user
	app.WriteJsonObject(w, resp, http.StatusAccepted)
}

// cretae function to get all data
func (app *Config) GetAllData(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	// get all data from database
	allData, err := app.DB.GetAllData()

	// check for an error
	if err != nil {
		log.Println("error when getting all data from database")
		app.ErrorJsonResponse(w, http.StatusInternalServerError, err)
		return
	}

	// create response
	resp := ResponseStatus{
		Status: http.StatusAccepted,
		Data:   allData,
	}

	// send response to user
	app.WriteJsonObject(w, resp, http.StatusAccepted)
}

// cretate function to delete data in database
func (app *Config) DeleteDataById(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	// get params from request
	urlQuery := r.URL.Query()

	// get id from params
	getIdStr := urlQuery.Get("id")

	// convert id string into integer
	getId, err := strconv.Atoi(getIdStr)

	// check for an error
	if err != nil {
		log.Println("error when converting id string into integer")
		app.ErrorJsonResponse(w, http.StatusInternalServerError, err)
		return
	}

	// delete data from datavase
	err = app.DB.DeletDataById(getId)

	// check for an error
	if err != nil {
		log.Println("error when deleting data from database")
		app.ErrorJsonResponse(w, http.StatusInternalServerError, err)
		return
	}

	// create response
	resp := ResponseStatus{
		Status: http.StatusAccepted,
		Data:   fmt.Sprintf("success deleting data with id : %d", getId),
	}

	// send response
	app.WriteJsonObject(w, resp, http.StatusAccepted)
}

/**
	-
	-
	-
	create helper function to all handler function
	-
	-
	-
**/

// create function to response an error
func (app *Config) ErrorJsonResponse(w http.ResponseWriter, httpStatus int, err error) {
	// set status
	w.WriteHeader(httpStatus)

	// creaet error payload
	jsonResp := ResponseStatus{
		Status: http.StatusInternalServerError,
		Data:   fmt.Sprintf("error happen with status code : %d and with error : %s", httpStatus, err.Error()),
	}

	// marshalling response
	objJson, err := json.MarshalIndent(jsonResp, "", "\t")

	// check for an error
	if err != nil {
		log.Println("error when marshalling obejct to json")
		return
	}

	// send json
	w.Write(objJson)
}

// create function to write obejct
func (app *Config) WriteJsonObject(w http.ResponseWriter, item interface{}, status int, header ...http.Header) error {
	// set header as json response
	w.Header().Set("Content-Type", "application/json")

	// check if there is header or not
	if len(header) > 0 {
		for k, v := range header[0] {
			w.Header()[k] = v
		}
	}

	// set header status
	w.WriteHeader(status)

	// create json object
	jsonObject, err := json.MarshalIndent(item, "", "\t")

	// check for an error
	if err != nil {
		log.Println("error when converting object to json")
		return err
	}

	// write to output
	_, err = w.Write(jsonObject)

	// check for an error
	if err != nil {
		log.Println("error when write to http output")
		return err
	}

	return nil
}
