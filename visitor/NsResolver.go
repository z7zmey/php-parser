// Package visitor contains walker.visitor implementations
package visitor

import (
	"strings"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/walker"
)

type Namespace struct {
	Namespace string
	Aliases   map[string]map[string]string
}

func NewNamespace(NSName string) *Namespace {
	return &Namespace{
		Namespace: "",
		Aliases: map[string]map[string]string{
			"":         {},
			"const":    {},
			"function": {},
		},
	}
}

func (ns *Namespace) AddAlias(aliasType string, aliasName string, alias string) {
	aliasType = strings.ToLower(aliasType)

	if aliasType == "const" {
		ns.Aliases[aliasType][alias] = aliasName
	} else {
		ns.Aliases[aliasType][strings.ToLower(alias)] = aliasName
	}
}

type NsResolver struct {
	Namespace     *Namespace
	ResolvedNames map[node.Node]string
}

// NewNsResolver NsResolver type constructor
func NewNsResolver() *NsResolver {
	return &NsResolver{
		Namespace:     NewNamespace(""),
		ResolvedNames: map[node.Node]string{},
	}
}

// EnterNode is invoked at every node in heirerchy
func (nsr *NsResolver) EnterNode(w walker.Walkable) bool {
	switch n := w.(type) {
	case *stmt.Namespace:
		if n.NamespaceName == nil {
			nsr.Namespace = NewNamespace("")
		} else {
			NSParts := n.NamespaceName.(*name.Name).Parts
			nsr.Namespace = NewNamespace(concatNameParts(NSParts))
		}

	case *stmt.UseList:
		useType := ""
		if n.UseType != nil {
			useType = n.UseType.(*node.Identifier).Value
		}

		for _, nn := range n.Uses {
			nsr.AddAlias(useType, nn, nil)
		}

		// no reason to iterate into depth
		return false

	case *stmt.GroupUse:
		useType := ""
		if n.UseType != nil {
			useType = n.UseType.(*node.Identifier).Value
		}

		for _, nn := range n.UseList {
			nsr.AddAlias(useType, nn, n.Prefix.(*name.Name).Parts)
		}

		// no reason to iterate into depth
		return false

	}

	return true
}

func (nsr *NsResolver) AddAlias(useType string, nn node.Node, prefix []node.Node) {
	switch use := nn.(type) {
	case *stmt.Use:
		if use.UseType != nil {
			useType = use.UseType.(*node.Identifier).Value
		}

		useNameParts := use.Use.(*name.Name).Parts
		var alias string
		if use.Alias == nil {
			alias = useNameParts[len(useNameParts)-1].(*name.NamePart).Value
		} else {
			alias = use.Alias.(*node.Identifier).Value
		}

		nsr.Namespace.AddAlias(useType, concatNameParts(prefix, useNameParts), alias)
	}
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
			nsr.Namespace = NewNamespace("")
		}
	}
}

func concatNameParts(parts ...[]node.Node) string {
	str := ""

	for _, p := range parts {
		for _, n := range p {
			if str == "" {
				str = n.(*name.NamePart).Value
			} else {
				str = str + "\\" + n.(*name.NamePart).Value
			}
		}
	}

	return str
}
