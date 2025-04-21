package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-gin/config"
	"github.com/jakkaphatminthana/go-gin/database"
	"github.com/jakkaphatminthana/go-gin/utils"
)

type ginServer struct {
	engine *gin.Engine
	conf   *config.Config
	db     database.Database
	server *http.Server
}

var (
	once   sync.Once
	server *ginServer
)

func NewGinServer(config *config.Config, db database.Database) *ginServer {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	once.Do(func() {
		server = &ginServer{
			engine: engine,
			conf:   config,
			db:     db,
		}
	})
	return server
}

func (s *ginServer) Start() {
	corsMiddleware := getCROSMiddleware(s.conf.Server.AllowOrigins)
	bodyLimitMiddleware := getBodyLimitMiddleware(s.conf.Server.BodyLimit)
	timeoutMiddleware := getTimeoutMiddleware(s.conf.Server.Timeout)

	s.engine.Use(gin.Recovery())
	s.engine.Use(gin.Logger())
	s.engine.Use(corsMiddleware)
	s.engine.Use(bodyLimitMiddleware)
	s.engine.Use(timeoutMiddleware)

	// routers
	s.engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	s.server = &http.Server{
		Addr:         ":" + s.conf.Server.Port,
		Handler:      s.engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// Start server in goroutine
	go s.startHTTPListener()

	// Gracefully shutdown
	s.gracefullyShutdown()
}

func getCROSMiddleware(allowOrigins []string) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func getBodyLimitMiddleware(bodyLimit string) gin.HandlerFunc {
	limitBytes := utils.ParseSize(bodyLimit)
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limitBytes)
		c.Next()
	}
}

func getTimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (s *ginServer) startHTTPListener() {
	log.Printf("✅ Server is running at port %s", s.conf.Server.Port)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("💥 Server failed: %v", err)
	}
}

func (s *ginServer) gracefullyShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("🚫 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("💥 Could not gracefully shutdown the server: %v", err)
	}

	log.Println("🚫 Server exited gracefully")
}
