package parser

import (
	"fmt"
	"strings"
)


var traceLevel int = 0
const traceIdentPlaceHolder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceHolder, traceLevel - 1 )
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n",identLevel(),fs)
}

func incIdent() {
	traceLevel = traceLevel + 1
}
func decIdent() {
	traceLevel = traceLevel - 1
}

func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN "+msg)
	return msg
}

func untrace(msg string) string {
	tracePrint("END "+msg)
	return msg
}