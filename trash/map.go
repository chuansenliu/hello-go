package main

import (
	"fmt"
	//"os"
)

func dump(s map[int]string) {

	fmt.Printf("len = %d\n", len(s))

	for in, v := range s {
		fmt.Printf("s[%d] = %v\n", in, v)
	}
}

func setTest(s map[int] string){

	s[1] = "s1new"
}


func main() {

	//	dict := map[int]string {}
	dict := map[int]string{1: "hello", 2: "world"}

	dump(dict)

	fmt.Printf("map[1] = %s\n", dict[1])

	v2, v2exists := dict[2]

	v3, v3exists := dict[3]

	fmt.Println(v2, v2exists)
	fmt.Println(v3, v3exists)

	s := map[int] string {}

	//copy(s, dict)

	setTest(dict)
	dump(s)
	dump(dict)
}
