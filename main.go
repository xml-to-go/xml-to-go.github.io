package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Golang WebAssembly main")

	done := make(chan struct{})
	js.Global().Set("xmlDataToGoTypeCode", js.FuncOf(xmlDataToGoTypeCode))
	<-done
}

func xmlDataToGoTypeCode(this js.Value, args []js.Value) interface{} {
	return args[0].String()
}
