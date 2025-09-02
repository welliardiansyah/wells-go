package http

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
	"wells-go/infrastructure/config"
	"wells-go/util/cors"
)

func init() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatalf("[FATAL] Failed to load timezone 'Asia/Jakarta': %v", err)
	}
	time.Local = loc

	if err := os.MkdirAll("app/logs", os.ModePerm); err != nil {
		log.Printf("[ERROR] Failed to create logs directory: %v", err)
	}

	logFileName := fmt.Sprintf("app/logs/%s.log", time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("[FATAL] Failed to open log file %s: %v", logFileName, err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	log.SetOutput(multiWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("[INFO] Logging initialized. Outputting to terminal and file.")
}

type Server struct {
	Config      *config.Config
	Store       *sql.DB
	Engine      *gin.Engine
	RouterGroup *gin.RouterGroup
}

func NewServer(db *sql.DB, config *config.Config, r *gin.Engine, cors gin.HandlerFunc) (*Server, error) {
	log.Println("[INFO] Initializing server...")

	if cors != nil {
		r.Use(cors)
	}

	server := &Server{
		Store:  db,
		Config: config,
		Engine: r,
	}

	log.Println("[INFO] Setting up router...")
	server.setupRouter(db)
	log.Println("[INFO] Router setup completed.")

	return server, nil
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("[DEBUG] Request to %s %s", c.Request.Method, c.Request.URL.Path)
		log.Printf("[DEBUG] Headers: %v", c.Request.Header)
		c.Next()
	}
}

func (server *Server) setupRouter(db *sql.DB) {
	router := gin.Default()
	log.Println("[INFO] Gin router initialized.")

	router.Use(cors.CORSMiddleware(server.Config))
	router.Use(LoggingMiddleware())

	// ====================================================================================
	// Public Routes Group: These routes do NOT require JWT or Android Key authentication.
	// ====================================================================================
	//publicRoutes := router.Group("")
	//log.Println("[INFO] Setting up BNI route (Public)...")
	//bniRoute.SetupBNIRouter(db, publicRoutes, server.Config)
}

func (server *Server) Start(address string) error {
	log.Printf("[INFO] Starting server at %s...", address)
	err := server.Engine.Run(address)
	if err != nil {
		log.Printf("[ERROR] Failed to start server: %v", err)
	} else {
		log.Println("[INFO] Server is running.")
	}
	return err
}
