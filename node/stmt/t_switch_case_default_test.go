package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/scalar"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestAltSwitch(t *testing.T) {
	src := `<? 
		switch (1) :
			case 1:
			default:
			case 2;
		endswitch;
	`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.AltSwitch{
				Cond: &scalar.Lnumber{Value: "1"},
				Cases: []node.Node{
					&stmt.Case{
						Cond:  &scalar.Lnumber{Value: "1"},
						Stmts: []node.Node{},
					},
					&stmt.Default{
						Stmts: []node.Node{},
					},
					&stmt.Case{
						Cond:  &scalar.Lnumber{Value: "2"},
						Stmts: []node.Node{},
					},
				},
			},
		},
	}

	actual, _, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestAltSwitchSemicolon(t *testing.T) {
	src := `<? 
		switch (1) :;
			case 1;
			case 2;
		endswitch;
	`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.AltSwitch{
				Cond: &scalar.Lnumber{Value: "1"},
				Cases: []node.Node{
					&stmt.Case{
						Cond:  &scalar.Lnumber{Value: "1"},
						Stmts: []node.Node{},
					},
					&stmt.Case{
						Cond:  &scalar.Lnumber{Value: "2"},
						Stmts: []node.Node{},
					},
				},
			},
		},
	}

	actual, _, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestSwitch(t *testing.T) {
	src := `<? 
		switch (1) {
			case 1: break;
			case 2: break;
		}
	`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Switch{
				Cond: &scalar.Lnumber{Value: "1"},
				Cases: []node.Node{
					&stmt.Case{
						Cond: &scalar.Lnumber{Value: "1"},
						Stmts: []node.Node{
							&stmt.Break{},
						},
					},
					&stmt.Case{
						Cond: &scalar.Lnumber{Value: "2"},
						Stmts: []node.Node{
							&stmt.Break{},
						},
					},
				},
			},
		},
	}

	actual, _, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestSwitchSemicolon(t *testing.T) {
	src := `<? 
		switch (1) {;
			case 1; break;
			case 2; break;
		}
	`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Switch{
				Cond: &scalar.Lnumber{Value: "1"},
				Cases: []node.Node{
					&stmt.Case{
						Cond: &scalar.Lnumber{Value: "1"},
						Stmts: []node.Node{
							&stmt.Break{},
						},
					},
					&stmt.Case{
						Cond: &scalar.Lnumber{Value: "2"},
						Stmts: []node.Node{
							&stmt.Break{},
						},
					},
				},
			},
		},
	}

	actual, _, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
