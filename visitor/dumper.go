// Package visitor contains walker.visitor implementations
package visitor

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/parser"

	"github.com/z7zmey/php-parser/walker"
)

// Dumper writes ast hierarchy to an io.Writer
// Also prints comments and positions attached to nodes
type Dumper struct {
	Writer     io.Writer
	Indent     string
	Comments   parser.Comments
	Positions  parser.Positions
	NsResolver *NamespaceResolver
}

// EnterNode is invoked at every node in hierarchy
func (d *Dumper) EnterNode(w walker.Walkable) bool {
	n := w.(node.Node)

	fmt.Fprintf(d.Writer, "%v[%v]\n", d.Indent, reflect.TypeOf(n))

	if d.Positions != nil {
		if p := d.Positions[n]; p != nil {
			fmt.Fprintf(d.Writer, "%v\"Position\": %s;\n", d.Indent+"  ", *p)
		}
	}

	if d.NsResolver != nil {
		if namespacedName, ok := d.NsResolver.ResolvedNames[n]; ok {
			fmt.Fprintf(d.Writer, "%v\"NamespacedName\": %s;\n", d.Indent+"  ", namespacedName)
		}
	}

	if d.Comments != nil {
		if c := d.Comments[n]; len(c) > 0 {
			fmt.Fprintf(d.Writer, "%v\"Comments\":\n", d.Indent+"  ")
			for _, cc := range c {
				fmt.Fprintf(d.Writer, "%v%q before %q\n", d.Indent+"    ", cc, comment.TokenNames[cc.TokenName()])
			}
		}
	}

	if a := n.Attributes(); len(a) > 0 {
		for key, attr := range a {
			fmt.Fprintf(d.Writer, "%v\"%v\": %v;\n", d.Indent+"  ", key, attr)
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
