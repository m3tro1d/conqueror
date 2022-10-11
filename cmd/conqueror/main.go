package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	conquerorapi "conqueror/api"
	"conqueror/pkg/conqueror/infrastructure"
	"conqueror/pkg/conqueror/infrastructure/transport"
	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	dependencyContainer, err := infrastructure.NewDependencyContainer(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}

	baseServer := grpc.NewServer()
	publicGRPCServer := transport.NewPublicGRPCServer(dependencyContainer)
	conquerorapi.RegisterConquerorServer(baseServer, publicGRPCServer)

	srv := transport.NewServer(dependencyContainer)

	server := startServer(":"+c.Port, srv)
	waitForKillSignal(getKillSignalChan())
	shutdownServer(server)
}

func startServer(serveURL string, srv *transport.Server) *http.Server {
	server := http.Server{
		Addr:    serveURL,
		Handler: srv.GetRouter(),
	}

	go func() {
		log.Print("Server is starting...")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(errors.Wrap(err, "error while serving HTTP"))
		}
	}()

	return &server
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	switch <-killSignalChan {
	case os.Interrupt:
		log.Print("got SIGINT...")
	case syscall.SIGTERM:
		log.Print("got SIGTERM...")
	}
}

func shutdownServer(server *http.Server) {
	log.Print("Server is shutting down...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}
