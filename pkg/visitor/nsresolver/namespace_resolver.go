// Package visitor contains walker.visitor implementations
package nsresolver

import (
	"errors"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/visitor"
	"strings"
)

// NamespaceResolver visitor
type NamespaceResolver struct {
	visitor.Null
	Namespace     *Namespace
	ResolvedNames map[ast.Vertex]string

	goDeep bool
}

// NewNamespaceResolver NamespaceResolver type constructor
func NewNamespaceResolver() *NamespaceResolver {
	return &NamespaceResolver{
		Namespace:     NewNamespace(""),
		ResolvedNames: map[ast.Vertex]string{},
		goDeep:        true,
	}
}

func (nsr *NamespaceResolver) EnterNode(n ast.Vertex) bool {
	n.Accept(nsr)

	if !nsr.goDeep {
		nsr.goDeep = true
		return false
	}

	return true
}

func (nsr *NamespaceResolver) StmtNamespace(n *ast.StmtNamespace) {
	if n.Name == nil {
		nsr.Namespace = NewNamespace("")
	} else {
		NSParts := n.Name.(*ast.NameName).Parts
		nsr.Namespace = NewNamespace(concatNameParts(NSParts))
	}
}

func (nsr *NamespaceResolver) StmtUse(n *ast.StmtUse) {
	useType := ""
	if n.Type != nil {
		useType = string(n.Type.(*ast.Identifier).Value)
	}

	for _, nn := range n.UseDeclarations {
		nsr.AddAlias(useType, nn, nil)
	}

	nsr.goDeep = false
}

func (nsr *NamespaceResolver) StmtGroupUse(n *ast.StmtGroupUse) {
	useType := ""
	if n.Type != nil {
		useType = string(n.Type.(*ast.Identifier).Value)
	}

	for _, nn := range n.UseDeclarations {
		nsr.AddAlias(useType, nn, n.Prefix.(*ast.NameName).Parts)
	}

	nsr.goDeep = false
}

func (nsr *NamespaceResolver) StmtClass(n *ast.StmtClass) {
	if n.Extends != nil {
		nsr.ResolveName(n.Extends, "")
	}

	if n.Implements != nil {
		for _, interfaceName := range n.Implements {
			nsr.ResolveName(interfaceName, "")
		}
	}

	if n.ClassName != nil {
		nsr.AddNamespacedName(n, string(n.ClassName.(*ast.Identifier).Value))
	}
}

func (nsr *NamespaceResolver) StmtInterface(n *ast.StmtInterface) {
	if n.Extends != nil {
		for _, interfaceName := range n.Extends {
			nsr.ResolveName(interfaceName, "")
		}
	}

	nsr.AddNamespacedName(n, string(n.InterfaceName.(*ast.Identifier).Value))
}

func (nsr *NamespaceResolver) StmtTrait(n *ast.StmtTrait) {
	nsr.AddNamespacedName(n, string(n.TraitName.(*ast.Identifier).Value))
}

func (nsr *NamespaceResolver) StmtFunction(n *ast.StmtFunction) {
	nsr.AddNamespacedName(n, string(n.FunctionName.(*ast.Identifier).Value))

	for _, parameter := range n.Params {
		nsr.ResolveType(parameter.(*ast.Parameter).Type)
	}

	if n.ReturnType != nil {
		nsr.ResolveType(n.ReturnType)
	}
}

func (nsr *NamespaceResolver) StmtClassMethod(n *ast.StmtClassMethod) {
	for _, parameter := range n.Params {
		nsr.ResolveType(parameter.(*ast.Parameter).Type)
	}

	if n.ReturnType != nil {
		nsr.ResolveType(n.ReturnType)
	}
}

func (nsr *NamespaceResolver) ExprClosure(n *ast.ExprClosure) {
	for _, parameter := range n.Params {
		nsr.ResolveType(parameter.(*ast.Parameter).Type)
	}

	if n.ReturnType != nil {
		nsr.ResolveType(n.ReturnType)
	}
}

func (nsr *NamespaceResolver) StmtPropertyList(n *ast.StmtPropertyList) {
	if n.Type != nil {
		nsr.ResolveType(n.Type)
	}
}

func (nsr *NamespaceResolver) StmtConstList(n *ast.StmtConstList) {
	for _, constant := range n.Consts {
		nsr.AddNamespacedName(constant, string(constant.(*ast.StmtConstant).Name.(*ast.Identifier).Value))
	}
}

func (nsr *NamespaceResolver) ExprStaticCall(n *ast.ExprStaticCall) {
	nsr.ResolveName(n.Class, "")
}

func (nsr *NamespaceResolver) ExprStaticPropertyFetch(n *ast.ExprStaticPropertyFetch) {
	nsr.ResolveName(n.Class, "")
}

func (nsr *NamespaceResolver) ExprClassConstFetch(n *ast.ExprClassConstFetch) {
	nsr.ResolveName(n.Class, "")
}

func (nsr *NamespaceResolver) ExprNew(n *ast.ExprNew) {
	nsr.ResolveName(n.Class, "")
}

func (nsr *NamespaceResolver) ExprInstanceOf(n *ast.ExprInstanceOf) {
	nsr.ResolveName(n.Class, "")
}

func (nsr *NamespaceResolver) StmtCatch(n *ast.StmtCatch) {
	for _, t := range n.Types {
		nsr.ResolveName(t, "")
	}
}

func (nsr *NamespaceResolver) ExprFunctionCall(n *ast.ExprFunctionCall) {
	nsr.ResolveName(n.Function, "function")
}

func (nsr *NamespaceResolver) ExprConstFetch(n *ast.ExprConstFetch) {
	nsr.ResolveName(n.Const, "const")
}

func (nsr *NamespaceResolver) StmtTraitUse(n *ast.StmtTraitUse) {
	for _, t := range n.Traits {
		nsr.ResolveName(t, "")
	}

	for _, a := range n.Adaptations {
		switch aa := a.(type) {
		case *ast.StmtTraitUsePrecedence:
			refTrait := aa.Trait
			if refTrait != nil {
				nsr.ResolveName(refTrait, "")
			}
			for _, insteadOf := range aa.Insteadof {
				nsr.ResolveName(insteadOf, "")
			}

		case *ast.StmtTraitUseAlias:
			refTrait := aa.Trait
			if refTrait != nil {
				nsr.ResolveName(refTrait, "")
			}
		}
	}
}

// LeaveNode is invoked after node process
func (nsr *NamespaceResolver) LeaveNode(n ast.Vertex) {
	switch nn := n.(type) {
	case *ast.StmtNamespace:
		if nn.Stmts != nil {
			nsr.Namespace = NewNamespace("")
		}
	}
}

// AddAlias adds a new alias
func (nsr *NamespaceResolver) AddAlias(useType string, nn ast.Vertex, prefix []ast.Vertex) {
	switch use := nn.(type) {
	case *ast.StmtUseDeclaration:
		if use.Type != nil {
			useType = string(use.Type.(*ast.Identifier).Value)
		}

		useNameParts := use.Use.(*ast.NameName).Parts
		var alias string
		if use.Alias == nil {
			alias = string(useNameParts[len(useNameParts)-1].(*ast.NameNamePart).Value)
		} else {
			alias = string(use.Alias.(*ast.Identifier).Value)
		}

		nsr.Namespace.AddAlias(useType, concatNameParts(prefix, useNameParts), alias)
	}
}

// AddNamespacedName adds namespaced name by node
func (nsr *NamespaceResolver) AddNamespacedName(nn ast.Vertex, nodeName string) {
	if nsr.Namespace.Namespace == "" {
		nsr.ResolvedNames[nn] = nodeName
	} else {
		nsr.ResolvedNames[nn] = nsr.Namespace.Namespace + "\\" + nodeName
	}
}

// ResolveName adds a resolved fully qualified name by node
func (nsr *NamespaceResolver) ResolveName(nameNode ast.Vertex, aliasType string) {
	resolved, err := nsr.Namespace.ResolveName(nameNode, aliasType)
	if err == nil {
		nsr.ResolvedNames[nameNode] = resolved
	}
}

// ResolveType adds a resolved fully qualified type name
func (nsr *NamespaceResolver) ResolveType(n ast.Vertex) {
	switch nn := n.(type) {
	case *ast.Nullable:
		nsr.ResolveType(nn.Expr)
	case *ast.NameName:
		nsr.ResolveName(n, "")
	case *ast.NameRelative:
		nsr.ResolveName(n, "")
	case *ast.NameFullyQualified:
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
func (ns *Namespace) ResolveName(nameNode ast.Vertex, aliasType string) (string, error) {
	switch n := nameNode.(type) {
	case *ast.NameFullyQualified:
		// Fully qualifid name is already resolved
		return concatNameParts(n.Parts), nil

	case *ast.NameRelative:
		if ns.Namespace == "" {
			return concatNameParts(n.Parts), nil
		}
		return ns.Namespace + "\\" + concatNameParts(n.Parts), nil

	case *ast.NameName:
		if aliasType == "const" && len(n.Parts) == 1 {
			part := strings.ToLower(string(n.Parts[0].(*ast.NameNamePart).Value))
			if part == "true" || part == "false" || part == "null" {
				return part, nil
			}
		}

		if aliasType == "" && len(n.Parts) == 1 {
			part := strings.ToLower(string(n.Parts[0].(*ast.NameNamePart).Value))

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
func (ns *Namespace) ResolveAlias(nameNode ast.Vertex, aliasType string) (string, error) {
	aliasType = strings.ToLower(aliasType)
	nameParts := nameNode.(*ast.NameName).Parts

	firstPartStr := string(nameParts[0].(*ast.NameNamePart).Value)

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

func concatNameParts(parts ...[]ast.Vertex) string {
	str := ""

	for _, p := range parts {
		for _, n := range p {
			if str == "" {
				str = string(n.(*ast.NameNamePart).Value)
			} else {
				str = str + "\\" + string(n.(*ast.NameNamePart).Value)
			}
		}
	}

	return str
}
