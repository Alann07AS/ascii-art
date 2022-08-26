package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	standard, _ := os.ReadFile("standard")
	s := strings.Split(string(standard), "\n")
	args := os.Args
	txt := args[1] // strings.Join(args[1:], " ")
	txtTable := strings.Split(txt, ("\\" + "n"))
	if strings.Join(txtTable, "") == "" {
		for i := 1; i <= len(txtTable)-1; i++ {
			fmt.Println()
		}
		return
	}
	for _, w := range txtTable {
		if w == "" {
			fmt.Println()
			continue
		}
		for i := 0; i != 8; i++ {
			for _, each := range w {
				sCache := s[(i + (int(each)-32)*9 + 1)]
				fmt.Print(string(sCache))
			}
			fmt.Println()
		}
	}
}
