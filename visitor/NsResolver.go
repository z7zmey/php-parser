// Package visitor contains walker.visitor implementations
package visitor

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/walker"
)

type NsResolver struct {
	Namespace     string
	ResolvedNames map[node.Node]string
	Aliases       map[string]map[string]string
}

// NewNsResolver NsResolver type constructor
func NewNsResolver() *NsResolver {
	return &NsResolver{
		Namespace:     "",
		ResolvedNames: map[node.Node]string{},
	}
}

// EnterNode is invoked at every node in heirerchy
func (nsr *NsResolver) EnterNode(w walker.Walkable) bool {
	switch n := w.(type) {
	case *stmt.Namespace:
		nsr.Namespace = concatNameParts(n.NamespaceName.(*name.Name).Parts)
	case *stmt.Use:

	}

	return true
}

// GetChildrenVisitor is invoked at every node parameter that contains children nodes
func (nsr *NsResolver) GetChildrenVisitor(key string) walker.Visitor {
	return nsr
}

// LeaveNode is invoked after node process
func (nsr *NsResolver) LeaveNode(w walker.Walkable) {
	switch n := w.(type) {
	case *stmt.Namespace:
		if n.Stmts != nil {
			nsr.Namespace = ""
		}
	}
}

func concatNameParts(parts []node.Node) string {
	str := ""

	for _, n := range parts {
		if str == "" {
			str = n.(*name.NamePart).Value
		} else {
			str = str + "\\" + n.(*name.NamePart).Value
		}
	}

	return str
}
