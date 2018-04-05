package cast_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/cast"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
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

func TestArray(t *testing.T) {
	src := `<? (array)$a;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &cast.Array{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestBool(t *testing.T) {
	src := `<? (boolean)$a;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &cast.Bool{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestBoolShort(t *testing.T) {
	src := `<? (bool)$a;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &cast.Bool{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestDouble(t *testing.T) {
	src := `<? (double)$a;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &cast.Double{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestCastFloat(t *testing.T) {
	src := `<? (float)$a;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &cast.Double{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestInt(t *testing.T) {
	src := `<? (integer)$a;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &cast.Int{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestIntShort(t *testing.T) {
	src := `<? (int)$a;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &cast.Int{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestObject(t *testing.T) {
	src := `<? (object)$a;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &cast.Object{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestString(t *testing.T) {
	src := `<? (string)$a;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &cast.String{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestUnset(t *testing.T) {
	src := `<? (unset)$a;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &cast.Unset{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
