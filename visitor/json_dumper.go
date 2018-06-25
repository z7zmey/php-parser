// Package visitor contains walker.visitor implementations
package visitor

import (
	"fmt"
	"io"
	"reflect"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/parser"

	"github.com/z7zmey/php-parser/walker"
)

type JsonDumper struct {
	Writer     io.Writer
	Comments   parser.Comments
	NsResolver *NamespaceResolver
}

// EnterNode is invoked at every node in hierarchy
func (d *JsonDumper) EnterNode(w walker.Walkable) bool {
	n := w.(node.Node)

	nodeType := reflect.TypeOf(n).String()

	fmt.Fprintf(d.Writer, "{%q:%q", "type", nodeType)

	if p := n.GetPosition(); p != nil {
		p := n.GetPosition()
		fmt.Fprintf(d.Writer, ",%q:{%q:%d,%q:%d,%q:%d,%q:%d}",
			"position",
			"startPos", p.StartPos,
			"endPos", p.EndPos,
			"startLine", p.StartLine,
			"endLine", p.EndLine)
	}

	if d.NsResolver != nil {
		if namespacedName, ok := d.NsResolver.ResolvedNames[n]; ok {
			fmt.Fprintf(d.Writer, ",%q:%q", "namespacedName", namespacedName)
		}
	}

	if c := n.GetComments(); len(c) > 0 {
		fmt.Fprintf(d.Writer, ",%q:[", "comments")

		for k, cc := range c {
			if k == 0 {
				fmt.Fprintf(d.Writer, "%q", cc)
			} else {
				fmt.Fprintf(d.Writer, ",%q", cc)
			}
		}

		fmt.Fprint(d.Writer, "]")
	}

	if a := n.Attributes(); len(a) > 0 {
		for key, attr := range a {
			switch attr.(type) {
			case string:
				fmt.Fprintf(d.Writer, ",\"%s\":%q", key, attr)
			default:
				fmt.Fprintf(d.Writer, ",\"%s\":%v", key, attr)
			}
		}
	}

	return true
}

// LeaveNode is invoked after node process
func (d *JsonDumper) LeaveNode(n walker.Walkable) {
	fmt.Fprint(d.Writer, "}")
}

func (d *JsonDumper) EnterChildNode(key string, w walker.Walkable) {
	fmt.Fprintf(d.Writer, ",%q:", key)
}

func (d *JsonDumper) LeaveChildNode(key string, w walker.Walkable) {
	// do nothing
}

func (d *JsonDumper) EnterChildList(key string, w walker.Walkable) {
	fmt.Fprintf(d.Writer, ",%q:[", key)

}

func (d *JsonDumper) LeaveChildList(key string, w walker.Walkable) {
	fmt.Fprint(d.Writer, "]")
}
