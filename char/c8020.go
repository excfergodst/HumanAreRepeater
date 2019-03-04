package char

import "math/rand"

type EightyYesTwentyNo struct{}

func (c EightyYesTwentyNo) Run(s Status) int64 {

	v := rand.Int63() % 10

	if v > 1 {
		return 1
	}

	return 0

}

func (c EightyYesTwentyNo) Name() string {

	return "EightyYesTwentyNo"
}
