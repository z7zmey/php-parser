package php7_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/internal/php7"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/errors"
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

	expected := &ast.Root{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   348,
				StartPos:  5,
				EndPos:    6319,
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
								Variadic:    false,
								IsReference: false,
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
								IsReference: false,
								Variadic:    false,
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
								Variadic:    false,
								IsReference: false,
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
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 9,
						EndLine:   9,
						StartPos:  160,
						EndPos:    185,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 9,
							EndLine:   9,
							StartPos:  160,
							EndPos:    184,
						},
					},
					Class: &ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  164,
								EndPos:    184,
							},
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  170,
									EndPos:    181,
								},
							},
							Arguments: []ast.Vertex{
								&ast.Argument{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  171,
											EndPos:    173,
										},
									},
									Variadic:    false,
									IsReference: false,
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  171,
												EndPos:    173,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 9,
													EndLine:   9,
													StartPos:  171,
													EndPos:    173,
												},
											},
											Value: []byte("a"),
										},
									},
								},
								&ast.Argument{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 9,
											EndLine:   9,
											StartPos:  175,
											EndPos:    180,
										},
									},
									Variadic:    true,
									IsReference: false,
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 9,
												EndLine:   9,
												StartPos:  178,
												EndPos:    180,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 9,
													EndLine:   9,
													StartPos:  178,
													EndPos:    180,
												},
											},
											Value: []byte("b"),
										},
									},
								},
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 10,
						EndLine:   10,
						StartPos:  188,
						EndPos:    201,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 10,
							EndLine:   10,
							StartPos:  188,
							EndPos:    200,
						},
					},
					Class: &ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 10,
								EndLine:   10,
								StartPos:  192,
								EndPos:    200,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 11,
						EndLine:   11,
						StartPos:  204,
						EndPos:    213,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 11,
							EndLine:   11,
							StartPos:  204,
							EndPos:    212,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 11,
								EndLine:   11,
								StartPos:  208,
								EndPos:    212,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  208,
									EndPos:    212,
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
						StartLine: 12,
						EndLine:   12,
						StartPos:  216,
						EndPos:    228,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 12,
							EndLine:   12,
							StartPos:  216,
							EndPos:    227,
						},
					},
					Class: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 12,
								EndLine:   12,
								StartPos:  220,
								EndPos:    227,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 12,
									EndLine:   12,
									StartPos:  220,
									EndPos:    224,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 12,
										EndLine:   12,
										StartPos:  220,
										EndPos:    224,
									},
								},
								Value: []byte("foo"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 12,
									EndLine:   12,
									StartPos:  225,
									EndPos:    226,
								},
							},
							Value: []byte("1"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 13,
						EndLine:   13,
						StartPos:  231,
						EndPos:    246,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 13,
							EndLine:   13,
							StartPos:  231,
							EndPos:    245,
						},
					},
					Class: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 13,
								EndLine:   13,
								StartPos:  235,
								EndPos:    245,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 13,
									EndLine:   13,
									StartPos:  235,
									EndPos:    239,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 13,
										EndLine:   13,
										StartPos:  235,
										EndPos:    239,
									},
								},
								Value: []byte("foo"),
							},
						},
						Dim: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 13,
									EndLine:   13,
									StartPos:  240,
									EndPos:    244,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 13,
										EndLine:   13,
										StartPos:  240,
										EndPos:    244,
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
						StartLine: 14,
						EndLine:   14,
						StartPos:  249,
						EndPos:    263,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 14,
							EndLine:   14,
							StartPos:  249,
							EndPos:    262,
						},
					},
					Class: &ast.ExprPropertyFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 14,
								EndLine:   14,
								StartPos:  253,
								EndPos:    262,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 14,
									EndLine:   14,
									StartPos:  253,
									EndPos:    257,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 14,
										EndLine:   14,
										StartPos:  253,
										EndPos:    257,
									},
								},
								Value: []byte("foo"),
							},
						},
						Property: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 14,
									EndLine:   14,
									StartPos:  259,
									EndPos:    262,
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
						StartLine: 15,
						EndLine:   15,
						StartPos:  266,
						EndPos:    281,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 15,
							EndLine:   15,
							StartPos:  266,
							EndPos:    280,
						},
					},
					Class: &ast.ExprStaticPropertyFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 15,
								EndLine:   15,
								StartPos:  270,
								EndPos:    280,
							},
						},
						Class: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 15,
									EndLine:   15,
									StartPos:  270,
									EndPos:    274,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 15,
										EndLine:   15,
										StartPos:  270,
										EndPos:    274,
									},
								},
								Value: []byte("foo"),
							},
						},
						Property: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 15,
									EndLine:   15,
									StartPos:  276,
									EndPos:    280,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 15,
										EndLine:   15,
										StartPos:  276,
										EndPos:    280,
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
						StartLine: 16,
						EndLine:   16,
						StartPos:  284,
						EndPos:    301,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 16,
							EndLine:   16,
							StartPos:  284,
							EndPos:    300,
						},
					},
					Class: &ast.ExprStaticPropertyFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 16,
								EndLine:   16,
								StartPos:  288,
								EndPos:    300,
							},
						},
						Class: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 16,
									EndLine:   16,
									StartPos:  288,
									EndPos:    294,
								},
							},
							Value: []byte("static"),
						},
						Property: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 16,
									EndLine:   16,
									StartPos:  296,
									EndPos:    300,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 16,
										EndLine:   16,
										StartPos:  296,
										EndPos:    300,
									},
								},
								Value: []byte("bar"),
							},
						},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 18,
						EndLine:   18,
						StartPos:  305,
						EndPos:    350,
					},
				},
				ReturnsRef: false,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 18,
							EndLine:   18,
							StartPos:  314,
							EndPos:    317,
						},
					},
					Value: []byte("foo"),
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 18,
								EndLine:   18,
								StartPos:  318,
								EndPos:    332,
							},
						},
						ByRef:    false,
						Variadic: false,
						Type: &ast.Nullable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  318,
									EndPos:    322,
								},
							},
							Expr: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 18,
										EndLine:   18,
										StartPos:  319,
										EndPos:    322,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 18,
												EndLine:   18,
												StartPos:  319,
												EndPos:    322,
											},
										},
										Value: []byte("bar"),
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  323,
									EndPos:    327,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 18,
										EndLine:   18,
										StartPos:  323,
										EndPos:    327,
									},
								},
								Value: []byte("bar"),
							},
						},
						DefaultValue: &ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  328,
									EndPos:    332,
								},
							},
							Const: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 18,
										EndLine:   18,
										StartPos:  328,
										EndPos:    332,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 18,
												EndLine:   18,
												StartPos:  328,
												EndPos:    332,
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
								StartLine: 18,
								EndLine:   18,
								StartPos:  334,
								EndPos:    346,
							},
						},
						Variadic: true,
						ByRef:    true,
						Type: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  334,
									EndPos:    337,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 18,
											EndLine:   18,
											StartPos:  334,
											EndPos:    337,
										},
									},
									Value: []byte("baz"),
								},
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 18,
									EndLine:   18,
									StartPos:  342,
									EndPos:    346,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 18,
										EndLine:   18,
										StartPos:  342,
										EndPos:    346,
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
						StartLine: 19,
						EndLine:   19,
						StartPos:  353,
						EndPos:    417,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 19,
							EndLine:   19,
							StartPos:  359,
							EndPos:    362,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 19,
								EndLine:   19,
								StartPos:  364,
								EndPos:    416,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 19,
									EndLine:   19,
									StartPos:  380,
									EndPos:    383,
								},
							},
							Value: []byte("foo"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 19,
										EndLine:   19,
										StartPos:  364,
										EndPos:    370,
									},
								},
								Value: []byte("public"),
							},
						},
						Params: []ast.Vertex{
							&ast.Parameter{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 19,
										EndLine:   19,
										StartPos:  384,
										EndPos:    398,
									},
								},
								ByRef:    false,
								Variadic: false,
								Type: &ast.Nullable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  384,
											EndPos:    388,
										},
									},
									Expr: &ast.NameName{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 19,
												EndLine:   19,
												StartPos:  385,
												EndPos:    388,
											},
										},
										Parts: []ast.Vertex{
											&ast.NameNamePart{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 19,
														EndLine:   19,
														StartPos:  385,
														EndPos:    388,
													},
												},
												Value: []byte("bar"),
											},
										},
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  389,
											EndPos:    393,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 19,
												EndLine:   19,
												StartPos:  389,
												EndPos:    393,
											},
										},
										Value: []byte("bar"),
									},
								},
								DefaultValue: &ast.ExprConstFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  394,
											EndPos:    398,
										},
									},
									Const: &ast.NameName{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 19,
												EndLine:   19,
												StartPos:  394,
												EndPos:    398,
											},
										},
										Parts: []ast.Vertex{
											&ast.NameNamePart{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 19,
														EndLine:   19,
														StartPos:  394,
														EndPos:    398,
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
										StartLine: 19,
										EndLine:   19,
										StartPos:  400,
										EndPos:    412,
									},
								},
								ByRef:    true,
								Variadic: true,
								Type: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  400,
											EndPos:    403,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 19,
													EndLine:   19,
													StartPos:  400,
													EndPos:    403,
												},
											},
											Value: []byte("baz"),
										},
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 19,
											EndLine:   19,
											StartPos:  408,
											EndPos:    412,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 19,
												EndLine:   19,
												StartPos:  408,
												EndPos:    412,
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
									StartLine: 19,
									EndLine:   19,
									StartPos:  414,
									EndPos:    416,
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
						StartLine: 20,
						EndLine:   20,
						StartPos:  420,
						EndPos:    462,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 20,
							EndLine:   20,
							StartPos:  420,
							EndPos:    461,
						},
					},
					ReturnsRef: false,
					Static:     false,
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 20,
									EndLine:   20,
									StartPos:  429,
									EndPos:    443,
								},
							},
							ByRef:    false,
							Variadic: false,
							Type: &ast.Nullable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  429,
										EndPos:    433,
									},
								},
								Expr: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 20,
											EndLine:   20,
											StartPos:  430,
											EndPos:    433,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 20,
													EndLine:   20,
													StartPos:  430,
													EndPos:    433,
												},
											},
											Value: []byte("bar"),
										},
									},
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  434,
										EndPos:    438,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 20,
											EndLine:   20,
											StartPos:  434,
											EndPos:    438,
										},
									},
									Value: []byte("bar"),
								},
							},
							DefaultValue: &ast.ExprConstFetch{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  439,
										EndPos:    443,
									},
								},
								Const: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 20,
											EndLine:   20,
											StartPos:  439,
											EndPos:    443,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 20,
													EndLine:   20,
													StartPos:  439,
													EndPos:    443,
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
									StartLine: 20,
									EndLine:   20,
									StartPos:  445,
									EndPos:    457,
								},
							},
							ByRef:    true,
							Variadic: true,
							Type: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  445,
										EndPos:    448,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 20,
												EndLine:   20,
												StartPos:  445,
												EndPos:    448,
											},
										},
										Value: []byte("baz"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 20,
										EndLine:   20,
										StartPos:  453,
										EndPos:    457,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 20,
											EndLine:   20,
											StartPos:  453,
											EndPos:    457,
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
						StartLine: 21,
						EndLine:   21,
						StartPos:  465,
						EndPos:    514,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 21,
							EndLine:   21,
							StartPos:  465,
							EndPos:    513,
						},
					},
					ReturnsRef: false,
					Static:     true,
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 21,
									EndLine:   21,
									StartPos:  481,
									EndPos:    495,
								},
							},
							ByRef:    false,
							Variadic: false,
							Type: &ast.Nullable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  481,
										EndPos:    485,
									},
								},
								Expr: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 21,
											EndLine:   21,
											StartPos:  482,
											EndPos:    485,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 21,
													EndLine:   21,
													StartPos:  482,
													EndPos:    485,
												},
											},
											Value: []byte("bar"),
										},
									},
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  486,
										EndPos:    490,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 21,
											EndLine:   21,
											StartPos:  486,
											EndPos:    490,
										},
									},
									Value: []byte("bar"),
								},
							},
							DefaultValue: &ast.ExprConstFetch{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  491,
										EndPos:    495,
									},
								},
								Const: &ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 21,
											EndLine:   21,
											StartPos:  491,
											EndPos:    495,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 21,
													EndLine:   21,
													StartPos:  491,
													EndPos:    495,
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
									StartLine: 21,
									EndLine:   21,
									StartPos:  497,
									EndPos:    509,
								},
							},
							ByRef:    true,
							Variadic: true,
							Type: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  497,
										EndPos:    500,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 21,
												EndLine:   21,
												StartPos:  497,
												EndPos:    500,
											},
										},
										Value: []byte("baz"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 21,
										EndLine:   21,
										StartPos:  505,
										EndPos:    509,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 21,
											EndLine:   21,
											StartPos:  505,
											EndPos:    509,
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
						StartLine: 23,
						EndLine:   23,
						StartPos:  518,
						EndPos:    538,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 23,
							EndLine:   23,
							StartPos:  518,
							EndPos:    537,
						},
					},
					Value: []byte("1234567890123456789"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 24,
						EndLine:   24,
						StartPos:  541,
						EndPos:    562,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 24,
							EndLine:   24,
							StartPos:  541,
							EndPos:    561,
						},
					},
					Value: []byte("12345678901234567890"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 25,
						EndLine:   25,
						StartPos:  565,
						EndPos:    568,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 25,
							EndLine:   25,
							StartPos:  565,
							EndPos:    567,
						},
					},
					Value: []byte("0."),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 26,
						EndLine:   26,
						StartPos:  571,
						EndPos:    638,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 26,
							EndLine:   26,
							StartPos:  571,
							EndPos:    637,
						},
					},
					Value: []byte("0b0111111111111111111111111111111111111111111111111111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 27,
						EndLine:   27,
						StartPos:  641,
						EndPos:    708,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 27,
							EndLine:   27,
							StartPos:  641,
							EndPos:    707,
						},
					},
					Value: []byte("0b1111111111111111111111111111111111111111111111111111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 28,
						EndLine:   28,
						StartPos:  711,
						EndPos:    732,
					},
				},
				Expr: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 28,
							EndLine:   28,
							StartPos:  711,
							EndPos:    731,
						},
					},
					Value: []byte("0x007111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 29,
						EndLine:   29,
						StartPos:  735,
						EndPos:    754,
					},
				},
				Expr: &ast.ScalarDnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 29,
							EndLine:   29,
							StartPos:  735,
							EndPos:    753,
						},
					},
					Value: []byte("0x8111111111111111"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 30,
						EndLine:   30,
						StartPos:  757,
						EndPos:    767,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 30,
							EndLine:   30,
							StartPos:  757,
							EndPos:    766,
						},
					},
					Value: []byte("__CLASS__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 31,
						EndLine:   31,
						StartPos:  770,
						EndPos:    778,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 31,
							EndLine:   31,
							StartPos:  770,
							EndPos:    777,
						},
					},
					Value: []byte("__DIR__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 32,
						EndLine:   32,
						StartPos:  781,
						EndPos:    790,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 32,
							EndLine:   32,
							StartPos:  781,
							EndPos:    789,
						},
					},
					Value: []byte("__FILE__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 33,
						EndLine:   33,
						StartPos:  793,
						EndPos:    806,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 33,
							EndLine:   33,
							StartPos:  793,
							EndPos:    805,
						},
					},
					Value: []byte("__FUNCTION__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 34,
						EndLine:   34,
						StartPos:  809,
						EndPos:    818,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 34,
							EndLine:   34,
							StartPos:  809,
							EndPos:    817,
						},
					},
					Value: []byte("__LINE__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 35,
						EndLine:   35,
						StartPos:  821,
						EndPos:    835,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 35,
							EndLine:   35,
							StartPos:  821,
							EndPos:    834,
						},
					},
					Value: []byte("__NAMESPACE__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 36,
						EndLine:   36,
						StartPos:  838,
						EndPos:    849,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 36,
							EndLine:   36,
							StartPos:  838,
							EndPos:    848,
						},
					},
					Value: []byte("__METHOD__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 37,
						EndLine:   37,
						StartPos:  852,
						EndPos:    862,
					},
				},
				Expr: &ast.ScalarMagicConstant{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 37,
							EndLine:   37,
							StartPos:  852,
							EndPos:    861,
						},
					},
					Value: []byte("__TRAIT__"),
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 39,
						EndLine:   39,
						StartPos:  866,
						EndPos:    878,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 39,
							EndLine:   39,
							StartPos:  866,
							EndPos:    877,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 39,
									EndLine:   39,
									StartPos:  867,
									EndPos:    872,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 39,
									EndLine:   39,
									StartPos:  872,
									EndPos:    876,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 39,
										EndLine:   39,
										StartPos:  872,
										EndPos:    876,
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
						StartLine: 40,
						EndLine:   40,
						StartPos:  881,
						EndPos:    896,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 40,
							EndLine:   40,
							StartPos:  881,
							EndPos:    895,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 40,
									EndLine:   40,
									StartPos:  882,
									EndPos:    887,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 40,
									EndLine:   40,
									StartPos:  887,
									EndPos:    894,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 40,
										EndLine:   40,
										StartPos:  887,
										EndPos:    891,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 40,
											EndLine:   40,
											StartPos:  887,
											EndPos:    891,
										},
									},
									Value: []byte("var"),
								},
							},
							Dim: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 40,
										EndLine:   40,
										StartPos:  892,
										EndPos:    893,
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
						StartLine: 41,
						EndLine:   41,
						StartPos:  899,
						EndPos:    915,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 41,
							EndLine:   41,
							StartPos:  899,
							EndPos:    914,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 41,
									EndLine:   41,
									StartPos:  900,
									EndPos:    905,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 41,
									EndLine:   41,
									StartPos:  905,
									EndPos:    913,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 41,
										EndLine:   41,
										StartPos:  905,
										EndPos:    909,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 41,
											EndLine:   41,
											StartPos:  905,
											EndPos:    909,
										},
									},
									Value: []byte("var"),
								},
							},
							Dim: &ast.ExprUnaryMinus{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 41,
										EndLine:   41,
										StartPos:  910,
										EndPos:    912,
									},
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 41,
											EndLine:   41,
											StartPos:  910,
											EndPos:    912,
										},
									},
									Value: []byte("1"),
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 42,
						EndLine:   42,
						StartPos:  918,
						EndPos:    972,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 42,
							EndLine:   42,
							StartPos:  918,
							EndPos:    971,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 42,
									EndLine:   42,
									StartPos:  919,
									EndPos:    924,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 42,
									EndLine:   42,
									StartPos:  924,
									EndPos:    970,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 42,
										EndLine:   42,
										StartPos:  924,
										EndPos:    928,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 42,
											EndLine:   42,
											StartPos:  924,
											EndPos:    928,
										},
									},
									Value: []byte("var"),
								},
							},
							Dim: &ast.ScalarString{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 42,
										EndLine:   42,
										StartPos:  929,
										EndPos:    969,
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
						StartLine: 43,
						EndLine:   43,
						StartPos:  975,
						EndPos:    1030,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 43,
							EndLine:   43,
							StartPos:  975,
							EndPos:    1029,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 43,
									EndLine:   43,
									StartPos:  976,
									EndPos:    981,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 43,
									EndLine:   43,
									StartPos:  981,
									EndPos:    1028,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 43,
										EndLine:   43,
										StartPos:  981,
										EndPos:    985,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 43,
											EndLine:   43,
											StartPos:  981,
											EndPos:    985,
										},
									},
									Value: []byte("var"),
								},
							},
							Dim: &ast.ScalarString{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 43,
										EndLine:   43,
										StartPos:  986,
										EndPos:    1027,
									},
								},
								Value: []byte("-1234567890123456789012345678901234567890"),
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 44,
						EndLine:   44,
						StartPos:  1033,
						EndPos:    1050,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 44,
							EndLine:   44,
							StartPos:  1033,
							EndPos:    1049,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 44,
									EndLine:   44,
									StartPos:  1034,
									EndPos:    1039,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 44,
									EndLine:   44,
									StartPos:  1039,
									EndPos:    1048,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 44,
										EndLine:   44,
										StartPos:  1039,
										EndPos:    1043,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 44,
											EndLine:   44,
											StartPos:  1039,
											EndPos:    1043,
										},
									},
									Value: []byte("var"),
								},
							},
							Dim: &ast.ScalarString{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 44,
										EndLine:   44,
										StartPos:  1044,
										EndPos:    1047,
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
						StartLine: 45,
						EndLine:   45,
						StartPos:  1053,
						EndPos:    1071,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 45,
							EndLine:   45,
							StartPos:  1053,
							EndPos:    1070,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 45,
									EndLine:   45,
									StartPos:  1054,
									EndPos:    1059,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 45,
									EndLine:   45,
									StartPos:  1059,
									EndPos:    1069,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 45,
										EndLine:   45,
										StartPos:  1059,
										EndPos:    1063,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 45,
											EndLine:   45,
											StartPos:  1059,
											EndPos:    1063,
										},
									},
									Value: []byte("var"),
								},
							},
							Dim: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 45,
										EndLine:   45,
										StartPos:  1064,
										EndPos:    1068,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 45,
											EndLine:   45,
											StartPos:  1064,
											EndPos:    1068,
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
						StartLine: 46,
						EndLine:   46,
						StartPos:  1074,
						EndPos:    1086,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 46,
							EndLine:   46,
							StartPos:  1074,
							EndPos:    1085,
						},
					},
					Parts: []ast.Vertex{
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 46,
									EndLine:   46,
									StartPos:  1075,
									EndPos:    1079,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 46,
										EndLine:   46,
										StartPos:  1075,
										EndPos:    1079,
									},
								},
								Value: []byte("foo"),
							},
						},
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 46,
									EndLine:   46,
									StartPos:  1079,
									EndPos:    1080,
								},
							},
							Value: []byte(" "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 46,
									EndLine:   46,
									StartPos:  1080,
									EndPos:    1084,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 46,
										EndLine:   46,
										StartPos:  1080,
										EndPos:    1084,
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
						StartLine: 47,
						EndLine:   47,
						StartPos:  1089,
						EndPos:    1108,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 47,
							EndLine:   47,
							StartPos:  1089,
							EndPos:    1107,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 47,
									EndLine:   47,
									StartPos:  1090,
									EndPos:    1095,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprPropertyFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 47,
									EndLine:   47,
									StartPos:  1095,
									EndPos:    1104,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 47,
										EndLine:   47,
										StartPos:  1095,
										EndPos:    1099,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 47,
											EndLine:   47,
											StartPos:  1095,
											EndPos:    1099,
										},
									},
									Value: []byte("foo"),
								},
							},
							Property: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 47,
										EndLine:   47,
										StartPos:  1101,
										EndPos:    1104,
									},
								},
								Value: []byte("bar"),
							},
						},
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 47,
									EndLine:   47,
									StartPos:  1104,
									EndPos:    1106,
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
						StartLine: 48,
						EndLine:   48,
						StartPos:  1111,
						EndPos:    1125,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 48,
							EndLine:   48,
							StartPos:  1111,
							EndPos:    1124,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 48,
									EndLine:   48,
									StartPos:  1112,
									EndPos:    1117,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 48,
									EndLine:   48,
									StartPos:  1117,
									EndPos:    1123,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 48,
										EndLine:   48,
										StartPos:  1119,
										EndPos:    1122,
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
						StartLine: 49,
						EndLine:   49,
						StartPos:  1128,
						EndPos:    1145,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 49,
							EndLine:   49,
							StartPos:  1128,
							EndPos:    1144,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 49,
									EndLine:   49,
									StartPos:  1129,
									EndPos:    1134,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprArrayDimFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 49,
									EndLine:   49,
									StartPos:  1134,
									EndPos:    1143,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 49,
										EndLine:   49,
										StartPos:  1136,
										EndPos:    1139,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 49,
											EndLine:   49,
											StartPos:  1136,
											EndPos:    1139,
										},
									},
									Value: []byte("foo"),
								},
							},
							Dim: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 49,
										EndLine:   49,
										StartPos:  1140,
										EndPos:    1141,
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
						StartLine: 50,
						EndLine:   50,
						StartPos:  1148,
						EndPos:    1163,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 50,
							EndLine:   50,
							StartPos:  1148,
							EndPos:    1162,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 50,
									EndLine:   50,
									StartPos:  1149,
									EndPos:    1154,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 50,
									EndLine:   50,
									StartPos:  1154,
									EndPos:    1161,
								},
							},
							VarName: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 50,
										EndLine:   50,
										StartPos:  1156,
										EndPos:    1160,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 50,
											EndLine:   50,
											StartPos:  1156,
											EndPos:    1160,
										},
									},
									Value: []byte("foo"),
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 51,
						EndLine:   51,
						StartPos:  1166,
						EndPos:    1187,
					},
				},
				Expr: &ast.ScalarEncapsed{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 51,
							EndLine:   51,
							StartPos:  1166,
							EndPos:    1186,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 51,
									EndLine:   51,
									StartPos:  1167,
									EndPos:    1172,
								},
							},
							Value: []byte("test "),
						},
						&ast.ExprMethodCall{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 51,
									EndLine:   51,
									StartPos:  1173,
									EndPos:    1184,
								},
							},
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 51,
										EndLine:   51,
										StartPos:  1173,
										EndPos:    1177,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 51,
											EndLine:   51,
											StartPos:  1173,
											EndPos:    1177,
										},
									},
									Value: []byte("foo"),
								},
							},
							Method: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 51,
										EndLine:   51,
										StartPos:  1179,
										EndPos:    1182,
									},
								},
								Value: []byte("bar"),
							},
							ArgumentList: &ast.ArgumentList{
								Node: ast.Node{
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
			},
			&ast.StmtAltIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 53,
						EndLine:   54,
						StartPos:  1191,
						EndPos:    1209,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 53,
							EndLine:   53,
							StartPos:  1195,
							EndPos:    1197,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 53,
								EndLine:   53,
								StartPos:  1195,
								EndPos:    1197,
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
						StartLine: 55,
						EndLine:   57,
						StartPos:  1212,
						EndPos:    1245,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 55,
							EndLine:   55,
							StartPos:  1216,
							EndPos:    1218,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 55,
								EndLine:   55,
								StartPos:  1216,
								EndPos:    1218,
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
								StartLine: 56,
								EndLine:   -1,
								StartPos:  1224,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 56,
									EndLine:   56,
									StartPos:  1232,
									EndPos:    1234,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 56,
										EndLine:   56,
										StartPos:  1232,
										EndPos:    1234,
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
						StartLine: 58,
						EndLine:   60,
						StartPos:  1248,
						EndPos:    1274,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 58,
							EndLine:   58,
							StartPos:  1252,
							EndPos:    1254,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 58,
								EndLine:   58,
								StartPos:  1252,
								EndPos:    1254,
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
							StartLine: 59,
							EndLine:   -1,
							StartPos:  1260,
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
						StartLine: 61,
						EndLine:   65,
						StartPos:  1277,
						EndPos:    1333,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 61,
							EndLine:   61,
							StartPos:  1281,
							EndPos:    1283,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 61,
								EndLine:   61,
								StartPos:  1281,
								EndPos:    1283,
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
								StartLine: 62,
								EndLine:   -1,
								StartPos:  1289,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 62,
									EndLine:   62,
									StartPos:  1297,
									EndPos:    1299,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 62,
										EndLine:   62,
										StartPos:  1297,
										EndPos:    1299,
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
								StartLine: 63,
								EndLine:   -1,
								StartPos:  1304,
								EndPos:    -1,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 63,
									EndLine:   63,
									StartPos:  1312,
									EndPos:    1314,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 63,
										EndLine:   63,
										StartPos:  1312,
										EndPos:    1314,
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
							StartLine: 64,
							EndLine:   -1,
							StartPos:  1319,
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
						StartLine: 67,
						EndLine:   67,
						StartPos:  1337,
						EndPos:    1357,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 67,
							EndLine:   67,
							StartPos:  1344,
							EndPos:    1345,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 67,
							EndLine:   67,
							StartPos:  1347,
							EndPos:    1357,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
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
			},
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 68,
						EndLine:   68,
						StartPos:  1360,
						EndPos:    1382,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 68,
							EndLine:   68,
							StartPos:  1367,
							EndPos:    1368,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 68,
							EndLine:   68,
							StartPos:  1370,
							EndPos:    1382,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 68,
									EndLine:   68,
									StartPos:  1372,
									EndPos:    1380,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 68,
										EndLine:   68,
										StartPos:  1378,
										EndPos:    1379,
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
						StartLine: 69,
						EndLine:   69,
						StartPos:  1385,
						EndPos:    1416,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 69,
							EndLine:   69,
							StartPos:  1392,
							EndPos:    1393,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 69,
							EndLine:   69,
							StartPos:  1397,
							EndPos:    1406,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtBreak{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 69,
									EndLine:   69,
									StartPos:  1397,
									EndPos:    1406,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 69,
										EndLine:   69,
										StartPos:  1403,
										EndPos:    1404,
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
						StartLine: 70,
						EndLine:   70,
						StartPos:  1419,
						EndPos:    1462,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 70,
							EndLine:   70,
							StartPos:  1425,
							EndPos:    1428,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassConstList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 70,
								EndLine:   70,
								StartPos:  1430,
								EndPos:    1460,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1430,
										EndPos:    1436,
									},
								},
								Value: []byte("public"),
							},
						},
						Consts: []ast.Vertex{
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1443,
										EndPos:    1450,
									},
								},
								ConstantName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 70,
											EndLine:   70,
											StartPos:  1443,
											EndPos:    1446,
										},
									},
									Value: []byte("FOO"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 70,
											EndLine:   70,
											StartPos:  1449,
											EndPos:    1450,
										},
									},
									Value: []byte("1"),
								},
							},
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 70,
										EndLine:   70,
										StartPos:  1452,
										EndPos:    1459,
									},
								},
								ConstantName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 70,
											EndLine:   70,
											StartPos:  1452,
											EndPos:    1455,
										},
									},
									Value: []byte("BAR"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 70,
											EndLine:   70,
											StartPos:  1458,
											EndPos:    1459,
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
						StartLine: 71,
						EndLine:   71,
						StartPos:  1465,
						EndPos:    1501,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 71,
							EndLine:   71,
							StartPos:  1471,
							EndPos:    1474,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassConstList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 71,
								EndLine:   71,
								StartPos:  1476,
								EndPos:    1499,
							},
						},
						Consts: []ast.Vertex{
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 71,
										EndLine:   71,
										StartPos:  1482,
										EndPos:    1489,
									},
								},
								ConstantName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 71,
											EndLine:   71,
											StartPos:  1482,
											EndPos:    1485,
										},
									},
									Value: []byte("FOO"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 71,
											EndLine:   71,
											StartPos:  1488,
											EndPos:    1489,
										},
									},
									Value: []byte("1"),
								},
							},
							&ast.StmtConstant{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 71,
										EndLine:   71,
										StartPos:  1491,
										EndPos:    1498,
									},
								},
								ConstantName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 71,
											EndLine:   71,
											StartPos:  1491,
											EndPos:    1494,
										},
									},
									Value: []byte("BAR"),
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 71,
											EndLine:   71,
											StartPos:  1497,
											EndPos:    1498,
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
						StartLine: 72,
						EndLine:   72,
						StartPos:  1504,
						EndPos:    1534,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 72,
							EndLine:   72,
							StartPos:  1510,
							EndPos:    1513,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 72,
								EndLine:   72,
								StartPos:  1515,
								EndPos:    1532,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 72,
									EndLine:   72,
									StartPos:  1524,
									EndPos:    1527,
								},
							},
							Value: []byte("bar"),
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 72,
									EndLine:   72,
									StartPos:  1530,
									EndPos:    1532,
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
						StartLine: 73,
						EndLine:   73,
						StartPos:  1537,
						EndPos:    1582,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 73,
							EndLine:   73,
							StartPos:  1543,
							EndPos:    1546,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 73,
								EndLine:   73,
								StartPos:  1548,
								EndPos:    1580,
							},
						},
						ReturnsRef: true,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 73,
									EndLine:   73,
									StartPos:  1572,
									EndPos:    1575,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 73,
										EndLine:   73,
										StartPos:  1548,
										EndPos:    1554,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 73,
										EndLine:   73,
										StartPos:  1555,
										EndPos:    1561,
									},
								},
								Value: []byte("static"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 73,
									EndLine:   73,
									StartPos:  1578,
									EndPos:    1580,
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
						StartLine: 74,
						EndLine:   74,
						StartPos:  1585,
						EndPos:    1636,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 74,
							EndLine:   74,
							StartPos:  1591,
							EndPos:    1594,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 74,
								EndLine:   74,
								StartPos:  1596,
								EndPos:    1634,
							},
						},
						ReturnsRef: true,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 74,
									EndLine:   74,
									StartPos:  1620,
									EndPos:    1623,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 74,
										EndLine:   74,
										StartPos:  1596,
										EndPos:    1602,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 74,
										EndLine:   74,
										StartPos:  1603,
										EndPos:    1609,
									},
								},
								Value: []byte("static"),
							},
						},
						ReturnType: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 74,
									EndLine:   74,
									StartPos:  1627,
									EndPos:    1631,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 74,
											EndLine:   74,
											StartPos:  1627,
											EndPos:    1631,
										},
									},
									Value: []byte("void"),
								},
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 74,
									EndLine:   74,
									StartPos:  1632,
									EndPos:    1634,
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
						StartLine: 75,
						EndLine:   75,
						StartPos:  1639,
						EndPos:    1660,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 75,
							EndLine:   75,
							StartPos:  1654,
							EndPos:    1657,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 75,
								EndLine:   75,
								StartPos:  1639,
								EndPos:    1647,
							},
						},
						Value: []byte("abstract"),
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 76,
						EndLine:   76,
						StartPos:  1663,
						EndPos:    1694,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 76,
							EndLine:   76,
							StartPos:  1675,
							EndPos:    1678,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 76,
								EndLine:   76,
								StartPos:  1663,
								EndPos:    1668,
							},
						},
						Value: []byte("final"),
					},
				},
				Extends: &ast.StmtClassExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 76,
							EndLine:   76,
							StartPos:  1679,
							EndPos:    1690,
						},
					},
					ClassName: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 76,
								EndLine:   76,
								StartPos:  1687,
								EndPos:    1690,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 76,
										EndLine:   76,
										StartPos:  1687,
										EndPos:    1690,
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
						StartLine: 77,
						EndLine:   77,
						StartPos:  1697,
						EndPos:    1731,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 77,
							EndLine:   77,
							StartPos:  1709,
							EndPos:    1712,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 77,
								EndLine:   77,
								StartPos:  1697,
								EndPos:    1702,
							},
						},
						Value: []byte("final"),
					},
				},
				Implements: &ast.StmtClassImplements{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 77,
							EndLine:   77,
							StartPos:  1713,
							EndPos:    1727,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 77,
									EndLine:   77,
									StartPos:  1724,
									EndPos:    1727,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 77,
											EndLine:   77,
											StartPos:  1724,
											EndPos:    1727,
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
						StartLine: 78,
						EndLine:   78,
						StartPos:  1734,
						EndPos:    1773,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 78,
							EndLine:   78,
							StartPos:  1746,
							EndPos:    1749,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 78,
								EndLine:   78,
								StartPos:  1734,
								EndPos:    1739,
							},
						},
						Value: []byte("final"),
					},
				},
				Implements: &ast.StmtClassImplements{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 78,
							EndLine:   78,
							StartPos:  1750,
							EndPos:    1769,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 78,
									EndLine:   78,
									StartPos:  1761,
									EndPos:    1764,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 78,
											EndLine:   78,
											StartPos:  1761,
											EndPos:    1764,
										},
									},
									Value: []byte("bar"),
								},
							},
						},
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 78,
									EndLine:   78,
									StartPos:  1766,
									EndPos:    1769,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 78,
											EndLine:   78,
											StartPos:  1766,
											EndPos:    1769,
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
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 79,
						EndLine:   79,
						StartPos:  1776,
						EndPos:    1824,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 79,
							EndLine:   79,
							StartPos:  1776,
							EndPos:    1823,
						},
					},
					Class: &ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 79,
								EndLine:   79,
								StartPos:  1780,
								EndPos:    1823,
							},
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 79,
									EndLine:   79,
									StartPos:  1785,
									EndPos:    1787,
								},
							},
						},
						Extends: &ast.StmtClassExtends{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 79,
									EndLine:   79,
									StartPos:  1788,
									EndPos:    1799,
								},
							},
							ClassName: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 79,
										EndLine:   79,
										StartPos:  1796,
										EndPos:    1799,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 79,
												EndLine:   79,
												StartPos:  1796,
												EndPos:    1799,
											},
										},
										Value: []byte("foo"),
									},
								},
							},
						},
						Implements: &ast.StmtClassImplements{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 79,
									EndLine:   79,
									StartPos:  1800,
									EndPos:    1819,
								},
							},
							InterfaceNames: []ast.Vertex{
								&ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 79,
											EndLine:   79,
											StartPos:  1811,
											EndPos:    1814,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 79,
													EndLine:   79,
													StartPos:  1811,
													EndPos:    1814,
												},
											},
											Value: []byte("bar"),
										},
									},
								},
								&ast.NameName{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 79,
											EndLine:   79,
											StartPos:  1816,
											EndPos:    1819,
										},
									},
									Parts: []ast.Vertex{
										&ast.NameNamePart{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 79,
													EndLine:   79,
													StartPos:  1816,
													EndPos:    1819,
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
				},
			},
			&ast.StmtConstList{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 81,
						EndLine:   81,
						StartPos:  1828,
						EndPos:    1851,
					},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1834,
								EndPos:    1841,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 81,
									EndLine:   81,
									StartPos:  1834,
									EndPos:    1837,
								},
							},
							Value: []byte("FOO"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 81,
									EndLine:   81,
									StartPos:  1840,
									EndPos:    1841,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 81,
								EndLine:   81,
								StartPos:  1843,
								EndPos:    1850,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 81,
									EndLine:   81,
									StartPos:  1843,
									EndPos:    1846,
								},
							},
							Value: []byte("BAR"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 81,
									EndLine:   81,
									StartPos:  1849,
									EndPos:    1850,
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
						StartLine: 82,
						EndLine:   82,
						StartPos:  1854,
						EndPos:    1877,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1861,
							EndPos:    1862,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 82,
							EndLine:   82,
							StartPos:  1864,
							EndPos:    1877,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
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
			},
			&ast.StmtWhile{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 83,
						EndLine:   83,
						StartPos:  1880,
						EndPos:    1905,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 83,
							EndLine:   83,
							StartPos:  1887,
							EndPos:    1888,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 83,
							EndLine:   83,
							StartPos:  1890,
							EndPos:    1905,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 83,
									EndLine:   83,
									StartPos:  1892,
									EndPos:    1903,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 83,
										EndLine:   83,
										StartPos:  1901,
										EndPos:    1902,
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
						StartLine: 84,
						EndLine:   84,
						StartPos:  1908,
						EndPos:    1934,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1915,
							EndPos:    1916,
						},
					},
					Value: []byte("1"),
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 84,
							EndLine:   84,
							StartPos:  1918,
							EndPos:    1934,
						},
					},
					Stmts: []ast.Vertex{
						&ast.StmtContinue{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 84,
									EndLine:   84,
									StartPos:  1920,
									EndPos:    1932,
								},
							},
							Expr: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 84,
										EndLine:   84,
										StartPos:  1929,
										EndPos:    1930,
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
						StartLine: 85,
						EndLine:   85,
						StartPos:  1937,
						EndPos:    1954,
					},
				},
				Alt: false,
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 85,
								EndLine:   85,
								StartPos:  1945,
								EndPos:    1952,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 85,
									EndLine:   85,
									StartPos:  1945,
									EndPos:    1950,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 85,
									EndLine:   85,
									StartPos:  1951,
									EndPos:    1952,
								},
							},
							Value: []byte("1"),
						},
					},
				},
				Stmt: &ast.StmtNop{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 85,
							EndLine:   85,
							StartPos:  1953,
							EndPos:    1954,
						},
					},
				},
			},
			&ast.StmtDeclare{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 86,
						EndLine:   86,
						StartPos:  1957,
						EndPos:    1976,
					},
				},
				Alt: false,
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 86,
								EndLine:   86,
								StartPos:  1965,
								EndPos:    1972,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 86,
									EndLine:   86,
									StartPos:  1965,
									EndPos:    1970,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 86,
									EndLine:   86,
									StartPos:  1971,
									EndPos:    1972,
								},
							},
							Value: []byte("1"),
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 86,
							EndLine:   86,
							StartPos:  1974,
							EndPos:    1976,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtDeclare{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 87,
						EndLine:   87,
						StartPos:  1979,
						EndPos:    2008,
					},
				},
				Alt: true,
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 87,
								EndLine:   87,
								StartPos:  1987,
								EndPos:    1994,
							},
						},
						ConstantName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 87,
									EndLine:   87,
									StartPos:  1987,
									EndPos:    1992,
								},
							},
							Value: []byte("ticks"),
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 87,
									EndLine:   87,
									StartPos:  1993,
									EndPos:    1994,
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
						StartLine: 88,
						EndLine:   88,
						StartPos:  2011,
						EndPos:    2026,
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 88,
							EndLine:   88,
							StartPos:  2014,
							EndPos:    2016,
						},
					},
					Stmts: []ast.Vertex{},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 88,
							EndLine:   88,
							StartPos:  2023,
							EndPos:    2024,
						},
					},
					Value: []byte("1"),
				},
			},
			&ast.StmtEcho{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 89,
						EndLine:   89,
						StartPos:  2029,
						EndPos:    2040,
					},
				},
				Exprs: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 89,
								EndLine:   89,
								StartPos:  2034,
								EndPos:    2036,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 89,
									EndLine:   89,
									StartPos:  2034,
									EndPos:    2036,
								},
							},
							Value: []byte("a"),
						},
					},
					&ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 89,
								EndLine:   89,
								StartPos:  2038,
								EndPos:    2039,
							},
						},
						Value: []byte("1"),
					},
				},
			},
			&ast.StmtEcho{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 90,
						EndLine:   90,
						StartPos:  2043,
						EndPos:    2052,
					},
				},
				Exprs: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 90,
								EndLine:   90,
								StartPos:  2048,
								EndPos:    2050,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 90,
									EndLine:   90,
									StartPos:  2048,
									EndPos:    2050,
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
						StartLine: 91,
						EndLine:   91,
						StartPos:  2055,
						EndPos:    2090,
					},
				},
				Init: []ast.Vertex{
					&ast.ExprAssign{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2059,
								EndPos:    2065,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2059,
									EndPos:    2061,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 91,
										EndLine:   91,
										StartPos:  2059,
										EndPos:    2061,
									},
								},
								Value: []byte("i"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2064,
									EndPos:    2065,
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
								StartLine: 91,
								EndLine:   91,
								StartPos:  2067,
								EndPos:    2074,
							},
						},
						Left: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2067,
									EndPos:    2069,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 91,
										EndLine:   91,
										StartPos:  2067,
										EndPos:    2069,
									},
								},
								Value: []byte("i"),
							},
						},
						Right: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2072,
									EndPos:    2074,
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
								StartLine: 91,
								EndLine:   91,
								StartPos:  2076,
								EndPos:    2080,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2076,
									EndPos:    2078,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 91,
										EndLine:   91,
										StartPos:  2076,
										EndPos:    2078,
									},
								},
								Value: []byte("i"),
							},
						},
					},
					&ast.ExprPostInc{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 91,
								EndLine:   91,
								StartPos:  2082,
								EndPos:    2086,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 91,
									EndLine:   91,
									StartPos:  2082,
									EndPos:    2084,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 91,
										EndLine:   91,
										StartPos:  2082,
										EndPos:    2084,
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
							StartLine: 91,
							EndLine:   91,
							StartPos:  2088,
							EndPos:    2090,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtAltFor{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 92,
						EndLine:   92,
						StartPos:  2093,
						EndPos:    2129,
					},
				},
				Cond: []ast.Vertex{
					&ast.ExprBinarySmaller{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 92,
								EndLine:   92,
								StartPos:  2099,
								EndPos:    2106,
							},
						},
						Left: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  2099,
									EndPos:    2101,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 92,
										EndLine:   92,
										StartPos:  2099,
										EndPos:    2101,
									},
								},
								Value: []byte("i"),
							},
						},
						Right: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  2104,
									EndPos:    2106,
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
								StartLine: 92,
								EndLine:   92,
								StartPos:  2108,
								EndPos:    2112,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  2108,
									EndPos:    2110,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 92,
										EndLine:   92,
										StartPos:  2108,
										EndPos:    2110,
									},
								},
								Value: []byte("i"),
							},
						},
					},
					&ast.ExprPostInc{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 92,
								EndLine:   92,
								StartPos:  2114,
								EndPos:    2118,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 92,
									EndLine:   92,
									StartPos:  2114,
									EndPos:    2116,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 92,
										EndLine:   92,
										StartPos:  2114,
										EndPos:    2116,
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
						StartLine: 93,
						EndLine:   93,
						StartPos:  2132,
						EndPos:    2153,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 93,
							EndLine:   93,
							StartPos:  2141,
							EndPos:    2143,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 93,
								EndLine:   93,
								StartPos:  2141,
								EndPos:    2143,
							},
						},
						Value: []byte("a"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 93,
							EndLine:   93,
							StartPos:  2147,
							EndPos:    2149,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 93,
								EndLine:   93,
								StartPos:  2147,
								EndPos:    2149,
							},
						},
						Value: []byte("v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 93,
							EndLine:   93,
							StartPos:  2151,
							EndPos:    2153,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtAltForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 94,
						EndLine:   94,
						StartPos:  2156,
						EndPos:    2188,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 94,
							EndLine:   94,
							StartPos:  2165,
							EndPos:    2167,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 94,
								EndLine:   94,
								StartPos:  2165,
								EndPos:    2167,
							},
						},
						Value: []byte("a"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 94,
							EndLine:   94,
							StartPos:  2171,
							EndPos:    2173,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 94,
								EndLine:   94,
								StartPos:  2171,
								EndPos:    2173,
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
						StartLine: 95,
						EndLine:   95,
						StartPos:  2191,
						EndPos:    2218,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2200,
							EndPos:    2202,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 95,
								EndLine:   95,
								StartPos:  2200,
								EndPos:    2202,
							},
						},
						Value: []byte("a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2206,
							EndPos:    2208,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 95,
								EndLine:   95,
								StartPos:  2206,
								EndPos:    2208,
							},
						},
						Value: []byte("k"),
					},
				},
				Var: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2212,
							EndPos:    2214,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 95,
								EndLine:   95,
								StartPos:  2212,
								EndPos:    2214,
							},
						},
						Value: []byte("v"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 95,
							EndLine:   95,
							StartPos:  2216,
							EndPos:    2218,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 96,
						EndLine:   96,
						StartPos:  2221,
						EndPos:    2249,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2230,
							EndPos:    2232,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 96,
								EndLine:   96,
								StartPos:  2230,
								EndPos:    2232,
							},
						},
						Value: []byte("a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2236,
							EndPos:    2238,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 96,
								EndLine:   96,
								StartPos:  2236,
								EndPos:    2238,
							},
						},
						Value: []byte("k"),
					},
				},
				Var: &ast.ExprReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2242,
							EndPos:    2245,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 96,
								EndLine:   96,
								StartPos:  2243,
								EndPos:    2245,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 96,
									EndLine:   96,
									StartPos:  2243,
									EndPos:    2245,
								},
							},
							Value: []byte("v"),
						},
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 96,
							EndLine:   96,
							StartPos:  2247,
							EndPos:    2249,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 97,
						EndLine:   97,
						StartPos:  2252,
						EndPos:    2285,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 97,
							EndLine:   97,
							StartPos:  2261,
							EndPos:    2263,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2261,
								EndPos:    2263,
							},
						},
						Value: []byte("a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 97,
							EndLine:   97,
							StartPos:  2267,
							EndPos:    2269,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 97,
								EndLine:   97,
								StartPos:  2267,
								EndPos:    2269,
							},
						},
						Value: []byte("k"),
					},
				},
				Var: &ast.ExprList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 97,
							EndLine:   97,
							StartPos:  2273,
							EndPos:    2281,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 97,
									EndLine:   97,
									StartPos:  2278,
									EndPos:    2280,
								},
							},
							Val: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 97,
										EndLine:   97,
										StartPos:  2278,
										EndPos:    2280,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 97,
											EndLine:   97,
											StartPos:  2278,
											EndPos:    2280,
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
							StartLine: 97,
							EndLine:   97,
							StartPos:  2283,
							EndPos:    2285,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtForeach{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 98,
						EndLine:   98,
						StartPos:  2288,
						EndPos:    2317,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2297,
							EndPos:    2299,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 98,
								EndLine:   98,
								StartPos:  2297,
								EndPos:    2299,
							},
						},
						Value: []byte("a"),
					},
				},
				Key: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2303,
							EndPos:    2305,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 98,
								EndLine:   98,
								StartPos:  2303,
								EndPos:    2305,
							},
						},
						Value: []byte("k"),
					},
				},
				Var: &ast.ExprShortList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 98,
							EndLine:   98,
							StartPos:  2309,
							EndPos:    2313,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 98,
									EndLine:   98,
									StartPos:  2310,
									EndPos:    2312,
								},
							},
							Val: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 98,
										EndLine:   98,
										StartPos:  2310,
										EndPos:    2312,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 98,
											EndLine:   98,
											StartPos:  2310,
											EndPos:    2312,
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
							StartLine: 98,
							EndLine:   98,
							StartPos:  2315,
							EndPos:    2317,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 99,
						EndLine:   99,
						StartPos:  2320,
						EndPos:    2337,
					},
				},
				ReturnsRef: false,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 99,
							EndLine:   99,
							StartPos:  2329,
							EndPos:    2332,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 100,
						EndLine:   100,
						StartPos:  2340,
						EndPos:    2364,
					},
				},
				ReturnsRef: false,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 100,
							EndLine:   100,
							StartPos:  2349,
							EndPos:    2352,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtReturn{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 100,
								EndLine:   100,
								StartPos:  2356,
								EndPos:    2363,
							},
						},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 101,
						EndLine:   101,
						StartPos:  2367,
						EndPos:    2394,
					},
				},
				ReturnsRef: true,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 101,
							EndLine:   101,
							StartPos:  2377,
							EndPos:    2380,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtReturn{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 101,
								EndLine:   101,
								StartPos:  2384,
								EndPos:    2393,
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 101,
									EndLine:   101,
									StartPos:  2391,
									EndPos:    2392,
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
						StartLine: 102,
						EndLine:   102,
						StartPos:  2397,
						EndPos:    2421,
					},
				},
				ReturnsRef: true,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 102,
							EndLine:   102,
							StartPos:  2407,
							EndPos:    2410,
						},
					},
					Value: []byte("foo"),
				},
				ReturnType: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 102,
							EndLine:   102,
							StartPos:  2414,
							EndPos:    2418,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 102,
									EndLine:   102,
									StartPos:  2414,
									EndPos:    2418,
								},
							},
							Value: []byte("void"),
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtGlobal{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 103,
						EndLine:   103,
						StartPos:  2424,
						EndPos:    2438,
					},
				},
				Vars: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2431,
								EndPos:    2433,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2431,
									EndPos:    2433,
								},
							},
							Value: []byte("a"),
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 103,
								EndLine:   103,
								StartPos:  2435,
								EndPos:    2437,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 103,
									EndLine:   103,
									StartPos:  2435,
									EndPos:    2437,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtLabel{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 104,
						EndLine:   104,
						StartPos:  2441,
						EndPos:    2443,
					},
				},
				LabelName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 104,
							EndLine:   104,
							StartPos:  2441,
							EndPos:    2442,
						},
					},
					Value: []byte("a"),
				},
			},
			&ast.StmtGoto{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 105,
						EndLine:   105,
						StartPos:  2447,
						EndPos:    2454,
					},
				},
				Label: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 105,
							EndLine:   105,
							StartPos:  2452,
							EndPos:    2453,
						},
					},
					Value: []byte("a"),
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 106,
						EndLine:   106,
						StartPos:  2457,
						EndPos:    2467,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 106,
							EndLine:   106,
							StartPos:  2461,
							EndPos:    2463,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 106,
								EndLine:   106,
								StartPos:  2461,
								EndPos:    2463,
							},
						},
						Value: []byte("a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 106,
							EndLine:   106,
							StartPos:  2465,
							EndPos:    2467,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 107,
						EndLine:   107,
						StartPos:  2470,
						EndPos:    2495,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 107,
							EndLine:   107,
							StartPos:  2474,
							EndPos:    2476,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 107,
								EndLine:   107,
								StartPos:  2474,
								EndPos:    2476,
							},
						},
						Value: []byte("a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 107,
							EndLine:   107,
							StartPos:  2478,
							EndPos:    2480,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 107,
								EndLine:   107,
								StartPos:  2481,
								EndPos:    2495,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 107,
									EndLine:   107,
									StartPos:  2489,
									EndPos:    2491,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 107,
										EndLine:   107,
										StartPos:  2489,
										EndPos:    2491,
									},
								},
								Value: []byte("b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 107,
									EndLine:   107,
									StartPos:  2493,
									EndPos:    2495,
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
						StartLine: 108,
						EndLine:   108,
						StartPos:  2498,
						EndPos:    2516,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 108,
							EndLine:   108,
							StartPos:  2502,
							EndPos:    2504,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 108,
								EndLine:   108,
								StartPos:  2502,
								EndPos:    2504,
							},
						},
						Value: []byte("a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 108,
							EndLine:   108,
							StartPos:  2506,
							EndPos:    2508,
						},
					},
					Stmts: []ast.Vertex{},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 108,
							EndLine:   108,
							StartPos:  2509,
							EndPos:    2516,
						},
					},
					Stmt: &ast.StmtStmtList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 108,
								EndLine:   108,
								StartPos:  2514,
								EndPos:    2516,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 109,
						EndLine:   109,
						StartPos:  2519,
						EndPos:    2567,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 109,
							EndLine:   109,
							StartPos:  2523,
							EndPos:    2525,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 109,
								EndLine:   109,
								StartPos:  2523,
								EndPos:    2525,
							},
						},
						Value: []byte("a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 109,
							EndLine:   109,
							StartPos:  2527,
							EndPos:    2529,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 109,
								EndLine:   109,
								StartPos:  2530,
								EndPos:    2544,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 109,
									EndLine:   109,
									StartPos:  2538,
									EndPos:    2540,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 109,
										EndLine:   109,
										StartPos:  2538,
										EndPos:    2540,
									},
								},
								Value: []byte("b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 109,
									EndLine:   109,
									StartPos:  2542,
									EndPos:    2544,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 109,
								EndLine:   109,
								StartPos:  2545,
								EndPos:    2559,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 109,
									EndLine:   109,
									StartPos:  2553,
									EndPos:    2555,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 109,
										EndLine:   109,
										StartPos:  2553,
										EndPos:    2555,
									},
								},
								Value: []byte("c"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 109,
									EndLine:   109,
									StartPos:  2557,
									EndPos:    2559,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 109,
							EndLine:   109,
							StartPos:  2560,
							EndPos:    2567,
						},
					},
					Stmt: &ast.StmtStmtList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 109,
								EndLine:   109,
								StartPos:  2565,
								EndPos:    2567,
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtIf{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 110,
						EndLine:   110,
						StartPos:  2570,
						EndPos:    2619,
					},
				},
				Cond: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2574,
							EndPos:    2576,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2574,
								EndPos:    2576,
							},
						},
						Value: []byte("a"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2578,
							EndPos:    2580,
						},
					},
					Stmts: []ast.Vertex{},
				},
				ElseIf: []ast.Vertex{
					&ast.StmtElseIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2581,
								EndPos:    2595,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2589,
									EndPos:    2591,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 110,
										EndLine:   110,
										StartPos:  2589,
										EndPos:    2591,
									},
								},
								Value: []byte("b"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2593,
									EndPos:    2595,
								},
							},
							Stmts: []ast.Vertex{},
						},
					},
				},
				Else: &ast.StmtElse{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 110,
							EndLine:   110,
							StartPos:  2596,
							EndPos:    2619,
						},
					},
					Stmt: &ast.StmtIf{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 110,
								EndLine:   110,
								StartPos:  2601,
								EndPos:    2619,
							},
						},
						Cond: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2605,
									EndPos:    2607,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 110,
										EndLine:   110,
										StartPos:  2605,
										EndPos:    2607,
									},
								},
								Value: []byte("c"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2609,
									EndPos:    2611,
								},
							},
							Stmts: []ast.Vertex{},
						},
						Else: &ast.StmtElse{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 110,
									EndLine:   110,
									StartPos:  2612,
									EndPos:    2619,
								},
							},
							Stmt: &ast.StmtStmtList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 110,
										EndLine:   110,
										StartPos:  2617,
										EndPos:    2619,
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
						StartLine: 111,
						EndLine:   111,
						StartPos:  2622,
						EndPos:    2624,
					},
				},
			},
			&ast.StmtInlineHtml{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 111,
						EndLine:   111,
						StartPos:  2624,
						EndPos:    2637,
					},
				},
				Value: []byte(" <div></div> "),
			},
			&ast.StmtInterface{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 112,
						EndLine:   112,
						StartPos:  2642,
						EndPos:    2658,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 112,
							EndLine:   112,
							StartPos:  2652,
							EndPos:    2655,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtInterface{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 113,
						EndLine:   113,
						StartPos:  2661,
						EndPos:    2689,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 113,
							EndLine:   113,
							StartPos:  2671,
							EndPos:    2674,
						},
					},
					Value: []byte("Foo"),
				},
				Extends: &ast.StmtInterfaceExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 113,
							EndLine:   113,
							StartPos:  2675,
							EndPos:    2686,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 113,
									EndLine:   113,
									StartPos:  2683,
									EndPos:    2686,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 113,
											EndLine:   113,
											StartPos:  2683,
											EndPos:    2686,
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
						StartLine: 114,
						EndLine:   114,
						StartPos:  2692,
						EndPos:    2725,
					},
				},
				InterfaceName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 114,
							EndLine:   114,
							StartPos:  2702,
							EndPos:    2705,
						},
					},
					Value: []byte("Foo"),
				},
				Extends: &ast.StmtInterfaceExtends{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 114,
							EndLine:   114,
							StartPos:  2706,
							EndPos:    2722,
						},
					},
					InterfaceNames: []ast.Vertex{
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 114,
									EndLine:   114,
									StartPos:  2714,
									EndPos:    2717,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 114,
											EndLine:   114,
											StartPos:  2714,
											EndPos:    2717,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						&ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 114,
									EndLine:   114,
									StartPos:  2719,
									EndPos:    2722,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 114,
											EndLine:   114,
											StartPos:  2719,
											EndPos:    2722,
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
						StartLine: 115,
						EndLine:   115,
						StartPos:  2728,
						EndPos:    2742,
					},
				},
				NamespaceName: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 115,
							EndLine:   115,
							StartPos:  2738,
							EndPos:    2741,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 115,
									EndLine:   115,
									StartPos:  2738,
									EndPos:    2741,
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
						StartLine: 116,
						EndLine:   116,
						StartPos:  2745,
						EndPos:    2761,
					},
				},
				NamespaceName: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 116,
							EndLine:   116,
							StartPos:  2755,
							EndPos:    2758,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 116,
									EndLine:   116,
									StartPos:  2755,
									EndPos:    2758,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtNamespace{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 117,
						EndLine:   117,
						StartPos:  2764,
						EndPos:    2776,
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 118,
						EndLine:   118,
						StartPos:  2779,
						EndPos:    2798,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 118,
							EndLine:   118,
							StartPos:  2785,
							EndPos:    2788,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtPropertyList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 118,
								EndLine:   118,
								StartPos:  2790,
								EndPos:    2797,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 118,
										EndLine:   118,
										StartPos:  2790,
										EndPos:    2793,
									},
								},
								Value: []byte("var"),
							},
						},
						Properties: []ast.Vertex{
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 118,
										EndLine:   118,
										StartPos:  2794,
										EndPos:    2796,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 118,
											EndLine:   118,
											StartPos:  2794,
											EndPos:    2796,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 118,
												EndLine:   118,
												StartPos:  2794,
												EndPos:    2796,
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
						StartLine: 119,
						EndLine:   119,
						StartPos:  2801,
						EndPos:    2838,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 119,
							EndLine:   119,
							StartPos:  2807,
							EndPos:    2810,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtPropertyList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 119,
								EndLine:   119,
								StartPos:  2812,
								EndPos:    2837,
							},
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 119,
										EndLine:   119,
										StartPos:  2812,
										EndPos:    2818,
									},
								},
								Value: []byte("public"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 119,
										EndLine:   119,
										StartPos:  2819,
										EndPos:    2825,
									},
								},
								Value: []byte("static"),
							},
						},
						Properties: []ast.Vertex{
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 119,
										EndLine:   119,
										StartPos:  2826,
										EndPos:    2828,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 119,
											EndLine:   119,
											StartPos:  2826,
											EndPos:    2828,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 119,
												EndLine:   119,
												StartPos:  2826,
												EndPos:    2828,
											},
										},
										Value: []byte("a"),
									},
								},
							},
							&ast.StmtProperty{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 119,
										EndLine:   119,
										StartPos:  2830,
										EndPos:    2836,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 119,
											EndLine:   119,
											StartPos:  2830,
											EndPos:    2832,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 119,
												EndLine:   119,
												StartPos:  2830,
												EndPos:    2832,
											},
										},
										Value: []byte("b"),
									},
								},
								Expr: &ast.ScalarLnumber{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 119,
											EndLine:   119,
											StartPos:  2835,
											EndPos:    2836,
										},
									},
									Value: []byte("1"),
								},
							},
						},
					},
				},
			},
			&ast.StmtStatic{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 120,
						EndLine:   120,
						StartPos:  2841,
						EndPos:    2859,
					},
				},
				Vars: []ast.Vertex{
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 120,
								EndLine:   120,
								StartPos:  2848,
								EndPos:    2850,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 120,
									EndLine:   120,
									StartPos:  2848,
									EndPos:    2850,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 120,
										EndLine:   120,
										StartPos:  2848,
										EndPos:    2850,
									},
								},
								Value: []byte("a"),
							},
						},
					},
					&ast.StmtStaticVar{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 120,
								EndLine:   120,
								StartPos:  2852,
								EndPos:    2858,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 120,
									EndLine:   120,
									StartPos:  2852,
									EndPos:    2854,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 120,
										EndLine:   120,
										StartPos:  2852,
										EndPos:    2854,
									},
								},
								Value: []byte("b"),
							},
						},
						Expr: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 120,
									EndLine:   120,
									StartPos:  2857,
									EndPos:    2858,
								},
							},
							Value: []byte("1"),
						},
					},
				},
			},
			&ast.StmtAltSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 122,
						EndLine:   126,
						StartPos:  2863,
						EndPos:    2922,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 122,
							EndLine:   122,
							StartPos:  2871,
							EndPos:    2872,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 123,
							EndLine:   -1,
							StartPos:  2879,
							EndPos:    -1,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 123,
									EndLine:   -1,
									StartPos:  2879,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 123,
										EndLine:   123,
										StartPos:  2884,
										EndPos:    2885,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtDefault{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 124,
									EndLine:   -1,
									StartPos:  2890,
									EndPos:    -1,
								},
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 125,
									EndLine:   -1,
									StartPos:  2902,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 125,
										EndLine:   125,
										StartPos:  2907,
										EndPos:    2908,
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
						StartLine: 128,
						EndLine:   131,
						StartPos:  2926,
						EndPos:    2974,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 128,
							EndLine:   128,
							StartPos:  2934,
							EndPos:    2935,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 129,
							EndLine:   -1,
							StartPos:  2943,
							EndPos:    -1,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 129,
									EndLine:   -1,
									StartPos:  2943,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 129,
										EndLine:   129,
										StartPos:  2948,
										EndPos:    2949,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 130,
									EndLine:   -1,
									StartPos:  2954,
									EndPos:    -1,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 130,
										EndLine:   130,
										StartPos:  2959,
										EndPos:    2960,
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
						StartLine: 133,
						EndLine:   136,
						StartPos:  2980,
						EndPos:    3032,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 133,
							EndLine:   133,
							StartPos:  2988,
							EndPos:    2989,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 133,
							EndLine:   136,
							StartPos:  2991,
							EndPos:    3032,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 134,
									EndLine:   134,
									StartPos:  2996,
									EndPos:    3010,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 134,
										EndLine:   134,
										StartPos:  3001,
										EndPos:    3002,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 134,
											EndLine:   134,
											StartPos:  3004,
											EndPos:    3010,
										},
									},
								},
							},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 135,
									EndLine:   135,
									StartPos:  3014,
									EndPos:    3028,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 135,
										EndLine:   135,
										StartPos:  3019,
										EndPos:    3020,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
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
			},
			&ast.StmtSwitch{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 138,
						EndLine:   141,
						StartPos:  3038,
						EndPos:    3091,
					},
				},
				Cond: &ast.ScalarLnumber{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 138,
							EndLine:   138,
							StartPos:  3046,
							EndPos:    3047,
						},
					},
					Value: []byte("1"),
				},
				CaseList: &ast.StmtCaseList{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 138,
							EndLine:   141,
							StartPos:  3049,
							EndPos:    3091,
						},
					},
					Cases: []ast.Vertex{
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 139,
									EndLine:   139,
									StartPos:  3055,
									EndPos:    3069,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 139,
										EndLine:   139,
										StartPos:  3060,
										EndPos:    3061,
									},
								},
								Value: []byte("1"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 139,
											EndLine:   139,
											StartPos:  3063,
											EndPos:    3069,
										},
									},
								},
							},
						},
						&ast.StmtCase{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 140,
									EndLine:   140,
									StartPos:  3073,
									EndPos:    3087,
								},
							},
							Cond: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 140,
										EndLine:   140,
										StartPos:  3078,
										EndPos:    3079,
									},
								},
								Value: []byte("2"),
							},
							Stmts: []ast.Vertex{
								&ast.StmtBreak{
									Node: ast.Node{
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
			},
			&ast.StmtThrow{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 143,
						EndLine:   143,
						StartPos:  3095,
						EndPos:    3104,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 143,
							EndLine:   143,
							StartPos:  3101,
							EndPos:    3103,
						},
					},
					VarName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 143,
								EndLine:   143,
								StartPos:  3101,
								EndPos:    3103,
							},
						},
						Value: []byte("e"),
					},
				},
			},
			&ast.StmtTrait{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 145,
						EndLine:   145,
						StartPos:  3108,
						EndPos:    3120,
					},
				},
				TraitName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 145,
							EndLine:   145,
							StartPos:  3114,
							EndPos:    3117,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 146,
						EndLine:   146,
						StartPos:  3123,
						EndPos:    3145,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 146,
							EndLine:   146,
							StartPos:  3129,
							EndPos:    3132,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 146,
								EndLine:   146,
								StartPos:  3135,
								EndPos:    3143,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 146,
										EndLine:   146,
										StartPos:  3139,
										EndPos:    3142,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 146,
												EndLine:   146,
												StartPos:  3139,
												EndPos:    3142,
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
									StartLine: 146,
									EndLine:   146,
									StartPos:  3142,
									EndPos:    3143,
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 147,
						EndLine:   147,
						StartPos:  3148,
						EndPos:    3177,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 147,
							EndLine:   147,
							StartPos:  3154,
							EndPos:    3157,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 147,
								EndLine:   147,
								StartPos:  3160,
								EndPos:    3175,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 147,
										EndLine:   147,
										StartPos:  3164,
										EndPos:    3167,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 147,
												EndLine:   147,
												StartPos:  3164,
												EndPos:    3167,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 147,
										EndLine:   147,
										StartPos:  3169,
										EndPos:    3172,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 147,
												EndLine:   147,
												StartPos:  3169,
												EndPos:    3172,
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
									StartLine: 147,
									EndLine:   147,
									StartPos:  3173,
									EndPos:    3175,
								},
							},
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 148,
						EndLine:   148,
						StartPos:  3180,
						EndPos:    3226,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 148,
							EndLine:   148,
							StartPos:  3186,
							EndPos:    3189,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 148,
								EndLine:   148,
								StartPos:  3192,
								EndPos:    3224,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 148,
										EndLine:   148,
										StartPos:  3196,
										EndPos:    3199,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 148,
												EndLine:   148,
												StartPos:  3196,
												EndPos:    3199,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 148,
										EndLine:   148,
										StartPos:  3201,
										EndPos:    3204,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 148,
												EndLine:   148,
												StartPos:  3201,
												EndPos:    3204,
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
									StartLine: 148,
									EndLine:   148,
									StartPos:  3205,
									EndPos:    3224,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUseAlias{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 148,
											EndLine:   148,
											StartPos:  3207,
											EndPos:    3221,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 148,
												EndLine:   148,
												StartPos:  3207,
												EndPos:    3210,
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 148,
													EndLine:   148,
													StartPos:  3207,
													EndPos:    3210,
												},
											},
											Value: []byte("one"),
										},
									},
									Alias: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 148,
												EndLine:   148,
												StartPos:  3214,
												EndPos:    3221,
											},
										},
										Value: []byte("include"),
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
						StartLine: 149,
						EndLine:   149,
						StartPos:  3229,
						EndPos:    3274,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 149,
							EndLine:   149,
							StartPos:  3235,
							EndPos:    3238,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 149,
								EndLine:   149,
								StartPos:  3241,
								EndPos:    3272,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3245,
										EndPos:    3248,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3245,
												EndPos:    3248,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 149,
										EndLine:   149,
										StartPos:  3250,
										EndPos:    3253,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3250,
												EndPos:    3253,
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
									StartLine: 149,
									EndLine:   149,
									StartPos:  3254,
									EndPos:    3272,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUseAlias{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 149,
											EndLine:   149,
											StartPos:  3256,
											EndPos:    3269,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3256,
												EndPos:    3259,
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 149,
													EndLine:   149,
													StartPos:  3256,
													EndPos:    3259,
												},
											},
											Value: []byte("one"),
										},
									},
									Modifier: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 149,
												EndLine:   149,
												StartPos:  3263,
												EndPos:    3269,
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
						StartLine: 150,
						EndLine:   150,
						StartPos:  3277,
						EndPos:    3326,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 150,
							EndLine:   150,
							StartPos:  3283,
							EndPos:    3286,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 150,
								EndLine:   150,
								StartPos:  3289,
								EndPos:    3324,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 150,
										EndLine:   150,
										StartPos:  3293,
										EndPos:    3296,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 150,
												EndLine:   150,
												StartPos:  3293,
												EndPos:    3296,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 150,
										EndLine:   150,
										StartPos:  3298,
										EndPos:    3301,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 150,
												EndLine:   150,
												StartPos:  3298,
												EndPos:    3301,
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
									StartLine: 150,
									EndLine:   150,
									StartPos:  3302,
									EndPos:    3324,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUseAlias{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 150,
											EndLine:   150,
											StartPos:  3304,
											EndPos:    3321,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 150,
												EndLine:   150,
												StartPos:  3304,
												EndPos:    3307,
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 150,
													EndLine:   150,
													StartPos:  3304,
													EndPos:    3307,
												},
											},
											Value: []byte("one"),
										},
									},
									Modifier: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 150,
												EndLine:   150,
												StartPos:  3311,
												EndPos:    3317,
											},
										},
										Value: []byte("public"),
									},
									Alias: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 150,
												EndLine:   150,
												StartPos:  3318,
												EndPos:    3321,
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
						StartLine: 151,
						EndLine:   151,
						StartPos:  3329,
						EndPos:    3406,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 151,
							EndLine:   151,
							StartPos:  3335,
							EndPos:    3338,
						},
					},
					Value: []byte("Foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtTraitUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 151,
								EndLine:   151,
								StartPos:  3341,
								EndPos:    3404,
							},
						},
						Traits: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 151,
										EndLine:   151,
										StartPos:  3345,
										EndPos:    3348,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3345,
												EndPos:    3348,
											},
										},
										Value: []byte("Bar"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 151,
										EndLine:   151,
										StartPos:  3350,
										EndPos:    3353,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3350,
												EndPos:    3353,
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
									StartLine: 151,
									EndLine:   151,
									StartPos:  3354,
									EndPos:    3404,
								},
							},
							Adaptations: []ast.Vertex{
								&ast.StmtTraitUsePrecedence{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 151,
											EndLine:   151,
											StartPos:  3356,
											EndPos:    3384,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3356,
												EndPos:    3364,
											},
										},
										Trait: &ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 151,
													EndLine:   151,
													StartPos:  3356,
													EndPos:    3359,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 151,
															EndLine:   151,
															StartPos:  3356,
															EndPos:    3359,
														},
													},
													Value: []byte("Bar"),
												},
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 151,
													EndLine:   151,
													StartPos:  3361,
													EndPos:    3364,
												},
											},
											Value: []byte("one"),
										},
									},
									Insteadof: []ast.Vertex{
										&ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 151,
													EndLine:   151,
													StartPos:  3375,
													EndPos:    3378,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 151,
															EndLine:   151,
															StartPos:  3375,
															EndPos:    3378,
														},
													},
													Value: []byte("Baz"),
												},
											},
										},
										&ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 151,
													EndLine:   151,
													StartPos:  3380,
													EndPos:    3384,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 151,
															EndLine:   151,
															StartPos:  3380,
															EndPos:    3384,
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
											StartLine: 151,
											EndLine:   151,
											StartPos:  3386,
											EndPos:    3401,
										},
									},
									Ref: &ast.StmtTraitMethodRef{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3386,
												EndPos:    3394,
											},
										},
										Trait: &ast.NameName{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 151,
													EndLine:   151,
													StartPos:  3386,
													EndPos:    3389,
												},
											},
											Parts: []ast.Vertex{
												&ast.NameNamePart{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 151,
															EndLine:   151,
															StartPos:  3386,
															EndPos:    3389,
														},
													},
													Value: []byte("Baz"),
												},
											},
										},
										Method: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 151,
													EndLine:   151,
													StartPos:  3391,
													EndPos:    3394,
												},
											},
											Value: []byte("one"),
										},
									},
									Alias: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 151,
												EndLine:   151,
												StartPos:  3398,
												EndPos:    3401,
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
						StartLine: 153,
						EndLine:   -1,
						StartPos:  3410,
						EndPos:    -1,
					},
				},
				Stmts:   []ast.Vertex{},
				Catches: []ast.Vertex{},
			},
			&ast.StmtTry{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 154,
						EndLine:   154,
						StartPos:  3419,
						EndPos:    3449,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 154,
								EndLine:   154,
								StartPos:  3426,
								EndPos:    3449,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 154,
										EndLine:   154,
										StartPos:  3433,
										EndPos:    3442,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 154,
												EndLine:   154,
												StartPos:  3433,
												EndPos:    3442,
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
									StartLine: 154,
									EndLine:   154,
									StartPos:  3443,
									EndPos:    3445,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 154,
										EndLine:   154,
										StartPos:  3443,
										EndPos:    3445,
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
						StartLine: 155,
						EndLine:   155,
						StartPos:  3452,
						EndPos:    3499,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 155,
								EndLine:   155,
								StartPos:  3459,
								EndPos:    3499,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 155,
										EndLine:   155,
										StartPos:  3466,
										EndPos:    3475,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 155,
												EndLine:   155,
												StartPos:  3466,
												EndPos:    3475,
											},
										},
										Value: []byte("Exception"),
									},
								},
							},
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 155,
										EndLine:   155,
										StartPos:  3476,
										EndPos:    3492,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 155,
												EndLine:   155,
												StartPos:  3476,
												EndPos:    3492,
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
									StartLine: 155,
									EndLine:   155,
									StartPos:  3493,
									EndPos:    3495,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 155,
										EndLine:   155,
										StartPos:  3493,
										EndPos:    3495,
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
						StartLine: 156,
						EndLine:   156,
						StartPos:  3502,
						EndPos:    3563,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 156,
								EndLine:   156,
								StartPos:  3509,
								EndPos:    3532,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 156,
										EndLine:   156,
										StartPos:  3516,
										EndPos:    3525,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 156,
												EndLine:   156,
												StartPos:  3516,
												EndPos:    3525,
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
									StartLine: 156,
									EndLine:   156,
									StartPos:  3526,
									EndPos:    3528,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 156,
										EndLine:   156,
										StartPos:  3526,
										EndPos:    3528,
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
								StartLine: 156,
								EndLine:   156,
								StartPos:  3533,
								EndPos:    3563,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 156,
										EndLine:   156,
										StartPos:  3540,
										EndPos:    3556,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 156,
												EndLine:   156,
												StartPos:  3540,
												EndPos:    3556,
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
									StartLine: 156,
									EndLine:   156,
									StartPos:  3557,
									EndPos:    3559,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 156,
										EndLine:   156,
										StartPos:  3557,
										EndPos:    3559,
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
						StartLine: 157,
						EndLine:   157,
						StartPos:  3566,
						EndPos:    3607,
					},
				},
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 157,
								EndLine:   157,
								StartPos:  3573,
								EndPos:    3596,
							},
						},
						Types: []ast.Vertex{
							&ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 157,
										EndLine:   157,
										StartPos:  3580,
										EndPos:    3589,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 157,
												EndLine:   157,
												StartPos:  3580,
												EndPos:    3589,
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
									StartLine: 157,
									EndLine:   157,
									StartPos:  3590,
									EndPos:    3592,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 157,
										EndLine:   157,
										StartPos:  3590,
										EndPos:    3592,
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
							StartLine: 157,
							EndLine:   157,
							StartPos:  3597,
							EndPos:    3607,
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtUnset{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 159,
						EndLine:   159,
						StartPos:  3611,
						EndPos:    3626,
					},
				},
				Vars: []ast.Vertex{
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 159,
								EndLine:   159,
								StartPos:  3617,
								EndPos:    3619,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 159,
									EndLine:   159,
									StartPos:  3617,
									EndPos:    3619,
								},
							},
							Value: []byte("a"),
						},
					},
					&ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 159,
								EndLine:   159,
								StartPos:  3621,
								EndPos:    3623,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 159,
									EndLine:   159,
									StartPos:  3621,
									EndPos:    3623,
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
						StartLine: 161,
						EndLine:   161,
						StartPos:  3630,
						EndPos:    3638,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 161,
								EndLine:   161,
								StartPos:  3634,
								EndPos:    3637,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 161,
									EndLine:   161,
									StartPos:  3634,
									EndPos:    3637,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 161,
											EndLine:   161,
											StartPos:  3634,
											EndPos:    3637,
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
						StartLine: 162,
						EndLine:   162,
						StartPos:  3641,
						EndPos:    3650,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 162,
								EndLine:   162,
								StartPos:  3646,
								EndPos:    3649,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 162,
									EndLine:   162,
									StartPos:  3646,
									EndPos:    3649,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 162,
											EndLine:   162,
											StartPos:  3646,
											EndPos:    3649,
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
						StartLine: 163,
						EndLine:   163,
						StartPos:  3653,
						EndPos:    3669,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 163,
								EndLine:   163,
								StartPos:  3658,
								EndPos:    3668,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 163,
									EndLine:   163,
									StartPos:  3658,
									EndPos:    3661,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 163,
											EndLine:   163,
											StartPos:  3658,
											EndPos:    3661,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 163,
									EndLine:   163,
									StartPos:  3665,
									EndPos:    3668,
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
						StartLine: 164,
						EndLine:   164,
						StartPos:  3672,
						EndPos:    3685,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 164,
								EndLine:   164,
								StartPos:  3676,
								EndPos:    3679,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 164,
									EndLine:   164,
									StartPos:  3676,
									EndPos:    3679,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 164,
											EndLine:   164,
											StartPos:  3676,
											EndPos:    3679,
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
								StartLine: 164,
								EndLine:   164,
								StartPos:  3681,
								EndPos:    3684,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 164,
									EndLine:   164,
									StartPos:  3681,
									EndPos:    3684,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 164,
											EndLine:   164,
											StartPos:  3681,
											EndPos:    3684,
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
						StartLine: 165,
						EndLine:   165,
						StartPos:  3688,
						EndPos:    3708,
					},
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 165,
								EndLine:   165,
								StartPos:  3692,
								EndPos:    3695,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 165,
									EndLine:   165,
									StartPos:  3692,
									EndPos:    3695,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 165,
											EndLine:   165,
											StartPos:  3692,
											EndPos:    3695,
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
								StartLine: 165,
								EndLine:   165,
								StartPos:  3697,
								EndPos:    3707,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 165,
									EndLine:   165,
									StartPos:  3697,
									EndPos:    3700,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 165,
											EndLine:   165,
											StartPos:  3697,
											EndPos:    3700,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 165,
									EndLine:   165,
									StartPos:  3704,
									EndPos:    3707,
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
						StartLine: 166,
						EndLine:   166,
						StartPos:  3711,
						EndPos:    3734,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 166,
							EndLine:   166,
							StartPos:  3715,
							EndPos:    3723,
						},
					},
					Value: []byte("function"),
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 166,
								EndLine:   166,
								StartPos:  3724,
								EndPos:    3727,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 166,
									EndLine:   166,
									StartPos:  3724,
									EndPos:    3727,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 166,
											EndLine:   166,
											StartPos:  3724,
											EndPos:    3727,
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
								StartLine: 166,
								EndLine:   166,
								StartPos:  3730,
								EndPos:    3733,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 166,
									EndLine:   166,
									StartPos:  3730,
									EndPos:    3733,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 166,
											EndLine:   166,
											StartPos:  3730,
											EndPos:    3733,
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
						StartLine: 167,
						EndLine:   167,
						StartPos:  3737,
						EndPos:    3774,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 167,
							EndLine:   167,
							StartPos:  3741,
							EndPos:    3749,
						},
					},
					Value: []byte("function"),
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 167,
								EndLine:   167,
								StartPos:  3750,
								EndPos:    3760,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 167,
									EndLine:   167,
									StartPos:  3750,
									EndPos:    3753,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 167,
											EndLine:   167,
											StartPos:  3750,
											EndPos:    3753,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 167,
									EndLine:   167,
									StartPos:  3757,
									EndPos:    3760,
								},
							},
							Value: []byte("foo"),
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 167,
								EndLine:   167,
								StartPos:  3763,
								EndPos:    3773,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 167,
									EndLine:   167,
									StartPos:  3763,
									EndPos:    3766,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 167,
											EndLine:   167,
											StartPos:  3763,
											EndPos:    3766,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 167,
									EndLine:   167,
									StartPos:  3770,
									EndPos:    3773,
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
						StartLine: 168,
						EndLine:   168,
						StartPos:  3777,
						EndPos:    3797,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 168,
							EndLine:   168,
							StartPos:  3781,
							EndPos:    3786,
						},
					},
					Value: []byte("const"),
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 168,
								EndLine:   168,
								StartPos:  3787,
								EndPos:    3790,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 168,
									EndLine:   168,
									StartPos:  3787,
									EndPos:    3790,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 168,
											EndLine:   168,
											StartPos:  3787,
											EndPos:    3790,
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
								StartLine: 168,
								EndLine:   168,
								StartPos:  3793,
								EndPos:    3796,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 168,
									EndLine:   168,
									StartPos:  3793,
									EndPos:    3796,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 168,
											EndLine:   168,
											StartPos:  3793,
											EndPos:    3796,
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
						StartLine: 169,
						EndLine:   169,
						StartPos:  3800,
						EndPos:    3834,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 169,
							EndLine:   169,
							StartPos:  3804,
							EndPos:    3809,
						},
					},
					Value: []byte("const"),
				},
				Uses: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 169,
								EndLine:   169,
								StartPos:  3810,
								EndPos:    3820,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 169,
									EndLine:   169,
									StartPos:  3810,
									EndPos:    3813,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 169,
											EndLine:   169,
											StartPos:  3810,
											EndPos:    3813,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 169,
									EndLine:   169,
									StartPos:  3817,
									EndPos:    3820,
								},
							},
							Value: []byte("foo"),
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 169,
								EndLine:   169,
								StartPos:  3823,
								EndPos:    3833,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 169,
									EndLine:   169,
									StartPos:  3823,
									EndPos:    3826,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 169,
											EndLine:   169,
											StartPos:  3823,
											EndPos:    3826,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 169,
									EndLine:   169,
									StartPos:  3830,
									EndPos:    3833,
								},
							},
							Value: []byte("bar"),
						},
					},
				},
			},
			&ast.StmtGroupUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 171,
						EndLine:   171,
						StartPos:  3838,
						EndPos:    3858,
					},
				},
				Prefix: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 171,
							EndLine:   171,
							StartPos:  3843,
							EndPos:    3846,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 171,
									EndLine:   171,
									StartPos:  3843,
									EndPos:    3846,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
				UseList: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 171,
								EndLine:   171,
								StartPos:  3848,
								EndPos:    3851,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 171,
									EndLine:   171,
									StartPos:  3848,
									EndPos:    3851,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 171,
											EndLine:   171,
											StartPos:  3848,
											EndPos:    3851,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 171,
								EndLine:   171,
								StartPos:  3853,
								EndPos:    3856,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 171,
									EndLine:   171,
									StartPos:  3853,
									EndPos:    3856,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 171,
											EndLine:   171,
											StartPos:  3853,
											EndPos:    3856,
										},
									},
									Value: []byte("Baz"),
								},
							},
						},
					},
				},
			},
			&ast.StmtGroupUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 172,
						EndLine:   172,
						StartPos:  3861,
						EndPos:    3888,
					},
				},
				Prefix: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 172,
							EndLine:   172,
							StartPos:  3865,
							EndPos:    3868,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 172,
									EndLine:   172,
									StartPos:  3865,
									EndPos:    3868,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
				UseList: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3870,
								EndPos:    3873,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 172,
									EndLine:   172,
									StartPos:  3870,
									EndPos:    3873,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 172,
											EndLine:   172,
											StartPos:  3870,
											EndPos:    3873,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 172,
								EndLine:   172,
								StartPos:  3875,
								EndPos:    3886,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 172,
									EndLine:   172,
									StartPos:  3875,
									EndPos:    3878,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 172,
											EndLine:   172,
											StartPos:  3875,
											EndPos:    3878,
										},
									},
									Value: []byte("Baz"),
								},
							},
						},
						Alias: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 172,
									EndLine:   172,
									StartPos:  3882,
									EndPos:    3886,
								},
							},
							Value: []byte("quux"),
						},
					},
				},
			},
			&ast.StmtGroupUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 173,
						EndLine:   173,
						StartPos:  3891,
						EndPos:    3919,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 173,
							EndLine:   173,
							StartPos:  3895,
							EndPos:    3903,
						},
					},
					Value: []byte("function"),
				},
				Prefix: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 173,
							EndLine:   173,
							StartPos:  3904,
							EndPos:    3907,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 173,
									EndLine:   173,
									StartPos:  3904,
									EndPos:    3907,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
				UseList: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 173,
								EndLine:   173,
								StartPos:  3909,
								EndPos:    3912,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 173,
									EndLine:   173,
									StartPos:  3909,
									EndPos:    3912,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 173,
											EndLine:   173,
											StartPos:  3909,
											EndPos:    3912,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 173,
								EndLine:   173,
								StartPos:  3914,
								EndPos:    3917,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 173,
									EndLine:   173,
									StartPos:  3914,
									EndPos:    3917,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 173,
											EndLine:   173,
											StartPos:  3914,
											EndPos:    3917,
										},
									},
									Value: []byte("Baz"),
								},
							},
						},
					},
				},
			},
			&ast.StmtGroupUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 174,
						EndLine:   174,
						StartPos:  3922,
						EndPos:    3948,
					},
				},
				UseType: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 174,
							EndLine:   174,
							StartPos:  3926,
							EndPos:    3931,
						},
					},
					Value: []byte("const"),
				},
				Prefix: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 174,
							EndLine:   174,
							StartPos:  3933,
							EndPos:    3936,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 174,
									EndLine:   174,
									StartPos:  3933,
									EndPos:    3936,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
				UseList: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3938,
								EndPos:    3941,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 174,
									EndLine:   174,
									StartPos:  3938,
									EndPos:    3941,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 174,
											EndLine:   174,
											StartPos:  3938,
											EndPos:    3941,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 174,
								EndLine:   174,
								StartPos:  3943,
								EndPos:    3946,
							},
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 174,
									EndLine:   174,
									StartPos:  3943,
									EndPos:    3946,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 174,
											EndLine:   174,
											StartPos:  3943,
											EndPos:    3946,
										},
									},
									Value: []byte("Baz"),
								},
							},
						},
					},
				},
			},
			&ast.StmtGroupUse{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 175,
						EndLine:   175,
						StartPos:  3951,
						EndPos:    3985,
					},
				},
				Prefix: &ast.NameName{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 175,
							EndLine:   175,
							StartPos:  3955,
							EndPos:    3958,
						},
					},
					Parts: []ast.Vertex{
						&ast.NameNamePart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 175,
									EndLine:   175,
									StartPos:  3955,
									EndPos:    3958,
								},
							},
							Value: []byte("Foo"),
						},
					},
				},
				UseList: []ast.Vertex{
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3966,
								EndPos:    3969,
							},
						},
						UseType: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 175,
									EndLine:   175,
									StartPos:  3960,
									EndPos:    3965,
								},
							},
							Value: []byte("const"),
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 175,
									EndLine:   175,
									StartPos:  3966,
									EndPos:    3969,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 175,
											EndLine:   175,
											StartPos:  3966,
											EndPos:    3969,
										},
									},
									Value: []byte("Bar"),
								},
							},
						},
					},
					&ast.StmtUse{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 175,
								EndLine:   175,
								StartPos:  3980,
								EndPos:    3983,
							},
						},
						UseType: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 175,
									EndLine:   175,
									StartPos:  3971,
									EndPos:    3979,
								},
							},
							Value: []byte("function"),
						},
						Use: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 175,
									EndLine:   175,
									StartPos:  3980,
									EndPos:    3983,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 175,
											EndLine:   175,
											StartPos:  3980,
											EndPos:    3983,
										},
									},
									Value: []byte("Baz"),
								},
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 177,
						EndLine:   177,
						StartPos:  3989,
						EndPos:    3995,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 177,
							EndLine:   177,
							StartPos:  3989,
							EndPos:    3994,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 177,
								EndLine:   177,
								StartPos:  3989,
								EndPos:    3991,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 177,
									EndLine:   177,
									StartPos:  3989,
									EndPos:    3991,
								},
							},
							Value: []byte("a"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 177,
								EndLine:   177,
								StartPos:  3992,
								EndPos:    3993,
							},
						},
						Value: []byte("1"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 178,
						EndLine:   178,
						StartPos:  3998,
						EndPos:    4007,
					},
				},
				Expr: &ast.ExprArrayDimFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 178,
							EndLine:   178,
							StartPos:  3998,
							EndPos:    4006,
						},
					},
					Var: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 178,
								EndLine:   178,
								StartPos:  3998,
								EndPos:    4003,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 178,
									EndLine:   178,
									StartPos:  3998,
									EndPos:    4000,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 178,
										EndLine:   178,
										StartPos:  3998,
										EndPos:    4000,
									},
								},
								Value: []byte("a"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 178,
									EndLine:   178,
									StartPos:  4001,
									EndPos:    4002,
								},
							},
							Value: []byte("1"),
						},
					},
					Dim: &ast.ScalarLnumber{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 178,
								EndLine:   178,
								StartPos:  4004,
								EndPos:    4005,
							},
						},
						Value: []byte("2"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 179,
						EndLine:   179,
						StartPos:  4010,
						EndPos:    4018,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 179,
							EndLine:   179,
							StartPos:  4010,
							EndPos:    4017,
						},
					},
					Items: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 180,
						EndLine:   180,
						StartPos:  4021,
						EndPos:    4030,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 180,
							EndLine:   180,
							StartPos:  4021,
							EndPos:    4029,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 180,
									EndLine:   180,
									StartPos:  4027,
									EndPos:    4028,
								},
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 180,
										EndLine:   180,
										StartPos:  4027,
										EndPos:    4028,
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
						StartLine: 181,
						EndLine:   181,
						StartPos:  4033,
						EndPos:    4051,
					},
				},
				Expr: &ast.ExprArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 181,
							EndLine:   181,
							StartPos:  4033,
							EndPos:    4050,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 181,
									EndLine:   181,
									StartPos:  4039,
									EndPos:    4043,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 181,
										EndLine:   181,
										StartPos:  4039,
										EndPos:    4040,
									},
								},
								Value: []byte("1"),
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 181,
										EndLine:   181,
										StartPos:  4042,
										EndPos:    4043,
									},
								},
								Value: []byte("1"),
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 181,
									EndLine:   181,
									StartPos:  4045,
									EndPos:    4048,
								},
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 181,
										EndLine:   181,
										StartPos:  4045,
										EndPos:    4048,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 181,
											EndLine:   181,
											StartPos:  4046,
											EndPos:    4048,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 181,
												EndLine:   181,
												StartPos:  4046,
												EndPos:    4048,
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
						StartLine: 182,
						EndLine:   182,
						StartPos:  4054,
						EndPos:    4058,
					},
				},
				Expr: &ast.ExprBitwiseNot{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 182,
							EndLine:   182,
							StartPos:  4054,
							EndPos:    4057,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 182,
								EndLine:   182,
								StartPos:  4055,
								EndPos:    4057,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 182,
									EndLine:   182,
									StartPos:  4055,
									EndPos:    4057,
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
						StartLine: 183,
						EndLine:   183,
						StartPos:  4061,
						EndPos:    4065,
					},
				},
				Expr: &ast.ExprBooleanNot{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 183,
							EndLine:   183,
							StartPos:  4061,
							EndPos:    4064,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 183,
								EndLine:   183,
								StartPos:  4062,
								EndPos:    4064,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 183,
									EndLine:   183,
									StartPos:  4062,
									EndPos:    4064,
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
						StartLine: 185,
						EndLine:   185,
						StartPos:  4069,
						EndPos:    4078,
					},
				},
				Expr: &ast.ExprClassConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 185,
							EndLine:   185,
							StartPos:  4069,
							EndPos:    4077,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 185,
								EndLine:   185,
								StartPos:  4069,
								EndPos:    4072,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 185,
										EndLine:   185,
										StartPos:  4069,
										EndPos:    4072,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ConstantName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 185,
								EndLine:   185,
								StartPos:  4074,
								EndPos:    4077,
							},
						},
						Value: []byte("Bar"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 186,
						EndLine:   186,
						StartPos:  4081,
						EndPos:    4091,
					},
				},
				Expr: &ast.ExprClassConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 186,
							EndLine:   186,
							StartPos:  4081,
							EndPos:    4090,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 186,
								EndLine:   186,
								StartPos:  4081,
								EndPos:    4085,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 186,
									EndLine:   186,
									StartPos:  4081,
									EndPos:    4085,
								},
							},
							Value: []byte("foo"),
						},
					},
					ConstantName: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 186,
								EndLine:   186,
								StartPos:  4087,
								EndPos:    4090,
							},
						},
						Value: []byte("Bar"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 187,
						EndLine:   187,
						StartPos:  4094,
						EndPos:    4104,
					},
				},
				Expr: &ast.ExprClone{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 187,
							EndLine:   187,
							StartPos:  4094,
							EndPos:    4102,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 187,
								EndLine:   187,
								StartPos:  4100,
								EndPos:    4102,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 187,
									EndLine:   187,
									StartPos:  4100,
									EndPos:    4102,
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
						StartPos:  4107,
						EndPos:    4116,
					},
				},
				Expr: &ast.ExprClone{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 188,
							EndLine:   188,
							StartPos:  4107,
							EndPos:    4115,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 188,
								EndLine:   188,
								StartPos:  4113,
								EndPos:    4115,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 188,
									EndLine:   188,
									StartPos:  4113,
									EndPos:    4115,
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
						StartPos:  4119,
						EndPos:    4132,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 189,
							EndLine:   189,
							StartPos:  4119,
							EndPos:    4131,
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
						StartLine: 190,
						EndLine:   190,
						StartPos:  4135,
						EndPos:    4169,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 190,
							EndLine:   190,
							StartPos:  4135,
							EndPos:    4168,
						},
					},
					Static:     false,
					ReturnsRef: false,
					Params: []ast.Vertex{
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 190,
									EndLine:   190,
									StartPos:  4144,
									EndPos:    4146,
								},
							},
							ByRef:    false,
							Variadic: false,
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 190,
										EndLine:   190,
										StartPos:  4144,
										EndPos:    4146,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 190,
											EndLine:   190,
											StartPos:  4144,
											EndPos:    4146,
										},
									},
									Value: []byte("a"),
								},
							},
						},
						&ast.Parameter{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 190,
									EndLine:   190,
									StartPos:  4148,
									EndPos:    4150,
								},
							},
							ByRef:    false,
							Variadic: false,
							Var: &ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 190,
										EndLine:   190,
										StartPos:  4148,
										EndPos:    4150,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 190,
											EndLine:   190,
											StartPos:  4148,
											EndPos:    4150,
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
								StartLine: 190,
								EndLine:   190,
								StartPos:  4152,
								EndPos:    4165,
							},
						},
						Uses: []ast.Vertex{
							&ast.ExprVariable{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 190,
										EndLine:   190,
										StartPos:  4157,
										EndPos:    4159,
									},
								},
								VarName: &ast.Identifier{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 190,
											EndLine:   190,
											StartPos:  4157,
											EndPos:    4159,
										},
									},
									Value: []byte("c"),
								},
							},
							&ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 190,
										EndLine:   190,
										StartPos:  4161,
										EndPos:    4164,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 190,
											EndLine:   190,
											StartPos:  4162,
											EndPos:    4164,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 190,
												EndLine:   190,
												StartPos:  4162,
												EndPos:    4164,
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
						StartLine: 191,
						EndLine:   191,
						StartPos:  4172,
						EndPos:    4192,
					},
				},
				Expr: &ast.ExprClosure{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 191,
							EndLine:   191,
							StartPos:  4172,
							EndPos:    4191,
						},
					},
					ReturnsRef: false,
					Static:     false,
					ReturnType: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 191,
								EndLine:   191,
								StartPos:  4184,
								EndPos:    4188,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 191,
										EndLine:   191,
										StartPos:  4184,
										EndPos:    4188,
									},
								},
								Value: []byte("void"),
							},
						},
					},
					Stmts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 192,
						EndLine:   192,
						StartPos:  4195,
						EndPos:    4199,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 192,
							EndLine:   192,
							StartPos:  4195,
							EndPos:    4198,
						},
					},
					Const: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 192,
								EndLine:   192,
								StartPos:  4195,
								EndPos:    4198,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 192,
										EndLine:   192,
										StartPos:  4195,
										EndPos:    4198,
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
						StartLine: 193,
						EndLine:   193,
						StartPos:  4202,
						EndPos:    4216,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 193,
							EndLine:   193,
							StartPos:  4202,
							EndPos:    4215,
						},
					},
					Const: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 193,
								EndLine:   193,
								StartPos:  4202,
								EndPos:    4215,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 193,
										EndLine:   193,
										StartPos:  4212,
										EndPos:    4215,
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
						StartLine: 194,
						EndLine:   194,
						StartPos:  4219,
						EndPos:    4224,
					},
				},
				Expr: &ast.ExprConstFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 194,
							EndLine:   194,
							StartPos:  4219,
							EndPos:    4223,
						},
					},
					Const: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 194,
								EndLine:   194,
								StartPos:  4219,
								EndPos:    4223,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 194,
										EndLine:   194,
										StartPos:  4220,
										EndPos:    4223,
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
						StartLine: 196,
						EndLine:   196,
						StartPos:  4228,
						EndPos:    4238,
					},
				},
				Expr: &ast.ExprEmpty{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 196,
							EndLine:   196,
							StartPos:  4228,
							EndPos:    4237,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 196,
								EndLine:   196,
								StartPos:  4234,
								EndPos:    4236,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 196,
									EndLine:   196,
									StartPos:  4234,
									EndPos:    4236,
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
						StartLine: 197,
						EndLine:   197,
						StartPos:  4241,
						EndPos:    4245,
					},
				},
				Expr: &ast.ExprErrorSuppress{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 197,
							EndLine:   197,
							StartPos:  4241,
							EndPos:    4244,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 197,
								EndLine:   197,
								StartPos:  4242,
								EndPos:    4244,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 197,
									EndLine:   197,
									StartPos:  4242,
									EndPos:    4244,
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
						StartLine: 198,
						EndLine:   198,
						StartPos:  4248,
						EndPos:    4257,
					},
				},
				Expr: &ast.ExprEval{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 198,
							EndLine:   198,
							StartPos:  4248,
							EndPos:    4256,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 198,
								EndLine:   198,
								StartPos:  4253,
								EndPos:    4255,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 198,
									EndLine:   198,
									StartPos:  4253,
									EndPos:    4255,
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
						StartPos:  4260,
						EndPos:    4265,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 199,
							EndLine:   199,
							StartPos:  4260,
							EndPos:    4264,
						},
					},
					Die: false,
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 200,
						EndLine:   200,
						StartPos:  4268,
						EndPos:    4277,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 200,
							EndLine:   200,
							StartPos:  4268,
							EndPos:    4276,
						},
					},
					Die: false,
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 200,
								EndLine:   200,
								StartPos:  4273,
								EndPos:    4275,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 200,
									EndLine:   200,
									StartPos:  4273,
									EndPos:    4275,
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
						StartPos:  4280,
						EndPos:    4284,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 201,
							EndLine:   201,
							StartPos:  4280,
							EndPos:    4283,
						},
					},
					Die: true,
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 202,
						EndLine:   202,
						StartPos:  4287,
						EndPos:    4295,
					},
				},
				Expr: &ast.ExprExit{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 202,
							EndLine:   202,
							StartPos:  4287,
							EndPos:    4294,
						},
					},
					Die: true,
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 202,
								EndLine:   202,
								StartPos:  4291,
								EndPos:    4293,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 202,
									EndLine:   202,
									StartPos:  4291,
									EndPos:    4293,
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
						StartPos:  4298,
						EndPos:    4304,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 203,
							EndLine:   203,
							StartPos:  4298,
							EndPos:    4303,
						},
					},
					Function: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 203,
								EndLine:   203,
								StartPos:  4298,
								EndPos:    4301,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 203,
										EndLine:   203,
										StartPos:  4298,
										EndPos:    4301,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 203,
								EndLine:   203,
								StartPos:  4301,
								EndPos:    4303,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 204,
						EndLine:   204,
						StartPos:  4307,
						EndPos:    4323,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 204,
							EndLine:   204,
							StartPos:  4307,
							EndPos:    4322,
						},
					},
					Function: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 204,
								EndLine:   204,
								StartPos:  4307,
								EndPos:    4320,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 204,
										EndLine:   204,
										StartPos:  4317,
										EndPos:    4320,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 204,
								EndLine:   204,
								StartPos:  4320,
								EndPos:    4322,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 205,
						EndLine:   205,
						StartPos:  4326,
						EndPos:    4333,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 205,
							EndLine:   205,
							StartPos:  4326,
							EndPos:    4332,
						},
					},
					Function: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 205,
								EndLine:   205,
								StartPos:  4326,
								EndPos:    4330,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 205,
										EndLine:   205,
										StartPos:  4327,
										EndPos:    4330,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 205,
								EndLine:   205,
								StartPos:  4330,
								EndPos:    4332,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 206,
						EndLine:   206,
						StartPos:  4336,
						EndPos:    4343,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 206,
							EndLine:   206,
							StartPos:  4336,
							EndPos:    4342,
						},
					},
					Function: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 206,
								EndLine:   206,
								StartPos:  4336,
								EndPos:    4340,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 206,
									EndLine:   206,
									StartPos:  4336,
									EndPos:    4340,
								},
							},
							Value: []byte("foo"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 206,
								EndLine:   206,
								StartPos:  4340,
								EndPos:    4342,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 208,
						EndLine:   208,
						StartPos:  4347,
						EndPos:    4352,
					},
				},
				Expr: &ast.ExprPostDec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 208,
							EndLine:   208,
							StartPos:  4347,
							EndPos:    4351,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 208,
								EndLine:   208,
								StartPos:  4347,
								EndPos:    4349,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 208,
									EndLine:   208,
									StartPos:  4347,
									EndPos:    4349,
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
						StartLine: 209,
						EndLine:   209,
						StartPos:  4355,
						EndPos:    4360,
					},
				},
				Expr: &ast.ExprPostInc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 209,
							EndLine:   209,
							StartPos:  4355,
							EndPos:    4359,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 209,
								EndLine:   209,
								StartPos:  4355,
								EndPos:    4357,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 209,
									EndLine:   209,
									StartPos:  4355,
									EndPos:    4357,
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
						StartLine: 210,
						EndLine:   210,
						StartPos:  4363,
						EndPos:    4368,
					},
				},
				Expr: &ast.ExprPreDec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 210,
							EndLine:   210,
							StartPos:  4363,
							EndPos:    4367,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 210,
								EndLine:   210,
								StartPos:  4365,
								EndPos:    4367,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 210,
									EndLine:   210,
									StartPos:  4365,
									EndPos:    4367,
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
						StartLine: 211,
						EndLine:   211,
						StartPos:  4371,
						EndPos:    4376,
					},
				},
				Expr: &ast.ExprPreInc{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 211,
							EndLine:   211,
							StartPos:  4371,
							EndPos:    4375,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 211,
								EndLine:   211,
								StartPos:  4373,
								EndPos:    4375,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 211,
									EndLine:   211,
									StartPos:  4373,
									EndPos:    4375,
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
						StartLine: 213,
						EndLine:   213,
						StartPos:  4380,
						EndPos:    4391,
					},
				},
				Expr: &ast.ExprInclude{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 213,
							EndLine:   213,
							StartPos:  4380,
							EndPos:    4390,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 213,
								EndLine:   213,
								StartPos:  4388,
								EndPos:    4390,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 213,
									EndLine:   213,
									StartPos:  4388,
									EndPos:    4390,
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
						StartLine: 214,
						EndLine:   214,
						StartPos:  4394,
						EndPos:    4410,
					},
				},
				Expr: &ast.ExprIncludeOnce{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 214,
							EndLine:   214,
							StartPos:  4394,
							EndPos:    4409,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 214,
								EndLine:   214,
								StartPos:  4407,
								EndPos:    4409,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 214,
									EndLine:   214,
									StartPos:  4407,
									EndPos:    4409,
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
						StartLine: 215,
						EndLine:   215,
						StartPos:  4413,
						EndPos:    4424,
					},
				},
				Expr: &ast.ExprRequire{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 215,
							EndLine:   215,
							StartPos:  4413,
							EndPos:    4423,
						},
					},
					Expr: &ast.ExprVariable{
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
							Value: []byte("a"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 216,
						EndLine:   216,
						StartPos:  4427,
						EndPos:    4443,
					},
				},
				Expr: &ast.ExprRequireOnce{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 216,
							EndLine:   216,
							StartPos:  4427,
							EndPos:    4442,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 216,
								EndLine:   216,
								StartPos:  4440,
								EndPos:    4442,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 216,
									EndLine:   216,
									StartPos:  4440,
									EndPos:    4442,
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
						StartLine: 218,
						EndLine:   218,
						StartPos:  4447,
						EndPos:    4465,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 218,
							EndLine:   218,
							StartPos:  4447,
							EndPos:    4464,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 218,
								EndLine:   218,
								StartPos:  4447,
								EndPos:    4449,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 218,
									EndLine:   218,
									StartPos:  4447,
									EndPos:    4449,
								},
							},
							Value: []byte("a"),
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 218,
								EndLine:   218,
								StartPos:  4461,
								EndPos:    4464,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 218,
										EndLine:   218,
										StartPos:  4461,
										EndPos:    4464,
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
						StartLine: 219,
						EndLine:   219,
						StartPos:  4468,
						EndPos:    4496,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 219,
							EndLine:   219,
							StartPos:  4468,
							EndPos:    4495,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 219,
								EndLine:   219,
								StartPos:  4468,
								EndPos:    4470,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 219,
									EndLine:   219,
									StartPos:  4468,
									EndPos:    4470,
								},
							},
							Value: []byte("a"),
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 219,
								EndLine:   219,
								StartPos:  4482,
								EndPos:    4495,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 219,
										EndLine:   219,
										StartPos:  4492,
										EndPos:    4495,
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
						StartLine: 220,
						EndLine:   220,
						StartPos:  4499,
						EndPos:    4518,
					},
				},
				Expr: &ast.ExprInstanceOf{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 220,
							EndLine:   220,
							StartPos:  4499,
							EndPos:    4517,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 220,
								EndLine:   220,
								StartPos:  4499,
								EndPos:    4501,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 220,
									EndLine:   220,
									StartPos:  4499,
									EndPos:    4501,
								},
							},
							Value: []byte("a"),
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 220,
								EndLine:   220,
								StartPos:  4513,
								EndPos:    4517,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 220,
										EndLine:   220,
										StartPos:  4514,
										EndPos:    4517,
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
						StartLine: 222,
						EndLine:   222,
						StartPos:  4522,
						EndPos:    4536,
					},
				},
				Expr: &ast.ExprIsset{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 222,
							EndLine:   222,
							StartPos:  4522,
							EndPos:    4535,
						},
					},
					Vars: []ast.Vertex{
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 222,
									EndLine:   222,
									StartPos:  4528,
									EndPos:    4530,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 222,
										EndLine:   222,
										StartPos:  4528,
										EndPos:    4530,
									},
								},
								Value: []byte("a"),
							},
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 222,
									EndLine:   222,
									StartPos:  4532,
									EndPos:    4534,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 222,
										EndLine:   222,
										StartPos:  4532,
										EndPos:    4534,
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
						StartLine: 223,
						EndLine:   223,
						StartPos:  4539,
						EndPos:    4553,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 223,
							EndLine:   223,
							StartPos:  4539,
							EndPos:    4552,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 223,
								EndLine:   223,
								StartPos:  4539,
								EndPos:    4547,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 223,
										EndLine:   223,
										StartPos:  4544,
										EndPos:    4546,
									},
								},
								Val: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 223,
											EndLine:   223,
											StartPos:  4544,
											EndPos:    4546,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 223,
												EndLine:   223,
												StartPos:  4544,
												EndPos:    4546,
											},
										},
										Value: []byte("a"),
									},
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 223,
								EndLine:   223,
								StartPos:  4550,
								EndPos:    4552,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 223,
									EndLine:   223,
									StartPos:  4550,
									EndPos:    4552,
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
						StartLine: 224,
						EndLine:   224,
						StartPos:  4556,
						EndPos:    4572,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 224,
							EndLine:   224,
							StartPos:  4556,
							EndPos:    4571,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 224,
								EndLine:   224,
								StartPos:  4556,
								EndPos:    4566,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 224,
										EndLine:   224,
										StartPos:  4561,
										EndPos:    4565,
									},
								},
								Val: &ast.ExprArrayDimFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 224,
											EndLine:   224,
											StartPos:  4561,
											EndPos:    4565,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 224,
												EndLine:   224,
												StartPos:  4561,
												EndPos:    4563,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 224,
													EndLine:   224,
													StartPos:  4561,
													EndPos:    4563,
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
								StartLine: 224,
								EndLine:   224,
								StartPos:  4569,
								EndPos:    4571,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 224,
									EndLine:   224,
									StartPos:  4569,
									EndPos:    4571,
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
						StartLine: 225,
						EndLine:   225,
						StartPos:  4575,
						EndPos:    4595,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 225,
							EndLine:   225,
							StartPos:  4575,
							EndPos:    4594,
						},
					},
					Var: &ast.ExprList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 225,
								EndLine:   225,
								StartPos:  4575,
								EndPos:    4589,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 225,
										EndLine:   225,
										StartPos:  4580,
										EndPos:    4588,
									},
								},
								Val: &ast.ExprList{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 225,
											EndLine:   225,
											StartPos:  4580,
											EndPos:    4588,
										},
									},
									Items: []ast.Vertex{
										&ast.ExprArrayItem{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 225,
													EndLine:   225,
													StartPos:  4585,
													EndPos:    4587,
												},
											},
											Val: &ast.ExprVariable{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 225,
														EndLine:   225,
														StartPos:  4585,
														EndPos:    4587,
													},
												},
												VarName: &ast.Identifier{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 225,
															EndLine:   225,
															StartPos:  4585,
															EndPos:    4587,
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
								StartLine: 225,
								EndLine:   225,
								StartPos:  4592,
								EndPos:    4594,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 225,
									EndLine:   225,
									StartPos:  4592,
									EndPos:    4594,
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
						StartLine: 227,
						EndLine:   227,
						StartPos:  4599,
						EndPos:    4609,
					},
				},
				Expr: &ast.ExprMethodCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 227,
							EndLine:   227,
							StartPos:  4599,
							EndPos:    4608,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 227,
								EndLine:   227,
								StartPos:  4599,
								EndPos:    4601,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 227,
									EndLine:   227,
									StartPos:  4599,
									EndPos:    4601,
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
								StartPos:  4603,
								EndPos:    4606,
							},
						},
						Value: []byte("foo"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 227,
								EndLine:   227,
								StartPos:  4606,
								EndPos:    4608,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 228,
						EndLine:   228,
						StartPos:  4612,
						EndPos:    4622,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 228,
							EndLine:   228,
							StartPos:  4612,
							EndPos:    4621,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 228,
								EndLine:   228,
								StartPos:  4616,
								EndPos:    4619,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 228,
										EndLine:   228,
										StartPos:  4616,
										EndPos:    4619,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 228,
								EndLine:   228,
								StartPos:  4619,
								EndPos:    4621,
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
						StartPos:  4625,
						EndPos:    4645,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 229,
							EndLine:   229,
							StartPos:  4625,
							EndPos:    4644,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 229,
								EndLine:   229,
								StartPos:  4629,
								EndPos:    4642,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 229,
										EndLine:   229,
										StartPos:  4639,
										EndPos:    4642,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 229,
								EndLine:   229,
								StartPos:  4642,
								EndPos:    4644,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 230,
						EndLine:   230,
						StartPos:  4648,
						EndPos:    4659,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 230,
							EndLine:   230,
							StartPos:  4648,
							EndPos:    4658,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 230,
								EndLine:   230,
								StartPos:  4652,
								EndPos:    4656,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 230,
										EndLine:   230,
										StartPos:  4653,
										EndPos:    4656,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 230,
								EndLine:   230,
								StartPos:  4656,
								EndPos:    4658,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 231,
						EndLine:   231,
						StartPos:  4662,
						EndPos:    4687,
					},
				},
				Expr: &ast.ExprNew{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 231,
							EndLine:   231,
							StartPos:  4662,
							EndPos:    4686,
						},
					},
					Class: &ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 231,
								EndLine:   231,
								StartPos:  4666,
								EndPos:    4686,
							},
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 231,
									EndLine:   231,
									StartPos:  4672,
									EndPos:    4683,
								},
							},
							Arguments: []ast.Vertex{
								&ast.Argument{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 231,
											EndLine:   231,
											StartPos:  4673,
											EndPos:    4675,
										},
									},
									IsReference: false,
									Variadic:    false,
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 231,
												EndLine:   231,
												StartPos:  4673,
												EndPos:    4675,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 231,
													EndLine:   231,
													StartPos:  4673,
													EndPos:    4675,
												},
											},
											Value: []byte("a"),
										},
									},
								},
								&ast.Argument{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 231,
											EndLine:   231,
											StartPos:  4677,
											EndPos:    4682,
										},
									},
									IsReference: false,
									Variadic:    true,
									Expr: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 231,
												EndLine:   231,
												StartPos:  4680,
												EndPos:    4682,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 231,
													EndLine:   231,
													StartPos:  4680,
													EndPos:    4682,
												},
											},
											Value: []byte("b"),
										},
									},
								},
							},
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 232,
						EndLine:   232,
						StartPos:  4690,
						EndPos:    4700,
					},
				},
				Expr: &ast.ExprPrint{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 232,
							EndLine:   232,
							StartPos:  4690,
							EndPos:    4698,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 232,
								EndLine:   232,
								StartPos:  4696,
								EndPos:    4698,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 232,
									EndLine:   232,
									StartPos:  4696,
									EndPos:    4698,
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
						StartLine: 233,
						EndLine:   233,
						StartPos:  4703,
						EndPos:    4711,
					},
				},
				Expr: &ast.ExprPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 233,
							EndLine:   233,
							StartPos:  4703,
							EndPos:    4710,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 233,
								EndLine:   233,
								StartPos:  4703,
								EndPos:    4705,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 233,
									EndLine:   233,
									StartPos:  4703,
									EndPos:    4705,
								},
							},
							Value: []byte("a"),
						},
					},
					Property: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 233,
								EndLine:   233,
								StartPos:  4707,
								EndPos:    4710,
							},
						},
						Value: []byte("foo"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 234,
						EndLine:   234,
						StartPos:  4714,
						EndPos:    4723,
					},
				},
				Expr: &ast.ExprShellExec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 234,
							EndLine:   234,
							StartPos:  4714,
							EndPos:    4722,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 234,
									EndLine:   234,
									StartPos:  4715,
									EndPos:    4719,
								},
							},
							Value: []byte("cmd "),
						},
						&ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 234,
									EndLine:   234,
									StartPos:  4719,
									EndPos:    4721,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 234,
										EndLine:   234,
										StartPos:  4719,
										EndPos:    4721,
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
						StartLine: 235,
						EndLine:   235,
						StartPos:  4726,
						EndPos:    4732,
					},
				},
				Expr: &ast.ExprShellExec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 235,
							EndLine:   235,
							StartPos:  4726,
							EndPos:    4731,
						},
					},
					Parts: []ast.Vertex{
						&ast.ScalarEncapsedStringPart{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 235,
									EndLine:   235,
									StartPos:  4727,
									EndPos:    4730,
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
						StartLine: 236,
						EndLine:   236,
						StartPos:  4735,
						EndPos:    4738,
					},
				},
				Expr: &ast.ExprShellExec{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 236,
							EndLine:   236,
							StartPos:  4735,
							EndPos:    4737,
						},
					},
					Parts: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 237,
						EndLine:   237,
						StartPos:  4741,
						EndPos:    4744,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 237,
							EndLine:   237,
							StartPos:  4741,
							EndPos:    4743,
						},
					},
					Items: []ast.Vertex{},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 238,
						EndLine:   238,
						StartPos:  4747,
						EndPos:    4751,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 238,
							EndLine:   238,
							StartPos:  4747,
							EndPos:    4750,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 238,
									EndLine:   238,
									StartPos:  4748,
									EndPos:    4749,
								},
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 238,
										EndLine:   238,
										StartPos:  4748,
										EndPos:    4749,
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
						StartLine: 239,
						EndLine:   239,
						StartPos:  4754,
						EndPos:    4767,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 239,
							EndLine:   239,
							StartPos:  4754,
							EndPos:    4766,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 239,
									EndLine:   239,
									StartPos:  4755,
									EndPos:    4759,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 239,
										EndLine:   239,
										StartPos:  4755,
										EndPos:    4756,
									},
								},
								Value: []byte("1"),
							},
							Val: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 239,
										EndLine:   239,
										StartPos:  4758,
										EndPos:    4759,
									},
								},
								Value: []byte("1"),
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 239,
									EndLine:   239,
									StartPos:  4761,
									EndPos:    4764,
								},
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 239,
										EndLine:   239,
										StartPos:  4761,
										EndPos:    4764,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 239,
											EndLine:   239,
											StartPos:  4762,
											EndPos:    4764,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 239,
												EndLine:   239,
												StartPos:  4762,
												EndPos:    4764,
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
						StartLine: 241,
						EndLine:   241,
						StartPos:  4771,
						EndPos:    4781,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 241,
							EndLine:   241,
							StartPos:  4771,
							EndPos:    4780,
						},
					},
					Var: &ast.ExprShortList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 241,
								EndLine:   241,
								StartPos:  4771,
								EndPos:    4775,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 241,
										EndLine:   241,
										StartPos:  4772,
										EndPos:    4774,
									},
								},
								Val: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 241,
											EndLine:   241,
											StartPos:  4772,
											EndPos:    4774,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 241,
												EndLine:   241,
												StartPos:  4772,
												EndPos:    4774,
											},
										},
										Value: []byte("a"),
									},
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 241,
								EndLine:   241,
								StartPos:  4778,
								EndPos:    4780,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 241,
									EndLine:   241,
									StartPos:  4778,
									EndPos:    4780,
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
						StartLine: 242,
						EndLine:   242,
						StartPos:  4784,
						EndPos:    4796,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 242,
							EndLine:   242,
							StartPos:  4784,
							EndPos:    4795,
						},
					},
					Var: &ast.ExprShortList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 242,
								EndLine:   242,
								StartPos:  4784,
								EndPos:    4790,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 242,
										EndLine:   242,
										StartPos:  4785,
										EndPos:    4789,
									},
								},
								Val: &ast.ExprArrayDimFetch{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 242,
											EndLine:   242,
											StartPos:  4785,
											EndPos:    4789,
										},
									},
									Var: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 242,
												EndLine:   242,
												StartPos:  4785,
												EndPos:    4787,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 242,
													EndLine:   242,
													StartPos:  4785,
													EndPos:    4787,
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
								StartLine: 242,
								EndLine:   242,
								StartPos:  4793,
								EndPos:    4795,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 242,
									EndLine:   242,
									StartPos:  4793,
									EndPos:    4795,
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
						StartLine: 243,
						EndLine:   243,
						StartPos:  4799,
						EndPos:    4815,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 243,
							EndLine:   243,
							StartPos:  4799,
							EndPos:    4814,
						},
					},
					Var: &ast.ExprShortList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 243,
								EndLine:   243,
								StartPos:  4799,
								EndPos:    4809,
							},
						},
						Items: []ast.Vertex{
							&ast.ExprArrayItem{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 243,
										EndLine:   243,
										StartPos:  4800,
										EndPos:    4808,
									},
								},
								Val: &ast.ExprList{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 243,
											EndLine:   243,
											StartPos:  4800,
											EndPos:    4808,
										},
									},
									Items: []ast.Vertex{
										&ast.ExprArrayItem{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 243,
													EndLine:   243,
													StartPos:  4805,
													EndPos:    4807,
												},
											},
											Val: &ast.ExprVariable{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 243,
														EndLine:   243,
														StartPos:  4805,
														EndPos:    4807,
													},
												},
												VarName: &ast.Identifier{
													Node: ast.Node{
														Position: &position.Position{
															StartLine: 243,
															EndLine:   243,
															StartPos:  4805,
															EndPos:    4807,
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
								StartLine: 243,
								EndLine:   243,
								StartPos:  4812,
								EndPos:    4814,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 243,
									EndLine:   243,
									StartPos:  4812,
									EndPos:    4814,
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
						StartLine: 244,
						EndLine:   244,
						StartPos:  4818,
						EndPos:    4829,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 244,
							EndLine:   244,
							StartPos:  4818,
							EndPos:    4828,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 244,
								EndLine:   244,
								StartPos:  4818,
								EndPos:    4821,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 244,
										EndLine:   244,
										StartPos:  4818,
										EndPos:    4821,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 244,
								EndLine:   244,
								StartPos:  4823,
								EndPos:    4826,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 244,
								EndLine:   244,
								StartPos:  4826,
								EndPos:    4828,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 245,
						EndLine:   245,
						StartPos:  4832,
						EndPos:    4853,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 245,
							EndLine:   245,
							StartPos:  4832,
							EndPos:    4852,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4832,
								EndPos:    4845,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 245,
										EndLine:   245,
										StartPos:  4842,
										EndPos:    4845,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4847,
								EndPos:    4850,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 245,
								EndLine:   245,
								StartPos:  4850,
								EndPos:    4852,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 246,
						EndLine:   246,
						StartPos:  4856,
						EndPos:    4868,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 246,
							EndLine:   246,
							StartPos:  4856,
							EndPos:    4867,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4856,
								EndPos:    4860,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 246,
										EndLine:   246,
										StartPos:  4857,
										EndPos:    4860,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4862,
								EndPos:    4865,
							},
						},
						Value: []byte("bar"),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 246,
								EndLine:   246,
								StartPos:  4865,
								EndPos:    4867,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 247,
						EndLine:   247,
						StartPos:  4871,
						EndPos:    4881,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 247,
							EndLine:   247,
							StartPos:  4871,
							EndPos:    4880,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 247,
								EndLine:   247,
								StartPos:  4871,
								EndPos:    4874,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 247,
										EndLine:   247,
										StartPos:  4871,
										EndPos:    4874,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 247,
								EndLine:   247,
								StartPos:  4876,
								EndPos:    4880,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 247,
									EndLine:   247,
									StartPos:  4876,
									EndPos:    4880,
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
						StartLine: 248,
						EndLine:   248,
						StartPos:  4884,
						EndPos:    4895,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 248,
							EndLine:   248,
							StartPos:  4884,
							EndPos:    4894,
						},
					},
					Class: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4884,
								EndPos:    4888,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 248,
									EndLine:   248,
									StartPos:  4884,
									EndPos:    4888,
								},
							},
							Value: []byte("foo"),
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 248,
								EndLine:   248,
								StartPos:  4890,
								EndPos:    4894,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 248,
									EndLine:   248,
									StartPos:  4890,
									EndPos:    4894,
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
						StartLine: 249,
						EndLine:   249,
						StartPos:  4898,
						EndPos:    4918,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 249,
							EndLine:   249,
							StartPos:  4898,
							EndPos:    4917,
						},
					},
					Class: &ast.NameRelative{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 249,
								EndLine:   249,
								StartPos:  4898,
								EndPos:    4911,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 249,
										EndLine:   249,
										StartPos:  4908,
										EndPos:    4911,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 249,
								EndLine:   249,
								StartPos:  4913,
								EndPos:    4917,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 249,
									EndLine:   249,
									StartPos:  4913,
									EndPos:    4917,
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
						StartLine: 250,
						EndLine:   250,
						StartPos:  4921,
						EndPos:    4932,
					},
				},
				Expr: &ast.ExprStaticPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 250,
							EndLine:   250,
							StartPos:  4921,
							EndPos:    4931,
						},
					},
					Class: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 250,
								EndLine:   250,
								StartPos:  4921,
								EndPos:    4925,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 250,
										EndLine:   250,
										StartPos:  4922,
										EndPos:    4925,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 250,
								EndLine:   250,
								StartPos:  4927,
								EndPos:    4931,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 250,
									EndLine:   250,
									StartPos:  4927,
									EndPos:    4931,
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
						StartLine: 251,
						EndLine:   251,
						StartPos:  4935,
						EndPos:    4948,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 251,
							EndLine:   251,
							StartPos:  4935,
							EndPos:    4947,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 251,
								EndLine:   251,
								StartPos:  4935,
								EndPos:    4937,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 251,
									EndLine:   251,
									StartPos:  4935,
									EndPos:    4937,
								},
							},
							Value: []byte("a"),
						},
					},
					IfTrue: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 251,
								EndLine:   251,
								StartPos:  4940,
								EndPos:    4942,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 251,
									EndLine:   251,
									StartPos:  4940,
									EndPos:    4942,
								},
							},
							Value: []byte("b"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 251,
								EndLine:   251,
								StartPos:  4945,
								EndPos:    4947,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 251,
									EndLine:   251,
									StartPos:  4945,
									EndPos:    4947,
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
						StartLine: 252,
						EndLine:   252,
						StartPos:  4951,
						EndPos:    4961,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 252,
							EndLine:   252,
							StartPos:  4951,
							EndPos:    4960,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 252,
								EndLine:   252,
								StartPos:  4951,
								EndPos:    4953,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 252,
									EndLine:   252,
									StartPos:  4951,
									EndPos:    4953,
								},
							},
							Value: []byte("a"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 252,
								EndLine:   252,
								StartPos:  4958,
								EndPos:    4960,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 252,
									EndLine:   252,
									StartPos:  4958,
									EndPos:    4960,
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
						StartLine: 253,
						EndLine:   253,
						StartPos:  4964,
						EndPos:    4987,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 253,
							EndLine:   253,
							StartPos:  4964,
							EndPos:    4986,
						},
					},
					Condition: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4964,
								EndPos:    4966,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4964,
									EndPos:    4966,
								},
							},
							Value: []byte("a"),
						},
					},
					IfTrue: &ast.ExprTernary{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4969,
								EndPos:    4981,
							},
						},
						Condition: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4969,
									EndPos:    4971,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 253,
										EndLine:   253,
										StartPos:  4969,
										EndPos:    4971,
									},
								},
								Value: []byte("b"),
							},
						},
						IfTrue: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4974,
									EndPos:    4976,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 253,
										EndLine:   253,
										StartPos:  4974,
										EndPos:    4976,
									},
								},
								Value: []byte("c"),
							},
						},
						IfFalse: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4979,
									EndPos:    4981,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 253,
										EndLine:   253,
										StartPos:  4979,
										EndPos:    4981,
									},
								},
								Value: []byte("d"),
							},
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 253,
								EndLine:   253,
								StartPos:  4984,
								EndPos:    4986,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 253,
									EndLine:   253,
									StartPos:  4984,
									EndPos:    4986,
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
						StartLine: 254,
						EndLine:   254,
						StartPos:  4990,
						EndPos:    5013,
					},
				},
				Expr: &ast.ExprTernary{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 254,
							EndLine:   254,
							StartPos:  4990,
							EndPos:    5012,
						},
					},
					Condition: &ast.ExprTernary{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  4990,
								EndPos:    5002,
							},
						},
						Condition: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4990,
									EndPos:    4992,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 254,
										EndLine:   254,
										StartPos:  4990,
										EndPos:    4992,
									},
								},
								Value: []byte("a"),
							},
						},
						IfTrue: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  4995,
									EndPos:    4997,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 254,
										EndLine:   254,
										StartPos:  4995,
										EndPos:    4997,
									},
								},
								Value: []byte("b"),
							},
						},
						IfFalse: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  5000,
									EndPos:    5002,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 254,
										EndLine:   254,
										StartPos:  5000,
										EndPos:    5002,
									},
								},
								Value: []byte("c"),
							},
						},
					},
					IfTrue: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  5005,
								EndPos:    5007,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  5005,
									EndPos:    5007,
								},
							},
							Value: []byte("d"),
						},
					},
					IfFalse: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 254,
								EndLine:   254,
								StartPos:  5010,
								EndPos:    5012,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 254,
									EndLine:   254,
									StartPos:  5010,
									EndPos:    5012,
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
						StartLine: 255,
						EndLine:   255,
						StartPos:  5016,
						EndPos:    5020,
					},
				},
				Expr: &ast.ExprUnaryMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 255,
							EndLine:   255,
							StartPos:  5016,
							EndPos:    5019,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 255,
								EndLine:   255,
								StartPos:  5017,
								EndPos:    5019,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 255,
									EndLine:   255,
									StartPos:  5017,
									EndPos:    5019,
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
						StartLine: 256,
						EndLine:   256,
						StartPos:  5023,
						EndPos:    5027,
					},
				},
				Expr: &ast.ExprUnaryPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 256,
							EndLine:   256,
							StartPos:  5023,
							EndPos:    5026,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 256,
								EndLine:   256,
								StartPos:  5024,
								EndPos:    5026,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 256,
									EndLine:   256,
									StartPos:  5024,
									EndPos:    5026,
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
						StartLine: 257,
						EndLine:   257,
						StartPos:  5030,
						EndPos:    5034,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 257,
							EndLine:   257,
							StartPos:  5030,
							EndPos:    5033,
						},
					},
					VarName: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 257,
								EndLine:   257,
								StartPos:  5031,
								EndPos:    5033,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 257,
									EndLine:   257,
									StartPos:  5031,
									EndPos:    5033,
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
						StartPos:  5037,
						EndPos:    5043,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 258,
							EndLine:   258,
							StartPos:  5037,
							EndPos:    5042,
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 259,
						EndLine:   259,
						StartPos:  5046,
						EndPos:    5055,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 259,
							EndLine:   259,
							StartPos:  5046,
							EndPos:    5054,
						},
					},
					Value: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 259,
								EndLine:   259,
								StartPos:  5052,
								EndPos:    5054,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 259,
									EndLine:   259,
									StartPos:  5052,
									EndPos:    5054,
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
						StartPos:  5058,
						EndPos:    5073,
					},
				},
				Expr: &ast.ExprYield{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 260,
							EndLine:   260,
							StartPos:  5058,
							EndPos:    5072,
						},
					},
					Key: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 260,
								EndLine:   260,
								StartPos:  5064,
								EndPos:    5066,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 260,
									EndLine:   260,
									StartPos:  5064,
									EndPos:    5066,
								},
							},
							Value: []byte("a"),
						},
					},
					Value: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 260,
								EndLine:   260,
								StartPos:  5070,
								EndPos:    5072,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 260,
									EndLine:   260,
									StartPos:  5070,
									EndPos:    5072,
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
						StartLine: 261,
						EndLine:   261,
						StartPos:  5076,
						EndPos:    5090,
					},
				},
				Expr: &ast.ExprYieldFrom{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 261,
							EndLine:   261,
							StartPos:  5076,
							EndPos:    5089,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 261,
								EndLine:   261,
								StartPos:  5087,
								EndPos:    5089,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 261,
									EndLine:   261,
									StartPos:  5087,
									EndPos:    5089,
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
						StartPos:  5096,
						EndPos:    5106,
					},
				},
				Expr: &ast.ExprCastArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 263,
							EndLine:   263,
							StartPos:  5096,
							EndPos:    5105,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 263,
								EndLine:   263,
								StartPos:  5103,
								EndPos:    5105,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 263,
									EndLine:   263,
									StartPos:  5103,
									EndPos:    5105,
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
						StartPos:  5109,
						EndPos:    5121,
					},
				},
				Expr: &ast.ExprCastBool{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 264,
							EndLine:   264,
							StartPos:  5109,
							EndPos:    5120,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 264,
								EndLine:   264,
								StartPos:  5118,
								EndPos:    5120,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 264,
									EndLine:   264,
									StartPos:  5118,
									EndPos:    5120,
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
						StartPos:  5124,
						EndPos:    5133,
					},
				},
				Expr: &ast.ExprCastBool{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 265,
							EndLine:   265,
							StartPos:  5124,
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
						EndPos:    5147,
					},
				},
				Expr: &ast.ExprCastDouble{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 266,
							EndLine:   266,
							StartPos:  5136,
							EndPos:    5146,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 266,
								EndLine:   266,
								StartPos:  5144,
								EndPos:    5146,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 266,
									EndLine:   266,
									StartPos:  5144,
									EndPos:    5146,
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
						StartLine: 267,
						EndLine:   267,
						StartPos:  5150,
						EndPos:    5160,
					},
				},
				Expr: &ast.ExprCastDouble{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 267,
							EndLine:   267,
							StartPos:  5150,
							EndPos:    5159,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 267,
								EndLine:   267,
								StartPos:  5157,
								EndPos:    5159,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 267,
									EndLine:   267,
									StartPos:  5157,
									EndPos:    5159,
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
						StartPos:  5163,
						EndPos:    5175,
					},
				},
				Expr: &ast.ExprCastInt{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 268,
							EndLine:   268,
							StartPos:  5163,
							EndPos:    5174,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 268,
								EndLine:   268,
								StartPos:  5172,
								EndPos:    5174,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 268,
									EndLine:   268,
									StartPos:  5172,
									EndPos:    5174,
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
						StartLine: 269,
						EndLine:   269,
						StartPos:  5178,
						EndPos:    5186,
					},
				},
				Expr: &ast.ExprCastInt{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 269,
							EndLine:   269,
							StartPos:  5178,
							EndPos:    5185,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 269,
								EndLine:   269,
								StartPos:  5183,
								EndPos:    5185,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 269,
									EndLine:   269,
									StartPos:  5183,
									EndPos:    5185,
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
						StartLine: 270,
						EndLine:   270,
						StartPos:  5189,
						EndPos:    5200,
					},
				},
				Expr: &ast.ExprCastObject{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 270,
							EndLine:   270,
							StartPos:  5189,
							EndPos:    5199,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 270,
								EndLine:   270,
								StartPos:  5197,
								EndPos:    5199,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 270,
									EndLine:   270,
									StartPos:  5197,
									EndPos:    5199,
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
						StartLine: 271,
						EndLine:   271,
						StartPos:  5203,
						EndPos:    5214,
					},
				},
				Expr: &ast.ExprCastString{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 271,
							EndLine:   271,
							StartPos:  5203,
							EndPos:    5213,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 271,
								EndLine:   271,
								StartPos:  5211,
								EndPos:    5213,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 271,
									EndLine:   271,
									StartPos:  5211,
									EndPos:    5213,
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
						StartLine: 272,
						EndLine:   272,
						StartPos:  5217,
						EndPos:    5227,
					},
				},
				Expr: &ast.ExprCastUnset{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 272,
							EndLine:   272,
							StartPos:  5217,
							EndPos:    5226,
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 272,
								EndLine:   272,
								StartPos:  5224,
								EndPos:    5226,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 272,
									EndLine:   272,
									StartPos:  5224,
									EndPos:    5226,
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
						StartLine: 274,
						EndLine:   274,
						StartPos:  5231,
						EndPos:    5239,
					},
				},
				Expr: &ast.ExprBinaryBitwiseAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 274,
							EndLine:   274,
							StartPos:  5231,
							EndPos:    5238,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 274,
								EndLine:   274,
								StartPos:  5231,
								EndPos:    5233,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 274,
									EndLine:   274,
									StartPos:  5231,
									EndPos:    5233,
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
								StartPos:  5236,
								EndPos:    5238,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 274,
									EndLine:   274,
									StartPos:  5236,
									EndPos:    5238,
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
						StartPos:  5242,
						EndPos:    5250,
					},
				},
				Expr: &ast.ExprBinaryBitwiseOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 275,
							EndLine:   275,
							StartPos:  5242,
							EndPos:    5249,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 275,
								EndLine:   275,
								StartPos:  5242,
								EndPos:    5244,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 275,
									EndLine:   275,
									StartPos:  5242,
									EndPos:    5244,
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
								StartPos:  5247,
								EndPos:    5249,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 275,
									EndLine:   275,
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
						StartLine: 276,
						EndLine:   276,
						StartPos:  5253,
						EndPos:    5261,
					},
				},
				Expr: &ast.ExprBinaryBitwiseXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 276,
							EndLine:   276,
							StartPos:  5253,
							EndPos:    5260,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 276,
								EndLine:   276,
								StartPos:  5253,
								EndPos:    5255,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 276,
									EndLine:   276,
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
								StartLine: 276,
								EndLine:   276,
								StartPos:  5258,
								EndPos:    5260,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 276,
									EndLine:   276,
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
						StartLine: 277,
						EndLine:   277,
						StartPos:  5264,
						EndPos:    5273,
					},
				},
				Expr: &ast.ExprBinaryBooleanAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 277,
							EndLine:   277,
							StartPos:  5264,
							EndPos:    5272,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 277,
								EndLine:   277,
								StartPos:  5264,
								EndPos:    5266,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 277,
									EndLine:   277,
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
								StartLine: 277,
								EndLine:   277,
								StartPos:  5270,
								EndPos:    5272,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 277,
									EndLine:   277,
									StartPos:  5270,
									EndPos:    5272,
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
						StartPos:  5276,
						EndPos:    5285,
					},
				},
				Expr: &ast.ExprBinaryBooleanOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 278,
							EndLine:   278,
							StartPos:  5276,
							EndPos:    5284,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 278,
								EndLine:   278,
								StartPos:  5276,
								EndPos:    5278,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 278,
									EndLine:   278,
									StartPos:  5276,
									EndPos:    5278,
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
								StartPos:  5282,
								EndPos:    5284,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 278,
									EndLine:   278,
									StartPos:  5282,
									EndPos:    5284,
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
						StartPos:  5288,
						EndPos:    5297,
					},
				},
				Expr: &ast.ExprBinaryCoalesce{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 279,
							EndLine:   279,
							StartPos:  5288,
							EndPos:    5296,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 279,
								EndLine:   279,
								StartPos:  5288,
								EndPos:    5290,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 279,
									EndLine:   279,
									StartPos:  5288,
									EndPos:    5290,
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
								StartPos:  5294,
								EndPos:    5296,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 279,
									EndLine:   279,
									StartPos:  5294,
									EndPos:    5296,
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
						StartPos:  5300,
						EndPos:    5308,
					},
				},
				Expr: &ast.ExprBinaryConcat{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 280,
							EndLine:   280,
							StartPos:  5300,
							EndPos:    5307,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 280,
								EndLine:   280,
								StartPos:  5300,
								EndPos:    5302,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 280,
									EndLine:   280,
									StartPos:  5300,
									EndPos:    5302,
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
								StartPos:  5305,
								EndPos:    5307,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 280,
									EndLine:   280,
									StartPos:  5305,
									EndPos:    5307,
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
						StartPos:  5311,
						EndPos:    5319,
					},
				},
				Expr: &ast.ExprBinaryDiv{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 281,
							EndLine:   281,
							StartPos:  5311,
							EndPos:    5318,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 281,
								EndLine:   281,
								StartPos:  5311,
								EndPos:    5313,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 281,
									EndLine:   281,
									StartPos:  5311,
									EndPos:    5313,
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
								StartPos:  5316,
								EndPos:    5318,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 281,
									EndLine:   281,
									StartPos:  5316,
									EndPos:    5318,
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
						StartPos:  5322,
						EndPos:    5331,
					},
				},
				Expr: &ast.ExprBinaryEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 282,
							EndLine:   282,
							StartPos:  5322,
							EndPos:    5330,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 282,
								EndLine:   282,
								StartPos:  5322,
								EndPos:    5324,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 282,
									EndLine:   282,
									StartPos:  5322,
									EndPos:    5324,
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
								StartPos:  5328,
								EndPos:    5330,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 282,
									EndLine:   282,
									StartPos:  5328,
									EndPos:    5330,
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
						StartPos:  5334,
						EndPos:    5343,
					},
				},
				Expr: &ast.ExprBinaryGreaterOrEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 283,
							EndLine:   283,
							StartPos:  5334,
							EndPos:    5342,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 283,
								EndLine:   283,
								StartPos:  5334,
								EndPos:    5336,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 283,
									EndLine:   283,
									StartPos:  5334,
									EndPos:    5336,
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
								StartPos:  5340,
								EndPos:    5342,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 283,
									EndLine:   283,
									StartPos:  5340,
									EndPos:    5342,
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
						StartPos:  5346,
						EndPos:    5354,
					},
				},
				Expr: &ast.ExprBinaryGreater{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 284,
							EndLine:   284,
							StartPos:  5346,
							EndPos:    5353,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 284,
								EndLine:   284,
								StartPos:  5346,
								EndPos:    5348,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 284,
									EndLine:   284,
									StartPos:  5346,
									EndPos:    5348,
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
								StartPos:  5351,
								EndPos:    5353,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 284,
									EndLine:   284,
									StartPos:  5351,
									EndPos:    5353,
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
						StartPos:  5357,
						EndPos:    5367,
					},
				},
				Expr: &ast.ExprBinaryIdentical{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 285,
							EndLine:   285,
							StartPos:  5357,
							EndPos:    5366,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 285,
								EndLine:   285,
								StartPos:  5357,
								EndPos:    5359,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 285,
									EndLine:   285,
									StartPos:  5357,
									EndPos:    5359,
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
								StartPos:  5364,
								EndPos:    5366,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 285,
									EndLine:   285,
									StartPos:  5364,
									EndPos:    5366,
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
						StartPos:  5370,
						EndPos:    5380,
					},
				},
				Expr: &ast.ExprBinaryLogicalAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 286,
							EndLine:   286,
							StartPos:  5370,
							EndPos:    5379,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 286,
								EndLine:   286,
								StartPos:  5370,
								EndPos:    5372,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 286,
									EndLine:   286,
									StartPos:  5370,
									EndPos:    5372,
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
								StartPos:  5377,
								EndPos:    5379,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 286,
									EndLine:   286,
									StartPos:  5377,
									EndPos:    5379,
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
						StartPos:  5383,
						EndPos:    5392,
					},
				},
				Expr: &ast.ExprBinaryLogicalOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 287,
							EndLine:   287,
							StartPos:  5383,
							EndPos:    5391,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 287,
								EndLine:   287,
								StartPos:  5383,
								EndPos:    5385,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 287,
									EndLine:   287,
									StartPos:  5383,
									EndPos:    5385,
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
								StartPos:  5389,
								EndPos:    5391,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 287,
									EndLine:   287,
									StartPos:  5389,
									EndPos:    5391,
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
						StartPos:  5395,
						EndPos:    5405,
					},
				},
				Expr: &ast.ExprBinaryLogicalXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 288,
							EndLine:   288,
							StartPos:  5395,
							EndPos:    5404,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 288,
								EndLine:   288,
								StartPos:  5395,
								EndPos:    5397,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 288,
									EndLine:   288,
									StartPos:  5395,
									EndPos:    5397,
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
								StartPos:  5402,
								EndPos:    5404,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 288,
									EndLine:   288,
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
						StartLine: 289,
						EndLine:   289,
						StartPos:  5408,
						EndPos:    5416,
					},
				},
				Expr: &ast.ExprBinaryMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 289,
							EndLine:   289,
							StartPos:  5408,
							EndPos:    5415,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 289,
								EndLine:   289,
								StartPos:  5408,
								EndPos:    5410,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 289,
									EndLine:   289,
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
								StartLine: 289,
								EndLine:   289,
								StartPos:  5413,
								EndPos:    5415,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 289,
									EndLine:   289,
									StartPos:  5413,
									EndPos:    5415,
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
						StartPos:  5419,
						EndPos:    5427,
					},
				},
				Expr: &ast.ExprBinaryMod{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 290,
							EndLine:   290,
							StartPos:  5419,
							EndPos:    5426,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 290,
								EndLine:   290,
								StartPos:  5419,
								EndPos:    5421,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 290,
									EndLine:   290,
									StartPos:  5419,
									EndPos:    5421,
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
								StartPos:  5424,
								EndPos:    5426,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 290,
									EndLine:   290,
									StartPos:  5424,
									EndPos:    5426,
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
						StartPos:  5430,
						EndPos:    5438,
					},
				},
				Expr: &ast.ExprBinaryMul{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 291,
							EndLine:   291,
							StartPos:  5430,
							EndPos:    5437,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 291,
								EndLine:   291,
								StartPos:  5430,
								EndPos:    5432,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 291,
									EndLine:   291,
									StartPos:  5430,
									EndPos:    5432,
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
								StartPos:  5435,
								EndPos:    5437,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 291,
									EndLine:   291,
									StartPos:  5435,
									EndPos:    5437,
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
						StartPos:  5441,
						EndPos:    5450,
					},
				},
				Expr: &ast.ExprBinaryNotEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 292,
							EndLine:   292,
							StartPos:  5441,
							EndPos:    5449,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 292,
								EndLine:   292,
								StartPos:  5441,
								EndPos:    5443,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 292,
									EndLine:   292,
									StartPos:  5441,
									EndPos:    5443,
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
								StartPos:  5447,
								EndPos:    5449,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 292,
									EndLine:   292,
									StartPos:  5447,
									EndPos:    5449,
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
						StartLine: 293,
						EndLine:   293,
						StartPos:  5453,
						EndPos:    5463,
					},
				},
				Expr: &ast.ExprBinaryNotIdentical{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 293,
							EndLine:   293,
							StartPos:  5453,
							EndPos:    5462,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 293,
								EndLine:   293,
								StartPos:  5453,
								EndPos:    5455,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 293,
									EndLine:   293,
									StartPos:  5453,
									EndPos:    5455,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 293,
								EndLine:   293,
								StartPos:  5460,
								EndPos:    5462,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 293,
									EndLine:   293,
									StartPos:  5460,
									EndPos:    5462,
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
						StartPos:  5466,
						EndPos:    5474,
					},
				},
				Expr: &ast.ExprBinaryPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 294,
							EndLine:   294,
							StartPos:  5466,
							EndPos:    5473,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  5466,
								EndPos:    5468,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 294,
									EndLine:   294,
									StartPos:  5466,
									EndPos:    5468,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 294,
								EndLine:   294,
								StartPos:  5471,
								EndPos:    5473,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 294,
									EndLine:   294,
									StartPos:  5471,
									EndPos:    5473,
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
						StartPos:  5477,
						EndPos:    5486,
					},
				},
				Expr: &ast.ExprBinaryPow{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 295,
							EndLine:   295,
							StartPos:  5477,
							EndPos:    5485,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 295,
								EndLine:   295,
								StartPos:  5477,
								EndPos:    5479,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 295,
									EndLine:   295,
									StartPos:  5477,
									EndPos:    5479,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 295,
								EndLine:   295,
								StartPos:  5483,
								EndPos:    5485,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 295,
									EndLine:   295,
									StartPos:  5483,
									EndPos:    5485,
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
						StartLine: 296,
						EndLine:   296,
						StartPos:  5489,
						EndPos:    5498,
					},
				},
				Expr: &ast.ExprBinaryShiftLeft{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 296,
							EndLine:   296,
							StartPos:  5489,
							EndPos:    5497,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  5489,
								EndPos:    5491,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 296,
									EndLine:   296,
									StartPos:  5489,
									EndPos:    5491,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 296,
								EndLine:   296,
								StartPos:  5495,
								EndPos:    5497,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 296,
									EndLine:   296,
									StartPos:  5495,
									EndPos:    5497,
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
						StartLine: 297,
						EndLine:   297,
						StartPos:  5501,
						EndPos:    5510,
					},
				},
				Expr: &ast.ExprBinaryShiftRight{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 297,
							EndLine:   297,
							StartPos:  5501,
							EndPos:    5509,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 297,
								EndLine:   297,
								StartPos:  5501,
								EndPos:    5503,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 297,
									EndLine:   297,
									StartPos:  5501,
									EndPos:    5503,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 297,
								EndLine:   297,
								StartPos:  5507,
								EndPos:    5509,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 297,
									EndLine:   297,
									StartPos:  5507,
									EndPos:    5509,
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
						StartPos:  5513,
						EndPos:    5522,
					},
				},
				Expr: &ast.ExprBinarySmallerOrEqual{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 298,
							EndLine:   298,
							StartPos:  5513,
							EndPos:    5521,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  5513,
								EndPos:    5515,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 298,
									EndLine:   298,
									StartPos:  5513,
									EndPos:    5515,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 298,
								EndLine:   298,
								StartPos:  5519,
								EndPos:    5521,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 298,
									EndLine:   298,
									StartPos:  5519,
									EndPos:    5521,
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
						StartPos:  5525,
						EndPos:    5533,
					},
				},
				Expr: &ast.ExprBinarySmaller{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 299,
							EndLine:   299,
							StartPos:  5525,
							EndPos:    5532,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5525,
								EndPos:    5527,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 299,
									EndLine:   299,
									StartPos:  5525,
									EndPos:    5527,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 299,
								EndLine:   299,
								StartPos:  5530,
								EndPos:    5532,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 299,
									EndLine:   299,
									StartPos:  5530,
									EndPos:    5532,
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
						StartPos:  5536,
						EndPos:    5546,
					},
				},
				Expr: &ast.ExprBinarySpaceship{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 300,
							EndLine:   300,
							StartPos:  5536,
							EndPos:    5545,
						},
					},
					Left: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5536,
								EndPos:    5538,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 300,
									EndLine:   300,
									StartPos:  5536,
									EndPos:    5538,
								},
							},
							Value: []byte("a"),
						},
					},
					Right: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 300,
								EndLine:   300,
								StartPos:  5543,
								EndPos:    5545,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 300,
									EndLine:   300,
									StartPos:  5543,
									EndPos:    5545,
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
						StartPos:  5550,
						EndPos:    5559,
					},
				},
				Expr: &ast.ExprAssignReference{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 302,
							EndLine:   302,
							StartPos:  5550,
							EndPos:    5558,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 302,
								EndLine:   302,
								StartPos:  5550,
								EndPos:    5552,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 302,
									EndLine:   302,
									StartPos:  5550,
									EndPos:    5552,
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
								StartPos:  5556,
								EndPos:    5558,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 302,
									EndLine:   302,
									StartPos:  5556,
									EndPos:    5558,
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
						StartPos:  5562,
						EndPos:    5570,
					},
				},
				Expr: &ast.ExprAssign{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 303,
							EndLine:   303,
							StartPos:  5562,
							EndPos:    5569,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 303,
								EndLine:   303,
								StartPos:  5562,
								EndPos:    5564,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 303,
									EndLine:   303,
									StartPos:  5562,
									EndPos:    5564,
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
								StartPos:  5567,
								EndPos:    5569,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 303,
									EndLine:   303,
									StartPos:  5567,
									EndPos:    5569,
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
						StartPos:  5573,
						EndPos:    5582,
					},
				},
				Expr: &ast.ExprAssignBitwiseAnd{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 304,
							EndLine:   304,
							StartPos:  5573,
							EndPos:    5581,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 304,
								EndLine:   304,
								StartPos:  5573,
								EndPos:    5575,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 304,
									EndLine:   304,
									StartPos:  5573,
									EndPos:    5575,
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
								StartPos:  5579,
								EndPos:    5581,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 304,
									EndLine:   304,
									StartPos:  5579,
									EndPos:    5581,
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
						StartPos:  5585,
						EndPos:    5594,
					},
				},
				Expr: &ast.ExprAssignBitwiseOr{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 305,
							EndLine:   305,
							StartPos:  5585,
							EndPos:    5593,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 305,
								EndLine:   305,
								StartPos:  5585,
								EndPos:    5587,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 305,
									EndLine:   305,
									StartPos:  5585,
									EndPos:    5587,
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
								StartPos:  5591,
								EndPos:    5593,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 305,
									EndLine:   305,
									StartPos:  5591,
									EndPos:    5593,
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
						StartPos:  5597,
						EndPos:    5606,
					},
				},
				Expr: &ast.ExprAssignBitwiseXor{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 306,
							EndLine:   306,
							StartPos:  5597,
							EndPos:    5605,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 306,
								EndLine:   306,
								StartPos:  5597,
								EndPos:    5599,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 306,
									EndLine:   306,
									StartPos:  5597,
									EndPos:    5599,
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
								StartPos:  5603,
								EndPos:    5605,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 306,
									EndLine:   306,
									StartPos:  5603,
									EndPos:    5605,
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
						StartPos:  5609,
						EndPos:    5618,
					},
				},
				Expr: &ast.ExprAssignConcat{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 307,
							EndLine:   307,
							StartPos:  5609,
							EndPos:    5617,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 307,
								EndLine:   307,
								StartPos:  5609,
								EndPos:    5611,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 307,
									EndLine:   307,
									StartPos:  5609,
									EndPos:    5611,
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
								StartPos:  5615,
								EndPos:    5617,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 307,
									EndLine:   307,
									StartPos:  5615,
									EndPos:    5617,
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
						StartPos:  5621,
						EndPos:    5630,
					},
				},
				Expr: &ast.ExprAssignDiv{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 308,
							EndLine:   308,
							StartPos:  5621,
							EndPos:    5629,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 308,
								EndLine:   308,
								StartPos:  5621,
								EndPos:    5623,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 308,
									EndLine:   308,
									StartPos:  5621,
									EndPos:    5623,
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
								StartPos:  5627,
								EndPos:    5629,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 308,
									EndLine:   308,
									StartPos:  5627,
									EndPos:    5629,
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
						StartPos:  5633,
						EndPos:    5642,
					},
				},
				Expr: &ast.ExprAssignMinus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 309,
							EndLine:   309,
							StartPos:  5633,
							EndPos:    5641,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 309,
								EndLine:   309,
								StartPos:  5633,
								EndPos:    5635,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 309,
									EndLine:   309,
									StartPos:  5633,
									EndPos:    5635,
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
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 310,
						EndLine:   310,
						StartPos:  5645,
						EndPos:    5654,
					},
				},
				Expr: &ast.ExprAssignMod{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 310,
							EndLine:   310,
							StartPos:  5645,
							EndPos:    5653,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 310,
								EndLine:   310,
								StartPos:  5645,
								EndPos:    5647,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 310,
									EndLine:   310,
									StartPos:  5645,
									EndPos:    5647,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 310,
								EndLine:   310,
								StartPos:  5651,
								EndPos:    5653,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 310,
									EndLine:   310,
									StartPos:  5651,
									EndPos:    5653,
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
						StartLine: 311,
						EndLine:   311,
						StartPos:  5657,
						EndPos:    5666,
					},
				},
				Expr: &ast.ExprAssignMul{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 311,
							EndLine:   311,
							StartPos:  5657,
							EndPos:    5665,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 311,
								EndLine:   311,
								StartPos:  5657,
								EndPos:    5659,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 311,
									EndLine:   311,
									StartPos:  5657,
									EndPos:    5659,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 311,
								EndLine:   311,
								StartPos:  5663,
								EndPos:    5665,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 311,
									EndLine:   311,
									StartPos:  5663,
									EndPos:    5665,
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
						StartPos:  5669,
						EndPos:    5678,
					},
				},
				Expr: &ast.ExprAssignPlus{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 312,
							EndLine:   312,
							StartPos:  5669,
							EndPos:    5677,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 312,
								EndLine:   312,
								StartPos:  5669,
								EndPos:    5671,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 312,
									EndLine:   312,
									StartPos:  5669,
									EndPos:    5671,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 312,
								EndLine:   312,
								StartPos:  5675,
								EndPos:    5677,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 312,
									EndLine:   312,
									StartPos:  5675,
									EndPos:    5677,
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
						StartLine: 313,
						EndLine:   313,
						StartPos:  5681,
						EndPos:    5691,
					},
				},
				Expr: &ast.ExprAssignPow{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 313,
							EndLine:   313,
							StartPos:  5681,
							EndPos:    5690,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5681,
								EndPos:    5683,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 313,
									EndLine:   313,
									StartPos:  5681,
									EndPos:    5683,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 313,
								EndLine:   313,
								StartPos:  5688,
								EndPos:    5690,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 313,
									EndLine:   313,
									StartPos:  5688,
									EndPos:    5690,
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
						StartLine: 314,
						EndLine:   314,
						StartPos:  5694,
						EndPos:    5704,
					},
				},
				Expr: &ast.ExprAssignShiftLeft{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 314,
							EndLine:   314,
							StartPos:  5694,
							EndPos:    5703,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 314,
								EndLine:   314,
								StartPos:  5694,
								EndPos:    5696,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 314,
									EndLine:   314,
									StartPos:  5694,
									EndPos:    5696,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 314,
								EndLine:   314,
								StartPos:  5701,
								EndPos:    5703,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 314,
									EndLine:   314,
									StartPos:  5701,
									EndPos:    5703,
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
						StartLine: 315,
						EndLine:   315,
						StartPos:  5707,
						EndPos:    5717,
					},
				},
				Expr: &ast.ExprAssignShiftRight{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 315,
							EndLine:   315,
							StartPos:  5707,
							EndPos:    5716,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5707,
								EndPos:    5709,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 315,
									EndLine:   315,
									StartPos:  5707,
									EndPos:    5709,
								},
							},
							Value: []byte("a"),
						},
					},
					Expr: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 315,
								EndLine:   315,
								StartPos:  5714,
								EndPos:    5716,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 315,
									EndLine:   315,
									StartPos:  5714,
									EndPos:    5716,
								},
							},
							Value: []byte("b"),
						},
					},
				},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 317,
						EndLine:   317,
						StartPos:  5721,
						EndPos:    5760,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 317,
							EndLine:   317,
							StartPos:  5727,
							EndPos:    5730,
						},
					},
					Value: []byte("foo"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 317,
								EndLine:   317,
								StartPos:  5732,
								EndPos:    5758,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 317,
									EndLine:   317,
									StartPos:  5748,
									EndPos:    5753,
								},
							},
							Value: []byte("class"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 317,
										EndLine:   317,
										StartPos:  5732,
										EndPos:    5738,
									},
								},
								Value: []byte("public"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 317,
									EndLine:   317,
									StartPos:  5756,
									EndPos:    5758,
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
						StartLine: 318,
						EndLine:   318,
						StartPos:  5763,
						EndPos:    5774,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 318,
							EndLine:   318,
							StartPos:  5763,
							EndPos:    5773,
						},
					},
					Function: &ast.NameFullyQualified{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 318,
								EndLine:   318,
								StartPos:  5763,
								EndPos:    5771,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 318,
										EndLine:   318,
										StartPos:  5764,
										EndPos:    5767,
									},
								},
								Value: []byte("foo"),
							},
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 318,
										EndLine:   318,
										StartPos:  5768,
										EndPos:    5771,
									},
								},
								Value: []byte("bar"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 318,
								EndLine:   318,
								StartPos:  5771,
								EndPos:    5773,
							},
						},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 320,
						EndLine:   326,
						StartPos:  5778,
						EndPos:    5905,
					},
				},
				ReturnsRef: false,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 320,
							EndLine:   320,
							StartPos:  5787,
							EndPos:    5790,
						},
					},
					Value: []byte("foo"),
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 320,
								EndLine:   320,
								StartPos:  5791,
								EndPos:    5794,
							},
						},
						ByRef:    true,
						Variadic: false,
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 320,
									EndLine:   320,
									StartPos:  5792,
									EndPos:    5794,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 320,
										EndLine:   320,
										StartPos:  5792,
										EndPos:    5794,
									},
								},
								Value: []byte("a"),
							},
						},
					},
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 320,
								EndLine:   320,
								StartPos:  5796,
								EndPos:    5801,
							},
						},
						ByRef:    false,
						Variadic: true,
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 320,
									EndLine:   320,
									StartPos:  5799,
									EndPos:    5801,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 320,
										EndLine:   320,
										StartPos:  5799,
										EndPos:    5801,
									},
								},
								Value: []byte("b"),
							},
						},
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtFunction{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 322,
								EndLine:   322,
								StartPos:  5830,
								EndPos:    5847,
							},
						},
						ReturnsRef: false,
						FunctionName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 322,
									EndLine:   322,
									StartPos:  5839,
									EndPos:    5842,
								},
							},
							Value: []byte("bar"),
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtClass{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 323,
								EndLine:   323,
								StartPos:  5851,
								EndPos:    5863,
							},
						},
						ClassName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 323,
									EndLine:   323,
									StartPos:  5857,
									EndPos:    5860,
								},
							},
							Value: []byte("Baz"),
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtTrait{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 324,
								EndLine:   324,
								StartPos:  5867,
								EndPos:    5879,
							},
						},
						TraitName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 324,
									EndLine:   324,
									StartPos:  5873,
									EndPos:    5877,
								},
							},
							Value: []byte("Quux"),
						},
						Stmts: []ast.Vertex{},
					},
					&ast.StmtInterface{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 325,
								EndLine:   325,
								StartPos:  5883,
								EndPos:    5901,
							},
						},
						InterfaceName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 325,
									EndLine:   325,
									StartPos:  5893,
									EndPos:    5898,
								},
							},
							Value: []byte("Quuux"),
						},
						Stmts: []ast.Vertex{},
					},
				},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 328,
						EndLine:   328,
						StartPos:  5911,
						EndPos:    5954,
					},
				},
				ReturnsRef: false,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 328,
							EndLine:   328,
							StartPos:  5920,
							EndPos:    5923,
						},
					},
					Value: []byte("foo"),
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5924,
								EndPos:    5931,
							},
						},
						ByRef:    true,
						Variadic: false,
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5925,
									EndPos:    5927,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 328,
										EndLine:   328,
										StartPos:  5925,
										EndPos:    5927,
									},
								},
								Value: []byte("a"),
							},
						},
						DefaultValue: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5930,
									EndPos:    5931,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5933,
								EndPos:    5942,
							},
						},
						ByRef:    false,
						Variadic: true,
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5936,
									EndPos:    5938,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 328,
										EndLine:   328,
										StartPos:  5936,
										EndPos:    5938,
									},
								},
								Value: []byte("b"),
							},
						},
						DefaultValue: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5941,
									EndPos:    5942,
								},
							},
							Value: []byte("1"),
						},
					},
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 328,
								EndLine:   328,
								StartPos:  5944,
								EndPos:    5950,
							},
						},
						ByRef:    false,
						Variadic: false,
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5944,
									EndPos:    5946,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 328,
										EndLine:   328,
										StartPos:  5944,
										EndPos:    5946,
									},
								},
								Value: []byte("c"),
							},
						},
						DefaultValue: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 328,
									EndLine:   328,
									StartPos:  5949,
									EndPos:    5950,
								},
							},
							Value: []byte("1"),
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtFunction{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 329,
						EndLine:   329,
						StartPos:  5957,
						EndPos:    5995,
					},
				},
				ReturnsRef: false,
				FunctionName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 329,
							EndLine:   329,
							StartPos:  5966,
							EndPos:    5969,
						},
					},
					Value: []byte("foo"),
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5970,
								EndPos:    5978,
							},
						},
						ByRef:    false,
						Variadic: false,
						Type: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5970,
									EndPos:    5975,
								},
							},
							Value: []byte("array"),
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5976,
									EndPos:    5978,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 329,
										EndLine:   329,
										StartPos:  5976,
										EndPos:    5978,
									},
								},
								Value: []byte("a"),
							},
						},
					},
					&ast.Parameter{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 329,
								EndLine:   329,
								StartPos:  5980,
								EndPos:    5991,
							},
						},
						Variadic: false,
						ByRef:    false,
						Type: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5980,
									EndPos:    5988,
								},
							},
							Value: []byte("callable"),
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 329,
									EndLine:   329,
									StartPos:  5989,
									EndPos:    5991,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 329,
										EndLine:   329,
										StartPos:  5989,
										EndPos:    5991,
									},
								},
								Value: []byte("b"),
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
			},
			&ast.StmtClass{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 330,
						EndLine:   330,
						StartPos:  5998,
						EndPos:    6100,
					},
				},
				ClassName: &ast.Identifier{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 330,
							EndLine:   330,
							StartPos:  6019,
							EndPos:    6022,
						},
					},
					Value: []byte("foo"),
				},
				Modifiers: []ast.Vertex{
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  5998,
								EndPos:    6006,
							},
						},
						Value: []byte("abstract"),
					},
					&ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  6007,
								EndPos:    6012,
							},
						},
						Value: []byte("final"),
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  6025,
								EndPos:    6066,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  6060,
									EndPos:    6063,
								},
							},
							Value: []byte("bar"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 330,
										EndLine:   330,
										StartPos:  6025,
										EndPos:    6033,
									},
								},
								Value: []byte("abstract"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 330,
										EndLine:   330,
										StartPos:  6034,
										EndPos:    6043,
									},
								},
								Value: []byte("protected"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 330,
										EndLine:   330,
										StartPos:  6044,
										EndPos:    6050,
									},
								},
								Value: []byte("static"),
							},
						},
						Stmt: &ast.StmtNop{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  6065,
									EndPos:    6066,
								},
							},
						},
					},
					&ast.StmtClassMethod{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 330,
								EndLine:   330,
								StartPos:  6067,
								EndPos:    6098,
							},
						},
						ReturnsRef: false,
						MethodName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  6090,
									EndPos:    6093,
								},
							},
							Value: []byte("baz"),
						},
						Modifiers: []ast.Vertex{
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 330,
										EndLine:   330,
										StartPos:  6067,
										EndPos:    6072,
									},
								},
								Value: []byte("final"),
							},
							&ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 330,
										EndLine:   330,
										StartPos:  6073,
										EndPos:    6080,
									},
								},
								Value: []byte("private"),
							},
						},
						Stmt: &ast.StmtStmtList{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 330,
									EndLine:   330,
									StartPos:  6096,
									EndPos:    6098,
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
						StartLine: 332,
						EndLine:   332,
						StartPos:  6105,
						EndPos:    6119,
					},
				},
				Expr: &ast.ExprPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 332,
							EndLine:   332,
							StartPos:  6105,
							EndPos:    6118,
						},
					},
					Var: &ast.ExprNew{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 332,
								EndLine:   332,
								StartPos:  6105,
								EndPos:    6112,
							},
						},
						Class: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 332,
									EndLine:   332,
									StartPos:  6109,
									EndPos:    6112,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 332,
											EndLine:   332,
											StartPos:  6109,
											EndPos:    6112,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
					Property: &ast.Identifier{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 332,
								EndLine:   332,
								StartPos:  6115,
								EndPos:    6118,
							},
						},
						Value: []byte("bar"),
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 333,
						EndLine:   333,
						StartPos:  6123,
						EndPos:    6134,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 333,
							EndLine:   333,
							StartPos:  6123,
							EndPos:    6133,
						},
					},
					Function: &ast.ExprNew{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 333,
								EndLine:   333,
								StartPos:  6123,
								EndPos:    6130,
							},
						},
						Class: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 333,
									EndLine:   333,
									StartPos:  6127,
									EndPos:    6130,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 333,
											EndLine:   333,
											StartPos:  6127,
											EndPos:    6130,
										},
									},
									Value: []byte("Foo"),
								},
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 333,
								EndLine:   333,
								StartPos:  6131,
								EndPos:    6133,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 334,
						EndLine:   334,
						StartPos:  6137,
						EndPos:    6149,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 334,
							EndLine:   334,
							StartPos:  6137,
							EndPos:    6148,
						},
					},
					Function: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 334,
								EndLine:   334,
								StartPos:  6137,
								EndPos:    6146,
							},
						},
						Var: &ast.ExprShortArray{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 334,
									EndLine:   334,
									StartPos:  6137,
									EndPos:    6143,
								},
							},
							Items: []ast.Vertex{
								&ast.ExprArrayItem{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 334,
											EndLine:   334,
											StartPos:  6138,
											EndPos:    6142,
										},
									},
									Val: &ast.ExprVariable{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 334,
												EndLine:   334,
												StartPos:  6138,
												EndPos:    6142,
											},
										},
										VarName: &ast.Identifier{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 334,
													EndLine:   334,
													StartPos:  6138,
													EndPos:    6142,
												},
											},
											Value: []byte("foo"),
										},
									},
								},
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 334,
									EndLine:   334,
									StartPos:  6144,
									EndPos:    6145,
								},
							},
							Value: []byte("0"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 334,
								EndLine:   334,
								StartPos:  6146,
								EndPos:    6148,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 335,
						EndLine:   335,
						StartPos:  6152,
						EndPos:    6161,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 335,
							EndLine:   335,
							StartPos:  6152,
							EndPos:    6160,
						},
					},
					Function: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 335,
								EndLine:   335,
								StartPos:  6152,
								EndPos:    6158,
							},
						},
						Var: &ast.ExprConstFetch{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 335,
									EndLine:   335,
									StartPos:  6152,
									EndPos:    6155,
								},
							},
							Const: &ast.NameName{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 335,
										EndLine:   335,
										StartPos:  6152,
										EndPos:    6155,
									},
								},
								Parts: []ast.Vertex{
									&ast.NameNamePart{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 335,
												EndLine:   335,
												StartPos:  6152,
												EndPos:    6155,
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
									StartLine: 335,
									EndLine:   335,
									StartPos:  6156,
									EndPos:    6157,
								},
							},
							Value: []byte("1"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 335,
								EndLine:   335,
								StartPos:  6158,
								EndPos:    6160,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 336,
						EndLine:   336,
						StartPos:  6164,
						EndPos:    6172,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 336,
							EndLine:   336,
							StartPos:  6164,
							EndPos:    6171,
						},
					},
					Function: &ast.ScalarString{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 336,
								EndLine:   336,
								StartPos:  6164,
								EndPos:    6169,
							},
						},
						Value: []byte("\"foo\""),
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 336,
								EndLine:   336,
								StartPos:  6169,
								EndPos:    6171,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 337,
						EndLine:   337,
						StartPos:  6175,
						EndPos:    6187,
					},
				},
				Expr: &ast.ExprFunctionCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 337,
							EndLine:   337,
							StartPos:  6175,
							EndPos:    6186,
						},
					},
					Function: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 337,
								EndLine:   337,
								StartPos:  6175,
								EndPos:    6184,
							},
						},
						Var: &ast.ExprShortArray{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 337,
									EndLine:   337,
									StartPos:  6175,
									EndPos:    6178,
								},
							},
							Items: []ast.Vertex{
								&ast.ExprArrayItem{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 337,
											EndLine:   337,
											StartPos:  6176,
											EndPos:    6177,
										},
									},
									Val: &ast.ScalarLnumber{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 337,
												EndLine:   337,
												StartPos:  6176,
												EndPos:    6177,
											},
										},
										Value: []byte("1"),
									},
								},
							},
						},
						Dim: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 337,
									EndLine:   337,
									StartPos:  6179,
									EndPos:    6183,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 337,
										EndLine:   337,
										StartPos:  6179,
										EndPos:    6183,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 337,
								EndLine:   337,
								StartPos:  6184,
								EndPos:    6186,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 338,
						EndLine:   338,
						StartPos:  6190,
						EndPos:    6199,
					},
				},
				Expr: &ast.ExprVariable{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 338,
							EndLine:   338,
							StartPos:  6190,
							EndPos:    6198,
						},
					},
					VarName: &ast.ExprFunctionCall{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 338,
								EndLine:   338,
								StartPos:  6192,
								EndPos:    6197,
							},
						},
						Function: &ast.NameName{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 338,
									EndLine:   338,
									StartPos:  6192,
									EndPos:    6195,
								},
							},
							Parts: []ast.Vertex{
								&ast.NameNamePart{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 338,
											EndLine:   338,
											StartPos:  6192,
											EndPos:    6195,
										},
									},
									Value: []byte("foo"),
								},
							},
						},
						ArgumentList: &ast.ArgumentList{
							Node: ast.Node{
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
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 340,
						EndLine:   340,
						StartPos:  6203,
						EndPos:    6215,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 340,
							EndLine:   340,
							StartPos:  6203,
							EndPos:    6214,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 340,
								EndLine:   340,
								StartPos:  6203,
								EndPos:    6206,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 340,
										EndLine:   340,
										StartPos:  6203,
										EndPos:    6206,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 340,
								EndLine:   340,
								StartPos:  6208,
								EndPos:    6212,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 340,
									EndLine:   340,
									StartPos:  6208,
									EndPos:    6212,
								},
							},
							Value: []byte("bar"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 340,
								EndLine:   340,
								StartPos:  6212,
								EndPos:    6214,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 341,
						EndLine:   341,
						StartPos:  6218,
						EndPos:    6235,
					},
				},
				Expr: &ast.ExprStaticCall{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 341,
							EndLine:   341,
							StartPos:  6218,
							EndPos:    6234,
						},
					},
					Class: &ast.NameName{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  6218,
								EndPos:    6221,
							},
						},
						Parts: []ast.Vertex{
							&ast.NameNamePart{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 341,
										EndLine:   341,
										StartPos:  6218,
										EndPos:    6221,
									},
								},
								Value: []byte("Foo"),
							},
						},
					},
					Call: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  6224,
								EndPos:    6231,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  6224,
									EndPos:    6228,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 341,
										EndLine:   341,
										StartPos:  6224,
										EndPos:    6228,
									},
								},
								Value: []byte("bar"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 341,
									EndLine:   341,
									StartPos:  6229,
									EndPos:    6230,
								},
							},
							Value: []byte("0"),
						},
					},
					ArgumentList: &ast.ArgumentList{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 341,
								EndLine:   341,
								StartPos:  6232,
								EndPos:    6234,
							},
						},
					},
				},
			},
			&ast.StmtExpression{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 343,
						EndLine:   343,
						StartPos:  6241,
						EndPos:    6252,
					},
				},
				Expr: &ast.ExprPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 343,
							EndLine:   343,
							StartPos:  6241,
							EndPos:    6251,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 343,
								EndLine:   343,
								StartPos:  6241,
								EndPos:    6245,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 343,
									EndLine:   343,
									StartPos:  6241,
									EndPos:    6245,
								},
							},
							Value: []byte("foo"),
						},
					},
					Property: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 343,
								EndLine:   343,
								StartPos:  6247,
								EndPos:    6251,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 343,
									EndLine:   343,
									StartPos:  6247,
									EndPos:    6251,
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
						StartLine: 344,
						EndLine:   344,
						StartPos:  6255,
						EndPos:    6271,
					},
				},
				Expr: &ast.ExprPropertyFetch{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 344,
							EndLine:   344,
							StartPos:  6255,
							EndPos:    6269,
						},
					},
					Var: &ast.ExprVariable{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 344,
								EndLine:   344,
								StartPos:  6255,
								EndPos:    6259,
							},
						},
						VarName: &ast.Identifier{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 344,
									EndLine:   344,
									StartPos:  6255,
									EndPos:    6259,
								},
							},
							Value: []byte("foo"),
						},
					},
					Property: &ast.ExprArrayDimFetch{
						Node: ast.Node{
							Position: &position.Position{
								StartLine: 344,
								EndLine:   344,
								StartPos:  6262,
								EndPos:    6269,
							},
						},
						Var: &ast.ExprVariable{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 344,
									EndLine:   344,
									StartPos:  6262,
									EndPos:    6266,
								},
							},
							VarName: &ast.Identifier{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 344,
										EndLine:   344,
										StartPos:  6262,
										EndPos:    6266,
									},
								},
								Value: []byte("bar"),
							},
						},
						Dim: &ast.ScalarLnumber{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 344,
									EndLine:   344,
									StartPos:  6267,
									EndPos:    6268,
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
						StartLine: 346,
						EndLine:   346,
						StartPos:  6275,
						EndPos:    6297,
					},
				},
				Expr: &ast.ExprShortArray{
					Node: ast.Node{
						Position: &position.Position{
							StartLine: 346,
							EndLine:   346,
							StartPos:  6275,
							EndPos:    6296,
						},
					},
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  6276,
									EndPos:    6282,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  6276,
										EndPos:    6277,
									},
								},
								Value: []byte("1"),
							},
							Val: &ast.ExprReference{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  6279,
										EndPos:    6282,
									},
								},
								Var: &ast.ExprVariable{
									Node: ast.Node{
										Position: &position.Position{
											StartLine: 346,
											EndLine:   346,
											StartPos:  6280,
											EndPos:    6282,
										},
									},
									VarName: &ast.Identifier{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 346,
												EndLine:   346,
												StartPos:  6280,
												EndPos:    6282,
											},
										},
										Value: []byte("a"),
									},
								},
							},
						},
						&ast.ExprArrayItem{
							Node: ast.Node{
								Position: &position.Position{
									StartLine: 346,
									EndLine:   346,
									StartPos:  6284,
									EndPos:    6295,
								},
							},
							Key: &ast.ScalarLnumber{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  6284,
										EndPos:    6285,
									},
								},
								Value: []byte("2"),
							},
							Val: &ast.ExprList{
								Node: ast.Node{
									Position: &position.Position{
										StartLine: 346,
										EndLine:   346,
										StartPos:  6287,
										EndPos:    6295,
									},
								},
								Items: []ast.Vertex{
									&ast.ExprArrayItem{
										Node: ast.Node{
											Position: &position.Position{
												StartLine: 346,
												EndLine:   346,
												StartPos:  6292,
												EndPos:    6294,
											},
										},
										Val: &ast.ExprVariable{
											Node: ast.Node{
												Position: &position.Position{
													StartLine: 346,
													EndLine:   346,
													StartPos:  6292,
													EndPos:    6294,
												},
											},
											VarName: &ast.Identifier{
												Node: ast.Node{
													Position: &position.Position{
														StartLine: 346,
														EndLine:   346,
														StartPos:  6292,
														EndPos:    6294,
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
			},
			&ast.StmtHaltCompiler{
				Node: ast.Node{
					Position: &position.Position{
						StartLine: 348,
						EndLine:   348,
						StartPos:  6301,
						EndPos:    6319,
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
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

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
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

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPhp7ControlCharsErrors(t *testing.T) {
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

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetErrors()
	assert.DeepEqual(t, expected, actual)
}
