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

// Dumper writes ast hierarchy to an io.Writer
// Also prints comments and positions attached to nodes
type Dumper struct {
	Writer     io.Writer
	Indent     string
	NsResolver *NamespaceResolver
}

// EnterNode is invoked at every node in hierarchy
func (d *Dumper) EnterNode(w walker.Walkable) bool {
	n := w.(node.Node)

	fmt.Fprintf(d.Writer, "%v[%v]\n", d.Indent, reflect.TypeOf(n))

	if n.GetPosition() != nil {
		fmt.Fprintf(d.Writer, "%v\"Position\": %s\n", d.Indent+"  ", n.GetPosition())
	}

	if d.NsResolver != nil {
		if namespacedName, ok := d.NsResolver.ResolvedNames[n]; ok {
			fmt.Fprintf(d.Writer, "%v\"NamespacedName\": %q\n", d.Indent+"  ", namespacedName)
		}
	}

	if !n.GetFreeFloating().IsEmpty() {
		fmt.Fprintf(d.Writer, "%v\"freefloating\":\n", d.Indent+"  ")
		for key, freeFloatingStrings := range *n.GetFreeFloating() {
			for _, freeFloatingString := range freeFloatingStrings {
				fmt.Fprintf(d.Writer, "%v%q: %q\n", d.Indent+"    ", key.String(), freeFloatingString.Value)
			}
		}
	}

	if a := n.Attributes(); len(a) > 0 {
		for key, attr := range a {
			switch attr.(type) {
			case string:
				fmt.Fprintf(d.Writer, "%v\"%v\": %q\n", d.Indent+"  ", key, attr)
			default:
				fmt.Fprintf(d.Writer, "%v\"%v\": %v\n", d.Indent+"  ", key, attr)
			}
		}
	}

	return true
}

// LeaveNode is invoked after node process
func (d *Dumper) LeaveNode(n walker.Walkable) {
	// do nothing
}

// GetChildrenVisitor is invoked at every node parameter that contains children nodes
func (d *Dumper) EnterChildNode(key string, w walker.Walkable) {
	fmt.Fprintf(d.Writer, "%v%q:\n", d.Indent+"  ", key)
	d.Indent = d.Indent + "    "
}

func (d *Dumper) LeaveChildNode(key string, w walker.Walkable) {
	d.Indent = strings.TrimSuffix(d.Indent, "    ")
}

func (d *Dumper) EnterChildList(key string, w walker.Walkable) {
	fmt.Fprintf(d.Writer, "%v%q:\n", d.Indent+"  ", key)
	d.Indent = d.Indent + "    "
}

func (d *Dumper) LeaveChildList(key string, w walker.Walkable) {
	d.Indent = strings.TrimSuffix(d.Indent, "    ")
}
