package main

import (
	"fmt"
	"strconv"
)

// func main(){
// //    BasicServer()
// 	// myErrors()
// 	arr := []int{11,22,33,44,545}
// 	for idx, vals := range arr{
// 		fmt.Println(idx, vals)
// 	}
// 	// in this loop we are getting values one is for the index and other onte is the values of that array.
// }

func main(){
	num, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)

	// so every error is of type interface and it has a method Error() which format the error message 
}