package main

import (
	"bufio"
	game "conn4"
	"fmt"
	"os"
	"strconv"
)

func main() {
	g := game.New()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nBegin Connect 4\nType Column # to Insert Token \n\n")
	g.PrintBoard()
	for !g.Over && scanner.Scan() {
		fmt.Printf("Player %s's Turn\n", g.Players[g.CurPlayer])
		g.CurPlayer *= -1
		col, err := parseInput(scanner)
		if err != nil {
			fmt.Println("Invalid Input:", err)
			continue
		}
		if col > 6 || len(g.Matrix[col]) == 6 {
			fmt.Print("\nInvalid Column, Pleases select another \n")
			continue
		}
		g.CurMove = col
		g.Move()
		g.PrintBoard()
		g.Over = g.WinChk()
	}
	fmt.Printf("\nPlayer %s Wins!!\n\n", g.Players[g.CurPlayer])
}

func parseInput(scanner *bufio.Scanner) (col int, err error) {
	input := scanner.Text()
	col, err = strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return col, err
}
