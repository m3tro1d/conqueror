package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"conqueror/pkg/conqueror/infrastructure"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	c, err := parseEnv()
	if err != nil {
		log.Fatal(err)
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", c.DBUser, c.DBPassword, c.DBHost, c.DBName)

	db, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	dependencyContainer := infrastructure.NewDependencyContainer(context.Background(), db)
	server := infrastructure.NewServer(dependencyContainer)

	err = http.ListenAndServe(":"+c.Port, server.GetRouter())
	if err != nil {
		log.Fatal(err)
	}
}
