package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type numericToken uint64
type commandToken string
type token interface{}
type tokenList []token
type numberStack []numericToken
type commandMap map[commandToken]tokenList
type commandInitializationMap map[string]string

var nStack numberStack
var cMap = make(commandMap)

func main() {
	fmt.Println("Starting GoForth")

	if e := initializeCommandMap(commandInitializationMap{"rickTest": "predefined1 321 pop predefined2"}); e != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", e)
		os.Exit(1)
	}

	testString := "predefined1 123 predefined2 rickTest 123456 654321 add"

	if tList, e := tokenizeString(testString); e != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", e)
		os.Exit(2)
	} else {
		if e := executeTokenList(tList); e != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", e)
			os.Exit(3)
		}
	}

	fmt.Println("Finished GoForth")
	fmt.Printf("Number stack looks like this: %v\n", nStack)
}

func tokenizeString(toTokenize string) (tokenList, error) {
	var retVal tokenList

	for _, token := range strings.Fields(toTokenize) {

		if n, e := strconv.ParseUint(token, 10, 64); e != nil {
			retVal = append(retVal, commandToken(token))
		} else {
			retVal = append(retVal, numericToken(n))
		}
	}

	return retVal, nil
}

func executeToken(token token) error {
	switch v := token.(type) {
	case numericToken:
		//		nStack = append(nStack, v)
		nStack = append(numberStack{v}, nStack...)
		fmt.Printf("Pushed %d, Number stack looks like this: %v\n\n", v, nStack)
	case commandToken:
		fmt.Printf("Running command: %v\n", v)
		switch v {
		case "predefined1":
			break
		case "predefined2":
			break
		case "add":
			var x numericToken
			var x2 numericToken
			x, nStack = nStack[0], nStack[1:]
			x2, nStack = nStack[0], nStack[1:]
			nStack = append(numberStack{x + x2}, nStack...)
			fmt.Printf("Added %d and %d to get %d\n", x, x2, x+x2)
		case "pop":
			var x numericToken
			x, nStack = nStack[0], nStack[1:]
			fmt.Printf("Popped %d off top of stack\n", x)
		default:
			tList := cMap[v]
			return executeTokenList(tList)
		}

		fmt.Printf("After command, number stack looks like this: %v\n\n", nStack)

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

func initializeCommandMap(cInitMap commandInitializationMap) error {
	for key, commandList := range cInitMap {
		tList, e := tokenizeString(commandList)
		if e != nil {
			return e
		}

		cMap[commandToken(key)] = tList
	}

	return nil
}
