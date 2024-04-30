package main

import (
	"log"

	"github.com/Angstreminus/exchanger/internal/server"
	"github.com/Angstreminus/exchanger/pkg/config"
	"github.com/Angstreminus/exchanger/pkg/logger"
	_ "github.com/swaggo/swag/example/celler/docs"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	log := logger.MustInitLogger(&cfg)
	server := server.NewServer(&cfg, log)
	server.MustRunWithGracefullShutdown()
}
