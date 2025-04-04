package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Lox struct {
}

func (lox *Lox) error(line int, msg string) {
	lox.report(line, "", msg)
}

func (lox *Lox) report(line int, where string, msg string) {
	fmt.Printf("[line %d] error where: %s %s\n", line, where, msg)
}

func (lox *Lox) RunFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return lox.run(string(bytes))
}

func (lox *Lox) RunPrompt() error {
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
		if err := lox.run(strings.TrimSpace(string(bytes))); err != nil {
			return err
		}
	}
}

func (lox *Lox) run(source string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	scanner := NewScanner(source)
	tokens := scanner.scanTokens()

	parser := NewParser(tokens)
	expr := parser.parse()

	interpreter := NewInterpreter(expr)
	value := interpreter.evaluate(expr)
	fmt.Printf("%f\n", value)
	return nil
}
