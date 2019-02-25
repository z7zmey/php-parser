package stmt_test

import (
	"bytes"
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/position"

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

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   6,
			StartPos:  7,
			EndPos:    65,
		},
		Stmts: []node.Node{
			&stmt.AltSwitch{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   6,
					StartPos:  7,
					EndPos:    65,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  15,
						EndPos:    15,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   -1,
						StartPos:  23,
						EndPos:    -1,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   -1,
								StartPos:  23,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  28,
									EndPos:    28,
								},
								Value: "1",
							},
							Stmts: []node.Node{},
						},
						&stmt.Default{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   -1,
								StartPos:  34,
								EndPos:    -1,
							},
							Stmts: []node.Node{},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   -1,
								StartPos:  46,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  51,
									EndPos:    51,
								},
								Value: "2",
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAltSwitchSemicolon(t *testing.T) {
	src := `<? 
		switch (1) :;
			case 1;
			case 2;
		endswitch;
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   5,
			StartPos:  7,
			EndPos:    54,
		},
		Stmts: []node.Node{
			&stmt.AltSwitch{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   5,
					StartPos:  7,
					EndPos:    54,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  15,
						EndPos:    15,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   -1,
						StartPos:  24,
						EndPos:    -1,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   -1,
								StartPos:  24,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  29,
									EndPos:    29,
								},
								Value: "1",
							},
							Stmts: []node.Node{},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   -1,
								StartPos:  35,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  40,
									EndPos:    40,
								},
								Value: "2",
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestSwitch(t *testing.T) {
	src := `<? 
		switch (1) {
			case 1: break;
			case 2: break;
		}
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   5,
			StartPos:  7,
			EndPos:    58,
		},
		Stmts: []node.Node{
			&stmt.Switch{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   5,
					StartPos:  7,
					EndPos:    58,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  15,
						EndPos:    15,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   5,
						StartPos:  18,
						EndPos:    58,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  23,
								EndPos:    36,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  28,
									EndPos:    28,
								},
								Value: "1",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  31,
										EndPos:    36,
									},
								},
							},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  41,
								EndPos:    54,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  46,
									EndPos:    46,
								},
								Value: "2",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  49,
										EndPos:    54,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestSwitchSemicolon(t *testing.T) {
	src := `<? 
		switch (1) {;
			case 1; break;
			case 2; break;
		}
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   5,
			StartPos:  7,
			EndPos:    59,
		},
		Stmts: []node.Node{
			&stmt.Switch{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   5,
					StartPos:  7,
					EndPos:    59,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  15,
						EndPos:    15,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   5,
						StartPos:  18,
						EndPos:    59,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  24,
								EndPos:    37,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  29,
									EndPos:    29,
								},
								Value: "1",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  32,
										EndPos:    37,
									},
								},
							},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  42,
								EndPos:    55,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  47,
									EndPos:    47,
								},
								Value: "2",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  50,
										EndPos:    55,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
