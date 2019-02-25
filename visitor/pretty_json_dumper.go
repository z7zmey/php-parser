// Package visitor contains walker.visitor implementations
package visitor

import (
	"fmt"
	"io"
	"reflect"

	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

type PrettyJsonDumper struct {
	Writer         io.Writer
	NsResolver     *NamespaceResolver
	depth          int
	isChildNode    bool
	isNotFirstNode bool
}

func NewPrettyJsonDumper(Writer io.Writer, NsResolver *NamespaceResolver) *PrettyJsonDumper {
	return &PrettyJsonDumper{
		Writer:         Writer,
		NsResolver:     NsResolver,
		depth:          0,
		isChildNode:    false,
		isNotFirstNode: false,
	}
}

func (d *PrettyJsonDumper) printIndent(w io.Writer) {
	for i := 0; i < d.depth; i++ {
		fmt.Fprint(d.Writer, "  ")
	}
}

// EnterNode is invoked at every node in hierarchy
func (d *PrettyJsonDumper) EnterNode(w walker.Walkable) bool {
	n := w.(node.Node)

	nodeType := reflect.TypeOf(n).String()

	if d.isChildNode {
		d.isChildNode = false
	} else if d.isNotFirstNode {
		fmt.Fprint(d.Writer, ",\n")
		d.printIndent(d.Writer)
	} else {
		d.printIndent(d.Writer)
		d.isNotFirstNode = true
	}

	fmt.Fprint(d.Writer, "{\n")
	d.depth++
	d.printIndent(d.Writer)
	fmt.Fprintf(d.Writer, "%q: %q", "type", nodeType)

	if p := n.GetPosition(); p != nil {
		fmt.Fprint(d.Writer, ",\n")
		d.printIndent(d.Writer)
		fmt.Fprintf(d.Writer, "%q: {\n", "position")
		d.depth++
		d.printIndent(d.Writer)
		fmt.Fprintf(d.Writer, "%q: %d,\n", "startPos", p.StartPos)
		d.printIndent(d.Writer)
		fmt.Fprintf(d.Writer, "%q: %d,\n", "endPos", p.EndPos)
		d.printIndent(d.Writer)
		fmt.Fprintf(d.Writer, "%q: %d,\n", "startLine", p.StartLine)
		d.printIndent(d.Writer)
		fmt.Fprintf(d.Writer, "%q: %d\n", "endLine", p.EndLine)
		d.depth--
		d.printIndent(d.Writer)
		fmt.Fprint(d.Writer, "}")
	}

	if d.NsResolver != nil {
		if namespacedName, ok := d.NsResolver.ResolvedNames[n]; ok {
			fmt.Fprint(d.Writer, ",\n")
			d.printIndent(d.Writer)
			fmt.Fprintf(d.Writer, "\"namespacedName\": %q", namespacedName)
		}
	}

	if !n.GetFreeFloating().IsEmpty() {
		fmt.Fprint(d.Writer, ",\n")
		d.printIndent(d.Writer)
		fmt.Fprint(d.Writer, "\"freefloating\": {\n")
		d.depth++
		i := 0
		for key, freeFloatingStrings := range *n.GetFreeFloating() {
			if i != 0 {
				fmt.Fprint(d.Writer, ",\n")
			}
			i++

			d.printIndent(d.Writer)
			fmt.Fprintf(d.Writer, "%q: [\n", key)
			d.depth++

			j := 0
			for _, freeFloatingString := range freeFloatingStrings {
				if j != 0 {
					fmt.Fprint(d.Writer, ",\n")
				}
				j++

				d.printIndent(d.Writer)
				fmt.Fprint(d.Writer, "{\n")
				d.depth++
				d.printIndent(d.Writer)
				switch freeFloatingString.StringType {
				case freefloating.CommentType:
					fmt.Fprintf(d.Writer, "%q: %q,\n", "type", "freefloating.CommentType")
				case freefloating.WhiteSpaceType:
					fmt.Fprintf(d.Writer, "%q: %q,\n", "type", "freefloating.WhiteSpaceType")
				case freefloating.TokenType:
					fmt.Fprintf(d.Writer, "%q: %q,\n", "type", "freefloating.TokenType")
				}
				d.printIndent(d.Writer)
				fmt.Fprintf(d.Writer, "%q: %q\n", "value", freeFloatingString.Value)
				d.depth--
				d.printIndent(d.Writer)
				fmt.Fprint(d.Writer, "}")
			}

			d.depth--
			fmt.Fprint(d.Writer, "\n")
			d.printIndent(d.Writer)
			fmt.Fprint(d.Writer, "]")
		}
		d.depth--
		fmt.Fprint(d.Writer, "\n")
		d.printIndent(d.Writer)
		fmt.Fprint(d.Writer, "}")
	}

	if a := n.Attributes(); len(a) > 0 {
		for key, attr := range a {
			fmt.Fprint(d.Writer, ",\n")
			d.printIndent(d.Writer)
			switch attr.(type) {
			case string:
				fmt.Fprintf(d.Writer, "\"%s\": %q", key, attr)
			default:
				fmt.Fprintf(d.Writer, "\"%s\": %v", key, attr)
			}
		}
	}

	return true
}

// LeaveNode is invoked after node process
func (d *PrettyJsonDumper) LeaveNode(n walker.Walkable) {
	d.depth--
	fmt.Fprint(d.Writer, "\n")
	d.printIndent(d.Writer)
	fmt.Fprint(d.Writer, "}")
}

func (d *PrettyJsonDumper) EnterChildNode(key string, w walker.Walkable) {
	fmt.Fprint(d.Writer, ",\n")
	d.printIndent(d.Writer)
	fmt.Fprintf(d.Writer, "%q: ", key)
	d.isChildNode = true
}

func (d *PrettyJsonDumper) LeaveChildNode(key string, w walker.Walkable) {
	// do nothing
}

func (d *PrettyJsonDumper) EnterChildList(key string, w walker.Walkable) {
	fmt.Fprint(d.Writer, ",\n")
	d.printIndent(d.Writer)
	fmt.Fprintf(d.Writer, "%q: [\n", key)
	d.depth++

	d.isNotFirstNode = false
}

func (d *PrettyJsonDumper) LeaveChildList(key string, w walker.Walkable) {
	d.depth--
	fmt.Fprint(d.Writer, "\n")
	d.printIndent(d.Writer)
	fmt.Fprint(d.Writer, "]")
}
