package main

import (
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"
)

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]

		fmt.Println("Generating QR code for:", arg)

		qrterminal.Generate(arg, qrterminal.L, os.Stdout)
	} else {
		fmt.Println("Plase typing url")
	}
}
