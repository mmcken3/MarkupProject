package src

import (
	"golang.org/x/net/html"
	"strings"
)

var scoreMap map[string]int // Map for the different scoring possibilities.

// Scorer scores information recieved.
type Scorer interface {
	Score(<-chan html.Token) <-chan int
}

// TagScorer implements Scorer.
type TagScorer struct {
	workers int
}

// NewTagScorer returns a new instance of TagScorer.
func NewTagScorer(numWorkers int) Scorer {
	return &TagScorer{
		workers: numWorkers,
	}
}

// Score scores tokens recieved on input channel and sends the scores on the output channel.
func (t TagScorer) Score(tokens <-chan html.Token) <-chan int {
	scores := make(chan int) // Creates the output channel.
	initializeMap()          // Initializes the scoring map.

	// Kicks off a routine so that the output channel can be read while still working.
	go func() {
		quit := make(chan bool)

		// Creates a worker pool of a passed worker ammount.
		for i := 0; i < t.workers; i++ {
			go worker(tokens, scores, quit)
		}
		for j := 0; j < t.workers; j++ {
			<-quit
		}
		close(scores)
	}()
	return scores
}

// worker is the method for scoring tokens from the input channel. This will be used
// to create a worker pool.
func worker(tokens <-chan html.Token, scores chan int, quit chan bool) {
	// This waits for tokens and then calls method to score them.
	for t := range tokens {
		s := scoreToken(t.String())
		scores <- s
	}
	quit <- true
}

// scoreToken scores each token passed in based on the scoring map.
func scoreToken(token string) int {
	splitToken := strings.Split(token, " ")

	// This scores a token based on whats in the map.
	for t, s := range scoreMap {
		if strings.Contains(strings.ToLower(splitToken[0]), t) {
			return s
		}
	}
	return 0 // Returns 0 if no score type is found for this token.
}

// initializeMap initializes scoring possibilities based on requirements.
func initializeMap() {
	scoreMap = make(map[string]int)
	scoreMap["div"] = 3
	scoreMap["p"] = 1
	scoreMap["h1"] = 3
	scoreMap["h2"] = 2
	scoreMap["html"] = 5
	scoreMap["body"] = 5
	scoreMap["header"] = 10
	scoreMap["footer"] = 10
	scoreMap["font"] = -1
	scoreMap["center"] = -2
	scoreMap["big"] = -2
	scoreMap["strike"] = -1
	scoreMap["tt"] = -2
	scoreMap["frameset"] = -5
	scoreMap["frame"] = -5
}
