package main

import "fmt"

func sameFreq(a int, b int) bool {
	objA := make(map[byte]int)
	objB := make(map[byte]int)

	as := []byte(fmt.Sprint(a))
	bs := []byte(fmt.Sprint(b))

	for _, v := range as {
		objA[v]++
	}

	for _, v := range bs {
		objB[v]++
	}

	// fmt.Println(objA, objB)

	for k, v := range objA {
		// fmt.Println(v, objA[k], objB[k])
		if v != objB[k] {
			return false
		}
	}

	return true
}

func main() {
	a := 1122334
	b := 123123
	fmt.Println(sameFreq(a, b))
	// c := string(a)
	// d := fmt.Sprintf(string(b))
	// fmt.Println("%v", "%T", c, c)
	// fmt.Println("%v", "%T", d, d)
}
