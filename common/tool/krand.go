package tool

import (
	"math/rand"
	"time"
)

const (
	KC_RAND_KIND_NUM   = 0 // only numbers
	KC_RAND_KIND_LOWER = 1 // lowercase characters
	KC_RAND_KIND_UPPER = 2 // uppercase characters
	KC_RAND_KIND_ALL   = 3 // numbers and characters
)

/* Generate a random string 
{10, 48}: Number, the base is 48 ('0' in ASCII code), the range is 10 (0-9).
{26, 97}: Lowercase letters, the base is 97 ('a' in ASCII code), the range is 26 (a-z).
{26, 65}: Uppercase letters, the base is 65 ('A' in ASCII code), the range is 26 (A-Z).
*/
func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		if isAll {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}