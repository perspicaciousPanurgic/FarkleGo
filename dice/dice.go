package dice

import "math/rand"
import "time"

//Roll to generate random numbers between 1-6
func Roll(dice [6]int) [6]int {

	// generate random source seed using time
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// fill each array index with random value between 1 & 6
	for i := range dice {
		dice[i] = r1.Intn(5) + 1
	}

	return dice
}
