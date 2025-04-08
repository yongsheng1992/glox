package main

import (
	"fmt"
	"github.com/yongsheng1992/glox/core"
	"io"
	"os"
)

func main() {
	lox := core.NewLox()
	fmt.Println(os.Args)
	if len(os.Args) > 2 {
		fmt.Println("Usage: lox [script]")
	} else if len(os.Args) == 2 {
		if err := lox.RunFile(os.Args[0]); err != nil {
			panic(err)
		}
	} else {
		if err := lox.RunPrompt(); err != nil {
			if err != io.EOF {
				panic(err)
			}
		}
	}
}
