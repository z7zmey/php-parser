package php5_test

import (
	"bytes"
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/assign"
	"github.com/z7zmey/php-parser/node/expr/binary"
	"github.com/z7zmey/php-parser/node/expr/cast"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/position"
)

func TestPhp5(t *testing.T) {
	src := `<?
		foo($a, ...$b);
		$foo($a, ...$b);
		$foo->bar($a, ...$b);
		foo::bar($a, ...$b);
		$foo::bar($a, ...$b);
		new foo($a, ...$b);

		function foo(bar $bar=null, baz &...$baz) {}
		class foo {public function foo(bar $bar=null, baz &...$baz) {}}
		function(bar $bar=null, baz &...$baz) {};
		static function(bar $bar=null, baz &...$baz) {};

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
		"test $var[1234567890123456789012345678901234567890]";
		"test $var[bar]";
		"test $var[$bar]";
		"$foo $bar";
		"test $foo->bar()";
		"test ${foo}";
		"test ${foo[0]}";
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
		class foo{ const FOO = 1, BAR = 2; }
		class foo{ function bar() {} }
		class foo{ public static function &bar() {} }
		class foo{ final private function bar() {} protected function baz() {} }
		abstract class foo{ abstract public function bar(); }
		final class foo extends bar { }
		final class foo implements bar { }
		final class foo implements bar, baz { }

		const FOO = 1, BAR = 2;
		while (1) { continue; }
		while (1) { continue 2; }
		while (1) { continue(3); }
		declare(ticks=1);
		declare(ticks=1, strict_types=1) {}
		declare(ticks=1): enddeclare;
		do {} while(1);
		echo $a, 1;
		echo($a);
		for($i = 0; $i < 10; $i++, $i++) {}
		for(; $i < 10; $i++) : endfor;
		foreach ($a as $v) {}
		foreach ([] as $v) {}
		foreach ($a as $v) : endforeach;
		foreach ($a as $k => $v) {}
		foreach ([] as $k => $v) {}
		foreach ($a as $k => &$v) {}
		foreach ($a as $k => list($v)) {}
		function foo() {}

		function foo() {
			function bar() {}
			class Baz {}
			return $a;
		}
		
		function foo(array $a, callable $b) {return;}
		function &foo() {return 1;}
		function &foo() {}
		global $a, $b, $$c, ${foo()};
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
		namespace Foo\Bar {}
		namespace {}
		class foo {var $a;}
		class foo {public static $a, $b = 1;}
		class foo {public static $a = 1, $b;}
		static $a, $b = 1;
		static $a = 1, $b;

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
		class Foo { use Bar, Baz { one as public; } }
		class Foo { use Bar, Baz { one as public two; } }
		class Foo { use Bar, Baz { Bar::one insteadof Baz, Quux; Baz::one as two; } }

		try {}
		try {} catch (Exception $e) {}
		try {} catch (Exception $e) {} catch (RuntimeException $e) {}
		try {} catch (Exception $e) {} catch (\RuntimeException $e) {} catch (namespace\AdditionException $e) {}
		try {} catch (Exception $e) {} finally {}

		unset($a, $b);

		use Foo;
		use \Foo;
		use \Foo as Bar;
		use Foo, Bar;
		use Foo, Bar as Baz;
		use function Foo, \Bar;
		use function Foo as foo, \Bar as bar;
		use const Foo, \Bar;
		use const Foo as foo, \Bar as bar;

		$a[1];
		$a[1][2];
		array();
		array(1);
		array(1=>1, &$b,);
		array(3 =>&$b);
		array(&$b, 1=>1, 1, 3 =>&$b);
		~$a;
		!$a;

		Foo::Bar;
		clone($a);
		clone $a;
		function(){};
		function($a, $b) use ($c, &$d) {};
		function($a, $b) use (&$c, $d) {};
		function() {};
		foo;
		namespace\foo;
		\foo;

		empty($a);
		empty(Foo);
		@$a;
		eval($a);
		exit;
		exit($a);
		die();
		die($a);
		foo();
		namespace\foo(&$a);
		\foo([]);
		$foo(yield $a);

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
		isset(Foo);
		list() = $b;
		list($a, $b) = $b;
		list($a[]) = $b;
		list(list($a)) = $b;

		$a->foo();
		new Foo;
		new namespace\Foo();
		new \Foo();
		print($a);
		$a->foo;
		$a->foo[1];
		$a->foo->bar->baz()->quux[0];
		$a->foo()[1][1];
		` + "`cmd $a`;" + `
		` + "`cmd`;" + `
		` + "``;" + `
		[];
		[1];
		[1=>1, &$b,];

		Foo::bar();
		namespace\Foo::bar();
		\Foo::bar();
		Foo::$bar();
		$foo::$bar();
		Foo::$bar;
		namespace\Foo::$bar;
		\Foo::$bar;
		$a ? $b : $c;
		$a ? : $c;
		$a ? $b ? $c : $d : $e;
		$a ? $b : $c ? $d : $e;
		-$a;
		+$a;
		$$a;
		$$$a;
		yield;
		yield $a;
		yield $a => $b;
		yield Foo::class;
		yield $a => Foo::class;
		
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

		$a =& $b;
		$a =& new Foo;
		$a =& new Foo($b);
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


		(new \Foo());
		(new \Foo())->bar()->baz;
		(new \Foo())[0][0];
		(new \Foo())[0]->bar();

		array([0])[0][0];
		"foo"[0];
		foo[0];
		static::foo;

		new $foo;
		new $foo::$bar;
		new $a->b[0];
		new $a->b{$b ?: null}->$c->d[0];static $a = [1][0];

		static $a = !1;
		static $a = ~1;
		static $a = +1;
		static $a = -1;
		static $a = (1);
		static $a = 1 ?: 2;
		static $a = 1 ? 2 : 3;
		static $a = 1 & 2;
		static $a = 1 | 2;
		static $a = 1 ^ 2;
		static $a = 1 && 2;
		static $a = 1 || 2;
		static $a = 1 . 2;
		static $a = 1 / 2;
		static $a = 1 == 2;
		static $a = 1 >= 2;
		static $a = 1 > 2;
		static $a = 1 === 2;
		static $a = 1 and 2;
		static $a = 1 or 2;
		static $a = 1 xor 2;
		static $a = 1 - 2;
		static $a = 1 % 2;
		static $a = 1 * 2;
		static $a = 1 != 2;
		static $a = 1 !== 2;
		static $a = 1 + 2;
		static $a = 1 ** 2;
		static $a = 1 << 2;
		static $a = 1 >> 2;
		static $a = 1 <= 2;
		static $a = 1 < 2;
		static $a = Foo::bar;
		static $a = Foo::class;
		static $a = __CLASS__;
		static $a = Foo;
		static $a = namespace\Foo;
		static $a = \Foo;
		static $a = array();
		static $a = array(1 => 1, 2);
		static $a = [1, 2 => 2][0];

		if (yield 1) {}
		Foo::$$bar;

		$foo();
		$foo()[0][0];
		$a{$b};
		${$a};
		$foo::{$bar}();
		$foo::bar;

		__halt_compiler();

		parsing process must be terminated
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   379,
			StartPos:  6,
			EndPos:    6944,
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
								Variadic:    true,
								IsReference: false,
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
								Variadic:    true,
								IsReference: false,
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
								Variadic:    true,
								IsReference: false,
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
								Variadic:    true,
								IsReference: false,
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
			&stmt.Function{
				Position: &position.Position{
					StartLine: 9,
					EndLine:   9,
					StartPos:  137,
					EndPos:    180,
				},
				PhpDocComment: "",
				ReturnsRef:    false,
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 9,
						EndLine:   9,
						StartPos:  146,
						EndPos:    148,
					},
					Value: "foo",
				},
				Params: []node.Node{
					&node.Parameter{
						Position: &position.Position{
							StartLine: 9,
							EndLine:   9,
							StartPos:  150,
							EndPos:    162,
						},
						ByRef:    false,
						Variadic: false,
						VariableType: &name.Name{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  150,
								EndPos:    152,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 9,
										EndLine:   9,
										StartPos:  150,
										EndPos:    152,
									},
									Value: "bar",
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  154,
								EndPos:    157,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  154,
									EndPos:    157,
								},
								Value: "bar",
							},
						},
						DefaultValue: &expr.ConstFetch{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  159,
								EndPos:    162,
							},
							Constant: &name.Name{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  159,
									EndPos:    162,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  159,
											EndPos:    162,
										},
										Value: "null",
									},
								},
							},
						},
					},
					&node.Parameter{
						Position: &position.Position{
							StartLine: 9,
							EndLine:   9,
							StartPos:  165,
							EndPos:    176,
						},
						ByRef:    true,
						Variadic: true,
						VariableType: &name.Name{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  165,
								EndPos:    167,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 9,
										EndLine:   9,
										StartPos:  165,
										EndPos:    167,
									},
									Value: "baz",
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  173,
								EndPos:    176,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  173,
									EndPos:    176,
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
					StartLine: 10,
					EndLine:   10,
					StartPos:  184,
					EndPos:    246,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 10,
						EndLine:   10,
						StartPos:  190,
						EndPos:    192,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 10,
							EndLine:   10,
							StartPos:  195,
							EndPos:    245,
						},
						PhpDocComment: "",
						ReturnsRef:    false,
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 10,
								EndLine:   10,
								StartPos:  211,
								EndPos:    213,
							},
							Value: "foo",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 10,
									EndLine:   10,
									StartPos:  195,
									EndPos:    200,
								},
								Value: "public",
							},
						},
						Params: []node.Node{
							&node.Parameter{
								Position: &position.Position{
									StartLine: 10,
									EndLine:   10,
									StartPos:  215,
									EndPos:    227,
								},
								ByRef:    false,
								Variadic: false,
								VariableType: &name.Name{
									Position: &position.Position{
										StartLine: 10,
										EndLine:   10,
										StartPos:  215,
										EndPos:    217,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 10,
												EndLine:   10,
												StartPos:  215,
												EndPos:    217,
											},
											Value: "bar",
										},
									},
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 10,
										EndLine:   10,
										StartPos:  219,
										EndPos:    222,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 10,
											EndLine:   10,
											StartPos:  219,
											EndPos:    222,
										},
										Value: "bar",
									},
								},
								DefaultValue: &expr.ConstFetch{
									Position: &position.Position{
										StartLine: 10,
										EndLine:   10,
										StartPos:  224,
										EndPos:    227,
									},
									Constant: &name.Name{
										Position: &position.Position{
											StartLine: 10,
											EndLine:   10,
											StartPos:  224,
											EndPos:    227,
										},
										Parts: []node.Node{
											&name.NamePart{
												Position: &position.Position{
													StartLine: 10,
													EndLine:   10,
													StartPos:  224,
													EndPos:    227,
												},
												Value: "null",
											},
										},
									},
								},
							},
							&node.Parameter{
								Position: &position.Position{
									StartLine: 10,
									EndLine:   10,
									StartPos:  230,
									EndPos:    241,
								},
								Variadic: true,
								ByRef:    true,
								VariableType: &name.Name{
									Position: &position.Position{
										StartLine: 10,
										EndLine:   10,
										StartPos:  230,
										EndPos:    232,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 10,
												EndLine:   10,
												StartPos:  230,
												EndPos:    232,
											},
											Value: "baz",
										},
									},
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 10,
										EndLine:   10,
										StartPos:  238,
										EndPos:    241,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 10,
											EndLine:   10,
											StartPos:  238,
											EndPos:    241,
										},
										Value: "baz",
									},
								},
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 10,
								EndLine:   10,
								StartPos:  244,
								EndPos:    245,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 11,
					EndLine:   11,
					StartPos:  250,
					EndPos:    290,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 11,
						EndLine:   11,
						StartPos:  250,
						EndPos:    289,
					},
					Static:        false,
					PhpDocComment: "",
					ReturnsRef:    false,
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 11,
								EndLine:   11,
								StartPos:  259,
								EndPos:    271,
							},
							ByRef:    false,
							Variadic: false,
							VariableType: &name.Name{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  259,
									EndPos:    261,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 11,
											EndLine:   11,
											StartPos:  259,
											EndPos:    261,
										},
										Value: "bar",
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  263,
									EndPos:    266,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  263,
										EndPos:    266,
									},
									Value: "bar",
								},
							},
							DefaultValue: &expr.ConstFetch{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  268,
									EndPos:    271,
								},
								Constant: &name.Name{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  268,
										EndPos:    271,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 11,
												EndLine:   11,
												StartPos:  268,
												EndPos:    271,
											},
											Value: "null",
										},
									},
								},
							},
						},
						&node.Parameter{
							Position: &position.Position{
								StartLine: 11,
								EndLine:   11,
								StartPos:  274,
								EndPos:    285,
							},
							Variadic: true,
							ByRef:    true,
							VariableType: &name.Name{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  274,
									EndPos:    276,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 11,
											EndLine:   11,
											StartPos:  274,
											EndPos:    276,
										},
										Value: "baz",
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  282,
									EndPos:    285,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  282,
										EndPos:    285,
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
					StartLine: 12,
					EndLine:   12,
					StartPos:  294,
					EndPos:    341,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 12,
						EndLine:   12,
						StartPos:  294,
						EndPos:    340,
					},
					PhpDocComment: "",
					ReturnsRef:    false,
					Static:        true,
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 12,
								EndLine:   12,
								StartPos:  310,
								EndPos:    322,
							},
							Variadic: false,
							ByRef:    false,
							VariableType: &name.Name{
								Position: &position.Position{
									StartLine: 12,
									EndLine:   12,
									StartPos:  310,
									EndPos:    312,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 12,
											EndLine:   12,
											StartPos:  310,
											EndPos:    312,
										},
										Value: "bar",
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 12,
									EndLine:   12,
									StartPos:  314,
									EndPos:    317,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  314,
										EndPos:    317,
									},
									Value: "bar",
								},
							},
							DefaultValue: &expr.ConstFetch{
								Position: &position.Position{
									StartLine: 12,
									EndLine:   12,
									StartPos:  319,
									EndPos:    322,
								},
								Constant: &name.Name{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  319,
										EndPos:    322,
									},
									Parts: []node.Node{
										&name.NamePart{
											Position: &position.Position{
												StartLine: 12,
												EndLine:   12,
												StartPos:  319,
												EndPos:    322,
											},
											Value: "null",
										},
									},
								},
							},
						},
						&node.Parameter{
							Position: &position.Position{
								StartLine: 12,
								EndLine:   12,
								StartPos:  325,
								EndPos:    336,
							},
							ByRef:    true,
							Variadic: true,
							VariableType: &name.Name{
								Position: &position.Position{
									StartLine: 12,
									EndLine:   12,
									StartPos:  325,
									EndPos:    327,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 12,
											EndLine:   12,
											StartPos:  325,
											EndPos:    327,
										},
										Value: "baz",
									},
								},
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 12,
									EndLine:   12,
									StartPos:  333,
									EndPos:    336,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  333,
										EndPos:    336,
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
					StartLine: 14,
					EndLine:   14,
					StartPos:  346,
					EndPos:    365,
				},
				Expr: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 14,
						EndLine:   14,
						StartPos:  346,
						EndPos:    364,
					},
					Value: "1234567890123456789",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 15,
					EndLine:   15,
					StartPos:  369,
					EndPos:    389,
				},
				Expr: &scalar.Dnumber{
					Position: &position.Position{
						StartLine: 15,
						EndLine:   15,
						StartPos:  369,
						EndPos:    388,
					},
					Value: "12345678901234567890",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 16,
					EndLine:   16,
					StartPos:  393,
					EndPos:    395,
				},
				Expr: &scalar.Dnumber{
					Position: &position.Position{
						StartLine: 16,
						EndLine:   16,
						StartPos:  393,
						EndPos:    394,
					},
					Value: "0.",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 17,
					EndLine:   17,
					StartPos:  399,
					EndPos:    465,
				},
				Expr: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 17,
						EndLine:   17,
						StartPos:  399,
						EndPos:    464,
					},
					Value: "0b0111111111111111111111111111111111111111111111111111111111111111",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 18,
					EndLine:   18,
					StartPos:  469,
					EndPos:    535,
				},
				Expr: &scalar.Dnumber{
					Position: &position.Position{
						StartLine: 18,
						EndLine:   18,
						StartPos:  469,
						EndPos:    534,
					},
					Value: "0b1111111111111111111111111111111111111111111111111111111111111111",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 19,
					EndLine:   19,
					StartPos:  539,
					EndPos:    559,
				},
				Expr: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 19,
						EndLine:   19,
						StartPos:  539,
						EndPos:    558,
					},
					Value: "0x007111111111111111",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 20,
					EndLine:   20,
					StartPos:  563,
					EndPos:    581,
				},
				Expr: &scalar.Dnumber{
					Position: &position.Position{
						StartLine: 20,
						EndLine:   20,
						StartPos:  563,
						EndPos:    580,
					},
					Value: "0x8111111111111111",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 21,
					EndLine:   21,
					StartPos:  585,
					EndPos:    594,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 21,
						EndLine:   21,
						StartPos:  585,
						EndPos:    593,
					},
					Value: "__CLASS__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 22,
					EndLine:   22,
					StartPos:  598,
					EndPos:    605,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 22,
						EndLine:   22,
						StartPos:  598,
						EndPos:    604,
					},
					Value: "__DIR__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 23,
					EndLine:   23,
					StartPos:  609,
					EndPos:    617,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 23,
						EndLine:   23,
						StartPos:  609,
						EndPos:    616,
					},
					Value: "__FILE__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 24,
					EndLine:   24,
					StartPos:  621,
					EndPos:    633,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 24,
						EndLine:   24,
						StartPos:  621,
						EndPos:    632,
					},
					Value: "__FUNCTION__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 25,
					EndLine:   25,
					StartPos:  637,
					EndPos:    645,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 25,
						EndLine:   25,
						StartPos:  637,
						EndPos:    644,
					},
					Value: "__LINE__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 26,
					EndLine:   26,
					StartPos:  649,
					EndPos:    662,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 26,
						EndLine:   26,
						StartPos:  649,
						EndPos:    661,
					},
					Value: "__NAMESPACE__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 27,
					EndLine:   27,
					StartPos:  666,
					EndPos:    676,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 27,
						EndLine:   27,
						StartPos:  666,
						EndPos:    675,
					},
					Value: "__METHOD__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 28,
					EndLine:   28,
					StartPos:  680,
					EndPos:    689,
				},
				Expr: &scalar.MagicConstant{
					Position: &position.Position{
						StartLine: 28,
						EndLine:   28,
						StartPos:  680,
						EndPos:    688,
					},
					Value: "__TRAIT__",
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 30,
					EndLine:   30,
					StartPos:  694,
					EndPos:    705,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 30,
						EndLine:   30,
						StartPos:  694,
						EndPos:    704,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 30,
								EndLine:   30,
								StartPos:  695,
								EndPos:    699,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 30,
								EndLine:   30,
								StartPos:  700,
								EndPos:    703,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 30,
									EndLine:   30,
									StartPos:  700,
									EndPos:    703,
								},
								Value: "var",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 31,
					EndLine:   31,
					StartPos:  709,
					EndPos:    723,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 31,
						EndLine:   31,
						StartPos:  709,
						EndPos:    722,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 31,
								EndLine:   31,
								StartPos:  710,
								EndPos:    714,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 31,
								EndLine:   31,
								StartPos:  715,
								EndPos:    721,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 31,
									EndLine:   31,
									StartPos:  715,
									EndPos:    718,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 31,
										EndLine:   31,
										StartPos:  715,
										EndPos:    718,
									},
									Value: "var",
								},
							},
							Dim: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 31,
									EndLine:   31,
									StartPos:  720,
									EndPos:    720,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 32,
					EndLine:   32,
					StartPos:  727,
					EndPos:    780,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 32,
						EndLine:   32,
						StartPos:  727,
						EndPos:    779,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 32,
								EndLine:   32,
								StartPos:  728,
								EndPos:    732,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 32,
								EndLine:   32,
								StartPos:  733,
								EndPos:    778,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 32,
									EndLine:   32,
									StartPos:  733,
									EndPos:    736,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 32,
										EndLine:   32,
										StartPos:  733,
										EndPos:    736,
									},
									Value: "var",
								},
							},
							Dim: &scalar.String{
								Position: &position.Position{
									StartLine: 32,
									EndLine:   32,
									StartPos:  738,
									EndPos:    777,
								},
								Value: "1234567890123456789012345678901234567890",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 33,
					EndLine:   33,
					StartPos:  784,
					EndPos:    800,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 33,
						EndLine:   33,
						StartPos:  784,
						EndPos:    799,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 33,
								EndLine:   33,
								StartPos:  785,
								EndPos:    789,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 33,
								EndLine:   33,
								StartPos:  790,
								EndPos:    798,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 33,
									EndLine:   33,
									StartPos:  790,
									EndPos:    793,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 33,
										EndLine:   33,
										StartPos:  790,
										EndPos:    793,
									},
									Value: "var",
								},
							},
							Dim: &scalar.String{
								Position: &position.Position{
									StartLine: 33,
									EndLine:   33,
									StartPos:  795,
									EndPos:    797,
								},
								Value: "bar",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 34,
					EndLine:   34,
					StartPos:  804,
					EndPos:    821,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 34,
						EndLine:   34,
						StartPos:  804,
						EndPos:    820,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 34,
								EndLine:   34,
								StartPos:  805,
								EndPos:    809,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 34,
								EndLine:   34,
								StartPos:  810,
								EndPos:    819,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 34,
									EndLine:   34,
									StartPos:  810,
									EndPos:    813,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 34,
										EndLine:   34,
										StartPos:  810,
										EndPos:    813,
									},
									Value: "var",
								},
							},
							Dim: &expr.Variable{
								Position: &position.Position{
									StartLine: 34,
									EndLine:   34,
									StartPos:  815,
									EndPos:    818,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 34,
										EndLine:   34,
										StartPos:  815,
										EndPos:    818,
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
					StartLine: 35,
					EndLine:   35,
					StartPos:  825,
					EndPos:    836,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 35,
						EndLine:   35,
						StartPos:  825,
						EndPos:    835,
					},
					Parts: []node.Node{
						&expr.Variable{
							Position: &position.Position{
								StartLine: 35,
								EndLine:   35,
								StartPos:  826,
								EndPos:    829,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 35,
									EndLine:   35,
									StartPos:  826,
									EndPos:    829,
								},
								Value: "foo",
							},
						},
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 35,
								EndLine:   35,
								StartPos:  830,
								EndPos:    830,
							},
							Value: " ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 35,
								EndLine:   35,
								StartPos:  831,
								EndPos:    834,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 35,
									EndLine:   35,
									StartPos:  831,
									EndPos:    834,
								},
								Value: "bar",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 36,
					EndLine:   36,
					StartPos:  840,
					EndPos:    858,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 36,
						EndLine:   36,
						StartPos:  840,
						EndPos:    857,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 36,
								EndLine:   36,
								StartPos:  841,
								EndPos:    845,
							},
							Value: "test ",
						},
						&expr.PropertyFetch{
							Position: &position.Position{
								StartLine: 36,
								EndLine:   36,
								StartPos:  846,
								EndPos:    854,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 36,
									EndLine:   36,
									StartPos:  846,
									EndPos:    849,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 36,
										EndLine:   36,
										StartPos:  846,
										EndPos:    849,
									},
									Value: "foo",
								},
							},
							Property: &node.Identifier{
								Position: &position.Position{
									StartLine: 36,
									EndLine:   36,
									StartPos:  852,
									EndPos:    854,
								},
								Value: "bar",
							},
						},
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 36,
								EndLine:   36,
								StartPos:  855,
								EndPos:    856,
							},
							Value: "()",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 37,
					EndLine:   37,
					StartPos:  862,
					EndPos:    875,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 37,
						EndLine:   37,
						StartPos:  862,
						EndPos:    874,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 37,
								EndLine:   37,
								StartPos:  863,
								EndPos:    867,
							},
							Value: "test ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 37,
								EndLine:   37,
								StartPos:  868,
								EndPos:    873,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 37,
									EndLine:   37,
									StartPos:  870,
									EndPos:    872,
								},
								Value: "foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 38,
					EndLine:   38,
					StartPos:  879,
					EndPos:    895,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 38,
						EndLine:   38,
						StartPos:  879,
						EndPos:    894,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 38,
								EndLine:   38,
								StartPos:  880,
								EndPos:    884,
							},
							Value: "test ",
						},
						&expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 38,
								EndLine:   38,
								StartPos:  885,
								EndPos:    893,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 38,
									EndLine:   38,
									StartPos:  887,
									EndPos:    889,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 38,
										EndLine:   38,
										StartPos:  887,
										EndPos:    889,
									},
									Value: "foo",
								},
							},
							Dim: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 38,
									EndLine:   38,
									StartPos:  891,
									EndPos:    891,
								},
								Value: "0",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 39,
					EndLine:   39,
					StartPos:  899,
					EndPos:    919,
				},
				Expr: &scalar.Encapsed{
					Position: &position.Position{
						StartLine: 39,
						EndLine:   39,
						StartPos:  899,
						EndPos:    918,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 39,
								EndLine:   39,
								StartPos:  900,
								EndPos:    904,
							},
							Value: "test ",
						},
						&expr.MethodCall{
							Position: &position.Position{
								StartLine: 39,
								EndLine:   39,
								StartPos:  906,
								EndPos:    916,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 39,
									EndLine:   39,
									StartPos:  906,
									EndPos:    909,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 39,
										EndLine:   39,
										StartPos:  906,
										EndPos:    909,
									},
									Value: "foo",
								},
							},
							Method: &node.Identifier{
								Position: &position.Position{
									StartLine: 39,
									EndLine:   39,
									StartPos:  912,
									EndPos:    914,
								},
								Value: "bar",
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 39,
									EndLine:   39,
									StartPos:  915,
									EndPos:    916,
								},
							},
						},
					},
				},
			},
			&stmt.AltIf{
				Position: &position.Position{
					StartLine: 41,
					EndLine:   42,
					StartPos:  924,
					EndPos:    941,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 41,
						EndLine:   41,
						StartPos:  928,
						EndPos:    929,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 41,
							EndLine:   41,
							StartPos:  928,
							EndPos:    929,
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
					StartLine: 43,
					EndLine:   45,
					StartPos:  945,
					EndPos:    977,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 43,
						EndLine:   43,
						StartPos:  949,
						EndPos:    950,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 43,
							EndLine:   43,
							StartPos:  949,
							EndPos:    950,
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
							StartLine: 44,
							EndLine:   -1,
							StartPos:  957,
							EndPos:    -1,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 44,
								EndLine:   44,
								StartPos:  965,
								EndPos:    966,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 44,
									EndLine:   44,
									StartPos:  965,
									EndPos:    966,
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
					StartLine: 46,
					EndLine:   48,
					StartPos:  981,
					EndPos:    1006,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 46,
						EndLine:   46,
						StartPos:  985,
						EndPos:    986,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 46,
							EndLine:   46,
							StartPos:  985,
							EndPos:    986,
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
						StartLine: 47,
						EndLine:   -1,
						StartPos:  993,
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
					StartLine: 49,
					EndLine:   53,
					StartPos:  1010,
					EndPos:    1065,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 49,
						EndLine:   49,
						StartPos:  1014,
						EndPos:    1015,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 49,
							EndLine:   49,
							StartPos:  1014,
							EndPos:    1015,
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
							StartLine: 50,
							EndLine:   -1,
							StartPos:  1022,
							EndPos:    -1,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 50,
								EndLine:   50,
								StartPos:  1030,
								EndPos:    1031,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 50,
									EndLine:   50,
									StartPos:  1030,
									EndPos:    1031,
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
							StartLine: 51,
							EndLine:   -1,
							StartPos:  1037,
							EndPos:    -1,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 51,
								EndLine:   51,
								StartPos:  1045,
								EndPos:    1046,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 51,
									EndLine:   51,
									StartPos:  1045,
									EndPos:    1046,
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
						StartLine: 52,
						EndLine:   -1,
						StartPos:  1052,
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
					StartLine: 55,
					EndLine:   55,
					StartPos:  1070,
					EndPos:    1089,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 55,
						EndLine:   55,
						StartPos:  1077,
						EndPos:    1077,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 55,
						EndLine:   55,
						StartPos:  1080,
						EndPos:    1089,
					},
					Stmts: []node.Node{
						&stmt.Break{
							Position: &position.Position{
								StartLine: 55,
								EndLine:   55,
								StartPos:  1082,
								EndPos:    1087,
							},
						},
					},
				},
			},
			&stmt.While{
				Position: &position.Position{
					StartLine: 56,
					EndLine:   56,
					StartPos:  1093,
					EndPos:    1114,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 56,
						EndLine:   56,
						StartPos:  1100,
						EndPos:    1100,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 56,
						EndLine:   56,
						StartPos:  1103,
						EndPos:    1114,
					},
					Stmts: []node.Node{
						&stmt.Break{
							Position: &position.Position{
								StartLine: 56,
								EndLine:   56,
								StartPos:  1105,
								EndPos:    1112,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 56,
									EndLine:   56,
									StartPos:  1111,
									EndPos:    1111,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.AltWhile{
				Position: &position.Position{
					StartLine: 57,
					EndLine:   57,
					StartPos:  1118,
					EndPos:    1148,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 57,
						EndLine:   57,
						StartPos:  1125,
						EndPos:    1125,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 57,
						EndLine:   57,
						StartPos:  1130,
						EndPos:    1138,
					},
					Stmts: []node.Node{
						&stmt.Break{
							Position: &position.Position{
								StartLine: 57,
								EndLine:   57,
								StartPos:  1130,
								EndPos:    1138,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 57,
									EndLine:   57,
									StartPos:  1136,
									EndPos:    1136,
								},
								Value: "3",
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 58,
					EndLine:   58,
					StartPos:  1152,
					EndPos:    1187,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 58,
						EndLine:   58,
						StartPos:  1158,
						EndPos:    1160,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassConstList{
						Position: &position.Position{
							StartLine: 58,
							EndLine:   58,
							StartPos:  1163,
							EndPos:    1185,
						},
						Consts: []node.Node{
							&stmt.Constant{
								Position: &position.Position{
									StartLine: 58,
									EndLine:   58,
									StartPos:  1169,
									EndPos:    1175,
								},
								PhpDocComment: "",
								ConstantName: &node.Identifier{
									Position: &position.Position{
										StartLine: 58,
										EndLine:   58,
										StartPos:  1169,
										EndPos:    1171,
									},
									Value: "FOO",
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 58,
										EndLine:   58,
										StartPos:  1175,
										EndPos:    1175,
									},
									Value: "1",
								},
							},
							&stmt.Constant{
								Position: &position.Position{
									StartLine: 58,
									EndLine:   58,
									StartPos:  1178,
									EndPos:    1184,
								},
								PhpDocComment: "",
								ConstantName: &node.Identifier{
									Position: &position.Position{
										StartLine: 58,
										EndLine:   58,
										StartPos:  1178,
										EndPos:    1180,
									},
									Value: "BAR",
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 58,
										EndLine:   58,
										StartPos:  1184,
										EndPos:    1184,
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
					StartLine: 59,
					EndLine:   59,
					StartPos:  1191,
					EndPos:    1220,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 59,
						EndLine:   59,
						StartPos:  1197,
						EndPos:    1199,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 59,
							EndLine:   59,
							StartPos:  1202,
							EndPos:    1218,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 59,
								EndLine:   59,
								StartPos:  1211,
								EndPos:    1213,
							},
							Value: "bar",
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 59,
								EndLine:   59,
								StartPos:  1217,
								EndPos:    1218,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 60,
					EndLine:   60,
					StartPos:  1224,
					EndPos:    1268,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 60,
						EndLine:   60,
						StartPos:  1230,
						EndPos:    1232,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 60,
							EndLine:   60,
							StartPos:  1235,
							EndPos:    1266,
						},
						ReturnsRef:    true,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 60,
								EndLine:   60,
								StartPos:  1259,
								EndPos:    1261,
							},
							Value: "bar",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 60,
									EndLine:   60,
									StartPos:  1235,
									EndPos:    1240,
								},
								Value: "public",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 60,
									EndLine:   60,
									StartPos:  1242,
									EndPos:    1247,
								},
								Value: "static",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 60,
								EndLine:   60,
								StartPos:  1265,
								EndPos:    1266,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 61,
					EndLine:   61,
					StartPos:  1272,
					EndPos:    1343,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 61,
						EndLine:   61,
						StartPos:  1278,
						EndPos:    1280,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 61,
							EndLine:   61,
							StartPos:  1283,
							EndPos:    1313,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 61,
								EndLine:   61,
								StartPos:  1306,
								EndPos:    1308,
							},
							Value: "bar",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 61,
									EndLine:   61,
									StartPos:  1283,
									EndPos:    1287,
								},
								Value: "final",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 61,
									EndLine:   61,
									StartPos:  1289,
									EndPos:    1295,
								},
								Value: "private",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 61,
								EndLine:   61,
								StartPos:  1312,
								EndPos:    1313,
							},
							Stmts: []node.Node{},
						},
					},
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 61,
							EndLine:   61,
							StartPos:  1315,
							EndPos:    1341,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 61,
								EndLine:   61,
								StartPos:  1334,
								EndPos:    1336,
							},
							Value: "baz",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 61,
									EndLine:   61,
									StartPos:  1315,
									EndPos:    1323,
								},
								Value: "protected",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 61,
								EndLine:   61,
								StartPos:  1340,
								EndPos:    1341,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 62,
					EndLine:   62,
					StartPos:  1347,
					EndPos:    1399,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 62,
						EndLine:   62,
						StartPos:  1362,
						EndPos:    1364,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 62,
							EndLine:   62,
							StartPos:  1347,
							EndPos:    1354,
						},
						Value: "abstract",
					},
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						Position: &position.Position{
							StartLine: 62,
							EndLine:   62,
							StartPos:  1367,
							EndPos:    1397,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						MethodName: &node.Identifier{
							Position: &position.Position{
								StartLine: 62,
								EndLine:   62,
								StartPos:  1392,
								EndPos:    1394,
							},
							Value: "bar",
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 62,
									EndLine:   62,
									StartPos:  1367,
									EndPos:    1374,
								},
								Value: "abstract",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 62,
									EndLine:   62,
									StartPos:  1376,
									EndPos:    1381,
								},
								Value: "public",
							},
						},
						Stmt: &stmt.Nop{
							Position: &position.Position{
								StartLine: 62,
								EndLine:   62,
								StartPos:  1397,
								EndPos:    1397,
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 63,
					EndLine:   63,
					StartPos:  1403,
					EndPos:    1433,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 63,
						EndLine:   63,
						StartPos:  1415,
						EndPos:    1417,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 63,
							EndLine:   63,
							StartPos:  1403,
							EndPos:    1407,
						},
						Value: "final",
					},
				},
				Extends: &stmt.ClassExtends{
					Position: &position.Position{
						StartLine: 63,
						EndLine:   63,
						StartPos:  1419,
						EndPos:    1429,
					},
					ClassName: &name.Name{
						Position: &position.Position{
							StartLine: 63,
							EndLine:   63,
							StartPos:  1427,
							EndPos:    1429,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 63,
									EndLine:   63,
									StartPos:  1427,
									EndPos:    1429,
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
					StartLine: 64,
					EndLine:   64,
					StartPos:  1437,
					EndPos:    1470,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 64,
						EndLine:   64,
						StartPos:  1449,
						EndPos:    1451,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 64,
							EndLine:   64,
							StartPos:  1437,
							EndPos:    1441,
						},
						Value: "final",
					},
				},
				Implements: &stmt.ClassImplements{
					Position: &position.Position{
						StartLine: 64,
						EndLine:   64,
						StartPos:  1453,
						EndPos:    1466,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 64,
								EndLine:   64,
								StartPos:  1464,
								EndPos:    1466,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 64,
										EndLine:   64,
										StartPos:  1464,
										EndPos:    1466,
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
					StartLine: 65,
					EndLine:   65,
					StartPos:  1474,
					EndPos:    1512,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 65,
						EndLine:   65,
						StartPos:  1486,
						EndPos:    1488,
					},
					Value: "foo",
				},
				Modifiers: []node.Node{
					&node.Identifier{
						Position: &position.Position{
							StartLine: 65,
							EndLine:   65,
							StartPos:  1474,
							EndPos:    1478,
						},
						Value: "final",
					},
				},
				Implements: &stmt.ClassImplements{
					Position: &position.Position{
						StartLine: 65,
						EndLine:   65,
						StartPos:  1490,
						EndPos:    1508,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 65,
								EndLine:   65,
								StartPos:  1501,
								EndPos:    1503,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 65,
										EndLine:   65,
										StartPos:  1501,
										EndPos:    1503,
									},
									Value: "bar",
								},
							},
						},
						&name.Name{
							Position: &position.Position{
								StartLine: 65,
								EndLine:   65,
								StartPos:  1506,
								EndPos:    1508,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 65,
										EndLine:   65,
										StartPos:  1506,
										EndPos:    1508,
									},
									Value: "baz",
								},
							},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.ConstList{
				Position: &position.Position{
					StartLine: 67,
					EndLine:   67,
					StartPos:  1517,
					EndPos:    1539,
				},
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 67,
							EndLine:   67,
							StartPos:  1523,
							EndPos:    1529,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 67,
								EndLine:   67,
								StartPos:  1523,
								EndPos:    1525,
							},
							Value: "FOO",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 67,
								EndLine:   67,
								StartPos:  1529,
								EndPos:    1529,
							},
							Value: "1",
						},
					},
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 67,
							EndLine:   67,
							StartPos:  1532,
							EndPos:    1538,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 67,
								EndLine:   67,
								StartPos:  1532,
								EndPos:    1534,
							},
							Value: "BAR",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 67,
								EndLine:   67,
								StartPos:  1538,
								EndPos:    1538,
							},
							Value: "2",
						},
					},
				},
			},
			&stmt.While{
				Position: &position.Position{
					StartLine: 68,
					EndLine:   68,
					StartPos:  1543,
					EndPos:    1565,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 68,
						EndLine:   68,
						StartPos:  1550,
						EndPos:    1550,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 68,
						EndLine:   68,
						StartPos:  1553,
						EndPos:    1565,
					},
					Stmts: []node.Node{
						&stmt.Continue{
							Position: &position.Position{
								StartLine: 68,
								EndLine:   68,
								StartPos:  1555,
								EndPos:    1563,
							},
						},
					},
				},
			},
			&stmt.While{
				Position: &position.Position{
					StartLine: 69,
					EndLine:   69,
					StartPos:  1569,
					EndPos:    1593,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 69,
						EndLine:   69,
						StartPos:  1576,
						EndPos:    1576,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 69,
						EndLine:   69,
						StartPos:  1579,
						EndPos:    1593,
					},
					Stmts: []node.Node{
						&stmt.Continue{
							Position: &position.Position{
								StartLine: 69,
								EndLine:   69,
								StartPos:  1581,
								EndPos:    1591,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 69,
									EndLine:   69,
									StartPos:  1590,
									EndPos:    1590,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.While{
				Position: &position.Position{
					StartLine: 70,
					EndLine:   70,
					StartPos:  1597,
					EndPos:    1622,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 70,
						EndLine:   70,
						StartPos:  1604,
						EndPos:    1604,
					},
					Value: "1",
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 70,
						EndLine:   70,
						StartPos:  1607,
						EndPos:    1622,
					},
					Stmts: []node.Node{
						&stmt.Continue{
							Position: &position.Position{
								StartLine: 70,
								EndLine:   70,
								StartPos:  1609,
								EndPos:    1620,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 70,
									EndLine:   70,
									StartPos:  1618,
									EndPos:    1618,
								},
								Value: "3",
							},
						},
					},
				},
			},
			&stmt.Declare{
				Position: &position.Position{
					StartLine: 71,
					EndLine:   71,
					StartPos:  1626,
					EndPos:    1642,
				},
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 71,
							EndLine:   71,
							StartPos:  1634,
							EndPos:    1640,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 71,
								EndLine:   71,
								StartPos:  1634,
								EndPos:    1638,
							},
							Value: "ticks",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 71,
								EndLine:   71,
								StartPos:  1640,
								EndPos:    1640,
							},
							Value: "1",
						},
					},
				},
				Stmt: &stmt.Nop{
					Position: &position.Position{
						StartLine: 71,
						EndLine:   71,
						StartPos:  1642,
						EndPos:    1642,
					},
				},
			},
			&stmt.Declare{
				Position: &position.Position{
					StartLine: 72,
					EndLine:   72,
					StartPos:  1646,
					EndPos:    1680,
				},
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 72,
							EndLine:   72,
							StartPos:  1654,
							EndPos:    1660,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 72,
								EndLine:   72,
								StartPos:  1654,
								EndPos:    1658,
							},
							Value: "ticks",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 72,
								EndLine:   72,
								StartPos:  1660,
								EndPos:    1660,
							},
							Value: "1",
						},
					},
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 72,
							EndLine:   72,
							StartPos:  1663,
							EndPos:    1676,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 72,
								EndLine:   72,
								StartPos:  1663,
								EndPos:    1674,
							},
							Value: "strict_types",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 72,
								EndLine:   72,
								StartPos:  1676,
								EndPos:    1676,
							},
							Value: "1",
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 72,
						EndLine:   72,
						StartPos:  1679,
						EndPos:    1680,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Declare{
				Position: &position.Position{
					StartLine: 73,
					EndLine:   73,
					StartPos:  1684,
					EndPos:    1712,
				},
				Alt: true,
				Consts: []node.Node{
					&stmt.Constant{
						Position: &position.Position{
							StartLine: 73,
							EndLine:   73,
							StartPos:  1692,
							EndPos:    1698,
						},
						PhpDocComment: "",
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 73,
								EndLine:   73,
								StartPos:  1692,
								EndPos:    1696,
							},
							Value: "ticks",
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 73,
								EndLine:   73,
								StartPos:  1698,
								EndPos:    1698,
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
					StartLine: 74,
					EndLine:   74,
					StartPos:  1716,
					EndPos:    1730,
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 74,
						EndLine:   74,
						StartPos:  1719,
						EndPos:    1720,
					},
					Stmts: []node.Node{},
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 74,
						EndLine:   74,
						StartPos:  1728,
						EndPos:    1728,
					},
					Value: "1",
				},
			},
			&stmt.Echo{
				Position: &position.Position{
					StartLine: 75,
					EndLine:   75,
					StartPos:  1734,
					EndPos:    1744,
				},
				Exprs: []node.Node{
					&expr.Variable{
						Position: &position.Position{
							StartLine: 75,
							EndLine:   75,
							StartPos:  1739,
							EndPos:    1740,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 75,
								EndLine:   75,
								StartPos:  1739,
								EndPos:    1740,
							},
							Value: "a",
						},
					},
					&scalar.Lnumber{
						Position: &position.Position{
							StartLine: 75,
							EndLine:   75,
							StartPos:  1743,
							EndPos:    1743,
						},
						Value: "1",
					},
				},
			},
			&stmt.Echo{
				Position: &position.Position{
					StartLine: 76,
					EndLine:   76,
					StartPos:  1748,
					EndPos:    1756,
				},
				Exprs: []node.Node{
					&expr.Variable{
						Position: &position.Position{
							StartLine: 76,
							EndLine:   76,
							StartPos:  1753,
							EndPos:    1754,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 76,
								EndLine:   76,
								StartPos:  1753,
								EndPos:    1754,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.For{
				Position: &position.Position{
					StartLine: 77,
					EndLine:   77,
					StartPos:  1760,
					EndPos:    1794,
				},
				Init: []node.Node{
					&assign.Assign{
						Position: &position.Position{
							StartLine: 77,
							EndLine:   77,
							StartPos:  1764,
							EndPos:    1769,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1764,
								EndPos:    1765,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1764,
									EndPos:    1765,
								},
								Value: "i",
							},
						},
						Expression: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1769,
								EndPos:    1769,
							},
							Value: "0",
						},
					},
				},
				Cond: []node.Node{
					&binary.Smaller{
						Position: &position.Position{
							StartLine: 77,
							EndLine:   77,
							StartPos:  1772,
							EndPos:    1778,
						},
						Left: &expr.Variable{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1772,
								EndPos:    1773,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1772,
									EndPos:    1773,
								},
								Value: "i",
							},
						},
						Right: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1777,
								EndPos:    1778,
							},
							Value: "10",
						},
					},
				},
				Loop: []node.Node{
					&expr.PostInc{
						Position: &position.Position{
							StartLine: 77,
							EndLine:   77,
							StartPos:  1781,
							EndPos:    1784,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1781,
								EndPos:    1782,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1781,
									EndPos:    1782,
								},
								Value: "i",
							},
						},
					},
					&expr.PostInc{
						Position: &position.Position{
							StartLine: 77,
							EndLine:   77,
							StartPos:  1787,
							EndPos:    1790,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1787,
								EndPos:    1788,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1787,
									EndPos:    1788,
								},
								Value: "i",
							},
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 77,
						EndLine:   77,
						StartPos:  1793,
						EndPos:    1794,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.AltFor{
				Position: &position.Position{
					StartLine: 78,
					EndLine:   78,
					StartPos:  1798,
					EndPos:    1827,
				},
				Cond: []node.Node{
					&binary.Smaller{
						Position: &position.Position{
							StartLine: 78,
							EndLine:   78,
							StartPos:  1804,
							EndPos:    1810,
						},
						Left: &expr.Variable{
							Position: &position.Position{
								StartLine: 78,
								EndLine:   78,
								StartPos:  1804,
								EndPos:    1805,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 78,
									EndLine:   78,
									StartPos:  1804,
									EndPos:    1805,
								},
								Value: "i",
							},
						},
						Right: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 78,
								EndLine:   78,
								StartPos:  1809,
								EndPos:    1810,
							},
							Value: "10",
						},
					},
				},
				Loop: []node.Node{
					&expr.PostInc{
						Position: &position.Position{
							StartLine: 78,
							EndLine:   78,
							StartPos:  1813,
							EndPos:    1816,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 78,
								EndLine:   78,
								StartPos:  1813,
								EndPos:    1814,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 78,
									EndLine:   78,
									StartPos:  1813,
									EndPos:    1814,
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
					StartLine: 79,
					EndLine:   79,
					StartPos:  1831,
					EndPos:    1851,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 79,
						EndLine:   79,
						StartPos:  1840,
						EndPos:    1841,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 79,
							EndLine:   79,
							StartPos:  1840,
							EndPos:    1841,
						},
						Value: "a",
					},
				},
				Variable: &expr.Variable{
					Position: &position.Position{
						StartLine: 79,
						EndLine:   79,
						StartPos:  1846,
						EndPos:    1847,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 79,
							EndLine:   79,
							StartPos:  1846,
							EndPos:    1847,
						},
						Value: "v",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 79,
						EndLine:   79,
						StartPos:  1850,
						EndPos:    1851,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Foreach{
				Position: &position.Position{
					StartLine: 80,
					EndLine:   80,
					StartPos:  1855,
					EndPos:    1875,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 80,
						EndLine:   80,
						StartPos:  1864,
						EndPos:    1865,
					},
					Items: []node.Node{},
				},
				Variable: &expr.Variable{
					Position: &position.Position{
						StartLine: 80,
						EndLine:   80,
						StartPos:  1870,
						EndPos:    1871,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 80,
							EndLine:   80,
							StartPos:  1870,
							EndPos:    1871,
						},
						Value: "v",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 80,
						EndLine:   80,
						StartPos:  1874,
						EndPos:    1875,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.AltForeach{
				Position: &position.Position{
					StartLine: 81,
					EndLine:   81,
					StartPos:  1879,
					EndPos:    1910,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 81,
						EndLine:   81,
						StartPos:  1888,
						EndPos:    1889,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 81,
							EndLine:   81,
							StartPos:  1888,
							EndPos:    1889,
						},
						Value: "a",
					},
				},
				Variable: &expr.Variable{
					Position: &position.Position{
						StartLine: 81,
						EndLine:   81,
						StartPos:  1894,
						EndPos:    1895,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 81,
							EndLine:   81,
							StartPos:  1894,
							EndPos:    1895,
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
					StartLine: 82,
					EndLine:   82,
					StartPos:  1914,
					EndPos:    1940,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 82,
						EndLine:   82,
						StartPos:  1923,
						EndPos:    1924,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1923,
							EndPos:    1924,
						},
						Value: "a",
					},
				},
				Key: &expr.Variable{
					Position: &position.Position{
						StartLine: 82,
						EndLine:   82,
						StartPos:  1929,
						EndPos:    1930,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1929,
							EndPos:    1930,
						},
						Value: "k",
					},
				},
				Variable: &expr.Variable{
					Position: &position.Position{
						StartLine: 82,
						EndLine:   82,
						StartPos:  1935,
						EndPos:    1936,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1935,
							EndPos:    1936,
						},
						Value: "v",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 82,
						EndLine:   82,
						StartPos:  1939,
						EndPos:    1940,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Foreach{
				Position: &position.Position{
					StartLine: 83,
					EndLine:   83,
					StartPos:  1944,
					EndPos:    1970,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 83,
						EndLine:   83,
						StartPos:  1953,
						EndPos:    1954,
					},
					Items: []node.Node{},
				},
				Key: &expr.Variable{
					Position: &position.Position{
						StartLine: 83,
						EndLine:   83,
						StartPos:  1959,
						EndPos:    1960,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 83,
							EndLine:   83,
							StartPos:  1959,
							EndPos:    1960,
						},
						Value: "k",
					},
				},
				Variable: &expr.Variable{
					Position: &position.Position{
						StartLine: 83,
						EndLine:   83,
						StartPos:  1965,
						EndPos:    1966,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 83,
							EndLine:   83,
							StartPos:  1965,
							EndPos:    1966,
						},
						Value: "v",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 83,
						EndLine:   83,
						StartPos:  1969,
						EndPos:    1970,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Foreach{
				Position: &position.Position{
					StartLine: 84,
					EndLine:   84,
					StartPos:  1974,
					EndPos:    2001,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 84,
						EndLine:   84,
						StartPos:  1983,
						EndPos:    1984,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1983,
							EndPos:    1984,
						},
						Value: "a",
					},
				},
				Key: &expr.Variable{
					Position: &position.Position{
						StartLine: 84,
						EndLine:   84,
						StartPos:  1989,
						EndPos:    1990,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1989,
							EndPos:    1990,
						},
						Value: "k",
					},
				},
				Variable: &expr.Reference{
					Position: &position.Position{
						StartLine: 84,
						EndLine:   84,
						StartPos:  1995,
						EndPos:    1997,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1996,
							EndPos:    1997,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 84,
								EndLine:   84,
								StartPos:  1996,
								EndPos:    1997,
							},
							Value: "v",
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 84,
						EndLine:   84,
						StartPos:  2000,
						EndPos:    2001,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Foreach{
				Position: &position.Position{
					StartLine: 85,
					EndLine:   85,
					StartPos:  2005,
					EndPos:    2037,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 85,
						EndLine:   85,
						StartPos:  2014,
						EndPos:    2015,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 85,
							EndLine:   85,
							StartPos:  2014,
							EndPos:    2015,
						},
						Value: "a",
					},
				},
				Key: &expr.Variable{
					Position: &position.Position{
						StartLine: 85,
						EndLine:   85,
						StartPos:  2020,
						EndPos:    2021,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 85,
							EndLine:   85,
							StartPos:  2020,
							EndPos:    2021,
						},
						Value: "k",
					},
				},
				Variable: &expr.List{
					Position: &position.Position{
						StartLine: 85,
						EndLine:   85,
						StartPos:  2026,
						EndPos:    2033,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 85,
								EndLine:   85,
								StartPos:  2031,
								EndPos:    2032,
							},
							Val: &expr.Variable{
								Position: &position.Position{
									StartLine: 85,
									EndLine:   85,
									StartPos:  2031,
									EndPos:    2032,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 85,
										EndLine:   85,
										StartPos:  2031,
										EndPos:    2032,
									},
									Value: "v",
								},
							},
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 85,
						EndLine:   85,
						StartPos:  2036,
						EndPos:    2037,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 86,
					EndLine:   86,
					StartPos:  2041,
					EndPos:    2057,
				},
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 86,
						EndLine:   86,
						StartPos:  2050,
						EndPos:    2052,
					},
					Value: "foo",
				},
				Stmts: []node.Node{},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 88,
					EndLine:   92,
					StartPos:  2062,
					EndPos:    2132,
				},
				PhpDocComment: "",
				ReturnsRef:    false,
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 88,
						EndLine:   88,
						StartPos:  2071,
						EndPos:    2073,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.Function{
						Position: &position.Position{
							StartLine: 89,
							EndLine:   89,
							StartPos:  2082,
							EndPos:    2098,
						},
						PhpDocComment: "",
						ReturnsRef:    false,
						FunctionName: &node.Identifier{
							Position: &position.Position{
								StartLine: 89,
								EndLine:   89,
								StartPos:  2091,
								EndPos:    2093,
							},
							Value: "bar",
						},
						Stmts: []node.Node{},
					},
					&stmt.Class{
						Position: &position.Position{
							StartLine: 90,
							EndLine:   90,
							StartPos:  2103,
							EndPos:    2114,
						},
						PhpDocComment: "",
						ClassName: &node.Identifier{
							Position: &position.Position{
								StartLine: 90,
								EndLine:   90,
								StartPos:  2109,
								EndPos:    2111,
							},
							Value: "Baz",
						},
						Stmts: []node.Node{},
					},
					&stmt.Return{
						Position: &position.Position{
							StartLine: 91,
							EndLine:   91,
							StartPos:  2119,
							EndPos:    2128,
						},
						Expr: &expr.Variable{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2126,
								EndPos:    2127,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2126,
									EndPos:    2127,
								},
								Value: "a",
							},
						},
					},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 94,
					EndLine:   94,
					StartPos:  2139,
					EndPos:    2183,
				},
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 94,
						EndLine:   94,
						StartPos:  2148,
						EndPos:    2150,
					},
					Value: "foo",
				},
				Params: []node.Node{
					&node.Parameter{
						Position: &position.Position{
							StartLine: 94,
							EndLine:   94,
							StartPos:  2152,
							EndPos:    2159,
						},
						ByRef:    false,
						Variadic: false,
						VariableType: &node.Identifier{
							Position: &position.Position{
								StartLine: 94,
								EndLine:   94,
								StartPos:  2152,
								EndPos:    2156,
							},
							Value: "array",
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 94,
								EndLine:   94,
								StartPos:  2158,
								EndPos:    2159,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 94,
									EndLine:   94,
									StartPos:  2158,
									EndPos:    2159,
								},
								Value: "a",
							},
						},
					},
					&node.Parameter{
						Position: &position.Position{
							StartLine: 94,
							EndLine:   94,
							StartPos:  2162,
							EndPos:    2172,
						},
						ByRef:    false,
						Variadic: false,
						VariableType: &node.Identifier{
							Position: &position.Position{
								StartLine: 94,
								EndLine:   94,
								StartPos:  2162,
								EndPos:    2169,
							},
							Value: "callable",
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 94,
								EndLine:   94,
								StartPos:  2171,
								EndPos:    2172,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 94,
									EndLine:   94,
									StartPos:  2171,
									EndPos:    2172,
								},
								Value: "b",
							},
						},
					},
				},
				Stmts: []node.Node{
					&stmt.Return{
						Position: &position.Position{
							StartLine: 94,
							EndLine:   94,
							StartPos:  2176,
							EndPos:    2182,
						},
					},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 95,
					EndLine:   95,
					StartPos:  2187,
					EndPos:    2213,
				},
				ReturnsRef:    true,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 95,
						EndLine:   95,
						StartPos:  2197,
						EndPos:    2199,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.Return{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2204,
							EndPos:    2212,
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 95,
								EndLine:   95,
								StartPos:  2211,
								EndPos:    2211,
							},
							Value: "1",
						},
					},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 96,
					EndLine:   96,
					StartPos:  2217,
					EndPos:    2234,
				},
				PhpDocComment: "",
				ReturnsRef:    true,
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 96,
						EndLine:   96,
						StartPos:  2227,
						EndPos:    2229,
					},
					Value: "foo",
				},
				Stmts: []node.Node{},
			},
			&stmt.Global{
				Position: &position.Position{
					StartLine: 97,
					EndLine:   97,
					StartPos:  2238,
					EndPos:    2266,
				},
				Vars: []node.Node{
					&expr.Variable{
						Position: &position.Position{
							StartLine: 97,
							EndLine:   97,
							StartPos:  2245,
							EndPos:    2246,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2245,
								EndPos:    2246,
							},
							Value: "a",
						},
					},
					&expr.Variable{
						Position: &position.Position{
							StartLine: 97,
							EndLine:   97,
							StartPos:  2249,
							EndPos:    2250,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2249,
								EndPos:    2250,
							},
							Value: "b",
						},
					},
					&expr.Variable{
						Position: &position.Position{
							StartLine: 97,
							EndLine:   97,
							StartPos:  2253,
							EndPos:    2255,
						},
						VarName: &expr.Variable{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2254,
								EndPos:    2255,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2254,
									EndPos:    2255,
								},
								Value: "c",
							},
						},
					},
					&expr.Variable{
						Position: &position.Position{
							StartLine: 97,
							EndLine:   97,
							StartPos:  2258,
							EndPos:    2265,
						},
						VarName: &expr.FunctionCall{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2260,
								EndPos:    2264,
							},
							Function: &name.Name{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2260,
									EndPos:    2262,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 97,
											EndLine:   97,
											StartPos:  2260,
											EndPos:    2262,
										},
										Value: "foo",
									},
								},
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2263,
									EndPos:    2264,
								},
							},
						},
					},
				},
			},
			&stmt.Label{
				Position: &position.Position{
					StartLine: 98,
					EndLine:   98,
					StartPos:  2270,
					EndPos:    2271,
				},
				LabelName: &node.Identifier{
					Position: &position.Position{
						StartLine: 98,
						EndLine:   98,
						StartPos:  2270,
						EndPos:    2270,
					},
					Value: "a",
				},
			},
			&stmt.Goto{
				Position: &position.Position{
					StartLine: 99,
					EndLine:   99,
					StartPos:  2276,
					EndPos:    2282,
				},
				Label: &node.Identifier{
					Position: &position.Position{
						StartLine: 99,
						EndLine:   99,
						StartPos:  2281,
						EndPos:    2281,
					},
					Value: "a",
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 100,
					EndLine:   100,
					StartPos:  2286,
					EndPos:    2295,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 100,
						EndLine:   100,
						StartPos:  2290,
						EndPos:    2291,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 100,
							EndLine:   100,
							StartPos:  2290,
							EndPos:    2291,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 100,
						EndLine:   100,
						StartPos:  2294,
						EndPos:    2295,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 101,
					EndLine:   101,
					StartPos:  2299,
					EndPos:    2323,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 101,
						EndLine:   101,
						StartPos:  2303,
						EndPos:    2304,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 101,
							EndLine:   101,
							StartPos:  2303,
							EndPos:    2304,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 101,
						EndLine:   101,
						StartPos:  2307,
						EndPos:    2308,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 101,
							EndLine:   101,
							StartPos:  2310,
							EndPos:    2323,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 101,
								EndLine:   101,
								StartPos:  2318,
								EndPos:    2319,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 101,
									EndLine:   101,
									StartPos:  2318,
									EndPos:    2319,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 101,
								EndLine:   101,
								StartPos:  2322,
								EndPos:    2323,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 102,
					EndLine:   102,
					StartPos:  2327,
					EndPos:    2344,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 102,
						EndLine:   102,
						StartPos:  2331,
						EndPos:    2332,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 102,
							EndLine:   102,
							StartPos:  2331,
							EndPos:    2332,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 102,
						EndLine:   102,
						StartPos:  2335,
						EndPos:    2336,
					},
					Stmts: []node.Node{},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 102,
						EndLine:   102,
						StartPos:  2338,
						EndPos:    2344,
					},
					Stmt: &stmt.StmtList{
						Position: &position.Position{
							StartLine: 102,
							EndLine:   102,
							StartPos:  2343,
							EndPos:    2344,
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 103,
					EndLine:   103,
					StartPos:  2348,
					EndPos:    2395,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 103,
						EndLine:   103,
						StartPos:  2352,
						EndPos:    2353,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2352,
							EndPos:    2353,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 103,
						EndLine:   103,
						StartPos:  2356,
						EndPos:    2357,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2359,
							EndPos:    2372,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2367,
								EndPos:    2368,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2367,
									EndPos:    2368,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2371,
								EndPos:    2372,
							},
							Stmts: []node.Node{},
						},
					},
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2374,
							EndPos:    2387,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2382,
								EndPos:    2383,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2382,
									EndPos:    2383,
								},
								Value: "c",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2386,
								EndPos:    2387,
							},
							Stmts: []node.Node{},
						},
					},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 103,
						EndLine:   103,
						StartPos:  2389,
						EndPos:    2395,
					},
					Stmt: &stmt.StmtList{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2394,
							EndPos:    2395,
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 104,
					EndLine:   104,
					StartPos:  2399,
					EndPos:    2447,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 104,
						EndLine:   104,
						StartPos:  2403,
						EndPos:    2404,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 104,
							EndLine:   104,
							StartPos:  2403,
							EndPos:    2404,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 104,
						EndLine:   104,
						StartPos:  2407,
						EndPos:    2408,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 104,
							EndLine:   104,
							StartPos:  2410,
							EndPos:    2423,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 104,
								EndLine:   104,
								StartPos:  2418,
								EndPos:    2419,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 104,
									EndLine:   104,
									StartPos:  2418,
									EndPos:    2419,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 104,
								EndLine:   104,
								StartPos:  2422,
								EndPos:    2423,
							},
							Stmts: []node.Node{},
						},
					},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 104,
						EndLine:   104,
						StartPos:  2425,
						EndPos:    2447,
					},
					Stmt: &stmt.If{
						Position: &position.Position{
							StartLine: 104,
							EndLine:   104,
							StartPos:  2430,
							EndPos:    2447,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 104,
								EndLine:   104,
								StartPos:  2434,
								EndPos:    2435,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 104,
									EndLine:   104,
									StartPos:  2434,
									EndPos:    2435,
								},
								Value: "c",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 104,
								EndLine:   104,
								StartPos:  2438,
								EndPos:    2439,
							},
							Stmts: []node.Node{},
						},
						Else: &stmt.Else{
							Position: &position.Position{
								StartLine: 104,
								EndLine:   104,
								StartPos:  2441,
								EndPos:    2447,
							},
							Stmt: &stmt.StmtList{
								Position: &position.Position{
									StartLine: 104,
									EndLine:   104,
									StartPos:  2446,
									EndPos:    2447,
								},
								Stmts: []node.Node{},
							},
						},
					},
				},
			},
			&stmt.Nop{
				Position: &position.Position{
					StartLine: 105,
					EndLine:   105,
					StartPos:  2451,
					EndPos:    2452,
				},
			},
			&stmt.InlineHtml{
				Position: &position.Position{
					StartLine: 105,
					EndLine:   105,
					StartPos:  2454,
					EndPos:    2465,
				},
				Value: "<div></div> ",
			},
			&stmt.Interface{
				Position: &position.Position{
					StartLine: 106,
					EndLine:   106,
					StartPos:  2471,
					EndPos:    2486,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 106,
						EndLine:   106,
						StartPos:  2481,
						EndPos:    2483,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{},
			},
			&stmt.Interface{
				Position: &position.Position{
					StartLine: 107,
					EndLine:   107,
					StartPos:  2490,
					EndPos:    2517,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 107,
						EndLine:   107,
						StartPos:  2500,
						EndPos:    2502,
					},
					Value: "Foo",
				},
				Extends: &stmt.InterfaceExtends{
					Position: &position.Position{
						StartLine: 107,
						EndLine:   107,
						StartPos:  2504,
						EndPos:    2514,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 107,
								EndLine:   107,
								StartPos:  2512,
								EndPos:    2514,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 107,
										EndLine:   107,
										StartPos:  2512,
										EndPos:    2514,
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
					StartLine: 108,
					EndLine:   108,
					StartPos:  2521,
					EndPos:    2553,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 108,
						EndLine:   108,
						StartPos:  2531,
						EndPos:    2533,
					},
					Value: "Foo",
				},
				Extends: &stmt.InterfaceExtends{
					Position: &position.Position{
						StartLine: 108,
						EndLine:   108,
						StartPos:  2535,
						EndPos:    2550,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 108,
								EndLine:   108,
								StartPos:  2543,
								EndPos:    2545,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 108,
										EndLine:   108,
										StartPos:  2543,
										EndPos:    2545,
									},
									Value: "Bar",
								},
							},
						},
						&name.Name{
							Position: &position.Position{
								StartLine: 108,
								EndLine:   108,
								StartPos:  2548,
								EndPos:    2550,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 108,
										EndLine:   108,
										StartPos:  2548,
										EndPos:    2550,
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
					StartLine: 109,
					EndLine:   109,
					StartPos:  2557,
					EndPos:    2570,
				},
				NamespaceName: &name.Name{
					Position: &position.Position{
						StartLine: 109,
						EndLine:   109,
						StartPos:  2567,
						EndPos:    2569,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 109,
								EndLine:   109,
								StartPos:  2567,
								EndPos:    2569,
							},
							Value: "Foo",
						},
					},
				},
			},
			&stmt.Namespace{
				Position: &position.Position{
					StartLine: 110,
					EndLine:   110,
					StartPos:  2574,
					EndPos:    2593,
				},
				NamespaceName: &name.Name{
					Position: &position.Position{
						StartLine: 110,
						EndLine:   110,
						StartPos:  2584,
						EndPos:    2590,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2584,
								EndPos:    2586,
							},
							Value: "Foo",
						},
						&name.NamePart{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2588,
								EndPos:    2590,
							},
							Value: "Bar",
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Namespace{
				Position: &position.Position{
					StartLine: 111,
					EndLine:   111,
					StartPos:  2597,
					EndPos:    2608,
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 112,
					EndLine:   112,
					StartPos:  2612,
					EndPos:    2630,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 112,
						EndLine:   112,
						StartPos:  2618,
						EndPos:    2620,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Position: &position.Position{
							StartLine: 112,
							EndLine:   112,
							StartPos:  2623,
							EndPos:    2629,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 112,
									EndLine:   112,
									StartPos:  2623,
									EndPos:    2625,
								},
								Value: "var",
							},
						},
						Properties: []node.Node{
							&stmt.Property{
								Position: &position.Position{
									StartLine: 112,
									EndLine:   112,
									StartPos:  2627,
									EndPos:    2628,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 112,
										EndLine:   112,
										StartPos:  2627,
										EndPos:    2628,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 112,
											EndLine:   112,
											StartPos:  2627,
											EndPos:    2628,
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
					StartLine: 113,
					EndLine:   113,
					StartPos:  2634,
					EndPos:    2670,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 113,
						EndLine:   113,
						StartPos:  2640,
						EndPos:    2642,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Position: &position.Position{
							StartLine: 113,
							EndLine:   113,
							StartPos:  2645,
							EndPos:    2669,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 113,
									EndLine:   113,
									StartPos:  2645,
									EndPos:    2650,
								},
								Value: "public",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 113,
									EndLine:   113,
									StartPos:  2652,
									EndPos:    2657,
								},
								Value: "static",
							},
						},
						Properties: []node.Node{
							&stmt.Property{
								Position: &position.Position{
									StartLine: 113,
									EndLine:   113,
									StartPos:  2659,
									EndPos:    2660,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 113,
										EndLine:   113,
										StartPos:  2659,
										EndPos:    2660,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 113,
											EndLine:   113,
											StartPos:  2659,
											EndPos:    2660,
										},
										Value: "a",
									},
								},
							},
							&stmt.Property{
								Position: &position.Position{
									StartLine: 113,
									EndLine:   113,
									StartPos:  2663,
									EndPos:    2668,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 113,
										EndLine:   113,
										StartPos:  2663,
										EndPos:    2664,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 113,
											EndLine:   113,
											StartPos:  2663,
											EndPos:    2664,
										},
										Value: "b",
									},
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 113,
										EndLine:   113,
										StartPos:  2668,
										EndPos:    2668,
									},
									Value: "1",
								},
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 114,
					EndLine:   114,
					StartPos:  2674,
					EndPos:    2710,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 114,
						EndLine:   114,
						StartPos:  2680,
						EndPos:    2682,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Position: &position.Position{
							StartLine: 114,
							EndLine:   114,
							StartPos:  2685,
							EndPos:    2709,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 114,
									EndLine:   114,
									StartPos:  2685,
									EndPos:    2690,
								},
								Value: "public",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 114,
									EndLine:   114,
									StartPos:  2692,
									EndPos:    2697,
								},
								Value: "static",
							},
						},
						Properties: []node.Node{
							&stmt.Property{
								Position: &position.Position{
									StartLine: 114,
									EndLine:   114,
									StartPos:  2699,
									EndPos:    2704,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2699,
										EndPos:    2700,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 114,
											EndLine:   114,
											StartPos:  2699,
											EndPos:    2700,
										},
										Value: "a",
									},
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2704,
										EndPos:    2704,
									},
									Value: "1",
								},
							},
							&stmt.Property{
								Position: &position.Position{
									StartLine: 114,
									EndLine:   114,
									StartPos:  2707,
									EndPos:    2708,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2707,
										EndPos:    2708,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 114,
											EndLine:   114,
											StartPos:  2707,
											EndPos:    2708,
										},
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 115,
					EndLine:   115,
					StartPos:  2714,
					EndPos:    2731,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 115,
							EndLine:   115,
							StartPos:  2721,
							EndPos:    2722,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 115,
								EndLine:   115,
								StartPos:  2721,
								EndPos:    2722,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 115,
									EndLine:   115,
									StartPos:  2721,
									EndPos:    2722,
								},
								Value: "a",
							},
						},
					},
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 115,
							EndLine:   115,
							StartPos:  2725,
							EndPos:    2730,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 115,
								EndLine:   115,
								StartPos:  2725,
								EndPos:    2726,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 115,
									EndLine:   115,
									StartPos:  2725,
									EndPos:    2726,
								},
								Value: "b",
							},
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 115,
								EndLine:   115,
								StartPos:  2730,
								EndPos:    2730,
							},
							Value: "1",
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 116,
					EndLine:   116,
					StartPos:  2735,
					EndPos:    2752,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 116,
							EndLine:   116,
							StartPos:  2742,
							EndPos:    2747,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 116,
								EndLine:   116,
								StartPos:  2742,
								EndPos:    2743,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 116,
									EndLine:   116,
									StartPos:  2742,
									EndPos:    2743,
								},
								Value: "a",
							},
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 116,
								EndLine:   116,
								StartPos:  2747,
								EndPos:    2747,
							},
							Value: "1",
						},
					},
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 116,
							EndLine:   116,
							StartPos:  2750,
							EndPos:    2751,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 116,
								EndLine:   116,
								StartPos:  2750,
								EndPos:    2751,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 116,
									EndLine:   116,
									StartPos:  2750,
									EndPos:    2751,
								},
								Value: "b",
							},
						},
					},
				},
			},
			&stmt.AltSwitch{
				Position: &position.Position{
					StartLine: 118,
					EndLine:   122,
					StartPos:  2757,
					EndPos:    2815,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 118,
						EndLine:   118,
						StartPos:  2765,
						EndPos:    2765,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 119,
						EndLine:   -1,
						StartPos:  2773,
						EndPos:    -1,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 119,
								EndLine:   -1,
								StartPos:  2773,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 119,
									EndLine:   119,
									StartPos:  2778,
									EndPos:    2778,
								},
								Value: "1",
							},
							Stmts: []node.Node{},
						},
						&stmt.Default{
							Position: &position.Position{
								StartLine: 120,
								EndLine:   -1,
								StartPos:  2784,
								EndPos:    -1,
							},
							Stmts: []node.Node{},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 121,
								EndLine:   -1,
								StartPos:  2796,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 121,
									EndLine:   121,
									StartPos:  2801,
									EndPos:    2801,
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
					StartLine: 124,
					EndLine:   127,
					StartPos:  2820,
					EndPos:    2867,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 124,
						EndLine:   124,
						StartPos:  2828,
						EndPos:    2828,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 125,
						EndLine:   -1,
						StartPos:  2837,
						EndPos:    -1,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 125,
								EndLine:   -1,
								StartPos:  2837,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 125,
									EndLine:   125,
									StartPos:  2842,
									EndPos:    2842,
								},
								Value: "1",
							},
							Stmts: []node.Node{},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 126,
								EndLine:   -1,
								StartPos:  2848,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 126,
									EndLine:   126,
									StartPos:  2853,
									EndPos:    2853,
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
					StartLine: 129,
					EndLine:   132,
					StartPos:  2874,
					EndPos:    2925,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 129,
						EndLine:   129,
						StartPos:  2882,
						EndPos:    2882,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 129,
						EndLine:   132,
						StartPos:  2885,
						EndPos:    2925,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 130,
								EndLine:   130,
								StartPos:  2890,
								EndPos:    2903,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 130,
									EndLine:   130,
									StartPos:  2895,
									EndPos:    2895,
								},
								Value: "1",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 130,
										EndLine:   130,
										StartPos:  2898,
										EndPos:    2903,
									},
								},
							},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 131,
								EndLine:   131,
								StartPos:  2908,
								EndPos:    2921,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 131,
									EndLine:   131,
									StartPos:  2913,
									EndPos:    2913,
								},
								Value: "2",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 131,
										EndLine:   131,
										StartPos:  2916,
										EndPos:    2921,
									},
								},
							},
						},
					},
				},
			},
			&stmt.Switch{
				Position: &position.Position{
					StartLine: 134,
					EndLine:   137,
					StartPos:  2932,
					EndPos:    2984,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 134,
						EndLine:   134,
						StartPos:  2940,
						EndPos:    2940,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 134,
						EndLine:   137,
						StartPos:  2943,
						EndPos:    2984,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 135,
								EndLine:   135,
								StartPos:  2949,
								EndPos:    2962,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 135,
									EndLine:   135,
									StartPos:  2954,
									EndPos:    2954,
								},
								Value: "1",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 135,
										EndLine:   135,
										StartPos:  2957,
										EndPos:    2962,
									},
								},
							},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 136,
								EndLine:   136,
								StartPos:  2967,
								EndPos:    2980,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 136,
									EndLine:   136,
									StartPos:  2972,
									EndPos:    2972,
								},
								Value: "2",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 136,
										EndLine:   136,
										StartPos:  2975,
										EndPos:    2980,
									},
								},
							},
						},
					},
				},
			},
			&stmt.Throw{
				Position: &position.Position{
					StartLine: 138,
					EndLine:   138,
					StartPos:  2988,
					EndPos:    2996,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 138,
						EndLine:   138,
						StartPos:  2994,
						EndPos:    2995,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 138,
							EndLine:   138,
							StartPos:  2994,
							EndPos:    2995,
						},
						Value: "e",
					},
				},
			},
			&stmt.Trait{
				Position: &position.Position{
					StartLine: 139,
					EndLine:   139,
					StartPos:  3000,
					EndPos:    3011,
				},
				PhpDocComment: "",
				TraitName: &node.Identifier{
					Position: &position.Position{
						StartLine: 139,
						EndLine:   139,
						StartPos:  3006,
						EndPos:    3008,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 140,
					EndLine:   140,
					StartPos:  3015,
					EndPos:    3036,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 140,
						EndLine:   140,
						StartPos:  3021,
						EndPos:    3023,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 140,
							EndLine:   140,
							StartPos:  3027,
							EndPos:    3034,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 140,
									EndLine:   140,
									StartPos:  3031,
									EndPos:    3033,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 140,
											EndLine:   140,
											StartPos:  3031,
											EndPos:    3033,
										},
										Value: "Bar",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.Nop{
							Position: &position.Position{
								StartLine: 140,
								EndLine:   140,
								StartPos:  3034,
								EndPos:    3034,
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 141,
					EndLine:   141,
					StartPos:  3040,
					EndPos:    3068,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 141,
						EndLine:   141,
						StartPos:  3046,
						EndPos:    3048,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 141,
							EndLine:   141,
							StartPos:  3052,
							EndPos:    3066,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 141,
									EndLine:   141,
									StartPos:  3056,
									EndPos:    3058,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 141,
											EndLine:   141,
											StartPos:  3056,
											EndPos:    3058,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 141,
									EndLine:   141,
									StartPos:  3061,
									EndPos:    3063,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 141,
											EndLine:   141,
											StartPos:  3061,
											EndPos:    3063,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 141,
								EndLine:   141,
								StartPos:  3065,
								EndPos:    3066,
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 142,
					EndLine:   142,
					StartPos:  3072,
					EndPos:    3116,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 142,
						EndLine:   142,
						StartPos:  3078,
						EndPos:    3080,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 142,
							EndLine:   142,
							StartPos:  3084,
							EndPos:    3114,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 142,
									EndLine:   142,
									StartPos:  3088,
									EndPos:    3090,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 142,
											EndLine:   142,
											StartPos:  3088,
											EndPos:    3090,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 142,
									EndLine:   142,
									StartPos:  3093,
									EndPos:    3095,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 142,
											EndLine:   142,
											StartPos:  3093,
											EndPos:    3095,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 142,
								EndLine:   142,
								StartPos:  3097,
								EndPos:    3114,
							},
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 142,
										EndLine:   142,
										StartPos:  3099,
										EndPos:    3111,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 142,
											EndLine:   142,
											StartPos:  3099,
											EndPos:    3101,
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 142,
												EndLine:   142,
												StartPos:  3099,
												EndPos:    3101,
											},
											Value: "one",
										},
									},
									Modifier: &node.Identifier{
										Position: &position.Position{
											StartLine: 142,
											EndLine:   142,
											StartPos:  3106,
											EndPos:    3111,
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
					StartLine: 143,
					EndLine:   143,
					StartPos:  3120,
					EndPos:    3168,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 143,
						EndLine:   143,
						StartPos:  3126,
						EndPos:    3128,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 143,
							EndLine:   143,
							StartPos:  3132,
							EndPos:    3166,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 143,
									EndLine:   143,
									StartPos:  3136,
									EndPos:    3138,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 143,
											EndLine:   143,
											StartPos:  3136,
											EndPos:    3138,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 143,
									EndLine:   143,
									StartPos:  3141,
									EndPos:    3143,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 143,
											EndLine:   143,
											StartPos:  3141,
											EndPos:    3143,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 143,
								EndLine:   143,
								StartPos:  3145,
								EndPos:    3166,
							},
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 143,
										EndLine:   143,
										StartPos:  3147,
										EndPos:    3163,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 143,
											EndLine:   143,
											StartPos:  3147,
											EndPos:    3149,
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 143,
												EndLine:   143,
												StartPos:  3147,
												EndPos:    3149,
											},
											Value: "one",
										},
									},
									Modifier: &node.Identifier{
										Position: &position.Position{
											StartLine: 143,
											EndLine:   143,
											StartPos:  3154,
											EndPos:    3159,
										},
										Value: "public",
									},
									Alias: &node.Identifier{
										Position: &position.Position{
											StartLine: 143,
											EndLine:   143,
											StartPos:  3161,
											EndPos:    3163,
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
					StartLine: 144,
					EndLine:   144,
					StartPos:  3172,
					EndPos:    3248,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 144,
						EndLine:   144,
						StartPos:  3178,
						EndPos:    3180,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 144,
							EndLine:   144,
							StartPos:  3184,
							EndPos:    3246,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 144,
									EndLine:   144,
									StartPos:  3188,
									EndPos:    3190,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 144,
											EndLine:   144,
											StartPos:  3188,
											EndPos:    3190,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 144,
									EndLine:   144,
									StartPos:  3193,
									EndPos:    3195,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 144,
											EndLine:   144,
											StartPos:  3193,
											EndPos:    3195,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 144,
								EndLine:   144,
								StartPos:  3197,
								EndPos:    3246,
							},
							Adaptations: []node.Node{
								&stmt.TraitUsePrecedence{
									Position: &position.Position{
										StartLine: 144,
										EndLine:   144,
										StartPos:  3199,
										EndPos:    3226,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 144,
											EndLine:   144,
											StartPos:  3199,
											EndPos:    3206,
										},
										Trait: &name.Name{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  3199,
												EndPos:    3201,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 144,
														EndLine:   144,
														StartPos:  3199,
														EndPos:    3201,
													},
													Value: "Bar",
												},
											},
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  3204,
												EndPos:    3206,
											},
											Value: "one",
										},
									},
									Insteadof: []node.Node{
										&name.Name{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  3218,
												EndPos:    3220,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 144,
														EndLine:   144,
														StartPos:  3218,
														EndPos:    3220,
													},
													Value: "Baz",
												},
											},
										},
										&name.Name{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  3223,
												EndPos:    3226,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 144,
														EndLine:   144,
														StartPos:  3223,
														EndPos:    3226,
													},
													Value: "Quux",
												},
											},
										},
									},
								},
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 144,
										EndLine:   144,
										StartPos:  3229,
										EndPos:    3243,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 144,
											EndLine:   144,
											StartPos:  3229,
											EndPos:    3236,
										},
										Trait: &name.Name{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  3229,
												EndPos:    3231,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 144,
														EndLine:   144,
														StartPos:  3229,
														EndPos:    3231,
													},
													Value: "Baz",
												},
											},
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  3234,
												EndPos:    3236,
											},
											Value: "one",
										},
									},
									Alias: &node.Identifier{
										Position: &position.Position{
											StartLine: 144,
											EndLine:   144,
											StartPos:  3241,
											EndPos:    3243,
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
					StartLine: 146,
					EndLine:   -1,
					StartPos:  3253,
					EndPos:    -1,
				},
				Stmts:   []node.Node{},
				Catches: []node.Node{},
			},
			&stmt.Try{
				Position: &position.Position{
					StartLine: 147,
					EndLine:   147,
					StartPos:  3262,
					EndPos:    3291,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 147,
							EndLine:   147,
							StartPos:  3269,
							EndPos:    3291,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 147,
									EndLine:   147,
									StartPos:  3276,
									EndPos:    3284,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 147,
											EndLine:   147,
											StartPos:  3276,
											EndPos:    3284,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 147,
								EndLine:   147,
								StartPos:  3286,
								EndPos:    3287,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 147,
									EndLine:   147,
									StartPos:  3286,
									EndPos:    3287,
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
					StartLine: 148,
					EndLine:   148,
					StartPos:  3295,
					EndPos:    3355,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 148,
							EndLine:   148,
							StartPos:  3302,
							EndPos:    3324,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 148,
									EndLine:   148,
									StartPos:  3309,
									EndPos:    3317,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 148,
											EndLine:   148,
											StartPos:  3309,
											EndPos:    3317,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 148,
								EndLine:   148,
								StartPos:  3319,
								EndPos:    3320,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 148,
									EndLine:   148,
									StartPos:  3319,
									EndPos:    3320,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 148,
							EndLine:   148,
							StartPos:  3326,
							EndPos:    3355,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 148,
									EndLine:   148,
									StartPos:  3333,
									EndPos:    3348,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 148,
											EndLine:   148,
											StartPos:  3333,
											EndPos:    3348,
										},
										Value: "RuntimeException",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 148,
								EndLine:   148,
								StartPos:  3350,
								EndPos:    3351,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 148,
									EndLine:   148,
									StartPos:  3350,
									EndPos:    3351,
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
					StartLine: 149,
					EndLine:   149,
					StartPos:  3359,
					EndPos:    3462,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 149,
							EndLine:   149,
							StartPos:  3366,
							EndPos:    3388,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3373,
									EndPos:    3381,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 149,
											EndLine:   149,
											StartPos:  3373,
											EndPos:    3381,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 149,
								EndLine:   149,
								StartPos:  3383,
								EndPos:    3384,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3383,
									EndPos:    3384,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 149,
							EndLine:   149,
							StartPos:  3390,
							EndPos:    3420,
						},
						Types: []node.Node{
							&name.FullyQualified{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3397,
									EndPos:    3413,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 149,
											EndLine:   149,
											StartPos:  3398,
											EndPos:    3413,
										},
										Value: "RuntimeException",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 149,
								EndLine:   149,
								StartPos:  3415,
								EndPos:    3416,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3415,
									EndPos:    3416,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 149,
							EndLine:   149,
							StartPos:  3422,
							EndPos:    3462,
						},
						Types: []node.Node{
							&name.Relative{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3429,
									EndPos:    3455,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 149,
											EndLine:   149,
											StartPos:  3439,
											EndPos:    3455,
										},
										Value: "AdditionException",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 149,
								EndLine:   149,
								StartPos:  3457,
								EndPos:    3458,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3457,
									EndPos:    3458,
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
					StartLine: 150,
					EndLine:   150,
					StartPos:  3466,
					EndPos:    3506,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 150,
							EndLine:   150,
							StartPos:  3473,
							EndPos:    3495,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 150,
									EndLine:   150,
									StartPos:  3480,
									EndPos:    3488,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 150,
											EndLine:   150,
											StartPos:  3480,
											EndPos:    3488,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 150,
								EndLine:   150,
								StartPos:  3490,
								EndPos:    3491,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 150,
									EndLine:   150,
									StartPos:  3490,
									EndPos:    3491,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
				},
				Finally: &stmt.Finally{
					Position: &position.Position{
						StartLine: 150,
						EndLine:   150,
						StartPos:  3497,
						EndPos:    3506,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Unset{
				Position: &position.Position{
					StartLine: 152,
					EndLine:   152,
					StartPos:  3511,
					EndPos:    3524,
				},
				Vars: []node.Node{
					&expr.Variable{
						Position: &position.Position{
							StartLine: 152,
							EndLine:   152,
							StartPos:  3517,
							EndPos:    3518,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 152,
								EndLine:   152,
								StartPos:  3517,
								EndPos:    3518,
							},
							Value: "a",
						},
					},
					&expr.Variable{
						Position: &position.Position{
							StartLine: 152,
							EndLine:   152,
							StartPos:  3521,
							EndPos:    3522,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 152,
								EndLine:   152,
								StartPos:  3521,
								EndPos:    3522,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 154,
					EndLine:   154,
					StartPos:  3529,
					EndPos:    3536,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 154,
							EndLine:   154,
							StartPos:  3533,
							EndPos:    3535,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 154,
								EndLine:   154,
								StartPos:  3533,
								EndPos:    3535,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 154,
										EndLine:   154,
										StartPos:  3533,
										EndPos:    3535,
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
					StartLine: 155,
					EndLine:   155,
					StartPos:  3540,
					EndPos:    3548,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 155,
							EndLine:   155,
							StartPos:  3545,
							EndPos:    3547,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 155,
								EndLine:   155,
								StartPos:  3545,
								EndPos:    3547,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 155,
										EndLine:   155,
										StartPos:  3545,
										EndPos:    3547,
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
					StartLine: 156,
					EndLine:   156,
					StartPos:  3552,
					EndPos:    3567,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 156,
							EndLine:   156,
							StartPos:  3557,
							EndPos:    3566,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 156,
								EndLine:   156,
								StartPos:  3557,
								EndPos:    3559,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 156,
										EndLine:   156,
										StartPos:  3557,
										EndPos:    3559,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 156,
								EndLine:   156,
								StartPos:  3564,
								EndPos:    3566,
							},
							Value: "Bar",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 157,
					EndLine:   157,
					StartPos:  3571,
					EndPos:    3583,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 157,
							EndLine:   157,
							StartPos:  3575,
							EndPos:    3577,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 157,
								EndLine:   157,
								StartPos:  3575,
								EndPos:    3577,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 157,
										EndLine:   157,
										StartPos:  3575,
										EndPos:    3577,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 157,
							EndLine:   157,
							StartPos:  3580,
							EndPos:    3582,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 157,
								EndLine:   157,
								StartPos:  3580,
								EndPos:    3582,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 157,
										EndLine:   157,
										StartPos:  3580,
										EndPos:    3582,
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
					StartLine: 158,
					EndLine:   158,
					StartPos:  3587,
					EndPos:    3606,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 158,
							EndLine:   158,
							StartPos:  3591,
							EndPos:    3593,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 158,
								EndLine:   158,
								StartPos:  3591,
								EndPos:    3593,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 158,
										EndLine:   158,
										StartPos:  3591,
										EndPos:    3593,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 158,
							EndLine:   158,
							StartPos:  3596,
							EndPos:    3605,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 158,
								EndLine:   158,
								StartPos:  3596,
								EndPos:    3598,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 158,
										EndLine:   158,
										StartPos:  3596,
										EndPos:    3598,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 158,
								EndLine:   158,
								StartPos:  3603,
								EndPos:    3605,
							},
							Value: "Baz",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 159,
					EndLine:   159,
					StartPos:  3610,
					EndPos:    3632,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 159,
						EndLine:   159,
						StartPos:  3614,
						EndPos:    3621,
					},
					Value: "function",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 159,
							EndLine:   159,
							StartPos:  3623,
							EndPos:    3625,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 159,
								EndLine:   159,
								StartPos:  3623,
								EndPos:    3625,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 159,
										EndLine:   159,
										StartPos:  3623,
										EndPos:    3625,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 159,
							EndLine:   159,
							StartPos:  3629,
							EndPos:    3631,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 159,
								EndLine:   159,
								StartPos:  3629,
								EndPos:    3631,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 159,
										EndLine:   159,
										StartPos:  3629,
										EndPos:    3631,
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
					StartLine: 160,
					EndLine:   160,
					StartPos:  3636,
					EndPos:    3672,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 160,
						EndLine:   160,
						StartPos:  3640,
						EndPos:    3647,
					},
					Value: "function",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 160,
							EndLine:   160,
							StartPos:  3649,
							EndPos:    3658,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 160,
								EndLine:   160,
								StartPos:  3649,
								EndPos:    3651,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 160,
										EndLine:   160,
										StartPos:  3649,
										EndPos:    3651,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 160,
								EndLine:   160,
								StartPos:  3656,
								EndPos:    3658,
							},
							Value: "foo",
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 160,
							EndLine:   160,
							StartPos:  3662,
							EndPos:    3671,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 160,
								EndLine:   160,
								StartPos:  3662,
								EndPos:    3664,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 160,
										EndLine:   160,
										StartPos:  3662,
										EndPos:    3664,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 160,
								EndLine:   160,
								StartPos:  3669,
								EndPos:    3671,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 161,
					EndLine:   161,
					StartPos:  3676,
					EndPos:    3695,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 161,
						EndLine:   161,
						StartPos:  3680,
						EndPos:    3684,
					},
					Value: "const",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 161,
							EndLine:   161,
							StartPos:  3686,
							EndPos:    3688,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 161,
								EndLine:   161,
								StartPos:  3686,
								EndPos:    3688,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 161,
										EndLine:   161,
										StartPos:  3686,
										EndPos:    3688,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 161,
							EndLine:   161,
							StartPos:  3692,
							EndPos:    3694,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 161,
								EndLine:   161,
								StartPos:  3692,
								EndPos:    3694,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 161,
										EndLine:   161,
										StartPos:  3692,
										EndPos:    3694,
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
					StartLine: 162,
					EndLine:   162,
					StartPos:  3699,
					EndPos:    3732,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 162,
						EndLine:   162,
						StartPos:  3703,
						EndPos:    3707,
					},
					Value: "const",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 162,
							EndLine:   162,
							StartPos:  3709,
							EndPos:    3718,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3709,
								EndPos:    3711,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 162,
										EndLine:   162,
										StartPos:  3709,
										EndPos:    3711,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3716,
								EndPos:    3718,
							},
							Value: "foo",
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 162,
							EndLine:   162,
							StartPos:  3722,
							EndPos:    3731,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3722,
								EndPos:    3724,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 162,
										EndLine:   162,
										StartPos:  3722,
										EndPos:    3724,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3729,
								EndPos:    3731,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 164,
					EndLine:   164,
					StartPos:  3737,
					EndPos:    3742,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 164,
						EndLine:   164,
						StartPos:  3737,
						EndPos:    3741,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 164,
							EndLine:   164,
							StartPos:  3737,
							EndPos:    3738,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 164,
								EndLine:   164,
								StartPos:  3737,
								EndPos:    3738,
							},
							Value: "a",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 164,
							EndLine:   164,
							StartPos:  3740,
							EndPos:    3740,
						},
						Value: "1",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 165,
					EndLine:   165,
					StartPos:  3746,
					EndPos:    3754,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 165,
						EndLine:   165,
						StartPos:  3746,
						EndPos:    3753,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 165,
							EndLine:   165,
							StartPos:  3746,
							EndPos:    3750,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 165,
								EndLine:   165,
								StartPos:  3746,
								EndPos:    3747,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 165,
									EndLine:   165,
									StartPos:  3746,
									EndPos:    3747,
								},
								Value: "a",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 165,
								EndLine:   165,
								StartPos:  3749,
								EndPos:    3749,
							},
							Value: "1",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 165,
							EndLine:   165,
							StartPos:  3752,
							EndPos:    3752,
						},
						Value: "2",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 166,
					EndLine:   166,
					StartPos:  3758,
					EndPos:    3765,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 166,
						EndLine:   166,
						StartPos:  3758,
						EndPos:    3764,
					},
					Items: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 167,
					EndLine:   167,
					StartPos:  3769,
					EndPos:    3777,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 167,
						EndLine:   167,
						StartPos:  3769,
						EndPos:    3776,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 167,
								EndLine:   167,
								StartPos:  3775,
								EndPos:    3775,
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 167,
									EndLine:   167,
									StartPos:  3775,
									EndPos:    3775,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 168,
					EndLine:   168,
					StartPos:  3781,
					EndPos:    3798,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 168,
						EndLine:   168,
						StartPos:  3781,
						EndPos:    3797,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 168,
								EndLine:   168,
								StartPos:  3787,
								EndPos:    3790,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 168,
									EndLine:   168,
									StartPos:  3787,
									EndPos:    3787,
								},
								Value: "1",
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 168,
									EndLine:   168,
									StartPos:  3790,
									EndPos:    3790,
								},
								Value: "1",
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 168,
								EndLine:   168,
								StartPos:  3793,
								EndPos:    3795,
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 168,
									EndLine:   168,
									StartPos:  3793,
									EndPos:    3795,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 168,
										EndLine:   168,
										StartPos:  3794,
										EndPos:    3795,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 168,
											EndLine:   168,
											StartPos:  3794,
											EndPos:    3795,
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
					StartLine: 169,
					EndLine:   169,
					StartPos:  3802,
					EndPos:    3816,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 169,
						EndLine:   169,
						StartPos:  3802,
						EndPos:    3815,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 169,
								EndLine:   169,
								StartPos:  3808,
								EndPos:    3814,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 169,
									EndLine:   169,
									StartPos:  3808,
									EndPos:    3808,
								},
								Value: "3",
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 169,
									EndLine:   169,
									StartPos:  3812,
									EndPos:    3814,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 169,
										EndLine:   169,
										StartPos:  3813,
										EndPos:    3814,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 169,
											EndLine:   169,
											StartPos:  3813,
											EndPos:    3814,
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
					StartLine: 170,
					EndLine:   170,
					StartPos:  3820,
					EndPos:    3848,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 170,
						EndLine:   170,
						StartPos:  3820,
						EndPos:    3847,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 170,
								EndLine:   170,
								StartPos:  3826,
								EndPos:    3828,
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3826,
									EndPos:    3828,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3827,
										EndPos:    3828,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 170,
											EndLine:   170,
											StartPos:  3827,
											EndPos:    3828,
										},
										Value: "b",
									},
								},
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 170,
								EndLine:   170,
								StartPos:  3831,
								EndPos:    3834,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3831,
									EndPos:    3831,
								},
								Value: "1",
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3834,
									EndPos:    3834,
								},
								Value: "1",
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 170,
								EndLine:   170,
								StartPos:  3837,
								EndPos:    3837,
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3837,
									EndPos:    3837,
								},
								Value: "1",
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 170,
								EndLine:   170,
								StartPos:  3840,
								EndPos:    3846,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3840,
									EndPos:    3840,
								},
								Value: "3",
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3844,
									EndPos:    3846,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3845,
										EndPos:    3846,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 170,
											EndLine:   170,
											StartPos:  3845,
											EndPos:    3846,
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
					StartLine: 171,
					EndLine:   171,
					StartPos:  3852,
					EndPos:    3855,
				},
				Expr: &expr.BitwiseNot{
					Position: &position.Position{
						StartLine: 171,
						EndLine:   171,
						StartPos:  3852,
						EndPos:    3854,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 171,
							EndLine:   171,
							StartPos:  3853,
							EndPos:    3854,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 171,
								EndLine:   171,
								StartPos:  3853,
								EndPos:    3854,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 172,
					EndLine:   172,
					StartPos:  3859,
					EndPos:    3862,
				},
				Expr: &expr.BooleanNot{
					Position: &position.Position{
						StartLine: 172,
						EndLine:   172,
						StartPos:  3859,
						EndPos:    3861,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 172,
							EndLine:   172,
							StartPos:  3860,
							EndPos:    3861,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3860,
								EndPos:    3861,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 174,
					EndLine:   174,
					StartPos:  3867,
					EndPos:    3875,
				},
				Expr: &expr.ClassConstFetch{
					Position: &position.Position{
						StartLine: 174,
						EndLine:   174,
						StartPos:  3867,
						EndPos:    3874,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 174,
							EndLine:   174,
							StartPos:  3867,
							EndPos:    3869,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 174,
									EndLine:   174,
									StartPos:  3867,
									EndPos:    3869,
								},
								Value: "Foo",
							},
						},
					},
					ConstantName: &node.Identifier{
						Position: &position.Position{
							StartLine: 174,
							EndLine:   174,
							StartPos:  3872,
							EndPos:    3874,
						},
						Value: "Bar",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 175,
					EndLine:   175,
					StartPos:  3879,
					EndPos:    3888,
				},
				Expr: &expr.Clone{
					Position: &position.Position{
						StartLine: 175,
						EndLine:   175,
						StartPos:  3879,
						EndPos:    3886,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 175,
							EndLine:   175,
							StartPos:  3885,
							EndPos:    3886,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3885,
								EndPos:    3886,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 176,
					EndLine:   176,
					StartPos:  3892,
					EndPos:    3900,
				},
				Expr: &expr.Clone{
					Position: &position.Position{
						StartLine: 176,
						EndLine:   176,
						StartPos:  3892,
						EndPos:    3899,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 176,
							EndLine:   176,
							StartPos:  3898,
							EndPos:    3899,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 176,
								EndLine:   176,
								StartPos:  3898,
								EndPos:    3899,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 177,
					EndLine:   177,
					StartPos:  3904,
					EndPos:    3916,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 177,
						EndLine:   177,
						StartPos:  3904,
						EndPos:    3915,
					},
					PhpDocComment: "",
					ReturnsRef:    false,
					Static:        false,
					Stmts:         []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 178,
					EndLine:   178,
					StartPos:  3920,
					EndPos:    3953,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 178,
						EndLine:   178,
						StartPos:  3920,
						EndPos:    3952,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 178,
								EndLine:   178,
								StartPos:  3929,
								EndPos:    3930,
							},
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 178,
									EndLine:   178,
									StartPos:  3929,
									EndPos:    3930,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 178,
										EndLine:   178,
										StartPos:  3929,
										EndPos:    3930,
									},
									Value: "a",
								},
							},
						},
						&node.Parameter{
							Position: &position.Position{
								StartLine: 178,
								EndLine:   178,
								StartPos:  3933,
								EndPos:    3934,
							},
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 178,
									EndLine:   178,
									StartPos:  3933,
									EndPos:    3934,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 178,
										EndLine:   178,
										StartPos:  3933,
										EndPos:    3934,
									},
									Value: "b",
								},
							},
						},
					},
					ClosureUse: &expr.ClosureUse{
						Position: &position.Position{
							StartLine: 178,
							EndLine:   178,
							StartPos:  3937,
							EndPos:    3949,
						},
						Uses: []node.Node{
							&expr.Variable{
								Position: &position.Position{
									StartLine: 178,
									EndLine:   178,
									StartPos:  3942,
									EndPos:    3943,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 178,
										EndLine:   178,
										StartPos:  3942,
										EndPos:    3943,
									},
									Value: "c",
								},
							},
							&expr.Reference{
								Position: &position.Position{
									StartLine: 178,
									EndLine:   178,
									StartPos:  3946,
									EndPos:    3948,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 178,
										EndLine:   178,
										StartPos:  3947,
										EndPos:    3948,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 178,
											EndLine:   178,
											StartPos:  3947,
											EndPos:    3948,
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
					StartLine: 179,
					EndLine:   179,
					StartPos:  3957,
					EndPos:    3990,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 179,
						EndLine:   179,
						StartPos:  3957,
						EndPos:    3989,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 179,
								EndLine:   179,
								StartPos:  3966,
								EndPos:    3967,
							},
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 179,
									EndLine:   179,
									StartPos:  3966,
									EndPos:    3967,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 179,
										EndLine:   179,
										StartPos:  3966,
										EndPos:    3967,
									},
									Value: "a",
								},
							},
						},
						&node.Parameter{
							Position: &position.Position{
								StartLine: 179,
								EndLine:   179,
								StartPos:  3970,
								EndPos:    3971,
							},
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 179,
									EndLine:   179,
									StartPos:  3970,
									EndPos:    3971,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 179,
										EndLine:   179,
										StartPos:  3970,
										EndPos:    3971,
									},
									Value: "b",
								},
							},
						},
					},
					ClosureUse: &expr.ClosureUse{
						Position: &position.Position{
							StartLine: 179,
							EndLine:   179,
							StartPos:  3974,
							EndPos:    3986,
						},
						Uses: []node.Node{
							&expr.Reference{
								Position: &position.Position{
									StartLine: 179,
									EndLine:   179,
									StartPos:  3979,
									EndPos:    3981,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 179,
										EndLine:   179,
										StartPos:  3980,
										EndPos:    3981,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 179,
											EndLine:   179,
											StartPos:  3980,
											EndPos:    3981,
										},
										Value: "c",
									},
								},
							},
							&expr.Variable{
								Position: &position.Position{
									StartLine: 179,
									EndLine:   179,
									StartPos:  3984,
									EndPos:    3985,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 179,
										EndLine:   179,
										StartPos:  3984,
										EndPos:    3985,
									},
									Value: "d",
								},
							},
						},
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 180,
					EndLine:   180,
					StartPos:  3994,
					EndPos:    4007,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 180,
						EndLine:   180,
						StartPos:  3994,
						EndPos:    4006,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Stmts:         []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 181,
					EndLine:   181,
					StartPos:  4011,
					EndPos:    4014,
				},
				Expr: &expr.ConstFetch{
					Position: &position.Position{
						StartLine: 181,
						EndLine:   181,
						StartPos:  4011,
						EndPos:    4013,
					},
					Constant: &name.Name{
						Position: &position.Position{
							StartLine: 181,
							EndLine:   181,
							StartPos:  4011,
							EndPos:    4013,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 181,
									EndLine:   181,
									StartPos:  4011,
									EndPos:    4013,
								},
								Value: "foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 182,
					EndLine:   182,
					StartPos:  4018,
					EndPos:    4031,
				},
				Expr: &expr.ConstFetch{
					Position: &position.Position{
						StartLine: 182,
						EndLine:   182,
						StartPos:  4018,
						EndPos:    4030,
					},
					Constant: &name.Relative{
						Position: &position.Position{
							StartLine: 182,
							EndLine:   182,
							StartPos:  4018,
							EndPos:    4030,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 182,
									EndLine:   182,
									StartPos:  4028,
									EndPos:    4030,
								},
								Value: "foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 183,
					EndLine:   183,
					StartPos:  4035,
					EndPos:    4039,
				},
				Expr: &expr.ConstFetch{
					Position: &position.Position{
						StartLine: 183,
						EndLine:   183,
						StartPos:  4035,
						EndPos:    4038,
					},
					Constant: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 183,
							EndLine:   183,
							StartPos:  4035,
							EndPos:    4038,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 183,
									EndLine:   183,
									StartPos:  4036,
									EndPos:    4038,
								},
								Value: "foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 185,
					EndLine:   185,
					StartPos:  4044,
					EndPos:    4053,
				},
				Expr: &expr.Empty{
					Position: &position.Position{
						StartLine: 185,
						EndLine:   185,
						StartPos:  4044,
						EndPos:    4052,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 185,
							EndLine:   185,
							StartPos:  4050,
							EndPos:    4051,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 185,
								EndLine:   185,
								StartPos:  4050,
								EndPos:    4051,
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
					StartPos:  4057,
					EndPos:    4067,
				},
				Expr: &expr.Empty{
					Position: &position.Position{
						StartLine: 186,
						EndLine:   186,
						StartPos:  4057,
						EndPos:    4066,
					},
					Expr: &expr.ConstFetch{
						Position: &position.Position{
							StartLine: 186,
							EndLine:   186,
							StartPos:  4063,
							EndPos:    4065,
						},
						Constant: &name.Name{
							Position: &position.Position{
								StartLine: 186,
								EndLine:   186,
								StartPos:  4063,
								EndPos:    4065,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 186,
										EndLine:   186,
										StartPos:  4063,
										EndPos:    4065,
									},
									Value: "Foo",
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 187,
					EndLine:   187,
					StartPos:  4071,
					EndPos:    4074,
				},
				Expr: &expr.ErrorSuppress{
					Position: &position.Position{
						StartLine: 187,
						EndLine:   187,
						StartPos:  4071,
						EndPos:    4073,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 187,
							EndLine:   187,
							StartPos:  4072,
							EndPos:    4073,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 187,
								EndLine:   187,
								StartPos:  4072,
								EndPos:    4073,
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
					StartPos:  4078,
					EndPos:    4086,
				},
				Expr: &expr.Eval{
					Position: &position.Position{
						StartLine: 188,
						EndLine:   188,
						StartPos:  4078,
						EndPos:    4085,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 188,
							EndLine:   188,
							StartPos:  4083,
							EndPos:    4084,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 188,
								EndLine:   188,
								StartPos:  4083,
								EndPos:    4084,
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
					StartPos:  4090,
					EndPos:    4094,
				},
				Expr: &expr.Exit{
					Position: &position.Position{
						StartLine: 189,
						EndLine:   189,
						StartPos:  4090,
						EndPos:    4093,
					},
					Die: false,
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 190,
					EndLine:   190,
					StartPos:  4098,
					EndPos:    4106,
				},
				Expr: &expr.Exit{
					Position: &position.Position{
						StartLine: 190,
						EndLine:   190,
						StartPos:  4098,
						EndPos:    4105,
					},
					Die: false,
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 190,
							EndLine:   190,
							StartPos:  4103,
							EndPos:    4104,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 190,
								EndLine:   190,
								StartPos:  4103,
								EndPos:    4104,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 191,
					EndLine:   191,
					StartPos:  4110,
					EndPos:    4115,
				},
				Expr: &expr.Exit{
					Position: &position.Position{
						StartLine: 191,
						EndLine:   191,
						StartPos:  4110,
						EndPos:    4114,
					},
					Die: true,
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 192,
					EndLine:   192,
					StartPos:  4119,
					EndPos:    4126,
				},
				Expr: &expr.Exit{
					Position: &position.Position{
						StartLine: 192,
						EndLine:   192,
						StartPos:  4119,
						EndPos:    4125,
					},
					Die: true,
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 192,
							EndLine:   192,
							StartPos:  4123,
							EndPos:    4124,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 192,
								EndLine:   192,
								StartPos:  4123,
								EndPos:    4124,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 193,
					EndLine:   193,
					StartPos:  4130,
					EndPos:    4135,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 193,
						EndLine:   193,
						StartPos:  4130,
						EndPos:    4134,
					},
					Function: &name.Name{
						Position: &position.Position{
							StartLine: 193,
							EndLine:   193,
							StartPos:  4130,
							EndPos:    4132,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 193,
									EndLine:   193,
									StartPos:  4130,
									EndPos:    4132,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 193,
							EndLine:   193,
							StartPos:  4133,
							EndPos:    4134,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 194,
					EndLine:   194,
					StartPos:  4139,
					EndPos:    4157,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 194,
						EndLine:   194,
						StartPos:  4139,
						EndPos:    4156,
					},
					Function: &name.Relative{
						Position: &position.Position{
							StartLine: 194,
							EndLine:   194,
							StartPos:  4139,
							EndPos:    4151,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 194,
									EndLine:   194,
									StartPos:  4149,
									EndPos:    4151,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 194,
							EndLine:   194,
							StartPos:  4152,
							EndPos:    4156,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 194,
									EndLine:   194,
									StartPos:  4154,
									EndPos:    4155,
								},
								Variadic:    false,
								IsReference: true,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 194,
										EndLine:   194,
										StartPos:  4154,
										EndPos:    4155,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 194,
											EndLine:   194,
											StartPos:  4154,
											EndPos:    4155,
										},
										Value: "a",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 195,
					EndLine:   195,
					StartPos:  4161,
					EndPos:    4169,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 195,
						EndLine:   195,
						StartPos:  4161,
						EndPos:    4168,
					},
					Function: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 195,
							EndLine:   195,
							StartPos:  4161,
							EndPos:    4164,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 195,
									EndLine:   195,
									StartPos:  4162,
									EndPos:    4164,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 195,
							EndLine:   195,
							StartPos:  4165,
							EndPos:    4168,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 195,
									EndLine:   195,
									StartPos:  4166,
									EndPos:    4167,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.ShortArray{
									Position: &position.Position{
										StartLine: 195,
										EndLine:   195,
										StartPos:  4166,
										EndPos:    4167,
									},
									Items: []node.Node{},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 196,
					EndLine:   196,
					StartPos:  4173,
					EndPos:    4187,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 196,
						EndLine:   196,
						StartPos:  4173,
						EndPos:    4186,
					},
					Function: &expr.Variable{
						Position: &position.Position{
							StartLine: 196,
							EndLine:   196,
							StartPos:  4173,
							EndPos:    4176,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 196,
								EndLine:   196,
								StartPos:  4173,
								EndPos:    4176,
							},
							Value: "foo",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 196,
							EndLine:   196,
							StartPos:  4177,
							EndPos:    4186,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 196,
									EndLine:   196,
									StartPos:  4178,
									EndPos:    4185,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Yield{
									Position: &position.Position{
										StartLine: 196,
										EndLine:   196,
										StartPos:  4178,
										EndPos:    4185,
									},
									Value: &expr.Variable{
										Position: &position.Position{
											StartLine: 196,
											EndLine:   196,
											StartPos:  4184,
											EndPos:    4185,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 196,
												EndLine:   196,
												StartPos:  4184,
												EndPos:    4185,
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
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 198,
					EndLine:   198,
					StartPos:  4192,
					EndPos:    4196,
				},
				Expr: &expr.PostDec{
					Position: &position.Position{
						StartLine: 198,
						EndLine:   198,
						StartPos:  4192,
						EndPos:    4195,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 198,
							EndLine:   198,
							StartPos:  4192,
							EndPos:    4193,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 198,
								EndLine:   198,
								StartPos:  4192,
								EndPos:    4193,
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
					StartPos:  4200,
					EndPos:    4204,
				},
				Expr: &expr.PostInc{
					Position: &position.Position{
						StartLine: 199,
						EndLine:   199,
						StartPos:  4200,
						EndPos:    4203,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 199,
							EndLine:   199,
							StartPos:  4200,
							EndPos:    4201,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 199,
								EndLine:   199,
								StartPos:  4200,
								EndPos:    4201,
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
					StartPos:  4208,
					EndPos:    4212,
				},
				Expr: &expr.PreDec{
					Position: &position.Position{
						StartLine: 200,
						EndLine:   200,
						StartPos:  4208,
						EndPos:    4211,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 200,
							EndLine:   200,
							StartPos:  4210,
							EndPos:    4211,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 200,
								EndLine:   200,
								StartPos:  4210,
								EndPos:    4211,
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
					StartPos:  4216,
					EndPos:    4220,
				},
				Expr: &expr.PreInc{
					Position: &position.Position{
						StartLine: 201,
						EndLine:   201,
						StartPos:  4216,
						EndPos:    4219,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 201,
							EndLine:   201,
							StartPos:  4218,
							EndPos:    4219,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 201,
								EndLine:   201,
								StartPos:  4218,
								EndPos:    4219,
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
					StartPos:  4225,
					EndPos:    4235,
				},
				Expr: &expr.Include{
					Position: &position.Position{
						StartLine: 203,
						EndLine:   203,
						StartPos:  4225,
						EndPos:    4234,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 203,
							EndLine:   203,
							StartPos:  4233,
							EndPos:    4234,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 203,
								EndLine:   203,
								StartPos:  4233,
								EndPos:    4234,
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
					StartPos:  4239,
					EndPos:    4254,
				},
				Expr: &expr.IncludeOnce{
					Position: &position.Position{
						StartLine: 204,
						EndLine:   204,
						StartPos:  4239,
						EndPos:    4253,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 204,
							EndLine:   204,
							StartPos:  4252,
							EndPos:    4253,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 204,
								EndLine:   204,
								StartPos:  4252,
								EndPos:    4253,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 205,
					EndLine:   205,
					StartPos:  4258,
					EndPos:    4268,
				},
				Expr: &expr.Require{
					Position: &position.Position{
						StartLine: 205,
						EndLine:   205,
						StartPos:  4258,
						EndPos:    4267,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 205,
							EndLine:   205,
							StartPos:  4266,
							EndPos:    4267,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 205,
								EndLine:   205,
								StartPos:  4266,
								EndPos:    4267,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 206,
					EndLine:   206,
					StartPos:  4272,
					EndPos:    4287,
				},
				Expr: &expr.RequireOnce{
					Position: &position.Position{
						StartLine: 206,
						EndLine:   206,
						StartPos:  4272,
						EndPos:    4286,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 206,
							EndLine:   206,
							StartPos:  4285,
							EndPos:    4286,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 206,
								EndLine:   206,
								StartPos:  4285,
								EndPos:    4286,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 208,
					EndLine:   208,
					StartPos:  4292,
					EndPos:    4309,
				},
				Expr: &expr.InstanceOf{
					Position: &position.Position{
						StartLine: 208,
						EndLine:   208,
						StartPos:  4292,
						EndPos:    4308,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 208,
							EndLine:   208,
							StartPos:  4292,
							EndPos:    4293,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 208,
								EndLine:   208,
								StartPos:  4292,
								EndPos:    4293,
							},
							Value: "a",
						},
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 208,
							EndLine:   208,
							StartPos:  4306,
							EndPos:    4308,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 208,
									EndLine:   208,
									StartPos:  4306,
									EndPos:    4308,
								},
								Value: "Foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 209,
					EndLine:   209,
					StartPos:  4313,
					EndPos:    4340,
				},
				Expr: &expr.InstanceOf{
					Position: &position.Position{
						StartLine: 209,
						EndLine:   209,
						StartPos:  4313,
						EndPos:    4339,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 209,
							EndLine:   209,
							StartPos:  4313,
							EndPos:    4314,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 209,
								EndLine:   209,
								StartPos:  4313,
								EndPos:    4314,
							},
							Value: "a",
						},
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 209,
							EndLine:   209,
							StartPos:  4327,
							EndPos:    4339,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 209,
									EndLine:   209,
									StartPos:  4337,
									EndPos:    4339,
								},
								Value: "Foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 210,
					EndLine:   210,
					StartPos:  4344,
					EndPos:    4362,
				},
				Expr: &expr.InstanceOf{
					Position: &position.Position{
						StartLine: 210,
						EndLine:   210,
						StartPos:  4344,
						EndPos:    4361,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 210,
							EndLine:   210,
							StartPos:  4344,
							EndPos:    4345,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 210,
								EndLine:   210,
								StartPos:  4344,
								EndPos:    4345,
							},
							Value: "a",
						},
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 210,
							EndLine:   210,
							StartPos:  4358,
							EndPos:    4361,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 210,
									EndLine:   210,
									StartPos:  4359,
									EndPos:    4361,
								},
								Value: "Foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 212,
					EndLine:   212,
					StartPos:  4367,
					EndPos:    4380,
				},
				Expr: &expr.Isset{
					Position: &position.Position{
						StartLine: 212,
						EndLine:   212,
						StartPos:  4367,
						EndPos:    4379,
					},
					Variables: []node.Node{
						&expr.Variable{
							Position: &position.Position{
								StartLine: 212,
								EndLine:   212,
								StartPos:  4373,
								EndPos:    4374,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 212,
									EndLine:   212,
									StartPos:  4373,
									EndPos:    4374,
								},
								Value: "a",
							},
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 212,
								EndLine:   212,
								StartPos:  4377,
								EndPos:    4378,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 212,
									EndLine:   212,
									StartPos:  4377,
									EndPos:    4378,
								},
								Value: "b",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 213,
					EndLine:   213,
					StartPos:  4384,
					EndPos:    4394,
				},
				Expr: &expr.Isset{
					Position: &position.Position{
						StartLine: 213,
						EndLine:   213,
						StartPos:  4384,
						EndPos:    4393,
					},
					Variables: []node.Node{
						&expr.ConstFetch{
							Position: &position.Position{
								StartLine: 213,
								EndLine:   213,
								StartPos:  4390,
								EndPos:    4392,
							},
							Constant: &name.Name{
								Position: &position.Position{
									StartLine: 213,
									EndLine:   213,
									StartPos:  4390,
									EndPos:    4392,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 213,
											EndLine:   213,
											StartPos:  4390,
											EndPos:    4392,
										},
										Value: "Foo",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 214,
					EndLine:   214,
					StartPos:  4398,
					EndPos:    4409,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 214,
						EndLine:   214,
						StartPos:  4398,
						EndPos:    4408,
					},
					Variable: &expr.List{
						Position: &position.Position{
							StartLine: 214,
							EndLine:   214,
							StartPos:  4398,
							EndPos:    4403,
						},
						Items: []node.Node{},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 214,
							EndLine:   214,
							StartPos:  4407,
							EndPos:    4408,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 214,
								EndLine:   214,
								StartPos:  4407,
								EndPos:    4408,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 215,
					EndLine:   215,
					StartPos:  4413,
					EndPos:    4430,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 215,
						EndLine:   215,
						StartPos:  4413,
						EndPos:    4429,
					},
					Variable: &expr.List{
						Position: &position.Position{
							StartLine: 215,
							EndLine:   215,
							StartPos:  4413,
							EndPos:    4424,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 215,
									EndLine:   215,
									StartPos:  4418,
									EndPos:    4419,
								},
								Val: &expr.Variable{
									Position: &position.Position{
										StartLine: 215,
										EndLine:   215,
										StartPos:  4418,
										EndPos:    4419,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 215,
											EndLine:   215,
											StartPos:  4418,
											EndPos:    4419,
										},
										Value: "a",
									},
								},
							},
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 215,
									EndLine:   215,
									StartPos:  4422,
									EndPos:    4423,
								},
								Val: &expr.Variable{
									Position: &position.Position{
										StartLine: 215,
										EndLine:   215,
										StartPos:  4422,
										EndPos:    4423,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 215,
											EndLine:   215,
											StartPos:  4422,
											EndPos:    4423,
										},
										Value: "b",
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 215,
							EndLine:   215,
							StartPos:  4428,
							EndPos:    4429,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 215,
								EndLine:   215,
								StartPos:  4428,
								EndPos:    4429,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 216,
					EndLine:   216,
					StartPos:  4434,
					EndPos:    4449,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 216,
						EndLine:   216,
						StartPos:  4434,
						EndPos:    4448,
					},
					Variable: &expr.List{
						Position: &position.Position{
							StartLine: 216,
							EndLine:   216,
							StartPos:  4434,
							EndPos:    4443,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 216,
									EndLine:   216,
									StartPos:  4439,
									EndPos:    4442,
								},
								Val: &expr.ArrayDimFetch{
									Position: &position.Position{
										StartLine: 216,
										EndLine:   216,
										StartPos:  4439,
										EndPos:    4442,
									},
									Variable: &expr.Variable{
										Position: &position.Position{
											StartLine: 216,
											EndLine:   216,
											StartPos:  4439,
											EndPos:    4440,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 216,
												EndLine:   216,
												StartPos:  4439,
												EndPos:    4440,
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
							StartLine: 216,
							EndLine:   216,
							StartPos:  4447,
							EndPos:    4448,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 216,
								EndLine:   216,
								StartPos:  4447,
								EndPos:    4448,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 217,
					EndLine:   217,
					StartPos:  4453,
					EndPos:    4472,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 217,
						EndLine:   217,
						StartPos:  4453,
						EndPos:    4471,
					},
					Variable: &expr.List{
						Position: &position.Position{
							StartLine: 217,
							EndLine:   217,
							StartPos:  4453,
							EndPos:    4466,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 217,
									EndLine:   217,
									StartPos:  4458,
									EndPos:    4465,
								},
								Val: &expr.List{
									Position: &position.Position{
										StartLine: 217,
										EndLine:   217,
										StartPos:  4458,
										EndPos:    4465,
									},
									Items: []node.Node{
										&expr.ArrayItem{
											Position: &position.Position{
												StartLine: 217,
												EndLine:   217,
												StartPos:  4463,
												EndPos:    4464,
											},
											Val: &expr.Variable{
												Position: &position.Position{
													StartLine: 217,
													EndLine:   217,
													StartPos:  4463,
													EndPos:    4464,
												},
												VarName: &node.Identifier{
													Position: &position.Position{
														StartLine: 217,
														EndLine:   217,
														StartPos:  4463,
														EndPos:    4464,
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
							StartLine: 217,
							EndLine:   217,
							StartPos:  4470,
							EndPos:    4471,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 217,
								EndLine:   217,
								StartPos:  4470,
								EndPos:    4471,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 219,
					EndLine:   219,
					StartPos:  4477,
					EndPos:    4486,
				},
				Expr: &expr.MethodCall{
					Position: &position.Position{
						StartLine: 219,
						EndLine:   219,
						StartPos:  4477,
						EndPos:    4485,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 219,
							EndLine:   219,
							StartPos:  4477,
							EndPos:    4478,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 219,
								EndLine:   219,
								StartPos:  4477,
								EndPos:    4478,
							},
							Value: "a",
						},
					},
					Method: &node.Identifier{
						Position: &position.Position{
							StartLine: 219,
							EndLine:   219,
							StartPos:  4481,
							EndPos:    4483,
						},
						Value: "foo",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 219,
							EndLine:   219,
							StartPos:  4484,
							EndPos:    4485,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 220,
					EndLine:   220,
					StartPos:  4490,
					EndPos:    4497,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 220,
						EndLine:   220,
						StartPos:  4490,
						EndPos:    4496,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 220,
							EndLine:   220,
							StartPos:  4494,
							EndPos:    4496,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 220,
									EndLine:   220,
									StartPos:  4494,
									EndPos:    4496,
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
					StartPos:  4501,
					EndPos:    4520,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 221,
						EndLine:   221,
						StartPos:  4501,
						EndPos:    4519,
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 221,
							EndLine:   221,
							StartPos:  4505,
							EndPos:    4517,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 221,
									EndLine:   221,
									StartPos:  4515,
									EndPos:    4517,
								},
								Value: "Foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 221,
							EndLine:   221,
							StartPos:  4518,
							EndPos:    4519,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 222,
					EndLine:   222,
					StartPos:  4524,
					EndPos:    4534,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 222,
						EndLine:   222,
						StartPos:  4524,
						EndPos:    4533,
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 222,
							EndLine:   222,
							StartPos:  4528,
							EndPos:    4531,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 222,
									EndLine:   222,
									StartPos:  4529,
									EndPos:    4531,
								},
								Value: "Foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 222,
							EndLine:   222,
							StartPos:  4532,
							EndPos:    4533,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 223,
					EndLine:   223,
					StartPos:  4538,
					EndPos:    4547,
				},
				Expr: &expr.Print{
					Position: &position.Position{
						StartLine: 223,
						EndLine:   223,
						StartPos:  4538,
						EndPos:    4545,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 223,
							EndLine:   223,
							StartPos:  4544,
							EndPos:    4545,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 223,
								EndLine:   223,
								StartPos:  4544,
								EndPos:    4545,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 224,
					EndLine:   224,
					StartPos:  4551,
					EndPos:    4558,
				},
				Expr: &expr.PropertyFetch{
					Position: &position.Position{
						StartLine: 224,
						EndLine:   224,
						StartPos:  4551,
						EndPos:    4557,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 224,
							EndLine:   224,
							StartPos:  4551,
							EndPos:    4552,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 224,
								EndLine:   224,
								StartPos:  4551,
								EndPos:    4552,
							},
							Value: "a",
						},
					},
					Property: &node.Identifier{
						Position: &position.Position{
							StartLine: 224,
							EndLine:   224,
							StartPos:  4555,
							EndPos:    4557,
						},
						Value: "foo",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 225,
					EndLine:   225,
					StartPos:  4562,
					EndPos:    4572,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 225,
						EndLine:   225,
						StartPos:  4562,
						EndPos:    4570,
					},
					Variable: &expr.PropertyFetch{
						Position: &position.Position{
							StartLine: 225,
							EndLine:   225,
							StartPos:  4562,
							EndPos:    4568,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 225,
								EndLine:   225,
								StartPos:  4562,
								EndPos:    4563,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 225,
									EndLine:   225,
									StartPos:  4562,
									EndPos:    4563,
								},
								Value: "a",
							},
						},
						Property: &node.Identifier{
							Position: &position.Position{
								StartLine: 225,
								EndLine:   225,
								StartPos:  4566,
								EndPos:    4568,
							},
							Value: "foo",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 225,
							EndLine:   225,
							StartPos:  4570,
							EndPos:    4570,
						},
						Value: "1",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 226,
					EndLine:   226,
					StartPos:  4576,
					EndPos:    4604,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 226,
						EndLine:   226,
						StartPos:  4576,
						EndPos:    4602,
					},
					Variable: &expr.PropertyFetch{
						Position: &position.Position{
							StartLine: 226,
							EndLine:   226,
							StartPos:  4576,
							EndPos:    4600,
						},
						Variable: &expr.MethodCall{
							Position: &position.Position{
								StartLine: 226,
								EndLine:   226,
								StartPos:  4576,
								EndPos:    4594,
							},
							Variable: &expr.PropertyFetch{
								Position: &position.Position{
									StartLine: 226,
									EndLine:   226,
									StartPos:  4576,
									EndPos:    4587,
								},
								Variable: &expr.PropertyFetch{
									Position: &position.Position{
										StartLine: 226,
										EndLine:   226,
										StartPos:  4576,
										EndPos:    4582,
									},
									Variable: &expr.Variable{
										Position: &position.Position{
											StartLine: 226,
											EndLine:   226,
											StartPos:  4576,
											EndPos:    4577,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 226,
												EndLine:   226,
												StartPos:  4576,
												EndPos:    4577,
											},
											Value: "a",
										},
									},
									Property: &node.Identifier{
										Position: &position.Position{
											StartLine: 226,
											EndLine:   226,
											StartPos:  4580,
											EndPos:    4582,
										},
										Value: "foo",
									},
								},
								Property: &node.Identifier{
									Position: &position.Position{
										StartLine: 226,
										EndLine:   226,
										StartPos:  4585,
										EndPos:    4587,
									},
									Value: "bar",
								},
							},
							Method: &node.Identifier{
								Position: &position.Position{
									StartLine: 226,
									EndLine:   226,
									StartPos:  4590,
									EndPos:    4592,
								},
								Value: "baz",
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 226,
									EndLine:   226,
									StartPos:  4593,
									EndPos:    4594,
								},
							},
						},
						Property: &node.Identifier{
							Position: &position.Position{
								StartLine: 226,
								EndLine:   226,
								StartPos:  4597,
								EndPos:    4600,
							},
							Value: "quux",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 226,
							EndLine:   226,
							StartPos:  4602,
							EndPos:    4602,
						},
						Value: "0",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 227,
					EndLine:   227,
					StartPos:  4608,
					EndPos:    4623,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 227,
						EndLine:   227,
						StartPos:  4608,
						EndPos:    4621,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 227,
							EndLine:   227,
							StartPos:  4608,
							EndPos:    4618,
						},
						Variable: &expr.MethodCall{
							Position: &position.Position{
								StartLine: 227,
								EndLine:   227,
								StartPos:  4608,
								EndPos:    4616,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 227,
									EndLine:   227,
									StartPos:  4608,
									EndPos:    4609,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 227,
										EndLine:   227,
										StartPos:  4608,
										EndPos:    4609,
									},
									Value: "a",
								},
							},
							Method: &node.Identifier{
								Position: &position.Position{
									StartLine: 227,
									EndLine:   227,
									StartPos:  4612,
									EndPos:    4614,
								},
								Value: "foo",
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 227,
									EndLine:   227,
									StartPos:  4615,
									EndPos:    4616,
								},
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 227,
								EndLine:   227,
								StartPos:  4618,
								EndPos:    4618,
							},
							Value: "1",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 227,
							EndLine:   227,
							StartPos:  4621,
							EndPos:    4621,
						},
						Value: "1",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 228,
					EndLine:   228,
					StartPos:  4627,
					EndPos:    4635,
				},
				Expr: &expr.ShellExec{
					Position: &position.Position{
						StartLine: 228,
						EndLine:   228,
						StartPos:  4627,
						EndPos:    4634,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 228,
								EndLine:   228,
								StartPos:  4628,
								EndPos:    4631,
							},
							Value: "cmd ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 228,
								EndLine:   228,
								StartPos:  4632,
								EndPos:    4633,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 228,
									EndLine:   228,
									StartPos:  4632,
									EndPos:    4633,
								},
								Value: "a",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 229,
					EndLine:   229,
					StartPos:  4639,
					EndPos:    4644,
				},
				Expr: &expr.ShellExec{
					Position: &position.Position{
						StartLine: 229,
						EndLine:   229,
						StartPos:  4639,
						EndPos:    4643,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Value: "cmd",
							Position: &position.Position{
								StartLine: 229,
								EndLine:   229,
								StartPos:  4640,
								EndPos:    4642,
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 230,
					EndLine:   230,
					StartPos:  4648,
					EndPos:    4650,
				},
				Expr: &expr.ShellExec{
					Position: &position.Position{
						StartLine: 230,
						EndLine:   230,
						StartPos:  4648,
						EndPos:    4649,
					},
					Parts: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 231,
					EndLine:   231,
					StartPos:  4654,
					EndPos:    4656,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 231,
						EndLine:   231,
						StartPos:  4654,
						EndPos:    4655,
					},
					Items: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 232,
					EndLine:   232,
					StartPos:  4660,
					EndPos:    4663,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 232,
						EndLine:   232,
						StartPos:  4660,
						EndPos:    4662,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 232,
								EndLine:   232,
								StartPos:  4661,
								EndPos:    4661,
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 232,
									EndLine:   232,
									StartPos:  4661,
									EndPos:    4661,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 233,
					EndLine:   233,
					StartPos:  4667,
					EndPos:    4679,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 233,
						EndLine:   233,
						StartPos:  4667,
						EndPos:    4678,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 233,
								EndLine:   233,
								StartPos:  4668,
								EndPos:    4671,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 233,
									EndLine:   233,
									StartPos:  4668,
									EndPos:    4668,
								},
								Value: "1",
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 233,
									EndLine:   233,
									StartPos:  4671,
									EndPos:    4671,
								},
								Value: "1",
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 233,
								EndLine:   233,
								StartPos:  4674,
								EndPos:    4676,
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 233,
									EndLine:   233,
									StartPos:  4674,
									EndPos:    4676,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 233,
										EndLine:   233,
										StartPos:  4675,
										EndPos:    4676,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 233,
											EndLine:   233,
											StartPos:  4675,
											EndPos:    4676,
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
					StartLine: 235,
					EndLine:   235,
					StartPos:  4684,
					EndPos:    4694,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 235,
						EndLine:   235,
						StartPos:  4684,
						EndPos:    4693,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 235,
							EndLine:   235,
							StartPos:  4684,
							EndPos:    4686,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 235,
									EndLine:   235,
									StartPos:  4684,
									EndPos:    4686,
								},
								Value: "Foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 235,
							EndLine:   235,
							StartPos:  4689,
							EndPos:    4691,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 235,
							EndLine:   235,
							StartPos:  4692,
							EndPos:    4693,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 236,
					EndLine:   236,
					StartPos:  4698,
					EndPos:    4718,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 236,
						EndLine:   236,
						StartPos:  4698,
						EndPos:    4717,
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 236,
							EndLine:   236,
							StartPos:  4698,
							EndPos:    4710,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 236,
									EndLine:   236,
									StartPos:  4708,
									EndPos:    4710,
								},
								Value: "Foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 236,
							EndLine:   236,
							StartPos:  4713,
							EndPos:    4715,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 236,
							EndLine:   236,
							StartPos:  4716,
							EndPos:    4717,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 237,
					EndLine:   237,
					StartPos:  4722,
					EndPos:    4733,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 237,
						EndLine:   237,
						StartPos:  4722,
						EndPos:    4732,
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 237,
							EndLine:   237,
							StartPos:  4722,
							EndPos:    4725,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 237,
									EndLine:   237,
									StartPos:  4723,
									EndPos:    4725,
								},
								Value: "Foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 237,
							EndLine:   237,
							StartPos:  4728,
							EndPos:    4730,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 237,
							EndLine:   237,
							StartPos:  4731,
							EndPos:    4732,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 238,
					EndLine:   238,
					StartPos:  4737,
					EndPos:    4748,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 238,
						EndLine:   238,
						StartPos:  4737,
						EndPos:    4747,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 238,
							EndLine:   238,
							StartPos:  4737,
							EndPos:    4739,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 238,
									EndLine:   238,
									StartPos:  4737,
									EndPos:    4739,
								},
								Value: "Foo",
							},
						},
					},
					Call: &expr.Variable{
						Position: &position.Position{
							StartLine: 238,
							EndLine:   238,
							StartPos:  4742,
							EndPos:    4745,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 238,
								EndLine:   238,
								StartPos:  4742,
								EndPos:    4745,
							},
							Value: "bar",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 238,
							EndLine:   238,
							StartPos:  4746,
							EndPos:    4747,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 239,
					EndLine:   239,
					StartPos:  4752,
					EndPos:    4764,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 239,
						EndLine:   239,
						StartPos:  4752,
						EndPos:    4763,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 239,
							EndLine:   239,
							StartPos:  4752,
							EndPos:    4755,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 239,
								EndLine:   239,
								StartPos:  4752,
								EndPos:    4755,
							},
							Value: "foo",
						},
					},
					Call: &expr.Variable{
						Position: &position.Position{
							StartLine: 239,
							EndLine:   239,
							StartPos:  4758,
							EndPos:    4761,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 239,
								EndLine:   239,
								StartPos:  4758,
								EndPos:    4761,
							},
							Value: "bar",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 239,
							EndLine:   239,
							StartPos:  4762,
							EndPos:    4763,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 240,
					EndLine:   240,
					StartPos:  4768,
					EndPos:    4777,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 240,
						EndLine:   240,
						StartPos:  4768,
						EndPos:    4776,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 240,
							EndLine:   240,
							StartPos:  4768,
							EndPos:    4770,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 240,
									EndLine:   240,
									StartPos:  4768,
									EndPos:    4770,
								},
								Value: "Foo",
							},
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 240,
							EndLine:   240,
							StartPos:  4773,
							EndPos:    4776,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 240,
								EndLine:   240,
								StartPos:  4773,
								EndPos:    4776,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 241,
					EndLine:   241,
					StartPos:  4781,
					EndPos:    4800,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 241,
						EndLine:   241,
						StartPos:  4781,
						EndPos:    4799,
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 241,
							EndLine:   241,
							StartPos:  4781,
							EndPos:    4793,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 241,
									EndLine:   241,
									StartPos:  4791,
									EndPos:    4793,
								},
								Value: "Foo",
							},
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 241,
							EndLine:   241,
							StartPos:  4796,
							EndPos:    4799,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 241,
								EndLine:   241,
								StartPos:  4796,
								EndPos:    4799,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 242,
					EndLine:   242,
					StartPos:  4804,
					EndPos:    4814,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 242,
						EndLine:   242,
						StartPos:  4804,
						EndPos:    4813,
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 242,
							EndLine:   242,
							StartPos:  4804,
							EndPos:    4807,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 242,
									EndLine:   242,
									StartPos:  4805,
									EndPos:    4807,
								},
								Value: "Foo",
							},
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 242,
							EndLine:   242,
							StartPos:  4810,
							EndPos:    4813,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 242,
								EndLine:   242,
								StartPos:  4810,
								EndPos:    4813,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 243,
					EndLine:   243,
					StartPos:  4818,
					EndPos:    4830,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 243,
						EndLine:   243,
						StartPos:  4818,
						EndPos:    4829,
					},
					Condition: &expr.Variable{
						Position: &position.Position{
							StartLine: 243,
							EndLine:   243,
							StartPos:  4818,
							EndPos:    4819,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 243,
								EndLine:   243,
								StartPos:  4818,
								EndPos:    4819,
							},
							Value: "a",
						},
					},
					IfTrue: &expr.Variable{
						Position: &position.Position{
							StartLine: 243,
							EndLine:   243,
							StartPos:  4823,
							EndPos:    4824,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 243,
								EndLine:   243,
								StartPos:  4823,
								EndPos:    4824,
							},
							Value: "b",
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 243,
							EndLine:   243,
							StartPos:  4828,
							EndPos:    4829,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 243,
								EndLine:   243,
								StartPos:  4828,
								EndPos:    4829,
							},
							Value: "c",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 244,
					EndLine:   244,
					StartPos:  4834,
					EndPos:    4843,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 244,
						EndLine:   244,
						StartPos:  4834,
						EndPos:    4842,
					},
					Condition: &expr.Variable{
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
							Value: "a",
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 244,
							EndLine:   244,
							StartPos:  4841,
							EndPos:    4842,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 244,
								EndLine:   244,
								StartPos:  4841,
								EndPos:    4842,
							},
							Value: "c",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 245,
					EndLine:   245,
					StartPos:  4847,
					EndPos:    4869,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 245,
						EndLine:   245,
						StartPos:  4847,
						EndPos:    4868,
					},
					Condition: &expr.Variable{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4847,
							EndPos:    4848,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4847,
								EndPos:    4848,
							},
							Value: "a",
						},
					},
					IfTrue: &expr.Ternary{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4852,
							EndPos:    4863,
						},
						Condition: &expr.Variable{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4852,
								EndPos:    4853,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4852,
									EndPos:    4853,
								},
								Value: "b",
							},
						},
						IfTrue: &expr.Variable{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4857,
								EndPos:    4858,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4857,
									EndPos:    4858,
								},
								Value: "c",
							},
						},
						IfFalse: &expr.Variable{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4862,
								EndPos:    4863,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4862,
									EndPos:    4863,
								},
								Value: "d",
							},
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4867,
							EndPos:    4868,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4867,
								EndPos:    4868,
							},
							Value: "e",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 246,
					EndLine:   246,
					StartPos:  4873,
					EndPos:    4895,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 246,
						EndLine:   246,
						StartPos:  4873,
						EndPos:    4894,
					},
					Condition: &expr.Ternary{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4873,
							EndPos:    4884,
						},
						Condition: &expr.Variable{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4873,
								EndPos:    4874,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4873,
									EndPos:    4874,
								},
								Value: "a",
							},
						},
						IfTrue: &expr.Variable{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4878,
								EndPos:    4879,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4878,
									EndPos:    4879,
								},
								Value: "b",
							},
						},
						IfFalse: &expr.Variable{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4883,
								EndPos:    4884,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4883,
									EndPos:    4884,
								},
								Value: "c",
							},
						},
					},
					IfTrue: &expr.Variable{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4888,
							EndPos:    4889,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4888,
								EndPos:    4889,
							},
							Value: "d",
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4893,
							EndPos:    4894,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4893,
								EndPos:    4894,
							},
							Value: "e",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 247,
					EndLine:   247,
					StartPos:  4899,
					EndPos:    4902,
				},
				Expr: &expr.UnaryMinus{
					Position: &position.Position{
						StartLine: 247,
						EndLine:   247,
						StartPos:  4899,
						EndPos:    4901,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 247,
							EndLine:   247,
							StartPos:  4900,
							EndPos:    4901,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 247,
								EndLine:   247,
								StartPos:  4900,
								EndPos:    4901,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 248,
					EndLine:   248,
					StartPos:  4906,
					EndPos:    4909,
				},
				Expr: &expr.UnaryPlus{
					Position: &position.Position{
						StartLine: 248,
						EndLine:   248,
						StartPos:  4906,
						EndPos:    4908,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 248,
							EndLine:   248,
							StartPos:  4907,
							EndPos:    4908,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4907,
								EndPos:    4908,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 249,
					EndLine:   249,
					StartPos:  4913,
					EndPos:    4916,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 249,
						EndLine:   249,
						StartPos:  4913,
						EndPos:    4915,
					},
					VarName: &expr.Variable{
						Position: &position.Position{
							StartLine: 249,
							EndLine:   249,
							StartPos:  4914,
							EndPos:    4915,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 249,
								EndLine:   249,
								StartPos:  4914,
								EndPos:    4915,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 250,
					EndLine:   250,
					StartPos:  4920,
					EndPos:    4924,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 250,
						EndLine:   250,
						StartPos:  4920,
						EndPos:    4923,
					},
					VarName: &expr.Variable{
						Position: &position.Position{
							StartLine: 250,
							EndLine:   250,
							StartPos:  4921,
							EndPos:    4923,
						},
						VarName: &expr.Variable{
							Position: &position.Position{
								StartLine: 250,
								EndLine:   250,
								StartPos:  4922,
								EndPos:    4923,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 250,
									EndLine:   250,
									StartPos:  4922,
									EndPos:    4923,
								},
								Value: "a",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 251,
					EndLine:   251,
					StartPos:  4928,
					EndPos:    4933,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 251,
						EndLine:   251,
						StartPos:  4928,
						EndPos:    4932,
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 252,
					EndLine:   252,
					StartPos:  4937,
					EndPos:    4945,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 252,
						EndLine:   252,
						StartPos:  4937,
						EndPos:    4944,
					},
					Value: &expr.Variable{
						Position: &position.Position{
							StartLine: 252,
							EndLine:   252,
							StartPos:  4943,
							EndPos:    4944,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 252,
								EndLine:   252,
								StartPos:  4943,
								EndPos:    4944,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 253,
					EndLine:   253,
					StartPos:  4949,
					EndPos:    4963,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 253,
						EndLine:   253,
						StartPos:  4949,
						EndPos:    4962,
					},
					Key: &expr.Variable{
						Position: &position.Position{
							StartLine: 253,
							EndLine:   253,
							StartPos:  4955,
							EndPos:    4956,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4955,
								EndPos:    4956,
							},
							Value: "a",
						},
					},
					Value: &expr.Variable{
						Position: &position.Position{
							StartLine: 253,
							EndLine:   253,
							StartPos:  4961,
							EndPos:    4962,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4961,
								EndPos:    4962,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 254,
					EndLine:   254,
					StartPos:  4967,
					EndPos:    4983,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 254,
						EndLine:   254,
						StartPos:  4967,
						EndPos:    4982,
					},
					Value: &expr.ClassConstFetch{
						Position: &position.Position{
							StartLine: 254,
							EndLine:   254,
							StartPos:  4973,
							EndPos:    4982,
						},
						Class: &name.Name{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  4973,
								EndPos:    4975,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 254,
										EndLine:   254,
										StartPos:  4973,
										EndPos:    4975,
									},
									Value: "Foo",
								},
							},
						},
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  4978,
								EndPos:    4982,
							},
							Value: "class",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 255,
					EndLine:   255,
					StartPos:  4987,
					EndPos:    5009,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 255,
						EndLine:   255,
						StartPos:  4987,
						EndPos:    5008,
					},
					Key: &expr.Variable{
						Position: &position.Position{
							StartLine: 255,
							EndLine:   255,
							StartPos:  4993,
							EndPos:    4994,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  4993,
								EndPos:    4994,
							},
							Value: "a",
						},
					},
					Value: &expr.ClassConstFetch{
						Position: &position.Position{
							StartLine: 255,
							EndLine:   255,
							StartPos:  4999,
							EndPos:    5008,
						},
						Class: &name.Name{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  4999,
								EndPos:    5001,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 255,
										EndLine:   255,
										StartPos:  4999,
										EndPos:    5001,
									},
									Value: "Foo",
								},
							},
						},
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  5004,
								EndPos:    5008,
							},
							Value: "class",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 257,
					EndLine:   257,
					StartPos:  5016,
					EndPos:    5025,
				},
				Expr: &cast.Array{
					Position: &position.Position{
						StartLine: 257,
						EndLine:   257,
						StartPos:  5016,
						EndPos:    5024,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 257,
							EndLine:   257,
							StartPos:  5023,
							EndPos:    5024,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 257,
								EndLine:   257,
								StartPos:  5023,
								EndPos:    5024,
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
					StartPos:  5029,
					EndPos:    5040,
				},
				Expr: &cast.Bool{
					Position: &position.Position{
						StartLine: 258,
						EndLine:   258,
						StartPos:  5029,
						EndPos:    5039,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 258,
							EndLine:   258,
							StartPos:  5038,
							EndPos:    5039,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 258,
								EndLine:   258,
								StartPos:  5038,
								EndPos:    5039,
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
					StartPos:  5044,
					EndPos:    5052,
				},
				Expr: &cast.Bool{
					Position: &position.Position{
						StartLine: 259,
						EndLine:   259,
						StartPos:  5044,
						EndPos:    5051,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 259,
							EndLine:   259,
							StartPos:  5050,
							EndPos:    5051,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 259,
								EndLine:   259,
								StartPos:  5050,
								EndPos:    5051,
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
					StartPos:  5056,
					EndPos:    5066,
				},
				Expr: &cast.Double{
					Position: &position.Position{
						StartLine: 260,
						EndLine:   260,
						StartPos:  5056,
						EndPos:    5065,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 260,
							EndLine:   260,
							StartPos:  5064,
							EndPos:    5065,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 260,
								EndLine:   260,
								StartPos:  5064,
								EndPos:    5065,
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
					StartPos:  5070,
					EndPos:    5079,
				},
				Expr: &cast.Double{
					Position: &position.Position{
						StartLine: 261,
						EndLine:   261,
						StartPos:  5070,
						EndPos:    5078,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 261,
							EndLine:   261,
							StartPos:  5077,
							EndPos:    5078,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 261,
								EndLine:   261,
								StartPos:  5077,
								EndPos:    5078,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 262,
					EndLine:   262,
					StartPos:  5083,
					EndPos:    5094,
				},
				Expr: &cast.Int{
					Position: &position.Position{
						StartLine: 262,
						EndLine:   262,
						StartPos:  5083,
						EndPos:    5093,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 262,
							EndLine:   262,
							StartPos:  5092,
							EndPos:    5093,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 262,
								EndLine:   262,
								StartPos:  5092,
								EndPos:    5093,
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
					StartPos:  5098,
					EndPos:    5105,
				},
				Expr: &cast.Int{
					Position: &position.Position{
						StartLine: 263,
						EndLine:   263,
						StartPos:  5098,
						EndPos:    5104,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 263,
							EndLine:   263,
							StartPos:  5103,
							EndPos:    5104,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 263,
								EndLine:   263,
								StartPos:  5103,
								EndPos:    5104,
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
					EndPos:    5119,
				},
				Expr: &cast.Object{
					Position: &position.Position{
						StartLine: 264,
						EndLine:   264,
						StartPos:  5109,
						EndPos:    5118,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 264,
							EndLine:   264,
							StartPos:  5117,
							EndPos:    5118,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 264,
								EndLine:   264,
								StartPos:  5117,
								EndPos:    5118,
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
					StartPos:  5123,
					EndPos:    5133,
				},
				Expr: &cast.String{
					Position: &position.Position{
						StartLine: 265,
						EndLine:   265,
						StartPos:  5123,
						EndPos:    5132,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 265,
							EndLine:   265,
							StartPos:  5131,
							EndPos:    5132,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 265,
								EndLine:   265,
								StartPos:  5131,
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
					StartPos:  5137,
					EndPos:    5146,
				},
				Expr: &cast.Unset{
					Position: &position.Position{
						StartLine: 266,
						EndLine:   266,
						StartPos:  5137,
						EndPos:    5145,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 266,
							EndLine:   266,
							StartPos:  5144,
							EndPos:    5145,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 266,
								EndLine:   266,
								StartPos:  5144,
								EndPos:    5145,
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
					StartPos:  5151,
					EndPos:    5158,
				},
				Expr: &binary.BitwiseAnd{
					Position: &position.Position{
						StartLine: 268,
						EndLine:   268,
						StartPos:  5151,
						EndPos:    5157,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 268,
							EndLine:   268,
							StartPos:  5151,
							EndPos:    5152,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 268,
								EndLine:   268,
								StartPos:  5151,
								EndPos:    5152,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 268,
							EndLine:   268,
							StartPos:  5156,
							EndPos:    5157,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 268,
								EndLine:   268,
								StartPos:  5156,
								EndPos:    5157,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 269,
					EndLine:   269,
					StartPos:  5162,
					EndPos:    5169,
				},
				Expr: &binary.BitwiseOr{
					Position: &position.Position{
						StartLine: 269,
						EndLine:   269,
						StartPos:  5162,
						EndPos:    5168,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 269,
							EndLine:   269,
							StartPos:  5162,
							EndPos:    5163,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 269,
								EndLine:   269,
								StartPos:  5162,
								EndPos:    5163,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 269,
							EndLine:   269,
							StartPos:  5167,
							EndPos:    5168,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 269,
								EndLine:   269,
								StartPos:  5167,
								EndPos:    5168,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 270,
					EndLine:   270,
					StartPos:  5173,
					EndPos:    5180,
				},
				Expr: &binary.BitwiseXor{
					Position: &position.Position{
						StartLine: 270,
						EndLine:   270,
						StartPos:  5173,
						EndPos:    5179,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 270,
							EndLine:   270,
							StartPos:  5173,
							EndPos:    5174,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 270,
								EndLine:   270,
								StartPos:  5173,
								EndPos:    5174,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 270,
							EndLine:   270,
							StartPos:  5178,
							EndPos:    5179,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 270,
								EndLine:   270,
								StartPos:  5178,
								EndPos:    5179,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 271,
					EndLine:   271,
					StartPos:  5184,
					EndPos:    5192,
				},
				Expr: &binary.BooleanAnd{
					Position: &position.Position{
						StartLine: 271,
						EndLine:   271,
						StartPos:  5184,
						EndPos:    5191,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 271,
							EndLine:   271,
							StartPos:  5184,
							EndPos:    5185,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 271,
								EndLine:   271,
								StartPos:  5184,
								EndPos:    5185,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 271,
							EndLine:   271,
							StartPos:  5190,
							EndPos:    5191,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 271,
								EndLine:   271,
								StartPos:  5190,
								EndPos:    5191,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 272,
					EndLine:   272,
					StartPos:  5196,
					EndPos:    5204,
				},
				Expr: &binary.BooleanOr{
					Position: &position.Position{
						StartLine: 272,
						EndLine:   272,
						StartPos:  5196,
						EndPos:    5203,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 272,
							EndLine:   272,
							StartPos:  5196,
							EndPos:    5197,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 272,
								EndLine:   272,
								StartPos:  5196,
								EndPos:    5197,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 272,
							EndLine:   272,
							StartPos:  5202,
							EndPos:    5203,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 272,
								EndLine:   272,
								StartPos:  5202,
								EndPos:    5203,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 273,
					EndLine:   273,
					StartPos:  5208,
					EndPos:    5215,
				},
				Expr: &binary.Concat{
					Position: &position.Position{
						StartLine: 273,
						EndLine:   273,
						StartPos:  5208,
						EndPos:    5214,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 273,
							EndLine:   273,
							StartPos:  5208,
							EndPos:    5209,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 273,
								EndLine:   273,
								StartPos:  5208,
								EndPos:    5209,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 273,
							EndLine:   273,
							StartPos:  5213,
							EndPos:    5214,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 273,
								EndLine:   273,
								StartPos:  5213,
								EndPos:    5214,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 274,
					EndLine:   274,
					StartPos:  5219,
					EndPos:    5226,
				},
				Expr: &binary.Div{
					Position: &position.Position{
						StartLine: 274,
						EndLine:   274,
						StartPos:  5219,
						EndPos:    5225,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 274,
							EndLine:   274,
							StartPos:  5219,
							EndPos:    5220,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 274,
								EndLine:   274,
								StartPos:  5219,
								EndPos:    5220,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 274,
							EndLine:   274,
							StartPos:  5224,
							EndPos:    5225,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 274,
								EndLine:   274,
								StartPos:  5224,
								EndPos:    5225,
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
					StartPos:  5230,
					EndPos:    5238,
				},
				Expr: &binary.Equal{
					Position: &position.Position{
						StartLine: 275,
						EndLine:   275,
						StartPos:  5230,
						EndPos:    5237,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 275,
							EndLine:   275,
							StartPos:  5230,
							EndPos:    5231,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 275,
								EndLine:   275,
								StartPos:  5230,
								EndPos:    5231,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 275,
							EndLine:   275,
							StartPos:  5236,
							EndPos:    5237,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 275,
								EndLine:   275,
								StartPos:  5236,
								EndPos:    5237,
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
					StartPos:  5242,
					EndPos:    5250,
				},
				Expr: &binary.GreaterOrEqual{
					Position: &position.Position{
						StartLine: 276,
						EndLine:   276,
						StartPos:  5242,
						EndPos:    5249,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 276,
							EndLine:   276,
							StartPos:  5242,
							EndPos:    5243,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 276,
								EndLine:   276,
								StartPos:  5242,
								EndPos:    5243,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 276,
							EndLine:   276,
							StartPos:  5248,
							EndPos:    5249,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 276,
								EndLine:   276,
								StartPos:  5248,
								EndPos:    5249,
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
					StartPos:  5254,
					EndPos:    5261,
				},
				Expr: &binary.Greater{
					Position: &position.Position{
						StartLine: 277,
						EndLine:   277,
						StartPos:  5254,
						EndPos:    5260,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 277,
							EndLine:   277,
							StartPos:  5254,
							EndPos:    5255,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 277,
								EndLine:   277,
								StartPos:  5254,
								EndPos:    5255,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 277,
							EndLine:   277,
							StartPos:  5259,
							EndPos:    5260,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 277,
								EndLine:   277,
								StartPos:  5259,
								EndPos:    5260,
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
					StartPos:  5265,
					EndPos:    5274,
				},
				Expr: &binary.Identical{
					Position: &position.Position{
						StartLine: 278,
						EndLine:   278,
						StartPos:  5265,
						EndPos:    5273,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 278,
							EndLine:   278,
							StartPos:  5265,
							EndPos:    5266,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 278,
								EndLine:   278,
								StartPos:  5265,
								EndPos:    5266,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 278,
							EndLine:   278,
							StartPos:  5272,
							EndPos:    5273,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 278,
								EndLine:   278,
								StartPos:  5272,
								EndPos:    5273,
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
					StartPos:  5278,
					EndPos:    5287,
				},
				Expr: &binary.LogicalAnd{
					Position: &position.Position{
						StartLine: 279,
						EndLine:   279,
						StartPos:  5278,
						EndPos:    5286,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 279,
							EndLine:   279,
							StartPos:  5278,
							EndPos:    5279,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 279,
								EndLine:   279,
								StartPos:  5278,
								EndPos:    5279,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 279,
							EndLine:   279,
							StartPos:  5285,
							EndPos:    5286,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 279,
								EndLine:   279,
								StartPos:  5285,
								EndPos:    5286,
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
					StartPos:  5291,
					EndPos:    5299,
				},
				Expr: &binary.LogicalOr{
					Position: &position.Position{
						StartLine: 280,
						EndLine:   280,
						StartPos:  5291,
						EndPos:    5298,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 280,
							EndLine:   280,
							StartPos:  5291,
							EndPos:    5292,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 280,
								EndLine:   280,
								StartPos:  5291,
								EndPos:    5292,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 280,
							EndLine:   280,
							StartPos:  5297,
							EndPos:    5298,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 280,
								EndLine:   280,
								StartPos:  5297,
								EndPos:    5298,
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
					StartPos:  5303,
					EndPos:    5312,
				},
				Expr: &binary.LogicalXor{
					Position: &position.Position{
						StartLine: 281,
						EndLine:   281,
						StartPos:  5303,
						EndPos:    5311,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 281,
							EndLine:   281,
							StartPos:  5303,
							EndPos:    5304,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 281,
								EndLine:   281,
								StartPos:  5303,
								EndPos:    5304,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 281,
							EndLine:   281,
							StartPos:  5310,
							EndPos:    5311,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 281,
								EndLine:   281,
								StartPos:  5310,
								EndPos:    5311,
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
					StartPos:  5316,
					EndPos:    5323,
				},
				Expr: &binary.Minus{
					Position: &position.Position{
						StartLine: 282,
						EndLine:   282,
						StartPos:  5316,
						EndPos:    5322,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 282,
							EndLine:   282,
							StartPos:  5316,
							EndPos:    5317,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 282,
								EndLine:   282,
								StartPos:  5316,
								EndPos:    5317,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 282,
							EndLine:   282,
							StartPos:  5321,
							EndPos:    5322,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 282,
								EndLine:   282,
								StartPos:  5321,
								EndPos:    5322,
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
					StartPos:  5327,
					EndPos:    5334,
				},
				Expr: &binary.Mod{
					Position: &position.Position{
						StartLine: 283,
						EndLine:   283,
						StartPos:  5327,
						EndPos:    5333,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 283,
							EndLine:   283,
							StartPos:  5327,
							EndPos:    5328,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 283,
								EndLine:   283,
								StartPos:  5327,
								EndPos:    5328,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 283,
							EndLine:   283,
							StartPos:  5332,
							EndPos:    5333,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 283,
								EndLine:   283,
								StartPos:  5332,
								EndPos:    5333,
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
					StartPos:  5338,
					EndPos:    5345,
				},
				Expr: &binary.Mul{
					Position: &position.Position{
						StartLine: 284,
						EndLine:   284,
						StartPos:  5338,
						EndPos:    5344,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 284,
							EndLine:   284,
							StartPos:  5338,
							EndPos:    5339,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 284,
								EndLine:   284,
								StartPos:  5338,
								EndPos:    5339,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 284,
							EndLine:   284,
							StartPos:  5343,
							EndPos:    5344,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 284,
								EndLine:   284,
								StartPos:  5343,
								EndPos:    5344,
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
					StartPos:  5349,
					EndPos:    5357,
				},
				Expr: &binary.NotEqual{
					Position: &position.Position{
						StartLine: 285,
						EndLine:   285,
						StartPos:  5349,
						EndPos:    5356,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 285,
							EndLine:   285,
							StartPos:  5349,
							EndPos:    5350,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 285,
								EndLine:   285,
								StartPos:  5349,
								EndPos:    5350,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 285,
							EndLine:   285,
							StartPos:  5355,
							EndPos:    5356,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 285,
								EndLine:   285,
								StartPos:  5355,
								EndPos:    5356,
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
					StartPos:  5361,
					EndPos:    5370,
				},
				Expr: &binary.NotIdentical{
					Position: &position.Position{
						StartLine: 286,
						EndLine:   286,
						StartPos:  5361,
						EndPos:    5369,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 286,
							EndLine:   286,
							StartPos:  5361,
							EndPos:    5362,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 286,
								EndLine:   286,
								StartPos:  5361,
								EndPos:    5362,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 286,
							EndLine:   286,
							StartPos:  5368,
							EndPos:    5369,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 286,
								EndLine:   286,
								StartPos:  5368,
								EndPos:    5369,
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
					StartPos:  5374,
					EndPos:    5381,
				},
				Expr: &binary.Plus{
					Position: &position.Position{
						StartLine: 287,
						EndLine:   287,
						StartPos:  5374,
						EndPos:    5380,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 287,
							EndLine:   287,
							StartPos:  5374,
							EndPos:    5375,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 287,
								EndLine:   287,
								StartPos:  5374,
								EndPos:    5375,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 287,
							EndLine:   287,
							StartPos:  5379,
							EndPos:    5380,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 287,
								EndLine:   287,
								StartPos:  5379,
								EndPos:    5380,
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
					StartPos:  5385,
					EndPos:    5393,
				},
				Expr: &binary.Pow{
					Position: &position.Position{
						StartLine: 288,
						EndLine:   288,
						StartPos:  5385,
						EndPos:    5392,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 288,
							EndLine:   288,
							StartPos:  5385,
							EndPos:    5386,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 288,
								EndLine:   288,
								StartPos:  5385,
								EndPos:    5386,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 288,
							EndLine:   288,
							StartPos:  5391,
							EndPos:    5392,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 288,
								EndLine:   288,
								StartPos:  5391,
								EndPos:    5392,
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
					StartPos:  5397,
					EndPos:    5405,
				},
				Expr: &binary.ShiftLeft{
					Position: &position.Position{
						StartLine: 289,
						EndLine:   289,
						StartPos:  5397,
						EndPos:    5404,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 289,
							EndLine:   289,
							StartPos:  5397,
							EndPos:    5398,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 289,
								EndLine:   289,
								StartPos:  5397,
								EndPos:    5398,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 289,
							EndLine:   289,
							StartPos:  5403,
							EndPos:    5404,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 289,
								EndLine:   289,
								StartPos:  5403,
								EndPos:    5404,
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
					StartPos:  5409,
					EndPos:    5417,
				},
				Expr: &binary.ShiftRight{
					Position: &position.Position{
						StartLine: 290,
						EndLine:   290,
						StartPos:  5409,
						EndPos:    5416,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 290,
							EndLine:   290,
							StartPos:  5409,
							EndPos:    5410,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 290,
								EndLine:   290,
								StartPos:  5409,
								EndPos:    5410,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 290,
							EndLine:   290,
							StartPos:  5415,
							EndPos:    5416,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 290,
								EndLine:   290,
								StartPos:  5415,
								EndPos:    5416,
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
					StartPos:  5421,
					EndPos:    5429,
				},
				Expr: &binary.SmallerOrEqual{
					Position: &position.Position{
						StartLine: 291,
						EndLine:   291,
						StartPos:  5421,
						EndPos:    5428,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 291,
							EndLine:   291,
							StartPos:  5421,
							EndPos:    5422,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 291,
								EndLine:   291,
								StartPos:  5421,
								EndPos:    5422,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 291,
							EndLine:   291,
							StartPos:  5427,
							EndPos:    5428,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 291,
								EndLine:   291,
								StartPos:  5427,
								EndPos:    5428,
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
					StartPos:  5433,
					EndPos:    5440,
				},
				Expr: &binary.Smaller{
					Position: &position.Position{
						StartLine: 292,
						EndLine:   292,
						StartPos:  5433,
						EndPos:    5439,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 292,
							EndLine:   292,
							StartPos:  5433,
							EndPos:    5434,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 292,
								EndLine:   292,
								StartPos:  5433,
								EndPos:    5434,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 292,
							EndLine:   292,
							StartPos:  5438,
							EndPos:    5439,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 292,
								EndLine:   292,
								StartPos:  5438,
								EndPos:    5439,
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
					StartPos:  5445,
					EndPos:    5453,
				},
				Expr: &assign.Reference{
					Position: &position.Position{
						StartLine: 294,
						EndLine:   294,
						StartPos:  5445,
						EndPos:    5452,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 294,
							EndLine:   294,
							StartPos:  5445,
							EndPos:    5446,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  5445,
								EndPos:    5446,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 294,
							EndLine:   294,
							StartPos:  5451,
							EndPos:    5452,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  5451,
								EndPos:    5452,
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
					StartPos:  5457,
					EndPos:    5470,
				},
				Expr: &assign.Reference{
					Position: &position.Position{
						StartLine: 295,
						EndLine:   295,
						StartPos:  5457,
						EndPos:    5469,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 295,
							EndLine:   295,
							StartPos:  5457,
							EndPos:    5458,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 295,
								EndLine:   295,
								StartPos:  5457,
								EndPos:    5458,
							},
							Value: "a",
						},
					},
					Expression: &expr.New{
						Position: &position.Position{
							StartLine: 295,
							EndLine:   295,
							StartPos:  5463,
							EndPos:    5469,
						},
						Class: &name.Name{
							Position: &position.Position{
								StartLine: 295,
								EndLine:   295,
								StartPos:  5467,
								EndPos:    5469,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 295,
										EndLine:   295,
										StartPos:  5467,
										EndPos:    5469,
									},
									Value: "Foo",
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 296,
					EndLine:   296,
					StartPos:  5474,
					EndPos:    5491,
				},
				Expr: &assign.Reference{
					Position: &position.Position{
						StartLine: 296,
						EndLine:   296,
						StartPos:  5474,
						EndPos:    5490,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 296,
							EndLine:   296,
							StartPos:  5474,
							EndPos:    5475,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  5474,
								EndPos:    5475,
							},
							Value: "a",
						},
					},
					Expression: &expr.New{
						Position: &position.Position{
							StartLine: 296,
							EndLine:   296,
							StartPos:  5480,
							EndPos:    5490,
						},
						Class: &name.Name{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  5484,
								EndPos:    5486,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 296,
										EndLine:   296,
										StartPos:  5484,
										EndPos:    5486,
									},
									Value: "Foo",
								},
							},
						},
						ArgumentList: &node.ArgumentList{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  5487,
								EndPos:    5490,
							},
							Arguments: []node.Node{
								&node.Argument{
									Position: &position.Position{
										StartLine: 296,
										EndLine:   296,
										StartPos:  5488,
										EndPos:    5489,
									},
									Variadic:    false,
									IsReference: false,
									Expr: &expr.Variable{
										Position: &position.Position{
											StartLine: 296,
											EndLine:   296,
											StartPos:  5488,
											EndPos:    5489,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 296,
												EndLine:   296,
												StartPos:  5488,
												EndPos:    5489,
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
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 297,
					EndLine:   297,
					StartPos:  5495,
					EndPos:    5502,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 297,
						EndLine:   297,
						StartPos:  5495,
						EndPos:    5501,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 297,
							EndLine:   297,
							StartPos:  5495,
							EndPos:    5496,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 297,
								EndLine:   297,
								StartPos:  5495,
								EndPos:    5496,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 297,
							EndLine:   297,
							StartPos:  5500,
							EndPos:    5501,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 297,
								EndLine:   297,
								StartPos:  5500,
								EndPos:    5501,
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
					StartPos:  5506,
					EndPos:    5514,
				},
				Expr: &assign.BitwiseAnd{
					Position: &position.Position{
						StartLine: 298,
						EndLine:   298,
						StartPos:  5506,
						EndPos:    5513,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 298,
							EndLine:   298,
							StartPos:  5506,
							EndPos:    5507,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  5506,
								EndPos:    5507,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 298,
							EndLine:   298,
							StartPos:  5512,
							EndPos:    5513,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  5512,
								EndPos:    5513,
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
					StartPos:  5518,
					EndPos:    5526,
				},
				Expr: &assign.BitwiseOr{
					Position: &position.Position{
						StartLine: 299,
						EndLine:   299,
						StartPos:  5518,
						EndPos:    5525,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 299,
							EndLine:   299,
							StartPos:  5518,
							EndPos:    5519,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5518,
								EndPos:    5519,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 299,
							EndLine:   299,
							StartPos:  5524,
							EndPos:    5525,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5524,
								EndPos:    5525,
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
					StartPos:  5530,
					EndPos:    5538,
				},
				Expr: &assign.BitwiseXor{
					Position: &position.Position{
						StartLine: 300,
						EndLine:   300,
						StartPos:  5530,
						EndPos:    5537,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 300,
							EndLine:   300,
							StartPos:  5530,
							EndPos:    5531,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5530,
								EndPos:    5531,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 300,
							EndLine:   300,
							StartPos:  5536,
							EndPos:    5537,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5536,
								EndPos:    5537,
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
					StartPos:  5542,
					EndPos:    5550,
				},
				Expr: &assign.Concat{
					Position: &position.Position{
						StartLine: 301,
						EndLine:   301,
						StartPos:  5542,
						EndPos:    5549,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 301,
							EndLine:   301,
							StartPos:  5542,
							EndPos:    5543,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 301,
								EndLine:   301,
								StartPos:  5542,
								EndPos:    5543,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 301,
							EndLine:   301,
							StartPos:  5548,
							EndPos:    5549,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 301,
								EndLine:   301,
								StartPos:  5548,
								EndPos:    5549,
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
					StartPos:  5554,
					EndPos:    5562,
				},
				Expr: &assign.Div{
					Position: &position.Position{
						StartLine: 302,
						EndLine:   302,
						StartPos:  5554,
						EndPos:    5561,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 302,
							EndLine:   302,
							StartPos:  5554,
							EndPos:    5555,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 302,
								EndLine:   302,
								StartPos:  5554,
								EndPos:    5555,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 302,
							EndLine:   302,
							StartPos:  5560,
							EndPos:    5561,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 302,
								EndLine:   302,
								StartPos:  5560,
								EndPos:    5561,
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
					StartPos:  5566,
					EndPos:    5574,
				},
				Expr: &assign.Minus{
					Position: &position.Position{
						StartLine: 303,
						EndLine:   303,
						StartPos:  5566,
						EndPos:    5573,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 303,
							EndLine:   303,
							StartPos:  5566,
							EndPos:    5567,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 303,
								EndLine:   303,
								StartPos:  5566,
								EndPos:    5567,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
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
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 304,
					EndLine:   304,
					StartPos:  5578,
					EndPos:    5586,
				},
				Expr: &assign.Mod{
					Position: &position.Position{
						StartLine: 304,
						EndLine:   304,
						StartPos:  5578,
						EndPos:    5585,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 304,
							EndLine:   304,
							StartPos:  5578,
							EndPos:    5579,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 304,
								EndLine:   304,
								StartPos:  5578,
								EndPos:    5579,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
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
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 305,
					EndLine:   305,
					StartPos:  5590,
					EndPos:    5598,
				},
				Expr: &assign.Mul{
					Position: &position.Position{
						StartLine: 305,
						EndLine:   305,
						StartPos:  5590,
						EndPos:    5597,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 305,
							EndLine:   305,
							StartPos:  5590,
							EndPos:    5591,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 305,
								EndLine:   305,
								StartPos:  5590,
								EndPos:    5591,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 305,
							EndLine:   305,
							StartPos:  5596,
							EndPos:    5597,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 305,
								EndLine:   305,
								StartPos:  5596,
								EndPos:    5597,
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
					StartPos:  5602,
					EndPos:    5610,
				},
				Expr: &assign.Plus{
					Position: &position.Position{
						StartLine: 306,
						EndLine:   306,
						StartPos:  5602,
						EndPos:    5609,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 306,
							EndLine:   306,
							StartPos:  5602,
							EndPos:    5603,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 306,
								EndLine:   306,
								StartPos:  5602,
								EndPos:    5603,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 306,
							EndLine:   306,
							StartPos:  5608,
							EndPos:    5609,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 306,
								EndLine:   306,
								StartPos:  5608,
								EndPos:    5609,
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
					StartPos:  5614,
					EndPos:    5623,
				},
				Expr: &assign.Pow{
					Position: &position.Position{
						StartLine: 307,
						EndLine:   307,
						StartPos:  5614,
						EndPos:    5622,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 307,
							EndLine:   307,
							StartPos:  5614,
							EndPos:    5615,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 307,
								EndLine:   307,
								StartPos:  5614,
								EndPos:    5615,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 307,
							EndLine:   307,
							StartPos:  5621,
							EndPos:    5622,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 307,
								EndLine:   307,
								StartPos:  5621,
								EndPos:    5622,
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
					StartPos:  5627,
					EndPos:    5636,
				},
				Expr: &assign.ShiftLeft{
					Position: &position.Position{
						StartLine: 308,
						EndLine:   308,
						StartPos:  5627,
						EndPos:    5635,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 308,
							EndLine:   308,
							StartPos:  5627,
							EndPos:    5628,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 308,
								EndLine:   308,
								StartPos:  5627,
								EndPos:    5628,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 308,
							EndLine:   308,
							StartPos:  5634,
							EndPos:    5635,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 308,
								EndLine:   308,
								StartPos:  5634,
								EndPos:    5635,
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
					StartPos:  5640,
					EndPos:    5649,
				},
				Expr: &assign.ShiftRight{
					Position: &position.Position{
						StartLine: 309,
						EndLine:   309,
						StartPos:  5640,
						EndPos:    5648,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 309,
							EndLine:   309,
							StartPos:  5640,
							EndPos:    5641,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 309,
								EndLine:   309,
								StartPos:  5640,
								EndPos:    5641,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 309,
							EndLine:   309,
							StartPos:  5647,
							EndPos:    5648,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 309,
								EndLine:   309,
								StartPos:  5647,
								EndPos:    5648,
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
					StartPos:  5656,
					EndPos:    5667,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 312,
						EndLine:   312,
						StartPos:  5656,
						EndPos:    5665,
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 312,
							EndLine:   312,
							StartPos:  5660,
							EndPos:    5663,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 312,
									EndLine:   312,
									StartPos:  5661,
									EndPos:    5663,
								},
								Value: "Foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 312,
							EndLine:   312,
							StartPos:  5664,
							EndPos:    5665,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 313,
					EndLine:   313,
					StartPos:  5692,
					EndPos:    5695,
				},
				Expr: &expr.PropertyFetch{
					Position: &position.Position{
						StartLine: 313,
						EndLine:   313,
						StartPos:  5692,
						EndPos:    5694,
					},
					Variable: &expr.MethodCall{
						Position: &position.Position{
							StartLine: 313,
							EndLine:   313,
							StartPos:  5688,
							EndPos:    5689,
						},
						Variable: &expr.New{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5672,
								EndPos:    5681,
							},
							Class: &name.FullyQualified{
								Position: &position.Position{
									StartLine: 313,
									EndLine:   313,
									StartPos:  5676,
									EndPos:    5679,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 313,
											EndLine:   313,
											StartPos:  5677,
											EndPos:    5679,
										},
										Value: "Foo",
									},
								},
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 313,
									EndLine:   313,
									StartPos:  5680,
									EndPos:    5681,
								},
							},
						},
						Method: &node.Identifier{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5685,
								EndPos:    5687,
							},
							Value: "bar",
						},
						ArgumentList: &node.ArgumentList{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5688,
								EndPos:    5689,
							},
						},
					},
					Property: &node.Identifier{
						Position: &position.Position{
							StartLine: 313,
							EndLine:   313,
							StartPos:  5692,
							EndPos:    5694,
						},
						Value: "baz",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 314,
					EndLine:   314,
					StartPos:  5715,
					EndPos:    5717,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 314,
						EndLine:   314,
						StartPos:  5715,
						EndPos:    5715,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 314,
							EndLine:   314,
							StartPos:  5712,
							EndPos:    5712,
						},
						Variable: &expr.New{
							Position: &position.Position{
								StartLine: 314,
								EndLine:   314,
								StartPos:  5700,
								EndPos:    5709,
							},
							Class: &name.FullyQualified{
								Position: &position.Position{
									StartLine: 314,
									EndLine:   314,
									StartPos:  5704,
									EndPos:    5707,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 314,
											EndLine:   314,
											StartPos:  5705,
											EndPos:    5707,
										},
										Value: "Foo",
									},
								},
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 314,
									EndLine:   314,
									StartPos:  5708,
									EndPos:    5709,
								},
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 314,
								EndLine:   314,
								StartPos:  5712,
								EndPos:    5712,
							},
							Value: "0",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 314,
							EndLine:   314,
							StartPos:  5715,
							EndPos:    5715,
						},
						Value: "0",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 315,
					EndLine:   315,
					StartPos:  5741,
					EndPos:    5743,
				},
				Expr: &expr.MethodCall{
					Position: &position.Position{
						StartLine: 315,
						EndLine:   315,
						StartPos:  5741,
						EndPos:    5742,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 315,
							EndLine:   315,
							StartPos:  5734,
							EndPos:    5734,
						},
						Variable: &expr.New{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5722,
								EndPos:    5731,
							},
							Class: &name.FullyQualified{
								Position: &position.Position{
									StartLine: 315,
									EndLine:   315,
									StartPos:  5726,
									EndPos:    5729,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 315,
											EndLine:   315,
											StartPos:  5727,
											EndPos:    5729,
										},
										Value: "Foo",
									},
								},
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 315,
									EndLine:   315,
									StartPos:  5730,
									EndPos:    5731,
								},
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5734,
								EndPos:    5734,
							},
							Value: "0",
						},
					},
					Method: &node.Identifier{
						Position: &position.Position{
							StartLine: 315,
							EndLine:   315,
							StartPos:  5738,
							EndPos:    5740,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 315,
							EndLine:   315,
							StartPos:  5741,
							EndPos:    5742,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 317,
					EndLine:   317,
					StartPos:  5748,
					EndPos:    5764,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 317,
						EndLine:   317,
						StartPos:  5748,
						EndPos:    5763,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 317,
							EndLine:   317,
							StartPos:  5748,
							EndPos:    5760,
						},
						Variable: &expr.Array{
							Position: &position.Position{
								StartLine: 317,
								EndLine:   317,
								StartPos:  5748,
								EndPos:    5757,
							},
							Items: []node.Node{
								&expr.ArrayItem{
									Position: &position.Position{
										StartLine: 317,
										EndLine:   317,
										StartPos:  5754,
										EndPos:    5756,
									},
									Val: &expr.ShortArray{
										Position: &position.Position{
											StartLine: 317,
											EndLine:   317,
											StartPos:  5754,
											EndPos:    5756,
										},
										Items: []node.Node{
											&expr.ArrayItem{
												Position: &position.Position{
													StartLine: 317,
													EndLine:   317,
													StartPos:  5755,
													EndPos:    5755,
												},
												Val: &scalar.Lnumber{
													Position: &position.Position{
														StartLine: 317,
														EndLine:   317,
														StartPos:  5755,
														EndPos:    5755,
													},
													Value: "0",
												},
											},
										},
									},
								},
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 317,
								EndLine:   317,
								StartPos:  5759,
								EndPos:    5759,
							},
							Value: "0",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 317,
							EndLine:   317,
							StartPos:  5762,
							EndPos:    5762,
						},
						Value: "0",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 318,
					EndLine:   318,
					StartPos:  5768,
					EndPos:    5776,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 318,
						EndLine:   318,
						StartPos:  5768,
						EndPos:    5775,
					},
					Variable: &scalar.String{
						Position: &position.Position{
							StartLine: 318,
							EndLine:   318,
							StartPos:  5768,
							EndPos:    5772,
						},
						Value: "\"foo\"",
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 318,
							EndLine:   318,
							StartPos:  5774,
							EndPos:    5774,
						},
						Value: "0",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 319,
					EndLine:   319,
					StartPos:  5780,
					EndPos:    5786,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 319,
						EndLine:   319,
						StartPos:  5780,
						EndPos:    5785,
					},
					Variable: &expr.ConstFetch{
						Position: &position.Position{
							StartLine: 319,
							EndLine:   319,
							StartPos:  5780,
							EndPos:    5782,
						},
						Constant: &name.Name{
							Position: &position.Position{
								StartLine: 319,
								EndLine:   319,
								StartPos:  5780,
								EndPos:    5782,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 319,
										EndLine:   319,
										StartPos:  5780,
										EndPos:    5782,
									},
									Value: "foo",
								},
							},
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 319,
							EndLine:   319,
							StartPos:  5784,
							EndPos:    5784,
						},
						Value: "0",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 320,
					EndLine:   320,
					StartPos:  5790,
					EndPos:    5801,
				},
				Expr: &expr.ClassConstFetch{
					Position: &position.Position{
						StartLine: 320,
						EndLine:   320,
						StartPos:  5790,
						EndPos:    5800,
					},
					Class: &node.Identifier{
						Position: &position.Position{
							StartLine: 320,
							EndLine:   320,
							StartPos:  5790,
							EndPos:    5795,
						},
						Value: "static",
					},
					ConstantName: &node.Identifier{
						Position: &position.Position{
							StartLine: 320,
							EndLine:   320,
							StartPos:  5798,
							EndPos:    5800,
						},
						Value: "foo",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 322,
					EndLine:   322,
					StartPos:  5806,
					EndPos:    5814,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 322,
						EndLine:   322,
						StartPos:  5806,
						EndPos:    5813,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 322,
							EndLine:   322,
							StartPos:  5810,
							EndPos:    5813,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 322,
								EndLine:   322,
								StartPos:  5810,
								EndPos:    5813,
							},
							Value: "foo",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 323,
					EndLine:   323,
					StartPos:  5818,
					EndPos:    5832,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 323,
						EndLine:   323,
						StartPos:  5818,
						EndPos:    5831,
					},
					Class: &expr.StaticPropertyFetch{
						Position: &position.Position{
							StartLine: 323,
							EndLine:   323,
							StartPos:  5822,
							EndPos:    5831,
						},
						Class: &expr.Variable{
							Position: &position.Position{
								StartLine: 323,
								EndLine:   323,
								StartPos:  5822,
								EndPos:    5825,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 323,
									EndLine:   323,
									StartPos:  5822,
									EndPos:    5825,
								},
								Value: "foo",
							},
						},
						Property: &expr.Variable{
							Position: &position.Position{
								StartLine: 323,
								EndLine:   323,
								StartPos:  5828,
								EndPos:    5831,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 323,
									EndLine:   323,
									StartPos:  5828,
									EndPos:    5831,
								},
								Value: "bar",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 324,
					EndLine:   324,
					StartPos:  5836,
					EndPos:    5848,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 324,
						EndLine:   324,
						StartPos:  5836,
						EndPos:    5846,
					},
					Class: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 324,
							EndLine:   324,
							StartPos:  5846,
							EndPos:    5846,
						},
						Variable: &expr.PropertyFetch{
							Position: &position.Position{
								StartLine: 324,
								EndLine:   324,
								StartPos:  5844,
								EndPos:    5846,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 324,
									EndLine:   324,
									StartPos:  5840,
									EndPos:    5844,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 324,
										EndLine:   324,
										StartPos:  5840,
										EndPos:    5841,
									},
									Value: "a",
								},
							},
							Property: &node.Identifier{
								Position: &position.Position{
									StartLine: 324,
									EndLine:   324,
									StartPos:  5844,
									EndPos:    5844,
								},
								Value: "b",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 324,
								EndLine:   324,
								StartPos:  5846,
								EndPos:    5846,
							},
							Value: "0",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 325,
					EndLine:   325,
					StartPos:  5852,
					EndPos:    5883,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 325,
						EndLine:   325,
						StartPos:  5852,
						EndPos:    5881,
					},
					Class: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 325,
							EndLine:   325,
							StartPos:  5881,
							EndPos:    5881,
						},
						Variable: &expr.PropertyFetch{
							Position: &position.Position{
								StartLine: 325,
								EndLine:   325,
								StartPos:  5879,
								EndPos:    5881,
							},
							Variable: &expr.PropertyFetch{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5875,
									EndPos:    5879,
								},
								Variable: &expr.ArrayDimFetch{
									Position: &position.Position{
										StartLine: 325,
										EndLine:   325,
										StartPos:  5862,
										EndPos:    5876,
									},
									Variable: &expr.PropertyFetch{
										Position: &position.Position{
											StartLine: 325,
											EndLine:   325,
											StartPos:  5860,
											EndPos:    5871,
										},
										Variable: &expr.Variable{
											Position: &position.Position{
												StartLine: 325,
												EndLine:   325,
												StartPos:  5856,
												EndPos:    5860,
											},
											VarName: &node.Identifier{
												Position: &position.Position{
													StartLine: 325,
													EndLine:   325,
													StartPos:  5856,
													EndPos:    5857,
												},
												Value: "a",
											},
										},
										Property: &node.Identifier{
											Position: &position.Position{
												StartLine: 325,
												EndLine:   325,
												StartPos:  5860,
												EndPos:    5860,
											},
											Value: "b",
										},
									},
									Dim: &expr.Ternary{
										Position: &position.Position{
											StartLine: 325,
											EndLine:   325,
											StartPos:  5862,
											EndPos:    5871,
										},
										Condition: &expr.Variable{
											Position: &position.Position{
												StartLine: 325,
												EndLine:   325,
												StartPos:  5862,
												EndPos:    5863,
											},
											VarName: &node.Identifier{
												Position: &position.Position{
													StartLine: 325,
													EndLine:   325,
													StartPos:  5862,
													EndPos:    5863,
												},
												Value: "b",
											},
										},
										IfFalse: &expr.ConstFetch{
											Position: &position.Position{
												StartLine: 325,
												EndLine:   325,
												StartPos:  5868,
												EndPos:    5871,
											},
											Constant: &name.Name{
												Position: &position.Position{
													StartLine: 325,
													EndLine:   325,
													StartPos:  5868,
													EndPos:    5871,
												},
												Parts: []node.Node{
													&name.NamePart{
														Position: &position.Position{
															StartLine: 325,
															EndLine:   325,
															StartPos:  5868,
															EndPos:    5871,
														},
														Value: "null",
													},
												},
											},
										},
									},
								},
								Property: &expr.Variable{
									Position: &position.Position{
										StartLine: 325,
										EndLine:   325,
										StartPos:  5875,
										EndPos:    5876,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 325,
											EndLine:   325,
											StartPos:  5875,
											EndPos:    5876,
										},
										Value: "c",
									},
								},
							},
							Property: &node.Identifier{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5879,
									EndPos:    5879,
								},
								Value: "d",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 325,
								EndLine:   325,
								StartPos:  5881,
								EndPos:    5881,
							},
							Value: "0",
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 325,
					EndLine:   325,
					StartPos:  5884,
					EndPos:    5902,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 325,
							EndLine:   325,
							StartPos:  5891,
							EndPos:    5901,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 325,
								EndLine:   325,
								StartPos:  5891,
								EndPos:    5892,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5891,
									EndPos:    5892,
								},
								Value: "a",
							},
						},
						Expr: &expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 325,
								EndLine:   325,
								StartPos:  5896,
								EndPos:    5901,
							},
							Variable: &expr.ShortArray{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5896,
									EndPos:    5898,
								},
								Items: []node.Node{
									&expr.ArrayItem{
										Position: &position.Position{
											StartLine: 325,
											EndLine:   325,
											StartPos:  5897,
											EndPos:    5897,
										},
										Val: &scalar.Lnumber{
											Position: &position.Position{
												StartLine: 325,
												EndLine:   325,
												StartPos:  5897,
												EndPos:    5897,
											},
											Value: "1",
										},
									},
								},
							},
							Dim: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5900,
									EndPos:    5900,
								},
								Value: "0",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 327,
					EndLine:   327,
					StartPos:  5907,
					EndPos:    5921,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 327,
							EndLine:   327,
							StartPos:  5914,
							EndPos:    5920,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 327,
								EndLine:   327,
								StartPos:  5914,
								EndPos:    5915,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 327,
									EndLine:   327,
									StartPos:  5914,
									EndPos:    5915,
								},
								Value: "a",
							},
						},
						Expr: &expr.BooleanNot{
							Position: &position.Position{
								StartLine: 327,
								EndLine:   327,
								StartPos:  5919,
								EndPos:    5920,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 327,
									EndLine:   327,
									StartPos:  5920,
									EndPos:    5920,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 328,
					EndLine:   328,
					StartPos:  5925,
					EndPos:    5939,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 328,
							EndLine:   328,
							StartPos:  5932,
							EndPos:    5938,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5932,
								EndPos:    5933,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5932,
									EndPos:    5933,
								},
								Value: "a",
							},
						},
						Expr: &expr.BitwiseNot{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5937,
								EndPos:    5938,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5938,
									EndPos:    5938,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 329,
					EndLine:   329,
					StartPos:  5943,
					EndPos:    5957,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 329,
							EndLine:   329,
							StartPos:  5950,
							EndPos:    5956,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5950,
								EndPos:    5951,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5950,
									EndPos:    5951,
								},
								Value: "a",
							},
						},
						Expr: &expr.UnaryPlus{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5955,
								EndPos:    5956,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5956,
									EndPos:    5956,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 330,
					EndLine:   330,
					StartPos:  5961,
					EndPos:    5975,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 330,
							EndLine:   330,
							StartPos:  5968,
							EndPos:    5974,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  5968,
								EndPos:    5969,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  5968,
									EndPos:    5969,
								},
								Value: "a",
							},
						},
						Expr: &expr.UnaryMinus{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  5973,
								EndPos:    5974,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  5974,
									EndPos:    5974,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 331,
					EndLine:   331,
					StartPos:  5979,
					EndPos:    5994,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 331,
							EndLine:   331,
							StartPos:  5986,
							EndPos:    5992,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 331,
								EndLine:   331,
								StartPos:  5986,
								EndPos:    5987,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 331,
									EndLine:   331,
									StartPos:  5986,
									EndPos:    5987,
								},
								Value: "a",
							},
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 331,
								EndLine:   331,
								StartPos:  5992,
								EndPos:    5992,
							},
							Value: "1",
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 332,
					EndLine:   332,
					StartPos:  5998,
					EndPos:    6016,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 332,
							EndLine:   332,
							StartPos:  6005,
							EndPos:    6015,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 332,
								EndLine:   332,
								StartPos:  6005,
								EndPos:    6006,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 332,
									EndLine:   332,
									StartPos:  6005,
									EndPos:    6006,
								},
								Value: "a",
							},
						},
						Expr: &expr.Ternary{
							Position: &position.Position{
								StartLine: 332,
								EndLine:   332,
								StartPos:  6010,
								EndPos:    6015,
							},
							Condition: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 332,
									EndLine:   332,
									StartPos:  6010,
									EndPos:    6010,
								},
								Value: "1",
							},
							IfFalse: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 332,
									EndLine:   332,
									StartPos:  6015,
									EndPos:    6015,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 333,
					EndLine:   333,
					StartPos:  6020,
					EndPos:    6041,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 333,
							EndLine:   333,
							StartPos:  6027,
							EndPos:    6040,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 333,
								EndLine:   333,
								StartPos:  6027,
								EndPos:    6028,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 333,
									EndLine:   333,
									StartPos:  6027,
									EndPos:    6028,
								},
								Value: "a",
							},
						},
						Expr: &expr.Ternary{
							Position: &position.Position{
								StartLine: 333,
								EndLine:   333,
								StartPos:  6032,
								EndPos:    6040,
							},
							Condition: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 333,
									EndLine:   333,
									StartPos:  6032,
									EndPos:    6032,
								},
								Value: "1",
							},
							IfTrue: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 333,
									EndLine:   333,
									StartPos:  6036,
									EndPos:    6036,
								},
								Value: "2",
							},
							IfFalse: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 333,
									EndLine:   333,
									StartPos:  6040,
									EndPos:    6040,
								},
								Value: "3",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 334,
					EndLine:   334,
					StartPos:  6045,
					EndPos:    6062,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 334,
							EndLine:   334,
							StartPos:  6052,
							EndPos:    6061,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 334,
								EndLine:   334,
								StartPos:  6052,
								EndPos:    6053,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 334,
									EndLine:   334,
									StartPos:  6052,
									EndPos:    6053,
								},
								Value: "a",
							},
						},
						Expr: &binary.BitwiseAnd{
							Position: &position.Position{
								StartLine: 334,
								EndLine:   334,
								StartPos:  6057,
								EndPos:    6061,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 334,
									EndLine:   334,
									StartPos:  6057,
									EndPos:    6057,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 334,
									EndLine:   334,
									StartPos:  6061,
									EndPos:    6061,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 335,
					EndLine:   335,
					StartPos:  6066,
					EndPos:    6083,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 335,
							EndLine:   335,
							StartPos:  6073,
							EndPos:    6082,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 335,
								EndLine:   335,
								StartPos:  6073,
								EndPos:    6074,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  6073,
									EndPos:    6074,
								},
								Value: "a",
							},
						},
						Expr: &binary.BitwiseOr{
							Position: &position.Position{
								StartLine: 335,
								EndLine:   335,
								StartPos:  6078,
								EndPos:    6082,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  6078,
									EndPos:    6078,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  6082,
									EndPos:    6082,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 336,
					EndLine:   336,
					StartPos:  6087,
					EndPos:    6104,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 336,
							EndLine:   336,
							StartPos:  6094,
							EndPos:    6103,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 336,
								EndLine:   336,
								StartPos:  6094,
								EndPos:    6095,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 336,
									EndLine:   336,
									StartPos:  6094,
									EndPos:    6095,
								},
								Value: "a",
							},
						},
						Expr: &binary.BitwiseXor{
							Position: &position.Position{
								StartLine: 336,
								EndLine:   336,
								StartPos:  6099,
								EndPos:    6103,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 336,
									EndLine:   336,
									StartPos:  6099,
									EndPos:    6099,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 336,
									EndLine:   336,
									StartPos:  6103,
									EndPos:    6103,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 337,
					EndLine:   337,
					StartPos:  6108,
					EndPos:    6126,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 337,
							EndLine:   337,
							StartPos:  6115,
							EndPos:    6125,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 337,
								EndLine:   337,
								StartPos:  6115,
								EndPos:    6116,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 337,
									EndLine:   337,
									StartPos:  6115,
									EndPos:    6116,
								},
								Value: "a",
							},
						},
						Expr: &binary.BooleanAnd{
							Position: &position.Position{
								StartLine: 337,
								EndLine:   337,
								StartPos:  6120,
								EndPos:    6125,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 337,
									EndLine:   337,
									StartPos:  6120,
									EndPos:    6120,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 337,
									EndLine:   337,
									StartPos:  6125,
									EndPos:    6125,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 338,
					EndLine:   338,
					StartPos:  6130,
					EndPos:    6148,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 338,
							EndLine:   338,
							StartPos:  6137,
							EndPos:    6147,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 338,
								EndLine:   338,
								StartPos:  6137,
								EndPos:    6138,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 338,
									EndLine:   338,
									StartPos:  6137,
									EndPos:    6138,
								},
								Value: "a",
							},
						},
						Expr: &binary.BooleanOr{
							Position: &position.Position{
								StartLine: 338,
								EndLine:   338,
								StartPos:  6142,
								EndPos:    6147,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 338,
									EndLine:   338,
									StartPos:  6142,
									EndPos:    6142,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 338,
									EndLine:   338,
									StartPos:  6147,
									EndPos:    6147,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 339,
					EndLine:   339,
					StartPos:  6152,
					EndPos:    6169,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 339,
							EndLine:   339,
							StartPos:  6159,
							EndPos:    6168,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 339,
								EndLine:   339,
								StartPos:  6159,
								EndPos:    6160,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 339,
									EndLine:   339,
									StartPos:  6159,
									EndPos:    6160,
								},
								Value: "a",
							},
						},
						Expr: &binary.Concat{
							Position: &position.Position{
								StartLine: 339,
								EndLine:   339,
								StartPos:  6164,
								EndPos:    6168,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 339,
									EndLine:   339,
									StartPos:  6164,
									EndPos:    6164,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 339,
									EndLine:   339,
									StartPos:  6168,
									EndPos:    6168,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 340,
					EndLine:   340,
					StartPos:  6173,
					EndPos:    6190,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 340,
							EndLine:   340,
							StartPos:  6180,
							EndPos:    6189,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 340,
								EndLine:   340,
								StartPos:  6180,
								EndPos:    6181,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 340,
									EndLine:   340,
									StartPos:  6180,
									EndPos:    6181,
								},
								Value: "a",
							},
						},
						Expr: &binary.Div{
							Position: &position.Position{
								StartLine: 340,
								EndLine:   340,
								StartPos:  6185,
								EndPos:    6189,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 340,
									EndLine:   340,
									StartPos:  6185,
									EndPos:    6185,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 340,
									EndLine:   340,
									StartPos:  6189,
									EndPos:    6189,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 341,
					EndLine:   341,
					StartPos:  6194,
					EndPos:    6212,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 341,
							EndLine:   341,
							StartPos:  6201,
							EndPos:    6211,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  6201,
								EndPos:    6202,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  6201,
									EndPos:    6202,
								},
								Value: "a",
							},
						},
						Expr: &binary.Equal{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  6206,
								EndPos:    6211,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  6206,
									EndPos:    6206,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  6211,
									EndPos:    6211,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 342,
					EndLine:   342,
					StartPos:  6216,
					EndPos:    6234,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 342,
							EndLine:   342,
							StartPos:  6223,
							EndPos:    6233,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 342,
								EndLine:   342,
								StartPos:  6223,
								EndPos:    6224,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 342,
									EndLine:   342,
									StartPos:  6223,
									EndPos:    6224,
								},
								Value: "a",
							},
						},
						Expr: &binary.GreaterOrEqual{
							Position: &position.Position{
								StartLine: 342,
								EndLine:   342,
								StartPos:  6228,
								EndPos:    6233,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 342,
									EndLine:   342,
									StartPos:  6228,
									EndPos:    6228,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 342,
									EndLine:   342,
									StartPos:  6233,
									EndPos:    6233,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 343,
					EndLine:   343,
					StartPos:  6238,
					EndPos:    6255,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 343,
							EndLine:   343,
							StartPos:  6245,
							EndPos:    6254,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 343,
								EndLine:   343,
								StartPos:  6245,
								EndPos:    6246,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 343,
									EndLine:   343,
									StartPos:  6245,
									EndPos:    6246,
								},
								Value: "a",
							},
						},
						Expr: &binary.Greater{
							Position: &position.Position{
								StartLine: 343,
								EndLine:   343,
								StartPos:  6250,
								EndPos:    6254,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 343,
									EndLine:   343,
									StartPos:  6250,
									EndPos:    6250,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 343,
									EndLine:   343,
									StartPos:  6254,
									EndPos:    6254,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 344,
					EndLine:   344,
					StartPos:  6259,
					EndPos:    6278,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 344,
							EndLine:   344,
							StartPos:  6266,
							EndPos:    6277,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 344,
								EndLine:   344,
								StartPos:  6266,
								EndPos:    6267,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 344,
									EndLine:   344,
									StartPos:  6266,
									EndPos:    6267,
								},
								Value: "a",
							},
						},
						Expr: &binary.Identical{
							Position: &position.Position{
								StartLine: 344,
								EndLine:   344,
								StartPos:  6271,
								EndPos:    6277,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 344,
									EndLine:   344,
									StartPos:  6271,
									EndPos:    6271,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 344,
									EndLine:   344,
									StartPos:  6277,
									EndPos:    6277,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 345,
					EndLine:   345,
					StartPos:  6282,
					EndPos:    6301,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 345,
							EndLine:   345,
							StartPos:  6289,
							EndPos:    6300,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 345,
								EndLine:   345,
								StartPos:  6289,
								EndPos:    6290,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 345,
									EndLine:   345,
									StartPos:  6289,
									EndPos:    6290,
								},
								Value: "a",
							},
						},
						Expr: &binary.LogicalAnd{
							Position: &position.Position{
								StartLine: 345,
								EndLine:   345,
								StartPos:  6294,
								EndPos:    6300,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 345,
									EndLine:   345,
									StartPos:  6294,
									EndPos:    6294,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 345,
									EndLine:   345,
									StartPos:  6300,
									EndPos:    6300,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 346,
					EndLine:   346,
					StartPos:  6305,
					EndPos:    6323,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 346,
							EndLine:   346,
							StartPos:  6312,
							EndPos:    6322,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 346,
								EndLine:   346,
								StartPos:  6312,
								EndPos:    6313,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  6312,
									EndPos:    6313,
								},
								Value: "a",
							},
						},
						Expr: &binary.LogicalOr{
							Position: &position.Position{
								StartLine: 346,
								EndLine:   346,
								StartPos:  6317,
								EndPos:    6322,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  6317,
									EndPos:    6317,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  6322,
									EndPos:    6322,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 347,
					EndLine:   347,
					StartPos:  6327,
					EndPos:    6346,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 347,
							EndLine:   347,
							StartPos:  6334,
							EndPos:    6345,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 347,
								EndLine:   347,
								StartPos:  6334,
								EndPos:    6335,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 347,
									EndLine:   347,
									StartPos:  6334,
									EndPos:    6335,
								},
								Value: "a",
							},
						},
						Expr: &binary.LogicalXor{
							Position: &position.Position{
								StartLine: 347,
								EndLine:   347,
								StartPos:  6339,
								EndPos:    6345,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 347,
									EndLine:   347,
									StartPos:  6339,
									EndPos:    6339,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 347,
									EndLine:   347,
									StartPos:  6345,
									EndPos:    6345,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 348,
					EndLine:   348,
					StartPos:  6350,
					EndPos:    6367,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 348,
							EndLine:   348,
							StartPos:  6357,
							EndPos:    6366,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 348,
								EndLine:   348,
								StartPos:  6357,
								EndPos:    6358,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 348,
									EndLine:   348,
									StartPos:  6357,
									EndPos:    6358,
								},
								Value: "a",
							},
						},
						Expr: &binary.Minus{
							Position: &position.Position{
								StartLine: 348,
								EndLine:   348,
								StartPos:  6362,
								EndPos:    6366,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 348,
									EndLine:   348,
									StartPos:  6362,
									EndPos:    6362,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 348,
									EndLine:   348,
									StartPos:  6366,
									EndPos:    6366,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 349,
					EndLine:   349,
					StartPos:  6371,
					EndPos:    6388,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 349,
							EndLine:   349,
							StartPos:  6378,
							EndPos:    6387,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 349,
								EndLine:   349,
								StartPos:  6378,
								EndPos:    6379,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 349,
									EndLine:   349,
									StartPos:  6378,
									EndPos:    6379,
								},
								Value: "a",
							},
						},
						Expr: &binary.Mod{
							Position: &position.Position{
								StartLine: 349,
								EndLine:   349,
								StartPos:  6383,
								EndPos:    6387,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 349,
									EndLine:   349,
									StartPos:  6383,
									EndPos:    6383,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 349,
									EndLine:   349,
									StartPos:  6387,
									EndPos:    6387,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 350,
					EndLine:   350,
					StartPos:  6392,
					EndPos:    6409,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 350,
							EndLine:   350,
							StartPos:  6399,
							EndPos:    6408,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 350,
								EndLine:   350,
								StartPos:  6399,
								EndPos:    6400,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 350,
									EndLine:   350,
									StartPos:  6399,
									EndPos:    6400,
								},
								Value: "a",
							},
						},
						Expr: &binary.Mul{
							Position: &position.Position{
								StartLine: 350,
								EndLine:   350,
								StartPos:  6404,
								EndPos:    6408,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 350,
									EndLine:   350,
									StartPos:  6404,
									EndPos:    6404,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 350,
									EndLine:   350,
									StartPos:  6408,
									EndPos:    6408,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 351,
					EndLine:   351,
					StartPos:  6413,
					EndPos:    6431,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 351,
							EndLine:   351,
							StartPos:  6420,
							EndPos:    6430,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 351,
								EndLine:   351,
								StartPos:  6420,
								EndPos:    6421,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 351,
									EndLine:   351,
									StartPos:  6420,
									EndPos:    6421,
								},
								Value: "a",
							},
						},
						Expr: &binary.NotEqual{
							Position: &position.Position{
								StartLine: 351,
								EndLine:   351,
								StartPos:  6425,
								EndPos:    6430,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 351,
									EndLine:   351,
									StartPos:  6425,
									EndPos:    6425,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 351,
									EndLine:   351,
									StartPos:  6430,
									EndPos:    6430,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 352,
					EndLine:   352,
					StartPos:  6435,
					EndPos:    6454,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 352,
							EndLine:   352,
							StartPos:  6442,
							EndPos:    6453,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 352,
								EndLine:   352,
								StartPos:  6442,
								EndPos:    6443,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 352,
									EndLine:   352,
									StartPos:  6442,
									EndPos:    6443,
								},
								Value: "a",
							},
						},
						Expr: &binary.NotIdentical{
							Position: &position.Position{
								StartLine: 352,
								EndLine:   352,
								StartPos:  6447,
								EndPos:    6453,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 352,
									EndLine:   352,
									StartPos:  6447,
									EndPos:    6447,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 352,
									EndLine:   352,
									StartPos:  6453,
									EndPos:    6453,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 353,
					EndLine:   353,
					StartPos:  6458,
					EndPos:    6475,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 353,
							EndLine:   353,
							StartPos:  6465,
							EndPos:    6474,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 353,
								EndLine:   353,
								StartPos:  6465,
								EndPos:    6466,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 353,
									EndLine:   353,
									StartPos:  6465,
									EndPos:    6466,
								},
								Value: "a",
							},
						},
						Expr: &binary.Plus{
							Position: &position.Position{
								StartLine: 353,
								EndLine:   353,
								StartPos:  6470,
								EndPos:    6474,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 353,
									EndLine:   353,
									StartPos:  6470,
									EndPos:    6470,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 353,
									EndLine:   353,
									StartPos:  6474,
									EndPos:    6474,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 354,
					EndLine:   354,
					StartPos:  6479,
					EndPos:    6497,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 354,
							EndLine:   354,
							StartPos:  6486,
							EndPos:    6496,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 354,
								EndLine:   354,
								StartPos:  6486,
								EndPos:    6487,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 354,
									EndLine:   354,
									StartPos:  6486,
									EndPos:    6487,
								},
								Value: "a",
							},
						},
						Expr: &binary.Pow{
							Position: &position.Position{
								StartLine: 354,
								EndLine:   354,
								StartPos:  6491,
								EndPos:    6496,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 354,
									EndLine:   354,
									StartPos:  6491,
									EndPos:    6491,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 354,
									EndLine:   354,
									StartPos:  6496,
									EndPos:    6496,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 355,
					EndLine:   355,
					StartPos:  6501,
					EndPos:    6519,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 355,
							EndLine:   355,
							StartPos:  6508,
							EndPos:    6518,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 355,
								EndLine:   355,
								StartPos:  6508,
								EndPos:    6509,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 355,
									EndLine:   355,
									StartPos:  6508,
									EndPos:    6509,
								},
								Value: "a",
							},
						},
						Expr: &binary.ShiftLeft{
							Position: &position.Position{
								StartLine: 355,
								EndLine:   355,
								StartPos:  6513,
								EndPos:    6518,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 355,
									EndLine:   355,
									StartPos:  6513,
									EndPos:    6513,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 355,
									EndLine:   355,
									StartPos:  6518,
									EndPos:    6518,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 356,
					EndLine:   356,
					StartPos:  6523,
					EndPos:    6541,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 356,
							EndLine:   356,
							StartPos:  6530,
							EndPos:    6540,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 356,
								EndLine:   356,
								StartPos:  6530,
								EndPos:    6531,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 356,
									EndLine:   356,
									StartPos:  6530,
									EndPos:    6531,
								},
								Value: "a",
							},
						},
						Expr: &binary.ShiftRight{
							Position: &position.Position{
								StartLine: 356,
								EndLine:   356,
								StartPos:  6535,
								EndPos:    6540,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 356,
									EndLine:   356,
									StartPos:  6535,
									EndPos:    6535,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 356,
									EndLine:   356,
									StartPos:  6540,
									EndPos:    6540,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 357,
					EndLine:   357,
					StartPos:  6545,
					EndPos:    6563,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 357,
							EndLine:   357,
							StartPos:  6552,
							EndPos:    6562,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 357,
								EndLine:   357,
								StartPos:  6552,
								EndPos:    6553,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 357,
									EndLine:   357,
									StartPos:  6552,
									EndPos:    6553,
								},
								Value: "a",
							},
						},
						Expr: &binary.SmallerOrEqual{
							Position: &position.Position{
								StartLine: 357,
								EndLine:   357,
								StartPos:  6557,
								EndPos:    6562,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 357,
									EndLine:   357,
									StartPos:  6557,
									EndPos:    6557,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 357,
									EndLine:   357,
									StartPos:  6562,
									EndPos:    6562,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 358,
					EndLine:   358,
					StartPos:  6567,
					EndPos:    6584,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 358,
							EndLine:   358,
							StartPos:  6574,
							EndPos:    6583,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 358,
								EndLine:   358,
								StartPos:  6574,
								EndPos:    6575,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 358,
									EndLine:   358,
									StartPos:  6574,
									EndPos:    6575,
								},
								Value: "a",
							},
						},
						Expr: &binary.Smaller{
							Position: &position.Position{
								StartLine: 358,
								EndLine:   358,
								StartPos:  6579,
								EndPos:    6583,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 358,
									EndLine:   358,
									StartPos:  6579,
									EndPos:    6579,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 358,
									EndLine:   358,
									StartPos:  6583,
									EndPos:    6583,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 359,
					EndLine:   359,
					StartPos:  6588,
					EndPos:    6608,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 359,
							EndLine:   359,
							StartPos:  6595,
							EndPos:    6607,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 359,
								EndLine:   359,
								StartPos:  6595,
								EndPos:    6596,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 359,
									EndLine:   359,
									StartPos:  6595,
									EndPos:    6596,
								},
								Value: "a",
							},
						},
						Expr: &expr.ClassConstFetch{
							Position: &position.Position{
								StartLine: 359,
								EndLine:   359,
								StartPos:  6600,
								EndPos:    6607,
							},
							Class: &name.Name{
								Position: &position.Position{
									StartLine: 359,
									EndLine:   359,
									StartPos:  6600,
									EndPos:    6602,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 359,
											EndLine:   359,
											StartPos:  6600,
											EndPos:    6602,
										},
										Value: "Foo",
									},
								},
							},
							ConstantName: &node.Identifier{
								Position: &position.Position{
									StartLine: 359,
									EndLine:   359,
									StartPos:  6605,
									EndPos:    6607,
								},
								Value: "bar",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 360,
					EndLine:   360,
					StartPos:  6612,
					EndPos:    6634,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 360,
							EndLine:   360,
							StartPos:  6619,
							EndPos:    6633,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 360,
								EndLine:   360,
								StartPos:  6619,
								EndPos:    6620,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 360,
									EndLine:   360,
									StartPos:  6619,
									EndPos:    6620,
								},
								Value: "a",
							},
						},
						Expr: &expr.ClassConstFetch{
							Position: &position.Position{
								StartLine: 360,
								EndLine:   360,
								StartPos:  6624,
								EndPos:    6633,
							},
							Class: &name.Name{
								Position: &position.Position{
									StartLine: 360,
									EndLine:   360,
									StartPos:  6624,
									EndPos:    6626,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 360,
											EndLine:   360,
											StartPos:  6624,
											EndPos:    6626,
										},
										Value: "Foo",
									},
								},
							},
							ConstantName: &node.Identifier{
								Position: &position.Position{
									StartLine: 360,
									EndLine:   360,
									StartPos:  6629,
									EndPos:    6633,
								},
								Value: "class",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 361,
					EndLine:   361,
					StartPos:  6638,
					EndPos:    6659,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 361,
							EndLine:   361,
							StartPos:  6645,
							EndPos:    6658,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 361,
								EndLine:   361,
								StartPos:  6645,
								EndPos:    6646,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 361,
									EndLine:   361,
									StartPos:  6645,
									EndPos:    6646,
								},
								Value: "a",
							},
						},
						Expr: &scalar.MagicConstant{
							Position: &position.Position{
								StartLine: 361,
								EndLine:   361,
								StartPos:  6650,
								EndPos:    6658,
							},
							Value: "__CLASS__",
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 362,
					EndLine:   362,
					StartPos:  6663,
					EndPos:    6678,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 362,
							EndLine:   362,
							StartPos:  6670,
							EndPos:    6677,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 362,
								EndLine:   362,
								StartPos:  6670,
								EndPos:    6671,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 362,
									EndLine:   362,
									StartPos:  6670,
									EndPos:    6671,
								},
								Value: "a",
							},
						},
						Expr: &expr.ConstFetch{
							Position: &position.Position{
								StartLine: 362,
								EndLine:   362,
								StartPos:  6675,
								EndPos:    6677,
							},
							Constant: &name.Name{
								Position: &position.Position{
									StartLine: 362,
									EndLine:   362,
									StartPos:  6675,
									EndPos:    6677,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 362,
											EndLine:   362,
											StartPos:  6675,
											EndPos:    6677,
										},
										Value: "Foo",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 363,
					EndLine:   363,
					StartPos:  6682,
					EndPos:    6707,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 363,
							EndLine:   363,
							StartPos:  6689,
							EndPos:    6706,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 363,
								EndLine:   363,
								StartPos:  6689,
								EndPos:    6690,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 363,
									EndLine:   363,
									StartPos:  6689,
									EndPos:    6690,
								},
								Value: "a",
							},
						},
						Expr: &expr.ConstFetch{
							Position: &position.Position{
								StartLine: 363,
								EndLine:   363,
								StartPos:  6694,
								EndPos:    6706,
							},
							Constant: &name.Relative{
								Position: &position.Position{
									StartLine: 363,
									EndLine:   363,
									StartPos:  6694,
									EndPos:    6706,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 363,
											EndLine:   363,
											StartPos:  6704,
											EndPos:    6706,
										},
										Value: "Foo",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 364,
					EndLine:   364,
					StartPos:  6711,
					EndPos:    6727,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 364,
							EndLine:   364,
							StartPos:  6718,
							EndPos:    6726,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 364,
								EndLine:   364,
								StartPos:  6718,
								EndPos:    6719,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 364,
									EndLine:   364,
									StartPos:  6718,
									EndPos:    6719,
								},
								Value: "a",
							},
						},
						Expr: &expr.ConstFetch{
							Position: &position.Position{
								StartLine: 364,
								EndLine:   364,
								StartPos:  6723,
								EndPos:    6726,
							},
							Constant: &name.FullyQualified{
								Position: &position.Position{
									StartLine: 364,
									EndLine:   364,
									StartPos:  6723,
									EndPos:    6726,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 364,
											EndLine:   364,
											StartPos:  6724,
											EndPos:    6726,
										},
										Value: "Foo",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 365,
					EndLine:   365,
					StartPos:  6731,
					EndPos:    6750,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 365,
							EndLine:   365,
							StartPos:  6738,
							EndPos:    6749,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 365,
								EndLine:   365,
								StartPos:  6738,
								EndPos:    6739,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 365,
									EndLine:   365,
									StartPos:  6738,
									EndPos:    6739,
								},
								Value: "a",
							},
						},
						Expr: &expr.Array{
							Position: &position.Position{
								StartLine: 365,
								EndLine:   365,
								StartPos:  6743,
								EndPos:    6749,
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 366,
					EndLine:   366,
					StartPos:  6754,
					EndPos:    6782,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 366,
							EndLine:   366,
							StartPos:  6761,
							EndPos:    6781,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 366,
								EndLine:   366,
								StartPos:  6761,
								EndPos:    6762,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 366,
									EndLine:   366,
									StartPos:  6761,
									EndPos:    6762,
								},
								Value: "a",
							},
						},
						Expr: &expr.Array{
							Position: &position.Position{
								StartLine: 366,
								EndLine:   366,
								StartPos:  6766,
								EndPos:    6781,
							},
							Items: []node.Node{
								&expr.ArrayItem{
									Position: &position.Position{
										StartLine: 366,
										EndLine:   366,
										StartPos:  6772,
										EndPos:    6777,
									},
									Key: &scalar.Lnumber{
										Position: &position.Position{
											StartLine: 366,
											EndLine:   366,
											StartPos:  6772,
											EndPos:    6772,
										},
										Value: "1",
									},
									Val: &scalar.Lnumber{
										Position: &position.Position{
											StartLine: 366,
											EndLine:   366,
											StartPos:  6777,
											EndPos:    6777,
										},
										Value: "1",
									},
								},
								&expr.ArrayItem{
									Position: &position.Position{
										StartLine: 366,
										EndLine:   366,
										StartPos:  6780,
										EndPos:    6780,
									},
									Val: &scalar.Lnumber{
										Position: &position.Position{
											StartLine: 366,
											EndLine:   366,
											StartPos:  6780,
											EndPos:    6780,
										},
										Value: "2",
									},
								},
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 367,
					EndLine:   367,
					StartPos:  6786,
					EndPos:    6812,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 367,
							EndLine:   367,
							StartPos:  6793,
							EndPos:    6811,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 367,
								EndLine:   367,
								StartPos:  6793,
								EndPos:    6794,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 367,
									EndLine:   367,
									StartPos:  6793,
									EndPos:    6794,
								},
								Value: "a",
							},
						},
						Expr: &expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 367,
								EndLine:   367,
								StartPos:  6798,
								EndPos:    6811,
							},
							Variable: &expr.ShortArray{
								Position: &position.Position{
									StartLine: 367,
									EndLine:   367,
									StartPos:  6798,
									EndPos:    6808,
								},
								Items: []node.Node{
									&expr.ArrayItem{
										Position: &position.Position{
											StartLine: 367,
											EndLine:   367,
											StartPos:  6799,
											EndPos:    6799,
										},
										Val: &scalar.Lnumber{
											Position: &position.Position{
												StartLine: 367,
												EndLine:   367,
												StartPos:  6799,
												EndPos:    6799,
											},
											Value: "1",
										},
									},
									&expr.ArrayItem{
										Position: &position.Position{
											StartLine: 367,
											EndLine:   367,
											StartPos:  6802,
											EndPos:    6807,
										},
										Key: &scalar.Lnumber{
											Position: &position.Position{
												StartLine: 367,
												EndLine:   367,
												StartPos:  6802,
												EndPos:    6802,
											},
											Value: "2",
										},
										Val: &scalar.Lnumber{
											Position: &position.Position{
												StartLine: 367,
												EndLine:   367,
												StartPos:  6807,
												EndPos:    6807,
											},
											Value: "2",
										},
									},
								},
							},
							Dim: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 367,
									EndLine:   367,
									StartPos:  6810,
									EndPos:    6810,
								},
								Value: "0",
							},
						},
					},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 369,
					EndLine:   369,
					StartPos:  6817,
					EndPos:    6831,
				},
				Cond: &expr.Yield{
					Position: &position.Position{
						StartLine: 369,
						EndLine:   369,
						StartPos:  6821,
						EndPos:    6827,
					},
					Value: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 369,
							EndLine:   369,
							StartPos:  6827,
							EndPos:    6827,
						},
						Value: "1",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 369,
						EndLine:   369,
						StartPos:  6830,
						EndPos:    6831,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 370,
					EndLine:   370,
					StartPos:  6835,
					EndPos:    6845,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 370,
						EndLine:   370,
						StartPos:  6835,
						EndPos:    6844,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 370,
							EndLine:   370,
							StartPos:  6835,
							EndPos:    6837,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 370,
									EndLine:   370,
									StartPos:  6835,
									EndPos:    6837,
								},
								Value: "Foo",
							},
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 370,
							EndLine:   370,
							StartPos:  6840,
							EndPos:    6844,
						},
						VarName: &expr.Variable{
							Position: &position.Position{
								StartLine: 370,
								EndLine:   370,
								StartPos:  6841,
								EndPos:    6844,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 370,
									EndLine:   370,
									StartPos:  6841,
									EndPos:    6844,
								},
								Value: "bar",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 372,
					EndLine:   372,
					StartPos:  6850,
					EndPos:    6856,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 372,
						EndLine:   372,
						StartPos:  6850,
						EndPos:    6855,
					},
					Function: &expr.Variable{
						Position: &position.Position{
							StartLine: 372,
							EndLine:   372,
							StartPos:  6850,
							EndPos:    6853,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 372,
								EndLine:   372,
								StartPos:  6850,
								EndPos:    6853,
							},
							Value: "foo",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 372,
							EndLine:   372,
							StartPos:  6854,
							EndPos:    6855,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 373,
					EndLine:   373,
					StartPos:  6860,
					EndPos:    6872,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 373,
						EndLine:   373,
						StartPos:  6860,
						EndPos:    6871,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 373,
							EndLine:   373,
							StartPos:  6860,
							EndPos:    6868,
						},
						Variable: &expr.FunctionCall{
							Position: &position.Position{
								StartLine: 373,
								EndLine:   373,
								StartPos:  6860,
								EndPos:    6865,
							},
							Function: &expr.Variable{
								Position: &position.Position{
									StartLine: 373,
									EndLine:   373,
									StartPos:  6860,
									EndPos:    6863,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 373,
										EndLine:   373,
										StartPos:  6860,
										EndPos:    6863,
									},
									Value: "foo",
								},
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 373,
									EndLine:   373,
									StartPos:  6864,
									EndPos:    6865,
								},
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 373,
								EndLine:   373,
								StartPos:  6867,
								EndPos:    6867,
							},
							Value: "0",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 373,
							EndLine:   373,
							StartPos:  6870,
							EndPos:    6870,
						},
						Value: "0",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 374,
					EndLine:   374,
					StartPos:  6876,
					EndPos:    6882,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 374,
						EndLine:   374,
						StartPos:  6876,
						EndPos:    6881,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 374,
							EndLine:   374,
							StartPos:  6876,
							EndPos:    6877,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 374,
								EndLine:   374,
								StartPos:  6876,
								EndPos:    6877,
							},
							Value: "a",
						},
					},
					Dim: &expr.Variable{
						Position: &position.Position{
							StartLine: 374,
							EndLine:   374,
							StartPos:  6879,
							EndPos:    6880,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 374,
								EndLine:   374,
								StartPos:  6879,
								EndPos:    6880,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 375,
					EndLine:   375,
					StartPos:  6886,
					EndPos:    6891,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 375,
						EndLine:   375,
						StartPos:  6886,
						EndPos:    6890,
					},
					VarName: &expr.Variable{
						Position: &position.Position{
							StartLine: 375,
							EndLine:   375,
							StartPos:  6888,
							EndPos:    6889,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 375,
								EndLine:   375,
								StartPos:  6888,
								EndPos:    6889,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 376,
					EndLine:   376,
					StartPos:  6895,
					EndPos:    6909,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 376,
						EndLine:   376,
						StartPos:  6895,
						EndPos:    6908,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 376,
							EndLine:   376,
							StartPos:  6895,
							EndPos:    6898,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 376,
								EndLine:   376,
								StartPos:  6895,
								EndPos:    6898,
							},
							Value: "foo",
						},
					},
					Call: &expr.Variable{
						Position: &position.Position{
							StartLine: 376,
							EndLine:   376,
							StartPos:  6901,
							EndPos:    6906,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 376,
								EndLine:   376,
								StartPos:  6902,
								EndPos:    6905,
							},
							Value: "bar",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 376,
							EndLine:   376,
							StartPos:  6907,
							EndPos:    6908,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 377,
					EndLine:   377,
					StartPos:  6913,
					EndPos:    6922,
				},
				Expr: &expr.ClassConstFetch{
					Position: &position.Position{
						StartLine: 377,
						EndLine:   377,
						StartPos:  6913,
						EndPos:    6921,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 377,
							EndLine:   377,
							StartPos:  6913,
							EndPos:    6916,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 377,
								EndLine:   377,
								StartPos:  6913,
								EndPos:    6916,
							},
							Value: "foo",
						},
					},
					ConstantName: &node.Identifier{
						Position: &position.Position{
							StartLine: 377,
							EndLine:   377,
							StartPos:  6919,
							EndPos:    6921,
						},
						Value: "bar",
					},
				},
			},
			&stmt.HaltCompiler{
				Position: &position.Position{
					StartLine: 379,
					EndLine:   379,
					StartPos:  6927,
					EndPos:    6944,
				},
			},
		},
	}

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual := php5parser.GetRootNode()
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

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual := php5parser.GetRootNode()
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

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual := php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPhp5ControlCharsErrors(t *testing.T) {
	src := "<?php \004 echo $b; \"$a[\005test]\";"

	expected := []*errors.Error{
		{
			Msg: "WARNING: Unexpected character in input: '\004' (ASCII=4)",
			Pos: &position.Position{1, 1, 7, 7},
		},
		{
			Msg: "WARNING: Unexpected character in input: '\005' (ASCII=5)",
			Pos: &position.Position{1, 1, 22, 22},
		},
	}

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual := php5parser.GetErrors()
	assert.DeepEqual(t, expected, actual)
}
