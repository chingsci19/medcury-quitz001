package main

import (
	"flag"

	"medcury/quitz001/modules/config"
	"medcury/quitz001/services/appointment"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	configPath = flag.String("config", "config_dev.toml", "Path of configuration file.")
	port       = flag.String("port", "8080", "Port for start server.")
	debug      = flag.Bool("debug", false, "sets log level to debug")
)

func main() {
	flag.Parse()
	config.LoadConfig(*configPath)
	logData := make(map[string]interface{})
	logData["tag"] = "main"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Info().Fields(logData).Msg("Server start at port : " + *port)
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4200, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(recover.New())
	app.Use(requestid.New())
	api := app.Group("/api")
	v1 := api.Group("/v1")

	appointment.Route(v1.Group("/appointment"))
	log.Fatal().Err(app.Listen(":" + *port))
}
