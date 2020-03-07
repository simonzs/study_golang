package main

import (
	"fmt"
	"io/ioutil"
)

// ifs if的使用
func ifs() {
	// if "abc.txt" is not found.
	// please check what current directory is,
	// and change filename accordingly.
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

// grade switch的使用
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic("score is unavailable value")
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	ifs()
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
	)
}
