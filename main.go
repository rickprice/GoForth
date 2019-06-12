package main

import (
	"errors"
	"fmt"
	"os"
)

type numericToken uint64
type commandToken string
type token interface{}
type tokenList []token
type numberStack []numericToken
type commandMap map[commandToken]tokenList

var nStack numberStack
var cMap commandMap

func main() {
	fmt.Println("Starting GoForth")

	if tList, e := tokenizeString("predefined1 123 predefined2"); e != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", e)
		os.Exit(1)
	} else {
		if e := executeTokenList(tList); e != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", e)
			os.Exit(1)
		}
	}
}

// +++ FIX THIS +++ Doesn't actually tokenize the string yet...
func tokenizeString(toTokenize string) (tokenList, error) {
	var retVal tokenList
	retVal = append(retVal, commandToken("predefined"))
	return retVal, nil
}

// +++ FIX THIS +++ Doesn't actually tokenize the string yet...

func executeToken(token token) error {
	switch v := token.(type) {
	case numericToken:
		nStack = append(nStack, v)
	case commandToken:
		switch v {
		case "predefined1":
			fmt.Println("predifined1 command run")
		case "predefined2":
			fmt.Println("predifined2 command run")
		default:
			tList := cMap[v]
			return executeTokenList(tList)
		}
	default:
		return errors.New("goforth: unknown token type")
	}

	return nil
}

func executeTokenList(tList tokenList) error {
	for _, token := range tList {
		if e := executeToken(token); e != nil {
			return e
		}
	}
	return nil
}
