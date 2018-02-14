// Package raindrops converts integer into raindrop string
package raindrops

import (
	"math"
	"strconv"
)

type Raindrop struct {
	num float64
	val string
}

var (
	r3        = Raindrop{3.0, "Pling"}
	r5        = Raindrop{5.0, "Plang"}
	r7        = Raindrop{7.0, "Plong"}
	raindrops = []Raindrop{r3, r5, r7}
)

// Convert converts integer into raindrop string
func Convert(i int) string {
	str, found := "", false

	for _, elem := range raindrops {
		if math.Mod(float64(i), elem.num) == 0 {
			str += elem.val
			found = true
		}
	}

	if found {
		return str
	}
	return strconv.Itoa(i)
}
