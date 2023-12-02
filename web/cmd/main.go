package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/backend-capstone/internal/driver"
	models "github.com/backend-capstone/internal/model"
)

// set web port
const webPort = "8080"

type Config struct {
	DB models.DbModel
}

func main() {
	fmt.Println("Starting backend server")

	dsn := "root:03052001ivan@tcp(localhost:3305)/widgets?parseTime=true&tls=false"
	conn, err := driver.InitConnection(dsn)

	// check there is connection
	if conn == nil {
		log.Fatal("connectio failed, returne empty connection to database")
		return
	}

	if err != nil {
		log.Println("error when connecting data")
		return
	}

	// create config
	myConfig := Config{
		DB: *models.InitDbModel(conn),
	}

	// create server
	serv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: myConfig.Routes(),
	}

	err = serv.ListenAndServe()

	if err != nil {
		log.Fatal("Error happen : ", err)
	}
}
