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

// NamespaceResolver visitor
type NamespaceResolver struct {
	Namespace     *Namespace
	ResolvedNames map[node.Node]string
}

// NewNamespaceResolver NamespaceResolver type constructor
func NewNamespaceResolver() *NamespaceResolver {
	return &NamespaceResolver{
		Namespace:     NewNamespace(""),
		ResolvedNames: map[node.Node]string{},
	}
}

// EnterNode is invoked at every node in heirerchy
func (nsr *NamespaceResolver) EnterNode(w walker.Walkable) bool {
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
			nsr.ResolveName(n.Extends.ClassName, "")
		}

		if n.Implements != nil {
			for _, interfaceName := range n.Implements.InterfaceNames {
				nsr.ResolveName(interfaceName, "")
			}
		}

		if n.ClassName != nil {
			nsr.AddNamespacedName(n, n.ClassName.(*node.Identifier).Value)
		}

	case *stmt.Interface:
		if n.Extends != nil {
			for _, interfaceName := range n.Extends.InterfaceNames {
				nsr.ResolveName(interfaceName, "")
			}
		}

		nsr.AddNamespacedName(n, n.InterfaceName.(*node.Identifier).Value)

	case *stmt.Trait:
		nsr.AddNamespacedName(n, n.TraitName.(*node.Identifier).Value)

	case *stmt.Function:
		nsr.AddNamespacedName(n, n.FunctionName.(*node.Identifier).Value)

		for _, parameter := range n.Params {
			nsr.ResolveType(parameter.(*node.Parameter).VariableType)
		}

		if n.ReturnType != nil {
			nsr.ResolveType(n.ReturnType)
		}

	case *stmt.ClassMethod:
		for _, parameter := range n.Params {
			nsr.ResolveType(parameter.(*node.Parameter).VariableType)
		}

		if n.ReturnType != nil {
			nsr.ResolveType(n.ReturnType)
		}

	case *expr.Closure:
		for _, parameter := range n.Params {
			nsr.ResolveType(parameter.(*node.Parameter).VariableType)
		}

		if n.ReturnType != nil {
			nsr.ResolveType(n.ReturnType)
		}

	case *stmt.ConstList:
		for _, constant := range n.Consts {
			nsr.AddNamespacedName(constant, constant.(*stmt.Constant).ConstantName.(*node.Identifier).Value)
		}

	case *expr.StaticCall:
		nsr.ResolveName(n.Class, "")

	case *expr.StaticPropertyFetch:
		nsr.ResolveName(n.Class, "")

	case *expr.ClassConstFetch:
		nsr.ResolveName(n.Class, "")

	case *expr.New:
		nsr.ResolveName(n.Class, "")

	case *expr.InstanceOf:
		nsr.ResolveName(n.Class, "")

	case *stmt.Catch:
		for _, t := range n.Types {
			nsr.ResolveName(t, "")
		}

	case *expr.FunctionCall:
		nsr.ResolveName(n.Function, "function")

	case *expr.ConstFetch:
		nsr.ResolveName(n.Constant, "const")

	case *stmt.TraitUse:
		for _, t := range n.Traits {
			nsr.ResolveName(t, "")
		}

		if n.TraitAdaptationList != nil {
			for _, a := range n.TraitAdaptationList.Adaptations {
				switch aa := a.(type) {
				case *stmt.TraitUsePrecedence:
					refTrait := aa.Ref.(*stmt.TraitMethodRef).Trait
					if refTrait != nil {
						nsr.ResolveName(refTrait, "")
					}
					for _, insteadOf := range aa.Insteadof {
						nsr.ResolveName(insteadOf, "")
					}

				case *stmt.TraitUseAlias:
					refTrait := aa.Ref.(*stmt.TraitMethodRef).Trait
					if refTrait != nil {
						nsr.ResolveName(refTrait, "")
					}
				}
			}
		}
	}

	return true
}

// GetChildrenVisitor is invoked at every node parameter that contains children nodes
func (nsr *NamespaceResolver) GetChildrenVisitor(key string) walker.Visitor {
	return nsr
}

// LeaveNode is invoked after node process
func (nsr *NamespaceResolver) LeaveNode(w walker.Walkable) {
	switch n := w.(type) {
	case *stmt.Namespace:
		if n.Stmts != nil {
			nsr.Namespace = NewNamespace("")
		}
	}
}

// AddAlias adds a new alias
func (nsr *NamespaceResolver) AddAlias(useType string, nn node.Node, prefix []node.Node) {
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

// AddNamespacedName adds namespaced name by node
func (nsr *NamespaceResolver) AddNamespacedName(nn node.Node, nodeName string) {
	if nsr.Namespace.Namespace == "" {
		nsr.ResolvedNames[nn] = nodeName
	} else {
		nsr.ResolvedNames[nn] = nsr.Namespace.Namespace + "\\" + nodeName
	}
}

// ResolveName adds a resolved fully qualified name by node
func (nsr *NamespaceResolver) ResolveName(nameNode node.Node, aliasType string) {
	resolved, err := nsr.Namespace.ResolveName(nameNode, aliasType)
	if err == nil {
		nsr.ResolvedNames[nameNode] = resolved
	}
}

// ResolveType adds a resolved fully qualified type name
func (nsr *NamespaceResolver) ResolveType(n node.Node) {
	switch nn := n.(type) {
	case *node.Nullable:
		nsr.ResolveType(nn.Expr)
	case name.Names:
		nsr.ResolveName(n, "")
	}
}

// Namespace context
type Namespace struct {
	Namespace string
	Aliases   map[string]map[string]string
}

// NewNamespace constructor
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

// AddAlias adds a new alias
func (ns *Namespace) AddAlias(aliasType string, aliasName string, alias string) {
	aliasType = strings.ToLower(aliasType)

	if aliasType == "const" {
		ns.Aliases[aliasType][alias] = aliasName
	} else {
		ns.Aliases[aliasType][strings.ToLower(alias)] = aliasName
	}
}

// ResolveName returns a resolved fully qualified name
func (ns *Namespace) ResolveName(nameNode node.Node, aliasType string) (string, error) {
	switch n := nameNode.(type) {
	case *name.FullyQualified:
		// Fully qualifid name is already resolved
		return concatNameParts(n.Parts), nil

	case *name.Relative:
		if ns.Namespace == "" {
			return concatNameParts(n.Parts), nil
		}
		return ns.Namespace + "\\" + concatNameParts(n.Parts), nil

	case *name.Name:
		if aliasType == "const" && len(n.Parts) == 1 {
			part := strings.ToLower(n.Parts[0].(*name.NamePart).Value)
			if part == "true" || part == "false" || part == "null" {
				return part, nil
			}
		}

		if aliasType == "" && len(n.Parts) == 1 {
			part := strings.ToLower(n.Parts[0].(*name.NamePart).Value)

			switch part {
			case "self":
				fallthrough
			case "static":
				fallthrough
			case "parent":
				fallthrough
			case "int":
				fallthrough
			case "float":
				fallthrough
			case "bool":
				fallthrough
			case "string":
				fallthrough
			case "void":
				fallthrough
			case "iterable":
				fallthrough
			case "object":
				return part, nil
			}
		}

		aliasName, err := ns.ResolveAlias(nameNode, aliasType)
		if err != nil {
			// resolve as relative name if alias not found
			if ns.Namespace == "" {
				return concatNameParts(n.Parts), nil
			}
			return ns.Namespace + "\\" + concatNameParts(n.Parts), nil
		}

		if len(n.Parts) > 1 {
			// if name qualified, replace first part by alias
			return aliasName + "\\" + concatNameParts(n.Parts[1:]), nil
		}

		return aliasName, nil
	}

	return "", errors.New("must be instance of name.Names")
}

// ResolveAlias returns alias or error if not found
func (ns *Namespace) ResolveAlias(nameNode node.Node, aliasType string) (string, error) {
	aliasType = strings.ToLower(aliasType)
	nameParts := nameNode.(*name.Name).Parts

	firstPartStr := nameParts[0].(*name.NamePart).Value

	if len(nameParts) > 1 { // resolve aliases for qualified names, always against class alias type
		firstPartStr = strings.ToLower(firstPartStr)
		aliasType = ""
	} else {
		if aliasType != "const" { // constants are case-sensitive
			firstPartStr = strings.ToLower(firstPartStr)
		}
	}

	aliasName, ok := ns.Aliases[aliasType][firstPartStr]
	if !ok {
		return "", errors.New("Not found")
	}

	return aliasName, nil
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
