package main

func checkScreen(lastSymbol string) bool {
	return lastSymbol == "\\"
}

type Token struct {
	Name        string
	Body        string
	Arg         string
	OpenSymbol  int
	CloseSymbol int
	TypeT       string
	TokenNumber int
	TG          string
}
