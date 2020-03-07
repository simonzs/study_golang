package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Fibonacci ...
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
}

func writeFile(filename string) {
	file, err := os.OpenFile(
		filename, os.O_EXCL|os.O_CREATE, 0666)

	err = errors.New("this is a custom error")

	if err != nil {
		if patherror, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(patherror.Op,
				patherror.Path,
				patherror.Err,
			)
		}
		fmt.Println("Error: ", err.Error())
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer()
	writeFile("abc.txt")
}
