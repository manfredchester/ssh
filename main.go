package main

import (
	"fmt"
	"os"
)

func main() {
	Unix()
	// fmt.Println(strings.IndexFunc("m*+&^]&./", s))
	// SSHMaker()
}

func Unix() {
	// func (in io.Reader, out io.Write, args []string )
	// app1 param1 | app2 param2
	// pipe(bind(app1, param1), bind(app2, param2))
	args := os.Args[1:]
	for _, v := range args {
		fmt.Println(v)
	}
	p := pipe(bind(app1, args), app2)
	p(os.Stdin, os.Stdout)
}

func s(c rune) bool {
	// if c != "]" {
	if c != ']' {
		return false
	}
	return true
}
