package recall

import "fmt"

var chessBoard [8][8]int

func EightQueen() {
	putQueueToLine(0)
}

func putQueueToLine(row int) {
	//fmt.Println("row:", row)
	if row == 8 {
		fmt.Println("Found 8 queen!")
		printChessBoard()
		return
	}
	for col := 0; col < 8; col++ {
		if checkChessBoard(row, col) {
			//fmt.Printf("put %d,%d=1\n", row, col)
			setLineToOne(row, col)
			putQueueToLine(row + 1)
		}
	}
}

func printChessBoard() {
	for _, line := range chessBoard {
		for _, v := range line {
			fmt.Printf("[%d]", v)
		}
		fmt.Println()
	}
}

func setLineToOne(row, col int) {
	for i := 0; i < 8; i++ {
		if i == col {
			chessBoard[row][i] = 1
		} else {
			chessBoard[row][i] = 0
		}
	}
}

func clearChessBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			chessBoard[i][j] = 0
		}
	}
}

func checkChessBoard(row, col int) bool {
	for i := 0; i < row; i++ {
		if chessBoard[i][col] == 1 {
			return false
		}
	}
	i := row
	j := col
	for i > 0 && j > 0 {
		i--
		j--
		if chessBoard[i][j] == 1 {
			return false
		}
	}
	i = row
	j = col
	for i > 0 && j < 7 {
		i--
		j++
		if chessBoard[i][j] == 1 {
			return false
		}
	}
	return true
}
