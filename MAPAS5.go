package main

import (
	"fmt"
	"strings"
)

func FreqCarac(str string) map[string]int {
	mapa := make(map[string]int)
	chave := strings.Split(str, "")
	for _, rang := range chave {
		mapa[rang] += 1
	}
	return mapa
}

func main() {
	fmt.Println(FreqCarac("Bom dia, flor do dia."))
}
