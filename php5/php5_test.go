package php5_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/assign"
	"github.com/z7zmey/php-parser/node/expr/binary"
	"github.com/z7zmey/php-parser/node/expr/cast"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
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
			__halt_compiler();
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
	`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   379,
			StartPos:  6,
			EndPos:    6965,
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
								IsReference: false,
								Variadic:    false,
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
						ReturnsRef:    false,
						PhpDocComment: "",
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
								ByRef:    true,
								Variadic: true,
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
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 11,
								EndLine:   11,
								StartPos:  259,
								EndPos:    271,
							},
							Variadic: false,
							ByRef:    false,
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
							ByRef:    true,
							Variadic: true,
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
					ReturnsRef:    false,
					Static:        true,
					PhpDocComment: "",
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 12,
								EndLine:   12,
								StartPos:  310,
								EndPos:    322,
							},
							ByRef:    false,
							Variadic: false,
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
						StartLine: 73,
						EndLine:   73,
						StartPos:  1700,
						EndPos:    1712,
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
					EndLine:   93,
					StartPos:  2062,
					EndPos:    2154,
				},
				ReturnsRef:    false,
				PhpDocComment: "",
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
					&stmt.HaltCompiler{
						Position: &position.Position{
							StartLine: 89,
							EndLine:   89,
							StartPos:  2082,
							EndPos:    2099,
						},
					},
					&stmt.Function{
						Position: &position.Position{
							StartLine: 90,
							EndLine:   90,
							StartPos:  2104,
							EndPos:    2120,
						},
						ReturnsRef:    false,
						PhpDocComment: "",
						FunctionName: &node.Identifier{
							Position: &position.Position{
								StartLine: 90,
								EndLine:   90,
								StartPos:  2113,
								EndPos:    2115,
							},
							Value: "bar",
						},
						Stmts: []node.Node{},
					},
					&stmt.Class{
						Position: &position.Position{
							StartLine: 91,
							EndLine:   91,
							StartPos:  2125,
							EndPos:    2136,
						},
						PhpDocComment: "",
						ClassName: &node.Identifier{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2131,
								EndPos:    2133,
							},
							Value: "Baz",
						},
						Stmts: []node.Node{},
					},
					&stmt.Return{
						Position: &position.Position{
							StartLine: 92,
							EndLine:   92,
							StartPos:  2141,
							EndPos:    2150,
						},
						Expr: &expr.Variable{
							Position: &position.Position{
								StartLine: 92,
								EndLine:   92,
								StartPos:  2148,
								EndPos:    2149,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  2148,
									EndPos:    2149,
								},
								Value: "a",
							},
						},
					},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 95,
					EndLine:   95,
					StartPos:  2161,
					EndPos:    2205,
				},
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 95,
						EndLine:   95,
						StartPos:  2170,
						EndPos:    2172,
					},
					Value: "foo",
				},
				Params: []node.Node{
					&node.Parameter{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2174,
							EndPos:    2181,
						},
						ByRef:    false,
						Variadic: false,
						VariableType: &node.Identifier{
							Position: &position.Position{
								StartLine: 95,
								EndLine:   95,
								StartPos:  2174,
								EndPos:    2178,
							},
							Value: "array",
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 95,
								EndLine:   95,
								StartPos:  2180,
								EndPos:    2181,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 95,
									EndLine:   95,
									StartPos:  2180,
									EndPos:    2181,
								},
								Value: "a",
							},
						},
					},
					&node.Parameter{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2184,
							EndPos:    2194,
						},
						ByRef:    false,
						Variadic: false,
						VariableType: &node.Identifier{
							Position: &position.Position{
								StartLine: 95,
								EndLine:   95,
								StartPos:  2184,
								EndPos:    2191,
							},
							Value: "callable",
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 95,
								EndLine:   95,
								StartPos:  2193,
								EndPos:    2194,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 95,
									EndLine:   95,
									StartPos:  2193,
									EndPos:    2194,
								},
								Value: "b",
							},
						},
					},
				},
				Stmts: []node.Node{
					&stmt.Return{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2198,
							EndPos:    2204,
						},
					},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 96,
					EndLine:   96,
					StartPos:  2209,
					EndPos:    2235,
				},
				ReturnsRef:    true,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 96,
						EndLine:   96,
						StartPos:  2219,
						EndPos:    2221,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.Return{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2226,
							EndPos:    2234,
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 96,
								EndLine:   96,
								StartPos:  2233,
								EndPos:    2233,
							},
							Value: "1",
						},
					},
				},
			},
			&stmt.Function{
				Position: &position.Position{
					StartLine: 97,
					EndLine:   97,
					StartPos:  2239,
					EndPos:    2256,
				},
				ReturnsRef:    true,
				PhpDocComment: "",
				FunctionName: &node.Identifier{
					Position: &position.Position{
						StartLine: 97,
						EndLine:   97,
						StartPos:  2249,
						EndPos:    2251,
					},
					Value: "foo",
				},
				Stmts: []node.Node{},
			},
			&stmt.Global{
				Position: &position.Position{
					StartLine: 98,
					EndLine:   98,
					StartPos:  2260,
					EndPos:    2288,
				},
				Vars: []node.Node{
					&expr.Variable{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2267,
							EndPos:    2268,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 98,
								EndLine:   98,
								StartPos:  2267,
								EndPos:    2268,
							},
							Value: "a",
						},
					},
					&expr.Variable{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2271,
							EndPos:    2272,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 98,
								EndLine:   98,
								StartPos:  2271,
								EndPos:    2272,
							},
							Value: "b",
						},
					},
					&expr.Variable{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2275,
							EndPos:    2277,
						},
						VarName: &expr.Variable{
							Position: &position.Position{
								StartLine: 98,
								EndLine:   98,
								StartPos:  2276,
								EndPos:    2277,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 98,
									EndLine:   98,
									StartPos:  2276,
									EndPos:    2277,
								},
								Value: "c",
							},
						},
					},
					&expr.Variable{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2280,
							EndPos:    2287,
						},
						VarName: &expr.FunctionCall{
							Position: &position.Position{
								StartLine: 98,
								EndLine:   98,
								StartPos:  2282,
								EndPos:    2286,
							},
							Function: &name.Name{
								Position: &position.Position{
									StartLine: 98,
									EndLine:   98,
									StartPos:  2282,
									EndPos:    2284,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 98,
											EndLine:   98,
											StartPos:  2282,
											EndPos:    2284,
										},
										Value: "foo",
									},
								},
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 98,
									EndLine:   98,
									StartPos:  2285,
									EndPos:    2286,
								},
							},
						},
					},
				},
			},
			&stmt.Label{
				Position: &position.Position{
					StartLine: 99,
					EndLine:   99,
					StartPos:  2292,
					EndPos:    2293,
				},
				LabelName: &node.Identifier{
					Position: &position.Position{
						StartLine: 99,
						EndLine:   99,
						StartPos:  2292,
						EndPos:    2292,
					},
					Value: "a",
				},
			},
			&stmt.Goto{
				Position: &position.Position{
					StartLine: 100,
					EndLine:   100,
					StartPos:  2298,
					EndPos:    2304,
				},
				Label: &node.Identifier{
					Position: &position.Position{
						StartLine: 100,
						EndLine:   100,
						StartPos:  2303,
						EndPos:    2303,
					},
					Value: "a",
				},
			},
			&stmt.HaltCompiler{
				Position: &position.Position{
					StartLine: 101,
					EndLine:   101,
					StartPos:  2308,
					EndPos:    2325,
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 102,
					EndLine:   102,
					StartPos:  2329,
					EndPos:    2338,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 102,
						EndLine:   102,
						StartPos:  2333,
						EndPos:    2334,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 102,
							EndLine:   102,
							StartPos:  2333,
							EndPos:    2334,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 102,
						EndLine:   102,
						StartPos:  2337,
						EndPos:    2338,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 103,
					EndLine:   103,
					StartPos:  2342,
					EndPos:    2366,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 103,
						EndLine:   103,
						StartPos:  2346,
						EndPos:    2347,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2346,
							EndPos:    2347,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 103,
						EndLine:   103,
						StartPos:  2350,
						EndPos:    2351,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2353,
							EndPos:    2366,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2361,
								EndPos:    2362,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2361,
									EndPos:    2362,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2365,
								EndPos:    2366,
							},
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 104,
					EndLine:   104,
					StartPos:  2370,
					EndPos:    2387,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 104,
						EndLine:   104,
						StartPos:  2374,
						EndPos:    2375,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 104,
							EndLine:   104,
							StartPos:  2374,
							EndPos:    2375,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 104,
						EndLine:   104,
						StartPos:  2378,
						EndPos:    2379,
					},
					Stmts: []node.Node{},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 104,
						EndLine:   104,
						StartPos:  2381,
						EndPos:    2387,
					},
					Stmt: &stmt.StmtList{
						Position: &position.Position{
							StartLine: 104,
							EndLine:   104,
							StartPos:  2386,
							EndPos:    2387,
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 105,
					EndLine:   105,
					StartPos:  2391,
					EndPos:    2438,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 105,
						EndLine:   105,
						StartPos:  2395,
						EndPos:    2396,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 105,
							EndLine:   105,
							StartPos:  2395,
							EndPos:    2396,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 105,
						EndLine:   105,
						StartPos:  2399,
						EndPos:    2400,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 105,
							EndLine:   105,
							StartPos:  2402,
							EndPos:    2415,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 105,
								EndLine:   105,
								StartPos:  2410,
								EndPos:    2411,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 105,
									EndLine:   105,
									StartPos:  2410,
									EndPos:    2411,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 105,
								EndLine:   105,
								StartPos:  2414,
								EndPos:    2415,
							},
							Stmts: []node.Node{},
						},
					},
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 105,
							EndLine:   105,
							StartPos:  2417,
							EndPos:    2430,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 105,
								EndLine:   105,
								StartPos:  2425,
								EndPos:    2426,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 105,
									EndLine:   105,
									StartPos:  2425,
									EndPos:    2426,
								},
								Value: "c",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 105,
								EndLine:   105,
								StartPos:  2429,
								EndPos:    2430,
							},
							Stmts: []node.Node{},
						},
					},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 105,
						EndLine:   105,
						StartPos:  2432,
						EndPos:    2438,
					},
					Stmt: &stmt.StmtList{
						Position: &position.Position{
							StartLine: 105,
							EndLine:   105,
							StartPos:  2437,
							EndPos:    2438,
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 106,
					EndLine:   106,
					StartPos:  2442,
					EndPos:    2490,
				},
				Cond: &expr.Variable{
					Position: &position.Position{
						StartLine: 106,
						EndLine:   106,
						StartPos:  2446,
						EndPos:    2447,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 106,
							EndLine:   106,
							StartPos:  2446,
							EndPos:    2447,
						},
						Value: "a",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 106,
						EndLine:   106,
						StartPos:  2450,
						EndPos:    2451,
					},
					Stmts: []node.Node{},
				},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Position: &position.Position{
							StartLine: 106,
							EndLine:   106,
							StartPos:  2453,
							EndPos:    2466,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 106,
								EndLine:   106,
								StartPos:  2461,
								EndPos:    2462,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 106,
									EndLine:   106,
									StartPos:  2461,
									EndPos:    2462,
								},
								Value: "b",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 106,
								EndLine:   106,
								StartPos:  2465,
								EndPos:    2466,
							},
							Stmts: []node.Node{},
						},
					},
				},
				Else: &stmt.Else{
					Position: &position.Position{
						StartLine: 106,
						EndLine:   106,
						StartPos:  2468,
						EndPos:    2490,
					},
					Stmt: &stmt.If{
						Position: &position.Position{
							StartLine: 106,
							EndLine:   106,
							StartPos:  2473,
							EndPos:    2490,
						},
						Cond: &expr.Variable{
							Position: &position.Position{
								StartLine: 106,
								EndLine:   106,
								StartPos:  2477,
								EndPos:    2478,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 106,
									EndLine:   106,
									StartPos:  2477,
									EndPos:    2478,
								},
								Value: "c",
							},
						},
						Stmt: &stmt.StmtList{
							Position: &position.Position{
								StartLine: 106,
								EndLine:   106,
								StartPos:  2481,
								EndPos:    2482,
							},
							Stmts: []node.Node{},
						},
						Else: &stmt.Else{
							Position: &position.Position{
								StartLine: 106,
								EndLine:   106,
								StartPos:  2484,
								EndPos:    2490,
							},
							Stmt: &stmt.StmtList{
								Position: &position.Position{
									StartLine: 106,
									EndLine:   106,
									StartPos:  2489,
									EndPos:    2490,
								},
								Stmts: []node.Node{},
							},
						},
					},
				},
			},
			&stmt.Nop{
				Position: &position.Position{
					StartLine: 107,
					EndLine:   107,
					StartPos:  2494,
					EndPos:    2495,
				},
			},
			&stmt.InlineHtml{
				Position: &position.Position{
					StartLine: 107,
					EndLine:   107,
					StartPos:  2497,
					EndPos:    2508,
				},
				Value: "<div></div> ",
			},
			&stmt.Interface{
				Position: &position.Position{
					StartLine: 108,
					EndLine:   108,
					StartPos:  2514,
					EndPos:    2529,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 108,
						EndLine:   108,
						StartPos:  2524,
						EndPos:    2526,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{},
			},
			&stmt.Interface{
				Position: &position.Position{
					StartLine: 109,
					EndLine:   109,
					StartPos:  2533,
					EndPos:    2560,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 109,
						EndLine:   109,
						StartPos:  2543,
						EndPos:    2545,
					},
					Value: "Foo",
				},
				Extends: &stmt.InterfaceExtends{
					Position: &position.Position{
						StartLine: 109,
						EndLine:   109,
						StartPos:  2547,
						EndPos:    2557,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 109,
								EndLine:   109,
								StartPos:  2555,
								EndPos:    2557,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 109,
										EndLine:   109,
										StartPos:  2555,
										EndPos:    2557,
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
					StartLine: 110,
					EndLine:   110,
					StartPos:  2564,
					EndPos:    2596,
				},
				PhpDocComment: "",
				InterfaceName: &node.Identifier{
					Position: &position.Position{
						StartLine: 110,
						EndLine:   110,
						StartPos:  2574,
						EndPos:    2576,
					},
					Value: "Foo",
				},
				Extends: &stmt.InterfaceExtends{
					Position: &position.Position{
						StartLine: 110,
						EndLine:   110,
						StartPos:  2578,
						EndPos:    2593,
					},
					InterfaceNames: []node.Node{
						&name.Name{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2586,
								EndPos:    2588,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 110,
										EndLine:   110,
										StartPos:  2586,
										EndPos:    2588,
									},
									Value: "Bar",
								},
							},
						},
						&name.Name{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2591,
								EndPos:    2593,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 110,
										EndLine:   110,
										StartPos:  2591,
										EndPos:    2593,
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
					StartLine: 111,
					EndLine:   111,
					StartPos:  2600,
					EndPos:    2613,
				},
				NamespaceName: &name.Name{
					Position: &position.Position{
						StartLine: 111,
						EndLine:   111,
						StartPos:  2610,
						EndPos:    2612,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 111,
								EndLine:   111,
								StartPos:  2610,
								EndPos:    2612,
							},
							Value: "Foo",
						},
					},
				},
			},
			&stmt.Namespace{
				Position: &position.Position{
					StartLine: 112,
					EndLine:   112,
					StartPos:  2617,
					EndPos:    2636,
				},
				NamespaceName: &name.Name{
					Position: &position.Position{
						StartLine: 112,
						EndLine:   112,
						StartPos:  2627,
						EndPos:    2633,
					},
					Parts: []node.Node{
						&name.NamePart{
							Position: &position.Position{
								StartLine: 112,
								EndLine:   112,
								StartPos:  2627,
								EndPos:    2629,
							},
							Value: "Foo",
						},
						&name.NamePart{
							Position: &position.Position{
								StartLine: 112,
								EndLine:   112,
								StartPos:  2631,
								EndPos:    2633,
							},
							Value: "Bar",
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Namespace{
				Position: &position.Position{
					StartLine: 113,
					EndLine:   113,
					StartPos:  2640,
					EndPos:    2651,
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 114,
					EndLine:   114,
					StartPos:  2655,
					EndPos:    2673,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 114,
						EndLine:   114,
						StartPos:  2661,
						EndPos:    2663,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Position: &position.Position{
							StartLine: 114,
							EndLine:   114,
							StartPos:  2666,
							EndPos:    2672,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 114,
									EndLine:   114,
									StartPos:  2666,
									EndPos:    2668,
								},
								Value: "var",
							},
						},
						Properties: []node.Node{
							&stmt.Property{
								Position: &position.Position{
									StartLine: 114,
									EndLine:   114,
									StartPos:  2670,
									EndPos:    2671,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2670,
										EndPos:    2671,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 114,
											EndLine:   114,
											StartPos:  2670,
											EndPos:    2671,
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
					StartLine: 115,
					EndLine:   115,
					StartPos:  2677,
					EndPos:    2713,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 115,
						EndLine:   115,
						StartPos:  2683,
						EndPos:    2685,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Position: &position.Position{
							StartLine: 115,
							EndLine:   115,
							StartPos:  2688,
							EndPos:    2712,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 115,
									EndLine:   115,
									StartPos:  2688,
									EndPos:    2693,
								},
								Value: "public",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 115,
									EndLine:   115,
									StartPos:  2695,
									EndPos:    2700,
								},
								Value: "static",
							},
						},
						Properties: []node.Node{
							&stmt.Property{
								Position: &position.Position{
									StartLine: 115,
									EndLine:   115,
									StartPos:  2702,
									EndPos:    2703,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 115,
										EndLine:   115,
										StartPos:  2702,
										EndPos:    2703,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 115,
											EndLine:   115,
											StartPos:  2702,
											EndPos:    2703,
										},
										Value: "a",
									},
								},
							},
							&stmt.Property{
								Position: &position.Position{
									StartLine: 115,
									EndLine:   115,
									StartPos:  2706,
									EndPos:    2711,
								},
								PhpDocComment: "",
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 115,
										EndLine:   115,
										StartPos:  2706,
										EndPos:    2707,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 115,
											EndLine:   115,
											StartPos:  2706,
											EndPos:    2707,
										},
										Value: "b",
									},
								},
								Expr: &scalar.Lnumber{
									Position: &position.Position{
										StartLine: 115,
										EndLine:   115,
										StartPos:  2711,
										EndPos:    2711,
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
					StartLine: 116,
					EndLine:   116,
					StartPos:  2717,
					EndPos:    2753,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 116,
						EndLine:   116,
						StartPos:  2723,
						EndPos:    2725,
					},
					Value: "foo",
				},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Position: &position.Position{
							StartLine: 116,
							EndLine:   116,
							StartPos:  2728,
							EndPos:    2752,
						},
						Modifiers: []node.Node{
							&node.Identifier{
								Position: &position.Position{
									StartLine: 116,
									EndLine:   116,
									StartPos:  2728,
									EndPos:    2733,
								},
								Value: "public",
							},
							&node.Identifier{
								Position: &position.Position{
									StartLine: 116,
									EndLine:   116,
									StartPos:  2735,
									EndPos:    2740,
								},
								Value: "static",
							},
						},
						Properties: []node.Node{
							&stmt.Property{
								Position: &position.Position{
									StartLine: 116,
									EndLine:   116,
									StartPos:  2742,
									EndPos:    2747,
								},
								PhpDocComment: "",
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
							&stmt.Property{
								Position: &position.Position{
									StartLine: 116,
									EndLine:   116,
									StartPos:  2750,
									EndPos:    2751,
								},
								PhpDocComment: "",
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
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 117,
					EndLine:   117,
					StartPos:  2757,
					EndPos:    2774,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 117,
							EndLine:   117,
							StartPos:  2764,
							EndPos:    2765,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 117,
								EndLine:   117,
								StartPos:  2764,
								EndPos:    2765,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 117,
									EndLine:   117,
									StartPos:  2764,
									EndPos:    2765,
								},
								Value: "a",
							},
						},
					},
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 117,
							EndLine:   117,
							StartPos:  2768,
							EndPos:    2773,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 117,
								EndLine:   117,
								StartPos:  2768,
								EndPos:    2769,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 117,
									EndLine:   117,
									StartPos:  2768,
									EndPos:    2769,
								},
								Value: "b",
							},
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 117,
								EndLine:   117,
								StartPos:  2773,
								EndPos:    2773,
							},
							Value: "1",
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 118,
					EndLine:   118,
					StartPos:  2778,
					EndPos:    2795,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 118,
							EndLine:   118,
							StartPos:  2785,
							EndPos:    2790,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 118,
								EndLine:   118,
								StartPos:  2785,
								EndPos:    2786,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 118,
									EndLine:   118,
									StartPos:  2785,
									EndPos:    2786,
								},
								Value: "a",
							},
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 118,
								EndLine:   118,
								StartPos:  2790,
								EndPos:    2790,
							},
							Value: "1",
						},
					},
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 118,
							EndLine:   118,
							StartPos:  2793,
							EndPos:    2794,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 118,
								EndLine:   118,
								StartPos:  2793,
								EndPos:    2794,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 118,
									EndLine:   118,
									StartPos:  2793,
									EndPos:    2794,
								},
								Value: "b",
							},
						},
					},
				},
			},
			&stmt.AltSwitch{
				Position: &position.Position{
					StartLine: 120,
					EndLine:   124,
					StartPos:  2800,
					EndPos:    2858,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 120,
						EndLine:   120,
						StartPos:  2808,
						EndPos:    2808,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 121,
						EndLine:   -1,
						StartPos:  2816,
						EndPos:    -1,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 121,
								EndLine:   -1,
								StartPos:  2816,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 121,
									EndLine:   121,
									StartPos:  2821,
									EndPos:    2821,
								},
								Value: "1",
							},
							Stmts: []node.Node{},
						},
						&stmt.Default{
							Position: &position.Position{
								StartLine: 122,
								EndLine:   -1,
								StartPos:  2827,
								EndPos:    -1,
							},
							Stmts: []node.Node{},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 123,
								EndLine:   -1,
								StartPos:  2839,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 123,
									EndLine:   123,
									StartPos:  2844,
									EndPos:    2844,
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
					StartLine: 126,
					EndLine:   129,
					StartPos:  2863,
					EndPos:    2910,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 126,
						EndLine:   126,
						StartPos:  2871,
						EndPos:    2871,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 127,
						EndLine:   -1,
						StartPos:  2880,
						EndPos:    -1,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 127,
								EndLine:   -1,
								StartPos:  2880,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 127,
									EndLine:   127,
									StartPos:  2885,
									EndPos:    2885,
								},
								Value: "1",
							},
							Stmts: []node.Node{},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 128,
								EndLine:   -1,
								StartPos:  2891,
								EndPos:    -1,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 128,
									EndLine:   128,
									StartPos:  2896,
									EndPos:    2896,
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
					StartLine: 131,
					EndLine:   134,
					StartPos:  2917,
					EndPos:    2968,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 131,
						EndLine:   131,
						StartPos:  2925,
						EndPos:    2925,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 131,
						EndLine:   134,
						StartPos:  2928,
						EndPos:    2968,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 132,
								EndLine:   132,
								StartPos:  2933,
								EndPos:    2946,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 132,
									EndLine:   132,
									StartPos:  2938,
									EndPos:    2938,
								},
								Value: "1",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 132,
										EndLine:   132,
										StartPos:  2941,
										EndPos:    2946,
									},
								},
							},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 133,
								EndLine:   133,
								StartPos:  2951,
								EndPos:    2964,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 133,
									EndLine:   133,
									StartPos:  2956,
									EndPos:    2956,
								},
								Value: "2",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 133,
										EndLine:   133,
										StartPos:  2959,
										EndPos:    2964,
									},
								},
							},
						},
					},
				},
			},
			&stmt.Switch{
				Position: &position.Position{
					StartLine: 136,
					EndLine:   139,
					StartPos:  2975,
					EndPos:    3027,
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 136,
						EndLine:   136,
						StartPos:  2983,
						EndPos:    2983,
					},
					Value: "1",
				},
				CaseList: &stmt.CaseList{
					Position: &position.Position{
						StartLine: 136,
						EndLine:   139,
						StartPos:  2986,
						EndPos:    3027,
					},
					Cases: []node.Node{
						&stmt.Case{
							Position: &position.Position{
								StartLine: 137,
								EndLine:   137,
								StartPos:  2992,
								EndPos:    3005,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 137,
									EndLine:   137,
									StartPos:  2997,
									EndPos:    2997,
								},
								Value: "1",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 137,
										EndLine:   137,
										StartPos:  3000,
										EndPos:    3005,
									},
								},
							},
						},
						&stmt.Case{
							Position: &position.Position{
								StartLine: 138,
								EndLine:   138,
								StartPos:  3010,
								EndPos:    3023,
							},
							Cond: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 138,
									EndLine:   138,
									StartPos:  3015,
									EndPos:    3015,
								},
								Value: "2",
							},
							Stmts: []node.Node{
								&stmt.Break{
									Position: &position.Position{
										StartLine: 138,
										EndLine:   138,
										StartPos:  3018,
										EndPos:    3023,
									},
								},
							},
						},
					},
				},
			},
			&stmt.Throw{
				Position: &position.Position{
					StartLine: 140,
					EndLine:   140,
					StartPos:  3031,
					EndPos:    3039,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 140,
						EndLine:   140,
						StartPos:  3037,
						EndPos:    3038,
					},
					VarName: &node.Identifier{
						Position: &position.Position{
							StartLine: 140,
							EndLine:   140,
							StartPos:  3037,
							EndPos:    3038,
						},
						Value: "e",
					},
				},
			},
			&stmt.Trait{
				Position: &position.Position{
					StartLine: 141,
					EndLine:   141,
					StartPos:  3043,
					EndPos:    3054,
				},
				PhpDocComment: "",
				TraitName: &node.Identifier{
					Position: &position.Position{
						StartLine: 141,
						EndLine:   141,
						StartPos:  3049,
						EndPos:    3051,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 142,
					EndLine:   142,
					StartPos:  3058,
					EndPos:    3079,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 142,
						EndLine:   142,
						StartPos:  3064,
						EndPos:    3066,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 142,
							EndLine:   142,
							StartPos:  3070,
							EndPos:    3077,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 142,
									EndLine:   142,
									StartPos:  3074,
									EndPos:    3076,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 142,
											EndLine:   142,
											StartPos:  3074,
											EndPos:    3076,
										},
										Value: "Bar",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.Nop{
							Position: &position.Position{
								StartLine: 142,
								EndLine:   142,
								StartPos:  3077,
								EndPos:    3077,
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 143,
					EndLine:   143,
					StartPos:  3083,
					EndPos:    3111,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 143,
						EndLine:   143,
						StartPos:  3089,
						EndPos:    3091,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 143,
							EndLine:   143,
							StartPos:  3095,
							EndPos:    3109,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 143,
									EndLine:   143,
									StartPos:  3099,
									EndPos:    3101,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 143,
											EndLine:   143,
											StartPos:  3099,
											EndPos:    3101,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 143,
									EndLine:   143,
									StartPos:  3104,
									EndPos:    3106,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 143,
											EndLine:   143,
											StartPos:  3104,
											EndPos:    3106,
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
								StartPos:  3108,
								EndPos:    3109,
							},
						},
					},
				},
			},
			&stmt.Class{
				Position: &position.Position{
					StartLine: 144,
					EndLine:   144,
					StartPos:  3115,
					EndPos:    3159,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 144,
						EndLine:   144,
						StartPos:  3121,
						EndPos:    3123,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 144,
							EndLine:   144,
							StartPos:  3127,
							EndPos:    3157,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 144,
									EndLine:   144,
									StartPos:  3131,
									EndPos:    3133,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 144,
											EndLine:   144,
											StartPos:  3131,
											EndPos:    3133,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 144,
									EndLine:   144,
									StartPos:  3136,
									EndPos:    3138,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 144,
											EndLine:   144,
											StartPos:  3136,
											EndPos:    3138,
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
								StartPos:  3140,
								EndPos:    3157,
							},
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 144,
										EndLine:   144,
										StartPos:  3142,
										EndPos:    3154,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 144,
											EndLine:   144,
											StartPos:  3142,
											EndPos:    3144,
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  3142,
												EndPos:    3144,
											},
											Value: "one",
										},
									},
									Modifier: &node.Identifier{
										Position: &position.Position{
											StartLine: 144,
											EndLine:   144,
											StartPos:  3149,
											EndPos:    3154,
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
					StartLine: 145,
					EndLine:   145,
					StartPos:  3163,
					EndPos:    3211,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 145,
						EndLine:   145,
						StartPos:  3169,
						EndPos:    3171,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 145,
							EndLine:   145,
							StartPos:  3175,
							EndPos:    3209,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 145,
									EndLine:   145,
									StartPos:  3179,
									EndPos:    3181,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 145,
											EndLine:   145,
											StartPos:  3179,
											EndPos:    3181,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 145,
									EndLine:   145,
									StartPos:  3184,
									EndPos:    3186,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 145,
											EndLine:   145,
											StartPos:  3184,
											EndPos:    3186,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 145,
								EndLine:   145,
								StartPos:  3188,
								EndPos:    3209,
							},
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 145,
										EndLine:   145,
										StartPos:  3190,
										EndPos:    3206,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 145,
											EndLine:   145,
											StartPos:  3190,
											EndPos:    3192,
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 145,
												EndLine:   145,
												StartPos:  3190,
												EndPos:    3192,
											},
											Value: "one",
										},
									},
									Modifier: &node.Identifier{
										Position: &position.Position{
											StartLine: 145,
											EndLine:   145,
											StartPos:  3197,
											EndPos:    3202,
										},
										Value: "public",
									},
									Alias: &node.Identifier{
										Position: &position.Position{
											StartLine: 145,
											EndLine:   145,
											StartPos:  3204,
											EndPos:    3206,
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
					StartLine: 146,
					EndLine:   146,
					StartPos:  3215,
					EndPos:    3291,
				},
				PhpDocComment: "",
				ClassName: &node.Identifier{
					Position: &position.Position{
						StartLine: 146,
						EndLine:   146,
						StartPos:  3221,
						EndPos:    3223,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Position: &position.Position{
							StartLine: 146,
							EndLine:   146,
							StartPos:  3227,
							EndPos:    3289,
						},
						Traits: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 146,
									EndLine:   146,
									StartPos:  3231,
									EndPos:    3233,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 146,
											EndLine:   146,
											StartPos:  3231,
											EndPos:    3233,
										},
										Value: "Bar",
									},
								},
							},
							&name.Name{
								Position: &position.Position{
									StartLine: 146,
									EndLine:   146,
									StartPos:  3236,
									EndPos:    3238,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 146,
											EndLine:   146,
											StartPos:  3236,
											EndPos:    3238,
										},
										Value: "Baz",
									},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Position: &position.Position{
								StartLine: 146,
								EndLine:   146,
								StartPos:  3240,
								EndPos:    3289,
							},
							Adaptations: []node.Node{
								&stmt.TraitUsePrecedence{
									Position: &position.Position{
										StartLine: 146,
										EndLine:   146,
										StartPos:  3242,
										EndPos:    3269,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 146,
											EndLine:   146,
											StartPos:  3242,
											EndPos:    3249,
										},
										Trait: &name.Name{
											Position: &position.Position{
												StartLine: 146,
												EndLine:   146,
												StartPos:  3242,
												EndPos:    3244,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 146,
														EndLine:   146,
														StartPos:  3242,
														EndPos:    3244,
													},
													Value: "Bar",
												},
											},
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 146,
												EndLine:   146,
												StartPos:  3247,
												EndPos:    3249,
											},
											Value: "one",
										},
									},
									Insteadof: []node.Node{
										&name.Name{
											Position: &position.Position{
												StartLine: 146,
												EndLine:   146,
												StartPos:  3261,
												EndPos:    3263,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 146,
														EndLine:   146,
														StartPos:  3261,
														EndPos:    3263,
													},
													Value: "Baz",
												},
											},
										},
										&name.Name{
											Position: &position.Position{
												StartLine: 146,
												EndLine:   146,
												StartPos:  3266,
												EndPos:    3269,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 146,
														EndLine:   146,
														StartPos:  3266,
														EndPos:    3269,
													},
													Value: "Quux",
												},
											},
										},
									},
								},
								&stmt.TraitUseAlias{
									Position: &position.Position{
										StartLine: 146,
										EndLine:   146,
										StartPos:  3272,
										EndPos:    3286,
									},
									Ref: &stmt.TraitMethodRef{
										Position: &position.Position{
											StartLine: 146,
											EndLine:   146,
											StartPos:  3272,
											EndPos:    3279,
										},
										Trait: &name.Name{
											Position: &position.Position{
												StartLine: 146,
												EndLine:   146,
												StartPos:  3272,
												EndPos:    3274,
											},
											Parts: []node.Node{
												&name.NamePart{
													Position: &position.Position{
														StartLine: 146,
														EndLine:   146,
														StartPos:  3272,
														EndPos:    3274,
													},
													Value: "Baz",
												},
											},
										},
										Method: &node.Identifier{
											Position: &position.Position{
												StartLine: 146,
												EndLine:   146,
												StartPos:  3277,
												EndPos:    3279,
											},
											Value: "one",
										},
									},
									Alias: &node.Identifier{
										Position: &position.Position{
											StartLine: 146,
											EndLine:   146,
											StartPos:  3284,
											EndPos:    3286,
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
					StartLine: 148,
					EndLine:   -1,
					StartPos:  3296,
					EndPos:    -1,
				},
				Stmts:   []node.Node{},
				Catches: []node.Node{},
			},
			&stmt.Try{
				Position: &position.Position{
					StartLine: 149,
					EndLine:   149,
					StartPos:  3305,
					EndPos:    3334,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 149,
							EndLine:   149,
							StartPos:  3312,
							EndPos:    3334,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3319,
									EndPos:    3327,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 149,
											EndLine:   149,
											StartPos:  3319,
											EndPos:    3327,
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
								StartPos:  3329,
								EndPos:    3330,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3329,
									EndPos:    3330,
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
					StartPos:  3338,
					EndPos:    3398,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 150,
							EndLine:   150,
							StartPos:  3345,
							EndPos:    3367,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 150,
									EndLine:   150,
									StartPos:  3352,
									EndPos:    3360,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 150,
											EndLine:   150,
											StartPos:  3352,
											EndPos:    3360,
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
								StartPos:  3362,
								EndPos:    3363,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 150,
									EndLine:   150,
									StartPos:  3362,
									EndPos:    3363,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 150,
							EndLine:   150,
							StartPos:  3369,
							EndPos:    3398,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 150,
									EndLine:   150,
									StartPos:  3376,
									EndPos:    3391,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 150,
											EndLine:   150,
											StartPos:  3376,
											EndPos:    3391,
										},
										Value: "RuntimeException",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 150,
								EndLine:   150,
								StartPos:  3393,
								EndPos:    3394,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 150,
									EndLine:   150,
									StartPos:  3393,
									EndPos:    3394,
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
					StartLine: 151,
					EndLine:   151,
					StartPos:  3402,
					EndPos:    3505,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 151,
							EndLine:   151,
							StartPos:  3409,
							EndPos:    3431,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 151,
									EndLine:   151,
									StartPos:  3416,
									EndPos:    3424,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3416,
											EndPos:    3424,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 151,
								EndLine:   151,
								StartPos:  3426,
								EndPos:    3427,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 151,
									EndLine:   151,
									StartPos:  3426,
									EndPos:    3427,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 151,
							EndLine:   151,
							StartPos:  3433,
							EndPos:    3463,
						},
						Types: []node.Node{
							&name.FullyQualified{
								Position: &position.Position{
									StartLine: 151,
									EndLine:   151,
									StartPos:  3440,
									EndPos:    3456,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3441,
											EndPos:    3456,
										},
										Value: "RuntimeException",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 151,
								EndLine:   151,
								StartPos:  3458,
								EndPos:    3459,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 151,
									EndLine:   151,
									StartPos:  3458,
									EndPos:    3459,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 151,
							EndLine:   151,
							StartPos:  3465,
							EndPos:    3505,
						},
						Types: []node.Node{
							&name.Relative{
								Position: &position.Position{
									StartLine: 151,
									EndLine:   151,
									StartPos:  3472,
									EndPos:    3498,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3482,
											EndPos:    3498,
										},
										Value: "AdditionException",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 151,
								EndLine:   151,
								StartPos:  3500,
								EndPos:    3501,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 151,
									EndLine:   151,
									StartPos:  3500,
									EndPos:    3501,
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
					StartLine: 152,
					EndLine:   152,
					StartPos:  3509,
					EndPos:    3549,
				},
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Position: &position.Position{
							StartLine: 152,
							EndLine:   152,
							StartPos:  3516,
							EndPos:    3538,
						},
						Types: []node.Node{
							&name.Name{
								Position: &position.Position{
									StartLine: 152,
									EndLine:   152,
									StartPos:  3523,
									EndPos:    3531,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 152,
											EndLine:   152,
											StartPos:  3523,
											EndPos:    3531,
										},
										Value: "Exception",
									},
								},
							},
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 152,
								EndLine:   152,
								StartPos:  3533,
								EndPos:    3534,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 152,
									EndLine:   152,
									StartPos:  3533,
									EndPos:    3534,
								},
								Value: "e",
							},
						},
						Stmts: []node.Node{},
					},
				},
				Finally: &stmt.Finally{
					Position: &position.Position{
						StartLine: 152,
						EndLine:   152,
						StartPos:  3540,
						EndPos:    3549,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Unset{
				Position: &position.Position{
					StartLine: 154,
					EndLine:   154,
					StartPos:  3554,
					EndPos:    3567,
				},
				Vars: []node.Node{
					&expr.Variable{
						Position: &position.Position{
							StartLine: 154,
							EndLine:   154,
							StartPos:  3560,
							EndPos:    3561,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 154,
								EndLine:   154,
								StartPos:  3560,
								EndPos:    3561,
							},
							Value: "a",
						},
					},
					&expr.Variable{
						Position: &position.Position{
							StartLine: 154,
							EndLine:   154,
							StartPos:  3564,
							EndPos:    3565,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 154,
								EndLine:   154,
								StartPos:  3564,
								EndPos:    3565,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 156,
					EndLine:   156,
					StartPos:  3572,
					EndPos:    3579,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 156,
							EndLine:   156,
							StartPos:  3576,
							EndPos:    3578,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 156,
								EndLine:   156,
								StartPos:  3576,
								EndPos:    3578,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 156,
										EndLine:   156,
										StartPos:  3576,
										EndPos:    3578,
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
					StartLine: 157,
					EndLine:   157,
					StartPos:  3583,
					EndPos:    3591,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 157,
							EndLine:   157,
							StartPos:  3588,
							EndPos:    3590,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 157,
								EndLine:   157,
								StartPos:  3588,
								EndPos:    3590,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 157,
										EndLine:   157,
										StartPos:  3588,
										EndPos:    3590,
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
					StartLine: 158,
					EndLine:   158,
					StartPos:  3595,
					EndPos:    3610,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 158,
							EndLine:   158,
							StartPos:  3600,
							EndPos:    3609,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 158,
								EndLine:   158,
								StartPos:  3600,
								EndPos:    3602,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 158,
										EndLine:   158,
										StartPos:  3600,
										EndPos:    3602,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 158,
								EndLine:   158,
								StartPos:  3607,
								EndPos:    3609,
							},
							Value: "Bar",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 159,
					EndLine:   159,
					StartPos:  3614,
					EndPos:    3626,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 159,
							EndLine:   159,
							StartPos:  3618,
							EndPos:    3620,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 159,
								EndLine:   159,
								StartPos:  3618,
								EndPos:    3620,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 159,
										EndLine:   159,
										StartPos:  3618,
										EndPos:    3620,
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
					StartPos:  3630,
					EndPos:    3649,
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 160,
							EndLine:   160,
							StartPos:  3634,
							EndPos:    3636,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 160,
								EndLine:   160,
								StartPos:  3634,
								EndPos:    3636,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 160,
										EndLine:   160,
										StartPos:  3634,
										EndPos:    3636,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 160,
							EndLine:   160,
							StartPos:  3639,
							EndPos:    3648,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 160,
								EndLine:   160,
								StartPos:  3639,
								EndPos:    3641,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 160,
										EndLine:   160,
										StartPos:  3639,
										EndPos:    3641,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 160,
								EndLine:   160,
								StartPos:  3646,
								EndPos:    3648,
							},
							Value: "Baz",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 161,
					EndLine:   161,
					StartPos:  3653,
					EndPos:    3675,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 161,
						EndLine:   161,
						StartPos:  3657,
						EndPos:    3664,
					},
					Value: "function",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 161,
							EndLine:   161,
							StartPos:  3666,
							EndPos:    3668,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 161,
								EndLine:   161,
								StartPos:  3666,
								EndPos:    3668,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 161,
										EndLine:   161,
										StartPos:  3666,
										EndPos:    3668,
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
							StartPos:  3672,
							EndPos:    3674,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 161,
								EndLine:   161,
								StartPos:  3672,
								EndPos:    3674,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 161,
										EndLine:   161,
										StartPos:  3672,
										EndPos:    3674,
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
					StartPos:  3679,
					EndPos:    3715,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 162,
						EndLine:   162,
						StartPos:  3683,
						EndPos:    3690,
					},
					Value: "function",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 162,
							EndLine:   162,
							StartPos:  3692,
							EndPos:    3701,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3692,
								EndPos:    3694,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 162,
										EndLine:   162,
										StartPos:  3692,
										EndPos:    3694,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3699,
								EndPos:    3701,
							},
							Value: "foo",
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 162,
							EndLine:   162,
							StartPos:  3705,
							EndPos:    3714,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3705,
								EndPos:    3707,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 162,
										EndLine:   162,
										StartPos:  3705,
										EndPos:    3707,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3712,
								EndPos:    3714,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.UseList{
				Position: &position.Position{
					StartLine: 163,
					EndLine:   163,
					StartPos:  3719,
					EndPos:    3738,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 163,
						EndLine:   163,
						StartPos:  3723,
						EndPos:    3727,
					},
					Value: "const",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 163,
							EndLine:   163,
							StartPos:  3729,
							EndPos:    3731,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 163,
								EndLine:   163,
								StartPos:  3729,
								EndPos:    3731,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 163,
										EndLine:   163,
										StartPos:  3729,
										EndPos:    3731,
									},
									Value: "Foo",
								},
							},
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 163,
							EndLine:   163,
							StartPos:  3735,
							EndPos:    3737,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 163,
								EndLine:   163,
								StartPos:  3735,
								EndPos:    3737,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 163,
										EndLine:   163,
										StartPos:  3735,
										EndPos:    3737,
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
					StartLine: 164,
					EndLine:   164,
					StartPos:  3742,
					EndPos:    3775,
				},
				UseType: &node.Identifier{
					Position: &position.Position{
						StartLine: 164,
						EndLine:   164,
						StartPos:  3746,
						EndPos:    3750,
					},
					Value: "const",
				},
				Uses: []node.Node{
					&stmt.Use{
						Position: &position.Position{
							StartLine: 164,
							EndLine:   164,
							StartPos:  3752,
							EndPos:    3761,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 164,
								EndLine:   164,
								StartPos:  3752,
								EndPos:    3754,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 164,
										EndLine:   164,
										StartPos:  3752,
										EndPos:    3754,
									},
									Value: "Foo",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 164,
								EndLine:   164,
								StartPos:  3759,
								EndPos:    3761,
							},
							Value: "foo",
						},
					},
					&stmt.Use{
						Position: &position.Position{
							StartLine: 164,
							EndLine:   164,
							StartPos:  3765,
							EndPos:    3774,
						},
						Use: &name.Name{
							Position: &position.Position{
								StartLine: 164,
								EndLine:   164,
								StartPos:  3765,
								EndPos:    3767,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 164,
										EndLine:   164,
										StartPos:  3765,
										EndPos:    3767,
									},
									Value: "Bar",
								},
							},
						},
						Alias: &node.Identifier{
							Position: &position.Position{
								StartLine: 164,
								EndLine:   164,
								StartPos:  3772,
								EndPos:    3774,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 166,
					EndLine:   166,
					StartPos:  3780,
					EndPos:    3785,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 166,
						EndLine:   166,
						StartPos:  3780,
						EndPos:    3784,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 166,
							EndLine:   166,
							StartPos:  3780,
							EndPos:    3781,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 166,
								EndLine:   166,
								StartPos:  3780,
								EndPos:    3781,
							},
							Value: "a",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 166,
							EndLine:   166,
							StartPos:  3783,
							EndPos:    3783,
						},
						Value: "1",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 167,
					EndLine:   167,
					StartPos:  3789,
					EndPos:    3797,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 167,
						EndLine:   167,
						StartPos:  3789,
						EndPos:    3796,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 167,
							EndLine:   167,
							StartPos:  3789,
							EndPos:    3793,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 167,
								EndLine:   167,
								StartPos:  3789,
								EndPos:    3790,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 167,
									EndLine:   167,
									StartPos:  3789,
									EndPos:    3790,
								},
								Value: "a",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 167,
								EndLine:   167,
								StartPos:  3792,
								EndPos:    3792,
							},
							Value: "1",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 167,
							EndLine:   167,
							StartPos:  3795,
							EndPos:    3795,
						},
						Value: "2",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 168,
					EndLine:   168,
					StartPos:  3801,
					EndPos:    3808,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 168,
						EndLine:   168,
						StartPos:  3801,
						EndPos:    3807,
					},
					Items: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 169,
					EndLine:   169,
					StartPos:  3812,
					EndPos:    3820,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 169,
						EndLine:   169,
						StartPos:  3812,
						EndPos:    3819,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 169,
								EndLine:   169,
								StartPos:  3818,
								EndPos:    3818,
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 169,
									EndLine:   169,
									StartPos:  3818,
									EndPos:    3818,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 170,
					EndLine:   170,
					StartPos:  3824,
					EndPos:    3841,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 170,
						EndLine:   170,
						StartPos:  3824,
						EndPos:    3840,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 170,
								EndLine:   170,
								StartPos:  3830,
								EndPos:    3833,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3830,
									EndPos:    3830,
								},
								Value: "1",
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3833,
									EndPos:    3833,
								},
								Value: "1",
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 170,
								EndLine:   170,
								StartPos:  3836,
								EndPos:    3838,
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3836,
									EndPos:    3838,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3837,
										EndPos:    3838,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 170,
											EndLine:   170,
											StartPos:  3837,
											EndPos:    3838,
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
					StartLine: 171,
					EndLine:   171,
					StartPos:  3845,
					EndPos:    3859,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 171,
						EndLine:   171,
						StartPos:  3845,
						EndPos:    3858,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 171,
								EndLine:   171,
								StartPos:  3851,
								EndPos:    3857,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 171,
									EndLine:   171,
									StartPos:  3851,
									EndPos:    3851,
								},
								Value: "3",
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 171,
									EndLine:   171,
									StartPos:  3855,
									EndPos:    3857,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 171,
										EndLine:   171,
										StartPos:  3856,
										EndPos:    3857,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 171,
											EndLine:   171,
											StartPos:  3856,
											EndPos:    3857,
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
					StartLine: 172,
					EndLine:   172,
					StartPos:  3863,
					EndPos:    3891,
				},
				Expr: &expr.Array{
					Position: &position.Position{
						StartLine: 172,
						EndLine:   172,
						StartPos:  3863,
						EndPos:    3890,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3869,
								EndPos:    3871,
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 172,
									EndLine:   172,
									StartPos:  3869,
									EndPos:    3871,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 172,
										EndLine:   172,
										StartPos:  3870,
										EndPos:    3871,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 172,
											EndLine:   172,
											StartPos:  3870,
											EndPos:    3871,
										},
										Value: "b",
									},
								},
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3874,
								EndPos:    3877,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 172,
									EndLine:   172,
									StartPos:  3874,
									EndPos:    3874,
								},
								Value: "1",
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 172,
									EndLine:   172,
									StartPos:  3877,
									EndPos:    3877,
								},
								Value: "1",
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3880,
								EndPos:    3880,
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 172,
									EndLine:   172,
									StartPos:  3880,
									EndPos:    3880,
								},
								Value: "1",
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3883,
								EndPos:    3889,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 172,
									EndLine:   172,
									StartPos:  3883,
									EndPos:    3883,
								},
								Value: "3",
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 172,
									EndLine:   172,
									StartPos:  3887,
									EndPos:    3889,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 172,
										EndLine:   172,
										StartPos:  3888,
										EndPos:    3889,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 172,
											EndLine:   172,
											StartPos:  3888,
											EndPos:    3889,
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
					StartLine: 173,
					EndLine:   173,
					StartPos:  3895,
					EndPos:    3898,
				},
				Expr: &expr.BitwiseNot{
					Position: &position.Position{
						StartLine: 173,
						EndLine:   173,
						StartPos:  3895,
						EndPos:    3897,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 173,
							EndLine:   173,
							StartPos:  3896,
							EndPos:    3897,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 173,
								EndLine:   173,
								StartPos:  3896,
								EndPos:    3897,
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
					StartPos:  3902,
					EndPos:    3905,
				},
				Expr: &expr.BooleanNot{
					Position: &position.Position{
						StartLine: 174,
						EndLine:   174,
						StartPos:  3902,
						EndPos:    3904,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 174,
							EndLine:   174,
							StartPos:  3903,
							EndPos:    3904,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3903,
								EndPos:    3904,
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
					StartPos:  3910,
					EndPos:    3918,
				},
				Expr: &expr.ClassConstFetch{
					Position: &position.Position{
						StartLine: 176,
						EndLine:   176,
						StartPos:  3910,
						EndPos:    3917,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 176,
							EndLine:   176,
							StartPos:  3910,
							EndPos:    3912,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 176,
									EndLine:   176,
									StartPos:  3910,
									EndPos:    3912,
								},
								Value: "Foo",
							},
						},
					},
					ConstantName: &node.Identifier{
						Position: &position.Position{
							StartLine: 176,
							EndLine:   176,
							StartPos:  3915,
							EndPos:    3917,
						},
						Value: "Bar",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 177,
					EndLine:   177,
					StartPos:  3922,
					EndPos:    3931,
				},
				Expr: &expr.Clone{
					Position: &position.Position{
						StartLine: 177,
						EndLine:   177,
						StartPos:  3922,
						EndPos:    3929,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 177,
							EndLine:   177,
							StartPos:  3928,
							EndPos:    3929,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 177,
								EndLine:   177,
								StartPos:  3928,
								EndPos:    3929,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 178,
					EndLine:   178,
					StartPos:  3935,
					EndPos:    3943,
				},
				Expr: &expr.Clone{
					Position: &position.Position{
						StartLine: 178,
						EndLine:   178,
						StartPos:  3935,
						EndPos:    3942,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 178,
							EndLine:   178,
							StartPos:  3941,
							EndPos:    3942,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 178,
								EndLine:   178,
								StartPos:  3941,
								EndPos:    3942,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 179,
					EndLine:   179,
					StartPos:  3947,
					EndPos:    3959,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 179,
						EndLine:   179,
						StartPos:  3947,
						EndPos:    3958,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Stmts:         []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 180,
					EndLine:   180,
					StartPos:  3963,
					EndPos:    3996,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 180,
						EndLine:   180,
						StartPos:  3963,
						EndPos:    3995,
					},
					Static:        false,
					PhpDocComment: "",
					ReturnsRef:    false,
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 180,
								EndLine:   180,
								StartPos:  3972,
								EndPos:    3973,
							},
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 180,
									EndLine:   180,
									StartPos:  3972,
									EndPos:    3973,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 180,
										EndLine:   180,
										StartPos:  3972,
										EndPos:    3973,
									},
									Value: "a",
								},
							},
						},
						&node.Parameter{
							Position: &position.Position{
								StartLine: 180,
								EndLine:   180,
								StartPos:  3976,
								EndPos:    3977,
							},
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 180,
									EndLine:   180,
									StartPos:  3976,
									EndPos:    3977,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 180,
										EndLine:   180,
										StartPos:  3976,
										EndPos:    3977,
									},
									Value: "b",
								},
							},
						},
					},
					ClosureUse: &expr.ClosureUse{
						Position: &position.Position{
							StartLine: 180,
							EndLine:   180,
							StartPos:  3980,
							EndPos:    3992,
						},
						Uses: []node.Node{
							&expr.Variable{
								Position: &position.Position{
									StartLine: 180,
									EndLine:   180,
									StartPos:  3985,
									EndPos:    3986,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 180,
										EndLine:   180,
										StartPos:  3985,
										EndPos:    3986,
									},
									Value: "c",
								},
							},
							&expr.Reference{
								Position: &position.Position{
									StartLine: 180,
									EndLine:   180,
									StartPos:  3989,
									EndPos:    3991,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 180,
										EndLine:   180,
										StartPos:  3990,
										EndPos:    3991,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 180,
											EndLine:   180,
											StartPos:  3990,
											EndPos:    3991,
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
					StartLine: 181,
					EndLine:   181,
					StartPos:  4000,
					EndPos:    4033,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 181,
						EndLine:   181,
						StartPos:  4000,
						EndPos:    4032,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Params: []node.Node{
						&node.Parameter{
							Position: &position.Position{
								StartLine: 181,
								EndLine:   181,
								StartPos:  4009,
								EndPos:    4010,
							},
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 181,
									EndLine:   181,
									StartPos:  4009,
									EndPos:    4010,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 181,
										EndLine:   181,
										StartPos:  4009,
										EndPos:    4010,
									},
									Value: "a",
								},
							},
						},
						&node.Parameter{
							Position: &position.Position{
								StartLine: 181,
								EndLine:   181,
								StartPos:  4013,
								EndPos:    4014,
							},
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 181,
									EndLine:   181,
									StartPos:  4013,
									EndPos:    4014,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 181,
										EndLine:   181,
										StartPos:  4013,
										EndPos:    4014,
									},
									Value: "b",
								},
							},
						},
					},
					ClosureUse: &expr.ClosureUse{
						Position: &position.Position{
							StartLine: 181,
							EndLine:   181,
							StartPos:  4017,
							EndPos:    4029,
						},
						Uses: []node.Node{
							&expr.Reference{
								Position: &position.Position{
									StartLine: 181,
									EndLine:   181,
									StartPos:  4022,
									EndPos:    4024,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 181,
										EndLine:   181,
										StartPos:  4023,
										EndPos:    4024,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 181,
											EndLine:   181,
											StartPos:  4023,
											EndPos:    4024,
										},
										Value: "c",
									},
								},
							},
							&expr.Variable{
								Position: &position.Position{
									StartLine: 181,
									EndLine:   181,
									StartPos:  4027,
									EndPos:    4028,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 181,
										EndLine:   181,
										StartPos:  4027,
										EndPos:    4028,
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
					StartLine: 182,
					EndLine:   182,
					StartPos:  4037,
					EndPos:    4050,
				},
				Expr: &expr.Closure{
					Position: &position.Position{
						StartLine: 182,
						EndLine:   182,
						StartPos:  4037,
						EndPos:    4049,
					},
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Stmts:         []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 183,
					EndLine:   183,
					StartPos:  4054,
					EndPos:    4057,
				},
				Expr: &expr.ConstFetch{
					Position: &position.Position{
						StartLine: 183,
						EndLine:   183,
						StartPos:  4054,
						EndPos:    4056,
					},
					Constant: &name.Name{
						Position: &position.Position{
							StartLine: 183,
							EndLine:   183,
							StartPos:  4054,
							EndPos:    4056,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 183,
									EndLine:   183,
									StartPos:  4054,
									EndPos:    4056,
								},
								Value: "foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 184,
					EndLine:   184,
					StartPos:  4061,
					EndPos:    4074,
				},
				Expr: &expr.ConstFetch{
					Position: &position.Position{
						StartLine: 184,
						EndLine:   184,
						StartPos:  4061,
						EndPos:    4073,
					},
					Constant: &name.Relative{
						Position: &position.Position{
							StartLine: 184,
							EndLine:   184,
							StartPos:  4061,
							EndPos:    4073,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 184,
									EndLine:   184,
									StartPos:  4071,
									EndPos:    4073,
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
					StartPos:  4078,
					EndPos:    4082,
				},
				Expr: &expr.ConstFetch{
					Position: &position.Position{
						StartLine: 185,
						EndLine:   185,
						StartPos:  4078,
						EndPos:    4081,
					},
					Constant: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 185,
							EndLine:   185,
							StartPos:  4078,
							EndPos:    4081,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 185,
									EndLine:   185,
									StartPos:  4079,
									EndPos:    4081,
								},
								Value: "foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 187,
					EndLine:   187,
					StartPos:  4087,
					EndPos:    4096,
				},
				Expr: &expr.Empty{
					Position: &position.Position{
						StartLine: 187,
						EndLine:   187,
						StartPos:  4087,
						EndPos:    4095,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 187,
							EndLine:   187,
							StartPos:  4093,
							EndPos:    4094,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 187,
								EndLine:   187,
								StartPos:  4093,
								EndPos:    4094,
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
					StartPos:  4100,
					EndPos:    4110,
				},
				Expr: &expr.Empty{
					Position: &position.Position{
						StartLine: 188,
						EndLine:   188,
						StartPos:  4100,
						EndPos:    4109,
					},
					Expr: &expr.ConstFetch{
						Position: &position.Position{
							StartLine: 188,
							EndLine:   188,
							StartPos:  4106,
							EndPos:    4108,
						},
						Constant: &name.Name{
							Position: &position.Position{
								StartLine: 188,
								EndLine:   188,
								StartPos:  4106,
								EndPos:    4108,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 188,
										EndLine:   188,
										StartPos:  4106,
										EndPos:    4108,
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
					StartLine: 189,
					EndLine:   189,
					StartPos:  4114,
					EndPos:    4117,
				},
				Expr: &expr.ErrorSuppress{
					Position: &position.Position{
						StartLine: 189,
						EndLine:   189,
						StartPos:  4114,
						EndPos:    4116,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 189,
							EndLine:   189,
							StartPos:  4115,
							EndPos:    4116,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 189,
								EndLine:   189,
								StartPos:  4115,
								EndPos:    4116,
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
					StartPos:  4121,
					EndPos:    4129,
				},
				Expr: &expr.Eval{
					Position: &position.Position{
						StartLine: 190,
						EndLine:   190,
						StartPos:  4121,
						EndPos:    4128,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 190,
							EndLine:   190,
							StartPos:  4126,
							EndPos:    4127,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 190,
								EndLine:   190,
								StartPos:  4126,
								EndPos:    4127,
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
					StartPos:  4133,
					EndPos:    4137,
				},
				Expr: &expr.Exit{
					Die: false,
					Position: &position.Position{
						StartLine: 191,
						EndLine:   191,
						StartPos:  4133,
						EndPos:    4136,
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 192,
					EndLine:   192,
					StartPos:  4141,
					EndPos:    4149,
				},
				Expr: &expr.Exit{
					Die: false,
					Position: &position.Position{
						StartLine: 192,
						EndLine:   192,
						StartPos:  4141,
						EndPos:    4148,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 192,
							EndLine:   192,
							StartPos:  4146,
							EndPos:    4147,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 192,
								EndLine:   192,
								StartPos:  4146,
								EndPos:    4147,
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
					StartPos:  4153,
					EndPos:    4158,
				},
				Expr: &expr.Exit{
					Die: true,
					Position: &position.Position{
						StartLine: 193,
						EndLine:   193,
						StartPos:  4153,
						EndPos:    4157,
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 194,
					EndLine:   194,
					StartPos:  4162,
					EndPos:    4169,
				},
				Expr: &expr.Exit{
					Die: true,
					Position: &position.Position{
						StartLine: 194,
						EndLine:   194,
						StartPos:  4162,
						EndPos:    4168,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 194,
							EndLine:   194,
							StartPos:  4166,
							EndPos:    4167,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 194,
								EndLine:   194,
								StartPos:  4166,
								EndPos:    4167,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 195,
					EndLine:   195,
					StartPos:  4173,
					EndPos:    4178,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 195,
						EndLine:   195,
						StartPos:  4173,
						EndPos:    4177,
					},
					Function: &name.Name{
						Position: &position.Position{
							StartLine: 195,
							EndLine:   195,
							StartPos:  4173,
							EndPos:    4175,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 195,
									EndLine:   195,
									StartPos:  4173,
									EndPos:    4175,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 195,
							EndLine:   195,
							StartPos:  4176,
							EndPos:    4177,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 196,
					EndLine:   196,
					StartPos:  4182,
					EndPos:    4200,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 196,
						EndLine:   196,
						StartPos:  4182,
						EndPos:    4199,
					},
					Function: &name.Relative{
						Position: &position.Position{
							StartLine: 196,
							EndLine:   196,
							StartPos:  4182,
							EndPos:    4194,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 196,
									EndLine:   196,
									StartPos:  4192,
									EndPos:    4194,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 196,
							EndLine:   196,
							StartPos:  4195,
							EndPos:    4199,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 196,
									EndLine:   196,
									StartPos:  4197,
									EndPos:    4198,
								},
								Variadic:    false,
								IsReference: true,
								Expr: &expr.Variable{
									Position: &position.Position{
										StartLine: 196,
										EndLine:   196,
										StartPos:  4197,
										EndPos:    4198,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 196,
											EndLine:   196,
											StartPos:  4197,
											EndPos:    4198,
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
					StartLine: 197,
					EndLine:   197,
					StartPos:  4204,
					EndPos:    4212,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 197,
						EndLine:   197,
						StartPos:  4204,
						EndPos:    4211,
					},
					Function: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 197,
							EndLine:   197,
							StartPos:  4204,
							EndPos:    4207,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 197,
									EndLine:   197,
									StartPos:  4205,
									EndPos:    4207,
								},
								Value: "foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 197,
							EndLine:   197,
							StartPos:  4208,
							EndPos:    4211,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 197,
									EndLine:   197,
									StartPos:  4209,
									EndPos:    4210,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.ShortArray{
									Position: &position.Position{
										StartLine: 197,
										EndLine:   197,
										StartPos:  4209,
										EndPos:    4210,
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
					StartLine: 198,
					EndLine:   198,
					StartPos:  4216,
					EndPos:    4230,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 198,
						EndLine:   198,
						StartPos:  4216,
						EndPos:    4229,
					},
					Function: &expr.Variable{
						Position: &position.Position{
							StartLine: 198,
							EndLine:   198,
							StartPos:  4216,
							EndPos:    4219,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 198,
								EndLine:   198,
								StartPos:  4216,
								EndPos:    4219,
							},
							Value: "foo",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 198,
							EndLine:   198,
							StartPos:  4220,
							EndPos:    4229,
						},
						Arguments: []node.Node{
							&node.Argument{
								Position: &position.Position{
									StartLine: 198,
									EndLine:   198,
									StartPos:  4221,
									EndPos:    4228,
								},
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Yield{
									Position: &position.Position{
										StartLine: 198,
										EndLine:   198,
										StartPos:  4221,
										EndPos:    4228,
									},
									Value: &expr.Variable{
										Position: &position.Position{
											StartLine: 198,
											EndLine:   198,
											StartPos:  4227,
											EndPos:    4228,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 198,
												EndLine:   198,
												StartPos:  4227,
												EndPos:    4228,
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
					StartLine: 200,
					EndLine:   200,
					StartPos:  4235,
					EndPos:    4239,
				},
				Expr: &expr.PostDec{
					Position: &position.Position{
						StartLine: 200,
						EndLine:   200,
						StartPos:  4235,
						EndPos:    4238,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 200,
							EndLine:   200,
							StartPos:  4235,
							EndPos:    4236,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 200,
								EndLine:   200,
								StartPos:  4235,
								EndPos:    4236,
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
					StartPos:  4243,
					EndPos:    4247,
				},
				Expr: &expr.PostInc{
					Position: &position.Position{
						StartLine: 201,
						EndLine:   201,
						StartPos:  4243,
						EndPos:    4246,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 201,
							EndLine:   201,
							StartPos:  4243,
							EndPos:    4244,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 201,
								EndLine:   201,
								StartPos:  4243,
								EndPos:    4244,
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
					StartPos:  4251,
					EndPos:    4255,
				},
				Expr: &expr.PreDec{
					Position: &position.Position{
						StartLine: 202,
						EndLine:   202,
						StartPos:  4251,
						EndPos:    4254,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 202,
							EndLine:   202,
							StartPos:  4253,
							EndPos:    4254,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 202,
								EndLine:   202,
								StartPos:  4253,
								EndPos:    4254,
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
					StartPos:  4259,
					EndPos:    4263,
				},
				Expr: &expr.PreInc{
					Position: &position.Position{
						StartLine: 203,
						EndLine:   203,
						StartPos:  4259,
						EndPos:    4262,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 203,
							EndLine:   203,
							StartPos:  4261,
							EndPos:    4262,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 203,
								EndLine:   203,
								StartPos:  4261,
								EndPos:    4262,
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
					StartPos:  4268,
					EndPos:    4278,
				},
				Expr: &expr.Include{
					Position: &position.Position{
						StartLine: 205,
						EndLine:   205,
						StartPos:  4268,
						EndPos:    4277,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 205,
							EndLine:   205,
							StartPos:  4276,
							EndPos:    4277,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 205,
								EndLine:   205,
								StartPos:  4276,
								EndPos:    4277,
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
					StartPos:  4282,
					EndPos:    4297,
				},
				Expr: &expr.IncludeOnce{
					Position: &position.Position{
						StartLine: 206,
						EndLine:   206,
						StartPos:  4282,
						EndPos:    4296,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 206,
							EndLine:   206,
							StartPos:  4295,
							EndPos:    4296,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 206,
								EndLine:   206,
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
					StartLine: 207,
					EndLine:   207,
					StartPos:  4301,
					EndPos:    4311,
				},
				Expr: &expr.Require{
					Position: &position.Position{
						StartLine: 207,
						EndLine:   207,
						StartPos:  4301,
						EndPos:    4310,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 207,
							EndLine:   207,
							StartPos:  4309,
							EndPos:    4310,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 207,
								EndLine:   207,
								StartPos:  4309,
								EndPos:    4310,
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
					StartPos:  4315,
					EndPos:    4330,
				},
				Expr: &expr.RequireOnce{
					Position: &position.Position{
						StartLine: 208,
						EndLine:   208,
						StartPos:  4315,
						EndPos:    4329,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 208,
							EndLine:   208,
							StartPos:  4328,
							EndPos:    4329,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 208,
								EndLine:   208,
								StartPos:  4328,
								EndPos:    4329,
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
					StartPos:  4335,
					EndPos:    4352,
				},
				Expr: &expr.InstanceOf{
					Position: &position.Position{
						StartLine: 210,
						EndLine:   210,
						StartPos:  4335,
						EndPos:    4351,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 210,
							EndLine:   210,
							StartPos:  4335,
							EndPos:    4336,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 210,
								EndLine:   210,
								StartPos:  4335,
								EndPos:    4336,
							},
							Value: "a",
						},
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 210,
							EndLine:   210,
							StartPos:  4349,
							EndPos:    4351,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 210,
									EndLine:   210,
									StartPos:  4349,
									EndPos:    4351,
								},
								Value: "Foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 211,
					EndLine:   211,
					StartPos:  4356,
					EndPos:    4383,
				},
				Expr: &expr.InstanceOf{
					Position: &position.Position{
						StartLine: 211,
						EndLine:   211,
						StartPos:  4356,
						EndPos:    4382,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 211,
							EndLine:   211,
							StartPos:  4356,
							EndPos:    4357,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 211,
								EndLine:   211,
								StartPos:  4356,
								EndPos:    4357,
							},
							Value: "a",
						},
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 211,
							EndLine:   211,
							StartPos:  4370,
							EndPos:    4382,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 211,
									EndLine:   211,
									StartPos:  4380,
									EndPos:    4382,
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
					StartPos:  4387,
					EndPos:    4405,
				},
				Expr: &expr.InstanceOf{
					Position: &position.Position{
						StartLine: 212,
						EndLine:   212,
						StartPos:  4387,
						EndPos:    4404,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 212,
							EndLine:   212,
							StartPos:  4387,
							EndPos:    4388,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 212,
								EndLine:   212,
								StartPos:  4387,
								EndPos:    4388,
							},
							Value: "a",
						},
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 212,
							EndLine:   212,
							StartPos:  4401,
							EndPos:    4404,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 212,
									EndLine:   212,
									StartPos:  4402,
									EndPos:    4404,
								},
								Value: "Foo",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 214,
					EndLine:   214,
					StartPos:  4410,
					EndPos:    4423,
				},
				Expr: &expr.Isset{
					Position: &position.Position{
						StartLine: 214,
						EndLine:   214,
						StartPos:  4410,
						EndPos:    4422,
					},
					Variables: []node.Node{
						&expr.Variable{
							Position: &position.Position{
								StartLine: 214,
								EndLine:   214,
								StartPos:  4416,
								EndPos:    4417,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 214,
									EndLine:   214,
									StartPos:  4416,
									EndPos:    4417,
								},
								Value: "a",
							},
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 214,
								EndLine:   214,
								StartPos:  4420,
								EndPos:    4421,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 214,
									EndLine:   214,
									StartPos:  4420,
									EndPos:    4421,
								},
								Value: "b",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 215,
					EndLine:   215,
					StartPos:  4427,
					EndPos:    4437,
				},
				Expr: &expr.Isset{
					Position: &position.Position{
						StartLine: 215,
						EndLine:   215,
						StartPos:  4427,
						EndPos:    4436,
					},
					Variables: []node.Node{
						&expr.ConstFetch{
							Position: &position.Position{
								StartLine: 215,
								EndLine:   215,
								StartPos:  4433,
								EndPos:    4435,
							},
							Constant: &name.Name{
								Position: &position.Position{
									StartLine: 215,
									EndLine:   215,
									StartPos:  4433,
									EndPos:    4435,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 215,
											EndLine:   215,
											StartPos:  4433,
											EndPos:    4435,
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
					StartLine: 216,
					EndLine:   216,
					StartPos:  4441,
					EndPos:    4452,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 216,
						EndLine:   216,
						StartPos:  4441,
						EndPos:    4451,
					},
					Variable: &expr.List{
						Position: &position.Position{
							StartLine: 216,
							EndLine:   216,
							StartPos:  4441,
							EndPos:    4446,
						},
						Items: []node.Node{},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 216,
							EndLine:   216,
							StartPos:  4450,
							EndPos:    4451,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 216,
								EndLine:   216,
								StartPos:  4450,
								EndPos:    4451,
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
					StartPos:  4456,
					EndPos:    4473,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 217,
						EndLine:   217,
						StartPos:  4456,
						EndPos:    4472,
					},
					Variable: &expr.List{
						Position: &position.Position{
							StartLine: 217,
							EndLine:   217,
							StartPos:  4456,
							EndPos:    4467,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 217,
									EndLine:   217,
									StartPos:  4461,
									EndPos:    4462,
								},
								Val: &expr.Variable{
									Position: &position.Position{
										StartLine: 217,
										EndLine:   217,
										StartPos:  4461,
										EndPos:    4462,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 217,
											EndLine:   217,
											StartPos:  4461,
											EndPos:    4462,
										},
										Value: "a",
									},
								},
							},
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 217,
									EndLine:   217,
									StartPos:  4465,
									EndPos:    4466,
								},
								Val: &expr.Variable{
									Position: &position.Position{
										StartLine: 217,
										EndLine:   217,
										StartPos:  4465,
										EndPos:    4466,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 217,
											EndLine:   217,
											StartPos:  4465,
											EndPos:    4466,
										},
										Value: "b",
									},
								},
							},
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 217,
							EndLine:   217,
							StartPos:  4471,
							EndPos:    4472,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 217,
								EndLine:   217,
								StartPos:  4471,
								EndPos:    4472,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 218,
					EndLine:   218,
					StartPos:  4477,
					EndPos:    4492,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 218,
						EndLine:   218,
						StartPos:  4477,
						EndPos:    4491,
					},
					Variable: &expr.List{
						Position: &position.Position{
							StartLine: 218,
							EndLine:   218,
							StartPos:  4477,
							EndPos:    4486,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 218,
									EndLine:   218,
									StartPos:  4482,
									EndPos:    4485,
								},
								Val: &expr.ArrayDimFetch{
									Position: &position.Position{
										StartLine: 218,
										EndLine:   218,
										StartPos:  4482,
										EndPos:    4485,
									},
									Variable: &expr.Variable{
										Position: &position.Position{
											StartLine: 218,
											EndLine:   218,
											StartPos:  4482,
											EndPos:    4483,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 218,
												EndLine:   218,
												StartPos:  4482,
												EndPos:    4483,
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
							StartLine: 218,
							EndLine:   218,
							StartPos:  4490,
							EndPos:    4491,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 218,
								EndLine:   218,
								StartPos:  4490,
								EndPos:    4491,
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
					StartPos:  4496,
					EndPos:    4515,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 219,
						EndLine:   219,
						StartPos:  4496,
						EndPos:    4514,
					},
					Variable: &expr.List{
						Position: &position.Position{
							StartLine: 219,
							EndLine:   219,
							StartPos:  4496,
							EndPos:    4509,
						},
						Items: []node.Node{
							&expr.ArrayItem{
								Position: &position.Position{
									StartLine: 219,
									EndLine:   219,
									StartPos:  4501,
									EndPos:    4508,
								},
								Val: &expr.List{
									Position: &position.Position{
										StartLine: 219,
										EndLine:   219,
										StartPos:  4501,
										EndPos:    4508,
									},
									Items: []node.Node{
										&expr.ArrayItem{
											Position: &position.Position{
												StartLine: 219,
												EndLine:   219,
												StartPos:  4506,
												EndPos:    4507,
											},
											Val: &expr.Variable{
												Position: &position.Position{
													StartLine: 219,
													EndLine:   219,
													StartPos:  4506,
													EndPos:    4507,
												},
												VarName: &node.Identifier{
													Position: &position.Position{
														StartLine: 219,
														EndLine:   219,
														StartPos:  4506,
														EndPos:    4507,
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
							StartLine: 219,
							EndLine:   219,
							StartPos:  4513,
							EndPos:    4514,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 219,
								EndLine:   219,
								StartPos:  4513,
								EndPos:    4514,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 221,
					EndLine:   221,
					StartPos:  4520,
					EndPos:    4529,
				},
				Expr: &expr.MethodCall{
					Position: &position.Position{
						StartLine: 221,
						EndLine:   221,
						StartPos:  4520,
						EndPos:    4528,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 221,
							EndLine:   221,
							StartPos:  4520,
							EndPos:    4521,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 221,
								EndLine:   221,
								StartPos:  4520,
								EndPos:    4521,
							},
							Value: "a",
						},
					},
					Method: &node.Identifier{
						Position: &position.Position{
							StartLine: 221,
							EndLine:   221,
							StartPos:  4524,
							EndPos:    4526,
						},
						Value: "foo",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 221,
							EndLine:   221,
							StartPos:  4527,
							EndPos:    4528,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 222,
					EndLine:   222,
					StartPos:  4533,
					EndPos:    4540,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 222,
						EndLine:   222,
						StartPos:  4533,
						EndPos:    4539,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 222,
							EndLine:   222,
							StartPos:  4537,
							EndPos:    4539,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 222,
									EndLine:   222,
									StartPos:  4537,
									EndPos:    4539,
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
					EndPos:    4563,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 223,
						EndLine:   223,
						StartPos:  4544,
						EndPos:    4562,
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 223,
							EndLine:   223,
							StartPos:  4548,
							EndPos:    4560,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 223,
									EndLine:   223,
									StartPos:  4558,
									EndPos:    4560,
								},
								Value: "Foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 223,
							EndLine:   223,
							StartPos:  4561,
							EndPos:    4562,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 224,
					EndLine:   224,
					StartPos:  4567,
					EndPos:    4577,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 224,
						EndLine:   224,
						StartPos:  4567,
						EndPos:    4576,
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 224,
							EndLine:   224,
							StartPos:  4571,
							EndPos:    4574,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 224,
									EndLine:   224,
									StartPos:  4572,
									EndPos:    4574,
								},
								Value: "Foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 224,
							EndLine:   224,
							StartPos:  4575,
							EndPos:    4576,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 225,
					EndLine:   225,
					StartPos:  4581,
					EndPos:    4590,
				},
				Expr: &expr.Print{
					Position: &position.Position{
						StartLine: 225,
						EndLine:   225,
						StartPos:  4581,
						EndPos:    4588,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 225,
							EndLine:   225,
							StartPos:  4587,
							EndPos:    4588,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 225,
								EndLine:   225,
								StartPos:  4587,
								EndPos:    4588,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 226,
					EndLine:   226,
					StartPos:  4594,
					EndPos:    4601,
				},
				Expr: &expr.PropertyFetch{
					Position: &position.Position{
						StartLine: 226,
						EndLine:   226,
						StartPos:  4594,
						EndPos:    4600,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 226,
							EndLine:   226,
							StartPos:  4594,
							EndPos:    4595,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 226,
								EndLine:   226,
								StartPos:  4594,
								EndPos:    4595,
							},
							Value: "a",
						},
					},
					Property: &node.Identifier{
						Position: &position.Position{
							StartLine: 226,
							EndLine:   226,
							StartPos:  4598,
							EndPos:    4600,
						},
						Value: "foo",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 227,
					EndLine:   227,
					StartPos:  4605,
					EndPos:    4615,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 227,
						EndLine:   227,
						StartPos:  4605,
						EndPos:    4613,
					},
					Variable: &expr.PropertyFetch{
						Position: &position.Position{
							StartLine: 227,
							EndLine:   227,
							StartPos:  4605,
							EndPos:    4611,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 227,
								EndLine:   227,
								StartPos:  4605,
								EndPos:    4606,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 227,
									EndLine:   227,
									StartPos:  4605,
									EndPos:    4606,
								},
								Value: "a",
							},
						},
						Property: &node.Identifier{
							Position: &position.Position{
								StartLine: 227,
								EndLine:   227,
								StartPos:  4609,
								EndPos:    4611,
							},
							Value: "foo",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 227,
							EndLine:   227,
							StartPos:  4613,
							EndPos:    4613,
						},
						Value: "1",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 228,
					EndLine:   228,
					StartPos:  4619,
					EndPos:    4647,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 228,
						EndLine:   228,
						StartPos:  4619,
						EndPos:    4645,
					},
					Variable: &expr.PropertyFetch{
						Position: &position.Position{
							StartLine: 228,
							EndLine:   228,
							StartPos:  4619,
							EndPos:    4643,
						},
						Variable: &expr.MethodCall{
							Position: &position.Position{
								StartLine: 228,
								EndLine:   228,
								StartPos:  4619,
								EndPos:    4637,
							},
							Variable: &expr.PropertyFetch{
								Position: &position.Position{
									StartLine: 228,
									EndLine:   228,
									StartPos:  4619,
									EndPos:    4630,
								},
								Variable: &expr.PropertyFetch{
									Position: &position.Position{
										StartLine: 228,
										EndLine:   228,
										StartPos:  4619,
										EndPos:    4625,
									},
									Variable: &expr.Variable{
										Position: &position.Position{
											StartLine: 228,
											EndLine:   228,
											StartPos:  4619,
											EndPos:    4620,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 228,
												EndLine:   228,
												StartPos:  4619,
												EndPos:    4620,
											},
											Value: "a",
										},
									},
									Property: &node.Identifier{
										Position: &position.Position{
											StartLine: 228,
											EndLine:   228,
											StartPos:  4623,
											EndPos:    4625,
										},
										Value: "foo",
									},
								},
								Property: &node.Identifier{
									Position: &position.Position{
										StartLine: 228,
										EndLine:   228,
										StartPos:  4628,
										EndPos:    4630,
									},
									Value: "bar",
								},
							},
							Method: &node.Identifier{
								Position: &position.Position{
									StartLine: 228,
									EndLine:   228,
									StartPos:  4633,
									EndPos:    4635,
								},
								Value: "baz",
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 228,
									EndLine:   228,
									StartPos:  4636,
									EndPos:    4637,
								},
							},
						},
						Property: &node.Identifier{
							Position: &position.Position{
								StartLine: 228,
								EndLine:   228,
								StartPos:  4640,
								EndPos:    4643,
							},
							Value: "quux",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 228,
							EndLine:   228,
							StartPos:  4645,
							EndPos:    4645,
						},
						Value: "0",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 229,
					EndLine:   229,
					StartPos:  4651,
					EndPos:    4666,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 229,
						EndLine:   229,
						StartPos:  4651,
						EndPos:    4664,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 229,
							EndLine:   229,
							StartPos:  4651,
							EndPos:    4661,
						},
						Variable: &expr.MethodCall{
							Position: &position.Position{
								StartLine: 229,
								EndLine:   229,
								StartPos:  4651,
								EndPos:    4659,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 229,
									EndLine:   229,
									StartPos:  4651,
									EndPos:    4652,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 229,
										EndLine:   229,
										StartPos:  4651,
										EndPos:    4652,
									},
									Value: "a",
								},
							},
							Method: &node.Identifier{
								Position: &position.Position{
									StartLine: 229,
									EndLine:   229,
									StartPos:  4655,
									EndPos:    4657,
								},
								Value: "foo",
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 229,
									EndLine:   229,
									StartPos:  4658,
									EndPos:    4659,
								},
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 229,
								EndLine:   229,
								StartPos:  4661,
								EndPos:    4661,
							},
							Value: "1",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 229,
							EndLine:   229,
							StartPos:  4664,
							EndPos:    4664,
						},
						Value: "1",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 230,
					EndLine:   230,
					StartPos:  4670,
					EndPos:    4678,
				},
				Expr: &expr.ShellExec{
					Position: &position.Position{
						StartLine: 230,
						EndLine:   230,
						StartPos:  4670,
						EndPos:    4677,
					},
					Parts: []node.Node{
						&scalar.EncapsedStringPart{
							Position: &position.Position{
								StartLine: 230,
								EndLine:   230,
								StartPos:  4671,
								EndPos:    4674,
							},
							Value: "cmd ",
						},
						&expr.Variable{
							Position: &position.Position{
								StartLine: 230,
								EndLine:   230,
								StartPos:  4675,
								EndPos:    4676,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 230,
									EndLine:   230,
									StartPos:  4675,
									EndPos:    4676,
								},
								Value: "a",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 231,
					EndLine:   231,
					StartPos:  4682,
					EndPos:    4687,
				},
				Expr: &expr.ShellExec{
					Position: &position.Position{
						StartLine: 231,
						EndLine:   231,
						StartPos:  4682,
						EndPos:    4686,
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
					StartLine: 232,
					EndLine:   232,
					StartPos:  4691,
					EndPos:    4693,
				},
				Expr: &expr.ShellExec{
					Position: &position.Position{
						StartLine: 232,
						EndLine:   232,
						StartPos:  4691,
						EndPos:    4692,
					},
					Parts: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 233,
					EndLine:   233,
					StartPos:  4697,
					EndPos:    4699,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 233,
						EndLine:   233,
						StartPos:  4697,
						EndPos:    4698,
					},
					Items: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 234,
					EndLine:   234,
					StartPos:  4703,
					EndPos:    4706,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 234,
						EndLine:   234,
						StartPos:  4703,
						EndPos:    4705,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 234,
								EndLine:   234,
								StartPos:  4704,
								EndPos:    4704,
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 234,
									EndLine:   234,
									StartPos:  4704,
									EndPos:    4704,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 235,
					EndLine:   235,
					StartPos:  4710,
					EndPos:    4722,
				},
				Expr: &expr.ShortArray{
					Position: &position.Position{
						StartLine: 235,
						EndLine:   235,
						StartPos:  4710,
						EndPos:    4721,
					},
					Items: []node.Node{
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 235,
								EndLine:   235,
								StartPos:  4711,
								EndPos:    4714,
							},
							Key: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 235,
									EndLine:   235,
									StartPos:  4711,
									EndPos:    4711,
								},
								Value: "1",
							},
							Val: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 235,
									EndLine:   235,
									StartPos:  4714,
									EndPos:    4714,
								},
								Value: "1",
							},
						},
						&expr.ArrayItem{
							Position: &position.Position{
								StartLine: 235,
								EndLine:   235,
								StartPos:  4717,
								EndPos:    4719,
							},
							Val: &expr.Reference{
								Position: &position.Position{
									StartLine: 235,
									EndLine:   235,
									StartPos:  4717,
									EndPos:    4719,
								},
								Variable: &expr.Variable{
									Position: &position.Position{
										StartLine: 235,
										EndLine:   235,
										StartPos:  4718,
										EndPos:    4719,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 235,
											EndLine:   235,
											StartPos:  4718,
											EndPos:    4719,
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
					StartLine: 237,
					EndLine:   237,
					StartPos:  4727,
					EndPos:    4737,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 237,
						EndLine:   237,
						StartPos:  4727,
						EndPos:    4736,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 237,
							EndLine:   237,
							StartPos:  4727,
							EndPos:    4729,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 237,
									EndLine:   237,
									StartPos:  4727,
									EndPos:    4729,
								},
								Value: "Foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 237,
							EndLine:   237,
							StartPos:  4732,
							EndPos:    4734,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 237,
							EndLine:   237,
							StartPos:  4735,
							EndPos:    4736,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 238,
					EndLine:   238,
					StartPos:  4741,
					EndPos:    4761,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 238,
						EndLine:   238,
						StartPos:  4741,
						EndPos:    4760,
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 238,
							EndLine:   238,
							StartPos:  4741,
							EndPos:    4753,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 238,
									EndLine:   238,
									StartPos:  4751,
									EndPos:    4753,
								},
								Value: "Foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 238,
							EndLine:   238,
							StartPos:  4756,
							EndPos:    4758,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 238,
							EndLine:   238,
							StartPos:  4759,
							EndPos:    4760,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 239,
					EndLine:   239,
					StartPos:  4765,
					EndPos:    4776,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 239,
						EndLine:   239,
						StartPos:  4765,
						EndPos:    4775,
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 239,
							EndLine:   239,
							StartPos:  4765,
							EndPos:    4768,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 239,
									EndLine:   239,
									StartPos:  4766,
									EndPos:    4768,
								},
								Value: "Foo",
							},
						},
					},
					Call: &node.Identifier{
						Position: &position.Position{
							StartLine: 239,
							EndLine:   239,
							StartPos:  4771,
							EndPos:    4773,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 239,
							EndLine:   239,
							StartPos:  4774,
							EndPos:    4775,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 240,
					EndLine:   240,
					StartPos:  4780,
					EndPos:    4791,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 240,
						EndLine:   240,
						StartPos:  4780,
						EndPos:    4790,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 240,
							EndLine:   240,
							StartPos:  4780,
							EndPos:    4782,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 240,
									EndLine:   240,
									StartPos:  4780,
									EndPos:    4782,
								},
								Value: "Foo",
							},
						},
					},
					Call: &expr.Variable{
						Position: &position.Position{
							StartLine: 240,
							EndLine:   240,
							StartPos:  4785,
							EndPos:    4788,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 240,
								EndLine:   240,
								StartPos:  4785,
								EndPos:    4788,
							},
							Value: "bar",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 240,
							EndLine:   240,
							StartPos:  4789,
							EndPos:    4790,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 241,
					EndLine:   241,
					StartPos:  4795,
					EndPos:    4807,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 241,
						EndLine:   241,
						StartPos:  4795,
						EndPos:    4806,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 241,
							EndLine:   241,
							StartPos:  4795,
							EndPos:    4798,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 241,
								EndLine:   241,
								StartPos:  4795,
								EndPos:    4798,
							},
							Value: "foo",
						},
					},
					Call: &expr.Variable{
						Position: &position.Position{
							StartLine: 241,
							EndLine:   241,
							StartPos:  4801,
							EndPos:    4804,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 241,
								EndLine:   241,
								StartPos:  4801,
								EndPos:    4804,
							},
							Value: "bar",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 241,
							EndLine:   241,
							StartPos:  4805,
							EndPos:    4806,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 242,
					EndLine:   242,
					StartPos:  4811,
					EndPos:    4820,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 242,
						EndLine:   242,
						StartPos:  4811,
						EndPos:    4819,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 242,
							EndLine:   242,
							StartPos:  4811,
							EndPos:    4813,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 242,
									EndLine:   242,
									StartPos:  4811,
									EndPos:    4813,
								},
								Value: "Foo",
							},
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 242,
							EndLine:   242,
							StartPos:  4816,
							EndPos:    4819,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 242,
								EndLine:   242,
								StartPos:  4816,
								EndPos:    4819,
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
					StartPos:  4824,
					EndPos:    4843,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 243,
						EndLine:   243,
						StartPos:  4824,
						EndPos:    4842,
					},
					Class: &name.Relative{
						Position: &position.Position{
							StartLine: 243,
							EndLine:   243,
							StartPos:  4824,
							EndPos:    4836,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 243,
									EndLine:   243,
									StartPos:  4834,
									EndPos:    4836,
								},
								Value: "Foo",
							},
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 243,
							EndLine:   243,
							StartPos:  4839,
							EndPos:    4842,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 243,
								EndLine:   243,
								StartPos:  4839,
								EndPos:    4842,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 244,
					EndLine:   244,
					StartPos:  4847,
					EndPos:    4857,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 244,
						EndLine:   244,
						StartPos:  4847,
						EndPos:    4856,
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 244,
							EndLine:   244,
							StartPos:  4847,
							EndPos:    4850,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 244,
									EndLine:   244,
									StartPos:  4848,
									EndPos:    4850,
								},
								Value: "Foo",
							},
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 244,
							EndLine:   244,
							StartPos:  4853,
							EndPos:    4856,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 244,
								EndLine:   244,
								StartPos:  4853,
								EndPos:    4856,
							},
							Value: "bar",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 245,
					EndLine:   245,
					StartPos:  4861,
					EndPos:    4873,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 245,
						EndLine:   245,
						StartPos:  4861,
						EndPos:    4872,
					},
					Condition: &expr.Variable{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4861,
							EndPos:    4862,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4861,
								EndPos:    4862,
							},
							Value: "a",
						},
					},
					IfTrue: &expr.Variable{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4866,
							EndPos:    4867,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4866,
								EndPos:    4867,
							},
							Value: "b",
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4871,
							EndPos:    4872,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4871,
								EndPos:    4872,
							},
							Value: "c",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 246,
					EndLine:   246,
					StartPos:  4877,
					EndPos:    4886,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 246,
						EndLine:   246,
						StartPos:  4877,
						EndPos:    4885,
					},
					Condition: &expr.Variable{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4877,
							EndPos:    4878,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4877,
								EndPos:    4878,
							},
							Value: "a",
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4884,
							EndPos:    4885,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4884,
								EndPos:    4885,
							},
							Value: "c",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 247,
					EndLine:   247,
					StartPos:  4890,
					EndPos:    4912,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 247,
						EndLine:   247,
						StartPos:  4890,
						EndPos:    4911,
					},
					Condition: &expr.Variable{
						Position: &position.Position{
							StartLine: 247,
							EndLine:   247,
							StartPos:  4890,
							EndPos:    4891,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 247,
								EndLine:   247,
								StartPos:  4890,
								EndPos:    4891,
							},
							Value: "a",
						},
					},
					IfTrue: &expr.Ternary{
						Position: &position.Position{
							StartLine: 247,
							EndLine:   247,
							StartPos:  4895,
							EndPos:    4906,
						},
						Condition: &expr.Variable{
							Position: &position.Position{
								StartLine: 247,
								EndLine:   247,
								StartPos:  4895,
								EndPos:    4896,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 247,
									EndLine:   247,
									StartPos:  4895,
									EndPos:    4896,
								},
								Value: "b",
							},
						},
						IfTrue: &expr.Variable{
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
								Value: "c",
							},
						},
						IfFalse: &expr.Variable{
							Position: &position.Position{
								StartLine: 247,
								EndLine:   247,
								StartPos:  4905,
								EndPos:    4906,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 247,
									EndLine:   247,
									StartPos:  4905,
									EndPos:    4906,
								},
								Value: "d",
							},
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 247,
							EndLine:   247,
							StartPos:  4910,
							EndPos:    4911,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 247,
								EndLine:   247,
								StartPos:  4910,
								EndPos:    4911,
							},
							Value: "e",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 248,
					EndLine:   248,
					StartPos:  4916,
					EndPos:    4938,
				},
				Expr: &expr.Ternary{
					Position: &position.Position{
						StartLine: 248,
						EndLine:   248,
						StartPos:  4916,
						EndPos:    4937,
					},
					Condition: &expr.Ternary{
						Position: &position.Position{
							StartLine: 248,
							EndLine:   248,
							StartPos:  4916,
							EndPos:    4927,
						},
						Condition: &expr.Variable{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4916,
								EndPos:    4917,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 248,
									EndLine:   248,
									StartPos:  4916,
									EndPos:    4917,
								},
								Value: "a",
							},
						},
						IfTrue: &expr.Variable{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4921,
								EndPos:    4922,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 248,
									EndLine:   248,
									StartPos:  4921,
									EndPos:    4922,
								},
								Value: "b",
							},
						},
						IfFalse: &expr.Variable{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4926,
								EndPos:    4927,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 248,
									EndLine:   248,
									StartPos:  4926,
									EndPos:    4927,
								},
								Value: "c",
							},
						},
					},
					IfTrue: &expr.Variable{
						Position: &position.Position{
							StartLine: 248,
							EndLine:   248,
							StartPos:  4931,
							EndPos:    4932,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4931,
								EndPos:    4932,
							},
							Value: "d",
						},
					},
					IfFalse: &expr.Variable{
						Position: &position.Position{
							StartLine: 248,
							EndLine:   248,
							StartPos:  4936,
							EndPos:    4937,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4936,
								EndPos:    4937,
							},
							Value: "e",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 249,
					EndLine:   249,
					StartPos:  4942,
					EndPos:    4945,
				},
				Expr: &expr.UnaryMinus{
					Position: &position.Position{
						StartLine: 249,
						EndLine:   249,
						StartPos:  4942,
						EndPos:    4944,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 249,
							EndLine:   249,
							StartPos:  4943,
							EndPos:    4944,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 249,
								EndLine:   249,
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
					StartLine: 250,
					EndLine:   250,
					StartPos:  4949,
					EndPos:    4952,
				},
				Expr: &expr.UnaryPlus{
					Position: &position.Position{
						StartLine: 250,
						EndLine:   250,
						StartPos:  4949,
						EndPos:    4951,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 250,
							EndLine:   250,
							StartPos:  4950,
							EndPos:    4951,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 250,
								EndLine:   250,
								StartPos:  4950,
								EndPos:    4951,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 251,
					EndLine:   251,
					StartPos:  4956,
					EndPos:    4959,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 251,
						EndLine:   251,
						StartPos:  4956,
						EndPos:    4958,
					},
					VarName: &expr.Variable{
						Position: &position.Position{
							StartLine: 251,
							EndLine:   251,
							StartPos:  4957,
							EndPos:    4958,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 251,
								EndLine:   251,
								StartPos:  4957,
								EndPos:    4958,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 252,
					EndLine:   252,
					StartPos:  4963,
					EndPos:    4967,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 252,
						EndLine:   252,
						StartPos:  4963,
						EndPos:    4966,
					},
					VarName: &expr.Variable{
						Position: &position.Position{
							StartLine: 252,
							EndLine:   252,
							StartPos:  4964,
							EndPos:    4966,
						},
						VarName: &expr.Variable{
							Position: &position.Position{
								StartLine: 252,
								EndLine:   252,
								StartPos:  4965,
								EndPos:    4966,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 252,
									EndLine:   252,
									StartPos:  4965,
									EndPos:    4966,
								},
								Value: "a",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 253,
					EndLine:   253,
					StartPos:  4971,
					EndPos:    4976,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 253,
						EndLine:   253,
						StartPos:  4971,
						EndPos:    4975,
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 254,
					EndLine:   254,
					StartPos:  4980,
					EndPos:    4988,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 254,
						EndLine:   254,
						StartPos:  4980,
						EndPos:    4987,
					},
					Value: &expr.Variable{
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
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 255,
					EndLine:   255,
					StartPos:  4992,
					EndPos:    5006,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 255,
						EndLine:   255,
						StartPos:  4992,
						EndPos:    5005,
					},
					Key: &expr.Variable{
						Position: &position.Position{
							StartLine: 255,
							EndLine:   255,
							StartPos:  4998,
							EndPos:    4999,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  4998,
								EndPos:    4999,
							},
							Value: "a",
						},
					},
					Value: &expr.Variable{
						Position: &position.Position{
							StartLine: 255,
							EndLine:   255,
							StartPos:  5004,
							EndPos:    5005,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  5004,
								EndPos:    5005,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 256,
					EndLine:   256,
					StartPos:  5010,
					EndPos:    5026,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 256,
						EndLine:   256,
						StartPos:  5010,
						EndPos:    5025,
					},
					Value: &expr.ClassConstFetch{
						Position: &position.Position{
							StartLine: 256,
							EndLine:   256,
							StartPos:  5016,
							EndPos:    5025,
						},
						Class: &name.Name{
							Position: &position.Position{
								StartLine: 256,
								EndLine:   256,
								StartPos:  5016,
								EndPos:    5018,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 256,
										EndLine:   256,
										StartPos:  5016,
										EndPos:    5018,
									},
									Value: "Foo",
								},
							},
						},
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 256,
								EndLine:   256,
								StartPos:  5021,
								EndPos:    5025,
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
					StartPos:  5030,
					EndPos:    5052,
				},
				Expr: &expr.Yield{
					Position: &position.Position{
						StartLine: 257,
						EndLine:   257,
						StartPos:  5030,
						EndPos:    5051,
					},
					Key: &expr.Variable{
						Position: &position.Position{
							StartLine: 257,
							EndLine:   257,
							StartPos:  5036,
							EndPos:    5037,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 257,
								EndLine:   257,
								StartPos:  5036,
								EndPos:    5037,
							},
							Value: "a",
						},
					},
					Value: &expr.ClassConstFetch{
						Position: &position.Position{
							StartLine: 257,
							EndLine:   257,
							StartPos:  5042,
							EndPos:    5051,
						},
						Class: &name.Name{
							Position: &position.Position{
								StartLine: 257,
								EndLine:   257,
								StartPos:  5042,
								EndPos:    5044,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 257,
										EndLine:   257,
										StartPos:  5042,
										EndPos:    5044,
									},
									Value: "Foo",
								},
							},
						},
						ConstantName: &node.Identifier{
							Position: &position.Position{
								StartLine: 257,
								EndLine:   257,
								StartPos:  5047,
								EndPos:    5051,
							},
							Value: "class",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 259,
					EndLine:   259,
					StartPos:  5059,
					EndPos:    5068,
				},
				Expr: &cast.Array{
					Position: &position.Position{
						StartLine: 259,
						EndLine:   259,
						StartPos:  5059,
						EndPos:    5067,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 259,
							EndLine:   259,
							StartPos:  5066,
							EndPos:    5067,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 259,
								EndLine:   259,
								StartPos:  5066,
								EndPos:    5067,
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
					StartPos:  5072,
					EndPos:    5083,
				},
				Expr: &cast.Bool{
					Position: &position.Position{
						StartLine: 260,
						EndLine:   260,
						StartPos:  5072,
						EndPos:    5082,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 260,
							EndLine:   260,
							StartPos:  5081,
							EndPos:    5082,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 260,
								EndLine:   260,
								StartPos:  5081,
								EndPos:    5082,
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
					StartPos:  5087,
					EndPos:    5095,
				},
				Expr: &cast.Bool{
					Position: &position.Position{
						StartLine: 261,
						EndLine:   261,
						StartPos:  5087,
						EndPos:    5094,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 261,
							EndLine:   261,
							StartPos:  5093,
							EndPos:    5094,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 261,
								EndLine:   261,
								StartPos:  5093,
								EndPos:    5094,
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
					StartPos:  5099,
					EndPos:    5109,
				},
				Expr: &cast.Double{
					Position: &position.Position{
						StartLine: 262,
						EndLine:   262,
						StartPos:  5099,
						EndPos:    5108,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 262,
							EndLine:   262,
							StartPos:  5107,
							EndPos:    5108,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 262,
								EndLine:   262,
								StartPos:  5107,
								EndPos:    5108,
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
					StartPos:  5113,
					EndPos:    5122,
				},
				Expr: &cast.Double{
					Position: &position.Position{
						StartLine: 263,
						EndLine:   263,
						StartPos:  5113,
						EndPos:    5121,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 263,
							EndLine:   263,
							StartPos:  5120,
							EndPos:    5121,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 263,
								EndLine:   263,
								StartPos:  5120,
								EndPos:    5121,
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
					StartPos:  5126,
					EndPos:    5137,
				},
				Expr: &cast.Int{
					Position: &position.Position{
						StartLine: 264,
						EndLine:   264,
						StartPos:  5126,
						EndPos:    5136,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 264,
							EndLine:   264,
							StartPos:  5135,
							EndPos:    5136,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 264,
								EndLine:   264,
								StartPos:  5135,
								EndPos:    5136,
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
					StartPos:  5141,
					EndPos:    5148,
				},
				Expr: &cast.Int{
					Position: &position.Position{
						StartLine: 265,
						EndLine:   265,
						StartPos:  5141,
						EndPos:    5147,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 265,
							EndLine:   265,
							StartPos:  5146,
							EndPos:    5147,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 265,
								EndLine:   265,
								StartPos:  5146,
								EndPos:    5147,
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
					StartPos:  5152,
					EndPos:    5162,
				},
				Expr: &cast.Object{
					Position: &position.Position{
						StartLine: 266,
						EndLine:   266,
						StartPos:  5152,
						EndPos:    5161,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 266,
							EndLine:   266,
							StartPos:  5160,
							EndPos:    5161,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 266,
								EndLine:   266,
								StartPos:  5160,
								EndPos:    5161,
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
					StartPos:  5166,
					EndPos:    5176,
				},
				Expr: &cast.String{
					Position: &position.Position{
						StartLine: 267,
						EndLine:   267,
						StartPos:  5166,
						EndPos:    5175,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 267,
							EndLine:   267,
							StartPos:  5174,
							EndPos:    5175,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 267,
								EndLine:   267,
								StartPos:  5174,
								EndPos:    5175,
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
					StartPos:  5180,
					EndPos:    5189,
				},
				Expr: &cast.Unset{
					Position: &position.Position{
						StartLine: 268,
						EndLine:   268,
						StartPos:  5180,
						EndPos:    5188,
					},
					Expr: &expr.Variable{
						Position: &position.Position{
							StartLine: 268,
							EndLine:   268,
							StartPos:  5187,
							EndPos:    5188,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 268,
								EndLine:   268,
								StartPos:  5187,
								EndPos:    5188,
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
					StartPos:  5194,
					EndPos:    5201,
				},
				Expr: &binary.BitwiseAnd{
					Position: &position.Position{
						StartLine: 270,
						EndLine:   270,
						StartPos:  5194,
						EndPos:    5200,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 270,
							EndLine:   270,
							StartPos:  5194,
							EndPos:    5195,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 270,
								EndLine:   270,
								StartPos:  5194,
								EndPos:    5195,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 270,
							EndLine:   270,
							StartPos:  5199,
							EndPos:    5200,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 270,
								EndLine:   270,
								StartPos:  5199,
								EndPos:    5200,
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
					StartPos:  5205,
					EndPos:    5212,
				},
				Expr: &binary.BitwiseOr{
					Position: &position.Position{
						StartLine: 271,
						EndLine:   271,
						StartPos:  5205,
						EndPos:    5211,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 271,
							EndLine:   271,
							StartPos:  5205,
							EndPos:    5206,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 271,
								EndLine:   271,
								StartPos:  5205,
								EndPos:    5206,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 271,
							EndLine:   271,
							StartPos:  5210,
							EndPos:    5211,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 271,
								EndLine:   271,
								StartPos:  5210,
								EndPos:    5211,
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
					StartPos:  5216,
					EndPos:    5223,
				},
				Expr: &binary.BitwiseXor{
					Position: &position.Position{
						StartLine: 272,
						EndLine:   272,
						StartPos:  5216,
						EndPos:    5222,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 272,
							EndLine:   272,
							StartPos:  5216,
							EndPos:    5217,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 272,
								EndLine:   272,
								StartPos:  5216,
								EndPos:    5217,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 272,
							EndLine:   272,
							StartPos:  5221,
							EndPos:    5222,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 272,
								EndLine:   272,
								StartPos:  5221,
								EndPos:    5222,
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
					StartPos:  5227,
					EndPos:    5235,
				},
				Expr: &binary.BooleanAnd{
					Position: &position.Position{
						StartLine: 273,
						EndLine:   273,
						StartPos:  5227,
						EndPos:    5234,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 273,
							EndLine:   273,
							StartPos:  5227,
							EndPos:    5228,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 273,
								EndLine:   273,
								StartPos:  5227,
								EndPos:    5228,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 273,
							EndLine:   273,
							StartPos:  5233,
							EndPos:    5234,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 273,
								EndLine:   273,
								StartPos:  5233,
								EndPos:    5234,
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
					StartPos:  5239,
					EndPos:    5247,
				},
				Expr: &binary.BooleanOr{
					Position: &position.Position{
						StartLine: 274,
						EndLine:   274,
						StartPos:  5239,
						EndPos:    5246,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 274,
							EndLine:   274,
							StartPos:  5239,
							EndPos:    5240,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 274,
								EndLine:   274,
								StartPos:  5239,
								EndPos:    5240,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 274,
							EndLine:   274,
							StartPos:  5245,
							EndPos:    5246,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 274,
								EndLine:   274,
								StartPos:  5245,
								EndPos:    5246,
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
					StartPos:  5251,
					EndPos:    5258,
				},
				Expr: &binary.Concat{
					Position: &position.Position{
						StartLine: 275,
						EndLine:   275,
						StartPos:  5251,
						EndPos:    5257,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 275,
							EndLine:   275,
							StartPos:  5251,
							EndPos:    5252,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 275,
								EndLine:   275,
								StartPos:  5251,
								EndPos:    5252,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 275,
							EndLine:   275,
							StartPos:  5256,
							EndPos:    5257,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 275,
								EndLine:   275,
								StartPos:  5256,
								EndPos:    5257,
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
					StartPos:  5262,
					EndPos:    5269,
				},
				Expr: &binary.Div{
					Position: &position.Position{
						StartLine: 276,
						EndLine:   276,
						StartPos:  5262,
						EndPos:    5268,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 276,
							EndLine:   276,
							StartPos:  5262,
							EndPos:    5263,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 276,
								EndLine:   276,
								StartPos:  5262,
								EndPos:    5263,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 276,
							EndLine:   276,
							StartPos:  5267,
							EndPos:    5268,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 276,
								EndLine:   276,
								StartPos:  5267,
								EndPos:    5268,
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
					StartPos:  5273,
					EndPos:    5281,
				},
				Expr: &binary.Equal{
					Position: &position.Position{
						StartLine: 277,
						EndLine:   277,
						StartPos:  5273,
						EndPos:    5280,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 277,
							EndLine:   277,
							StartPos:  5273,
							EndPos:    5274,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 277,
								EndLine:   277,
								StartPos:  5273,
								EndPos:    5274,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 277,
							EndLine:   277,
							StartPos:  5279,
							EndPos:    5280,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 277,
								EndLine:   277,
								StartPos:  5279,
								EndPos:    5280,
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
					StartPos:  5285,
					EndPos:    5293,
				},
				Expr: &binary.GreaterOrEqual{
					Position: &position.Position{
						StartLine: 278,
						EndLine:   278,
						StartPos:  5285,
						EndPos:    5292,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 278,
							EndLine:   278,
							StartPos:  5285,
							EndPos:    5286,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 278,
								EndLine:   278,
								StartPos:  5285,
								EndPos:    5286,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 278,
							EndLine:   278,
							StartPos:  5291,
							EndPos:    5292,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 278,
								EndLine:   278,
								StartPos:  5291,
								EndPos:    5292,
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
					StartPos:  5297,
					EndPos:    5304,
				},
				Expr: &binary.Greater{
					Position: &position.Position{
						StartLine: 279,
						EndLine:   279,
						StartPos:  5297,
						EndPos:    5303,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 279,
							EndLine:   279,
							StartPos:  5297,
							EndPos:    5298,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 279,
								EndLine:   279,
								StartPos:  5297,
								EndPos:    5298,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 279,
							EndLine:   279,
							StartPos:  5302,
							EndPos:    5303,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 279,
								EndLine:   279,
								StartPos:  5302,
								EndPos:    5303,
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
					StartPos:  5308,
					EndPos:    5317,
				},
				Expr: &binary.Identical{
					Position: &position.Position{
						StartLine: 280,
						EndLine:   280,
						StartPos:  5308,
						EndPos:    5316,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 280,
							EndLine:   280,
							StartPos:  5308,
							EndPos:    5309,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 280,
								EndLine:   280,
								StartPos:  5308,
								EndPos:    5309,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 280,
							EndLine:   280,
							StartPos:  5315,
							EndPos:    5316,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 280,
								EndLine:   280,
								StartPos:  5315,
								EndPos:    5316,
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
					StartPos:  5321,
					EndPos:    5330,
				},
				Expr: &binary.LogicalAnd{
					Position: &position.Position{
						StartLine: 281,
						EndLine:   281,
						StartPos:  5321,
						EndPos:    5329,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 281,
							EndLine:   281,
							StartPos:  5321,
							EndPos:    5322,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 281,
								EndLine:   281,
								StartPos:  5321,
								EndPos:    5322,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 281,
							EndLine:   281,
							StartPos:  5328,
							EndPos:    5329,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 281,
								EndLine:   281,
								StartPos:  5328,
								EndPos:    5329,
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
					StartPos:  5334,
					EndPos:    5342,
				},
				Expr: &binary.LogicalOr{
					Position: &position.Position{
						StartLine: 282,
						EndLine:   282,
						StartPos:  5334,
						EndPos:    5341,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 282,
							EndLine:   282,
							StartPos:  5334,
							EndPos:    5335,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 282,
								EndLine:   282,
								StartPos:  5334,
								EndPos:    5335,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 282,
							EndLine:   282,
							StartPos:  5340,
							EndPos:    5341,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 282,
								EndLine:   282,
								StartPos:  5340,
								EndPos:    5341,
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
					StartPos:  5346,
					EndPos:    5355,
				},
				Expr: &binary.LogicalXor{
					Position: &position.Position{
						StartLine: 283,
						EndLine:   283,
						StartPos:  5346,
						EndPos:    5354,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 283,
							EndLine:   283,
							StartPos:  5346,
							EndPos:    5347,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 283,
								EndLine:   283,
								StartPos:  5346,
								EndPos:    5347,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 283,
							EndLine:   283,
							StartPos:  5353,
							EndPos:    5354,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 283,
								EndLine:   283,
								StartPos:  5353,
								EndPos:    5354,
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
					StartPos:  5359,
					EndPos:    5366,
				},
				Expr: &binary.Minus{
					Position: &position.Position{
						StartLine: 284,
						EndLine:   284,
						StartPos:  5359,
						EndPos:    5365,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 284,
							EndLine:   284,
							StartPos:  5359,
							EndPos:    5360,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 284,
								EndLine:   284,
								StartPos:  5359,
								EndPos:    5360,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 284,
							EndLine:   284,
							StartPos:  5364,
							EndPos:    5365,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 284,
								EndLine:   284,
								StartPos:  5364,
								EndPos:    5365,
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
					StartPos:  5370,
					EndPos:    5377,
				},
				Expr: &binary.Mod{
					Position: &position.Position{
						StartLine: 285,
						EndLine:   285,
						StartPos:  5370,
						EndPos:    5376,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 285,
							EndLine:   285,
							StartPos:  5370,
							EndPos:    5371,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 285,
								EndLine:   285,
								StartPos:  5370,
								EndPos:    5371,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 285,
							EndLine:   285,
							StartPos:  5375,
							EndPos:    5376,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 285,
								EndLine:   285,
								StartPos:  5375,
								EndPos:    5376,
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
					StartPos:  5381,
					EndPos:    5388,
				},
				Expr: &binary.Mul{
					Position: &position.Position{
						StartLine: 286,
						EndLine:   286,
						StartPos:  5381,
						EndPos:    5387,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 286,
							EndLine:   286,
							StartPos:  5381,
							EndPos:    5382,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 286,
								EndLine:   286,
								StartPos:  5381,
								EndPos:    5382,
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
					EndPos:    5400,
				},
				Expr: &binary.NotEqual{
					Position: &position.Position{
						StartLine: 287,
						EndLine:   287,
						StartPos:  5392,
						EndPos:    5399,
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
							StartPos:  5398,
							EndPos:    5399,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 287,
								EndLine:   287,
								StartPos:  5398,
								EndPos:    5399,
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
					StartPos:  5404,
					EndPos:    5413,
				},
				Expr: &binary.NotIdentical{
					Position: &position.Position{
						StartLine: 288,
						EndLine:   288,
						StartPos:  5404,
						EndPos:    5412,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 288,
							EndLine:   288,
							StartPos:  5404,
							EndPos:    5405,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 288,
								EndLine:   288,
								StartPos:  5404,
								EndPos:    5405,
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
					EndPos:    5424,
				},
				Expr: &binary.Plus{
					Position: &position.Position{
						StartLine: 289,
						EndLine:   289,
						StartPos:  5417,
						EndPos:    5423,
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
							StartPos:  5422,
							EndPos:    5423,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 289,
								EndLine:   289,
								StartPos:  5422,
								EndPos:    5423,
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
					StartPos:  5428,
					EndPos:    5436,
				},
				Expr: &binary.Pow{
					Position: &position.Position{
						StartLine: 290,
						EndLine:   290,
						StartPos:  5428,
						EndPos:    5435,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 290,
							EndLine:   290,
							StartPos:  5428,
							EndPos:    5429,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 290,
								EndLine:   290,
								StartPos:  5428,
								EndPos:    5429,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 290,
							EndLine:   290,
							StartPos:  5434,
							EndPos:    5435,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 290,
								EndLine:   290,
								StartPos:  5434,
								EndPos:    5435,
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
					StartPos:  5440,
					EndPos:    5448,
				},
				Expr: &binary.ShiftLeft{
					Position: &position.Position{
						StartLine: 291,
						EndLine:   291,
						StartPos:  5440,
						EndPos:    5447,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 291,
							EndLine:   291,
							StartPos:  5440,
							EndPos:    5441,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 291,
								EndLine:   291,
								StartPos:  5440,
								EndPos:    5441,
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
					EndPos:    5460,
				},
				Expr: &binary.ShiftRight{
					Position: &position.Position{
						StartLine: 292,
						EndLine:   292,
						StartPos:  5452,
						EndPos:    5459,
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
							StartPos:  5458,
							EndPos:    5459,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 292,
								EndLine:   292,
								StartPos:  5458,
								EndPos:    5459,
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
					StartPos:  5464,
					EndPos:    5472,
				},
				Expr: &binary.SmallerOrEqual{
					Position: &position.Position{
						StartLine: 293,
						EndLine:   293,
						StartPos:  5464,
						EndPos:    5471,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 293,
							EndLine:   293,
							StartPos:  5464,
							EndPos:    5465,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 293,
								EndLine:   293,
								StartPos:  5464,
								EndPos:    5465,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 293,
							EndLine:   293,
							StartPos:  5470,
							EndPos:    5471,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 293,
								EndLine:   293,
								StartPos:  5470,
								EndPos:    5471,
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
					StartPos:  5476,
					EndPos:    5483,
				},
				Expr: &binary.Smaller{
					Position: &position.Position{
						StartLine: 294,
						EndLine:   294,
						StartPos:  5476,
						EndPos:    5482,
					},
					Left: &expr.Variable{
						Position: &position.Position{
							StartLine: 294,
							EndLine:   294,
							StartPos:  5476,
							EndPos:    5477,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  5476,
								EndPos:    5477,
							},
							Value: "a",
						},
					},
					Right: &expr.Variable{
						Position: &position.Position{
							StartLine: 294,
							EndLine:   294,
							StartPos:  5481,
							EndPos:    5482,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  5481,
								EndPos:    5482,
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
					StartPos:  5488,
					EndPos:    5496,
				},
				Expr: &assign.Reference{
					Position: &position.Position{
						StartLine: 296,
						EndLine:   296,
						StartPos:  5488,
						EndPos:    5495,
					},
					Variable: &expr.Variable{
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
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 296,
							EndLine:   296,
							StartPos:  5494,
							EndPos:    5495,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  5494,
								EndPos:    5495,
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
					StartPos:  5500,
					EndPos:    5513,
				},
				Expr: &assign.Reference{
					Position: &position.Position{
						StartLine: 297,
						EndLine:   297,
						StartPos:  5500,
						EndPos:    5512,
					},
					Variable: &expr.Variable{
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
							Value: "a",
						},
					},
					Expression: &expr.New{
						Position: &position.Position{
							StartLine: 297,
							EndLine:   297,
							StartPos:  5506,
							EndPos:    5512,
						},
						Class: &name.Name{
							Position: &position.Position{
								StartLine: 297,
								EndLine:   297,
								StartPos:  5510,
								EndPos:    5512,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 297,
										EndLine:   297,
										StartPos:  5510,
										EndPos:    5512,
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
					StartLine: 298,
					EndLine:   298,
					StartPos:  5517,
					EndPos:    5534,
				},
				Expr: &assign.Reference{
					Position: &position.Position{
						StartLine: 298,
						EndLine:   298,
						StartPos:  5517,
						EndPos:    5533,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 298,
							EndLine:   298,
							StartPos:  5517,
							EndPos:    5518,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  5517,
								EndPos:    5518,
							},
							Value: "a",
						},
					},
					Expression: &expr.New{
						Position: &position.Position{
							StartLine: 298,
							EndLine:   298,
							StartPos:  5523,
							EndPos:    5533,
						},
						Class: &name.Name{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  5527,
								EndPos:    5529,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 298,
										EndLine:   298,
										StartPos:  5527,
										EndPos:    5529,
									},
									Value: "Foo",
								},
							},
						},
						ArgumentList: &node.ArgumentList{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  5530,
								EndPos:    5533,
							},
							Arguments: []node.Node{
								&node.Argument{
									Position: &position.Position{
										StartLine: 298,
										EndLine:   298,
										StartPos:  5531,
										EndPos:    5532,
									},
									Variadic:    false,
									IsReference: false,
									Expr: &expr.Variable{
										Position: &position.Position{
											StartLine: 298,
											EndLine:   298,
											StartPos:  5531,
											EndPos:    5532,
										},
										VarName: &node.Identifier{
											Position: &position.Position{
												StartLine: 298,
												EndLine:   298,
												StartPos:  5531,
												EndPos:    5532,
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
					StartLine: 299,
					EndLine:   299,
					StartPos:  5538,
					EndPos:    5545,
				},
				Expr: &assign.Assign{
					Position: &position.Position{
						StartLine: 299,
						EndLine:   299,
						StartPos:  5538,
						EndPos:    5544,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 299,
							EndLine:   299,
							StartPos:  5538,
							EndPos:    5539,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5538,
								EndPos:    5539,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 299,
							EndLine:   299,
							StartPos:  5543,
							EndPos:    5544,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5543,
								EndPos:    5544,
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
					StartPos:  5549,
					EndPos:    5557,
				},
				Expr: &assign.BitwiseAnd{
					Position: &position.Position{
						StartLine: 300,
						EndLine:   300,
						StartPos:  5549,
						EndPos:    5556,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 300,
							EndLine:   300,
							StartPos:  5549,
							EndPos:    5550,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5549,
								EndPos:    5550,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 300,
							EndLine:   300,
							StartPos:  5555,
							EndPos:    5556,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5555,
								EndPos:    5556,
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
					StartPos:  5561,
					EndPos:    5569,
				},
				Expr: &assign.BitwiseOr{
					Position: &position.Position{
						StartLine: 301,
						EndLine:   301,
						StartPos:  5561,
						EndPos:    5568,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 301,
							EndLine:   301,
							StartPos:  5561,
							EndPos:    5562,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 301,
								EndLine:   301,
								StartPos:  5561,
								EndPos:    5562,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 301,
							EndLine:   301,
							StartPos:  5567,
							EndPos:    5568,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 301,
								EndLine:   301,
								StartPos:  5567,
								EndPos:    5568,
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
					StartPos:  5573,
					EndPos:    5581,
				},
				Expr: &assign.BitwiseXor{
					Position: &position.Position{
						StartLine: 302,
						EndLine:   302,
						StartPos:  5573,
						EndPos:    5580,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 302,
							EndLine:   302,
							StartPos:  5573,
							EndPos:    5574,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 302,
								EndLine:   302,
								StartPos:  5573,
								EndPos:    5574,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 302,
							EndLine:   302,
							StartPos:  5579,
							EndPos:    5580,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 302,
								EndLine:   302,
								StartPos:  5579,
								EndPos:    5580,
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
					StartPos:  5585,
					EndPos:    5593,
				},
				Expr: &assign.Concat{
					Position: &position.Position{
						StartLine: 303,
						EndLine:   303,
						StartPos:  5585,
						EndPos:    5592,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 303,
							EndLine:   303,
							StartPos:  5585,
							EndPos:    5586,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 303,
								EndLine:   303,
								StartPos:  5585,
								EndPos:    5586,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 303,
							EndLine:   303,
							StartPos:  5591,
							EndPos:    5592,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 303,
								EndLine:   303,
								StartPos:  5591,
								EndPos:    5592,
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
					StartPos:  5597,
					EndPos:    5605,
				},
				Expr: &assign.Div{
					Position: &position.Position{
						StartLine: 304,
						EndLine:   304,
						StartPos:  5597,
						EndPos:    5604,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 304,
							EndLine:   304,
							StartPos:  5597,
							EndPos:    5598,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 304,
								EndLine:   304,
								StartPos:  5597,
								EndPos:    5598,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 304,
							EndLine:   304,
							StartPos:  5603,
							EndPos:    5604,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 304,
								EndLine:   304,
								StartPos:  5603,
								EndPos:    5604,
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
					StartPos:  5609,
					EndPos:    5617,
				},
				Expr: &assign.Minus{
					Position: &position.Position{
						StartLine: 305,
						EndLine:   305,
						StartPos:  5609,
						EndPos:    5616,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 305,
							EndLine:   305,
							StartPos:  5609,
							EndPos:    5610,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 305,
								EndLine:   305,
								StartPos:  5609,
								EndPos:    5610,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 305,
							EndLine:   305,
							StartPos:  5615,
							EndPos:    5616,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 305,
								EndLine:   305,
								StartPos:  5615,
								EndPos:    5616,
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
					StartPos:  5621,
					EndPos:    5629,
				},
				Expr: &assign.Mod{
					Position: &position.Position{
						StartLine: 306,
						EndLine:   306,
						StartPos:  5621,
						EndPos:    5628,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 306,
							EndLine:   306,
							StartPos:  5621,
							EndPos:    5622,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 306,
								EndLine:   306,
								StartPos:  5621,
								EndPos:    5622,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 306,
							EndLine:   306,
							StartPos:  5627,
							EndPos:    5628,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 306,
								EndLine:   306,
								StartPos:  5627,
								EndPos:    5628,
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
					StartPos:  5633,
					EndPos:    5641,
				},
				Expr: &assign.Mul{
					Position: &position.Position{
						StartLine: 307,
						EndLine:   307,
						StartPos:  5633,
						EndPos:    5640,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 307,
							EndLine:   307,
							StartPos:  5633,
							EndPos:    5634,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 307,
								EndLine:   307,
								StartPos:  5633,
								EndPos:    5634,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 307,
							EndLine:   307,
							StartPos:  5639,
							EndPos:    5640,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 307,
								EndLine:   307,
								StartPos:  5639,
								EndPos:    5640,
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
					StartPos:  5645,
					EndPos:    5653,
				},
				Expr: &assign.Plus{
					Position: &position.Position{
						StartLine: 308,
						EndLine:   308,
						StartPos:  5645,
						EndPos:    5652,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 308,
							EndLine:   308,
							StartPos:  5645,
							EndPos:    5646,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 308,
								EndLine:   308,
								StartPos:  5645,
								EndPos:    5646,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 308,
							EndLine:   308,
							StartPos:  5651,
							EndPos:    5652,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 308,
								EndLine:   308,
								StartPos:  5651,
								EndPos:    5652,
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
					StartPos:  5657,
					EndPos:    5666,
				},
				Expr: &assign.Pow{
					Position: &position.Position{
						StartLine: 309,
						EndLine:   309,
						StartPos:  5657,
						EndPos:    5665,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 309,
							EndLine:   309,
							StartPos:  5657,
							EndPos:    5658,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 309,
								EndLine:   309,
								StartPos:  5657,
								EndPos:    5658,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 309,
							EndLine:   309,
							StartPos:  5664,
							EndPos:    5665,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 309,
								EndLine:   309,
								StartPos:  5664,
								EndPos:    5665,
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
					StartPos:  5670,
					EndPos:    5679,
				},
				Expr: &assign.ShiftLeft{
					Position: &position.Position{
						StartLine: 310,
						EndLine:   310,
						StartPos:  5670,
						EndPos:    5678,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 310,
							EndLine:   310,
							StartPos:  5670,
							EndPos:    5671,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 310,
								EndLine:   310,
								StartPos:  5670,
								EndPos:    5671,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 310,
							EndLine:   310,
							StartPos:  5677,
							EndPos:    5678,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 310,
								EndLine:   310,
								StartPos:  5677,
								EndPos:    5678,
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
					StartPos:  5683,
					EndPos:    5692,
				},
				Expr: &assign.ShiftRight{
					Position: &position.Position{
						StartLine: 311,
						EndLine:   311,
						StartPos:  5683,
						EndPos:    5691,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 311,
							EndLine:   311,
							StartPos:  5683,
							EndPos:    5684,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 311,
								EndLine:   311,
								StartPos:  5683,
								EndPos:    5684,
							},
							Value: "a",
						},
					},
					Expression: &expr.Variable{
						Position: &position.Position{
							StartLine: 311,
							EndLine:   311,
							StartPos:  5690,
							EndPos:    5691,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 311,
								EndLine:   311,
								StartPos:  5690,
								EndPos:    5691,
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
					StartPos:  5699,
					EndPos:    5710,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 314,
						EndLine:   314,
						StartPos:  5699,
						EndPos:    5708,
					},
					Class: &name.FullyQualified{
						Position: &position.Position{
							StartLine: 314,
							EndLine:   314,
							StartPos:  5703,
							EndPos:    5706,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 314,
									EndLine:   314,
									StartPos:  5704,
									EndPos:    5706,
								},
								Value: "Foo",
							},
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 314,
							EndLine:   314,
							StartPos:  5707,
							EndPos:    5708,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 315,
					EndLine:   315,
					StartPos:  5735,
					EndPos:    5738,
				},
				Expr: &expr.PropertyFetch{
					Position: &position.Position{
						StartLine: 315,
						EndLine:   315,
						StartPos:  5735,
						EndPos:    5737,
					},
					Variable: &expr.MethodCall{
						Position: &position.Position{
							StartLine: 315,
							EndLine:   315,
							StartPos:  5731,
							EndPos:    5732,
						},
						Variable: &expr.New{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5715,
								EndPos:    5724,
							},
							Class: &name.FullyQualified{
								Position: &position.Position{
									StartLine: 315,
									EndLine:   315,
									StartPos:  5719,
									EndPos:    5722,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 315,
											EndLine:   315,
											StartPos:  5720,
											EndPos:    5722,
										},
										Value: "Foo",
									},
								},
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 315,
									EndLine:   315,
									StartPos:  5723,
									EndPos:    5724,
								},
							},
						},
						Method: &node.Identifier{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5728,
								EndPos:    5730,
							},
							Value: "bar",
						},
						ArgumentList: &node.ArgumentList{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5731,
								EndPos:    5732,
							},
						},
					},
					Property: &node.Identifier{
						Position: &position.Position{
							StartLine: 315,
							EndLine:   315,
							StartPos:  5735,
							EndPos:    5737,
						},
						Value: "baz",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 316,
					EndLine:   316,
					StartPos:  5758,
					EndPos:    5760,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 316,
						EndLine:   316,
						StartPos:  5758,
						EndPos:    5758,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 316,
							EndLine:   316,
							StartPos:  5755,
							EndPos:    5755,
						},
						Variable: &expr.New{
							Position: &position.Position{
								StartLine: 316,
								EndLine:   316,
								StartPos:  5743,
								EndPos:    5752,
							},
							Class: &name.FullyQualified{
								Position: &position.Position{
									StartLine: 316,
									EndLine:   316,
									StartPos:  5747,
									EndPos:    5750,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 316,
											EndLine:   316,
											StartPos:  5748,
											EndPos:    5750,
										},
										Value: "Foo",
									},
								},
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 316,
									EndLine:   316,
									StartPos:  5751,
									EndPos:    5752,
								},
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 316,
								EndLine:   316,
								StartPos:  5755,
								EndPos:    5755,
							},
							Value: "0",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 316,
							EndLine:   316,
							StartPos:  5758,
							EndPos:    5758,
						},
						Value: "0",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 317,
					EndLine:   317,
					StartPos:  5784,
					EndPos:    5786,
				},
				Expr: &expr.MethodCall{
					Position: &position.Position{
						StartLine: 317,
						EndLine:   317,
						StartPos:  5784,
						EndPos:    5785,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 317,
							EndLine:   317,
							StartPos:  5777,
							EndPos:    5777,
						},
						Variable: &expr.New{
							Position: &position.Position{
								StartLine: 317,
								EndLine:   317,
								StartPos:  5765,
								EndPos:    5774,
							},
							Class: &name.FullyQualified{
								Position: &position.Position{
									StartLine: 317,
									EndLine:   317,
									StartPos:  5769,
									EndPos:    5772,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 317,
											EndLine:   317,
											StartPos:  5770,
											EndPos:    5772,
										},
										Value: "Foo",
									},
								},
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 317,
									EndLine:   317,
									StartPos:  5773,
									EndPos:    5774,
								},
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 317,
								EndLine:   317,
								StartPos:  5777,
								EndPos:    5777,
							},
							Value: "0",
						},
					},
					Method: &node.Identifier{
						Position: &position.Position{
							StartLine: 317,
							EndLine:   317,
							StartPos:  5781,
							EndPos:    5783,
						},
						Value: "bar",
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 317,
							EndLine:   317,
							StartPos:  5784,
							EndPos:    5785,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 319,
					EndLine:   319,
					StartPos:  5791,
					EndPos:    5807,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 319,
						EndLine:   319,
						StartPos:  5791,
						EndPos:    5806,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 319,
							EndLine:   319,
							StartPos:  5791,
							EndPos:    5803,
						},
						Variable: &expr.Array{
							Position: &position.Position{
								StartLine: 319,
								EndLine:   319,
								StartPos:  5791,
								EndPos:    5800,
							},
							Items: []node.Node{
								&expr.ArrayItem{
									Position: &position.Position{
										StartLine: 319,
										EndLine:   319,
										StartPos:  5797,
										EndPos:    5799,
									},
									Val: &expr.ShortArray{
										Position: &position.Position{
											StartLine: 319,
											EndLine:   319,
											StartPos:  5797,
											EndPos:    5799,
										},
										Items: []node.Node{
											&expr.ArrayItem{
												Position: &position.Position{
													StartLine: 319,
													EndLine:   319,
													StartPos:  5798,
													EndPos:    5798,
												},
												Val: &scalar.Lnumber{
													Position: &position.Position{
														StartLine: 319,
														EndLine:   319,
														StartPos:  5798,
														EndPos:    5798,
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
								StartLine: 319,
								EndLine:   319,
								StartPos:  5802,
								EndPos:    5802,
							},
							Value: "0",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 319,
							EndLine:   319,
							StartPos:  5805,
							EndPos:    5805,
						},
						Value: "0",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 320,
					EndLine:   320,
					StartPos:  5811,
					EndPos:    5819,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 320,
						EndLine:   320,
						StartPos:  5811,
						EndPos:    5818,
					},
					Variable: &scalar.String{
						Position: &position.Position{
							StartLine: 320,
							EndLine:   320,
							StartPos:  5811,
							EndPos:    5815,
						},
						Value: "\"foo\"",
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 320,
							EndLine:   320,
							StartPos:  5817,
							EndPos:    5817,
						},
						Value: "0",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 321,
					EndLine:   321,
					StartPos:  5823,
					EndPos:    5829,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 321,
						EndLine:   321,
						StartPos:  5823,
						EndPos:    5828,
					},
					Variable: &expr.ConstFetch{
						Position: &position.Position{
							StartLine: 321,
							EndLine:   321,
							StartPos:  5823,
							EndPos:    5825,
						},
						Constant: &name.Name{
							Position: &position.Position{
								StartLine: 321,
								EndLine:   321,
								StartPos:  5823,
								EndPos:    5825,
							},
							Parts: []node.Node{
								&name.NamePart{
									Position: &position.Position{
										StartLine: 321,
										EndLine:   321,
										StartPos:  5823,
										EndPos:    5825,
									},
									Value: "foo",
								},
							},
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 321,
							EndLine:   321,
							StartPos:  5827,
							EndPos:    5827,
						},
						Value: "0",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 322,
					EndLine:   322,
					StartPos:  5833,
					EndPos:    5844,
				},
				Expr: &expr.ClassConstFetch{
					Position: &position.Position{
						StartLine: 322,
						EndLine:   322,
						StartPos:  5833,
						EndPos:    5843,
					},
					Class: &node.Identifier{
						Position: &position.Position{
							StartLine: 322,
							EndLine:   322,
							StartPos:  5833,
							EndPos:    5838,
						},
						Value: "static",
					},
					ConstantName: &node.Identifier{
						Position: &position.Position{
							StartLine: 322,
							EndLine:   322,
							StartPos:  5841,
							EndPos:    5843,
						},
						Value: "foo",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 324,
					EndLine:   324,
					StartPos:  5849,
					EndPos:    5857,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 324,
						EndLine:   324,
						StartPos:  5849,
						EndPos:    5856,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 324,
							EndLine:   324,
							StartPos:  5853,
							EndPos:    5856,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 324,
								EndLine:   324,
								StartPos:  5853,
								EndPos:    5856,
							},
							Value: "foo",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 325,
					EndLine:   325,
					StartPos:  5861,
					EndPos:    5875,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 325,
						EndLine:   325,
						StartPos:  5861,
						EndPos:    5874,
					},
					Class: &expr.StaticPropertyFetch{
						Position: &position.Position{
							StartLine: 325,
							EndLine:   325,
							StartPos:  5865,
							EndPos:    5874,
						},
						Class: &expr.Variable{
							Position: &position.Position{
								StartLine: 325,
								EndLine:   325,
								StartPos:  5865,
								EndPos:    5868,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5865,
									EndPos:    5868,
								},
								Value: "foo",
							},
						},
						Property: &expr.Variable{
							Position: &position.Position{
								StartLine: 325,
								EndLine:   325,
								StartPos:  5871,
								EndPos:    5874,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5871,
									EndPos:    5874,
								},
								Value: "bar",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 326,
					EndLine:   326,
					StartPos:  5879,
					EndPos:    5891,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 326,
						EndLine:   326,
						StartPos:  5879,
						EndPos:    5889,
					},
					Class: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 326,
							EndLine:   326,
							StartPos:  5889,
							EndPos:    5889,
						},
						Variable: &expr.PropertyFetch{
							Position: &position.Position{
								StartLine: 326,
								EndLine:   326,
								StartPos:  5887,
								EndPos:    5889,
							},
							Variable: &expr.Variable{
								Position: &position.Position{
									StartLine: 326,
									EndLine:   326,
									StartPos:  5883,
									EndPos:    5887,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 326,
										EndLine:   326,
										StartPos:  5883,
										EndPos:    5884,
									},
									Value: "a",
								},
							},
							Property: &node.Identifier{
								Position: &position.Position{
									StartLine: 326,
									EndLine:   326,
									StartPos:  5887,
									EndPos:    5887,
								},
								Value: "b",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 326,
								EndLine:   326,
								StartPos:  5889,
								EndPos:    5889,
							},
							Value: "0",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 327,
					EndLine:   327,
					StartPos:  5895,
					EndPos:    5926,
				},
				Expr: &expr.New{
					Position: &position.Position{
						StartLine: 327,
						EndLine:   327,
						StartPos:  5895,
						EndPos:    5924,
					},
					Class: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 327,
							EndLine:   327,
							StartPos:  5924,
							EndPos:    5924,
						},
						Variable: &expr.PropertyFetch{
							Position: &position.Position{
								StartLine: 327,
								EndLine:   327,
								StartPos:  5922,
								EndPos:    5924,
							},
							Variable: &expr.PropertyFetch{
								Position: &position.Position{
									StartLine: 327,
									EndLine:   327,
									StartPos:  5918,
									EndPos:    5922,
								},
								Variable: &expr.ArrayDimFetch{
									Position: &position.Position{
										StartLine: 327,
										EndLine:   327,
										StartPos:  5905,
										EndPos:    5919,
									},
									Variable: &expr.PropertyFetch{
										Position: &position.Position{
											StartLine: 327,
											EndLine:   327,
											StartPos:  5903,
											EndPos:    5914,
										},
										Variable: &expr.Variable{
											Position: &position.Position{
												StartLine: 327,
												EndLine:   327,
												StartPos:  5899,
												EndPos:    5903,
											},
											VarName: &node.Identifier{
												Position: &position.Position{
													StartLine: 327,
													EndLine:   327,
													StartPos:  5899,
													EndPos:    5900,
												},
												Value: "a",
											},
										},
										Property: &node.Identifier{
											Position: &position.Position{
												StartLine: 327,
												EndLine:   327,
												StartPos:  5903,
												EndPos:    5903,
											},
											Value: "b",
										},
									},
									Dim: &expr.Ternary{
										Position: &position.Position{
											StartLine: 327,
											EndLine:   327,
											StartPos:  5905,
											EndPos:    5914,
										},
										Condition: &expr.Variable{
											Position: &position.Position{
												StartLine: 327,
												EndLine:   327,
												StartPos:  5905,
												EndPos:    5906,
											},
											VarName: &node.Identifier{
												Position: &position.Position{
													StartLine: 327,
													EndLine:   327,
													StartPos:  5905,
													EndPos:    5906,
												},
												Value: "b",
											},
										},
										IfFalse: &expr.ConstFetch{
											Position: &position.Position{
												StartLine: 327,
												EndLine:   327,
												StartPos:  5911,
												EndPos:    5914,
											},
											Constant: &name.Name{
												Position: &position.Position{
													StartLine: 327,
													EndLine:   327,
													StartPos:  5911,
													EndPos:    5914,
												},
												Parts: []node.Node{
													&name.NamePart{
														Position: &position.Position{
															StartLine: 327,
															EndLine:   327,
															StartPos:  5911,
															EndPos:    5914,
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
										StartLine: 327,
										EndLine:   327,
										StartPos:  5918,
										EndPos:    5919,
									},
									VarName: &node.Identifier{
										Position: &position.Position{
											StartLine: 327,
											EndLine:   327,
											StartPos:  5918,
											EndPos:    5919,
										},
										Value: "c",
									},
								},
							},
							Property: &node.Identifier{
								Position: &position.Position{
									StartLine: 327,
									EndLine:   327,
									StartPos:  5922,
									EndPos:    5922,
								},
								Value: "d",
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 327,
								EndLine:   327,
								StartPos:  5924,
								EndPos:    5924,
							},
							Value: "0",
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 327,
					EndLine:   327,
					StartPos:  5927,
					EndPos:    5945,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 327,
							EndLine:   327,
							StartPos:  5934,
							EndPos:    5944,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 327,
								EndLine:   327,
								StartPos:  5934,
								EndPos:    5935,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 327,
									EndLine:   327,
									StartPos:  5934,
									EndPos:    5935,
								},
								Value: "a",
							},
						},
						Expr: &expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 327,
								EndLine:   327,
								StartPos:  5939,
								EndPos:    5944,
							},
							Variable: &expr.ShortArray{
								Position: &position.Position{
									StartLine: 327,
									EndLine:   327,
									StartPos:  5939,
									EndPos:    5941,
								},
								Items: []node.Node{
									&expr.ArrayItem{
										Position: &position.Position{
											StartLine: 327,
											EndLine:   327,
											StartPos:  5940,
											EndPos:    5940,
										},
										Val: &scalar.Lnumber{
											Position: &position.Position{
												StartLine: 327,
												EndLine:   327,
												StartPos:  5940,
												EndPos:    5940,
											},
											Value: "1",
										},
									},
								},
							},
							Dim: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 327,
									EndLine:   327,
									StartPos:  5943,
									EndPos:    5943,
								},
								Value: "0",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 329,
					EndLine:   329,
					StartPos:  5950,
					EndPos:    5964,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 329,
							EndLine:   329,
							StartPos:  5957,
							EndPos:    5963,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5957,
								EndPos:    5958,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5957,
									EndPos:    5958,
								},
								Value: "a",
							},
						},
						Expr: &expr.BooleanNot{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5962,
								EndPos:    5963,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5963,
									EndPos:    5963,
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
					StartPos:  5968,
					EndPos:    5982,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 330,
							EndLine:   330,
							StartPos:  5975,
							EndPos:    5981,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  5975,
								EndPos:    5976,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  5975,
									EndPos:    5976,
								},
								Value: "a",
							},
						},
						Expr: &expr.BitwiseNot{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  5980,
								EndPos:    5981,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  5981,
									EndPos:    5981,
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
					StartPos:  5986,
					EndPos:    6000,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 331,
							EndLine:   331,
							StartPos:  5993,
							EndPos:    5999,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 331,
								EndLine:   331,
								StartPos:  5993,
								EndPos:    5994,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 331,
									EndLine:   331,
									StartPos:  5993,
									EndPos:    5994,
								},
								Value: "a",
							},
						},
						Expr: &expr.UnaryPlus{
							Position: &position.Position{
								StartLine: 331,
								EndLine:   331,
								StartPos:  5998,
								EndPos:    5999,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 331,
									EndLine:   331,
									StartPos:  5999,
									EndPos:    5999,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 332,
					EndLine:   332,
					StartPos:  6004,
					EndPos:    6018,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 332,
							EndLine:   332,
							StartPos:  6011,
							EndPos:    6017,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 332,
								EndLine:   332,
								StartPos:  6011,
								EndPos:    6012,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 332,
									EndLine:   332,
									StartPos:  6011,
									EndPos:    6012,
								},
								Value: "a",
							},
						},
						Expr: &expr.UnaryMinus{
							Position: &position.Position{
								StartLine: 332,
								EndLine:   332,
								StartPos:  6016,
								EndPos:    6017,
							},
							Expr: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 332,
									EndLine:   332,
									StartPos:  6017,
									EndPos:    6017,
								},
								Value: "1",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 333,
					EndLine:   333,
					StartPos:  6022,
					EndPos:    6037,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 333,
							EndLine:   333,
							StartPos:  6029,
							EndPos:    6035,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 333,
								EndLine:   333,
								StartPos:  6029,
								EndPos:    6030,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 333,
									EndLine:   333,
									StartPos:  6029,
									EndPos:    6030,
								},
								Value: "a",
							},
						},
						Expr: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 333,
								EndLine:   333,
								StartPos:  6035,
								EndPos:    6035,
							},
							Value: "1",
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 334,
					EndLine:   334,
					StartPos:  6041,
					EndPos:    6059,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 334,
							EndLine:   334,
							StartPos:  6048,
							EndPos:    6058,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 334,
								EndLine:   334,
								StartPos:  6048,
								EndPos:    6049,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 334,
									EndLine:   334,
									StartPos:  6048,
									EndPos:    6049,
								},
								Value: "a",
							},
						},
						Expr: &expr.Ternary{
							Position: &position.Position{
								StartLine: 334,
								EndLine:   334,
								StartPos:  6053,
								EndPos:    6058,
							},
							Condition: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 334,
									EndLine:   334,
									StartPos:  6053,
									EndPos:    6053,
								},
								Value: "1",
							},
							IfFalse: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 334,
									EndLine:   334,
									StartPos:  6058,
									EndPos:    6058,
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
					StartPos:  6063,
					EndPos:    6084,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 335,
							EndLine:   335,
							StartPos:  6070,
							EndPos:    6083,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 335,
								EndLine:   335,
								StartPos:  6070,
								EndPos:    6071,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  6070,
									EndPos:    6071,
								},
								Value: "a",
							},
						},
						Expr: &expr.Ternary{
							Position: &position.Position{
								StartLine: 335,
								EndLine:   335,
								StartPos:  6075,
								EndPos:    6083,
							},
							Condition: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  6075,
									EndPos:    6075,
								},
								Value: "1",
							},
							IfTrue: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  6079,
									EndPos:    6079,
								},
								Value: "2",
							},
							IfFalse: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  6083,
									EndPos:    6083,
								},
								Value: "3",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 336,
					EndLine:   336,
					StartPos:  6088,
					EndPos:    6105,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 336,
							EndLine:   336,
							StartPos:  6095,
							EndPos:    6104,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 336,
								EndLine:   336,
								StartPos:  6095,
								EndPos:    6096,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 336,
									EndLine:   336,
									StartPos:  6095,
									EndPos:    6096,
								},
								Value: "a",
							},
						},
						Expr: &binary.BitwiseAnd{
							Position: &position.Position{
								StartLine: 336,
								EndLine:   336,
								StartPos:  6100,
								EndPos:    6104,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 336,
									EndLine:   336,
									StartPos:  6100,
									EndPos:    6100,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 336,
									EndLine:   336,
									StartPos:  6104,
									EndPos:    6104,
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
					StartPos:  6109,
					EndPos:    6126,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 337,
							EndLine:   337,
							StartPos:  6116,
							EndPos:    6125,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 337,
								EndLine:   337,
								StartPos:  6116,
								EndPos:    6117,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 337,
									EndLine:   337,
									StartPos:  6116,
									EndPos:    6117,
								},
								Value: "a",
							},
						},
						Expr: &binary.BitwiseOr{
							Position: &position.Position{
								StartLine: 337,
								EndLine:   337,
								StartPos:  6121,
								EndPos:    6125,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 337,
									EndLine:   337,
									StartPos:  6121,
									EndPos:    6121,
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
					EndPos:    6147,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 338,
							EndLine:   338,
							StartPos:  6137,
							EndPos:    6146,
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
						Expr: &binary.BitwiseXor{
							Position: &position.Position{
								StartLine: 338,
								EndLine:   338,
								StartPos:  6142,
								EndPos:    6146,
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
									StartPos:  6146,
									EndPos:    6146,
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
					StartPos:  6151,
					EndPos:    6169,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 339,
							EndLine:   339,
							StartPos:  6158,
							EndPos:    6168,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 339,
								EndLine:   339,
								StartPos:  6158,
								EndPos:    6159,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 339,
									EndLine:   339,
									StartPos:  6158,
									EndPos:    6159,
								},
								Value: "a",
							},
						},
						Expr: &binary.BooleanAnd{
							Position: &position.Position{
								StartLine: 339,
								EndLine:   339,
								StartPos:  6163,
								EndPos:    6168,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 339,
									EndLine:   339,
									StartPos:  6163,
									EndPos:    6163,
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
					EndPos:    6191,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 340,
							EndLine:   340,
							StartPos:  6180,
							EndPos:    6190,
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
						Expr: &binary.BooleanOr{
							Position: &position.Position{
								StartLine: 340,
								EndLine:   340,
								StartPos:  6185,
								EndPos:    6190,
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
									StartPos:  6190,
									EndPos:    6190,
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
					StartPos:  6195,
					EndPos:    6212,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 341,
							EndLine:   341,
							StartPos:  6202,
							EndPos:    6211,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  6202,
								EndPos:    6203,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  6202,
									EndPos:    6203,
								},
								Value: "a",
							},
						},
						Expr: &binary.Concat{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  6207,
								EndPos:    6211,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  6207,
									EndPos:    6207,
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
					EndPos:    6233,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 342,
							EndLine:   342,
							StartPos:  6223,
							EndPos:    6232,
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
						Expr: &binary.Div{
							Position: &position.Position{
								StartLine: 342,
								EndLine:   342,
								StartPos:  6228,
								EndPos:    6232,
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
									StartPos:  6232,
									EndPos:    6232,
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
					StartPos:  6237,
					EndPos:    6255,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 343,
							EndLine:   343,
							StartPos:  6244,
							EndPos:    6254,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 343,
								EndLine:   343,
								StartPos:  6244,
								EndPos:    6245,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 343,
									EndLine:   343,
									StartPos:  6244,
									EndPos:    6245,
								},
								Value: "a",
							},
						},
						Expr: &binary.Equal{
							Position: &position.Position{
								StartLine: 343,
								EndLine:   343,
								StartPos:  6249,
								EndPos:    6254,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 343,
									EndLine:   343,
									StartPos:  6249,
									EndPos:    6249,
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
					EndPos:    6277,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 344,
							EndLine:   344,
							StartPos:  6266,
							EndPos:    6276,
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
						Expr: &binary.GreaterOrEqual{
							Position: &position.Position{
								StartLine: 344,
								EndLine:   344,
								StartPos:  6271,
								EndPos:    6276,
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
									StartPos:  6276,
									EndPos:    6276,
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
					StartPos:  6281,
					EndPos:    6298,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 345,
							EndLine:   345,
							StartPos:  6288,
							EndPos:    6297,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 345,
								EndLine:   345,
								StartPos:  6288,
								EndPos:    6289,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 345,
									EndLine:   345,
									StartPos:  6288,
									EndPos:    6289,
								},
								Value: "a",
							},
						},
						Expr: &binary.Greater{
							Position: &position.Position{
								StartLine: 345,
								EndLine:   345,
								StartPos:  6293,
								EndPos:    6297,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 345,
									EndLine:   345,
									StartPos:  6293,
									EndPos:    6293,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 345,
									EndLine:   345,
									StartPos:  6297,
									EndPos:    6297,
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
					StartPos:  6302,
					EndPos:    6321,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 346,
							EndLine:   346,
							StartPos:  6309,
							EndPos:    6320,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 346,
								EndLine:   346,
								StartPos:  6309,
								EndPos:    6310,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  6309,
									EndPos:    6310,
								},
								Value: "a",
							},
						},
						Expr: &binary.Identical{
							Position: &position.Position{
								StartLine: 346,
								EndLine:   346,
								StartPos:  6314,
								EndPos:    6320,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  6314,
									EndPos:    6314,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  6320,
									EndPos:    6320,
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
					StartPos:  6325,
					EndPos:    6344,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 347,
							EndLine:   347,
							StartPos:  6332,
							EndPos:    6343,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 347,
								EndLine:   347,
								StartPos:  6332,
								EndPos:    6333,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 347,
									EndLine:   347,
									StartPos:  6332,
									EndPos:    6333,
								},
								Value: "a",
							},
						},
						Expr: &binary.LogicalAnd{
							Position: &position.Position{
								StartLine: 347,
								EndLine:   347,
								StartPos:  6337,
								EndPos:    6343,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 347,
									EndLine:   347,
									StartPos:  6337,
									EndPos:    6337,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 347,
									EndLine:   347,
									StartPos:  6343,
									EndPos:    6343,
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
					StartPos:  6348,
					EndPos:    6366,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 348,
							EndLine:   348,
							StartPos:  6355,
							EndPos:    6365,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 348,
								EndLine:   348,
								StartPos:  6355,
								EndPos:    6356,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 348,
									EndLine:   348,
									StartPos:  6355,
									EndPos:    6356,
								},
								Value: "a",
							},
						},
						Expr: &binary.LogicalOr{
							Position: &position.Position{
								StartLine: 348,
								EndLine:   348,
								StartPos:  6360,
								EndPos:    6365,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 348,
									EndLine:   348,
									StartPos:  6360,
									EndPos:    6360,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 348,
									EndLine:   348,
									StartPos:  6365,
									EndPos:    6365,
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
					StartPos:  6370,
					EndPos:    6389,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 349,
							EndLine:   349,
							StartPos:  6377,
							EndPos:    6388,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 349,
								EndLine:   349,
								StartPos:  6377,
								EndPos:    6378,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 349,
									EndLine:   349,
									StartPos:  6377,
									EndPos:    6378,
								},
								Value: "a",
							},
						},
						Expr: &binary.LogicalXor{
							Position: &position.Position{
								StartLine: 349,
								EndLine:   349,
								StartPos:  6382,
								EndPos:    6388,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 349,
									EndLine:   349,
									StartPos:  6382,
									EndPos:    6382,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 349,
									EndLine:   349,
									StartPos:  6388,
									EndPos:    6388,
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
					StartPos:  6393,
					EndPos:    6410,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 350,
							EndLine:   350,
							StartPos:  6400,
							EndPos:    6409,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 350,
								EndLine:   350,
								StartPos:  6400,
								EndPos:    6401,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 350,
									EndLine:   350,
									StartPos:  6400,
									EndPos:    6401,
								},
								Value: "a",
							},
						},
						Expr: &binary.Minus{
							Position: &position.Position{
								StartLine: 350,
								EndLine:   350,
								StartPos:  6405,
								EndPos:    6409,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 350,
									EndLine:   350,
									StartPos:  6405,
									EndPos:    6405,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 350,
									EndLine:   350,
									StartPos:  6409,
									EndPos:    6409,
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
					StartPos:  6414,
					EndPos:    6431,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 351,
							EndLine:   351,
							StartPos:  6421,
							EndPos:    6430,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 351,
								EndLine:   351,
								StartPos:  6421,
								EndPos:    6422,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 351,
									EndLine:   351,
									StartPos:  6421,
									EndPos:    6422,
								},
								Value: "a",
							},
						},
						Expr: &binary.Mod{
							Position: &position.Position{
								StartLine: 351,
								EndLine:   351,
								StartPos:  6426,
								EndPos:    6430,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 351,
									EndLine:   351,
									StartPos:  6426,
									EndPos:    6426,
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
					EndPos:    6452,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 352,
							EndLine:   352,
							StartPos:  6442,
							EndPos:    6451,
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
						Expr: &binary.Mul{
							Position: &position.Position{
								StartLine: 352,
								EndLine:   352,
								StartPos:  6447,
								EndPos:    6451,
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
									StartPos:  6451,
									EndPos:    6451,
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
					StartPos:  6456,
					EndPos:    6474,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 353,
							EndLine:   353,
							StartPos:  6463,
							EndPos:    6473,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 353,
								EndLine:   353,
								StartPos:  6463,
								EndPos:    6464,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 353,
									EndLine:   353,
									StartPos:  6463,
									EndPos:    6464,
								},
								Value: "a",
							},
						},
						Expr: &binary.NotEqual{
							Position: &position.Position{
								StartLine: 353,
								EndLine:   353,
								StartPos:  6468,
								EndPos:    6473,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 353,
									EndLine:   353,
									StartPos:  6468,
									EndPos:    6468,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 353,
									EndLine:   353,
									StartPos:  6473,
									EndPos:    6473,
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
					StartPos:  6478,
					EndPos:    6497,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 354,
							EndLine:   354,
							StartPos:  6485,
							EndPos:    6496,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 354,
								EndLine:   354,
								StartPos:  6485,
								EndPos:    6486,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 354,
									EndLine:   354,
									StartPos:  6485,
									EndPos:    6486,
								},
								Value: "a",
							},
						},
						Expr: &binary.NotIdentical{
							Position: &position.Position{
								StartLine: 354,
								EndLine:   354,
								StartPos:  6490,
								EndPos:    6496,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 354,
									EndLine:   354,
									StartPos:  6490,
									EndPos:    6490,
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
					EndPos:    6518,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 355,
							EndLine:   355,
							StartPos:  6508,
							EndPos:    6517,
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
						Expr: &binary.Plus{
							Position: &position.Position{
								StartLine: 355,
								EndLine:   355,
								StartPos:  6513,
								EndPos:    6517,
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
									StartPos:  6517,
									EndPos:    6517,
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
					StartPos:  6522,
					EndPos:    6540,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 356,
							EndLine:   356,
							StartPos:  6529,
							EndPos:    6539,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 356,
								EndLine:   356,
								StartPos:  6529,
								EndPos:    6530,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 356,
									EndLine:   356,
									StartPos:  6529,
									EndPos:    6530,
								},
								Value: "a",
							},
						},
						Expr: &binary.Pow{
							Position: &position.Position{
								StartLine: 356,
								EndLine:   356,
								StartPos:  6534,
								EndPos:    6539,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 356,
									EndLine:   356,
									StartPos:  6534,
									EndPos:    6534,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 356,
									EndLine:   356,
									StartPos:  6539,
									EndPos:    6539,
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
					StartPos:  6544,
					EndPos:    6562,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 357,
							EndLine:   357,
							StartPos:  6551,
							EndPos:    6561,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 357,
								EndLine:   357,
								StartPos:  6551,
								EndPos:    6552,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 357,
									EndLine:   357,
									StartPos:  6551,
									EndPos:    6552,
								},
								Value: "a",
							},
						},
						Expr: &binary.ShiftLeft{
							Position: &position.Position{
								StartLine: 357,
								EndLine:   357,
								StartPos:  6556,
								EndPos:    6561,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 357,
									EndLine:   357,
									StartPos:  6556,
									EndPos:    6556,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 357,
									EndLine:   357,
									StartPos:  6561,
									EndPos:    6561,
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
					StartPos:  6566,
					EndPos:    6584,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 358,
							EndLine:   358,
							StartPos:  6573,
							EndPos:    6583,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 358,
								EndLine:   358,
								StartPos:  6573,
								EndPos:    6574,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 358,
									EndLine:   358,
									StartPos:  6573,
									EndPos:    6574,
								},
								Value: "a",
							},
						},
						Expr: &binary.ShiftRight{
							Position: &position.Position{
								StartLine: 358,
								EndLine:   358,
								StartPos:  6578,
								EndPos:    6583,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 358,
									EndLine:   358,
									StartPos:  6578,
									EndPos:    6578,
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
					EndPos:    6606,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 359,
							EndLine:   359,
							StartPos:  6595,
							EndPos:    6605,
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
						Expr: &binary.SmallerOrEqual{
							Position: &position.Position{
								StartLine: 359,
								EndLine:   359,
								StartPos:  6600,
								EndPos:    6605,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 359,
									EndLine:   359,
									StartPos:  6600,
									EndPos:    6600,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 359,
									EndLine:   359,
									StartPos:  6605,
									EndPos:    6605,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 360,
					EndLine:   360,
					StartPos:  6610,
					EndPos:    6627,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 360,
							EndLine:   360,
							StartPos:  6617,
							EndPos:    6626,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 360,
								EndLine:   360,
								StartPos:  6617,
								EndPos:    6618,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 360,
									EndLine:   360,
									StartPos:  6617,
									EndPos:    6618,
								},
								Value: "a",
							},
						},
						Expr: &binary.Smaller{
							Position: &position.Position{
								StartLine: 360,
								EndLine:   360,
								StartPos:  6622,
								EndPos:    6626,
							},
							Left: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 360,
									EndLine:   360,
									StartPos:  6622,
									EndPos:    6622,
								},
								Value: "1",
							},
							Right: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 360,
									EndLine:   360,
									StartPos:  6626,
									EndPos:    6626,
								},
								Value: "2",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 361,
					EndLine:   361,
					StartPos:  6631,
					EndPos:    6651,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 361,
							EndLine:   361,
							StartPos:  6638,
							EndPos:    6650,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 361,
								EndLine:   361,
								StartPos:  6638,
								EndPos:    6639,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 361,
									EndLine:   361,
									StartPos:  6638,
									EndPos:    6639,
								},
								Value: "a",
							},
						},
						Expr: &expr.ClassConstFetch{
							Position: &position.Position{
								StartLine: 361,
								EndLine:   361,
								StartPos:  6643,
								EndPos:    6650,
							},
							Class: &name.Name{
								Position: &position.Position{
									StartLine: 361,
									EndLine:   361,
									StartPos:  6643,
									EndPos:    6645,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 361,
											EndLine:   361,
											StartPos:  6643,
											EndPos:    6645,
										},
										Value: "Foo",
									},
								},
							},
							ConstantName: &node.Identifier{
								Position: &position.Position{
									StartLine: 361,
									EndLine:   361,
									StartPos:  6648,
									EndPos:    6650,
								},
								Value: "bar",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 362,
					EndLine:   362,
					StartPos:  6655,
					EndPos:    6677,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 362,
							EndLine:   362,
							StartPos:  6662,
							EndPos:    6676,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 362,
								EndLine:   362,
								StartPos:  6662,
								EndPos:    6663,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 362,
									EndLine:   362,
									StartPos:  6662,
									EndPos:    6663,
								},
								Value: "a",
							},
						},
						Expr: &expr.ClassConstFetch{
							Position: &position.Position{
								StartLine: 362,
								EndLine:   362,
								StartPos:  6667,
								EndPos:    6676,
							},
							Class: &name.Name{
								Position: &position.Position{
									StartLine: 362,
									EndLine:   362,
									StartPos:  6667,
									EndPos:    6669,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 362,
											EndLine:   362,
											StartPos:  6667,
											EndPos:    6669,
										},
										Value: "Foo",
									},
								},
							},
							ConstantName: &node.Identifier{
								Position: &position.Position{
									StartLine: 362,
									EndLine:   362,
									StartPos:  6672,
									EndPos:    6676,
								},
								Value: "class",
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 363,
					EndLine:   363,
					StartPos:  6681,
					EndPos:    6702,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 363,
							EndLine:   363,
							StartPos:  6688,
							EndPos:    6701,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 363,
								EndLine:   363,
								StartPos:  6688,
								EndPos:    6689,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 363,
									EndLine:   363,
									StartPos:  6688,
									EndPos:    6689,
								},
								Value: "a",
							},
						},
						Expr: &scalar.MagicConstant{
							Position: &position.Position{
								StartLine: 363,
								EndLine:   363,
								StartPos:  6693,
								EndPos:    6701,
							},
							Value: "__CLASS__",
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 364,
					EndLine:   364,
					StartPos:  6706,
					EndPos:    6721,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 364,
							EndLine:   364,
							StartPos:  6713,
							EndPos:    6720,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 364,
								EndLine:   364,
								StartPos:  6713,
								EndPos:    6714,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 364,
									EndLine:   364,
									StartPos:  6713,
									EndPos:    6714,
								},
								Value: "a",
							},
						},
						Expr: &expr.ConstFetch{
							Position: &position.Position{
								StartLine: 364,
								EndLine:   364,
								StartPos:  6718,
								EndPos:    6720,
							},
							Constant: &name.Name{
								Position: &position.Position{
									StartLine: 364,
									EndLine:   364,
									StartPos:  6718,
									EndPos:    6720,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 364,
											EndLine:   364,
											StartPos:  6718,
											EndPos:    6720,
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
					StartPos:  6725,
					EndPos:    6750,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 365,
							EndLine:   365,
							StartPos:  6732,
							EndPos:    6749,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 365,
								EndLine:   365,
								StartPos:  6732,
								EndPos:    6733,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 365,
									EndLine:   365,
									StartPos:  6732,
									EndPos:    6733,
								},
								Value: "a",
							},
						},
						Expr: &expr.ConstFetch{
							Position: &position.Position{
								StartLine: 365,
								EndLine:   365,
								StartPos:  6737,
								EndPos:    6749,
							},
							Constant: &name.Relative{
								Position: &position.Position{
									StartLine: 365,
									EndLine:   365,
									StartPos:  6737,
									EndPos:    6749,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 365,
											EndLine:   365,
											StartPos:  6747,
											EndPos:    6749,
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
					StartLine: 366,
					EndLine:   366,
					StartPos:  6754,
					EndPos:    6770,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 366,
							EndLine:   366,
							StartPos:  6761,
							EndPos:    6769,
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
						Expr: &expr.ConstFetch{
							Position: &position.Position{
								StartLine: 366,
								EndLine:   366,
								StartPos:  6766,
								EndPos:    6769,
							},
							Constant: &name.FullyQualified{
								Position: &position.Position{
									StartLine: 366,
									EndLine:   366,
									StartPos:  6766,
									EndPos:    6769,
								},
								Parts: []node.Node{
									&name.NamePart{
										Position: &position.Position{
											StartLine: 366,
											EndLine:   366,
											StartPos:  6767,
											EndPos:    6769,
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
					StartLine: 367,
					EndLine:   367,
					StartPos:  6774,
					EndPos:    6793,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 367,
							EndLine:   367,
							StartPos:  6781,
							EndPos:    6792,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 367,
								EndLine:   367,
								StartPos:  6781,
								EndPos:    6782,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 367,
									EndLine:   367,
									StartPos:  6781,
									EndPos:    6782,
								},
								Value: "a",
							},
						},
						Expr: &expr.Array{
							Position: &position.Position{
								StartLine: 367,
								EndLine:   367,
								StartPos:  6786,
								EndPos:    6792,
							},
						},
					},
				},
			},
			&stmt.Static{
				Position: &position.Position{
					StartLine: 368,
					EndLine:   368,
					StartPos:  6797,
					EndPos:    6825,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 368,
							EndLine:   368,
							StartPos:  6804,
							EndPos:    6824,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 368,
								EndLine:   368,
								StartPos:  6804,
								EndPos:    6805,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 368,
									EndLine:   368,
									StartPos:  6804,
									EndPos:    6805,
								},
								Value: "a",
							},
						},
						Expr: &expr.Array{
							Position: &position.Position{
								StartLine: 368,
								EndLine:   368,
								StartPos:  6809,
								EndPos:    6824,
							},
							Items: []node.Node{
								&expr.ArrayItem{
									Position: &position.Position{
										StartLine: 368,
										EndLine:   368,
										StartPos:  6815,
										EndPos:    6820,
									},
									Key: &scalar.Lnumber{
										Position: &position.Position{
											StartLine: 368,
											EndLine:   368,
											StartPos:  6815,
											EndPos:    6815,
										},
										Value: "1",
									},
									Val: &scalar.Lnumber{
										Position: &position.Position{
											StartLine: 368,
											EndLine:   368,
											StartPos:  6820,
											EndPos:    6820,
										},
										Value: "1",
									},
								},
								&expr.ArrayItem{
									Position: &position.Position{
										StartLine: 368,
										EndLine:   368,
										StartPos:  6823,
										EndPos:    6823,
									},
									Val: &scalar.Lnumber{
										Position: &position.Position{
											StartLine: 368,
											EndLine:   368,
											StartPos:  6823,
											EndPos:    6823,
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
					StartLine: 369,
					EndLine:   369,
					StartPos:  6829,
					EndPos:    6855,
				},
				Vars: []node.Node{
					&stmt.StaticVar{
						Position: &position.Position{
							StartLine: 369,
							EndLine:   369,
							StartPos:  6836,
							EndPos:    6854,
						},
						Variable: &expr.Variable{
							Position: &position.Position{
								StartLine: 369,
								EndLine:   369,
								StartPos:  6836,
								EndPos:    6837,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 369,
									EndLine:   369,
									StartPos:  6836,
									EndPos:    6837,
								},
								Value: "a",
							},
						},
						Expr: &expr.ArrayDimFetch{
							Position: &position.Position{
								StartLine: 369,
								EndLine:   369,
								StartPos:  6841,
								EndPos:    6854,
							},
							Variable: &expr.ShortArray{
								Position: &position.Position{
									StartLine: 369,
									EndLine:   369,
									StartPos:  6841,
									EndPos:    6851,
								},
								Items: []node.Node{
									&expr.ArrayItem{
										Position: &position.Position{
											StartLine: 369,
											EndLine:   369,
											StartPos:  6842,
											EndPos:    6842,
										},
										Val: &scalar.Lnumber{
											Position: &position.Position{
												StartLine: 369,
												EndLine:   369,
												StartPos:  6842,
												EndPos:    6842,
											},
											Value: "1",
										},
									},
									&expr.ArrayItem{
										Position: &position.Position{
											StartLine: 369,
											EndLine:   369,
											StartPos:  6845,
											EndPos:    6850,
										},
										Key: &scalar.Lnumber{
											Position: &position.Position{
												StartLine: 369,
												EndLine:   369,
												StartPos:  6845,
												EndPos:    6845,
											},
											Value: "2",
										},
										Val: &scalar.Lnumber{
											Position: &position.Position{
												StartLine: 369,
												EndLine:   369,
												StartPos:  6850,
												EndPos:    6850,
											},
											Value: "2",
										},
									},
								},
							},
							Dim: &scalar.Lnumber{
								Position: &position.Position{
									StartLine: 369,
									EndLine:   369,
									StartPos:  6853,
									EndPos:    6853,
								},
								Value: "0",
							},
						},
					},
				},
			},
			&stmt.If{
				Position: &position.Position{
					StartLine: 371,
					EndLine:   371,
					StartPos:  6860,
					EndPos:    6874,
				},
				Cond: &expr.Yield{
					Position: &position.Position{
						StartLine: 371,
						EndLine:   371,
						StartPos:  6864,
						EndPos:    6870,
					},
					Value: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 371,
							EndLine:   371,
							StartPos:  6870,
							EndPos:    6870,
						},
						Value: "1",
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 371,
						EndLine:   371,
						StartPos:  6873,
						EndPos:    6874,
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 372,
					EndLine:   372,
					StartPos:  6878,
					EndPos:    6888,
				},
				Expr: &expr.StaticPropertyFetch{
					Position: &position.Position{
						StartLine: 372,
						EndLine:   372,
						StartPos:  6878,
						EndPos:    6887,
					},
					Class: &name.Name{
						Position: &position.Position{
							StartLine: 372,
							EndLine:   372,
							StartPos:  6878,
							EndPos:    6880,
						},
						Parts: []node.Node{
							&name.NamePart{
								Position: &position.Position{
									StartLine: 372,
									EndLine:   372,
									StartPos:  6878,
									EndPos:    6880,
								},
								Value: "Foo",
							},
						},
					},
					Property: &expr.Variable{
						Position: &position.Position{
							StartLine: 372,
							EndLine:   372,
							StartPos:  6883,
							EndPos:    6887,
						},
						VarName: &expr.Variable{
							Position: &position.Position{
								StartLine: 372,
								EndLine:   372,
								StartPos:  6884,
								EndPos:    6887,
							},
							VarName: &node.Identifier{
								Position: &position.Position{
									StartLine: 372,
									EndLine:   372,
									StartPos:  6884,
									EndPos:    6887,
								},
								Value: "bar",
							},
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 374,
					EndLine:   374,
					StartPos:  6893,
					EndPos:    6899,
				},
				Expr: &expr.FunctionCall{
					Position: &position.Position{
						StartLine: 374,
						EndLine:   374,
						StartPos:  6893,
						EndPos:    6898,
					},
					Function: &expr.Variable{
						Position: &position.Position{
							StartLine: 374,
							EndLine:   374,
							StartPos:  6893,
							EndPos:    6896,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 374,
								EndLine:   374,
								StartPos:  6893,
								EndPos:    6896,
							},
							Value: "foo",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 374,
							EndLine:   374,
							StartPos:  6897,
							EndPos:    6898,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 375,
					EndLine:   375,
					StartPos:  6903,
					EndPos:    6915,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 375,
						EndLine:   375,
						StartPos:  6903,
						EndPos:    6914,
					},
					Variable: &expr.ArrayDimFetch{
						Position: &position.Position{
							StartLine: 375,
							EndLine:   375,
							StartPos:  6903,
							EndPos:    6911,
						},
						Variable: &expr.FunctionCall{
							Position: &position.Position{
								StartLine: 375,
								EndLine:   375,
								StartPos:  6903,
								EndPos:    6908,
							},
							Function: &expr.Variable{
								Position: &position.Position{
									StartLine: 375,
									EndLine:   375,
									StartPos:  6903,
									EndPos:    6906,
								},
								VarName: &node.Identifier{
									Position: &position.Position{
										StartLine: 375,
										EndLine:   375,
										StartPos:  6903,
										EndPos:    6906,
									},
									Value: "foo",
								},
							},
							ArgumentList: &node.ArgumentList{
								Position: &position.Position{
									StartLine: 375,
									EndLine:   375,
									StartPos:  6907,
									EndPos:    6908,
								},
							},
						},
						Dim: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 375,
								EndLine:   375,
								StartPos:  6910,
								EndPos:    6910,
							},
							Value: "0",
						},
					},
					Dim: &scalar.Lnumber{
						Position: &position.Position{
							StartLine: 375,
							EndLine:   375,
							StartPos:  6913,
							EndPos:    6913,
						},
						Value: "0",
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 376,
					EndLine:   376,
					StartPos:  6919,
					EndPos:    6925,
				},
				Expr: &expr.ArrayDimFetch{
					Position: &position.Position{
						StartLine: 376,
						EndLine:   376,
						StartPos:  6919,
						EndPos:    6924,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 376,
							EndLine:   376,
							StartPos:  6919,
							EndPos:    6920,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 376,
								EndLine:   376,
								StartPos:  6919,
								EndPos:    6920,
							},
							Value: "a",
						},
					},
					Dim: &expr.Variable{
						Position: &position.Position{
							StartLine: 376,
							EndLine:   376,
							StartPos:  6922,
							EndPos:    6923,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 376,
								EndLine:   376,
								StartPos:  6922,
								EndPos:    6923,
							},
							Value: "b",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 377,
					EndLine:   377,
					StartPos:  6929,
					EndPos:    6934,
				},
				Expr: &expr.Variable{
					Position: &position.Position{
						StartLine: 377,
						EndLine:   377,
						StartPos:  6929,
						EndPos:    6933,
					},
					VarName: &expr.Variable{
						Position: &position.Position{
							StartLine: 377,
							EndLine:   377,
							StartPos:  6931,
							EndPos:    6932,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 377,
								EndLine:   377,
								StartPos:  6931,
								EndPos:    6932,
							},
							Value: "a",
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 378,
					EndLine:   378,
					StartPos:  6938,
					EndPos:    6952,
				},
				Expr: &expr.StaticCall{
					Position: &position.Position{
						StartLine: 378,
						EndLine:   378,
						StartPos:  6938,
						EndPos:    6951,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 378,
							EndLine:   378,
							StartPos:  6938,
							EndPos:    6941,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 378,
								EndLine:   378,
								StartPos:  6938,
								EndPos:    6941,
							},
							Value: "foo",
						},
					},
					Call: &expr.Variable{
						Position: &position.Position{
							StartLine: 378,
							EndLine:   378,
							StartPos:  6944,
							EndPos:    6949,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 378,
								EndLine:   378,
								StartPos:  6945,
								EndPos:    6948,
							},
							Value: "bar",
						},
					},
					ArgumentList: &node.ArgumentList{
						Position: &position.Position{
							StartLine: 378,
							EndLine:   378,
							StartPos:  6950,
							EndPos:    6951,
						},
					},
				},
			},
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 379,
					EndLine:   379,
					StartPos:  6956,
					EndPos:    6965,
				},
				Expr: &expr.ClassConstFetch{
					Position: &position.Position{
						StartLine: 379,
						EndLine:   379,
						StartPos:  6956,
						EndPos:    6964,
					},
					Class: &expr.Variable{
						Position: &position.Position{
							StartLine: 379,
							EndLine:   379,
							StartPos:  6956,
							EndPos:    6959,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 379,
								EndLine:   379,
								StartPos:  6956,
								EndPos:    6959,
							},
							Value: "foo",
						},
					},
					ConstantName: &node.Identifier{
						Position: &position.Position{
							StartLine: 379,
							EndLine:   379,
							StartPos:  6962,
							EndPos:    6964,
						},
						Value: "bar",
					},
				},
			},
		},
	}

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual := php5parser.GetRootNode()
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

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual := php5parser.GetRootNode()
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

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual := php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}
