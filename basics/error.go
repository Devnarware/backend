package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func myErrors(){

	n, err := readInt("Enter a number: ")

	if err != nil{
		fmt.Println("Invalid Input")
	}else{
		fmt.Println(n)
	}
}


func readInt(prompt string) (int, error){

	fmt.Print(prompt)
	sc.Scan()

	input := strings.TrimSpace(sc.Text())

	return strconv.Atoi(input)
}

// ============= INPUTS ================

// 1. .scan
// 2. .scanln
// 3. .scanf
// 4. bufio.scanner -> only this take the full line as an input