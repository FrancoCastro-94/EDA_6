package main

import (
	"flag"
	"log"

	json_models "github.com/FrancoCastro-94/web-socket/models"
	"github.com/FrancoCastro-94/web-socket/strategy"
	"golang.org/x/net/websocket"
)

// const token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyIjoiZnJhbmNvOTRjYXN0cm9AaG90bWFpbC5jb20ifQ.MwGdE1LICawwRX8Cp0G71q1isYYeYiGYaqhiBjdaunI"
// const origin = "wss://4yyity02md.execute-api.us-east-1.amazonaws.com"
// const url = "wss://4yyity02md.execute-api.us-east-1.amazonaws.com/ws?token=" + token

func main() {

	var token string
	// Set token with "bot_token" argument key and "default_value" as default value
	flag.StringVar(&token, "bot_token", "default_value", "Bot token to use")
	flag.Parse()
	// Check token
	if token == "default_value" {
		log.Fatal("Bot token is required, try: \n > app -bot_token <your_bot_token_here>")
	}
	// token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyIjoiRnJhbmNvMiJ9.m8RWDdgXuuknzfw9ewt_PdzD9VskMw7861iox16NNYI"
	origin := "wss://4yyity02md.execute-api.us-east-1.amazonaws.com"
	url := "wss://4yyity02md.execute-api.us-east-1.amazonaws.com/ws?token=" + token

	var ws, _ = websocket.Dial(url, "", origin)

	event := json_models.Event{}

	// Receive messages
	log.Println("  !!! Start !!!  ")
	for {

		if err := websocket.JSON.Receive(ws, &event); err != nil {
			log.Println("Error :" + err.Error())
			// Create new connection
			ws, _ = websocket.Dial(url, "", origin)
			continue
		}

		if event.Event == "your_turn" {
			go responseMyTurn(ws, event)
			continue
		}

		if event.Event == "challenge" {
			// if event.Data.Opponent == "Franco2" {
				messageAcceptChallenge := getMessageAcceptChallenge(event.Data.Challenge_id)
				websocket.JSON.Send(ws, &messageAcceptChallenge)
			// }
			continue

		}
		// if event.Event == "list_users" {
		// 	// log.Println(event)
		// 	continue
		// }
	}

}

// Return the struc for accept the challenge
func getMessageAcceptChallenge(challenge_id string) json_models.AcceptChallenge {
	acceptChallenge := json_models.AcceptChallenge{}
	acceptChallenge.Consturctor(challenge_id)
	return acceptChallenge
}

// Send JSON of response for "your_turn" event
func responseMyTurn(ws *websocket.Conn, e json_models.Event) {
	wall := json_models.PutWall{}
	move := json_models.MovePawn{}
	move, wall = strategy.GetMessageMove(&e)
	if wall.Action == "wall" {
		err := websocket.JSON.Send(ws, &wall)
		if err != nil {
			log.Println("Error :" + err.Error())
		}
	} else {
		err := websocket.JSON.Send(ws, &move)
		if err != nil {
			log.Println("Error :" + err.Error())
		}
	}
}
