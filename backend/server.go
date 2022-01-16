package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/williamjoseph77/evos/handlers"
	"github.com/williamjoseph77/evos/postgres"

	"github.com/gorilla/mux"
)

func main() {
	db := postgres.Connect()
	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	defer db.Close()

	db.AddQueryHook(postgres.DBLogger{})

	h := handlers.Handler{
		Database: db,
	}

	var router = mux.NewRouter()
	const port string = ":8080"

	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Currently up and running..")
	}).Methods("GET")

	router.HandleFunc("/api/characters", h.HandleGetCharacterList).Methods("GET")
	router.HandleFunc("/api/characters/secure", h.HandleCreateCharacterSecure).Methods("POST")
	router.HandleFunc("/api/characters/non-secure", h.HandleCreateCharacterNonSecure).Methods("POST")
	router.HandleFunc("/api/roles", h.HandleGetRoleList).Methods("GET")

	log.Println("Server listening on port ", port)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	corsHandler := c.Handler(router)
	log.Fatalln(http.ListenAndServe(port, corsHandler))
}
