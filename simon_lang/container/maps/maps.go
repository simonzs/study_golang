package main

import "fmt"

func main() {
	m1 := map[string]string{
		"name":    "simon",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]string) // m2 == empty map

	var m3 map[string]int // m3 == nil
	fmt.Println("m1, m2, m3 = ")
	fmt.Println(m1, m2, m3)

	fmt.Println("Traversing map m:")
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m1["course"]
	fmt.Println("m['course'] =", courseName)

	// 检测
	if causeName, ok := m1["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key 'cause' dose not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m1["name"]
	fmt.Printf("m1['%q'] before delete: %q, %v\n", "name", name, ok)

	// 删除
	delete(m1, "name")
	name, ok = m1["name"]
	fmt.Printf("m1['%q'] after delete: %q, %v\n", "name", name, ok)

}
