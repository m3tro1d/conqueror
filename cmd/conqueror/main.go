package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"conqueror/pkg/conqueror/infrastructure"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	c, err := parseEnv()
	if err != nil {
		panic(err)
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", c.DBName, c.DBPassword, c.DBHost, c.DBName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dependencyContainer := infrastructure.NewDependencyContainer()
	server := infrastructure.NewServer(dependencyContainer)
	log.Fatal(http.ListenAndServe(":8080", server.GetRouter()))
}
