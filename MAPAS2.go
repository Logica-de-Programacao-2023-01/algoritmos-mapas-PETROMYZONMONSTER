package main

import "fmt"

func main() {
	map1 := map[string]int{
		"val1": 11,
		"val2": 22,
		"val3": 23,
	}
	map2 := map[string]int{
		"val3": 33,
		"val4": 44,
	}
	mapfinal := mapasjuntos(map1, map2)
	fmt.Println("Seu mapa:", mapfinal)
}

func mapasjuntos(map1 map[string]int, map2 map[string]int) map[string]int {
	final := make(map[string]int)
	for var1, range1 := range map1 {
		final[var1] = range1
	}
	for var2, range2 := range map2 {
		final[var2] = range2
	}
	return final
}
