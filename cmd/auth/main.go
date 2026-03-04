package main

import (
	"fmt"

	"github.com/evtepo/auth/internal/config"
	uHandler "github.com/evtepo/auth/internal/delivery/http/user"
	"github.com/evtepo/auth/internal/infrastructure/db"
	uRepo "github.com/evtepo/auth/internal/infrastructure/db/postgres"
	uUsecase "github.com/evtepo/auth/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

func main() {
	serviceCfg := config.NewConfig()
	config.LoadConfig("/internal/config/config.toml", serviceCfg)

	dbConnect := db.BuildConnection(&serviceCfg.Db)
	defer db.CloseConnection(dbConnect)

	userUC := uUsecase.NewUserUsecase(uRepo.NewRepository(dbConnect), uRepo.NewRoleRepository(dbConnect))
	userHandler := uHandler.NewHandler(userUC)

	server := gin.Default()
	userHandler.RegisterUserRoutes(server)

	server.Run(fmt.Sprintf(":%d", serviceCfg.App.Port))
}
