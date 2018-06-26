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

	expectedParams := []node.Node{
		&node.Parameter{
			ByRef:        false,
			Variadic:     false,
			VariableType: &name.Name{Parts: []node.Node{&name.NamePart{Value: "bar"}}},
			Variable:     &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
			DefaultValue: &expr.ConstFetch{Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "null"}}}},
		},
		&node.Parameter{
			ByRef:        true,
			Variadic:     true,
			VariableType: &name.Name{Parts: []node.Node{&name.NamePart{Value: "baz"}}},
			Variable:     &expr.Variable{VarName: &node.Identifier{Value: "baz"}},
		},
	}

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.Name{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.MethodCall{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					Method:   &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class: &name.Name{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
					Call:  &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					Call:  &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.New{
					Class: &name.Name{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{Variadic: false, Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
							&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Function{
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName:  &node.Identifier{Value: "foo"},
				Params:        expectedParams,
				Stmts:         []node.Node{},
			},
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						MethodName: &node.Identifier{Value: "foo"},
						Modifiers:  []node.Node{&node.Identifier{Value: "public"}},
						Params:     expectedParams,
						Stmt: &stmt.StmtList{
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.Closure{
					Params: expectedParams,
					Stmts:  []node.Node{},
				},
			},
			&stmt.Expression{
				Expr: &expr.Closure{
					Static: true,
					Params: expectedParams,
					Stmts:  []node.Node{},
				},
			},
			&stmt.Expression{
				Expr: &scalar.Lnumber{Value: "1234567890123456789"},
			},
			&stmt.Expression{
				Expr: &scalar.Dnumber{Value: "12345678901234567890"},
			},
			&stmt.Expression{
				Expr: &scalar.Dnumber{Value: "0."},
			},
			&stmt.Expression{
				Expr: &scalar.Lnumber{Value: "0b0111111111111111111111111111111111111111111111111111111111111111"},
			},
			&stmt.Expression{
				Expr: &scalar.Dnumber{Value: "0b1111111111111111111111111111111111111111111111111111111111111111"},
			},
			&stmt.Expression{
				Expr: &scalar.Lnumber{Value: "0x007111111111111111"},
			},
			&stmt.Expression{
				Expr: &scalar.Dnumber{Value: "0x8111111111111111"},
			},

			&stmt.Expression{
				Expr: &scalar.MagicConstant{Value: "__CLASS__"},
			},
			&stmt.Expression{
				Expr: &scalar.MagicConstant{Value: "__DIR__"},
			},
			&stmt.Expression{
				Expr: &scalar.MagicConstant{Value: "__FILE__"},
			},
			&stmt.Expression{
				Expr: &scalar.MagicConstant{Value: "__FUNCTION__"},
			},
			&stmt.Expression{
				Expr: &scalar.MagicConstant{Value: "__LINE__"},
			},
			&stmt.Expression{
				Expr: &scalar.MagicConstant{Value: "__NAMESPACE__"},
			},
			&stmt.Expression{
				Expr: &scalar.MagicConstant{Value: "__METHOD__"},
			},
			&stmt.Expression{
				Expr: &scalar.MagicConstant{Value: "__TRAIT__"},
			},

			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.Variable{VarName: &node.Identifier{Value: "var"}},
					},
				},
			},
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.ArrayDimFetch{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
							Dim:      &scalar.Lnumber{Value: "1"},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.ArrayDimFetch{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
							Dim:      &scalar.String{Value: "1234567890123456789012345678901234567890"},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.ArrayDimFetch{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
							Dim:      &scalar.String{Value: "bar"},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.ArrayDimFetch{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "var"}},
							Dim:      &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&expr.Variable{VarName: &node.Identifier{Value: "foo"}},
						&scalar.EncapsedStringPart{Value: " "},
						&expr.Variable{VarName: &node.Identifier{Value: "bar"}},
					},
				},
			},
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.PropertyFetch{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
							Property: &node.Identifier{Value: "bar"},
						},
						&scalar.EncapsedStringPart{Value: "()"},
					},
				},
			},
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					},
				},
			},
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.ArrayDimFetch{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
							Dim:      &scalar.Lnumber{Value: "0"},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.MethodCall{
							Variable:     &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
							Method:       &node.Identifier{Value: "bar"},
							ArgumentList: &node.ArgumentList{},
						},
					},
				},
			},

			&stmt.AltIf{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
			},
			&stmt.AltIf{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				ElseIf: []node.Node{
					&stmt.AltElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
				},
			},
			&stmt.AltIf{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				Else: &stmt.AltElse{
					Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				},
			},
			&stmt.AltIf{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				ElseIf: []node.Node{
					&stmt.AltElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
					&stmt.AltElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "c"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
				},
				Else: &stmt.AltElse{
					Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				},
			},
			&stmt.While{
				Cond: &scalar.Lnumber{Value: "1"},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Break{},
					},
				},
			},
			&stmt.While{
				Cond: &scalar.Lnumber{Value: "1"},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Break{
							Expr: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.AltWhile{
				Cond: &scalar.Lnumber{Value: "1"},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Break{
							Expr: &scalar.Lnumber{Value: "3"},
						},
					},
				},
			},
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.ClassConstList{
						Consts: []node.Node{
							&stmt.Constant{
								PhpDocComment: "",
								ConstantName:  &node.Identifier{Value: "FOO"},
								Expr:          &scalar.Lnumber{Value: "1"},
							},
							&stmt.Constant{
								PhpDocComment: "",
								ConstantName:  &node.Identifier{Value: "BAR"},
								Expr:          &scalar.Lnumber{Value: "2"},
							},
						},
					},
				},
			},
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						PhpDocComment: "",
						MethodName:    &node.Identifier{Value: "bar"},
						Stmt: &stmt.StmtList{
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						PhpDocComment: "",
						ReturnsRef:    true,
						MethodName:    &node.Identifier{Value: "bar"},
						Modifiers: []node.Node{
							&node.Identifier{Value: "public"},
							&node.Identifier{Value: "static"},
						},
						Stmt: &stmt.StmtList{
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						PhpDocComment: "",
						ReturnsRef:    false,
						MethodName:    &node.Identifier{Value: "bar"},
						Modifiers: []node.Node{
							&node.Identifier{Value: "final"},
							&node.Identifier{Value: "private"},
						},
						Stmt: &stmt.StmtList{
							Stmts: []node.Node{},
						},
					},
					&stmt.ClassMethod{
						PhpDocComment: "",
						ReturnsRef:    false,
						MethodName:    &node.Identifier{Value: "baz"},
						Modifiers: []node.Node{
							&node.Identifier{Value: "protected"},
						},
						Stmt: &stmt.StmtList{
							Stmts: []node.Node{},
						},
					},
				},
			},
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Modifiers: []node.Node{
					&node.Identifier{Value: "abstract"},
				},
				Stmts: []node.Node{
					&stmt.ClassMethod{
						PhpDocComment: "",
						ReturnsRef:    false,
						MethodName:    &node.Identifier{Value: "bar"},
						Modifiers: []node.Node{
							&node.Identifier{Value: "abstract"},
							&node.Identifier{Value: "public"},
						},
						Stmt: &stmt.Nop{},
					},
				},
			},
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Modifiers: []node.Node{
					&node.Identifier{Value: "final"},
				},
				Extends: &stmt.ClassExtends{
					ClassName: &name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "bar"},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Modifiers: []node.Node{
					&node.Identifier{Value: "final"},
				},
				Implements: &stmt.ClassImplements{
					InterfaceNames: []node.Node{
						&name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "bar"},
							},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Modifiers: []node.Node{
					&node.Identifier{Value: "final"},
				},
				Implements: &stmt.ClassImplements{
					InterfaceNames: []node.Node{
						&name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "bar"},
							},
						},
						&name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "baz"},
							},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.ConstList{
				Consts: []node.Node{
					&stmt.Constant{
						PhpDocComment: "",
						ConstantName:  &node.Identifier{Value: "FOO"},
						Expr:          &scalar.Lnumber{Value: "1"},
					},
					&stmt.Constant{
						PhpDocComment: "",
						ConstantName:  &node.Identifier{Value: "BAR"},
						Expr:          &scalar.Lnumber{Value: "2"},
					},
				},
			},
			&stmt.While{
				Cond: &scalar.Lnumber{Value: "1"},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Continue{Expr: nil},
					},
				},
			},
			&stmt.While{
				Cond: &scalar.Lnumber{Value: "1"},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Continue{
							Expr: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.While{
				Cond: &scalar.Lnumber{Value: "1"},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Continue{
							Expr: &scalar.Lnumber{Value: "3"},
						},
					},
				},
			},
			&stmt.Declare{
				Consts: []node.Node{
					&stmt.Constant{
						PhpDocComment: "",
						ConstantName:  &node.Identifier{Value: "ticks"},
						Expr:          &scalar.Lnumber{Value: "1"},
					},
				},
				Stmt: &stmt.Nop{},
			},
			&stmt.Declare{
				Consts: []node.Node{
					&stmt.Constant{
						PhpDocComment: "",
						ConstantName:  &node.Identifier{Value: "ticks"},
						Expr:          &scalar.Lnumber{Value: "1"},
					},
					&stmt.Constant{
						PhpDocComment: "",
						ConstantName:  &node.Identifier{Value: "strict_types"},
						Expr:          &scalar.Lnumber{Value: "1"},
					},
				},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{},
				},
			},
			&stmt.Declare{
				Consts: []node.Node{
					&stmt.Constant{
						PhpDocComment: "",
						ConstantName:  &node.Identifier{Value: "ticks"},
						Expr:          &scalar.Lnumber{Value: "1"},
					},
				},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{},
				},
			},
			&stmt.Do{
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{},
				},
				Cond: &scalar.Lnumber{Value: "1"},
			},
			&stmt.Echo{
				Exprs: []node.Node{
					&expr.Variable{
						VarName: &node.Identifier{Value: "a"},
					},
					&scalar.Lnumber{Value: "1"},
				},
			},
			&stmt.Echo{
				Exprs: []node.Node{
					&expr.Variable{
						VarName: &node.Identifier{Value: "a"},
					},
				},
			},
			&stmt.For{
				Init: []node.Node{
					&assign.Assign{
						Variable:   &expr.Variable{VarName: &node.Identifier{Value: "i"}},
						Expression: &scalar.Lnumber{Value: "0"},
					},
				},
				Cond: []node.Node{
					&binary.Smaller{
						Left:  &expr.Variable{VarName: &node.Identifier{Value: "i"}},
						Right: &scalar.Lnumber{Value: "10"},
					},
				},
				Loop: []node.Node{
					&expr.PostInc{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "i"}},
					},
					&expr.PostInc{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "i"}},
					},
				},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
			},
			&stmt.AltFor{
				Cond: []node.Node{
					&binary.Smaller{
						Left:  &expr.Variable{VarName: &node.Identifier{Value: "i"}},
						Right: &scalar.Lnumber{Value: "10"},
					},
				},
				Loop: []node.Node{
					&expr.PostInc{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "i"}},
					},
				},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
			},
			&stmt.Foreach{
				Expr:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "v"}},
				Stmt:     &stmt.StmtList{Stmts: []node.Node{}},
			},
			&stmt.Foreach{
				Expr:     &expr.ShortArray{Items: []node.Node{}},
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "v"}},
				Stmt:     &stmt.StmtList{Stmts: []node.Node{}},
			},
			&stmt.AltForeach{
				Expr:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "v"}},
				Stmt:     &stmt.StmtList{Stmts: []node.Node{}},
			},
			&stmt.Foreach{
				Expr:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Key:      &expr.Variable{VarName: &node.Identifier{Value: "k"}},
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "v"}},
				Stmt:     &stmt.StmtList{Stmts: []node.Node{}},
			},
			&stmt.Foreach{
				Expr:     &expr.ShortArray{Items: []node.Node{}},
				Key:      &expr.Variable{VarName: &node.Identifier{Value: "k"}},
				Variable: &expr.Variable{VarName: &node.Identifier{Value: "v"}},
				Stmt:     &stmt.StmtList{Stmts: []node.Node{}},
			},
			&stmt.Foreach{
				Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Key:  &expr.Variable{VarName: &node.Identifier{Value: "k"}},
				Variable: &expr.Reference{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "v"}},
				},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
			},
			&stmt.Foreach{
				Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Key:  &expr.Variable{VarName: &node.Identifier{Value: "k"}},
				Variable: &expr.List{
					Items: []node.Node{
						&expr.ArrayItem{
							Val: &expr.Variable{VarName: &node.Identifier{Value: "v"}},
						},
					},
				},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
			},
			&stmt.Function{
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName:  &node.Identifier{Value: "foo"},
				Stmts:         []node.Node{},
			},
			&stmt.Function{
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName:  &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.HaltCompiler{},
					&stmt.Function{
						ReturnsRef:    false,
						PhpDocComment: "",
						FunctionName:  &node.Identifier{Value: "bar"},
						Stmts:         []node.Node{},
					},
					&stmt.Class{
						PhpDocComment: "",
						ClassName:     &node.Identifier{Value: "Baz"},
						Stmts:         []node.Node{},
					},
					&stmt.Return{
						Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					},
				},
			},
			&stmt.Function{
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName:  &node.Identifier{Value: "foo"},
				Params: []node.Node{
					&node.Parameter{
						ByRef:        false,
						Variadic:     false,
						VariableType: &node.Identifier{Value: "array"},
						Variable:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					},
					&node.Parameter{
						ByRef:        false,
						Variadic:     false,
						VariableType: &node.Identifier{Value: "callable"},
						Variable:     &expr.Variable{VarName: &node.Identifier{Value: "b"}},
					},
				},
				Stmts: []node.Node{
					&stmt.Return{},
				},
			},
			&stmt.Function{
				ReturnsRef:    true,
				PhpDocComment: "",
				FunctionName:  &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.Return{
						Expr: &scalar.Lnumber{Value: "1"},
					},
				},
			},
			&stmt.Function{
				ReturnsRef:    true,
				PhpDocComment: "",
				FunctionName:  &node.Identifier{Value: "foo"},
				Stmts:         []node.Node{},
			},
			&stmt.Global{
				Vars: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "a"}},
					&expr.Variable{VarName: &node.Identifier{Value: "b"}},
					&expr.Variable{VarName: &expr.Variable{VarName: &node.Identifier{Value: "c"}}},
					&expr.Variable{
						VarName: &expr.FunctionCall{
							Function: &name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "foo"},
								},
							},
							ArgumentList: &node.ArgumentList{},
						},
					},
				},
			},
			&stmt.Label{
				LabelName: &node.Identifier{Value: "a"},
			},
			&stmt.Goto{
				Label: &node.Identifier{Value: "a"},
			},
			&stmt.HaltCompiler{},
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
			},
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
				},
			},
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				Else: &stmt.Else{
					Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				},
			},
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
					&stmt.ElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "c"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
				},
				Else: &stmt.Else{
					Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				},
			},
			&stmt.If{
				Cond: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				Stmt: &stmt.StmtList{Stmts: []node.Node{}},
				ElseIf: []node.Node{
					&stmt.ElseIf{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
					},
				},
				Else: &stmt.Else{
					Stmt: &stmt.If{
						Cond: &expr.Variable{VarName: &node.Identifier{Value: "c"}},
						Stmt: &stmt.StmtList{Stmts: []node.Node{}},
						Else: &stmt.Else{
							Stmt: &stmt.StmtList{Stmts: []node.Node{}},
						},
					},
				},
			},
			&stmt.Nop{},
			&stmt.InlineHtml{Value: "<div></div> "},
			&stmt.Interface{
				PhpDocComment: "",
				InterfaceName: &node.Identifier{Value: "Foo"},
				Stmts:         []node.Node{},
			},
			&stmt.Interface{
				PhpDocComment: "",
				InterfaceName: &node.Identifier{Value: "Foo"},
				Extends: &stmt.InterfaceExtends{
					InterfaceNames: []node.Node{
						&name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Interface{
				PhpDocComment: "",
				InterfaceName: &node.Identifier{Value: "Foo"},
				Extends: &stmt.InterfaceExtends{
					InterfaceNames: []node.Node{
						&name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
						&name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Baz"},
							},
						},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Namespace{
				NamespaceName: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "Foo"},
					},
				},
			},
			&stmt.Namespace{
				NamespaceName: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "Foo"},
						&name.NamePart{Value: "Bar"},
					},
				},
				Stmts: []node.Node{},
			},
			&stmt.Namespace{
				Stmts: []node.Node{},
			},
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Modifiers: []node.Node{
							&node.Identifier{Value: "var"},
						},
						Properties: []node.Node{
							&stmt.Property{
								PhpDocComment: "",
								Variable:      &expr.Variable{VarName: &node.Identifier{Value: "a"}},
							},
						},
					},
				},
			},
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Modifiers: []node.Node{
							&node.Identifier{Value: "public"},
							&node.Identifier{Value: "static"},
						},
						Properties: []node.Node{
							&stmt.Property{
								PhpDocComment: "",
								Variable:      &expr.Variable{VarName: &node.Identifier{Value: "a"}},
							},
							&stmt.Property{
								PhpDocComment: "",
								Variable:      &expr.Variable{VarName: &node.Identifier{Value: "b"}},
								Expr:          &scalar.Lnumber{Value: "1"},
							},
						},
					},
				},
			},
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Modifiers: []node.Node{
							&node.Identifier{Value: "public"},
							&node.Identifier{Value: "static"},
						},
						Properties: []node.Node{
							&stmt.Property{
								PhpDocComment: "",
								Variable:      &expr.Variable{VarName: &node.Identifier{Value: "a"}},
								Expr:          &scalar.Lnumber{Value: "1"},
							},
							&stmt.Property{
								PhpDocComment: "",
								Variable:      &expr.Variable{VarName: &node.Identifier{Value: "b"}},
							},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					},
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						Expr:     &scalar.Lnumber{Value: "1"},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr:     &scalar.Lnumber{Value: "1"},
					},
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
					},
				},
			},
			&stmt.AltSwitch{
				Cond: &scalar.Lnumber{Value: "1"},
				CaseList: &stmt.CaseList{
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
			&stmt.AltSwitch{
				Cond: &scalar.Lnumber{Value: "1"},
				CaseList: &stmt.CaseList{
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
			&stmt.Switch{
				Cond: &scalar.Lnumber{Value: "1"},
				CaseList: &stmt.CaseList{
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
			&stmt.Switch{
				Cond: &scalar.Lnumber{Value: "1"},
				CaseList: &stmt.CaseList{
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
			&stmt.Throw{
				Expr: &expr.Variable{VarName: &node.Identifier{Value: "e"}},
			},
			&stmt.Trait{
				PhpDocComment: "",
				TraitName:     &node.Identifier{Value: "Foo"},
				Stmts:         []node.Node{},
			},
			&stmt.Class{
				PhpDocComment: "",
				ClassName:     &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Traits: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Bar"},
								},
							},
						},
					},
				},
			},
			&stmt.Class{
				PhpDocComment: "",
				ClassName:     &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Traits: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Bar"},
								},
							},
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Baz"},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{},
					},
				},
			},
			&stmt.Class{
				PhpDocComment: "",
				ClassName:     &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Traits: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Bar"},
								},
							},
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Baz"},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Ref: &stmt.TraitMethodRef{
										Method: &node.Identifier{Value: "one"},
									},
									Modifier: &node.Identifier{Value: "public"},
								},
							},
						},
					},
				},
			},
			&stmt.Class{
				PhpDocComment: "",
				ClassName:     &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Traits: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Bar"},
								},
							},
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Baz"},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Adaptations: []node.Node{
								&stmt.TraitUseAlias{
									Ref: &stmt.TraitMethodRef{
										Method: &node.Identifier{Value: "one"},
									},
									Modifier: &node.Identifier{Value: "public"},
									Alias:    &node.Identifier{Value: "two"},
								},
							},
						},
					},
				},
			},
			&stmt.Class{
				PhpDocComment: "",
				ClassName:     &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Traits: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Bar"},
								},
							},
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Baz"},
								},
							},
						},
						TraitAdaptationList: &stmt.TraitAdaptationList{
							Adaptations: []node.Node{
								&stmt.TraitUsePrecedence{
									Ref: &stmt.TraitMethodRef{
										Trait: &name.Name{
											Parts: []node.Node{
												&name.NamePart{Value: "Bar"},
											},
										},
										Method: &node.Identifier{Value: "one"},
									},
									Insteadof: []node.Node{
										&name.Name{
											Parts: []node.Node{
												&name.NamePart{Value: "Baz"},
											},
										},
										&name.Name{
											Parts: []node.Node{
												&name.NamePart{Value: "Quux"},
											},
										},
									},
								},
								&stmt.TraitUseAlias{
									Ref: &stmt.TraitMethodRef{
										Trait: &name.Name{
											Parts: []node.Node{
												&name.NamePart{Value: "Baz"},
											},
										},
										Method: &node.Identifier{Value: "one"},
									},
									Alias: &node.Identifier{Value: "two"},
								},
							},
						},
					},
				},
			},
			&stmt.Try{
				Stmts:   []node.Node{},
				Catches: []node.Node{},
			},
			&stmt.Try{
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Exception"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.Try{
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Exception"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "RuntimeException"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.Try{
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Exception"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Types: []node.Node{
							&name.FullyQualified{
								Parts: []node.Node{
									&name.NamePart{Value: "RuntimeException"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Types: []node.Node{
							&name.Relative{
								Parts: []node.Node{
									&name.NamePart{Value: "AdditionException"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
					},
				},
			},
			&stmt.Try{
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Exception"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
					},
				},
				Finally: &stmt.Finally{
					Stmts: []node.Node{},
				},
			},
			&stmt.Unset{
				Vars: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "a"}},
					&expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
					},
				},
			},
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
					},
				},
			},
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
						Alias: &node.Identifier{Value: "Bar"},
					},
				},
			},
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
					},
				},
			},
			&stmt.UseList{
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
						Alias: &node.Identifier{Value: "Baz"},
					},
				},
			},
			&stmt.UseList{
				UseType: &node.Identifier{Value: "function"},
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
					},
				},
			},
			&stmt.UseList{
				UseType: &node.Identifier{Value: "function"},
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
						Alias: &node.Identifier{Value: "foo"},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
						Alias: &node.Identifier{Value: "bar"},
					},
				},
			},
			&stmt.UseList{
				UseType: &node.Identifier{Value: "const"},
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
					},
				},
			},
			&stmt.UseList{
				UseType: &node.Identifier{Value: "const"},
				Uses: []node.Node{
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
						Alias: &node.Identifier{Value: "foo"},
					},
					&stmt.Use{
						Use: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Bar"},
							},
						},
						Alias: &node.Identifier{Value: "bar"},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.ArrayDimFetch{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Dim:      &scalar.Lnumber{Value: "1"},
				},
			},
			&stmt.Expression{
				Expr: &expr.ArrayDimFetch{
					Variable: &expr.ArrayDimFetch{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Dim:      &scalar.Lnumber{Value: "1"},
					},
					Dim: &scalar.Lnumber{Value: "2"},
				},
			},
			&stmt.Expression{
				Expr: &expr.Array{
					Items: []node.Node{},
				},
			},
			&stmt.Expression{
				Expr: &expr.Array{
					Items: []node.Node{
						&expr.ArrayItem{
							Val: &scalar.Lnumber{Value: "1"},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.Array{
					Items: []node.Node{
						&expr.ArrayItem{
							Key: &scalar.Lnumber{Value: "1"},
							Val: &scalar.Lnumber{Value: "1"},
						},
						&expr.ArrayItem{
							Val: &expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
						nil,
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.Array{
					Items: []node.Node{
						&expr.ArrayItem{
							Key: &scalar.Lnumber{Value: "3"},
							Val: &expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.Array{
					Items: []node.Node{
						&expr.ArrayItem{
							Val: &expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
						&expr.ArrayItem{
							Key: &scalar.Lnumber{Value: "1"},
							Val: &scalar.Lnumber{Value: "1"},
						},
						&expr.ArrayItem{
							Val: &scalar.Lnumber{Value: "1"},
						},
						&expr.ArrayItem{
							Key: &scalar.Lnumber{Value: "3"},
							Val: &expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.BitwiseNot{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.BooleanNot{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.ClassConstFetch{
					Class: &name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					ConstantName: &node.Identifier{Value: "Bar"},
				},
			},
			&stmt.Expression{
				Expr: &expr.Clone{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Clone{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Closure{
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Stmts:         []node.Node{},
				},
			},
			&stmt.Expression{
				Expr: &expr.Closure{
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Params: []node.Node{
						&node.Parameter{
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						},
						&node.Parameter{
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						},
					},
					ClosureUse: &expr.ClosureUse{
						Uses: []node.Node{
							&expr.Variable{VarName: &node.Identifier{Value: "c"}},
							&expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "d"}}},
						},
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Expression{
				Expr: &expr.Closure{
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Params: []node.Node{
						&node.Parameter{
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						},
						&node.Parameter{
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						},
					},
					ClosureUse: &expr.ClosureUse{
						Uses: []node.Node{
							&expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "c"}}},
							&expr.Variable{VarName: &node.Identifier{Value: "d"}},
						},
					},
					Stmts: []node.Node{},
				},
			},
			&stmt.Expression{
				Expr: &expr.Closure{
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Stmts:         []node.Node{},
				},
			},
			&stmt.Expression{
				Expr: &expr.ConstFetch{
					Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
				},
			},
			&stmt.Expression{
				Expr: &expr.ConstFetch{
					Constant: &name.Relative{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
				},
			},
			&stmt.Expression{
				Expr: &expr.ConstFetch{
					Constant: &name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Empty{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Empty{
					Expr: &expr.ConstFetch{
						Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.ErrorSuppress{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Eval{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Exit{},
			},
			&stmt.Expression{
				Expr: &expr.Exit{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Die{},
			},
			&stmt.Expression{
				Expr: &expr.Die{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "foo"},
						},
					},
					ArgumentList: &node.ArgumentList{},
				},
			},
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.Relative{
						Parts: []node.Node{
							&name.NamePart{Value: "foo"},
						},
					},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{
								Variadic:    false,
								IsReference: true,
								Expr:        &expr.Variable{VarName: &node.Identifier{Value: "a"}},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.FullyQualified{
						Parts: []node.Node{
							&name.NamePart{Value: "foo"},
						},
					},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{
								Variadic:    false,
								IsReference: false,
								Expr: &expr.ShortArray{
									Items: []node.Node{},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					ArgumentList: &node.ArgumentList{
						Arguments: []node.Node{
							&node.Argument{
								Variadic:    false,
								IsReference: false,
								Expr: &expr.Yield{
									Value: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.PostDec{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.PostInc{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.PreDec{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.PreInc{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Include{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.IncludeOnce{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Require{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.RequireOnce{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.InstanceOf{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Class: &name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.InstanceOf{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Class: &name.Relative{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.InstanceOf{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Class: &name.FullyQualified{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.Isset{
					Variables: []node.Node{
						&expr.Variable{VarName: &node.Identifier{Value: "a"}},
						&expr.Variable{VarName: &node.Identifier{Value: "b"}},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.Isset{
					Variables: []node.Node{
						&expr.ConstFetch{
							Constant: &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.List{
						Items: []node.Node{},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.List{
						Items: []node.Node{
							&expr.ArrayItem{
								Val: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
							},
							&expr.ArrayItem{
								Val: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.List{
						Items: []node.Node{
							&expr.ArrayItem{
								Val: &expr.ArrayDimFetch{
									Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
								},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.List{
						Items: []node.Node{
							&expr.ArrayItem{
								Val: &expr.List{
									Items: []node.Node{
										&expr.ArrayItem{
											Val: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
										},
									},
								},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.MethodCall{
					Variable:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Method:       &node.Identifier{Value: "foo"},
					ArgumentList: &node.ArgumentList{},
				},
			},
			&stmt.Expression{
				Expr: &expr.New{
					Class: &name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.New{
					Class: &name.Relative{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					ArgumentList: &node.ArgumentList{},
				},
			},
			&stmt.Expression{
				Expr: &expr.New{
					Class: &name.FullyQualified{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					ArgumentList: &node.ArgumentList{},
				},
			},
			&stmt.Expression{
				Expr: &expr.Print{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.PropertyFetch{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Property: &node.Identifier{Value: "foo"},
				},
			},
			&stmt.Expression{
				Expr: &expr.ArrayDimFetch{
					Variable: &expr.PropertyFetch{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Property: &node.Identifier{Value: "foo"},
					},
					Dim: &scalar.Lnumber{Value: "1"},
				},
			},
			&stmt.Expression{
				Expr: &expr.ArrayDimFetch{
					Variable: &expr.PropertyFetch{
						Variable: &expr.MethodCall{
							Variable: &expr.PropertyFetch{
								Variable: &expr.PropertyFetch{
									Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
									Property: &node.Identifier{Value: "foo"},
								},
								Property: &node.Identifier{Value: "bar"},
							},
							Method:       &node.Identifier{Value: "baz"},
							ArgumentList: &node.ArgumentList{},
						},
						Property: &node.Identifier{Value: "quux"},
					},
					Dim: &scalar.Lnumber{Value: "0"},
				},
			},
			&stmt.Expression{
				Expr: &expr.ArrayDimFetch{
					Variable: &expr.ArrayDimFetch{
						Variable: &expr.MethodCall{
							Variable:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
							Method:       &node.Identifier{Value: "foo"},
							ArgumentList: &node.ArgumentList{},
						},
						Dim: &scalar.Lnumber{Value: "1"},
					},
					Dim: &scalar.Lnumber{Value: "1"},
				},
			},
			&stmt.Expression{
				Expr: &expr.ShellExec{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "cmd "},
						&expr.Variable{VarName: &node.Identifier{Value: "a"}},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.ShellExec{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "cmd"},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.ShellExec{
					Parts: []node.Node{},
				},
			},
			&stmt.Expression{
				Expr: &expr.ShortArray{
					Items: []node.Node{},
				},
			},
			&stmt.Expression{
				Expr: &expr.ShortArray{
					Items: []node.Node{
						&expr.ArrayItem{
							Val: &scalar.Lnumber{Value: "1"},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.ShortArray{
					Items: []node.Node{
						&expr.ArrayItem{
							Key: &scalar.Lnumber{Value: "1"},
							Val: &scalar.Lnumber{Value: "1"},
						},
						&expr.ArrayItem{
							Val: &expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
						nil,
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class: &name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					Call:         &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class: &name.Relative{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					Call:         &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class: &name.FullyQualified{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					Call:         &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class: &name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					Call:         &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
					ArgumentList: &node.ArgumentList{},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class:        &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					Call:         &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
					ArgumentList: &node.ArgumentList{},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticPropertyFetch{
					Class: &name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					Property: &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticPropertyFetch{
					Class: &name.Relative{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					Property: &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticPropertyFetch{
					Class: &name.FullyQualified{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					Property: &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Ternary{
					Condition: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					IfTrue:    &expr.Variable{VarName: &node.Identifier{Value: "b"}},
					IfFalse:   &expr.Variable{VarName: &node.Identifier{Value: "c"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Ternary{
					Condition: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					IfFalse:   &expr.Variable{VarName: &node.Identifier{Value: "c"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Ternary{
					Condition: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					IfTrue: &expr.Ternary{
						Condition: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						IfTrue:    &expr.Variable{VarName: &node.Identifier{Value: "c"}},
						IfFalse:   &expr.Variable{VarName: &node.Identifier{Value: "d"}},
					},
					IfFalse: &expr.Variable{VarName: &node.Identifier{Value: "e"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Ternary{
					Condition: &expr.Ternary{
						Condition: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						IfTrue:    &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						IfFalse:   &expr.Variable{VarName: &node.Identifier{Value: "c"}},
					},
					IfTrue:  &expr.Variable{VarName: &node.Identifier{Value: "d"}},
					IfFalse: &expr.Variable{VarName: &node.Identifier{Value: "e"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.UnaryMinus{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.UnaryPlus{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Variable{VarName: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
			},
			&stmt.Expression{
				Expr: &expr.Variable{VarName: &expr.Variable{VarName: &expr.Variable{VarName: &node.Identifier{Value: "a"}}}},
			},
			&stmt.Expression{
				Expr: &expr.Yield{},
			},
			&stmt.Expression{
				Expr: &expr.Yield{
					Value: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Yield{
					Key:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Value: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Yield{
					Value: &expr.ClassConstFetch{
						Class: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
						ConstantName: &node.Identifier{Value: "class"},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.Yield{
					Key: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Value: &expr.ClassConstFetch{
						Class: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
						ConstantName: &node.Identifier{Value: "class"},
					},
				},
			},
			&stmt.Expression{
				Expr: &cast.Array{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &cast.Bool{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &cast.Bool{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &cast.Double{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &cast.Double{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &cast.Int{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &cast.Int{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &cast.Object{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &cast.String{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &cast.Unset{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.BitwiseAnd{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.BitwiseOr{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.BitwiseXor{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.BooleanAnd{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.BooleanOr{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.Concat{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.Div{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.Equal{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.GreaterOrEqual{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.Greater{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.Identical{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.LogicalAnd{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.LogicalOr{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.LogicalXor{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.Minus{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.Mod{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.Mul{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.NotEqual{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.NotIdentical{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.Plus{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.Pow{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.ShiftLeft{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.ShiftRight{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.SmallerOrEqual{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &binary.Smaller{
					Left:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Right: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.Reference{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.Reference{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.New{
						Class: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &assign.Reference{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.New{
						Class: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "Foo"},
							},
						},
						ArgumentList: &node.ArgumentList{
							Arguments: []node.Node{
								&node.Argument{
									Variadic:    false,
									IsReference: false,
									Expr:        &expr.Variable{VarName: &node.Identifier{Value: "b"}},
								},
							},
						},
					},
				},
			},
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.BitwiseAnd{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.BitwiseOr{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.BitwiseXor{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.Concat{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.Div{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.Minus{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.Mod{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.Mul{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.Plus{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.Pow{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.ShiftLeft{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &assign.ShiftRight{
					Variable:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.New{
					Class: &name.FullyQualified{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					ArgumentList: &node.ArgumentList{},
				},
			},
			&stmt.Expression{
				Expr: &expr.PropertyFetch{
					Variable: &expr.MethodCall{
						Variable: &expr.New{
							Class: &name.FullyQualified{
								Parts: []node.Node{
									&name.NamePart{Value: "Foo"},
								},
							},
							ArgumentList: &node.ArgumentList{},
						},
						Method:       &node.Identifier{Value: "bar"},
						ArgumentList: &node.ArgumentList{},
					},
					Property: &node.Identifier{Value: "baz"},
				},
			},
			&stmt.Expression{
				Expr: &expr.ArrayDimFetch{
					Variable: &expr.ArrayDimFetch{
						Variable: &expr.New{
							Class: &name.FullyQualified{
								Parts: []node.Node{
									&name.NamePart{Value: "Foo"},
								},
							},
							ArgumentList: &node.ArgumentList{},
						},
						Dim: &scalar.Lnumber{Value: "0"},
					},
					Dim: &scalar.Lnumber{Value: "0"},
				},
			},
			&stmt.Expression{
				Expr: &expr.MethodCall{
					Variable: &expr.ArrayDimFetch{
						Variable: &expr.New{
							Class: &name.FullyQualified{
								Parts: []node.Node{
									&name.NamePart{Value: "Foo"},
								},
							},
							ArgumentList: &node.ArgumentList{},
						},
						Dim: &scalar.Lnumber{Value: "0"},
					},
					Method:       &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{},
				},
			},
			&stmt.Expression{
				Expr: &expr.ArrayDimFetch{
					Variable: &expr.ArrayDimFetch{
						Variable: &expr.Array{
							Items: []node.Node{
								&expr.ArrayItem{
									Val: &expr.ShortArray{
										Items: []node.Node{
											&expr.ArrayItem{
												Val: &scalar.Lnumber{Value: "0"},
											},
										},
									},
								},
							},
						},
						Dim: &scalar.Lnumber{Value: "0"},
					},
					Dim: &scalar.Lnumber{Value: "0"},
				},
			},
			&stmt.Expression{
				Expr: &expr.ArrayDimFetch{
					Variable: &scalar.String{Value: "\"foo\""},
					Dim:      &scalar.Lnumber{Value: "0"},
				},
			},
			&stmt.Expression{
				Expr: &expr.ArrayDimFetch{
					Variable: &expr.ConstFetch{
						Constant: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "foo"},
							},
						},
					},
					Dim: &scalar.Lnumber{Value: "0"},
				},
			},
			&stmt.Expression{
				Expr: &expr.ClassConstFetch{
					Class:        &node.Identifier{Value: "static"},
					ConstantName: &node.Identifier{Value: "foo"},
				},
			},
			&stmt.Expression{
				Expr: &expr.New{
					Class: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.New{
					Class: &expr.StaticPropertyFetch{
						Class:    &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
						Property: &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.New{
					Class: &expr.ArrayDimFetch{
						Variable: &expr.PropertyFetch{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
							Property: &node.Identifier{Value: "b"},
						},
						Dim: &scalar.Lnumber{Value: "0"},
					},
				},
			},
			&stmt.Expression{
				Expr: &expr.New{
					Class: &expr.ArrayDimFetch{
						Variable: &expr.PropertyFetch{
							Variable: &expr.PropertyFetch{
								Variable: &expr.ArrayDimFetch{
									Variable: &expr.PropertyFetch{
										Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
										Property: &node.Identifier{Value: "b"},
									},
									Dim: &expr.Ternary{
										Condition: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
										IfFalse: &expr.ConstFetch{
											Constant: &name.Name{
												Parts: []node.Node{
													&name.NamePart{Value: "null"},
												},
											},
										},
									},
								},
								Property: &expr.Variable{VarName: &node.Identifier{Value: "c"}},
							},
							Property: &node.Identifier{Value: "d"},
						},
						Dim: &scalar.Lnumber{Value: "0"},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &expr.ArrayDimFetch{
							Variable: &expr.ShortArray{
								Items: []node.Node{
									&expr.ArrayItem{
										Val: &scalar.Lnumber{Value: "1"},
									},
								},
							},
							Dim: &scalar.Lnumber{Value: "0"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &expr.BooleanNot{
							Expr: &scalar.Lnumber{Value: "1"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &expr.BitwiseNot{
							Expr: &scalar.Lnumber{Value: "1"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &expr.UnaryPlus{
							Expr: &scalar.Lnumber{Value: "1"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &expr.UnaryMinus{
							Expr: &scalar.Lnumber{Value: "1"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr:     &scalar.Lnumber{Value: "1"},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &expr.Ternary{
							Condition: &scalar.Lnumber{Value: "1"},
							IfFalse:   &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &expr.Ternary{
							Condition: &scalar.Lnumber{Value: "1"},
							IfTrue:    &scalar.Lnumber{Value: "2"},
							IfFalse:   &scalar.Lnumber{Value: "3"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.BitwiseAnd{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.BitwiseOr{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.BitwiseXor{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.BooleanAnd{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.BooleanOr{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.Concat{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.Div{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.Equal{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.GreaterOrEqual{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.Greater{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.Identical{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.LogicalAnd{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.LogicalOr{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.LogicalXor{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.Minus{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.Mod{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.Mul{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.NotEqual{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.NotIdentical{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.Plus{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.Pow{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.ShiftLeft{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.ShiftRight{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.SmallerOrEqual{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &binary.Smaller{
							Left:  &scalar.Lnumber{Value: "1"},
							Right: &scalar.Lnumber{Value: "2"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &expr.ClassConstFetch{
							Class: &name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Foo"},
								},
							},
							ConstantName: &node.Identifier{Value: "bar"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &expr.ClassConstFetch{
							Class: &name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Foo"},
								},
							},
							ConstantName: &node.Identifier{Value: "class"},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr:     &scalar.MagicConstant{Value: "__CLASS__"},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &expr.ConstFetch{
							Constant: &name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Foo"},
								},
							},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &expr.ConstFetch{
							Constant: &name.Relative{
								Parts: []node.Node{
									&name.NamePart{Value: "Foo"},
								},
							},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &expr.ConstFetch{
							Constant: &name.FullyQualified{
								Parts: []node.Node{
									&name.NamePart{Value: "Foo"},
								},
							},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr:     &expr.Array{},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &expr.Array{
							Items: []node.Node{
								&expr.ArrayItem{
									Key: &scalar.Lnumber{Value: "1"},
									Val: &scalar.Lnumber{Value: "1"},
								},
								&expr.ArrayItem{
									Val: &scalar.Lnumber{Value: "2"},
								},
							},
						},
					},
				},
			},
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr: &expr.ArrayDimFetch{
							Variable: &expr.ShortArray{
								Items: []node.Node{
									&expr.ArrayItem{
										Val: &scalar.Lnumber{Value: "1"},
									},
									&expr.ArrayItem{
										Key: &scalar.Lnumber{Value: "2"},
										Val: &scalar.Lnumber{Value: "2"},
									},
								},
							},
							Dim: &scalar.Lnumber{Value: "0"},
						},
					},
				},
			},
			&stmt.If{
				Cond: &expr.Yield{
					Value: &scalar.Lnumber{Value: "1"},
				},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{},
				},
			},
			&stmt.Expression{
				Expr: &expr.StaticPropertyFetch{
					Class:    &name.Name{Parts: []node.Node{&name.NamePart{Value: "Foo"}}},
					Property: &expr.Variable{VarName: &expr.Variable{VarName: &node.Identifier{Value: "bar"}}},
				},
			},
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function:     &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					ArgumentList: &node.ArgumentList{},
				},
			},
			&stmt.Expression{
				Expr: &expr.ArrayDimFetch{
					Variable: &expr.ArrayDimFetch{
						Variable: &expr.FunctionCall{
							Function:     &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
							ArgumentList: &node.ArgumentList{},
						},
						Dim: &scalar.Lnumber{Value: "0"},
					},
					Dim: &scalar.Lnumber{Value: "0"},
				},
			},
			&stmt.Expression{
				Expr: &expr.ArrayDimFetch{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Dim:      &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
			&stmt.Expression{
				Expr: &expr.Variable{VarName: &expr.Variable{VarName: &node.Identifier{Value: "a"}}},
			},
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class:        &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					Call:         &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
					ArgumentList: &node.ArgumentList{},
				},
			},
			&stmt.Expression{
				Expr: &expr.ClassConstFetch{
					Class:        &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					ConstantName: &node.Identifier{Value: "bar"},
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
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.String{Value: "\"test\""},
			},
			&stmt.Expression{
				Expr: &scalar.String{Value: "\"\\$test\""},
			},
			&stmt.Expression{
				Expr: &scalar.String{Value: "\"\n\t\t\ttest\n\t\t\""},
			},
			&stmt.Expression{
				Expr: &scalar.String{Value: "'$test'"},
			},
			&stmt.Expression{
				Expr: &scalar.String{Value: "'\n\t\t\t$test\n\t\t'"},
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
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Heredoc{
					Label: "CAD",
				},
			},
			&stmt.Expression{
				Expr: &scalar.Heredoc{
					Label: "CAD",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "\thello\n"},
					},
				},
			},
			&stmt.Expression{
				Expr: &scalar.Heredoc{
					Label: "\"CAD\"",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "\thello\n"},
					},
				},
			},
			&stmt.Expression{
				Expr: &scalar.Heredoc{
					Label: "\"CAD\"",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "\thello "},
						&expr.Variable{VarName: &node.Identifier{Value: "world"}},
						&scalar.EncapsedStringPart{Value: "\n"},
					},
				},
			},
			&stmt.Expression{
				Expr: &scalar.Heredoc{
					Label: "'CAD'",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "\thello $world\n"},
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
