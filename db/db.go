package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/purp1eeeee/go-tech-dojo/config"
)

func NewDB(config config.DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v sslmode=%s",
		config.Host, config.User, config.Password, config.DB, config.Port, "disable")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Fatalln("PingError: ", err)
	}
	return db, nil
}
