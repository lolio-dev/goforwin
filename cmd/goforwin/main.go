package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"goforwin/pkg/server"
	"log"
	"net/http"
	"os"
)

var (
	headersOk = handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk = handlers.AllowedOrigins([]string{os.Getenv("http://localhost:3000")})
	methodsOk = handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/socket", server.WsEndpoint)
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headersOk, originsOk, methodsOk)(router)))
}

// 	game := core.NewGame()
//	p1 := core.NewPlayer("Lolio")
//	p2 := core.NewPlayer("Soso")
//
//	p1.JoinGame(game)
//	p2.JoinGame(game)
//
//	game.PlacePawn(6, p1)
//	game.PlacePawn(5, p2)
//	game.PlacePawn(5, p1)
//	game.PlacePawn(4, p1)
//	game.PlacePawn(4, p1)
//	game.PlacePawn(4, p1)
//	game.PlacePawn(3, p2)
//	game.PlacePawn(3, p1)
//	game.PlacePawn(3, p1)
//	game.PlacePawn(3, p1)
//
//	for _, raw := range game.Grid {
//		fmt.Println(raw)
//	}
//	fmt.Println(game.CheckPawnWin(&core.PawnPosition{6, 5}))
