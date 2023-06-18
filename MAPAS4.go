package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	mapa1 := []string{"amor", "maluco", "fim", "roma", "mif", "columa"}
	fmt.Println(anagramas(mapa1))
}

func anagramas(mapa1 []string) map[string][]string {
	anagramas := make(map[string][]string)
	for _, mapa := range mapa1 {
		sorted := sorting(mapa)
		anagramas[sorted] = append(anagramas[sorted], mapa)
	}
	return anagramas
}

func sorting(str string) string {
	sorts := strings.Split(str, "")
	sort.Strings(sorts)
	return strings.Join(sorts, "")
}
