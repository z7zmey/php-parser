package printer_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/printer"
)

func ExamplePrinter() {
	src := `<?php

namespace Foo;

abstract class Bar extends Baz
{
    public function greet()
    {
        echo "Hello";
        // some comment
    }
}
	`

	// parse

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.WithMeta()
	php7parser.Parse()

	rootNode := php7parser.GetRootNode()

	// change

	parts := &rootNode.(*node.Root).Stmts[0].(*stmt.Namespace).NamespaceName.(*name.Name).Parts
	*parts = append(*parts, &name.NamePart{Value: "Quuz"})

	// print

	p := printer.NewPrinter(os.Stdout)
	p.Print(rootNode)

	// Output:
	//<?php
	//
	// namespace Foo\Quuz;
	//
	// abstract class Bar extends Baz
	// {
	//     public function greet()
	//     {
	//         echo "Hello";
	//         // some comment
	//     }
	// }
}

func parse(src string) node.Node {
	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.WithMeta()
	php7parser.Parse()

	return php7parser.GetRootNode()
}

func print(n node.Node) string {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(n)

	return o.String()
}

func TestParseAndPrintRoot(t *testing.T) {

	src := `<?php

	namespace Foo;

	abstract class Bar extends Baz
	{
		public function greet()
		{
			echo "Hello World";
		}
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintRootWithMeta(t *testing.T) {

	src := `<?php

	namespace Foo;

	abstract class Bar extends Baz
	{
		public function greet()
		{
			echo "Hello World";
		}
	}
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintRootWithHtml(t *testing.T) {

	src := `<div>Hello</div>
	<?php
	$a;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintFunctionAndParameter(t *testing.T) {

	src := `<?php
	function & foo (
		? int $a , & $b = null
		, \ Foo ...$c
	) : namespac \ Bar {
		;
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintFunctionCallAndArgument(t *testing.T) {

	src := `<?php
	foo ( $a , $b
		, ... $c
	) ; `

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIncludeAndRequire(t *testing.T) {

	src := `<?php
	include 'foo' ;
	include_once 'bar' ;
	
	require __DIR__ . '/folder' ;
	require_once $a ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIfInstanceOfMethodCallIssetListEcho(t *testing.T) {
	src := `<?php
	
	if ( $a instanceof Foo ) {
		$a -> bar ( ' ' ) ;
	} elseif ( isset ( $b [ 2 ] ) )
		list( , $c , ) = $b;
	else if ( 1 );
	else {
		echo '' ;
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNew(t *testing.T) {
	src := `<?php

	new Foo ;

	new Foo ( $a, $b ) ;

	new class ( $c ) extends Foo implements Bar , Baz {

	} ; `

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIncDec(t *testing.T) {
	src := `<?php
	++ $a ;
	-- $a ;
	$a ++ ;
	$a -- ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPropertyFetchPrint(t *testing.T) {
	src := `<?php
	$a -> b ;
	print ( $a ) ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintForeachReferenceShellExec(t *testing.T) {
	src := `<?php
	foreach ( $a as $k => & $v ) {
		` + "` $v cmd `" + ` ;
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintShortListShortArray(t *testing.T) {
	src := `<?php
	$a = [0, 1, 2] ;
	[ , $b ,] = $a ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintStaticCallStaticPropertyFetch(t *testing.T) {
	src := `<?php
	Foo :: bar ( $a , $b ) ;
	Foo :: $bar ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintTernary(t *testing.T) {
	src := `<?php
	$a ? $b : $c ;
	$a ? : $c ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintUnaryPlusMinus(t *testing.T) {
	src := `<?php
	+ $a ;
	- $a ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintYield(t *testing.T) {
	src := `<?php
	yield $a ;
	yield $k => $v ;
	yield from $a ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAltIfExitDie(t *testing.T) {
	src := `<?php
	if ( 1 ) :
		exit ( 1 ) ;
	elseif ( 2 ) :
		die ( 2 );
	else :
		;
	endif ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAltFor(t *testing.T) {
	src := `<?php
	for ( $a ; $b ; $c ) :
	endfor ;
	
	for ( ; ; ) :
	endfor ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAltForeach(t *testing.T) {
	src := `<?php
	foreach ( $a as $k => & $v ) :
		echo $v ;
	endforeach ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAltSwitchCaseBreakContinueDefault(t *testing.T) {
	src := `<?php

	switch ( $a ) : ;
	case 1 : 
		break ;
	case 2 :
		break ( 2 ) ;
	case 3 :
		continue ;
	default :
		continue ( 2 ) ;
	endswitch ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAltWhile(t *testing.T) {
	src := `<?php

	while ( $a ) :
		;
	endwhile ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintTryCatchFinally(t *testing.T) {
	src := `<?php

	try {

	} catch ( \ Exception | \ Foo \ Bar $e) {

	} finally {

	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassMethodProperty(t *testing.T) {
	src := `<?php
	abstract final class Foo extends Bar implements Baz , Quuz {
		public const FOO = 'f' , BAR = 'b' ;

		var $a = '' , $b = null , $c ;
		
		public static function & greet ( ? Foo $a ) : void ;
		
		function greet ( string $a )
		{
			return 'hello';
		}
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintDeclare(t *testing.T) {
	src := `<?php
	declare ( FOO = 'bar' ) {
		;
	}
	
	declare ( FOO = 'bar' ) ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintEmpty(t *testing.T) {
	src := `<?php
	empty ( $a ) ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIfElseIfElse(t *testing.T) {
	src := `<?php
	if ( 1 ) ;
	elseif ( 2 ) {
		;
	}
	else if ( 3 ) $a;
	else { }`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintFor(t *testing.T) {
	src := `<?php
	for ( $i = 0 ; $i < 3 ; $i ++ ) 
		echo $i . PHP_EOL;
	
	for ( ; ; ) {

	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintGlobal(t *testing.T) {
	src := `<?php
	global $a , $b ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintGotoLabel(t *testing.T) {
	src := `<?php
	Foo :
	goto Foo ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintUseGroupUse(t *testing.T) {
	src := `<?php
	use const Foo as Bar, baz ;
	use function Foo \ { Bar as Baz , Quuz } ;
	use Foo \ { function Bar as Baz , const Quuz } ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintHaltCompiller(t *testing.T) {
	src := `<?php
	__halt_compiler ( ) ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintInlineHtml(t *testing.T) {
	src := `<?php
	$a;?>test<?php `

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintInterface(t *testing.T) {
	src := `<?php
	interface Foo extends Bar , Baz {

	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNamespace(t *testing.T) {
	src := `<?php
	namespace Foo \ Bar ; 
	namespace Baz {

	}
	namespace {
		
	}
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintStaticVar(t *testing.T) {
	src := `<?php
	static $a , $b = foo ( ) ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintSwitchCaseDefault(t *testing.T) {
	src := `<?php

	switch ( $a ) {
		case 1 : ;
		default : ;
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintThrow(t *testing.T) {
	src := `<?php
	throw new \ Exception ( "msg" ) ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintTraitUse(t *testing.T) {
	src := `<?php
	class foo {
		use foo , bar ;
		use foo , bar { }
		use foo , bar {
			foo :: a as b ;
			bar :: a insteadof foo ;
			foo :: c as public ;
			foo :: d as public e;
			f as g ;
		}
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintTrait(t *testing.T) {
	src := `<?php
	trait foo {
		function bar ( ) { }
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintUnset(t *testing.T) {
	src := `<?php
	unset ( $a ) ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintWhile(t *testing.T) {
	src := `<?php
	while ( $a ) echo '' ;
	while ( $a ) { }
	while ( $a ) ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintDoWhile(t *testing.T) {
	src := `<?php
	do { ; } while ( $a ) ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNames(t *testing.T) {
	src := `<?php
	new \ Exception ;
	new namespace \ Exception ;
	new Exception ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintScalar(t *testing.T) {
	src := `<?php
	1 ;
	.2 ;
	0.2 ;
	'Hello' ;
	"Hello $world";
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintHeredoc(t *testing.T) {
	src := `<?php
	foo(<<<EAP
test
EAP
, 'test'
);

<<<EAP
test
EAP;

<<<'EAP'
test
EAP;

<<<"EAP"
test
EAP;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintMagicConstant(t *testing.T) {
	src := `<?php
	__CLASS__ ;
	__DIR__ ;
	__FILE__ ;
	__FUNCTION__ ;
	__LINE__ ;
	__NAMESPACE__ ;
	__METHOD__ ;
	__TRAIT__ ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAssign(t *testing.T) {
	src := `<?php
	$a = $b ;
	$a = & $b ;
	$a &= $b ;
	$a |= $b ;
	$a ^= $b ;
	$a .= $b ;
	$a /= $b ;
	$a -= $b ;
	$a %= $b ;
	$a *= $b ;
	$a += $b ;
	$a **= $b ;
	$a <<= $b ;
	$a >>= $b ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintBinary(t *testing.T) {
	src := `<?php
	$a & $b ;
	$a | $b ;
	$a ^ $b ;
	$a && $b ;
	$a || $b ;
	$a ?? $b ;
	$a . $b ;
	$a / $b ;
	$a == $b ;
	$a >= $b ;
	$a > $b ;
	$a === $b ;
	$a and $b ;
	$a or $b ;
	$a xor $b ;
	$a - $b ;
	$a % $b ;
	$a * $b ;
	$a != $b ;
	$a !== $b ;
	$a + $b ;
	$a ** $b ;
	$a << $b ;
	$a >> $b ;
	$a <= $b ;
	$a < $b ;
	$a <=> $b ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintUnary(t *testing.T) {
	src := `<?php
	~ $var ;
	! $var ;
	+ $var ;
	- $var ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintCast(t *testing.T) {
	src := `<?php
	(array) $a ;
	(bool) $a ;
	(float) $a ;
	(int) $a ;
	(object) $a ;
	(string) $a ;
	(unset) $a ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintArray(t *testing.T) {
	src := `<?php
	$a [ 0 ] ;

	$a = [ 0 , 2 => 2 ] ;
	$a = array( 0 , 2 => 2 ) ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassConstFetch(t *testing.T) {
	src := `<?php
	$var :: CONST ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClone(t *testing.T) {
	src := `<?php
	clone $var ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClosure(t *testing.T) {
	src := `<?php
	$a  = static function & ( ) use ( $a , & $b ) : void { } ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintConstFetch(t *testing.T) {
	src := `<?php
	null ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintErrorSupress(t *testing.T) {
	src := `<?php
	@ foo ( ) ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintEval(t *testing.T) {
	src := `<?php
	eval ( " " ) ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}
