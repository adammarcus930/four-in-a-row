package main

import (
	"bufio"
	game "conn4"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	g := game.New()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nBegin Connect 4\nType Column # to Insert Token \n\n")
	g.Board.PrintBoard()
	for !g.Over && scanner.Scan() {
		fmt.Printf("Player %s's Turn\n", g.Board.Players[g.CurPlayer])
		g.CurPlayer *= -1
		col, err := parseInput(scanner)
		if err != nil {
			fmt.Println("Invalid Input:", err)
			continue
		}
		if col > 6 || len(g.Board.Matrix[col]) == 6 {
			fmt.Print("\nInvalid Column, Pleases select another \n")
			continue
		}
		g.CurMove = col
		g.Move()
		g.Board.PrintBoard()
		g.Over = g.WinChk()
	}
	fmt.Printf("\nPlayer %s Wins!!\n\n", g.Board.Players[g.CurPlayer])
}

func parseInput(scanner *bufio.Scanner) (col int, err error) {
	input := scanner.Text()
	strs := strings.Split(input, ",")
	col, err = strconv.Atoi(strs[0])
	if err != nil {
		return 0, err
	}
	return col, err
}
