package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Golang WebAssembly main")

	done := make(chan struct{})
	js.Global().Set("xmlDataToGoTypeCode", js.FuncOf(xmlDataToGoTypeCodeWasmWrapper))
	<-done
}

func xmlDataToGoTypeCodeWasmWrapper(this js.Value, args []js.Value) interface{} {
	var (
		content  = args[0].String()
		inline   = args[1].Bool()
		compact  = args[2].Bool()
		withJSON = args[3].Bool()
	)

	return xmlDataToGoTypeCode(content, inline, compact, withJSON)
}
