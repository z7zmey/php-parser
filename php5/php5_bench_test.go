package php5_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/php5"
)

func BenchmarkPhp5(b *testing.B) {
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

		"test";
		"\$test";
		"
			test
		";
		'$test';
		'
			$test
		';
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

	for n := 0; n < b.N; n++ {
		php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
		php5parser.Parse()
	}
}
