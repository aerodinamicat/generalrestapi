package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aerodinamicat/generalrestapi/database"
	"github.com/aerodinamicat/generalrestapi/repositories"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgresDockerContainerName string `envconfig:"POSGRES_DOCKERCONTAINER_NAME"`
	PostgresDB                  string `envconfig:"POSTGRES_DB"`
	PostgresUser                string `envconfig:"POSTGRES_USER"`
	PostgresPassword            string `envconfig:"POSTGRES_PASSWORD"`
	Port                        string `envconfig:"PORT"`
	JWTSecret                   string `envconfig:"JWT_SECRET"`
}

func main() {
	var config *Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("%v", err)
	}

	dbAddress := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode_disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDockerContainerName,
		config.PostgresDB,
	)
	dbServer, err := database.NewPostgresDatabaseRepository(dbAddress)
	if err != nil {
		log.Fatalf("%v", err)
	}
	repositories.SetDatabaseRepository(dbServer)
	defer repositories.CloseDatabaseConnection()

	router := mux.NewRouter()
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router); err != nil {
		log.Fatal(err)
	}
}
