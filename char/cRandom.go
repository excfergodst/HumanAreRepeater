package char

import "math/rand"

type Random struct{}

func (c Random) Run(s Status) int64 {

	return rand.Int63() % 2
}

func (c Random) Name() string {

	return "Random"
}
