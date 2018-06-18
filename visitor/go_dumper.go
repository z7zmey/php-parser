// Package visitor contains walker.visitor implementations
package visitor

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/z7zmey/php-parser/node"

	"github.com/z7zmey/php-parser/walker"
)

// GoDumper writes ast hierarchy to an io.Writer as native Golang struct
type GoDumper struct {
	Writer      io.Writer
	depth       int
	isChildNode bool
}

func printIndent(w io.Writer, d int) {
	for i := 0; i < d; i++ {
		io.WriteString(w, "\t")
	}
}

// EnterNode is invoked at every node in hierarchy
func (d *GoDumper) EnterNode(w walker.Walkable) bool {
	n := w.(node.Node)

	nodeType := reflect.TypeOf(n).String()
	nodeType = strings.Replace(nodeType, "*", "&", 1)

	if d.isChildNode {
		d.isChildNode = false
	} else {
		printIndent(d.Writer, d.depth)
	}

	io.WriteString(d.Writer, nodeType+"{\n")

	d.depth++

	if a := n.Attributes(); len(a) > 0 {
		for key, attr := range a {
			printIndent(d.Writer, d.depth)
			switch attr.(type) {
			case string:
				fmt.Fprintf(d.Writer, "%s: %q,\n", key, attr)
			default:
				fmt.Fprintf(d.Writer, "%s: %v,\n", key, attr)
			}
		}
	}

	return true
}

// LeaveNode is invoked after node process
func (d *GoDumper) LeaveNode(n walker.Walkable) {
	d.depth--
	printIndent(d.Writer, d.depth)
	if d.depth != 0 {
		io.WriteString(d.Writer, "},\n")
	} else {
		io.WriteString(d.Writer, "}\n")
	}
}

func (d *GoDumper) EnterChildNode(key string, w walker.Walkable) {
	printIndent(d.Writer, d.depth)
	io.WriteString(d.Writer, key+": ")
	d.isChildNode = true
}

func (d *GoDumper) LeaveChildNode(key string, w walker.Walkable) {
	// do nothing
}

func (d *GoDumper) EnterChildList(key string, w walker.Walkable) {
	printIndent(d.Writer, d.depth)
	io.WriteString(d.Writer, key+": []node.Node{\n")
	d.depth++
}

func (d *GoDumper) LeaveChildList(key string, w walker.Walkable) {
	d.depth--
	printIndent(d.Writer, d.depth)
	if d.depth != 0 {
		io.WriteString(d.Writer, "},\n")
	}
}
