package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Lox struct {
}

func (g *Lox) error(line int, msg string) {
	g.report(line, "", msg)
}

func (g *Lox) report(line int, where string, msg string) {
	fmt.Printf("[line %d] error where: %s %s\n", line, where, msg)
}

func (g *Lox) RunFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return g.run(string(bytes))
}

func (g *Lox) RunPrompt() error {
	for {
		fmt.Printf("> ")
		ioScanner := bufio.NewReader(os.Stdin)
		bytes, err := ioScanner.ReadBytes('\n')
		if err != nil {
			return err
		}
		if len(bytes) == 1 && bytes[0] == '\x03' {
			return nil
		}
		fmt.Printf("%c", bytes[0])
		if err := g.run(strings.TrimSpace(string(bytes))); err != nil {
			return err
		}
	}
}

func (g *Lox) run(source string) error {
	scanner := NewScanner(source)
	tokens := scanner.scanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}

	return nil
}
