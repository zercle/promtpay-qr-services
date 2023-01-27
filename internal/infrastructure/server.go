package infrastructure

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/segmentio/encoding/json"
	"github.com/spf13/viper"
)

type Server struct {
	PrdMode bool
	Version string
	Build   string
	RunEnv  string
}

func NewServer(version, buildTag, runEnv string) (server *Server, err error) {
	// Init server
	server = &Server{
		PrdMode: (viper.GetString("app.env") == "production"),
		Version: version,
		Build:   buildTag,
		RunEnv:  runEnv,
	}

	return
}

func (s *Server) Run() (err error) {
	app := fiber.New(fiber.Config{
		ErrorHandler:   customErrorHandler,
		ReadTimeout:    60 * time.Second,
		ReadBufferSize: 8 * 1024,
		Prefork:        s.PrdMode,
		// speed up json with segmentio/encoding
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Logger middleware for Fiber that logs HTTP request/response details.
	app.Use(logger.New())

	// Recover middleware for Fiber that recovers from panics anywhere in the stack chain and handles the control to the centralized ErrorHandler.
	app.Use(recover.New(recover.Config{EnableStackTrace: !s.PrdMode}))

	// CORS middleware for Fiber that that can be used to enable Cross-Origin Resource Sharing with various options.
	app.Use(cors.New())

	// App Handlers
	s.SetupRoutes(app)

	// Log GO_ENV
	log.Printf("Runtime ENV: %s", viper.GetString("app.env"))
	log.Printf("Version: %s", s.Version)
	log.Printf("Build: %s", s.Build)

	// Listen from a different goroutine

	// Listen HTTP
	go func() {
		if err := app.Listen(":" + viper.GetString("app.port.http")); err != nil {
			log.Panic(err)
		}
	}()

	// Create channel to signify a signal being sent
	quit := make(chan os.Signal, 1)
	// When an interrupt or termination signal is sent, notify the channel
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// This blocks the main thread until an interrupt is received
	<-quit
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")
	// Your cleanup tasks go here
	fmt.Println("Successful shutdown.")
	return
}
