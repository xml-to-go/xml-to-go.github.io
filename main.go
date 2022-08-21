package main

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"
	"syscall/js"

	"github.com/miku/zek"
)

func main() {
	fmt.Println("Golang WebAssembly main")

	done := make(chan struct{})
	js.Global().Set("xmlDataToGoTypeCode", js.FuncOf(xmlDataToGoTypeCode))
	<-done
}

func xmlDataToGoTypeCode(this js.Value, args []js.Value) interface{} {
	var (
		buffer   = new(bytes.Buffer)
		rootNode = new(zek.Node)
		sw       = zek.NewStructWriter(buffer)
	)

	_, err := rootNode.ReadFrom(strings.NewReader(args[0].String()))
	if err != nil {
		fmt.Printf("Cannot read Node from source XML with err: %s\n", err)

		return ""
	}

	err = sw.WriteNode(rootNode)
	if err != nil {
		fmt.Printf("Cannot write Node from source XML with err: %s\n", err)

		return ""
	}

	source, err := format.Source(buffer.Bytes())
	if err != nil {
		fmt.Printf("Format source code with err: %s\n", err)

		return ""
	}

	return string(source)
}
