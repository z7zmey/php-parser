// Package visitor contains walker.visitor implementations
package visitor

import (
	"fmt"
	"reflect"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/walker"
)

// Dumper prints ast hierarchy to stdout
// Also prints comments and positions attached to nodes
type Dumper struct {
	Indent     string
	Comments   comment.Comments
	Positions  position.Positions
	NsResolver *NamespaceResolver
}

// EnterNode is invoked at every node in heirerchy
func (d Dumper) EnterNode(w walker.Walkable) bool {
	n := w.(node.Node)

	fmt.Printf("%v[%v]\n", d.Indent, reflect.TypeOf(n))

	if d.Positions != nil {
		if p := d.Positions[n]; p != nil {
			fmt.Printf("%v\"Position\": %s;\n", d.Indent+"  ", *p)
		}
	}

	if d.NsResolver != nil {
		if namespacedName, ok := d.NsResolver.ResolvedNames[n]; ok {
			fmt.Printf("%v\"NamespacedName\": %s;\n", d.Indent+"  ", namespacedName)
		}
	}

	if d.Comments != nil {
		if c := d.Comments[n]; len(c) > 0 {
			fmt.Printf("%v\"Comments\":\n", d.Indent+"  ")
			for _, cc := range c {
				fmt.Printf("%v%q\n", d.Indent+"    ", cc)
			}
		}
	}

	if a := n.Attributes(); len(a) > 0 {
		for key, attr := range a {
			fmt.Printf("%v\"%v\": %v;\n", d.Indent+"  ", key, attr)
		}
	}

	return true
}

// GetChildrenVisitor is invoked at every node parameter that contains children nodes
func (d Dumper) GetChildrenVisitor(key string) walker.Visitor {
	fmt.Printf("%v%q:\n", d.Indent+"  ", key)
	return Dumper{d.Indent + "    ", d.Comments, d.Positions, d.NsResolver}
}

// LeaveNode is invoked after node process
func (d Dumper) LeaveNode(n walker.Walkable) {
	// do nothing
}
