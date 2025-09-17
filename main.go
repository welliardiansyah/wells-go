package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"wells-go/domain/migrate"
	"wells-go/infrastructure/config"
	"wells-go/infrastructure/database"
	"wells-go/infrastructure/redis"
	"wells-go/interfaces/http"
	"wells-go/util/cors"
	"wells-go/util/security"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})

	redis.InitRedis()

	cfg := config.GetConfig()
	fmt.Println("✅ Environment Loaded:", cfg.Environment)

	port, err := strconv.Atoi(cfg.DbPort)
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid DB port")
	}

	err = database.ConnectGorm(cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName, port)
	if err != nil {
		log.Fatal().Err(err).Msg("❌ Error connecting to PostgreSQL DB (GORM)")
	}
	fmt.Println("✅ PostgreSQL (GORM) connected successfully")

	db := database.GetGormDB()
	if err := migrate.Migrate(db); err != nil {
		log.Fatal().Err(err).Msg("❌ Gagal menjalankan migrasi")
	}

	r := gin.Default()
	r.Use(cors.CORSMiddleware(cfg))

	jwtMaker, err := security.NewJWTMaker(cfg.JWTSecret, cfg.JWTIssuer, redis.Rdb)
	if err != nil {
		log.Fatal().Err(err).Msg("❌ Cannot create JWT maker")
	}

	server, err := http.NewServer(db, cfg, r, cors.CORSMiddleware(cfg), jwtMaker)
	if err != nil {
		log.Fatal().Err(err).Msg("❌ Cannot create HTTP server")
	}

	fmt.Println("🚀 Starting HTTP server on", cfg.HTTPServerAddress)
	err = server.Start(cfg.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("❌ Cannot start HTTP server")
	}
}
