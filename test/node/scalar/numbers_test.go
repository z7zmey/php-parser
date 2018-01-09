package test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/parser"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		diff := pretty.Compare(expected, actual)

		if diff != "" {
			t.Errorf("diff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("expected and actual are not equal\n")
		}

	}
}

func TestLNumber(t *testing.T) {
	src := `<? 1234567890123456789;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Lnumber{Value: "1234567890123456789"},
			},
		},
	}

	actual, _, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	assertEqual(t, expected, actual)
}

func TestDNumber(t *testing.T) {
	src := `<? 12345678901234567890;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Dnumber{Value: "12345678901234567890"},
			},
		},
	}

	actual, _, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	assertEqual(t, expected, actual)
}

func TestFloat(t *testing.T) {
	src := `<? 0.;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Dnumber{Value: "0."},
			},
		},
	}

	actual, _, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	assertEqual(t, expected, actual)
}

func TestBinaryLNumber(t *testing.T) {
	src := `<? 0b0111111111111111111111111111111111111111111111111111111111111111;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Lnumber{Value: "0b0111111111111111111111111111111111111111111111111111111111111111"},
			},
		},
	}

	actual, _, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	assertEqual(t, expected, actual)
}

func TestBinaryDNumber(t *testing.T) {
	src := `<? 0b1111111111111111111111111111111111111111111111111111111111111111;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Dnumber{Value: "0b1111111111111111111111111111111111111111111111111111111111111111"},
			},
		},
	}

	actual, _, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	assertEqual(t, expected, actual)
}

func TestHLNumber(t *testing.T) {
	src := `<? 0x007111111111111111;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Lnumber{Value: "0x007111111111111111"},
			},
		},
	}

	actual, _, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	assertEqual(t, expected, actual)
}

func TestHDNumber(t *testing.T) {
	src := `<? 0x8111111111111111;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Dnumber{Value: "0x8111111111111111"},
			},
		},
	}

	actual, _, _ := parser.Parse(bytes.NewBufferString(src), "test.php")

	assertEqual(t, expected, actual)
}
