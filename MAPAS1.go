package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Digite a sua string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(ocorrencias(s))
}

func ocorrencias(s string) map[string]int {
	palavra := strings.Fields(s)
	numeropalavras := make(map[string]int)
	for _, palavras := range palavra {
		numeropalavras[palavras]++
	}
	return numeropalavras
}
