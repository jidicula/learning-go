package main

import "fmt"

func main() {
	// map is nil
	var nilMap map[string]int
	mapPrint(nilMap)

	totalWins := map[string]int{} // not nil!
	mapPrint(totalWins)

	teams := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)
	fmt.Println(teams == nil)
	fmt.Println(teams["Lions"])

	ages := make(map[int][]string, 10) // create map with known size
	fmt.Println(ages)

}

// mapPrint
func mapPrint(inputMap map[string]int) {
	fmt.Println(inputMap)
	fmt.Println(inputMap == nil)
	fmt.Println(inputMap["hello"])
}
