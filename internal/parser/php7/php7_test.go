package php7_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/internal/graph"
	"github.com/z7zmey/php-parser/internal/parser/php7"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/node"
	"github.com/z7zmey/php-parser/pkg/node/expr"
	"github.com/z7zmey/php-parser/pkg/node/expr/assign"
	"github.com/z7zmey/php-parser/pkg/node/expr/binary"
	"github.com/z7zmey/php-parser/pkg/node/expr/cast"
	"github.com/z7zmey/php-parser/pkg/node/name"
	"github.com/z7zmey/php-parser/pkg/node/scalar"
	"github.com/z7zmey/php-parser/pkg/node/stmt"
	"github.com/z7zmey/php-parser/pkg/position"
)

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

		__halt_compiler();

		parsing process must be terminated
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   348,
			StartPos:  5,
			EndPos:    6319,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  5,
					EndPos:    20,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  5,
						EndPos:    19,
					},
					Function: &name.Name{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  5,
							EndPos:    8,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  5,
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
							StartPos:  8,
							EndPos:    19,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  9,
									EndPos:    11,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  9,
										EndPos:    11,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  9,
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
									StartPos:  13,
									EndPos:    18,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  16,
										EndPos:    18,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  16,
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
					StartPos:  23,
					EndPos:    39,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  23,
						EndPos:    38,
					},
					Function: &expr.Variable{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  23,
							EndPos:    27,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  23,
								EndPos:    27,
							},
							Value: "foo",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  27,
							EndPos:    38,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  28,
									EndPos:    30,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  28,
										EndPos:    30,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  28,
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
									StartPos:  32,
									EndPos:    37,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  35,
										EndPos:    37,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  35,
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
					StartPos:  42,
					EndPos:    63,
				},
				Expr: &expr.MethodCall{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   4,
						StartPos:  42,
						EndPos:    62,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  42,
							EndPos:    46,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  42,
								EndPos:    46,
							},
							Value: "foo",
						},
					},
					Method: &node.Identifier{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  48,
							EndPos:    51,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  51,
							EndPos:    62,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  52,
									EndPos:    54,
								},
								IsReference: false,
								Variadic:    false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  52,
										EndPos:    54,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  52,
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
									StartPos:  56,
									EndPos:    61,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  59,
										EndPos:    61,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  59,
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
					StartPos:  66,
					EndPos:    86,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 5,
						EndLine:   5,
						StartPos:  66,
						EndPos:    85,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  66,
							EndPos:    69,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  66,
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
							StartPos:  71,
							EndPos:    74,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  74,
							EndPos:    85,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  75,
									EndPos:    77,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  75,
										EndPos:    77,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  75,
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
									StartPos:  79,
									EndPos:    84,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  82,
										EndPos:    84,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  82,
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
					StartPos:  89,
					EndPos:    110,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 6,
						EndLine:   6,
						StartPos:  89,
						EndPos:    109,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  89,
							EndPos:    93,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 6,
								EndLine:   6,
								StartPos:  89,
								EndPos:    93,
							},
							Value: "foo",
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  95,
							EndPos:    98,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  98,
							EndPos:    109,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 6,
									EndLine:   6,
									StartPos:  99,
									EndPos:    101,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 6,
										EndLine:   6,
										StartPos:  99,
										EndPos:    101,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 6,
											EndLine:   6,
											StartPos:  99,
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
									StartPos:  103,
									EndPos:    108,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 6,
										EndLine:   6,
										StartPos:  106,
										EndPos:    108,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 6,
											EndLine:   6,
											StartPos:  106,
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
					StartPos:  113,
					EndPos:    132,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   7,
						StartPos:  113,
						EndPos:    131,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   7,
							StartPos:  117,
							EndPos:    120,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 7,
									EndLine:   7,
									StartPos:  117,
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
							StartPos:  120,
							EndPos:    131,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 7,
									EndLine:   7,
									StartPos:  121,
									EndPos:    123,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  121,
										EndPos:    123,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 7,
											EndLine:   7,
											StartPos:  121,
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
									StartPos:  125,
									EndPos:    130,
								},
								Variadic:    true,
								IsReference: false,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  128,
										EndPos:    130,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 7,
											EndLine:   7,
											StartPos:  128,
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
					StartPos:  160,
					EndPos:    185,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 9,
						EndLine:   9,
						StartPos:  160,
						EndPos:    184,
					},
					Class: &stmt.Class{
						Position: &position.Position{
							StartLine: 9,
							EndLine:   9,
							StartPos:  164,
							EndPos:    184,
						},
						PhpDocComment: "/** anonymous class */",
						ArgumentList: &node.ArgumentList{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  170,
								EndPos:    181,
							},
							Arguments: []node.Node{
								&node.Argument{
									Position: &position.Position{
										StartLine: 9,
										EndLine:   9,
										StartPos:  171,
										EndPos:    173,
									},
									Variadic:    false,
									IsReference: false,
									Expr: &expr.Variable{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  171,
											EndPos:    173,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  171,
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
										StartPos:  175,
										EndPos:    180,
									},
									Variadic:    true,
									IsReference: false,
									Expr: &expr.Variable{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  178,
											EndPos:    180,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  178,
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
					StartPos:  188,
					EndPos:    201,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 10,
						EndLine:   10,
						StartPos:  188,
						EndPos:    200,
					},
					Class: &stmt.Class{
						Position: &position.Position{
							StartLine: 10,
							EndLine:   10,
							StartPos:  192,
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
					StartPos:  204,
					EndPos:    213,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 11,
						EndLine:   11,
						StartPos:  204,
						EndPos:    212,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 11,
							EndLine:   11,
							StartPos:  208,
							EndPos:    212,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 11,
								EndLine:   11,
								StartPos:  208,
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
					StartPos:  216,
					EndPos:    228,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 12,
						EndLine:   12,
						StartPos:  216,
						EndPos:    227,
					},
					Class: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 12,
							EndLine:   12,
							StartPos:  220,
							EndPos:    227,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 12,
								EndLine:   12,
								StartPos:  220,
								EndPos:    224,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 12,
									EndLine:   12,
									StartPos:  220,
									EndPos:    224,
								},
								Value: "foo",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 12,
								EndLine:   12,
								StartPos:  225,
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
					StartPos:  231,
					EndPos:    246,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 13,
						EndLine:   13,
						StartPos:  231,
						EndPos:    245,
					},
					Class: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 13,
							EndLine:   13,
							StartPos:  235,
							EndPos:    245,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 13,
								EndLine:   13,
								StartPos:  235,
								EndPos:    239,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 13,
									EndLine:   13,
									StartPos:  235,
									EndPos:    239,
								},
								Value: "foo",
							},
						},
						Dim: &expr.Variable{
							Position: &position.Position{
								StartLine: 13,
								EndLine:   13,
								StartPos:  240,
								EndPos:    244,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 13,
									EndLine:   13,
									StartPos:  240,
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
					StartPos:  249,
					EndPos:    263,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 14,
						EndLine:   14,
						StartPos:  249,
						EndPos:    262,
					},
					Class: &expr.PropertyFetch{
						Position: &position.Position{
							StartLine: 14,
							EndLine:   14,
							StartPos:  253,
							EndPos:    262,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 14,
								EndLine:   14,
								StartPos:  253,
								EndPos:    257,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 14,
									EndLine:   14,
									StartPos:  253,
									EndPos:    257,
								},
								Value: "foo",
							},
						},
						Property: &node.Identifier{
							Position: &position.Position{
								StartLine: 14,
								EndLine:   14,
								StartPos:  259,
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
					StartPos:  266,
					EndPos:    281,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 15,
						EndLine:   15,
						StartPos:  266,
						EndPos:    280,
					},
					Class: &expr.StaticPropertyFetch{
						Position: &position.Position{
							StartLine: 15,
							EndLine:   15,
							StartPos:  270,
							EndPos:    280,
						},
						Class: &expr.Variable{
							Position: &position.Position{
								StartLine: 15,
								EndLine:   15,
								StartPos:  270,
								EndPos:    274,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 15,
									EndLine:   15,
									StartPos:  270,
									EndPos:    274,
								},
								Value: "foo",
							},
						},
						Property: &expr.Variable{
							Position: &position.Position{
								StartLine: 15,
								EndLine:   15,
								StartPos:  276,
								EndPos:    280,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 15,
									EndLine:   15,
									StartPos:  276,
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
					StartPos:  284,
					EndPos:    301,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 16,
						EndLine:   16,
						StartPos:  284,
						EndPos:    300,
					},
					Class: &expr.StaticPropertyFetch{
						Position: &position.Position{
							StartLine: 16,
							EndLine:   16,
							StartPos:  288,
							EndPos:    300,
						},
						Class: &node.Identifier{
							Position: &position.Position{
								StartLine: 16,
								EndLine:   16,
								StartPos:  288,
								EndPos:    294,
							},
							Value: "static",
						},
						Property: &expr.Variable{
							Position: &position.Position{
								StartLine: 16,
								EndLine:   16,
								StartPos:  296,
								EndPos:    300,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 16,
									EndLine:   16,
									StartPos:  296,
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
					StartPos:  305,
					EndPos:    350,
				},
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 18,
						EndLine:   18,
						StartPos:  314,
						EndPos:    317,
					},
					Value: "foo",
				},
				Params: []node.Node{
					&node.Parameter{
						Position: &position.Position{
							StartLine: 18,
							EndLine:   18,
							StartPos:  318,
							EndPos:    332,
						},
						ByRef:    false,
						Variadic: false,
						VariableType: &node.Nullable{
							Position: &position.Position{
								StartLine: 18,
								EndLine:   18,
								StartPos:  318,
								EndPos:    322,
							},
							Expr: &name.Name{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  319,
									EndPos:    322,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 18,
											EndLine:   18,
											StartPos:  319,
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
								StartPos:  323,
								EndPos:    327,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  323,
									EndPos:    327,
								},
								Value: "bar",
							},
						},
						DefaultValue: &expr.ConstFetch{
							Position: &position.Position{
								StartLine: 18,
								EndLine:   18,
								StartPos:  328,
								EndPos:    332,
							},
							Constant: &name.Name{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  328,
									EndPos:    332,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 18,
											EndLine:   18,
											StartPos:  328,
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
							StartPos:  334,
							EndPos:    346,
						},
						Variadic: true,
						ByRef:    true,
						VariableType: &name.Name{
							Position: &position.Position{
								StartLine: 18,
								EndLine:   18,
								StartPos:  334,
								EndPos:    337,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 18,
										EndLine:   18,
										StartPos:  334,
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
								StartPos:  342,
								EndPos:    346,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  342,
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
					StartPos:  353,
					EndPos:    417,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 19,
						EndLine:   19,
						StartPos:  359,
						EndPos:    362,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 19,
							EndLine:   19,
							StartPos:  364,
							EndPos:    416,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 19,
								EndLine:   19,
								StartPos:  380,
								EndPos:    383,
							},
							Value: "foo",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 19,
									EndLine:   19,
									StartPos:  364,
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
									StartPos:  384,
									EndPos:    398,
								},
								ByRef:    false,
								Variadic: false,
								VariableType: &node.Nullable{
									Position: &position.Position{
										StartLine: 19,
										EndLine:   19,
										StartPos:  384,
										EndPos:    388,
									},
									Expr: &name.Name{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  385,
											EndPos:    388,
										},
										Parts: []node.Node{
											&name.NamePart{
												Position: &position.Position{
													StartLine: 19,
													EndLine:   19,
													StartPos:  385,
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
										StartPos:  389,
										EndPos:    393,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  389,
											EndPos:    393,
										},
										Value: "bar",
									},
								},
								DefaultValue: &expr.ConstFetch{
									Position: &position.Position{
										StartLine: 19,
										EndLine:   19,
										StartPos:  394,
										EndPos:    398,
									},
									Constant: &name.Name{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  394,
											EndPos:    398,
										},
										Parts: []node.Node{
											&name.NamePart{
												Position: &position.Position{
													StartLine: 19,
													EndLine:   19,
													StartPos:  394,
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
									StartPos:  400,
									EndPos:    412,
								},
								ByRef:    true,
								Variadic: true,
								VariableType: &name.Name{
									Position: &position.Position{
										StartLine: 19,
										EndLine:   19,
										StartPos:  400,
										EndPos:    403,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 19,
												EndLine:   19,
												StartPos:  400,
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
										StartPos:  408,
										EndPos:    412,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  408,
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
								StartPos:  414,
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
					StartPos:  420,
					EndPos:    462,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 20,
						EndLine:   20,
						StartPos:  420,
						EndPos:    461,
					},
					PhpDocComment: "",
					ReturnsRef:    false,
					Static:        false,
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 20,
								EndLine:   20,
								StartPos:  429,
								EndPos:    443,
							},
							ByRef:    false,
							Variadic: false,
							VariableType: &node.Nullable{
								Position: &position.Position{
									StartLine: 20,
									EndLine:   20,
									StartPos:  429,
									EndPos:    433,
								},
								Expr: &name.Name{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  430,
										EndPos:    433,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 20,
												EndLine:   20,
												StartPos:  430,
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
									StartPos:  434,
									EndPos:    438,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  434,
										EndPos:    438,
									},
									Value: "bar",
								},
							},
							DefaultValue: &expr.ConstFetch{
								Position: &position.Position{
									StartLine: 20,
									EndLine:   20,
									StartPos:  439,
									EndPos:    443,
								},
								Constant: &name.Name{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  439,
										EndPos:    443,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 20,
												EndLine:   20,
												StartPos:  439,
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
								StartPos:  445,
								EndPos:    457,
							},
							ByRef:    true,
							Variadic: true,
							VariableType: &name.Name{
								Position: &position.Position{
									StartLine: 20,
									EndLine:   20,
									StartPos:  445,
									EndPos:    448,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 20,
											EndLine:   20,
											StartPos:  445,
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
									StartPos:  453,
									EndPos:    457,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  453,
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
					StartPos:  465,
					EndPos:    514,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 21,
						EndLine:   21,
						StartPos:  465,
						EndPos:    513,
					},
					PhpDocComment: "",
					ReturnsRef:    false,
					Static:        true,
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 21,
								EndLine:   21,
								StartPos:  481,
								EndPos:    495,
							},
							ByRef:    false,
							Variadic: false,
							VariableType: &node.Nullable{
								Position: &position.Position{
									StartLine: 21,
									EndLine:   21,
									StartPos:  481,
									EndPos:    485,
								},
								Expr: &name.Name{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  482,
										EndPos:    485,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 21,
												EndLine:   21,
												StartPos:  482,
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
									StartPos:  486,
									EndPos:    490,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  486,
										EndPos:    490,
									},
									Value: "bar",
								},
							},
							DefaultValue: &expr.ConstFetch{
								Position: &position.Position{
									StartLine: 21,
									EndLine:   21,
									StartPos:  491,
									EndPos:    495,
								},
								Constant: &name.Name{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  491,
										EndPos:    495,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 21,
												EndLine:   21,
												StartPos:  491,
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
								StartPos:  497,
								EndPos:    509,
							},
							ByRef:    true,
							Variadic: true,
							VariableType: &name.Name{
								Position: &position.Position{
									StartLine: 21,
									EndLine:   21,
									StartPos:  497,
									EndPos:    500,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 21,
											EndLine:   21,
											StartPos:  497,
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
									StartPos:  505,
									EndPos:    509,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  505,
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
					StartPos:  518,
					EndPos:    538,
				},
				Expr: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 23,
						EndLine:   23,
						StartPos:  518,
						EndPos:    537,
					},
					Value: "1234567890123456789",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 24,
					EndLine:   24,
					StartPos:  541,
					EndPos:    562,
				},
				Expr: &scalar.Dnumber{
					Position: &position.Position{
						StartLine: 24,
						EndLine:   24,
						StartPos:  541,
						EndPos:    561,
					},
					Value: "12345678901234567890",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 25,
					EndLine:   25,
					StartPos:  565,
					EndPos:    568,
				},
				Expr: &scalar.Dnumber{
					Position: &position.Position{
						StartLine: 25,
						EndLine:   25,
						StartPos:  565,
						EndPos:    567,
					},
					Value: "0.",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 26,
					EndLine:   26,
					StartPos:  571,
					EndPos:    638,
				},
				Expr: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 26,
						EndLine:   26,
						StartPos:  571,
						EndPos:    637,
					},
					Value: "0b0111111111111111111111111111111111111111111111111111111111111111",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 27,
					EndLine:   27,
					StartPos:  641,
					EndPos:    708,
				},
				Expr: &scalar.Dnumber{
					Position: &position.Position{
						StartLine: 27,
						EndLine:   27,
						StartPos:  641,
						EndPos:    707,
					},
					Value: "0b1111111111111111111111111111111111111111111111111111111111111111",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 28,
					EndLine:   28,
					StartPos:  711,
					EndPos:    732,
				},
				Expr: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 28,
						EndLine:   28,
						StartPos:  711,
						EndPos:    731,
					},
					Value: "0x007111111111111111",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 29,
					EndLine:   29,
					StartPos:  735,
					EndPos:    754,
				},
				Expr: &scalar.Dnumber{
					Position: &position.Position{
						StartLine: 29,
						EndLine:   29,
						StartPos:  735,
						EndPos:    753,
					},
					Value: "0x8111111111111111",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 30,
					EndLine:   30,
					StartPos:  757,
					EndPos:    767,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 30,
						EndLine:   30,
						StartPos:  757,
						EndPos:    766,
					},
					Value: "__CLASS__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 31,
					EndLine:   31,
					StartPos:  770,
					EndPos:    778,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 31,
						EndLine:   31,
						StartPos:  770,
						EndPos:    777,
					},
					Value: "__DIR__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 32,
					EndLine:   32,
					StartPos:  781,
					EndPos:    790,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 32,
						EndLine:   32,
						StartPos:  781,
						EndPos:    789,
					},
					Value: "__FILE__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 33,
					EndLine:   33,
					StartPos:  793,
					EndPos:    806,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 33,
						EndLine:   33,
						StartPos:  793,
						EndPos:    805,
					},
					Value: "__FUNCTION__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 34,
					EndLine:   34,
					StartPos:  809,
					EndPos:    818,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 34,
						EndLine:   34,
						StartPos:  809,
						EndPos:    817,
					},
					Value: "__LINE__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 35,
					EndLine:   35,
					StartPos:  821,
					EndPos:    835,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 35,
						EndLine:   35,
						StartPos:  821,
						EndPos:    834,
					},
					Value: "__NAMESPACE__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 36,
					EndLine:   36,
					StartPos:  838,
					EndPos:    849,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 36,
						EndLine:   36,
						StartPos:  838,
						EndPos:    848,
					},
					Value: "__METHOD__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 37,
					EndLine:   37,
					StartPos:  852,
					EndPos:    862,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 37,
						EndLine:   37,
						StartPos:  852,
						EndPos:    861,
					},
					Value: "__TRAIT__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 39,
					EndLine:   39,
					StartPos:  866,
					EndPos:    878,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 39,
						EndLine:   39,
						StartPos:  866,
						EndPos:    877,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 39,
								EndLine:   39,
								StartPos:  867,
								EndPos:    872,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 39,
								EndLine:   39,
								StartPos:  872,
								EndPos:    876,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 39,
									EndLine:   39,
									StartPos:  872,
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
					StartPos:  881,
					EndPos:    896,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 40,
						EndLine:   40,
						StartPos:  881,
						EndPos:    895,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 40,
								EndLine:   40,
								StartPos:  882,
								EndPos:    887,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 40,
								EndLine:   40,
								StartPos:  887,
								EndPos:    894,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 40,
									EndLine:   40,
									StartPos:  887,
									EndPos:    891,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 40,
										EndLine:   40,
										StartPos:  887,
										EndPos:    891,
									},
									Value: "var",
								},
							},
							Dim: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 40,
									EndLine:   40,
									StartPos:  892,
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
					StartPos:  899,
					EndPos:    915,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 41,
						EndLine:   41,
						StartPos:  899,
						EndPos:    914,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 41,
								EndLine:   41,
								StartPos:  900,
								EndPos:    905,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 41,
								EndLine:   41,
								StartPos:  905,
								EndPos:    913,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 41,
									EndLine:   41,
									StartPos:  905,
									EndPos:    909,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 41,
										EndLine:   41,
										StartPos:  905,
										EndPos:    909,
									},
									Value: "var",
								},
							},
							Dim: &expr.UnaryMinus{
								Position: &position.Position{
									StartLine: 41,
									EndLine:   41,
									StartPos:  910,
									EndPos:    912,
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 41,
										EndLine:   41,
										StartPos:  910,
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
					StartPos:  918,
					EndPos:    972,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 42,
						EndLine:   42,
						StartPos:  918,
						EndPos:    971,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 42,
								EndLine:   42,
								StartPos:  919,
								EndPos:    924,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 42,
								EndLine:   42,
								StartPos:  924,
								EndPos:    970,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 42,
									EndLine:   42,
									StartPos:  924,
									EndPos:    928,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 42,
										EndLine:   42,
										StartPos:  924,
										EndPos:    928,
									},
									Value: "var",
								},
							},
							Dim: &scalar.String{
								Position: &position.Position{
									StartLine: 42,
									EndLine:   42,
									StartPos:  929,
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
					StartPos:  975,
					EndPos:    1030,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 43,
						EndLine:   43,
						StartPos:  975,
						EndPos:    1029,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 43,
								EndLine:   43,
								StartPos:  976,
								EndPos:    981,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 43,
								EndLine:   43,
								StartPos:  981,
								EndPos:    1028,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 43,
									EndLine:   43,
									StartPos:  981,
									EndPos:    985,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 43,
										EndLine:   43,
										StartPos:  981,
										EndPos:    985,
									},
									Value: "var",
								},
							},
							Dim: &scalar.String{
								Position: &position.Position{
									StartLine: 43,
									EndLine:   43,
									StartPos:  986,
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
					StartPos:  1033,
					EndPos:    1050,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 44,
						EndLine:   44,
						StartPos:  1033,
						EndPos:    1049,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 44,
								EndLine:   44,
								StartPos:  1034,
								EndPos:    1039,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 44,
								EndLine:   44,
								StartPos:  1039,
								EndPos:    1048,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 44,
									EndLine:   44,
									StartPos:  1039,
									EndPos:    1043,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 44,
										EndLine:   44,
										StartPos:  1039,
										EndPos:    1043,
									},
									Value: "var",
								},
							},
							Dim: &scalar.String{
								Position: &position.Position{
									StartLine: 44,
									EndLine:   44,
									StartPos:  1044,
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
					StartPos:  1053,
					EndPos:    1071,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 45,
						EndLine:   45,
						StartPos:  1053,
						EndPos:    1070,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 45,
								EndLine:   45,
								StartPos:  1054,
								EndPos:    1059,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 45,
								EndLine:   45,
								StartPos:  1059,
								EndPos:    1069,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 45,
									EndLine:   45,
									StartPos:  1059,
									EndPos:    1063,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 45,
										EndLine:   45,
										StartPos:  1059,
										EndPos:    1063,
									},
									Value: "var",
								},
							},
							Dim: &expr.Variable{
								Position: &position.Position{
									StartLine: 45,
									EndLine:   45,
									StartPos:  1064,
									EndPos:    1068,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 45,
										EndLine:   45,
										StartPos:  1064,
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
					StartPos:  1074,
					EndPos:    1086,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 46,
						EndLine:   46,
						StartPos:  1074,
						EndPos:    1085,
					},
					Parts: []node.Node{
						&expr.Variable{
							Position: &position.Position{
								StartLine: 46,
								EndLine:   46,
								StartPos:  1075,
								EndPos:    1079,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 46,
									EndLine:   46,
									StartPos:  1075,
									EndPos:    1079,
								},
								Value: "foo",
							},
						},
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 46,
								EndLine:   46,
								StartPos:  1079,
								EndPos:    1080,
							},
							Value: " ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 46,
								EndLine:   46,
								StartPos:  1080,
								EndPos:    1084,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 46,
									EndLine:   46,
									StartPos:  1080,
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
					StartPos:  1089,
					EndPos:    1108,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 47,
						EndLine:   47,
						StartPos:  1089,
						EndPos:    1107,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 47,
								EndLine:   47,
								StartPos:  1090,
								EndPos:    1095,
							},
							Value: "test ",
						},
						&expr.PropertyFetch{
							Position: &position.Position{
								StartLine: 47,
								EndLine:   47,
								StartPos:  1095,
								EndPos:    1104,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 47,
									EndLine:   47,
									StartPos:  1095,
									EndPos:    1099,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 47,
										EndLine:   47,
										StartPos:  1095,
										EndPos:    1099,
									},
									Value: "foo",
								},
							},
							Property: &node.Identifier{
								Position: &position.Position{
									StartLine: 47,
									EndLine:   47,
									StartPos:  1101,
									EndPos:    1104,
								},
								Value: "bar",
							},
						},
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 47,
								EndLine:   47,
								StartPos:  1104,
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
					StartPos:  1111,
					EndPos:    1125,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 48,
						EndLine:   48,
						StartPos:  1111,
						EndPos:    1124,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 48,
								EndLine:   48,
								StartPos:  1112,
								EndPos:    1117,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 48,
								EndLine:   48,
								StartPos:  1117,
								EndPos:    1123,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 48,
									EndLine:   48,
									StartPos:  1119,
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
					StartPos:  1128,
					EndPos:    1145,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 49,
						EndLine:   49,
						StartPos:  1128,
						EndPos:    1144,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 49,
								EndLine:   49,
								StartPos:  1129,
								EndPos:    1134,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 49,
								EndLine:   49,
								StartPos:  1134,
								EndPos:    1143,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 49,
									EndLine:   49,
									StartPos:  1136,
									EndPos:    1139,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 49,
										EndLine:   49,
										StartPos:  1136,
										EndPos:    1139,
									},
									Value: "foo",
								},
							},
							Dim: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 49,
									EndLine:   49,
									StartPos:  1140,
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
					StartPos:  1148,
					EndPos:    1163,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 50,
						EndLine:   50,
						StartPos:  1148,
						EndPos:    1162,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 50,
								EndLine:   50,
								StartPos:  1149,
								EndPos:    1154,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 50,
								EndLine:   50,
								StartPos:  1154,
								EndPos:    1161,
							},
							VarName: &expr.Variable{
								Position: &position.Position{
									StartLine: 50,
									EndLine:   50,
									StartPos:  1156,
									EndPos:    1160,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 50,
										EndLine:   50,
										StartPos:  1156,
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
					StartPos:  1166,
					EndPos:    1187,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 51,
						EndLine:   51,
						StartPos:  1166,
						EndPos:    1186,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 51,
								EndLine:   51,
								StartPos:  1167,
								EndPos:    1172,
							},
							Value: "test ",
						},
						&expr.MethodCall{
							Position: &position.Position{
								StartLine: 51,
								EndLine:   51,
								StartPos:  1173,
								EndPos:    1184,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 51,
									EndLine:   51,
									StartPos:  1173,
									EndPos:    1177,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 51,
										EndLine:   51,
										StartPos:  1173,
										EndPos:    1177,
									},
									Value: "foo",
								},
							},
							Method: &node.Identifier{
								Position: &position.Position{
									StartLine: 51,
									EndLine:   51,
									StartPos:  1179,
									EndPos:    1182,
								},
								Value: "bar",
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 51,
									EndLine:   51,
									StartPos:  1182,
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
					StartPos:  1191,
					EndPos:    1209,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 53,
						EndLine:   53,
						StartPos:  1195,
						EndPos:    1197,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 53,
							EndLine:   53,
							StartPos:  1195,
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
					StartPos:  1212,
					EndPos:    1245,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 55,
						EndLine:   55,
						StartPos:  1216,
						EndPos:    1218,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 55,
							EndLine:   55,
							StartPos:  1216,
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
							StartPos:  1224,
							EndPos:    -1,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 56,
								EndLine:   56,
								StartPos:  1232,
								EndPos:    1234,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 56,
									EndLine:   56,
									StartPos:  1232,
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
					StartPos:  1248,
					EndPos:    1274,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 58,
						EndLine:   58,
						StartPos:  1252,
						EndPos:    1254,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 58,
							EndLine:   58,
							StartPos:  1252,
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
						StartPos:  1260,
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
					StartPos:  1277,
					EndPos:    1333,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 61,
						EndLine:   61,
						StartPos:  1281,
						EndPos:    1283,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 61,
							EndLine:   61,
							StartPos:  1281,
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
							StartPos:  1289,
							EndPos:    -1,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 62,
								EndLine:   62,
								StartPos:  1297,
								EndPos:    1299,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 62,
									EndLine:   62,
									StartPos:  1297,
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
							StartPos:  1304,
							EndPos:    -1,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 63,
								EndLine:   63,
								StartPos:  1312,
								EndPos:    1314,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 63,
									EndLine:   63,
									StartPos:  1312,
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
						StartPos:  1319,
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
					StartPos:  1337,
					EndPos:    1357,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 67,
						EndLine:   67,
						StartPos:  1344,
						EndPos:    1345,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 67,
						EndLine:   67,
						StartPos:  1347,
						EndPos:    1357,
					},
					Stmts: []node.Node{
						&stmt.Break{
							Position: &position.Position{
								StartLine: 67,
								EndLine:   67,
								StartPos:  1349,
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
					StartPos:  1360,
					EndPos:    1382,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 68,
						EndLine:   68,
						StartPos:  1367,
						EndPos:    1368,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 68,
						EndLine:   68,
						StartPos:  1370,
						EndPos:    1382,
					},
					Stmts: []node.Node{
						&stmt.Break{
							Position: &position.Position{
								StartLine: 68,
								EndLine:   68,
								StartPos:  1372,
								EndPos:    1380,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 68,
									EndLine:   68,
									StartPos:  1378,
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
					StartPos:  1385,
					EndPos:    1416,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 69,
						EndLine:   69,
						StartPos:  1392,
						EndPos:    1393,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 69,
						EndLine:   69,
						StartPos:  1397,
						EndPos:    1406,
					},
					Stmts: []node.Node{
						&stmt.Break{
							Position: &position.Position{
								StartLine: 69,
								EndLine:   69,
								StartPos:  1397,
								EndPos:    1406,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 69,
									EndLine:   69,
									StartPos:  1403,
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
					StartPos:  1419,
					EndPos:    1462,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 70,
						EndLine:   70,
						StartPos:  1425,
						EndPos:    1428,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassConstList{
						Position: &position.Position{
							StartLine: 70,
							EndLine:   70,
							StartPos:  1430,
							EndPos:    1460,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 70,
									EndLine:   70,
									StartPos:  1430,
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
									StartPos:  1443,
									EndPos:    1450,
								},
								PhpDocComment: "",
								ConstantName: &node.Identifier{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1443,
										EndPos:    1446,
									},
									Value: "FOO",
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1449,
										EndPos:    1450,
									},
									Value: "1",
								},
							},
							&stmt.Constant{
								Position: &position.Position{
									StartLine: 70,
									EndLine:   70,
									StartPos:  1452,
									EndPos:    1459,
								},
								PhpDocComment: "",
								ConstantName: &node.Identifier{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1452,
										EndPos:    1455,
									},
									Value: "BAR",
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1458,
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
					StartPos:  1465,
					EndPos:    1501,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 71,
						EndLine:   71,
						StartPos:  1471,
						EndPos:    1474,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassConstList{
						Position: &position.Position{
							StartLine: 71,
							EndLine:   71,
							StartPos:  1476,
							EndPos:    1499,
						},
						Consts: []node.Node{
							&stmt.Constant{
								Position: &position.Position{
									StartLine: 71,
									EndLine:   71,
									StartPos:  1482,
									EndPos:    1489,
								},
								PhpDocComment: "",
								ConstantName: &node.Identifier{
									Position: &position.Position{
										StartLine: 71,
										EndLine:   71,
										StartPos:  1482,
										EndPos:    1485,
									},
									Value: "FOO",
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 71,
										EndLine:   71,
										StartPos:  1488,
										EndPos:    1489,
									},
									Value: "1",
								},
							},
							&stmt.Constant{
								Position: &position.Position{
									StartLine: 71,
									EndLine:   71,
									StartPos:  1491,
									EndPos:    1498,
								},
								PhpDocComment: "",
								ConstantName: &node.Identifier{
									Position: &position.Position{
										StartLine: 71,
										EndLine:   71,
										StartPos:  1491,
										EndPos:    1494,
									},
									Value: "BAR",
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 71,
										EndLine:   71,
										StartPos:  1497,
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
					StartPos:  1504,
					EndPos:    1534,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 72,
						EndLine:   72,
						StartPos:  1510,
						EndPos:    1513,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 72,
							EndLine:   72,
							StartPos:  1515,
							EndPos:    1532,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 72,
								EndLine:   72,
								StartPos:  1524,
								EndPos:    1527,
							},
							Value: "bar",
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 72,
								EndLine:   72,
								StartPos:  1530,
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
					StartPos:  1537,
					EndPos:    1582,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 73,
						EndLine:   73,
						StartPos:  1543,
						EndPos:    1546,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 73,
							EndLine:   73,
							StartPos:  1548,
							EndPos:    1580,
						},
						ReturnsRef:    true,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 73,
								EndLine:   73,
								StartPos:  1572,
								EndPos:    1575,
							},
							Value: "bar",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 73,
									EndLine:   73,
									StartPos:  1548,
									EndPos:    1554,
								},
								Value: "public",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 73,
									EndLine:   73,
									StartPos:  1555,
									EndPos:    1561,
								},
								Value: "static",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 73,
								EndLine:   73,
								StartPos:  1578,
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
					StartPos:  1585,
					EndPos:    1636,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 74,
						EndLine:   74,
						StartPos:  1591,
						EndPos:    1594,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 74,
							EndLine:   74,
							StartPos:  1596,
							EndPos:    1634,
						},
						ReturnsRef:    true,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 74,
								EndLine:   74,
								StartPos:  1620,
								EndPos:    1623,
							},
							Value: "bar",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 74,
									EndLine:   74,
									StartPos:  1596,
									EndPos:    1602,
								},
								Value: "public",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 74,
									EndLine:   74,
									StartPos:  1603,
									EndPos:    1609,
								},
								Value: "static",
							},
						},
						ReturnType: &name.Name{
							Position: &position.Position{
								StartLine: 74,
								EndLine:   74,
								StartPos:  1627,
								EndPos:    1631,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 74,
										EndLine:   74,
										StartPos:  1627,
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
								StartPos:  1632,
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
					StartPos:  1639,
					EndPos:    1660,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 75,
						EndLine:   75,
						StartPos:  1654,
						EndPos:    1657,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 75,
							EndLine:   75,
							StartPos:  1639,
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
					StartPos:  1663,
					EndPos:    1694,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 76,
						EndLine:   76,
						StartPos:  1675,
						EndPos:    1678,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 76,
							EndLine:   76,
							StartPos:  1663,
							EndPos:    1668,
						},
						Value: "final",
					},
				},
				Extends: &stmt.ClassExtends{
					Position: &position.Position{
						StartLine: 76,
						EndLine:   76,
						StartPos:  1679,
						EndPos:    1690,
					},
					ClassName: &name.Name{
						Position: &position.Position{
							StartLine: 76,
							EndLine:   76,
							StartPos:  1687,
							EndPos:    1690,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 76,
									EndLine:   76,
									StartPos:  1687,
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
					StartPos:  1697,
					EndPos:    1731,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 77,
						EndLine:   77,
						StartPos:  1709,
						EndPos:    1712,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 77,
							EndLine:   77,
							StartPos:  1697,
							EndPos:    1702,
						},
						Value: "final",
					},
				},
				Implements: &stmt.ClassImplements{
					Position: &position.Position{
						StartLine: 77,
						EndLine:   77,
						StartPos:  1713,
						EndPos:    1727,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1724,
								EndPos:    1727,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 77,
										EndLine:   77,
										StartPos:  1724,
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
					StartPos:  1734,
					EndPos:    1773,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 78,
						EndLine:   78,
						StartPos:  1746,
						EndPos:    1749,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 78,
							EndLine:   78,
							StartPos:  1734,
							EndPos:    1739,
						},
						Value: "final",
					},
				},
				Implements: &stmt.ClassImplements{
					Position: &position.Position{
						StartLine: 78,
						EndLine:   78,
						StartPos:  1750,
						EndPos:    1769,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 78,
								EndLine:   78,
								StartPos:  1761,
								EndPos:    1764,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 78,
										EndLine:   78,
										StartPos:  1761,
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
								StartPos:  1766,
								EndPos:    1769,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 78,
										EndLine:   78,
										StartPos:  1766,
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
					StartPos:  1776,
					EndPos:    1824,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 79,
						EndLine:   79,
						StartPos:  1776,
						EndPos:    1823,
					},
					Class: &stmt.Class{
						Position: &position.Position{
							StartLine: 79,
							EndLine:   79,
							StartPos:  1780,
							EndPos:    1823,
						},
						PhpDocComment: "",
						ArgumentList: &node.ArgumentList{
							Position: &position.Position{
								StartLine: 79,
								EndLine:   79,
								StartPos:  1785,
								EndPos:    1787,
							},
						},
						Extends: &stmt.ClassExtends{
							Position: &position.Position{
								StartLine: 79,
								EndLine:   79,
								StartPos:  1788,
								EndPos:    1799,
							},
							ClassName: &name.Name{
								Position: &position.Position{
									StartLine: 79,
									EndLine:   79,
									StartPos:  1796,
									EndPos:    1799,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 79,
											EndLine:   79,
											StartPos:  1796,
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
								StartPos:  1800,
								EndPos:    1819,
							},
							InterfaceNames: []node.Node{
								&name.Name{
									Position: &position.Position{
										StartLine: 79,
										EndLine:   79,
										StartPos:  1811,
										EndPos:    1814,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 79,
												EndLine:   79,
												StartPos:  1811,
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
										StartPos:  1816,
										EndPos:    1819,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 79,
												EndLine:   79,
												StartPos:  1816,
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
					StartPos:  1828,
					EndPos:    1851,
				},
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 81,
							EndLine:   81,
							StartPos:  1834,
							EndPos:    1841,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1834,
								EndPos:    1837,
							},
							Value: "FOO",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1840,
								EndPos:    1841,
							},
							Value: "1",
						},
					},
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 81,
							EndLine:   81,
							StartPos:  1843,
							EndPos:    1850,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1843,
								EndPos:    1846,
							},
							Value: "BAR",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1849,
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
					StartPos:  1854,
					EndPos:    1877,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 82,
						EndLine:   82,
						StartPos:  1861,
						EndPos:    1862,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 82,
						EndLine:   82,
						StartPos:  1864,
						EndPos:    1877,
					},
					Stmts: []node.Node{
						&stmt.Continue{
							Position: &position.Position{
								StartLine: 82,
								EndLine:   82,
								StartPos:  1866,
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
					StartPos:  1880,
					EndPos:    1905,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 83,
						EndLine:   83,
						StartPos:  1887,
						EndPos:    1888,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 83,
						EndLine:   83,
						StartPos:  1890,
						EndPos:    1905,
					},
					Stmts: []node.Node{
						&stmt.Continue{
							Position: &position.Position{
								StartLine: 83,
								EndLine:   83,
								StartPos:  1892,
								EndPos:    1903,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 83,
									EndLine:   83,
									StartPos:  1901,
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
					StartPos:  1908,
					EndPos:    1934,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 84,
						EndLine:   84,
						StartPos:  1915,
						EndPos:    1916,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 84,
						EndLine:   84,
						StartPos:  1918,
						EndPos:    1934,
					},
					Stmts: []node.Node{
						&stmt.Continue{
							Position: &position.Position{
								StartLine: 84,
								EndLine:   84,
								StartPos:  1920,
								EndPos:    1932,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 84,
									EndLine:   84,
									StartPos:  1929,
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
					StartPos:  1937,
					EndPos:    1954,
				},
				Alt: false,
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 85,
							EndLine:   85,
							StartPos:  1945,
							EndPos:    1952,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 85,
								EndLine:   85,
								StartPos:  1945,
								EndPos:    1950,
							},
							Value: "ticks",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 85,
								EndLine:   85,
								StartPos:  1951,
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
						StartPos:  1953,
						EndPos:    1954,
					},
				},
			},
			&stmt.Declare{
				Position: &position.Position{
					StartLine: 86,
					EndLine:   86,
					StartPos:  1957,
					EndPos:    1976,
				},
				Alt: false,
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 86,
							EndLine:   86,
							StartPos:  1965,
							EndPos:    1972,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 86,
								EndLine:   86,
								StartPos:  1965,
								EndPos:    1970,
							},
							Value: "ticks",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 86,
								EndLine:   86,
								StartPos:  1971,
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
						StartPos:  1974,
						EndPos:    1976,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Declare{
				Position: &position.Position{
					StartLine: 87,
					EndLine:   87,
					StartPos:  1979,
					EndPos:    2008,
				},
				Alt: true,
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 87,
							EndLine:   87,
							StartPos:  1987,
							EndPos:    1994,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 87,
								EndLine:   87,
								StartPos:  1987,
								EndPos:    1992,
							},
							Value: "ticks",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 87,
								EndLine:   87,
								StartPos:  1993,
								EndPos:    1994,
							},
							Value: "1",
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
			&stmt.Do{
				Position: &position.Position{
					StartLine: 88,
					EndLine:   88,
					StartPos:  2011,
					EndPos:    2026,
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 88,
						EndLine:   88,
						StartPos:  2014,
						EndPos:    2016,
					},
					Stmts: []node.Node{},
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 88,
						EndLine:   88,
						StartPos:  2023,
						EndPos:    2024,
					},
					Value: "1",
				},
			},
			&stmt.Echo{
				Position: &position.Position{
					StartLine: 89,
					EndLine:   89,
					StartPos:  2029,
					EndPos:    2040,
				},
				Exprs: []node.Node{
					&expr.Variable{
						Position: &position.Position{
							StartLine: 89,
							EndLine:   89,
							StartPos:  2034,
							EndPos:    2036,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 89,
								EndLine:   89,
								StartPos:  2034,
								EndPos:    2036,
							},
							Value: "a",
						},
					},
					&scalar.Lnumber{
						Position: &position.Position{
							StartLine: 89,
							EndLine:   89,
							StartPos:  2038,
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
					StartPos:  2043,
					EndPos:    2052,
				},
				Exprs: []node.Node{
					&expr.Variable{
						Position: &position.Position{
							StartLine: 90,
							EndLine:   90,
							StartPos:  2048,
							EndPos:    2050,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 90,
								EndLine:   90,
								StartPos:  2048,
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
					StartPos:  2055,
					EndPos:    2090,
				},
				Init: []node.Node{
					&assign.Assign{
						Position: &position.Position{
							StartLine: 91,
							EndLine:   91,
							StartPos:  2059,
							EndPos:    2065,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2059,
								EndPos:    2061,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2059,
									EndPos:    2061,
								},
								Value: "i",
							},
						},
						Expression: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2064,
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
							StartPos:  2067,
							EndPos:    2074,
						},
						Left: &expr.Variable{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2067,
								EndPos:    2069,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2067,
									EndPos:    2069,
								},
								Value: "i",
							},
						},
						Right: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2072,
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
							StartPos:  2076,
							EndPos:    2080,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2076,
								EndPos:    2078,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2076,
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
							StartPos:  2082,
							EndPos:    2086,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2082,
								EndPos:    2084,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2082,
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
						StartPos:  2088,
						EndPos:    2090,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.AltFor{
				Position: &position.Position{
					StartLine: 92,
					EndLine:   92,
					StartPos:  2093,
					EndPos:    2129,
				},
				Cond: []node.Node{
					&binary.Smaller{
						Position: &position.Position{
							StartLine: 92,
							EndLine:   92,
							StartPos:  2099,
							EndPos:    2106,
						},
						Left: &expr.Variable{
							Position: &position.Position{
								StartLine: 92,
								EndLine:   92,
								StartPos:  2099,
								EndPos:    2101,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  2099,
									EndPos:    2101,
								},
								Value: "i",
							},
						},
						Right: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 92,
								EndLine:   92,
								StartPos:  2104,
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
							StartPos:  2108,
							EndPos:    2112,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 92,
								EndLine:   92,
								StartPos:  2108,
								EndPos:    2110,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  2108,
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
							StartPos:  2114,
							EndPos:    2118,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 92,
								EndLine:   92,
								StartPos:  2114,
								EndPos:    2116,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  2114,
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
					StartPos:  2132,
					EndPos:    2153,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 93,
						EndLine:   93,
						StartPos:  2141,
						EndPos:    2143,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 93,
							EndLine:   93,
							StartPos:  2141,
							EndPos:    2143,
						},
						Value: "a",
					},
				},
				Variable: &expr.Variable{
					Position: &position.Position{
						StartLine: 93,
						EndLine:   93,
						StartPos:  2147,
						EndPos:    2149,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 93,
							EndLine:   93,
							StartPos:  2147,
							EndPos:    2149,
						},
						Value: "v",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 93,
						EndLine:   93,
						StartPos:  2151,
						EndPos:    2153,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.AltForeach{
				Position: &position.Position{
					StartLine: 94,
					EndLine:   94,
					StartPos:  2156,
					EndPos:    2188,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 94,
						EndLine:   94,
						StartPos:  2165,
						EndPos:    2167,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 94,
							EndLine:   94,
							StartPos:  2165,
							EndPos:    2167,
						},
						Value: "a",
					},
				},
				Variable: &expr.Variable{
					Position: &position.Position{
						StartLine: 94,
						EndLine:   94,
						StartPos:  2171,
						EndPos:    2173,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 94,
							EndLine:   94,
							StartPos:  2171,
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
					StartPos:  2191,
					EndPos:    2218,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 95,
						EndLine:   95,
						StartPos:  2200,
						EndPos:    2202,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2200,
							EndPos:    2202,
						},
						Value: "a",
					},
				},
				Key: &expr.Variable{
					Position: &position.Position{
						StartLine: 95,
						EndLine:   95,
						StartPos:  2206,
						EndPos:    2208,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2206,
							EndPos:    2208,
						},
						Value: "k",
					},
				},
				Variable: &expr.Variable{
					Position: &position.Position{
						StartLine: 95,
						EndLine:   95,
						StartPos:  2212,
						EndPos:    2214,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2212,
							EndPos:    2214,
						},
						Value: "v",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 95,
						EndLine:   95,
						StartPos:  2216,
						EndPos:    2218,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Foreach{
				Position: &position.Position{
					StartLine: 96,
					EndLine:   96,
					StartPos:  2221,
					EndPos:    2249,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 96,
						EndLine:   96,
						StartPos:  2230,
						EndPos:    2232,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2230,
							EndPos:    2232,
						},
						Value: "a",
					},
				},
				Key: &expr.Variable{
					Position: &position.Position{
						StartLine: 96,
						EndLine:   96,
						StartPos:  2236,
						EndPos:    2238,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2236,
							EndPos:    2238,
						},
						Value: "k",
					},
				},
				Variable: &expr.Reference{
					Position: &position.Position{
						StartLine: 96,
						EndLine:   96,
						StartPos:  2242,
						EndPos:    2245,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2243,
							EndPos:    2245,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 96,
								EndLine:   96,
								StartPos:  2243,
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
						StartPos:  2247,
						EndPos:    2249,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Foreach{
				Position: &position.Position{
					StartLine: 97,
					EndLine:   97,
					StartPos:  2252,
					EndPos:    2285,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 97,
						EndLine:   97,
						StartPos:  2261,
						EndPos:    2263,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 97,
							EndLine:   97,
							StartPos:  2261,
							EndPos:    2263,
						},
						Value: "a",
					},
				},
				Key: &expr.Variable{
					Position: &position.Position{
						StartLine: 97,
						EndLine:   97,
						StartPos:  2267,
						EndPos:    2269,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 97,
							EndLine:   97,
							StartPos:  2267,
							EndPos:    2269,
						},
						Value: "k",
					},
				},
				Variable: &expr.List{
					Position: &position.Position{
						StartLine: 97,
						EndLine:   97,
						StartPos:  2273,
						EndPos:    2281,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2278,
								EndPos:    2280,
							},
							Val: &expr.Variable{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2278,
									EndPos:    2280,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 97,
										EndLine:   97,
										StartPos:  2278,
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
						StartPos:  2283,
						EndPos:    2285,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Foreach{
				Position: &position.Position{
					StartLine: 98,
					EndLine:   98,
					StartPos:  2288,
					EndPos:    2317,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 98,
						EndLine:   98,
						StartPos:  2297,
						EndPos:    2299,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2297,
							EndPos:    2299,
						},
						Value: "a",
					},
				},
				Key: &expr.Variable{
					Position: &position.Position{
						StartLine: 98,
						EndLine:   98,
						StartPos:  2303,
						EndPos:    2305,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2303,
							EndPos:    2305,
						},
						Value: "k",
					},
				},
				Variable: &expr.ShortList{
					Position: &position.Position{
						StartLine: 98,
						EndLine:   98,
						StartPos:  2309,
						EndPos:    2313,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 98,
								EndLine:   98,
								StartPos:  2310,
								EndPos:    2312,
							},
							Val: &expr.Variable{
								Position: &position.Position{
									StartLine: 98,
									EndLine:   98,
									StartPos:  2310,
									EndPos:    2312,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 98,
										EndLine:   98,
										StartPos:  2310,
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
						StartPos:  2315,
						EndPos:    2317,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 99,
					EndLine:   99,
					StartPos:  2320,
					EndPos:    2337,
				},
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 99,
						EndLine:   99,
						StartPos:  2329,
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
					StartPos:  2340,
					EndPos:    2364,
				},
				PhpDocComment: "",
				ReturnsRef:    false,
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 100,
						EndLine:   100,
						StartPos:  2349,
						EndPos:    2352,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.Return{
						Position: &position.Position{
							StartLine: 100,
							EndLine:   100,
							StartPos:  2356,
							EndPos:    2363,
						},
					},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 101,
					EndLine:   101,
					StartPos:  2367,
					EndPos:    2394,
				},
				ReturnsRef:    true,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 101,
						EndLine:   101,
						StartPos:  2377,
						EndPos:    2380,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.Return{
						Position: &position.Position{
							StartLine: 101,
							EndLine:   101,
							StartPos:  2384,
							EndPos:    2393,
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 101,
								EndLine:   101,
								StartPos:  2391,
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
					StartPos:  2397,
					EndPos:    2421,
				},
				ReturnsRef:    true,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 102,
						EndLine:   102,
						StartPos:  2407,
						EndPos:    2410,
					},
					Value: "foo",
				},
				ReturnType: &name.Name{
					Position: &position.Position{
						StartLine: 102,
						EndLine:   102,
						StartPos:  2414,
						EndPos:    2418,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 102,
								EndLine:   102,
								StartPos:  2414,
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
					StartPos:  2424,
					EndPos:    2438,
				},
				Vars: []node.Node{
					&expr.Variable{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2431,
							EndPos:    2433,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2431,
								EndPos:    2433,
							},
							Value: "a",
						},
					},
					&expr.Variable{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2435,
							EndPos:    2437,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2435,
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
					StartPos:  2441,
					EndPos:    2443,
				},
				LabelName: &node.Identifier{
					Position: &position.Position{
						StartLine: 104,
						EndLine:   104,
						StartPos:  2441,
						EndPos:    2442,
					},
					Value: "a",
				},
			},
			&stmt.Goto{
				Position: &position.Position{
					StartLine: 105,
					EndLine:   105,
					StartPos:  2447,
					EndPos:    2454,
				},
				Label: &node.Identifier{
					Position: &position.Position{
						StartLine: 105,
						EndLine:   105,
						StartPos:  2452,
						EndPos:    2453,
					},
					Value: "a",
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 106,
					EndLine:   106,
					StartPos:  2457,
					EndPos:    2467,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 106,
						EndLine:   106,
						StartPos:  2461,
						EndPos:    2463,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 106,
							EndLine:   106,
							StartPos:  2461,
							EndPos:    2463,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 106,
						EndLine:   106,
						StartPos:  2465,
						EndPos:    2467,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 107,
					EndLine:   107,
					StartPos:  2470,
					EndPos:    2495,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 107,
						EndLine:   107,
						StartPos:  2474,
						EndPos:    2476,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 107,
							EndLine:   107,
							StartPos:  2474,
							EndPos:    2476,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 107,
						EndLine:   107,
						StartPos:  2478,
						EndPos:    2480,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 107,
							EndLine:   107,
							StartPos:  2481,
							EndPos:    2495,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 107,
								EndLine:   107,
								StartPos:  2489,
								EndPos:    2491,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 107,
									EndLine:   107,
									StartPos:  2489,
									EndPos:    2491,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 107,
								EndLine:   107,
								StartPos:  2493,
								EndPos:    2495,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 108,
					EndLine:   108,
					StartPos:  2498,
					EndPos:    2516,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 108,
						EndLine:   108,
						StartPos:  2502,
						EndPos:    2504,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 108,
							EndLine:   108,
							StartPos:  2502,
							EndPos:    2504,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 108,
						EndLine:   108,
						StartPos:  2506,
						EndPos:    2508,
					},
					Stmts: []node.Node{},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 108,
						EndLine:   108,
						StartPos:  2509,
						EndPos:    2516,
					},
					Stmt: &stmt.StmtList{
						Position: &position.Position{
							StartLine: 108,
							EndLine:   108,
							StartPos:  2514,
							EndPos:    2516,
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 109,
					EndLine:   109,
					StartPos:  2519,
					EndPos:    2567,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 109,
						EndLine:   109,
						StartPos:  2523,
						EndPos:    2525,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 109,
							EndLine:   109,
							StartPos:  2523,
							EndPos:    2525,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 109,
						EndLine:   109,
						StartPos:  2527,
						EndPos:    2529,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 109,
							EndLine:   109,
							StartPos:  2530,
							EndPos:    2544,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 109,
								EndLine:   109,
								StartPos:  2538,
								EndPos:    2540,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 109,
									EndLine:   109,
									StartPos:  2538,
									EndPos:    2540,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 109,
								EndLine:   109,
								StartPos:  2542,
								EndPos:    2544,
							},
							Stmts: []node.Node{},
						},
					},
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 109,
							EndLine:   109,
							StartPos:  2545,
							EndPos:    2559,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 109,
								EndLine:   109,
								StartPos:  2553,
								EndPos:    2555,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 109,
									EndLine:   109,
									StartPos:  2553,
									EndPos:    2555,
								},
								Value: "c",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 109,
								EndLine:   109,
								StartPos:  2557,
								EndPos:    2559,
							},
							Stmts: []node.Node{},
						},
					},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 109,
						EndLine:   109,
						StartPos:  2560,
						EndPos:    2567,
					},
					Stmt: &stmt.StmtList{
						Position: &position.Position{
							StartLine: 109,
							EndLine:   109,
							StartPos:  2565,
							EndPos:    2567,
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 110,
					EndLine:   110,
					StartPos:  2570,
					EndPos:    2619,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 110,
						EndLine:   110,
						StartPos:  2574,
						EndPos:    2576,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2574,
							EndPos:    2576,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 110,
						EndLine:   110,
						StartPos:  2578,
						EndPos:    2580,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2581,
							EndPos:    2595,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2589,
								EndPos:    2591,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2589,
									EndPos:    2591,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2593,
								EndPos:    2595,
							},
							Stmts: []node.Node{},
						},
					},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 110,
						EndLine:   110,
						StartPos:  2596,
						EndPos:    2619,
					},
					Stmt: &stmt.If{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2601,
							EndPos:    2619,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2605,
								EndPos:    2607,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2605,
									EndPos:    2607,
								},
								Value: "c",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2609,
								EndPos:    2611,
							},
							Stmts: []node.Node{},
						},
						Else: &stmt.Else{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2612,
								EndPos:    2619,
							},
							Stmt: &stmt.StmtList{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2617,
									EndPos:    2619,
								},
								Stmts: []node.Node{},
							},
						},
					},
				},
			},
			&stmt.Nop{
				Position: &position.Position{
					StartLine: 111,
					EndLine:   111,
					StartPos:  2622,
					EndPos:    2624,
				},
			},
			&stmt.InlineHtml{
				Position: &position.Position{
					StartLine: 111,
					EndLine:   111,
					StartPos:  2624,
					EndPos:    2637,
				},
				Value: " <div></div> ",
			},
			&stmt.Interface{
				Position: &position.Position{
					StartLine: 112,
					EndLine:   112,
					StartPos:  2642,
					EndPos:    2658,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 112,
						EndLine:   112,
						StartPos:  2652,
						EndPos:    2655,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{},
			},
			&stmt.Interface{
				Position: &position.Position{
					StartLine: 113,
					EndLine:   113,
					StartPos:  2661,
					EndPos:    2689,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 113,
						EndLine:   113,
						StartPos:  2671,
						EndPos:    2674,
					},
					Value: "Foo",
				},
				Extends: &stmt.InterfaceExtends{
					Position: &position.Position{
						StartLine: 113,
						EndLine:   113,
						StartPos:  2675,
						EndPos:    2686,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 113,
								EndLine:   113,
								StartPos:  2683,
								EndPos:    2686,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 113,
										EndLine:   113,
										StartPos:  2683,
										EndPos:    2686,
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
					StartLine: 114,
					EndLine:   114,
					StartPos:  2692,
					EndPos:    2725,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 114,
						EndLine:   114,
						StartPos:  2702,
						EndPos:    2705,
					},
					Value: "Foo",
				},
				Extends: &stmt.InterfaceExtends{
					Position: &position.Position{
						StartLine: 114,
						EndLine:   114,
						StartPos:  2706,
						EndPos:    2722,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 114,
								EndLine:   114,
								StartPos:  2714,
								EndPos:    2717,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2714,
										EndPos:    2717,
									},
									Value: "Bar",
								},
							},
						},
						&name.Name{
							Position: &position.Position{
								StartLine: 114,
								EndLine:   114,
								StartPos:  2719,
								EndPos:    2722,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2719,
										EndPos:    2722,
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
					StartLine: 115,
					EndLine:   115,
					StartPos:  2728,
					EndPos:    2742,
				},
				NamespaceName: &name.Name{
					Position: &position.Position{
						StartLine: 115,
						EndLine:   115,
						StartPos:  2738,
						EndPos:    2741,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 115,
								EndLine:   115,
								StartPos:  2738,
								EndPos:    2741,
							},
							Value: "Foo",
						},
					},
				},
			},
			&stmt.Namespace{
				Position: &position.Position{
					StartLine: 116,
					EndLine:   116,
					StartPos:  2745,
					EndPos:    2761,
				},
				NamespaceName: &name.Name{
					Position: &position.Position{
						StartLine: 116,
						EndLine:   116,
						StartPos:  2755,
						EndPos:    2758,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 116,
								EndLine:   116,
								StartPos:  2755,
								EndPos:    2758,
							},
							Value: "Foo",
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Namespace{
				Position: &position.Position{
					StartLine: 117,
					EndLine:   117,
					StartPos:  2764,
					EndPos:    2776,
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 118,
					EndLine:   118,
					StartPos:  2779,
					EndPos:    2798,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 118,
						EndLine:   118,
						StartPos:  2785,
						EndPos:    2788,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Position: &position.Position{
							StartLine: 118,
							EndLine:   118,
							StartPos:  2790,
							EndPos:    2797,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 118,
									EndLine:   118,
									StartPos:  2790,
									EndPos:    2793,
								},
								Value: "var",
							},
						},
						Properties: []node.Node{
							&stmt.Property{
								Position: &position.Position{
									StartLine: 118,
									EndLine:   118,
									StartPos:  2794,
									EndPos:    2796,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 118,
										EndLine:   118,
										StartPos:  2794,
										EndPos:    2796,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 118,
											EndLine:   118,
											StartPos:  2794,
											EndPos:    2796,
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
					StartLine: 119,
					EndLine:   119,
					StartPos:  2801,
					EndPos:    2838,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 119,
						EndLine:   119,
						StartPos:  2807,
						EndPos:    2810,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Position: &position.Position{
							StartLine: 119,
							EndLine:   119,
							StartPos:  2812,
							EndPos:    2837,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 119,
									EndLine:   119,
									StartPos:  2812,
									EndPos:    2818,
								},
								Value: "public",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 119,
									EndLine:   119,
									StartPos:  2819,
									EndPos:    2825,
								},
								Value: "static",
							},
						},
						Properties: []node.Node{
							&stmt.Property{
								Position: &position.Position{
									StartLine: 119,
									EndLine:   119,
									StartPos:  2826,
									EndPos:    2828,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 119,
										EndLine:   119,
										StartPos:  2826,
										EndPos:    2828,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 119,
											EndLine:   119,
											StartPos:  2826,
											EndPos:    2828,
										},
										Value: "a",
									},
								},
							},
							&stmt.Property{
								Position: &position.Position{
									StartLine: 119,
									EndLine:   119,
									StartPos:  2830,
									EndPos:    2836,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 119,
										EndLine:   119,
										StartPos:  2830,
										EndPos:    2832,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 119,
											EndLine:   119,
											StartPos:  2830,
											EndPos:    2832,
										},
										Value: "b",
									},
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 119,
										EndLine:   119,
										StartPos:  2835,
										EndPos:    2836,
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
					StartLine: 120,
					EndLine:   120,
					StartPos:  2841,
					EndPos:    2859,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 120,
							EndLine:   120,
							StartPos:  2848,
							EndPos:    2850,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 120,
								EndLine:   120,
								StartPos:  2848,
								EndPos:    2850,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 120,
									EndLine:   120,
									StartPos:  2848,
									EndPos:    2850,
								},
								Value: "a",
							},
						},
					},
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 120,
							EndLine:   120,
							StartPos:  2852,
							EndPos:    2858,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 120,
								EndLine:   120,
								StartPos:  2852,
								EndPos:    2854,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 120,
									EndLine:   120,
									StartPos:  2852,
									EndPos:    2854,
								},
								Value: "b",
							},
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 120,
								EndLine:   120,
								StartPos:  2857,
								EndPos:    2858,
							},
							Value: "1",
						},
					},
				},
			},
			&stmt.AltSwitch{
				Position: &position.Position{
					StartLine: 122,
					EndLine:   126,
					StartPos:  2863,
					EndPos:    2922,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 122,
						EndLine:   122,
						StartPos:  2871,
						EndPos:    2872,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 123,
						EndLine:   -1,
						StartPos:  2879,
						EndPos:    -1,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 123,
								EndLine:   -1,
								StartPos:  2879,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 123,
									EndLine:   123,
									StartPos:  2884,
									EndPos:    2885,
								},
								Value: "1",
							},
							Stmts: []node.Node{},
						},
						&stmt.Default{
							Position: &position.Position{
								StartLine: 124,
								EndLine:   -1,
								StartPos:  2890,
								EndPos:    -1,
							},
							Stmts: []node.Node{},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 125,
								EndLine:   -1,
								StartPos:  2902,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 125,
									EndLine:   125,
									StartPos:  2907,
									EndPos:    2908,
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
					StartLine: 128,
					EndLine:   131,
					StartPos:  2926,
					EndPos:    2974,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 128,
						EndLine:   128,
						StartPos:  2934,
						EndPos:    2935,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 129,
						EndLine:   -1,
						StartPos:  2943,
						EndPos:    -1,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 129,
								EndLine:   -1,
								StartPos:  2943,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 129,
									EndLine:   129,
									StartPos:  2948,
									EndPos:    2949,
								},
								Value: "1",
							},
							Stmts: []node.Node{},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 130,
								EndLine:   -1,
								StartPos:  2954,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 130,
									EndLine:   130,
									StartPos:  2959,
									EndPos:    2960,
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
					StartLine: 133,
					EndLine:   136,
					StartPos:  2980,
					EndPos:    3032,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 133,
						EndLine:   133,
						StartPos:  2988,
						EndPos:    2989,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 133,
						EndLine:   136,
						StartPos:  2991,
						EndPos:    3032,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 134,
								EndLine:   134,
								StartPos:  2996,
								EndPos:    3010,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 134,
									EndLine:   134,
									StartPos:  3001,
									EndPos:    3002,
								},
								Value: "1",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 134,
										EndLine:   134,
										StartPos:  3004,
										EndPos:    3010,
									},
								},
							},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 135,
								EndLine:   135,
								StartPos:  3014,
								EndPos:    3028,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 135,
									EndLine:   135,
									StartPos:  3019,
									EndPos:    3020,
								},
								Value: "2",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 135,
										EndLine:   135,
										StartPos:  3022,
										EndPos:    3028,
									},
								},
							},
						},
					},
				},
			},
			&stmt.Switch{
				Position: &position.Position{
					StartLine: 138,
					EndLine:   141,
					StartPos:  3038,
					EndPos:    3091,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 138,
						EndLine:   138,
						StartPos:  3046,
						EndPos:    3047,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 138,
						EndLine:   141,
						StartPos:  3049,
						EndPos:    3091,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 139,
								EndLine:   139,
								StartPos:  3055,
								EndPos:    3069,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 139,
									EndLine:   139,
									StartPos:  3060,
									EndPos:    3061,
								},
								Value: "1",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 139,
										EndLine:   139,
										StartPos:  3063,
										EndPos:    3069,
									},
								},
							},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 140,
								EndLine:   140,
								StartPos:  3073,
								EndPos:    3087,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 140,
									EndLine:   140,
									StartPos:  3078,
									EndPos:    3079,
								},
								Value: "2",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 140,
										EndLine:   140,
										StartPos:  3081,
										EndPos:    3087,
									},
								},
							},
						},
					},
				},
			},
			&stmt.Throw{
				Position: &position.Position{
					StartLine: 143,
					EndLine:   143,
					StartPos:  3095,
					EndPos:    3104,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 143,
						EndLine:   143,
						StartPos:  3101,
						EndPos:    3103,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 143,
							EndLine:   143,
							StartPos:  3101,
							EndPos:    3103,
						},
						Value: "e",
					},
				},
			},
			&stmt.Trait{
				Position: &position.Position{
					StartLine: 145,
					EndLine:   145,
					StartPos:  3108,
					EndPos:    3120,
				},
				PhpDocComment: "",
				TraitName: &node.Identifier{
					Position: &position.Position{
						StartLine: 145,
						EndLine:   145,
						StartPos:  3114,
						EndPos:    3117,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 146,
					EndLine:   146,
					StartPos:  3123,
					EndPos:    3145,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 146,
						EndLine:   146,
						StartPos:  3129,
						EndPos:    3132,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 146,
							EndLine:   146,
							StartPos:  3135,
							EndPos:    3143,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 146,
									EndLine:   146,
									StartPos:  3139,
									EndPos:    3142,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 146,
											EndLine:   146,
											StartPos:  3139,
											EndPos:    3142,
										},
										Value: "Bar",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.Nop{
							Position: &position.Position{
								StartLine: 146,
								EndLine:   146,
								StartPos:  3142,
								EndPos:    3143,
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 147,
					EndLine:   147,
					StartPos:  3148,
					EndPos:    3177,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 147,
						EndLine:   147,
						StartPos:  3154,
						EndPos:    3157,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 147,
							EndLine:   147,
							StartPos:  3160,
							EndPos:    3175,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 147,
									EndLine:   147,
									StartPos:  3164,
									EndPos:    3167,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 147,
											EndLine:   147,
											StartPos:  3164,
											EndPos:    3167,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 147,
									EndLine:   147,
									StartPos:  3169,
									EndPos:    3172,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 147,
											EndLine:   147,
											StartPos:  3169,
											EndPos:    3172,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 147,
								EndLine:   147,
								StartPos:  3173,
								EndPos:    3175,
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 148,
					EndLine:   148,
					StartPos:  3180,
					EndPos:    3226,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 148,
						EndLine:   148,
						StartPos:  3186,
						EndPos:    3189,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 148,
							EndLine:   148,
							StartPos:  3192,
							EndPos:    3224,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 148,
									EndLine:   148,
									StartPos:  3196,
									EndPos:    3199,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 148,
											EndLine:   148,
											StartPos:  3196,
											EndPos:    3199,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 148,
									EndLine:   148,
									StartPos:  3201,
									EndPos:    3204,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 148,
											EndLine:   148,
											StartPos:  3201,
											EndPos:    3204,
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
								StartPos:  3205,
								EndPos:    3224,
							},
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 148,
										EndLine:   148,
										StartPos:  3207,
										EndPos:    3221,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 148,
											EndLine:   148,
											StartPos:  3207,
											EndPos:    3210,
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 148,
												EndLine:   148,
												StartPos:  3207,
												EndPos:    3210,
											},
											Value: "one",
										},
									},
									Alias: &node.Identifier{
										Position: &position.Position{
											StartLine: 148,
											EndLine:   148,
											StartPos:  3214,
											EndPos:    3221,
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
					StartLine: 149,
					EndLine:   149,
					StartPos:  3229,
					EndPos:    3274,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 149,
						EndLine:   149,
						StartPos:  3235,
						EndPos:    3238,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 149,
							EndLine:   149,
							StartPos:  3241,
							EndPos:    3272,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3245,
									EndPos:    3248,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 149,
											EndLine:   149,
											StartPos:  3245,
											EndPos:    3248,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3250,
									EndPos:    3253,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 149,
											EndLine:   149,
											StartPos:  3250,
											EndPos:    3253,
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
								StartPos:  3254,
								EndPos:    3272,
							},
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3256,
										EndPos:    3269,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 149,
											EndLine:   149,
											StartPos:  3256,
											EndPos:    3259,
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3256,
												EndPos:    3259,
											},
											Value: "one",
										},
									},
									Modifier: &node.Identifier{
										Position: &position.Position{
											StartLine: 149,
											EndLine:   149,
											StartPos:  3263,
											EndPos:    3269,
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
					StartLine: 150,
					EndLine:   150,
					StartPos:  3277,
					EndPos:    3326,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 150,
						EndLine:   150,
						StartPos:  3283,
						EndPos:    3286,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 150,
							EndLine:   150,
							StartPos:  3289,
							EndPos:    3324,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 150,
									EndLine:   150,
									StartPos:  3293,
									EndPos:    3296,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 150,
											EndLine:   150,
											StartPos:  3293,
											EndPos:    3296,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 150,
									EndLine:   150,
									StartPos:  3298,
									EndPos:    3301,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 150,
											EndLine:   150,
											StartPos:  3298,
											EndPos:    3301,
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
								StartPos:  3302,
								EndPos:    3324,
							},
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 150,
										EndLine:   150,
										StartPos:  3304,
										EndPos:    3321,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 150,
											EndLine:   150,
											StartPos:  3304,
											EndPos:    3307,
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 150,
												EndLine:   150,
												StartPos:  3304,
												EndPos:    3307,
											},
											Value: "one",
										},
									},
									Modifier: &node.Identifier{
										Position: &position.Position{
											StartLine: 150,
											EndLine:   150,
											StartPos:  3311,
											EndPos:    3317,
										},
										Value: "public",
									},
									Alias: &node.Identifier{
										Position: &position.Position{
											StartLine: 150,
											EndLine:   150,
											StartPos:  3318,
											EndPos:    3321,
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
					StartLine: 151,
					EndLine:   151,
					StartPos:  3329,
					EndPos:    3406,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 151,
						EndLine:   151,
						StartPos:  3335,
						EndPos:    3338,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 151,
							EndLine:   151,
							StartPos:  3341,
							EndPos:    3404,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 151,
									EndLine:   151,
									StartPos:  3345,
									EndPos:    3348,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3345,
											EndPos:    3348,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 151,
									EndLine:   151,
									StartPos:  3350,
									EndPos:    3353,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3350,
											EndPos:    3353,
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
								StartPos:  3354,
								EndPos:    3404,
							},
							Adaptations: []node.Node{
								&stmt.TraitUsePrecedence{
									Position: &position.Position{
										StartLine: 151,
										EndLine:   151,
										StartPos:  3356,
										EndPos:    3384,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3356,
											EndPos:    3364,
										},
										Trait: &name.Name{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3356,
												EndPos:    3359,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 151,
														EndLine:   151,
														StartPos:  3356,
														EndPos:    3359,
													},
													Value: "Bar",
												},
											},
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3361,
												EndPos:    3364,
											},
											Value: "one",
										},
									},
									Insteadof: []node.Node{
										&name.Name{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3375,
												EndPos:    3378,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 151,
														EndLine:   151,
														StartPos:  3375,
														EndPos:    3378,
													},
													Value: "Baz",
												},
											},
										},
										&name.Name{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3380,
												EndPos:    3384,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 151,
														EndLine:   151,
														StartPos:  3380,
														EndPos:    3384,
													},
													Value: "Quux",
												},
											},
										},
									},
								},
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 151,
										EndLine:   151,
										StartPos:  3386,
										EndPos:    3401,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3386,
											EndPos:    3394,
										},
										Trait: &name.Name{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3386,
												EndPos:    3389,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 151,
														EndLine:   151,
														StartPos:  3386,
														EndPos:    3389,
													},
													Value: "Baz",
												},
											},
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3391,
												EndPos:    3394,
											},
											Value: "one",
										},
									},
									Alias: &node.Identifier{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3398,
											EndPos:    3401,
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
					StartLine: 153,
					EndLine:   -1,
					StartPos:  3410,
					EndPos:    -1,
				},
				Stmts:   []node.Node{},
				Catches: []node.Node{},
			},
			&stmt.Try{
				Position: &position.Position{
					StartLine: 154,
					EndLine:   154,
					StartPos:  3419,
					EndPos:    3449,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 154,
							EndLine:   154,
							StartPos:  3426,
							EndPos:    3449,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 154,
									EndLine:   154,
									StartPos:  3433,
									EndPos:    3442,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 154,
											EndLine:   154,
											StartPos:  3433,
											EndPos:    3442,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 154,
								EndLine:   154,
								StartPos:  3443,
								EndPos:    3445,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 154,
									EndLine:   154,
									StartPos:  3443,
									EndPos:    3445,
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
					StartLine: 155,
					EndLine:   155,
					StartPos:  3452,
					EndPos:    3499,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 155,
							EndLine:   155,
							StartPos:  3459,
							EndPos:    3499,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 155,
									EndLine:   155,
									StartPos:  3466,
									EndPos:    3475,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 155,
											EndLine:   155,
											StartPos:  3466,
											EndPos:    3475,
										},
										Value: "Exception",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 155,
									EndLine:   155,
									StartPos:  3476,
									EndPos:    3492,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 155,
											EndLine:   155,
											StartPos:  3476,
											EndPos:    3492,
										},
										Value: "RuntimeException",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 155,
								EndLine:   155,
								StartPos:  3493,
								EndPos:    3495,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 155,
									EndLine:   155,
									StartPos:  3493,
									EndPos:    3495,
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
					StartPos:  3502,
					EndPos:    3563,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 156,
							EndLine:   156,
							StartPos:  3509,
							EndPos:    3532,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 156,
									EndLine:   156,
									StartPos:  3516,
									EndPos:    3525,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 156,
											EndLine:   156,
											StartPos:  3516,
											EndPos:    3525,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 156,
								EndLine:   156,
								StartPos:  3526,
								EndPos:    3528,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 156,
									EndLine:   156,
									StartPos:  3526,
									EndPos:    3528,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 156,
							EndLine:   156,
							StartPos:  3533,
							EndPos:    3563,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 156,
									EndLine:   156,
									StartPos:  3540,
									EndPos:    3556,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 156,
											EndLine:   156,
											StartPos:  3540,
											EndPos:    3556,
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
								StartPos:  3557,
								EndPos:    3559,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 156,
									EndLine:   156,
									StartPos:  3557,
									EndPos:    3559,
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
					StartPos:  3566,
					EndPos:    3607,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 157,
							EndLine:   157,
							StartPos:  3573,
							EndPos:    3596,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 157,
									EndLine:   157,
									StartPos:  3580,
									EndPos:    3589,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 157,
											EndLine:   157,
											StartPos:  3580,
											EndPos:    3589,
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
								StartPos:  3590,
								EndPos:    3592,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 157,
									EndLine:   157,
									StartPos:  3590,
									EndPos:    3592,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
				},
				Finally: &stmt.Finally{
					Position: &position.Position{
						StartLine: 157,
						EndLine:   157,
						StartPos:  3597,
						EndPos:    3607,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Unset{
				Position: &position.Position{
					StartLine: 159,
					EndLine:   159,
					StartPos:  3611,
					EndPos:    3626,
				},
				Vars: []node.Node{
					&expr.Variable{
						Position: &position.Position{
							StartLine: 159,
							EndLine:   159,
							StartPos:  3617,
							EndPos:    3619,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 159,
								EndLine:   159,
								StartPos:  3617,
								EndPos:    3619,
							},
							Value: "a",
						},
					},
					&expr.Variable{
						Position: &position.Position{
							StartLine: 159,
							EndLine:   159,
							StartPos:  3621,
							EndPos:    3623,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 159,
								EndLine:   159,
								StartPos:  3621,
								EndPos:    3623,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 161,
					EndLine:   161,
					StartPos:  3630,
					EndPos:    3638,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 161,
							EndLine:   161,
							StartPos:  3634,
							EndPos:    3637,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 161,
								EndLine:   161,
								StartPos:  3634,
								EndPos:    3637,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 161,
										EndLine:   161,
										StartPos:  3634,
										EndPos:    3637,
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
					StartLine: 162,
					EndLine:   162,
					StartPos:  3641,
					EndPos:    3650,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 162,
							EndLine:   162,
							StartPos:  3646,
							EndPos:    3649,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3646,
								EndPos:    3649,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 162,
										EndLine:   162,
										StartPos:  3646,
										EndPos:    3649,
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
					StartPos:  3653,
					EndPos:    3669,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 163,
							EndLine:   163,
							StartPos:  3658,
							EndPos:    3668,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 163,
								EndLine:   163,
								StartPos:  3658,
								EndPos:    3661,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 163,
										EndLine:   163,
										StartPos:  3658,
										EndPos:    3661,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 163,
								EndLine:   163,
								StartPos:  3665,
								EndPos:    3668,
							},
							Value: "Bar",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 164,
					EndLine:   164,
					StartPos:  3672,
					EndPos:    3685,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 164,
							EndLine:   164,
							StartPos:  3676,
							EndPos:    3679,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 164,
								EndLine:   164,
								StartPos:  3676,
								EndPos:    3679,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 164,
										EndLine:   164,
										StartPos:  3676,
										EndPos:    3679,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 164,
							EndLine:   164,
							StartPos:  3681,
							EndPos:    3684,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 164,
								EndLine:   164,
								StartPos:  3681,
								EndPos:    3684,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 164,
										EndLine:   164,
										StartPos:  3681,
										EndPos:    3684,
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
					StartLine: 165,
					EndLine:   165,
					StartPos:  3688,
					EndPos:    3708,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 165,
							EndLine:   165,
							StartPos:  3692,
							EndPos:    3695,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 165,
								EndLine:   165,
								StartPos:  3692,
								EndPos:    3695,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 165,
										EndLine:   165,
										StartPos:  3692,
										EndPos:    3695,
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
							StartPos:  3697,
							EndPos:    3707,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 165,
								EndLine:   165,
								StartPos:  3697,
								EndPos:    3700,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 165,
										EndLine:   165,
										StartPos:  3697,
										EndPos:    3700,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 165,
								EndLine:   165,
								StartPos:  3704,
								EndPos:    3707,
							},
							Value: "Baz",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 166,
					EndLine:   166,
					StartPos:  3711,
					EndPos:    3734,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 166,
						EndLine:   166,
						StartPos:  3715,
						EndPos:    3723,
					},
					Value: "function",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 166,
							EndLine:   166,
							StartPos:  3724,
							EndPos:    3727,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 166,
								EndLine:   166,
								StartPos:  3724,
								EndPos:    3727,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 166,
										EndLine:   166,
										StartPos:  3724,
										EndPos:    3727,
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
							StartPos:  3730,
							EndPos:    3733,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 166,
								EndLine:   166,
								StartPos:  3730,
								EndPos:    3733,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 166,
										EndLine:   166,
										StartPos:  3730,
										EndPos:    3733,
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
					StartLine: 167,
					EndLine:   167,
					StartPos:  3737,
					EndPos:    3774,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 167,
						EndLine:   167,
						StartPos:  3741,
						EndPos:    3749,
					},
					Value: "function",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 167,
							EndLine:   167,
							StartPos:  3750,
							EndPos:    3760,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 167,
								EndLine:   167,
								StartPos:  3750,
								EndPos:    3753,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 167,
										EndLine:   167,
										StartPos:  3750,
										EndPos:    3753,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 167,
								EndLine:   167,
								StartPos:  3757,
								EndPos:    3760,
							},
							Value: "foo",
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 167,
							EndLine:   167,
							StartPos:  3763,
							EndPos:    3773,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 167,
								EndLine:   167,
								StartPos:  3763,
								EndPos:    3766,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 167,
										EndLine:   167,
										StartPos:  3763,
										EndPos:    3766,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 167,
								EndLine:   167,
								StartPos:  3770,
								EndPos:    3773,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 168,
					EndLine:   168,
					StartPos:  3777,
					EndPos:    3797,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 168,
						EndLine:   168,
						StartPos:  3781,
						EndPos:    3786,
					},
					Value: "const",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 168,
							EndLine:   168,
							StartPos:  3787,
							EndPos:    3790,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 168,
								EndLine:   168,
								StartPos:  3787,
								EndPos:    3790,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 168,
										EndLine:   168,
										StartPos:  3787,
										EndPos:    3790,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 168,
							EndLine:   168,
							StartPos:  3793,
							EndPos:    3796,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 168,
								EndLine:   168,
								StartPos:  3793,
								EndPos:    3796,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 168,
										EndLine:   168,
										StartPos:  3793,
										EndPos:    3796,
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
					StartLine: 169,
					EndLine:   169,
					StartPos:  3800,
					EndPos:    3834,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 169,
						EndLine:   169,
						StartPos:  3804,
						EndPos:    3809,
					},
					Value: "const",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 169,
							EndLine:   169,
							StartPos:  3810,
							EndPos:    3820,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 169,
								EndLine:   169,
								StartPos:  3810,
								EndPos:    3813,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 169,
										EndLine:   169,
										StartPos:  3810,
										EndPos:    3813,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 169,
								EndLine:   169,
								StartPos:  3817,
								EndPos:    3820,
							},
							Value: "foo",
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 169,
							EndLine:   169,
							StartPos:  3823,
							EndPos:    3833,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 169,
								EndLine:   169,
								StartPos:  3823,
								EndPos:    3826,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 169,
										EndLine:   169,
										StartPos:  3823,
										EndPos:    3826,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 169,
								EndLine:   169,
								StartPos:  3830,
								EndPos:    3833,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.GroupUse{
				Position: &position.Position{
					StartLine: 171,
					EndLine:   171,
					StartPos:  3838,
					EndPos:    3858,
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 171,
						EndLine:   171,
						StartPos:  3843,
						EndPos:    3846,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 171,
								EndLine:   171,
								StartPos:  3843,
								EndPos:    3846,
							},
							Value: "Foo",
						},
					},
				},
				UseList: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 171,
							EndLine:   171,
							StartPos:  3848,
							EndPos:    3851,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 171,
								EndLine:   171,
								StartPos:  3848,
								EndPos:    3851,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 171,
										EndLine:   171,
										StartPos:  3848,
										EndPos:    3851,
									},
									Value: "Bar",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 171,
							EndLine:   171,
							StartPos:  3853,
							EndPos:    3856,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 171,
								EndLine:   171,
								StartPos:  3853,
								EndPos:    3856,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 171,
										EndLine:   171,
										StartPos:  3853,
										EndPos:    3856,
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
					StartLine: 172,
					EndLine:   172,
					StartPos:  3861,
					EndPos:    3888,
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 172,
						EndLine:   172,
						StartPos:  3865,
						EndPos:    3868,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3865,
								EndPos:    3868,
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
							EndPos:    3873,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3870,
								EndPos:    3873,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 172,
										EndLine:   172,
										StartPos:  3870,
										EndPos:    3873,
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
							EndPos:    3886,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3875,
								EndPos:    3878,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 172,
										EndLine:   172,
										StartPos:  3875,
										EndPos:    3878,
									},
									Value: "Baz",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3882,
								EndPos:    3886,
							},
							Value: "quux",
						},
					},
				},
			},
			&stmt.GroupUse{
				Position: &position.Position{
					StartLine: 173,
					EndLine:   173,
					StartPos:  3891,
					EndPos:    3919,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 173,
						EndLine:   173,
						StartPos:  3895,
						EndPos:    3903,
					},
					Value: "function",
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 173,
						EndLine:   173,
						StartPos:  3904,
						EndPos:    3907,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 173,
								EndLine:   173,
								StartPos:  3904,
								EndPos:    3907,
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
							StartPos:  3909,
							EndPos:    3912,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 173,
								EndLine:   173,
								StartPos:  3909,
								EndPos:    3912,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 173,
										EndLine:   173,
										StartPos:  3909,
										EndPos:    3912,
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
							StartPos:  3914,
							EndPos:    3917,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 173,
								EndLine:   173,
								StartPos:  3914,
								EndPos:    3917,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 173,
										EndLine:   173,
										StartPos:  3914,
										EndPos:    3917,
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
					StartLine: 174,
					EndLine:   174,
					StartPos:  3922,
					EndPos:    3948,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 174,
						EndLine:   174,
						StartPos:  3926,
						EndPos:    3931,
					},
					Value: "const",
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 174,
						EndLine:   174,
						StartPos:  3933,
						EndPos:    3936,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3933,
								EndPos:    3936,
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
							StartPos:  3938,
							EndPos:    3941,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3938,
								EndPos:    3941,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 174,
										EndLine:   174,
										StartPos:  3938,
										EndPos:    3941,
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
							StartPos:  3943,
							EndPos:    3946,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3943,
								EndPos:    3946,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 174,
										EndLine:   174,
										StartPos:  3943,
										EndPos:    3946,
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
					StartPos:  3951,
					EndPos:    3985,
				},
				Prefix: &name.Name{
					Position: &position.Position{
						StartLine: 175,
						EndLine:   175,
						StartPos:  3955,
						EndPos:    3958,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3955,
								EndPos:    3958,
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
							StartPos:  3966,
							EndPos:    3969,
						},
						UseType: &node.Identifier{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3960,
								EndPos:    3965,
							},
							Value: "const",
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3966,
								EndPos:    3969,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 175,
										EndLine:   175,
										StartPos:  3966,
										EndPos:    3969,
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
							StartPos:  3980,
							EndPos:    3983,
						},
						UseType: &node.Identifier{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3971,
								EndPos:    3979,
							},
							Value: "function",
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3980,
								EndPos:    3983,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 175,
										EndLine:   175,
										StartPos:  3980,
										EndPos:    3983,
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
					StartLine: 177,
					EndLine:   177,
					StartPos:  3989,
					EndPos:    3995,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 177,
						EndLine:   177,
						StartPos:  3989,
						EndPos:    3994,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 177,
							EndLine:   177,
							StartPos:  3989,
							EndPos:    3991,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 177,
								EndLine:   177,
								StartPos:  3989,
								EndPos:    3991,
							},
							Value: "a",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 177,
							EndLine:   177,
							StartPos:  3992,
							EndPos:    3993,
						},
						Value: "1",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 178,
					EndLine:   178,
					StartPos:  3998,
					EndPos:    4007,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 178,
						EndLine:   178,
						StartPos:  3998,
						EndPos:    4006,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 178,
							EndLine:   178,
							StartPos:  3998,
							EndPos:    4003,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 178,
								EndLine:   178,
								StartPos:  3998,
								EndPos:    4000,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 178,
									EndLine:   178,
									StartPos:  3998,
									EndPos:    4000,
								},
								Value: "a",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 178,
								EndLine:   178,
								StartPos:  4001,
								EndPos:    4002,
							},
							Value: "1",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 178,
							EndLine:   178,
							StartPos:  4004,
							EndPos:    4005,
						},
						Value: "2",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 179,
					EndLine:   179,
					StartPos:  4010,
					EndPos:    4018,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 179,
						EndLine:   179,
						StartPos:  4010,
						EndPos:    4017,
					},
					Items: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 180,
					EndLine:   180,
					StartPos:  4021,
					EndPos:    4030,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 180,
						EndLine:   180,
						StartPos:  4021,
						EndPos:    4029,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 180,
								EndLine:   180,
								StartPos:  4027,
								EndPos:    4028,
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 180,
									EndLine:   180,
									StartPos:  4027,
									EndPos:    4028,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 181,
					EndLine:   181,
					StartPos:  4033,
					EndPos:    4051,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 181,
						EndLine:   181,
						StartPos:  4033,
						EndPos:    4050,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 181,
								EndLine:   181,
								StartPos:  4039,
								EndPos:    4043,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 181,
									EndLine:   181,
									StartPos:  4039,
									EndPos:    4040,
								},
								Value: "1",
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 181,
									EndLine:   181,
									StartPos:  4042,
									EndPos:    4043,
								},
								Value: "1",
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 181,
								EndLine:   181,
								StartPos:  4045,
								EndPos:    4048,
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 181,
									EndLine:   181,
									StartPos:  4045,
									EndPos:    4048,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 181,
										EndLine:   181,
										StartPos:  4046,
										EndPos:    4048,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 181,
											EndLine:   181,
											StartPos:  4046,
											EndPos:    4048,
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
					StartLine: 182,
					EndLine:   182,
					StartPos:  4054,
					EndPos:    4058,
				},
				Expr: &expr.BitwiseNot{
					Position: &position.Position{
						StartLine: 182,
						EndLine:   182,
						StartPos:  4054,
						EndPos:    4057,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 182,
							EndLine:   182,
							StartPos:  4055,
							EndPos:    4057,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 182,
								EndLine:   182,
								StartPos:  4055,
								EndPos:    4057,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 183,
					EndLine:   183,
					StartPos:  4061,
					EndPos:    4065,
				},
				Expr: &expr.BooleanNot{
					Position: &position.Position{
						StartLine: 183,
						EndLine:   183,
						StartPos:  4061,
						EndPos:    4064,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 183,
							EndLine:   183,
							StartPos:  4062,
							EndPos:    4064,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 183,
								EndLine:   183,
								StartPos:  4062,
								EndPos:    4064,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 185,
					EndLine:   185,
					StartPos:  4069,
					EndPos:    4078,
				},
				Expr: &expr.ClassConstFetch{
					Position: &position.Position{
						StartLine: 185,
						EndLine:   185,
						StartPos:  4069,
						EndPos:    4077,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 185,
							EndLine:   185,
							StartPos:  4069,
							EndPos:    4072,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 185,
									EndLine:   185,
									StartPos:  4069,
									EndPos:    4072,
								},
								Value: "Foo",
							},
						},
					},
					ConstantName: &node.Identifier{
						Position: &position.Position{
							StartLine: 185,
							EndLine:   185,
							StartPos:  4074,
							EndPos:    4077,
						},
						Value: "Bar",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 186,
					EndLine:   186,
					StartPos:  4081,
					EndPos:    4091,
				},
				Expr: &expr.ClassConstFetch{
					Position: &position.Position{
						StartLine: 186,
						EndLine:   186,
						StartPos:  4081,
						EndPos:    4090,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 186,
							EndLine:   186,
							StartPos:  4081,
							EndPos:    4085,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 186,
								EndLine:   186,
								StartPos:  4081,
								EndPos:    4085,
							},
							Value: "foo",
						},
					},
					ConstantName: &node.Identifier{
						Position: &position.Position{
							StartLine: 186,
							EndLine:   186,
							StartPos:  4087,
							EndPos:    4090,
						},
						Value: "Bar",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 187,
					EndLine:   187,
					StartPos:  4094,
					EndPos:    4104,
				},
				Expr: &expr.Clone{
					Position: &position.Position{
						StartLine: 187,
						EndLine:   187,
						StartPos:  4094,
						EndPos:    4102,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 187,
							EndLine:   187,
							StartPos:  4100,
							EndPos:    4102,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 187,
								EndLine:   187,
								StartPos:  4100,
								EndPos:    4102,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 188,
					EndLine:   188,
					StartPos:  4107,
					EndPos:    4116,
				},
				Expr: &expr.Clone{
					Position: &position.Position{
						StartLine: 188,
						EndLine:   188,
						StartPos:  4107,
						EndPos:    4115,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 188,
							EndLine:   188,
							StartPos:  4113,
							EndPos:    4115,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 188,
								EndLine:   188,
								StartPos:  4113,
								EndPos:    4115,
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
					StartPos:  4119,
					EndPos:    4132,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 189,
						EndLine:   189,
						StartPos:  4119,
						EndPos:    4131,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Stmts:         []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 190,
					EndLine:   190,
					StartPos:  4135,
					EndPos:    4169,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 190,
						EndLine:   190,
						StartPos:  4135,
						EndPos:    4168,
					},
					Static:        false,
					PhpDocComment: "",
					ReturnsRef:    false,
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 190,
								EndLine:   190,
								StartPos:  4144,
								EndPos:    4146,
							},
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 190,
									EndLine:   190,
									StartPos:  4144,
									EndPos:    4146,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 190,
										EndLine:   190,
										StartPos:  4144,
										EndPos:    4146,
									},
									Value: "a",
								},
							},
						},
						&node.Parameter{
							Position: &position.Position{
								StartLine: 190,
								EndLine:   190,
								StartPos:  4148,
								EndPos:    4150,
							},
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 190,
									EndLine:   190,
									StartPos:  4148,
									EndPos:    4150,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 190,
										EndLine:   190,
										StartPos:  4148,
										EndPos:    4150,
									},
									Value: "b",
								},
							},
						},
					},
					ClosureUse: &expr.ClosureUse{
						Position: &position.Position{
							StartLine: 190,
							EndLine:   190,
							StartPos:  4152,
							EndPos:    4165,
						},
						Uses: []node.Node{
							&expr.Variable{
								Position: &position.Position{
									StartLine: 190,
									EndLine:   190,
									StartPos:  4157,
									EndPos:    4159,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 190,
										EndLine:   190,
										StartPos:  4157,
										EndPos:    4159,
									},
									Value: "c",
								},
							},
							&expr.Reference{
								Position: &position.Position{
									StartLine: 190,
									EndLine:   190,
									StartPos:  4161,
									EndPos:    4164,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 190,
										EndLine:   190,
										StartPos:  4162,
										EndPos:    4164,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 190,
											EndLine:   190,
											StartPos:  4162,
											EndPos:    4164,
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
					StartLine: 191,
					EndLine:   191,
					StartPos:  4172,
					EndPos:    4192,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 191,
						EndLine:   191,
						StartPos:  4172,
						EndPos:    4191,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					ReturnType: &name.Name{
						Position: &position.Position{
							StartLine: 191,
							EndLine:   191,
							StartPos:  4184,
							EndPos:    4188,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 191,
									EndLine:   191,
									StartPos:  4184,
									EndPos:    4188,
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
					StartLine: 192,
					EndLine:   192,
					StartPos:  4195,
					EndPos:    4199,
				},
				Expr: &expr.ConstFetch{
					Position: &position.Position{
						StartLine: 192,
						EndLine:   192,
						StartPos:  4195,
						EndPos:    4198,
					},
					Constant: &name.Name{
						Position: &position.Position{
							StartLine: 192,
							EndLine:   192,
							StartPos:  4195,
							EndPos:    4198,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 192,
									EndLine:   192,
									StartPos:  4195,
									EndPos:    4198,
								},
								Value: "foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 193,
					EndLine:   193,
					StartPos:  4202,
					EndPos:    4216,
				},
				Expr: &expr.ConstFetch{
					Position: &position.Position{
						StartLine: 193,
						EndLine:   193,
						StartPos:  4202,
						EndPos:    4215,
					},
					Constant: &name.Relative{
						Position: &position.Position{
							StartLine: 193,
							EndLine:   193,
							StartPos:  4202,
							EndPos:    4215,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 193,
									EndLine:   193,
									StartPos:  4212,
									EndPos:    4215,
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
					StartPos:  4219,
					EndPos:    4224,
				},
				Expr: &expr.ConstFetch{
					Position: &position.Position{
						StartLine: 194,
						EndLine:   194,
						StartPos:  4219,
						EndPos:    4223,
					},
					Constant: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 194,
							EndLine:   194,
							StartPos:  4219,
							EndPos:    4223,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 194,
									EndLine:   194,
									StartPos:  4220,
									EndPos:    4223,
								},
								Value: "foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 196,
					EndLine:   196,
					StartPos:  4228,
					EndPos:    4238,
				},
				Expr: &expr.Empty{
					Position: &position.Position{
						StartLine: 196,
						EndLine:   196,
						StartPos:  4228,
						EndPos:    4237,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 196,
							EndLine:   196,
							StartPos:  4234,
							EndPos:    4236,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 196,
								EndLine:   196,
								StartPos:  4234,
								EndPos:    4236,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 197,
					EndLine:   197,
					StartPos:  4241,
					EndPos:    4245,
				},
				Expr: &expr.ErrorSuppress{
					Position: &position.Position{
						StartLine: 197,
						EndLine:   197,
						StartPos:  4241,
						EndPos:    4244,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 197,
							EndLine:   197,
							StartPos:  4242,
							EndPos:    4244,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 197,
								EndLine:   197,
								StartPos:  4242,
								EndPos:    4244,
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
					StartPos:  4248,
					EndPos:    4257,
				},
				Expr: &expr.Eval{
					Position: &position.Position{
						StartLine: 198,
						EndLine:   198,
						StartPos:  4248,
						EndPos:    4256,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 198,
							EndLine:   198,
							StartPos:  4253,
							EndPos:    4255,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 198,
								EndLine:   198,
								StartPos:  4253,
								EndPos:    4255,
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
					StartPos:  4260,
					EndPos:    4265,
				},
				Expr: &expr.Exit{
					Position: &position.Position{
						StartLine: 199,
						EndLine:   199,
						StartPos:  4260,
						EndPos:    4264,
					},
					Die: false,
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 200,
					EndLine:   200,
					StartPos:  4268,
					EndPos:    4277,
				},
				Expr: &expr.Exit{
					Position: &position.Position{
						StartLine: 200,
						EndLine:   200,
						StartPos:  4268,
						EndPos:    4276,
					},
					Die: false,
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 200,
							EndLine:   200,
							StartPos:  4273,
							EndPos:    4275,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 200,
								EndLine:   200,
								StartPos:  4273,
								EndPos:    4275,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 201,
					EndLine:   201,
					StartPos:  4280,
					EndPos:    4284,
				},
				Expr: &expr.Exit{
					Position: &position.Position{
						StartLine: 201,
						EndLine:   201,
						StartPos:  4280,
						EndPos:    4283,
					},
					Die: true,
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 202,
					EndLine:   202,
					StartPos:  4287,
					EndPos:    4295,
				},
				Expr: &expr.Exit{
					Position: &position.Position{
						StartLine: 202,
						EndLine:   202,
						StartPos:  4287,
						EndPos:    4294,
					},
					Die: true,
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 202,
							EndLine:   202,
							StartPos:  4291,
							EndPos:    4293,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 202,
								EndLine:   202,
								StartPos:  4291,
								EndPos:    4293,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 203,
					EndLine:   203,
					StartPos:  4298,
					EndPos:    4304,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 203,
						EndLine:   203,
						StartPos:  4298,
						EndPos:    4303,
					},
					Function: &name.Name{
						Position: &position.Position{
							StartLine: 203,
							EndLine:   203,
							StartPos:  4298,
							EndPos:    4301,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 203,
									EndLine:   203,
									StartPos:  4298,
									EndPos:    4301,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 203,
							EndLine:   203,
							StartPos:  4301,
							EndPos:    4303,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 204,
					EndLine:   204,
					StartPos:  4307,
					EndPos:    4323,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 204,
						EndLine:   204,
						StartPos:  4307,
						EndPos:    4322,
					},
					Function: &name.Relative{
						Position: &position.Position{
							StartLine: 204,
							EndLine:   204,
							StartPos:  4307,
							EndPos:    4320,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 204,
									EndLine:   204,
									StartPos:  4317,
									EndPos:    4320,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 204,
							EndLine:   204,
							StartPos:  4320,
							EndPos:    4322,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 205,
					EndLine:   205,
					StartPos:  4326,
					EndPos:    4333,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 205,
						EndLine:   205,
						StartPos:  4326,
						EndPos:    4332,
					},
					Function: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 205,
							EndLine:   205,
							StartPos:  4326,
							EndPos:    4330,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 205,
									EndLine:   205,
									StartPos:  4327,
									EndPos:    4330,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 205,
							EndLine:   205,
							StartPos:  4330,
							EndPos:    4332,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 206,
					EndLine:   206,
					StartPos:  4336,
					EndPos:    4343,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 206,
						EndLine:   206,
						StartPos:  4336,
						EndPos:    4342,
					},
					Function: &expr.Variable{
						Position: &position.Position{
							StartLine: 206,
							EndLine:   206,
							StartPos:  4336,
							EndPos:    4340,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 206,
								EndLine:   206,
								StartPos:  4336,
								EndPos:    4340,
							},
							Value: "foo",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 206,
							EndLine:   206,
							StartPos:  4340,
							EndPos:    4342,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 208,
					EndLine:   208,
					StartPos:  4347,
					EndPos:    4352,
				},
				Expr: &expr.PostDec{
					Position: &position.Position{
						StartLine: 208,
						EndLine:   208,
						StartPos:  4347,
						EndPos:    4351,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 208,
							EndLine:   208,
							StartPos:  4347,
							EndPos:    4349,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 208,
								EndLine:   208,
								StartPos:  4347,
								EndPos:    4349,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 209,
					EndLine:   209,
					StartPos:  4355,
					EndPos:    4360,
				},
				Expr: &expr.PostInc{
					Position: &position.Position{
						StartLine: 209,
						EndLine:   209,
						StartPos:  4355,
						EndPos:    4359,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 209,
							EndLine:   209,
							StartPos:  4355,
							EndPos:    4357,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 209,
								EndLine:   209,
								StartPos:  4355,
								EndPos:    4357,
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
					StartPos:  4363,
					EndPos:    4368,
				},
				Expr: &expr.PreDec{
					Position: &position.Position{
						StartLine: 210,
						EndLine:   210,
						StartPos:  4363,
						EndPos:    4367,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 210,
							EndLine:   210,
							StartPos:  4365,
							EndPos:    4367,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 210,
								EndLine:   210,
								StartPos:  4365,
								EndPos:    4367,
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
					StartPos:  4371,
					EndPos:    4376,
				},
				Expr: &expr.PreInc{
					Position: &position.Position{
						StartLine: 211,
						EndLine:   211,
						StartPos:  4371,
						EndPos:    4375,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 211,
							EndLine:   211,
							StartPos:  4373,
							EndPos:    4375,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 211,
								EndLine:   211,
								StartPos:  4373,
								EndPos:    4375,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 213,
					EndLine:   213,
					StartPos:  4380,
					EndPos:    4391,
				},
				Expr: &expr.Include{
					Position: &position.Position{
						StartLine: 213,
						EndLine:   213,
						StartPos:  4380,
						EndPos:    4390,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 213,
							EndLine:   213,
							StartPos:  4388,
							EndPos:    4390,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 213,
								EndLine:   213,
								StartPos:  4388,
								EndPos:    4390,
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
					StartPos:  4394,
					EndPos:    4410,
				},
				Expr: &expr.IncludeOnce{
					Position: &position.Position{
						StartLine: 214,
						EndLine:   214,
						StartPos:  4394,
						EndPos:    4409,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 214,
							EndLine:   214,
							StartPos:  4407,
							EndPos:    4409,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 214,
								EndLine:   214,
								StartPos:  4407,
								EndPos:    4409,
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
					StartPos:  4413,
					EndPos:    4424,
				},
				Expr: &expr.Require{
					Position: &position.Position{
						StartLine: 215,
						EndLine:   215,
						StartPos:  4413,
						EndPos:    4423,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 215,
							EndLine:   215,
							StartPos:  4421,
							EndPos:    4423,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 215,
								EndLine:   215,
								StartPos:  4421,
								EndPos:    4423,
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
					StartPos:  4427,
					EndPos:    4443,
				},
				Expr: &expr.RequireOnce{
					Position: &position.Position{
						StartLine: 216,
						EndLine:   216,
						StartPos:  4427,
						EndPos:    4442,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 216,
							EndLine:   216,
							StartPos:  4440,
							EndPos:    4442,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 216,
								EndLine:   216,
								StartPos:  4440,
								EndPos:    4442,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 218,
					EndLine:   218,
					StartPos:  4447,
					EndPos:    4465,
				},
				Expr: &expr.InstanceOf{
					Position: &position.Position{
						StartLine: 218,
						EndLine:   218,
						StartPos:  4447,
						EndPos:    4464,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 218,
							EndLine:   218,
							StartPos:  4447,
							EndPos:    4449,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 218,
								EndLine:   218,
								StartPos:  4447,
								EndPos:    4449,
							},
							Value: "a",
						},
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 218,
							EndLine:   218,
							StartPos:  4461,
							EndPos:    4464,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 218,
									EndLine:   218,
									StartPos:  4461,
									EndPos:    4464,
								},
								Value: "Foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 219,
					EndLine:   219,
					StartPos:  4468,
					EndPos:    4496,
				},
				Expr: &expr.InstanceOf{
					Position: &position.Position{
						StartLine: 219,
						EndLine:   219,
						StartPos:  4468,
						EndPos:    4495,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 219,
							EndLine:   219,
							StartPos:  4468,
							EndPos:    4470,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 219,
								EndLine:   219,
								StartPos:  4468,
								EndPos:    4470,
							},
							Value: "a",
						},
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 219,
							EndLine:   219,
							StartPos:  4482,
							EndPos:    4495,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 219,
									EndLine:   219,
									StartPos:  4492,
									EndPos:    4495,
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
					StartPos:  4499,
					EndPos:    4518,
				},
				Expr: &expr.InstanceOf{
					Position: &position.Position{
						StartLine: 220,
						EndLine:   220,
						StartPos:  4499,
						EndPos:    4517,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 220,
							EndLine:   220,
							StartPos:  4499,
							EndPos:    4501,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 220,
								EndLine:   220,
								StartPos:  4499,
								EndPos:    4501,
							},
							Value: "a",
						},
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 220,
							EndLine:   220,
							StartPos:  4513,
							EndPos:    4517,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 220,
									EndLine:   220,
									StartPos:  4514,
									EndPos:    4517,
								},
								Value: "Foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 222,
					EndLine:   222,
					StartPos:  4522,
					EndPos:    4536,
				},
				Expr: &expr.Isset{
					Position: &position.Position{
						StartLine: 222,
						EndLine:   222,
						StartPos:  4522,
						EndPos:    4535,
					},
					Variables: []node.Node{
						&expr.Variable{
							Position: &position.Position{
								StartLine: 222,
								EndLine:   222,
								StartPos:  4528,
								EndPos:    4530,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 222,
									EndLine:   222,
									StartPos:  4528,
									EndPos:    4530,
								},
								Value: "a",
							},
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 222,
								EndLine:   222,
								StartPos:  4532,
								EndPos:    4534,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 222,
									EndLine:   222,
									StartPos:  4532,
									EndPos:    4534,
								},
								Value: "b",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 223,
					EndLine:   223,
					StartPos:  4539,
					EndPos:    4553,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 223,
						EndLine:   223,
						StartPos:  4539,
						EndPos:    4552,
					},
					Variable: &expr.List{
						Position: &position.Position{
							StartLine: 223,
							EndLine:   223,
							StartPos:  4539,
							EndPos:    4547,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 223,
									EndLine:   223,
									StartPos:  4544,
									EndPos:    4546,
								},
								Val: &expr.Variable{
									Position: &position.Position{
										StartLine: 223,
										EndLine:   223,
										StartPos:  4544,
										EndPos:    4546,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 223,
											EndLine:   223,
											StartPos:  4544,
											EndPos:    4546,
										},
										Value: "a",
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 223,
							EndLine:   223,
							StartPos:  4550,
							EndPos:    4552,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 223,
								EndLine:   223,
								StartPos:  4550,
								EndPos:    4552,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 224,
					EndLine:   224,
					StartPos:  4556,
					EndPos:    4572,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 224,
						EndLine:   224,
						StartPos:  4556,
						EndPos:    4571,
					},
					Variable: &expr.List{
						Position: &position.Position{
							StartLine: 224,
							EndLine:   224,
							StartPos:  4556,
							EndPos:    4566,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 224,
									EndLine:   224,
									StartPos:  4561,
									EndPos:    4565,
								},
								Val: &expr.ArrayDimFetch{
									Position: &position.Position{
										StartLine: 224,
										EndLine:   224,
										StartPos:  4561,
										EndPos:    4565,
									},
									Variable: &expr.Variable{
										Position: &position.Position{
											StartLine: 224,
											EndLine:   224,
											StartPos:  4561,
											EndPos:    4563,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 224,
												EndLine:   224,
												StartPos:  4561,
												EndPos:    4563,
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
							StartLine: 224,
							EndLine:   224,
							StartPos:  4569,
							EndPos:    4571,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 224,
								EndLine:   224,
								StartPos:  4569,
								EndPos:    4571,
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
					StartPos:  4575,
					EndPos:    4595,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 225,
						EndLine:   225,
						StartPos:  4575,
						EndPos:    4594,
					},
					Variable: &expr.List{
						Position: &position.Position{
							StartLine: 225,
							EndLine:   225,
							StartPos:  4575,
							EndPos:    4589,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 225,
									EndLine:   225,
									StartPos:  4580,
									EndPos:    4588,
								},
								Val: &expr.List{
									Position: &position.Position{
										StartLine: 225,
										EndLine:   225,
										StartPos:  4580,
										EndPos:    4588,
									},
									Items: []node.Node{
										&expr.ArrayItem{
											Position: &position.Position{
												StartLine: 225,
												EndLine:   225,
												StartPos:  4585,
												EndPos:    4587,
											},
											Val: &expr.Variable{
												Position: &position.Position{
													StartLine: 225,
													EndLine:   225,
													StartPos:  4585,
													EndPos:    4587,
												},
												VarName: &node.Identifier{
													Position: &position.Position{
														StartLine: 225,
														EndLine:   225,
														StartPos:  4585,
														EndPos:    4587,
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
							StartLine: 225,
							EndLine:   225,
							StartPos:  4592,
							EndPos:    4594,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 225,
								EndLine:   225,
								StartPos:  4592,
								EndPos:    4594,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 227,
					EndLine:   227,
					StartPos:  4599,
					EndPos:    4609,
				},
				Expr: &expr.MethodCall{
					Position: &position.Position{
						StartLine: 227,
						EndLine:   227,
						StartPos:  4599,
						EndPos:    4608,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 227,
							EndLine:   227,
							StartPos:  4599,
							EndPos:    4601,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 227,
								EndLine:   227,
								StartPos:  4599,
								EndPos:    4601,
							},
							Value: "a",
						},
					},
					Method: &node.Identifier{
						Position: &position.Position{
							StartLine: 227,
							EndLine:   227,
							StartPos:  4603,
							EndPos:    4606,
						},
						Value: "foo",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 227,
							EndLine:   227,
							StartPos:  4606,
							EndPos:    4608,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 228,
					EndLine:   228,
					StartPos:  4612,
					EndPos:    4622,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 228,
						EndLine:   228,
						StartPos:  4612,
						EndPos:    4621,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 228,
							EndLine:   228,
							StartPos:  4616,
							EndPos:    4619,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 228,
									EndLine:   228,
									StartPos:  4616,
									EndPos:    4619,
								},
								Value: "Foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 228,
							EndLine:   228,
							StartPos:  4619,
							EndPos:    4621,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 229,
					EndLine:   229,
					StartPos:  4625,
					EndPos:    4645,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 229,
						EndLine:   229,
						StartPos:  4625,
						EndPos:    4644,
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 229,
							EndLine:   229,
							StartPos:  4629,
							EndPos:    4642,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 229,
									EndLine:   229,
									StartPos:  4639,
									EndPos:    4642,
								},
								Value: "Foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 229,
							EndLine:   229,
							StartPos:  4642,
							EndPos:    4644,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 230,
					EndLine:   230,
					StartPos:  4648,
					EndPos:    4659,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 230,
						EndLine:   230,
						StartPos:  4648,
						EndPos:    4658,
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 230,
							EndLine:   230,
							StartPos:  4652,
							EndPos:    4656,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 230,
									EndLine:   230,
									StartPos:  4653,
									EndPos:    4656,
								},
								Value: "Foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 230,
							EndLine:   230,
							StartPos:  4656,
							EndPos:    4658,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 231,
					EndLine:   231,
					StartPos:  4662,
					EndPos:    4687,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 231,
						EndLine:   231,
						StartPos:  4662,
						EndPos:    4686,
					},
					Class: &stmt.Class{
						Position: &position.Position{
							StartLine: 231,
							EndLine:   231,
							StartPos:  4666,
							EndPos:    4686,
						},
						PhpDocComment: "",
						ArgumentList: &node.ArgumentList{
							Position: &position.Position{
								StartLine: 231,
								EndLine:   231,
								StartPos:  4672,
								EndPos:    4683,
							},
							Arguments: []node.Node{
								&node.Argument{
									Position: &position.Position{
										StartLine: 231,
										EndLine:   231,
										StartPos:  4673,
										EndPos:    4675,
									},
									IsReference: false,
									Variadic:    false,
									Expr: &expr.Variable{
										Position: &position.Position{
											StartLine: 231,
											EndLine:   231,
											StartPos:  4673,
											EndPos:    4675,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 231,
												EndLine:   231,
												StartPos:  4673,
												EndPos:    4675,
											},
											Value: "a",
										},
									},
								},
								&node.Argument{
									Position: &position.Position{
										StartLine: 231,
										EndLine:   231,
										StartPos:  4677,
										EndPos:    4682,
									},
									IsReference: false,
									Variadic:    true,
									Expr: &expr.Variable{
										Position: &position.Position{
											StartLine: 231,
											EndLine:   231,
											StartPos:  4680,
											EndPos:    4682,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 231,
												EndLine:   231,
												StartPos:  4680,
												EndPos:    4682,
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
					StartLine: 232,
					EndLine:   232,
					StartPos:  4690,
					EndPos:    4700,
				},
				Expr: &expr.Print{
					Position: &position.Position{
						StartLine: 232,
						EndLine:   232,
						StartPos:  4690,
						EndPos:    4698,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 232,
							EndLine:   232,
							StartPos:  4696,
							EndPos:    4698,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 232,
								EndLine:   232,
								StartPos:  4696,
								EndPos:    4698,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 233,
					EndLine:   233,
					StartPos:  4703,
					EndPos:    4711,
				},
				Expr: &expr.PropertyFetch{
					Position: &position.Position{
						StartLine: 233,
						EndLine:   233,
						StartPos:  4703,
						EndPos:    4710,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 233,
							EndLine:   233,
							StartPos:  4703,
							EndPos:    4705,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 233,
								EndLine:   233,
								StartPos:  4703,
								EndPos:    4705,
							},
							Value: "a",
						},
					},
					Property: &node.Identifier{
						Position: &position.Position{
							StartLine: 233,
							EndLine:   233,
							StartPos:  4707,
							EndPos:    4710,
						},
						Value: "foo",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 234,
					EndLine:   234,
					StartPos:  4714,
					EndPos:    4723,
				},
				Expr: &expr.ShellExec{
					Position: &position.Position{
						StartLine: 234,
						EndLine:   234,
						StartPos:  4714,
						EndPos:    4722,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 234,
								EndLine:   234,
								StartPos:  4715,
								EndPos:    4719,
							},
							Value: "cmd ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 234,
								EndLine:   234,
								StartPos:  4719,
								EndPos:    4721,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 234,
									EndLine:   234,
									StartPos:  4719,
									EndPos:    4721,
								},
								Value: "a",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 235,
					EndLine:   235,
					StartPos:  4726,
					EndPos:    4732,
				},
				Expr: &expr.ShellExec{
					Position: &position.Position{
						StartLine: 235,
						EndLine:   235,
						StartPos:  4726,
						EndPos:    4731,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 235,
								EndLine:   235,
								StartPos:  4727,
								EndPos:    4730,
							},
							Value: "cmd",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 236,
					EndLine:   236,
					StartPos:  4735,
					EndPos:    4738,
				},
				Expr: &expr.ShellExec{
					Position: &position.Position{
						StartLine: 236,
						EndLine:   236,
						StartPos:  4735,
						EndPos:    4737,
					},
					Parts: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 237,
					EndLine:   237,
					StartPos:  4741,
					EndPos:    4744,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 237,
						EndLine:   237,
						StartPos:  4741,
						EndPos:    4743,
					},
					Items: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 238,
					EndLine:   238,
					StartPos:  4747,
					EndPos:    4751,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 238,
						EndLine:   238,
						StartPos:  4747,
						EndPos:    4750,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 238,
								EndLine:   238,
								StartPos:  4748,
								EndPos:    4749,
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 238,
									EndLine:   238,
									StartPos:  4748,
									EndPos:    4749,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 239,
					EndLine:   239,
					StartPos:  4754,
					EndPos:    4767,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 239,
						EndLine:   239,
						StartPos:  4754,
						EndPos:    4766,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 239,
								EndLine:   239,
								StartPos:  4755,
								EndPos:    4759,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 239,
									EndLine:   239,
									StartPos:  4755,
									EndPos:    4756,
								},
								Value: "1",
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 239,
									EndLine:   239,
									StartPos:  4758,
									EndPos:    4759,
								},
								Value: "1",
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 239,
								EndLine:   239,
								StartPos:  4761,
								EndPos:    4764,
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 239,
									EndLine:   239,
									StartPos:  4761,
									EndPos:    4764,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 239,
										EndLine:   239,
										StartPos:  4762,
										EndPos:    4764,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 239,
											EndLine:   239,
											StartPos:  4762,
											EndPos:    4764,
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
					StartLine: 241,
					EndLine:   241,
					StartPos:  4771,
					EndPos:    4781,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 241,
						EndLine:   241,
						StartPos:  4771,
						EndPos:    4780,
					},
					Variable: &expr.ShortList{
						Position: &position.Position{
							StartLine: 241,
							EndLine:   241,
							StartPos:  4771,
							EndPos:    4775,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 241,
									EndLine:   241,
									StartPos:  4772,
									EndPos:    4774,
								},
								Val: &expr.Variable{
									Position: &position.Position{
										StartLine: 241,
										EndLine:   241,
										StartPos:  4772,
										EndPos:    4774,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 241,
											EndLine:   241,
											StartPos:  4772,
											EndPos:    4774,
										},
										Value: "a",
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 241,
							EndLine:   241,
							StartPos:  4778,
							EndPos:    4780,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 241,
								EndLine:   241,
								StartPos:  4778,
								EndPos:    4780,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 242,
					EndLine:   242,
					StartPos:  4784,
					EndPos:    4796,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 242,
						EndLine:   242,
						StartPos:  4784,
						EndPos:    4795,
					},
					Variable: &expr.ShortList{
						Position: &position.Position{
							StartLine: 242,
							EndLine:   242,
							StartPos:  4784,
							EndPos:    4790,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 242,
									EndLine:   242,
									StartPos:  4785,
									EndPos:    4789,
								},
								Val: &expr.ArrayDimFetch{
									Position: &position.Position{
										StartLine: 242,
										EndLine:   242,
										StartPos:  4785,
										EndPos:    4789,
									},
									Variable: &expr.Variable{
										Position: &position.Position{
											StartLine: 242,
											EndLine:   242,
											StartPos:  4785,
											EndPos:    4787,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 242,
												EndLine:   242,
												StartPos:  4785,
												EndPos:    4787,
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
							StartLine: 242,
							EndLine:   242,
							StartPos:  4793,
							EndPos:    4795,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 242,
								EndLine:   242,
								StartPos:  4793,
								EndPos:    4795,
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
					StartPos:  4799,
					EndPos:    4815,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 243,
						EndLine:   243,
						StartPos:  4799,
						EndPos:    4814,
					},
					Variable: &expr.ShortList{
						Position: &position.Position{
							StartLine: 243,
							EndLine:   243,
							StartPos:  4799,
							EndPos:    4809,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 243,
									EndLine:   243,
									StartPos:  4800,
									EndPos:    4808,
								},
								Val: &expr.List{
									Position: &position.Position{
										StartLine: 243,
										EndLine:   243,
										StartPos:  4800,
										EndPos:    4808,
									},
									Items: []node.Node{
										&expr.ArrayItem{
											Position: &position.Position{
												StartLine: 243,
												EndLine:   243,
												StartPos:  4805,
												EndPos:    4807,
											},
											Val: &expr.Variable{
												Position: &position.Position{
													StartLine: 243,
													EndLine:   243,
													StartPos:  4805,
													EndPos:    4807,
												},
												VarName: &node.Identifier{
													Position: &position.Position{
														StartLine: 243,
														EndLine:   243,
														StartPos:  4805,
														EndPos:    4807,
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
							StartLine: 243,
							EndLine:   243,
							StartPos:  4812,
							EndPos:    4814,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 243,
								EndLine:   243,
								StartPos:  4812,
								EndPos:    4814,
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
					StartPos:  4818,
					EndPos:    4829,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 244,
						EndLine:   244,
						StartPos:  4818,
						EndPos:    4828,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 244,
							EndLine:   244,
							StartPos:  4818,
							EndPos:    4821,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 244,
									EndLine:   244,
									StartPos:  4818,
									EndPos:    4821,
								},
								Value: "Foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 244,
							EndLine:   244,
							StartPos:  4823,
							EndPos:    4826,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 244,
							EndLine:   244,
							StartPos:  4826,
							EndPos:    4828,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 245,
					EndLine:   245,
					StartPos:  4832,
					EndPos:    4853,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 245,
						EndLine:   245,
						StartPos:  4832,
						EndPos:    4852,
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4832,
							EndPos:    4845,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4842,
									EndPos:    4845,
								},
								Value: "Foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4847,
							EndPos:    4850,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4850,
							EndPos:    4852,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 246,
					EndLine:   246,
					StartPos:  4856,
					EndPos:    4868,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 246,
						EndLine:   246,
						StartPos:  4856,
						EndPos:    4867,
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4856,
							EndPos:    4860,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4857,
									EndPos:    4860,
								},
								Value: "Foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4862,
							EndPos:    4865,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4865,
							EndPos:    4867,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 247,
					EndLine:   247,
					StartPos:  4871,
					EndPos:    4881,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 247,
						EndLine:   247,
						StartPos:  4871,
						EndPos:    4880,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 247,
							EndLine:   247,
							StartPos:  4871,
							EndPos:    4874,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 247,
									EndLine:   247,
									StartPos:  4871,
									EndPos:    4874,
								},
								Value: "Foo",
							},
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 247,
							EndLine:   247,
							StartPos:  4876,
							EndPos:    4880,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 247,
								EndLine:   247,
								StartPos:  4876,
								EndPos:    4880,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 248,
					EndLine:   248,
					StartPos:  4884,
					EndPos:    4895,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 248,
						EndLine:   248,
						StartPos:  4884,
						EndPos:    4894,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 248,
							EndLine:   248,
							StartPos:  4884,
							EndPos:    4888,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4884,
								EndPos:    4888,
							},
							Value: "foo",
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 248,
							EndLine:   248,
							StartPos:  4890,
							EndPos:    4894,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4890,
								EndPos:    4894,
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
					StartPos:  4898,
					EndPos:    4918,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 249,
						EndLine:   249,
						StartPos:  4898,
						EndPos:    4917,
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 249,
							EndLine:   249,
							StartPos:  4898,
							EndPos:    4911,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 249,
									EndLine:   249,
									StartPos:  4908,
									EndPos:    4911,
								},
								Value: "Foo",
							},
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 249,
							EndLine:   249,
							StartPos:  4913,
							EndPos:    4917,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 249,
								EndLine:   249,
								StartPos:  4913,
								EndPos:    4917,
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
					StartPos:  4921,
					EndPos:    4932,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 250,
						EndLine:   250,
						StartPos:  4921,
						EndPos:    4931,
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 250,
							EndLine:   250,
							StartPos:  4921,
							EndPos:    4925,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 250,
									EndLine:   250,
									StartPos:  4922,
									EndPos:    4925,
								},
								Value: "Foo",
							},
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 250,
							EndLine:   250,
							StartPos:  4927,
							EndPos:    4931,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 250,
								EndLine:   250,
								StartPos:  4927,
								EndPos:    4931,
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
					StartPos:  4935,
					EndPos:    4948,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 251,
						EndLine:   251,
						StartPos:  4935,
						EndPos:    4947,
					},
					Condition: &expr.Variable{
						Position: &position.Position{
							StartLine: 251,
							EndLine:   251,
							StartPos:  4935,
							EndPos:    4937,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 251,
								EndLine:   251,
								StartPos:  4935,
								EndPos:    4937,
							},
							Value: "a",
						},
					},
					IfTrue: &expr.Variable{
						Position: &position.Position{
							StartLine: 251,
							EndLine:   251,
							StartPos:  4940,
							EndPos:    4942,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 251,
								EndLine:   251,
								StartPos:  4940,
								EndPos:    4942,
							},
							Value: "b",
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 251,
							EndLine:   251,
							StartPos:  4945,
							EndPos:    4947,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 251,
								EndLine:   251,
								StartPos:  4945,
								EndPos:    4947,
							},
							Value: "c",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 252,
					EndLine:   252,
					StartPos:  4951,
					EndPos:    4961,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 252,
						EndLine:   252,
						StartPos:  4951,
						EndPos:    4960,
					},
					Condition: &expr.Variable{
						Position: &position.Position{
							StartLine: 252,
							EndLine:   252,
							StartPos:  4951,
							EndPos:    4953,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 252,
								EndLine:   252,
								StartPos:  4951,
								EndPos:    4953,
							},
							Value: "a",
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 252,
							EndLine:   252,
							StartPos:  4958,
							EndPos:    4960,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 252,
								EndLine:   252,
								StartPos:  4958,
								EndPos:    4960,
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
					StartPos:  4964,
					EndPos:    4987,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 253,
						EndLine:   253,
						StartPos:  4964,
						EndPos:    4986,
					},
					Condition: &expr.Variable{
						Position: &position.Position{
							StartLine: 253,
							EndLine:   253,
							StartPos:  4964,
							EndPos:    4966,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4964,
								EndPos:    4966,
							},
							Value: "a",
						},
					},
					IfTrue: &expr.Ternary{
						Position: &position.Position{
							StartLine: 253,
							EndLine:   253,
							StartPos:  4969,
							EndPos:    4981,
						},
						Condition: &expr.Variable{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4969,
								EndPos:    4971,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4969,
									EndPos:    4971,
								},
								Value: "b",
							},
						},
						IfTrue: &expr.Variable{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4974,
								EndPos:    4976,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4974,
									EndPos:    4976,
								},
								Value: "c",
							},
						},
						IfFalse: &expr.Variable{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4979,
								EndPos:    4981,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4979,
									EndPos:    4981,
								},
								Value: "d",
							},
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 253,
							EndLine:   253,
							StartPos:  4984,
							EndPos:    4986,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4984,
								EndPos:    4986,
							},
							Value: "e",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 254,
					EndLine:   254,
					StartPos:  4990,
					EndPos:    5013,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 254,
						EndLine:   254,
						StartPos:  4990,
						EndPos:    5012,
					},
					Condition: &expr.Ternary{
						Position: &position.Position{
							StartLine: 254,
							EndLine:   254,
							StartPos:  4990,
							EndPos:    5002,
						},
						Condition: &expr.Variable{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  4990,
								EndPos:    4992,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4990,
									EndPos:    4992,
								},
								Value: "a",
							},
						},
						IfTrue: &expr.Variable{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  4995,
								EndPos:    4997,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4995,
									EndPos:    4997,
								},
								Value: "b",
							},
						},
						IfFalse: &expr.Variable{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  5000,
								EndPos:    5002,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  5000,
									EndPos:    5002,
								},
								Value: "c",
							},
						},
					},
					IfTrue: &expr.Variable{
						Position: &position.Position{
							StartLine: 254,
							EndLine:   254,
							StartPos:  5005,
							EndPos:    5007,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  5005,
								EndPos:    5007,
							},
							Value: "d",
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 254,
							EndLine:   254,
							StartPos:  5010,
							EndPos:    5012,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  5010,
								EndPos:    5012,
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
					StartPos:  5016,
					EndPos:    5020,
				},
				Expr: &expr.UnaryMinus{
					Position: &position.Position{
						StartLine: 255,
						EndLine:   255,
						StartPos:  5016,
						EndPos:    5019,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 255,
							EndLine:   255,
							StartPos:  5017,
							EndPos:    5019,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  5017,
								EndPos:    5019,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 256,
					EndLine:   256,
					StartPos:  5023,
					EndPos:    5027,
				},
				Expr: &expr.UnaryPlus{
					Position: &position.Position{
						StartLine: 256,
						EndLine:   256,
						StartPos:  5023,
						EndPos:    5026,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 256,
							EndLine:   256,
							StartPos:  5024,
							EndPos:    5026,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 256,
								EndLine:   256,
								StartPos:  5024,
								EndPos:    5026,
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
					StartPos:  5030,
					EndPos:    5034,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 257,
						EndLine:   257,
						StartPos:  5030,
						EndPos:    5033,
					},
					VarName: &expr.Variable{
						Position: &position.Position{
							StartLine: 257,
							EndLine:   257,
							StartPos:  5031,
							EndPos:    5033,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 257,
								EndLine:   257,
								StartPos:  5031,
								EndPos:    5033,
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
					StartPos:  5037,
					EndPos:    5043,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 258,
						EndLine:   258,
						StartPos:  5037,
						EndPos:    5042,
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 259,
					EndLine:   259,
					StartPos:  5046,
					EndPos:    5055,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 259,
						EndLine:   259,
						StartPos:  5046,
						EndPos:    5054,
					},
					Value: &expr.Variable{
						Position: &position.Position{
							StartLine: 259,
							EndLine:   259,
							StartPos:  5052,
							EndPos:    5054,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 259,
								EndLine:   259,
								StartPos:  5052,
								EndPos:    5054,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 260,
					EndLine:   260,
					StartPos:  5058,
					EndPos:    5073,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 260,
						EndLine:   260,
						StartPos:  5058,
						EndPos:    5072,
					},
					Key: &expr.Variable{
						Position: &position.Position{
							StartLine: 260,
							EndLine:   260,
							StartPos:  5064,
							EndPos:    5066,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 260,
								EndLine:   260,
								StartPos:  5064,
								EndPos:    5066,
							},
							Value: "a",
						},
					},
					Value: &expr.Variable{
						Position: &position.Position{
							StartLine: 260,
							EndLine:   260,
							StartPos:  5070,
							EndPos:    5072,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 260,
								EndLine:   260,
								StartPos:  5070,
								EndPos:    5072,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 261,
					EndLine:   261,
					StartPos:  5076,
					EndPos:    5090,
				},
				Expr: &expr.YieldFrom{
					Position: &position.Position{
						StartLine: 261,
						EndLine:   261,
						StartPos:  5076,
						EndPos:    5089,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 261,
							EndLine:   261,
							StartPos:  5087,
							EndPos:    5089,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 261,
								EndLine:   261,
								StartPos:  5087,
								EndPos:    5089,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 263,
					EndLine:   263,
					StartPos:  5096,
					EndPos:    5106,
				},
				Expr: &cast.Array{
					Position: &position.Position{
						StartLine: 263,
						EndLine:   263,
						StartPos:  5096,
						EndPos:    5105,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 263,
							EndLine:   263,
							StartPos:  5103,
							EndPos:    5105,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 263,
								EndLine:   263,
								StartPos:  5103,
								EndPos:    5105,
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
					StartPos:  5109,
					EndPos:    5121,
				},
				Expr: &cast.Bool{
					Position: &position.Position{
						StartLine: 264,
						EndLine:   264,
						StartPos:  5109,
						EndPos:    5120,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 264,
							EndLine:   264,
							StartPos:  5118,
							EndPos:    5120,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 264,
								EndLine:   264,
								StartPos:  5118,
								EndPos:    5120,
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
					StartPos:  5124,
					EndPos:    5133,
				},
				Expr: &cast.Bool{
					Position: &position.Position{
						StartLine: 265,
						EndLine:   265,
						StartPos:  5124,
						EndPos:    5132,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 265,
							EndLine:   265,
							StartPos:  5130,
							EndPos:    5132,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 265,
								EndLine:   265,
								StartPos:  5130,
								EndPos:    5132,
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
					StartPos:  5136,
					EndPos:    5147,
				},
				Expr: &cast.Double{
					Position: &position.Position{
						StartLine: 266,
						EndLine:   266,
						StartPos:  5136,
						EndPos:    5146,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 266,
							EndLine:   266,
							StartPos:  5144,
							EndPos:    5146,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 266,
								EndLine:   266,
								StartPos:  5144,
								EndPos:    5146,
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
					StartPos:  5150,
					EndPos:    5160,
				},
				Expr: &cast.Double{
					Position: &position.Position{
						StartLine: 267,
						EndLine:   267,
						StartPos:  5150,
						EndPos:    5159,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 267,
							EndLine:   267,
							StartPos:  5157,
							EndPos:    5159,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 267,
								EndLine:   267,
								StartPos:  5157,
								EndPos:    5159,
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
					StartPos:  5163,
					EndPos:    5175,
				},
				Expr: &cast.Int{
					Position: &position.Position{
						StartLine: 268,
						EndLine:   268,
						StartPos:  5163,
						EndPos:    5174,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 268,
							EndLine:   268,
							StartPos:  5172,
							EndPos:    5174,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 268,
								EndLine:   268,
								StartPos:  5172,
								EndPos:    5174,
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
					StartPos:  5178,
					EndPos:    5186,
				},
				Expr: &cast.Int{
					Position: &position.Position{
						StartLine: 269,
						EndLine:   269,
						StartPos:  5178,
						EndPos:    5185,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 269,
							EndLine:   269,
							StartPos:  5183,
							EndPos:    5185,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 269,
								EndLine:   269,
								StartPos:  5183,
								EndPos:    5185,
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
					StartPos:  5189,
					EndPos:    5200,
				},
				Expr: &cast.Object{
					Position: &position.Position{
						StartLine: 270,
						EndLine:   270,
						StartPos:  5189,
						EndPos:    5199,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 270,
							EndLine:   270,
							StartPos:  5197,
							EndPos:    5199,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 270,
								EndLine:   270,
								StartPos:  5197,
								EndPos:    5199,
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
					StartPos:  5203,
					EndPos:    5214,
				},
				Expr: &cast.String{
					Position: &position.Position{
						StartLine: 271,
						EndLine:   271,
						StartPos:  5203,
						EndPos:    5213,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 271,
							EndLine:   271,
							StartPos:  5211,
							EndPos:    5213,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 271,
								EndLine:   271,
								StartPos:  5211,
								EndPos:    5213,
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
					StartPos:  5217,
					EndPos:    5227,
				},
				Expr: &cast.Unset{
					Position: &position.Position{
						StartLine: 272,
						EndLine:   272,
						StartPos:  5217,
						EndPos:    5226,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 272,
							EndLine:   272,
							StartPos:  5224,
							EndPos:    5226,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 272,
								EndLine:   272,
								StartPos:  5224,
								EndPos:    5226,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 274,
					EndLine:   274,
					StartPos:  5231,
					EndPos:    5239,
				},
				Expr: &binary.BitwiseAnd{
					Position: &position.Position{
						StartLine: 274,
						EndLine:   274,
						StartPos:  5231,
						EndPos:    5238,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 274,
							EndLine:   274,
							StartPos:  5231,
							EndPos:    5233,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 274,
								EndLine:   274,
								StartPos:  5231,
								EndPos:    5233,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 274,
							EndLine:   274,
							StartPos:  5236,
							EndPos:    5238,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 274,
								EndLine:   274,
								StartPos:  5236,
								EndPos:    5238,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 275,
					EndLine:   275,
					StartPos:  5242,
					EndPos:    5250,
				},
				Expr: &binary.BitwiseOr{
					Position: &position.Position{
						StartLine: 275,
						EndLine:   275,
						StartPos:  5242,
						EndPos:    5249,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 275,
							EndLine:   275,
							StartPos:  5242,
							EndPos:    5244,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 275,
								EndLine:   275,
								StartPos:  5242,
								EndPos:    5244,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 275,
							EndLine:   275,
							StartPos:  5247,
							EndPos:    5249,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 275,
								EndLine:   275,
								StartPos:  5247,
								EndPos:    5249,
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
					StartPos:  5253,
					EndPos:    5261,
				},
				Expr: &binary.BitwiseXor{
					Position: &position.Position{
						StartLine: 276,
						EndLine:   276,
						StartPos:  5253,
						EndPos:    5260,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 276,
							EndLine:   276,
							StartPos:  5253,
							EndPos:    5255,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 276,
								EndLine:   276,
								StartPos:  5253,
								EndPos:    5255,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 276,
							EndLine:   276,
							StartPos:  5258,
							EndPos:    5260,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 276,
								EndLine:   276,
								StartPos:  5258,
								EndPos:    5260,
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
					StartPos:  5264,
					EndPos:    5273,
				},
				Expr: &binary.BooleanAnd{
					Position: &position.Position{
						StartLine: 277,
						EndLine:   277,
						StartPos:  5264,
						EndPos:    5272,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 277,
							EndLine:   277,
							StartPos:  5264,
							EndPos:    5266,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 277,
								EndLine:   277,
								StartPos:  5264,
								EndPos:    5266,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 277,
							EndLine:   277,
							StartPos:  5270,
							EndPos:    5272,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 277,
								EndLine:   277,
								StartPos:  5270,
								EndPos:    5272,
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
					StartPos:  5276,
					EndPos:    5285,
				},
				Expr: &binary.BooleanOr{
					Position: &position.Position{
						StartLine: 278,
						EndLine:   278,
						StartPos:  5276,
						EndPos:    5284,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 278,
							EndLine:   278,
							StartPos:  5276,
							EndPos:    5278,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 278,
								EndLine:   278,
								StartPos:  5276,
								EndPos:    5278,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 278,
							EndLine:   278,
							StartPos:  5282,
							EndPos:    5284,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 278,
								EndLine:   278,
								StartPos:  5282,
								EndPos:    5284,
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
					StartPos:  5288,
					EndPos:    5297,
				},
				Expr: &binary.Coalesce{
					Position: &position.Position{
						StartLine: 279,
						EndLine:   279,
						StartPos:  5288,
						EndPos:    5296,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 279,
							EndLine:   279,
							StartPos:  5288,
							EndPos:    5290,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 279,
								EndLine:   279,
								StartPos:  5288,
								EndPos:    5290,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 279,
							EndLine:   279,
							StartPos:  5294,
							EndPos:    5296,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 279,
								EndLine:   279,
								StartPos:  5294,
								EndPos:    5296,
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
					StartPos:  5300,
					EndPos:    5308,
				},
				Expr: &binary.Concat{
					Position: &position.Position{
						StartLine: 280,
						EndLine:   280,
						StartPos:  5300,
						EndPos:    5307,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 280,
							EndLine:   280,
							StartPos:  5300,
							EndPos:    5302,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 280,
								EndLine:   280,
								StartPos:  5300,
								EndPos:    5302,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 280,
							EndLine:   280,
							StartPos:  5305,
							EndPos:    5307,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 280,
								EndLine:   280,
								StartPos:  5305,
								EndPos:    5307,
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
					StartPos:  5311,
					EndPos:    5319,
				},
				Expr: &binary.Div{
					Position: &position.Position{
						StartLine: 281,
						EndLine:   281,
						StartPos:  5311,
						EndPos:    5318,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 281,
							EndLine:   281,
							StartPos:  5311,
							EndPos:    5313,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 281,
								EndLine:   281,
								StartPos:  5311,
								EndPos:    5313,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 281,
							EndLine:   281,
							StartPos:  5316,
							EndPos:    5318,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 281,
								EndLine:   281,
								StartPos:  5316,
								EndPos:    5318,
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
					StartPos:  5322,
					EndPos:    5331,
				},
				Expr: &binary.Equal{
					Position: &position.Position{
						StartLine: 282,
						EndLine:   282,
						StartPos:  5322,
						EndPos:    5330,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 282,
							EndLine:   282,
							StartPos:  5322,
							EndPos:    5324,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 282,
								EndLine:   282,
								StartPos:  5322,
								EndPos:    5324,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 282,
							EndLine:   282,
							StartPos:  5328,
							EndPos:    5330,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 282,
								EndLine:   282,
								StartPos:  5328,
								EndPos:    5330,
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
					StartPos:  5334,
					EndPos:    5343,
				},
				Expr: &binary.GreaterOrEqual{
					Position: &position.Position{
						StartLine: 283,
						EndLine:   283,
						StartPos:  5334,
						EndPos:    5342,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 283,
							EndLine:   283,
							StartPos:  5334,
							EndPos:    5336,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 283,
								EndLine:   283,
								StartPos:  5334,
								EndPos:    5336,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 283,
							EndLine:   283,
							StartPos:  5340,
							EndPos:    5342,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 283,
								EndLine:   283,
								StartPos:  5340,
								EndPos:    5342,
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
					StartPos:  5346,
					EndPos:    5354,
				},
				Expr: &binary.Greater{
					Position: &position.Position{
						StartLine: 284,
						EndLine:   284,
						StartPos:  5346,
						EndPos:    5353,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 284,
							EndLine:   284,
							StartPos:  5346,
							EndPos:    5348,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 284,
								EndLine:   284,
								StartPos:  5346,
								EndPos:    5348,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 284,
							EndLine:   284,
							StartPos:  5351,
							EndPos:    5353,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 284,
								EndLine:   284,
								StartPos:  5351,
								EndPos:    5353,
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
					StartPos:  5357,
					EndPos:    5367,
				},
				Expr: &binary.Identical{
					Position: &position.Position{
						StartLine: 285,
						EndLine:   285,
						StartPos:  5357,
						EndPos:    5366,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 285,
							EndLine:   285,
							StartPos:  5357,
							EndPos:    5359,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 285,
								EndLine:   285,
								StartPos:  5357,
								EndPos:    5359,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 285,
							EndLine:   285,
							StartPos:  5364,
							EndPos:    5366,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 285,
								EndLine:   285,
								StartPos:  5364,
								EndPos:    5366,
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
					StartPos:  5370,
					EndPos:    5380,
				},
				Expr: &binary.LogicalAnd{
					Position: &position.Position{
						StartLine: 286,
						EndLine:   286,
						StartPos:  5370,
						EndPos:    5379,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 286,
							EndLine:   286,
							StartPos:  5370,
							EndPos:    5372,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 286,
								EndLine:   286,
								StartPos:  5370,
								EndPos:    5372,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 286,
							EndLine:   286,
							StartPos:  5377,
							EndPos:    5379,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 286,
								EndLine:   286,
								StartPos:  5377,
								EndPos:    5379,
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
					StartPos:  5383,
					EndPos:    5392,
				},
				Expr: &binary.LogicalOr{
					Position: &position.Position{
						StartLine: 287,
						EndLine:   287,
						StartPos:  5383,
						EndPos:    5391,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 287,
							EndLine:   287,
							StartPos:  5383,
							EndPos:    5385,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 287,
								EndLine:   287,
								StartPos:  5383,
								EndPos:    5385,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 287,
							EndLine:   287,
							StartPos:  5389,
							EndPos:    5391,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 287,
								EndLine:   287,
								StartPos:  5389,
								EndPos:    5391,
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
					StartPos:  5395,
					EndPos:    5405,
				},
				Expr: &binary.LogicalXor{
					Position: &position.Position{
						StartLine: 288,
						EndLine:   288,
						StartPos:  5395,
						EndPos:    5404,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 288,
							EndLine:   288,
							StartPos:  5395,
							EndPos:    5397,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 288,
								EndLine:   288,
								StartPos:  5395,
								EndPos:    5397,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 288,
							EndLine:   288,
							StartPos:  5402,
							EndPos:    5404,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 288,
								EndLine:   288,
								StartPos:  5402,
								EndPos:    5404,
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
					StartPos:  5408,
					EndPos:    5416,
				},
				Expr: &binary.Minus{
					Position: &position.Position{
						StartLine: 289,
						EndLine:   289,
						StartPos:  5408,
						EndPos:    5415,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 289,
							EndLine:   289,
							StartPos:  5408,
							EndPos:    5410,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 289,
								EndLine:   289,
								StartPos:  5408,
								EndPos:    5410,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 289,
							EndLine:   289,
							StartPos:  5413,
							EndPos:    5415,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 289,
								EndLine:   289,
								StartPos:  5413,
								EndPos:    5415,
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
					StartPos:  5419,
					EndPos:    5427,
				},
				Expr: &binary.Mod{
					Position: &position.Position{
						StartLine: 290,
						EndLine:   290,
						StartPos:  5419,
						EndPos:    5426,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 290,
							EndLine:   290,
							StartPos:  5419,
							EndPos:    5421,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 290,
								EndLine:   290,
								StartPos:  5419,
								EndPos:    5421,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 290,
							EndLine:   290,
							StartPos:  5424,
							EndPos:    5426,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 290,
								EndLine:   290,
								StartPos:  5424,
								EndPos:    5426,
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
					StartPos:  5430,
					EndPos:    5438,
				},
				Expr: &binary.Mul{
					Position: &position.Position{
						StartLine: 291,
						EndLine:   291,
						StartPos:  5430,
						EndPos:    5437,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 291,
							EndLine:   291,
							StartPos:  5430,
							EndPos:    5432,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 291,
								EndLine:   291,
								StartPos:  5430,
								EndPos:    5432,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 291,
							EndLine:   291,
							StartPos:  5435,
							EndPos:    5437,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 291,
								EndLine:   291,
								StartPos:  5435,
								EndPos:    5437,
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
					StartPos:  5441,
					EndPos:    5450,
				},
				Expr: &binary.NotEqual{
					Position: &position.Position{
						StartLine: 292,
						EndLine:   292,
						StartPos:  5441,
						EndPos:    5449,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 292,
							EndLine:   292,
							StartPos:  5441,
							EndPos:    5443,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 292,
								EndLine:   292,
								StartPos:  5441,
								EndPos:    5443,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 292,
							EndLine:   292,
							StartPos:  5447,
							EndPos:    5449,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 292,
								EndLine:   292,
								StartPos:  5447,
								EndPos:    5449,
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
					StartPos:  5453,
					EndPos:    5463,
				},
				Expr: &binary.NotIdentical{
					Position: &position.Position{
						StartLine: 293,
						EndLine:   293,
						StartPos:  5453,
						EndPos:    5462,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 293,
							EndLine:   293,
							StartPos:  5453,
							EndPos:    5455,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 293,
								EndLine:   293,
								StartPos:  5453,
								EndPos:    5455,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 293,
							EndLine:   293,
							StartPos:  5460,
							EndPos:    5462,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 293,
								EndLine:   293,
								StartPos:  5460,
								EndPos:    5462,
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
					StartPos:  5466,
					EndPos:    5474,
				},
				Expr: &binary.Plus{
					Position: &position.Position{
						StartLine: 294,
						EndLine:   294,
						StartPos:  5466,
						EndPos:    5473,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 294,
							EndLine:   294,
							StartPos:  5466,
							EndPos:    5468,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  5466,
								EndPos:    5468,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 294,
							EndLine:   294,
							StartPos:  5471,
							EndPos:    5473,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  5471,
								EndPos:    5473,
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
					StartPos:  5477,
					EndPos:    5486,
				},
				Expr: &binary.Pow{
					Position: &position.Position{
						StartLine: 295,
						EndLine:   295,
						StartPos:  5477,
						EndPos:    5485,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 295,
							EndLine:   295,
							StartPos:  5477,
							EndPos:    5479,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 295,
								EndLine:   295,
								StartPos:  5477,
								EndPos:    5479,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 295,
							EndLine:   295,
							StartPos:  5483,
							EndPos:    5485,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 295,
								EndLine:   295,
								StartPos:  5483,
								EndPos:    5485,
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
					StartPos:  5489,
					EndPos:    5498,
				},
				Expr: &binary.ShiftLeft{
					Position: &position.Position{
						StartLine: 296,
						EndLine:   296,
						StartPos:  5489,
						EndPos:    5497,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 296,
							EndLine:   296,
							StartPos:  5489,
							EndPos:    5491,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  5489,
								EndPos:    5491,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 296,
							EndLine:   296,
							StartPos:  5495,
							EndPos:    5497,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  5495,
								EndPos:    5497,
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
					StartPos:  5501,
					EndPos:    5510,
				},
				Expr: &binary.ShiftRight{
					Position: &position.Position{
						StartLine: 297,
						EndLine:   297,
						StartPos:  5501,
						EndPos:    5509,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 297,
							EndLine:   297,
							StartPos:  5501,
							EndPos:    5503,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 297,
								EndLine:   297,
								StartPos:  5501,
								EndPos:    5503,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 297,
							EndLine:   297,
							StartPos:  5507,
							EndPos:    5509,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 297,
								EndLine:   297,
								StartPos:  5507,
								EndPos:    5509,
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
					StartPos:  5513,
					EndPos:    5522,
				},
				Expr: &binary.SmallerOrEqual{
					Position: &position.Position{
						StartLine: 298,
						EndLine:   298,
						StartPos:  5513,
						EndPos:    5521,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 298,
							EndLine:   298,
							StartPos:  5513,
							EndPos:    5515,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  5513,
								EndPos:    5515,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 298,
							EndLine:   298,
							StartPos:  5519,
							EndPos:    5521,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  5519,
								EndPos:    5521,
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
					StartPos:  5525,
					EndPos:    5533,
				},
				Expr: &binary.Smaller{
					Position: &position.Position{
						StartLine: 299,
						EndLine:   299,
						StartPos:  5525,
						EndPos:    5532,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 299,
							EndLine:   299,
							StartPos:  5525,
							EndPos:    5527,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5525,
								EndPos:    5527,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 299,
							EndLine:   299,
							StartPos:  5530,
							EndPos:    5532,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5530,
								EndPos:    5532,
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
					StartPos:  5536,
					EndPos:    5546,
				},
				Expr: &binary.Spaceship{
					Position: &position.Position{
						StartLine: 300,
						EndLine:   300,
						StartPos:  5536,
						EndPos:    5545,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 300,
							EndLine:   300,
							StartPos:  5536,
							EndPos:    5538,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5536,
								EndPos:    5538,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 300,
							EndLine:   300,
							StartPos:  5543,
							EndPos:    5545,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5543,
								EndPos:    5545,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 302,
					EndLine:   302,
					StartPos:  5550,
					EndPos:    5559,
				},
				Expr: &assign.Reference{
					Position: &position.Position{
						StartLine: 302,
						EndLine:   302,
						StartPos:  5550,
						EndPos:    5558,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 302,
							EndLine:   302,
							StartPos:  5550,
							EndPos:    5552,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 302,
								EndLine:   302,
								StartPos:  5550,
								EndPos:    5552,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 302,
							EndLine:   302,
							StartPos:  5556,
							EndPos:    5558,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 302,
								EndLine:   302,
								StartPos:  5556,
								EndPos:    5558,
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
					StartPos:  5562,
					EndPos:    5570,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 303,
						EndLine:   303,
						StartPos:  5562,
						EndPos:    5569,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 303,
							EndLine:   303,
							StartPos:  5562,
							EndPos:    5564,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 303,
								EndLine:   303,
								StartPos:  5562,
								EndPos:    5564,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 303,
							EndLine:   303,
							StartPos:  5567,
							EndPos:    5569,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 303,
								EndLine:   303,
								StartPos:  5567,
								EndPos:    5569,
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
					StartPos:  5573,
					EndPos:    5582,
				},
				Expr: &assign.BitwiseAnd{
					Position: &position.Position{
						StartLine: 304,
						EndLine:   304,
						StartPos:  5573,
						EndPos:    5581,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 304,
							EndLine:   304,
							StartPos:  5573,
							EndPos:    5575,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 304,
								EndLine:   304,
								StartPos:  5573,
								EndPos:    5575,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 304,
							EndLine:   304,
							StartPos:  5579,
							EndPos:    5581,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 304,
								EndLine:   304,
								StartPos:  5579,
								EndPos:    5581,
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
					StartPos:  5585,
					EndPos:    5594,
				},
				Expr: &assign.BitwiseOr{
					Position: &position.Position{
						StartLine: 305,
						EndLine:   305,
						StartPos:  5585,
						EndPos:    5593,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 305,
							EndLine:   305,
							StartPos:  5585,
							EndPos:    5587,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 305,
								EndLine:   305,
								StartPos:  5585,
								EndPos:    5587,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 305,
							EndLine:   305,
							StartPos:  5591,
							EndPos:    5593,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 305,
								EndLine:   305,
								StartPos:  5591,
								EndPos:    5593,
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
					StartPos:  5597,
					EndPos:    5606,
				},
				Expr: &assign.BitwiseXor{
					Position: &position.Position{
						StartLine: 306,
						EndLine:   306,
						StartPos:  5597,
						EndPos:    5605,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 306,
							EndLine:   306,
							StartPos:  5597,
							EndPos:    5599,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 306,
								EndLine:   306,
								StartPos:  5597,
								EndPos:    5599,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 306,
							EndLine:   306,
							StartPos:  5603,
							EndPos:    5605,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 306,
								EndLine:   306,
								StartPos:  5603,
								EndPos:    5605,
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
					StartPos:  5609,
					EndPos:    5618,
				},
				Expr: &assign.Concat{
					Position: &position.Position{
						StartLine: 307,
						EndLine:   307,
						StartPos:  5609,
						EndPos:    5617,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 307,
							EndLine:   307,
							StartPos:  5609,
							EndPos:    5611,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 307,
								EndLine:   307,
								StartPos:  5609,
								EndPos:    5611,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 307,
							EndLine:   307,
							StartPos:  5615,
							EndPos:    5617,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 307,
								EndLine:   307,
								StartPos:  5615,
								EndPos:    5617,
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
					StartPos:  5621,
					EndPos:    5630,
				},
				Expr: &assign.Div{
					Position: &position.Position{
						StartLine: 308,
						EndLine:   308,
						StartPos:  5621,
						EndPos:    5629,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 308,
							EndLine:   308,
							StartPos:  5621,
							EndPos:    5623,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 308,
								EndLine:   308,
								StartPos:  5621,
								EndPos:    5623,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 308,
							EndLine:   308,
							StartPos:  5627,
							EndPos:    5629,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 308,
								EndLine:   308,
								StartPos:  5627,
								EndPos:    5629,
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
					StartPos:  5633,
					EndPos:    5642,
				},
				Expr: &assign.Minus{
					Position: &position.Position{
						StartLine: 309,
						EndLine:   309,
						StartPos:  5633,
						EndPos:    5641,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 309,
							EndLine:   309,
							StartPos:  5633,
							EndPos:    5635,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 309,
								EndLine:   309,
								StartPos:  5633,
								EndPos:    5635,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 309,
							EndLine:   309,
							StartPos:  5639,
							EndPos:    5641,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 309,
								EndLine:   309,
								StartPos:  5639,
								EndPos:    5641,
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
					StartPos:  5645,
					EndPos:    5654,
				},
				Expr: &assign.Mod{
					Position: &position.Position{
						StartLine: 310,
						EndLine:   310,
						StartPos:  5645,
						EndPos:    5653,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 310,
							EndLine:   310,
							StartPos:  5645,
							EndPos:    5647,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 310,
								EndLine:   310,
								StartPos:  5645,
								EndPos:    5647,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 310,
							EndLine:   310,
							StartPos:  5651,
							EndPos:    5653,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 310,
								EndLine:   310,
								StartPos:  5651,
								EndPos:    5653,
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
					StartPos:  5657,
					EndPos:    5666,
				},
				Expr: &assign.Mul{
					Position: &position.Position{
						StartLine: 311,
						EndLine:   311,
						StartPos:  5657,
						EndPos:    5665,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 311,
							EndLine:   311,
							StartPos:  5657,
							EndPos:    5659,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 311,
								EndLine:   311,
								StartPos:  5657,
								EndPos:    5659,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 311,
							EndLine:   311,
							StartPos:  5663,
							EndPos:    5665,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 311,
								EndLine:   311,
								StartPos:  5663,
								EndPos:    5665,
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
					StartPos:  5669,
					EndPos:    5678,
				},
				Expr: &assign.Plus{
					Position: &position.Position{
						StartLine: 312,
						EndLine:   312,
						StartPos:  5669,
						EndPos:    5677,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 312,
							EndLine:   312,
							StartPos:  5669,
							EndPos:    5671,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 312,
								EndLine:   312,
								StartPos:  5669,
								EndPos:    5671,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 312,
							EndLine:   312,
							StartPos:  5675,
							EndPos:    5677,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 312,
								EndLine:   312,
								StartPos:  5675,
								EndPos:    5677,
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
					StartPos:  5681,
					EndPos:    5691,
				},
				Expr: &assign.Pow{
					Position: &position.Position{
						StartLine: 313,
						EndLine:   313,
						StartPos:  5681,
						EndPos:    5690,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 313,
							EndLine:   313,
							StartPos:  5681,
							EndPos:    5683,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5681,
								EndPos:    5683,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 313,
							EndLine:   313,
							StartPos:  5688,
							EndPos:    5690,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5688,
								EndPos:    5690,
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
					StartPos:  5694,
					EndPos:    5704,
				},
				Expr: &assign.ShiftLeft{
					Position: &position.Position{
						StartLine: 314,
						EndLine:   314,
						StartPos:  5694,
						EndPos:    5703,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 314,
							EndLine:   314,
							StartPos:  5694,
							EndPos:    5696,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 314,
								EndLine:   314,
								StartPos:  5694,
								EndPos:    5696,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 314,
							EndLine:   314,
							StartPos:  5701,
							EndPos:    5703,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 314,
								EndLine:   314,
								StartPos:  5701,
								EndPos:    5703,
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
					StartPos:  5707,
					EndPos:    5717,
				},
				Expr: &assign.ShiftRight{
					Position: &position.Position{
						StartLine: 315,
						EndLine:   315,
						StartPos:  5707,
						EndPos:    5716,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 315,
							EndLine:   315,
							StartPos:  5707,
							EndPos:    5709,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5707,
								EndPos:    5709,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 315,
							EndLine:   315,
							StartPos:  5714,
							EndPos:    5716,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5714,
								EndPos:    5716,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 317,
					EndLine:   317,
					StartPos:  5721,
					EndPos:    5760,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 317,
						EndLine:   317,
						StartPos:  5727,
						EndPos:    5730,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 317,
							EndLine:   317,
							StartPos:  5732,
							EndPos:    5758,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 317,
								EndLine:   317,
								StartPos:  5748,
								EndPos:    5753,
							},
							Value: "class",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 317,
									EndLine:   317,
									StartPos:  5732,
									EndPos:    5738,
								},
								Value: "public",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 317,
								EndLine:   317,
								StartPos:  5756,
								EndPos:    5758,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 318,
					EndLine:   318,
					StartPos:  5763,
					EndPos:    5774,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 318,
						EndLine:   318,
						StartPos:  5763,
						EndPos:    5773,
					},
					Function: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 318,
							EndLine:   318,
							StartPos:  5763,
							EndPos:    5771,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 318,
									EndLine:   318,
									StartPos:  5764,
									EndPos:    5767,
								},
								Value: "foo",
							},
							&name.NamePart{
								Position: &position.Position{
									StartLine: 318,
									EndLine:   318,
									StartPos:  5768,
									EndPos:    5771,
								},
								Value: "bar",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 318,
							EndLine:   318,
							StartPos:  5771,
							EndPos:    5773,
						},
					},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 320,
					EndLine:   326,
					StartPos:  5778,
					EndPos:    5905,
				},
				PhpDocComment: "",
				ReturnsRef:    false,
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 320,
						EndLine:   320,
						StartPos:  5787,
						EndPos:    5790,
					},
					Value: "foo",
				},
				Params: []node.Node{
					&node.Parameter{
						Position: &position.Position{
							StartLine: 320,
							EndLine:   320,
							StartPos:  5791,
							EndPos:    5794,
						},
						ByRef:    true,
						Variadic: false,
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 320,
								EndLine:   320,
								StartPos:  5792,
								EndPos:    5794,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 320,
									EndLine:   320,
									StartPos:  5792,
									EndPos:    5794,
								},
								Value: "a",
							},
						},
					},
					&node.Parameter{
						Position: &position.Position{
							StartLine: 320,
							EndLine:   320,
							StartPos:  5796,
							EndPos:    5801,
						},
						ByRef:    false,
						Variadic: true,
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 320,
								EndLine:   320,
								StartPos:  5799,
								EndPos:    5801,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 320,
									EndLine:   320,
									StartPos:  5799,
									EndPos:    5801,
								},
								Value: "b",
							},
						},
					},
				},
				Stmts: []node.Node{
					&stmt.Function{
						Position: &position.Position{
							StartLine: 322,
							EndLine:   322,
							StartPos:  5830,
							EndPos:    5847,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						FunctionName: &node.Identifier{
							Position: &position.Position{
								StartLine: 322,
								EndLine:   322,
								StartPos:  5839,
								EndPos:    5842,
							},
							Value: "bar",
						},
						Stmts: []node.Node{},
					},
					&stmt.Class{
						Position: &position.Position{
							StartLine: 323,
							EndLine:   323,
							StartPos:  5851,
							EndPos:    5863,
						},
						PhpDocComment: "",
						ClassName: &node.Identifier{
							Position: &position.Position{
								StartLine: 323,
								EndLine:   323,
								StartPos:  5857,
								EndPos:    5860,
							},
							Value: "Baz",
						},
						Stmts: []node.Node{},
					},
					&stmt.Trait{
						Position: &position.Position{
							StartLine: 324,
							EndLine:   324,
							StartPos:  5867,
							EndPos:    5879,
						},
						PhpDocComment: "",
						TraitName: &node.Identifier{
							Position: &position.Position{
								StartLine: 324,
								EndLine:   324,
								StartPos:  5873,
								EndPos:    5877,
							},
							Value: "Quux",
						},
						Stmts: []node.Node{},
					},
					&stmt.Interface{
						Position: &position.Position{
							StartLine: 325,
							EndLine:   325,
							StartPos:  5883,
							EndPos:    5901,
						},
						PhpDocComment: "",
						InterfaceName: &node.Identifier{
							Position: &position.Position{
								StartLine: 325,
								EndLine:   325,
								StartPos:  5893,
								EndPos:    5898,
							},
							Value: "Quuux",
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 328,
					EndLine:   328,
					StartPos:  5911,
					EndPos:    5954,
				},
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 328,
						EndLine:   328,
						StartPos:  5920,
						EndPos:    5923,
					},
					Value: "foo",
				},
				Params: []node.Node{
					&node.Parameter{
						Position: &position.Position{
							StartLine: 328,
							EndLine:   328,
							StartPos:  5924,
							EndPos:    5931,
						},
						ByRef:    true,
						Variadic: false,
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5925,
								EndPos:    5927,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5925,
									EndPos:    5927,
								},
								Value: "a",
							},
						},
						DefaultValue: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5930,
								EndPos:    5931,
							},
							Value: "1",
						},
					},
					&node.Parameter{
						Position: &position.Position{
							StartLine: 328,
							EndLine:   328,
							StartPos:  5933,
							EndPos:    5942,
						},
						ByRef:    false,
						Variadic: true,
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5936,
								EndPos:    5938,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5936,
									EndPos:    5938,
								},
								Value: "b",
							},
						},
						DefaultValue: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5941,
								EndPos:    5942,
							},
							Value: "1",
						},
					},
					&node.Parameter{
						Position: &position.Position{
							StartLine: 328,
							EndLine:   328,
							StartPos:  5944,
							EndPos:    5950,
						},
						ByRef:    false,
						Variadic: false,
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5944,
								EndPos:    5946,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5944,
									EndPos:    5946,
								},
								Value: "c",
							},
						},
						DefaultValue: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5949,
								EndPos:    5950,
							},
							Value: "1",
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 329,
					EndLine:   329,
					StartPos:  5957,
					EndPos:    5995,
				},
				PhpDocComment: "",
				ReturnsRef:    false,
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 329,
						EndLine:   329,
						StartPos:  5966,
						EndPos:    5969,
					},
					Value: "foo",
				},
				Params: []node.Node{
					&node.Parameter{
						Position: &position.Position{
							StartLine: 329,
							EndLine:   329,
							StartPos:  5970,
							EndPos:    5978,
						},
						ByRef:    false,
						Variadic: false,
						VariableType: &node.Identifier{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5970,
								EndPos:    5975,
							},
							Value: "array",
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5976,
								EndPos:    5978,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5976,
									EndPos:    5978,
								},
								Value: "a",
							},
						},
					},
					&node.Parameter{
						Position: &position.Position{
							StartLine: 329,
							EndLine:   329,
							StartPos:  5980,
							EndPos:    5991,
						},
						Variadic: false,
						ByRef:    false,
						VariableType: &node.Identifier{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5980,
								EndPos:    5988,
							},
							Value: "callable",
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5989,
								EndPos:    5991,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5989,
									EndPos:    5991,
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
					StartLine: 330,
					EndLine:   330,
					StartPos:  5998,
					EndPos:    6100,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 330,
						EndLine:   330,
						StartPos:  6019,
						EndPos:    6022,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 330,
							EndLine:   330,
							StartPos:  5998,
							EndPos:    6006,
						},
						Value: "abstract",
					},
					&node.Identifier{
						Position: &position.Position{
							StartLine: 330,
							EndLine:   330,
							StartPos:  6007,
							EndPos:    6012,
						},
						Value: "final",
					},
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 330,
							EndLine:   330,
							StartPos:  6025,
							EndPos:    6066,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  6060,
								EndPos:    6063,
							},
							Value: "bar",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  6025,
									EndPos:    6033,
								},
								Value: "abstract",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  6034,
									EndPos:    6043,
								},
								Value: "protected",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  6044,
									EndPos:    6050,
								},
								Value: "static",
							},
						},
						Stmt: &stmt.Nop{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  6065,
								EndPos:    6066,
							},
						},
					},
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 330,
							EndLine:   330,
							StartPos:  6067,
							EndPos:    6098,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  6090,
								EndPos:    6093,
							},
							Value: "baz",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  6067,
									EndPos:    6072,
								},
								Value: "final",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  6073,
									EndPos:    6080,
								},
								Value: "private",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  6096,
								EndPos:    6098,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 332,
					EndLine:   332,
					StartPos:  6105,
					EndPos:    6119,
				},
				Expr: &expr.PropertyFetch{
					Position: &position.Position{
						StartLine: 332,
						EndLine:   332,
						StartPos:  6105,
						EndPos:    6118,
					},
					Variable: &expr.New{
						Position: &position.Position{
							StartLine: 332,
							EndLine:   332,
							StartPos:  6105,
							EndPos:    6112,
						},
						Class: &name.Name{
							Position: &position.Position{
								StartLine: 332,
								EndLine:   332,
								StartPos:  6109,
								EndPos:    6112,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 332,
										EndLine:   332,
										StartPos:  6109,
										EndPos:    6112,
									},
									Value: "Foo",
								},
							},
						},
					},
					Property: &node.Identifier{
						Position: &position.Position{
							StartLine: 332,
							EndLine:   332,
							StartPos:  6115,
							EndPos:    6118,
						},
						Value: "bar",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 333,
					EndLine:   333,
					StartPos:  6123,
					EndPos:    6134,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 333,
						EndLine:   333,
						StartPos:  6123,
						EndPos:    6133,
					},
					Function: &expr.New{
						Position: &position.Position{
							StartLine: 333,
							EndLine:   333,
							StartPos:  6123,
							EndPos:    6130,
						},
						Class: &name.Name{
							Position: &position.Position{
								StartLine: 333,
								EndLine:   333,
								StartPos:  6127,
								EndPos:    6130,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 333,
										EndLine:   333,
										StartPos:  6127,
										EndPos:    6130,
									},
									Value: "Foo",
								},
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 333,
							EndLine:   333,
							StartPos:  6131,
							EndPos:    6133,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 334,
					EndLine:   334,
					StartPos:  6137,
					EndPos:    6149,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 334,
						EndLine:   334,
						StartPos:  6137,
						EndPos:    6148,
					},
					Function: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 334,
							EndLine:   334,
							StartPos:  6137,
							EndPos:    6146,
						},
						Variable: &expr.ShortArray{
							Position: &position.Position{
								StartLine: 334,
								EndLine:   334,
								StartPos:  6137,
								EndPos:    6143,
							},
							Items: []node.Node{
								&expr.ArrayItem{
									Position: &position.Position{
										StartLine: 334,
										EndLine:   334,
										StartPos:  6138,
										EndPos:    6142,
									},
									Val: &expr.Variable{
										Position: &position.Position{
											StartLine: 334,
											EndLine:   334,
											StartPos:  6138,
											EndPos:    6142,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 334,
												EndLine:   334,
												StartPos:  6138,
												EndPos:    6142,
											},
											Value: "foo",
										},
									},
								},
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 334,
								EndLine:   334,
								StartPos:  6144,
								EndPos:    6145,
							},
							Value: "0",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 334,
							EndLine:   334,
							StartPos:  6146,
							EndPos:    6148,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 335,
					EndLine:   335,
					StartPos:  6152,
					EndPos:    6161,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 335,
						EndLine:   335,
						StartPos:  6152,
						EndPos:    6160,
					},
					Function: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 335,
							EndLine:   335,
							StartPos:  6152,
							EndPos:    6158,
						},
						Variable: &expr.ConstFetch{
							Position: &position.Position{
								StartLine: 335,
								EndLine:   335,
								StartPos:  6152,
								EndPos:    6155,
							},
							Constant: &name.Name{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  6152,
									EndPos:    6155,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 335,
											EndLine:   335,
											StartPos:  6152,
											EndPos:    6155,
										},
										Value: "foo",
									},
								},
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 335,
								EndLine:   335,
								StartPos:  6156,
								EndPos:    6157,
							},
							Value: "1",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 335,
							EndLine:   335,
							StartPos:  6158,
							EndPos:    6160,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 336,
					EndLine:   336,
					StartPos:  6164,
					EndPos:    6172,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 336,
						EndLine:   336,
						StartPos:  6164,
						EndPos:    6171,
					},
					Function: &scalar.String{
						Position: &position.Position{
							StartLine: 336,
							EndLine:   336,
							StartPos:  6164,
							EndPos:    6169,
						},
						Value: "\"foo\"",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 336,
							EndLine:   336,
							StartPos:  6169,
							EndPos:    6171,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 337,
					EndLine:   337,
					StartPos:  6175,
					EndPos:    6187,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 337,
						EndLine:   337,
						StartPos:  6175,
						EndPos:    6186,
					},
					Function: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 337,
							EndLine:   337,
							StartPos:  6175,
							EndPos:    6184,
						},
						Variable: &expr.ShortArray{
							Position: &position.Position{
								StartLine: 337,
								EndLine:   337,
								StartPos:  6175,
								EndPos:    6178,
							},
							Items: []node.Node{
								&expr.ArrayItem{
									Position: &position.Position{
										StartLine: 337,
										EndLine:   337,
										StartPos:  6176,
										EndPos:    6177,
									},
									Val: &scalar.Lnumber{
										Position: &position.Position{
											StartLine: 337,
											EndLine:   337,
											StartPos:  6176,
											EndPos:    6177,
										},
										Value: "1",
									},
								},
							},
						},
						Dim: &expr.Variable{
							Position: &position.Position{
								StartLine: 337,
								EndLine:   337,
								StartPos:  6179,
								EndPos:    6183,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 337,
									EndLine:   337,
									StartPos:  6179,
									EndPos:    6183,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 337,
							EndLine:   337,
							StartPos:  6184,
							EndPos:    6186,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 338,
					EndLine:   338,
					StartPos:  6190,
					EndPos:    6199,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 338,
						EndLine:   338,
						StartPos:  6190,
						EndPos:    6198,
					},
					VarName: &expr.FunctionCall{
						Position: &position.Position{
							StartLine: 338,
							EndLine:   338,
							StartPos:  6192,
							EndPos:    6197,
						},
						Function: &name.Name{
							Position: &position.Position{
								StartLine: 338,
								EndLine:   338,
								StartPos:  6192,
								EndPos:    6195,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 338,
										EndLine:   338,
										StartPos:  6192,
										EndPos:    6195,
									},
									Value: "foo",
								},
							},
						},
						ArgumentList: &node.ArgumentList{
							Position: &position.Position{
								StartLine: 338,
								EndLine:   338,
								StartPos:  6195,
								EndPos:    6197,
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 340,
					EndLine:   340,
					StartPos:  6203,
					EndPos:    6215,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 340,
						EndLine:   340,
						StartPos:  6203,
						EndPos:    6214,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 340,
							EndLine:   340,
							StartPos:  6203,
							EndPos:    6206,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 340,
									EndLine:   340,
									StartPos:  6203,
									EndPos:    6206,
								},
								Value: "Foo",
							},
						},
					},
					Call: &expr.Variable{
						Position: &position.Position{
							StartLine: 340,
							EndLine:   340,
							StartPos:  6208,
							EndPos:    6212,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 340,
								EndLine:   340,
								StartPos:  6208,
								EndPos:    6212,
							},
							Value: "bar",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 340,
							EndLine:   340,
							StartPos:  6212,
							EndPos:    6214,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 341,
					EndLine:   341,
					StartPos:  6218,
					EndPos:    6235,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 341,
						EndLine:   341,
						StartPos:  6218,
						EndPos:    6234,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 341,
							EndLine:   341,
							StartPos:  6218,
							EndPos:    6221,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  6218,
									EndPos:    6221,
								},
								Value: "Foo",
							},
						},
					},
					Call: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 341,
							EndLine:   341,
							StartPos:  6224,
							EndPos:    6231,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  6224,
								EndPos:    6228,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  6224,
									EndPos:    6228,
								},
								Value: "bar",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  6229,
								EndPos:    6230,
							},
							Value: "0",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 341,
							EndLine:   341,
							StartPos:  6232,
							EndPos:    6234,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 343,
					EndLine:   343,
					StartPos:  6241,
					EndPos:    6252,
				},
				Expr: &expr.PropertyFetch{
					Position: &position.Position{
						StartLine: 343,
						EndLine:   343,
						StartPos:  6241,
						EndPos:    6251,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 343,
							EndLine:   343,
							StartPos:  6241,
							EndPos:    6245,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 343,
								EndLine:   343,
								StartPos:  6241,
								EndPos:    6245,
							},
							Value: "foo",
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 343,
							EndLine:   343,
							StartPos:  6247,
							EndPos:    6251,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 343,
								EndLine:   343,
								StartPos:  6247,
								EndPos:    6251,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 344,
					EndLine:   344,
					StartPos:  6255,
					EndPos:    6271,
				},
				Expr: &expr.PropertyFetch{
					Position: &position.Position{
						StartLine: 344,
						EndLine:   344,
						StartPos:  6255,
						EndPos:    6269,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 344,
							EndLine:   344,
							StartPos:  6255,
							EndPos:    6259,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 344,
								EndLine:   344,
								StartPos:  6255,
								EndPos:    6259,
							},
							Value: "foo",
						},
					},
					Property: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 344,
							EndLine:   344,
							StartPos:  6262,
							EndPos:    6269,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 344,
								EndLine:   344,
								StartPos:  6262,
								EndPos:    6266,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 344,
									EndLine:   344,
									StartPos:  6262,
									EndPos:    6266,
								},
								Value: "bar",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 344,
								EndLine:   344,
								StartPos:  6267,
								EndPos:    6268,
							},
							Value: "0",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 346,
					EndLine:   346,
					StartPos:  6275,
					EndPos:    6297,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 346,
						EndLine:   346,
						StartPos:  6275,
						EndPos:    6296,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 346,
								EndLine:   346,
								StartPos:  6276,
								EndPos:    6282,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  6276,
									EndPos:    6277,
								},
								Value: "1",
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  6279,
									EndPos:    6282,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  6280,
										EndPos:    6282,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 346,
											EndLine:   346,
											StartPos:  6280,
											EndPos:    6282,
										},
										Value: "a",
									},
								},
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 346,
								EndLine:   346,
								StartPos:  6284,
								EndPos:    6295,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  6284,
									EndPos:    6285,
								},
								Value: "2",
							},
							Val: &expr.List{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  6287,
									EndPos:    6295,
								},
								Items: []node.Node{
									&expr.ArrayItem{
										Position: &position.Position{
											StartLine: 346,
											EndLine:   346,
											StartPos:  6292,
											EndPos:    6294,
										},
										Val: &expr.Variable{
											Position: &position.Position{
												StartLine: 346,
												EndLine:   346,
												StartPos:  6292,
												EndPos:    6294,
											},
											VarName: &node.Identifier{
												Position: &position.Position{
													StartLine: 346,
													EndLine:   346,
													StartPos:  6292,
													EndPos:    6294,
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
			&stmt.HaltCompiler{
				Position: &position.Position{
					StartLine: 348,
					EndLine:   348,
					StartPos:  6301,
					EndPos:    6319,
				},
			},
		},
	}

	php7parser := php7.NewParser()
	php7parser.Parse([]byte(src), &graph.AST{})
	actual := php7parser.Ast
	assert.DeepEqual(t, expected, actual)
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
			StartPos:  5,
			EndPos:    70,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  5,
					EndPos:    12,
				},
				Expr: &scalar.String{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  5,
						EndPos:    11,
					},
					Value: "\"test\"",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 3,
					EndLine:   3,
					StartPos:  15,
					EndPos:    24,
				},
				Expr: &scalar.String{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  15,
						EndPos:    23,
					},
					Value: "\"\\$test\"",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 4,
					EndLine:   6,
					StartPos:  27,
					EndPos:    41,
				},
				Expr: &scalar.String{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   6,
						StartPos:  27,
						EndPos:    40,
					},
					Value: "\"\n\t\t\ttest\n\t\t\"",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 7,
					EndLine:   7,
					StartPos:  44,
					EndPos:    52,
				},
				Expr: &scalar.String{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   7,
						StartPos:  44,
						EndPos:    51,
					},
					Value: "'$test'",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 8,
					EndLine:   10,
					StartPos:  55,
					EndPos:    70,
				},
				Expr: &scalar.String{
					Position: &position.Position{
						StartLine: 8,
						EndLine:   10,
						StartPos:  55,
						EndPos:    69,
					},
					Value: "'\n\t\t\t$test\n\t\t'",
				},
			},
		},
	}

	php7parser := php7.NewParser()
	php7parser.Parse([]byte(src), &graph.AST{})
	actual := php7parser.Ast
	assert.DeepEqual(t, expected, actual)
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
			StartPos:  5,
			EndPos:    120,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   3,
					StartPos:  5,
					EndPos:    16,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   3,
						StartPos:  5,
						EndPos:    15,
					},
					Label: "<<<CAD\n",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 4,
					EndLine:   6,
					StartPos:  19,
					EndPos:    37,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   6,
						StartPos:  19,
						EndPos:    36,
					},
					Label: "<<<CAD\n",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  26,
								EndPos:    33,
							},
							Value: "\thello\n",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 7,
					EndLine:   9,
					StartPos:  40,
					EndPos:    60,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   9,
						StartPos:  40,
						EndPos:    59,
					},
					Label: "<<<\"CAD\"\n",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 8,
								EndLine:   8,
								StartPos:  49,
								EndPos:    56,
							},
							Value: "\thello\n",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 10,
					EndLine:   12,
					StartPos:  63,
					EndPos:    90,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 10,
						EndLine:   12,
						StartPos:  63,
						EndPos:    89,
					},
					Label: "<<<\"CAD\"\n",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 11,
								EndLine:   11,
								StartPos:  72,
								EndPos:    79,
							},
							Value: "\thello ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 11,
								EndLine:   11,
								StartPos:  79,
								EndPos:    85,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  79,
									EndPos:    85,
								},
								Value: "world",
							},
						},
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 11,
								EndLine:   11,
								StartPos:  85,
								EndPos:    86,
							},
							Value: "\n",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 13,
					EndLine:   15,
					StartPos:  93,
					EndPos:    120,
				},
				Expr: &scalar.Heredoc{
					Position: &position.Position{
						StartLine: 13,
						EndLine:   15,
						StartPos:  93,
						EndPos:    119,
					},
					Label: "<<<'CAD'\n",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 14,
								EndLine:   14,
								StartPos:  102,
								EndPos:    116,
							},
							Value: "\thello $world\n",
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser()
	php7parser.Parse([]byte(src), &graph.AST{})
	actual := php7parser.Ast
	assert.DeepEqual(t, expected, actual)
}

func TestPhp7ControlCharsErrors(t *testing.T) {
	src := "<?php \004 echo $b; \"$a[\005test]\";"

	expected := []*errors.Error{
		{
			Msg: "WARNING: Unexpected character in input: '\004' (ASCII=4)",
			Pos: &position.Position{1, 1, 6, 7},
		},
		{
			Msg: "WARNING: Unexpected character in input: '\005' (ASCII=5)",
			Pos: &position.Position{1, 1, 21, 22},
		},
	}

	php7parser := php7.NewParser()
	php7parser.Parse([]byte(src), &graph.AST{})
	actual := php7parser.GetErrors()
	assert.DeepEqual(t, expected, actual)
}
