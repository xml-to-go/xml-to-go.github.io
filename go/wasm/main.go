package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Golang WebAssembly main")

	// globalThis.xmlDataToGoTypeCode = function () {}
	js.Global().Set("xmlDataToGoTypeCode", js.FuncOf(xmlDataToGoTypeCodeWasmWrapper))

	done := make(chan struct{})
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
