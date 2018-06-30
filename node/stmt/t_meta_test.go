package stmt_test

import (
	"testing"

	"github.com/z7zmey/php-parser/node/stmt"

	"github.com/z7zmey/php-parser/meta"

	"github.com/z7zmey/php-parser/node"
)

var nodes = []node.Node{
	&stmt.AltElseIf{},
	&stmt.AltElse{},
	&stmt.AltFor{},
	&stmt.AltForeach{},
	&stmt.AltIf{},
	&stmt.AltSwitch{},
	&stmt.AltWhile{},
	&stmt.Break{},
	&stmt.CaseList{},
	&stmt.Case{},
	&stmt.Catch{},
	&stmt.ClassConstList{},
	&stmt.ClassExtends{},
	&stmt.ClassImplements{},
	&stmt.ClassMethod{},
	&stmt.Class{},
	&stmt.ConstList{},
	&stmt.Constant{},
	&stmt.Continue{},
	&stmt.Declare{},
	&stmt.Default{},
	&stmt.Do{},
	&stmt.Echo{},
	&stmt.ElseIf{},
	&stmt.Else{},
	&stmt.Expression{},
	&stmt.Finally{},
	&stmt.For{},
	&stmt.Foreach{},
	&stmt.Function{},
	&stmt.Global{},
	&stmt.Goto{},
	&stmt.GroupUse{},
	&stmt.HaltCompiler{},
	&stmt.If{},
	&stmt.InlineHtml{},
	&stmt.InterfaceExtends{},
	&stmt.Interface{},
	&stmt.Label{},
	&stmt.Namespace{},
	&stmt.Nop{},
	&stmt.PropertyList{},
	&stmt.Property{},
	&stmt.Return{},
	&stmt.StaticVar{},
	&stmt.Static{},
	&stmt.StmtList{},
	&stmt.Switch{},
	&stmt.Throw{},
	&stmt.TraitAdaptationList{},
	&stmt.TraitMethodRef{},
	&stmt.TraitUseAlias{},
	&stmt.TraitUsePrecedence{},
	&stmt.TraitUse{},
	&stmt.Trait{},
	&stmt.Try{},
	&stmt.Unset{},
	&stmt.UseList{},
	&stmt.Use{},
	&stmt.While{},
}

func TestMeta(t *testing.T) {
	expected := []meta.Meta{
		meta.NewComment("//comment\n", nil),
		meta.NewWhiteSpace("    ", nil),
	}
	for _, n := range nodes {
		n.AddMeta(expected)
		actual := n.GetMeta()
		assertEqual(t, expected, actual)
	}
}
