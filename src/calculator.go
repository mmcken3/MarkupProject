package src

// Calculator calculates a total score.
type Calculator interface {
	Calc(<-chan int) int
}

// AddCalculator implements Calculator.
type AddCalculator struct {
	total int
}

// NewAddCalculator returns a new instance of AddCalculator.
func NewAddCalculator() Calculator {
	return &AddCalculator{
		total: 0,
	}
}

// Calc totals up the score from ints passed in on an input channel.
func (t AddCalculator) Calc(scores <-chan int) int {
	// Monitor channel for new ints.
	for s := range scores {
		t.total += s
	}
	return t.total // Returns the total score.
}
