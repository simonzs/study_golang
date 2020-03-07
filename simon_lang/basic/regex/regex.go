package main

import (
	"fmt"
	"regexp"
)

const text = `
my email is ccmouse@gmail.com@abc.com
email1 is abc@def.org
email2 is    kkk@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	// re, err := regexp.Compile("ccmouse@gmail.com@abc.com")
	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(.[a-zA-Z0-9.]+)`)
	matchs := re.FindAllStringSubmatch(text, -1)
	// match := re.FindString(text)
	for _, m := range matchs {
		fmt.Println(m)
	}
}
