package models

import (
	"testing"
)

const uotOfBoard = "Pawn tries move out of the boar; "

var board_1 = [17][17]string{
	//0   0     1    1    2    2   3    3     4    4    5   5    6    6    7    7    8
	//0    1    2    3    4    5    6    7    8    9   10   11   12   13   14   15   16
	{"N", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "N"}, //0    0
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //1    0
	{" ", " ", "N", " ", " ", " ", "N", " ", "N", "|", "N", " ", " ", " ", " ", " ", " "}, //2    1
	{" ", " ", " ", " ", "-", "*", "-", " ", " ", "*", "-", "*", "-", " ", " ", " ", " "}, //3    1
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", "|", " ", " ", " ", " ", " ", " ", " "}, //4    2
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //5    2
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //6    3
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //7    3
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //8    4
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //9    4
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //10   5
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //11   5
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", "|", " ", " ", " ", " ", " ", " ", " "}, //12   6
	{" ", " ", " ", " ", "-", "*", "-", " ", " ", "*", "-", "*", "-", " ", " ", " ", " "}, //13   6
	{" ", " ", "S", " ", " ", " ", "S", " ", "S", "|", "S", " ", " ", " ", " ", " ", " "}, //14   7
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //15   7
	{"S", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "S"}} //16   8

func TestPackageFunctions(t *testing.T) {
	// Create pawns
	Npawn := Pawn{}
	Spawn := Pawn{}
	// Set values for pawns
	Npawn.Constructor(0, 0, "N", "S", &board_1)
	Spawn.Constructor(0, 0, "S", "N", &board_1)

	// parsePosition function tests
	for i := 0; i < 17; i++ {
		if parsePosition(i) != int(i/2) {
			t.Error("Error to parse position")
		}

	}

	if parsePosition(9) != 4 {
		t.Error("Error to parse position")
	}
	if parsePosition(8) != 4 {
		t.Error("Error to parse position")
	}
	if parsePosition(0) != 0 {
		t.Error("Error to parse position")
	}
	if parsePosition(1) != 0 {
		t.Error("Error to parse position")
	}
	if parsePosition(5) != 2 {
		t.Error("Error to parse position")
	}
	if parsePosition(7) != 3 {
		t.Error("Error to parse position")
	}
	if parsePosition(15) != 7 {
		t.Error("Error to parse position")
	}
	if parsePosition(16) != 8 {
		t.Error("Error to parse position")
	}

	// isInBoard function tests
	for i := 0; i < 17; i++ {
		if !isInBoard(i, i) {
			t.Error("Error position is out of board_1")
		}
		for j := 0; j < 17; j++ {
			if !isInBoard(i, j) {
				t.Error("Error position is out of board_1")
			}
		}
	}

	if isInBoard(17, 0) || isInBoard(0, 17) {
		t.Error("Error position is out of board_1")
	}

	if isInBoard(-1, 0) || isInBoard(0, -1) {
		t.Error("Error position is out of board_1")
	}

}

func TestPawnMoves(t *testing.T) {
	// Create pawns
	Npawn := Pawn{}
	Spawn := Pawn{}
	// Set values for N pawn
	Npawn.Constructor(0, 0, "N", "S", &board_1)
	NpawnPosition := "Npawn row = 0 and col = 0; "
	if Npawn.canMoveLeft() > 0 {
		t.Error(uotOfBoard + NpawnPosition + "pawn tries move to left")
	}
	if Npawn.canMoveRight() != 1 {
		t.Error(NpawnPosition + "pawn can not move right")
	}
	if Npawn.canMoveNorth() > 0 {
		t.Error(uotOfBoard + NpawnPosition + "pawn tries move to north")
	}
	if Npawn.canMoveSouth() != 1 {
		t.Error(NpawnPosition + "pawn can not move south")
	}

	// Set values for S pawn
	Spawn.Constructor(16, 0, "S", "N", &board_1)
	SpawnPosition := "Spawn row = 0 and col = 0; "
	if Spawn.canMoveLeft() > 0 {
		t.Error(uotOfBoard + SpawnPosition + "pawn tries move to left")
	}
	if Spawn.canMoveRight() != 1 {
		t.Error(SpawnPosition + "pawn can not move right")
	}
	if Spawn.canMoveNorth() != 1 {
		t.Error(SpawnPosition + "pawn can not move north")
	}
	if Spawn.canMoveSouth() > 0 {
		t.Error(uotOfBoard + SpawnPosition + "pawn tries move to south")
	}

	// Set values for N pawn
	Npawn.Row = 2
	Npawn.Col = 2
	NpawnPosition = "Npawn row = 2 and col = 2; "
	if Npawn.canMoveLeft() != 1 {
		t.Error(NpawnPosition + "pawn can not move left")
	}
	if Npawn.canMoveRight() != 1 {
		t.Error(NpawnPosition + "pawn can not move right")
	}
	if Npawn.canMoveNorth() != 1 {
		t.Error(NpawnPosition + "pawn can not move north")
	}
	if Npawn.canMoveSouth() != 1 {
		t.Error(NpawnPosition + "pawn can not move south")
	}

	// Set values for S pawn
	Spawn.Row = 14
	Spawn.Col = 2
	SpawnPosition = "Spawn row = 14 and col = 2; "
	if Spawn.canMoveLeft() != 1 {
		t.Error(SpawnPosition + "pawn can not move left")
	}
	if Spawn.canMoveRight() != 1 {
		t.Error(SpawnPosition + "pawn can not move right")
	}
	if Spawn.canMoveNorth() != 1 {
		t.Error(SpawnPosition + "pawn can not move north")
	}
	if Spawn.canMoveSouth() != 1 {
		t.Error(SpawnPosition + "pawn can not move south")
	}

	// Set values for N pawn
	Npawn.Row = 2
	Npawn.Col = 6
	NpawnPosition = "Npawn row = 2 and col = 6; "
	if Npawn.canMoveLeft() != 1 {
		t.Error(NpawnPosition + "pawn can not move left")
	}
	if Npawn.canMoveRight() > 0 {
		t.Error(NpawnPosition + "pawn tries move to right, there is other pawnn")
	}
	if Npawn.canMoveNorth() != 1 {
		t.Error(NpawnPosition + "pawn can not move north")
	}
	if Npawn.canMoveSouth() > 0 {
		t.Error(NpawnPosition + "pawn tries move to south, there is a wall")
	}

	// Set values for S pawn
	Spawn.Row = 14
	Spawn.Col = 6
	SpawnPosition = "Spawn row = 14 and col = 6; "
	if Spawn.canMoveLeft() != 1 {
		t.Error(SpawnPosition + "pawn can not move left")
	}
	if Spawn.canMoveRight() > 0 {
		t.Error(SpawnPosition + "pawn tries move to right, there is other pawnn")
	}
	if Spawn.canMoveNorth() > 0 {
		t.Error(SpawnPosition + "pawn tries move to north, there is a wall")
	}
	if Spawn.canMoveSouth() != 1 {
		t.Error(SpawnPosition + "pawn can not move south")
	}

	// Set values for N pawn
	Npawn.Row = 2
	Npawn.Col = 8
	NpawnPosition = "Npawn row = 2 and col = 8; "
	if Npawn.canMoveLeft() > 0 {
		t.Error(NpawnPosition + "pawn tries move to left, there is other pawn")
	}
	if Npawn.canMoveRight() > 0 {
		t.Error(NpawnPosition + "pawn tries move to right, there is a wall")
	}
	if Npawn.canMoveNorth() != 1 {
		t.Error(NpawnPosition + "pawn can not move north")
	}
	if Npawn.canMoveSouth() != 1 {
		t.Error(NpawnPosition + "pawn can not move south")
	}

	// Set values for S pawn
	Spawn.Row = 14
	Spawn.Col = 8
	SpawnPosition = "Spawn row = 14 and col = 8; "
	if Spawn.canMoveLeft() > 0 {
		t.Error(SpawnPosition + "can not move left, there is other pawn")
	}
	if Spawn.canMoveRight() > 0 {
		t.Error(SpawnPosition + "pawn tries move to right, there is a wall")
	}
	if Spawn.canMoveNorth() != 1 {
		t.Error(SpawnPosition + "pawn can not move north")
	}
	if Spawn.canMoveSouth() != 1 {
		t.Error(SpawnPosition + "pawn  can not move south")
	}

	// Set values for N pawn
	Npawn.Row = 2
	Npawn.Col = 10
	NpawnPosition = "Npawn row = 2 and col = 10; "
	if Npawn.canMoveLeft() > 0 {
		t.Error(NpawnPosition + "pawn tries move to left, there is a wall")
	}
	if Npawn.canMoveRight() != 1 {
		t.Error(NpawnPosition + "pawn can not move right")
	}
	if Npawn.canMoveNorth() != 1 {
		t.Error(NpawnPosition + "pawn can not move north")
	}
	if Npawn.canMoveSouth() > 0 {
		t.Error(NpawnPosition + "pawn tries move to south, there is a wall")
	}

	// Set values for S pawn
	Spawn.Row = 14
	Spawn.Col = 10
	SpawnPosition = "Spawn row = 14 and col = 10; "
	if Spawn.canMoveLeft() > 0 {
		t.Error(SpawnPosition + "pawn tries move to left, there is a wall")
	}
	if Spawn.canMoveRight() != 1 {
		t.Error(SpawnPosition + "pawn can not move right")
	}
	if Spawn.canMoveNorth() > 0 {
		t.Error(SpawnPosition + "pawn tries move to north, there is a wall")
	}
	if Spawn.canMoveSouth() != 1 {
		t.Error(SpawnPosition + "pawn can not move south")
	}

	// Set values for N pawn
	Npawn.Row = 0
	Npawn.Col = 16
	NpawnPosition = "Npawn row = 0 and col = 16; "
	if Npawn.canMoveLeft() != 1 {
		t.Error(NpawnPosition + "pawn can not move to left")
	}
	if Npawn.canMoveRight() > 0 {
		t.Error(uotOfBoard + NpawnPosition + "pawn tries move to right")
	}
	if Npawn.canMoveNorth() > 0 {
		t.Error(uotOfBoard + NpawnPosition + "pawn tries move to north")
	}
	if Npawn.canMoveSouth() != 1 {
		t.Error(NpawnPosition + "pawn can not move to south")
	}

	// Set values for S pawn
	Spawn.Row = 16
	Spawn.Col = 16
	SpawnPosition = "Spawn row = 16 and col = 16; "
	if Spawn.canMoveLeft() != 1 {
		t.Error(SpawnPosition + "pawn can not move to left")
	}
	if Spawn.canMoveRight() > 0 {
		t.Error(uotOfBoard + SpawnPosition + "pawn tries move to right")
	}
	if Spawn.canMoveNorth() != 1 {
		t.Error(SpawnPosition + "pawn can not move to north")
	}
	if Spawn.canMoveSouth() > 0 {
		t.Error(uotOfBoard + SpawnPosition + "pawn tries move to south")
	}

}

var board_2 = [17][17]string{
	//0   0     1    1    2    2   3    3     4    4    5   5    6    6    7    7    8
	//0    1    2    3    4    5    6    7    8    9   10   11   12   13   14   15   16
	{"N", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "N"}, //0    0
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //1    0
	{"S", " ", "N", " ", "N", " ", " ", " ", "N", "|", "N", " ", " ", " ", " ", " ", " "}, //2    1
	{" ", " ", " ", " ", "-", "*", "-", " ", " ", "*", "-", "*", "-", " ", " ", " ", " "}, //3    1
	{" ", " ", "S", " ", "S", " ", "N", " ", " ", "|", " ", " ", " ", " ", " ", " ", " "}, //4    2
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //5    2
	{" ", " ", " ", " ", " ", " ", "S", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //6    3
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //7    3
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //8    4
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //9    4
	{" ", " ", " ", " ", " ", " ", "N", " ", " ", " ", " ", " ", " ", " ", "N", " ", " "}, //10   5
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //11   5
	{" ", " ", "N", " ", "N", " ", "S", " ", " ", "|", " ", " ", " ", " ", "S", " ", " "}, //12   6
	{" ", " ", " ", " ", "-", "*", "-", " ", " ", "*", "-", "*", "-", " ", " ", " ", " "}, //13   6
	{"N", " ", "S", " ", "S", " ", " ", " ", "S", "|", "S", " ", " ", " ", " ", " ", " "}, //14   7
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}, //15   7
	{"S", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "S"}} //16   8

// test for pawn jumps
func TestJumOpponent(t *testing.T) {

	Npawn := Pawn{}
	Spawn := Pawn{}
	// Set values for pawns
	Npawn.Constructor(10, 14, "N", "S", &board_2)
	Spawn.Constructor(12, 14, "S", "N", &board_2)

	NpawnPosition := "Npawn row = 10 and col = 14; "
	SpawnPosition := "Spawn row = 14 and col = 12; "

	if Npawn.canJumpOpponentToSouth() < 0 {
		t.Error(NpawnPosition + "pawn can not jump opponent")
	}
	if Spawn.canJumpOpponentToNorth() < 0 {
		t.Error(SpawnPosition + "pawn can not jump opponent")
	}

	Npawn.Row = 0
	Npawn.Col = 0
	NpawnPosition = "Npawn row = 0 and col = 0; "
	Spawn.Row = 2
	Spawn.Col = 0
	SpawnPosition = "Spawn row = 2 and col = 0; "
	if Npawn.canJumpOpponentToSouth() != 2 {
		t.Error(NpawnPosition + "pawn can not jump opponent")
	}
	if Spawn.canJumpOpponentToNorth() > 0 {
		t.Error(SpawnPosition + "pawn tries jump out the board")
	}
	Npawn.Row = 14
	Npawn.Col = 0
	NpawnPosition = "Npawn row = 14 and col = 0; "
	Spawn.Row = 16
	Spawn.Col = 0
	SpawnPosition = "Spawn row = 16 and col = 0; "
	if Npawn.canJumpOpponentToSouth() > 0 {
		t.Error(NpawnPosition + "pawn tries jump out the board")
	}
	if Spawn.canJumpOpponentToNorth() != 2 {
		t.Error(SpawnPosition + "pawn can not jump opponent")
	}
	Npawn.Row = 2
	Npawn.Col = 2
	NpawnPosition = "Npawn row = 2 and col = 2; "
	Spawn.Row = 4
	Spawn.Col = 2
	SpawnPosition = "Spawn row = 4 and col = 2; "
	if Npawn.canJumpOpponentToSouth() != 2 {
		t.Error(NpawnPosition + "pawn can not jump opponent")
	}
	if Spawn.canJumpOpponentToNorth() != 2 {
		t.Error(SpawnPosition + "pawn can not jump opponent")
	}
	Npawn.Row = 12
	Npawn.Col = 2
	NpawnPosition = "Npawn row = 12 and col = 2; "
	Spawn.Row = 14
	Spawn.Col = 2
	SpawnPosition = "Spawn row = 14 and col = 2; "
	if Npawn.canJumpOpponentToSouth() != 2 {
		t.Error(NpawnPosition + "pawn can not jump opponent")
	}
	if Spawn.canJumpOpponentToNorth() != 2 {
		t.Error(SpawnPosition + "pawn can not jump opponent")
	}
	Npawn.Row = 2
	Npawn.Col = 4
	NpawnPosition = "Npawn row = 2 and col = 4; "
	Spawn.Row = 4
	Spawn.Col = 4
	SpawnPosition = "Spawn row = 4 and col = 4; "
	if Npawn.canJumpOpponentToSouth() > 0 {
		t.Error(NpawnPosition + "pawn tries jump a wall")
	}
	if Spawn.canJumpOpponentToNorth() > 0 {
		t.Error(SpawnPosition + "pawn tries jump a wall")
	}
	Npawn.Row = 12
	Npawn.Col = 4
	NpawnPosition = "Npawn row = 12 and col = 4; "
	Spawn.Row = 14
	Spawn.Col = 4
	SpawnPosition = "Spawn row = 14 and col = 4; "
	if Npawn.canJumpOpponentToSouth() > 0 {
		t.Error(NpawnPosition + "pawn tries jump a wall")
	}
	if Spawn.canJumpOpponentToNorth() > 0 {
		t.Error(SpawnPosition + "pawn tries jump a wall")
	}
	Npawn.Row = 4
	Npawn.Col = 6
	NpawnPosition = "Npawn row = 4 and col = 6; "
	Spawn.Row = 6
	Spawn.Col = 6
	SpawnPosition = "Spawn row = 6 and col = 6; "
	if Npawn.canJumpOpponentToSouth() != 2 {
		t.Error(NpawnPosition + "pawn can not jump opponent")
	}
	if Spawn.canJumpOpponentToNorth() > 0 {
		t.Error(SpawnPosition + "pawn tries jump a wall")
	}
	Npawn.Row = 10
	Npawn.Col = 6
	NpawnPosition = "Npawn row = 10 and col = 6; "
	Spawn.Row = 12
	Spawn.Col = 6
	SpawnPosition = "Spawn row = 12 and col = 6; "
	if Npawn.canJumpOpponentToSouth() > 0 {
		t.Error(NpawnPosition + "pawn tries jump a wall")
	}
	if Spawn.canJumpOpponentToNorth() != 2 {
		t.Error(SpawnPosition + "pawn can not jump opponent")
	}

}
