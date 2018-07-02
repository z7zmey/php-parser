// Package visitor contains walker.visitor implementations
package visitor

import (
	"fmt"
	"io"
	"reflect"
	"sort"

	"github.com/z7zmey/php-parser/meta"

	"github.com/z7zmey/php-parser/node"

	"github.com/z7zmey/php-parser/walker"
)

type JsonDumper struct {
	Writer     io.Writer
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

	if mm := n.GetMeta(); len(mm) > 0 {
		fmt.Fprintf(d.Writer, ",%q:[", "meta")

		for k, m := range mm {
			if k != 0 {
				fmt.Fprint(d.Writer, ",")
			}

			switch m.(type) {
			case *meta.Comment:
				fmt.Fprintf(d.Writer, "{%q:%q,%q:%q,%q:%q}", "type", "*meta.Comment", "value", m.String(), "tokenName", m.GetTokenName().String())
			case *meta.WhiteSpace:
				fmt.Fprintf(d.Writer, "{%q:%q,%q:%q,%q:%q}", "type", "*meta.WhiteSpace", "value", m.String(), "tokenName", m.GetTokenName().String())
			}
		}

		fmt.Fprint(d.Writer, "]")
	}

	if a := n.Attributes(); len(a) > 0 {
		var attributes []string
		for key := range n.Attributes() {
			attributes = append(attributes, key)
		}

		sort.Strings(attributes)

		for _, attributeName := range attributes {
			attr := a[attributeName]
			switch attr.(type) {
			case string:
				fmt.Fprintf(d.Writer, ",\"%s\":%q", attributeName, attr)
			default:
				fmt.Fprintf(d.Writer, ",\"%s\":%v", attributeName, attr)
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
