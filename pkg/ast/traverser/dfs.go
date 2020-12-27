package traverser

import "github.com/z7zmey/php-parser/pkg/ast"

type DFS struct {
	visitor ast.Visitor
}

func NewDFS(visitor ast.Visitor) *DFS {
	return &DFS{
		visitor: visitor,
	}
}

func (t *DFS) Traverse(n ast.Vertex) {
	switch nn := n.(type) {
	case *ast.Root:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Stmts != nil {
			t.visitor.Enter("Stmts", false)
			for _, c := range nn.Stmts {
				t.Traverse(c)
			}
			t.visitor.Leave("Stmts", false)
		}
	case *ast.Nullable:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.Parameter:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Type != nil {
			t.visitor.Enter("Type", true)
			t.Traverse(nn.Type)
			t.visitor.Leave("Type", true)
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.DefaultValue != nil {
			t.visitor.Enter("DefaultValue", true)
			t.Traverse(nn.DefaultValue)
			t.visitor.Leave("DefaultValue", true)
		}
	case *ast.Identifier:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
	case *ast.Argument:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.StmtBreak:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.StmtCase:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Cond != nil {
			t.visitor.Enter("Cond", true)
			t.Traverse(nn.Cond)
			t.visitor.Leave("Cond", true)
		}
		if nn.Stmts != nil {
			t.visitor.Enter("Stmts", false)
			for _, c := range nn.Stmts {
				t.Traverse(c)
			}
			t.visitor.Leave("Stmts", false)
		}
	case *ast.StmtCatch:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Types != nil {
			t.visitor.Enter("Types", false)
			for _, c := range nn.Types {
				t.Traverse(c)
			}
			t.visitor.Leave("Types", false)
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Stmts != nil {
			t.visitor.Enter("Stmts", false)
			for _, c := range nn.Stmts {
				t.Traverse(c)
			}
			t.visitor.Leave("Stmts", false)
		}
	case *ast.StmtClass:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.ClassName != nil {
			t.visitor.Enter("ClassName", true)
			t.Traverse(nn.ClassName)
			t.visitor.Leave("ClassName", true)
		}
		if nn.Modifiers != nil {
			t.visitor.Enter("Modifiers", false)
			for _, c := range nn.Modifiers {
				t.Traverse(c)
			}
			t.visitor.Leave("Modifiers", false)
		}
		if nn.Arguments != nil {
			t.visitor.Enter("Arguments", false)
			for _, c := range nn.Arguments {
				t.Traverse(c)
			}
			t.visitor.Leave("Arguments", false)
		}
		if nn.Extends != nil {
			t.visitor.Enter("Extends", true)
			t.Traverse(nn.Extends)
			t.visitor.Leave("Extends", true)
		}
		if nn.Implements != nil {
			t.visitor.Enter("Implements", true)
			t.Traverse(nn.Implements)
			t.visitor.Leave("Implements", true)
		}
		if nn.Stmts != nil {
			t.visitor.Enter("Stmts", false)
			for _, c := range nn.Stmts {
				t.Traverse(c)
			}
			t.visitor.Leave("Stmts", false)
		}
	case *ast.StmtClassConstList:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Modifiers != nil {
			t.visitor.Enter("Modifiers", false)
			for _, c := range nn.Modifiers {
				t.Traverse(c)
			}
			t.visitor.Leave("Modifiers", false)
		}
		if nn.Consts != nil {
			t.visitor.Enter("Consts", false)
			for _, c := range nn.Consts {
				t.Traverse(c)
			}
			t.visitor.Leave("Consts", false)
		}
	case *ast.StmtClassExtends:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.ClassName != nil {
			t.visitor.Enter("ClassName", true)
			t.Traverse(nn.ClassName)
			t.visitor.Leave("ClassName", true)
		}
	case *ast.StmtClassImplements:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.InterfaceNames != nil {
			t.visitor.Enter("InterfaceNames", false)
			for _, c := range nn.InterfaceNames {
				t.Traverse(c)
			}
			t.visitor.Leave("InterfaceNames", false)
		}
	case *ast.StmtClassMethod:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.MethodName != nil {
			t.visitor.Enter("MethodName", true)
			t.Traverse(nn.MethodName)
			t.visitor.Leave("MethodName", true)
		}
		if nn.Modifiers != nil {
			t.visitor.Enter("Modifiers", false)
			for _, c := range nn.Modifiers {
				t.Traverse(c)
			}
			t.visitor.Leave("Modifiers", false)
		}
		if nn.Params != nil {
			t.visitor.Enter("Params", false)
			for _, c := range nn.Params {
				t.Traverse(c)
			}
			t.visitor.Leave("Params", false)
		}
		if nn.ReturnType != nil {
			t.visitor.Enter("ReturnType", true)
			t.Traverse(nn.ReturnType)
			t.visitor.Leave("ReturnType", true)
		}
		if nn.Stmt != nil {
			t.visitor.Enter("Stmt", true)
			t.Traverse(nn.Stmt)
			t.visitor.Leave("Stmt", true)
		}
	case *ast.StmtConstList:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Consts != nil {
			t.visitor.Enter("Consts", false)
			for _, c := range nn.Consts {
				t.Traverse(c)
			}
			t.visitor.Leave("Consts", false)
		}
	case *ast.StmtConstant:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Name != nil {
			t.visitor.Enter("Name", true)
			t.Traverse(nn.Name)
			t.visitor.Leave("Name", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.StmtContinue:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.StmtDeclare:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Consts != nil {
			t.visitor.Enter("Consts", false)
			for _, c := range nn.Consts {
				t.Traverse(c)
			}
			t.visitor.Leave("Consts", false)
		}
		if nn.Stmt != nil {
			t.visitor.Enter("Stmt", true)
			t.Traverse(nn.Stmt)
			t.visitor.Leave("Stmt", true)
		}
	case *ast.StmtDefault:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Stmts != nil {
			t.visitor.Enter("Stmts", false)
			for _, c := range nn.Stmts {
				t.Traverse(c)
			}
			t.visitor.Leave("Stmts", false)
		}
	case *ast.StmtDo:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Stmt != nil {
			t.visitor.Enter("Stmt", true)
			t.Traverse(nn.Stmt)
			t.visitor.Leave("Stmt", true)
		}
		if nn.Cond != nil {
			t.visitor.Enter("Cond", true)
			t.Traverse(nn.Cond)
			t.visitor.Leave("Cond", true)
		}
	case *ast.StmtEcho:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Exprs != nil {
			t.visitor.Enter("Exprs", false)
			for _, c := range nn.Exprs {
				t.Traverse(c)
			}
			t.visitor.Leave("Exprs", false)
		}
	case *ast.StmtElse:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Stmt != nil {
			t.visitor.Enter("Stmt", true)
			t.Traverse(nn.Stmt)
			t.visitor.Leave("Stmt", true)
		}
	case *ast.StmtElseIf:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Cond != nil {
			t.visitor.Enter("Cond", true)
			t.Traverse(nn.Cond)
			t.visitor.Leave("Cond", true)
		}
		if nn.Stmt != nil {
			t.visitor.Enter("Stmt", true)
			t.Traverse(nn.Stmt)
			t.visitor.Leave("Stmt", true)
		}
	case *ast.StmtExpression:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.StmtFinally:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Stmts != nil {
			t.visitor.Enter("Stmts", false)
			for _, c := range nn.Stmts {
				t.Traverse(c)
			}
			t.visitor.Leave("Stmts", false)
		}
	case *ast.StmtFor:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Init != nil {
			t.visitor.Enter("Init", false)
			for _, c := range nn.Init {
				t.Traverse(c)
			}
			t.visitor.Leave("Init", false)
		}
		if nn.Cond != nil {
			t.visitor.Enter("Cond", false)
			for _, c := range nn.Cond {
				t.Traverse(c)
			}
			t.visitor.Leave("Cond", false)
		}
		if nn.Loop != nil {
			t.visitor.Enter("Loop", false)
			for _, c := range nn.Loop {
				t.Traverse(c)
			}
			t.visitor.Leave("Loop", false)
		}
		if nn.Stmt != nil {
			t.visitor.Enter("Stmt", true)
			t.Traverse(nn.Stmt)
			t.visitor.Leave("Stmt", true)
		}
	case *ast.StmtForeach:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
		if nn.Key != nil {
			t.visitor.Enter("Key", true)
			t.Traverse(nn.Key)
			t.visitor.Leave("Key", true)
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Stmt != nil {
			t.visitor.Enter("Stmt", true)
			t.Traverse(nn.Stmt)
			t.visitor.Leave("Stmt", true)
		}
	case *ast.StmtFunction:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.FunctionName != nil {
			t.visitor.Enter("FunctionName", true)
			t.Traverse(nn.FunctionName)
			t.visitor.Leave("FunctionName", true)
		}
		if nn.Params != nil {
			t.visitor.Enter("Params", false)
			for _, c := range nn.Params {
				t.Traverse(c)
			}
			t.visitor.Leave("Params", false)
		}
		if nn.ReturnType != nil {
			t.visitor.Enter("ReturnType", true)
			t.Traverse(nn.ReturnType)
			t.visitor.Leave("ReturnType", true)
		}
		if nn.Stmts != nil {
			t.visitor.Enter("Stmts", false)
			for _, c := range nn.Stmts {
				t.Traverse(c)
			}
			t.visitor.Leave("Stmts", false)
		}
	case *ast.StmtGlobal:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Vars != nil {
			t.visitor.Enter("Vars", false)
			for _, c := range nn.Vars {
				t.Traverse(c)
			}
			t.visitor.Leave("Vars", false)
		}
	case *ast.StmtGoto:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Label != nil {
			t.visitor.Enter("Label", true)
			t.Traverse(nn.Label)
			t.visitor.Leave("Label", true)
		}
	case *ast.StmtHaltCompiler:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
	case *ast.StmtIf:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Cond != nil {
			t.visitor.Enter("Cond", true)
			t.Traverse(nn.Cond)
			t.visitor.Leave("Cond", true)
		}
		if nn.Stmt != nil {
			t.visitor.Enter("Stmt", true)
			t.Traverse(nn.Stmt)
			t.visitor.Leave("Stmt", true)
		}
		if nn.ElseIf != nil {
			t.visitor.Enter("ElseIf", false)
			for _, c := range nn.ElseIf {
				t.Traverse(c)
			}
			t.visitor.Leave("ElseIf", false)
		}
		if nn.Else != nil {
			t.visitor.Enter("Else", true)
			t.Traverse(nn.Else)
			t.visitor.Leave("Else", true)
		}
	case *ast.StmtInlineHtml:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
	case *ast.StmtInterface:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.InterfaceName != nil {
			t.visitor.Enter("InterfaceName", true)
			t.Traverse(nn.InterfaceName)
			t.visitor.Leave("InterfaceName", true)
		}
		if nn.Extends != nil {
			t.visitor.Enter("Extends", true)
			t.Traverse(nn.Extends)
			t.visitor.Leave("Extends", true)
		}
		if nn.Stmts != nil {
			t.visitor.Enter("Stmts", false)
			for _, c := range nn.Stmts {
				t.Traverse(c)
			}
			t.visitor.Leave("Stmts", false)
		}
	case *ast.StmtInterfaceExtends:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.InterfaceNames != nil {
			t.visitor.Enter("InterfaceNames", false)
			for _, c := range nn.InterfaceNames {
				t.Traverse(c)
			}
			t.visitor.Leave("InterfaceNames", false)
		}
	case *ast.StmtLabel:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.LabelName != nil {
			t.visitor.Enter("LabelName", true)
			t.Traverse(nn.LabelName)
			t.visitor.Leave("LabelName", true)
		}
	case *ast.StmtNamespace:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Name != nil {
			t.visitor.Enter("Name", true)
			t.Traverse(nn.Name)
			t.visitor.Leave("Name", true)
		}
		if nn.Stmts != nil {
			t.visitor.Enter("Stmts", false)
			for _, c := range nn.Stmts {
				t.Traverse(c)
			}
			t.visitor.Leave("Stmts", false)
		}
	case *ast.StmtNop:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
	case *ast.StmtProperty:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.StmtPropertyList:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Modifiers != nil {
			t.visitor.Enter("Modifiers", false)
			for _, c := range nn.Modifiers {
				t.Traverse(c)
			}
			t.visitor.Leave("Modifiers", false)
		}
		if nn.Type != nil {
			t.visitor.Enter("Type", true)
			t.Traverse(nn.Type)
			t.visitor.Leave("Type", true)
		}
		if nn.Properties != nil {
			t.visitor.Enter("Properties", false)
			for _, c := range nn.Properties {
				t.Traverse(c)
			}
			t.visitor.Leave("Properties", false)
		}
	case *ast.StmtReturn:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.StmtStatic:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Vars != nil {
			t.visitor.Enter("Vars", false)
			for _, c := range nn.Vars {
				t.Traverse(c)
			}
			t.visitor.Leave("Vars", false)
		}
	case *ast.StmtStaticVar:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.StmtStmtList:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Stmts != nil {
			t.visitor.Enter("Stmts", false)
			for _, c := range nn.Stmts {
				t.Traverse(c)
			}
			t.visitor.Leave("Stmts", false)
		}
	case *ast.StmtSwitch:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Cond != nil {
			t.visitor.Enter("Cond", true)
			t.Traverse(nn.Cond)
			t.visitor.Leave("Cond", true)
		}
		if nn.CaseList != nil {
			t.visitor.Enter("CaseList", false)
			for _, c := range nn.CaseList {
				t.Traverse(c)
			}
			t.visitor.Leave("CaseList", false)
		}
	case *ast.StmtThrow:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.StmtTrait:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.TraitName != nil {
			t.visitor.Enter("TraitName", true)
			t.Traverse(nn.TraitName)
			t.visitor.Leave("TraitName", true)
		}
		if nn.Stmts != nil {
			t.visitor.Enter("Stmts", false)
			for _, c := range nn.Stmts {
				t.Traverse(c)
			}
			t.visitor.Leave("Stmts", false)
		}
	case *ast.StmtTraitMethodRef:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Trait != nil {
			t.visitor.Enter("Trait", true)
			t.Traverse(nn.Trait)
			t.visitor.Leave("Trait", true)
		}
		if nn.Method != nil {
			t.visitor.Enter("Method", true)
			t.Traverse(nn.Method)
			t.visitor.Leave("Method", true)
		}
	case *ast.StmtTraitUse:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Traits != nil {
			t.visitor.Enter("Traits", false)
			for _, c := range nn.Traits {
				t.Traverse(c)
			}
			t.visitor.Leave("Traits", false)
		}
		if nn.Adaptations != nil {
			t.visitor.Enter("Adaptations", false)
			for _, c := range nn.Adaptations {
				t.Traverse(c)
			}
			t.visitor.Leave("Adaptations", false)
		}
	case *ast.StmtTraitUseAlias:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Ref != nil {
			t.visitor.Enter("Ref", true)
			t.Traverse(nn.Ref)
			t.visitor.Leave("Ref", true)
		}
		if nn.Modifier != nil {
			t.visitor.Enter("Modifier", true)
			t.Traverse(nn.Modifier)
			t.visitor.Leave("Modifier", true)
		}
		if nn.Alias != nil {
			t.visitor.Enter("Alias", true)
			t.Traverse(nn.Alias)
			t.visitor.Leave("Alias", true)
		}
	case *ast.StmtTraitUsePrecedence:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Ref != nil {
			t.visitor.Enter("Ref", true)
			t.Traverse(nn.Ref)
			t.visitor.Leave("Ref", true)
		}
		if nn.Insteadof != nil {
			t.visitor.Enter("Insteadof", false)
			for _, c := range nn.Insteadof {
				t.Traverse(c)
			}
			t.visitor.Leave("Insteadof", false)
		}
	case *ast.StmtTry:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Stmts != nil {
			t.visitor.Enter("Stmts", false)
			for _, c := range nn.Stmts {
				t.Traverse(c)
			}
			t.visitor.Leave("Stmts", false)
		}
		if nn.Catches != nil {
			t.visitor.Enter("Catches", false)
			for _, c := range nn.Catches {
				t.Traverse(c)
			}
			t.visitor.Leave("Catches", false)
		}
		if nn.Finally != nil {
			t.visitor.Enter("Finally", true)
			t.Traverse(nn.Finally)
			t.visitor.Leave("Finally", true)
		}
	case *ast.StmtUnset:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Vars != nil {
			t.visitor.Enter("Vars", false)
			for _, c := range nn.Vars {
				t.Traverse(c)
			}
			t.visitor.Leave("Vars", false)
		}
	case *ast.StmtUse:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Type != nil {
			t.visitor.Enter("Type", true)
			t.Traverse(nn.Type)
			t.visitor.Leave("Type", true)
		}
		if nn.UseDeclarations != nil {
			t.visitor.Enter("UseDeclarations", false)
			for _, c := range nn.UseDeclarations {
				t.Traverse(c)
			}
			t.visitor.Leave("UseDeclarations", false)
		}
	case *ast.StmtGroupUse:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Type != nil {
			t.visitor.Enter("Type", true)
			t.Traverse(nn.Type)
			t.visitor.Leave("Type", true)
		}
		if nn.Prefix != nil {
			t.visitor.Enter("Prefix", true)
			t.Traverse(nn.Prefix)
			t.visitor.Leave("Prefix", true)
		}
		if nn.UseDeclarations != nil {
			t.visitor.Enter("UseDeclarations", false)
			for _, c := range nn.UseDeclarations {
				t.Traverse(c)
			}
			t.visitor.Leave("UseDeclarations", false)
		}
	case *ast.StmtUseDeclaration:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Type != nil {
			t.visitor.Enter("Type", true)
			t.Traverse(nn.Type)
			t.visitor.Leave("Type", true)
		}
		if nn.Use != nil {
			t.visitor.Enter("Use", true)
			t.Traverse(nn.Use)
			t.visitor.Leave("Use", true)
		}
		if nn.Alias != nil {
			t.visitor.Enter("Alias", true)
			t.Traverse(nn.Alias)
			t.visitor.Leave("Alias", true)
		}
	case *ast.StmtWhile:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Cond != nil {
			t.visitor.Enter("Cond", true)
			t.Traverse(nn.Cond)
			t.visitor.Leave("Cond", true)
		}
		if nn.Stmt != nil {
			t.visitor.Enter("Stmt", true)
			t.Traverse(nn.Stmt)
			t.visitor.Leave("Stmt", true)
		}
	case *ast.ExprArray:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Items != nil {
			t.visitor.Enter("Items", false)
			for _, c := range nn.Items {
				t.Traverse(c)
			}
			t.visitor.Leave("Items", false)
		}
	case *ast.ExprArrayDimFetch:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Dim != nil {
			t.visitor.Enter("Dim", true)
			t.Traverse(nn.Dim)
			t.visitor.Leave("Dim", true)
		}
	case *ast.ExprArrayItem:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Key != nil {
			t.visitor.Enter("Key", true)
			t.Traverse(nn.Key)
			t.visitor.Leave("Key", true)
		}
		if nn.Val != nil {
			t.visitor.Enter("Val", true)
			t.Traverse(nn.Val)
			t.visitor.Leave("Val", true)
		}
	case *ast.ExprArrowFunction:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Params != nil {
			t.visitor.Enter("Params", false)
			for _, c := range nn.Params {
				t.Traverse(c)
			}
			t.visitor.Leave("Params", false)
		}
		if nn.ReturnType != nil {
			t.visitor.Enter("ReturnType", true)
			t.Traverse(nn.ReturnType)
			t.visitor.Leave("ReturnType", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprBitwiseNot:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprBooleanNot:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprBrackets:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprClassConstFetch:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Class != nil {
			t.visitor.Enter("Class", true)
			t.Traverse(nn.Class)
			t.visitor.Leave("Class", true)
		}
		if nn.ConstantName != nil {
			t.visitor.Enter("Name", true)
			t.Traverse(nn.ConstantName)
			t.visitor.Leave("Name", true)
		}
	case *ast.ExprClone:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprClosure:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Params != nil {
			t.visitor.Enter("Params", false)
			for _, c := range nn.Params {
				t.Traverse(c)
			}
			t.visitor.Leave("Params", false)
		}
		if nn.Use != nil {
			t.visitor.Enter("Use", false)
			for _, c := range nn.Use {
				t.Traverse(c)
			}
			t.visitor.Leave("Use", false)
		}
		if nn.ReturnType != nil {
			t.visitor.Enter("ReturnType", true)
			t.Traverse(nn.ReturnType)
			t.visitor.Leave("ReturnType", true)
		}
		if nn.Stmts != nil {
			t.visitor.Enter("Stmts", false)
			for _, c := range nn.Stmts {
				t.Traverse(c)
			}
			t.visitor.Leave("Stmts", false)
		}
	case *ast.ExprClosureUse:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
	case *ast.ExprConstFetch:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Const != nil {
			t.visitor.Enter("Const", true)
			t.Traverse(nn.Const)
			t.visitor.Leave("Const", true)
		}
	case *ast.ExprEmpty:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprErrorSuppress:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprEval:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprExit:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprFunctionCall:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Function != nil {
			t.visitor.Enter("Function", true)
			t.Traverse(nn.Function)
			t.visitor.Leave("Function", true)
		}
		if nn.Arguments != nil {
			t.visitor.Enter("Arguments", false)
			for _, c := range nn.Arguments {
				t.Traverse(c)
			}
			t.visitor.Leave("Arguments", false)
		}
	case *ast.ExprInclude:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprIncludeOnce:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprInstanceOf:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
		if nn.Class != nil {
			t.visitor.Enter("Class", true)
			t.Traverse(nn.Class)
			t.visitor.Leave("Class", true)
		}
	case *ast.ExprIsset:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Vars != nil {
			t.visitor.Enter("Vars", false)
			for _, c := range nn.Vars {
				t.Traverse(c)
			}
			t.visitor.Leave("Vars", false)
		}
	case *ast.ExprList:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Items != nil {
			t.visitor.Enter("Items", false)
			for _, c := range nn.Items {
				t.Traverse(c)
			}
			t.visitor.Leave("Items", false)
		}
	case *ast.ExprMethodCall:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Method != nil {
			t.visitor.Enter("Method", true)
			t.Traverse(nn.Method)
			t.visitor.Leave("Method", true)
		}
		if nn.Arguments != nil {
			t.visitor.Enter("Arguments", false)
			for _, c := range nn.Arguments {
				t.Traverse(c)
			}
			t.visitor.Leave("Arguments", false)
		}
	case *ast.ExprNew:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Class != nil {
			t.visitor.Enter("Class", true)
			t.Traverse(nn.Class)
			t.visitor.Leave("Class", true)
		}
		if nn.Arguments != nil {
			t.visitor.Enter("Arguments", false)
			for _, c := range nn.Arguments {
				t.Traverse(c)
			}
			t.visitor.Leave("Arguments", false)
		}
	case *ast.ExprPostDec:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
	case *ast.ExprPostInc:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
	case *ast.ExprPreDec:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
	case *ast.ExprPreInc:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
	case *ast.ExprPrint:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprPropertyFetch:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Property != nil {
			t.visitor.Enter("Property", true)
			t.Traverse(nn.Property)
			t.visitor.Leave("Property", true)
		}
	case *ast.ExprRequire:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprRequireOnce:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprShellExec:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Parts != nil {
			t.visitor.Enter("Parts", false)
			for _, c := range nn.Parts {
				t.Traverse(c)
			}
			t.visitor.Leave("Parts", false)
		}
	case *ast.ExprStaticCall:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Class != nil {
			t.visitor.Enter("Class", true)
			t.Traverse(nn.Class)
			t.visitor.Leave("Class", true)
		}
		if nn.Call != nil {
			t.visitor.Enter("Call", true)
			t.Traverse(nn.Call)
			t.visitor.Leave("Call", true)
		}
		if nn.Arguments != nil {
			t.visitor.Enter("Arguments", false)
			for _, c := range nn.Arguments {
				t.Traverse(c)
			}
			t.visitor.Leave("Arguments", false)
		}
	case *ast.ExprStaticPropertyFetch:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Class != nil {
			t.visitor.Enter("Class", true)
			t.Traverse(nn.Class)
			t.visitor.Leave("Class", true)
		}
		if nn.Property != nil {
			t.visitor.Enter("Property", true)
			t.Traverse(nn.Property)
			t.visitor.Leave("Property", true)
		}
	case *ast.ExprTernary:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Condition != nil {
			t.visitor.Enter("Condition", true)
			t.Traverse(nn.Condition)
			t.visitor.Leave("Condition", true)
		}
		if nn.IfTrue != nil {
			t.visitor.Enter("IfTrue", true)
			t.Traverse(nn.IfTrue)
			t.visitor.Leave("IfTrue", true)
		}
		if nn.IfFalse != nil {
			t.visitor.Enter("IfFalse", true)
			t.Traverse(nn.IfFalse)
			t.visitor.Leave("IfFalse", true)
		}
	case *ast.ExprUnaryMinus:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprUnaryPlus:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprVariable:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.VarName != nil {
			t.visitor.Enter("VarName", true)
			t.Traverse(nn.VarName)
			t.visitor.Leave("VarName", true)
		}
	case *ast.ExprYield:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Key != nil {
			t.visitor.Enter("Key", true)
			t.Traverse(nn.Key)
			t.visitor.Leave("Key", true)
		}
		if nn.Value != nil {
			t.visitor.Enter("Value", true)
			t.Traverse(nn.Value)
			t.visitor.Leave("Value", true)
		}
	case *ast.ExprYieldFrom:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssign:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssignReference:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssignBitwiseAnd:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssignBitwiseOr:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssignBitwiseXor:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssignCoalesce:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssignConcat:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssignDiv:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssignMinus:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssignMod:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssignMul:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssignPlus:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssignPow:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssignShiftLeft:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprAssignShiftRight:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprBinaryBitwiseAnd:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryBitwiseOr:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryBitwiseXor:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryBooleanAnd:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryBooleanOr:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryCoalesce:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryConcat:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryDiv:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryEqual:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryGreater:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryGreaterOrEqual:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryIdentical:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryLogicalAnd:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryLogicalOr:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryLogicalXor:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryMinus:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryMod:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryMul:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryNotEqual:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryNotIdentical:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryPlus:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryPow:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryShiftLeft:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinaryShiftRight:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinarySmaller:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinarySmallerOrEqual:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprBinarySpaceship:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Left != nil {
			t.visitor.Enter("Left", true)
			t.Traverse(nn.Left)
			t.visitor.Leave("Left", true)
		}
		if nn.Right != nil {
			t.visitor.Enter("Right", true)
			t.Traverse(nn.Right)
			t.visitor.Leave("Right", true)
		}
	case *ast.ExprCastArray:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprCastBool:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprCastDouble:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprCastInt:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprCastObject:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprCastString:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ExprCastUnset:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Expr != nil {
			t.visitor.Enter("Expr", true)
			t.Traverse(nn.Expr)
			t.visitor.Leave("Expr", true)
		}
	case *ast.ScalarDnumber:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
	case *ast.ScalarEncapsed:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Parts != nil {
			t.visitor.Enter("Parts", false)
			for _, c := range nn.Parts {
				t.Traverse(c)
			}
			t.visitor.Leave("Parts", false)
		}
	case *ast.ScalarEncapsedStringPart:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
	case *ast.ScalarEncapsedStringVar:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.VarName != nil {
			t.visitor.Enter("VarName", true)
			t.Traverse(nn.VarName)
			t.visitor.Leave("VarName", true)
		}
		if nn.Dim != nil {
			t.visitor.Enter("Dim", true)
			t.Traverse(nn.Dim)
			t.visitor.Leave("Dim", true)
		}
	case *ast.ScalarEncapsedStringBrackets:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Var != nil {
			t.visitor.Enter("Var", true)
			t.Traverse(nn.Var)
			t.visitor.Leave("Var", true)
		}
	case *ast.ScalarHeredoc:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Parts != nil {
			t.visitor.Enter("Parts", false)
			for _, c := range nn.Parts {
				t.Traverse(c)
			}
			t.visitor.Leave("Parts", false)
		}
	case *ast.ScalarLnumber:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
	case *ast.ScalarMagicConstant:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
	case *ast.ScalarString:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
	case *ast.NameName:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Parts != nil {
			t.visitor.Enter("Parts", false)
			for _, c := range nn.Parts {
				t.Traverse(c)
			}
			t.visitor.Leave("Parts", false)
		}
	case *ast.NameFullyQualified:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Parts != nil {
			t.visitor.Enter("Parts", false)
			for _, c := range nn.Parts {
				t.Traverse(c)
			}
			t.visitor.Leave("Parts", false)
		}
	case *ast.NameRelative:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
		if nn.Parts != nil {
			t.visitor.Enter("Parts", false)
			for _, c := range nn.Parts {
				t.Traverse(c)
			}
			t.visitor.Leave("Parts", false)
		}
	case *ast.NameNamePart:
		if nn == nil {
			return
		}
		if !t.visitor.EnterNode(nn) {
			return
		}
	default:
		panic("unexpected type of node")
	}

	t.visitor.LeaveNode(n)
}
