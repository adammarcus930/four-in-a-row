package game

import "fmt"

type Board struct {
	Matrix  [7][]int
	Players map[int]string
}

func NewBoard() Board {
	matrix := [7][]int{}
	return Board{matrix, map[int]string{1: "R", -1: "Y"}}
}

func (b *Board) PrintBoard() {
	for i := 5; i >= 0; i-- {
		rowStr := ""
		div := " "
		for _, col := range b.Matrix {
			if len(col)-1 < i {
				rowStr += div + " "
			} else {
				rowStr += div + b.Players[col[i]]
			}
			div = " | "
		}
		fmt.Println(rowStr)
		fmt.Println("---+---+---+---+---+---+---")
	}
	fmt.Print(" 0   1   2   3   4   5   6\n\n")
}
