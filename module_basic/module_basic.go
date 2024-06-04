package module_basic

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("Module basic activado ...")
	rand.Seed(time.Now().Unix())
}

// RandBytes generates a random string of length n using characters from the input string.
//
// n: the length of the random string to generate.
// input: the string containing the characters to use for generating the random string.
// returns: the randomly generated string.
func RandBytes(n int, input string) string {
	b := make([]byte, n)
	num_bytes := len(input)
	for i := 0; i < n; i++ {
		b[i] = input[rand.Intn(num_bytes)]
	}
	return string(b)
}

// RandRunes generates a random string of n runes using characters from the input string.
//
// n: the length of the random string to be generated.
// input: the string containing the characters to be used in the random string.
// Returns: the randomly generated string.
func RandRunes(n int, input string) string {
	b := make([]rune, n)
	r := []rune(input)
	num_runes := len(r)
	for i := 0; i < n; i++ {
		b[i] = r[rand.Intn(num_runes)]
	}
	return string(b)
}

func CalculatePi(iterations int) float32 {

	inside := 0
	radius := 0.5

	for i := 0; i < iterations; i++ {
		x := rand.Float64() * 2 * radius
		y := rand.Float64() * 2 * radius
		if x*x+y*y < radius*radius {
			inside++
		}
	}
	return 16 * float32(inside) / float32(iterations)
}

// type SetAcv struct {
// 	storage []int
// }

// func NewSetAcv() SetAcv {
// 	return SetAcv{}
// }

// func (s SetAcv) Print() string {
// 	return fmt.Sprintf("%v", s.storage)
// }

// func (s *SetAcv) Añadir(val int) {
// 	if !s.Existe(val) {
// 		s.storage = append(s.storage, val)
// 	} else {
// 		fmt.Println("No inserto el valor porque ya existe", val)
// 	}
// }

// func (s *SetAcv) Borrar(val int) {
// 	storage := []int{}
// 	for i := 0; i < len(s.storage); i++ {
// 		if s.storage[i] != val {
// 			storage = append(storage, s.storage[i])
// 		}
// 	}
// 	s.storage = storage
// }

// func (s *SetAcv) Existe(val int) bool {
// 	for i := 0; i < len(s.storage); i++ {
// 		if s.storage[i] == val {
// 			return true
// 		}
// 	}
// 	return false
// }

type SetAcv struct {
	storage map[int]bool
}

func NewSetAcv() SetAcv {
	return SetAcv{
		storage: make(map[int]bool),
	}
}

func (s SetAcv) Print() string {
	return fmt.Sprintf("%v", s.storage)
}

func (s *SetAcv) Add(val int) {
	s.storage[val] = true
}

func (s *SetAcv) Remove(val int) {
	delete(s.storage, val)
}

func (s *SetAcv) Existe(val int) bool {
	_, ok := s.storage[val]
	return ok
}
