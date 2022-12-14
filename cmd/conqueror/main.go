package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"conqueror/pkg/conqueror/infrastructure"
	"conqueror/pkg/conqueror/infrastructure/mysql"
	"conqueror/pkg/conqueror/infrastructure/transport"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func main() {
	c, err := parseEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := connectDB(c)
	if err != nil {
		log.Fatal(err)
	}

	err = migrateDB(c, db)
	if err != nil {
		log.Fatal(err)
	}

	publicAPI, err := createPublicAPI(db, c)
	if err != nil {
		log.Fatal(err)
	}

	server := startServer(":"+c.Port, publicAPI)
	waitForKillSignal(getKillSignalChan())
	shutdownServer(server)
}

func connectDB(c *config) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBName,
	)

	db, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrateDB(c *config, db *sqlx.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := db.Connx(ctx)
	if err != nil {
		return err
	}

	migrator := mysql.NewMigrator(ctx, c.MigrationsDir, conn)
	err = migrator.MigrateUp()
	if err != nil {
		return err
	}

	return err
}

func createPublicAPI(db *sqlx.DB, c *config) (transport.PublicAPI, error) {
	dependencyContainer, err := infrastructure.NewDependencyContainer(context.Background(), db, c.FilesDir)
	if err != nil {
		return nil, err
	}

	return transport.NewPublicAPI(dependencyContainer, []byte(c.Secret)), nil
}

func startServer(serveURL string, publicAPI transport.PublicAPI) *http.Server {
	server := http.Server{
		Addr:    serveURL,
		Handler: transport.NewRouter(publicAPI),
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
