package forms

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

func PrintMenu(menu string) {
	fmt.Printf("\n%s\n", menu)
	fmt.Printf("\nEscolha a ação desejada (1-6): ")
}
