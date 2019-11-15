

package main

import (
	"fmt"
	"os"
	"strconv"
)





func dumpSlice(s [] string){

	fmt.Printf("len = %d\n", len(s))

	for in,v := range(s){
		fmt.Printf("s[%d] = %v\n", in, v)
	}
}

func dumpSliceP(s *[] string){

	fmt.Printf("len = %d\n", len(*s))

	for in,v := range(*s){
		fmt.Printf("*s[%d] = %v\n", in, v)
	}
}


func main(){

	args := len(os.Args)
	fmt.Printf("Input %d args\n", args)
	for in, v := range(os.Args) {
		fmt.Printf("args[%d] = %s\n", in, v)
	}

	fmt.Printf("os.Getuid() = %v\n", os.Getuid())
	dir,_ := os.Getwd()
	fmt.Printf("os.Getwd() = %v\n", dir)


	s := make([]string, 10, 20)

	for i := 0; i < 10; i ++ {
		s[i] = "v:" + strconv.Itoa(i)
	}

	dumpSlice(s)


	c := make([]string, len(s), cap(s))
	copy(c, s)

	s1 := s[2:4:4]

	dumpSlice(s1)
	//s1 = append(s1, "append")

	s1[1] = "modify"

	s3 := c[:]

	s3[0] = "s3"

	dumpSlice(s3)
	dumpSlice(s1)
	dumpSlice(s)
	dumpSlice(c)

	//s = append(s, "nihao")

	//dumpSlice(n)
	//dumpSlice(s)

	s2 := c[0:len(c) - 1]

	dumpSliceP(&s2)
}