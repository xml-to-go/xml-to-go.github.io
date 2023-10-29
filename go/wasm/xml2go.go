package main

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"
	"time"

	"github.com/miku/zek"
)

func xmlDataToGoTypeCode(content string, inline, compact, withJSON bool) string {
	var rootNode = new(zek.Node)
	_, err := rootNode.ReadFrom(strings.NewReader(content), &zek.ReadOpts{})
	if err != nil {
		fmt.Printf("github.com/miku/zek::Node.ReadFrom cannot decode XML with err: %s\n", err)

		return ""
	}

	var (
		buffer = new(bytes.Buffer)
		sw     = zek.NewStructWriter(buffer)
	)
	sw.Banner = fmt.Sprintf(
		"generated %s by %s in Ukraine.",
		time.Now().Format("2006-01-02 15:04:05"),
		"https://xml-to-go.github.io/",
	)
	_ = inline // @TODO sw.Inline = inline after https://github.com/miku/zek/issues/14
	sw.Compact = compact
	sw.WithJSONTags = withJSON

	err = sw.WriteNode(rootNode)
	if err != nil {
		fmt.Printf("github.com/miku/zek::StructWriter.WriteNode cannot generate Go code with err: %s\n", err)

		return ""
	}

	source, err := format.Source(buffer.Bytes())
	if err != nil {
		fmt.Printf("go/format::.Source cannot format Go code with err: %s\n", err)

		return ""
	}

	return string(source)
}
