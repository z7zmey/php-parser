package node_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/node/name"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/kylelemons/godebug/pretty"
	"github.com/z7zmey/php-parser/node"
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

func TestIdentifier(t *testing.T) {
	src := `<? $foo;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Variable{
					VarName: &node.Identifier{Value: "foo"},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestPhp7ArgumentNode(t *testing.T) {
	src := `<? 
		foo($a, ...$b);
		$foo($a, ...$b);
		$foo->bar($a, ...$b);
		foo::bar($a, ...$b);
		$foo::bar($a, ...$b);
		new foo($a, ...$b);
		/** anonymous class */
		new class ($a, ...$b) {};
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.Name{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.MethodCall{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					Method:   &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class: &name.Name{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
					Call:  &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					Call:  &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.New{
					Class: &name.Name{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.New{
					Class: &stmt.Class{
						PhpDocComment: "/** anonymous class */",
						ArgumentList: &node.ArgumentList{
							Arguments: []node.Node{
								&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
								&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
							},
						},
						Stmts: []node.Node{},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestPhp5ArgumentNode(t *testing.T) {
	src := `<? 
		foo($a, ...$b);
		$foo($a, ...$b);
		$foo->bar($a, ...$b);
		foo::bar($a, ...$b);
		$foo::bar($a, ...$b);
		new foo($a, ...$b);
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.Name{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.MethodCall{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					Method:   &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class: &name.Name{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
					Call:  &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					Call:  &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.New{
					Class: &name.Name{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
		},
	}

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual := php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestPhp7ParameterNode(t *testing.T) {
	src := `<? 
		function foo(?bar $bar=null, baz &...$baz) {}
		class foo {public function foo(?bar $bar=null, baz &...$baz) {}}
		function(?bar $bar=null, baz &...$baz) {};
		static function(?bar $bar=null, baz &...$baz) {};
	`

	expectedParams := []node.Node{
		&node.Parameter{
			ByRef:        false,
			Variadic:     false,
			VariableType: &node.Nullable{Expr: &name.Name{Parts: []node.Node{&name.NamePart{Value: "bar"}}}},
			Variable:     &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
			DefaultValue: &expr.ConstFetch{Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "null"}}}},
		},
		&node.Parameter{
			ByRef:        true,
			Variadic:     true,
			VariableType: &name.Name{Parts: []node.Node{&name.NamePart{Value: "baz"}}},
			Variable:     &expr.Variable{VarName: &node.Identifier{Value: "baz"}},
		},
	}

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Function{
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName:  &node.Identifier{Value: "foo"},
				Params:        expectedParams,
				Stmts:         []node.Node{},
			},
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						MethodName: &node.Identifier{Value: "foo"},
						Modifiers:  []node.Node{&node.Identifier{Value: "public"}},
						Params:     expectedParams,
						Stmts:      []node.Node{},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.Closure{
					Params: expectedParams,
					Uses:   []node.Node{},
					Stmts:  []node.Node{},
				},
			},
			&stmt.Expression{
				Expr: &expr.Closure{
					Static: true,
					Params: expectedParams,
					Uses:   []node.Node{},
					Stmts:  []node.Node{},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestPhp5ParameterNode(t *testing.T) {
	src := `<? 
		function foo(bar $bar=null, baz &...$baz) {}
		class foo {public function foo(bar $bar=null, baz &...$baz) {}}
		function(bar $bar=null, baz &...$baz) {};
		static function(bar $bar=null, baz &...$baz) {};
	`

	expectedParams := []node.Node{
		&node.Parameter{
			ByRef:        false,
			Variadic:     false,
			VariableType: &name.Name{Parts: []node.Node{&name.NamePart{Value: "bar"}}},
			Variable:     &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
			DefaultValue: &expr.ConstFetch{Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "null"}}}},
		},
		&node.Parameter{
			ByRef:        true,
			Variadic:     true,
			VariableType: &name.Name{Parts: []node.Node{&name.NamePart{Value: "baz"}}},
			Variable:     &expr.Variable{VarName: &node.Identifier{Value: "baz"}},
		},
	}

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Function{
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName:  &node.Identifier{Value: "foo"},
				Params:        expectedParams,
				Stmts:         []node.Node{},
			},
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						MethodName: &node.Identifier{Value: "foo"},
						Modifiers:  []node.Node{&node.Identifier{Value: "public"}},
						Params:     expectedParams,
						Stmts:      []node.Node{},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.Closure{
					Params: expectedParams,
					Uses:   []node.Node{},
					Stmts:  []node.Node{},
				},
			},
			&stmt.Expression{
				Expr: &expr.Closure{
					Static: true,
					Params: expectedParams,
					Uses:   []node.Node{},
					Stmts:  []node.Node{},
				},
			},
		},
	}

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual := php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestCommentEndFile(t *testing.T) {
	src := `<? //comment at the end)`

	expected := &node.Root{
		Stmts: []node.Node{},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}
