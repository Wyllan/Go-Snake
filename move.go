package main

import (
	"fmt"
	"math/rand"
)

func move(r MoveReq) string {
	obstacles := occupiedSpaces(r) //doesnt include tails
	// make an array of tails and explicitly check for them in areas?
	possibleMoves := checkWalls(r, obstacles)
	move := possibleMoves[rand.Intn(len(possibleMoves))]
	fmt.Println(r.Turn, possibleMoves, obstacles)
	return move
}

func checkWalls(r MoveReq, obstacles []point) []string {
	// check walls
	moves := []string{"up", "down", "left", "right"}
	head := r.You.Body[0]
	if head.X == r.Board.Width-1 {
		moves = remove(moves, "right")
	}
	if head.Y == r.Board.Height-1 {
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

func minmax() {
	// checks if will die or can kill in future
}

func toDir() {
	// resolve coords to a direction
	// params: from coord, to coord
	// returns: dir string
}

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

func occupiedSpaces(r MoveReq) []point {
	filled := []point{}
	for _, s := range r.Board.Snakes {
		filled = append(filled, s.Body...)
		filled = filled[:len(filled)-1] // make tail option
	}
	return filled
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func inArray(s []point, r point) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}
