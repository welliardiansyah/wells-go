package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
	"time"
	"wells-go/infrastructure/config"
	"wells-go/interfaces/http/permission"
	"wells-go/interfaces/http/role"
	"wells-go/interfaces/http/users"
	"wells-go/util/cors"
	"wells-go/util/security"
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
	Config   *config.Config
	DB       *gorm.DB
	Engine   *gin.Engine
	Security security.Maker
}

func NewServer(db *gorm.DB, config *config.Config, r *gin.Engine, corsMiddleware gin.HandlerFunc, securityMaker security.Maker) (*Server, error) {
	log.Println("[INFO] Initializing server...")

	if corsMiddleware != nil {
		r.Use(corsMiddleware)
	}

	server := &Server{
		DB:       db,
		Config:   config,
		Engine:   r,
		Security: securityMaker,
	}

	log.Println("[INFO] Setting up router...")
	server.setupRouter()
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

func (server *Server) setupRouter() {
	router := server.Engine

	router.Use(cors.CORSMiddleware(server.Config))
	router.Use(LoggingMiddleware())

	// ====================================================================================
	// Routes Group
	// ====================================================================================
	publicRoutes := router.Group("")

	users.RouteUsers(server.DB, publicRoutes, server.Config, server.Security)
	role.RouteRoles(server.DB, publicRoutes, server.Config, server.Security)
	permission.RoutePermissions(server.DB, publicRoutes, server.Config, server.Security)

	server.Engine = router
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
