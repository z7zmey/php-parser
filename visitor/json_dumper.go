// Package visitor contains walker.visitor implementations
package visitor

import (
	"fmt"
	"io"
	"reflect"
	"sort"

	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

type JsonDumper struct {
	Writer         io.Writer
	NsResolver     *NamespaceResolver
	isChildNode    bool
	isNotFirstNode bool
}

// EnterNode is invoked at every node in hierarchy
func (d *JsonDumper) EnterNode(w walker.Walkable) bool {
	n := w.(node.Node)

	nodeType := reflect.TypeOf(n).String()

	if d.isChildNode {
		d.isChildNode = false
	} else if d.isNotFirstNode {
		fmt.Fprint(d.Writer, ",")
	} else {
		d.isNotFirstNode = true
	}

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

	if !n.GetFreeFloating().IsEmpty() {
		fmt.Fprintf(d.Writer, ",%q:{", "freefloating")

		var freefloatingStringsKeys []int
		for key := range *n.GetFreeFloating() {
			freefloatingStringsKeys = append(freefloatingStringsKeys, int(key))
		}

		sort.Ints(freefloatingStringsKeys)

		i := 0
		for _, k := range freefloatingStringsKeys {
			key := freefloating.Position(k)
			freeFloatingStrings := (*n.GetFreeFloating())[key]
			if i != 0 {
				fmt.Fprint(d.Writer, ",")
			}
			i++

			fmt.Fprintf(d.Writer, "%q: [", key.String())

			j := 0
			for _, freeFloatingString := range freeFloatingStrings {
				if j != 0 {
					fmt.Fprint(d.Writer, ",")
				}
				j++

				switch freeFloatingString.StringType {
				case freefloating.CommentType:
					fmt.Fprintf(d.Writer, "{%q:%q,%q:%q}", "type", "freefloating.CommentType", "value", freeFloatingString.Value)
				case freefloating.WhiteSpaceType:
					fmt.Fprintf(d.Writer, "{%q:%q,%q:%q}", "type", "freefloating.WhiteSpaceType", "value", freeFloatingString.Value)
				case freefloating.TokenType:
					fmt.Fprintf(d.Writer, "{%q:%q,%q:%q}", "type", "freefloating.TokenType", "value", freeFloatingString.Value)
				}
			}

			fmt.Fprint(d.Writer, "]")
		}

		fmt.Fprint(d.Writer, "}")
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
	d.isChildNode = true
}

func (d *JsonDumper) LeaveChildNode(key string, w walker.Walkable) {
	// do nothing
}

func (d *JsonDumper) EnterChildList(key string, w walker.Walkable) {
	fmt.Fprintf(d.Writer, ",%q:[", key)
	d.isNotFirstNode = false

}

func (d *JsonDumper) LeaveChildList(key string, w walker.Walkable) {
	fmt.Fprint(d.Writer, "]")
}
