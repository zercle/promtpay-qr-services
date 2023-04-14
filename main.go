package main

import (
	"flag"
	"log"
	"os"

	"github.com/zercle/promtpay-qr-services/pkg/config"
	server "github.com/zercle/promtpay-qr-services/internal/infrastructure"
)

var (
	version string
	build   string
	runEnv  string
)

// @title Promptpay QR Services
// @version 0.1
// @description This is a sample promtpay QR generator in GO
// @contact.name Kawin Viriyaprasopsook
// @contact.email kawin.vir@zercle.tech
// @contact.url https://github.com/zercle/promtpay-qr-services
// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	// Running flag
	if len(os.Getenv("ENV")) != 0 {
		runEnv = os.Getenv("ENV")
	} else {
		flagEnv := flag.String("env", "dev", "A config file name without .env")
		flag.Parse()
		runEnv = *flagEnv
	}
	if err := config.LoadConfig(runEnv); err != nil {
		log.Fatalf("error while loading the env:\n %+v", err)
	}

	server, err := server.NewServer(version, build, runEnv)
	if err != nil {
		log.Fatalf("error while create server:\n %+v", err)
	}

	server.Run()
}
