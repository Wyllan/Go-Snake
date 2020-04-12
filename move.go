package main
import (
	// "encoding/json"
	// "fmt"
	// "log"
	"math/rand"
)

func move() string {
	possibleMoves := []string{"up", "down", "left", "right"}
	move := possibleMoves[rand.Intn(len(possibleMoves))]
	return move
}