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

	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-gin/config"
	"github.com/jakkaphatminthana/go-gin/database"
	"github.com/jakkaphatminthana/go-gin/middleware"
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
	gin.SetMode(gin.DebugMode)
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
	corsMiddleware := middleware.CROSMiddleware(s.conf.Server.AllowOrigins)
	bodyLimitMiddleware := middleware.BodyLimitMiddleware(s.conf.Server.BodyLimit)
	timeoutMiddleware := middleware.TimeoutMiddleware(s.conf.Server.Timeout)
	errorHandlerMiddleware := middleware.ErrorHandlerMiddleware()
	authMiddleware := middleware.NewAuthorizationMiddleware(s.conf)

	s.engine.Use(gin.Recovery())
	s.engine.Use(gin.Logger())
	s.engine.Use(corsMiddleware)
	s.engine.Use(bodyLimitMiddleware)
	s.engine.Use(timeoutMiddleware)
	s.engine.Use(errorHandlerMiddleware)

	// routers
	s.engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})
	s.initTaskRouter(authMiddleware)
	s.initOAuth2Router()

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
