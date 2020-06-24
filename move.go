package main

import (
	"math"
)

type Tree struct {
	Up    *Tree
	Down  *Tree
	Left  *Tree
	Right *Tree
	Value point
}

func move(r MoveReq) string {
	// make an array of tails and explicitly check for them in areas?
	// fmt.Println(r)
	// moves := []string{"up", "down", "left", "right"}
	b := boardData{r.You, r.Board}
	// bodies, tails, heads := occupiedSpaces(b)
	// possibleMoves := checkMoves(b, bodies, b.You)
	// move := possibleMoves[rand.Intn(len(possibleMoves))]
	// fmt.Println(r.Turn, possibleMoves, tails, b.You.Health)
	// updateSnake(move, &b.You, &b.Board.Food)

	switch move := minimax(b, 7, true, 0, 0, true); move {
	case 1:
		return "up"
	case 2:
		return "down"
	case 3:
		return "left"
	case 4:
		return "right"
	}
	// fmt.Println(r.Turn, possibleMoves, b.You.Health, heads)
	// return move
}

func minimax(b boardData, maxDepth int, isMain bool, alpha int, beta int, first bool) float64 { //, possible_moves []string)
	// checks if will die or can kill in future

	possibleMoves := []string{"up", "down", "left", "right"}
	currMove := "none"
	// filter moves
	// bodies, tails, heads
	bodies, _, _ := occupiedSpaces(b)
	possibleMoves = checkMoves(b, bodies, b.You)

	if maxDepth == 0 || inArray(bodies, b.You.Body[0]) ||
		len(possibleMoves) == 0 || b.You.Health == 0 {
		return evalState(b, isMain)
	}

	if isMain {
		value := -math.Inf(-1)
		for _, d := range possibleMoves {
			updateSnake(d, &b.You, &b.Board.Food) // apply state
			eval := minimax(b, maxDepth-1, false, alpha, beta, false)
			if beta <= alpha {
				break
			}
			value = math.Max(value, eval)
			if first {
				if value < eval {
					currMove = d
				}
			}
		}
		if first {
			switch currMove {
			case "up":
				return 1
			case "down":
				return 2
			case "left":
				return 3
			case "right":
				return 4
			}
		}
		return value
	}

	value := math.Inf(1)
	for _, d := range enemyMoves(b, len(b.Board.Snakes)) {
		eval := minimax(d, maxDepth-1, true, alpha, beta, false)
		if beta <= alpha {
			break
		}
		value = math.Min(value, eval)
	}
	return value
}

func evalState(b boardData, main bool) float64 {
	val := 0.0
	head := b.You.Body[0]
	bodies, _, heads := occupiedSpaces(b)
	if main {
		if inArray(bodies, head) {
			val = math.Inf(-1)
		}
		for _, h := range heads {
			if inArray(bodies, h) {
				val++
			}
		}

		return val
	}
	if inArray(bodies, head) {
		val--
	}

	// other snakes die
	// possible move
	// food ideally
	// nom snakes x<y
	// in a space surrounded by obstacles

	// min moves

	return val
}

// could be named enemyStates
func enemyMoves(b boardData, numSnakes int) []boardData {
	// t := Tree{nil, nil, nil, nil, point{-1, -1}}
	moves := []string{"up", "down", "left", "right"}

	// value will be the coords of the prev snake given a dir
	if b.You.ID == b.Board.Snakes[numSnakes].ID {
		numSnakes--
	}

	// for a snake, do all its moves
	for _, d := range moves {
		// show the state of the board if that move
		// this var seems fucky
		t := b
		updateSnake(d, &t.Board.Snakes[numSnakes], &t.Board.Food)
		if numSnakes == 0 { // if last snake we dunnit
			return []boardData{b}
		}
		// now check what other snakes can do and combine
		return append(enemyMoves(b, numSnakes-1), b)
	}

	/*
		//can eventually merge, but for now how to keep track of board states?
		for _, m in range moves {
			snake do move
			update board
			if not last snake
				return append(move, enemyMoves())
			else
				return board
		}
	*/
	return []boardData{b}
}

// make a powery thing for the snakes by recursion

func updateSnake(dir string, s *snake, f *[]point) {
	// for minimax -
	// add move point to obstacles, remove tails, minus health
	dest := []point{toPoint(dir, s.Body[0])}
	s.Body = append(dest, s.Body[:len(s.Body)-1]...)

	if inArray(*f, dest[0]) {
		*f = removePoint(*f, dest[0])
	} else {
		s.Health--
	}
	// return filled, tails, heads
	// somehow compute the move.
	// can use toPoint to add
}

func toDir() {
	// resolve coords to a direction
	// params: from coord, to coord
	// returns: dir string
}

func checkMoves(b boardData, obstacles []point, s snake) []string {
	// check walls
	moves := []string{"up", "down", "left", "right"}
	head := s.Body[0]
	if head.X == b.Board.Width-1 {
		moves = remove(moves, "right")
	}
	if head.Y == b.Board.Height-1 {
		moves = remove(moves, "down")
	}
	if head.X == 0 {
		moves = remove(moves, "left")
	}
	if head.Y == 0 {
		moves = remove(moves, "up")
	}
	// check for body
	badMoves := []string{}
	for _, d := range moves {
		if inArray(obstacles, toPoint(d, head)) {
			badMoves = append(badMoves, d)
		}
	}
	for _, d := range badMoves {
		moves = remove(moves, d)
	}
	return moves
}

// resolves a direction to the point it will give
func toPoint(dir string, head point) point {
	if dir == "up" {
		return point{head.X, head.Y - 1}
	}
	if dir == "down" {
		return point{head.X, head.Y + 1}
	}
	if dir == "left" {
		return point{head.X - 1, head.Y}
	}
	if dir == "right" {
		return point{head.X + 1, head.Y}
	}
	return point{-1, -1}
}

// returns the locations of bodies, heads, and tails
func occupiedSpaces(b boardData) ([]point, []point, []point) {
	filled := []point{}
	tails := []point{}
	heads := []point{}
	for _, s := range b.Board.Snakes {
		filled = append(filled, s.Body...)
		tails = filled[len(filled)-1:]
		heads = filled[0:1]
		filled = filled[:len(filled)-1]
	}
	return filled, tails, heads
}

// removes a string from an array
func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func removePoint(s []point, r point) []point {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// checks if a point is in a point slice (good for bodies)
func inArray(s []point, r point) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}
