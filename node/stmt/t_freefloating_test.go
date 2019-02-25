package stmt_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
)

var expected freefloating.Collection = freefloating.Collection{
	freefloating.Start: []freefloating.String{
		{
			StringType: freefloating.WhiteSpaceType,
			Value:      "    ",
			Position:   nil,
		},
		{
			StringType: freefloating.CommentType,
			Value:      "//comment\n",
			Position:   nil,
		},
	},
}

var nodes = []node.Node{
	&stmt.AltElseIf{
		FreeFloating: expected,
	},
	&stmt.AltElse{
		FreeFloating: expected,
	},
	&stmt.AltFor{
		FreeFloating: expected,
	},
	&stmt.AltForeach{
		FreeFloating: expected,
	},
	&stmt.AltIf{
		FreeFloating: expected,
	},
	&stmt.AltSwitch{
		FreeFloating: expected,
	},
	&stmt.AltWhile{
		FreeFloating: expected,
	},
	&stmt.Break{
		FreeFloating: expected,
	},
	&stmt.CaseList{
		FreeFloating: expected,
	},
	&stmt.Case{
		FreeFloating: expected,
	},
	&stmt.Catch{
		FreeFloating: expected,
	},
	&stmt.ClassConstList{
		FreeFloating: expected,
	},
	&stmt.ClassExtends{
		FreeFloating: expected,
	},
	&stmt.ClassImplements{
		FreeFloating: expected,
	},
	&stmt.ClassMethod{
		FreeFloating: expected,
	},
	&stmt.Class{
		FreeFloating: expected,
	},
	&stmt.ConstList{
		FreeFloating: expected,
	},
	&stmt.Constant{
		FreeFloating: expected,
	},
	&stmt.Continue{
		FreeFloating: expected,
	},
	&stmt.Declare{
		FreeFloating: expected,
	},
	&stmt.Default{
		FreeFloating: expected,
	},
	&stmt.Do{
		FreeFloating: expected,
	},
	&stmt.Echo{
		FreeFloating: expected,
	},
	&stmt.ElseIf{
		FreeFloating: expected,
	},
	&stmt.Else{
		FreeFloating: expected,
	},
	&stmt.Expression{
		FreeFloating: expected,
	},
	&stmt.Finally{
		FreeFloating: expected,
	},
	&stmt.For{
		FreeFloating: expected,
	},
	&stmt.Foreach{
		FreeFloating: expected,
	},
	&stmt.Function{
		FreeFloating: expected,
	},
	&stmt.Global{
		FreeFloating: expected,
	},
	&stmt.Goto{
		FreeFloating: expected,
	},
	&stmt.GroupUse{
		FreeFloating: expected,
	},
	&stmt.HaltCompiler{
		FreeFloating: expected,
	},
	&stmt.If{
		FreeFloating: expected,
	},
	&stmt.InlineHtml{
		FreeFloating: expected,
	},
	&stmt.InterfaceExtends{
		FreeFloating: expected,
	},
	&stmt.Interface{
		FreeFloating: expected,
	},
	&stmt.Label{
		FreeFloating: expected,
	},
	&stmt.Namespace{
		FreeFloating: expected,
	},
	&stmt.Nop{
		FreeFloating: expected,
	},
	&stmt.PropertyList{
		FreeFloating: expected,
	},
	&stmt.Property{
		FreeFloating: expected,
	},
	&stmt.Return{
		FreeFloating: expected,
	},
	&stmt.StaticVar{
		FreeFloating: expected,
	},
	&stmt.Static{
		FreeFloating: expected,
	},
	&stmt.StmtList{
		FreeFloating: expected,
	},
	&stmt.Switch{
		FreeFloating: expected,
	},
	&stmt.Throw{
		FreeFloating: expected,
	},
	&stmt.TraitAdaptationList{
		FreeFloating: expected,
	},
	&stmt.TraitMethodRef{
		FreeFloating: expected,
	},
	&stmt.TraitUseAlias{
		FreeFloating: expected,
	},
	&stmt.TraitUsePrecedence{
		FreeFloating: expected,
	},
	&stmt.TraitUse{
		FreeFloating: expected,
	},
	&stmt.Trait{
		FreeFloating: expected,
	},
	&stmt.Try{
		FreeFloating: expected,
	},
	&stmt.Unset{
		FreeFloating: expected,
	},
	&stmt.UseList{
		FreeFloating: expected,
	},
	&stmt.Use{
		FreeFloating: expected,
	},
	&stmt.While{
		FreeFloating: expected,
	},
}

func TestMeta(t *testing.T) {
	for _, n := range nodes {
		actual := *n.GetFreeFloating()
		assert.DeepEqual(t, expected, actual)
	}
}
