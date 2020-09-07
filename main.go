package main

import (
	"log"
	"os"

	"github.com/takeru56/t/parser"
	"github.com/takeru56/t/token"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing argument error")
		return
	}
	tok := token.New(os.Args[1])
	parser := parser.New(tok)
	Program(parser.Program())
}
