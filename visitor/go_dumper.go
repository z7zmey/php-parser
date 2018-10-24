// Package visitor contains walker.visitor implementations
package visitor

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/z7zmey/php-parser/meta"

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

	if p := n.GetPosition(); p != nil {
		printIndent(d.Writer, d.depth)
		fmt.Fprint(d.Writer, "Position: &position.Position{\n")
		d.depth++
		printIndent(d.Writer, d.depth)
		fmt.Fprintf(d.Writer, "StartLine: %d,\n", p.StartLine)
		printIndent(d.Writer, d.depth)
		fmt.Fprintf(d.Writer, "EndLine: %d,\n", p.EndLine)
		printIndent(d.Writer, d.depth)
		fmt.Fprintf(d.Writer, "StartPos: %d,\n", p.StartPos)
		printIndent(d.Writer, d.depth)
		fmt.Fprintf(d.Writer, "EndPos: %d,\n", p.EndPos)
		d.depth--
		printIndent(d.Writer, d.depth)
		fmt.Fprint(d.Writer, "},\n")
	}

	if mm := n.GetMeta(); len(*mm) > 0 {
		printIndent(d.Writer, d.depth)
		fmt.Fprint(d.Writer, "Meta: meta.Collection{\n")
		d.depth++
		for _, m := range *mm {
			printIndent(d.Writer, d.depth)
			fmt.Fprint(d.Writer, "&meta.Data{\n")
			d.depth++

			printIndent(d.Writer, d.depth)

			switch m.Type {
			case meta.CommentType:
				fmt.Fprint(d.Writer, "Type: meta.CommentType,\n")
			case meta.WhiteSpaceType:
				fmt.Fprint(d.Writer, "Type: meta.WhiteSpaceType,\n")
			case meta.TokenType:
				fmt.Fprint(d.Writer, "Type: meta.TokenType,\n")
			}

			printIndent(d.Writer, d.depth)

			if m.Position != nil {
				fmt.Fprint(d.Writer, "Position: &position.Position{\n")
				d.depth++
				printIndent(d.Writer, d.depth)
				fmt.Fprintf(d.Writer, "StartLine: %d,\n", m.Position.StartLine)
				printIndent(d.Writer, d.depth)
				fmt.Fprintf(d.Writer, "EndLine: %d,\n", m.Position.EndLine)
				printIndent(d.Writer, d.depth)
				fmt.Fprintf(d.Writer, "StartPos: %d,\n", m.Position.StartPos)
				printIndent(d.Writer, d.depth)
				fmt.Fprintf(d.Writer, "EndPos: %d,\n", m.Position.EndPos)
				d.depth--
				printIndent(d.Writer, d.depth)
				fmt.Fprint(d.Writer, "},\n")
			} else {
				fmt.Fprint(d.Writer, "Position: nil,\n")
			}

			printIndent(d.Writer, d.depth)
			fmt.Fprintf(d.Writer, "Value: %q,\n", m.String())
			printIndent(d.Writer, d.depth)
			fmt.Fprintf(d.Writer, "TokenName: meta.%s,\n", m.TokenName.String())

			d.depth--
			printIndent(d.Writer, d.depth)
			fmt.Fprint(d.Writer, "},\n")
		}
		d.depth--
		printIndent(d.Writer, d.depth)
		fmt.Fprint(d.Writer, "},\n")
	}

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
