package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// convertToBin 转换成二进制
func convertToBin(n int) (result string) {
	for ; n > 0; n = n / 2 {
		remainder := n % 2
		result = strconv.Itoa(remainder) + result
	}
	return
}

// printFile 打印出文件
func printFile(filename string) {
	if file, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		printFileContents(file)
	}
}

// printFileContents 打印[]byte
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// forever 死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println("convertToBin result:")
	fmt.Println(
		convertToBin(12),
		convertToBin(13),
		convertToBin(72387885),
		convertToBin(0),
	)

	const filename = "abc.txt"
	fmt.Println()
	fmt.Println("abc.txt contents:")
	printFile(filename)

	// Uncomment to see it runs forever
	// forever()
}
