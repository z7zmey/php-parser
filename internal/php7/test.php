<?
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
`cmd $a`;
`cmd`;
``;
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