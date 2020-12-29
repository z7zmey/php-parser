package nsresolver_test

import (
	"github.com/z7zmey/php-parser/pkg/visitor/nsresolver"
	"github.com/z7zmey/php-parser/pkg/visitor/traverser"
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/pkg/ast"
)

func TestResolveStaticCall(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameBC := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}, &ast.NamePart{Value: []byte("C")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Use: nameAB,
					},
				},
			},
			&ast.ExprStaticCall{
				Class: nameBC,
				Call:  &ast.Identifier{Value: []byte("foo")},
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveStaticPropertyFetch(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameBC := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}, &ast.NamePart{Value: []byte("C")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Use: nameAB,
					},
				},
			},
			&ast.ExprStaticPropertyFetch{
				Class: nameBC,
				Prop:  &ast.Identifier{Value: []byte("foo")},
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveClassConstFetch(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameBC := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}, &ast.NamePart{Value: []byte("C")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Use: nameAB,
					},
				},
			},
			&ast.ExprClassConstFetch{
				Class: nameBC,
				Const: &ast.Identifier{Value: []byte("FOO")},
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveNew(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameBC := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}, &ast.NamePart{Value: []byte("C")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Use: nameAB,
					},
				},
			},
			&ast.ExprNew{
				Class: nameBC,
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveInstanceOf(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameBC := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}, &ast.NamePart{Value: []byte("C")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Use: nameAB,
					},
				},
			},
			&ast.ExprInstanceOf{
				Expr:  &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("foo")}},
				Class: nameBC,
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveInstanceCatch(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameBC := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}, &ast.NamePart{Value: []byte("C")}}}

	nameDE := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("D")}, &ast.NamePart{Value: []byte("E")}}}
	nameF := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("F")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Use: nameAB,
					},
					&ast.StmtUse{
						Use:   nameDE,
						Alias: &ast.Identifier{Value: []byte("F")},
					},
				},
			},
			&ast.StmtTry{
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Types: []ast.Vertex{
							nameBC,
							nameF,
						},
						Var:   &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("foo")}},
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameBC: "A\\B\\C",
		nameF:  "D\\E",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveFunctionCall(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Type: &ast.Identifier{Value: []byte("function")},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Use: nameAB,
					},
				},
			},
			&ast.ExprFunctionCall{
				Function: nameB,
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameB: "A\\B",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveConstFetch(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Type: &ast.Identifier{Value: []byte("const")},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Use: nameAB,
					},
				},
			},
			&ast.ExprConstFetch{
				Const: nameB,
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameB: "A\\B",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveGroupUse(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameBD := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}, &ast.NamePart{Value: []byte("D")}}}
	nameE := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("E")}}}
	nameC := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("C")}}}
	nameF := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("F")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtGroupUseList{
				Prefix: nameAB,
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Type: &ast.Identifier{Value: []byte("Function")},
						Use:  nameF,
					},
					&ast.StmtUse{
						Type: &ast.Identifier{Value: []byte("const")},
						Use:  nameC,
					},
				},
			},
			&ast.StmtGroupUseList{
				Prefix: nameBD,
				Type:   &ast.Identifier{Value: []byte("Function")},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Use: nameE,
					},
				},
			},
			&ast.ExprConstFetch{
				Const: nameC,
			},
			&ast.ExprFunctionCall{
				Function: nameF,
			},
			&ast.ExprFunctionCall{
				Function: nameE,
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameC: "A\\B\\C",
		nameF: "A\\B\\F",
		nameE: "B\\D\\E",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveTraitUse(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}}}
	nameD := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("D")}}}

	fullyQualifiedNameB := &ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}}}
	fullyQualifiedNameBC := &ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}, &ast.NamePart{Value: []byte("C")}}}
	relativeNameB := &ast.NameRelative{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}}}
	relativeNameBC := &ast.NameRelative{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}, &ast.NamePart{Value: []byte("C")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUseList{
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Use: nameAB,
					},
				},
			},
			&ast.StmtTraitUse{
				Traits: []ast.Vertex{
					nameB,
					relativeNameB,
				},
				Adaptations: []ast.Vertex{
					&ast.StmtTraitUsePrecedence{
						Trait:     fullyQualifiedNameB,
						Method:    &ast.Identifier{Value: []byte("foo")},
						Insteadof: []ast.Vertex{fullyQualifiedNameBC},
					},
					&ast.StmtTraitUseAlias{
						Trait:  relativeNameBC,
						Method: &ast.Identifier{Value: []byte("foo")},
						Alias:  &ast.Identifier{Value: []byte("bar")},
					},
				},
			},
			&ast.StmtTraitUse{
				Traits: []ast.Vertex{
					nameD,
				},
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameB:                "A\\B",
		nameD:                "D",
		relativeNameB:        "B",
		fullyQualifiedNameB:  "B",
		fullyQualifiedNameBC: "B\\C",
		relativeNameBC:       "B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveClassName(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameBC := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}, &ast.NamePart{Value: []byte("C")}}}

	class := &ast.StmtClass{
		Name:    &ast.Identifier{Value: []byte("A")},
		Extends: nameAB,
		Implements: []ast.Vertex{
			nameBC,
		},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			class,
		},
	}

	expected := map[ast.Vertex]string{
		class:  "A",
		nameAB: "A\\B",
		nameBC: "B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveInterfaceName(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameBC := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}, &ast.NamePart{Value: []byte("C")}}}

	interfaceNode := &ast.StmtInterface{
		Name: &ast.Identifier{Value: []byte("A")},
		Extends: []ast.Vertex{
			nameAB,
			nameBC,
		},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			interfaceNode,
		},
	}

	expected := map[ast.Vertex]string{
		interfaceNode: "A",
		nameAB:        "A\\B",
		nameBC:        "B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveTraitName(t *testing.T) {
	traitNode := &ast.StmtTrait{
		Name:  &ast.Identifier{Value: []byte("A")},
		Stmts: []ast.Vertex{},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			traitNode,
		},
	}

	expected := map[ast.Vertex]string{
		traitNode: "A",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveFunctionName(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameBC := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}, &ast.NamePart{Value: []byte("C")}}}

	functionNode := &ast.StmtFunction{
		Name: &ast.Identifier{Value: []byte("A")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Type: nameAB,
				Var:  &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("foo")}},
			},
		},
		ReturnType: &ast.Nullable{Expr: nameBC},
		Stmts:      []ast.Vertex{},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			functionNode,
		},
	}

	expected := map[ast.Vertex]string{
		functionNode: "A",
		nameAB:       "A\\B",
		nameBC:       "B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveMethodName(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameBC := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}, &ast.NamePart{Value: []byte("C")}}}

	methodNode := &ast.StmtClassMethod{
		Name: &ast.Identifier{Value: []byte("A")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Type: nameAB,
				Var:  &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("foo")}},
			},
		},
		ReturnType: &ast.Nullable{Expr: nameBC},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{},
		},
	}

	expected := map[ast.Vertex]string{
		nameAB: "A\\B",
		nameBC: "B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(methodNode)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveClosureName(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameBC := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("B")}, &ast.NamePart{Value: []byte("C")}}}

	closureNode := &ast.ExprClosure{
		Params: []ast.Vertex{
			&ast.Parameter{
				Type: nameAB,
				Var:  &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("foo")}},
			},
		},
		ReturnType: &ast.Nullable{Expr: nameBC},
		Stmts:      []ast.Vertex{},
	}

	expected := map[ast.Vertex]string{
		nameAB: "A\\B",
		nameBC: "B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(closureNode)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveConstantsName(t *testing.T) {
	nameAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}

	constantB := &ast.StmtConstant{
		Name: &ast.Identifier{Value: []byte("B")},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
	}
	constantC := &ast.StmtConstant{
		Name: &ast.Identifier{Value: []byte("C")},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: nameAB,
			},
			&ast.StmtConstList{
				Consts: []ast.Vertex{
					constantB,
					constantC,
				},
			},
		},
	}

	expected := map[ast.Vertex]string{
		constantB: "A\\B\\B",
		constantC: "A\\B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveNamespaces(t *testing.T) {
	namespaceAB := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	namespaceCD := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("C")}, &ast.NamePart{Value: []byte("D")}}}

	nameAC := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("C")}}}
	nameCF := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("C")}, &ast.NamePart{Value: []byte("F")}}}
	nameFG := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("F")}, &ast.NamePart{Value: []byte("G")}}}
	relativeNameCE := &ast.NameRelative{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("C")}, &ast.NamePart{Value: []byte("E")}}}

	constantB := &ast.StmtConstant{
		Name: &ast.Identifier{Value: []byte("B")},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
	}
	constantC := &ast.StmtConstant{
		Name: &ast.Identifier{Value: []byte("C")},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: namespaceAB,
			},
			&ast.StmtConstList{
				Consts: []ast.Vertex{
					constantB,
					constantC,
				},
			},
			&ast.ExprStaticCall{
				Class: nameFG,
				Call:  &ast.Identifier{Value: []byte("foo")},
			},
			&ast.StmtNamespace{
				Stmts: []ast.Vertex{},
			},
			&ast.StmtNamespace{
				Name: namespaceCD,
				Stmts: []ast.Vertex{
					&ast.StmtUseList{
						Uses: []ast.Vertex{
							&ast.StmtUse{
								Use: nameAC,
							},
						},
					},
					&ast.ExprStaticCall{
						Class: relativeNameCE,
						Call:  &ast.Identifier{Value: []byte("foo")},
					},
					&ast.ExprStaticCall{
						Class: nameCF,
						Call:  &ast.Identifier{Value: []byte("foo")},
					},
				},
			},
		},
	}

	expected := map[ast.Vertex]string{
		constantB:      "A\\B\\B",
		constantC:      "A\\B\\C",
		nameFG:         "A\\B\\F\\G",
		relativeNameCE: "C\\D\\C\\E",
		nameCF:         "A\\C\\F",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveStaticCallDinamicClassName(t *testing.T) {
	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.ExprStaticCall{
				Class: &ast.ExprVariable{Name: &ast.Identifier{Value: []byte("foo")}},
				Call:  &ast.Identifier{Value: []byte("foo")},
			},
		},
	}

	expected := map[ast.Vertex]string{}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestDoNotResolveReservedConstants(t *testing.T) {
	namespaceName := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}}

	constantTrue := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{Value: []byte("True")},
		},
	}

	constantFalse := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{Value: []byte("False")},
		},
	}

	constantNull := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{Value: []byte("NULL")},
		},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: namespaceName,
			},
			&ast.StmtExpression{
				Expr: &ast.ExprConstFetch{
					Const: constantTrue,
				},
			},
			&ast.StmtExpression{
				Expr: &ast.ExprConstFetch{
					Const: constantFalse,
				},
			},
			&ast.StmtExpression{
				Expr: &ast.ExprConstFetch{
					Const: constantNull,
				},
			},
		},
	}

	expected := map[ast.Vertex]string{
		constantTrue:  "true",
		constantFalse: "false",
		constantNull:  "null",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestDoNotResolveReservedNames(t *testing.T) {

	nameInt := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{Value: []byte("int")},
		},
	}

	nameFloat := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{Value: []byte("float")},
		},
	}

	nameBool := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{Value: []byte("bool")},
		},
	}

	nameString := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{Value: []byte("string")},
		},
	}

	nameVoid := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{Value: []byte("void")},
		},
	}

	nameIterable := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{Value: []byte("iterable")},
		},
	}

	nameObject := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{Value: []byte("object")},
		},
	}

	function := &ast.StmtFunction{
		Name: &ast.Identifier{Value: []byte("bar")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Type: nameInt,
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("Int")},
				},
			},
			&ast.Parameter{
				Type: nameFloat,
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("Float")},
				},
			},
			&ast.Parameter{
				Type: nameBool,
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("Bool")},
				},
			},
			&ast.Parameter{
				Type: nameString,
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("String")},
				},
			},
			&ast.Parameter{
				Type: nameVoid,
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("Void")},
				},
			},
			&ast.Parameter{
				Type: nameIterable,
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("Iterable")},
				},
			},
			&ast.Parameter{
				Type: nameObject,
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("Object")},
				},
			},
		},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{Value: []byte("Foo")},
					},
				},
			},
			function,
		},
	}

	expected := map[ast.Vertex]string{
		function:     "Foo\\bar",
		nameInt:      "int",
		nameFloat:    "float",
		nameBool:     "bool",
		nameString:   "string",
		nameVoid:     "void",
		nameIterable: "iterable",
		nameObject:   "object",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestDoNotResolveReservedSpecialNames(t *testing.T) {

	nameSelf := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{Value: []byte("Self")},
		},
	}

	nameStatic := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{Value: []byte("Static")},
		},
	}

	nameParent := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{Value: []byte("Parent")},
		},
	}

	cls := &ast.StmtClass{
		Name: &ast.Identifier{Value: []byte("Bar")},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Expr: &ast.ExprStaticCall{
					Class: nameSelf,
					Call:  &ast.Identifier{Value: []byte("func")},
				},
			},
			&ast.StmtExpression{
				Expr: &ast.ExprStaticCall{
					Class: nameStatic,
					Call:  &ast.Identifier{Value: []byte("func")},
				},
			},
			&ast.StmtExpression{
				Expr: &ast.ExprStaticCall{
					Class: nameParent,
					Call:  &ast.Identifier{Value: []byte("func")},
				},
			},
		},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{Value: []byte("Foo")},
					},
				},
			},
			cls,
		},
	}

	expected := map[ast.Vertex]string{
		cls:        "Foo\\Bar",
		nameSelf:   "self",
		nameStatic: "static",
		nameParent: "parent",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}
func TestResolvePropertyTypeName(t *testing.T) {
	nameSimple := &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameRelative := &ast.NameRelative{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}
	nameFullyQualified := &ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("A")}, &ast.NamePart{Value: []byte("B")}}}

	propertyNodeSimple := &ast.StmtPropertyList{
		Type: nameSimple,
	}

	propertyNodeRelative := &ast.StmtPropertyList{
		Type: nameRelative,
	}

	propertyNodeFullyQualified := &ast.StmtPropertyList{
		Type: nameFullyQualified,
	}

	classNode := &ast.StmtClass{
		Name: &ast.Identifier{Value: []byte("Bar")},
		Stmts: []ast.Vertex{
			propertyNodeSimple,
			propertyNodeRelative,
			propertyNodeFullyQualified,
		},
	}

	stmts := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{Value: []byte("Foo")},
					},
				},
			},
			classNode,
		},
	}

	expected := map[ast.Vertex]string{
		nameSimple:         "Foo\\A\\B",
		nameRelative:       "Foo\\A\\B",
		nameFullyQualified: "A\\B",
		classNode:          "Foo\\Bar",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stmts)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}
