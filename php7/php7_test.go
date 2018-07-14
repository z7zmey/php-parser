package php7_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/assign"
	"github.com/z7zmey/php-parser/node/expr/binary"
	"github.com/z7zmey/php-parser/node/expr/cast"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/position"
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

func TestPhp7(t *testing.T) {
	src := `<?
		foo($a, ...$b);
		$foo($a, ...$b);
		$foo->bar($a, ...$b);
		foo::bar($a, ...$b);
		$foo::bar($a, ...$b);
		new foo($a, ...$b);
		/** anonymous class */
		new class ($a, ...$b) {};
		new class {};
		new $foo;
		new $foo[1];
		new $foo{$bar};
		new $foo->bar;
		new $foo::$bar;
		new static::$bar;

		function foo(?bar $bar=null, baz &...$baz) {}
		class foo {public function foo(?bar $bar=null, baz &...$baz) {}}
		function(?bar $bar=null, baz &...$baz) {};
		static function(?bar $bar=null, baz &...$baz) {};

		1234567890123456789;
		12345678901234567890;
		0.;
		0b0111111111111111111111111111111111111111111111111111111111111111;
		0b1111111111111111111111111111111111111111111111111111111111111111;
		0x007111111111111111;
		0x8111111111111111;
		__CLASS__;
		__DIR__;
		__FILE__;
		__FUNCTION__;
		__LINE__;
		__NAMESPACE__;
		__METHOD__;
		__TRAIT__;

		"test $var";
		"test $var[1]";
		"test $var[-1]";
		"test $var[1234567890123456789012345678901234567890]";
		"test $var[-1234567890123456789012345678901234567890]";
		"test $var[bar]";
		"test $var[$bar]";
		"$foo $bar";
		"test $foo->bar()";
		"test ${foo}";
		"test ${foo[0]}";
		"test ${$foo}";
		"test {$foo->bar()}";

		if ($a) :
		endif;
		if ($a) :
		elseif ($b):
		endif;
		if ($a) :
		else:
		endif;
		if ($a) :
		elseif ($b):
		elseif ($c):
		else:
		endif;

		while (1) { break; }
		while (1) { break 2; }
		while (1) : break(3); endwhile;
		class foo{ public const FOO = 1, BAR = 2; }
		class foo{ const FOO = 1, BAR = 2; }
		class foo{ function bar() {} }
		class foo{ public static function &bar() {} }
		class foo{ public static function &bar(): void {} }
		abstract class foo{ }
		final class foo extends bar { }
		final class foo implements bar { }
		final class foo implements bar, baz { }
		new class() extends foo implements bar, baz { };

		const FOO = 1, BAR = 2;
		while (1) { continue; }
		while (1) { continue 2; }
		while (1) { continue(3); }
		declare(ticks=1);
		declare(ticks=1) {}
		declare(ticks=1): enddeclare;
		do {} while(1);
		echo $a, 1;
		echo($a);
		for($i = 0; $i < 10; $i++, $i++) {}
		for(; $i < 10; $i++, $i++) : endfor;
		foreach ($a as $v) {}
		foreach ($a as $v) : endforeach;
		foreach ($a as $k => $v) {}
		foreach ($a as $k => &$v) {}
		foreach ($a as $k => list($v)) {}
		foreach ($a as $k => [$v]) {}
		function foo() {}
		function foo() {return;}
		function &foo() {return 1;}
		function &foo(): void {}
		global $a, $b;
		a: 
		goto a;
		__halt_compiler();
		if ($a) {}
		if ($a) {} elseif ($b) {}
		if ($a) {} else {}
		if ($a) {} elseif ($b) {} elseif ($c) {} else {}
		if ($a) {} elseif ($b) {} else if ($c) {} else {}
		?> <div></div> <?
		interface Foo {}
		interface Foo extends Bar {}
		interface Foo extends Bar, Baz {}
		namespace Foo;
		namespace Foo {}
		namespace {}
		class foo {var $a;}
		class foo {public static $a, $b = 1;}
		static $a, $b = 1;

		switch (1) :
			case 1:
			default:
			case 2:
		endswitch;

		switch (1) :;
			case 1;
			case 2;
		endswitch;
		
		switch (1) {
			case 1: break;
			case 2: break;
		}
		
		switch (1) {;
			case 1; break;
			case 2; break;
		}

		throw $e;

		trait Foo {}
		class Foo { use Bar; }
		class Foo { use Bar, Baz {} }
		class Foo { use Bar, Baz { one as include; } }
		class Foo { use Bar, Baz { one as public; } }
		class Foo { use Bar, Baz { one as public two; } }
		class Foo { use Bar, Baz { Bar::one insteadof Baz, Quux; Baz::one as two; } }

		try {}
		try {} catch (Exception $e) {}
		try {} catch (Exception|RuntimeException $e) {}
		try {} catch (Exception $e) {} catch (RuntimeException $e) {}
		try {} catch (Exception $e) {} finally {}

		unset($a, $b,);

		use Foo;
		use \Foo;
		use \Foo as Bar;
		use Foo, Bar;
		use Foo, Bar as Baz;
		use function Foo, \Bar;
		use function Foo as foo, \Bar as bar;
		use const Foo, \Bar;
		use const Foo as foo, \Bar as bar;

		use \Foo\{Bar, Baz};
		use Foo\{Bar, Baz as quux};
		use function Foo\{Bar, Baz};
		use const \Foo\{Bar, Baz};
		use Foo\{const Bar, function Baz};

		$a[1];
		$a[1][2];
		array();
		array(1);
		array(1=>1, &$b,);
		~$a;
		!$a;

		Foo::Bar;
		$foo::Bar;
		clone($a);
		clone $a;
		function(){};
		function($a, $b) use ($c, &$d) {};
		function(): void {};
		foo;
		namespace\foo;
		\foo;

		empty($a);
		@$a;
		eval($a);
		exit;
		exit($a);
		die;
		die($a);
		foo();
		namespace\foo();
		\foo();
		$foo();

		$a--;
		$a++;
		--$a;
		++$a;

		include $a;
		include_once $a;
		require $a;
		require_once $a;

		$a instanceof Foo;
		$a instanceof namespace\Foo;
		$a instanceof \Foo;

		isset($a, $b);
		list($a) = $b;
		list($a[]) = $b;
		list(list($a)) = $b;

		$a->foo();
		new Foo();
		new namespace\Foo();
		new \Foo();
		new class ($a, ...$b) {};
		print($a);
		$a->foo;
		` + "`cmd $a`;" + `
		` + "`cmd`;" + `
		` + "``;" + `
		[];
		[1];
		[1=>1, &$b,];

		[$a] = $b;
		[$a[]] = $b;
		[list($a)] = $b;
		Foo::bar();
		namespace\Foo::bar();
		\Foo::bar();
		Foo::$bar;
		$foo::$bar;
		namespace\Foo::$bar;
		\Foo::$bar;
		$a ? $b : $c;
		$a ? : $c;
		$a ? $b ? $c : $d : $e;
		$a ? $b : $c ? $d : $e;
		-$a;
		+$a;
		$$a;
		yield;
		yield $a;
		yield $a => $b;
		yield from $a;
		
		(array)$a;
		(boolean)$a;
		(bool)$a;
		(double)$a;
		(float)$a;
		(integer)$a;
		(int)$a;
		(object)$a;
		(string)$a;
		(unset)$a;

		$a & $b;
		$a | $b;
		$a ^ $b;
		$a && $b;
		$a || $b;
		$a ?? $b;
		$a . $b;
		$a / $b;
		$a == $b;
		$a >= $b;
		$a > $b;
		$a === $b;
		$a and $b;
		$a or $b;
		$a xor $b;
		$a - $b;
		$a % $b;
		$a * $b;
		$a != $b;
		$a !== $b;
		$a + $b;
		$a ** $b;
		$a << $b;
		$a >> $b;
		$a <= $b;
		$a < $b;
		$a <=> $b;

		$a =& $b;
		$a = $b;
		$a &= $b;
		$a |= $b;
		$a ^= $b;
		$a .= $b;
		$a /= $b;
		$a -= $b;
		$a %= $b;
		$a *= $b;
		$a += $b;
		$a **= $b;
		$a <<= $b;
		$a >>= $b;

		class foo {public function class() {} }
		\foo\bar();

		function foo(&$a, ...$b) {
			__halt_compiler();
			function bar() {}
			class Baz {}
			trait Quux{}
			interface Quuux {}
		}
		
		function foo(&$a = 1, ...$b = 1, $c = 1) {}
		function foo(array $a, callable $b) {}
		abstract final class foo { abstract protected static function bar(); final private function baz() {} }

		(new Foo)->bar;
		(new Foo)();
		[$foo][0]();
		foo[1]();
		"foo"();
		[1]{$foo}();
		${foo()};

		Foo::$bar();
		Foo::{$bar[0]}();
		
		$foo->$bar;
		$foo->{$bar[0]};

		[1=>&$a, 2=>list($b)];
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   347,
			StartPos:  6,
			EndPos:    6318,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  6,
					EndPos:    20,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  6,
						EndPos:    19,
					},
					Function: &name.Name{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  6,
							EndPos:    8,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  6,
									EndPos:    8,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  9,
							EndPos:    19,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  10,
									EndPos:    11,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  10,
										EndPos:    11,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  10,
											EndPos:    11,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  14,
									EndPos:    18,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  17,
										EndPos:    18,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  17,
											EndPos:    18,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 3,
					EndLine:   3,
					StartPos:  24,
					EndPos:    39,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  24,
						EndPos:    38,
					},
					Function: &expr.Variable{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  24,
							EndPos:    27,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  24,
								EndPos:    27,
							},
							Value: "foo",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  28,
							EndPos:    38,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  29,
									EndPos:    30,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  29,
										EndPos:    30,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  29,
											EndPos:    30,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  33,
									EndPos:    37,
								},
								IsReference: false,
								Variadic:    true,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  36,
										EndPos:    37,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  36,
											EndPos:    37,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 4,
					EndLine:   4,
					StartPos:  43,
					EndPos:    63,
				},
				Expr: &expr.MethodCall{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   4,
						StartPos:  43,
						EndPos:    62,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  43,
							EndPos:    46,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  43,
								EndPos:    46,
							},
							Value: "foo",
						},
					},
					Method: &node.Identifier{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  49,
							EndPos:    51,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  52,
							EndPos:    62,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  53,
									EndPos:    54,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  53,
										EndPos:    54,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  53,
											EndPos:    54,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  57,
									EndPos:    61,
								},
								IsReference: false,
								Variadic:    true,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  60,
										EndPos:    61,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  60,
											EndPos:    61,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 5,
					EndLine:   5,
					StartPos:  67,
					EndPos:    86,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 5,
						EndLine:   5,
						StartPos:  67,
						EndPos:    85,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  67,
							EndPos:    69,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  67,
									EndPos:    69,
								},
								Value: "foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  72,
							EndPos:    74,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  75,
							EndPos:    85,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  76,
									EndPos:    77,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  76,
										EndPos:    77,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  76,
											EndPos:    77,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  80,
									EndPos:    84,
								},
								IsReference: false,
								Variadic:    true,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  83,
										EndPos:    84,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  83,
											EndPos:    84,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 6,
					EndLine:   6,
					StartPos:  90,
					EndPos:    110,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 6,
						EndLine:   6,
						StartPos:  90,
						EndPos:    109,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  90,
							EndPos:    93,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 6,
								EndLine:   6,
								StartPos:  90,
								EndPos:    93,
							},
							Value: "foo",
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  96,
							EndPos:    98,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  99,
							EndPos:    109,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 6,
									EndLine:   6,
									StartPos:  100,
									EndPos:    101,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 6,
										EndLine:   6,
										StartPos:  100,
										EndPos:    101,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 6,
											EndLine:   6,
											StartPos:  100,
											EndPos:    101,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 6,
									EndLine:   6,
									StartPos:  104,
									EndPos:    108,
								},
								IsReference: false,
								Variadic:    true,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 6,
										EndLine:   6,
										StartPos:  107,
										EndPos:    108,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 6,
											EndLine:   6,
											StartPos:  107,
											EndPos:    108,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 7,
					EndLine:   7,
					StartPos:  114,
					EndPos:    132,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   7,
						StartPos:  114,
						EndPos:    131,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   7,
							StartPos:  118,
							EndPos:    120,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 7,
									EndLine:   7,
									StartPos:  118,
									EndPos:    120,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   7,
							StartPos:  121,
							EndPos:    131,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 7,
									EndLine:   7,
									StartPos:  122,
									EndPos:    123,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  122,
										EndPos:    123,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 7,
											EndLine:   7,
											StartPos:  122,
											EndPos:    123,
										},
										Value: "a",
									},
								},
							},
							&node.Argument{
								Position: &position.Position{
									StartLine: 7,
									EndLine:   7,
									StartPos:  126,
									EndPos:    130,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  129,
										EndPos:    130,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 7,
											EndLine:   7,
											StartPos:  129,
											EndPos:    130,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 9,
					EndLine:   9,
					StartPos:  161,
					EndPos:    185,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 9,
						EndLine:   9,
						StartPos:  161,
						EndPos:    184,
					},
					Class: &stmt.Class{
						Position: &position.Position{
							StartLine: 9,
							EndLine:   9,
							StartPos:  165,
							EndPos:    184,
						},
						PhpDocComment: "/** anonymous class */",
						ArgumentList: &node.ArgumentList{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  171,
								EndPos:    181,
							},
							Arguments: []node.Node{
								&node.Argument{
									Position: &position.Position{
										StartLine: 9,
										EndLine:   9,
										StartPos:  172,
										EndPos:    173,
									},
									Variadic:    false,
									IsReference: false,
									Expr: &expr.Variable{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  172,
											EndPos:    173,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  172,
												EndPos:    173,
											},
											Value: "a",
										},
									},
								},
								&node.Argument{
									Position: &position.Position{
										StartLine: 9,
										EndLine:   9,
										StartPos:  176,
										EndPos:    180,
									},
									Variadic:    true,
									IsReference: false,
									Expr: &expr.Variable{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  179,
											EndPos:    180,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  179,
												EndPos:    180,
											},
											Value: "b",
										},
									},
								},
							},
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 10,
					EndLine:   10,
					StartPos:  189,
					EndPos:    201,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 10,
						EndLine:   10,
						StartPos:  189,
						EndPos:    200,
					},
					Class: &stmt.Class{
						Position: &position.Position{
							StartLine: 10,
							EndLine:   10,
							StartPos:  193,
							EndPos:    200,
						},
						PhpDocComment: "",
						Stmts:         []node.Node{},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 11,
					EndLine:   11,
					StartPos:  205,
					EndPos:    213,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 11,
						EndLine:   11,
						StartPos:  205,
						EndPos:    212,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 11,
							EndLine:   11,
							StartPos:  209,
							EndPos:    212,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 11,
								EndLine:   11,
								StartPos:  209,
								EndPos:    212,
							},
							Value: "foo",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 12,
					EndLine:   12,
					StartPos:  217,
					EndPos:    228,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 12,
						EndLine:   12,
						StartPos:  217,
						EndPos:    227,
					},
					Class: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 12,
							EndLine:   12,
							StartPos:  221,
							EndPos:    227,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 12,
								EndLine:   12,
								StartPos:  221,
								EndPos:    224,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 12,
									EndLine:   12,
									StartPos:  221,
									EndPos:    224,
								},
								Value: "foo",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 12,
								EndLine:   12,
								StartPos:  226,
								EndPos:    226,
							},
							Value: "1",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 13,
					EndLine:   13,
					StartPos:  232,
					EndPos:    246,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 13,
						EndLine:   13,
						StartPos:  232,
						EndPos:    245,
					},
					Class: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 13,
							EndLine:   13,
							StartPos:  236,
							EndPos:    245,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 13,
								EndLine:   13,
								StartPos:  236,
								EndPos:    239,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 13,
									EndLine:   13,
									StartPos:  236,
									EndPos:    239,
								},
								Value: "foo",
							},
						},
						Dim: &expr.Variable{
							Position: &position.Position{
								StartLine: 13,
								EndLine:   13,
								StartPos:  241,
								EndPos:    244,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 13,
									EndLine:   13,
									StartPos:  241,
									EndPos:    244,
								},
								Value: "bar",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 14,
					EndLine:   14,
					StartPos:  250,
					EndPos:    263,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 14,
						EndLine:   14,
						StartPos:  250,
						EndPos:    262,
					},
					Class: &expr.PropertyFetch{
						Position: &position.Position{
							StartLine: 14,
							EndLine:   14,
							StartPos:  254,
							EndPos:    262,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 14,
								EndLine:   14,
								StartPos:  254,
								EndPos:    257,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 14,
									EndLine:   14,
									StartPos:  254,
									EndPos:    257,
								},
								Value: "foo",
							},
						},
						Property: &node.Identifier{
							Position: &position.Position{
								StartLine: 14,
								EndLine:   14,
								StartPos:  260,
								EndPos:    262,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 15,
					EndLine:   15,
					StartPos:  267,
					EndPos:    281,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 15,
						EndLine:   15,
						StartPos:  267,
						EndPos:    280,
					},
					Class: &expr.StaticPropertyFetch{
						Position: &position.Position{
							StartLine: 15,
							EndLine:   15,
							StartPos:  271,
							EndPos:    280,
						},
						Class: &expr.Variable{
							Position: &position.Position{
								StartLine: 15,
								EndLine:   15,
								StartPos:  271,
								EndPos:    274,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 15,
									EndLine:   15,
									StartPos:  271,
									EndPos:    274,
								},
								Value: "foo",
							},
						},
						Property: &expr.Variable{
							Position: &position.Position{
								StartLine: 15,
								EndLine:   15,
								StartPos:  277,
								EndPos:    280,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 15,
									EndLine:   15,
									StartPos:  277,
									EndPos:    280,
								},
								Value: "bar",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 16,
					EndLine:   16,
					StartPos:  285,
					EndPos:    301,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 16,
						EndLine:   16,
						StartPos:  285,
						EndPos:    300,
					},
					Class: &expr.StaticPropertyFetch{
						Position: &position.Position{
							StartLine: 16,
							EndLine:   16,
							StartPos:  289,
							EndPos:    300,
						},
						Class: &node.Identifier{
							Position: &position.Position{
								StartLine: 16,
								EndLine:   16,
								StartPos:  289,
								EndPos:    294,
							},
							Value: "static",
						},
						Property: &expr.Variable{
							Position: &position.Position{
								StartLine: 16,
								EndLine:   16,
								StartPos:  297,
								EndPos:    300,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 16,
									EndLine:   16,
									StartPos:  297,
									EndPos:    300,
								},
								Value: "bar",
							},
						},
					},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 18,
					EndLine:   18,
					StartPos:  306,
					EndPos:    350,
				},
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 18,
						EndLine:   18,
						StartPos:  315,
						EndPos:    317,
					},
					Value: "foo",
				},
				Params: []node.Node{
					&node.Parameter{
						Position: &position.Position{
							StartLine: 18,
							EndLine:   18,
							StartPos:  319,
							EndPos:    332,
						},
						ByRef:    false,
						Variadic: false,
						VariableType: &node.Nullable{
							Position: &position.Position{
								StartLine: 18,
								EndLine:   18,
								StartPos:  319,
								EndPos:    322,
							},
							Expr: &name.Name{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  320,
									EndPos:    322,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 18,
											EndLine:   18,
											StartPos:  320,
											EndPos:    322,
										},
										Value: "bar",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 18,
								EndLine:   18,
								StartPos:  324,
								EndPos:    327,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  324,
									EndPos:    327,
								},
								Value: "bar",
							},
						},
						DefaultValue: &expr.ConstFetch{
							Position: &position.Position{
								StartLine: 18,
								EndLine:   18,
								StartPos:  329,
								EndPos:    332,
							},
							Constant: &name.Name{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  329,
									EndPos:    332,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 18,
											EndLine:   18,
											StartPos:  329,
											EndPos:    332,
										},
										Value: "null",
									},
								},
							},
						},
					},
					&node.Parameter{
						Position: &position.Position{
							StartLine: 18,
							EndLine:   18,
							StartPos:  335,
							EndPos:    346,
						},
						ByRef:    true,
						Variadic: true,
						VariableType: &name.Name{
							Position: &position.Position{
								StartLine: 18,
								EndLine:   18,
								StartPos:  335,
								EndPos:    337,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 18,
										EndLine:   18,
										StartPos:  335,
										EndPos:    337,
									},
									Value: "baz",
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 18,
								EndLine:   18,
								StartPos:  343,
								EndPos:    346,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  343,
									EndPos:    346,
								},
								Value: "baz",
							},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 19,
					EndLine:   19,
					StartPos:  354,
					EndPos:    417,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 19,
						EndLine:   19,
						StartPos:  360,
						EndPos:    362,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 19,
							EndLine:   19,
							StartPos:  365,
							EndPos:    416,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 19,
								EndLine:   19,
								StartPos:  381,
								EndPos:    383,
							},
							Value: "foo",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 19,
									EndLine:   19,
									StartPos:  365,
									EndPos:    370,
								},
								Value: "public",
							},
						},
						Params: []node.Node{
							&node.Parameter{
								Position: &position.Position{
									StartLine: 19,
									EndLine:   19,
									StartPos:  385,
									EndPos:    398,
								},
								ByRef:    false,
								Variadic: false,
								VariableType: &node.Nullable{
									Position: &position.Position{
										StartLine: 19,
										EndLine:   19,
										StartPos:  385,
										EndPos:    388,
									},
									Expr: &name.Name{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  386,
											EndPos:    388,
										},
										Parts: []node.Node{
											&name.NamePart{
												Position: &position.Position{
													StartLine: 19,
													EndLine:   19,
													StartPos:  386,
													EndPos:    388,
												},
												Value: "bar",
											},
										},
									},
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 19,
										EndLine:   19,
										StartPos:  390,
										EndPos:    393,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  390,
											EndPos:    393,
										},
										Value: "bar",
									},
								},
								DefaultValue: &expr.ConstFetch{
									Position: &position.Position{
										StartLine: 19,
										EndLine:   19,
										StartPos:  395,
										EndPos:    398,
									},
									Constant: &name.Name{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  395,
											EndPos:    398,
										},
										Parts: []node.Node{
											&name.NamePart{
												Position: &position.Position{
													StartLine: 19,
													EndLine:   19,
													StartPos:  395,
													EndPos:    398,
												},
												Value: "null",
											},
										},
									},
								},
							},
							&node.Parameter{
								Position: &position.Position{
									StartLine: 19,
									EndLine:   19,
									StartPos:  401,
									EndPos:    412,
								},
								ByRef:    true,
								Variadic: true,
								VariableType: &name.Name{
									Position: &position.Position{
										StartLine: 19,
										EndLine:   19,
										StartPos:  401,
										EndPos:    403,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 19,
												EndLine:   19,
												StartPos:  401,
												EndPos:    403,
											},
											Value: "baz",
										},
									},
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 19,
										EndLine:   19,
										StartPos:  409,
										EndPos:    412,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  409,
											EndPos:    412,
										},
										Value: "baz",
									},
								},
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 19,
								EndLine:   19,
								StartPos:  415,
								EndPos:    416,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 20,
					EndLine:   20,
					StartPos:  421,
					EndPos:    462,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 20,
						EndLine:   20,
						StartPos:  421,
						EndPos:    461,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 20,
								EndLine:   20,
								StartPos:  430,
								EndPos:    443,
							},
							ByRef:    false,
							Variadic: false,
							VariableType: &node.Nullable{
								Position: &position.Position{
									StartLine: 20,
									EndLine:   20,
									StartPos:  430,
									EndPos:    433,
								},
								Expr: &name.Name{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  431,
										EndPos:    433,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 20,
												EndLine:   20,
												StartPos:  431,
												EndPos:    433,
											},
											Value: "bar",
										},
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 20,
									EndLine:   20,
									StartPos:  435,
									EndPos:    438,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  435,
										EndPos:    438,
									},
									Value: "bar",
								},
							},
							DefaultValue: &expr.ConstFetch{
								Position: &position.Position{
									StartLine: 20,
									EndLine:   20,
									StartPos:  440,
									EndPos:    443,
								},
								Constant: &name.Name{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  440,
										EndPos:    443,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 20,
												EndLine:   20,
												StartPos:  440,
												EndPos:    443,
											},
											Value: "null",
										},
									},
								},
							},
						},
						&node.Parameter{
							Position: &position.Position{
								StartLine: 20,
								EndLine:   20,
								StartPos:  446,
								EndPos:    457,
							},
							Variadic: true,
							ByRef:    true,
							VariableType: &name.Name{
								Position: &position.Position{
									StartLine: 20,
									EndLine:   20,
									StartPos:  446,
									EndPos:    448,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 20,
											EndLine:   20,
											StartPos:  446,
											EndPos:    448,
										},
										Value: "baz",
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 20,
									EndLine:   20,
									StartPos:  454,
									EndPos:    457,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  454,
										EndPos:    457,
									},
									Value: "baz",
								},
							},
						},
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 21,
					EndLine:   21,
					StartPos:  466,
					EndPos:    514,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 21,
						EndLine:   21,
						StartPos:  466,
						EndPos:    513,
					},
					ReturnsRef:    false,
					Static:        true,
					PhpDocComment: "",
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 21,
								EndLine:   21,
								StartPos:  482,
								EndPos:    495,
							},
							ByRef:    false,
							Variadic: false,
							VariableType: &node.Nullable{
								Position: &position.Position{
									StartLine: 21,
									EndLine:   21,
									StartPos:  482,
									EndPos:    485,
								},
								Expr: &name.Name{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  483,
										EndPos:    485,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 21,
												EndLine:   21,
												StartPos:  483,
												EndPos:    485,
											},
											Value: "bar",
										},
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 21,
									EndLine:   21,
									StartPos:  487,
									EndPos:    490,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  487,
										EndPos:    490,
									},
									Value: "bar",
								},
							},
							DefaultValue: &expr.ConstFetch{
								Position: &position.Position{
									StartLine: 21,
									EndLine:   21,
									StartPos:  492,
									EndPos:    495,
								},
								Constant: &name.Name{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  492,
										EndPos:    495,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 21,
												EndLine:   21,
												StartPos:  492,
												EndPos:    495,
											},
											Value: "null",
										},
									},
								},
							},
						},
						&node.Parameter{
							Position: &position.Position{
								StartLine: 21,
								EndLine:   21,
								StartPos:  498,
								EndPos:    509,
							},
							ByRef:    true,
							Variadic: true,
							VariableType: &name.Name{
								Position: &position.Position{
									StartLine: 21,
									EndLine:   21,
									StartPos:  498,
									EndPos:    500,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 21,
											EndLine:   21,
											StartPos:  498,
											EndPos:    500,
										},
										Value: "baz",
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 21,
									EndLine:   21,
									StartPos:  506,
									EndPos:    509,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  506,
										EndPos:    509,
									},
									Value: "baz",
								},
							},
						},
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 23,
					EndLine:   23,
					StartPos:  519,
					EndPos:    538,
				},
				Expr: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 23,
						EndLine:   23,
						StartPos:  519,
						EndPos:    537,
					},
					Value: "1234567890123456789",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 24,
					EndLine:   24,
					StartPos:  542,
					EndPos:    562,
				},
				Expr: &scalar.Dnumber{
					Position: &position.Position{
						StartLine: 24,
						EndLine:   24,
						StartPos:  542,
						EndPos:    561,
					},
					Value: "12345678901234567890",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 25,
					EndLine:   25,
					StartPos:  566,
					EndPos:    568,
				},
				Expr: &scalar.Dnumber{
					Position: &position.Position{
						StartLine: 25,
						EndLine:   25,
						StartPos:  566,
						EndPos:    567,
					},
					Value: "0.",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 26,
					EndLine:   26,
					StartPos:  572,
					EndPos:    638,
				},
				Expr: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 26,
						EndLine:   26,
						StartPos:  572,
						EndPos:    637,
					},
					Value: "0b0111111111111111111111111111111111111111111111111111111111111111",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 27,
					EndLine:   27,
					StartPos:  642,
					EndPos:    708,
				},
				Expr: &scalar.Dnumber{
					Position: &position.Position{
						StartLine: 27,
						EndLine:   27,
						StartPos:  642,
						EndPos:    707,
					},
					Value: "0b1111111111111111111111111111111111111111111111111111111111111111",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 28,
					EndLine:   28,
					StartPos:  712,
					EndPos:    732,
				},
				Expr: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 28,
						EndLine:   28,
						StartPos:  712,
						EndPos:    731,
					},
					Value: "0x007111111111111111",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 29,
					EndLine:   29,
					StartPos:  736,
					EndPos:    754,
				},
				Expr: &scalar.Dnumber{
					Position: &position.Position{
						StartLine: 29,
						EndLine:   29,
						StartPos:  736,
						EndPos:    753,
					},
					Value: "0x8111111111111111",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 30,
					EndLine:   30,
					StartPos:  758,
					EndPos:    767,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 30,
						EndLine:   30,
						StartPos:  758,
						EndPos:    766,
					},
					Value: "__CLASS__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 31,
					EndLine:   31,
					StartPos:  771,
					EndPos:    778,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 31,
						EndLine:   31,
						StartPos:  771,
						EndPos:    777,
					},
					Value: "__DIR__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 32,
					EndLine:   32,
					StartPos:  782,
					EndPos:    790,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 32,
						EndLine:   32,
						StartPos:  782,
						EndPos:    789,
					},
					Value: "__FILE__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 33,
					EndLine:   33,
					StartPos:  794,
					EndPos:    806,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 33,
						EndLine:   33,
						StartPos:  794,
						EndPos:    805,
					},
					Value: "__FUNCTION__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 34,
					EndLine:   34,
					StartPos:  810,
					EndPos:    818,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 34,
						EndLine:   34,
						StartPos:  810,
						EndPos:    817,
					},
					Value: "__LINE__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 35,
					EndLine:   35,
					StartPos:  822,
					EndPos:    835,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 35,
						EndLine:   35,
						StartPos:  822,
						EndPos:    834,
					},
					Value: "__NAMESPACE__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 36,
					EndLine:   36,
					StartPos:  839,
					EndPos:    849,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 36,
						EndLine:   36,
						StartPos:  839,
						EndPos:    848,
					},
					Value: "__METHOD__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 37,
					EndLine:   37,
					StartPos:  853,
					EndPos:    862,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 37,
						EndLine:   37,
						StartPos:  853,
						EndPos:    861,
					},
					Value: "__TRAIT__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 39,
					EndLine:   39,
					StartPos:  867,
					EndPos:    878,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 39,
						EndLine:   39,
						StartPos:  867,
						EndPos:    877,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 39,
								EndLine:   39,
								StartPos:  868,
								EndPos:    872,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 39,
								EndLine:   39,
								StartPos:  873,
								EndPos:    876,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 39,
									EndLine:   39,
									StartPos:  873,
									EndPos:    876,
								},
								Value: "var",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 40,
					EndLine:   40,
					StartPos:  882,
					EndPos:    896,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 40,
						EndLine:   40,
						StartPos:  882,
						EndPos:    895,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 40,
								EndLine:   40,
								StartPos:  883,
								EndPos:    887,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 40,
								EndLine:   40,
								StartPos:  888,
								EndPos:    894,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 40,
									EndLine:   40,
									StartPos:  888,
									EndPos:    891,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 40,
										EndLine:   40,
										StartPos:  888,
										EndPos:    891,
									},
									Value: "var",
								},
							},
							Dim: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 40,
									EndLine:   40,
									StartPos:  893,
									EndPos:    893,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 41,
					EndLine:   41,
					StartPos:  900,
					EndPos:    915,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 41,
						EndLine:   41,
						StartPos:  900,
						EndPos:    914,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 41,
								EndLine:   41,
								StartPos:  901,
								EndPos:    905,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 41,
								EndLine:   41,
								StartPos:  906,
								EndPos:    913,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 41,
									EndLine:   41,
									StartPos:  906,
									EndPos:    909,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 41,
										EndLine:   41,
										StartPos:  906,
										EndPos:    909,
									},
									Value: "var",
								},
							},
							Dim: &expr.UnaryMinus{
								Position: &position.Position{
									StartLine: 41,
									EndLine:   41,
									StartPos:  911,
									EndPos:    912,
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 41,
										EndLine:   41,
										StartPos:  911,
										EndPos:    912,
									},
									Value: "1",
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 42,
					EndLine:   42,
					StartPos:  919,
					EndPos:    972,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 42,
						EndLine:   42,
						StartPos:  919,
						EndPos:    971,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 42,
								EndLine:   42,
								StartPos:  920,
								EndPos:    924,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 42,
								EndLine:   42,
								StartPos:  925,
								EndPos:    970,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 42,
									EndLine:   42,
									StartPos:  925,
									EndPos:    928,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 42,
										EndLine:   42,
										StartPos:  925,
										EndPos:    928,
									},
									Value: "var",
								},
							},
							Dim: &scalar.String{
								Position: &position.Position{
									StartLine: 42,
									EndLine:   42,
									StartPos:  930,
									EndPos:    969,
								},
								Value: "1234567890123456789012345678901234567890",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 43,
					EndLine:   43,
					StartPos:  976,
					EndPos:    1030,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 43,
						EndLine:   43,
						StartPos:  976,
						EndPos:    1029,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 43,
								EndLine:   43,
								StartPos:  977,
								EndPos:    981,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 43,
								EndLine:   43,
								StartPos:  982,
								EndPos:    1028,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 43,
									EndLine:   43,
									StartPos:  982,
									EndPos:    985,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 43,
										EndLine:   43,
										StartPos:  982,
										EndPos:    985,
									},
									Value: "var",
								},
							},
							Dim: &scalar.String{
								Position: &position.Position{
									StartLine: 43,
									EndLine:   43,
									StartPos:  987,
									EndPos:    1027,
								},
								Value: "-1234567890123456789012345678901234567890",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 44,
					EndLine:   44,
					StartPos:  1034,
					EndPos:    1050,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 44,
						EndLine:   44,
						StartPos:  1034,
						EndPos:    1049,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 44,
								EndLine:   44,
								StartPos:  1035,
								EndPos:    1039,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 44,
								EndLine:   44,
								StartPos:  1040,
								EndPos:    1048,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 44,
									EndLine:   44,
									StartPos:  1040,
									EndPos:    1043,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 44,
										EndLine:   44,
										StartPos:  1040,
										EndPos:    1043,
									},
									Value: "var",
								},
							},
							Dim: &scalar.String{
								Position: &position.Position{
									StartLine: 44,
									EndLine:   44,
									StartPos:  1045,
									EndPos:    1047,
								},
								Value: "bar",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 45,
					EndLine:   45,
					StartPos:  1054,
					EndPos:    1071,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 45,
						EndLine:   45,
						StartPos:  1054,
						EndPos:    1070,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 45,
								EndLine:   45,
								StartPos:  1055,
								EndPos:    1059,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 45,
								EndLine:   45,
								StartPos:  1060,
								EndPos:    1069,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 45,
									EndLine:   45,
									StartPos:  1060,
									EndPos:    1063,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 45,
										EndLine:   45,
										StartPos:  1060,
										EndPos:    1063,
									},
									Value: "var",
								},
							},
							Dim: &expr.Variable{
								Position: &position.Position{
									StartLine: 45,
									EndLine:   45,
									StartPos:  1065,
									EndPos:    1068,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 45,
										EndLine:   45,
										StartPos:  1065,
										EndPos:    1068,
									},
									Value: "bar",
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 46,
					EndLine:   46,
					StartPos:  1075,
					EndPos:    1086,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 46,
						EndLine:   46,
						StartPos:  1075,
						EndPos:    1085,
					},
					Parts: []node.Node{
						&expr.Variable{
							Position: &position.Position{
								StartLine: 46,
								EndLine:   46,
								StartPos:  1076,
								EndPos:    1079,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 46,
									EndLine:   46,
									StartPos:  1076,
									EndPos:    1079,
								},
								Value: "foo",
							},
						},
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 46,
								EndLine:   46,
								StartPos:  1080,
								EndPos:    1080,
							},
							Value: " ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 46,
								EndLine:   46,
								StartPos:  1081,
								EndPos:    1084,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 46,
									EndLine:   46,
									StartPos:  1081,
									EndPos:    1084,
								},
								Value: "bar",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 47,
					EndLine:   47,
					StartPos:  1090,
					EndPos:    1108,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 47,
						EndLine:   47,
						StartPos:  1090,
						EndPos:    1107,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 47,
								EndLine:   47,
								StartPos:  1091,
								EndPos:    1095,
							},
							Value: "test ",
						},
						&expr.PropertyFetch{
							Position: &position.Position{
								StartLine: 47,
								EndLine:   47,
								StartPos:  1096,
								EndPos:    1104,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 47,
									EndLine:   47,
									StartPos:  1096,
									EndPos:    1099,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 47,
										EndLine:   47,
										StartPos:  1096,
										EndPos:    1099,
									},
									Value: "foo",
								},
							},
							Property: &node.Identifier{
								Position: &position.Position{
									StartLine: 47,
									EndLine:   47,
									StartPos:  1102,
									EndPos:    1104,
								},
								Value: "bar",
							},
						},
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 47,
								EndLine:   47,
								StartPos:  1105,
								EndPos:    1106,
							},
							Value: "()",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 48,
					EndLine:   48,
					StartPos:  1112,
					EndPos:    1125,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 48,
						EndLine:   48,
						StartPos:  1112,
						EndPos:    1124,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 48,
								EndLine:   48,
								StartPos:  1113,
								EndPos:    1117,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 48,
								EndLine:   48,
								StartPos:  1118,
								EndPos:    1123,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 48,
									EndLine:   48,
									StartPos:  1120,
									EndPos:    1122,
								},
								Value: "foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 49,
					EndLine:   49,
					StartPos:  1129,
					EndPos:    1145,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 49,
						EndLine:   49,
						StartPos:  1129,
						EndPos:    1144,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 49,
								EndLine:   49,
								StartPos:  1130,
								EndPos:    1134,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 49,
								EndLine:   49,
								StartPos:  1135,
								EndPos:    1143,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 49,
									EndLine:   49,
									StartPos:  1137,
									EndPos:    1139,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 49,
										EndLine:   49,
										StartPos:  1137,
										EndPos:    1139,
									},
									Value: "foo",
								},
							},
							Dim: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 49,
									EndLine:   49,
									StartPos:  1141,
									EndPos:    1141,
								},
								Value: "0",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 50,
					EndLine:   50,
					StartPos:  1149,
					EndPos:    1163,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 50,
						EndLine:   50,
						StartPos:  1149,
						EndPos:    1162,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 50,
								EndLine:   50,
								StartPos:  1150,
								EndPos:    1154,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 50,
								EndLine:   50,
								StartPos:  1155,
								EndPos:    1161,
							},
							VarName: &expr.Variable{
								Position: &position.Position{
									StartLine: 50,
									EndLine:   50,
									StartPos:  1157,
									EndPos:    1160,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 50,
										EndLine:   50,
										StartPos:  1157,
										EndPos:    1160,
									},
									Value: "foo",
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 51,
					EndLine:   51,
					StartPos:  1167,
					EndPos:    1187,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 51,
						EndLine:   51,
						StartPos:  1167,
						EndPos:    1186,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 51,
								EndLine:   51,
								StartPos:  1168,
								EndPos:    1172,
							},
							Value: "test ",
						},
						&expr.MethodCall{
							Position: &position.Position{
								StartLine: 51,
								EndLine:   51,
								StartPos:  1174,
								EndPos:    1184,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 51,
									EndLine:   51,
									StartPos:  1174,
									EndPos:    1177,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 51,
										EndLine:   51,
										StartPos:  1174,
										EndPos:    1177,
									},
									Value: "foo",
								},
							},
							Method: &node.Identifier{
								Position: &position.Position{
									StartLine: 51,
									EndLine:   51,
									StartPos:  1180,
									EndPos:    1182,
								},
								Value: "bar",
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 51,
									EndLine:   51,
									StartPos:  1183,
									EndPos:    1184,
								},
							},
						},
					},
				},
			},
			&stmt.AltIf{
				Position: &position.Position{
					StartLine: 53,
					EndLine:   54,
					StartPos:  1192,
					EndPos:    1209,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 53,
						EndLine:   53,
						StartPos:  1196,
						EndPos:    1197,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 53,
							EndLine:   53,
							StartPos:  1196,
							EndPos:    1197,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: -1,
						EndLine:   -1,
						StartPos:  -1,
						EndPos:    -1,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.AltIf{
				Position: &position.Position{
					StartLine: 55,
					EndLine:   57,
					StartPos:  1213,
					EndPos:    1245,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 55,
						EndLine:   55,
						StartPos:  1217,
						EndPos:    1218,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 55,
							EndLine:   55,
							StartPos:  1217,
							EndPos:    1218,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: -1,
						EndLine:   -1,
						StartPos:  -1,
						EndPos:    -1,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.AltElseIf{
						Position: &position.Position{
							StartLine: 56,
							EndLine:   -1,
							StartPos:  1225,
							EndPos:    -1,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 56,
								EndLine:   56,
								StartPos:  1233,
								EndPos:    1234,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 56,
									EndLine:   56,
									StartPos:  1233,
									EndPos:    1234,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: -1,
								EndLine:   -1,
								StartPos:  -1,
								EndPos:    -1,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.AltIf{
				Position: &position.Position{
					StartLine: 58,
					EndLine:   60,
					StartPos:  1249,
					EndPos:    1274,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 58,
						EndLine:   58,
						StartPos:  1253,
						EndPos:    1254,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 58,
							EndLine:   58,
							StartPos:  1253,
							EndPos:    1254,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: -1,
						EndLine:   -1,
						StartPos:  -1,
						EndPos:    -1,
					},
					Stmts: []node.Node{},
				},
				Else: &stmt.AltElse{
					Position: &position.Position{
						StartLine: 59,
						EndLine:   -1,
						StartPos:  1261,
						EndPos:    -1,
					},
					Stmt: &stmt.StmtList{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.AltIf{
				Position: &position.Position{
					StartLine: 61,
					EndLine:   65,
					StartPos:  1278,
					EndPos:    1333,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 61,
						EndLine:   61,
						StartPos:  1282,
						EndPos:    1283,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 61,
							EndLine:   61,
							StartPos:  1282,
							EndPos:    1283,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: -1,
						EndLine:   -1,
						StartPos:  -1,
						EndPos:    -1,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.AltElseIf{
						Position: &position.Position{
							StartLine: 62,
							EndLine:   -1,
							StartPos:  1290,
							EndPos:    -1,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 62,
								EndLine:   62,
								StartPos:  1298,
								EndPos:    1299,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 62,
									EndLine:   62,
									StartPos:  1298,
									EndPos:    1299,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: -1,
								EndLine:   -1,
								StartPos:  -1,
								EndPos:    -1,
							},
							Stmts: []node.Node{},
						},
					},
					&stmt.AltElseIf{
						Position: &position.Position{
							StartLine: 63,
							EndLine:   -1,
							StartPos:  1305,
							EndPos:    -1,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 63,
								EndLine:   63,
								StartPos:  1313,
								EndPos:    1314,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 63,
									EndLine:   63,
									StartPos:  1313,
									EndPos:    1314,
								},
								Value: "c",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: -1,
								EndLine:   -1,
								StartPos:  -1,
								EndPos:    -1,
							},
							Stmts: []node.Node{},
						},
					},
				},
				Else: &stmt.AltElse{
					Position: &position.Position{
						StartLine: 64,
						EndLine:   -1,
						StartPos:  1320,
						EndPos:    -1,
					},
					Stmt: &stmt.StmtList{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.While{
				Position: &position.Position{
					StartLine: 67,
					EndLine:   67,
					StartPos:  1338,
					EndPos:    1357,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 67,
						EndLine:   67,
						StartPos:  1345,
						EndPos:    1345,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 67,
						EndLine:   67,
						StartPos:  1348,
						EndPos:    1357,
					},
					Stmts: []node.Node{
						&stmt.Break{
							Position: &position.Position{
								StartLine: 67,
								EndLine:   67,
								StartPos:  1350,
								EndPos:    1355,
							},
						},
					},
				},
			},
			&stmt.While{
				Position: &position.Position{
					StartLine: 68,
					EndLine:   68,
					StartPos:  1361,
					EndPos:    1382,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 68,
						EndLine:   68,
						StartPos:  1368,
						EndPos:    1368,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 68,
						EndLine:   68,
						StartPos:  1371,
						EndPos:    1382,
					},
					Stmts: []node.Node{
						&stmt.Break{
							Position: &position.Position{
								StartLine: 68,
								EndLine:   68,
								StartPos:  1373,
								EndPos:    1380,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 68,
									EndLine:   68,
									StartPos:  1379,
									EndPos:    1379,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.AltWhile{
				Position: &position.Position{
					StartLine: 69,
					EndLine:   69,
					StartPos:  1386,
					EndPos:    1416,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 69,
						EndLine:   69,
						StartPos:  1393,
						EndPos:    1393,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 69,
						EndLine:   69,
						StartPos:  1398,
						EndPos:    1406,
					},
					Stmts: []node.Node{
						&stmt.Break{
							Position: &position.Position{
								StartLine: 69,
								EndLine:   69,
								StartPos:  1398,
								EndPos:    1406,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 69,
									EndLine:   69,
									StartPos:  1404,
									EndPos:    1404,
								},
								Value: "3",
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 70,
					EndLine:   70,
					StartPos:  1420,
					EndPos:    1462,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 70,
						EndLine:   70,
						StartPos:  1426,
						EndPos:    1428,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassConstList{
						Position: &position.Position{
							StartLine: 70,
							EndLine:   70,
							StartPos:  1431,
							EndPos:    1460,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 70,
									EndLine:   70,
									StartPos:  1431,
									EndPos:    1436,
								},
								Value: "public",
							},
						},
						Consts: []node.Node{
							&stmt.Constant{
								Position: &position.Position{
									StartLine: 70,
									EndLine:   70,
									StartPos:  1444,
									EndPos:    1450,
								},
								PhpDocComment: "",
								ConstantName: &node.Identifier{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1444,
										EndPos:    1446,
									},
									Value: "FOO",
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1450,
										EndPos:    1450,
									},
									Value: "1",
								},
							},
							&stmt.Constant{
								Position: &position.Position{
									StartLine: 70,
									EndLine:   70,
									StartPos:  1453,
									EndPos:    1459,
								},
								PhpDocComment: "",
								ConstantName: &node.Identifier{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1453,
										EndPos:    1455,
									},
									Value: "BAR",
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1459,
										EndPos:    1459,
									},
									Value: "2",
								},
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 71,
					EndLine:   71,
					StartPos:  1466,
					EndPos:    1501,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 71,
						EndLine:   71,
						StartPos:  1472,
						EndPos:    1474,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassConstList{
						Position: &position.Position{
							StartLine: 71,
							EndLine:   71,
							StartPos:  1477,
							EndPos:    1499,
						},
						Consts: []node.Node{
							&stmt.Constant{
								Position: &position.Position{
									StartLine: 71,
									EndLine:   71,
									StartPos:  1483,
									EndPos:    1489,
								},
								PhpDocComment: "",
								ConstantName: &node.Identifier{
									Position: &position.Position{
										StartLine: 71,
										EndLine:   71,
										StartPos:  1483,
										EndPos:    1485,
									},
									Value: "FOO",
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 71,
										EndLine:   71,
										StartPos:  1489,
										EndPos:    1489,
									},
									Value: "1",
								},
							},
							&stmt.Constant{
								Position: &position.Position{
									StartLine: 71,
									EndLine:   71,
									StartPos:  1492,
									EndPos:    1498,
								},
								PhpDocComment: "",
								ConstantName: &node.Identifier{
									Position: &position.Position{
										StartLine: 71,
										EndLine:   71,
										StartPos:  1492,
										EndPos:    1494,
									},
									Value: "BAR",
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 71,
										EndLine:   71,
										StartPos:  1498,
										EndPos:    1498,
									},
									Value: "2",
								},
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 72,
					EndLine:   72,
					StartPos:  1505,
					EndPos:    1534,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 72,
						EndLine:   72,
						StartPos:  1511,
						EndPos:    1513,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 72,
							EndLine:   72,
							StartPos:  1516,
							EndPos:    1532,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 72,
								EndLine:   72,
								StartPos:  1525,
								EndPos:    1527,
							},
							Value: "bar",
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 72,
								EndLine:   72,
								StartPos:  1531,
								EndPos:    1532,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 73,
					EndLine:   73,
					StartPos:  1538,
					EndPos:    1582,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 73,
						EndLine:   73,
						StartPos:  1544,
						EndPos:    1546,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 73,
							EndLine:   73,
							StartPos:  1549,
							EndPos:    1580,
						},
						ReturnsRef:    true,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 73,
								EndLine:   73,
								StartPos:  1573,
								EndPos:    1575,
							},
							Value: "bar",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 73,
									EndLine:   73,
									StartPos:  1549,
									EndPos:    1554,
								},
								Value: "public",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 73,
									EndLine:   73,
									StartPos:  1556,
									EndPos:    1561,
								},
								Value: "static",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 73,
								EndLine:   73,
								StartPos:  1579,
								EndPos:    1580,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 74,
					EndLine:   74,
					StartPos:  1586,
					EndPos:    1636,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 74,
						EndLine:   74,
						StartPos:  1592,
						EndPos:    1594,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 74,
							EndLine:   74,
							StartPos:  1597,
							EndPos:    1634,
						},
						PhpDocComment: "",
						ReturnsRef:    true,
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 74,
								EndLine:   74,
								StartPos:  1621,
								EndPos:    1623,
							},
							Value: "bar",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 74,
									EndLine:   74,
									StartPos:  1597,
									EndPos:    1602,
								},
								Value: "public",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 74,
									EndLine:   74,
									StartPos:  1604,
									EndPos:    1609,
								},
								Value: "static",
							},
						},
						ReturnType: &name.Name{
							Position: &position.Position{
								StartLine: 74,
								EndLine:   74,
								StartPos:  1628,
								EndPos:    1631,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 74,
										EndLine:   74,
										StartPos:  1628,
										EndPos:    1631,
									},
									Value: "void",
								},
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 74,
								EndLine:   74,
								StartPos:  1633,
								EndPos:    1634,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 75,
					EndLine:   75,
					StartPos:  1640,
					EndPos:    1660,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 75,
						EndLine:   75,
						StartPos:  1655,
						EndPos:    1657,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 75,
							EndLine:   75,
							StartPos:  1640,
							EndPos:    1647,
						},
						Value: "abstract",
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 76,
					EndLine:   76,
					StartPos:  1664,
					EndPos:    1694,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 76,
						EndLine:   76,
						StartPos:  1676,
						EndPos:    1678,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 76,
							EndLine:   76,
							StartPos:  1664,
							EndPos:    1668,
						},
						Value: "final",
					},
				},
				Extends: &stmt.ClassExtends{
					Position: &position.Position{
						StartLine: 76,
						EndLine:   76,
						StartPos:  1680,
						EndPos:    1690,
					},
					ClassName: &name.Name{
						Position: &position.Position{
							StartLine: 76,
							EndLine:   76,
							StartPos:  1688,
							EndPos:    1690,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 76,
									EndLine:   76,
									StartPos:  1688,
									EndPos:    1690,
								},
								Value: "bar",
							},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 77,
					EndLine:   77,
					StartPos:  1698,
					EndPos:    1731,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 77,
						EndLine:   77,
						StartPos:  1710,
						EndPos:    1712,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 77,
							EndLine:   77,
							StartPos:  1698,
							EndPos:    1702,
						},
						Value: "final",
					},
				},
				Implements: &stmt.ClassImplements{
					Position: &position.Position{
						StartLine: 77,
						EndLine:   77,
						StartPos:  1714,
						EndPos:    1727,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1725,
								EndPos:    1727,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 77,
										EndLine:   77,
										StartPos:  1725,
										EndPos:    1727,
									},
									Value: "bar",
								},
							},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 78,
					EndLine:   78,
					StartPos:  1735,
					EndPos:    1773,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 78,
						EndLine:   78,
						StartPos:  1747,
						EndPos:    1749,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 78,
							EndLine:   78,
							StartPos:  1735,
							EndPos:    1739,
						},
						Value: "final",
					},
				},
				Implements: &stmt.ClassImplements{
					Position: &position.Position{
						StartLine: 78,
						EndLine:   78,
						StartPos:  1751,
						EndPos:    1769,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 78,
								EndLine:   78,
								StartPos:  1762,
								EndPos:    1764,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 78,
										EndLine:   78,
										StartPos:  1762,
										EndPos:    1764,
									},
									Value: "bar",
								},
							},
						},
						&name.Name{
							Position: &position.Position{
								StartLine: 78,
								EndLine:   78,
								StartPos:  1767,
								EndPos:    1769,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 78,
										EndLine:   78,
										StartPos:  1767,
										EndPos:    1769,
									},
									Value: "baz",
								},
							},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 79,
					EndLine:   79,
					StartPos:  1777,
					EndPos:    1824,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 79,
						EndLine:   79,
						StartPos:  1777,
						EndPos:    1823,
					},
					Class: &stmt.Class{
						Position: &position.Position{
							StartLine: 79,
							EndLine:   79,
							StartPos:  1781,
							EndPos:    1823,
						},
						PhpDocComment: "",
						ArgumentList: &node.ArgumentList{
							Position: &position.Position{
								StartLine: 79,
								EndLine:   79,
								StartPos:  1786,
								EndPos:    1787,
							},
						},
						Extends: &stmt.ClassExtends{
							Position: &position.Position{
								StartLine: 79,
								EndLine:   79,
								StartPos:  1789,
								EndPos:    1799,
							},
							ClassName: &name.Name{
								Position: &position.Position{
									StartLine: 79,
									EndLine:   79,
									StartPos:  1797,
									EndPos:    1799,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 79,
											EndLine:   79,
											StartPos:  1797,
											EndPos:    1799,
										},
										Value: "foo",
									},
								},
							},
						},
						Implements: &stmt.ClassImplements{
							Position: &position.Position{
								StartLine: 79,
								EndLine:   79,
								StartPos:  1801,
								EndPos:    1819,
							},
							InterfaceNames: []node.Node{
								&name.Name{
									Position: &position.Position{
										StartLine: 79,
										EndLine:   79,
										StartPos:  1812,
										EndPos:    1814,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 79,
												EndLine:   79,
												StartPos:  1812,
												EndPos:    1814,
											},
											Value: "bar",
										},
									},
								},
								&name.Name{
									Position: &position.Position{
										StartLine: 79,
										EndLine:   79,
										StartPos:  1817,
										EndPos:    1819,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 79,
												EndLine:   79,
												StartPos:  1817,
												EndPos:    1819,
											},
											Value: "baz",
										},
									},
								},
							},
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.ConstList{
				Position: &position.Position{
					StartLine: 81,
					EndLine:   81,
					StartPos:  1829,
					EndPos:    1851,
				},
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 81,
							EndLine:   81,
							StartPos:  1835,
							EndPos:    1841,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1835,
								EndPos:    1837,
							},
							Value: "FOO",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1841,
								EndPos:    1841,
							},
							Value: "1",
						},
					},
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 81,
							EndLine:   81,
							StartPos:  1844,
							EndPos:    1850,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1844,
								EndPos:    1846,
							},
							Value: "BAR",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1850,
								EndPos:    1850,
							},
							Value: "2",
						},
					},
				},
			},
			&stmt.While{
				Position: &position.Position{
					StartLine: 82,
					EndLine:   82,
					StartPos:  1855,
					EndPos:    1877,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 82,
						EndLine:   82,
						StartPos:  1862,
						EndPos:    1862,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 82,
						EndLine:   82,
						StartPos:  1865,
						EndPos:    1877,
					},
					Stmts: []node.Node{
						&stmt.Continue{
							Position: &position.Position{
								StartLine: 82,
								EndLine:   82,
								StartPos:  1867,
								EndPos:    1875,
							},
						},
					},
				},
			},
			&stmt.While{
				Position: &position.Position{
					StartLine: 83,
					EndLine:   83,
					StartPos:  1881,
					EndPos:    1905,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 83,
						EndLine:   83,
						StartPos:  1888,
						EndPos:    1888,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 83,
						EndLine:   83,
						StartPos:  1891,
						EndPos:    1905,
					},
					Stmts: []node.Node{
						&stmt.Continue{
							Position: &position.Position{
								StartLine: 83,
								EndLine:   83,
								StartPos:  1893,
								EndPos:    1903,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 83,
									EndLine:   83,
									StartPos:  1902,
									EndPos:    1902,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.While{
				Position: &position.Position{
					StartLine: 84,
					EndLine:   84,
					StartPos:  1909,
					EndPos:    1934,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 84,
						EndLine:   84,
						StartPos:  1916,
						EndPos:    1916,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 84,
						EndLine:   84,
						StartPos:  1919,
						EndPos:    1934,
					},
					Stmts: []node.Node{
						&stmt.Continue{
							Position: &position.Position{
								StartLine: 84,
								EndLine:   84,
								StartPos:  1921,
								EndPos:    1932,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 84,
									EndLine:   84,
									StartPos:  1930,
									EndPos:    1930,
								},
								Value: "3",
							},
						},
					},
				},
			},
			&stmt.Declare{
				Position: &position.Position{
					StartLine: 85,
					EndLine:   85,
					StartPos:  1938,
					EndPos:    1954,
				},
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 85,
							EndLine:   85,
							StartPos:  1946,
							EndPos:    1952,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 85,
								EndLine:   85,
								StartPos:  1946,
								EndPos:    1950,
							},
							Value: "ticks",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 85,
								EndLine:   85,
								StartPos:  1952,
								EndPos:    1952,
							},
							Value: "1",
						},
					},
				},
				Stmt: &stmt.Nop{
					Position: &position.Position{
						StartLine: 85,
						EndLine:   85,
						StartPos:  1954,
						EndPos:    1954,
					},
				},
			},
			&stmt.Declare{
				Position: &position.Position{
					StartLine: 86,
					EndLine:   86,
					StartPos:  1958,
					EndPos:    1976,
				},
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 86,
							EndLine:   86,
							StartPos:  1966,
							EndPos:    1972,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 86,
								EndLine:   86,
								StartPos:  1966,
								EndPos:    1970,
							},
							Value: "ticks",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 86,
								EndLine:   86,
								StartPos:  1972,
								EndPos:    1972,
							},
							Value: "1",
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 86,
						EndLine:   86,
						StartPos:  1975,
						EndPos:    1976,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Declare{
				Position: &position.Position{
					StartLine: 87,
					EndLine:   87,
					StartPos:  1980,
					EndPos:    2008,
				},
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 87,
							EndLine:   87,
							StartPos:  1988,
							EndPos:    1994,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 87,
								EndLine:   87,
								StartPos:  1988,
								EndPos:    1992,
							},
							Value: "ticks",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 87,
								EndLine:   87,
								StartPos:  1994,
								EndPos:    1994,
							},
							Value: "1",
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 87,
						EndLine:   87,
						StartPos:  1996,
						EndPos:    2008,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Do{
				Position: &position.Position{
					StartLine: 88,
					EndLine:   88,
					StartPos:  2012,
					EndPos:    2026,
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 88,
						EndLine:   88,
						StartPos:  2015,
						EndPos:    2016,
					},
					Stmts: []node.Node{},
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 88,
						EndLine:   88,
						StartPos:  2024,
						EndPos:    2024,
					},
					Value: "1",
				},
			},
			&stmt.Echo{
				Position: &position.Position{
					StartLine: 89,
					EndLine:   89,
					StartPos:  2030,
					EndPos:    2040,
				},
				Exprs: []node.Node{
					&expr.Variable{
						Position: &position.Position{
							StartLine: 89,
							EndLine:   89,
							StartPos:  2035,
							EndPos:    2036,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 89,
								EndLine:   89,
								StartPos:  2035,
								EndPos:    2036,
							},
							Value: "a",
						},
					},
					&scalar.Lnumber{
						Position: &position.Position{
							StartLine: 89,
							EndLine:   89,
							StartPos:  2039,
							EndPos:    2039,
						},
						Value: "1",
					},
				},
			},
			&stmt.Echo{
				Position: &position.Position{
					StartLine: 90,
					EndLine:   90,
					StartPos:  2044,
					EndPos:    2052,
				},
				Exprs: []node.Node{
					&expr.Variable{
						Position: &position.Position{
							StartLine: 90,
							EndLine:   90,
							StartPos:  2049,
							EndPos:    2050,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 90,
								EndLine:   90,
								StartPos:  2049,
								EndPos:    2050,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.For{
				Position: &position.Position{
					StartLine: 91,
					EndLine:   91,
					StartPos:  2056,
					EndPos:    2090,
				},
				Init: []node.Node{
					&assign.Assign{
						Position: &position.Position{
							StartLine: 91,
							EndLine:   91,
							StartPos:  2060,
							EndPos:    2065,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2060,
								EndPos:    2061,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2060,
									EndPos:    2061,
								},
								Value: "i",
							},
						},
						Expression: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2065,
								EndPos:    2065,
							},
							Value: "0",
						},
					},
				},
				Cond: []node.Node{
					&binary.Smaller{
						Position: &position.Position{
							StartLine: 91,
							EndLine:   91,
							StartPos:  2068,
							EndPos:    2074,
						},
						Left: &expr.Variable{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2068,
								EndPos:    2069,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2068,
									EndPos:    2069,
								},
								Value: "i",
							},
						},
						Right: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2073,
								EndPos:    2074,
							},
							Value: "10",
						},
					},
				},
				Loop: []node.Node{
					&expr.PostInc{
						Position: &position.Position{
							StartLine: 91,
							EndLine:   91,
							StartPos:  2077,
							EndPos:    2080,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2077,
								EndPos:    2078,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2077,
									EndPos:    2078,
								},
								Value: "i",
							},
						},
					},
					&expr.PostInc{
						Position: &position.Position{
							StartLine: 91,
							EndLine:   91,
							StartPos:  2083,
							EndPos:    2086,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2083,
								EndPos:    2084,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2083,
									EndPos:    2084,
								},
								Value: "i",
							},
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 91,
						EndLine:   91,
						StartPos:  2089,
						EndPos:    2090,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.AltFor{
				Position: &position.Position{
					StartLine: 92,
					EndLine:   92,
					StartPos:  2094,
					EndPos:    2129,
				},
				Cond: []node.Node{
					&binary.Smaller{
						Position: &position.Position{
							StartLine: 92,
							EndLine:   92,
							StartPos:  2100,
							EndPos:    2106,
						},
						Left: &expr.Variable{
							Position: &position.Position{
								StartLine: 92,
								EndLine:   92,
								StartPos:  2100,
								EndPos:    2101,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  2100,
									EndPos:    2101,
								},
								Value: "i",
							},
						},
						Right: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 92,
								EndLine:   92,
								StartPos:  2105,
								EndPos:    2106,
							},
							Value: "10",
						},
					},
				},
				Loop: []node.Node{
					&expr.PostInc{
						Position: &position.Position{
							StartLine: 92,
							EndLine:   92,
							StartPos:  2109,
							EndPos:    2112,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 92,
								EndLine:   92,
								StartPos:  2109,
								EndPos:    2110,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  2109,
									EndPos:    2110,
								},
								Value: "i",
							},
						},
					},
					&expr.PostInc{
						Position: &position.Position{
							StartLine: 92,
							EndLine:   92,
							StartPos:  2115,
							EndPos:    2118,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 92,
								EndLine:   92,
								StartPos:  2115,
								EndPos:    2116,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  2115,
									EndPos:    2116,
								},
								Value: "i",
							},
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: -1,
						EndLine:   -1,
						StartPos:  -1,
						EndPos:    -1,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Foreach{
				Position: &position.Position{
					StartLine: 93,
					EndLine:   93,
					StartPos:  2133,
					EndPos:    2153,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 93,
						EndLine:   93,
						StartPos:  2142,
						EndPos:    2143,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 93,
							EndLine:   93,
							StartPos:  2142,
							EndPos:    2143,
						},
						Value: "a",
					},
				},
				Variable: &expr.Variable{
					Position: &position.Position{
						StartLine: 93,
						EndLine:   93,
						StartPos:  2148,
						EndPos:    2149,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 93,
							EndLine:   93,
							StartPos:  2148,
							EndPos:    2149,
						},
						Value: "v",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 93,
						EndLine:   93,
						StartPos:  2152,
						EndPos:    2153,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.AltForeach{
				Position: &position.Position{
					StartLine: 94,
					EndLine:   94,
					StartPos:  2157,
					EndPos:    2188,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 94,
						EndLine:   94,
						StartPos:  2166,
						EndPos:    2167,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 94,
							EndLine:   94,
							StartPos:  2166,
							EndPos:    2167,
						},
						Value: "a",
					},
				},
				Variable: &expr.Variable{
					Position: &position.Position{
						StartLine: 94,
						EndLine:   94,
						StartPos:  2172,
						EndPos:    2173,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 94,
							EndLine:   94,
							StartPos:  2172,
							EndPos:    2173,
						},
						Value: "v",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: -1,
						EndLine:   -1,
						StartPos:  -1,
						EndPos:    -1,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Foreach{
				Position: &position.Position{
					StartLine: 95,
					EndLine:   95,
					StartPos:  2192,
					EndPos:    2218,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 95,
						EndLine:   95,
						StartPos:  2201,
						EndPos:    2202,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2201,
							EndPos:    2202,
						},
						Value: "a",
					},
				},
				Key: &expr.Variable{
					Position: &position.Position{
						StartLine: 95,
						EndLine:   95,
						StartPos:  2207,
						EndPos:    2208,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2207,
							EndPos:    2208,
						},
						Value: "k",
					},
				},
				Variable: &expr.Variable{
					Position: &position.Position{
						StartLine: 95,
						EndLine:   95,
						StartPos:  2213,
						EndPos:    2214,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2213,
							EndPos:    2214,
						},
						Value: "v",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 95,
						EndLine:   95,
						StartPos:  2217,
						EndPos:    2218,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Foreach{
				Position: &position.Position{
					StartLine: 96,
					EndLine:   96,
					StartPos:  2222,
					EndPos:    2249,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 96,
						EndLine:   96,
						StartPos:  2231,
						EndPos:    2232,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2231,
							EndPos:    2232,
						},
						Value: "a",
					},
				},
				Key: &expr.Variable{
					Position: &position.Position{
						StartLine: 96,
						EndLine:   96,
						StartPos:  2237,
						EndPos:    2238,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2237,
							EndPos:    2238,
						},
						Value: "k",
					},
				},
				Variable: &expr.Reference{
					Position: &position.Position{
						StartLine: 96,
						EndLine:   96,
						StartPos:  2243,
						EndPos:    2245,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2244,
							EndPos:    2245,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 96,
								EndLine:   96,
								StartPos:  2244,
								EndPos:    2245,
							},
							Value: "v",
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 96,
						EndLine:   96,
						StartPos:  2248,
						EndPos:    2249,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Foreach{
				Position: &position.Position{
					StartLine: 97,
					EndLine:   97,
					StartPos:  2253,
					EndPos:    2285,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 97,
						EndLine:   97,
						StartPos:  2262,
						EndPos:    2263,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 97,
							EndLine:   97,
							StartPos:  2262,
							EndPos:    2263,
						},
						Value: "a",
					},
				},
				Key: &expr.Variable{
					Position: &position.Position{
						StartLine: 97,
						EndLine:   97,
						StartPos:  2268,
						EndPos:    2269,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 97,
							EndLine:   97,
							StartPos:  2268,
							EndPos:    2269,
						},
						Value: "k",
					},
				},
				Variable: &expr.List{
					Position: &position.Position{
						StartLine: 97,
						EndLine:   97,
						StartPos:  2274,
						EndPos:    2281,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2279,
								EndPos:    2280,
							},
							Val: &expr.Variable{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2279,
									EndPos:    2280,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 97,
										EndLine:   97,
										StartPos:  2279,
										EndPos:    2280,
									},
									Value: "v",
								},
							},
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 97,
						EndLine:   97,
						StartPos:  2284,
						EndPos:    2285,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Foreach{
				Position: &position.Position{
					StartLine: 98,
					EndLine:   98,
					StartPos:  2289,
					EndPos:    2317,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 98,
						EndLine:   98,
						StartPos:  2298,
						EndPos:    2299,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2298,
							EndPos:    2299,
						},
						Value: "a",
					},
				},
				Key: &expr.Variable{
					Position: &position.Position{
						StartLine: 98,
						EndLine:   98,
						StartPos:  2304,
						EndPos:    2305,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2304,
							EndPos:    2305,
						},
						Value: "k",
					},
				},
				Variable: &expr.ShortList{
					Position: &position.Position{
						StartLine: 98,
						EndLine:   98,
						StartPos:  2310,
						EndPos:    2313,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 98,
								EndLine:   98,
								StartPos:  2311,
								EndPos:    2312,
							},
							Val: &expr.Variable{
								Position: &position.Position{
									StartLine: 98,
									EndLine:   98,
									StartPos:  2311,
									EndPos:    2312,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 98,
										EndLine:   98,
										StartPos:  2311,
										EndPos:    2312,
									},
									Value: "v",
								},
							},
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 98,
						EndLine:   98,
						StartPos:  2316,
						EndPos:    2317,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 99,
					EndLine:   99,
					StartPos:  2321,
					EndPos:    2337,
				},
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 99,
						EndLine:   99,
						StartPos:  2330,
						EndPos:    2332,
					},
					Value: "foo",
				},
				Stmts: []node.Node{},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 100,
					EndLine:   100,
					StartPos:  2341,
					EndPos:    2364,
				},
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 100,
						EndLine:   100,
						StartPos:  2350,
						EndPos:    2352,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.Return{
						Position: &position.Position{
							StartLine: 100,
							EndLine:   100,
							StartPos:  2357,
							EndPos:    2363,
						},
					},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 101,
					EndLine:   101,
					StartPos:  2368,
					EndPos:    2394,
				},
				ReturnsRef:    true,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 101,
						EndLine:   101,
						StartPos:  2378,
						EndPos:    2380,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.Return{
						Position: &position.Position{
							StartLine: 101,
							EndLine:   101,
							StartPos:  2385,
							EndPos:    2393,
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 101,
								EndLine:   101,
								StartPos:  2392,
								EndPos:    2392,
							},
							Value: "1",
						},
					},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 102,
					EndLine:   102,
					StartPos:  2398,
					EndPos:    2421,
				},
				ReturnsRef:    true,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 102,
						EndLine:   102,
						StartPos:  2408,
						EndPos:    2410,
					},
					Value: "foo",
				},
				ReturnType: &name.Name{
					Position: &position.Position{
						StartLine: 102,
						EndLine:   102,
						StartPos:  2415,
						EndPos:    2418,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 102,
								EndLine:   102,
								StartPos:  2415,
								EndPos:    2418,
							},
							Value: "void",
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Global{
				Position: &position.Position{
					StartLine: 103,
					EndLine:   103,
					StartPos:  2425,
					EndPos:    2438,
				},
				Vars: []node.Node{
					&expr.Variable{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2432,
							EndPos:    2433,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2432,
								EndPos:    2433,
							},
							Value: "a",
						},
					},
					&expr.Variable{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2436,
							EndPos:    2437,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2436,
								EndPos:    2437,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Label{
				Position: &position.Position{
					StartLine: 104,
					EndLine:   104,
					StartPos:  2442,
					EndPos:    2443,
				},
				LabelName: &node.Identifier{
					Position: &position.Position{
						StartLine: 104,
						EndLine:   104,
						StartPos:  2442,
						EndPos:    2442,
					},
					Value: "a",
				},
			},
			&stmt.Goto{
				Position: &position.Position{
					StartLine: 105,
					EndLine:   105,
					StartPos:  2448,
					EndPos:    2454,
				},
				Label: &node.Identifier{
					Position: &position.Position{
						StartLine: 105,
						EndLine:   105,
						StartPos:  2453,
						EndPos:    2453,
					},
					Value: "a",
				},
			},
			&stmt.HaltCompiler{
				Position: &position.Position{
					StartLine: 106,
					EndLine:   106,
					StartPos:  2458,
					EndPos:    2475,
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 107,
					EndLine:   107,
					StartPos:  2479,
					EndPos:    2488,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 107,
						EndLine:   107,
						StartPos:  2483,
						EndPos:    2484,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 107,
							EndLine:   107,
							StartPos:  2483,
							EndPos:    2484,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 107,
						EndLine:   107,
						StartPos:  2487,
						EndPos:    2488,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 108,
					EndLine:   108,
					StartPos:  2492,
					EndPos:    2516,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 108,
						EndLine:   108,
						StartPos:  2496,
						EndPos:    2497,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 108,
							EndLine:   108,
							StartPos:  2496,
							EndPos:    2497,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 108,
						EndLine:   108,
						StartPos:  2500,
						EndPos:    2501,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 108,
							EndLine:   108,
							StartPos:  2503,
							EndPos:    2516,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 108,
								EndLine:   108,
								StartPos:  2511,
								EndPos:    2512,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 108,
									EndLine:   108,
									StartPos:  2511,
									EndPos:    2512,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 108,
								EndLine:   108,
								StartPos:  2515,
								EndPos:    2516,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 109,
					EndLine:   109,
					StartPos:  2520,
					EndPos:    2537,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 109,
						EndLine:   109,
						StartPos:  2524,
						EndPos:    2525,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 109,
							EndLine:   109,
							StartPos:  2524,
							EndPos:    2525,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 109,
						EndLine:   109,
						StartPos:  2528,
						EndPos:    2529,
					},
					Stmts: []node.Node{},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 109,
						EndLine:   109,
						StartPos:  2531,
						EndPos:    2537,
					},
					Stmt: &stmt.StmtList{
						Position: &position.Position{
							StartLine: 109,
							EndLine:   109,
							StartPos:  2536,
							EndPos:    2537,
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 110,
					EndLine:   110,
					StartPos:  2541,
					EndPos:    2588,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 110,
						EndLine:   110,
						StartPos:  2545,
						EndPos:    2546,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2545,
							EndPos:    2546,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 110,
						EndLine:   110,
						StartPos:  2549,
						EndPos:    2550,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2552,
							EndPos:    2565,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2560,
								EndPos:    2561,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2560,
									EndPos:    2561,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2564,
								EndPos:    2565,
							},
							Stmts: []node.Node{},
						},
					},
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2567,
							EndPos:    2580,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2575,
								EndPos:    2576,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2575,
									EndPos:    2576,
								},
								Value: "c",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2579,
								EndPos:    2580,
							},
							Stmts: []node.Node{},
						},
					},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 110,
						EndLine:   110,
						StartPos:  2582,
						EndPos:    2588,
					},
					Stmt: &stmt.StmtList{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2587,
							EndPos:    2588,
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 111,
					EndLine:   111,
					StartPos:  2592,
					EndPos:    2640,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 111,
						EndLine:   111,
						StartPos:  2596,
						EndPos:    2597,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 111,
							EndLine:   111,
							StartPos:  2596,
							EndPos:    2597,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 111,
						EndLine:   111,
						StartPos:  2600,
						EndPos:    2601,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 111,
							EndLine:   111,
							StartPos:  2603,
							EndPos:    2616,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 111,
								EndLine:   111,
								StartPos:  2611,
								EndPos:    2612,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 111,
									EndLine:   111,
									StartPos:  2611,
									EndPos:    2612,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 111,
								EndLine:   111,
								StartPos:  2615,
								EndPos:    2616,
							},
							Stmts: []node.Node{},
						},
					},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 111,
						EndLine:   111,
						StartPos:  2618,
						EndPos:    2640,
					},
					Stmt: &stmt.If{
						Position: &position.Position{
							StartLine: 111,
							EndLine:   111,
							StartPos:  2623,
							EndPos:    2640,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 111,
								EndLine:   111,
								StartPos:  2627,
								EndPos:    2628,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 111,
									EndLine:   111,
									StartPos:  2627,
									EndPos:    2628,
								},
								Value: "c",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 111,
								EndLine:   111,
								StartPos:  2631,
								EndPos:    2632,
							},
							Stmts: []node.Node{},
						},
						Else: &stmt.Else{
							Position: &position.Position{
								StartLine: 111,
								EndLine:   111,
								StartPos:  2634,
								EndPos:    2640,
							},
							Stmt: &stmt.StmtList{
								Position: &position.Position{
									StartLine: 111,
									EndLine:   111,
									StartPos:  2639,
									EndPos:    2640,
								},
								Stmts: []node.Node{},
							},
						},
					},
				},
			},
			&stmt.Nop{
				Position: &position.Position{
					StartLine: 112,
					EndLine:   112,
					StartPos:  2644,
					EndPos:    2645,
				},
			},
			&stmt.InlineHtml{
				Position: &position.Position{
					StartLine: 112,
					EndLine:   112,
					StartPos:  2647,
					EndPos:    2658,
				},
				Value: "<div></div> ",
			},
			&stmt.Interface{
				Position: &position.Position{
					StartLine: 113,
					EndLine:   113,
					StartPos:  2664,
					EndPos:    2679,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 113,
						EndLine:   113,
						StartPos:  2674,
						EndPos:    2676,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{},
			},
			&stmt.Interface{
				Position: &position.Position{
					StartLine: 114,
					EndLine:   114,
					StartPos:  2683,
					EndPos:    2710,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 114,
						EndLine:   114,
						StartPos:  2693,
						EndPos:    2695,
					},
					Value: "Foo",
				},
				Extends: &stmt.InterfaceExtends{
					Position: &position.Position{
						StartLine: 114,
						EndLine:   114,
						StartPos:  2697,
						EndPos:    2707,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 114,
								EndLine:   114,
								StartPos:  2705,
								EndPos:    2707,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2705,
										EndPos:    2707,
									},
									Value: "Bar",
								},
							},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Interface{
				Position: &position.Position{
					StartLine: 115,
					EndLine:   115,
					StartPos:  2714,
					EndPos:    2746,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 115,
						EndLine:   115,
						StartPos:  2724,
						EndPos:    2726,
					},
					Value: "Foo",
				},
				Extends: &stmt.InterfaceExtends{
					Position: &position.Position{
						StartLine: 115,
						EndLine:   115,
						StartPos:  2728,
						EndPos:    2743,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 115,
								EndLine:   115,
								StartPos:  2736,
								EndPos:    2738,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 115,
										EndLine:   115,
										StartPos:  2736,
										EndPos:    2738,
									},
									Value: "Bar",
								},
							},
						},
						&name.Name{
							Position: &position.Position{
								StartLine: 115,
								EndLine:   115,
								StartPos:  2741,
								EndPos:    2743,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 115,
										EndLine:   115,
										StartPos:  2741,
										EndPos:    2743,
									},
									Value: "Baz",
								},
							},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Namespace{
				Position: &position.Position{
					StartLine: 116,
					EndLine:   116,
					StartPos:  2750,
					EndPos:    2763,
				},
				NamespaceName: &name.Name{
					Position: &position.Position{
						StartLine: 116,
						EndLine:   116,
						StartPos:  2760,
						EndPos:    2762,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 116,
								EndLine:   116,
								StartPos:  2760,
								EndPos:    2762,
							},
							Value: "Foo",
						},
					},
				},
			},
			&stmt.Namespace{
				Position: &position.Position{
					StartLine: 117,
					EndLine:   117,
					StartPos:  2767,
					EndPos:    2782,
				},
				NamespaceName: &name.Name{
					Position: &position.Position{
						StartLine: 117,
						EndLine:   117,
						StartPos:  2777,
						EndPos:    2779,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 117,
								EndLine:   117,
								StartPos:  2777,
								EndPos:    2779,
							},
							Value: "Foo",
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Namespace{
				Position: &position.Position{
					StartLine: 118,
					EndLine:   118,
					StartPos:  2786,
					EndPos:    2797,
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 119,
					EndLine:   119,
					StartPos:  2801,
					EndPos:    2819,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 119,
						EndLine:   119,
						StartPos:  2807,
						EndPos:    2809,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Position: &position.Position{
							StartLine: 119,
							EndLine:   119,
							StartPos:  2812,
							EndPos:    2818,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 119,
									EndLine:   119,
									StartPos:  2812,
									EndPos:    2814,
								},
								Value: "var",
							},
						},
						Properties: []node.Node{
							&stmt.Property{
								Position: &position.Position{
									StartLine: 119,
									EndLine:   119,
									StartPos:  2816,
									EndPos:    2817,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 119,
										EndLine:   119,
										StartPos:  2816,
										EndPos:    2817,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 119,
											EndLine:   119,
											StartPos:  2816,
											EndPos:    2817,
										},
										Value: "a",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 120,
					EndLine:   120,
					StartPos:  2823,
					EndPos:    2859,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 120,
						EndLine:   120,
						StartPos:  2829,
						EndPos:    2831,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Position: &position.Position{
							StartLine: 120,
							EndLine:   120,
							StartPos:  2834,
							EndPos:    2858,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 120,
									EndLine:   120,
									StartPos:  2834,
									EndPos:    2839,
								},
								Value: "public",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 120,
									EndLine:   120,
									StartPos:  2841,
									EndPos:    2846,
								},
								Value: "static",
							},
						},
						Properties: []node.Node{
							&stmt.Property{
								Position: &position.Position{
									StartLine: 120,
									EndLine:   120,
									StartPos:  2848,
									EndPos:    2849,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 120,
										EndLine:   120,
										StartPos:  2848,
										EndPos:    2849,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 120,
											EndLine:   120,
											StartPos:  2848,
											EndPos:    2849,
										},
										Value: "a",
									},
								},
							},
							&stmt.Property{
								Position: &position.Position{
									StartLine: 120,
									EndLine:   120,
									StartPos:  2852,
									EndPos:    2857,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 120,
										EndLine:   120,
										StartPos:  2852,
										EndPos:    2853,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 120,
											EndLine:   120,
											StartPos:  2852,
											EndPos:    2853,
										},
										Value: "b",
									},
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 120,
										EndLine:   120,
										StartPos:  2857,
										EndPos:    2857,
									},
									Value: "1",
								},
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 121,
					EndLine:   121,
					StartPos:  2863,
					EndPos:    2880,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 121,
							EndLine:   121,
							StartPos:  2870,
							EndPos:    2871,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 121,
								EndLine:   121,
								StartPos:  2870,
								EndPos:    2871,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 121,
									EndLine:   121,
									StartPos:  2870,
									EndPos:    2871,
								},
								Value: "a",
							},
						},
					},
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 121,
							EndLine:   121,
							StartPos:  2874,
							EndPos:    2879,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 121,
								EndLine:   121,
								StartPos:  2874,
								EndPos:    2875,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 121,
									EndLine:   121,
									StartPos:  2874,
									EndPos:    2875,
								},
								Value: "b",
							},
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 121,
								EndLine:   121,
								StartPos:  2879,
								EndPos:    2879,
							},
							Value: "1",
						},
					},
				},
			},
			&stmt.AltSwitch{
				Position: &position.Position{
					StartLine: 123,
					EndLine:   127,
					StartPos:  2885,
					EndPos:    2943,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 123,
						EndLine:   123,
						StartPos:  2893,
						EndPos:    2893,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 124,
						EndLine:   -1,
						StartPos:  2901,
						EndPos:    -1,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 124,
								EndLine:   -1,
								StartPos:  2901,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 124,
									EndLine:   124,
									StartPos:  2906,
									EndPos:    2906,
								},
								Value: "1",
							},
							Stmts: []node.Node{},
						},
						&stmt.Default{
							Position: &position.Position{
								StartLine: 125,
								EndLine:   -1,
								StartPos:  2912,
								EndPos:    -1,
							},
							Stmts: []node.Node{},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 126,
								EndLine:   -1,
								StartPos:  2924,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 126,
									EndLine:   126,
									StartPos:  2929,
									EndPos:    2929,
								},
								Value: "2",
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.AltSwitch{
				Position: &position.Position{
					StartLine: 129,
					EndLine:   132,
					StartPos:  2948,
					EndPos:    2995,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 129,
						EndLine:   129,
						StartPos:  2956,
						EndPos:    2956,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 130,
						EndLine:   -1,
						StartPos:  2965,
						EndPos:    -1,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 130,
								EndLine:   -1,
								StartPos:  2965,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 130,
									EndLine:   130,
									StartPos:  2970,
									EndPos:    2970,
								},
								Value: "1",
							},
							Stmts: []node.Node{},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 131,
								EndLine:   -1,
								StartPos:  2976,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 131,
									EndLine:   131,
									StartPos:  2981,
									EndPos:    2981,
								},
								Value: "2",
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Switch{
				Position: &position.Position{
					StartLine: 134,
					EndLine:   137,
					StartPos:  3002,
					EndPos:    3053,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 134,
						EndLine:   134,
						StartPos:  3010,
						EndPos:    3010,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 134,
						EndLine:   137,
						StartPos:  3013,
						EndPos:    3053,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 135,
								EndLine:   135,
								StartPos:  3018,
								EndPos:    3031,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 135,
									EndLine:   135,
									StartPos:  3023,
									EndPos:    3023,
								},
								Value: "1",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 135,
										EndLine:   135,
										StartPos:  3026,
										EndPos:    3031,
									},
								},
							},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 136,
								EndLine:   136,
								StartPos:  3036,
								EndPos:    3049,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 136,
									EndLine:   136,
									StartPos:  3041,
									EndPos:    3041,
								},
								Value: "2",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 136,
										EndLine:   136,
										StartPos:  3044,
										EndPos:    3049,
									},
								},
							},
						},
					},
				},
			},
			&stmt.Switch{
				Position: &position.Position{
					StartLine: 139,
					EndLine:   142,
					StartPos:  3060,
					EndPos:    3112,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 139,
						EndLine:   139,
						StartPos:  3068,
						EndPos:    3068,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 139,
						EndLine:   142,
						StartPos:  3071,
						EndPos:    3112,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 140,
								EndLine:   140,
								StartPos:  3077,
								EndPos:    3090,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 140,
									EndLine:   140,
									StartPos:  3082,
									EndPos:    3082,
								},
								Value: "1",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 140,
										EndLine:   140,
										StartPos:  3085,
										EndPos:    3090,
									},
								},
							},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 141,
								EndLine:   141,
								StartPos:  3095,
								EndPos:    3108,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 141,
									EndLine:   141,
									StartPos:  3100,
									EndPos:    3100,
								},
								Value: "2",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 141,
										EndLine:   141,
										StartPos:  3103,
										EndPos:    3108,
									},
								},
							},
						},
					},
				},
			},
			&stmt.Throw{
				Position: &position.Position{
					StartLine: 144,
					EndLine:   144,
					StartPos:  3117,
					EndPos:    3125,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 144,
						EndLine:   144,
						StartPos:  3123,
						EndPos:    3124,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 144,
							EndLine:   144,
							StartPos:  3123,
							EndPos:    3124,
						},
						Value: "e",
					},
				},
			},
			&stmt.Trait{
				Position: &position.Position{
					StartLine: 146,
					EndLine:   146,
					StartPos:  3130,
					EndPos:    3141,
				},
				PhpDocComment: "",
				TraitName: &node.Identifier{
					Position: &position.Position{
						StartLine: 146,
						EndLine:   146,
						StartPos:  3136,
						EndPos:    3138,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 147,
					EndLine:   147,
					StartPos:  3145,
					EndPos:    3166,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 147,
						EndLine:   147,
						StartPos:  3151,
						EndPos:    3153,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 147,
							EndLine:   147,
							StartPos:  3157,
							EndPos:    3164,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 147,
									EndLine:   147,
									StartPos:  3161,
									EndPos:    3163,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 147,
											EndLine:   147,
											StartPos:  3161,
											EndPos:    3163,
										},
										Value: "Bar",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.Nop{
							Position: &position.Position{
								StartLine: 147,
								EndLine:   147,
								StartPos:  3164,
								EndPos:    3164,
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 148,
					EndLine:   148,
					StartPos:  3170,
					EndPos:    3198,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 148,
						EndLine:   148,
						StartPos:  3176,
						EndPos:    3178,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 148,
							EndLine:   148,
							StartPos:  3182,
							EndPos:    3196,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 148,
									EndLine:   148,
									StartPos:  3186,
									EndPos:    3188,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 148,
											EndLine:   148,
											StartPos:  3186,
											EndPos:    3188,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 148,
									EndLine:   148,
									StartPos:  3191,
									EndPos:    3193,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 148,
											EndLine:   148,
											StartPos:  3191,
											EndPos:    3193,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 148,
								EndLine:   148,
								StartPos:  3195,
								EndPos:    3196,
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 149,
					EndLine:   149,
					StartPos:  3202,
					EndPos:    3247,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 149,
						EndLine:   149,
						StartPos:  3208,
						EndPos:    3210,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 149,
							EndLine:   149,
							StartPos:  3214,
							EndPos:    3245,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3218,
									EndPos:    3220,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 149,
											EndLine:   149,
											StartPos:  3218,
											EndPos:    3220,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3223,
									EndPos:    3225,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 149,
											EndLine:   149,
											StartPos:  3223,
											EndPos:    3225,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 149,
								EndLine:   149,
								StartPos:  3227,
								EndPos:    3245,
							},
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3229,
										EndPos:    3242,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 149,
											EndLine:   149,
											StartPos:  3229,
											EndPos:    3231,
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3229,
												EndPos:    3231,
											},
											Value: "one",
										},
									},
									Alias: &node.Identifier{
										Position: &position.Position{
											StartLine: 149,
											EndLine:   149,
											StartPos:  3236,
											EndPos:    3242,
										},
										Value: "include",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 150,
					EndLine:   150,
					StartPos:  3251,
					EndPos:    3295,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 150,
						EndLine:   150,
						StartPos:  3257,
						EndPos:    3259,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 150,
							EndLine:   150,
							StartPos:  3263,
							EndPos:    3293,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 150,
									EndLine:   150,
									StartPos:  3267,
									EndPos:    3269,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 150,
											EndLine:   150,
											StartPos:  3267,
											EndPos:    3269,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 150,
									EndLine:   150,
									StartPos:  3272,
									EndPos:    3274,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 150,
											EndLine:   150,
											StartPos:  3272,
											EndPos:    3274,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 150,
								EndLine:   150,
								StartPos:  3276,
								EndPos:    3293,
							},
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 150,
										EndLine:   150,
										StartPos:  3278,
										EndPos:    3290,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 150,
											EndLine:   150,
											StartPos:  3278,
											EndPos:    3280,
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 150,
												EndLine:   150,
												StartPos:  3278,
												EndPos:    3280,
											},
											Value: "one",
										},
									},
									Modifier: &node.Identifier{
										Position: &position.Position{
											StartLine: 150,
											EndLine:   150,
											StartPos:  3285,
											EndPos:    3290,
										},
										Value: "public",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 151,
					EndLine:   151,
					StartPos:  3299,
					EndPos:    3347,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 151,
						EndLine:   151,
						StartPos:  3305,
						EndPos:    3307,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 151,
							EndLine:   151,
							StartPos:  3311,
							EndPos:    3345,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 151,
									EndLine:   151,
									StartPos:  3315,
									EndPos:    3317,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3315,
											EndPos:    3317,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 151,
									EndLine:   151,
									StartPos:  3320,
									EndPos:    3322,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3320,
											EndPos:    3322,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 151,
								EndLine:   151,
								StartPos:  3324,
								EndPos:    3345,
							},
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 151,
										EndLine:   151,
										StartPos:  3326,
										EndPos:    3342,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3326,
											EndPos:    3328,
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3326,
												EndPos:    3328,
											},
											Value: "one",
										},
									},
									Modifier: &node.Identifier{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3333,
											EndPos:    3338,
										},
										Value: "public",
									},
									Alias: &node.Identifier{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3340,
											EndPos:    3342,
										},
										Value: "two",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 152,
					EndLine:   152,
					StartPos:  3351,
					EndPos:    3427,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 152,
						EndLine:   152,
						StartPos:  3357,
						EndPos:    3359,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 152,
							EndLine:   152,
							StartPos:  3363,
							EndPos:    3425,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 152,
									EndLine:   152,
									StartPos:  3367,
									EndPos:    3369,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 152,
											EndLine:   152,
											StartPos:  3367,
											EndPos:    3369,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 152,
									EndLine:   152,
									StartPos:  3372,
									EndPos:    3374,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 152,
											EndLine:   152,
											StartPos:  3372,
											EndPos:    3374,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 152,
								EndLine:   152,
								StartPos:  3376,
								EndPos:    3425,
							},
							Adaptations: []node.Node{
								&stmt.TraitUsePrecedence{
									Position: &position.Position{
										StartLine: 152,
										EndLine:   152,
										StartPos:  3378,
										EndPos:    3405,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 152,
											EndLine:   152,
											StartPos:  3378,
											EndPos:    3385,
										},
										Trait: &name.Name{
											Position: &position.Position{
												StartLine: 152,
												EndLine:   152,
												StartPos:  3378,
												EndPos:    3380,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 152,
														EndLine:   152,
														StartPos:  3378,
														EndPos:    3380,
													},
													Value: "Bar",
												},
											},
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 152,
												EndLine:   152,
												StartPos:  3383,
												EndPos:    3385,
											},
											Value: "one",
										},
									},
									Insteadof: []node.Node{
										&name.Name{
											Position: &position.Position{
												StartLine: 152,
												EndLine:   152,
												StartPos:  3397,
												EndPos:    3399,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 152,
														EndLine:   152,
														StartPos:  3397,
														EndPos:    3399,
													},
													Value: "Baz",
												},
											},
										},
										&name.Name{
											Position: &position.Position{
												StartLine: 152,
												EndLine:   152,
												StartPos:  3402,
												EndPos:    3405,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 152,
														EndLine:   152,
														StartPos:  3402,
														EndPos:    3405,
													},
													Value: "Quux",
												},
											},
										},
									},
								},
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 152,
										EndLine:   152,
										StartPos:  3408,
										EndPos:    3422,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 152,
											EndLine:   152,
											StartPos:  3408,
											EndPos:    3415,
										},
										Trait: &name.Name{
											Position: &position.Position{
												StartLine: 152,
												EndLine:   152,
												StartPos:  3408,
												EndPos:    3410,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 152,
														EndLine:   152,
														StartPos:  3408,
														EndPos:    3410,
													},
													Value: "Baz",
												},
											},
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 152,
												EndLine:   152,
												StartPos:  3413,
												EndPos:    3415,
											},
											Value: "one",
										},
									},
									Alias: &node.Identifier{
										Position: &position.Position{
											StartLine: 152,
											EndLine:   152,
											StartPos:  3420,
											EndPos:    3422,
										},
										Value: "two",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Try{
				Position: &position.Position{
					StartLine: 154,
					EndLine:   -1,
					StartPos:  3432,
					EndPos:    -1,
				},
				Stmts:   []node.Node{},
				Catches: []node.Node{},
			},
			&stmt.Try{
				Position: &position.Position{
					StartLine: 155,
					EndLine:   155,
					StartPos:  3441,
					EndPos:    3470,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 155,
							EndLine:   155,
							StartPos:  3448,
							EndPos:    3470,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 155,
									EndLine:   155,
									StartPos:  3455,
									EndPos:    3463,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 155,
											EndLine:   155,
											StartPos:  3455,
											EndPos:    3463,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 155,
								EndLine:   155,
								StartPos:  3465,
								EndPos:    3466,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 155,
									EndLine:   155,
									StartPos:  3465,
									EndPos:    3466,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.Try{
				Position: &position.Position{
					StartLine: 156,
					EndLine:   156,
					StartPos:  3474,
					EndPos:    3520,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 156,
							EndLine:   156,
							StartPos:  3481,
							EndPos:    3520,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 156,
									EndLine:   156,
									StartPos:  3488,
									EndPos:    3496,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 156,
											EndLine:   156,
											StartPos:  3488,
											EndPos:    3496,
										},
										Value: "Exception",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 156,
									EndLine:   156,
									StartPos:  3498,
									EndPos:    3513,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 156,
											EndLine:   156,
											StartPos:  3498,
											EndPos:    3513,
										},
										Value: "RuntimeException",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 156,
								EndLine:   156,
								StartPos:  3515,
								EndPos:    3516,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 156,
									EndLine:   156,
									StartPos:  3515,
									EndPos:    3516,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.Try{
				Position: &position.Position{
					StartLine: 157,
					EndLine:   157,
					StartPos:  3524,
					EndPos:    3584,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 157,
							EndLine:   157,
							StartPos:  3531,
							EndPos:    3553,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 157,
									EndLine:   157,
									StartPos:  3538,
									EndPos:    3546,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 157,
											EndLine:   157,
											StartPos:  3538,
											EndPos:    3546,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 157,
								EndLine:   157,
								StartPos:  3548,
								EndPos:    3549,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 157,
									EndLine:   157,
									StartPos:  3548,
									EndPos:    3549,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 157,
							EndLine:   157,
							StartPos:  3555,
							EndPos:    3584,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 157,
									EndLine:   157,
									StartPos:  3562,
									EndPos:    3577,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 157,
											EndLine:   157,
											StartPos:  3562,
											EndPos:    3577,
										},
										Value: "RuntimeException",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 157,
								EndLine:   157,
								StartPos:  3579,
								EndPos:    3580,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 157,
									EndLine:   157,
									StartPos:  3579,
									EndPos:    3580,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.Try{
				Position: &position.Position{
					StartLine: 158,
					EndLine:   158,
					StartPos:  3588,
					EndPos:    3628,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 158,
							EndLine:   158,
							StartPos:  3595,
							EndPos:    3617,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 158,
									EndLine:   158,
									StartPos:  3602,
									EndPos:    3610,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 158,
											EndLine:   158,
											StartPos:  3602,
											EndPos:    3610,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 158,
								EndLine:   158,
								StartPos:  3612,
								EndPos:    3613,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 158,
									EndLine:   158,
									StartPos:  3612,
									EndPos:    3613,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
				},
				Finally: &stmt.Finally{
					Position: &position.Position{
						StartLine: 158,
						EndLine:   158,
						StartPos:  3619,
						EndPos:    3628,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Unset{
				Position: &position.Position{
					StartLine: 160,
					EndLine:   160,
					StartPos:  3633,
					EndPos:    3647,
				},
				Vars: []node.Node{
					&expr.Variable{
						Position: &position.Position{
							StartLine: 160,
							EndLine:   160,
							StartPos:  3639,
							EndPos:    3640,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 160,
								EndLine:   160,
								StartPos:  3639,
								EndPos:    3640,
							},
							Value: "a",
						},
					},
					&expr.Variable{
						Position: &position.Position{
							StartLine: 160,
							EndLine:   160,
							StartPos:  3643,
							EndPos:    3644,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 160,
								EndLine:   160,
								StartPos:  3643,
								EndPos:    3644,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 162,
					EndLine:   162,
					StartPos:  3652,
					EndPos:    3659,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 162,
							EndLine:   162,
							StartPos:  3656,
							EndPos:    3658,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3656,
								EndPos:    3658,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 162,
										EndLine:   162,
										StartPos:  3656,
										EndPos:    3658,
									},
									Value: "Foo",
								},
							},
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 163,
					EndLine:   163,
					StartPos:  3663,
					EndPos:    3671,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 163,
							EndLine:   163,
							StartPos:  3668,
							EndPos:    3670,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 163,
								EndLine:   163,
								StartPos:  3668,
								EndPos:    3670,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 163,
										EndLine:   163,
										StartPos:  3668,
										EndPos:    3670,
									},
									Value: "Foo",
								},
							},
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 164,
					EndLine:   164,
					StartPos:  3675,
					EndPos:    3690,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 164,
							EndLine:   164,
							StartPos:  3680,
							EndPos:    3689,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 164,
								EndLine:   164,
								StartPos:  3680,
								EndPos:    3682,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 164,
										EndLine:   164,
										StartPos:  3680,
										EndPos:    3682,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 164,
								EndLine:   164,
								StartPos:  3687,
								EndPos:    3689,
							},
							Value: "Bar",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 165,
					EndLine:   165,
					StartPos:  3694,
					EndPos:    3706,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 165,
							EndLine:   165,
							StartPos:  3698,
							EndPos:    3700,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 165,
								EndLine:   165,
								StartPos:  3698,
								EndPos:    3700,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 165,
										EndLine:   165,
										StartPos:  3698,
										EndPos:    3700,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 165,
							EndLine:   165,
							StartPos:  3703,
							EndPos:    3705,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 165,
								EndLine:   165,
								StartPos:  3703,
								EndPos:    3705,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 165,
										EndLine:   165,
										StartPos:  3703,
										EndPos:    3705,
									},
									Value: "Bar",
								},
							},
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 166,
					EndLine:   166,
					StartPos:  3710,
					EndPos:    3729,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 166,
							EndLine:   166,
							StartPos:  3714,
							EndPos:    3716,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 166,
								EndLine:   166,
								StartPos:  3714,
								EndPos:    3716,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 166,
										EndLine:   166,
										StartPos:  3714,
										EndPos:    3716,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 166,
							EndLine:   166,
							StartPos:  3719,
							EndPos:    3728,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 166,
								EndLine:   166,
								StartPos:  3719,
								EndPos:    3721,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 166,
										EndLine:   166,
										StartPos:  3719,
										EndPos:    3721,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 166,
								EndLine:   166,
								StartPos:  3726,
								EndPos:    3728,
							},
							Value: "Baz",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 167,
					EndLine:   167,
					StartPos:  3733,
					EndPos:    3755,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 167,
						EndLine:   167,
						StartPos:  3737,
						EndPos:    3744,
					},
					Value: "function",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 167,
							EndLine:   167,
							StartPos:  3746,
							EndPos:    3748,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 167,
								EndLine:   167,
								StartPos:  3746,
								EndPos:    3748,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 167,
										EndLine:   167,
										StartPos:  3746,
										EndPos:    3748,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 167,
							EndLine:   167,
							StartPos:  3752,
							EndPos:    3754,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 167,
								EndLine:   167,
								StartPos:  3752,
								EndPos:    3754,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 167,
										EndLine:   167,
										StartPos:  3752,
										EndPos:    3754,
									},
									Value: "Bar",
								},
							},
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 168,
					EndLine:   168,
					StartPos:  3759,
					EndPos:    3795,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 168,
						EndLine:   168,
						StartPos:  3763,
						EndPos:    3770,
					},
					Value: "function",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 168,
							EndLine:   168,
							StartPos:  3772,
							EndPos:    3781,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 168,
								EndLine:   168,
								StartPos:  3772,
								EndPos:    3774,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 168,
										EndLine:   168,
										StartPos:  3772,
										EndPos:    3774,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 168,
								EndLine:   168,
								StartPos:  3779,
								EndPos:    3781,
							},
							Value: "foo",
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 168,
							EndLine:   168,
							StartPos:  3785,
							EndPos:    3794,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 168,
								EndLine:   168,
								StartPos:  3785,
								EndPos:    3787,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 168,
										EndLine:   168,
										StartPos:  3785,
										EndPos:    3787,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 168,
								EndLine:   168,
								StartPos:  3792,
								EndPos:    3794,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 169,
					EndLine:   169,
					StartPos:  3799,
					EndPos:    3818,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 169,
						EndLine:   169,
						StartPos:  3803,
						EndPos:    3807,
					},
					Value: "const",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 169,
							EndLine:   169,
							StartPos:  3809,
							EndPos:    3811,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 169,
								EndLine:   169,
								StartPos:  3809,
								EndPos:    3811,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 169,
										EndLine:   169,
										StartPos:  3809,
										EndPos:    3811,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 169,
							EndLine:   169,
							StartPos:  3815,
							EndPos:    3817,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 169,
								EndLine:   169,
								StartPos:  3815,
								EndPos:    3817,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 169,
										EndLine:   169,
										StartPos:  3815,
										EndPos:    3817,
									},
									Value: "Bar",
								},
							},
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 170,
					EndLine:   170,
					StartPos:  3822,
					EndPos:    3855,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 170,
						EndLine:   170,
						StartPos:  3826,
						EndPos:    3830,
					},
					Value: "const",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 170,
							EndLine:   170,
							StartPos:  3832,
							EndPos:    3841,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 170,
								EndLine:   170,
								StartPos:  3832,
								EndPos:    3834,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3832,
										EndPos:    3834,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 170,
								EndLine:   170,
								StartPos:  3839,
								EndPos:    3841,
							},
							Value: "foo",
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 170,
							EndLine:   170,
							StartPos:  3845,
							EndPos:    3854,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 170,
								EndLine:   170,
								StartPos:  3845,
								EndPos:    3847,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3845,
										EndPos:    3847,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 170,
								EndLine:   170,
								StartPos:  3852,
								EndPos:    3854,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.GroupUse{
				Position: &position.Position{
					StartLine: 172,
					EndLine:   172,
					StartPos:  3860,
					EndPos:    3879,
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 172,
						EndLine:   172,
						StartPos:  3865,
						EndPos:    3867,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3865,
								EndPos:    3867,
							},
							Value: "Foo",
						},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 172,
							EndLine:   172,
							StartPos:  3870,
							EndPos:    3872,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3870,
								EndPos:    3872,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 172,
										EndLine:   172,
										StartPos:  3870,
										EndPos:    3872,
									},
									Value: "Bar",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 172,
							EndLine:   172,
							StartPos:  3875,
							EndPos:    3877,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3875,
								EndPos:    3877,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 172,
										EndLine:   172,
										StartPos:  3875,
										EndPos:    3877,
									},
									Value: "Baz",
								},
							},
						},
					},
				},
			},
			&stmt.GroupUse{
				Position: &position.Position{
					StartLine: 173,
					EndLine:   173,
					StartPos:  3883,
					EndPos:    3909,
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 173,
						EndLine:   173,
						StartPos:  3887,
						EndPos:    3889,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 173,
								EndLine:   173,
								StartPos:  3887,
								EndPos:    3889,
							},
							Value: "Foo",
						},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 173,
							EndLine:   173,
							StartPos:  3892,
							EndPos:    3894,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 173,
								EndLine:   173,
								StartPos:  3892,
								EndPos:    3894,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 173,
										EndLine:   173,
										StartPos:  3892,
										EndPos:    3894,
									},
									Value: "Bar",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 173,
							EndLine:   173,
							StartPos:  3897,
							EndPos:    3907,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 173,
								EndLine:   173,
								StartPos:  3897,
								EndPos:    3899,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 173,
										EndLine:   173,
										StartPos:  3897,
										EndPos:    3899,
									},
									Value: "Baz",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 173,
								EndLine:   173,
								StartPos:  3904,
								EndPos:    3907,
							},
							Value: "quux",
						},
					},
				},
			},
			&stmt.GroupUse{
				Position: &position.Position{
					StartLine: 174,
					EndLine:   174,
					StartPos:  3913,
					EndPos:    3940,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 174,
						EndLine:   174,
						StartPos:  3917,
						EndPos:    3924,
					},
					Value: "function",
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 174,
						EndLine:   174,
						StartPos:  3926,
						EndPos:    3928,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3926,
								EndPos:    3928,
							},
							Value: "Foo",
						},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 174,
							EndLine:   174,
							StartPos:  3931,
							EndPos:    3933,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3931,
								EndPos:    3933,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 174,
										EndLine:   174,
										StartPos:  3931,
										EndPos:    3933,
									},
									Value: "Bar",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 174,
							EndLine:   174,
							StartPos:  3936,
							EndPos:    3938,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3936,
								EndPos:    3938,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 174,
										EndLine:   174,
										StartPos:  3936,
										EndPos:    3938,
									},
									Value: "Baz",
								},
							},
						},
					},
				},
			},
			&stmt.GroupUse{
				Position: &position.Position{
					StartLine: 175,
					EndLine:   175,
					StartPos:  3944,
					EndPos:    3969,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 175,
						EndLine:   175,
						StartPos:  3948,
						EndPos:    3952,
					},
					Value: "const",
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 175,
						EndLine:   175,
						StartPos:  3955,
						EndPos:    3957,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3955,
								EndPos:    3957,
							},
							Value: "Foo",
						},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 175,
							EndLine:   175,
							StartPos:  3960,
							EndPos:    3962,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3960,
								EndPos:    3962,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 175,
										EndLine:   175,
										StartPos:  3960,
										EndPos:    3962,
									},
									Value: "Bar",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 175,
							EndLine:   175,
							StartPos:  3965,
							EndPos:    3967,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3965,
								EndPos:    3967,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 175,
										EndLine:   175,
										StartPos:  3965,
										EndPos:    3967,
									},
									Value: "Baz",
								},
							},
						},
					},
				},
			},
			&stmt.GroupUse{
				Position: &position.Position{
					StartLine: 176,
					EndLine:   176,
					StartPos:  3973,
					EndPos:    4006,
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 176,
						EndLine:   176,
						StartPos:  3977,
						EndPos:    3979,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 176,
								EndLine:   176,
								StartPos:  3977,
								EndPos:    3979,
							},
							Value: "Foo",
						},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 176,
							EndLine:   176,
							StartPos:  3988,
							EndPos:    3990,
						},
						UseType: &node.Identifier{
							Position: &position.Position{
								StartLine: 176,
								EndLine:   176,
								StartPos:  3982,
								EndPos:    3986,
							},
							Value: "const",
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 176,
								EndLine:   176,
								StartPos:  3988,
								EndPos:    3990,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 176,
										EndLine:   176,
										StartPos:  3988,
										EndPos:    3990,
									},
									Value: "Bar",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 176,
							EndLine:   176,
							StartPos:  4002,
							EndPos:    4004,
						},
						UseType: &node.Identifier{
							Position: &position.Position{
								StartLine: 176,
								EndLine:   176,
								StartPos:  3993,
								EndPos:    4000,
							},
							Value: "function",
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 176,
								EndLine:   176,
								StartPos:  4002,
								EndPos:    4004,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 176,
										EndLine:   176,
										StartPos:  4002,
										EndPos:    4004,
									},
									Value: "Baz",
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 178,
					EndLine:   178,
					StartPos:  4011,
					EndPos:    4016,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 178,
						EndLine:   178,
						StartPos:  4011,
						EndPos:    4015,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 178,
							EndLine:   178,
							StartPos:  4011,
							EndPos:    4012,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 178,
								EndLine:   178,
								StartPos:  4011,
								EndPos:    4012,
							},
							Value: "a",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 178,
							EndLine:   178,
							StartPos:  4014,
							EndPos:    4014,
						},
						Value: "1",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 179,
					EndLine:   179,
					StartPos:  4020,
					EndPos:    4028,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 179,
						EndLine:   179,
						StartPos:  4020,
						EndPos:    4027,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 179,
							EndLine:   179,
							StartPos:  4020,
							EndPos:    4024,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 179,
								EndLine:   179,
								StartPos:  4020,
								EndPos:    4021,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 179,
									EndLine:   179,
									StartPos:  4020,
									EndPos:    4021,
								},
								Value: "a",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 179,
								EndLine:   179,
								StartPos:  4023,
								EndPos:    4023,
							},
							Value: "1",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 179,
							EndLine:   179,
							StartPos:  4026,
							EndPos:    4026,
						},
						Value: "2",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 180,
					EndLine:   180,
					StartPos:  4032,
					EndPos:    4039,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 180,
						EndLine:   180,
						StartPos:  4032,
						EndPos:    4038,
					},
					Items: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 181,
					EndLine:   181,
					StartPos:  4043,
					EndPos:    4051,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 181,
						EndLine:   181,
						StartPos:  4043,
						EndPos:    4050,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 181,
								EndLine:   181,
								StartPos:  4049,
								EndPos:    4049,
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 181,
									EndLine:   181,
									StartPos:  4049,
									EndPos:    4049,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 182,
					EndLine:   182,
					StartPos:  4055,
					EndPos:    4072,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 182,
						EndLine:   182,
						StartPos:  4055,
						EndPos:    4071,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 182,
								EndLine:   182,
								StartPos:  4061,
								EndPos:    4064,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 182,
									EndLine:   182,
									StartPos:  4061,
									EndPos:    4061,
								},
								Value: "1",
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 182,
									EndLine:   182,
									StartPos:  4064,
									EndPos:    4064,
								},
								Value: "1",
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 182,
								EndLine:   182,
								StartPos:  4067,
								EndPos:    4069,
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 182,
									EndLine:   182,
									StartPos:  4067,
									EndPos:    4069,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 182,
										EndLine:   182,
										StartPos:  4068,
										EndPos:    4069,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 182,
											EndLine:   182,
											StartPos:  4068,
											EndPos:    4069,
										},
										Value: "b",
									},
								},
							},
						},
						&expr.ArrayItem{},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 183,
					EndLine:   183,
					StartPos:  4076,
					EndPos:    4079,
				},
				Expr: &expr.BitwiseNot{
					Position: &position.Position{
						StartLine: 183,
						EndLine:   183,
						StartPos:  4076,
						EndPos:    4078,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 183,
							EndLine:   183,
							StartPos:  4077,
							EndPos:    4078,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 183,
								EndLine:   183,
								StartPos:  4077,
								EndPos:    4078,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 184,
					EndLine:   184,
					StartPos:  4083,
					EndPos:    4086,
				},
				Expr: &expr.BooleanNot{
					Position: &position.Position{
						StartLine: 184,
						EndLine:   184,
						StartPos:  4083,
						EndPos:    4085,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 184,
							EndLine:   184,
							StartPos:  4084,
							EndPos:    4085,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 184,
								EndLine:   184,
								StartPos:  4084,
								EndPos:    4085,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 186,
					EndLine:   186,
					StartPos:  4091,
					EndPos:    4099,
				},
				Expr: &expr.ClassConstFetch{
					Position: &position.Position{
						StartLine: 186,
						EndLine:   186,
						StartPos:  4091,
						EndPos:    4098,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 186,
							EndLine:   186,
							StartPos:  4091,
							EndPos:    4093,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 186,
									EndLine:   186,
									StartPos:  4091,
									EndPos:    4093,
								},
								Value: "Foo",
							},
						},
					},
					ConstantName: &node.Identifier{
						Position: &position.Position{
							StartLine: 186,
							EndLine:   186,
							StartPos:  4096,
							EndPos:    4098,
						},
						Value: "Bar",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 187,
					EndLine:   187,
					StartPos:  4103,
					EndPos:    4112,
				},
				Expr: &expr.ClassConstFetch{
					Position: &position.Position{
						StartLine: 187,
						EndLine:   187,
						StartPos:  4103,
						EndPos:    4111,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 187,
							EndLine:   187,
							StartPos:  4103,
							EndPos:    4106,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 187,
								EndLine:   187,
								StartPos:  4103,
								EndPos:    4106,
							},
							Value: "foo",
						},
					},
					ConstantName: &node.Identifier{
						Position: &position.Position{
							StartLine: 187,
							EndLine:   187,
							StartPos:  4109,
							EndPos:    4111,
						},
						Value: "Bar",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 188,
					EndLine:   188,
					StartPos:  4116,
					EndPos:    4125,
				},
				Expr: &expr.Clone{
					Position: &position.Position{
						StartLine: 188,
						EndLine:   188,
						StartPos:  4116,
						EndPos:    4123,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 188,
							EndLine:   188,
							StartPos:  4122,
							EndPos:    4123,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 188,
								EndLine:   188,
								StartPos:  4122,
								EndPos:    4123,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 189,
					EndLine:   189,
					StartPos:  4129,
					EndPos:    4137,
				},
				Expr: &expr.Clone{
					Position: &position.Position{
						StartLine: 189,
						EndLine:   189,
						StartPos:  4129,
						EndPos:    4136,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 189,
							EndLine:   189,
							StartPos:  4135,
							EndPos:    4136,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 189,
								EndLine:   189,
								StartPos:  4135,
								EndPos:    4136,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 190,
					EndLine:   190,
					StartPos:  4141,
					EndPos:    4153,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 190,
						EndLine:   190,
						StartPos:  4141,
						EndPos:    4152,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Stmts:         []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 191,
					EndLine:   191,
					StartPos:  4157,
					EndPos:    4190,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 191,
						EndLine:   191,
						StartPos:  4157,
						EndPos:    4189,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 191,
								EndLine:   191,
								StartPos:  4166,
								EndPos:    4167,
							},
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 191,
									EndLine:   191,
									StartPos:  4166,
									EndPos:    4167,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 191,
										EndLine:   191,
										StartPos:  4166,
										EndPos:    4167,
									},
									Value: "a",
								},
							},
						},
						&node.Parameter{
							Position: &position.Position{
								StartLine: 191,
								EndLine:   191,
								StartPos:  4170,
								EndPos:    4171,
							},
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 191,
									EndLine:   191,
									StartPos:  4170,
									EndPos:    4171,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 191,
										EndLine:   191,
										StartPos:  4170,
										EndPos:    4171,
									},
									Value: "b",
								},
							},
						},
					},
					ClosureUse: &expr.ClosureUse{
						Position: &position.Position{
							StartLine: 191,
							EndLine:   191,
							StartPos:  4174,
							EndPos:    4186,
						},
						Uses: []node.Node{
							&expr.Variable{
								Position: &position.Position{
									StartLine: 191,
									EndLine:   191,
									StartPos:  4179,
									EndPos:    4180,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 191,
										EndLine:   191,
										StartPos:  4179,
										EndPos:    4180,
									},
									Value: "c",
								},
							},
							&expr.Reference{
								Position: &position.Position{
									StartLine: 191,
									EndLine:   191,
									StartPos:  4183,
									EndPos:    4185,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 191,
										EndLine:   191,
										StartPos:  4184,
										EndPos:    4185,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 191,
											EndLine:   191,
											StartPos:  4184,
											EndPos:    4185,
										},
										Value: "d",
									},
								},
							},
						},
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 192,
					EndLine:   192,
					StartPos:  4194,
					EndPos:    4213,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 192,
						EndLine:   192,
						StartPos:  4194,
						EndPos:    4212,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					ReturnType: &name.Name{
						Position: &position.Position{
							StartLine: 192,
							EndLine:   192,
							StartPos:  4206,
							EndPos:    4209,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 192,
									EndLine:   192,
									StartPos:  4206,
									EndPos:    4209,
								},
								Value: "void",
							},
						},
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 193,
					EndLine:   193,
					StartPos:  4217,
					EndPos:    4220,
				},
				Expr: &expr.ConstFetch{
					Position: &position.Position{
						StartLine: 193,
						EndLine:   193,
						StartPos:  4217,
						EndPos:    4219,
					},
					Constant: &name.Name{
						Position: &position.Position{
							StartLine: 193,
							EndLine:   193,
							StartPos:  4217,
							EndPos:    4219,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 193,
									EndLine:   193,
									StartPos:  4217,
									EndPos:    4219,
								},
								Value: "foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 194,
					EndLine:   194,
					StartPos:  4224,
					EndPos:    4237,
				},
				Expr: &expr.ConstFetch{
					Position: &position.Position{
						StartLine: 194,
						EndLine:   194,
						StartPos:  4224,
						EndPos:    4236,
					},
					Constant: &name.Relative{
						Position: &position.Position{
							StartLine: 194,
							EndLine:   194,
							StartPos:  4224,
							EndPos:    4236,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 194,
									EndLine:   194,
									StartPos:  4234,
									EndPos:    4236,
								},
								Value: "foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 195,
					EndLine:   195,
					StartPos:  4241,
					EndPos:    4245,
				},
				Expr: &expr.ConstFetch{
					Position: &position.Position{
						StartLine: 195,
						EndLine:   195,
						StartPos:  4241,
						EndPos:    4244,
					},
					Constant: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 195,
							EndLine:   195,
							StartPos:  4241,
							EndPos:    4244,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 195,
									EndLine:   195,
									StartPos:  4242,
									EndPos:    4244,
								},
								Value: "foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 197,
					EndLine:   197,
					StartPos:  4250,
					EndPos:    4259,
				},
				Expr: &expr.Empty{
					Position: &position.Position{
						StartLine: 197,
						EndLine:   197,
						StartPos:  4250,
						EndPos:    4258,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 197,
							EndLine:   197,
							StartPos:  4256,
							EndPos:    4257,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 197,
								EndLine:   197,
								StartPos:  4256,
								EndPos:    4257,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 198,
					EndLine:   198,
					StartPos:  4263,
					EndPos:    4266,
				},
				Expr: &expr.ErrorSuppress{
					Position: &position.Position{
						StartLine: 198,
						EndLine:   198,
						StartPos:  4263,
						EndPos:    4265,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 198,
							EndLine:   198,
							StartPos:  4264,
							EndPos:    4265,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 198,
								EndLine:   198,
								StartPos:  4264,
								EndPos:    4265,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 199,
					EndLine:   199,
					StartPos:  4270,
					EndPos:    4278,
				},
				Expr: &expr.Eval{
					Position: &position.Position{
						StartLine: 199,
						EndLine:   199,
						StartPos:  4270,
						EndPos:    4277,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 199,
							EndLine:   199,
							StartPos:  4275,
							EndPos:    4276,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 199,
								EndLine:   199,
								StartPos:  4275,
								EndPos:    4276,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 200,
					EndLine:   200,
					StartPos:  4282,
					EndPos:    4286,
				},
				Expr: &expr.Exit{
					Die: false,
					Position: &position.Position{
						StartLine: 200,
						EndLine:   200,
						StartPos:  4282,
						EndPos:    4285,
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 201,
					EndLine:   201,
					StartPos:  4290,
					EndPos:    4298,
				},
				Expr: &expr.Exit{
					Die: false,
					Position: &position.Position{
						StartLine: 201,
						EndLine:   201,
						StartPos:  4290,
						EndPos:    4297,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 201,
							EndLine:   201,
							StartPos:  4295,
							EndPos:    4296,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 201,
								EndLine:   201,
								StartPos:  4295,
								EndPos:    4296,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 202,
					EndLine:   202,
					StartPos:  4302,
					EndPos:    4305,
				},
				Expr: &expr.Exit{
					Die: true,
					Position: &position.Position{
						StartLine: 202,
						EndLine:   202,
						StartPos:  4302,
						EndPos:    4304,
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 203,
					EndLine:   203,
					StartPos:  4309,
					EndPos:    4316,
				},
				Expr: &expr.Exit{
					Die: true,
					Position: &position.Position{
						StartLine: 203,
						EndLine:   203,
						StartPos:  4309,
						EndPos:    4315,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 203,
							EndLine:   203,
							StartPos:  4313,
							EndPos:    4314,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 203,
								EndLine:   203,
								StartPos:  4313,
								EndPos:    4314,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 204,
					EndLine:   204,
					StartPos:  4320,
					EndPos:    4325,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 204,
						EndLine:   204,
						StartPos:  4320,
						EndPos:    4324,
					},
					Function: &name.Name{
						Position: &position.Position{
							StartLine: 204,
							EndLine:   204,
							StartPos:  4320,
							EndPos:    4322,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 204,
									EndLine:   204,
									StartPos:  4320,
									EndPos:    4322,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 204,
							EndLine:   204,
							StartPos:  4323,
							EndPos:    4324,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 205,
					EndLine:   205,
					StartPos:  4329,
					EndPos:    4344,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 205,
						EndLine:   205,
						StartPos:  4329,
						EndPos:    4343,
					},
					Function: &name.Relative{
						Position: &position.Position{
							StartLine: 205,
							EndLine:   205,
							StartPos:  4329,
							EndPos:    4341,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 205,
									EndLine:   205,
									StartPos:  4339,
									EndPos:    4341,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 205,
							EndLine:   205,
							StartPos:  4342,
							EndPos:    4343,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 206,
					EndLine:   206,
					StartPos:  4348,
					EndPos:    4354,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 206,
						EndLine:   206,
						StartPos:  4348,
						EndPos:    4353,
					},
					Function: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 206,
							EndLine:   206,
							StartPos:  4348,
							EndPos:    4351,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 206,
									EndLine:   206,
									StartPos:  4349,
									EndPos:    4351,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 206,
							EndLine:   206,
							StartPos:  4352,
							EndPos:    4353,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 207,
					EndLine:   207,
					StartPos:  4358,
					EndPos:    4364,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 207,
						EndLine:   207,
						StartPos:  4358,
						EndPos:    4363,
					},
					Function: &expr.Variable{
						Position: &position.Position{
							StartLine: 207,
							EndLine:   207,
							StartPos:  4358,
							EndPos:    4361,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 207,
								EndLine:   207,
								StartPos:  4358,
								EndPos:    4361,
							},
							Value: "foo",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 207,
							EndLine:   207,
							StartPos:  4362,
							EndPos:    4363,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 209,
					EndLine:   209,
					StartPos:  4369,
					EndPos:    4373,
				},
				Expr: &expr.PostDec{
					Position: &position.Position{
						StartLine: 209,
						EndLine:   209,
						StartPos:  4369,
						EndPos:    4372,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 209,
							EndLine:   209,
							StartPos:  4369,
							EndPos:    4370,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 209,
								EndLine:   209,
								StartPos:  4369,
								EndPos:    4370,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 210,
					EndLine:   210,
					StartPos:  4377,
					EndPos:    4381,
				},
				Expr: &expr.PostInc{
					Position: &position.Position{
						StartLine: 210,
						EndLine:   210,
						StartPos:  4377,
						EndPos:    4380,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 210,
							EndLine:   210,
							StartPos:  4377,
							EndPos:    4378,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 210,
								EndLine:   210,
								StartPos:  4377,
								EndPos:    4378,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 211,
					EndLine:   211,
					StartPos:  4385,
					EndPos:    4389,
				},
				Expr: &expr.PreDec{
					Position: &position.Position{
						StartLine: 211,
						EndLine:   211,
						StartPos:  4385,
						EndPos:    4388,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 211,
							EndLine:   211,
							StartPos:  4387,
							EndPos:    4388,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 211,
								EndLine:   211,
								StartPos:  4387,
								EndPos:    4388,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 212,
					EndLine:   212,
					StartPos:  4393,
					EndPos:    4397,
				},
				Expr: &expr.PreInc{
					Position: &position.Position{
						StartLine: 212,
						EndLine:   212,
						StartPos:  4393,
						EndPos:    4396,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 212,
							EndLine:   212,
							StartPos:  4395,
							EndPos:    4396,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 212,
								EndLine:   212,
								StartPos:  4395,
								EndPos:    4396,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 214,
					EndLine:   214,
					StartPos:  4402,
					EndPos:    4412,
				},
				Expr: &expr.Include{
					Position: &position.Position{
						StartLine: 214,
						EndLine:   214,
						StartPos:  4402,
						EndPos:    4411,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 214,
							EndLine:   214,
							StartPos:  4410,
							EndPos:    4411,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 214,
								EndLine:   214,
								StartPos:  4410,
								EndPos:    4411,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 215,
					EndLine:   215,
					StartPos:  4416,
					EndPos:    4431,
				},
				Expr: &expr.IncludeOnce{
					Position: &position.Position{
						StartLine: 215,
						EndLine:   215,
						StartPos:  4416,
						EndPos:    4430,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 215,
							EndLine:   215,
							StartPos:  4429,
							EndPos:    4430,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 215,
								EndLine:   215,
								StartPos:  4429,
								EndPos:    4430,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 216,
					EndLine:   216,
					StartPos:  4435,
					EndPos:    4445,
				},
				Expr: &expr.Require{
					Position: &position.Position{
						StartLine: 216,
						EndLine:   216,
						StartPos:  4435,
						EndPos:    4444,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 216,
							EndLine:   216,
							StartPos:  4443,
							EndPos:    4444,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 216,
								EndLine:   216,
								StartPos:  4443,
								EndPos:    4444,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 217,
					EndLine:   217,
					StartPos:  4449,
					EndPos:    4464,
				},
				Expr: &expr.RequireOnce{
					Position: &position.Position{
						StartLine: 217,
						EndLine:   217,
						StartPos:  4449,
						EndPos:    4463,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 217,
							EndLine:   217,
							StartPos:  4462,
							EndPos:    4463,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 217,
								EndLine:   217,
								StartPos:  4462,
								EndPos:    4463,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 219,
					EndLine:   219,
					StartPos:  4469,
					EndPos:    4486,
				},
				Expr: &expr.InstanceOf{
					Position: &position.Position{
						StartLine: 219,
						EndLine:   219,
						StartPos:  4469,
						EndPos:    4485,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 219,
							EndLine:   219,
							StartPos:  4469,
							EndPos:    4470,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 219,
								EndLine:   219,
								StartPos:  4469,
								EndPos:    4470,
							},
							Value: "a",
						},
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 219,
							EndLine:   219,
							StartPos:  4483,
							EndPos:    4485,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 219,
									EndLine:   219,
									StartPos:  4483,
									EndPos:    4485,
								},
								Value: "Foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 220,
					EndLine:   220,
					StartPos:  4490,
					EndPos:    4517,
				},
				Expr: &expr.InstanceOf{
					Position: &position.Position{
						StartLine: 220,
						EndLine:   220,
						StartPos:  4490,
						EndPos:    4516,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 220,
							EndLine:   220,
							StartPos:  4490,
							EndPos:    4491,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 220,
								EndLine:   220,
								StartPos:  4490,
								EndPos:    4491,
							},
							Value: "a",
						},
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 220,
							EndLine:   220,
							StartPos:  4504,
							EndPos:    4516,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 220,
									EndLine:   220,
									StartPos:  4514,
									EndPos:    4516,
								},
								Value: "Foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 221,
					EndLine:   221,
					StartPos:  4521,
					EndPos:    4539,
				},
				Expr: &expr.InstanceOf{
					Position: &position.Position{
						StartLine: 221,
						EndLine:   221,
						StartPos:  4521,
						EndPos:    4538,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 221,
							EndLine:   221,
							StartPos:  4521,
							EndPos:    4522,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 221,
								EndLine:   221,
								StartPos:  4521,
								EndPos:    4522,
							},
							Value: "a",
						},
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 221,
							EndLine:   221,
							StartPos:  4535,
							EndPos:    4538,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 221,
									EndLine:   221,
									StartPos:  4536,
									EndPos:    4538,
								},
								Value: "Foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 223,
					EndLine:   223,
					StartPos:  4544,
					EndPos:    4557,
				},
				Expr: &expr.Isset{
					Position: &position.Position{
						StartLine: 223,
						EndLine:   223,
						StartPos:  4544,
						EndPos:    4556,
					},
					Variables: []node.Node{
						&expr.Variable{
							Position: &position.Position{
								StartLine: 223,
								EndLine:   223,
								StartPos:  4550,
								EndPos:    4551,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 223,
									EndLine:   223,
									StartPos:  4550,
									EndPos:    4551,
								},
								Value: "a",
							},
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 223,
								EndLine:   223,
								StartPos:  4554,
								EndPos:    4555,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 223,
									EndLine:   223,
									StartPos:  4554,
									EndPos:    4555,
								},
								Value: "b",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 224,
					EndLine:   224,
					StartPos:  4561,
					EndPos:    4574,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 224,
						EndLine:   224,
						StartPos:  4561,
						EndPos:    4573,
					},
					Variable: &expr.List{
						Position: &position.Position{
							StartLine: 224,
							EndLine:   224,
							StartPos:  4561,
							EndPos:    4568,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 224,
									EndLine:   224,
									StartPos:  4566,
									EndPos:    4567,
								},
								Val: &expr.Variable{
									Position: &position.Position{
										StartLine: 224,
										EndLine:   224,
										StartPos:  4566,
										EndPos:    4567,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 224,
											EndLine:   224,
											StartPos:  4566,
											EndPos:    4567,
										},
										Value: "a",
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 224,
							EndLine:   224,
							StartPos:  4572,
							EndPos:    4573,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 224,
								EndLine:   224,
								StartPos:  4572,
								EndPos:    4573,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 225,
					EndLine:   225,
					StartPos:  4578,
					EndPos:    4593,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 225,
						EndLine:   225,
						StartPos:  4578,
						EndPos:    4592,
					},
					Variable: &expr.List{
						Position: &position.Position{
							StartLine: 225,
							EndLine:   225,
							StartPos:  4578,
							EndPos:    4587,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 225,
									EndLine:   225,
									StartPos:  4583,
									EndPos:    4586,
								},
								Val: &expr.ArrayDimFetch{
									Position: &position.Position{
										StartLine: 225,
										EndLine:   225,
										StartPos:  4583,
										EndPos:    4586,
									},
									Variable: &expr.Variable{
										Position: &position.Position{
											StartLine: 225,
											EndLine:   225,
											StartPos:  4583,
											EndPos:    4584,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 225,
												EndLine:   225,
												StartPos:  4583,
												EndPos:    4584,
											},
											Value: "a",
										},
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 225,
							EndLine:   225,
							StartPos:  4591,
							EndPos:    4592,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 225,
								EndLine:   225,
								StartPos:  4591,
								EndPos:    4592,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 226,
					EndLine:   226,
					StartPos:  4597,
					EndPos:    4616,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 226,
						EndLine:   226,
						StartPos:  4597,
						EndPos:    4615,
					},
					Variable: &expr.List{
						Position: &position.Position{
							StartLine: 226,
							EndLine:   226,
							StartPos:  4597,
							EndPos:    4610,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 226,
									EndLine:   226,
									StartPos:  4602,
									EndPos:    4609,
								},
								Val: &expr.List{
									Position: &position.Position{
										StartLine: 226,
										EndLine:   226,
										StartPos:  4602,
										EndPos:    4609,
									},
									Items: []node.Node{
										&expr.ArrayItem{
											Position: &position.Position{
												StartLine: 226,
												EndLine:   226,
												StartPos:  4607,
												EndPos:    4608,
											},
											Val: &expr.Variable{
												Position: &position.Position{
													StartLine: 226,
													EndLine:   226,
													StartPos:  4607,
													EndPos:    4608,
												},
												VarName: &node.Identifier{
													Position: &position.Position{
														StartLine: 226,
														EndLine:   226,
														StartPos:  4607,
														EndPos:    4608,
													},
													Value: "a",
												},
											},
										},
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 226,
							EndLine:   226,
							StartPos:  4614,
							EndPos:    4615,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 226,
								EndLine:   226,
								StartPos:  4614,
								EndPos:    4615,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 228,
					EndLine:   228,
					StartPos:  4621,
					EndPos:    4630,
				},
				Expr: &expr.MethodCall{
					Position: &position.Position{
						StartLine: 228,
						EndLine:   228,
						StartPos:  4621,
						EndPos:    4629,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 228,
							EndLine:   228,
							StartPos:  4621,
							EndPos:    4622,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 228,
								EndLine:   228,
								StartPos:  4621,
								EndPos:    4622,
							},
							Value: "a",
						},
					},
					Method: &node.Identifier{
						Position: &position.Position{
							StartLine: 228,
							EndLine:   228,
							StartPos:  4625,
							EndPos:    4627,
						},
						Value: "foo",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 228,
							EndLine:   228,
							StartPos:  4628,
							EndPos:    4629,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 229,
					EndLine:   229,
					StartPos:  4634,
					EndPos:    4643,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 229,
						EndLine:   229,
						StartPos:  4634,
						EndPos:    4642,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 229,
							EndLine:   229,
							StartPos:  4638,
							EndPos:    4640,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 229,
									EndLine:   229,
									StartPos:  4638,
									EndPos:    4640,
								},
								Value: "Foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 229,
							EndLine:   229,
							StartPos:  4641,
							EndPos:    4642,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 230,
					EndLine:   230,
					StartPos:  4647,
					EndPos:    4666,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 230,
						EndLine:   230,
						StartPos:  4647,
						EndPos:    4665,
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 230,
							EndLine:   230,
							StartPos:  4651,
							EndPos:    4663,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 230,
									EndLine:   230,
									StartPos:  4661,
									EndPos:    4663,
								},
								Value: "Foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 230,
							EndLine:   230,
							StartPos:  4664,
							EndPos:    4665,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 231,
					EndLine:   231,
					StartPos:  4670,
					EndPos:    4680,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 231,
						EndLine:   231,
						StartPos:  4670,
						EndPos:    4679,
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 231,
							EndLine:   231,
							StartPos:  4674,
							EndPos:    4677,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 231,
									EndLine:   231,
									StartPos:  4675,
									EndPos:    4677,
								},
								Value: "Foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 231,
							EndLine:   231,
							StartPos:  4678,
							EndPos:    4679,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 232,
					EndLine:   232,
					StartPos:  4684,
					EndPos:    4708,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 232,
						EndLine:   232,
						StartPos:  4684,
						EndPos:    4707,
					},
					Class: &stmt.Class{
						Position: &position.Position{
							StartLine: 232,
							EndLine:   232,
							StartPos:  4688,
							EndPos:    4707,
						},
						PhpDocComment: "",
						ArgumentList: &node.ArgumentList{
							Position: &position.Position{
								StartLine: 232,
								EndLine:   232,
								StartPos:  4694,
								EndPos:    4704,
							},
							Arguments: []node.Node{
								&node.Argument{
									Position: &position.Position{
										StartLine: 232,
										EndLine:   232,
										StartPos:  4695,
										EndPos:    4696,
									},
									Variadic:    false,
									IsReference: false,
									Expr: &expr.Variable{
										Position: &position.Position{
											StartLine: 232,
											EndLine:   232,
											StartPos:  4695,
											EndPos:    4696,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 232,
												EndLine:   232,
												StartPos:  4695,
												EndPos:    4696,
											},
											Value: "a",
										},
									},
								},
								&node.Argument{
									Position: &position.Position{
										StartLine: 232,
										EndLine:   232,
										StartPos:  4699,
										EndPos:    4703,
									},
									Variadic:    true,
									IsReference: false,
									Expr: &expr.Variable{
										Position: &position.Position{
											StartLine: 232,
											EndLine:   232,
											StartPos:  4702,
											EndPos:    4703,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 232,
												EndLine:   232,
												StartPos:  4702,
												EndPos:    4703,
											},
											Value: "b",
										},
									},
								},
							},
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 233,
					EndLine:   233,
					StartPos:  4712,
					EndPos:    4721,
				},
				Expr: &expr.Print{
					Position: &position.Position{
						StartLine: 233,
						EndLine:   233,
						StartPos:  4712,
						EndPos:    4719,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 233,
							EndLine:   233,
							StartPos:  4718,
							EndPos:    4719,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 233,
								EndLine:   233,
								StartPos:  4718,
								EndPos:    4719,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 234,
					EndLine:   234,
					StartPos:  4725,
					EndPos:    4732,
				},
				Expr: &expr.PropertyFetch{
					Position: &position.Position{
						StartLine: 234,
						EndLine:   234,
						StartPos:  4725,
						EndPos:    4731,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 234,
							EndLine:   234,
							StartPos:  4725,
							EndPos:    4726,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 234,
								EndLine:   234,
								StartPos:  4725,
								EndPos:    4726,
							},
							Value: "a",
						},
					},
					Property: &node.Identifier{
						Position: &position.Position{
							StartLine: 234,
							EndLine:   234,
							StartPos:  4729,
							EndPos:    4731,
						},
						Value: "foo",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 235,
					EndLine:   235,
					StartPos:  4736,
					EndPos:    4744,
				},
				Expr: &expr.ShellExec{
					Position: &position.Position{
						StartLine: 235,
						EndLine:   235,
						StartPos:  4736,
						EndPos:    4743,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 235,
								EndLine:   235,
								StartPos:  4737,
								EndPos:    4740,
							},
							Value: "cmd ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 235,
								EndLine:   235,
								StartPos:  4741,
								EndPos:    4742,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 235,
									EndLine:   235,
									StartPos:  4741,
									EndPos:    4742,
								},
								Value: "a",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 236,
					EndLine:   236,
					StartPos:  4748,
					EndPos:    4753,
				},
				Expr: &expr.ShellExec{
					Position: &position.Position{
						StartLine: 236,
						EndLine:   236,
						StartPos:  4748,
						EndPos:    4752,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Value: "cmd",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 237,
					EndLine:   237,
					StartPos:  4757,
					EndPos:    4759,
				},
				Expr: &expr.ShellExec{
					Position: &position.Position{
						StartLine: 237,
						EndLine:   237,
						StartPos:  4757,
						EndPos:    4758,
					},
					Parts: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 238,
					EndLine:   238,
					StartPos:  4763,
					EndPos:    4765,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 238,
						EndLine:   238,
						StartPos:  4763,
						EndPos:    4764,
					},
					Items: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 239,
					EndLine:   239,
					StartPos:  4769,
					EndPos:    4772,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 239,
						EndLine:   239,
						StartPos:  4769,
						EndPos:    4771,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 239,
								EndLine:   239,
								StartPos:  4770,
								EndPos:    4770,
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 239,
									EndLine:   239,
									StartPos:  4770,
									EndPos:    4770,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 240,
					EndLine:   240,
					StartPos:  4776,
					EndPos:    4788,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 240,
						EndLine:   240,
						StartPos:  4776,
						EndPos:    4787,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 240,
								EndLine:   240,
								StartPos:  4777,
								EndPos:    4780,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 240,
									EndLine:   240,
									StartPos:  4777,
									EndPos:    4777,
								},
								Value: "1",
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 240,
									EndLine:   240,
									StartPos:  4780,
									EndPos:    4780,
								},
								Value: "1",
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 240,
								EndLine:   240,
								StartPos:  4783,
								EndPos:    4785,
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 240,
									EndLine:   240,
									StartPos:  4783,
									EndPos:    4785,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 240,
										EndLine:   240,
										StartPos:  4784,
										EndPos:    4785,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 240,
											EndLine:   240,
											StartPos:  4784,
											EndPos:    4785,
										},
										Value: "b",
									},
								},
							},
						},
						&expr.ArrayItem{},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 242,
					EndLine:   242,
					StartPos:  4793,
					EndPos:    4802,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 242,
						EndLine:   242,
						StartPos:  4793,
						EndPos:    4801,
					},
					Variable: &expr.ShortList{
						Position: &position.Position{
							StartLine: 242,
							EndLine:   242,
							StartPos:  4793,
							EndPos:    4796,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 242,
									EndLine:   242,
									StartPos:  4794,
									EndPos:    4795,
								},
								Val: &expr.Variable{
									Position: &position.Position{
										StartLine: 242,
										EndLine:   242,
										StartPos:  4794,
										EndPos:    4795,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 242,
											EndLine:   242,
											StartPos:  4794,
											EndPos:    4795,
										},
										Value: "a",
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 242,
							EndLine:   242,
							StartPos:  4800,
							EndPos:    4801,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 242,
								EndLine:   242,
								StartPos:  4800,
								EndPos:    4801,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 243,
					EndLine:   243,
					StartPos:  4806,
					EndPos:    4817,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 243,
						EndLine:   243,
						StartPos:  4806,
						EndPos:    4816,
					},
					Variable: &expr.ShortList{
						Position: &position.Position{
							StartLine: 243,
							EndLine:   243,
							StartPos:  4806,
							EndPos:    4811,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 243,
									EndLine:   243,
									StartPos:  4807,
									EndPos:    4810,
								},
								Val: &expr.ArrayDimFetch{
									Position: &position.Position{
										StartLine: 243,
										EndLine:   243,
										StartPos:  4807,
										EndPos:    4810,
									},
									Variable: &expr.Variable{
										Position: &position.Position{
											StartLine: 243,
											EndLine:   243,
											StartPos:  4807,
											EndPos:    4808,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 243,
												EndLine:   243,
												StartPos:  4807,
												EndPos:    4808,
											},
											Value: "a",
										},
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 243,
							EndLine:   243,
							StartPos:  4815,
							EndPos:    4816,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 243,
								EndLine:   243,
								StartPos:  4815,
								EndPos:    4816,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 244,
					EndLine:   244,
					StartPos:  4821,
					EndPos:    4836,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 244,
						EndLine:   244,
						StartPos:  4821,
						EndPos:    4835,
					},
					Variable: &expr.ShortList{
						Position: &position.Position{
							StartLine: 244,
							EndLine:   244,
							StartPos:  4821,
							EndPos:    4830,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 244,
									EndLine:   244,
									StartPos:  4822,
									EndPos:    4829,
								},
								Val: &expr.List{
									Position: &position.Position{
										StartLine: 244,
										EndLine:   244,
										StartPos:  4822,
										EndPos:    4829,
									},
									Items: []node.Node{
										&expr.ArrayItem{
											Position: &position.Position{
												StartLine: 244,
												EndLine:   244,
												StartPos:  4827,
												EndPos:    4828,
											},
											Val: &expr.Variable{
												Position: &position.Position{
													StartLine: 244,
													EndLine:   244,
													StartPos:  4827,
													EndPos:    4828,
												},
												VarName: &node.Identifier{
													Position: &position.Position{
														StartLine: 244,
														EndLine:   244,
														StartPos:  4827,
														EndPos:    4828,
													},
													Value: "a",
												},
											},
										},
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 244,
							EndLine:   244,
							StartPos:  4834,
							EndPos:    4835,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 244,
								EndLine:   244,
								StartPos:  4834,
								EndPos:    4835,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 245,
					EndLine:   245,
					StartPos:  4840,
					EndPos:    4850,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 245,
						EndLine:   245,
						StartPos:  4840,
						EndPos:    4849,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4840,
							EndPos:    4842,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4840,
									EndPos:    4842,
								},
								Value: "Foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4845,
							EndPos:    4847,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4848,
							EndPos:    4849,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 246,
					EndLine:   246,
					StartPos:  4854,
					EndPos:    4874,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 246,
						EndLine:   246,
						StartPos:  4854,
						EndPos:    4873,
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4854,
							EndPos:    4866,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4864,
									EndPos:    4866,
								},
								Value: "Foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4869,
							EndPos:    4871,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4872,
							EndPos:    4873,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 247,
					EndLine:   247,
					StartPos:  4878,
					EndPos:    4889,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 247,
						EndLine:   247,
						StartPos:  4878,
						EndPos:    4888,
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 247,
							EndLine:   247,
							StartPos:  4878,
							EndPos:    4881,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 247,
									EndLine:   247,
									StartPos:  4879,
									EndPos:    4881,
								},
								Value: "Foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 247,
							EndLine:   247,
							StartPos:  4884,
							EndPos:    4886,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 247,
							EndLine:   247,
							StartPos:  4887,
							EndPos:    4888,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 248,
					EndLine:   248,
					StartPos:  4893,
					EndPos:    4902,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 248,
						EndLine:   248,
						StartPos:  4893,
						EndPos:    4901,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 248,
							EndLine:   248,
							StartPos:  4893,
							EndPos:    4895,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 248,
									EndLine:   248,
									StartPos:  4893,
									EndPos:    4895,
								},
								Value: "Foo",
							},
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 248,
							EndLine:   248,
							StartPos:  4898,
							EndPos:    4901,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4898,
								EndPos:    4901,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 249,
					EndLine:   249,
					StartPos:  4906,
					EndPos:    4916,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 249,
						EndLine:   249,
						StartPos:  4906,
						EndPos:    4915,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 249,
							EndLine:   249,
							StartPos:  4906,
							EndPos:    4909,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 249,
								EndLine:   249,
								StartPos:  4906,
								EndPos:    4909,
							},
							Value: "foo",
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 249,
							EndLine:   249,
							StartPos:  4912,
							EndPos:    4915,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 249,
								EndLine:   249,
								StartPos:  4912,
								EndPos:    4915,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 250,
					EndLine:   250,
					StartPos:  4920,
					EndPos:    4939,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 250,
						EndLine:   250,
						StartPos:  4920,
						EndPos:    4938,
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 250,
							EndLine:   250,
							StartPos:  4920,
							EndPos:    4932,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 250,
									EndLine:   250,
									StartPos:  4930,
									EndPos:    4932,
								},
								Value: "Foo",
							},
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 250,
							EndLine:   250,
							StartPos:  4935,
							EndPos:    4938,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 250,
								EndLine:   250,
								StartPos:  4935,
								EndPos:    4938,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 251,
					EndLine:   251,
					StartPos:  4943,
					EndPos:    4953,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 251,
						EndLine:   251,
						StartPos:  4943,
						EndPos:    4952,
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 251,
							EndLine:   251,
							StartPos:  4943,
							EndPos:    4946,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 251,
									EndLine:   251,
									StartPos:  4944,
									EndPos:    4946,
								},
								Value: "Foo",
							},
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 251,
							EndLine:   251,
							StartPos:  4949,
							EndPos:    4952,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 251,
								EndLine:   251,
								StartPos:  4949,
								EndPos:    4952,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 252,
					EndLine:   252,
					StartPos:  4957,
					EndPos:    4969,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 252,
						EndLine:   252,
						StartPos:  4957,
						EndPos:    4968,
					},
					Condition: &expr.Variable{
						Position: &position.Position{
							StartLine: 252,
							EndLine:   252,
							StartPos:  4957,
							EndPos:    4958,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 252,
								EndLine:   252,
								StartPos:  4957,
								EndPos:    4958,
							},
							Value: "a",
						},
					},
					IfTrue: &expr.Variable{
						Position: &position.Position{
							StartLine: 252,
							EndLine:   252,
							StartPos:  4962,
							EndPos:    4963,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 252,
								EndLine:   252,
								StartPos:  4962,
								EndPos:    4963,
							},
							Value: "b",
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 252,
							EndLine:   252,
							StartPos:  4967,
							EndPos:    4968,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 252,
								EndLine:   252,
								StartPos:  4967,
								EndPos:    4968,
							},
							Value: "c",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 253,
					EndLine:   253,
					StartPos:  4973,
					EndPos:    4982,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 253,
						EndLine:   253,
						StartPos:  4973,
						EndPos:    4981,
					},
					Condition: &expr.Variable{
						Position: &position.Position{
							StartLine: 253,
							EndLine:   253,
							StartPos:  4973,
							EndPos:    4974,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4973,
								EndPos:    4974,
							},
							Value: "a",
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 253,
							EndLine:   253,
							StartPos:  4980,
							EndPos:    4981,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4980,
								EndPos:    4981,
							},
							Value: "c",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 254,
					EndLine:   254,
					StartPos:  4986,
					EndPos:    5008,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 254,
						EndLine:   254,
						StartPos:  4986,
						EndPos:    5007,
					},
					Condition: &expr.Variable{
						Position: &position.Position{
							StartLine: 254,
							EndLine:   254,
							StartPos:  4986,
							EndPos:    4987,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  4986,
								EndPos:    4987,
							},
							Value: "a",
						},
					},
					IfTrue: &expr.Ternary{
						Position: &position.Position{
							StartLine: 254,
							EndLine:   254,
							StartPos:  4991,
							EndPos:    5002,
						},
						Condition: &expr.Variable{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  4991,
								EndPos:    4992,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4991,
									EndPos:    4992,
								},
								Value: "b",
							},
						},
						IfTrue: &expr.Variable{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  4996,
								EndPos:    4997,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4996,
									EndPos:    4997,
								},
								Value: "c",
							},
						},
						IfFalse: &expr.Variable{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  5001,
								EndPos:    5002,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  5001,
									EndPos:    5002,
								},
								Value: "d",
							},
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 254,
							EndLine:   254,
							StartPos:  5006,
							EndPos:    5007,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  5006,
								EndPos:    5007,
							},
							Value: "e",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 255,
					EndLine:   255,
					StartPos:  5012,
					EndPos:    5034,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 255,
						EndLine:   255,
						StartPos:  5012,
						EndPos:    5033,
					},
					Condition: &expr.Ternary{
						Position: &position.Position{
							StartLine: 255,
							EndLine:   255,
							StartPos:  5012,
							EndPos:    5023,
						},
						Condition: &expr.Variable{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  5012,
								EndPos:    5013,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 255,
									EndLine:   255,
									StartPos:  5012,
									EndPos:    5013,
								},
								Value: "a",
							},
						},
						IfTrue: &expr.Variable{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  5017,
								EndPos:    5018,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 255,
									EndLine:   255,
									StartPos:  5017,
									EndPos:    5018,
								},
								Value: "b",
							},
						},
						IfFalse: &expr.Variable{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  5022,
								EndPos:    5023,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 255,
									EndLine:   255,
									StartPos:  5022,
									EndPos:    5023,
								},
								Value: "c",
							},
						},
					},
					IfTrue: &expr.Variable{
						Position: &position.Position{
							StartLine: 255,
							EndLine:   255,
							StartPos:  5027,
							EndPos:    5028,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  5027,
								EndPos:    5028,
							},
							Value: "d",
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 255,
							EndLine:   255,
							StartPos:  5032,
							EndPos:    5033,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  5032,
								EndPos:    5033,
							},
							Value: "e",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 256,
					EndLine:   256,
					StartPos:  5038,
					EndPos:    5041,
				},
				Expr: &expr.UnaryMinus{
					Position: &position.Position{
						StartLine: 256,
						EndLine:   256,
						StartPos:  5038,
						EndPos:    5040,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 256,
							EndLine:   256,
							StartPos:  5039,
							EndPos:    5040,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 256,
								EndLine:   256,
								StartPos:  5039,
								EndPos:    5040,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 257,
					EndLine:   257,
					StartPos:  5045,
					EndPos:    5048,
				},
				Expr: &expr.UnaryPlus{
					Position: &position.Position{
						StartLine: 257,
						EndLine:   257,
						StartPos:  5045,
						EndPos:    5047,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 257,
							EndLine:   257,
							StartPos:  5046,
							EndPos:    5047,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 257,
								EndLine:   257,
								StartPos:  5046,
								EndPos:    5047,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 258,
					EndLine:   258,
					StartPos:  5052,
					EndPos:    5055,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 258,
						EndLine:   258,
						StartPos:  5052,
						EndPos:    5054,
					},
					VarName: &expr.Variable{
						Position: &position.Position{
							StartLine: 258,
							EndLine:   258,
							StartPos:  5053,
							EndPos:    5054,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 258,
								EndLine:   258,
								StartPos:  5053,
								EndPos:    5054,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 259,
					EndLine:   259,
					StartPos:  5059,
					EndPos:    5064,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 259,
						EndLine:   259,
						StartPos:  5059,
						EndPos:    5063,
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 260,
					EndLine:   260,
					StartPos:  5068,
					EndPos:    5076,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 260,
						EndLine:   260,
						StartPos:  5068,
						EndPos:    5075,
					},
					Value: &expr.Variable{
						Position: &position.Position{
							StartLine: 260,
							EndLine:   260,
							StartPos:  5074,
							EndPos:    5075,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 260,
								EndLine:   260,
								StartPos:  5074,
								EndPos:    5075,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 261,
					EndLine:   261,
					StartPos:  5080,
					EndPos:    5094,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 261,
						EndLine:   261,
						StartPos:  5080,
						EndPos:    5093,
					},
					Key: &expr.Variable{
						Position: &position.Position{
							StartLine: 261,
							EndLine:   261,
							StartPos:  5086,
							EndPos:    5087,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 261,
								EndLine:   261,
								StartPos:  5086,
								EndPos:    5087,
							},
							Value: "a",
						},
					},
					Value: &expr.Variable{
						Position: &position.Position{
							StartLine: 261,
							EndLine:   261,
							StartPos:  5092,
							EndPos:    5093,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 261,
								EndLine:   261,
								StartPos:  5092,
								EndPos:    5093,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 262,
					EndLine:   262,
					StartPos:  5098,
					EndPos:    5111,
				},
				Expr: &expr.YieldFrom{
					Position: &position.Position{
						StartLine: 262,
						EndLine:   262,
						StartPos:  5098,
						EndPos:    5110,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 262,
							EndLine:   262,
							StartPos:  5109,
							EndPos:    5110,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 262,
								EndLine:   262,
								StartPos:  5109,
								EndPos:    5110,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 264,
					EndLine:   264,
					StartPos:  5118,
					EndPos:    5127,
				},
				Expr: &cast.Array{
					Position: &position.Position{
						StartLine: 264,
						EndLine:   264,
						StartPos:  5118,
						EndPos:    5126,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 264,
							EndLine:   264,
							StartPos:  5125,
							EndPos:    5126,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 264,
								EndLine:   264,
								StartPos:  5125,
								EndPos:    5126,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 265,
					EndLine:   265,
					StartPos:  5131,
					EndPos:    5142,
				},
				Expr: &cast.Bool{
					Position: &position.Position{
						StartLine: 265,
						EndLine:   265,
						StartPos:  5131,
						EndPos:    5141,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 265,
							EndLine:   265,
							StartPos:  5140,
							EndPos:    5141,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 265,
								EndLine:   265,
								StartPos:  5140,
								EndPos:    5141,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 266,
					EndLine:   266,
					StartPos:  5146,
					EndPos:    5154,
				},
				Expr: &cast.Bool{
					Position: &position.Position{
						StartLine: 266,
						EndLine:   266,
						StartPos:  5146,
						EndPos:    5153,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 266,
							EndLine:   266,
							StartPos:  5152,
							EndPos:    5153,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 266,
								EndLine:   266,
								StartPos:  5152,
								EndPos:    5153,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 267,
					EndLine:   267,
					StartPos:  5158,
					EndPos:    5168,
				},
				Expr: &cast.Double{
					Position: &position.Position{
						StartLine: 267,
						EndLine:   267,
						StartPos:  5158,
						EndPos:    5167,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 267,
							EndLine:   267,
							StartPos:  5166,
							EndPos:    5167,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 267,
								EndLine:   267,
								StartPos:  5166,
								EndPos:    5167,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 268,
					EndLine:   268,
					StartPos:  5172,
					EndPos:    5181,
				},
				Expr: &cast.Double{
					Position: &position.Position{
						StartLine: 268,
						EndLine:   268,
						StartPos:  5172,
						EndPos:    5180,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 268,
							EndLine:   268,
							StartPos:  5179,
							EndPos:    5180,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 268,
								EndLine:   268,
								StartPos:  5179,
								EndPos:    5180,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 269,
					EndLine:   269,
					StartPos:  5185,
					EndPos:    5196,
				},
				Expr: &cast.Int{
					Position: &position.Position{
						StartLine: 269,
						EndLine:   269,
						StartPos:  5185,
						EndPos:    5195,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 269,
							EndLine:   269,
							StartPos:  5194,
							EndPos:    5195,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 269,
								EndLine:   269,
								StartPos:  5194,
								EndPos:    5195,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 270,
					EndLine:   270,
					StartPos:  5200,
					EndPos:    5207,
				},
				Expr: &cast.Int{
					Position: &position.Position{
						StartLine: 270,
						EndLine:   270,
						StartPos:  5200,
						EndPos:    5206,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 270,
							EndLine:   270,
							StartPos:  5205,
							EndPos:    5206,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 270,
								EndLine:   270,
								StartPos:  5205,
								EndPos:    5206,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 271,
					EndLine:   271,
					StartPos:  5211,
					EndPos:    5221,
				},
				Expr: &cast.Object{
					Position: &position.Position{
						StartLine: 271,
						EndLine:   271,
						StartPos:  5211,
						EndPos:    5220,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 271,
							EndLine:   271,
							StartPos:  5219,
							EndPos:    5220,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 271,
								EndLine:   271,
								StartPos:  5219,
								EndPos:    5220,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 272,
					EndLine:   272,
					StartPos:  5225,
					EndPos:    5235,
				},
				Expr: &cast.String{
					Position: &position.Position{
						StartLine: 272,
						EndLine:   272,
						StartPos:  5225,
						EndPos:    5234,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 272,
							EndLine:   272,
							StartPos:  5233,
							EndPos:    5234,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 272,
								EndLine:   272,
								StartPos:  5233,
								EndPos:    5234,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 273,
					EndLine:   273,
					StartPos:  5239,
					EndPos:    5248,
				},
				Expr: &cast.Unset{
					Position: &position.Position{
						StartLine: 273,
						EndLine:   273,
						StartPos:  5239,
						EndPos:    5247,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 273,
							EndLine:   273,
							StartPos:  5246,
							EndPos:    5247,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 273,
								EndLine:   273,
								StartPos:  5246,
								EndPos:    5247,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 275,
					EndLine:   275,
					StartPos:  5253,
					EndPos:    5260,
				},
				Expr: &binary.BitwiseAnd{
					Position: &position.Position{
						StartLine: 275,
						EndLine:   275,
						StartPos:  5253,
						EndPos:    5259,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 275,
							EndLine:   275,
							StartPos:  5253,
							EndPos:    5254,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 275,
								EndLine:   275,
								StartPos:  5253,
								EndPos:    5254,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 275,
							EndLine:   275,
							StartPos:  5258,
							EndPos:    5259,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 275,
								EndLine:   275,
								StartPos:  5258,
								EndPos:    5259,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 276,
					EndLine:   276,
					StartPos:  5264,
					EndPos:    5271,
				},
				Expr: &binary.BitwiseOr{
					Position: &position.Position{
						StartLine: 276,
						EndLine:   276,
						StartPos:  5264,
						EndPos:    5270,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 276,
							EndLine:   276,
							StartPos:  5264,
							EndPos:    5265,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 276,
								EndLine:   276,
								StartPos:  5264,
								EndPos:    5265,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 276,
							EndLine:   276,
							StartPos:  5269,
							EndPos:    5270,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 276,
								EndLine:   276,
								StartPos:  5269,
								EndPos:    5270,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 277,
					EndLine:   277,
					StartPos:  5275,
					EndPos:    5282,
				},
				Expr: &binary.BitwiseXor{
					Position: &position.Position{
						StartLine: 277,
						EndLine:   277,
						StartPos:  5275,
						EndPos:    5281,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 277,
							EndLine:   277,
							StartPos:  5275,
							EndPos:    5276,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 277,
								EndLine:   277,
								StartPos:  5275,
								EndPos:    5276,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 277,
							EndLine:   277,
							StartPos:  5280,
							EndPos:    5281,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 277,
								EndLine:   277,
								StartPos:  5280,
								EndPos:    5281,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 278,
					EndLine:   278,
					StartPos:  5286,
					EndPos:    5294,
				},
				Expr: &binary.BooleanAnd{
					Position: &position.Position{
						StartLine: 278,
						EndLine:   278,
						StartPos:  5286,
						EndPos:    5293,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 278,
							EndLine:   278,
							StartPos:  5286,
							EndPos:    5287,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 278,
								EndLine:   278,
								StartPos:  5286,
								EndPos:    5287,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 278,
							EndLine:   278,
							StartPos:  5292,
							EndPos:    5293,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 278,
								EndLine:   278,
								StartPos:  5292,
								EndPos:    5293,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 279,
					EndLine:   279,
					StartPos:  5298,
					EndPos:    5306,
				},
				Expr: &binary.BooleanOr{
					Position: &position.Position{
						StartLine: 279,
						EndLine:   279,
						StartPos:  5298,
						EndPos:    5305,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 279,
							EndLine:   279,
							StartPos:  5298,
							EndPos:    5299,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 279,
								EndLine:   279,
								StartPos:  5298,
								EndPos:    5299,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 279,
							EndLine:   279,
							StartPos:  5304,
							EndPos:    5305,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 279,
								EndLine:   279,
								StartPos:  5304,
								EndPos:    5305,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 280,
					EndLine:   280,
					StartPos:  5310,
					EndPos:    5318,
				},
				Expr: &binary.Coalesce{
					Position: &position.Position{
						StartLine: 280,
						EndLine:   280,
						StartPos:  5310,
						EndPos:    5317,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 280,
							EndLine:   280,
							StartPos:  5310,
							EndPos:    5311,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 280,
								EndLine:   280,
								StartPos:  5310,
								EndPos:    5311,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 280,
							EndLine:   280,
							StartPos:  5316,
							EndPos:    5317,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 280,
								EndLine:   280,
								StartPos:  5316,
								EndPos:    5317,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 281,
					EndLine:   281,
					StartPos:  5322,
					EndPos:    5329,
				},
				Expr: &binary.Concat{
					Position: &position.Position{
						StartLine: 281,
						EndLine:   281,
						StartPos:  5322,
						EndPos:    5328,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 281,
							EndLine:   281,
							StartPos:  5322,
							EndPos:    5323,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 281,
								EndLine:   281,
								StartPos:  5322,
								EndPos:    5323,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 281,
							EndLine:   281,
							StartPos:  5327,
							EndPos:    5328,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 281,
								EndLine:   281,
								StartPos:  5327,
								EndPos:    5328,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 282,
					EndLine:   282,
					StartPos:  5333,
					EndPos:    5340,
				},
				Expr: &binary.Div{
					Position: &position.Position{
						StartLine: 282,
						EndLine:   282,
						StartPos:  5333,
						EndPos:    5339,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 282,
							EndLine:   282,
							StartPos:  5333,
							EndPos:    5334,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 282,
								EndLine:   282,
								StartPos:  5333,
								EndPos:    5334,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 282,
							EndLine:   282,
							StartPos:  5338,
							EndPos:    5339,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 282,
								EndLine:   282,
								StartPos:  5338,
								EndPos:    5339,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 283,
					EndLine:   283,
					StartPos:  5344,
					EndPos:    5352,
				},
				Expr: &binary.Equal{
					Position: &position.Position{
						StartLine: 283,
						EndLine:   283,
						StartPos:  5344,
						EndPos:    5351,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 283,
							EndLine:   283,
							StartPos:  5344,
							EndPos:    5345,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 283,
								EndLine:   283,
								StartPos:  5344,
								EndPos:    5345,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 283,
							EndLine:   283,
							StartPos:  5350,
							EndPos:    5351,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 283,
								EndLine:   283,
								StartPos:  5350,
								EndPos:    5351,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 284,
					EndLine:   284,
					StartPos:  5356,
					EndPos:    5364,
				},
				Expr: &binary.GreaterOrEqual{
					Position: &position.Position{
						StartLine: 284,
						EndLine:   284,
						StartPos:  5356,
						EndPos:    5363,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 284,
							EndLine:   284,
							StartPos:  5356,
							EndPos:    5357,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 284,
								EndLine:   284,
								StartPos:  5356,
								EndPos:    5357,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 284,
							EndLine:   284,
							StartPos:  5362,
							EndPos:    5363,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 284,
								EndLine:   284,
								StartPos:  5362,
								EndPos:    5363,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 285,
					EndLine:   285,
					StartPos:  5368,
					EndPos:    5375,
				},
				Expr: &binary.Greater{
					Position: &position.Position{
						StartLine: 285,
						EndLine:   285,
						StartPos:  5368,
						EndPos:    5374,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 285,
							EndLine:   285,
							StartPos:  5368,
							EndPos:    5369,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 285,
								EndLine:   285,
								StartPos:  5368,
								EndPos:    5369,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 285,
							EndLine:   285,
							StartPos:  5373,
							EndPos:    5374,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 285,
								EndLine:   285,
								StartPos:  5373,
								EndPos:    5374,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 286,
					EndLine:   286,
					StartPos:  5379,
					EndPos:    5388,
				},
				Expr: &binary.Identical{
					Position: &position.Position{
						StartLine: 286,
						EndLine:   286,
						StartPos:  5379,
						EndPos:    5387,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 286,
							EndLine:   286,
							StartPos:  5379,
							EndPos:    5380,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 286,
								EndLine:   286,
								StartPos:  5379,
								EndPos:    5380,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 286,
							EndLine:   286,
							StartPos:  5386,
							EndPos:    5387,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 286,
								EndLine:   286,
								StartPos:  5386,
								EndPos:    5387,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 287,
					EndLine:   287,
					StartPos:  5392,
					EndPos:    5401,
				},
				Expr: &binary.LogicalAnd{
					Position: &position.Position{
						StartLine: 287,
						EndLine:   287,
						StartPos:  5392,
						EndPos:    5400,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 287,
							EndLine:   287,
							StartPos:  5392,
							EndPos:    5393,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 287,
								EndLine:   287,
								StartPos:  5392,
								EndPos:    5393,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 287,
							EndLine:   287,
							StartPos:  5399,
							EndPos:    5400,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 287,
								EndLine:   287,
								StartPos:  5399,
								EndPos:    5400,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 288,
					EndLine:   288,
					StartPos:  5405,
					EndPos:    5413,
				},
				Expr: &binary.LogicalOr{
					Position: &position.Position{
						StartLine: 288,
						EndLine:   288,
						StartPos:  5405,
						EndPos:    5412,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 288,
							EndLine:   288,
							StartPos:  5405,
							EndPos:    5406,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 288,
								EndLine:   288,
								StartPos:  5405,
								EndPos:    5406,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 288,
							EndLine:   288,
							StartPos:  5411,
							EndPos:    5412,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 288,
								EndLine:   288,
								StartPos:  5411,
								EndPos:    5412,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 289,
					EndLine:   289,
					StartPos:  5417,
					EndPos:    5426,
				},
				Expr: &binary.LogicalXor{
					Position: &position.Position{
						StartLine: 289,
						EndLine:   289,
						StartPos:  5417,
						EndPos:    5425,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 289,
							EndLine:   289,
							StartPos:  5417,
							EndPos:    5418,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 289,
								EndLine:   289,
								StartPos:  5417,
								EndPos:    5418,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 289,
							EndLine:   289,
							StartPos:  5424,
							EndPos:    5425,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 289,
								EndLine:   289,
								StartPos:  5424,
								EndPos:    5425,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 290,
					EndLine:   290,
					StartPos:  5430,
					EndPos:    5437,
				},
				Expr: &binary.Minus{
					Position: &position.Position{
						StartLine: 290,
						EndLine:   290,
						StartPos:  5430,
						EndPos:    5436,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 290,
							EndLine:   290,
							StartPos:  5430,
							EndPos:    5431,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 290,
								EndLine:   290,
								StartPos:  5430,
								EndPos:    5431,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 290,
							EndLine:   290,
							StartPos:  5435,
							EndPos:    5436,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 290,
								EndLine:   290,
								StartPos:  5435,
								EndPos:    5436,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 291,
					EndLine:   291,
					StartPos:  5441,
					EndPos:    5448,
				},
				Expr: &binary.Mod{
					Position: &position.Position{
						StartLine: 291,
						EndLine:   291,
						StartPos:  5441,
						EndPos:    5447,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 291,
							EndLine:   291,
							StartPos:  5441,
							EndPos:    5442,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 291,
								EndLine:   291,
								StartPos:  5441,
								EndPos:    5442,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 291,
							EndLine:   291,
							StartPos:  5446,
							EndPos:    5447,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 291,
								EndLine:   291,
								StartPos:  5446,
								EndPos:    5447,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 292,
					EndLine:   292,
					StartPos:  5452,
					EndPos:    5459,
				},
				Expr: &binary.Mul{
					Position: &position.Position{
						StartLine: 292,
						EndLine:   292,
						StartPos:  5452,
						EndPos:    5458,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 292,
							EndLine:   292,
							StartPos:  5452,
							EndPos:    5453,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 292,
								EndLine:   292,
								StartPos:  5452,
								EndPos:    5453,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 292,
							EndLine:   292,
							StartPos:  5457,
							EndPos:    5458,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 292,
								EndLine:   292,
								StartPos:  5457,
								EndPos:    5458,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 293,
					EndLine:   293,
					StartPos:  5463,
					EndPos:    5471,
				},
				Expr: &binary.NotEqual{
					Position: &position.Position{
						StartLine: 293,
						EndLine:   293,
						StartPos:  5463,
						EndPos:    5470,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 293,
							EndLine:   293,
							StartPos:  5463,
							EndPos:    5464,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 293,
								EndLine:   293,
								StartPos:  5463,
								EndPos:    5464,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 293,
							EndLine:   293,
							StartPos:  5469,
							EndPos:    5470,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 293,
								EndLine:   293,
								StartPos:  5469,
								EndPos:    5470,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 294,
					EndLine:   294,
					StartPos:  5475,
					EndPos:    5484,
				},
				Expr: &binary.NotIdentical{
					Position: &position.Position{
						StartLine: 294,
						EndLine:   294,
						StartPos:  5475,
						EndPos:    5483,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 294,
							EndLine:   294,
							StartPos:  5475,
							EndPos:    5476,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  5475,
								EndPos:    5476,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 294,
							EndLine:   294,
							StartPos:  5482,
							EndPos:    5483,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  5482,
								EndPos:    5483,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 295,
					EndLine:   295,
					StartPos:  5488,
					EndPos:    5495,
				},
				Expr: &binary.Plus{
					Position: &position.Position{
						StartLine: 295,
						EndLine:   295,
						StartPos:  5488,
						EndPos:    5494,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 295,
							EndLine:   295,
							StartPos:  5488,
							EndPos:    5489,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 295,
								EndLine:   295,
								StartPos:  5488,
								EndPos:    5489,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 295,
							EndLine:   295,
							StartPos:  5493,
							EndPos:    5494,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 295,
								EndLine:   295,
								StartPos:  5493,
								EndPos:    5494,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 296,
					EndLine:   296,
					StartPos:  5499,
					EndPos:    5507,
				},
				Expr: &binary.Pow{
					Position: &position.Position{
						StartLine: 296,
						EndLine:   296,
						StartPos:  5499,
						EndPos:    5506,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 296,
							EndLine:   296,
							StartPos:  5499,
							EndPos:    5500,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  5499,
								EndPos:    5500,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 296,
							EndLine:   296,
							StartPos:  5505,
							EndPos:    5506,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  5505,
								EndPos:    5506,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 297,
					EndLine:   297,
					StartPos:  5511,
					EndPos:    5519,
				},
				Expr: &binary.ShiftLeft{
					Position: &position.Position{
						StartLine: 297,
						EndLine:   297,
						StartPos:  5511,
						EndPos:    5518,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 297,
							EndLine:   297,
							StartPos:  5511,
							EndPos:    5512,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 297,
								EndLine:   297,
								StartPos:  5511,
								EndPos:    5512,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 297,
							EndLine:   297,
							StartPos:  5517,
							EndPos:    5518,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 297,
								EndLine:   297,
								StartPos:  5517,
								EndPos:    5518,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 298,
					EndLine:   298,
					StartPos:  5523,
					EndPos:    5531,
				},
				Expr: &binary.ShiftRight{
					Position: &position.Position{
						StartLine: 298,
						EndLine:   298,
						StartPos:  5523,
						EndPos:    5530,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 298,
							EndLine:   298,
							StartPos:  5523,
							EndPos:    5524,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  5523,
								EndPos:    5524,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 298,
							EndLine:   298,
							StartPos:  5529,
							EndPos:    5530,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  5529,
								EndPos:    5530,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 299,
					EndLine:   299,
					StartPos:  5535,
					EndPos:    5543,
				},
				Expr: &binary.SmallerOrEqual{
					Position: &position.Position{
						StartLine: 299,
						EndLine:   299,
						StartPos:  5535,
						EndPos:    5542,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 299,
							EndLine:   299,
							StartPos:  5535,
							EndPos:    5536,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5535,
								EndPos:    5536,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 299,
							EndLine:   299,
							StartPos:  5541,
							EndPos:    5542,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5541,
								EndPos:    5542,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 300,
					EndLine:   300,
					StartPos:  5547,
					EndPos:    5554,
				},
				Expr: &binary.Smaller{
					Position: &position.Position{
						StartLine: 300,
						EndLine:   300,
						StartPos:  5547,
						EndPos:    5553,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 300,
							EndLine:   300,
							StartPos:  5547,
							EndPos:    5548,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5547,
								EndPos:    5548,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 300,
							EndLine:   300,
							StartPos:  5552,
							EndPos:    5553,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5552,
								EndPos:    5553,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 301,
					EndLine:   301,
					StartPos:  5558,
					EndPos:    5567,
				},
				Expr: &binary.Spaceship{
					Position: &position.Position{
						StartLine: 301,
						EndLine:   301,
						StartPos:  5558,
						EndPos:    5566,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 301,
							EndLine:   301,
							StartPos:  5558,
							EndPos:    5559,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 301,
								EndLine:   301,
								StartPos:  5558,
								EndPos:    5559,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 301,
							EndLine:   301,
							StartPos:  5565,
							EndPos:    5566,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 301,
								EndLine:   301,
								StartPos:  5565,
								EndPos:    5566,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 303,
					EndLine:   303,
					StartPos:  5572,
					EndPos:    5580,
				},
				Expr: &assign.Reference{
					Position: &position.Position{
						StartLine: 303,
						EndLine:   303,
						StartPos:  5572,
						EndPos:    5579,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 303,
							EndLine:   303,
							StartPos:  5572,
							EndPos:    5573,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 303,
								EndLine:   303,
								StartPos:  5572,
								EndPos:    5573,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 303,
							EndLine:   303,
							StartPos:  5578,
							EndPos:    5579,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 303,
								EndLine:   303,
								StartPos:  5578,
								EndPos:    5579,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 304,
					EndLine:   304,
					StartPos:  5584,
					EndPos:    5591,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 304,
						EndLine:   304,
						StartPos:  5584,
						EndPos:    5590,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 304,
							EndLine:   304,
							StartPos:  5584,
							EndPos:    5585,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 304,
								EndLine:   304,
								StartPos:  5584,
								EndPos:    5585,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 304,
							EndLine:   304,
							StartPos:  5589,
							EndPos:    5590,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 304,
								EndLine:   304,
								StartPos:  5589,
								EndPos:    5590,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 305,
					EndLine:   305,
					StartPos:  5595,
					EndPos:    5603,
				},
				Expr: &assign.BitwiseAnd{
					Position: &position.Position{
						StartLine: 305,
						EndLine:   305,
						StartPos:  5595,
						EndPos:    5602,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 305,
							EndLine:   305,
							StartPos:  5595,
							EndPos:    5596,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 305,
								EndLine:   305,
								StartPos:  5595,
								EndPos:    5596,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 305,
							EndLine:   305,
							StartPos:  5601,
							EndPos:    5602,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 305,
								EndLine:   305,
								StartPos:  5601,
								EndPos:    5602,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 306,
					EndLine:   306,
					StartPos:  5607,
					EndPos:    5615,
				},
				Expr: &assign.BitwiseOr{
					Position: &position.Position{
						StartLine: 306,
						EndLine:   306,
						StartPos:  5607,
						EndPos:    5614,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 306,
							EndLine:   306,
							StartPos:  5607,
							EndPos:    5608,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 306,
								EndLine:   306,
								StartPos:  5607,
								EndPos:    5608,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 306,
							EndLine:   306,
							StartPos:  5613,
							EndPos:    5614,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 306,
								EndLine:   306,
								StartPos:  5613,
								EndPos:    5614,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 307,
					EndLine:   307,
					StartPos:  5619,
					EndPos:    5627,
				},
				Expr: &assign.BitwiseXor{
					Position: &position.Position{
						StartLine: 307,
						EndLine:   307,
						StartPos:  5619,
						EndPos:    5626,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 307,
							EndLine:   307,
							StartPos:  5619,
							EndPos:    5620,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 307,
								EndLine:   307,
								StartPos:  5619,
								EndPos:    5620,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 307,
							EndLine:   307,
							StartPos:  5625,
							EndPos:    5626,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 307,
								EndLine:   307,
								StartPos:  5625,
								EndPos:    5626,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 308,
					EndLine:   308,
					StartPos:  5631,
					EndPos:    5639,
				},
				Expr: &assign.Concat{
					Position: &position.Position{
						StartLine: 308,
						EndLine:   308,
						StartPos:  5631,
						EndPos:    5638,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 308,
							EndLine:   308,
							StartPos:  5631,
							EndPos:    5632,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 308,
								EndLine:   308,
								StartPos:  5631,
								EndPos:    5632,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 308,
							EndLine:   308,
							StartPos:  5637,
							EndPos:    5638,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 308,
								EndLine:   308,
								StartPos:  5637,
								EndPos:    5638,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 309,
					EndLine:   309,
					StartPos:  5643,
					EndPos:    5651,
				},
				Expr: &assign.Div{
					Position: &position.Position{
						StartLine: 309,
						EndLine:   309,
						StartPos:  5643,
						EndPos:    5650,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 309,
							EndLine:   309,
							StartPos:  5643,
							EndPos:    5644,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 309,
								EndLine:   309,
								StartPos:  5643,
								EndPos:    5644,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 309,
							EndLine:   309,
							StartPos:  5649,
							EndPos:    5650,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 309,
								EndLine:   309,
								StartPos:  5649,
								EndPos:    5650,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 310,
					EndLine:   310,
					StartPos:  5655,
					EndPos:    5663,
				},
				Expr: &assign.Minus{
					Position: &position.Position{
						StartLine: 310,
						EndLine:   310,
						StartPos:  5655,
						EndPos:    5662,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 310,
							EndLine:   310,
							StartPos:  5655,
							EndPos:    5656,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 310,
								EndLine:   310,
								StartPos:  5655,
								EndPos:    5656,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 310,
							EndLine:   310,
							StartPos:  5661,
							EndPos:    5662,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 310,
								EndLine:   310,
								StartPos:  5661,
								EndPos:    5662,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 311,
					EndLine:   311,
					StartPos:  5667,
					EndPos:    5675,
				},
				Expr: &assign.Mod{
					Position: &position.Position{
						StartLine: 311,
						EndLine:   311,
						StartPos:  5667,
						EndPos:    5674,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 311,
							EndLine:   311,
							StartPos:  5667,
							EndPos:    5668,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 311,
								EndLine:   311,
								StartPos:  5667,
								EndPos:    5668,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 311,
							EndLine:   311,
							StartPos:  5673,
							EndPos:    5674,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 311,
								EndLine:   311,
								StartPos:  5673,
								EndPos:    5674,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 312,
					EndLine:   312,
					StartPos:  5679,
					EndPos:    5687,
				},
				Expr: &assign.Mul{
					Position: &position.Position{
						StartLine: 312,
						EndLine:   312,
						StartPos:  5679,
						EndPos:    5686,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 312,
							EndLine:   312,
							StartPos:  5679,
							EndPos:    5680,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 312,
								EndLine:   312,
								StartPos:  5679,
								EndPos:    5680,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 312,
							EndLine:   312,
							StartPos:  5685,
							EndPos:    5686,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 312,
								EndLine:   312,
								StartPos:  5685,
								EndPos:    5686,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 313,
					EndLine:   313,
					StartPos:  5691,
					EndPos:    5699,
				},
				Expr: &assign.Plus{
					Position: &position.Position{
						StartLine: 313,
						EndLine:   313,
						StartPos:  5691,
						EndPos:    5698,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 313,
							EndLine:   313,
							StartPos:  5691,
							EndPos:    5692,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5691,
								EndPos:    5692,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 313,
							EndLine:   313,
							StartPos:  5697,
							EndPos:    5698,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5697,
								EndPos:    5698,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 314,
					EndLine:   314,
					StartPos:  5703,
					EndPos:    5712,
				},
				Expr: &assign.Pow{
					Position: &position.Position{
						StartLine: 314,
						EndLine:   314,
						StartPos:  5703,
						EndPos:    5711,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 314,
							EndLine:   314,
							StartPos:  5703,
							EndPos:    5704,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 314,
								EndLine:   314,
								StartPos:  5703,
								EndPos:    5704,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 314,
							EndLine:   314,
							StartPos:  5710,
							EndPos:    5711,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 314,
								EndLine:   314,
								StartPos:  5710,
								EndPos:    5711,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 315,
					EndLine:   315,
					StartPos:  5716,
					EndPos:    5725,
				},
				Expr: &assign.ShiftLeft{
					Position: &position.Position{
						StartLine: 315,
						EndLine:   315,
						StartPos:  5716,
						EndPos:    5724,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 315,
							EndLine:   315,
							StartPos:  5716,
							EndPos:    5717,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5716,
								EndPos:    5717,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 315,
							EndLine:   315,
							StartPos:  5723,
							EndPos:    5724,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5723,
								EndPos:    5724,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 316,
					EndLine:   316,
					StartPos:  5729,
					EndPos:    5738,
				},
				Expr: &assign.ShiftRight{
					Position: &position.Position{
						StartLine: 316,
						EndLine:   316,
						StartPos:  5729,
						EndPos:    5737,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 316,
							EndLine:   316,
							StartPos:  5729,
							EndPos:    5730,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 316,
								EndLine:   316,
								StartPos:  5729,
								EndPos:    5730,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 316,
							EndLine:   316,
							StartPos:  5736,
							EndPos:    5737,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 316,
								EndLine:   316,
								StartPos:  5736,
								EndPos:    5737,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 318,
					EndLine:   318,
					StartPos:  5743,
					EndPos:    5781,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 318,
						EndLine:   318,
						StartPos:  5749,
						EndPos:    5751,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 318,
							EndLine:   318,
							StartPos:  5754,
							EndPos:    5779,
						},
						PhpDocComment: "",
						ReturnsRef:    false,
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 318,
								EndLine:   318,
								StartPos:  5770,
								EndPos:    5774,
							},
							Value: "class",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 318,
									EndLine:   318,
									StartPos:  5754,
									EndPos:    5759,
								},
								Value: "public",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 318,
								EndLine:   318,
								StartPos:  5778,
								EndPos:    5779,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 319,
					EndLine:   319,
					StartPos:  5785,
					EndPos:    5795,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 319,
						EndLine:   319,
						StartPos:  5785,
						EndPos:    5794,
					},
					Function: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 319,
							EndLine:   319,
							StartPos:  5785,
							EndPos:    5792,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 319,
									EndLine:   319,
									StartPos:  5786,
									EndPos:    5788,
								},
								Value: "foo",
							},
							&name.NamePart{
								Position: &position.Position{
									StartLine: 319,
									EndLine:   319,
									StartPos:  5790,
									EndPos:    5792,
								},
								Value: "bar",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 319,
							EndLine:   319,
							StartPos:  5793,
							EndPos:    5794,
						},
					},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 321,
					EndLine:   327,
					StartPos:  5800,
					EndPos:    5926,
				},
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 321,
						EndLine:   321,
						StartPos:  5809,
						EndPos:    5811,
					},
					Value: "foo",
				},
				Params: []node.Node{
					&node.Parameter{
						Position: &position.Position{
							StartLine: 321,
							EndLine:   321,
							StartPos:  5813,
							EndPos:    5815,
						},
						ByRef:    true,
						Variadic: false,
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 321,
								EndLine:   321,
								StartPos:  5814,
								EndPos:    5815,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 321,
									EndLine:   321,
									StartPos:  5814,
									EndPos:    5815,
								},
								Value: "a",
							},
						},
					},
					&node.Parameter{
						Position: &position.Position{
							StartLine: 321,
							EndLine:   321,
							StartPos:  5818,
							EndPos:    5822,
						},
						ByRef:    false,
						Variadic: true,
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 321,
								EndLine:   321,
								StartPos:  5821,
								EndPos:    5822,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 321,
									EndLine:   321,
									StartPos:  5821,
									EndPos:    5822,
								},
								Value: "b",
							},
						},
					},
				},
				Stmts: []node.Node{
					&stmt.HaltCompiler{
						Position: &position.Position{
							StartLine: 322,
							EndLine:   322,
							StartPos:  5830,
							EndPos:    5847,
						},
					},
					&stmt.Function{
						Position: &position.Position{
							StartLine: 323,
							EndLine:   323,
							StartPos:  5852,
							EndPos:    5868,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						FunctionName: &node.Identifier{
							Position: &position.Position{
								StartLine: 323,
								EndLine:   323,
								StartPos:  5861,
								EndPos:    5863,
							},
							Value: "bar",
						},
						Stmts: []node.Node{},
					},
					&stmt.Class{
						Position: &position.Position{
							StartLine: 324,
							EndLine:   324,
							StartPos:  5873,
							EndPos:    5884,
						},
						PhpDocComment: "",
						ClassName: &node.Identifier{
							Position: &position.Position{
								StartLine: 324,
								EndLine:   324,
								StartPos:  5879,
								EndPos:    5881,
							},
							Value: "Baz",
						},
						Stmts: []node.Node{},
					},
					&stmt.Trait{
						Position: &position.Position{
							StartLine: 325,
							EndLine:   325,
							StartPos:  5889,
							EndPos:    5900,
						},
						PhpDocComment: "",
						TraitName: &node.Identifier{
							Position: &position.Position{
								StartLine: 325,
								EndLine:   325,
								StartPos:  5895,
								EndPos:    5898,
							},
							Value: "Quux",
						},
						Stmts: []node.Node{},
					},
					&stmt.Interface{
						Position: &position.Position{
							StartLine: 326,
							EndLine:   326,
							StartPos:  5905,
							EndPos:    5922,
						},
						PhpDocComment: "",
						InterfaceName: &node.Identifier{
							Position: &position.Position{
								StartLine: 326,
								EndLine:   326,
								StartPos:  5915,
								EndPos:    5919,
							},
							Value: "Quuux",
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 329,
					EndLine:   329,
					StartPos:  5933,
					EndPos:    5975,
				},
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 329,
						EndLine:   329,
						StartPos:  5942,
						EndPos:    5944,
					},
					Value: "foo",
				},
				Params: []node.Node{
					&node.Parameter{
						Position: &position.Position{
							StartLine: 329,
							EndLine:   329,
							StartPos:  5946,
							EndPos:    5952,
						},
						ByRef:    true,
						Variadic: false,
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5947,
								EndPos:    5948,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5947,
									EndPos:    5948,
								},
								Value: "a",
							},
						},
						DefaultValue: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5952,
								EndPos:    5952,
							},
							Value: "1",
						},
					},
					&node.Parameter{
						Position: &position.Position{
							StartLine: 329,
							EndLine:   329,
							StartPos:  5955,
							EndPos:    5963,
						},
						ByRef:    false,
						Variadic: true,
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5958,
								EndPos:    5959,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5958,
									EndPos:    5959,
								},
								Value: "b",
							},
						},
						DefaultValue: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5963,
								EndPos:    5963,
							},
							Value: "1",
						},
					},
					&node.Parameter{
						Position: &position.Position{
							StartLine: 329,
							EndLine:   329,
							StartPos:  5966,
							EndPos:    5971,
						},
						ByRef:    false,
						Variadic: false,
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5966,
								EndPos:    5967,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5966,
									EndPos:    5967,
								},
								Value: "c",
							},
						},
						DefaultValue: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5971,
								EndPos:    5971,
							},
							Value: "1",
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 330,
					EndLine:   330,
					StartPos:  5979,
					EndPos:    6016,
				},
				PhpDocComment: "",
				ReturnsRef:    false,
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 330,
						EndLine:   330,
						StartPos:  5988,
						EndPos:    5990,
					},
					Value: "foo",
				},
				Params: []node.Node{
					&node.Parameter{
						Position: &position.Position{
							StartLine: 330,
							EndLine:   330,
							StartPos:  5992,
							EndPos:    5999,
						},
						ByRef:    false,
						Variadic: false,
						VariableType: &node.Identifier{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  5992,
								EndPos:    5996,
							},
							Value: "array",
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  5998,
								EndPos:    5999,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  5998,
									EndPos:    5999,
								},
								Value: "a",
							},
						},
					},
					&node.Parameter{
						Position: &position.Position{
							StartLine: 330,
							EndLine:   330,
							StartPos:  6002,
							EndPos:    6012,
						},
						ByRef:    false,
						Variadic: false,
						VariableType: &node.Identifier{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  6002,
								EndPos:    6009,
							},
							Value: "callable",
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  6011,
								EndPos:    6012,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  6011,
									EndPos:    6012,
								},
								Value: "b",
							},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 331,
					EndLine:   331,
					StartPos:  6020,
					EndPos:    6121,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 331,
						EndLine:   331,
						StartPos:  6041,
						EndPos:    6043,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 331,
							EndLine:   331,
							StartPos:  6020,
							EndPos:    6027,
						},
						Value: "abstract",
					},
					&node.Identifier{
						Position: &position.Position{
							StartLine: 331,
							EndLine:   331,
							StartPos:  6029,
							EndPos:    6033,
						},
						Value: "final",
					},
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 331,
							EndLine:   331,
							StartPos:  6047,
							EndPos:    6087,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 331,
								EndLine:   331,
								StartPos:  6082,
								EndPos:    6084,
							},
							Value: "bar",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 331,
									EndLine:   331,
									StartPos:  6047,
									EndPos:    6054,
								},
								Value: "abstract",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 331,
									EndLine:   331,
									StartPos:  6056,
									EndPos:    6064,
								},
								Value: "protected",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 331,
									EndLine:   331,
									StartPos:  6066,
									EndPos:    6071,
								},
								Value: "static",
							},
						},
						Stmt: &stmt.Nop{
							Position: &position.Position{
								StartLine: 331,
								EndLine:   331,
								StartPos:  6087,
								EndPos:    6087,
							},
						},
					},
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 331,
							EndLine:   331,
							StartPos:  6089,
							EndPos:    6119,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 331,
								EndLine:   331,
								StartPos:  6112,
								EndPos:    6114,
							},
							Value: "baz",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 331,
									EndLine:   331,
									StartPos:  6089,
									EndPos:    6093,
								},
								Value: "final",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 331,
									EndLine:   331,
									StartPos:  6095,
									EndPos:    6101,
								},
								Value: "private",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 331,
								EndLine:   331,
								StartPos:  6118,
								EndPos:    6119,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 333,
					EndLine:   333,
					StartPos:  6127,
					EndPos:    6140,
				},
				Expr: &expr.PropertyFetch{
					Position: &position.Position{
						StartLine: 333,
						EndLine:   333,
						StartPos:  6127,
						EndPos:    6139,
					},
					Variable: &expr.New{
						Position: &position.Position{
							StartLine: 333,
							EndLine:   333,
							StartPos:  6127,
							EndPos:    6133,
						},
						Class: &name.Name{
							Position: &position.Position{
								StartLine: 333,
								EndLine:   333,
								StartPos:  6131,
								EndPos:    6133,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 333,
										EndLine:   333,
										StartPos:  6131,
										EndPos:    6133,
									},
									Value: "Foo",
								},
							},
						},
					},
					Property: &node.Identifier{
						Position: &position.Position{
							StartLine: 333,
							EndLine:   333,
							StartPos:  6137,
							EndPos:    6139,
						},
						Value: "bar",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 334,
					EndLine:   334,
					StartPos:  6145,
					EndPos:    6155,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 334,
						EndLine:   334,
						StartPos:  6145,
						EndPos:    6154,
					},
					Function: &expr.New{
						Position: &position.Position{
							StartLine: 334,
							EndLine:   334,
							StartPos:  6145,
							EndPos:    6151,
						},
						Class: &name.Name{
							Position: &position.Position{
								StartLine: 334,
								EndLine:   334,
								StartPos:  6149,
								EndPos:    6151,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 334,
										EndLine:   334,
										StartPos:  6149,
										EndPos:    6151,
									},
									Value: "Foo",
								},
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 334,
							EndLine:   334,
							StartPos:  6153,
							EndPos:    6154,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 335,
					EndLine:   335,
					StartPos:  6159,
					EndPos:    6170,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 335,
						EndLine:   335,
						StartPos:  6159,
						EndPos:    6169,
					},
					Function: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 335,
							EndLine:   335,
							StartPos:  6159,
							EndPos:    6167,
						},
						Variable: &expr.ShortArray{
							Position: &position.Position{
								StartLine: 335,
								EndLine:   335,
								StartPos:  6159,
								EndPos:    6164,
							},
							Items: []node.Node{
								&expr.ArrayItem{
									Position: &position.Position{
										StartLine: 335,
										EndLine:   335,
										StartPos:  6160,
										EndPos:    6163,
									},
									Val: &expr.Variable{
										Position: &position.Position{
											StartLine: 335,
											EndLine:   335,
											StartPos:  6160,
											EndPos:    6163,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 335,
												EndLine:   335,
												StartPos:  6160,
												EndPos:    6163,
											},
											Value: "foo",
										},
									},
								},
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 335,
								EndLine:   335,
								StartPos:  6166,
								EndPos:    6166,
							},
							Value: "0",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 335,
							EndLine:   335,
							StartPos:  6168,
							EndPos:    6169,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 336,
					EndLine:   336,
					StartPos:  6174,
					EndPos:    6182,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 336,
						EndLine:   336,
						StartPos:  6174,
						EndPos:    6181,
					},
					Function: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 336,
							EndLine:   336,
							StartPos:  6174,
							EndPos:    6179,
						},
						Variable: &expr.ConstFetch{
							Position: &position.Position{
								StartLine: 336,
								EndLine:   336,
								StartPos:  6174,
								EndPos:    6176,
							},
							Constant: &name.Name{
								Position: &position.Position{
									StartLine: 336,
									EndLine:   336,
									StartPos:  6174,
									EndPos:    6176,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 336,
											EndLine:   336,
											StartPos:  6174,
											EndPos:    6176,
										},
										Value: "foo",
									},
								},
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 336,
								EndLine:   336,
								StartPos:  6178,
								EndPos:    6178,
							},
							Value: "1",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 336,
							EndLine:   336,
							StartPos:  6180,
							EndPos:    6181,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 337,
					EndLine:   337,
					StartPos:  6186,
					EndPos:    6193,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 337,
						EndLine:   337,
						StartPos:  6186,
						EndPos:    6192,
					},
					Function: &scalar.String{
						Position: &position.Position{
							StartLine: 337,
							EndLine:   337,
							StartPos:  6186,
							EndPos:    6190,
						},
						Value: "\"foo\"",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 337,
							EndLine:   337,
							StartPos:  6191,
							EndPos:    6192,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 338,
					EndLine:   338,
					StartPos:  6197,
					EndPos:    6208,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 338,
						EndLine:   338,
						StartPos:  6197,
						EndPos:    6207,
					},
					Function: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 338,
							EndLine:   338,
							StartPos:  6197,
							EndPos:    6205,
						},
						Variable: &expr.ShortArray{
							Position: &position.Position{
								StartLine: 338,
								EndLine:   338,
								StartPos:  6197,
								EndPos:    6199,
							},
							Items: []node.Node{
								&expr.ArrayItem{
									Position: &position.Position{
										StartLine: 338,
										EndLine:   338,
										StartPos:  6198,
										EndPos:    6198,
									},
									Val: &scalar.Lnumber{
										Position: &position.Position{
											StartLine: 338,
											EndLine:   338,
											StartPos:  6198,
											EndPos:    6198,
										},
										Value: "1",
									},
								},
							},
						},
						Dim: &expr.Variable{
							Position: &position.Position{
								StartLine: 338,
								EndLine:   338,
								StartPos:  6201,
								EndPos:    6204,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 338,
									EndLine:   338,
									StartPos:  6201,
									EndPos:    6204,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 338,
							EndLine:   338,
							StartPos:  6206,
							EndPos:    6207,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 339,
					EndLine:   339,
					StartPos:  6212,
					EndPos:    6220,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 339,
						EndLine:   339,
						StartPos:  6212,
						EndPos:    6219,
					},
					VarName: &expr.FunctionCall{
						Position: &position.Position{
							StartLine: 339,
							EndLine:   339,
							StartPos:  6214,
							EndPos:    6218,
						},
						Function: &name.Name{
							Position: &position.Position{
								StartLine: 339,
								EndLine:   339,
								StartPos:  6214,
								EndPos:    6216,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 339,
										EndLine:   339,
										StartPos:  6214,
										EndPos:    6216,
									},
									Value: "foo",
								},
							},
						},
						ArgumentList: &node.ArgumentList{
							Position: &position.Position{
								StartLine: 339,
								EndLine:   339,
								StartPos:  6217,
								EndPos:    6218,
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 341,
					EndLine:   341,
					StartPos:  6225,
					EndPos:    6236,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 341,
						EndLine:   341,
						StartPos:  6225,
						EndPos:    6235,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 341,
							EndLine:   341,
							StartPos:  6225,
							EndPos:    6227,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  6225,
									EndPos:    6227,
								},
								Value: "Foo",
							},
						},
					},
					Call: &expr.Variable{
						Position: &position.Position{
							StartLine: 341,
							EndLine:   341,
							StartPos:  6230,
							EndPos:    6233,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  6230,
								EndPos:    6233,
							},
							Value: "bar",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 341,
							EndLine:   341,
							StartPos:  6234,
							EndPos:    6235,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 342,
					EndLine:   342,
					StartPos:  6240,
					EndPos:    6256,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 342,
						EndLine:   342,
						StartPos:  6240,
						EndPos:    6255,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 342,
							EndLine:   342,
							StartPos:  6240,
							EndPos:    6242,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 342,
									EndLine:   342,
									StartPos:  6240,
									EndPos:    6242,
								},
								Value: "Foo",
							},
						},
					},
					Call: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 342,
							EndLine:   342,
							StartPos:  6246,
							EndPos:    6252,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 342,
								EndLine:   342,
								StartPos:  6246,
								EndPos:    6249,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 342,
									EndLine:   342,
									StartPos:  6246,
									EndPos:    6249,
								},
								Value: "bar",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 342,
								EndLine:   342,
								StartPos:  6251,
								EndPos:    6251,
							},
							Value: "0",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 342,
							EndLine:   342,
							StartPos:  6254,
							EndPos:    6255,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 344,
					EndLine:   344,
					StartPos:  6263,
					EndPos:    6273,
				},
				Expr: &expr.PropertyFetch{
					Position: &position.Position{
						StartLine: 344,
						EndLine:   344,
						StartPos:  6263,
						EndPos:    6272,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 344,
							EndLine:   344,
							StartPos:  6263,
							EndPos:    6266,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 344,
								EndLine:   344,
								StartPos:  6263,
								EndPos:    6266,
							},
							Value: "foo",
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 344,
							EndLine:   344,
							StartPos:  6269,
							EndPos:    6272,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 344,
								EndLine:   344,
								StartPos:  6269,
								EndPos:    6272,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 345,
					EndLine:   345,
					StartPos:  6277,
					EndPos:    6292,
				},
				Expr: &expr.PropertyFetch{
					Position: &position.Position{
						StartLine: 345,
						EndLine:   345,
						StartPos:  6277,
						EndPos:    6290,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 345,
							EndLine:   345,
							StartPos:  6277,
							EndPos:    6280,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 345,
								EndLine:   345,
								StartPos:  6277,
								EndPos:    6280,
							},
							Value: "foo",
						},
					},
					Property: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 345,
							EndLine:   345,
							StartPos:  6284,
							EndPos:    6290,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 345,
								EndLine:   345,
								StartPos:  6284,
								EndPos:    6287,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 345,
									EndLine:   345,
									StartPos:  6284,
									EndPos:    6287,
								},
								Value: "bar",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 345,
								EndLine:   345,
								StartPos:  6289,
								EndPos:    6289,
							},
							Value: "0",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 347,
					EndLine:   347,
					StartPos:  6297,
					EndPos:    6318,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 347,
						EndLine:   347,
						StartPos:  6297,
						EndPos:    6317,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 347,
								EndLine:   347,
								StartPos:  6298,
								EndPos:    6303,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 347,
									EndLine:   347,
									StartPos:  6298,
									EndPos:    6298,
								},
								Value: "1",
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 347,
									EndLine:   347,
									StartPos:  6301,
									EndPos:    6303,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 347,
										EndLine:   347,
										StartPos:  6302,
										EndPos:    6303,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 347,
											EndLine:   347,
											StartPos:  6302,
											EndPos:    6303,
										},
										Value: "a",
									},
								},
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 347,
								EndLine:   347,
								StartPos:  6306,
								EndPos:    6316,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 347,
									EndLine:   347,
									StartPos:  6306,
									EndPos:    6306,
								},
								Value: "2",
							},
							Val: &expr.List{
								Position: &position.Position{
									StartLine: 347,
									EndLine:   347,
									StartPos:  6309,
									EndPos:    6316,
								},
								Items: []node.Node{
									&expr.ArrayItem{
										Position: &position.Position{
											StartLine: 347,
											EndLine:   347,
											StartPos:  6314,
											EndPos:    6315,
										},
										Val: &expr.Variable{
											Position: &position.Position{
												StartLine: 347,
												EndLine:   347,
												StartPos:  6314,
												EndPos:    6315,
											},
											VarName: &node.Identifier{
												Position: &position.Position{
													StartLine: 347,
													EndLine:   347,
													StartPos:  6314,
													EndPos:    6315,
												},
												Value: "b",
											},
										},
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
	assertEqual(t, expected, actual)
}

func TestPhp5Strings(t *testing.T) {
	src := `<?
		"test";
		"\$test";
		"
			test
		";
		'$test';
		'
			$test
		';
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   10,
			StartPos:  6,
			EndPos:    70,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  6,
					EndPos:    12,
				},
				Expr: &scalar.String{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  6,
						EndPos:    11,
					},
					Value: "\"test\"",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 3,
					EndLine:   3,
					StartPos:  16,
					EndPos:    24,
				},
				Expr: &scalar.String{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  16,
						EndPos:    23,
					},
					Value: "\"\\$test\"",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 4,
					EndLine:   6,
					StartPos:  28,
					EndPos:    41,
				},
				Expr: &scalar.String{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   6,
						StartPos:  28,
						EndPos:    40,
					},
					Value: "\"\n\t\t\ttest\n\t\t\"",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 7,
					EndLine:   7,
					StartPos:  45,
					EndPos:    52,
				},
				Expr: &scalar.String{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   7,
						StartPos:  45,
						EndPos:    51,
					},
					Value: "'$test'",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 8,
					EndLine:   10,
					StartPos:  56,
					EndPos:    70,
				},
				Expr: &scalar.String{
					Position: &position.Position{
						StartLine: 8,
						EndLine:   10,
						StartPos:  56,
						EndPos:    69,
					},
					Value: "'\n\t\t\t$test\n\t\t'",
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestPhp5Heredoc(t *testing.T) {
	src := `<?
		<<<CAD
CAD;
		<<<CAD
	hello
CAD;
		<<<"CAD"
	hello
CAD;
		<<<"CAD"
	hello $world
CAD;
		<<<'CAD'
	hello $world
CAD;
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   15,
			StartPos:  9,
			EndPos:    120,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   3,
					StartPos:  9,
					EndPos:    16,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   3,
						StartPos:  9,
						EndPos:    15,
					},
					Label: "CAD",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 4,
					EndLine:   6,
					StartPos:  23,
					EndPos:    37,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   6,
						StartPos:  23,
						EndPos:    36,
					},
					Label: "CAD",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  27,
								EndPos:    32,
							},
							Value: "\thello",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 7,
					EndLine:   9,
					StartPos:  44,
					EndPos:    60,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   9,
						StartPos:  44,
						EndPos:    59,
					},
					Label: "\"CAD\"",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 8,
								EndLine:   8,
								StartPos:  50,
								EndPos:    55,
							},
							Value: "\thello",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 10,
					EndLine:   12,
					StartPos:  67,
					EndPos:    90,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 10,
						EndLine:   12,
						StartPos:  67,
						EndPos:    89,
					},
					Label: "\"CAD\"",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 11,
								EndLine:   11,
								StartPos:  73,
								EndPos:    79,
							},
							Value: "\thello ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 11,
								EndLine:   11,
								StartPos:  80,
								EndPos:    85,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  80,
									EndPos:    85,
								},
								Value: "world",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 13,
					EndLine:   15,
					StartPos:  97,
					EndPos:    120,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 13,
						EndLine:   15,
						StartPos:  97,
						EndPos:    119,
					},
					Label: "'CAD'",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 14,
								EndLine:   14,
								StartPos:  103,
								EndPos:    115,
							},
							Value: "\thello $world",
						},
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
