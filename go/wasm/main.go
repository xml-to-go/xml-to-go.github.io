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

func xmlDataToGoTypeCode(content string, inline, compact, withJSON bool) string {
	_ = inline // @TODO after https://github.com/miku/zek/issues/14

	var (
		buffer   = new(bytes.Buffer)
		rootNode = new(zek.Node)
		sw       = zek.NewStructWriter(buffer)
	)

	_, err := rootNode.ReadFrom(strings.NewReader(content))
	if err != nil {
		fmt.Printf("Cannot read Node from source XML with err: %s\n", err)

		return ""
	}

	sw.Banner = fmt.Sprintf(
		"generated %s by %s in Ukraine.",
		time.Now().Format("2006-01-02 15:04:05"),
		"https://xml-to-go.github.io/",
	)

	sw.Compact = compact
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
