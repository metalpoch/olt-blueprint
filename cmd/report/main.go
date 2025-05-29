package main

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"

	"github.com/goccy/go-json"
	"github.com/metalpoch/olt-blueprint/common/database"
	"github.com/metalpoch/olt-blueprint/common/model"
	"github.com/metalpoch/olt-blueprint/common/pkg/openstreetmap"
	"github.com/metalpoch/olt-blueprint/common/pkg/tracking"
	"github.com/metalpoch/olt-blueprint/report/router"
)

var cfg model.Config

func init() {
	fileJSON := os.Getenv("CONFIG_JSON")
	if fileJSON == "" {
		log.Panicln("CONFIG_JSON env is required")
	}

	f, err := os.ReadFile(fileJSON)
	if err != nil {
		log.Panicln(err)
	}

	json.Unmarshal([]byte(f), &cfg)
}

func main() {
	telegram := tracking.SmartModule{
		URL: cfg.SmartModuleTelegramURL,
	}

	openstreetmap := openstreetmap.OSM{
		URL: cfg.SmartModuleOSMURL,
	}

	db := database.Connect(cfg.DatabaseURI, cfg.IsProduction)
	server := fiber.New(fiber.Config{
		StructValidator: &model.StructValidator{Validator: validator.New()},
		JSONEncoder:     json.Marshal,
		JSONDecoder:     json.Unmarshal,
		BodyLimit:       100 * 1024 * 1024, // 100 mb
	})

	server.Use(logger.New())
	server.Use(cors.New())

	router.Setup(server, db, telegram, openstreetmap, cfg.StaticReportDirectory)
	server.Listen(":3002")
}
