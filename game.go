package game

type Engine struct {
	Board     *Board
	CurPlayer int
	CurMove   int
	Over      bool
}

func New() Engine {
	b := NewBoard()
	return Engine{&b, -1, 0, false}
}

func (g *Engine) Move() {
	col := g.CurMove
	g.Board.Matrix[col] = append(g.Board.Matrix[col], g.CurPlayer)
}

func (g *Engine) WinChk() bool {
	return g.checkHorizontal() || g.checkVertical() || g.checkDiagonalLR() || g.checkDiagonalRL()
}

func (g *Engine) checkVertical() bool {
	col := g.Board.Matrix[g.CurMove]
	row := len(col) - 1
	if row < 3 {
		return false
	}
	for i := row; i >= row-3; i-- {
		if col[i] != g.CurPlayer {
			return false
		}
	}
	return true
}

func (g *Engine) checkHorizontal() bool {
	col := g.CurMove
	row := len(g.Board.Matrix[col]) - 1
	for i := 0; i < 4; i++ {
		sum := 0
		for j := i; j < 4+i; j++ {
			if len(g.Board.Matrix[j])-1 < row {
				break
			}
			sum += g.Board.Matrix[j][row]
		}
		if sum == (4 * g.CurPlayer) {
			return true
		}
	}
	return false
}

func (g *Engine) checkDiagonalLR() bool {
	col := g.CurMove
	row := len(g.Board.Matrix[g.CurMove]) - 1
	tot := row + col
	// Get starting point for diagonols
	x, y := 0, 0
	if tot <= 5 {
		x, y = 0, tot
	} else {
		x, y = (tot - 5), 5
	}
	for x, y := x, y; y >= 3 && x <= 3; x, y = x+1, y-1 {
		sum := 0
		i, j := x, y
		for k := 0; k < 4; k++ {
			if len(g.Board.Matrix[i])-1 < j {
				break
			}
			sum += g.Board.Matrix[i][j]
			i++
			j--
		}
		if sum == (4 * g.CurPlayer) {
			return true
		}

	}

	return false
}

func (g *Engine) checkDiagonalRL() bool {
	col := g.CurMove
	row := len(g.Board.Matrix[g.CurMove]) - 1
	dif := col - row
	// Get starting point for diagonols
	x, y := 0, 0
	if dif < 0 {
		x, y = dif, 5
	} else {
		x, y = 6, 6-dif
	}
	for x, y := x, y; y >= 3 && x >= 3; x, y = x-1, y-1 {
		sum := 0
		i, j := x, y
		for k := 0; k < 4; k++ {
			if len(g.Board.Matrix[i])-1 < j {
				break
			}
			sum += g.Board.Matrix[i][j]
			i--
			j--
		}
		if sum == (4 * g.CurPlayer) {
			return true
		}

	}

	return false
}
