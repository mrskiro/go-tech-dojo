package main

import (
	"log"

	"github.com/k0kubun/sqldef"
	"github.com/k0kubun/sqldef/adapter"
	"github.com/k0kubun/sqldef/adapter/postgres"
	"github.com/k0kubun/sqldef/schema"
	"github.com/purp1eeeee/go-tech-dojo/config"
)

func main() {
	config, err := config.NewDBConfig()
	if err != nil {
		log.Fatalln(err)
		return
	}
	db, err := postgres.NewDatabase(adapter.Config{
		DbName:   config.DB,
		User:     config.User,
		Password: config.Password,
		Host:     config.Host,
		Port:     config.Port,
	})

	defer func() {
		_ = db.Close()
	}()

	if err != nil {
		log.Fatalln(err)
		return
	}
	sqldef.Run(schema.GeneratorModePostgres, db, &options)
}

var options = sqldef.Options{
	DesiredFile: "./db/schema.sql",
}
