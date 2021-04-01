package main

import "fmt"

func main() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])
	foo := totalWins["Orcas"]
	fmt.Println(foo)

	m := map[string]int{"hello": 5, "world": 0}
	v, ok := m["hello"]
	fmt.Println(v, ok)
	v, ok = m["world"]
	fmt.Println(v, ok)
	v, ok = m["goodbye"]
	fmt.Println(v, ok)
	_, ok = m["foobar"]
	fmt.Println(ok)

}
