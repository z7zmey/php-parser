package test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/kylelemons/godebug/pretty"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/parser"
)

func TestSimpleVar(t *testing.T) {
	src := `<? "test $var";`

	varName := node.NewIdentifier("$var").SetPosition(&node.Position{1, 1, 10, 13})
	parts := []node.Node{
		scalar.NewEncapsedStringPart("test ").SetPosition(&node.Position{1, 1, 5, 9}),
		expr.NewVariable(varName).SetPosition(&node.Position{1, 1, 10, 13}),
	}
	encapsed := scalar.NewEncapsed(parts).SetPosition(&node.Position{1, 1, 4, 14})
	expr := stmt.NewExpression(encapsed).SetPosition(&node.Position{1, 1, 4, 15})
	expected := stmt.NewStmtList([]node.Node{expr}).SetPosition(&node.Position{1, 1, 4, 15})

	actual, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestSimpleVarPropertyFetch(t *testing.T) {
	src := `<? "test $foo->bar()";`

	varName := node.NewIdentifier("$foo").SetPosition(&node.Position{1, 1, 10, 13})
	variable := expr.NewVariable(varName).SetPosition(&node.Position{1, 1, 10, 13})

	property := node.NewIdentifier("bar").SetPosition(&node.Position{1, 1, 16, 18})
	propertyFetch := expr.NewPropertyFetch(variable, property).SetPosition(&node.Position{1, 1, 10, 18})

	parts := []node.Node{
		scalar.NewEncapsedStringPart("test ").SetPosition(&node.Position{1, 1, 5, 9}),
		propertyFetch,
		scalar.NewEncapsedStringPart("()").SetPosition(&node.Position{1, 1, 19, 20}),
	}
	encapsed := scalar.NewEncapsed(parts).SetPosition(&node.Position{1, 1, 4, 21})
	expr := stmt.NewExpression(encapsed).SetPosition(&node.Position{1, 1, 4, 22})
	expected := stmt.NewStmtList([]node.Node{expr}).SetPosition(&node.Position{1, 1, 4, 22})

	actual, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestDollarOpenCurlyBraces(t *testing.T) {
	src := `<? "test ${foo}";`

	varName := node.NewIdentifier("foo").SetPosition(&node.Position{1, 1, 12, 14})
	variable := expr.NewVariable(varName).SetPosition(&node.Position{1, 1, 10, 15})
	parts := []node.Node{
		scalar.NewEncapsedStringPart("test ").SetPosition(&node.Position{1, 1, 5, 9}),
		variable,
	}
	encapsed := scalar.NewEncapsed(parts).SetPosition(&node.Position{1, 1, 4, 16})
	expr := stmt.NewExpression(encapsed).SetPosition(&node.Position{1, 1, 4, 17})
	expected := stmt.NewStmtList([]node.Node{expr}).SetPosition(&node.Position{1, 1, 4, 17})

	actual, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestDollarOpenCurlyBracesDimNumber(t *testing.T) {
	src := `<? "test ${foo[0]}";`

	varName := node.NewIdentifier("foo").SetPosition(&node.Position{1, 1, 12, 14})
	variable := expr.NewVariable(varName).SetPosition(&node.Position{1, 1, 12, 14})

	lNumber := scalar.NewLnumber("0").SetPosition(&node.Position{1, 1, 16, 16})

	arrayDimFetch := expr.NewArrayDimFetch(variable, lNumber).SetPosition(&node.Position{1, 1, 10, 18})

	parts := []node.Node{
		scalar.NewEncapsedStringPart("test ").SetPosition(&node.Position{1, 1, 5, 9}),
		arrayDimFetch,
	}
	encapsed := scalar.NewEncapsed(parts).SetPosition(&node.Position{1, 1, 4, 19})
	expr := stmt.NewExpression(encapsed).SetPosition(&node.Position{1, 1, 4, 20})
	expected := stmt.NewStmtList([]node.Node{expr}).SetPosition(&node.Position{1, 1, 4, 20})

	actual, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}

func TestCurlyOpenMethodCall(t *testing.T) {
	src := `<? "test {$foo->bar()}";`

	varName := node.NewIdentifier("$foo").SetPosition(&node.Position{1, 1, 11, 14})
	variable := expr.NewVariable(varName).SetPosition(&node.Position{1, 1, 11, 14})

	method := node.NewIdentifier("bar").SetPosition(&node.Position{1, 1, 17, 19})

	methodCall := expr.NewMethodCall(variable, method, nil).SetPosition(&node.Position{1, 1, 11, 21})

	parts := []node.Node{
		scalar.NewEncapsedStringPart("test ").SetPosition(&node.Position{1, 1, 5, 9}),
		methodCall,
	}
	encapsed := scalar.NewEncapsed(parts).SetPosition(&node.Position{1, 1, 4, 23})
	expr := stmt.NewExpression(encapsed).SetPosition(&node.Position{1, 1, 4, 24})
	expected := stmt.NewStmtList([]node.Node{expr}).SetPosition(&node.Position{1, 1, 4, 24})

	actual, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	if diff := pretty.Compare(expected, actual); diff != "" {
		t.Errorf("diff: (-expected +actual)\n%s", diff)
	}
}
