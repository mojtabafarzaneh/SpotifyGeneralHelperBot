package main

import (
	"log"
	"net/http"

	"github.com/mojtabafarzaneh/SpotifyGeneralHelperBot/handlers"
)

func main() {

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/callback", handlers.CallbackHandler)
	http.HandleFunc("/albums", handlers.AlbumsHandler)

	log.Println("Server listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
