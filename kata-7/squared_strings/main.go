package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcd\nefgh\nijkl\nmnop"
	// strs := strings.Split(s, "\n")
	// fmt.Printf("%v \n", strs)
	resVertMirror := VertMirror(s)
	resHorMirror := HorMirror(s)
	fmt.Printf("%v\n\n", resVertMirror)
	fmt.Printf("%v\n\n", resHorMirror)
}

func VertMirror(s string) string {
	reverseString := func(s string) string {
		n := 0
		rune := make([]rune, len(s))

		for _, r := range s {
			rune[n] = r
			n++
		}

		rune = rune[0:n]

		for i := 0; i < n/2; i++ {
			rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
		}

		res := string(rune)
		return res
	}

	var res []string
	strs := strings.Split(s, "\n")

	for _, v := range strs {
		res = append(res, reverseString(v))
	}

	return strings.Join(res, "\n")
}

// HorMirror returns strings with reverse separated by "\n"
func HorMirror(s string) string {
	var res []string
	strs := strings.Split(s, "\n")

	for i := range strs {
		index := len(strs) - i - 1
		res = append(res, strs[index])
	}

	return strings.Join(res, "\n")
}

// FParam is shit
type FParam func(string) string

// Oper was moron
func Oper(f FParam, x string) string {
	return f(x)
}
