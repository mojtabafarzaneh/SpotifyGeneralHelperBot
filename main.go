package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"

	_ "github.com/lib/pq"
	"github.com/mojtabafarzaneh/SpotifyGeneralHelperBot/handlers"
)

func main() {
	// Connect to DB
	// dsn := "host=localhost port=5432 user=postgres password=password dbname=mydb sslmode=disable"
	// db, err := sql.Open("postgres", dsn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432/mydb")
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	clientID := "712630d2b9424750a0ea6c9af45a2165"
	clientSecret := "2166d6c5eec547ddb79b069e2869ffd9"
	clientGrantType := "client_credentials"
	redirectURI := "http://localhost:8080/callback"

	// Register callback handler
	http.HandleFunc("/callback", handlers.CallbackHandler(conn, clientID, clientGrantType, clientSecret, redirectURI))

	log.Println("Server running at :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
