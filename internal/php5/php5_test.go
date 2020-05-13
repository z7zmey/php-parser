package php5_test

import (
	"gotest.tools/assert"
	"testing"

	"github.com/z7zmey/php-parser/internal/php5"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/position"
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

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   379,
				StartPos:  5,
				EndPos:    6944,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  5,
						EndPos:    20,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  5,
							EndPos:    19,
						},
					},
					Function: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  5,
								EndPos:    8,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  5,
										EndPos:    8,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  8,
								EndPos:    19,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  9,
										EndPos:    11,
									},
								},
								IsReference: false,
								Variadic:    false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  9,
											EndPos:    11,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  9,
												EndPos:    11,
											},
										},
										Value: []byte("a"),
									},
								},
							},
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  13,
										EndPos:    18,
									},
								},
								Variadic:    true,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  16,
											EndPos:    18,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  16,
												EndPos:    18,
											},
										},
										Value: []byte("b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  23,
						EndPos:    39,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  23,
							EndPos:    38,
						},
					},
					Function: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  23,
								EndPos:    27,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  23,
									EndPos:    27,
								},
							},
							Value: []byte("foo"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  27,
								EndPos:    38,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  28,
										EndPos:    30,
									},
								},
								Variadic:    false,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  28,
											EndPos:    30,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  28,
												EndPos:    30,
											},
										},
										Value: []byte("a"),
									},
								},
							},
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  32,
										EndPos:    37,
									},
								},
								Variadic:    true,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  35,
											EndPos:    37,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  35,
												EndPos:    37,
											},
										},
										Value: []byte("b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   4,
						StartPos:  42,
						EndPos:    63,
					},
				},
				Expr: &ast.ExprMethodCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  42,
							EndPos:    62,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  42,
								EndPos:    46,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  42,
									EndPos:    46,
								},
							},
							Value: []byte("foo"),
						},
					},
					Method: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  48,
								EndPos:    51,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  51,
								EndPos:    62,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  52,
										EndPos:    54,
									},
								},
								Variadic:    false,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  52,
											EndPos:    54,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 4,
												EndLine:   4,
												StartPos:  52,
												EndPos:    54,
											},
										},
										Value: []byte("a"),
									},
								},
							},
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  56,
										EndPos:    61,
									},
								},
								Variadic:    true,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  59,
											EndPos:    61,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 4,
												EndLine:   4,
												StartPos:  59,
												EndPos:    61,
											},
										},
										Value: []byte("b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 5,
						EndLine:   5,
						StartPos:  66,
						EndPos:    86,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  66,
							EndPos:    85,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  66,
								EndPos:    69,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  66,
										EndPos:    69,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  71,
								EndPos:    74,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  74,
								EndPos:    85,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  75,
										EndPos:    77,
									},
								},
								Variadic:    false,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  75,
											EndPos:    77,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 5,
												EndLine:   5,
												StartPos:  75,
												EndPos:    77,
											},
										},
										Value: []byte("a"),
									},
								},
							},
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 5,
										EndLine:   5,
										StartPos:  79,
										EndPos:    84,
									},
								},
								Variadic:    true,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 5,
											EndLine:   5,
											StartPos:  82,
											EndPos:    84,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 5,
												EndLine:   5,
												StartPos:  82,
												EndPos:    84,
											},
										},
										Value: []byte("b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 6,
						EndLine:   6,
						StartPos:  89,
						EndPos:    110,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  89,
							EndPos:    109,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 6,
								EndLine:   6,
								StartPos:  89,
								EndPos:    93,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 6,
									EndLine:   6,
									StartPos:  89,
									EndPos:    93,
								},
							},
							Value: []byte("foo"),
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 6,
								EndLine:   6,
								StartPos:  95,
								EndPos:    98,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 6,
								EndLine:   6,
								StartPos:  98,
								EndPos:    109,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 6,
										EndLine:   6,
										StartPos:  99,
										EndPos:    101,
									},
								},
								Variadic:    false,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 6,
											EndLine:   6,
											StartPos:  99,
											EndPos:    101,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 6,
												EndLine:   6,
												StartPos:  99,
												EndPos:    101,
											},
										},
										Value: []byte("a"),
									},
								},
							},
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 6,
										EndLine:   6,
										StartPos:  103,
										EndPos:    108,
									},
								},
								Variadic:    true,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 6,
											EndLine:   6,
											StartPos:  106,
											EndPos:    108,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 6,
												EndLine:   6,
												StartPos:  106,
												EndPos:    108,
											},
										},
										Value: []byte("b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   7,
						StartPos:  113,
						EndPos:    132,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   7,
							StartPos:  113,
							EndPos:    131,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 7,
								EndLine:   7,
								StartPos:  117,
								EndPos:    120,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  117,
										EndPos:    120,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 7,
								EndLine:   7,
								StartPos:  120,
								EndPos:    131,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  121,
										EndPos:    123,
									},
								},
								IsReference: false,
								Variadic:    false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 7,
											EndLine:   7,
											StartPos:  121,
											EndPos:    123,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 7,
												EndLine:   7,
												StartPos:  121,
												EndPos:    123,
											},
										},
										Value: []byte("a"),
									},
								},
							},
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 7,
										EndLine:   7,
										StartPos:  125,
										EndPos:    130,
									},
								},
								Variadic:    true,
								IsReference: false,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 7,
											EndLine:   7,
											StartPos:  128,
											EndPos:    130,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 7,
												EndLine:   7,
												StartPos:  128,
												EndPos:    130,
											},
										},
										Value: []byte("b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 9,
						EndLine:   9,
						StartPos:  136,
						EndPos:    180,
					},
				},
				ReturnsRef: false,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 9,
							EndLine:   9,
							StartPos:  145,
							EndPos:    148,
						},
					},
					Value: []byte("foo"),
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  149,
								EndPos:    162,
							},
						},
						ByRef:    false,
						Variadic: false,
						Type: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  149,
									EndPos:    152,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  149,
											EndPos:    152,
										},
									},
									Value: []byte("bar"),
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  153,
									EndPos:    157,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 9,
										EndLine:   9,
										StartPos:  153,
										EndPos:    157,
									},
								},
								Value: []byte("bar"),
							},
						},
						DefaultValue: &ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  158,
									EndPos:    162,
								},
							},
							Const: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 9,
										EndLine:   9,
										StartPos:  158,
										EndPos:    162,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  158,
												EndPos:    162,
											},
										},
										Value: []byte("null"),
									},
								},
							},
						},
					},
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  164,
								EndPos:    176,
							},
						},
						ByRef:    true,
						Variadic: true,
						Type: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  164,
									EndPos:    167,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  164,
											EndPos:    167,
										},
									},
									Value: []byte("baz"),
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  172,
									EndPos:    176,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 9,
										EndLine:   9,
										StartPos:  172,
										EndPos:    176,
									},
								},
								Value: []byte("baz"),
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 10,
						EndLine:   10,
						StartPos:  183,
						EndPos:    246,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 10,
							EndLine:   10,
							StartPos:  189,
							EndPos:    192,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 10,
								EndLine:   10,
								StartPos:  194,
								EndPos:    245,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 10,
									EndLine:   10,
									StartPos:  210,
									EndPos:    213,
								},
							},
							Value: []byte("foo"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 10,
										EndLine:   10,
										StartPos:  194,
										EndPos:    200,
									},
								},
								Value: []byte("public"),
							},
						},
						Params: []ast.Vertex{
							&ast.Parameter{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 10,
										EndLine:   10,
										StartPos:  214,
										EndPos:    227,
									},
								},
								ByRef:    false,
								Variadic: false,
								Type: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 10,
											EndLine:   10,
											StartPos:  214,
											EndPos:    217,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 10,
													EndLine:   10,
													StartPos:  214,
													EndPos:    217,
												},
											},
											Value: []byte("bar"),
										},
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 10,
											EndLine:   10,
											StartPos:  218,
											EndPos:    222,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 10,
												EndLine:   10,
												StartPos:  218,
												EndPos:    222,
											},
										},
										Value: []byte("bar"),
									},
								},
								DefaultValue: &ast.ExprConstFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 10,
											EndLine:   10,
											StartPos:  223,
											EndPos:    227,
										},
									},
									Const: &ast.NameName{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 10,
												EndLine:   10,
												StartPos:  223,
												EndPos:    227,
											},
										},
										Parts: []ast.Vertex{
											&ast.NameNamePart{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 10,
														EndLine:   10,
														StartPos:  223,
														EndPos:    227,
													},
												},
												Value: []byte("null"),
											},
										},
									},
								},
							},
							&ast.Parameter{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 10,
										EndLine:   10,
										StartPos:  229,
										EndPos:    241,
									},
								},
								ByRef:    true,
								Variadic: true,
								Type: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 10,
											EndLine:   10,
											StartPos:  229,
											EndPos:    232,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 10,
													EndLine:   10,
													StartPos:  229,
													EndPos:    232,
												},
											},
											Value: []byte("baz"),
										},
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 10,
											EndLine:   10,
											StartPos:  237,
											EndPos:    241,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 10,
												EndLine:   10,
												StartPos:  237,
												EndPos:    241,
											},
										},
										Value: []byte("baz"),
									},
								},
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 10,
									EndLine:   10,
									StartPos:  243,
									EndPos:    245,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 11,
						EndLine:   11,
						StartPos:  249,
						EndPos:    290,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 11,
							EndLine:   11,
							StartPos:  249,
							EndPos:    289,
						},
					},
					ReturnsRef: false,
					Static:     false,
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  258,
									EndPos:    271,
								},
							},
							ByRef:    false,
							Variadic: false,
							Type: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  258,
										EndPos:    261,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 11,
												EndLine:   11,
												StartPos:  258,
												EndPos:    261,
											},
										},
										Value: []byte("bar"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  262,
										EndPos:    266,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 11,
											EndLine:   11,
											StartPos:  262,
											EndPos:    266,
										},
									},
									Value: []byte("bar"),
								},
							},
							DefaultValue: &ast.ExprConstFetch{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  267,
										EndPos:    271,
									},
								},
								Const: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 11,
											EndLine:   11,
											StartPos:  267,
											EndPos:    271,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 11,
													EndLine:   11,
													StartPos:  267,
													EndPos:    271,
												},
											},
											Value: []byte("null"),
										},
									},
								},
							},
						},
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  273,
									EndPos:    285,
								},
							},
							ByRef:    true,
							Variadic: true,
							Type: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  273,
										EndPos:    276,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 11,
												EndLine:   11,
												StartPos:  273,
												EndPos:    276,
											},
										},
										Value: []byte("baz"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  281,
										EndPos:    285,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 11,
											EndLine:   11,
											StartPos:  281,
											EndPos:    285,
										},
									},
									Value: []byte("baz"),
								},
							},
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 12,
						EndLine:   12,
						StartPos:  293,
						EndPos:    341,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 12,
							EndLine:   12,
							StartPos:  293,
							EndPos:    340,
						},
					},
					ReturnsRef: false,
					Static:     true,
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 12,
									EndLine:   12,
									StartPos:  309,
									EndPos:    322,
								},
							},
							ByRef:    false,
							Variadic: false,
							Type: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  309,
										EndPos:    312,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 12,
												EndLine:   12,
												StartPos:  309,
												EndPos:    312,
											},
										},
										Value: []byte("bar"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  313,
										EndPos:    317,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 12,
											EndLine:   12,
											StartPos:  313,
											EndPos:    317,
										},
									},
									Value: []byte("bar"),
								},
							},
							DefaultValue: &ast.ExprConstFetch{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  318,
										EndPos:    322,
									},
								},
								Const: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 12,
											EndLine:   12,
											StartPos:  318,
											EndPos:    322,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 12,
													EndLine:   12,
													StartPos:  318,
													EndPos:    322,
												},
											},
											Value: []byte("null"),
										},
									},
								},
							},
						},
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 12,
									EndLine:   12,
									StartPos:  324,
									EndPos:    336,
								},
							},
							ByRef:    true,
							Variadic: true,
							Type: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  324,
										EndPos:    327,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 12,
												EndLine:   12,
												StartPos:  324,
												EndPos:    327,
											},
										},
										Value: []byte("baz"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  332,
										EndPos:    336,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 12,
											EndLine:   12,
											StartPos:  332,
											EndPos:    336,
										},
									},
									Value: []byte("baz"),
								},
							},
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 14,
						EndLine:   14,
						StartPos:  345,
						EndPos:    365,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 14,
							EndLine:   14,
							StartPos:  345,
							EndPos:    364,
						},
					},
					Value: []byte("1234567890123456789"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 15,
						EndLine:   15,
						StartPos:  368,
						EndPos:    389,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 15,
							EndLine:   15,
							StartPos:  368,
							EndPos:    388,
						},
					},
					Value: []byte("12345678901234567890"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 16,
						EndLine:   16,
						StartPos:  392,
						EndPos:    395,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 16,
							EndLine:   16,
							StartPos:  392,
							EndPos:    394,
						},
					},
					Value: []byte("0."),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 17,
						EndLine:   17,
						StartPos:  398,
						EndPos:    465,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 17,
							EndLine:   17,
							StartPos:  398,
							EndPos:    464,
						},
					},
					Value: []byte("0b0111111111111111111111111111111111111111111111111111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 18,
						EndLine:   18,
						StartPos:  468,
						EndPos:    535,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 18,
							EndLine:   18,
							StartPos:  468,
							EndPos:    534,
						},
					},
					Value: []byte("0b1111111111111111111111111111111111111111111111111111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 19,
						EndLine:   19,
						StartPos:  538,
						EndPos:    559,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 19,
							EndLine:   19,
							StartPos:  538,
							EndPos:    558,
						},
					},
					Value: []byte("0x007111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 20,
						EndLine:   20,
						StartPos:  562,
						EndPos:    581,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 20,
							EndLine:   20,
							StartPos:  562,
							EndPos:    580,
						},
					},
					Value: []byte("0x8111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 21,
						EndLine:   21,
						StartPos:  584,
						EndPos:    594,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 21,
							EndLine:   21,
							StartPos:  584,
							EndPos:    593,
						},
					},
					Value: []byte("__CLASS__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 22,
						EndLine:   22,
						StartPos:  597,
						EndPos:    605,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 22,
							EndLine:   22,
							StartPos:  597,
							EndPos:    604,
						},
					},
					Value: []byte("__DIR__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 23,
						EndLine:   23,
						StartPos:  608,
						EndPos:    617,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 23,
							EndLine:   23,
							StartPos:  608,
							EndPos:    616,
						},
					},
					Value: []byte("__FILE__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 24,
						EndLine:   24,
						StartPos:  620,
						EndPos:    633,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 24,
							EndLine:   24,
							StartPos:  620,
							EndPos:    632,
						},
					},
					Value: []byte("__FUNCTION__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 25,
						EndLine:   25,
						StartPos:  636,
						EndPos:    645,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 25,
							EndLine:   25,
							StartPos:  636,
							EndPos:    644,
						},
					},
					Value: []byte("__LINE__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 26,
						EndLine:   26,
						StartPos:  648,
						EndPos:    662,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 26,
							EndLine:   26,
							StartPos:  648,
							EndPos:    661,
						},
					},
					Value: []byte("__NAMESPACE__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 27,
						EndLine:   27,
						StartPos:  665,
						EndPos:    676,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 27,
							EndLine:   27,
							StartPos:  665,
							EndPos:    675,
						},
					},
					Value: []byte("__METHOD__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 28,
						EndLine:   28,
						StartPos:  679,
						EndPos:    689,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 28,
							EndLine:   28,
							StartPos:  679,
							EndPos:    688,
						},
					},
					Value: []byte("__TRAIT__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 30,
						EndLine:   30,
						StartPos:  693,
						EndPos:    705,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 30,
							EndLine:   30,
							StartPos:  693,
							EndPos:    704,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 30,
									EndLine:   30,
									StartPos:  694,
									EndPos:    699,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 30,
									EndLine:   30,
									StartPos:  699,
									EndPos:    703,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 30,
										EndLine:   30,
										StartPos:  699,
										EndPos:    703,
									},
								},
								Value: []byte("var"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 31,
						EndLine:   31,
						StartPos:  708,
						EndPos:    723,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 31,
							EndLine:   31,
							StartPos:  708,
							EndPos:    722,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 31,
									EndLine:   31,
									StartPos:  709,
									EndPos:    714,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 31,
									EndLine:   31,
									StartPos:  714,
									EndPos:    721,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 31,
										EndLine:   31,
										StartPos:  714,
										EndPos:    718,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 31,
											EndLine:   31,
											StartPos:  714,
											EndPos:    718,
										},
									},
									Value: []byte("var"),
								},
							},
							Dim: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 31,
										EndLine:   31,
										StartPos:  719,
										EndPos:    720,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 32,
						EndLine:   32,
						StartPos:  726,
						EndPos:    780,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 32,
							EndLine:   32,
							StartPos:  726,
							EndPos:    779,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 32,
									EndLine:   32,
									StartPos:  727,
									EndPos:    732,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 32,
									EndLine:   32,
									StartPos:  732,
									EndPos:    778,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 32,
										EndLine:   32,
										StartPos:  732,
										EndPos:    736,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 32,
											EndLine:   32,
											StartPos:  732,
											EndPos:    736,
										},
									},
									Value: []byte("var"),
								},
							},
							Dim: &ast.ScalarString{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 32,
										EndLine:   32,
										StartPos:  737,
										EndPos:    777,
									},
								},
								Value: []byte("1234567890123456789012345678901234567890"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 33,
						EndLine:   33,
						StartPos:  783,
						EndPos:    800,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 33,
							EndLine:   33,
							StartPos:  783,
							EndPos:    799,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 33,
									EndLine:   33,
									StartPos:  784,
									EndPos:    789,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 33,
									EndLine:   33,
									StartPos:  789,
									EndPos:    798,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 33,
										EndLine:   33,
										StartPos:  789,
										EndPos:    793,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 33,
											EndLine:   33,
											StartPos:  789,
											EndPos:    793,
										},
									},
									Value: []byte("var"),
								},
							},
							Dim: &ast.ScalarString{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 33,
										EndLine:   33,
										StartPos:  794,
										EndPos:    797,
									},
								},
								Value: []byte("bar"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 34,
						EndLine:   34,
						StartPos:  803,
						EndPos:    821,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 34,
							EndLine:   34,
							StartPos:  803,
							EndPos:    820,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 34,
									EndLine:   34,
									StartPos:  804,
									EndPos:    809,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 34,
									EndLine:   34,
									StartPos:  809,
									EndPos:    819,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 34,
										EndLine:   34,
										StartPos:  809,
										EndPos:    813,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 34,
											EndLine:   34,
											StartPos:  809,
											EndPos:    813,
										},
									},
									Value: []byte("var"),
								},
							},
							Dim: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 34,
										EndLine:   34,
										StartPos:  814,
										EndPos:    818,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 34,
											EndLine:   34,
											StartPos:  814,
											EndPos:    818,
										},
									},
									Value: []byte("bar"),
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 35,
						EndLine:   35,
						StartPos:  824,
						EndPos:    836,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 35,
							EndLine:   35,
							StartPos:  824,
							EndPos:    835,
						},
					},
					Parts: []ast.Vertex{
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 35,
									EndLine:   35,
									StartPos:  825,
									EndPos:    829,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 35,
										EndLine:   35,
										StartPos:  825,
										EndPos:    829,
									},
								},
								Value: []byte("foo"),
							},
						},
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 35,
									EndLine:   35,
									StartPos:  829,
									EndPos:    830,
								},
							},
							Value: []byte(" "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 35,
									EndLine:   35,
									StartPos:  830,
									EndPos:    834,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 35,
										EndLine:   35,
										StartPos:  830,
										EndPos:    834,
									},
								},
								Value: []byte("bar"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 36,
						EndLine:   36,
						StartPos:  839,
						EndPos:    858,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 36,
							EndLine:   36,
							StartPos:  839,
							EndPos:    857,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 36,
									EndLine:   36,
									StartPos:  840,
									EndPos:    845,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprPropertyFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 36,
									EndLine:   36,
									StartPos:  845,
									EndPos:    854,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 36,
										EndLine:   36,
										StartPos:  845,
										EndPos:    849,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 36,
											EndLine:   36,
											StartPos:  845,
											EndPos:    849,
										},
									},
									Value: []byte("foo"),
								},
							},
							Property: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 36,
										EndLine:   36,
										StartPos:  851,
										EndPos:    854,
									},
								},
								Value: []byte("bar"),
							},
						},
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 36,
									EndLine:   36,
									StartPos:  854,
									EndPos:    856,
								},
							},
							Value: []byte("()"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 37,
						EndLine:   37,
						StartPos:  861,
						EndPos:    875,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 37,
							EndLine:   37,
							StartPos:  861,
							EndPos:    874,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 37,
									EndLine:   37,
									StartPos:  862,
									EndPos:    867,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 37,
									EndLine:   37,
									StartPos:  867,
									EndPos:    873,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 37,
										EndLine:   37,
										StartPos:  869,
										EndPos:    872,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 38,
						EndLine:   38,
						StartPos:  878,
						EndPos:    895,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 38,
							EndLine:   38,
							StartPos:  878,
							EndPos:    894,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 38,
									EndLine:   38,
									StartPos:  879,
									EndPos:    884,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 38,
									EndLine:   38,
									StartPos:  884,
									EndPos:    893,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 38,
										EndLine:   38,
										StartPos:  886,
										EndPos:    889,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 38,
											EndLine:   38,
											StartPos:  886,
											EndPos:    889,
										},
									},
									Value: []byte("foo"),
								},
							},
							Dim: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 38,
										EndLine:   38,
										StartPos:  890,
										EndPos:    891,
									},
								},
								Value: []byte("0"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 39,
						EndLine:   39,
						StartPos:  898,
						EndPos:    919,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 39,
							EndLine:   39,
							StartPos:  898,
							EndPos:    918,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 39,
									EndLine:   39,
									StartPos:  899,
									EndPos:    904,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprMethodCall{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 39,
									EndLine:   39,
									StartPos:  905,
									EndPos:    916,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 39,
										EndLine:   39,
										StartPos:  905,
										EndPos:    909,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 39,
											EndLine:   39,
											StartPos:  905,
											EndPos:    909,
										},
									},
									Value: []byte("foo"),
								},
							},
							Method: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 39,
										EndLine:   39,
										StartPos:  911,
										EndPos:    914,
									},
								},
								Value: []byte("bar"),
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 39,
										EndLine:   39,
										StartPos:  914,
										EndPos:    916,
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtAltIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 41,
						EndLine:   42,
						StartPos:  923,
						EndPos:    941,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 41,
							EndLine:   41,
							StartPos:  927,
							EndPos:    929,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 41,
								EndLine:   41,
								StartPos:  927,
								EndPos:    929,
							},
						},
						Value: []byte("a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtAltIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 43,
						EndLine:   45,
						StartPos:  944,
						EndPos:    977,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 43,
							EndLine:   43,
							StartPos:  948,
							EndPos:    950,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 43,
								EndLine:   43,
								StartPos:  948,
								EndPos:    950,
							},
						},
						Value: []byte("a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtAltElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 44,
								EndLine:   -1,
								StartPos:  956,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 44,
									EndLine:   44,
									StartPos:  964,
									EndPos:    966,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 44,
										EndLine:   44,
										StartPos:  964,
										EndPos:    966,
									},
								},
								Value: []byte("b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: -1,
									EndLine:   -1,
									StartPos:  -1,
									EndPos:    -1,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtAltIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 46,
						EndLine:   48,
						StartPos:  980,
						EndPos:    1006,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 46,
							EndLine:   46,
							StartPos:  984,
							EndPos:    986,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 46,
								EndLine:   46,
								StartPos:  984,
								EndPos:    986,
							},
						},
						Value: []byte("a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
					},
					Stmts: []ast.Vertex{},
				},
				Else: &ast.StmtAltElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 47,
							EndLine:   -1,
							StartPos:  992,
							EndPos:    -1,
						},
					},
					Stmt: &ast.StmtStmtList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: -1,
								EndLine:   -1,
								StartPos:  -1,
								EndPos:    -1,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtAltIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 49,
						EndLine:   53,
						StartPos:  1009,
						EndPos:    1065,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 49,
							EndLine:   49,
							StartPos:  1013,
							EndPos:    1015,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 49,
								EndLine:   49,
								StartPos:  1013,
								EndPos:    1015,
							},
						},
						Value: []byte("a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtAltElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 50,
								EndLine:   -1,
								StartPos:  1021,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 50,
									EndLine:   50,
									StartPos:  1029,
									EndPos:    1031,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 50,
										EndLine:   50,
										StartPos:  1029,
										EndPos:    1031,
									},
								},
								Value: []byte("b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: -1,
									EndLine:   -1,
									StartPos:  -1,
									EndPos:    -1,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
					&ast.StmtAltElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 51,
								EndLine:   -1,
								StartPos:  1036,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 51,
									EndLine:   51,
									StartPos:  1044,
									EndPos:    1046,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 51,
										EndLine:   51,
										StartPos:  1044,
										EndPos:    1046,
									},
								},
								Value: []byte("c"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: -1,
									EndLine:   -1,
									StartPos:  -1,
									EndPos:    -1,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
				Else: &ast.StmtAltElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 52,
							EndLine:   -1,
							StartPos:  1051,
							EndPos:    -1,
						},
					},
					Stmt: &ast.StmtStmtList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: -1,
								EndLine:   -1,
								StartPos:  -1,
								EndPos:    -1,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 55,
						EndLine:   55,
						StartPos:  1069,
						EndPos:    1089,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 55,
							EndLine:   55,
							StartPos:  1076,
							EndPos:    1077,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 55,
							EndLine:   55,
							StartPos:  1079,
							EndPos:    1089,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 55,
									EndLine:   55,
									StartPos:  1081,
									EndPos:    1087,
								},
							},
						},
					},
				},
			},
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 56,
						EndLine:   56,
						StartPos:  1092,
						EndPos:    1114,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 56,
							EndLine:   56,
							StartPos:  1099,
							EndPos:    1100,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 56,
							EndLine:   56,
							StartPos:  1102,
							EndPos:    1114,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 56,
									EndLine:   56,
									StartPos:  1104,
									EndPos:    1112,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 56,
										EndLine:   56,
										StartPos:  1110,
										EndPos:    1111,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtAltWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 57,
						EndLine:   57,
						StartPos:  1117,
						EndPos:    1148,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 57,
							EndLine:   57,
							StartPos:  1124,
							EndPos:    1125,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 57,
							EndLine:   57,
							StartPos:  1129,
							EndPos:    1138,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 57,
									EndLine:   57,
									StartPos:  1129,
									EndPos:    1138,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 57,
										EndLine:   57,
										StartPos:  1135,
										EndPos:    1136,
									},
								},
								Value: []byte("3"),
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 58,
						EndLine:   58,
						StartPos:  1151,
						EndPos:    1187,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 58,
							EndLine:   58,
							StartPos:  1157,
							EndPos:    1160,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassConstList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 58,
								EndLine:   58,
								StartPos:  1162,
								EndPos:    1185,
							},
						},
						Consts: []ast.Vertex{
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 58,
										EndLine:   58,
										StartPos:  1168,
										EndPos:    1175,
									},
								},
								ConstantName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 58,
											EndLine:   58,
											StartPos:  1168,
											EndPos:    1171,
										},
									},
									Value: []byte("FOO"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 58,
											EndLine:   58,
											StartPos:  1174,
											EndPos:    1175,
										},
									},
									Value: []byte("1"),
								},
							},
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 58,
										EndLine:   58,
										StartPos:  1177,
										EndPos:    1184,
									},
								},
								ConstantName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 58,
											EndLine:   58,
											StartPos:  1177,
											EndPos:    1180,
										},
									},
									Value: []byte("BAR"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 58,
											EndLine:   58,
											StartPos:  1183,
											EndPos:    1184,
										},
									},
									Value: []byte("2"),
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 59,
						EndLine:   59,
						StartPos:  1190,
						EndPos:    1220,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 59,
							EndLine:   59,
							StartPos:  1196,
							EndPos:    1199,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 59,
								EndLine:   59,
								StartPos:  1201,
								EndPos:    1218,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 59,
									EndLine:   59,
									StartPos:  1210,
									EndPos:    1213,
								},
							},
							Value: []byte("bar"),
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 59,
									EndLine:   59,
									StartPos:  1216,
									EndPos:    1218,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 60,
						EndLine:   60,
						StartPos:  1223,
						EndPos:    1268,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 60,
							EndLine:   60,
							StartPos:  1229,
							EndPos:    1232,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 60,
								EndLine:   60,
								StartPos:  1234,
								EndPos:    1266,
							},
						},
						ReturnsRef: true,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 60,
									EndLine:   60,
									StartPos:  1258,
									EndPos:    1261,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 60,
										EndLine:   60,
										StartPos:  1234,
										EndPos:    1240,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 60,
										EndLine:   60,
										StartPos:  1241,
										EndPos:    1247,
									},
								},
								Value: []byte("static"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 60,
									EndLine:   60,
									StartPos:  1264,
									EndPos:    1266,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 61,
						EndLine:   61,
						StartPos:  1271,
						EndPos:    1343,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 61,
							EndLine:   61,
							StartPos:  1277,
							EndPos:    1280,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 61,
								EndLine:   61,
								StartPos:  1282,
								EndPos:    1313,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 61,
									EndLine:   61,
									StartPos:  1305,
									EndPos:    1308,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 61,
										EndLine:   61,
										StartPos:  1282,
										EndPos:    1287,
									},
								},
								Value: []byte("final"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 61,
										EndLine:   61,
										StartPos:  1288,
										EndPos:    1295,
									},
								},
								Value: []byte("private"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 61,
									EndLine:   61,
									StartPos:  1311,
									EndPos:    1313,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 61,
								EndLine:   61,
								StartPos:  1314,
								EndPos:    1341,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 61,
									EndLine:   61,
									StartPos:  1333,
									EndPos:    1336,
								},
							},
							Value: []byte("baz"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 61,
										EndLine:   61,
										StartPos:  1314,
										EndPos:    1323,
									},
								},
								Value: []byte("protected"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 61,
									EndLine:   61,
									StartPos:  1339,
									EndPos:    1341,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 62,
						EndLine:   62,
						StartPos:  1346,
						EndPos:    1399,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 62,
							EndLine:   62,
							StartPos:  1361,
							EndPos:    1364,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 62,
								EndLine:   62,
								StartPos:  1346,
								EndPos:    1354,
							},
						},
						Value: []byte("abstract"),
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 62,
								EndLine:   62,
								StartPos:  1366,
								EndPos:    1397,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 62,
									EndLine:   62,
									StartPos:  1391,
									EndPos:    1394,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 62,
										EndLine:   62,
										StartPos:  1366,
										EndPos:    1374,
									},
								},
								Value: []byte("abstract"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 62,
										EndLine:   62,
										StartPos:  1375,
										EndPos:    1381,
									},
								},
								Value: []byte("public"),
							},
						},
						Stmt: &ast.StmtNop{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 62,
									EndLine:   62,
									StartPos:  1396,
									EndPos:    1397,
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 63,
						EndLine:   63,
						StartPos:  1402,
						EndPos:    1433,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 63,
							EndLine:   63,
							StartPos:  1414,
							EndPos:    1417,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 63,
								EndLine:   63,
								StartPos:  1402,
								EndPos:    1407,
							},
						},
						Value: []byte("final"),
					},
				},
				Extends: &ast.StmtClassExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 63,
							EndLine:   63,
							StartPos:  1418,
							EndPos:    1429,
						},
					},
					ClassName: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 63,
								EndLine:   63,
								StartPos:  1426,
								EndPos:    1429,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 63,
										EndLine:   63,
										StartPos:  1426,
										EndPos:    1429,
									},
								},
								Value: []byte("bar"),
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 64,
						EndLine:   64,
						StartPos:  1436,
						EndPos:    1470,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 64,
							EndLine:   64,
							StartPos:  1448,
							EndPos:    1451,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 64,
								EndLine:   64,
								StartPos:  1436,
								EndPos:    1441,
							},
						},
						Value: []byte("final"),
					},
				},
				Implements: &ast.StmtClassImplements{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 64,
							EndLine:   64,
							StartPos:  1452,
							EndPos:    1466,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 64,
									EndLine:   64,
									StartPos:  1463,
									EndPos:    1466,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 64,
											EndLine:   64,
											StartPos:  1463,
											EndPos:    1466,
										},
									},
									Value: []byte("bar"),
								},
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 65,
						EndLine:   65,
						StartPos:  1473,
						EndPos:    1512,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 65,
							EndLine:   65,
							StartPos:  1485,
							EndPos:    1488,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 65,
								EndLine:   65,
								StartPos:  1473,
								EndPos:    1478,
							},
						},
						Value: []byte("final"),
					},
				},
				Implements: &ast.StmtClassImplements{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 65,
							EndLine:   65,
							StartPos:  1489,
							EndPos:    1508,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 65,
									EndLine:   65,
									StartPos:  1500,
									EndPos:    1503,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 65,
											EndLine:   65,
											StartPos:  1500,
											EndPos:    1503,
										},
									},
									Value: []byte("bar"),
								},
							},
						},
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 65,
									EndLine:   65,
									StartPos:  1505,
									EndPos:    1508,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 65,
											EndLine:   65,
											StartPos:  1505,
											EndPos:    1508,
										},
									},
									Value: []byte("baz"),
								},
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtConstList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 67,
						EndLine:   67,
						StartPos:  1516,
						EndPos:    1539,
					},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 67,
								EndLine:   67,
								StartPos:  1522,
								EndPos:    1529,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 67,
									EndLine:   67,
									StartPos:  1522,
									EndPos:    1525,
								},
							},
							Value: []byte("FOO"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 67,
									EndLine:   67,
									StartPos:  1528,
									EndPos:    1529,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 67,
								EndLine:   67,
								StartPos:  1531,
								EndPos:    1538,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 67,
									EndLine:   67,
									StartPos:  1531,
									EndPos:    1534,
								},
							},
							Value: []byte("BAR"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 67,
									EndLine:   67,
									StartPos:  1537,
									EndPos:    1538,
								},
							},
							Value: []byte("2"),
						},
					},
				},
			},
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 68,
						EndLine:   68,
						StartPos:  1542,
						EndPos:    1565,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 68,
							EndLine:   68,
							StartPos:  1549,
							EndPos:    1550,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 68,
							EndLine:   68,
							StartPos:  1552,
							EndPos:    1565,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 68,
									EndLine:   68,
									StartPos:  1554,
									EndPos:    1563,
								},
							},
						},
					},
				},
			},
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 69,
						EndLine:   69,
						StartPos:  1568,
						EndPos:    1593,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 69,
							EndLine:   69,
							StartPos:  1575,
							EndPos:    1576,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 69,
							EndLine:   69,
							StartPos:  1578,
							EndPos:    1593,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 69,
									EndLine:   69,
									StartPos:  1580,
									EndPos:    1591,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 69,
										EndLine:   69,
										StartPos:  1589,
										EndPos:    1590,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 70,
						EndLine:   70,
						StartPos:  1596,
						EndPos:    1622,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 70,
							EndLine:   70,
							StartPos:  1603,
							EndPos:    1604,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 70,
							EndLine:   70,
							StartPos:  1606,
							EndPos:    1622,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 70,
									EndLine:   70,
									StartPos:  1608,
									EndPos:    1620,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1617,
										EndPos:    1618,
									},
								},
								Value: []byte("3"),
							},
						},
					},
				},
			},
			&ast.StmtDeclare{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 71,
						EndLine:   71,
						StartPos:  1625,
						EndPos:    1642,
					},
				},
				Alt: false,
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 71,
								EndLine:   71,
								StartPos:  1633,
								EndPos:    1640,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 71,
									EndLine:   71,
									StartPos:  1633,
									EndPos:    1638,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 71,
									EndLine:   71,
									StartPos:  1639,
									EndPos:    1640,
								},
							},
							Value: []byte("1"),
						},
					},
				},
				Stmt: &ast.StmtNop{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 71,
							EndLine:   71,
							StartPos:  1641,
							EndPos:    1642,
						},
					},
				},
			},
			&ast.StmtDeclare{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 72,
						EndLine:   72,
						StartPos:  1645,
						EndPos:    1680,
					},
				},
				Alt: false,
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 72,
								EndLine:   72,
								StartPos:  1653,
								EndPos:    1660,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 72,
									EndLine:   72,
									StartPos:  1653,
									EndPos:    1658,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 72,
									EndLine:   72,
									StartPos:  1659,
									EndPos:    1660,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 72,
								EndLine:   72,
								StartPos:  1662,
								EndPos:    1676,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 72,
									EndLine:   72,
									StartPos:  1662,
									EndPos:    1674,
								},
							},
							Value: []byte("strict_types"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 72,
									EndLine:   72,
									StartPos:  1675,
									EndPos:    1676,
								},
							},
							Value: []byte("1"),
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 72,
							EndLine:   72,
							StartPos:  1678,
							EndPos:    1680,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtDeclare{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 73,
						EndLine:   73,
						StartPos:  1683,
						EndPos:    1712,
					},
				},
				Alt: true,
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 73,
								EndLine:   73,
								StartPos:  1691,
								EndPos:    1698,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 73,
									EndLine:   73,
									StartPos:  1691,
									EndPos:    1696,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 73,
									EndLine:   73,
									StartPos:  1697,
									EndPos:    1698,
								},
							},
							Value: []byte("1"),
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtDo{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 74,
						EndLine:   74,
						StartPos:  1715,
						EndPos:    1730,
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 74,
							EndLine:   74,
							StartPos:  1718,
							EndPos:    1720,
						},
					},
					Stmts: []ast.Vertex{},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 74,
							EndLine:   74,
							StartPos:  1727,
							EndPos:    1728,
						},
					},
					Value: []byte("1"),
				},
			},
			&ast.StmtEcho{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 75,
						EndLine:   75,
						StartPos:  1733,
						EndPos:    1744,
					},
				},
				Exprs: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 75,
								EndLine:   75,
								StartPos:  1738,
								EndPos:    1740,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 75,
									EndLine:   75,
									StartPos:  1738,
									EndPos:    1740,
								},
							},
							Value: []byte("a"),
						},
					},
					&ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 75,
								EndLine:   75,
								StartPos:  1742,
								EndPos:    1743,
							},
						},
						Value: []byte("1"),
					},
				},
			},
			&ast.StmtEcho{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 76,
						EndLine:   76,
						StartPos:  1747,
						EndPos:    1756,
					},
				},
				Exprs: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 76,
								EndLine:   76,
								StartPos:  1752,
								EndPos:    1754,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 76,
									EndLine:   76,
									StartPos:  1752,
									EndPos:    1754,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtFor{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 77,
						EndLine:   77,
						StartPos:  1759,
						EndPos:    1794,
					},
				},
				Init: []ast.Vertex{
					&ast.ExprAssign{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1763,
								EndPos:    1769,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1763,
									EndPos:    1765,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 77,
										EndLine:   77,
										StartPos:  1763,
										EndPos:    1765,
									},
								},
								Value: []byte("i"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1768,
									EndPos:    1769,
								},
							},
							Value: []byte("0"),
						},
					},
				},
				Cond: []ast.Vertex{
					&ast.ExprBinarySmaller{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1771,
								EndPos:    1778,
							},
						},
						Left: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1771,
									EndPos:    1773,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 77,
										EndLine:   77,
										StartPos:  1771,
										EndPos:    1773,
									},
								},
								Value: []byte("i"),
							},
						},
						Right: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1776,
									EndPos:    1778,
								},
							},
							Value: []byte("10"),
						},
					},
				},
				Loop: []ast.Vertex{
					&ast.ExprPostInc{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1780,
								EndPos:    1784,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1780,
									EndPos:    1782,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 77,
										EndLine:   77,
										StartPos:  1780,
										EndPos:    1782,
									},
								},
								Value: []byte("i"),
							},
						},
					},
					&ast.ExprPostInc{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1786,
								EndPos:    1790,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1786,
									EndPos:    1788,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 77,
										EndLine:   77,
										StartPos:  1786,
										EndPos:    1788,
									},
								},
								Value: []byte("i"),
							},
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 77,
							EndLine:   77,
							StartPos:  1792,
							EndPos:    1794,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtAltFor{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 78,
						EndLine:   78,
						StartPos:  1797,
						EndPos:    1827,
					},
				},
				Cond: []ast.Vertex{
					&ast.ExprBinarySmaller{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 78,
								EndLine:   78,
								StartPos:  1803,
								EndPos:    1810,
							},
						},
						Left: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 78,
									EndLine:   78,
									StartPos:  1803,
									EndPos:    1805,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 78,
										EndLine:   78,
										StartPos:  1803,
										EndPos:    1805,
									},
								},
								Value: []byte("i"),
							},
						},
						Right: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 78,
									EndLine:   78,
									StartPos:  1808,
									EndPos:    1810,
								},
							},
							Value: []byte("10"),
						},
					},
				},
				Loop: []ast.Vertex{
					&ast.ExprPostInc{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 78,
								EndLine:   78,
								StartPos:  1812,
								EndPos:    1816,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 78,
									EndLine:   78,
									StartPos:  1812,
									EndPos:    1814,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 78,
										EndLine:   78,
										StartPos:  1812,
										EndPos:    1814,
									},
								},
								Value: []byte("i"),
							},
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 79,
						EndLine:   79,
						StartPos:  1830,
						EndPos:    1851,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 79,
							EndLine:   79,
							StartPos:  1839,
							EndPos:    1841,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 79,
								EndLine:   79,
								StartPos:  1839,
								EndPos:    1841,
							},
						},
						Value: []byte("a"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 79,
							EndLine:   79,
							StartPos:  1845,
							EndPos:    1847,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 79,
								EndLine:   79,
								StartPos:  1845,
								EndPos:    1847,
							},
						},
						Value: []byte("v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 79,
							EndLine:   79,
							StartPos:  1849,
							EndPos:    1851,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 80,
						EndLine:   80,
						StartPos:  1854,
						EndPos:    1875,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 80,
							EndLine:   80,
							StartPos:  1863,
							EndPos:    1865,
						},
					},
					Items: []ast.Vertex{},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 80,
							EndLine:   80,
							StartPos:  1869,
							EndPos:    1871,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 80,
								EndLine:   80,
								StartPos:  1869,
								EndPos:    1871,
							},
						},
						Value: []byte("v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 80,
							EndLine:   80,
							StartPos:  1873,
							EndPos:    1875,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtAltForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 81,
						EndLine:   81,
						StartPos:  1878,
						EndPos:    1910,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 81,
							EndLine:   81,
							StartPos:  1887,
							EndPos:    1889,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1887,
								EndPos:    1889,
							},
						},
						Value: []byte("a"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 81,
							EndLine:   81,
							StartPos:  1893,
							EndPos:    1895,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1893,
								EndPos:    1895,
							},
						},
						Value: []byte("v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: -1,
							EndLine:   -1,
							StartPos:  -1,
							EndPos:    -1,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 82,
						EndLine:   82,
						StartPos:  1913,
						EndPos:    1940,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1922,
							EndPos:    1924,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 82,
								EndLine:   82,
								StartPos:  1922,
								EndPos:    1924,
							},
						},
						Value: []byte("a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1928,
							EndPos:    1930,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 82,
								EndLine:   82,
								StartPos:  1928,
								EndPos:    1930,
							},
						},
						Value: []byte("k"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1934,
							EndPos:    1936,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 82,
								EndLine:   82,
								StartPos:  1934,
								EndPos:    1936,
							},
						},
						Value: []byte("v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1938,
							EndPos:    1940,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 83,
						EndLine:   83,
						StartPos:  1943,
						EndPos:    1970,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 83,
							EndLine:   83,
							StartPos:  1952,
							EndPos:    1954,
						},
					},
					Items: []ast.Vertex{},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 83,
							EndLine:   83,
							StartPos:  1958,
							EndPos:    1960,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 83,
								EndLine:   83,
								StartPos:  1958,
								EndPos:    1960,
							},
						},
						Value: []byte("k"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 83,
							EndLine:   83,
							StartPos:  1964,
							EndPos:    1966,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 83,
								EndLine:   83,
								StartPos:  1964,
								EndPos:    1966,
							},
						},
						Value: []byte("v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 83,
							EndLine:   83,
							StartPos:  1968,
							EndPos:    1970,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 84,
						EndLine:   84,
						StartPos:  1973,
						EndPos:    2001,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1982,
							EndPos:    1984,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 84,
								EndLine:   84,
								StartPos:  1982,
								EndPos:    1984,
							},
						},
						Value: []byte("a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1988,
							EndPos:    1990,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 84,
								EndLine:   84,
								StartPos:  1988,
								EndPos:    1990,
							},
						},
						Value: []byte("k"),
					},
				},
				Var: &ast.ExprReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1994,
							EndPos:    1997,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 84,
								EndLine:   84,
								StartPos:  1995,
								EndPos:    1997,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 84,
									EndLine:   84,
									StartPos:  1995,
									EndPos:    1997,
								},
							},
							Value: []byte("v"),
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1999,
							EndPos:    2001,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 85,
						EndLine:   85,
						StartPos:  2004,
						EndPos:    2037,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 85,
							EndLine:   85,
							StartPos:  2013,
							EndPos:    2015,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 85,
								EndLine:   85,
								StartPos:  2013,
								EndPos:    2015,
							},
						},
						Value: []byte("a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 85,
							EndLine:   85,
							StartPos:  2019,
							EndPos:    2021,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 85,
								EndLine:   85,
								StartPos:  2019,
								EndPos:    2021,
							},
						},
						Value: []byte("k"),
					},
				},
				Var: &ast.ExprList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 85,
							EndLine:   85,
							StartPos:  2025,
							EndPos:    2033,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 85,
									EndLine:   85,
									StartPos:  2030,
									EndPos:    2032,
								},
							},
							Val: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 85,
										EndLine:   85,
										StartPos:  2030,
										EndPos:    2032,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 85,
											EndLine:   85,
											StartPos:  2030,
											EndPos:    2032,
										},
									},
									Value: []byte("v"),
								},
							},
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 85,
							EndLine:   85,
							StartPos:  2035,
							EndPos:    2037,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 86,
						EndLine:   86,
						StartPos:  2040,
						EndPos:    2057,
					},
				},
				ReturnsRef: false,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 86,
							EndLine:   86,
							StartPos:  2049,
							EndPos:    2052,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 88,
						EndLine:   92,
						StartPos:  2061,
						EndPos:    2132,
					},
				},
				ReturnsRef: false,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 88,
							EndLine:   88,
							StartPos:  2070,
							EndPos:    2073,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtFunction{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 89,
								EndLine:   89,
								StartPos:  2081,
								EndPos:    2098,
							},
						},
						ReturnsRef: false,
						FunctionName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 89,
									EndLine:   89,
									StartPos:  2090,
									EndPos:    2093,
								},
							},
							Value: []byte("bar"),
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 90,
								EndLine:   90,
								StartPos:  2102,
								EndPos:    2114,
							},
						},
						ClassName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 90,
									EndLine:   90,
									StartPos:  2108,
									EndPos:    2111,
								},
							},
							Value: []byte("Baz"),
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtReturn{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2118,
								EndPos:    2128,
							},
						},
						Expr: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2125,
									EndPos:    2127,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 91,
										EndLine:   91,
										StartPos:  2125,
										EndPos:    2127,
									},
								},
								Value: []byte("a"),
							},
						},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 94,
						EndLine:   94,
						StartPos:  2138,
						EndPos:    2183,
					},
				},
				ReturnsRef: false,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 94,
							EndLine:   94,
							StartPos:  2147,
							EndPos:    2150,
						},
					},
					Value: []byte("foo"),
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 94,
								EndLine:   94,
								StartPos:  2151,
								EndPos:    2159,
							},
						},
						Variadic: false,
						ByRef:    false,
						Type: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 94,
									EndLine:   94,
									StartPos:  2151,
									EndPos:    2156,
								},
							},
							Value: []byte("array"),
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 94,
									EndLine:   94,
									StartPos:  2157,
									EndPos:    2159,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 94,
										EndLine:   94,
										StartPos:  2157,
										EndPos:    2159,
									},
								},
								Value: []byte("a"),
							},
						},
					},
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 94,
								EndLine:   94,
								StartPos:  2161,
								EndPos:    2172,
							},
						},
						ByRef:    false,
						Variadic: false,
						Type: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 94,
									EndLine:   94,
									StartPos:  2161,
									EndPos:    2169,
								},
							},
							Value: []byte("callable"),
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 94,
									EndLine:   94,
									StartPos:  2170,
									EndPos:    2172,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 94,
										EndLine:   94,
										StartPos:  2170,
										EndPos:    2172,
									},
								},
								Value: []byte("b"),
							},
						},
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtReturn{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 94,
								EndLine:   94,
								StartPos:  2175,
								EndPos:    2182,
							},
						},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 95,
						EndLine:   95,
						StartPos:  2186,
						EndPos:    2213,
					},
				},
				ReturnsRef: true,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2196,
							EndPos:    2199,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtReturn{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 95,
								EndLine:   95,
								StartPos:  2203,
								EndPos:    2212,
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 95,
									EndLine:   95,
									StartPos:  2210,
									EndPos:    2211,
								},
							},
							Value: []byte("1"),
						},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 96,
						EndLine:   96,
						StartPos:  2216,
						EndPos:    2234,
					},
				},
				ReturnsRef: true,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2226,
							EndPos:    2229,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtGlobal{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 97,
						EndLine:   97,
						StartPos:  2237,
						EndPos:    2266,
					},
				},
				Vars: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2244,
								EndPos:    2246,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2244,
									EndPos:    2246,
								},
							},
							Value: []byte("a"),
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2248,
								EndPos:    2250,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2248,
									EndPos:    2250,
								},
							},
							Value: []byte("b"),
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2252,
								EndPos:    2255,
							},
						},
						VarName: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2253,
									EndPos:    2255,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 97,
										EndLine:   97,
										StartPos:  2253,
										EndPos:    2255,
									},
								},
								Value: []byte("c"),
							},
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2257,
								EndPos:    2265,
							},
						},
						VarName: &ast.ExprFunctionCall{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2259,
									EndPos:    2264,
								},
							},
							Function: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 97,
										EndLine:   97,
										StartPos:  2259,
										EndPos:    2262,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 97,
												EndLine:   97,
												StartPos:  2259,
												EndPos:    2262,
											},
										},
										Value: []byte("foo"),
									},
								},
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 97,
										EndLine:   97,
										StartPos:  2262,
										EndPos:    2264,
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtLabel{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 98,
						EndLine:   98,
						StartPos:  2269,
						EndPos:    2271,
					},
				},
				LabelName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2269,
							EndPos:    2270,
						},
					},
					Value: []byte("a"),
				},
			},
			&ast.StmtGoto{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 99,
						EndLine:   99,
						StartPos:  2275,
						EndPos:    2282,
					},
				},
				Label: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 99,
							EndLine:   99,
							StartPos:  2280,
							EndPos:    2281,
						},
					},
					Value: []byte("a"),
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 100,
						EndLine:   100,
						StartPos:  2285,
						EndPos:    2295,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 100,
							EndLine:   100,
							StartPos:  2289,
							EndPos:    2291,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 100,
								EndLine:   100,
								StartPos:  2289,
								EndPos:    2291,
							},
						},
						Value: []byte("a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 100,
							EndLine:   100,
							StartPos:  2293,
							EndPos:    2295,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 101,
						EndLine:   101,
						StartPos:  2298,
						EndPos:    2323,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 101,
							EndLine:   101,
							StartPos:  2302,
							EndPos:    2304,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 101,
								EndLine:   101,
								StartPos:  2302,
								EndPos:    2304,
							},
						},
						Value: []byte("a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 101,
							EndLine:   101,
							StartPos:  2306,
							EndPos:    2308,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 101,
								EndLine:   101,
								StartPos:  2309,
								EndPos:    2323,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 101,
									EndLine:   101,
									StartPos:  2317,
									EndPos:    2319,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 101,
										EndLine:   101,
										StartPos:  2317,
										EndPos:    2319,
									},
								},
								Value: []byte("b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 101,
									EndLine:   101,
									StartPos:  2321,
									EndPos:    2323,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 102,
						EndLine:   102,
						StartPos:  2326,
						EndPos:    2344,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 102,
							EndLine:   102,
							StartPos:  2330,
							EndPos:    2332,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 102,
								EndLine:   102,
								StartPos:  2330,
								EndPos:    2332,
							},
						},
						Value: []byte("a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 102,
							EndLine:   102,
							StartPos:  2334,
							EndPos:    2336,
						},
					},
					Stmts: []ast.Vertex{},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 102,
							EndLine:   102,
							StartPos:  2337,
							EndPos:    2344,
						},
					},
					Stmt: &ast.StmtStmtList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 102,
								EndLine:   102,
								StartPos:  2342,
								EndPos:    2344,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 103,
						EndLine:   103,
						StartPos:  2347,
						EndPos:    2395,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2351,
							EndPos:    2353,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2351,
								EndPos:    2353,
							},
						},
						Value: []byte("a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2355,
							EndPos:    2357,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2358,
								EndPos:    2372,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2366,
									EndPos:    2368,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 103,
										EndLine:   103,
										StartPos:  2366,
										EndPos:    2368,
									},
								},
								Value: []byte("b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2370,
									EndPos:    2372,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2373,
								EndPos:    2387,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2381,
									EndPos:    2383,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 103,
										EndLine:   103,
										StartPos:  2381,
										EndPos:    2383,
									},
								},
								Value: []byte("c"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2385,
									EndPos:    2387,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 103,
							EndLine:   103,
							StartPos:  2388,
							EndPos:    2395,
						},
					},
					Stmt: &ast.StmtStmtList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2393,
								EndPos:    2395,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 104,
						EndLine:   104,
						StartPos:  2398,
						EndPos:    2447,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 104,
							EndLine:   104,
							StartPos:  2402,
							EndPos:    2404,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 104,
								EndLine:   104,
								StartPos:  2402,
								EndPos:    2404,
							},
						},
						Value: []byte("a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 104,
							EndLine:   104,
							StartPos:  2406,
							EndPos:    2408,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 104,
								EndLine:   104,
								StartPos:  2409,
								EndPos:    2423,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 104,
									EndLine:   104,
									StartPos:  2417,
									EndPos:    2419,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 104,
										EndLine:   104,
										StartPos:  2417,
										EndPos:    2419,
									},
								},
								Value: []byte("b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 104,
									EndLine:   104,
									StartPos:  2421,
									EndPos:    2423,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 104,
							EndLine:   104,
							StartPos:  2424,
							EndPos:    2447,
						},
					},
					Stmt: &ast.StmtIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 104,
								EndLine:   104,
								StartPos:  2429,
								EndPos:    2447,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 104,
									EndLine:   104,
									StartPos:  2433,
									EndPos:    2435,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 104,
										EndLine:   104,
										StartPos:  2433,
										EndPos:    2435,
									},
								},
								Value: []byte("c"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 104,
									EndLine:   104,
									StartPos:  2437,
									EndPos:    2439,
								},
							},
							Stmts: []ast.Vertex{},
						},
						Else: &ast.StmtElse{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 104,
									EndLine:   104,
									StartPos:  2440,
									EndPos:    2447,
								},
							},
							Stmt: &ast.StmtStmtList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 104,
										EndLine:   104,
										StartPos:  2445,
										EndPos:    2447,
									},
								},
								Stmts: []ast.Vertex{},
							},
						},
					},
				},
			},
			&ast.StmtNop{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 105,
						EndLine:   105,
						StartPos:  2450,
						EndPos:    2452,
					},
				},
			},
			&ast.StmtInlineHtml{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 105,
						EndLine:   105,
						StartPos:  2452,
						EndPos:    2465,
					},
				},
				Value: []byte(" <div></div> "),
			},
			&ast.StmtInterface{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 106,
						EndLine:   106,
						StartPos:  2470,
						EndPos:    2486,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 106,
							EndLine:   106,
							StartPos:  2480,
							EndPos:    2483,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtInterface{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 107,
						EndLine:   107,
						StartPos:  2489,
						EndPos:    2517,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 107,
							EndLine:   107,
							StartPos:  2499,
							EndPos:    2502,
						},
					},
					Value: []byte("Foo"),
				},
				Extends: &ast.StmtInterfaceExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 107,
							EndLine:   107,
							StartPos:  2503,
							EndPos:    2514,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 107,
									EndLine:   107,
									StartPos:  2511,
									EndPos:    2514,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 107,
											EndLine:   107,
											StartPos:  2511,
											EndPos:    2514,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtInterface{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 108,
						EndLine:   108,
						StartPos:  2520,
						EndPos:    2553,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 108,
							EndLine:   108,
							StartPos:  2530,
							EndPos:    2533,
						},
					},
					Value: []byte("Foo"),
				},
				Extends: &ast.StmtInterfaceExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 108,
							EndLine:   108,
							StartPos:  2534,
							EndPos:    2550,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 108,
									EndLine:   108,
									StartPos:  2542,
									EndPos:    2545,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 108,
											EndLine:   108,
											StartPos:  2542,
											EndPos:    2545,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 108,
									EndLine:   108,
									StartPos:  2547,
									EndPos:    2550,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 108,
											EndLine:   108,
											StartPos:  2547,
											EndPos:    2550,
										},
									},
									Value: []byte("Baz"),
								},
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtNamespace{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 109,
						EndLine:   109,
						StartPos:  2556,
						EndPos:    2570,
					},
				},
				NamespaceName: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 109,
							EndLine:   109,
							StartPos:  2566,
							EndPos:    2569,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 109,
									EndLine:   109,
									StartPos:  2566,
									EndPos:    2569,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
			},
			&ast.StmtNamespace{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 110,
						EndLine:   110,
						StartPos:  2573,
						EndPos:    2593,
					},
				},
				NamespaceName: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2583,
							EndPos:    2590,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2583,
									EndPos:    2586,
								},
							},
							Value: []byte("Foo"),
						},
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2587,
									EndPos:    2590,
								},
							},
							Value: []byte("Bar"),
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtNamespace{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 111,
						EndLine:   111,
						StartPos:  2596,
						EndPos:    2608,
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 112,
						EndLine:   112,
						StartPos:  2611,
						EndPos:    2630,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 112,
							EndLine:   112,
							StartPos:  2617,
							EndPos:    2620,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtPropertyList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 112,
								EndLine:   112,
								StartPos:  2622,
								EndPos:    2629,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 112,
										EndLine:   112,
										StartPos:  2622,
										EndPos:    2625,
									},
								},
								Value: []byte("var"),
							},
						},
						Properties: []ast.Vertex{
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 112,
										EndLine:   112,
										StartPos:  2626,
										EndPos:    2628,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 112,
											EndLine:   112,
											StartPos:  2626,
											EndPos:    2628,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 112,
												EndLine:   112,
												StartPos:  2626,
												EndPos:    2628,
											},
										},
										Value: []byte("a"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 113,
						EndLine:   113,
						StartPos:  2633,
						EndPos:    2670,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 113,
							EndLine:   113,
							StartPos:  2639,
							EndPos:    2642,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtPropertyList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 113,
								EndLine:   113,
								StartPos:  2644,
								EndPos:    2669,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 113,
										EndLine:   113,
										StartPos:  2644,
										EndPos:    2650,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 113,
										EndLine:   113,
										StartPos:  2651,
										EndPos:    2657,
									},
								},
								Value: []byte("static"),
							},
						},
						Properties: []ast.Vertex{
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 113,
										EndLine:   113,
										StartPos:  2658,
										EndPos:    2660,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 113,
											EndLine:   113,
											StartPos:  2658,
											EndPos:    2660,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 113,
												EndLine:   113,
												StartPos:  2658,
												EndPos:    2660,
											},
										},
										Value: []byte("a"),
									},
								},
							},
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 113,
										EndLine:   113,
										StartPos:  2662,
										EndPos:    2668,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 113,
											EndLine:   113,
											StartPos:  2662,
											EndPos:    2664,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 113,
												EndLine:   113,
												StartPos:  2662,
												EndPos:    2664,
											},
										},
										Value: []byte("b"),
									},
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 113,
											EndLine:   113,
											StartPos:  2667,
											EndPos:    2668,
										},
									},
									Value: []byte("1"),
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 114,
						EndLine:   114,
						StartPos:  2673,
						EndPos:    2710,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 114,
							EndLine:   114,
							StartPos:  2679,
							EndPos:    2682,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtPropertyList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 114,
								EndLine:   114,
								StartPos:  2684,
								EndPos:    2709,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2684,
										EndPos:    2690,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2691,
										EndPos:    2697,
									},
								},
								Value: []byte("static"),
							},
						},
						Properties: []ast.Vertex{
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2698,
										EndPos:    2704,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 114,
											EndLine:   114,
											StartPos:  2698,
											EndPos:    2700,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 114,
												EndLine:   114,
												StartPos:  2698,
												EndPos:    2700,
											},
										},
										Value: []byte("a"),
									},
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 114,
											EndLine:   114,
											StartPos:  2703,
											EndPos:    2704,
										},
									},
									Value: []byte("1"),
								},
							},
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 114,
										EndLine:   114,
										StartPos:  2706,
										EndPos:    2708,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 114,
											EndLine:   114,
											StartPos:  2706,
											EndPos:    2708,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 114,
												EndLine:   114,
												StartPos:  2706,
												EndPos:    2708,
											},
										},
										Value: []byte("b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 115,
						EndLine:   115,
						StartPos:  2713,
						EndPos:    2731,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 115,
								EndLine:   115,
								StartPos:  2720,
								EndPos:    2722,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 115,
									EndLine:   115,
									StartPos:  2720,
									EndPos:    2722,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 115,
										EndLine:   115,
										StartPos:  2720,
										EndPos:    2722,
									},
								},
								Value: []byte("a"),
							},
						},
					},
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 115,
								EndLine:   115,
								StartPos:  2724,
								EndPos:    2730,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 115,
									EndLine:   115,
									StartPos:  2724,
									EndPos:    2726,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 115,
										EndLine:   115,
										StartPos:  2724,
										EndPos:    2726,
									},
								},
								Value: []byte("b"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 115,
									EndLine:   115,
									StartPos:  2729,
									EndPos:    2730,
								},
							},
							Value: []byte("1"),
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 116,
						EndLine:   116,
						StartPos:  2734,
						EndPos:    2752,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 116,
								EndLine:   116,
								StartPos:  2741,
								EndPos:    2747,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 116,
									EndLine:   116,
									StartPos:  2741,
									EndPos:    2743,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 116,
										EndLine:   116,
										StartPos:  2741,
										EndPos:    2743,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 116,
									EndLine:   116,
									StartPos:  2746,
									EndPos:    2747,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 116,
								EndLine:   116,
								StartPos:  2749,
								EndPos:    2751,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 116,
									EndLine:   116,
									StartPos:  2749,
									EndPos:    2751,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 116,
										EndLine:   116,
										StartPos:  2749,
										EndPos:    2751,
									},
								},
								Value: []byte("b"),
							},
						},
					},
				},
			},
			&ast.StmtAltSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 118,
						EndLine:   122,
						StartPos:  2756,
						EndPos:    2815,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 118,
							EndLine:   118,
							StartPos:  2764,
							EndPos:    2765,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 119,
							EndLine:   -1,
							StartPos:  2772,
							EndPos:    -1,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 119,
									EndLine:   -1,
									StartPos:  2772,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 119,
										EndLine:   119,
										StartPos:  2777,
										EndPos:    2778,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtDefault{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 120,
									EndLine:   -1,
									StartPos:  2783,
									EndPos:    -1,
								},
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 121,
									EndLine:   -1,
									StartPos:  2795,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 121,
										EndLine:   121,
										StartPos:  2800,
										EndPos:    2801,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtAltSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 124,
						EndLine:   127,
						StartPos:  2819,
						EndPos:    2867,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 124,
							EndLine:   124,
							StartPos:  2827,
							EndPos:    2828,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 125,
							EndLine:   -1,
							StartPos:  2836,
							EndPos:    -1,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 125,
									EndLine:   -1,
									StartPos:  2836,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 125,
										EndLine:   125,
										StartPos:  2841,
										EndPos:    2842,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 126,
									EndLine:   -1,
									StartPos:  2847,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 126,
										EndLine:   126,
										StartPos:  2852,
										EndPos:    2853,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
			},
			&ast.StmtSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 129,
						EndLine:   132,
						StartPos:  2873,
						EndPos:    2925,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 129,
							EndLine:   129,
							StartPos:  2881,
							EndPos:    2882,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 129,
							EndLine:   132,
							StartPos:  2884,
							EndPos:    2925,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 130,
									EndLine:   130,
									StartPos:  2889,
									EndPos:    2903,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 130,
										EndLine:   130,
										StartPos:  2894,
										EndPos:    2895,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 130,
											EndLine:   130,
											StartPos:  2897,
											EndPos:    2903,
										},
									},
								},
							},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 131,
									EndLine:   131,
									StartPos:  2907,
									EndPos:    2921,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 131,
										EndLine:   131,
										StartPos:  2912,
										EndPos:    2913,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 131,
											EndLine:   131,
											StartPos:  2915,
											EndPos:    2921,
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 134,
						EndLine:   137,
						StartPos:  2931,
						EndPos:    2984,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 134,
							EndLine:   134,
							StartPos:  2939,
							EndPos:    2940,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 134,
							EndLine:   137,
							StartPos:  2942,
							EndPos:    2984,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 135,
									EndLine:   135,
									StartPos:  2948,
									EndPos:    2962,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 135,
										EndLine:   135,
										StartPos:  2953,
										EndPos:    2954,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 135,
											EndLine:   135,
											StartPos:  2956,
											EndPos:    2962,
										},
									},
								},
							},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 136,
									EndLine:   136,
									StartPos:  2966,
									EndPos:    2980,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 136,
										EndLine:   136,
										StartPos:  2971,
										EndPos:    2972,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 136,
											EndLine:   136,
											StartPos:  2974,
											EndPos:    2980,
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtThrow{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 138,
						EndLine:   138,
						StartPos:  2987,
						EndPos:    2996,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 138,
							EndLine:   138,
							StartPos:  2993,
							EndPos:    2995,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 138,
								EndLine:   138,
								StartPos:  2993,
								EndPos:    2995,
							},
						},
						Value: []byte("e"),
					},
				},
			},
			&ast.StmtTrait{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 139,
						EndLine:   139,
						StartPos:  2999,
						EndPos:    3011,
					},
				},
				TraitName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 139,
							EndLine:   139,
							StartPos:  3005,
							EndPos:    3008,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 140,
						EndLine:   140,
						StartPos:  3014,
						EndPos:    3036,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 140,
							EndLine:   140,
							StartPos:  3020,
							EndPos:    3023,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 140,
								EndLine:   140,
								StartPos:  3026,
								EndPos:    3034,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 140,
										EndLine:   140,
										StartPos:  3030,
										EndPos:    3033,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 140,
												EndLine:   140,
												StartPos:  3030,
												EndPos:    3033,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
						},
						TraitAdaptationList: &ast.StmtNop{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 140,
									EndLine:   140,
									StartPos:  3033,
									EndPos:    3034,
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 141,
						EndLine:   141,
						StartPos:  3039,
						EndPos:    3068,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 141,
							EndLine:   141,
							StartPos:  3045,
							EndPos:    3048,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 141,
								EndLine:   141,
								StartPos:  3051,
								EndPos:    3066,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 141,
										EndLine:   141,
										StartPos:  3055,
										EndPos:    3058,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 141,
												EndLine:   141,
												StartPos:  3055,
												EndPos:    3058,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 141,
										EndLine:   141,
										StartPos:  3060,
										EndPos:    3063,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 141,
												EndLine:   141,
												StartPos:  3060,
												EndPos:    3063,
											},
										},
										Value: []byte("Baz"),
									},
								},
							},
						},
						TraitAdaptationList: &ast.StmtTraitAdaptationList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 141,
									EndLine:   141,
									StartPos:  3064,
									EndPos:    3066,
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 142,
						EndLine:   142,
						StartPos:  3071,
						EndPos:    3116,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 142,
							EndLine:   142,
							StartPos:  3077,
							EndPos:    3080,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 142,
								EndLine:   142,
								StartPos:  3083,
								EndPos:    3114,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 142,
										EndLine:   142,
										StartPos:  3087,
										EndPos:    3090,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 142,
												EndLine:   142,
												StartPos:  3087,
												EndPos:    3090,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 142,
										EndLine:   142,
										StartPos:  3092,
										EndPos:    3095,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 142,
												EndLine:   142,
												StartPos:  3092,
												EndPos:    3095,
											},
										},
										Value: []byte("Baz"),
									},
								},
							},
						},
						TraitAdaptationList: &ast.StmtTraitAdaptationList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 142,
									EndLine:   142,
									StartPos:  3096,
									EndPos:    3114,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUseAlias{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 142,
											EndLine:   142,
											StartPos:  3098,
											EndPos:    3111,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 142,
												EndLine:   142,
												StartPos:  3098,
												EndPos:    3101,
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 142,
													EndLine:   142,
													StartPos:  3098,
													EndPos:    3101,
												},
											},
											Value: []byte("one"),
										},
									},
									Modifier: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 142,
												EndLine:   142,
												StartPos:  3105,
												EndPos:    3111,
											},
										},
										Value: []byte("public"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 143,
						EndLine:   143,
						StartPos:  3119,
						EndPos:    3168,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 143,
							EndLine:   143,
							StartPos:  3125,
							EndPos:    3128,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 143,
								EndLine:   143,
								StartPos:  3131,
								EndPos:    3166,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 143,
										EndLine:   143,
										StartPos:  3135,
										EndPos:    3138,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 143,
												EndLine:   143,
												StartPos:  3135,
												EndPos:    3138,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 143,
										EndLine:   143,
										StartPos:  3140,
										EndPos:    3143,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 143,
												EndLine:   143,
												StartPos:  3140,
												EndPos:    3143,
											},
										},
										Value: []byte("Baz"),
									},
								},
							},
						},
						TraitAdaptationList: &ast.StmtTraitAdaptationList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 143,
									EndLine:   143,
									StartPos:  3144,
									EndPos:    3166,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUseAlias{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 143,
											EndLine:   143,
											StartPos:  3146,
											EndPos:    3163,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 143,
												EndLine:   143,
												StartPos:  3146,
												EndPos:    3149,
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 143,
													EndLine:   143,
													StartPos:  3146,
													EndPos:    3149,
												},
											},
											Value: []byte("one"),
										},
									},
									Modifier: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 143,
												EndLine:   143,
												StartPos:  3153,
												EndPos:    3159,
											},
										},
										Value: []byte("public"),
									},
									Alias: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 143,
												EndLine:   143,
												StartPos:  3160,
												EndPos:    3163,
											},
										},
										Value: []byte("two"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 144,
						EndLine:   144,
						StartPos:  3171,
						EndPos:    3248,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 144,
							EndLine:   144,
							StartPos:  3177,
							EndPos:    3180,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 144,
								EndLine:   144,
								StartPos:  3183,
								EndPos:    3246,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 144,
										EndLine:   144,
										StartPos:  3187,
										EndPos:    3190,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  3187,
												EndPos:    3190,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 144,
										EndLine:   144,
										StartPos:  3192,
										EndPos:    3195,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  3192,
												EndPos:    3195,
											},
										},
										Value: []byte("Baz"),
									},
								},
							},
						},
						TraitAdaptationList: &ast.StmtTraitAdaptationList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 144,
									EndLine:   144,
									StartPos:  3196,
									EndPos:    3246,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUsePrecedence{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 144,
											EndLine:   144,
											StartPos:  3198,
											EndPos:    3226,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  3198,
												EndPos:    3206,
											},
										},
										Trait: &ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 144,
													EndLine:   144,
													StartPos:  3198,
													EndPos:    3201,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 144,
															EndLine:   144,
															StartPos:  3198,
															EndPos:    3201,
														},
													},
													Value: []byte("Bar"),
												},
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 144,
													EndLine:   144,
													StartPos:  3203,
													EndPos:    3206,
												},
											},
											Value: []byte("one"),
										},
									},
									Insteadof: []ast.Vertex{
										&ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 144,
													EndLine:   144,
													StartPos:  3217,
													EndPos:    3220,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 144,
															EndLine:   144,
															StartPos:  3217,
															EndPos:    3220,
														},
													},
													Value: []byte("Baz"),
												},
											},
										},
										&ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 144,
													EndLine:   144,
													StartPos:  3222,
													EndPos:    3226,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 144,
															EndLine:   144,
															StartPos:  3222,
															EndPos:    3226,
														},
													},
													Value: []byte("Quux"),
												},
											},
										},
									},
								},
								&ast.StmtTraitUseAlias{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 144,
											EndLine:   144,
											StartPos:  3228,
											EndPos:    3243,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  3228,
												EndPos:    3236,
											},
										},
										Trait: &ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 144,
													EndLine:   144,
													StartPos:  3228,
													EndPos:    3231,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 144,
															EndLine:   144,
															StartPos:  3228,
															EndPos:    3231,
														},
													},
													Value: []byte("Baz"),
												},
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 144,
													EndLine:   144,
													StartPos:  3233,
													EndPos:    3236,
												},
											},
											Value: []byte("one"),
										},
									},
									Alias: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 144,
												EndLine:   144,
												StartPos:  3240,
												EndPos:    3243,
											},
										},
										Value: []byte("two"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 146,
						EndLine:   -1,
						StartPos:  3252,
						EndPos:    -1,
					},
				},
				Stmts:   []ast.Vertex{},
				Catches: []ast.Vertex{},
			},
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 147,
						EndLine:   147,
						StartPos:  3261,
						EndPos:    3291,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 147,
								EndLine:   147,
								StartPos:  3268,
								EndPos:    3291,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 147,
										EndLine:   147,
										StartPos:  3275,
										EndPos:    3284,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 147,
												EndLine:   147,
												StartPos:  3275,
												EndPos:    3284,
											},
										},
										Value: []byte("Exception"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 147,
									EndLine:   147,
									StartPos:  3285,
									EndPos:    3287,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 147,
										EndLine:   147,
										StartPos:  3285,
										EndPos:    3287,
									},
								},
								Value: []byte("e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 148,
						EndLine:   148,
						StartPos:  3294,
						EndPos:    3355,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 148,
								EndLine:   148,
								StartPos:  3301,
								EndPos:    3324,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 148,
										EndLine:   148,
										StartPos:  3308,
										EndPos:    3317,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 148,
												EndLine:   148,
												StartPos:  3308,
												EndPos:    3317,
											},
										},
										Value: []byte("Exception"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 148,
									EndLine:   148,
									StartPos:  3318,
									EndPos:    3320,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 148,
										EndLine:   148,
										StartPos:  3318,
										EndPos:    3320,
									},
								},
								Value: []byte("e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 148,
								EndLine:   148,
								StartPos:  3325,
								EndPos:    3355,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 148,
										EndLine:   148,
										StartPos:  3332,
										EndPos:    3348,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 148,
												EndLine:   148,
												StartPos:  3332,
												EndPos:    3348,
											},
										},
										Value: []byte("RuntimeException"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 148,
									EndLine:   148,
									StartPos:  3349,
									EndPos:    3351,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 148,
										EndLine:   148,
										StartPos:  3349,
										EndPos:    3351,
									},
								},
								Value: []byte("e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 149,
						EndLine:   149,
						StartPos:  3358,
						EndPos:    3462,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 149,
								EndLine:   149,
								StartPos:  3365,
								EndPos:    3388,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3372,
										EndPos:    3381,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3372,
												EndPos:    3381,
											},
										},
										Value: []byte("Exception"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3382,
									EndPos:    3384,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3382,
										EndPos:    3384,
									},
								},
								Value: []byte("e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 149,
								EndLine:   149,
								StartPos:  3389,
								EndPos:    3420,
							},
						},
						Types: []ast.Vertex{
							&ast.NameFullyQualified{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3396,
										EndPos:    3413,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3397,
												EndPos:    3413,
											},
										},
										Value: []byte("RuntimeException"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3414,
									EndPos:    3416,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3414,
										EndPos:    3416,
									},
								},
								Value: []byte("e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 149,
								EndLine:   149,
								StartPos:  3421,
								EndPos:    3462,
							},
						},
						Types: []ast.Vertex{
							&ast.NameRelative{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3428,
										EndPos:    3455,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3438,
												EndPos:    3455,
											},
										},
										Value: []byte("AdditionException"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 149,
									EndLine:   149,
									StartPos:  3456,
									EndPos:    3458,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3456,
										EndPos:    3458,
									},
								},
								Value: []byte("e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 150,
						EndLine:   150,
						StartPos:  3465,
						EndPos:    3506,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 150,
								EndLine:   150,
								StartPos:  3472,
								EndPos:    3495,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 150,
										EndLine:   150,
										StartPos:  3479,
										EndPos:    3488,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 150,
												EndLine:   150,
												StartPos:  3479,
												EndPos:    3488,
											},
										},
										Value: []byte("Exception"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 150,
									EndLine:   150,
									StartPos:  3489,
									EndPos:    3491,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 150,
										EndLine:   150,
										StartPos:  3489,
										EndPos:    3491,
									},
								},
								Value: []byte("e"),
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
				Finally: &ast.StmtFinally{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 150,
							EndLine:   150,
							StartPos:  3496,
							EndPos:    3506,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtUnset{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 152,
						EndLine:   152,
						StartPos:  3510,
						EndPos:    3524,
					},
				},
				Vars: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 152,
								EndLine:   152,
								StartPos:  3516,
								EndPos:    3518,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 152,
									EndLine:   152,
									StartPos:  3516,
									EndPos:    3518,
								},
							},
							Value: []byte("a"),
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 152,
								EndLine:   152,
								StartPos:  3520,
								EndPos:    3522,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 152,
									EndLine:   152,
									StartPos:  3520,
									EndPos:    3522,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 154,
						EndLine:   154,
						StartPos:  3528,
						EndPos:    3536,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 154,
								EndLine:   154,
								StartPos:  3532,
								EndPos:    3535,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 154,
									EndLine:   154,
									StartPos:  3532,
									EndPos:    3535,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 154,
											EndLine:   154,
											StartPos:  3532,
											EndPos:    3535,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
				},
			},
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 155,
						EndLine:   155,
						StartPos:  3539,
						EndPos:    3548,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 155,
								EndLine:   155,
								StartPos:  3544,
								EndPos:    3547,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 155,
									EndLine:   155,
									StartPos:  3544,
									EndPos:    3547,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 155,
											EndLine:   155,
											StartPos:  3544,
											EndPos:    3547,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
				},
			},
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 156,
						EndLine:   156,
						StartPos:  3551,
						EndPos:    3567,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 156,
								EndLine:   156,
								StartPos:  3556,
								EndPos:    3566,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 156,
									EndLine:   156,
									StartPos:  3556,
									EndPos:    3559,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 156,
											EndLine:   156,
											StartPos:  3556,
											EndPos:    3559,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 156,
									EndLine:   156,
									StartPos:  3563,
									EndPos:    3566,
								},
							},
							Value: []byte("Bar"),
						},
					},
				},
			},
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 157,
						EndLine:   157,
						StartPos:  3570,
						EndPos:    3583,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 157,
								EndLine:   157,
								StartPos:  3574,
								EndPos:    3577,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 157,
									EndLine:   157,
									StartPos:  3574,
									EndPos:    3577,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 157,
											EndLine:   157,
											StartPos:  3574,
											EndPos:    3577,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 157,
								EndLine:   157,
								StartPos:  3579,
								EndPos:    3582,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 157,
									EndLine:   157,
									StartPos:  3579,
									EndPos:    3582,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 157,
											EndLine:   157,
											StartPos:  3579,
											EndPos:    3582,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
				},
			},
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 158,
						EndLine:   158,
						StartPos:  3586,
						EndPos:    3606,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 158,
								EndLine:   158,
								StartPos:  3590,
								EndPos:    3593,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 158,
									EndLine:   158,
									StartPos:  3590,
									EndPos:    3593,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 158,
											EndLine:   158,
											StartPos:  3590,
											EndPos:    3593,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 158,
								EndLine:   158,
								StartPos:  3595,
								EndPos:    3605,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 158,
									EndLine:   158,
									StartPos:  3595,
									EndPos:    3598,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 158,
											EndLine:   158,
											StartPos:  3595,
											EndPos:    3598,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 158,
									EndLine:   158,
									StartPos:  3602,
									EndPos:    3605,
								},
							},
							Value: []byte("Baz"),
						},
					},
				},
			},
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 159,
						EndLine:   159,
						StartPos:  3609,
						EndPos:    3632,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 159,
							EndLine:   159,
							StartPos:  3613,
							EndPos:    3621,
						},
					},
					Value: []byte("function"),
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 159,
								EndLine:   159,
								StartPos:  3622,
								EndPos:    3625,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 159,
									EndLine:   159,
									StartPos:  3622,
									EndPos:    3625,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 159,
											EndLine:   159,
											StartPos:  3622,
											EndPos:    3625,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 159,
								EndLine:   159,
								StartPos:  3628,
								EndPos:    3631,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 159,
									EndLine:   159,
									StartPos:  3628,
									EndPos:    3631,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 159,
											EndLine:   159,
											StartPos:  3628,
											EndPos:    3631,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
				},
			},
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 160,
						EndLine:   160,
						StartPos:  3635,
						EndPos:    3672,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 160,
							EndLine:   160,
							StartPos:  3639,
							EndPos:    3647,
						},
					},
					Value: []byte("function"),
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 160,
								EndLine:   160,
								StartPos:  3648,
								EndPos:    3658,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 160,
									EndLine:   160,
									StartPos:  3648,
									EndPos:    3651,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 160,
											EndLine:   160,
											StartPos:  3648,
											EndPos:    3651,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 160,
									EndLine:   160,
									StartPos:  3655,
									EndPos:    3658,
								},
							},
							Value: []byte("foo"),
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 160,
								EndLine:   160,
								StartPos:  3661,
								EndPos:    3671,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 160,
									EndLine:   160,
									StartPos:  3661,
									EndPos:    3664,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 160,
											EndLine:   160,
											StartPos:  3661,
											EndPos:    3664,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 160,
									EndLine:   160,
									StartPos:  3668,
									EndPos:    3671,
								},
							},
							Value: []byte("bar"),
						},
					},
				},
			},
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 161,
						EndLine:   161,
						StartPos:  3675,
						EndPos:    3695,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 161,
							EndLine:   161,
							StartPos:  3679,
							EndPos:    3684,
						},
					},
					Value: []byte("const"),
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 161,
								EndLine:   161,
								StartPos:  3685,
								EndPos:    3688,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 161,
									EndLine:   161,
									StartPos:  3685,
									EndPos:    3688,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 161,
											EndLine:   161,
											StartPos:  3685,
											EndPos:    3688,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 161,
								EndLine:   161,
								StartPos:  3691,
								EndPos:    3694,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 161,
									EndLine:   161,
									StartPos:  3691,
									EndPos:    3694,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 161,
											EndLine:   161,
											StartPos:  3691,
											EndPos:    3694,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
				},
			},
			&ast.StmtUseList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 162,
						EndLine:   162,
						StartPos:  3698,
						EndPos:    3732,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 162,
							EndLine:   162,
							StartPos:  3702,
							EndPos:    3707,
						},
					},
					Value: []byte("const"),
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3708,
								EndPos:    3718,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 162,
									EndLine:   162,
									StartPos:  3708,
									EndPos:    3711,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 162,
											EndLine:   162,
											StartPos:  3708,
											EndPos:    3711,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 162,
									EndLine:   162,
									StartPos:  3715,
									EndPos:    3718,
								},
							},
							Value: []byte("foo"),
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3721,
								EndPos:    3731,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 162,
									EndLine:   162,
									StartPos:  3721,
									EndPos:    3724,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 162,
											EndLine:   162,
											StartPos:  3721,
											EndPos:    3724,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 162,
									EndLine:   162,
									StartPos:  3728,
									EndPos:    3731,
								},
							},
							Value: []byte("bar"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 164,
						EndLine:   164,
						StartPos:  3736,
						EndPos:    3742,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 164,
							EndLine:   164,
							StartPos:  3736,
							EndPos:    3741,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 164,
								EndLine:   164,
								StartPos:  3736,
								EndPos:    3738,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 164,
									EndLine:   164,
									StartPos:  3736,
									EndPos:    3738,
								},
							},
							Value: []byte("a"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 164,
								EndLine:   164,
								StartPos:  3739,
								EndPos:    3740,
							},
						},
						Value: []byte("1"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 165,
						EndLine:   165,
						StartPos:  3745,
						EndPos:    3754,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 165,
							EndLine:   165,
							StartPos:  3745,
							EndPos:    3753,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 165,
								EndLine:   165,
								StartPos:  3745,
								EndPos:    3750,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 165,
									EndLine:   165,
									StartPos:  3745,
									EndPos:    3747,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 165,
										EndLine:   165,
										StartPos:  3745,
										EndPos:    3747,
									},
								},
								Value: []byte("a"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 165,
									EndLine:   165,
									StartPos:  3748,
									EndPos:    3749,
								},
							},
							Value: []byte("1"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 165,
								EndLine:   165,
								StartPos:  3751,
								EndPos:    3752,
							},
						},
						Value: []byte("2"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 166,
						EndLine:   166,
						StartPos:  3757,
						EndPos:    3765,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 166,
							EndLine:   166,
							StartPos:  3757,
							EndPos:    3764,
						},
					},
					Items: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 167,
						EndLine:   167,
						StartPos:  3768,
						EndPos:    3777,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 167,
							EndLine:   167,
							StartPos:  3768,
							EndPos:    3776,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 167,
									EndLine:   167,
									StartPos:  3774,
									EndPos:    3775,
								},
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 167,
										EndLine:   167,
										StartPos:  3774,
										EndPos:    3775,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 168,
						EndLine:   168,
						StartPos:  3780,
						EndPos:    3798,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 168,
							EndLine:   168,
							StartPos:  3780,
							EndPos:    3797,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 168,
									EndLine:   168,
									StartPos:  3786,
									EndPos:    3790,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 168,
										EndLine:   168,
										StartPos:  3786,
										EndPos:    3787,
									},
								},
								Value: []byte("1"),
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 168,
										EndLine:   168,
										StartPos:  3789,
										EndPos:    3790,
									},
								},
								Value: []byte("1"),
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 168,
									EndLine:   168,
									StartPos:  3792,
									EndPos:    3795,
								},
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 168,
										EndLine:   168,
										StartPos:  3792,
										EndPos:    3795,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 168,
											EndLine:   168,
											StartPos:  3793,
											EndPos:    3795,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 168,
												EndLine:   168,
												StartPos:  3793,
												EndPos:    3795,
											},
										},
										Value: []byte("b"),
									},
								},
							},
						},
						&ast.ExprArrayItem{},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 169,
						EndLine:   169,
						StartPos:  3801,
						EndPos:    3816,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 169,
							EndLine:   169,
							StartPos:  3801,
							EndPos:    3815,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 169,
									EndLine:   169,
									StartPos:  3807,
									EndPos:    3814,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 169,
										EndLine:   169,
										StartPos:  3807,
										EndPos:    3808,
									},
								},
								Value: []byte("3"),
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 169,
										EndLine:   169,
										StartPos:  3811,
										EndPos:    3814,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 169,
											EndLine:   169,
											StartPos:  3812,
											EndPos:    3814,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 169,
												EndLine:   169,
												StartPos:  3812,
												EndPos:    3814,
											},
										},
										Value: []byte("b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 170,
						EndLine:   170,
						StartPos:  3819,
						EndPos:    3848,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 170,
							EndLine:   170,
							StartPos:  3819,
							EndPos:    3847,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3825,
									EndPos:    3828,
								},
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3825,
										EndPos:    3828,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 170,
											EndLine:   170,
											StartPos:  3826,
											EndPos:    3828,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 170,
												EndLine:   170,
												StartPos:  3826,
												EndPos:    3828,
											},
										},
										Value: []byte("b"),
									},
								},
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3830,
									EndPos:    3834,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3830,
										EndPos:    3831,
									},
								},
								Value: []byte("1"),
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3833,
										EndPos:    3834,
									},
								},
								Value: []byte("1"),
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3836,
									EndPos:    3837,
								},
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3836,
										EndPos:    3837,
									},
								},
								Value: []byte("1"),
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 170,
									EndLine:   170,
									StartPos:  3839,
									EndPos:    3846,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3839,
										EndPos:    3840,
									},
								},
								Value: []byte("3"),
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 170,
										EndLine:   170,
										StartPos:  3843,
										EndPos:    3846,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 170,
											EndLine:   170,
											StartPos:  3844,
											EndPos:    3846,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 170,
												EndLine:   170,
												StartPos:  3844,
												EndPos:    3846,
											},
										},
										Value: []byte("b"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 171,
						EndLine:   171,
						StartPos:  3851,
						EndPos:    3855,
					},
				},
				Expr: &ast.ExprBitwiseNot{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 171,
							EndLine:   171,
							StartPos:  3851,
							EndPos:    3854,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 171,
								EndLine:   171,
								StartPos:  3852,
								EndPos:    3854,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 171,
									EndLine:   171,
									StartPos:  3852,
									EndPos:    3854,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 172,
						EndLine:   172,
						StartPos:  3858,
						EndPos:    3862,
					},
				},
				Expr: &ast.ExprBooleanNot{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 172,
							EndLine:   172,
							StartPos:  3858,
							EndPos:    3861,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3859,
								EndPos:    3861,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 172,
									EndLine:   172,
									StartPos:  3859,
									EndPos:    3861,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 174,
						EndLine:   174,
						StartPos:  3866,
						EndPos:    3875,
					},
				},
				Expr: &ast.ExprClassConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 174,
							EndLine:   174,
							StartPos:  3866,
							EndPos:    3874,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3866,
								EndPos:    3869,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 174,
										EndLine:   174,
										StartPos:  3866,
										EndPos:    3869,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ConstantName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3871,
								EndPos:    3874,
							},
						},
						Value: []byte("Bar"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 175,
						EndLine:   175,
						StartPos:  3878,
						EndPos:    3888,
					},
				},
				Expr: &ast.ExprClone{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 175,
							EndLine:   175,
							StartPos:  3878,
							EndPos:    3886,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3884,
								EndPos:    3886,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 175,
									EndLine:   175,
									StartPos:  3884,
									EndPos:    3886,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 176,
						EndLine:   176,
						StartPos:  3891,
						EndPos:    3900,
					},
				},
				Expr: &ast.ExprClone{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 176,
							EndLine:   176,
							StartPos:  3891,
							EndPos:    3899,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 176,
								EndLine:   176,
								StartPos:  3897,
								EndPos:    3899,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 176,
									EndLine:   176,
									StartPos:  3897,
									EndPos:    3899,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 177,
						EndLine:   177,
						StartPos:  3903,
						EndPos:    3916,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 177,
							EndLine:   177,
							StartPos:  3903,
							EndPos:    3915,
						},
					},
					Static:     false,
					ReturnsRef: false,
					Stmts:      []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 178,
						EndLine:   178,
						StartPos:  3919,
						EndPos:    3953,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 178,
							EndLine:   178,
							StartPos:  3919,
							EndPos:    3952,
						},
					},
					ReturnsRef: false,
					Static:     false,
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 178,
									EndLine:   178,
									StartPos:  3928,
									EndPos:    3930,
								},
							},
							ByRef:    false,
							Variadic: false,
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 178,
										EndLine:   178,
										StartPos:  3928,
										EndPos:    3930,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 178,
											EndLine:   178,
											StartPos:  3928,
											EndPos:    3930,
										},
									},
									Value: []byte("a"),
								},
							},
						},
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 178,
									EndLine:   178,
									StartPos:  3932,
									EndPos:    3934,
								},
							},
							Variadic: false,
							ByRef:    false,
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 178,
										EndLine:   178,
										StartPos:  3932,
										EndPos:    3934,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 178,
											EndLine:   178,
											StartPos:  3932,
											EndPos:    3934,
										},
									},
									Value: []byte("b"),
								},
							},
						},
					},
					ClosureUse: &ast.ExprClosureUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 178,
								EndLine:   178,
								StartPos:  3936,
								EndPos:    3949,
							},
						},
						Uses: []ast.Vertex{
							&ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 178,
										EndLine:   178,
										StartPos:  3941,
										EndPos:    3943,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 178,
											EndLine:   178,
											StartPos:  3941,
											EndPos:    3943,
										},
									},
									Value: []byte("c"),
								},
							},
							&ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 178,
										EndLine:   178,
										StartPos:  3945,
										EndPos:    3948,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 178,
											EndLine:   178,
											StartPos:  3946,
											EndPos:    3948,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 178,
												EndLine:   178,
												StartPos:  3946,
												EndPos:    3948,
											},
										},
										Value: []byte("d"),
									},
								},
							},
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 179,
						EndLine:   179,
						StartPos:  3956,
						EndPos:    3990,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 179,
							EndLine:   179,
							StartPos:  3956,
							EndPos:    3989,
						},
					},
					ReturnsRef: false,
					Static:     false,
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 179,
									EndLine:   179,
									StartPos:  3965,
									EndPos:    3967,
								},
							},
							ByRef:    false,
							Variadic: false,
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 179,
										EndLine:   179,
										StartPos:  3965,
										EndPos:    3967,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 179,
											EndLine:   179,
											StartPos:  3965,
											EndPos:    3967,
										},
									},
									Value: []byte("a"),
								},
							},
						},
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 179,
									EndLine:   179,
									StartPos:  3969,
									EndPos:    3971,
								},
							},
							ByRef:    false,
							Variadic: false,
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 179,
										EndLine:   179,
										StartPos:  3969,
										EndPos:    3971,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 179,
											EndLine:   179,
											StartPos:  3969,
											EndPos:    3971,
										},
									},
									Value: []byte("b"),
								},
							},
						},
					},
					ClosureUse: &ast.ExprClosureUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 179,
								EndLine:   179,
								StartPos:  3973,
								EndPos:    3986,
							},
						},
						Uses: []ast.Vertex{
							&ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 179,
										EndLine:   179,
										StartPos:  3978,
										EndPos:    3981,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 179,
											EndLine:   179,
											StartPos:  3979,
											EndPos:    3981,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 179,
												EndLine:   179,
												StartPos:  3979,
												EndPos:    3981,
											},
										},
										Value: []byte("c"),
									},
								},
							},
							&ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 179,
										EndLine:   179,
										StartPos:  3983,
										EndPos:    3985,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 179,
											EndLine:   179,
											StartPos:  3983,
											EndPos:    3985,
										},
									},
									Value: []byte("d"),
								},
							},
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 180,
						EndLine:   180,
						StartPos:  3993,
						EndPos:    4007,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 180,
							EndLine:   180,
							StartPos:  3993,
							EndPos:    4006,
						},
					},
					ReturnsRef: false,
					Static:     false,
					Stmts:      []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 181,
						EndLine:   181,
						StartPos:  4010,
						EndPos:    4014,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 181,
							EndLine:   181,
							StartPos:  4010,
							EndPos:    4013,
						},
					},
					Const: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 181,
								EndLine:   181,
								StartPos:  4010,
								EndPos:    4013,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 181,
										EndLine:   181,
										StartPos:  4010,
										EndPos:    4013,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 182,
						EndLine:   182,
						StartPos:  4017,
						EndPos:    4031,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 182,
							EndLine:   182,
							StartPos:  4017,
							EndPos:    4030,
						},
					},
					Const: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 182,
								EndLine:   182,
								StartPos:  4017,
								EndPos:    4030,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 182,
										EndLine:   182,
										StartPos:  4027,
										EndPos:    4030,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 183,
						EndLine:   183,
						StartPos:  4034,
						EndPos:    4039,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 183,
							EndLine:   183,
							StartPos:  4034,
							EndPos:    4038,
						},
					},
					Const: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 183,
								EndLine:   183,
								StartPos:  4034,
								EndPos:    4038,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 183,
										EndLine:   183,
										StartPos:  4035,
										EndPos:    4038,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 185,
						EndLine:   185,
						StartPos:  4043,
						EndPos:    4053,
					},
				},
				Expr: &ast.ExprEmpty{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 185,
							EndLine:   185,
							StartPos:  4043,
							EndPos:    4052,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 185,
								EndLine:   185,
								StartPos:  4049,
								EndPos:    4051,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 185,
									EndLine:   185,
									StartPos:  4049,
									EndPos:    4051,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 186,
						EndLine:   186,
						StartPos:  4056,
						EndPos:    4067,
					},
				},
				Expr: &ast.ExprEmpty{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 186,
							EndLine:   186,
							StartPos:  4056,
							EndPos:    4066,
						},
					},
					Expr: &ast.ExprConstFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 186,
								EndLine:   186,
								StartPos:  4062,
								EndPos:    4065,
							},
						},
						Const: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 186,
									EndLine:   186,
									StartPos:  4062,
									EndPos:    4065,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 186,
											EndLine:   186,
											StartPos:  4062,
											EndPos:    4065,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 187,
						EndLine:   187,
						StartPos:  4070,
						EndPos:    4074,
					},
				},
				Expr: &ast.ExprErrorSuppress{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 187,
							EndLine:   187,
							StartPos:  4070,
							EndPos:    4073,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 187,
								EndLine:   187,
								StartPos:  4071,
								EndPos:    4073,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 187,
									EndLine:   187,
									StartPos:  4071,
									EndPos:    4073,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 188,
						EndLine:   188,
						StartPos:  4077,
						EndPos:    4086,
					},
				},
				Expr: &ast.ExprEval{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 188,
							EndLine:   188,
							StartPos:  4077,
							EndPos:    4085,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 188,
								EndLine:   188,
								StartPos:  4082,
								EndPos:    4084,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 188,
									EndLine:   188,
									StartPos:  4082,
									EndPos:    4084,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 189,
						EndLine:   189,
						StartPos:  4089,
						EndPos:    4094,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 189,
							EndLine:   189,
							StartPos:  4089,
							EndPos:    4093,
						},
					},
					Die: false,
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 190,
						EndLine:   190,
						StartPos:  4097,
						EndPos:    4106,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 190,
							EndLine:   190,
							StartPos:  4097,
							EndPos:    4105,
						},
					},
					Die: false,
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 190,
								EndLine:   190,
								StartPos:  4102,
								EndPos:    4104,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 190,
									EndLine:   190,
									StartPos:  4102,
									EndPos:    4104,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 191,
						EndLine:   191,
						StartPos:  4109,
						EndPos:    4115,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 191,
							EndLine:   191,
							StartPos:  4109,
							EndPos:    4114,
						},
					},
					Die: true,
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 192,
						EndLine:   192,
						StartPos:  4118,
						EndPos:    4126,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 192,
							EndLine:   192,
							StartPos:  4118,
							EndPos:    4125,
						},
					},
					Die: true,
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 192,
								EndLine:   192,
								StartPos:  4122,
								EndPos:    4124,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 192,
									EndLine:   192,
									StartPos:  4122,
									EndPos:    4124,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 193,
						EndLine:   193,
						StartPos:  4129,
						EndPos:    4135,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 193,
							EndLine:   193,
							StartPos:  4129,
							EndPos:    4134,
						},
					},
					Function: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 193,
								EndLine:   193,
								StartPos:  4129,
								EndPos:    4132,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 193,
										EndLine:   193,
										StartPos:  4129,
										EndPos:    4132,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 193,
								EndLine:   193,
								StartPos:  4132,
								EndPos:    4134,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 194,
						EndLine:   194,
						StartPos:  4138,
						EndPos:    4157,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 194,
							EndLine:   194,
							StartPos:  4138,
							EndPos:    4156,
						},
					},
					Function: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 194,
								EndLine:   194,
								StartPos:  4138,
								EndPos:    4151,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 194,
										EndLine:   194,
										StartPos:  4148,
										EndPos:    4151,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 194,
								EndLine:   194,
								StartPos:  4151,
								EndPos:    4156,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 194,
										EndLine:   194,
										StartPos:  4153,
										EndPos:    4155,
									},
								},
								Variadic:    false,
								IsReference: true,
								Expr: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 194,
											EndLine:   194,
											StartPos:  4153,
											EndPos:    4155,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 194,
												EndLine:   194,
												StartPos:  4153,
												EndPos:    4155,
											},
										},
										Value: []byte("a"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 195,
						EndLine:   195,
						StartPos:  4160,
						EndPos:    4169,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 195,
							EndLine:   195,
							StartPos:  4160,
							EndPos:    4168,
						},
					},
					Function: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 195,
								EndLine:   195,
								StartPos:  4160,
								EndPos:    4164,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 195,
										EndLine:   195,
										StartPos:  4161,
										EndPos:    4164,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 195,
								EndLine:   195,
								StartPos:  4164,
								EndPos:    4168,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 195,
										EndLine:   195,
										StartPos:  4165,
										EndPos:    4167,
									},
								},
								Variadic:    false,
								IsReference: false,
								Expr: &ast.ExprShortArray{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 195,
											EndLine:   195,
											StartPos:  4165,
											EndPos:    4167,
										},
									},
									Items: []ast.Vertex{},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 196,
						EndLine:   196,
						StartPos:  4172,
						EndPos:    4187,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 196,
							EndLine:   196,
							StartPos:  4172,
							EndPos:    4186,
						},
					},
					Function: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 196,
								EndLine:   196,
								StartPos:  4172,
								EndPos:    4176,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 196,
									EndLine:   196,
									StartPos:  4172,
									EndPos:    4176,
								},
							},
							Value: []byte("foo"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 196,
								EndLine:   196,
								StartPos:  4176,
								EndPos:    4186,
							},
						},
						Arguments: []ast.Vertex{
							&ast.Argument{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 196,
										EndLine:   196,
										StartPos:  4177,
										EndPos:    4185,
									},
								},
								IsReference: false,
								Variadic:    false,
								Expr: &ast.ExprYield{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 196,
											EndLine:   196,
											StartPos:  4177,
											EndPos:    4185,
										},
									},
									Value: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 196,
												EndLine:   196,
												StartPos:  4183,
												EndPos:    4185,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 196,
													EndLine:   196,
													StartPos:  4183,
													EndPos:    4185,
												},
											},
											Value: []byte("a"),
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 198,
						EndLine:   198,
						StartPos:  4191,
						EndPos:    4196,
					},
				},
				Expr: &ast.ExprPostDec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 198,
							EndLine:   198,
							StartPos:  4191,
							EndPos:    4195,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 198,
								EndLine:   198,
								StartPos:  4191,
								EndPos:    4193,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 198,
									EndLine:   198,
									StartPos:  4191,
									EndPos:    4193,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 199,
						EndLine:   199,
						StartPos:  4199,
						EndPos:    4204,
					},
				},
				Expr: &ast.ExprPostInc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 199,
							EndLine:   199,
							StartPos:  4199,
							EndPos:    4203,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 199,
								EndLine:   199,
								StartPos:  4199,
								EndPos:    4201,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 199,
									EndLine:   199,
									StartPos:  4199,
									EndPos:    4201,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 200,
						EndLine:   200,
						StartPos:  4207,
						EndPos:    4212,
					},
				},
				Expr: &ast.ExprPreDec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 200,
							EndLine:   200,
							StartPos:  4207,
							EndPos:    4211,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 200,
								EndLine:   200,
								StartPos:  4209,
								EndPos:    4211,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 200,
									EndLine:   200,
									StartPos:  4209,
									EndPos:    4211,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 201,
						EndLine:   201,
						StartPos:  4215,
						EndPos:    4220,
					},
				},
				Expr: &ast.ExprPreInc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 201,
							EndLine:   201,
							StartPos:  4215,
							EndPos:    4219,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 201,
								EndLine:   201,
								StartPos:  4217,
								EndPos:    4219,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 201,
									EndLine:   201,
									StartPos:  4217,
									EndPos:    4219,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 203,
						EndLine:   203,
						StartPos:  4224,
						EndPos:    4235,
					},
				},
				Expr: &ast.ExprInclude{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 203,
							EndLine:   203,
							StartPos:  4224,
							EndPos:    4234,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 203,
								EndLine:   203,
								StartPos:  4232,
								EndPos:    4234,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 203,
									EndLine:   203,
									StartPos:  4232,
									EndPos:    4234,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 204,
						EndLine:   204,
						StartPos:  4238,
						EndPos:    4254,
					},
				},
				Expr: &ast.ExprIncludeOnce{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 204,
							EndLine:   204,
							StartPos:  4238,
							EndPos:    4253,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 204,
								EndLine:   204,
								StartPos:  4251,
								EndPos:    4253,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 204,
									EndLine:   204,
									StartPos:  4251,
									EndPos:    4253,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 205,
						EndLine:   205,
						StartPos:  4257,
						EndPos:    4268,
					},
				},
				Expr: &ast.ExprRequire{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 205,
							EndLine:   205,
							StartPos:  4257,
							EndPos:    4267,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 205,
								EndLine:   205,
								StartPos:  4265,
								EndPos:    4267,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 205,
									EndLine:   205,
									StartPos:  4265,
									EndPos:    4267,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 206,
						EndLine:   206,
						StartPos:  4271,
						EndPos:    4287,
					},
				},
				Expr: &ast.ExprRequireOnce{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 206,
							EndLine:   206,
							StartPos:  4271,
							EndPos:    4286,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 206,
								EndLine:   206,
								StartPos:  4284,
								EndPos:    4286,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 206,
									EndLine:   206,
									StartPos:  4284,
									EndPos:    4286,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 208,
						EndLine:   208,
						StartPos:  4291,
						EndPos:    4309,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 208,
							EndLine:   208,
							StartPos:  4291,
							EndPos:    4308,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 208,
								EndLine:   208,
								StartPos:  4291,
								EndPos:    4293,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 208,
									EndLine:   208,
									StartPos:  4291,
									EndPos:    4293,
								},
							},
							Value: []byte("a"),
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 208,
								EndLine:   208,
								StartPos:  4305,
								EndPos:    4308,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 208,
										EndLine:   208,
										StartPos:  4305,
										EndPos:    4308,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 209,
						EndLine:   209,
						StartPos:  4312,
						EndPos:    4340,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 209,
							EndLine:   209,
							StartPos:  4312,
							EndPos:    4339,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 209,
								EndLine:   209,
								StartPos:  4312,
								EndPos:    4314,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 209,
									EndLine:   209,
									StartPos:  4312,
									EndPos:    4314,
								},
							},
							Value: []byte("a"),
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 209,
								EndLine:   209,
								StartPos:  4326,
								EndPos:    4339,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 209,
										EndLine:   209,
										StartPos:  4336,
										EndPos:    4339,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 210,
						EndLine:   210,
						StartPos:  4343,
						EndPos:    4362,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 210,
							EndLine:   210,
							StartPos:  4343,
							EndPos:    4361,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 210,
								EndLine:   210,
								StartPos:  4343,
								EndPos:    4345,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 210,
									EndLine:   210,
									StartPos:  4343,
									EndPos:    4345,
								},
							},
							Value: []byte("a"),
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 210,
								EndLine:   210,
								StartPos:  4357,
								EndPos:    4361,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 210,
										EndLine:   210,
										StartPos:  4358,
										EndPos:    4361,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 212,
						EndLine:   212,
						StartPos:  4366,
						EndPos:    4380,
					},
				},
				Expr: &ast.ExprIsset{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 212,
							EndLine:   212,
							StartPos:  4366,
							EndPos:    4379,
						},
					},
					Vars: []ast.Vertex{
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 212,
									EndLine:   212,
									StartPos:  4372,
									EndPos:    4374,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 212,
										EndLine:   212,
										StartPos:  4372,
										EndPos:    4374,
									},
								},
								Value: []byte("a"),
							},
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 212,
									EndLine:   212,
									StartPos:  4376,
									EndPos:    4378,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 212,
										EndLine:   212,
										StartPos:  4376,
										EndPos:    4378,
									},
								},
								Value: []byte("b"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 213,
						EndLine:   213,
						StartPos:  4383,
						EndPos:    4394,
					},
				},
				Expr: &ast.ExprIsset{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 213,
							EndLine:   213,
							StartPos:  4383,
							EndPos:    4393,
						},
					},
					Vars: []ast.Vertex{
						&ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 213,
									EndLine:   213,
									StartPos:  4389,
									EndPos:    4392,
								},
							},
							Const: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 213,
										EndLine:   213,
										StartPos:  4389,
										EndPos:    4392,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 213,
												EndLine:   213,
												StartPos:  4389,
												EndPos:    4392,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 214,
						EndLine:   214,
						StartPos:  4397,
						EndPos:    4409,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 214,
							EndLine:   214,
							StartPos:  4397,
							EndPos:    4408,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 214,
								EndLine:   214,
								StartPos:  4397,
								EndPos:    4403,
							},
						},
						Items: []ast.Vertex{},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 214,
								EndLine:   214,
								StartPos:  4406,
								EndPos:    4408,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 214,
									EndLine:   214,
									StartPos:  4406,
									EndPos:    4408,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 215,
						EndLine:   215,
						StartPos:  4412,
						EndPos:    4430,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 215,
							EndLine:   215,
							StartPos:  4412,
							EndPos:    4429,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 215,
								EndLine:   215,
								StartPos:  4412,
								EndPos:    4424,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 215,
										EndLine:   215,
										StartPos:  4417,
										EndPos:    4419,
									},
								},
								Val: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 215,
											EndLine:   215,
											StartPos:  4417,
											EndPos:    4419,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 215,
												EndLine:   215,
												StartPos:  4417,
												EndPos:    4419,
											},
										},
										Value: []byte("a"),
									},
								},
							},
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 215,
										EndLine:   215,
										StartPos:  4421,
										EndPos:    4423,
									},
								},
								Val: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 215,
											EndLine:   215,
											StartPos:  4421,
											EndPos:    4423,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 215,
												EndLine:   215,
												StartPos:  4421,
												EndPos:    4423,
											},
										},
										Value: []byte("b"),
									},
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 215,
								EndLine:   215,
								StartPos:  4427,
								EndPos:    4429,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 215,
									EndLine:   215,
									StartPos:  4427,
									EndPos:    4429,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 216,
						EndLine:   216,
						StartPos:  4433,
						EndPos:    4449,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 216,
							EndLine:   216,
							StartPos:  4433,
							EndPos:    4448,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 216,
								EndLine:   216,
								StartPos:  4433,
								EndPos:    4443,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 216,
										EndLine:   216,
										StartPos:  4438,
										EndPos:    4442,
									},
								},
								Val: &ast.ExprArrayDimFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 216,
											EndLine:   216,
											StartPos:  4438,
											EndPos:    4442,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 216,
												EndLine:   216,
												StartPos:  4438,
												EndPos:    4440,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 216,
													EndLine:   216,
													StartPos:  4438,
													EndPos:    4440,
												},
											},
											Value: []byte("a"),
										},
									},
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 216,
								EndLine:   216,
								StartPos:  4446,
								EndPos:    4448,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 216,
									EndLine:   216,
									StartPos:  4446,
									EndPos:    4448,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 217,
						EndLine:   217,
						StartPos:  4452,
						EndPos:    4472,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 217,
							EndLine:   217,
							StartPos:  4452,
							EndPos:    4471,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 217,
								EndLine:   217,
								StartPos:  4452,
								EndPos:    4466,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 217,
										EndLine:   217,
										StartPos:  4457,
										EndPos:    4465,
									},
								},
								Val: &ast.ExprList{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 217,
											EndLine:   217,
											StartPos:  4457,
											EndPos:    4465,
										},
									},
									Items: []ast.Vertex{
										&ast.ExprArrayItem{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 217,
													EndLine:   217,
													StartPos:  4462,
													EndPos:    4464,
												},
											},
											Val: &ast.ExprVariable{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 217,
														EndLine:   217,
														StartPos:  4462,
														EndPos:    4464,
													},
												},
												VarName: &ast.Identifier{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 217,
															EndLine:   217,
															StartPos:  4462,
															EndPos:    4464,
														},
													},
													Value: []byte("a"),
												},
											},
										},
									},
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 217,
								EndLine:   217,
								StartPos:  4469,
								EndPos:    4471,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 217,
									EndLine:   217,
									StartPos:  4469,
									EndPos:    4471,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 219,
						EndLine:   219,
						StartPos:  4476,
						EndPos:    4486,
					},
				},
				Expr: &ast.ExprMethodCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 219,
							EndLine:   219,
							StartPos:  4476,
							EndPos:    4485,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 219,
								EndLine:   219,
								StartPos:  4476,
								EndPos:    4478,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 219,
									EndLine:   219,
									StartPos:  4476,
									EndPos:    4478,
								},
							},
							Value: []byte("a"),
						},
					},
					Method: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 219,
								EndLine:   219,
								StartPos:  4480,
								EndPos:    4483,
							},
						},
						Value: []byte("foo"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 219,
								EndLine:   219,
								StartPos:  4483,
								EndPos:    4485,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 220,
						EndLine:   220,
						StartPos:  4489,
						EndPos:    4497,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 220,
							EndLine:   220,
							StartPos:  4489,
							EndPos:    4496,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 220,
								EndLine:   220,
								StartPos:  4493,
								EndPos:    4496,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 220,
										EndLine:   220,
										StartPos:  4493,
										EndPos:    4496,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 221,
						EndLine:   221,
						StartPos:  4500,
						EndPos:    4520,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 221,
							EndLine:   221,
							StartPos:  4500,
							EndPos:    4519,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 221,
								EndLine:   221,
								StartPos:  4504,
								EndPos:    4517,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 221,
										EndLine:   221,
										StartPos:  4514,
										EndPos:    4517,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 221,
								EndLine:   221,
								StartPos:  4517,
								EndPos:    4519,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 222,
						EndLine:   222,
						StartPos:  4523,
						EndPos:    4534,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 222,
							EndLine:   222,
							StartPos:  4523,
							EndPos:    4533,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 222,
								EndLine:   222,
								StartPos:  4527,
								EndPos:    4531,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 222,
										EndLine:   222,
										StartPos:  4528,
										EndPos:    4531,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 222,
								EndLine:   222,
								StartPos:  4531,
								EndPos:    4533,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 223,
						EndLine:   223,
						StartPos:  4537,
						EndPos:    4547,
					},
				},
				Expr: &ast.ExprPrint{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 223,
							EndLine:   223,
							StartPos:  4537,
							EndPos:    4545,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 223,
								EndLine:   223,
								StartPos:  4543,
								EndPos:    4545,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 223,
									EndLine:   223,
									StartPos:  4543,
									EndPos:    4545,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 224,
						EndLine:   224,
						StartPos:  4550,
						EndPos:    4558,
					},
				},
				Expr: &ast.ExprPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 224,
							EndLine:   224,
							StartPos:  4550,
							EndPos:    4557,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 224,
								EndLine:   224,
								StartPos:  4550,
								EndPos:    4552,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 224,
									EndLine:   224,
									StartPos:  4550,
									EndPos:    4552,
								},
							},
							Value: []byte("a"),
						},
					},
					Property: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 224,
								EndLine:   224,
								StartPos:  4554,
								EndPos:    4557,
							},
						},
						Value: []byte("foo"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 225,
						EndLine:   225,
						StartPos:  4561,
						EndPos:    4572,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 225,
							EndLine:   225,
							StartPos:  4561,
							EndPos:    4570,
						},
					},
					Var: &ast.ExprPropertyFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 225,
								EndLine:   225,
								StartPos:  4561,
								EndPos:    4568,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 225,
									EndLine:   225,
									StartPos:  4561,
									EndPos:    4563,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 225,
										EndLine:   225,
										StartPos:  4561,
										EndPos:    4563,
									},
								},
								Value: []byte("a"),
							},
						},
						Property: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 225,
									EndLine:   225,
									StartPos:  4565,
									EndPos:    4568,
								},
							},
							Value: []byte("foo"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 225,
								EndLine:   225,
								StartPos:  4569,
								EndPos:    4570,
							},
						},
						Value: []byte("1"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 226,
						EndLine:   226,
						StartPos:  4575,
						EndPos:    4604,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 226,
							EndLine:   226,
							StartPos:  4575,
							EndPos:    4602,
						},
					},
					Var: &ast.ExprPropertyFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 226,
								EndLine:   226,
								StartPos:  4575,
								EndPos:    4600,
							},
						},
						Var: &ast.ExprMethodCall{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 226,
									EndLine:   226,
									StartPos:  4575,
									EndPos:    4594,
								},
							},
							Var: &ast.ExprPropertyFetch{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 226,
										EndLine:   226,
										StartPos:  4575,
										EndPos:    4587,
									},
								},
								Var: &ast.ExprPropertyFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 226,
											EndLine:   226,
											StartPos:  4575,
											EndPos:    4582,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 226,
												EndLine:   226,
												StartPos:  4575,
												EndPos:    4577,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 226,
													EndLine:   226,
													StartPos:  4575,
													EndPos:    4577,
												},
											},
											Value: []byte("a"),
										},
									},
									Property: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 226,
												EndLine:   226,
												StartPos:  4579,
												EndPos:    4582,
											},
										},
										Value: []byte("foo"),
									},
								},
								Property: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 226,
											EndLine:   226,
											StartPos:  4584,
											EndPos:    4587,
										},
									},
									Value: []byte("bar"),
								},
							},
							Method: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 226,
										EndLine:   226,
										StartPos:  4589,
										EndPos:    4592,
									},
								},
								Value: []byte("baz"),
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 226,
										EndLine:   226,
										StartPos:  4592,
										EndPos:    4594,
									},
								},
							},
						},
						Property: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 226,
									EndLine:   226,
									StartPos:  4596,
									EndPos:    4600,
								},
							},
							Value: []byte("quux"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 226,
								EndLine:   226,
								StartPos:  4601,
								EndPos:    4602,
							},
						},
						Value: []byte("0"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 227,
						EndLine:   227,
						StartPos:  4607,
						EndPos:    4623,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 227,
							EndLine:   227,
							StartPos:  4607,
							EndPos:    4621,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 227,
								EndLine:   227,
								StartPos:  4607,
								EndPos:    4618,
							},
						},
						Var: &ast.ExprMethodCall{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 227,
									EndLine:   227,
									StartPos:  4607,
									EndPos:    4616,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 227,
										EndLine:   227,
										StartPos:  4607,
										EndPos:    4609,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 227,
											EndLine:   227,
											StartPos:  4607,
											EndPos:    4609,
										},
									},
									Value: []byte("a"),
								},
							},
							Method: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 227,
										EndLine:   227,
										StartPos:  4611,
										EndPos:    4614,
									},
								},
								Value: []byte("foo"),
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 227,
										EndLine:   227,
										StartPos:  4614,
										EndPos:    4616,
									},
								},
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 227,
									EndLine:   227,
									StartPos:  4617,
									EndPos:    4618,
								},
							},
							Value: []byte("1"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 227,
								EndLine:   227,
								StartPos:  4620,
								EndPos:    4621,
							},
						},
						Value: []byte("1"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 228,
						EndLine:   228,
						StartPos:  4626,
						EndPos:    4635,
					},
				},
				Expr: &ast.ExprShellExec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 228,
							EndLine:   228,
							StartPos:  4626,
							EndPos:    4634,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 228,
									EndLine:   228,
									StartPos:  4627,
									EndPos:    4631,
								},
							},
							Value: []byte("cmd "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 228,
									EndLine:   228,
									StartPos:  4631,
									EndPos:    4633,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 228,
										EndLine:   228,
										StartPos:  4631,
										EndPos:    4633,
									},
								},
								Value: []byte("a"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 229,
						EndLine:   229,
						StartPos:  4638,
						EndPos:    4644,
					},
				},
				Expr: &ast.ExprShellExec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 229,
							EndLine:   229,
							StartPos:  4638,
							EndPos:    4643,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 229,
									EndLine:   229,
									StartPos:  4639,
									EndPos:    4642,
								},
							},
							Value: []byte("cmd"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 230,
						EndLine:   230,
						StartPos:  4647,
						EndPos:    4650,
					},
				},
				Expr: &ast.ExprShellExec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 230,
							EndLine:   230,
							StartPos:  4647,
							EndPos:    4649,
						},
					},
					Parts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 231,
						EndLine:   231,
						StartPos:  4653,
						EndPos:    4656,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 231,
							EndLine:   231,
							StartPos:  4653,
							EndPos:    4655,
						},
					},
					Items: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 232,
						EndLine:   232,
						StartPos:  4659,
						EndPos:    4663,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 232,
							EndLine:   232,
							StartPos:  4659,
							EndPos:    4662,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 232,
									EndLine:   232,
									StartPos:  4660,
									EndPos:    4661,
								},
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 232,
										EndLine:   232,
										StartPos:  4660,
										EndPos:    4661,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 233,
						EndLine:   233,
						StartPos:  4666,
						EndPos:    4679,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 233,
							EndLine:   233,
							StartPos:  4666,
							EndPos:    4678,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 233,
									EndLine:   233,
									StartPos:  4667,
									EndPos:    4671,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 233,
										EndLine:   233,
										StartPos:  4667,
										EndPos:    4668,
									},
								},
								Value: []byte("1"),
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 233,
										EndLine:   233,
										StartPos:  4670,
										EndPos:    4671,
									},
								},
								Value: []byte("1"),
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 233,
									EndLine:   233,
									StartPos:  4673,
									EndPos:    4676,
								},
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 233,
										EndLine:   233,
										StartPos:  4673,
										EndPos:    4676,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 233,
											EndLine:   233,
											StartPos:  4674,
											EndPos:    4676,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 233,
												EndLine:   233,
												StartPos:  4674,
												EndPos:    4676,
											},
										},
										Value: []byte("b"),
									},
								},
							},
						},
						&ast.ExprArrayItem{},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 235,
						EndLine:   235,
						StartPos:  4683,
						EndPos:    4694,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 235,
							EndLine:   235,
							StartPos:  4683,
							EndPos:    4693,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 235,
								EndLine:   235,
								StartPos:  4683,
								EndPos:    4686,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 235,
										EndLine:   235,
										StartPos:  4683,
										EndPos:    4686,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 235,
								EndLine:   235,
								StartPos:  4688,
								EndPos:    4691,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 235,
								EndLine:   235,
								StartPos:  4691,
								EndPos:    4693,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 236,
						EndLine:   236,
						StartPos:  4697,
						EndPos:    4718,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 236,
							EndLine:   236,
							StartPos:  4697,
							EndPos:    4717,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 236,
								EndLine:   236,
								StartPos:  4697,
								EndPos:    4710,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 236,
										EndLine:   236,
										StartPos:  4707,
										EndPos:    4710,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 236,
								EndLine:   236,
								StartPos:  4712,
								EndPos:    4715,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 236,
								EndLine:   236,
								StartPos:  4715,
								EndPos:    4717,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 237,
						EndLine:   237,
						StartPos:  4721,
						EndPos:    4733,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 237,
							EndLine:   237,
							StartPos:  4721,
							EndPos:    4732,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 237,
								EndLine:   237,
								StartPos:  4721,
								EndPos:    4725,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 237,
										EndLine:   237,
										StartPos:  4722,
										EndPos:    4725,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 237,
								EndLine:   237,
								StartPos:  4727,
								EndPos:    4730,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 237,
								EndLine:   237,
								StartPos:  4730,
								EndPos:    4732,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 238,
						EndLine:   238,
						StartPos:  4736,
						EndPos:    4748,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 238,
							EndLine:   238,
							StartPos:  4736,
							EndPos:    4747,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 238,
								EndLine:   238,
								StartPos:  4736,
								EndPos:    4739,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 238,
										EndLine:   238,
										StartPos:  4736,
										EndPos:    4739,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 238,
								EndLine:   238,
								StartPos:  4741,
								EndPos:    4745,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 238,
									EndLine:   238,
									StartPos:  4741,
									EndPos:    4745,
								},
							},
							Value: []byte("bar"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 238,
								EndLine:   238,
								StartPos:  4745,
								EndPos:    4747,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 239,
						EndLine:   239,
						StartPos:  4751,
						EndPos:    4764,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 239,
							EndLine:   239,
							StartPos:  4751,
							EndPos:    4763,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 239,
								EndLine:   239,
								StartPos:  4751,
								EndPos:    4755,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 239,
									EndLine:   239,
									StartPos:  4751,
									EndPos:    4755,
								},
							},
							Value: []byte("foo"),
						},
					},
					Call: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 239,
								EndLine:   239,
								StartPos:  4757,
								EndPos:    4761,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 239,
									EndLine:   239,
									StartPos:  4757,
									EndPos:    4761,
								},
							},
							Value: []byte("bar"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 239,
								EndLine:   239,
								StartPos:  4761,
								EndPos:    4763,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 240,
						EndLine:   240,
						StartPos:  4767,
						EndPos:    4777,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 240,
							EndLine:   240,
							StartPos:  4767,
							EndPos:    4776,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 240,
								EndLine:   240,
								StartPos:  4767,
								EndPos:    4770,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 240,
										EndLine:   240,
										StartPos:  4767,
										EndPos:    4770,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 240,
								EndLine:   240,
								StartPos:  4772,
								EndPos:    4776,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 240,
									EndLine:   240,
									StartPos:  4772,
									EndPos:    4776,
								},
							},
							Value: []byte("bar"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 241,
						EndLine:   241,
						StartPos:  4780,
						EndPos:    4800,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 241,
							EndLine:   241,
							StartPos:  4780,
							EndPos:    4799,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 241,
								EndLine:   241,
								StartPos:  4780,
								EndPos:    4793,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 241,
										EndLine:   241,
										StartPos:  4790,
										EndPos:    4793,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 241,
								EndLine:   241,
								StartPos:  4795,
								EndPos:    4799,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 241,
									EndLine:   241,
									StartPos:  4795,
									EndPos:    4799,
								},
							},
							Value: []byte("bar"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 242,
						EndLine:   242,
						StartPos:  4803,
						EndPos:    4814,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 242,
							EndLine:   242,
							StartPos:  4803,
							EndPos:    4813,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 242,
								EndLine:   242,
								StartPos:  4803,
								EndPos:    4807,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 242,
										EndLine:   242,
										StartPos:  4804,
										EndPos:    4807,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 242,
								EndLine:   242,
								StartPos:  4809,
								EndPos:    4813,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 242,
									EndLine:   242,
									StartPos:  4809,
									EndPos:    4813,
								},
							},
							Value: []byte("bar"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 243,
						EndLine:   243,
						StartPos:  4817,
						EndPos:    4830,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 243,
							EndLine:   243,
							StartPos:  4817,
							EndPos:    4829,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 243,
								EndLine:   243,
								StartPos:  4817,
								EndPos:    4819,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 243,
									EndLine:   243,
									StartPos:  4817,
									EndPos:    4819,
								},
							},
							Value: []byte("a"),
						},
					},
					IfTrue: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 243,
								EndLine:   243,
								StartPos:  4822,
								EndPos:    4824,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 243,
									EndLine:   243,
									StartPos:  4822,
									EndPos:    4824,
								},
							},
							Value: []byte("b"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 243,
								EndLine:   243,
								StartPos:  4827,
								EndPos:    4829,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 243,
									EndLine:   243,
									StartPos:  4827,
									EndPos:    4829,
								},
							},
							Value: []byte("c"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 244,
						EndLine:   244,
						StartPos:  4833,
						EndPos:    4843,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 244,
							EndLine:   244,
							StartPos:  4833,
							EndPos:    4842,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 244,
								EndLine:   244,
								StartPos:  4833,
								EndPos:    4835,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 244,
									EndLine:   244,
									StartPos:  4833,
									EndPos:    4835,
								},
							},
							Value: []byte("a"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 244,
								EndLine:   244,
								StartPos:  4840,
								EndPos:    4842,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 244,
									EndLine:   244,
									StartPos:  4840,
									EndPos:    4842,
								},
							},
							Value: []byte("c"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 245,
						EndLine:   245,
						StartPos:  4846,
						EndPos:    4869,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4846,
							EndPos:    4868,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4846,
								EndPos:    4848,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4846,
									EndPos:    4848,
								},
							},
							Value: []byte("a"),
						},
					},
					IfTrue: &ast.ExprTernary{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4851,
								EndPos:    4863,
							},
						},
						Condition: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4851,
									EndPos:    4853,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 245,
										EndLine:   245,
										StartPos:  4851,
										EndPos:    4853,
									},
								},
								Value: []byte("b"),
							},
						},
						IfTrue: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4856,
									EndPos:    4858,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 245,
										EndLine:   245,
										StartPos:  4856,
										EndPos:    4858,
									},
								},
								Value: []byte("c"),
							},
						},
						IfFalse: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4861,
									EndPos:    4863,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 245,
										EndLine:   245,
										StartPos:  4861,
										EndPos:    4863,
									},
								},
								Value: []byte("d"),
							},
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4866,
								EndPos:    4868,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 245,
									EndLine:   245,
									StartPos:  4866,
									EndPos:    4868,
								},
							},
							Value: []byte("e"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 246,
						EndLine:   246,
						StartPos:  4872,
						EndPos:    4895,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4872,
							EndPos:    4894,
						},
					},
					Condition: &ast.ExprTernary{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4872,
								EndPos:    4884,
							},
						},
						Condition: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4872,
									EndPos:    4874,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 246,
										EndLine:   246,
										StartPos:  4872,
										EndPos:    4874,
									},
								},
								Value: []byte("a"),
							},
						},
						IfTrue: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4877,
									EndPos:    4879,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 246,
										EndLine:   246,
										StartPos:  4877,
										EndPos:    4879,
									},
								},
								Value: []byte("b"),
							},
						},
						IfFalse: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4882,
									EndPos:    4884,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 246,
										EndLine:   246,
										StartPos:  4882,
										EndPos:    4884,
									},
								},
								Value: []byte("c"),
							},
						},
					},
					IfTrue: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4887,
								EndPos:    4889,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4887,
									EndPos:    4889,
								},
							},
							Value: []byte("d"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4892,
								EndPos:    4894,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 246,
									EndLine:   246,
									StartPos:  4892,
									EndPos:    4894,
								},
							},
							Value: []byte("e"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 247,
						EndLine:   247,
						StartPos:  4898,
						EndPos:    4902,
					},
				},
				Expr: &ast.ExprUnaryMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 247,
							EndLine:   247,
							StartPos:  4898,
							EndPos:    4901,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 247,
								EndLine:   247,
								StartPos:  4899,
								EndPos:    4901,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 247,
									EndLine:   247,
									StartPos:  4899,
									EndPos:    4901,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 248,
						EndLine:   248,
						StartPos:  4905,
						EndPos:    4909,
					},
				},
				Expr: &ast.ExprUnaryPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 248,
							EndLine:   248,
							StartPos:  4905,
							EndPos:    4908,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4906,
								EndPos:    4908,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 248,
									EndLine:   248,
									StartPos:  4906,
									EndPos:    4908,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 249,
						EndLine:   249,
						StartPos:  4912,
						EndPos:    4916,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 249,
							EndLine:   249,
							StartPos:  4912,
							EndPos:    4915,
						},
					},
					VarName: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 249,
								EndLine:   249,
								StartPos:  4913,
								EndPos:    4915,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 249,
									EndLine:   249,
									StartPos:  4913,
									EndPos:    4915,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 250,
						EndLine:   250,
						StartPos:  4919,
						EndPos:    4924,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 250,
							EndLine:   250,
							StartPos:  4919,
							EndPos:    4923,
						},
					},
					VarName: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 250,
								EndLine:   250,
								StartPos:  4920,
								EndPos:    4923,
							},
						},
						VarName: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 250,
									EndLine:   250,
									StartPos:  4921,
									EndPos:    4923,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 250,
										EndLine:   250,
										StartPos:  4921,
										EndPos:    4923,
									},
								},
								Value: []byte("a"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 251,
						EndLine:   251,
						StartPos:  4927,
						EndPos:    4933,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 251,
							EndLine:   251,
							StartPos:  4927,
							EndPos:    4932,
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 252,
						EndLine:   252,
						StartPos:  4936,
						EndPos:    4945,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 252,
							EndLine:   252,
							StartPos:  4936,
							EndPos:    4944,
						},
					},
					Value: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 252,
								EndLine:   252,
								StartPos:  4942,
								EndPos:    4944,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 252,
									EndLine:   252,
									StartPos:  4942,
									EndPos:    4944,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 253,
						EndLine:   253,
						StartPos:  4948,
						EndPos:    4963,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 253,
							EndLine:   253,
							StartPos:  4948,
							EndPos:    4962,
						},
					},
					Key: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4954,
								EndPos:    4956,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4954,
									EndPos:    4956,
								},
							},
							Value: []byte("a"),
						},
					},
					Value: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4960,
								EndPos:    4962,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4960,
									EndPos:    4962,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 254,
						EndLine:   254,
						StartPos:  4966,
						EndPos:    4983,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 254,
							EndLine:   254,
							StartPos:  4966,
							EndPos:    4982,
						},
					},
					Value: &ast.ExprClassConstFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  4972,
								EndPos:    4982,
							},
						},
						Class: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4972,
									EndPos:    4975,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 254,
											EndLine:   254,
											StartPos:  4972,
											EndPos:    4975,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4977,
									EndPos:    4982,
								},
							},
							Value: []byte("class"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 255,
						EndLine:   255,
						StartPos:  4986,
						EndPos:    5009,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 255,
							EndLine:   255,
							StartPos:  4986,
							EndPos:    5008,
						},
					},
					Key: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  4992,
								EndPos:    4994,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 255,
									EndLine:   255,
									StartPos:  4992,
									EndPos:    4994,
								},
							},
							Value: []byte("a"),
						},
					},
					Value: &ast.ExprClassConstFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  4998,
								EndPos:    5008,
							},
						},
						Class: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 255,
									EndLine:   255,
									StartPos:  4998,
									EndPos:    5001,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 255,
											EndLine:   255,
											StartPos:  4998,
											EndPos:    5001,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 255,
									EndLine:   255,
									StartPos:  5003,
									EndPos:    5008,
								},
							},
							Value: []byte("class"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 257,
						EndLine:   257,
						StartPos:  5015,
						EndPos:    5025,
					},
				},
				Expr: &ast.ExprCastArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 257,
							EndLine:   257,
							StartPos:  5015,
							EndPos:    5024,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 257,
								EndLine:   257,
								StartPos:  5022,
								EndPos:    5024,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 257,
									EndLine:   257,
									StartPos:  5022,
									EndPos:    5024,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 258,
						EndLine:   258,
						StartPos:  5028,
						EndPos:    5040,
					},
				},
				Expr: &ast.ExprCastBool{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 258,
							EndLine:   258,
							StartPos:  5028,
							EndPos:    5039,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 258,
								EndLine:   258,
								StartPos:  5037,
								EndPos:    5039,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 258,
									EndLine:   258,
									StartPos:  5037,
									EndPos:    5039,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 259,
						EndLine:   259,
						StartPos:  5043,
						EndPos:    5052,
					},
				},
				Expr: &ast.ExprCastBool{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 259,
							EndLine:   259,
							StartPos:  5043,
							EndPos:    5051,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 259,
								EndLine:   259,
								StartPos:  5049,
								EndPos:    5051,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 259,
									EndLine:   259,
									StartPos:  5049,
									EndPos:    5051,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 260,
						EndLine:   260,
						StartPos:  5055,
						EndPos:    5066,
					},
				},
				Expr: &ast.ExprCastDouble{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 260,
							EndLine:   260,
							StartPos:  5055,
							EndPos:    5065,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 260,
								EndLine:   260,
								StartPos:  5063,
								EndPos:    5065,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 260,
									EndLine:   260,
									StartPos:  5063,
									EndPos:    5065,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 261,
						EndLine:   261,
						StartPos:  5069,
						EndPos:    5079,
					},
				},
				Expr: &ast.ExprCastDouble{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 261,
							EndLine:   261,
							StartPos:  5069,
							EndPos:    5078,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 261,
								EndLine:   261,
								StartPos:  5076,
								EndPos:    5078,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 261,
									EndLine:   261,
									StartPos:  5076,
									EndPos:    5078,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 262,
						EndLine:   262,
						StartPos:  5082,
						EndPos:    5094,
					},
				},
				Expr: &ast.ExprCastInt{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 262,
							EndLine:   262,
							StartPos:  5082,
							EndPos:    5093,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 262,
								EndLine:   262,
								StartPos:  5091,
								EndPos:    5093,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 262,
									EndLine:   262,
									StartPos:  5091,
									EndPos:    5093,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 263,
						EndLine:   263,
						StartPos:  5097,
						EndPos:    5105,
					},
				},
				Expr: &ast.ExprCastInt{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 263,
							EndLine:   263,
							StartPos:  5097,
							EndPos:    5104,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 263,
								EndLine:   263,
								StartPos:  5102,
								EndPos:    5104,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 263,
									EndLine:   263,
									StartPos:  5102,
									EndPos:    5104,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 264,
						EndLine:   264,
						StartPos:  5108,
						EndPos:    5119,
					},
				},
				Expr: &ast.ExprCastObject{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 264,
							EndLine:   264,
							StartPos:  5108,
							EndPos:    5118,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 264,
								EndLine:   264,
								StartPos:  5116,
								EndPos:    5118,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 264,
									EndLine:   264,
									StartPos:  5116,
									EndPos:    5118,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 265,
						EndLine:   265,
						StartPos:  5122,
						EndPos:    5133,
					},
				},
				Expr: &ast.ExprCastString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 265,
							EndLine:   265,
							StartPos:  5122,
							EndPos:    5132,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 265,
								EndLine:   265,
								StartPos:  5130,
								EndPos:    5132,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 265,
									EndLine:   265,
									StartPos:  5130,
									EndPos:    5132,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 266,
						EndLine:   266,
						StartPos:  5136,
						EndPos:    5146,
					},
				},
				Expr: &ast.ExprCastUnset{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 266,
							EndLine:   266,
							StartPos:  5136,
							EndPos:    5145,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 266,
								EndLine:   266,
								StartPos:  5143,
								EndPos:    5145,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 266,
									EndLine:   266,
									StartPos:  5143,
									EndPos:    5145,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 268,
						EndLine:   268,
						StartPos:  5150,
						EndPos:    5158,
					},
				},
				Expr: &ast.ExprBinaryBitwiseAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 268,
							EndLine:   268,
							StartPos:  5150,
							EndPos:    5157,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 268,
								EndLine:   268,
								StartPos:  5150,
								EndPos:    5152,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 268,
									EndLine:   268,
									StartPos:  5150,
									EndPos:    5152,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 268,
								EndLine:   268,
								StartPos:  5155,
								EndPos:    5157,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 268,
									EndLine:   268,
									StartPos:  5155,
									EndPos:    5157,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 269,
						EndLine:   269,
						StartPos:  5161,
						EndPos:    5169,
					},
				},
				Expr: &ast.ExprBinaryBitwiseOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 269,
							EndLine:   269,
							StartPos:  5161,
							EndPos:    5168,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 269,
								EndLine:   269,
								StartPos:  5161,
								EndPos:    5163,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 269,
									EndLine:   269,
									StartPos:  5161,
									EndPos:    5163,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 269,
								EndLine:   269,
								StartPos:  5166,
								EndPos:    5168,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 269,
									EndLine:   269,
									StartPos:  5166,
									EndPos:    5168,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 270,
						EndLine:   270,
						StartPos:  5172,
						EndPos:    5180,
					},
				},
				Expr: &ast.ExprBinaryBitwiseXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 270,
							EndLine:   270,
							StartPos:  5172,
							EndPos:    5179,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 270,
								EndLine:   270,
								StartPos:  5172,
								EndPos:    5174,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 270,
									EndLine:   270,
									StartPos:  5172,
									EndPos:    5174,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 270,
								EndLine:   270,
								StartPos:  5177,
								EndPos:    5179,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 270,
									EndLine:   270,
									StartPos:  5177,
									EndPos:    5179,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 271,
						EndLine:   271,
						StartPos:  5183,
						EndPos:    5192,
					},
				},
				Expr: &ast.ExprBinaryBooleanAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 271,
							EndLine:   271,
							StartPos:  5183,
							EndPos:    5191,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 271,
								EndLine:   271,
								StartPos:  5183,
								EndPos:    5185,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 271,
									EndLine:   271,
									StartPos:  5183,
									EndPos:    5185,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 271,
								EndLine:   271,
								StartPos:  5189,
								EndPos:    5191,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 271,
									EndLine:   271,
									StartPos:  5189,
									EndPos:    5191,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 272,
						EndLine:   272,
						StartPos:  5195,
						EndPos:    5204,
					},
				},
				Expr: &ast.ExprBinaryBooleanOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 272,
							EndLine:   272,
							StartPos:  5195,
							EndPos:    5203,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 272,
								EndLine:   272,
								StartPos:  5195,
								EndPos:    5197,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 272,
									EndLine:   272,
									StartPos:  5195,
									EndPos:    5197,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 272,
								EndLine:   272,
								StartPos:  5201,
								EndPos:    5203,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 272,
									EndLine:   272,
									StartPos:  5201,
									EndPos:    5203,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 273,
						EndLine:   273,
						StartPos:  5207,
						EndPos:    5215,
					},
				},
				Expr: &ast.ExprBinaryConcat{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 273,
							EndLine:   273,
							StartPos:  5207,
							EndPos:    5214,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 273,
								EndLine:   273,
								StartPos:  5207,
								EndPos:    5209,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 273,
									EndLine:   273,
									StartPos:  5207,
									EndPos:    5209,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 273,
								EndLine:   273,
								StartPos:  5212,
								EndPos:    5214,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 273,
									EndLine:   273,
									StartPos:  5212,
									EndPos:    5214,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 274,
						EndLine:   274,
						StartPos:  5218,
						EndPos:    5226,
					},
				},
				Expr: &ast.ExprBinaryDiv{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 274,
							EndLine:   274,
							StartPos:  5218,
							EndPos:    5225,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 274,
								EndLine:   274,
								StartPos:  5218,
								EndPos:    5220,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 274,
									EndLine:   274,
									StartPos:  5218,
									EndPos:    5220,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 274,
								EndLine:   274,
								StartPos:  5223,
								EndPos:    5225,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 274,
									EndLine:   274,
									StartPos:  5223,
									EndPos:    5225,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 275,
						EndLine:   275,
						StartPos:  5229,
						EndPos:    5238,
					},
				},
				Expr: &ast.ExprBinaryEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 275,
							EndLine:   275,
							StartPos:  5229,
							EndPos:    5237,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 275,
								EndLine:   275,
								StartPos:  5229,
								EndPos:    5231,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 275,
									EndLine:   275,
									StartPos:  5229,
									EndPos:    5231,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 275,
								EndLine:   275,
								StartPos:  5235,
								EndPos:    5237,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 275,
									EndLine:   275,
									StartPos:  5235,
									EndPos:    5237,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 276,
						EndLine:   276,
						StartPos:  5241,
						EndPos:    5250,
					},
				},
				Expr: &ast.ExprBinaryGreaterOrEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 276,
							EndLine:   276,
							StartPos:  5241,
							EndPos:    5249,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 276,
								EndLine:   276,
								StartPos:  5241,
								EndPos:    5243,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 276,
									EndLine:   276,
									StartPos:  5241,
									EndPos:    5243,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 276,
								EndLine:   276,
								StartPos:  5247,
								EndPos:    5249,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 276,
									EndLine:   276,
									StartPos:  5247,
									EndPos:    5249,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 277,
						EndLine:   277,
						StartPos:  5253,
						EndPos:    5261,
					},
				},
				Expr: &ast.ExprBinaryGreater{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 277,
							EndLine:   277,
							StartPos:  5253,
							EndPos:    5260,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 277,
								EndLine:   277,
								StartPos:  5253,
								EndPos:    5255,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 277,
									EndLine:   277,
									StartPos:  5253,
									EndPos:    5255,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 277,
								EndLine:   277,
								StartPos:  5258,
								EndPos:    5260,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 277,
									EndLine:   277,
									StartPos:  5258,
									EndPos:    5260,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 278,
						EndLine:   278,
						StartPos:  5264,
						EndPos:    5274,
					},
				},
				Expr: &ast.ExprBinaryIdentical{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 278,
							EndLine:   278,
							StartPos:  5264,
							EndPos:    5273,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 278,
								EndLine:   278,
								StartPos:  5264,
								EndPos:    5266,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 278,
									EndLine:   278,
									StartPos:  5264,
									EndPos:    5266,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 278,
								EndLine:   278,
								StartPos:  5271,
								EndPos:    5273,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 278,
									EndLine:   278,
									StartPos:  5271,
									EndPos:    5273,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 279,
						EndLine:   279,
						StartPos:  5277,
						EndPos:    5287,
					},
				},
				Expr: &ast.ExprBinaryLogicalAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 279,
							EndLine:   279,
							StartPos:  5277,
							EndPos:    5286,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 279,
								EndLine:   279,
								StartPos:  5277,
								EndPos:    5279,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 279,
									EndLine:   279,
									StartPos:  5277,
									EndPos:    5279,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 279,
								EndLine:   279,
								StartPos:  5284,
								EndPos:    5286,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 279,
									EndLine:   279,
									StartPos:  5284,
									EndPos:    5286,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 280,
						EndLine:   280,
						StartPos:  5290,
						EndPos:    5299,
					},
				},
				Expr: &ast.ExprBinaryLogicalOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 280,
							EndLine:   280,
							StartPos:  5290,
							EndPos:    5298,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 280,
								EndLine:   280,
								StartPos:  5290,
								EndPos:    5292,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 280,
									EndLine:   280,
									StartPos:  5290,
									EndPos:    5292,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 280,
								EndLine:   280,
								StartPos:  5296,
								EndPos:    5298,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 280,
									EndLine:   280,
									StartPos:  5296,
									EndPos:    5298,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 281,
						EndLine:   281,
						StartPos:  5302,
						EndPos:    5312,
					},
				},
				Expr: &ast.ExprBinaryLogicalXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 281,
							EndLine:   281,
							StartPos:  5302,
							EndPos:    5311,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 281,
								EndLine:   281,
								StartPos:  5302,
								EndPos:    5304,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 281,
									EndLine:   281,
									StartPos:  5302,
									EndPos:    5304,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 281,
								EndLine:   281,
								StartPos:  5309,
								EndPos:    5311,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 281,
									EndLine:   281,
									StartPos:  5309,
									EndPos:    5311,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 282,
						EndLine:   282,
						StartPos:  5315,
						EndPos:    5323,
					},
				},
				Expr: &ast.ExprBinaryMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 282,
							EndLine:   282,
							StartPos:  5315,
							EndPos:    5322,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 282,
								EndLine:   282,
								StartPos:  5315,
								EndPos:    5317,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 282,
									EndLine:   282,
									StartPos:  5315,
									EndPos:    5317,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 282,
								EndLine:   282,
								StartPos:  5320,
								EndPos:    5322,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 282,
									EndLine:   282,
									StartPos:  5320,
									EndPos:    5322,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 283,
						EndLine:   283,
						StartPos:  5326,
						EndPos:    5334,
					},
				},
				Expr: &ast.ExprBinaryMod{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 283,
							EndLine:   283,
							StartPos:  5326,
							EndPos:    5333,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 283,
								EndLine:   283,
								StartPos:  5326,
								EndPos:    5328,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 283,
									EndLine:   283,
									StartPos:  5326,
									EndPos:    5328,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 283,
								EndLine:   283,
								StartPos:  5331,
								EndPos:    5333,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 283,
									EndLine:   283,
									StartPos:  5331,
									EndPos:    5333,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 284,
						EndLine:   284,
						StartPos:  5337,
						EndPos:    5345,
					},
				},
				Expr: &ast.ExprBinaryMul{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 284,
							EndLine:   284,
							StartPos:  5337,
							EndPos:    5344,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 284,
								EndLine:   284,
								StartPos:  5337,
								EndPos:    5339,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 284,
									EndLine:   284,
									StartPos:  5337,
									EndPos:    5339,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 284,
								EndLine:   284,
								StartPos:  5342,
								EndPos:    5344,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 284,
									EndLine:   284,
									StartPos:  5342,
									EndPos:    5344,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 285,
						EndLine:   285,
						StartPos:  5348,
						EndPos:    5357,
					},
				},
				Expr: &ast.ExprBinaryNotEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 285,
							EndLine:   285,
							StartPos:  5348,
							EndPos:    5356,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 285,
								EndLine:   285,
								StartPos:  5348,
								EndPos:    5350,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 285,
									EndLine:   285,
									StartPos:  5348,
									EndPos:    5350,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 285,
								EndLine:   285,
								StartPos:  5354,
								EndPos:    5356,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 285,
									EndLine:   285,
									StartPos:  5354,
									EndPos:    5356,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 286,
						EndLine:   286,
						StartPos:  5360,
						EndPos:    5370,
					},
				},
				Expr: &ast.ExprBinaryNotIdentical{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 286,
							EndLine:   286,
							StartPos:  5360,
							EndPos:    5369,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 286,
								EndLine:   286,
								StartPos:  5360,
								EndPos:    5362,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 286,
									EndLine:   286,
									StartPos:  5360,
									EndPos:    5362,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 286,
								EndLine:   286,
								StartPos:  5367,
								EndPos:    5369,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 286,
									EndLine:   286,
									StartPos:  5367,
									EndPos:    5369,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 287,
						EndLine:   287,
						StartPos:  5373,
						EndPos:    5381,
					},
				},
				Expr: &ast.ExprBinaryPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 287,
							EndLine:   287,
							StartPos:  5373,
							EndPos:    5380,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 287,
								EndLine:   287,
								StartPos:  5373,
								EndPos:    5375,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 287,
									EndLine:   287,
									StartPos:  5373,
									EndPos:    5375,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 287,
								EndLine:   287,
								StartPos:  5378,
								EndPos:    5380,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 287,
									EndLine:   287,
									StartPos:  5378,
									EndPos:    5380,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 288,
						EndLine:   288,
						StartPos:  5384,
						EndPos:    5393,
					},
				},
				Expr: &ast.ExprBinaryPow{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 288,
							EndLine:   288,
							StartPos:  5384,
							EndPos:    5392,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 288,
								EndLine:   288,
								StartPos:  5384,
								EndPos:    5386,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 288,
									EndLine:   288,
									StartPos:  5384,
									EndPos:    5386,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 288,
								EndLine:   288,
								StartPos:  5390,
								EndPos:    5392,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 288,
									EndLine:   288,
									StartPos:  5390,
									EndPos:    5392,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 289,
						EndLine:   289,
						StartPos:  5396,
						EndPos:    5405,
					},
				},
				Expr: &ast.ExprBinaryShiftLeft{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 289,
							EndLine:   289,
							StartPos:  5396,
							EndPos:    5404,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 289,
								EndLine:   289,
								StartPos:  5396,
								EndPos:    5398,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 289,
									EndLine:   289,
									StartPos:  5396,
									EndPos:    5398,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 289,
								EndLine:   289,
								StartPos:  5402,
								EndPos:    5404,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 289,
									EndLine:   289,
									StartPos:  5402,
									EndPos:    5404,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 290,
						EndLine:   290,
						StartPos:  5408,
						EndPos:    5417,
					},
				},
				Expr: &ast.ExprBinaryShiftRight{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 290,
							EndLine:   290,
							StartPos:  5408,
							EndPos:    5416,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 290,
								EndLine:   290,
								StartPos:  5408,
								EndPos:    5410,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 290,
									EndLine:   290,
									StartPos:  5408,
									EndPos:    5410,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 290,
								EndLine:   290,
								StartPos:  5414,
								EndPos:    5416,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 290,
									EndLine:   290,
									StartPos:  5414,
									EndPos:    5416,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 291,
						EndLine:   291,
						StartPos:  5420,
						EndPos:    5429,
					},
				},
				Expr: &ast.ExprBinarySmallerOrEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 291,
							EndLine:   291,
							StartPos:  5420,
							EndPos:    5428,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 291,
								EndLine:   291,
								StartPos:  5420,
								EndPos:    5422,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 291,
									EndLine:   291,
									StartPos:  5420,
									EndPos:    5422,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 291,
								EndLine:   291,
								StartPos:  5426,
								EndPos:    5428,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 291,
									EndLine:   291,
									StartPos:  5426,
									EndPos:    5428,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 292,
						EndLine:   292,
						StartPos:  5432,
						EndPos:    5440,
					},
				},
				Expr: &ast.ExprBinarySmaller{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 292,
							EndLine:   292,
							StartPos:  5432,
							EndPos:    5439,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 292,
								EndLine:   292,
								StartPos:  5432,
								EndPos:    5434,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 292,
									EndLine:   292,
									StartPos:  5432,
									EndPos:    5434,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 292,
								EndLine:   292,
								StartPos:  5437,
								EndPos:    5439,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 292,
									EndLine:   292,
									StartPos:  5437,
									EndPos:    5439,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 294,
						EndLine:   294,
						StartPos:  5444,
						EndPos:    5453,
					},
				},
				Expr: &ast.ExprAssignReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 294,
							EndLine:   294,
							StartPos:  5444,
							EndPos:    5452,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  5444,
								EndPos:    5446,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 294,
									EndLine:   294,
									StartPos:  5444,
									EndPos:    5446,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  5450,
								EndPos:    5452,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 294,
									EndLine:   294,
									StartPos:  5450,
									EndPos:    5452,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 295,
						EndLine:   295,
						StartPos:  5456,
						EndPos:    5470,
					},
				},
				Expr: &ast.ExprAssignReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 295,
							EndLine:   295,
							StartPos:  5456,
							EndPos:    5469,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 295,
								EndLine:   295,
								StartPos:  5456,
								EndPos:    5458,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 295,
									EndLine:   295,
									StartPos:  5456,
									EndPos:    5458,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprNew{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 295,
								EndLine:   295,
								StartPos:  5462,
								EndPos:    5469,
							},
						},
						Class: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 295,
									EndLine:   295,
									StartPos:  5466,
									EndPos:    5469,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 295,
											EndLine:   295,
											StartPos:  5466,
											EndPos:    5469,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 296,
						EndLine:   296,
						StartPos:  5473,
						EndPos:    5491,
					},
				},
				Expr: &ast.ExprAssignReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 296,
							EndLine:   296,
							StartPos:  5473,
							EndPos:    5490,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  5473,
								EndPos:    5475,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 296,
									EndLine:   296,
									StartPos:  5473,
									EndPos:    5475,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprNew{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  5479,
								EndPos:    5490,
							},
						},
						Class: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 296,
									EndLine:   296,
									StartPos:  5483,
									EndPos:    5486,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 296,
											EndLine:   296,
											StartPos:  5483,
											EndPos:    5486,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 296,
									EndLine:   296,
									StartPos:  5486,
									EndPos:    5490,
								},
							},
							Arguments: []ast.Vertex{
								&ast.Argument{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 296,
											EndLine:   296,
											StartPos:  5487,
											EndPos:    5489,
										},
									},
									IsReference: false,
									Variadic:    false,
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 296,
												EndLine:   296,
												StartPos:  5487,
												EndPos:    5489,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 296,
													EndLine:   296,
													StartPos:  5487,
													EndPos:    5489,
												},
											},
											Value: []byte("b"),
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 297,
						EndLine:   297,
						StartPos:  5494,
						EndPos:    5502,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 297,
							EndLine:   297,
							StartPos:  5494,
							EndPos:    5501,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 297,
								EndLine:   297,
								StartPos:  5494,
								EndPos:    5496,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 297,
									EndLine:   297,
									StartPos:  5494,
									EndPos:    5496,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 297,
								EndLine:   297,
								StartPos:  5499,
								EndPos:    5501,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 297,
									EndLine:   297,
									StartPos:  5499,
									EndPos:    5501,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 298,
						EndLine:   298,
						StartPos:  5505,
						EndPos:    5514,
					},
				},
				Expr: &ast.ExprAssignBitwiseAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 298,
							EndLine:   298,
							StartPos:  5505,
							EndPos:    5513,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  5505,
								EndPos:    5507,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 298,
									EndLine:   298,
									StartPos:  5505,
									EndPos:    5507,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  5511,
								EndPos:    5513,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 298,
									EndLine:   298,
									StartPos:  5511,
									EndPos:    5513,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 299,
						EndLine:   299,
						StartPos:  5517,
						EndPos:    5526,
					},
				},
				Expr: &ast.ExprAssignBitwiseOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 299,
							EndLine:   299,
							StartPos:  5517,
							EndPos:    5525,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5517,
								EndPos:    5519,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 299,
									EndLine:   299,
									StartPos:  5517,
									EndPos:    5519,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5523,
								EndPos:    5525,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 299,
									EndLine:   299,
									StartPos:  5523,
									EndPos:    5525,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 300,
						EndLine:   300,
						StartPos:  5529,
						EndPos:    5538,
					},
				},
				Expr: &ast.ExprAssignBitwiseXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 300,
							EndLine:   300,
							StartPos:  5529,
							EndPos:    5537,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5529,
								EndPos:    5531,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 300,
									EndLine:   300,
									StartPos:  5529,
									EndPos:    5531,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5535,
								EndPos:    5537,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 300,
									EndLine:   300,
									StartPos:  5535,
									EndPos:    5537,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 301,
						EndLine:   301,
						StartPos:  5541,
						EndPos:    5550,
					},
				},
				Expr: &ast.ExprAssignConcat{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 301,
							EndLine:   301,
							StartPos:  5541,
							EndPos:    5549,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 301,
								EndLine:   301,
								StartPos:  5541,
								EndPos:    5543,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 301,
									EndLine:   301,
									StartPos:  5541,
									EndPos:    5543,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 301,
								EndLine:   301,
								StartPos:  5547,
								EndPos:    5549,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 301,
									EndLine:   301,
									StartPos:  5547,
									EndPos:    5549,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 302,
						EndLine:   302,
						StartPos:  5553,
						EndPos:    5562,
					},
				},
				Expr: &ast.ExprAssignDiv{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 302,
							EndLine:   302,
							StartPos:  5553,
							EndPos:    5561,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 302,
								EndLine:   302,
								StartPos:  5553,
								EndPos:    5555,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 302,
									EndLine:   302,
									StartPos:  5553,
									EndPos:    5555,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 302,
								EndLine:   302,
								StartPos:  5559,
								EndPos:    5561,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 302,
									EndLine:   302,
									StartPos:  5559,
									EndPos:    5561,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 303,
						EndLine:   303,
						StartPos:  5565,
						EndPos:    5574,
					},
				},
				Expr: &ast.ExprAssignMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 303,
							EndLine:   303,
							StartPos:  5565,
							EndPos:    5573,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 303,
								EndLine:   303,
								StartPos:  5565,
								EndPos:    5567,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 303,
									EndLine:   303,
									StartPos:  5565,
									EndPos:    5567,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 303,
								EndLine:   303,
								StartPos:  5571,
								EndPos:    5573,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 303,
									EndLine:   303,
									StartPos:  5571,
									EndPos:    5573,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 304,
						EndLine:   304,
						StartPos:  5577,
						EndPos:    5586,
					},
				},
				Expr: &ast.ExprAssignMod{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 304,
							EndLine:   304,
							StartPos:  5577,
							EndPos:    5585,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 304,
								EndLine:   304,
								StartPos:  5577,
								EndPos:    5579,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 304,
									EndLine:   304,
									StartPos:  5577,
									EndPos:    5579,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 304,
								EndLine:   304,
								StartPos:  5583,
								EndPos:    5585,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 304,
									EndLine:   304,
									StartPos:  5583,
									EndPos:    5585,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 305,
						EndLine:   305,
						StartPos:  5589,
						EndPos:    5598,
					},
				},
				Expr: &ast.ExprAssignMul{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 305,
							EndLine:   305,
							StartPos:  5589,
							EndPos:    5597,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 305,
								EndLine:   305,
								StartPos:  5589,
								EndPos:    5591,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 305,
									EndLine:   305,
									StartPos:  5589,
									EndPos:    5591,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 305,
								EndLine:   305,
								StartPos:  5595,
								EndPos:    5597,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 305,
									EndLine:   305,
									StartPos:  5595,
									EndPos:    5597,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 306,
						EndLine:   306,
						StartPos:  5601,
						EndPos:    5610,
					},
				},
				Expr: &ast.ExprAssignPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 306,
							EndLine:   306,
							StartPos:  5601,
							EndPos:    5609,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 306,
								EndLine:   306,
								StartPos:  5601,
								EndPos:    5603,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 306,
									EndLine:   306,
									StartPos:  5601,
									EndPos:    5603,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 306,
								EndLine:   306,
								StartPos:  5607,
								EndPos:    5609,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 306,
									EndLine:   306,
									StartPos:  5607,
									EndPos:    5609,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 307,
						EndLine:   307,
						StartPos:  5613,
						EndPos:    5623,
					},
				},
				Expr: &ast.ExprAssignPow{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 307,
							EndLine:   307,
							StartPos:  5613,
							EndPos:    5622,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 307,
								EndLine:   307,
								StartPos:  5613,
								EndPos:    5615,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 307,
									EndLine:   307,
									StartPos:  5613,
									EndPos:    5615,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 307,
								EndLine:   307,
								StartPos:  5620,
								EndPos:    5622,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 307,
									EndLine:   307,
									StartPos:  5620,
									EndPos:    5622,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 308,
						EndLine:   308,
						StartPos:  5626,
						EndPos:    5636,
					},
				},
				Expr: &ast.ExprAssignShiftLeft{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 308,
							EndLine:   308,
							StartPos:  5626,
							EndPos:    5635,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 308,
								EndLine:   308,
								StartPos:  5626,
								EndPos:    5628,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 308,
									EndLine:   308,
									StartPos:  5626,
									EndPos:    5628,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 308,
								EndLine:   308,
								StartPos:  5633,
								EndPos:    5635,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 308,
									EndLine:   308,
									StartPos:  5633,
									EndPos:    5635,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 309,
						EndLine:   309,
						StartPos:  5639,
						EndPos:    5649,
					},
				},
				Expr: &ast.ExprAssignShiftRight{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 309,
							EndLine:   309,
							StartPos:  5639,
							EndPos:    5648,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 309,
								EndLine:   309,
								StartPos:  5639,
								EndPos:    5641,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 309,
									EndLine:   309,
									StartPos:  5639,
									EndPos:    5641,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 309,
								EndLine:   309,
								StartPos:  5646,
								EndPos:    5648,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 309,
									EndLine:   309,
									StartPos:  5646,
									EndPos:    5648,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 312,
						EndLine:   312,
						StartPos:  5655,
						EndPos:    5667,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 312,
							EndLine:   312,
							StartPos:  5655,
							EndPos:    5665,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 312,
								EndLine:   312,
								StartPos:  5659,
								EndPos:    5663,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 312,
										EndLine:   312,
										StartPos:  5660,
										EndPos:    5663,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 312,
								EndLine:   312,
								StartPos:  5663,
								EndPos:    5665,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 313,
						EndLine:   313,
						StartPos:  5691,
						EndPos:    5695,
					},
				},
				Expr: &ast.ExprPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 313,
							EndLine:   313,
							StartPos:  5691,
							EndPos:    5694,
						},
					},
					Var: &ast.ExprMethodCall{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5687,
								EndPos:    5689,
							},
						},
						Var: &ast.ExprNew{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 313,
									EndLine:   313,
									StartPos:  5671,
									EndPos:    5681,
								},
							},
							Class: &ast.NameFullyQualified{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 313,
										EndLine:   313,
										StartPos:  5675,
										EndPos:    5679,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 313,
												EndLine:   313,
												StartPos:  5676,
												EndPos:    5679,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 313,
										EndLine:   313,
										StartPos:  5679,
										EndPos:    5681,
									},
								},
							},
						},
						Method: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 313,
									EndLine:   313,
									StartPos:  5684,
									EndPos:    5687,
								},
							},
							Value: []byte("bar"),
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 313,
									EndLine:   313,
									StartPos:  5687,
									EndPos:    5689,
								},
							},
						},
					},
					Property: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5691,
								EndPos:    5694,
							},
						},
						Value: []byte("baz"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 314,
						EndLine:   314,
						StartPos:  5714,
						EndPos:    5717,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 314,
							EndLine:   314,
							StartPos:  5714,
							EndPos:    5715,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 314,
								EndLine:   314,
								StartPos:  5711,
								EndPos:    5712,
							},
						},
						Var: &ast.ExprNew{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 314,
									EndLine:   314,
									StartPos:  5699,
									EndPos:    5709,
								},
							},
							Class: &ast.NameFullyQualified{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 314,
										EndLine:   314,
										StartPos:  5703,
										EndPos:    5707,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 314,
												EndLine:   314,
												StartPos:  5704,
												EndPos:    5707,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 314,
										EndLine:   314,
										StartPos:  5707,
										EndPos:    5709,
									},
								},
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 314,
									EndLine:   314,
									StartPos:  5711,
									EndPos:    5712,
								},
							},
							Value: []byte("0"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 314,
								EndLine:   314,
								StartPos:  5714,
								EndPos:    5715,
							},
						},
						Value: []byte("0"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 315,
						EndLine:   315,
						StartPos:  5740,
						EndPos:    5743,
					},
				},
				Expr: &ast.ExprMethodCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 315,
							EndLine:   315,
							StartPos:  5740,
							EndPos:    5742,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5733,
								EndPos:    5734,
							},
						},
						Var: &ast.ExprNew{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 315,
									EndLine:   315,
									StartPos:  5721,
									EndPos:    5731,
								},
							},
							Class: &ast.NameFullyQualified{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 315,
										EndLine:   315,
										StartPos:  5725,
										EndPos:    5729,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 315,
												EndLine:   315,
												StartPos:  5726,
												EndPos:    5729,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 315,
										EndLine:   315,
										StartPos:  5729,
										EndPos:    5731,
									},
								},
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 315,
									EndLine:   315,
									StartPos:  5733,
									EndPos:    5734,
								},
							},
							Value: []byte("0"),
						},
					},
					Method: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5737,
								EndPos:    5740,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5740,
								EndPos:    5742,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 317,
						EndLine:   317,
						StartPos:  5747,
						EndPos:    5764,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 317,
							EndLine:   317,
							StartPos:  5747,
							EndPos:    5763,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 317,
								EndLine:   317,
								StartPos:  5747,
								EndPos:    5760,
							},
						},
						Var: &ast.ExprArray{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 317,
									EndLine:   317,
									StartPos:  5747,
									EndPos:    5757,
								},
							},
							Items: []ast.Vertex{
								&ast.ExprArrayItem{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 317,
											EndLine:   317,
											StartPos:  5753,
											EndPos:    5756,
										},
									},
									Val: &ast.ExprShortArray{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 317,
												EndLine:   317,
												StartPos:  5753,
												EndPos:    5756,
											},
										},
										Items: []ast.Vertex{
											&ast.ExprArrayItem{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 317,
														EndLine:   317,
														StartPos:  5754,
														EndPos:    5755,
													},
												},
												Val: &ast.ScalarLnumber{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 317,
															EndLine:   317,
															StartPos:  5754,
															EndPos:    5755,
														},
													},
													Value: []byte("0"),
												},
											},
										},
									},
								},
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 317,
									EndLine:   317,
									StartPos:  5758,
									EndPos:    5759,
								},
							},
							Value: []byte("0"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 317,
								EndLine:   317,
								StartPos:  5761,
								EndPos:    5762,
							},
						},
						Value: []byte("0"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 318,
						EndLine:   318,
						StartPos:  5767,
						EndPos:    5776,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 318,
							EndLine:   318,
							StartPos:  5767,
							EndPos:    5775,
						},
					},
					Var: &ast.ScalarString{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 318,
								EndLine:   318,
								StartPos:  5767,
								EndPos:    5772,
							},
						},
						Value: []byte("\"foo\""),
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 318,
								EndLine:   318,
								StartPos:  5773,
								EndPos:    5774,
							},
						},
						Value: []byte("0"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 319,
						EndLine:   319,
						StartPos:  5779,
						EndPos:    5786,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 319,
							EndLine:   319,
							StartPos:  5779,
							EndPos:    5785,
						},
					},
					Var: &ast.ExprConstFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 319,
								EndLine:   319,
								StartPos:  5779,
								EndPos:    5782,
							},
						},
						Const: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 319,
									EndLine:   319,
									StartPos:  5779,
									EndPos:    5782,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 319,
											EndLine:   319,
											StartPos:  5779,
											EndPos:    5782,
										},
									},
									Value: []byte("foo"),
								},
							},
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 319,
								EndLine:   319,
								StartPos:  5783,
								EndPos:    5784,
							},
						},
						Value: []byte("0"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 320,
						EndLine:   320,
						StartPos:  5789,
						EndPos:    5801,
					},
				},
				Expr: &ast.ExprClassConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 320,
							EndLine:   320,
							StartPos:  5789,
							EndPos:    5800,
						},
					},
					Class: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 320,
								EndLine:   320,
								StartPos:  5789,
								EndPos:    5795,
							},
						},
						Value: []byte("static"),
					},
					ConstantName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 320,
								EndLine:   320,
								StartPos:  5797,
								EndPos:    5800,
							},
						},
						Value: []byte("foo"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 322,
						EndLine:   322,
						StartPos:  5805,
						EndPos:    5814,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 322,
							EndLine:   322,
							StartPos:  5805,
							EndPos:    5813,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 322,
								EndLine:   322,
								StartPos:  5809,
								EndPos:    5813,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 322,
									EndLine:   322,
									StartPos:  5809,
									EndPos:    5813,
								},
							},
							Value: []byte("foo"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 323,
						EndLine:   323,
						StartPos:  5817,
						EndPos:    5832,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 323,
							EndLine:   323,
							StartPos:  5817,
							EndPos:    5831,
						},
					},
					Class: &ast.ExprStaticPropertyFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 323,
								EndLine:   323,
								StartPos:  5821,
								EndPos:    5831,
							},
						},
						Class: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 323,
									EndLine:   323,
									StartPos:  5821,
									EndPos:    5825,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 323,
										EndLine:   323,
										StartPos:  5821,
										EndPos:    5825,
									},
								},
								Value: []byte("foo"),
							},
						},
						Property: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 323,
									EndLine:   323,
									StartPos:  5827,
									EndPos:    5831,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 323,
										EndLine:   323,
										StartPos:  5827,
										EndPos:    5831,
									},
								},
								Value: []byte("bar"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 324,
						EndLine:   324,
						StartPos:  5835,
						EndPos:    5848,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 324,
							EndLine:   324,
							StartPos:  5835,
							EndPos:    5846,
						},
					},
					Class: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 324,
								EndLine:   324,
								StartPos:  5845,
								EndPos:    5846,
							},
						},
						Var: &ast.ExprPropertyFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 324,
									EndLine:   324,
									StartPos:  5843,
									EndPos:    5846,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 324,
										EndLine:   324,
										StartPos:  5839,
										EndPos:    5844,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 324,
											EndLine:   324,
											StartPos:  5839,
											EndPos:    5841,
										},
									},
									Value: []byte("a"),
								},
							},
							Property: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 324,
										EndLine:   324,
										StartPos:  5843,
										EndPos:    5844,
									},
								},
								Value: []byte("b"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 324,
									EndLine:   324,
									StartPos:  5845,
									EndPos:    5846,
								},
							},
							Value: []byte("0"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 325,
						EndLine:   325,
						StartPos:  5851,
						EndPos:    5883,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 325,
							EndLine:   325,
							StartPos:  5851,
							EndPos:    5881,
						},
					},
					Class: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 325,
								EndLine:   325,
								StartPos:  5880,
								EndPos:    5881,
							},
						},
						Var: &ast.ExprPropertyFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5878,
									EndPos:    5881,
								},
							},
							Var: &ast.ExprPropertyFetch{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 325,
										EndLine:   325,
										StartPos:  5874,
										EndPos:    5879,
									},
								},
								Var: &ast.ExprArrayDimFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 325,
											EndLine:   325,
											StartPos:  5861,
											EndPos:    5876,
										},
									},
									Var: &ast.ExprPropertyFetch{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 325,
												EndLine:   325,
												StartPos:  5859,
												EndPos:    5871,
											},
										},
										Var: &ast.ExprVariable{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 325,
													EndLine:   325,
													StartPos:  5855,
													EndPos:    5860,
												},
											},
											VarName: &ast.Identifier{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 325,
														EndLine:   325,
														StartPos:  5855,
														EndPos:    5857,
													},
												},
												Value: []byte("a"),
											},
										},
										Property: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 325,
													EndLine:   325,
													StartPos:  5859,
													EndPos:    5860,
												},
											},
											Value: []byte("b"),
										},
									},
									Dim: &ast.ExprTernary{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 325,
												EndLine:   325,
												StartPos:  5861,
												EndPos:    5871,
											},
										},
										Condition: &ast.ExprVariable{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 325,
													EndLine:   325,
													StartPos:  5861,
													EndPos:    5863,
												},
											},
											VarName: &ast.Identifier{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 325,
														EndLine:   325,
														StartPos:  5861,
														EndPos:    5863,
													},
												},
												Value: []byte("b"),
											},
										},
										IfFalse: &ast.ExprConstFetch{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 325,
													EndLine:   325,
													StartPos:  5867,
													EndPos:    5871,
												},
											},
											Const: &ast.NameName{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 325,
														EndLine:   325,
														StartPos:  5867,
														EndPos:    5871,
													},
												},
												Parts: []ast.Vertex{
													&ast.NameNamePart{
														Node: ast.Node{
															Position: &position.Position{
																StartLine: 325,
																EndLine:   325,
																StartPos:  5867,
																EndPos:    5871,
															},
														},
														Value: []byte("null"),
													},
												},
											},
										},
									},
								},
								Property: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 325,
											EndLine:   325,
											StartPos:  5874,
											EndPos:    5876,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 325,
												EndLine:   325,
												StartPos:  5874,
												EndPos:    5876,
											},
										},
										Value: []byte("c"),
									},
								},
							},
							Property: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 325,
										EndLine:   325,
										StartPos:  5878,
										EndPos:    5879,
									},
								},
								Value: []byte("d"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5880,
									EndPos:    5881,
								},
							},
							Value: []byte("0"),
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 325,
						EndLine:   325,
						StartPos:  5883,
						EndPos:    5902,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 325,
								EndLine:   325,
								StartPos:  5890,
								EndPos:    5901,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5890,
									EndPos:    5892,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 325,
										EndLine:   325,
										StartPos:  5890,
										EndPos:    5892,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5895,
									EndPos:    5901,
								},
							},
							Var: &ast.ExprShortArray{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 325,
										EndLine:   325,
										StartPos:  5895,
										EndPos:    5898,
									},
								},
								Items: []ast.Vertex{
									&ast.ExprArrayItem{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 325,
												EndLine:   325,
												StartPos:  5896,
												EndPos:    5897,
											},
										},
										Val: &ast.ScalarLnumber{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 325,
													EndLine:   325,
													StartPos:  5896,
													EndPos:    5897,
												},
											},
											Value: []byte("1"),
										},
									},
								},
							},
							Dim: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 325,
										EndLine:   325,
										StartPos:  5899,
										EndPos:    5900,
									},
								},
								Value: []byte("0"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 327,
						EndLine:   327,
						StartPos:  5906,
						EndPos:    5921,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 327,
								EndLine:   327,
								StartPos:  5913,
								EndPos:    5920,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 327,
									EndLine:   327,
									StartPos:  5913,
									EndPos:    5915,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 327,
										EndLine:   327,
										StartPos:  5913,
										EndPos:    5915,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBooleanNot{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 327,
									EndLine:   327,
									StartPos:  5918,
									EndPos:    5920,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 327,
										EndLine:   327,
										StartPos:  5919,
										EndPos:    5920,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 328,
						EndLine:   328,
						StartPos:  5924,
						EndPos:    5939,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5931,
								EndPos:    5938,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5931,
									EndPos:    5933,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 328,
										EndLine:   328,
										StartPos:  5931,
										EndPos:    5933,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBitwiseNot{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5936,
									EndPos:    5938,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 328,
										EndLine:   328,
										StartPos:  5937,
										EndPos:    5938,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 329,
						EndLine:   329,
						StartPos:  5942,
						EndPos:    5957,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5949,
								EndPos:    5956,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5949,
									EndPos:    5951,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 329,
										EndLine:   329,
										StartPos:  5949,
										EndPos:    5951,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprUnaryPlus{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5954,
									EndPos:    5956,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 329,
										EndLine:   329,
										StartPos:  5955,
										EndPos:    5956,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 330,
						EndLine:   330,
						StartPos:  5960,
						EndPos:    5975,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  5967,
								EndPos:    5974,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  5967,
									EndPos:    5969,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 330,
										EndLine:   330,
										StartPos:  5967,
										EndPos:    5969,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprUnaryMinus{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  5972,
									EndPos:    5974,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 330,
										EndLine:   330,
										StartPos:  5973,
										EndPos:    5974,
									},
								},
								Value: []byte("1"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 331,
						EndLine:   331,
						StartPos:  5978,
						EndPos:    5994,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 331,
								EndLine:   331,
								StartPos:  5985,
								EndPos:    5992,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 331,
									EndLine:   331,
									StartPos:  5985,
									EndPos:    5987,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 331,
										EndLine:   331,
										StartPos:  5985,
										EndPos:    5987,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 331,
									EndLine:   331,
									StartPos:  5991,
									EndPos:    5992,
								},
							},
							Value: []byte("1"),
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 332,
						EndLine:   332,
						StartPos:  5997,
						EndPos:    6016,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 332,
								EndLine:   332,
								StartPos:  6004,
								EndPos:    6015,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 332,
									EndLine:   332,
									StartPos:  6004,
									EndPos:    6006,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 332,
										EndLine:   332,
										StartPos:  6004,
										EndPos:    6006,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprTernary{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 332,
									EndLine:   332,
									StartPos:  6009,
									EndPos:    6015,
								},
							},
							Condition: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 332,
										EndLine:   332,
										StartPos:  6009,
										EndPos:    6010,
									},
								},
								Value: []byte("1"),
							},
							IfFalse: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 332,
										EndLine:   332,
										StartPos:  6014,
										EndPos:    6015,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 333,
						EndLine:   333,
						StartPos:  6019,
						EndPos:    6041,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 333,
								EndLine:   333,
								StartPos:  6026,
								EndPos:    6040,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 333,
									EndLine:   333,
									StartPos:  6026,
									EndPos:    6028,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 333,
										EndLine:   333,
										StartPos:  6026,
										EndPos:    6028,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprTernary{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 333,
									EndLine:   333,
									StartPos:  6031,
									EndPos:    6040,
								},
							},
							Condition: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 333,
										EndLine:   333,
										StartPos:  6031,
										EndPos:    6032,
									},
								},
								Value: []byte("1"),
							},
							IfTrue: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 333,
										EndLine:   333,
										StartPos:  6035,
										EndPos:    6036,
									},
								},
								Value: []byte("2"),
							},
							IfFalse: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 333,
										EndLine:   333,
										StartPos:  6039,
										EndPos:    6040,
									},
								},
								Value: []byte("3"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 334,
						EndLine:   334,
						StartPos:  6044,
						EndPos:    6062,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 334,
								EndLine:   334,
								StartPos:  6051,
								EndPos:    6061,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 334,
									EndLine:   334,
									StartPos:  6051,
									EndPos:    6053,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 334,
										EndLine:   334,
										StartPos:  6051,
										EndPos:    6053,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryBitwiseAnd{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 334,
									EndLine:   334,
									StartPos:  6056,
									EndPos:    6061,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 334,
										EndLine:   334,
										StartPos:  6056,
										EndPos:    6057,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 334,
										EndLine:   334,
										StartPos:  6060,
										EndPos:    6061,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 335,
						EndLine:   335,
						StartPos:  6065,
						EndPos:    6083,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 335,
								EndLine:   335,
								StartPos:  6072,
								EndPos:    6082,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  6072,
									EndPos:    6074,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 335,
										EndLine:   335,
										StartPos:  6072,
										EndPos:    6074,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryBitwiseOr{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  6077,
									EndPos:    6082,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 335,
										EndLine:   335,
										StartPos:  6077,
										EndPos:    6078,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 335,
										EndLine:   335,
										StartPos:  6081,
										EndPos:    6082,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 336,
						EndLine:   336,
						StartPos:  6086,
						EndPos:    6104,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 336,
								EndLine:   336,
								StartPos:  6093,
								EndPos:    6103,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 336,
									EndLine:   336,
									StartPos:  6093,
									EndPos:    6095,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 336,
										EndLine:   336,
										StartPos:  6093,
										EndPos:    6095,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryBitwiseXor{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 336,
									EndLine:   336,
									StartPos:  6098,
									EndPos:    6103,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 336,
										EndLine:   336,
										StartPos:  6098,
										EndPos:    6099,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 336,
										EndLine:   336,
										StartPos:  6102,
										EndPos:    6103,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 337,
						EndLine:   337,
						StartPos:  6107,
						EndPos:    6126,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 337,
								EndLine:   337,
								StartPos:  6114,
								EndPos:    6125,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 337,
									EndLine:   337,
									StartPos:  6114,
									EndPos:    6116,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 337,
										EndLine:   337,
										StartPos:  6114,
										EndPos:    6116,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryBooleanAnd{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 337,
									EndLine:   337,
									StartPos:  6119,
									EndPos:    6125,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 337,
										EndLine:   337,
										StartPos:  6119,
										EndPos:    6120,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 337,
										EndLine:   337,
										StartPos:  6124,
										EndPos:    6125,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 338,
						EndLine:   338,
						StartPos:  6129,
						EndPos:    6148,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 338,
								EndLine:   338,
								StartPos:  6136,
								EndPos:    6147,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 338,
									EndLine:   338,
									StartPos:  6136,
									EndPos:    6138,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 338,
										EndLine:   338,
										StartPos:  6136,
										EndPos:    6138,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryBooleanOr{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 338,
									EndLine:   338,
									StartPos:  6141,
									EndPos:    6147,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 338,
										EndLine:   338,
										StartPos:  6141,
										EndPos:    6142,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 338,
										EndLine:   338,
										StartPos:  6146,
										EndPos:    6147,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 339,
						EndLine:   339,
						StartPos:  6151,
						EndPos:    6169,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 339,
								EndLine:   339,
								StartPos:  6158,
								EndPos:    6168,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 339,
									EndLine:   339,
									StartPos:  6158,
									EndPos:    6160,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 339,
										EndLine:   339,
										StartPos:  6158,
										EndPos:    6160,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryConcat{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 339,
									EndLine:   339,
									StartPos:  6163,
									EndPos:    6168,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 339,
										EndLine:   339,
										StartPos:  6163,
										EndPos:    6164,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 339,
										EndLine:   339,
										StartPos:  6167,
										EndPos:    6168,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 340,
						EndLine:   340,
						StartPos:  6172,
						EndPos:    6190,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 340,
								EndLine:   340,
								StartPos:  6179,
								EndPos:    6189,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 340,
									EndLine:   340,
									StartPos:  6179,
									EndPos:    6181,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 340,
										EndLine:   340,
										StartPos:  6179,
										EndPos:    6181,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryDiv{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 340,
									EndLine:   340,
									StartPos:  6184,
									EndPos:    6189,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 340,
										EndLine:   340,
										StartPos:  6184,
										EndPos:    6185,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 340,
										EndLine:   340,
										StartPos:  6188,
										EndPos:    6189,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 341,
						EndLine:   341,
						StartPos:  6193,
						EndPos:    6212,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  6200,
								EndPos:    6211,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  6200,
									EndPos:    6202,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 341,
										EndLine:   341,
										StartPos:  6200,
										EndPos:    6202,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryEqual{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  6205,
									EndPos:    6211,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 341,
										EndLine:   341,
										StartPos:  6205,
										EndPos:    6206,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 341,
										EndLine:   341,
										StartPos:  6210,
										EndPos:    6211,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 342,
						EndLine:   342,
						StartPos:  6215,
						EndPos:    6234,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 342,
								EndLine:   342,
								StartPos:  6222,
								EndPos:    6233,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 342,
									EndLine:   342,
									StartPos:  6222,
									EndPos:    6224,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 342,
										EndLine:   342,
										StartPos:  6222,
										EndPos:    6224,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryGreaterOrEqual{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 342,
									EndLine:   342,
									StartPos:  6227,
									EndPos:    6233,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 342,
										EndLine:   342,
										StartPos:  6227,
										EndPos:    6228,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 342,
										EndLine:   342,
										StartPos:  6232,
										EndPos:    6233,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 343,
						EndLine:   343,
						StartPos:  6237,
						EndPos:    6255,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 343,
								EndLine:   343,
								StartPos:  6244,
								EndPos:    6254,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 343,
									EndLine:   343,
									StartPos:  6244,
									EndPos:    6246,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 343,
										EndLine:   343,
										StartPos:  6244,
										EndPos:    6246,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryGreater{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 343,
									EndLine:   343,
									StartPos:  6249,
									EndPos:    6254,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 343,
										EndLine:   343,
										StartPos:  6249,
										EndPos:    6250,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 343,
										EndLine:   343,
										StartPos:  6253,
										EndPos:    6254,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 344,
						EndLine:   344,
						StartPos:  6258,
						EndPos:    6278,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 344,
								EndLine:   344,
								StartPos:  6265,
								EndPos:    6277,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 344,
									EndLine:   344,
									StartPos:  6265,
									EndPos:    6267,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 344,
										EndLine:   344,
										StartPos:  6265,
										EndPos:    6267,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryIdentical{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 344,
									EndLine:   344,
									StartPos:  6270,
									EndPos:    6277,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 344,
										EndLine:   344,
										StartPos:  6270,
										EndPos:    6271,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 344,
										EndLine:   344,
										StartPos:  6276,
										EndPos:    6277,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 345,
						EndLine:   345,
						StartPos:  6281,
						EndPos:    6301,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 345,
								EndLine:   345,
								StartPos:  6288,
								EndPos:    6300,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 345,
									EndLine:   345,
									StartPos:  6288,
									EndPos:    6290,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 345,
										EndLine:   345,
										StartPos:  6288,
										EndPos:    6290,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryLogicalAnd{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 345,
									EndLine:   345,
									StartPos:  6293,
									EndPos:    6300,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 345,
										EndLine:   345,
										StartPos:  6293,
										EndPos:    6294,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 345,
										EndLine:   345,
										StartPos:  6299,
										EndPos:    6300,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 346,
						EndLine:   346,
						StartPos:  6304,
						EndPos:    6323,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 346,
								EndLine:   346,
								StartPos:  6311,
								EndPos:    6322,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  6311,
									EndPos:    6313,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  6311,
										EndPos:    6313,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryLogicalOr{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  6316,
									EndPos:    6322,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  6316,
										EndPos:    6317,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  6321,
										EndPos:    6322,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 347,
						EndLine:   347,
						StartPos:  6326,
						EndPos:    6346,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 347,
								EndLine:   347,
								StartPos:  6333,
								EndPos:    6345,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 347,
									EndLine:   347,
									StartPos:  6333,
									EndPos:    6335,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 347,
										EndLine:   347,
										StartPos:  6333,
										EndPos:    6335,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryLogicalXor{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 347,
									EndLine:   347,
									StartPos:  6338,
									EndPos:    6345,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 347,
										EndLine:   347,
										StartPos:  6338,
										EndPos:    6339,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 347,
										EndLine:   347,
										StartPos:  6344,
										EndPos:    6345,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 348,
						EndLine:   348,
						StartPos:  6349,
						EndPos:    6367,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 348,
								EndLine:   348,
								StartPos:  6356,
								EndPos:    6366,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 348,
									EndLine:   348,
									StartPos:  6356,
									EndPos:    6358,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 348,
										EndLine:   348,
										StartPos:  6356,
										EndPos:    6358,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryMinus{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 348,
									EndLine:   348,
									StartPos:  6361,
									EndPos:    6366,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 348,
										EndLine:   348,
										StartPos:  6361,
										EndPos:    6362,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 348,
										EndLine:   348,
										StartPos:  6365,
										EndPos:    6366,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 349,
						EndLine:   349,
						StartPos:  6370,
						EndPos:    6388,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 349,
								EndLine:   349,
								StartPos:  6377,
								EndPos:    6387,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 349,
									EndLine:   349,
									StartPos:  6377,
									EndPos:    6379,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 349,
										EndLine:   349,
										StartPos:  6377,
										EndPos:    6379,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryMod{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 349,
									EndLine:   349,
									StartPos:  6382,
									EndPos:    6387,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 349,
										EndLine:   349,
										StartPos:  6382,
										EndPos:    6383,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 349,
										EndLine:   349,
										StartPos:  6386,
										EndPos:    6387,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 350,
						EndLine:   350,
						StartPos:  6391,
						EndPos:    6409,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 350,
								EndLine:   350,
								StartPos:  6398,
								EndPos:    6408,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 350,
									EndLine:   350,
									StartPos:  6398,
									EndPos:    6400,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 350,
										EndLine:   350,
										StartPos:  6398,
										EndPos:    6400,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryMul{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 350,
									EndLine:   350,
									StartPos:  6403,
									EndPos:    6408,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 350,
										EndLine:   350,
										StartPos:  6403,
										EndPos:    6404,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 350,
										EndLine:   350,
										StartPos:  6407,
										EndPos:    6408,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 351,
						EndLine:   351,
						StartPos:  6412,
						EndPos:    6431,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 351,
								EndLine:   351,
								StartPos:  6419,
								EndPos:    6430,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 351,
									EndLine:   351,
									StartPos:  6419,
									EndPos:    6421,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 351,
										EndLine:   351,
										StartPos:  6419,
										EndPos:    6421,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryNotEqual{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 351,
									EndLine:   351,
									StartPos:  6424,
									EndPos:    6430,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 351,
										EndLine:   351,
										StartPos:  6424,
										EndPos:    6425,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 351,
										EndLine:   351,
										StartPos:  6429,
										EndPos:    6430,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 352,
						EndLine:   352,
						StartPos:  6434,
						EndPos:    6454,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 352,
								EndLine:   352,
								StartPos:  6441,
								EndPos:    6453,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 352,
									EndLine:   352,
									StartPos:  6441,
									EndPos:    6443,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 352,
										EndLine:   352,
										StartPos:  6441,
										EndPos:    6443,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryNotIdentical{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 352,
									EndLine:   352,
									StartPos:  6446,
									EndPos:    6453,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 352,
										EndLine:   352,
										StartPos:  6446,
										EndPos:    6447,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 352,
										EndLine:   352,
										StartPos:  6452,
										EndPos:    6453,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 353,
						EndLine:   353,
						StartPos:  6457,
						EndPos:    6475,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 353,
								EndLine:   353,
								StartPos:  6464,
								EndPos:    6474,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 353,
									EndLine:   353,
									StartPos:  6464,
									EndPos:    6466,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 353,
										EndLine:   353,
										StartPos:  6464,
										EndPos:    6466,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryPlus{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 353,
									EndLine:   353,
									StartPos:  6469,
									EndPos:    6474,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 353,
										EndLine:   353,
										StartPos:  6469,
										EndPos:    6470,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 353,
										EndLine:   353,
										StartPos:  6473,
										EndPos:    6474,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 354,
						EndLine:   354,
						StartPos:  6478,
						EndPos:    6497,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 354,
								EndLine:   354,
								StartPos:  6485,
								EndPos:    6496,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 354,
									EndLine:   354,
									StartPos:  6485,
									EndPos:    6487,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 354,
										EndLine:   354,
										StartPos:  6485,
										EndPos:    6487,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryPow{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 354,
									EndLine:   354,
									StartPos:  6490,
									EndPos:    6496,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 354,
										EndLine:   354,
										StartPos:  6490,
										EndPos:    6491,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 354,
										EndLine:   354,
										StartPos:  6495,
										EndPos:    6496,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 355,
						EndLine:   355,
						StartPos:  6500,
						EndPos:    6519,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 355,
								EndLine:   355,
								StartPos:  6507,
								EndPos:    6518,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 355,
									EndLine:   355,
									StartPos:  6507,
									EndPos:    6509,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 355,
										EndLine:   355,
										StartPos:  6507,
										EndPos:    6509,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryShiftLeft{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 355,
									EndLine:   355,
									StartPos:  6512,
									EndPos:    6518,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 355,
										EndLine:   355,
										StartPos:  6512,
										EndPos:    6513,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 355,
										EndLine:   355,
										StartPos:  6517,
										EndPos:    6518,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 356,
						EndLine:   356,
						StartPos:  6522,
						EndPos:    6541,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 356,
								EndLine:   356,
								StartPos:  6529,
								EndPos:    6540,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 356,
									EndLine:   356,
									StartPos:  6529,
									EndPos:    6531,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 356,
										EndLine:   356,
										StartPos:  6529,
										EndPos:    6531,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinaryShiftRight{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 356,
									EndLine:   356,
									StartPos:  6534,
									EndPos:    6540,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 356,
										EndLine:   356,
										StartPos:  6534,
										EndPos:    6535,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 356,
										EndLine:   356,
										StartPos:  6539,
										EndPos:    6540,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 357,
						EndLine:   357,
						StartPos:  6544,
						EndPos:    6563,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 357,
								EndLine:   357,
								StartPos:  6551,
								EndPos:    6562,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 357,
									EndLine:   357,
									StartPos:  6551,
									EndPos:    6553,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 357,
										EndLine:   357,
										StartPos:  6551,
										EndPos:    6553,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinarySmallerOrEqual{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 357,
									EndLine:   357,
									StartPos:  6556,
									EndPos:    6562,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 357,
										EndLine:   357,
										StartPos:  6556,
										EndPos:    6557,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 357,
										EndLine:   357,
										StartPos:  6561,
										EndPos:    6562,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 358,
						EndLine:   358,
						StartPos:  6566,
						EndPos:    6584,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 358,
								EndLine:   358,
								StartPos:  6573,
								EndPos:    6583,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 358,
									EndLine:   358,
									StartPos:  6573,
									EndPos:    6575,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 358,
										EndLine:   358,
										StartPos:  6573,
										EndPos:    6575,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprBinarySmaller{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 358,
									EndLine:   358,
									StartPos:  6578,
									EndPos:    6583,
								},
							},
							Left: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 358,
										EndLine:   358,
										StartPos:  6578,
										EndPos:    6579,
									},
								},
								Value: []byte("1"),
							},
							Right: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 358,
										EndLine:   358,
										StartPos:  6582,
										EndPos:    6583,
									},
								},
								Value: []byte("2"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 359,
						EndLine:   359,
						StartPos:  6587,
						EndPos:    6608,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 359,
								EndLine:   359,
								StartPos:  6594,
								EndPos:    6607,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 359,
									EndLine:   359,
									StartPos:  6594,
									EndPos:    6596,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 359,
										EndLine:   359,
										StartPos:  6594,
										EndPos:    6596,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprClassConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 359,
									EndLine:   359,
									StartPos:  6599,
									EndPos:    6607,
								},
							},
							Class: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 359,
										EndLine:   359,
										StartPos:  6599,
										EndPos:    6602,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 359,
												EndLine:   359,
												StartPos:  6599,
												EndPos:    6602,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
							ConstantName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 359,
										EndLine:   359,
										StartPos:  6604,
										EndPos:    6607,
									},
								},
								Value: []byte("bar"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 360,
						EndLine:   360,
						StartPos:  6611,
						EndPos:    6634,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 360,
								EndLine:   360,
								StartPos:  6618,
								EndPos:    6633,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 360,
									EndLine:   360,
									StartPos:  6618,
									EndPos:    6620,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 360,
										EndLine:   360,
										StartPos:  6618,
										EndPos:    6620,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprClassConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 360,
									EndLine:   360,
									StartPos:  6623,
									EndPos:    6633,
								},
							},
							Class: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 360,
										EndLine:   360,
										StartPos:  6623,
										EndPos:    6626,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 360,
												EndLine:   360,
												StartPos:  6623,
												EndPos:    6626,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
							ConstantName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 360,
										EndLine:   360,
										StartPos:  6628,
										EndPos:    6633,
									},
								},
								Value: []byte("class"),
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 361,
						EndLine:   361,
						StartPos:  6637,
						EndPos:    6659,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 361,
								EndLine:   361,
								StartPos:  6644,
								EndPos:    6658,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 361,
									EndLine:   361,
									StartPos:  6644,
									EndPos:    6646,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 361,
										EndLine:   361,
										StartPos:  6644,
										EndPos:    6646,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ScalarMagicConstant{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 361,
									EndLine:   361,
									StartPos:  6649,
									EndPos:    6658,
								},
							},
							Value: []byte("__CLASS__"),
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 362,
						EndLine:   362,
						StartPos:  6662,
						EndPos:    6678,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 362,
								EndLine:   362,
								StartPos:  6669,
								EndPos:    6677,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 362,
									EndLine:   362,
									StartPos:  6669,
									EndPos:    6671,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 362,
										EndLine:   362,
										StartPos:  6669,
										EndPos:    6671,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 362,
									EndLine:   362,
									StartPos:  6674,
									EndPos:    6677,
								},
							},
							Const: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 362,
										EndLine:   362,
										StartPos:  6674,
										EndPos:    6677,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 362,
												EndLine:   362,
												StartPos:  6674,
												EndPos:    6677,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 363,
						EndLine:   363,
						StartPos:  6681,
						EndPos:    6707,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 363,
								EndLine:   363,
								StartPos:  6688,
								EndPos:    6706,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 363,
									EndLine:   363,
									StartPos:  6688,
									EndPos:    6690,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 363,
										EndLine:   363,
										StartPos:  6688,
										EndPos:    6690,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 363,
									EndLine:   363,
									StartPos:  6693,
									EndPos:    6706,
								},
							},
							Const: &ast.NameRelative{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 363,
										EndLine:   363,
										StartPos:  6693,
										EndPos:    6706,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 363,
												EndLine:   363,
												StartPos:  6703,
												EndPos:    6706,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 364,
						EndLine:   364,
						StartPos:  6710,
						EndPos:    6727,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 364,
								EndLine:   364,
								StartPos:  6717,
								EndPos:    6726,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 364,
									EndLine:   364,
									StartPos:  6717,
									EndPos:    6719,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 364,
										EndLine:   364,
										StartPos:  6717,
										EndPos:    6719,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 364,
									EndLine:   364,
									StartPos:  6722,
									EndPos:    6726,
								},
							},
							Const: &ast.NameFullyQualified{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 364,
										EndLine:   364,
										StartPos:  6722,
										EndPos:    6726,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 364,
												EndLine:   364,
												StartPos:  6723,
												EndPos:    6726,
											},
										},
										Value: []byte("Foo"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 365,
						EndLine:   365,
						StartPos:  6730,
						EndPos:    6750,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 365,
								EndLine:   365,
								StartPos:  6737,
								EndPos:    6749,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 365,
									EndLine:   365,
									StartPos:  6737,
									EndPos:    6739,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 365,
										EndLine:   365,
										StartPos:  6737,
										EndPos:    6739,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprArray{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 365,
									EndLine:   365,
									StartPos:  6742,
									EndPos:    6749,
								},
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 366,
						EndLine:   366,
						StartPos:  6753,
						EndPos:    6782,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 366,
								EndLine:   366,
								StartPos:  6760,
								EndPos:    6781,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 366,
									EndLine:   366,
									StartPos:  6760,
									EndPos:    6762,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 366,
										EndLine:   366,
										StartPos:  6760,
										EndPos:    6762,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprArray{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 366,
									EndLine:   366,
									StartPos:  6765,
									EndPos:    6781,
								},
							},
							Items: []ast.Vertex{
								&ast.ExprArrayItem{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 366,
											EndLine:   366,
											StartPos:  6771,
											EndPos:    6777,
										},
									},
									Key: &ast.ScalarLnumber{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 366,
												EndLine:   366,
												StartPos:  6771,
												EndPos:    6772,
											},
										},
										Value: []byte("1"),
									},
									Val: &ast.ScalarLnumber{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 366,
												EndLine:   366,
												StartPos:  6776,
												EndPos:    6777,
											},
										},
										Value: []byte("1"),
									},
								},
								&ast.ExprArrayItem{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 366,
											EndLine:   366,
											StartPos:  6779,
											EndPos:    6780,
										},
									},
									Val: &ast.ScalarLnumber{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 366,
												EndLine:   366,
												StartPos:  6779,
												EndPos:    6780,
											},
										},
										Value: []byte("2"),
									},
								},
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 367,
						EndLine:   367,
						StartPos:  6785,
						EndPos:    6812,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 367,
								EndLine:   367,
								StartPos:  6792,
								EndPos:    6811,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 367,
									EndLine:   367,
									StartPos:  6792,
									EndPos:    6794,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 367,
										EndLine:   367,
										StartPos:  6792,
										EndPos:    6794,
									},
								},
								Value: []byte("a"),
							},
						},
						Expr: &ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 367,
									EndLine:   367,
									StartPos:  6797,
									EndPos:    6811,
								},
							},
							Var: &ast.ExprShortArray{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 367,
										EndLine:   367,
										StartPos:  6797,
										EndPos:    6808,
									},
								},
								Items: []ast.Vertex{
									&ast.ExprArrayItem{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 367,
												EndLine:   367,
												StartPos:  6798,
												EndPos:    6799,
											},
										},
										Val: &ast.ScalarLnumber{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 367,
													EndLine:   367,
													StartPos:  6798,
													EndPos:    6799,
												},
											},
											Value: []byte("1"),
										},
									},
									&ast.ExprArrayItem{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 367,
												EndLine:   367,
												StartPos:  6801,
												EndPos:    6807,
											},
										},
										Key: &ast.ScalarLnumber{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 367,
													EndLine:   367,
													StartPos:  6801,
													EndPos:    6802,
												},
											},
											Value: []byte("2"),
										},
										Val: &ast.ScalarLnumber{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 367,
													EndLine:   367,
													StartPos:  6806,
													EndPos:    6807,
												},
											},
											Value: []byte("2"),
										},
									},
								},
							},
							Dim: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 367,
										EndLine:   367,
										StartPos:  6809,
										EndPos:    6810,
									},
								},
								Value: []byte("0"),
							},
						},
					},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 369,
						EndLine:   369,
						StartPos:  6816,
						EndPos:    6831,
					},
				},
				Cond: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 369,
							EndLine:   369,
							StartPos:  6820,
							EndPos:    6827,
						},
					},
					Value: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 369,
								EndLine:   369,
								StartPos:  6826,
								EndPos:    6827,
							},
						},
						Value: []byte("1"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 369,
							EndLine:   369,
							StartPos:  6829,
							EndPos:    6831,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 370,
						EndLine:   370,
						StartPos:  6834,
						EndPos:    6845,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 370,
							EndLine:   370,
							StartPos:  6834,
							EndPos:    6844,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 370,
								EndLine:   370,
								StartPos:  6834,
								EndPos:    6837,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 370,
										EndLine:   370,
										StartPos:  6834,
										EndPos:    6837,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 370,
								EndLine:   370,
								StartPos:  6839,
								EndPos:    6844,
							},
						},
						VarName: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 370,
									EndLine:   370,
									StartPos:  6840,
									EndPos:    6844,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 370,
										EndLine:   370,
										StartPos:  6840,
										EndPos:    6844,
									},
								},
								Value: []byte("bar"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 372,
						EndLine:   372,
						StartPos:  6849,
						EndPos:    6856,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 372,
							EndLine:   372,
							StartPos:  6849,
							EndPos:    6855,
						},
					},
					Function: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 372,
								EndLine:   372,
								StartPos:  6849,
								EndPos:    6853,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 372,
									EndLine:   372,
									StartPos:  6849,
									EndPos:    6853,
								},
							},
							Value: []byte("foo"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 372,
								EndLine:   372,
								StartPos:  6853,
								EndPos:    6855,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 373,
						EndLine:   373,
						StartPos:  6859,
						EndPos:    6872,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 373,
							EndLine:   373,
							StartPos:  6859,
							EndPos:    6871,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 373,
								EndLine:   373,
								StartPos:  6859,
								EndPos:    6868,
							},
						},
						Var: &ast.ExprFunctionCall{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 373,
									EndLine:   373,
									StartPos:  6859,
									EndPos:    6865,
								},
							},
							Function: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 373,
										EndLine:   373,
										StartPos:  6859,
										EndPos:    6863,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 373,
											EndLine:   373,
											StartPos:  6859,
											EndPos:    6863,
										},
									},
									Value: []byte("foo"),
								},
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 373,
										EndLine:   373,
										StartPos:  6863,
										EndPos:    6865,
									},
								},
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 373,
									EndLine:   373,
									StartPos:  6866,
									EndPos:    6867,
								},
							},
							Value: []byte("0"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 373,
								EndLine:   373,
								StartPos:  6869,
								EndPos:    6870,
							},
						},
						Value: []byte("0"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 374,
						EndLine:   374,
						StartPos:  6875,
						EndPos:    6882,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 374,
							EndLine:   374,
							StartPos:  6875,
							EndPos:    6881,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 374,
								EndLine:   374,
								StartPos:  6875,
								EndPos:    6877,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 374,
									EndLine:   374,
									StartPos:  6875,
									EndPos:    6877,
								},
							},
							Value: []byte("a"),
						},
					},
					Dim: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 374,
								EndLine:   374,
								StartPos:  6878,
								EndPos:    6880,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 374,
									EndLine:   374,
									StartPos:  6878,
									EndPos:    6880,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 375,
						EndLine:   375,
						StartPos:  6885,
						EndPos:    6891,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 375,
							EndLine:   375,
							StartPos:  6885,
							EndPos:    6890,
						},
					},
					VarName: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 375,
								EndLine:   375,
								StartPos:  6887,
								EndPos:    6889,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 375,
									EndLine:   375,
									StartPos:  6887,
									EndPos:    6889,
								},
							},
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 376,
						EndLine:   376,
						StartPos:  6894,
						EndPos:    6909,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 376,
							EndLine:   376,
							StartPos:  6894,
							EndPos:    6908,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 376,
								EndLine:   376,
								StartPos:  6894,
								EndPos:    6898,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 376,
									EndLine:   376,
									StartPos:  6894,
									EndPos:    6898,
								},
							},
							Value: []byte("foo"),
						},
					},
					Call: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 376,
								EndLine:   376,
								StartPos:  6900,
								EndPos:    6906,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 376,
									EndLine:   376,
									StartPos:  6901,
									EndPos:    6905,
								},
							},
							Value: []byte("bar"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 376,
								EndLine:   376,
								StartPos:  6906,
								EndPos:    6908,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 377,
						EndLine:   377,
						StartPos:  6912,
						EndPos:    6922,
					},
				},
				Expr: &ast.ExprClassConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 377,
							EndLine:   377,
							StartPos:  6912,
							EndPos:    6921,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 377,
								EndLine:   377,
								StartPos:  6912,
								EndPos:    6916,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 377,
									EndLine:   377,
									StartPos:  6912,
									EndPos:    6916,
								},
							},
							Value: []byte("foo"),
						},
					},
					ConstantName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 377,
								EndLine:   377,
								StartPos:  6918,
								EndPos:    6921,
							},
						},
						Value: []byte("bar"),
					},
				},
			},
			&ast.StmtHaltCompiler{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 379,
						EndLine:   379,
						StartPos:  6926,
						EndPos:    6944,
					},
				},
			},
		},
	}

	php5parser := php5.NewParser([]byte(src), "5.6")
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

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   10,
				StartPos:  5,
				EndPos:    70,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  5,
						EndPos:    12,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  5,
							EndPos:    11,
						},
					},
					Value: []byte("\"test\""),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  15,
						EndPos:    24,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  15,
							EndPos:    23,
						},
					},
					Value: []byte("\"\\$test\""),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   6,
						StartPos:  27,
						EndPos:    41,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   6,
							StartPos:  27,
							EndPos:    40,
						},
					},
					Value: []byte("\"\n\t\t\ttest\n\t\t\""),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   7,
						StartPos:  44,
						EndPos:    52,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   7,
							StartPos:  44,
							EndPos:    51,
						},
					},
					Value: []byte("'$test'"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 8,
						EndLine:   10,
						StartPos:  55,
						EndPos:    70,
					},
				},
				Expr: &ast.ScalarString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 8,
							EndLine:   10,
							StartPos:  55,
							EndPos:    69,
						},
					},
					Value: []byte("'\n\t\t\t$test\n\t\t'"),
				},
			},
		},
	}

	php5parser := php5.NewParser([]byte(src), "5.6")
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

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   15,
				StartPos:  5,
				EndPos:    120,
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   3,
						StartPos:  5,
						EndPos:    16,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   3,
							StartPos:  5,
							EndPos:    15,
						},
					},
					Label: []byte("<<<CAD\n"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   6,
						StartPos:  19,
						EndPos:    37,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   6,
							StartPos:  19,
							EndPos:    36,
						},
					},
					Label: []byte("<<<CAD\n"),
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  26,
									EndPos:    33,
								},
							},
							Value: []byte("\thello\n"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   9,
						StartPos:  40,
						EndPos:    60,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   9,
							StartPos:  40,
							EndPos:    59,
						},
					},
					Label: []byte("<<<\"CAD\"\n"),
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 8,
									EndLine:   8,
									StartPos:  49,
									EndPos:    56,
								},
							},
							Value: []byte("\thello\n"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 10,
						EndLine:   12,
						StartPos:  63,
						EndPos:    90,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 10,
							EndLine:   12,
							StartPos:  63,
							EndPos:    89,
						},
					},
					Label: []byte("<<<\"CAD\"\n"),
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  72,
									EndPos:    79,
								},
							},
							Value: []byte("\thello "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  79,
									EndPos:    85,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 11,
										EndLine:   11,
										StartPos:  79,
										EndPos:    85,
									},
								},
								Value: []byte("world"),
							},
						},
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  85,
									EndPos:    86,
								},
							},
							Value: []byte("\n"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 13,
						EndLine:   15,
						StartPos:  93,
						EndPos:    120,
					},
				},
				Expr: &ast.ScalarHeredoc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 13,
							EndLine:   15,
							StartPos:  93,
							EndPos:    119,
						},
					},
					Label: []byte("<<<'CAD'\n"),
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 14,
									EndLine:   14,
									StartPos:  102,
									EndPos:    116,
								},
							},
							Value: []byte("\thello $world\n"),
						},
					},
				},
			},
		},
	}

	php5parser := php5.NewParser([]byte(src), "5.6")
	php5parser.Parse()
	actual := php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPhp5ControlCharsErrors(t *testing.T) {
	src := "<?php \004 echo $b; \"$a[\005test]\";"

	expected := []*errors.Error{
		{
			Msg: "WARNING: Unexpected character in input: '\004' (ASCII=4)",
			Pos: &position.Position{StartLine: 1, EndLine: 1, StartPos: 6, EndPos: 7},
		},
		{
			Msg: "WARNING: Unexpected character in input: '\005' (ASCII=5)",
			Pos: &position.Position{StartLine: 1, EndLine: 1, StartPos: 21, EndPos: 22},
		},
	}

	php5parser := php5.NewParser([]byte(src), "5.6")
	php5parser.Parse()
	actual := php5parser.GetErrors()
	assert.DeepEqual(t, expected, actual)
}
