package main

import (
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"
)

func main() {

	arg := os.Args[1]

	fmt.Println("Generating QR code for:", arg)

	qrterminal.Generate(arg, qrterminal.L, os.Stdout)

}
