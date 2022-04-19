package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/ulascansenturk/go-clean-architecture/config"
	"github.com/ulascansenturk/go-clean-architecture/infrastructure/datastore"
	"github.com/ulascansenturk/go-clean-architecture/infrastructure/router"
	"github.com/ulascansenturk/go-clean-architecture/registry"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()
	r := registry.NewRegistry(db)
	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())
	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
