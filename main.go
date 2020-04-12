
package main
import (
	"encoding/json"
	"fmt"
	// "log"
	// "math/rand"
	"net/http"
	"os"
	// "io/ioutil"
)

type StartReq struct {
	Game  game  `json:"game"`
	Turn  int   `json:"turn"`
	Board board `json:"board"`
	You   snake `json:"you"`
}

type StartResp struct {
	Color string `json:"color"`
	HeadType string `json:"headType"`
	TailType string `json:"tailType"`
}

type MoveReq struct {
	Game  game  `json:"game"`
	Turn  int   `json:"turn"`
	Board board `json:"board"`
	You   snake `json:"you"`
}

type MoveResp struct {
  Move string `json:"move"`
  Shout string `json:"shout,omitempty"`
}

type game struct {
	ID string `json:"id"`
}

type board struct {
  Height int `json:"height"`
  Width int `json:"width"`
  Food food `json:"food"`
}

type snake struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Health int `json:"health"`
  Body body `json:"body"`
}

type body struct {
  Body []point `json:"body"`
}

type food struct {
  Food []point
}

type point struct {
  X int `json:"x"`
  Y int `json:"y"`
}

func handler(rw http.ResponseWriter, r *http.Request) {
  switch path := r.URL.Path; path {
    case "/start":
      fmt.Println("sending start req")
      b, err := json.Marshal(StartResp{"#ff00ff", "bendr", "bolt"})
      if err != nil {
        panic(err)
      }
      rw.Header().Set("Content-Type", "application/json")
      rw.WriteHeader(http.StatusCreated)
      rw.Write(b)
      // json.NewEncoder(rw).Encode(response)
    case "/move":
      request := MoveReq{}
	    json.NewDecoder(r.Body).Decode(&request)

      // Choose a random direction to move in

      response := MoveResp{
        Move: move(),
      }
  
      rw.Header().Set("Content-Type", "application/json")
      json.NewEncoder(rw).Encode(response)
    case "/end":
      fmt.Println("game has ended")
    case "/ping":
      rw.Write([]byte("ping"))
      fmt.Println("pong")
    }
}

func main() {
  port := os.Getenv("PORT")
  if len(port) == 0 {
    port = "8080"
  }

  http.HandleFunc("/", handler)

  fmt.Printf("Starting Server at http://127.0.0.1:%s...\n", port)

  if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
    panic(err)
  }

	// fmt.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
