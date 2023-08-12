package main

import (
	"context"
	"time"

	"github.com/Sup3r-Us3r/barber-server/config"
	_ "github.com/Sup3r-Us3r/barber-server/docs"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/database"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/web/handler"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/web/webserver"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase"
	"github.com/Sup3r-Us3r/barber-server/log"
)

// @title                       BarberShop API
// @version                     1.0
// @description                 BarberShop API
// @termsOfService              http://swagger.io/terms
// @contact.name                Mayderson Mello
// @contact.url                 https://mayderson.me
// @contact.email               maydersonmello@gmail.com
// @license.name                MIT
// @license.url                 https://github.com/Sup3r-Us3r/go-clean-arch-template/blob/main/LICENSE
// @host                        localhost:8080
// @BasePath                    /
// @securityDefinitions.apikey   BearerAuth
// @in                          header
// @name                        Authorization
func main() {
	configs := config.LoadConfig(".")

	log.NewLogger(configs)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.NewMongoDBConnection(ctx, configs.MongoURI, configs.MongoDBName)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	webserver := webserver.NewWebServer(":" + configs.ServerPort)

	repositoryContainer := repository.GetMongoDBRepositories(db)
	useCaseContainer := usecase.GetUseCases(*repositoryContainer)
	handler.NewHandlerContainer(*useCaseContainer, *webserver)

	log.Info("HTTP SERVER RUNNING ON PORT :" + configs.ServerPort)
	webserver.StartServer()
}
