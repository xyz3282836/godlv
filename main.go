package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)
var _xxb = 100

func count(i, j int) int {
	yz := 5
	a := i + j
	result := a * yz
	b := result + _xxb
	return b
}

func randHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var (
		i = rand.Intn(20)
		j = rand.Intn(20)
	)

	result := count(i, j)
	_, _ = w.Write([]byte(fmt.Sprintf("%d", result)))
}

func main() {
	http.HandleFunc("/rand", randHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("start server fail: %v", err)
	}
}
