// Package visitor contains walker.visitor implementations
package visitor

import (
	"errors"
	"strings"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/walker"
)

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

type Namespace struct {
	Namespace string
	Aliases   map[string]map[string]string
}

func NewNamespace(NSName string) *Namespace {
	return &Namespace{
		Namespace: NSName,
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

func (ns *Namespace) resolveName(nameNode node.Node, aliasType string) string {
	switch n := nameNode.(type) {
	case *name.FullyQualified:
		return concatNameParts(n.Parts)

	case *name.Relative:
		return ns.Namespace + "\\" + concatNameParts(n.Parts)

	case *name.Name:
		aliasName, err := ns.resolveAlias(nameNode, aliasType)
		if err != nil {
			return ns.Namespace + "\\" + concatNameParts(n.Parts)
		}

		if len(n.Parts) > 1 {
			return aliasName + "\\" + concatNameParts(n.Parts[1:])
		}

		return aliasName
	}

	panic("invalid nameNode variable type")
}

func (ns *Namespace) resolveAlias(nameNode node.Node, aliasType string) (string, error) {
	aliasType = strings.ToLower(aliasType)
	nameParts := nameNode.(*name.Name).Parts

	firstPartStr := nameParts[0].(*name.NamePart).Value

	if len(nameParts) > 1 { // resolve aliases for qualified names, always against class alias table
		firstPartStr = strings.ToLower(firstPartStr)
		aliasType = ""
	} else {
		if aliasType != "const" { // constans are case-sensitive
			firstPartStr = strings.ToLower(firstPartStr)
		}
	}

	aliasName, ok := ns.Aliases[aliasType][firstPartStr]
	if !ok {
		return "", errors.New("Not found")
	}

	return aliasName, nil
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

	case *stmt.Class:
		if n.Extends != nil {
			nsr.resolveName(n.Extends, "")
		}

		for _, interfaceName := range n.Implements {
			nsr.resolveName(interfaceName, "")
		}

		nsr.addNamespacedName(n, n.ClassName.(*node.Identifier).Value)

	case *stmt.Interface:
		for _, interfaceName := range n.Extends {
			nsr.resolveName(interfaceName, "")
		}

		nsr.addNamespacedName(n, n.InterfaceName.(*node.Identifier).Value)

	case *stmt.Trait:
		nsr.addNamespacedName(n, n.TraitName.(*node.Identifier).Value)

	case *stmt.Function:
		nsr.addNamespacedName(n, n.FunctionName.(*node.Identifier).Value)

		for _, parameter := range n.Params {
			nsr.resolveType(parameter.(*node.Parameter).VariableType)
		}

		if n.ReturnType != nil {
			nsr.resolveType(n.ReturnType)
		}

	case *stmt.ClassMethod:
		for _, parameter := range n.Params {
			nsr.resolveType(parameter.(*node.Parameter).VariableType)
		}

		if n.ReturnType != nil {
			nsr.resolveType(n.ReturnType)
		}

	case *expr.Closure:
		for _, parameter := range n.Params {
			nsr.resolveType(parameter.(*node.Parameter).VariableType)
		}

		if n.ReturnType != nil {
			nsr.resolveType(n.ReturnType)
		}

	case *stmt.ConstList:
		for _, constant := range n.Consts {
			nsr.addNamespacedName(constant, constant.(*stmt.Constant).ConstantName.(*node.Identifier).Value)
		}

	case *expr.StaticCall:
		switch nn := n.Class.(type) {
		case *name.Name:
			nsr.resolveName(nn, "")
		case *name.Relative:
			nsr.resolveName(nn, "")
		case *name.FullyQualified:
			nsr.resolveName(nn, "")
		}

	case *expr.StaticPropertyFetch:
		switch nn := n.Class.(type) {
		case *name.Name:
			nsr.resolveName(nn, "")
		case *name.Relative:
			nsr.resolveName(nn, "")
		case *name.FullyQualified:
			nsr.resolveName(nn, "")
		}

	case *expr.ClassConstFetch:
		switch nn := n.Class.(type) {
		case *name.Name:
			nsr.resolveName(nn, "")
		case *name.Relative:
			nsr.resolveName(nn, "")
		case *name.FullyQualified:
			nsr.resolveName(nn, "")
		}

	case *expr.New:
		switch nn := n.Class.(type) {
		case *name.Name:
			nsr.resolveName(nn, "")
		case *name.Relative:
			nsr.resolveName(nn, "")
		case *name.FullyQualified:
			nsr.resolveName(nn, "")
		}

	case *expr.InstanceOf:
		switch nn := n.Class.(type) {
		case *name.Name:
			nsr.resolveName(nn, "")
		case *name.Relative:
			nsr.resolveName(nn, "")
		case *name.FullyQualified:
			nsr.resolveName(nn, "")
		}

	case *stmt.Catch:
		for _, t := range n.Types {
			nsr.resolveName(t, "")
		}

	case *expr.FunctionCall:
		switch nn := n.Function.(type) {
		case *name.Name:
			nsr.resolveName(nn, "function")
		case *name.Relative:
			nsr.resolveName(nn, "function")
		case *name.FullyQualified:
			nsr.resolveName(nn, "function")
		}

	case *expr.ConstFetch:
		nsr.resolveName(n.Constant, "const")

	case *stmt.TraitUse:
		for _, t := range n.Traits {
			nsr.resolveName(t, "")
		}

		for _, a := range n.Adaptations {
			switch aa := a.(type) {
			case *stmt.TraitUsePrecedence:
				refTrait := aa.Ref.(*stmt.TraitMethodRef).Trait
				if refTrait != nil {
					nsr.resolveName(refTrait, "")
				}
				for _, insteadOf := range aa.Insteadof {
					nsr.resolveName(insteadOf, "")
				}

			case *stmt.TraitUseAlias:
				refTrait := aa.Ref.(*stmt.TraitMethodRef).Trait
				if refTrait != nil {
					nsr.resolveName(refTrait, "")
				}
			}
		}

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
			nsr.Namespace = NewNamespace("")
		}
	}
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

func (nsr *NsResolver) addNamespacedName(nn node.Node, nodeName string) {
	if nsr.Namespace.Namespace == "" {
		nsr.ResolvedNames[nn] = nodeName
	} else {
		nsr.ResolvedNames[nn] = nsr.Namespace.Namespace + "\\" + nodeName
	}
}

func (nsr *NsResolver) resolveName(nameNode node.Node, aliasType string) {
	nsr.ResolvedNames[nameNode] = nsr.Namespace.resolveName(nameNode, aliasType)
}

func (nsr *NsResolver) resolveType(n node.Node) {
	switch nn := n.(type) {
	case *node.Nullable:
		nsr.resolveType(nn.Expr)
	case *name.Name:
		nsr.resolveName(n, "")
	case *name.Relative:
		nsr.resolveName(n, "")
	case *name.FullyQualified:
		nsr.resolveName(n, "")
	}
}
