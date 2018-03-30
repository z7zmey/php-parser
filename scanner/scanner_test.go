package scanner_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/scanner"
	"github.com/z7zmey/php-parser/token"

	"github.com/kylelemons/godebug/pretty"
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

type lval struct {
	Tkn token.Token
}

func (lv *lval) Token(t token.Token) {
	lv.Tkn = t
}

func TestTokens(t *testing.T) {
	src := `inline html - 
		<? ?>
		<?= ?>
		<?php
		
		0.1
		.1
		1e10
		.1e10

		0b1
		0b1111111111111111111111111111111111111111111111111111111111111111

		0x007111111111111111
		0x8111111111111111

		1234567890123456789
		12345678901234567890

		abstract
		array
		as
		break
		callable
		case
		catch
		class
		clone
		const
		continue
		declare
		default
		do
		echo
		else
		elseif
		empty
		enddeclare
		endfor
		endforeach
		endif
		endswitch
		endwhile
		eval
		exit
		extends
		final
		finally
		for
		foreach
		function
		cfunction
		global
		goto
		if
		isset
		implements
		instanceof
		insteadof
		interface
		list
		namespace
		private
		public
		print
		protected
		return
		static
		switch
		throw
		trait
		try
		unset
		use
		var
		while
		yield ` + "\t\r\n" + ` from
		yield
		include
		include_once
		require
		require_once

		__CLASS__
		__DIR__
		__FILE__
		__FUNCTION__
		__LINE__
		__NAMESPACE__
		__METHOD__
		__TRAIT__
		__halt_compiler

		new
		and
		or
		xor

		\
		...
		::
		&&
		||
		&=
		|=
		.=
		*=
		**=
		/=
		+=
		-=
		^=
		%=
		--
		++
		=>
		<=>
		!=
		<>
		!==
		==
		===
		<<=
		>>=
		>=
		<=
		**
		<<
		>>
		??

		#  inline comment
		// inline comment

		/*
			multiline comment
		*/

		/**
		 * PHP Doc comment
		 */

		;
		:
		,
		.
		[
		]
		(
		)
		|
		/
		^
		&
		+
		-
		*
		=
		%
		!
		~
		$
		<
		>
		?
		@
		{
		}

		$var
		str

		-> ` + "\t\r\n" + ` ->prop

		'adsf\'adsf\''

		"test"
		b"\$var $4 {a"

		( array )
		( bool )
		( boolean )
		( real )
		( double )
		( float )
		( int )
		( integer )
		( object )
		( string )
		( unset )

	`

	expected := []int{
		scanner.T_INLINE_HTML,
		scanner.Rune2Class(';'),
		scanner.T_ECHO,
		scanner.Rune2Class(';'),

		scanner.T_DNUMBER,
		scanner.T_DNUMBER,
		scanner.T_DNUMBER,
		scanner.T_DNUMBER,

		scanner.T_LNUMBER,
		scanner.T_DNUMBER,

		scanner.T_LNUMBER,
		scanner.T_DNUMBER,

		scanner.T_LNUMBER,
		scanner.T_DNUMBER,

		scanner.T_ABSTRACT,
		scanner.T_ARRAY,
		scanner.T_AS,
		scanner.T_BREAK,
		scanner.T_CALLABLE,
		scanner.T_CASE,
		scanner.T_CATCH,
		scanner.T_CLASS,
		scanner.T_CLONE,
		scanner.T_CONST,
		scanner.T_CONTINUE,
		scanner.T_DECLARE,
		scanner.T_DEFAULT,
		scanner.T_DO,
		scanner.T_ECHO,
		scanner.T_ELSE,
		scanner.T_ELSEIF,
		scanner.T_EMPTY,
		scanner.T_ENDDECLARE,
		scanner.T_ENDFOR,
		scanner.T_ENDFOREACH,
		scanner.T_ENDIF,
		scanner.T_ENDSWITCH,
		scanner.T_ENDWHILE,
		scanner.T_EVAL,
		scanner.T_EXIT,
		scanner.T_EXTENDS,
		scanner.T_FINAL,
		scanner.T_FINALLY,
		scanner.T_FOR,
		scanner.T_FOREACH,
		scanner.T_FUNCTION,
		scanner.T_FUNCTION,
		scanner.T_GLOBAL,
		scanner.T_GOTO,
		scanner.T_IF,
		scanner.T_ISSET,
		scanner.T_IMPLEMENTS,
		scanner.T_INSTANCEOF,
		scanner.T_INSTEADOF,
		scanner.T_INTERFACE,
		scanner.T_LIST,
		scanner.T_NAMESPACE,
		scanner.T_PRIVATE,
		scanner.T_PUBLIC,
		scanner.T_PRINT,
		scanner.T_PROTECTED,
		scanner.T_RETURN,
		scanner.T_STATIC,
		scanner.T_SWITCH,
		scanner.T_THROW,
		scanner.T_TRAIT,
		scanner.T_TRY,
		scanner.T_UNSET,
		scanner.T_USE,
		scanner.T_VAR,
		scanner.T_WHILE,
		scanner.T_YIELD_FROM,
		scanner.T_YIELD,
		scanner.T_INCLUDE,
		scanner.T_INCLUDE_ONCE,
		scanner.T_REQUIRE,
		scanner.T_REQUIRE_ONCE,

		scanner.T_CLASS_C,
		scanner.T_DIR,
		scanner.T_FILE,
		scanner.T_FUNC_C,
		scanner.T_LINE,
		scanner.T_NS_C,
		scanner.T_METHOD_C,
		scanner.T_TRAIT_C,
		scanner.T_HALT_COMPILER,

		scanner.T_NEW,
		scanner.T_LOGICAL_AND,
		scanner.T_LOGICAL_OR,
		scanner.T_LOGICAL_XOR,

		scanner.T_NS_SEPARATOR,
		scanner.T_ELLIPSIS,
		scanner.T_PAAMAYIM_NEKUDOTAYIM,
		scanner.T_BOOLEAN_AND,
		scanner.T_BOOLEAN_OR,
		scanner.T_AND_EQUAL,
		scanner.T_OR_EQUAL,
		scanner.T_CONCAT_EQUAL,
		scanner.T_MUL_EQUAL,
		scanner.T_POW_EQUAL,
		scanner.T_DIV_EQUAL,
		scanner.T_PLUS_EQUAL,
		scanner.T_MINUS_EQUAL,
		scanner.T_XOR_EQUAL,
		scanner.T_MOD_EQUAL,
		scanner.T_DEC,
		scanner.T_INC,
		scanner.T_DOUBLE_ARROW,
		scanner.T_SPACESHIP,
		scanner.T_IS_NOT_EQUAL,
		scanner.T_IS_NOT_EQUAL,
		scanner.T_IS_NOT_IDENTICAL,
		scanner.T_IS_EQUAL,
		scanner.T_IS_IDENTICAL,
		scanner.T_SL_EQUAL,
		scanner.T_SR_EQUAL,
		scanner.T_IS_GREATER_OR_EQUAL,
		scanner.T_IS_SMALLER_OR_EQUAL,
		scanner.T_POW,
		scanner.T_SL,
		scanner.T_SR,
		scanner.T_COALESCE,

		scanner.Rune2Class(';'),
		scanner.Rune2Class(':'),
		scanner.Rune2Class(','),
		scanner.Rune2Class('.'),
		scanner.Rune2Class('['),
		scanner.Rune2Class(']'),
		scanner.Rune2Class('('),
		scanner.Rune2Class(')'),
		scanner.Rune2Class('|'),
		scanner.Rune2Class('/'),
		scanner.Rune2Class('^'),
		scanner.Rune2Class('&'),
		scanner.Rune2Class('+'),
		scanner.Rune2Class('-'),
		scanner.Rune2Class('*'),
		scanner.Rune2Class('='),
		scanner.Rune2Class('%'),
		scanner.Rune2Class('!'),
		scanner.Rune2Class('~'),
		scanner.Rune2Class('$'),
		scanner.Rune2Class('<'),
		scanner.Rune2Class('>'),
		scanner.Rune2Class('?'),
		scanner.Rune2Class('@'),
		scanner.Rune2Class('{'),
		scanner.Rune2Class('}'),

		scanner.T_VARIABLE,
		scanner.T_STRING,

		scanner.T_OBJECT_OPERATOR,
		scanner.T_OBJECT_OPERATOR,
		scanner.T_STRING,

		scanner.T_CONSTANT_ENCAPSED_STRING,
		scanner.T_CONSTANT_ENCAPSED_STRING,
		scanner.T_CONSTANT_ENCAPSED_STRING,

		scanner.T_ARRAY_CAST,
		scanner.T_BOOL_CAST,
		scanner.T_BOOL_CAST,
		scanner.T_DOUBLE_CAST,
		scanner.T_DOUBLE_CAST,
		scanner.T_DOUBLE_CAST,
		scanner.T_INT_CAST,
		scanner.T_INT_CAST,
		scanner.T_OBJECT_CAST,
		scanner.T_STRING_CAST,
		scanner.T_UNSET_CAST,
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lv := &lval{}
	actual := []int{}

	for {
		token := lexer.Lex(lv)
		if token < 0 {
			break
		}

		actual = append(actual, token)
	}

	assertEqual(t, expected, actual)
}

func TestTeplateStringTokens(t *testing.T) {
	src := `<?php
		` + "`test $var {$var} ${var_name} {s $ \\$a `" + `

		"test $var {$var} ${var_name} {s $ \$a "
		
		"{$var}"
	`

	expected := []int{
		scanner.Rune2Class('`'),
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_VARIABLE,
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_CURLY_OPEN,
		scanner.T_VARIABLE,
		scanner.Rune2Class('}'),
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_DOLLAR_OPEN_CURLY_BRACES,
		scanner.T_STRING_VARNAME,
		scanner.Rune2Class('}'),
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.Rune2Class('`'),

		scanner.Rune2Class('"'),
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_VARIABLE,
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_CURLY_OPEN,
		scanner.T_VARIABLE,
		scanner.Rune2Class('}'),
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_DOLLAR_OPEN_CURLY_BRACES,
		scanner.T_STRING_VARNAME,
		scanner.Rune2Class('}'),
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.Rune2Class('"'),

		scanner.Rune2Class('"'),
		scanner.T_CURLY_OPEN,
		scanner.T_VARIABLE,
		scanner.Rune2Class('}'),
		scanner.Rune2Class('"'),
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lv := &lval{}
	actual := []int{}

	for {
		token := lexer.Lex(lv)
		if token < 0 {
			break
		}

		actual = append(actual, token)
	}

	assertEqual(t, expected, actual)
}

func TestHereDocTokens(t *testing.T) {
	src := `<?php
	<<<CAT
	test
CAT;

	<<<'CAT'
	test
CAT;

	<<<"CAT"
	$var->prop
	$var[1]
	$var[0x1]
	$var[0b1]
	$var[var_name]
	$var[$var]

	{$var}
	${var_name}
	{s $ \$a 
CAT;
	`

	expected := []int{
		scanner.T_START_HEREDOC,
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_END_HEREDOC,
		scanner.Rune2Class(';'),

		scanner.T_START_HEREDOC,
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_END_HEREDOC,
		scanner.Rune2Class(';'),

		scanner.T_START_HEREDOC,
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_VARIABLE,
		scanner.T_OBJECT_OPERATOR,
		scanner.T_STRING,
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_VARIABLE,
		scanner.Rune2Class('['),
		scanner.T_NUM_STRING,
		scanner.Rune2Class(']'),
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_VARIABLE,
		scanner.Rune2Class('['),
		scanner.T_NUM_STRING,
		scanner.Rune2Class(']'),
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_VARIABLE,
		scanner.Rune2Class('['),
		scanner.T_NUM_STRING,
		scanner.Rune2Class(']'),
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_VARIABLE,
		scanner.Rune2Class('['),
		scanner.T_STRING,
		scanner.Rune2Class(']'),
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_VARIABLE,
		scanner.Rune2Class('['),
		scanner.T_VARIABLE,
		scanner.Rune2Class(']'),
		scanner.T_ENCAPSED_AND_WHITESPACE, scanner.T_CURLY_OPEN,
		scanner.T_VARIABLE,
		scanner.Rune2Class('}'),
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_DOLLAR_OPEN_CURLY_BRACES,
		scanner.T_STRING_VARNAME,
		scanner.Rune2Class('}'),
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_END_HEREDOC,
		scanner.Rune2Class(';'),
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lv := &lval{}
	actual := []int{}

	for {
		token := lexer.Lex(lv)
		if token < 0 {
			break
		}

		actual = append(actual, token)
	}

	assertEqual(t, expected, actual)
}

func TestStringTokensAfterVariable(t *testing.T) {
	src := `<?php "test \"$var\""`

	expected := []int{
		scanner.Rune2Class('"'),
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.T_VARIABLE,
		scanner.T_ENCAPSED_AND_WHITESPACE,
		scanner.Rune2Class('"'),
	}

	expectedTokens := []string{
		"\"",
		"test \\\"",
		"$var",
		"\\\"",
		"\"",
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lv := &lval{}
	actual := []int{}
	actualTokens := []string{}

	for {
		token := lexer.Lex(lv)
		if token < 0 {
			break
		}

		actualTokens = append(actualTokens, lv.Tkn.Value)
		actual = append(actual, token)
	}

	assertEqual(t, expected, actual)
	assertEqual(t, expectedTokens, actualTokens)
}
