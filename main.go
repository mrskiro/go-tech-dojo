package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/purp1eeeee/go-tech-dojo/config"
	"github.com/purp1eeeee/go-tech-dojo/db"
	httpHandler "github.com/purp1eeeee/go-tech-dojo/handlers/http"
	"github.com/purp1eeeee/go-tech-dojo/handlers/http/middlewares"
	"github.com/purp1eeeee/go-tech-dojo/oapi"
	"github.com/purp1eeeee/go-tech-dojo/repositories/tokens"
	"github.com/purp1eeeee/go-tech-dojo/repositories/tx"
	"github.com/purp1eeeee/go-tech-dojo/repositories/users"
	"github.com/purp1eeeee/go-tech-dojo/usecases/user"
)

func main() {
	config, err := config.NewDBConfig()
	if err != nil {
		log.Fatalln(err)
		return
	}
	db, err := db.NewDB(config)
	if err != nil {
		log.Fatalln(err)
		return
	}

	txRepo := tx.NewTxRepository(db)
	tokensRepo := tokens.NewTokenRepository(db)
	usersRepo := users.NewUserRepository(db)

	userUseCase := user.NewUserUseCase(txRepo, tokensRepo, usersRepo)
	handlers := httpHandler.NewHandlers(userUseCase)

	middlewares := middlewares.NewMiddlewares(tokensRepo)

	r := chi.NewRouter()
	r.Use(middlewares.SetOperaterID)
	r.Use(middlewares.SetUserIDByToken)

	handler := oapi.HandlerFromMux(handlers, r)

	mux := http.NewServeMux()
	mux.Handle("/", handler)
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("ping")
	})
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
