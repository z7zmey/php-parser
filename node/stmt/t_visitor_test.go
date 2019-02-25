package stmt_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node/name"

	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/stmt"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

var nodesToTest = []struct {
	node                node.Node // node
	expectedVisitedKeys []string  // visited keys
	expectedAttributes  map[string]interface{}
}{
	{
		&stmt.AltIf{
			Cond:   &stmt.Expression{},
			Stmt:   &stmt.StmtList{},
			ElseIf: []node.Node{&stmt.ElseIf{}},
			Else:   &stmt.Else{},
		},
		[]string{"Cond", "Stmt", "ElseIf", "Else"},
		nil,
	},
	{
		&stmt.AltElse{
			Stmt: &stmt.StmtList{},
		},
		[]string{"Stmt"},
		nil,
	},
	{
		&stmt.AltElseIf{
			Cond: &stmt.Expression{},
			Stmt: &stmt.StmtList{},
		},
		[]string{"Cond", "Stmt"},
		nil,
	},
	{
		&stmt.Break{
			Expr: &stmt.Expression{},
		},
		[]string{"Expr"},
		nil,
	},
	{
		&stmt.Case{
			Cond:  &stmt.Expression{},
			Stmts: []node.Node{&stmt.Expression{}},
		},
		[]string{"Cond", "Stmts"},
		nil,
	},
	{
		&stmt.Catch{
			Types:    []node.Node{&stmt.Expression{}},
			Variable: &expr.Variable{},
			Stmts:    []node.Node{&stmt.Expression{}},
		},
		[]string{"Types", "Variable", "Stmts"},
		nil,
	},
	{
		&stmt.ClassConstList{
			Modifiers: []node.Node{&stmt.Expression{}},
			Consts:    []node.Node{&stmt.Expression{}},
		},
		[]string{"Modifiers", "Consts"},
		nil,
	},
	{
		&stmt.ClassMethod{
			ReturnsRef:    true,
			PhpDocComment: "/** */",
			MethodName:    &node.Identifier{},
			Modifiers:     []node.Node{&stmt.Expression{}},
			Params:        []node.Node{&stmt.Expression{}},
			ReturnType:    &node.Identifier{},
			Stmt:          &stmt.StmtList{},
		},
		[]string{"MethodName", "Modifiers", "Params", "ReturnType", "Stmt"},
		map[string]interface{}{"ReturnsRef": true, "PhpDocComment": "/** */"},
	},
	{
		&stmt.Class{
			PhpDocComment: "/** */",
			ClassName:     &node.Identifier{},
			Modifiers:     []node.Node{&stmt.Expression{}},
			ArgumentList:  &node.ArgumentList{},
			Extends:       &stmt.ClassExtends{},
			Implements:    &stmt.ClassImplements{},
			Stmts:         []node.Node{&stmt.Expression{}},
		},
		[]string{"ClassName", "Modifiers", "ArgumentList", "Extends", "Implements", "Stmts"},
		map[string]interface{}{"PhpDocComment": "/** */"},
	},
	{
		&stmt.ConstList{
			Consts: []node.Node{&stmt.Expression{}},
		},
		[]string{"Consts"},
		nil,
	},
	{
		&stmt.Constant{
			PhpDocComment: "/** */",
			ConstantName:  &node.Identifier{},
			Expr:          &stmt.Expression{},
		},
		[]string{"ConstantName", "Expr"},
		map[string]interface{}{"PhpDocComment": "/** */"},
	},
	{
		&stmt.Continue{
			Expr: &stmt.Expression{},
		},
		[]string{"Expr"},
		nil,
	},
	{
		&stmt.Declare{
			Consts: []node.Node{&stmt.Expression{}},
			Stmt:   &stmt.StmtList{},
		},
		[]string{"Consts", "Stmt"},
		nil,
	},
	{
		&stmt.Default{
			Stmts: []node.Node{&stmt.Expression{}},
		},
		[]string{"Stmts"},
		nil,
	},
	{
		&stmt.Do{
			Stmt: &stmt.StmtList{},
			Cond: &expr.Variable{},
		},
		[]string{"Stmt", "Cond"},
		nil,
	},
	{
		&stmt.Do{
			Stmt: &stmt.StmtList{},
			Cond: &expr.Variable{},
		},
		[]string{"Stmt", "Cond"},
		nil,
	},
	{
		&stmt.Echo{
			Exprs: []node.Node{&stmt.Expression{}},
		},
		[]string{"Exprs"},
		nil,
	},
	{
		&stmt.If{
			Cond:   &stmt.Expression{},
			Stmt:   &stmt.StmtList{},
			ElseIf: []node.Node{&stmt.ElseIf{}},
			Else:   &stmt.Else{},
		},
		[]string{"Cond", "Stmt", "ElseIf", "Else"},
		nil,
	},
	{
		&stmt.Else{
			Stmt: &stmt.StmtList{},
		},
		[]string{"Stmt"},
		nil,
	},
	{
		&stmt.ElseIf{
			Cond: &stmt.Expression{},
			Stmt: &stmt.StmtList{},
		},
		[]string{"Cond", "Stmt"},
		nil,
	},
	{
		&stmt.Expression{
			Expr: &stmt.Expression{},
		},
		[]string{"Expr"},
		nil,
	},
	{
		&stmt.Finally{
			Stmts: []node.Node{&stmt.Expression{}},
		},
		[]string{"Stmts"},
		nil,
	},
	{
		&stmt.For{
			Init: []node.Node{&stmt.Expression{}},
			Cond: []node.Node{&stmt.Expression{}},
			Loop: []node.Node{&stmt.Expression{}},
			Stmt: &stmt.StmtList{},
		},
		[]string{"Init", "Cond", "Loop", "Stmt"},
		nil,
	},
	{
		&stmt.AltFor{
			Init: []node.Node{&stmt.Expression{}},
			Cond: []node.Node{&stmt.Expression{}},
			Loop: []node.Node{&stmt.Expression{}},
			Stmt: &stmt.StmtList{},
		},
		[]string{"Init", "Cond", "Loop", "Stmt"},
		nil,
	},
	{
		&stmt.Foreach{
			Expr:     &stmt.Expression{},
			Key:      &expr.Variable{},
			Variable: &expr.Variable{},
			Stmt:     &stmt.StmtList{},
		},
		[]string{"Expr", "Key", "Variable", "Stmt"},
		nil,
	},
	{
		&stmt.AltForeach{
			Expr:     &stmt.Expression{},
			Key:      &expr.Variable{},
			Variable: &expr.Variable{},
			Stmt:     &stmt.StmtList{},
		},
		[]string{"Expr", "Key", "Variable", "Stmt"},
		nil,
	},
	{
		&stmt.Function{
			ReturnsRef:    true,
			PhpDocComment: "/** */",
			FunctionName:  &node.Identifier{},
			Params:        []node.Node{&stmt.Expression{}},
			ReturnType:    &node.Identifier{},
			Stmts:         []node.Node{&stmt.Expression{}},
		},
		[]string{"FunctionName", "Params", "ReturnType", "Stmts"},
		map[string]interface{}{"ReturnsRef": true, "PhpDocComment": "/** */"},
	},
	{
		&stmt.Global{
			Vars: []node.Node{&stmt.Expression{}},
		},
		[]string{"Vars"},
		nil,
	},
	{
		&stmt.Goto{
			Label: &node.Identifier{},
		},
		[]string{"Label"},
		nil,
	},
	{
		&stmt.GroupUse{
			UseType: &node.Identifier{},
			Prefix:  &node.Identifier{},
			UseList: []node.Node{&stmt.Expression{}},
		},
		[]string{"UseType", "Prefix", "UseList"},
		nil,
	},
	{
		&stmt.HaltCompiler{},
		[]string{},
		nil,
	},
	{
		&stmt.InlineHtml{
			Value: "hello",
		},
		[]string{},
		map[string]interface{}{"Value": "hello"},
	},
	{
		&stmt.Interface{
			PhpDocComment: "/** */",
			InterfaceName: &node.Identifier{},
			Extends:       &stmt.InterfaceExtends{},
			Stmts:         []node.Node{&stmt.Expression{}},
		},
		[]string{"InterfaceName", "Extends", "Stmts"},
		map[string]interface{}{"PhpDocComment": "/** */"},
	},
	{
		&stmt.Label{
			LabelName: &node.Identifier{},
		},
		[]string{"LabelName"},
		nil,
	},
	{
		&stmt.Namespace{
			NamespaceName: &node.Identifier{},
			Stmts:         []node.Node{&stmt.Expression{}},
		},
		[]string{"NamespaceName", "Stmts"},
		nil,
	},
	{
		&stmt.Nop{},
		[]string{},
		nil,
	},
	{
		&stmt.PropertyList{
			Modifiers:  []node.Node{&stmt.Expression{}},
			Properties: []node.Node{&stmt.Expression{}},
		},
		[]string{"Modifiers", "Properties"},
		nil,
	},
	{
		&stmt.Property{
			PhpDocComment: "/** */",
			Variable:      &expr.Variable{},
			Expr:          &stmt.Expression{},
		},
		[]string{"Variable", "Expr"},
		map[string]interface{}{"PhpDocComment": "/** */"},
	},
	{
		&stmt.Return{
			Expr: &stmt.Expression{},
		},
		[]string{"Expr"},
		nil,
	},
	{
		&stmt.StaticVar{
			Variable: &expr.Variable{},
			Expr:     &stmt.Expression{},
		},
		[]string{"Variable", "Expr"},
		nil,
	},
	{
		&stmt.Static{
			Vars: []node.Node{&stmt.Expression{}},
		},
		[]string{"Vars"},
		nil,
	},
	{
		&stmt.Switch{
			Cond:     &expr.Variable{},
			CaseList: &stmt.CaseList{},
		},
		[]string{"Cond", "CaseList"},
		nil,
	},
	{
		&stmt.AltSwitch{
			Cond:     &expr.Variable{},
			CaseList: &stmt.CaseList{},
		},
		[]string{"Cond", "CaseList"},
		nil,
	},
	{
		&stmt.Throw{
			Expr: &stmt.Expression{},
		},
		[]string{"Expr"},
		nil,
	},
	{
		&stmt.TraitMethodRef{
			Trait:  &node.Identifier{},
			Method: &node.Identifier{},
		},
		[]string{"Trait", "Method"},
		nil,
	},
	{
		&stmt.TraitUseAlias{
			Ref:      &node.Identifier{},
			Modifier: &node.Identifier{},
			Alias:    &node.Identifier{},
		},
		[]string{"Ref", "Modifier", "Alias"},
		nil,
	},
	{
		&stmt.TraitUsePrecedence{
			Ref:       &node.Identifier{},
			Insteadof: []node.Node{&node.Identifier{}},
		},
		[]string{"Ref", "Insteadof"},
		nil,
	},
	{
		&stmt.TraitUse{
			Traits:              []node.Node{&stmt.Expression{}},
			TraitAdaptationList: &stmt.TraitAdaptationList{},
		},
		[]string{"Traits", "TraitAdaptationList"},
		nil,
	},
	{
		&stmt.Trait{
			PhpDocComment: "/** */",
			TraitName:     &node.Identifier{},
			Stmts:         []node.Node{&stmt.Expression{}},
		},
		[]string{"TraitName", "Stmts"},
		map[string]interface{}{"PhpDocComment": "/** */"},
	},
	{
		&stmt.Try{
			Stmts:   []node.Node{&stmt.Expression{}},
			Catches: []node.Node{&stmt.Expression{}},
			Finally: &stmt.Finally{},
		},
		[]string{"Stmts", "Catches", "Finally"},
		nil,
	},
	{
		&stmt.Unset{
			Vars: []node.Node{&stmt.Expression{}},
		},
		[]string{"Vars"},
		nil,
	},
	{
		&stmt.UseList{
			UseType: &node.Identifier{},
			Uses:    []node.Node{&stmt.Expression{}},
		},
		[]string{"UseType", "Uses"},
		nil,
	},
	{
		&stmt.Use{
			UseType: &node.Identifier{},
			Use:     &node.Identifier{},
			Alias:   &node.Identifier{},
		},
		[]string{"UseType", "Use", "Alias"},
		nil,
	},
	{
		&stmt.While{
			Cond: &expr.Variable{},
			Stmt: &stmt.StmtList{},
		},
		[]string{"Cond", "Stmt"},
		nil,
	},
	{
		&stmt.AltWhile{
			Cond: &expr.Variable{},
			Stmt: &stmt.StmtList{},
		},
		[]string{"Cond", "Stmt"},
		nil,
	},
	{
		&stmt.StmtList{
			Stmts: []node.Node{&stmt.Expression{}},
		},
		[]string{"Stmts"},
		nil,
	},
	{
		&stmt.CaseList{
			Cases: []node.Node{&stmt.Expression{}},
		},
		[]string{"Cases"},
		nil,
	},
	{
		&stmt.TraitAdaptationList{
			Adaptations: []node.Node{&stmt.TraitUsePrecedence{}},
		},
		[]string{"Adaptations"},
		nil,
	},
	{
		&stmt.ClassExtends{
			ClassName: &name.Name{},
		},
		[]string{"ClassName"},
		nil,
	},
	{
		&stmt.ClassImplements{
			InterfaceNames: []node.Node{
				&name.Name{},
			},
		},
		[]string{"InterfaceNames"},
		nil,
	},
	{
		&stmt.InterfaceExtends{
			InterfaceNames: []node.Node{
				&name.Name{},
			},
		},
		[]string{"InterfaceNames"},
		nil,
	},
}

type visitorMock struct {
	visitChildren bool
	visitedKeys   []string
}

func (v *visitorMock) EnterNode(n walker.Walkable) bool { return v.visitChildren }
func (v *visitorMock) LeaveNode(n walker.Walkable)      {}
func (v *visitorMock) EnterChildNode(key string, w walker.Walkable) {
	v.visitedKeys = append(v.visitedKeys, key)
}
func (v *visitorMock) LeaveChildNode(key string, w walker.Walkable) {}
func (v *visitorMock) EnterChildList(key string, w walker.Walkable) {
	v.visitedKeys = append(v.visitedKeys, key)
}
func (v *visitorMock) LeaveChildList(key string, w walker.Walkable) {}

func TestVisitorDisableChildren(t *testing.T) {
	for _, tt := range nodesToTest {
		v := &visitorMock{false, []string{}}
		tt.node.Walk(v)

		expected := []string{}
		actual := v.visitedKeys

		assert.DeepEqual(t, expected, actual)
	}
}

func TestVisitor(t *testing.T) {
	for _, tt := range nodesToTest {
		v := &visitorMock{true, []string{}}
		tt.node.Walk(v)

		expected := tt.expectedVisitedKeys
		actual := v.visitedKeys

		assert.DeepEqual(t, expected, actual)
	}
}

// test Attributes()

func TestNameAttributes(t *testing.T) {
	for _, tt := range nodesToTest {
		expected := tt.expectedAttributes
		actual := tt.node.Attributes()

		assert.DeepEqual(t, expected, actual)
	}
}
