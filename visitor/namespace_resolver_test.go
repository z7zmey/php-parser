// Package visitor contains walker.visitor implementations
package visitor_test

import (
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/z7zmey/php-parser/node/scalar"

	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/visitor"

	"github.com/z7zmey/php-parser/node/stmt"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/name"
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

func TestResolveStaticCall(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameBC := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}, &name.NamePart{Value: "C"}}}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: nameAB,
					},
				},
			},
			&expr.StaticCall{
				Class:     nameBC,
				Call:      &node.Identifier{Value: "foo"},
				Arguments: []node.Node{},
			},
		},
	}

	expected := map[node.Node]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveStaticPropertyFetch(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameBC := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}, &name.NamePart{Value: "C"}}}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: nameAB,
					},
				},
			},
			&expr.StaticPropertyFetch{
				Class:    nameBC,
				Property: &node.Identifier{Value: "$foo"},
			},
		},
	}

	expected := map[node.Node]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveClassConstFetch(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameBC := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}, &name.NamePart{Value: "C"}}}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: nameAB,
					},
				},
			},
			&expr.ClassConstFetch{
				Class:        nameBC,
				ConstantName: &node.Identifier{Value: "FOO"},
			},
		},
	}

	expected := map[node.Node]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveNew(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameBC := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}, &name.NamePart{Value: "C"}}}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: nameAB,
					},
				},
			},
			&expr.New{
				Class:     nameBC,
				Arguments: []node.Node{},
			},
		},
	}

	expected := map[node.Node]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveInstanceOf(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameBC := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}, &name.NamePart{Value: "C"}}}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: nameAB,
					},
				},
			},
			&expr.InstanceOf{
				Expr:  &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
				Class: nameBC,
			},
		},
	}

	expected := map[node.Node]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveInstanceCatch(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameBC := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}, &name.NamePart{Value: "C"}}}

	nameDE := &name.Name{Parts: []node.Node{&name.NamePart{Value: "D"}, &name.NamePart{Value: "E"}}}
	nameF := &name.Name{Parts: []node.Node{&name.NamePart{Value: "F"}}}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: nameAB,
					},
					&stmt.Use{
						Use:   nameDE,
						Alias: &node.Identifier{Value: "F"},
					},
				},
			},
			&stmt.Try{
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							nameBC,
							nameF,
						},
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
						Stmts:    []node.Node{},
					},
				},
			},
		},
	}

	expected := map[node.Node]string{
		nameBC: "A\\B\\C",
		nameF:  "D\\E",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveFunctionCall(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}}}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.UseList{
				UseType: &node.Identifier{Value: "function"},
				Uses: []node.Node{
					&stmt.Use{
						Use: nameAB,
					},
				},
			},
			&expr.FunctionCall{
				Function:  nameB,
				Arguments: []node.Node{},
			},
		},
	}

	expected := map[node.Node]string{
		nameB: "A\\B",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveConstFetch(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}}}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.UseList{
				UseType: &node.Identifier{Value: "const"},
				Uses: []node.Node{
					&stmt.Use{
						Use: nameAB,
					},
				},
			},
			&expr.ConstFetch{
				Constant: nameB,
			},
		},
	}

	expected := map[node.Node]string{
		nameB: "A\\B",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveGroupUse(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameBD := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}, &name.NamePart{Value: "D"}}}
	nameE := &name.Name{Parts: []node.Node{&name.NamePart{Value: "E"}}}
	nameC := &name.Name{Parts: []node.Node{&name.NamePart{Value: "C"}}}
	nameF := &name.Name{Parts: []node.Node{&name.NamePart{Value: "F"}}}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.GroupUse{
				Prefix: nameAB,
				UseList: []node.Node{
					&stmt.Use{
						UseType: &node.Identifier{Value: "Function"},
						Use:     nameF,
					},
					&stmt.Use{
						UseType: &node.Identifier{Value: "const"},
						Use:     nameC,
					},
				},
			},
			&stmt.GroupUse{
				Prefix:  nameBD,
				UseType: &node.Identifier{Value: "Function"},
				UseList: []node.Node{
					&stmt.Use{
						Use: nameE,
					},
				},
			},
			&expr.ConstFetch{
				Constant: nameC,
			},
			&expr.FunctionCall{
				Function:  nameF,
				Arguments: []node.Node{},
			},
			&expr.FunctionCall{
				Function:  nameE,
				Arguments: []node.Node{},
			},
		},
	}

	expected := map[node.Node]string{
		nameC: "A\\B\\C",
		nameF: "A\\B\\F",
		nameE: "B\\D\\E",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveTraitUse(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}}}

	fullyQualifiedNameB := &name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "B"}}}
	fullyQualifiedNameBC := &name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "B"}, &name.NamePart{Value: "C"}}}
	relativeNameB := &name.Relative{Parts: []node.Node{&name.NamePart{Value: "B"}}}
	relativeNameBC := &name.Relative{Parts: []node.Node{&name.NamePart{Value: "B"}, &name.NamePart{Value: "C"}}}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: nameAB,
					},
				},
			},
			&stmt.TraitUse{
				Traits: []node.Node{
					nameB,
					relativeNameB,
				},
				Adaptations: []node.Node{
					&stmt.TraitUsePrecedence{
						Ref: &stmt.TraitMethodRef{
							Trait:  fullyQualifiedNameB,
							Method: &node.Identifier{Value: "foo"},
						},
						Insteadof: []node.Node{fullyQualifiedNameBC},
					},
					&stmt.TraitUseAlias{
						Ref: &stmt.TraitMethodRef{
							Trait:  relativeNameBC,
							Method: &node.Identifier{Value: "foo"},
						},
						Alias: &node.Identifier{Value: "bar"},
					},
				},
			},
		},
	}

	expected := map[node.Node]string{
		nameB:                "A\\B",
		relativeNameB:        "B",
		fullyQualifiedNameB:  "B",
		fullyQualifiedNameBC: "B\\C",
		relativeNameBC:       "B\\C",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveClassName(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameBC := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}, &name.NamePart{Value: "C"}}}

	class := &stmt.Class{
		PhpDocComment: "",
		ClassName:     &node.Identifier{Value: "A"},
		Extends:       nameAB,
		Implements: []node.Node{
			nameBC,
		},
	}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			class,
		},
	}

	expected := map[node.Node]string{
		class:  "A",
		nameAB: "A\\B",
		nameBC: "B\\C",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveInterfaceName(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameBC := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}, &name.NamePart{Value: "C"}}}

	interfaceNode := &stmt.Interface{
		PhpDocComment: "",
		InterfaceName: &node.Identifier{Value: "A"},
		Extends: []node.Node{
			nameAB,
			nameBC,
		},
	}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			interfaceNode,
		},
	}

	expected := map[node.Node]string{
		interfaceNode: "A",
		nameAB:        "A\\B",
		nameBC:        "B\\C",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveTraitName(t *testing.T) {
	traitNode := &stmt.Trait{
		PhpDocComment: "",
		TraitName:     &node.Identifier{Value: "A"},
		Stmts:         []node.Node{},
	}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			traitNode,
		},
	}

	expected := map[node.Node]string{
		traitNode: "A",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveFunctionName(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameBC := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}, &name.NamePart{Value: "C"}}}

	functionNode := &stmt.Function{
		ReturnsRef:    false,
		PhpDocComment: "",
		FunctionName:  &node.Identifier{Value: "A"},
		Params: []node.Node{
			&node.Parameter{
				ByRef:        false,
				Variadic:     false,
				VariableType: nameAB,
				Variable:     &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
			},
		},
		ReturnType: &node.Nullable{Expr: nameBC},
		Stmts:      []node.Node{},
	}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			functionNode,
		},
	}

	expected := map[node.Node]string{
		functionNode: "A",
		nameAB:       "A\\B",
		nameBC:       "B\\C",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveMethodName(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameBC := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}, &name.NamePart{Value: "C"}}}

	methodNode := &stmt.ClassMethod{
		ReturnsRef:    false,
		PhpDocComment: "",
		MethodName:    &node.Identifier{Value: "A"},
		Params: []node.Node{
			&node.Parameter{
				ByRef:        false,
				Variadic:     false,
				VariableType: nameAB,
				Variable:     &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
			},
		},
		ReturnType: &node.Nullable{Expr: nameBC},
		Stmts:      []node.Node{},
	}

	expected := map[node.Node]string{
		nameAB: "A\\B",
		nameBC: "B\\C",
	}

	nsResolver := visitor.NewNamespaceResolver()
	methodNode.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveClosureName(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	nameBC := &name.Name{Parts: []node.Node{&name.NamePart{Value: "B"}, &name.NamePart{Value: "C"}}}

	closureNode := &expr.Closure{
		ReturnsRef:    false,
		Static:        false,
		PhpDocComment: "",
		Params: []node.Node{
			&node.Parameter{
				ByRef:        false,
				Variadic:     false,
				VariableType: nameAB,
				Variable:     &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
			},
		},
		ReturnType: &node.Nullable{Expr: nameBC},
		Stmts:      []node.Node{},
	}

	expected := map[node.Node]string{
		nameAB: "A\\B",
		nameBC: "B\\C",
	}

	nsResolver := visitor.NewNamespaceResolver()
	closureNode.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveConstantsName(t *testing.T) {
	nameAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}

	constantB := &stmt.Constant{
		PhpDocComment: "",
		ConstantName:  &node.Identifier{Value: "B"},
		Expr:          &scalar.Lnumber{Value: "1"},
	}
	constantC := &stmt.Constant{
		PhpDocComment: "",
		ConstantName:  &node.Identifier{Value: "C"},
		Expr:          &scalar.Lnumber{Value: "1"},
	}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Namespace{
				NamespaceName: nameAB,
			},
			&stmt.ConstList{
				Consts: []node.Node{
					constantB,
					constantC,
				},
			},
		},
	}

	expected := map[node.Node]string{
		constantB: "A\\B\\B",
		constantC: "A\\B\\C",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveNamespaces(t *testing.T) {
	namespaceAB := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "B"}}}
	namespaceCD := &name.Name{Parts: []node.Node{&name.NamePart{Value: "C"}, &name.NamePart{Value: "D"}}}

	nameAC := &name.Name{Parts: []node.Node{&name.NamePart{Value: "A"}, &name.NamePart{Value: "C"}}}
	nameCF := &name.Name{Parts: []node.Node{&name.NamePart{Value: "C"}, &name.NamePart{Value: "F"}}}
	nameFG := &name.Name{Parts: []node.Node{&name.NamePart{Value: "F"}, &name.NamePart{Value: "G"}}}
	relativeNameCE := &name.Relative{Parts: []node.Node{&name.NamePart{Value: "C"}, &name.NamePart{Value: "E"}}}

	constantB := &stmt.Constant{
		PhpDocComment: "",
		ConstantName:  &node.Identifier{Value: "B"},
		Expr:          &scalar.Lnumber{Value: "1"},
	}
	constantC := &stmt.Constant{
		PhpDocComment: "",
		ConstantName:  &node.Identifier{Value: "C"},
		Expr:          &scalar.Lnumber{Value: "1"},
	}

	ast := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Namespace{
				NamespaceName: namespaceAB,
			},
			&stmt.ConstList{
				Consts: []node.Node{
					constantB,
					constantC,
				},
			},
			&expr.StaticCall{
				Class:     nameFG,
				Call:      &node.Identifier{Value: "foo"},
				Arguments: []node.Node{},
			},
			&stmt.Namespace{
				Stmts: []node.Node{},
			},
			&stmt.Namespace{
				NamespaceName: namespaceCD,
				Stmts: []node.Node{
					&stmt.UseList{
						Uses: []node.Node{
							&stmt.Use{
								Use: nameAC,
							},
						},
					},
					&expr.StaticCall{
						Class:     relativeNameCE,
						Call:      &node.Identifier{Value: "foo"},
						Arguments: []node.Node{},
					},
					&expr.StaticCall{
						Class:     nameCF,
						Call:      &node.Identifier{Value: "foo"},
						Arguments: []node.Node{},
					},
				},
			},
		},
	}

	expected := map[node.Node]string{
		constantB:      "A\\B\\B",
		constantC:      "A\\B\\C",
		nameFG:         "A\\B\\F\\G",
		relativeNameCE: "C\\D\\C\\E",
		nameCF:         "A\\C\\F",
	}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveStaticCallDinamicClassName(t *testing.T) {
	ast := &stmt.StmtList{
		Stmts: []node.Node{
			&expr.StaticCall{
				Class:     &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
				Call:      &node.Identifier{Value: "foo"},
				Arguments: []node.Node{},
			},
		},
	}

	expected := map[node.Node]string{}

	nsResolver := visitor.NewNamespaceResolver()
	ast.Walk(nsResolver)

	assertEqual(t, expected, nsResolver.ResolvedNames)
}
