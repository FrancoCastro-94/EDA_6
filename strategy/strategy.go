package strategy

import (
	"fmt"
	"strings"

	models "github.com/FrancoCastro-94/web-socket/models"
)

// Return action wall or move for response "your_turn" event
func GetMessageMove(turn *models.Event) (models.MovePawn, models.PutWall) {
	var side, opponent_side string
	if turn.Data.Side == "S" {
		opponent_side = "N"
		side = "S"
	} else {
		opponent_side = "S"
		side = "N"
	}

	board, my_positions, _ /*opponent_positions*/ := makeBoardMatrix(turn.Data.Board, side, opponent_side)

	// printBoard(board)

	move := models.MovePawn{}
	wall := models.PutWall{}

	if turn.Data.Side == "S" {
		w_row, w_col, orientation := chectAndPutWallForSouthSide(&board)
		if w_row > -1 && w_col > -1 {
			wall.Consturctor(turn.Data.Game_id, turn.Data.Turn_token, orientation, w_row, w_col)
			return move, wall
		}
	} else {
		w_row, w_col, orientation := chectAndPutWallForNorthSide(&board)
		if w_row > -1 && w_col > -1 {
			wall.Consturctor(turn.Data.Game_id, turn.Data.Turn_token, orientation, w_row, w_col)
			return move, wall
		}
	}
	my_pawns := createPawns(side, opponent_side, &board, my_positions)

	from_row, from_col, to_row, to_col := getMoveFromPawns(&my_pawns)

	// opponent_pawns := createPawns(opponent_side, side, &board, opponent_positions)

	// fmt.Println("----------board-----------")
	// printBoard(board)
	// fmt.Println("----------board-----------")

	move.Consturctor(turn.Data.Game_id, turn.Data.Turn_token, from_row, from_col, to_row, to_col)
	// fmt.Println("----------move-----------")
	// fmt.Println("Side: ", turn.Data.Side)
	// fmt.Println("Remaining moves: ", turn.Data.Remaining_moves)
	// fmt.Println("----------move-----------")
	// w_row, w_col := tryPutBackWall(my_pawns)

	return move, wall
}

// Create and return 17x17 matrix board, my pawns positiones and opponent pawns positions
func makeBoardMatrix(srt_board, side, opponent_side string) (board [17][17]string, my_positions, opponent_positions []int) {
	slice_board := strings.Split(srt_board, "")
	count := 0
	if side == "S" {
		for row := 0; row < 17; row++ {
			for col := 0; col < 17; col++ {
				switch slice_board[count] {
				case side:
					my_positions = append(my_positions, int(row))
					my_positions = append(my_positions, int(col))
				case opponent_side:
					opponent_positions = append(opponent_positions, int(row))
					opponent_positions = append(opponent_positions, int(col))
				}
				board[row][col] = slice_board[count]
				count++
			}
		}
	} else {
		count = 288
		for row := 16; row > -1; row-- {
			for col := 16; col > -1; col-- {
				switch slice_board[count] {
				case side:
					my_positions = append(my_positions, int(row))
					my_positions = append(my_positions, int(col))
				case opponent_side:
					opponent_positions = append(opponent_positions, int(row))
					opponent_positions = append(opponent_positions, int(col))
				}
				board[row][col] = slice_board[count]
				count--
			}
		}
	}
	return
}

// Pint visual board
func printBoard(board [17][17]string) {
	for i, row := range board {
		if i < 10 {
			fmt.Print(i, "  |")
		} else {
			fmt.Print(i, " |")
		}
		for j, value := range row {
			if value != " " && i%2 == 0 && j%2 == 0 {
				switch value {
				case "S":
					fmt.Print("|", value, "|")
				case "N":
					fmt.Print("|", value, "|")
				default:
					fmt.Print(" ", value, " ")
				}
			} else if value == " " && i%2 == 0 && j%2 == 0 {
				fmt.Print("| |")
			} else {
				fmt.Print(" ", value, " ")
			}
		}
		fmt.Print("|\n")
	}
}

// Return slice of pawns from received positions
func createPawns(side, opponent_side string, board *[17][17]string, positions []int) (pawns []models.Pawn) {
	for i := 0; i < len(positions); i += 2 {
		pawn := models.Pawn{}
		pawn.Constructor(positions[i], positions[i+1], side, opponent_side, board)
		pawns = append(pawns, pawn)
	}
	return
}

// Select the best move according to the priority returned by the pawns
func getMoveFromPawns(my_pawns *[]models.Pawn) (from_row, from_col, to_row, to_col int) {
	moves := [3][4]int{}
	priority := 0
	lastPriority := 0
	// index of max priority
	iMax := 0
	// Save in moves the pawns moves
	for i, p := range *my_pawns {
		moves[i][0] = int(p.Row / 2)
		moves[i][1] = int(p.Col / 2)
		moves[i][2], moves[i][3], priority = p.GetMove()
		if priority > lastPriority {
			lastPriority = priority
			iMax = i
		}
	}
	// Select the move with max priority
	from_row = moves[iMax][0]
	from_col = moves[iMax][1]
	to_row = moves[iMax][2]
	to_col = moves[iMax][3]
	return
}

// Gets the number of row or column in the board 17x17 and return number for 9x9 board
func parsePosition(position int) int {
	return int(position / 2)
}

// position of horizontal walls for N side
var n_h_wall_positions = [][]int{
	{3, 1},
	{3, 5},
	{3, 9},
	{3, 13},
}

// position of vertical walls for N side
var n_v_wall_positions = [][]int{
	{5, 15},
	{9, 15},
	{13, 15},
}

// Takes walls positions for N side and check if can put one if can then return position and orientation
func chectAndPutWallForNorthSide(board *[17][17]string) (int, int, string) {
	for _, h := range n_h_wall_positions {
		if checkIsFreeForHorizontalWall(h[0], h[1], board) {
			return parsePosition(h[0]), parsePosition(h[1]), "h"
		}
	}
	for _, v := range n_v_wall_positions {
		if checkIsFreeForVerticalWall(v[0], v[1], board) {
			return parsePosition(v[0]), parsePosition(v[1]), "v"
		}
	}
	return -1, -1, ""
}

// position of horizontal walls for S side
var s_h_wall_positions = [][]int{
	{13, 15},
	{13, 11},
	{13, 7},
	{13, 3},
}

// position of vertical walls for S side
var s_v_wall_positions = [][]int{
	{11, 1},
	{7, 1},
	{3, 1},
}

// Takes walls positions for S side and check if can put one if can then return position and orientation
func chectAndPutWallForSouthSide(board *[17][17]string) (int, int, string) {
	for _, h := range s_h_wall_positions {
		if checkIsFreeForHorizontalWall(h[0], h[1], board) {
			return parsePosition(h[0]), parsePosition(h[1]), "h"
		}
	}
	for _, v := range s_v_wall_positions {
		if checkIsFreeForVerticalWall(v[0], v[1], board) {
			return parsePosition(v[0]), parsePosition(v[1]), "v"
		}
	}
	return -1, -1, ""
}

// Check all horizontal wall positions is free
func checkIsFreeForHorizontalWall(row, col int, board *[17][17]string) bool {
	return board[row][col+1] == " " && board[row][col] == " " && board[row][col-1] == " "
}

// Check all vertical wall positions is free
func checkIsFreeForVerticalWall(row, col int, board *[17][17]string) bool {
	return board[row+1][col] == " " && board[row][col] == " " && board[row-1][col] == " "
}
