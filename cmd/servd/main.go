package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pr1s10n3r/microurl/cmd/servd/routes"
	"github.com/pr1s10n3r/microurl/internal/platform/database"
	"github.com/pr1s10n3r/microurl/internal/url"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Infof("unable to load dotenv file: %s. Ignoring", err)
	}

	dbConn := database.Connection{
		URL: os.Getenv("DATABASE_URL"),
	}

	if err := dbConn.Connect(); err != nil {
		log.Fatalf("unable to connect with database: %s", err)
	}
	defer dbConn.Close()

	router := routes.Router{
		UrlRepo: url.NewUrlRepositoryImpl(dbConn),
	}

	if err := router.Start(); err != nil {
		log.Fatalf("unable to start router: %s", err)
	}
}
