package main

import (
	"fmt"

	"github.com/evtepo/auth/internal/config"
	"github.com/evtepo/auth/internal/infrastructure/db"
)

func main() {
	serviceCfg := config.NewConfig()
	config.LoadConfig("/internal/config/config.toml", serviceCfg)

	dbConnect := db.BuildConnection(&serviceCfg.Db)
	fmt.Println(*dbConnect)
	
}