package main

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"
	"syscall/js"
	"time"

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

	var (
		content  = args[0].String()
		inline   = args[1].Bool()
		withJSON = args[2].Bool()
	)

	_, err := rootNode.ReadFrom(strings.NewReader(content))
	if err != nil {
		fmt.Printf("Cannot read Node from source XML with err: %s\n", err)

		return ""
	}

	_ = inline

	sw.Banner = fmt.Sprintf("generated %s by %s in Ukraine.",
		time.Now().Format("2006-01-02 15:04:05"),
		"https://xml-to-go.github.io/",
	)
	sw.WithJSONTags = withJSON

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
