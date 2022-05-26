package models

// Event
// Struct for revice json events
type Event struct {
	Event string    `json:"event"`
	Data  DataEvent `json:"data"`
}

type DataEvent struct {
	Users           []string `json:"users"`
	Challenge_id    string   `json:"challenge_id"`
	Opponent        string   `json:"opponent"`
	Player_1        string   `json:"player_1"`
	Player_2        string   `json:"player_2"`
	Walls           float32  `json:"walls"`
	Score_2         float32  `json:"score_2"`
	Score_1         float32  `json:"score_1"`
	Side            string   `json:"side"`
	Remaining_moves float32  `json:"remaining_moves"`
	Board           string   `json:"board"`
	Turn_token      string   `json:"turn_token"`
	Game_id         string   `json:"game_id"`
}

// Move Pawn
//This struct representing json response for a pawn move
type MovePawn struct {
	Action string       `json:"action"`
	Data   DataMovePawn `json:"data"`
}

func (m *MovePawn) Consturctor(game_id, turn_token string, from_row, from_col, to_row, to_col int) {
	m.Action = "move"
	m.Data.Game_id = game_id
	m.Data.Turn_token = turn_token
	m.Data.From_row = from_row
	m.Data.From_col = from_col
	m.Data.To_row = to_row
	m.Data.To_col = to_col

}

type DataMovePawn struct {
	Turn_token string `json:"turn_token"`
	Game_id    string `json:"game_id"`
	From_row   int    `json:"from_row"`
	From_col   int    `json:"from_col"`
	To_row     int    `json:"to_row"`
	To_col     int    `json:"to_col"`
}

// Put Wall
//This struct representing json response for put a wall
type PutWall struct {
	Action string      `json:"action"`
	Data   DataPutWall `json:"data"`
}

func (p *PutWall) Consturctor(game_id, turn_token, orientation string, row, col int) {
	p.Action = "wall"
	p.Data.Game_id = game_id
	p.Data.Turn_token = turn_token
	p.Data.Row = row
	p.Data.Col = col
	p.Data.Orientation = orientation
}

type DataPutWall struct {
	Game_id     string `json:"game_id"`
	Turn_token  string `json:"turn_token"`
	Row         int    `json:"row"`
	Col         int    `json:"col"`
	Orientation string `json:"orientation"`
}

// Accept Challenge
//This struct representing json response for put a wall
type AcceptChallenge struct {
	Action string              `json:"action"`
	Data   DataAcceptChallenge `json:"data"`
}

type DataAcceptChallenge struct {
	Challenge_id string `json:"challenge_id"`
}

func (a *AcceptChallenge) Consturctor(challenge_id string) {
	a.Action = "accept_challenge"
	a.Data.Challenge_id = challenge_id

}
