package stmt_test

import (
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/stmt"

	"github.com/kylelemons/godebug/pretty"
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
		map[string]interface{}{},
	},
	{
		&stmt.AltElse{
			Stmt: &stmt.StmtList{},
		},
		[]string{"Stmt"},
		map[string]interface{}{},
	},
	{
		&stmt.AltElseIf{
			Cond: &stmt.Expression{},
			Stmt: &stmt.StmtList{},
		},
		[]string{"Cond", "Stmt"},
		map[string]interface{}{},
	},
	{
		&stmt.Break{
			Expr: &stmt.Expression{},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&stmt.Case{
			Cond:  &stmt.Expression{},
			Stmts: []node.Node{&stmt.Expression{}},
		},
		[]string{"Cond", "Stmts"},
		map[string]interface{}{},
	},
	{
		&stmt.Catch{
			Types:    []node.Node{&stmt.Expression{}},
			Variable: &expr.Variable{},
			Stmts:    []node.Node{&stmt.Expression{}},
		},
		[]string{"Types", "Variable", "Stmts"},
		map[string]interface{}{},
	},
	{
		&stmt.ClassConstList{
			Modifiers: []node.Node{&stmt.Expression{}},
			Consts:    []node.Node{&stmt.Expression{}},
		},
		[]string{"Modifiers", "Consts"},
		map[string]interface{}{},
	},
	{
		&stmt.ClassMethod{
			ReturnsRef:    true,
			PhpDocComment: "/** */",
			MethodName:    &node.Identifier{},
			Modifiers:     []node.Node{&stmt.Expression{}},
			Params:        []node.Node{&stmt.Expression{}},
			ReturnType:    &node.Identifier{},
			Stmts:         []node.Node{&stmt.Expression{}},
		},
		[]string{"MethodName", "Modifiers", "Params", "ReturnType", "Stmts"},
		map[string]interface{}{"ReturnsRef": true, "PhpDocComment": "/** */"},
	},
	{
		&stmt.Class{
			PhpDocComment: "/** */",
			ClassName:     &node.Identifier{},
			Modifiers:     []node.Node{&stmt.Expression{}},
			Args:          []node.Node{&stmt.Expression{}},
			Extends:       &node.Identifier{},
			Implements:    []node.Node{&stmt.Expression{}},
			Stmts:         []node.Node{&stmt.Expression{}},
		},
		[]string{"ClassName", "Modifiers", "Args", "Extends", "Implements", "Stmts"},
		map[string]interface{}{"PhpDocComment": "/** */"},
	},
	{
		&stmt.ConstList{
			Consts: []node.Node{&stmt.Expression{}},
		},
		[]string{"Consts"},
		map[string]interface{}{},
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
		map[string]interface{}{},
	},
	{
		&stmt.Declare{
			Consts: []node.Node{&stmt.Expression{}},
			Stmt:   &stmt.StmtList{},
		},
		[]string{"Consts", "Stmt"},
		map[string]interface{}{},
	},
	{
		&stmt.Default{
			Stmts: []node.Node{&stmt.Expression{}},
		},
		[]string{"Stmts"},
		map[string]interface{}{},
	},
	{
		&stmt.Do{
			Stmt: &stmt.StmtList{},
			Cond: &expr.Variable{},
		},
		[]string{"Stmt", "Cond"},
		map[string]interface{}{},
	},
	{
		&stmt.Do{
			Stmt: &stmt.StmtList{},
			Cond: &expr.Variable{},
		},
		[]string{"Stmt", "Cond"},
		map[string]interface{}{},
	},
	{
		&stmt.Echo{
			Exprs: []node.Node{&stmt.Expression{}},
		},
		[]string{"Exprs"},
		map[string]interface{}{},
	},
	{
		&stmt.If{
			Cond:   &stmt.Expression{},
			Stmt:   &stmt.StmtList{},
			ElseIf: []node.Node{&stmt.ElseIf{}},
			Else:   &stmt.Else{},
		},
		[]string{"Cond", "Stmt", "ElseIf", "Else"},
		map[string]interface{}{},
	},
	{
		&stmt.Else{
			Stmt: &stmt.StmtList{},
		},
		[]string{"Stmt"},
		map[string]interface{}{},
	},
	{
		&stmt.ElseIf{
			Cond: &stmt.Expression{},
			Stmt: &stmt.StmtList{},
		},
		[]string{"Cond", "Stmt"},
		map[string]interface{}{},
	},
	{
		&stmt.Expression{
			Expr: &stmt.Expression{},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&stmt.Finally{
			Stmts: []node.Node{&stmt.Expression{}},
		},
		[]string{"Stmts"},
		map[string]interface{}{},
	},
	{
		&stmt.For{
			Init: []node.Node{&stmt.Expression{}},
			Cond: []node.Node{&stmt.Expression{}},
			Loop: []node.Node{&stmt.Expression{}},
			Stmt: &stmt.StmtList{},
		},
		[]string{"Init", "Cond", "Loop", "Stmt"},
		map[string]interface{}{},
	},
	{
		&stmt.AltFor{
			Init: []node.Node{&stmt.Expression{}},
			Cond: []node.Node{&stmt.Expression{}},
			Loop: []node.Node{&stmt.Expression{}},
			Stmt: &stmt.StmtList{},
		},
		[]string{"Init", "Cond", "Loop", "Stmt"},
		map[string]interface{}{},
	},
	{
		&stmt.Foreach{
			ByRef:    true,
			Expr:     &stmt.Expression{},
			Key:      &expr.Variable{},
			Variable: &expr.Variable{},
			Stmt:     &stmt.StmtList{},
		},
		[]string{"Expr", "Key", "Variable", "Stmt"},
		map[string]interface{}{"ByRef": true},
	},
	{
		&stmt.AltForeach{
			ByRef:    true,
			Expr:     &stmt.Expression{},
			Key:      &expr.Variable{},
			Variable: &expr.Variable{},
			Stmt:     &stmt.StmtList{},
		},
		[]string{"Expr", "Key", "Variable", "Stmt"},
		map[string]interface{}{"ByRef": true},
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
		map[string]interface{}{},
	},
	{
		&stmt.Goto{
			Label: &node.Identifier{},
		},
		[]string{"Label"},
		map[string]interface{}{},
	},
	{
		&stmt.GroupUse{
			UseType: &node.Identifier{},
			Prefix:  &node.Identifier{},
			UseList: []node.Node{&stmt.Expression{}},
		},
		[]string{"UseType", "Prefix", "UseList"},
		map[string]interface{}{},
	},
	{
		&stmt.HaltCompiler{},
		[]string{},
		map[string]interface{}{},
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
			Extends:       []node.Node{&stmt.Expression{}},
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
		map[string]interface{}{},
	},
	{
		&stmt.Namespace{
			NamespaceName: &node.Identifier{},
			Stmts:         []node.Node{&stmt.Expression{}},
		},
		[]string{"NamespaceName", "Stmts"},
		map[string]interface{}{},
	},
	{
		&stmt.Nop{},
		[]string{},
		map[string]interface{}{},
	},
	{
		&stmt.PropertyList{
			Modifiers:  []node.Node{&stmt.Expression{}},
			Properties: []node.Node{&stmt.Expression{}},
		},
		[]string{"Modifiers", "Properties"},
		map[string]interface{}{},
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
		map[string]interface{}{},
	},
	{
		&stmt.StaticVar{
			Variable: &expr.Variable{},
			Expr:     &stmt.Expression{},
		},
		[]string{"Variable", "Expr"},
		map[string]interface{}{},
	},
	{
		&stmt.Static{
			Vars: []node.Node{&stmt.Expression{}},
		},
		[]string{"Vars"},
		map[string]interface{}{},
	},
	{
		&stmt.Switch{
			Cond:  &expr.Variable{},
			Cases: []node.Node{&stmt.Expression{}},
		},
		[]string{"Cond", "Cases"},
		map[string]interface{}{},
	},
	{
		&stmt.AltSwitch{
			Cond:  &expr.Variable{},
			Cases: []node.Node{&stmt.Expression{}},
		},
		[]string{"Cond", "Cases"},
		map[string]interface{}{},
	},
	{
		&stmt.Throw{
			Expr: &stmt.Expression{},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&stmt.TraitMethodRef{
			Trait:  &node.Identifier{},
			Method: &node.Identifier{},
		},
		[]string{"Trait", "Method"},
		map[string]interface{}{},
	},
	{
		&stmt.TraitUseAlias{
			Ref:      &node.Identifier{},
			Modifier: &node.Identifier{},
			Alias:    &node.Identifier{},
		},
		[]string{"Ref", "Modifier", "Alias"},
		map[string]interface{}{},
	},
	{
		&stmt.TraitUsePrecedence{
			Ref:       &node.Identifier{},
			Insteadof: []node.Node{&node.Identifier{}},
		},
		[]string{"Ref", "Insteadof"},
		map[string]interface{}{},
	},
	{
		&stmt.TraitUse{
			Traits:      []node.Node{&stmt.Expression{}},
			Adaptations: []node.Node{&stmt.Expression{}},
		},
		[]string{"Traits", "Adaptations"},
		map[string]interface{}{},
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
		map[string]interface{}{},
	},
	{
		&stmt.Unset{
			Vars: []node.Node{&stmt.Expression{}},
		},
		[]string{"Vars"},
		map[string]interface{}{},
	},
	{
		&stmt.UseList{
			UseType: &node.Identifier{},
			Uses:    []node.Node{&stmt.Expression{}},
		},
		[]string{"UseType", "Uses"},
		map[string]interface{}{},
	},
	{
		&stmt.Use{
			UseType: &node.Identifier{},
			Use:     &node.Identifier{},
			Alias:   &node.Identifier{},
		},
		[]string{"UseType", "Use", "Alias"},
		map[string]interface{}{},
	},
	{
		&stmt.While{
			Cond: &expr.Variable{},
			Stmt: &stmt.StmtList{},
		},
		[]string{"Cond", "Stmt"},
		map[string]interface{}{},
	},
	{
		&stmt.AltWhile{
			Cond: &expr.Variable{},
			Stmt: &stmt.StmtList{},
		},
		[]string{"Cond", "Stmt"},
		map[string]interface{}{},
	},
	{
		&stmt.StmtList{
			Stmts: []node.Node{&stmt.Expression{}},
		},
		[]string{"Stmts"},
		map[string]interface{}{},
	},
}

type visitorMock struct {
	visitChildren bool
	visitedKeys   []string
}

func (v *visitorMock) EnterNode(n walker.Walkable) bool { return v.visitChildren }
func (v *visitorMock) GetChildrenVisitor(key string) walker.Visitor {
	v.visitedKeys = append(v.visitedKeys, key)
	return &visitorMock{v.visitChildren, nil}
}
func (v *visitorMock) LeaveNode(n walker.Walkable) {}

func TestVisitorDisableChildren(t *testing.T) {
	for _, tt := range nodesToTest {
		v := &visitorMock{false, nil}
		tt.node.Walk(v)

		expected := []string{}
		actual := v.visitedKeys

		diff := pretty.Compare(expected, actual)
		if diff != "" {
			t.Errorf("%s diff: (-expected +actual)\n%s", reflect.TypeOf(tt.node), diff)
		}
	}
}

func TestVisitor(t *testing.T) {
	for _, tt := range nodesToTest {
		v := &visitorMock{true, nil}
		tt.node.Walk(v)

		expected := tt.expectedVisitedKeys
		actual := v.visitedKeys

		diff := pretty.Compare(expected, actual)
		if diff != "" {
			t.Errorf("%s diff: (-expected +actual)\n%s", reflect.TypeOf(tt.node), diff)
		}
	}
}

// test Attributes()

func TestNameAttributes(t *testing.T) {
	for _, tt := range nodesToTest {
		expected := tt.expectedAttributes
		actual := tt.node.Attributes()

		diff := pretty.Compare(expected, actual)
		if diff != "" {
			t.Errorf("%s diff: (-expected +actual)\n%s", reflect.TypeOf(tt.node), diff)
		}
	}
}
