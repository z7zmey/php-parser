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
	php7parser.WithFreeFloating()
	php7parser.Parse()

	rootNode := php7parser.GetRootNode()

	// change namespace

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
	php7parser.WithFreeFloating()
	php7parser.Parse()

	return php7parser.GetRootNode()
}

func print(n node.Node) string {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	p.Print(n)

	return o.String()
}

// test node

func TestParseAndPrintRoot(t *testing.T) {

	src := ` <div>Hello</div> 
	<?php
	$a;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIdentifier(t *testing.T) {

	src := `<?
	/* Foo */
	Foo ( ) ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintParameter(t *testing.T) {

	src := `<?php
	function & foo (
		? int $a , & $b = null
		, \ Foo ...$c
	) : namespace  \ Bar \  baz \ quuz{
		;
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNullable(t *testing.T) {

	src := `<?php
	function & foo ( ? int $a ) {
		/* do nothing */
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintArgument(t *testing.T) {

	src := `<?php
	foo ( $a , $b
		, ... $c ,
	) ; `

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test name

func TestParseAndPrintNames(t *testing.T) {
	src := `<?php
	foo ( ) ;
	\ foo ( ) ;
	namespace \ foo ( ) ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test scalar

func TestParseAndPrintMagicConstant(t *testing.T) {
	src := `<?php
	__CLASS__     ;
	__DIR__       ;
	__FILE__      ;
	__FUNCTION__  ;
	__LINE__      ;
	__NAMESPACE__ ;
	__METHOD__    ;
	__TRAIT__     ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNumber(t *testing.T) {
	src := `<?php
	// LNumber
	1234567890123456789 ;

	// DNumber
	12345678901234567890 ;
	0. ;
	.2 ;
	0.2 ;

	// binary LNumber
	0b0111111111111111111111111111111111111111111111111111111111111111 ;

	// binary DNumber
	0b1111111111111111111111111111111111111111111111111111111111111111 ;

	// HLNumber
	0x007111111111111111 ;

	// HDNumber
	0x8111111111111111 ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintString(t *testing.T) {
	src := `<?php
	'Hello' ;
	"Hello {$world } " ;
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

// test assign

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

// test binary

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
	$a <> $b ;
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

// test cast

func TestParseAndPrintCast(t *testing.T) {
	src := `<?php
	(  array     ) $a ;
	(  bool      ) $a ;
	(  boolean   ) $a ;
	(  real      ) $a ;
	(  double    ) $a ;
	(  float     ) $a ;
	(  int       ) $a ;
	(  integer   ) $a ;
	(  object    ) $a ;
	(  string    ) $a ;
	(  binary    ) $a ;
	(  unset     ) $a ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test expr

func TestParseAndPrintArrayDimFetch(t *testing.T) {
	src := `<?php
	FOO [ ] ;
	FOO [ 1 ] ;
	$a [ ] ;
	$a [ 1 ] ;
	$a { 1 } ;
	new $a [ ] ;
	new $a [ 1 ] ;
	new $a { 1 } ;
	"$a[1]test" ;
	"${ a [ 1 ] }test" ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintArrayItem(t *testing.T) {
	src := `<?php
	$foo = [
		$world ,
		& $world ,
		'Hello' => $world ,
	] ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintArray(t *testing.T) {
	src := `<?php
	array ( /* empty array */ ) ;
	array ( 0 , 2 => 2 ) ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintBitwiseNot(t *testing.T) {
	src := `<?php
	~ $var ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintBooleanNot(t *testing.T) {
	src := `<?php
	! $var ;
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

func TestParseAndPrintClosureUse(t *testing.T) {
	src := `<?php
	$a = function ( ) use ( $a , & $b ) {
		// do nothing
	} ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClosure(t *testing.T) {
	src := `<?php
	$a  = static function & ( ) : void {
		// do nothing
	} ;
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

func TestParseAndPrintEmpty(t *testing.T) {
	src := `<?php
	empty ( $a ) ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintErrorSuppress(t *testing.T) {
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

func TestParseAndPrintExit(t *testing.T) {
	src := `<?php
	exit ;
	exit ( ) ;
	exit (1) ;
	exit ( 1 );
	die ;
	die ( ) ;
	die (1) ;
	die ( 1 );
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintFunctionCall(t *testing.T) {
	src := `<?php
	foo ( ) ;
	$var ( $a , ... $b , $c ) ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintInclude(t *testing.T) {

	src := `<?php
	include 'foo' ;
	include_once 'bar' ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintInstanceOf(t *testing.T) {
	src := `<?php
	$a instanceof Foo ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIsset(t *testing.T) {
	src := `<?php
	isset ( $a , $b [ 2 ] , ) ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintList(t *testing.T) {
	src := `<?php
	list( , $var , ) = $b ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintMethodCall(t *testing.T) {
	src := `<?php
	$a -> bar ( $arg , ) ;`

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

func TestParseAndPrintPrint(t *testing.T) {
	src := `<?php
	print $a ;
	print ( $a ) ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPropertyFetch(t *testing.T) {
	src := `<?php
	$a -> b ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintReference(t *testing.T) {
	src := `<?php
	$a = & $b ;
	$a = [ & $b ] ;
	$a = [ $b => & $c ] ;

	$a = function ( ) use ( & $b ) {
		// do nothing
	} ;

	foreach ( $a as & $b ) {
		// do nothing
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintRequire(t *testing.T) {

	src := `<?php
	require __DIR__ . '/folder' ;
	require_once $a ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintShellExec(t *testing.T) {
	src := "<?php ` {$v} cmd ` ; "

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintShortArray(t *testing.T) {
	src := `<?php
	$a = [ ] ;
	$a = [ 0 ] ;
	$a = [
		1 => & $b , // one
		$c       , /* two */ 
	] ;
	$a = [0, 1, 2] ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintShortList(t *testing.T) {
	src := `<?php
	[ 
		/* skip */,
		$b 
		/* skip */,
	] = $a ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintStaticCall(t *testing.T) {
	src := `<?php
	Foo :: bar ( $a , $b ) ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintStaticPropertyFetch(t *testing.T) {
	src := `<?php
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

func TestParseAndPrintUnary(t *testing.T) {
	src := `<?php
	- $a ;
	+ $a ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintVariable(t *testing.T) {
	src := `<?php
	$ /* variable variable comment */ $var ; `

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

// test stmt

func TestParseAndPrintAltIf(t *testing.T) {
	src := `<?php
	if ( 1 ) :
		// do nothing
	elseif ( 2 ) :
	elseif ( 3 ) :
		;
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

func TestParseAndPrintAltSwitch(t *testing.T) {
	src := `<?php

	switch ( $a ) : 
	case 1 : 
		;
	case 2 : ;
	case 3 :
		 ;
	default :
		;
	endswitch ;
	

	switch ( $a ) : ;
	case 1 ; ;
	default ; ;
	endswitch ;

	switch ( $a ) :
	endswitch ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAltWhile(t *testing.T) {
	src := `<?php

	while ( $a ) :
		// do nothing
	endwhile ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintBreak(t *testing.T) {
	src := `<?php

	break ;
	break 1 ;
	break ( 2 ) ;
`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassMethod(t *testing.T) {
	src := `<?php
	class Foo {
		/**
		 * abstract method
		 */
		public static function & greet ( ? Foo $a ) : void ;
		
		function greet ( string $a )
		{
			return 'hello' ;
		}
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClass(t *testing.T) {
	src := `<?php
	abstract final class Foo extends Bar implements Baz , Quuz {
		
	}

	new class ( $c ) extends Foo implements Bar , Baz {

	} ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassConstList(t *testing.T) {
	src := `<?php
	class Foo {
		public const FOO = 'f' , BAR = 'b' ;
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintConstList(t *testing.T) {
	src := `<?php
	const FOO = 1 , BAR = 2 ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintContinue(t *testing.T) {
	src := `<?php

	continue ;
	continue 1 ;
	continue ( 2 ) ;
`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintDeclare(t *testing.T) {
	src := `<?php
	declare ( FOO = 'bar' ) ;
	declare ( FOO = 'bar' ) $a ;
	declare ( FOO = 'bar' ) { }

	declare ( FOO = 'bar' ) : enddeclare ;
	declare ( FOO = 'bar' ) :
		;
	enddeclare ;
`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintDoWhile(t *testing.T) {
	src := `<?php
	do {
		;
	} while ( $a ) ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintEcho(t *testing.T) {
	src := `<?php
	echo '' ;
	echo $a , ' ' , PHP_EOL;

	?>

	<?= $a, $b ?>
	<?= $c ; 

`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIfExpression(t *testing.T) {
	src := `<?php
	$a ; `

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

func TestParseAndPrintForeach(t *testing.T) {
	src := `<?php
	foreach ( $a as $k => & $v ) {
		;
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintFunction(t *testing.T) {

	src := `<?php
	function & foo ( ) : void {
		;
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

func TestParseAndPrintGoto(t *testing.T) {
	src := `<?php
	goto Foo ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintGroupUse(t *testing.T) {
	src := `<?php
	use function Foo \ { Bar as Baz , Quuz , } ;
	use Foo \ { function Bar as Baz , const Quuz } ;
	use \ Foo \ { function Bar as Baz , const Quuz , } ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintHaltCompiler(t *testing.T) {
	src := `<?php
	__halt_compiler ( ) ;
	this text is ignored by parser
	`

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

func TestParseAndPrintInlineHtml(t *testing.T) {
	src := `<?php
	$a;?>test<? `

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

func TestParseAndPrintGotoLabel(t *testing.T) {
	src := `<?php
	Foo : $b ; `

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

func TestParseAndPrintNop(t *testing.T) {
	src := `<?php ; `

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPropertyList(t *testing.T) {
	src := `<?php
	class Foo {
		var $a = '' , $b = null ;
		private $c ;
		public static $d ;
		
	}`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintReturn(t *testing.T) {
	src := `<?php
	class Foo {
		function bar ( )
		{
			return null ;
		}
	}

	function foo ( )
	{
		return $a ;
	}

	function bar ( )
	{
		return ;
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

func TestParseAndPrintStmtList(t *testing.T) {
	src := `<?php
	{
		;
	}
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintSwitch(t *testing.T) {
	src := `<?php

	switch ( $a ) {
		case 1 : ;
		default : ;
	}
	switch ( $a ) { ;
		case 1 ; ;
		default ; ;
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
		use \ foo , bar ;
		use foo , \ bar { }
		use \ foo , \ bar {
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

func TestParseAndPrintUnset(t *testing.T) {
	src := `<?php
	unset ( $a ) ;
	unset ( $a , $b , ) ;`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintUseList(t *testing.T) {
	src := `<?php
	use Foo ;
	use \ Foo as Bar ;
	use function \ Foo as Bar ;
	use const Foo as Bar, baz ;
	`

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

// other

func TestParseAndPrintParentheses(t *testing.T) {
	src := `<?php
	$b = (($a));
	$b = ( ($a) );
	$b = ( ( $a ) );
	$b = ( ($a ));
	$b = (( $a) );

	( $a + $b ) * 2 ;
	( $ $foo . 'foo' ) :: { $bar . 'bar' } ( ) ;
	( $ $foo ) [ 'bar' ] ;
	$ { $a . 'b' } -> call ( ) ;
	$a -> { $b . 'b' } ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintComplexString1(t *testing.T) {
	src := `<?php
	// "test $foo" ;
	"test $foo[1]" ;
	"test $foo[-1]" ;
	"test $foo[112345678901234567890] " ;
	"test $foo[-112345678901234567890] " ;
	"test $foo[a]" ;
	"test $foo[$bar]" ;
	"test $foo->bar" ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintComplexString2(t *testing.T) {
	src := `<?php
	"test ${ foo }" ;
	"test ${ foo . 'bar' }" ;
	"test ${ foo [ ] }" ;
	"test ${ foo [ $b ] }" ;
	"test ${ foo [ 1 ] }" ;
	"test ${ foo [ 'expr' . $bar ] }" ;
	"test ${ $foo }" ;
	"test ${ $foo -> bar }" ;
	"test ${ $foo -> bar ( ) }" ;
	"test ${ $a . '' }" ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintComplexString3(t *testing.T) {
	src := `<?php
	"test ${foo}" ;
	"test ${foo[0]}";
	"test ${foo::$bar}";
	"test ${foo }" ;
	"test ${foo . 'bar' }" ;
	"test ${foo [ ] }" ;
	"test ${foo [ $b ] }" ;
	"test ${foo [ 1 ] }" ;
	"test ${foo [ 'expr' . $bar ] }" ;
	"test ${$foo }" ;
	"test ${$foo -> bar }" ;
	"test ${$foo -> bar ( ) }" ;
	"test ${$a . '' }" ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintComplexString4(t *testing.T) {
	src := `<?php
	"test {$foo }" ;
	"test {$foo [ ] }" ;
	"test {$foo [ 1 ] }" ;
	"test {$foo -> bar }" ;
	"test {$foo -> bar ( ) }" ;
	`

	actual := print(parse(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}
