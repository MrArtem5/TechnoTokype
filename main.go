package main

import (
	"TechnoTokype/evalut"
	"TechnoTokype/lexer"
	"TechnoTokype/optimizer"
	"TechnoTokype/parser"
	"fmt"
	"os"
)

const Version = "0.0.2 (unstable)"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tokype <file.tp>")
		return
	}

	if os.Args[1] == "version" {
		fmt.Printf("Tokype version %s\n", Version)
		return
	}
	if os.Args[1] == "help" {
		fmt.Println("tokype <filename>.tp - run tokype script")
		fmt.Println("tokype help - show this help")
		fmt.Println("tokype version - show version")
		fmt.Println("tokype license - show license")
		return
	}
	if os.Args[1] == "license" {
		fmt.Println("License information not available in this build")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	code := string(data)

	l := lexer.NewLexer(code)
	p := parser.NewParser(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		fmt.Println("Parser errors:")
		for _, err := range p.Errors() {
			fmt.Printf("  %s\n", err)
		}
		return
	}

	env := evalut.NewEnvironment()
	optimizedProgram := optimizer.OptimizeAST(program)

	result := evalut.Eval(optimizedProgram, env)

	if !result.IsNil() {
		fmt.Printf("Result: %v\n", result.String())
	}
}
