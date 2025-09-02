package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"wells-go/infrastructure/config"
	"wells-go/infrastructure/database"
	"wells-go/interfaces/http"
	"wells-go/util/cors"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})

	cfg := config.GetConfig()
	fmt.Println("✅ Environment Loaded:", cfg.Environment)
	if cfg.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}

	fmt.Println("🔧 DB Config:")
	fmt.Println("DB_HOST:", cfg.DbHost)
	fmt.Println("DB_PORT:", cfg.DbPort)
	fmt.Println("DB_USER:", cfg.DbUser)
	fmt.Println("DB_NAME:", cfg.DbName)

	port, err := strconv.Atoi(cfg.DbPort)
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid DB port")
	}

	err = database.ConnectRawSQL(cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName, port)
	if err != nil {
		log.Fatal().Err(err).Msg("❌ Error connecting to PostgreSQL DB")
	}
	fmt.Println("✅ PostgreSQL (Raw SQL) connected successfully")

	r := gin.Default()
	r.Use(cors.CORSMiddleware(cfg))

	server, err := http.NewServer(database.GetPool(), cfg, r, cors.CORSMiddleware(cfg))
	if err != nil {
		log.Fatal().Err(err).Msg("❌ Cannot create HTTP server")
	}

	fmt.Println("🚀 Starting HTTP server on", cfg.HTTPServerAddress)
	err = server.Start(cfg.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("❌ Cannot start HTTP server")
	}
}
