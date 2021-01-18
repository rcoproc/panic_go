package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(fileName string) *os.File {
	f, err := os.Open(fileName)

	defer func() {
		if error := recover(); error != nil {
			fmt.Println("defer", error)
      panic(error)
		} else {
			fmt.Println("defer no erro")
    }
	}()

	check(err)
	return f
}

func main() {

	f := readFile("test1.txt")
	defer f.Close()
	fmt.Printf("end of file open\n")

	b1 := make([]byte, 30)
	n1, err := f.Read(b1)
	check(err)

  fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
}
