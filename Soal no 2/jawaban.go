package main

import "fmt"

type UserInput interface {
	Add(rune)
	GetValue() string
}

type NumericInput struct {
	input string
}

func (n *NumericInput) Add(r rune) {
	if r >= '0' && r <= '9' {
		n.input += string(r)
	}
}

func (n *NumericInput) GetValue() string {
	return n.input
}

func main() {
	var input UserInput = &NumericInput{}
	input.Add('1')
	input.Add('a')
	input.Add('0')
	fmt.Println(input.GetValue()) // Output: "10"
}
