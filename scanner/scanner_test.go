package scanner_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/scanner"
	"gotest.tools/assert"
)

type lval struct {
	Tkn *scanner.Token
}

func (lv *lval) Token(t *scanner.Token) {
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
		( binary )
		( unset )

	`

	expected := []string{
		scanner.T_INLINE_HTML.String(),
		scanner.LexerToken(scanner.Rune2Class(';')).String(),
		scanner.T_ECHO.String(),
		scanner.LexerToken(scanner.Rune2Class(';')).String(),

		scanner.T_DNUMBER.String(),
		scanner.T_DNUMBER.String(),
		scanner.T_DNUMBER.String(),
		scanner.T_DNUMBER.String(),

		scanner.T_LNUMBER.String(),
		scanner.T_DNUMBER.String(),

		scanner.T_LNUMBER.String(),
		scanner.T_DNUMBER.String(),

		scanner.T_LNUMBER.String(),
		scanner.T_DNUMBER.String(),

		scanner.T_ABSTRACT.String(),
		scanner.T_ARRAY.String(),
		scanner.T_AS.String(),
		scanner.T_BREAK.String(),
		scanner.T_CALLABLE.String(),
		scanner.T_CASE.String(),
		scanner.T_CATCH.String(),
		scanner.T_CLASS.String(),
		scanner.T_CLONE.String(),
		scanner.T_CONST.String(),
		scanner.T_CONTINUE.String(),
		scanner.T_DECLARE.String(),
		scanner.T_DEFAULT.String(),
		scanner.T_DO.String(),
		scanner.T_ECHO.String(),
		scanner.T_ELSE.String(),
		scanner.T_ELSEIF.String(),
		scanner.T_EMPTY.String(),
		scanner.T_ENDDECLARE.String(),
		scanner.T_ENDFOR.String(),
		scanner.T_ENDFOREACH.String(),
		scanner.T_ENDIF.String(),
		scanner.T_ENDSWITCH.String(),
		scanner.T_ENDWHILE.String(),
		scanner.T_EVAL.String(),
		scanner.T_EXIT.String(),
		scanner.T_EXTENDS.String(),
		scanner.T_FINAL.String(),
		scanner.T_FINALLY.String(),
		scanner.T_FOR.String(),
		scanner.T_FOREACH.String(),
		scanner.T_FUNCTION.String(),
		scanner.T_FUNCTION.String(),
		scanner.T_GLOBAL.String(),
		scanner.T_GOTO.String(),
		scanner.T_IF.String(),
		scanner.T_ISSET.String(),
		scanner.T_IMPLEMENTS.String(),
		scanner.T_INSTANCEOF.String(),
		scanner.T_INSTEADOF.String(),
		scanner.T_INTERFACE.String(),
		scanner.T_LIST.String(),
		scanner.T_NAMESPACE.String(),
		scanner.T_PRIVATE.String(),
		scanner.T_PUBLIC.String(),
		scanner.T_PRINT.String(),
		scanner.T_PROTECTED.String(),
		scanner.T_RETURN.String(),
		scanner.T_STATIC.String(),
		scanner.T_SWITCH.String(),
		scanner.T_THROW.String(),
		scanner.T_TRAIT.String(),
		scanner.T_TRY.String(),
		scanner.T_UNSET.String(),
		scanner.T_USE.String(),
		scanner.T_VAR.String(),
		scanner.T_WHILE.String(),
		scanner.T_YIELD_FROM.String(),
		scanner.T_YIELD.String(),
		scanner.T_INCLUDE.String(),
		scanner.T_INCLUDE_ONCE.String(),
		scanner.T_REQUIRE.String(),
		scanner.T_REQUIRE_ONCE.String(),

		scanner.T_CLASS_C.String(),
		scanner.T_DIR.String(),
		scanner.T_FILE.String(),
		scanner.T_FUNC_C.String(),
		scanner.T_LINE.String(),
		scanner.T_NS_C.String(),
		scanner.T_METHOD_C.String(),
		scanner.T_TRAIT_C.String(),
		scanner.T_HALT_COMPILER.String(),

		scanner.T_NEW.String(),
		scanner.T_LOGICAL_AND.String(),
		scanner.T_LOGICAL_OR.String(),
		scanner.T_LOGICAL_XOR.String(),

		scanner.T_NS_SEPARATOR.String(),
		scanner.T_ELLIPSIS.String(),
		scanner.T_PAAMAYIM_NEKUDOTAYIM.String(),
		scanner.T_BOOLEAN_AND.String(),
		scanner.T_BOOLEAN_OR.String(),
		scanner.T_AND_EQUAL.String(),
		scanner.T_OR_EQUAL.String(),
		scanner.T_CONCAT_EQUAL.String(),
		scanner.T_MUL_EQUAL.String(),
		scanner.T_POW_EQUAL.String(),
		scanner.T_DIV_EQUAL.String(),
		scanner.T_PLUS_EQUAL.String(),
		scanner.T_MINUS_EQUAL.String(),
		scanner.T_XOR_EQUAL.String(),
		scanner.T_MOD_EQUAL.String(),
		scanner.T_DEC.String(),
		scanner.T_INC.String(),
		scanner.T_DOUBLE_ARROW.String(),
		scanner.T_SPACESHIP.String(),
		scanner.T_IS_NOT_EQUAL.String(),
		scanner.T_IS_NOT_EQUAL.String(),
		scanner.T_IS_NOT_IDENTICAL.String(),
		scanner.T_IS_EQUAL.String(),
		scanner.T_IS_IDENTICAL.String(),
		scanner.T_SL_EQUAL.String(),
		scanner.T_SR_EQUAL.String(),
		scanner.T_IS_GREATER_OR_EQUAL.String(),
		scanner.T_IS_SMALLER_OR_EQUAL.String(),
		scanner.T_POW.String(),
		scanner.T_SL.String(),
		scanner.T_SR.String(),
		scanner.T_COALESCE.String(),

		scanner.LexerToken(scanner.Rune2Class(';')).String(),
		scanner.LexerToken(scanner.Rune2Class(':')).String(),
		scanner.LexerToken(scanner.Rune2Class(',')).String(),
		scanner.LexerToken(scanner.Rune2Class('.')).String(),
		scanner.LexerToken(scanner.Rune2Class('[')).String(),
		scanner.LexerToken(scanner.Rune2Class(']')).String(),
		scanner.LexerToken(scanner.Rune2Class('(')).String(),
		scanner.LexerToken(scanner.Rune2Class(')')).String(),
		scanner.LexerToken(scanner.Rune2Class('|')).String(),
		scanner.LexerToken(scanner.Rune2Class('/')).String(),
		scanner.LexerToken(scanner.Rune2Class('^')).String(),
		scanner.LexerToken(scanner.Rune2Class('&')).String(),
		scanner.LexerToken(scanner.Rune2Class('+')).String(),
		scanner.LexerToken(scanner.Rune2Class('-')).String(),
		scanner.LexerToken(scanner.Rune2Class('*')).String(),
		scanner.LexerToken(scanner.Rune2Class('=')).String(),
		scanner.LexerToken(scanner.Rune2Class('%')).String(),
		scanner.LexerToken(scanner.Rune2Class('!')).String(),
		scanner.LexerToken(scanner.Rune2Class('~')).String(),
		scanner.LexerToken(scanner.Rune2Class('$')).String(),
		scanner.LexerToken(scanner.Rune2Class('<')).String(),
		scanner.LexerToken(scanner.Rune2Class('>')).String(),
		scanner.LexerToken(scanner.Rune2Class('?')).String(),
		scanner.LexerToken(scanner.Rune2Class('@')).String(),
		scanner.LexerToken(scanner.Rune2Class('{')).String(),
		scanner.LexerToken(scanner.Rune2Class('}')).String(),

		scanner.T_VARIABLE.String(),
		scanner.T_STRING.String(),

		scanner.T_OBJECT_OPERATOR.String(),
		scanner.T_OBJECT_OPERATOR.String(),
		scanner.T_STRING.String(),

		scanner.T_CONSTANT_ENCAPSED_STRING.String(),
		scanner.T_CONSTANT_ENCAPSED_STRING.String(),
		scanner.T_CONSTANT_ENCAPSED_STRING.String(),

		scanner.T_ARRAY_CAST.String(),
		scanner.T_BOOL_CAST.String(),
		scanner.T_BOOL_CAST.String(),
		scanner.T_DOUBLE_CAST.String(),
		scanner.T_DOUBLE_CAST.String(),
		scanner.T_DOUBLE_CAST.String(),
		scanner.T_INT_CAST.String(),
		scanner.T_INT_CAST.String(),
		scanner.T_OBJECT_CAST.String(),
		scanner.T_STRING_CAST.String(),
		scanner.T_STRING_CAST.String(),
		scanner.T_UNSET_CAST.String(),
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token < 0 {
			break
		}

		actual = append(actual, scanner.LexerToken(token).String())
	}

	assert.DeepEqual(t, expected, actual)
}

func TestSingleQuoteStringTokens(t *testing.T) {
	src := `<?php
		'str $var str'
		
		'\''
		
		'\'
		'
		
		'\
		\''
		
		'\\'
		
		'\\
		'

		'\	
		\''
	`

	expected := []string{
		scanner.T_CONSTANT_ENCAPSED_STRING.String(),
		scanner.T_CONSTANT_ENCAPSED_STRING.String(),
		scanner.T_CONSTANT_ENCAPSED_STRING.String(),
		scanner.T_CONSTANT_ENCAPSED_STRING.String(),
		scanner.T_CONSTANT_ENCAPSED_STRING.String(),
		scanner.T_CONSTANT_ENCAPSED_STRING.String(),
		scanner.T_CONSTANT_ENCAPSED_STRING.String(),
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token < 0 {
			break
		}

		actual = append(actual, scanner.LexerToken(token).String())
	}

	assert.DeepEqual(t, expected, actual)
}

func TestTeplateStringTokens(t *testing.T) {
	src := `<?php
		"foo $a"

		"foo $a{$b}"

		"test $var {$var} ${var_name} {s $ \$a "
		
		"{$var}"
		
		"$foo/"
		"$foo/100;"

		"$/$foo"
		"$0$foo"
	`

	expected := []string{
		scanner.LexerToken(scanner.Rune2Class('"')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('"')).String(),

		scanner.LexerToken(scanner.Rune2Class('"')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.T_CURLY_OPEN.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('}')).String(),
		scanner.LexerToken(scanner.Rune2Class('"')).String(),

		scanner.LexerToken(scanner.Rune2Class('"')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_CURLY_OPEN.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('}')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_DOLLAR_OPEN_CURLY_BRACES.String(),
		scanner.T_STRING_VARNAME.String(),
		scanner.LexerToken(scanner.Rune2Class('}')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.LexerToken(scanner.Rune2Class('"')).String(),

		scanner.LexerToken(scanner.Rune2Class('"')).String(),
		scanner.T_CURLY_OPEN.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('}')).String(),
		scanner.LexerToken(scanner.Rune2Class('"')).String(),

		scanner.LexerToken(scanner.Rune2Class('"')).String(),
		scanner.T_VARIABLE.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.LexerToken(scanner.Rune2Class('"')).String(),

		scanner.LexerToken(scanner.Rune2Class('"')).String(),
		scanner.T_VARIABLE.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.LexerToken(scanner.Rune2Class('"')).String(),

		scanner.LexerToken(scanner.Rune2Class('"')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('"')).String(),

		scanner.LexerToken(scanner.Rune2Class('"')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('"')).String(),
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token < 0 {
			break
		}

		actual = append(actual, scanner.LexerToken(token).String())
	}

	assert.DeepEqual(t, expected, actual)
}

func TestBackquoteStringTokens(t *testing.T) {
	src := `<?php
		` + "`foo $a`" + `
		` + "`foo $a{$b}`" + `

		` + "`test $var {$var} ${var_name} {s $ \\$a `" + `
		
		` + "`{$var}`" + `
		` + "`$foo/`" + `
		` + "`$foo/100`" + `
		` + "`$/$foo`" + `
		` + "`$0$foo`" + `
	`

	expected := []string{
		scanner.LexerToken(scanner.Rune2Class('`')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('`')).String(),

		scanner.LexerToken(scanner.Rune2Class('`')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.T_CURLY_OPEN.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('}')).String(),
		scanner.LexerToken(scanner.Rune2Class('`')).String(),

		scanner.LexerToken(scanner.Rune2Class('`')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_CURLY_OPEN.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('}')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_DOLLAR_OPEN_CURLY_BRACES.String(),
		scanner.T_STRING_VARNAME.String(),
		scanner.LexerToken(scanner.Rune2Class('}')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.LexerToken(scanner.Rune2Class('`')).String(),

		scanner.LexerToken(scanner.Rune2Class('`')).String(),
		scanner.T_CURLY_OPEN.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('}')).String(),
		scanner.LexerToken(scanner.Rune2Class('`')).String(),

		scanner.LexerToken(scanner.Rune2Class('`')).String(),
		scanner.T_VARIABLE.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.LexerToken(scanner.Rune2Class('`')).String(),

		scanner.LexerToken(scanner.Rune2Class('`')).String(),
		scanner.T_VARIABLE.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.LexerToken(scanner.Rune2Class('`')).String(),

		scanner.LexerToken(scanner.Rune2Class('`')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('`')).String(),

		scanner.LexerToken(scanner.Rune2Class('`')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('`')).String(),
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token < 0 {
			break
		}

		actual = append(actual, scanner.LexerToken(token).String())
	}

	assert.DeepEqual(t, expected, actual)
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

	expected := []string{
		scanner.T_START_HEREDOC.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_END_HEREDOC.String(),
		scanner.LexerToken(scanner.Rune2Class(';')).String(),

		scanner.T_START_HEREDOC.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_END_HEREDOC.String(),
		scanner.LexerToken(scanner.Rune2Class(';')).String(),

		scanner.T_START_HEREDOC.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.T_OBJECT_OPERATOR.String(),
		scanner.T_STRING.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('[')).String(),
		scanner.T_NUM_STRING.String(),
		scanner.LexerToken(scanner.Rune2Class(']')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('[')).String(),
		scanner.T_NUM_STRING.String(),
		scanner.LexerToken(scanner.Rune2Class(']')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('[')).String(),
		scanner.T_NUM_STRING.String(),
		scanner.LexerToken(scanner.Rune2Class(']')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('[')).String(),
		scanner.T_STRING.String(),
		scanner.LexerToken(scanner.Rune2Class(']')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('[')).String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class(']')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_CURLY_OPEN.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('}')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_DOLLAR_OPEN_CURLY_BRACES.String(),
		scanner.T_STRING_VARNAME.String(),
		scanner.LexerToken(scanner.Rune2Class('}')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_END_HEREDOC.String(),
		scanner.LexerToken(scanner.Rune2Class(';')).String(),
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token < 0 {
			break
		}

		actual = append(actual, scanner.LexerToken(token).String())
	}

	assert.DeepEqual(t, expected, actual)
}

func TestHereDocTokens2(t *testing.T) {
	src := `<?php
	<<<CAT
$foo/
CAT;

	<<<CAT
$foo/100
CAT;

	<<<CAT
$/$foo
CAT;

	<<<CAT
$0$foo
CAT;

	<<<CAT
$foo$bar\
CAT
`

	expected := []string{
		scanner.T_START_HEREDOC.String(),
		scanner.T_VARIABLE.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_END_HEREDOC.String(),
		scanner.LexerToken(scanner.Rune2Class(';')).String(),

		scanner.T_START_HEREDOC.String(),
		scanner.T_VARIABLE.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_END_HEREDOC.String(),
		scanner.LexerToken(scanner.Rune2Class(';')).String(),

		scanner.T_START_HEREDOC.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.T_END_HEREDOC.String(),
		scanner.LexerToken(scanner.Rune2Class(';')).String(),

		scanner.T_START_HEREDOC.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.T_END_HEREDOC.String(),
		scanner.LexerToken(scanner.Rune2Class(';')).String(),

		scanner.T_START_HEREDOC.String(),
		scanner.T_VARIABLE.String(),
		scanner.T_VARIABLE.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_END_HEREDOC.String(),
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token < 0 {
			break
		}

		actual = append(actual, scanner.LexerToken(token).String())
	}

	assert.DeepEqual(t, expected, actual)
}

func TestHereDocTokens3(t *testing.T) {
	src := `<?php

	<<<"CAT"
\\{$a['b']}
CAT;
	`

	expected := []string{

		scanner.T_START_HEREDOC.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_CURLY_OPEN.String(),
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('[')).String(),
		scanner.T_CONSTANT_ENCAPSED_STRING.String(),
		scanner.LexerToken(scanner.Rune2Class(']')).String(),
		scanner.LexerToken(scanner.Rune2Class('}')).String(),
		scanner.T_END_HEREDOC.String(),
		scanner.LexerToken(scanner.Rune2Class(';')).String(),
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token < 0 {
			break
		}

		actual = append(actual, scanner.LexerToken(token).String())
	}

	assert.DeepEqual(t, expected, actual)
}

func TestInlineHtmlNopTokens(t *testing.T) {
	src := `<?php
		$a; ?> test <?php
		$a ?> test
	`

	expected := []string{
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class(';')).String(),
		scanner.T_INLINE_HTML.String(),

		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class(';')).String(),
		scanner.T_INLINE_HTML.String(),
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token < 0 {
			break
		}

		actual = append(actual, scanner.LexerToken(token).String())
	}

	assert.DeepEqual(t, expected, actual)
}

func TestStringTokensAfterVariable(t *testing.T) {
	src := `<?php "test \"$var\""`

	expected := []string{
		scanner.LexerToken(scanner.Rune2Class('"')).String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.T_VARIABLE.String(),
		scanner.T_ENCAPSED_AND_WHITESPACE.String(),
		scanner.LexerToken(scanner.Rune2Class('"')).String(),
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
	actual := []string{}
	actualTokens := []string{}

	for {
		token := lexer.Lex(lv)
		if token < 0 {
			break
		}

		actualTokens = append(actualTokens, lv.Tkn.Value)
		actual = append(actual, scanner.LexerToken(token).String())
	}

	assert.DeepEqual(t, expected, actual)
	assert.DeepEqual(t, expectedTokens, actualTokens)
}

func TestSlashAfterVariable(t *testing.T) {
	src := `<?php $foo/3`

	expected := []string{
		scanner.T_VARIABLE.String(),
		scanner.LexerToken(scanner.Rune2Class('/')).String(),
		scanner.T_LNUMBER.String(),
	}

	expectedTokens := []string{
		"$foo",
		"/",
		"3",
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lv := &lval{}
	actual := []string{}
	actualTokens := []string{}

	for {
		token := lexer.Lex(lv)
		if token < 0 {
			break
		}

		actualTokens = append(actualTokens, lv.Tkn.Value)
		actual = append(actual, scanner.LexerToken(token).String())
	}

	assert.DeepEqual(t, expected, actual)
	assert.DeepEqual(t, expectedTokens, actualTokens)
}

func TestCommentEnd(t *testing.T) {
	src := `<?php //test`

	expected := []freefloating.String{
		{
			Value:      "<?php",
			StringType: freefloating.TokenType,
			Position:   position.NewPosition(1, 1, 1, 5),
		},
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 1, 6, 6),
		},
		{
			Value:      "//test",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(1, 1, 7, 12),
		},
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}

	lexer.Lex(lv)

	actual := lexer.FreeFloating

	assert.DeepEqual(t, expected, actual)
}

func TestCommentNewLine(t *testing.T) {
	src := "<?php //test\n$a"

	expected := []freefloating.String{
		{
			Value:      "<?php",
			StringType: freefloating.TokenType,
			Position:   position.NewPosition(1, 1, 1, 5),
		},
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 1, 6, 6),
		},
		{
			Value:      "//test\n",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(1, 1, 7, 13),
		},
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}

	lexer.Lex(lv)

	actual := lv.Tkn.FreeFloating

	assert.DeepEqual(t, expected, actual)
}

func TestCommentNewLine1(t *testing.T) {
	src := "<?php //test\r$a"

	expected := []freefloating.String{
		{
			Value:      "<?php",
			StringType: freefloating.TokenType,
			Position:   position.NewPosition(1, 1, 1, 5),
		},
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 1, 6, 6),
		},
		{
			Value:      "//test\r",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(1, 1, 7, 13),
		},
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}

	lexer.Lex(lv)

	actual := lv.Tkn.FreeFloating

	assert.DeepEqual(t, expected, actual)
}

func TestCommentNewLine2(t *testing.T) {
	src := "<?php #test\r\n$a"

	expected := []freefloating.String{
		{
			Value:      "<?php",
			StringType: freefloating.TokenType,
			Position:   position.NewPosition(1, 1, 1, 5),
		},
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 1, 6, 6),
		},
		{
			Value:      "#test\r\n",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(1, 1, 7, 13),
		},
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}

	lexer.Lex(lv)

	actual := lv.Tkn.FreeFloating

	assert.DeepEqual(t, expected, actual)
}

func TestCommentWithPhpEndTag(t *testing.T) {
	src := `<?php
	//test?> test`

	expected := []freefloating.String{
		{
			Value:      "<?php",
			StringType: freefloating.TokenType,
			Position:   position.NewPosition(1, 1, 1, 5),
		},
		{
			Value:      "\n\t",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 2, 6, 7),
		},
		{
			Value:      "//test",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(2, 2, 8, 13),
		},
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}

	lexer.Lex(lv)

	actual := lv.Tkn.FreeFloating

	assert.DeepEqual(t, expected, actual)
}

func TestInlineComment(t *testing.T) {
	src := `<?php
	/*test*/`

	expected := []freefloating.String{
		{
			Value:      "<?php",
			StringType: freefloating.TokenType,
			Position:   position.NewPosition(1, 1, 1, 5),
		},
		{
			Value:      "\n\t",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 2, 6, 7),
		},
		{
			Value:      "/*test*/",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(2, 2, 8, 15),
		},
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}

	lexer.Lex(lv)

	actual := lv.Tkn.FreeFloating

	assert.DeepEqual(t, expected, actual)
}

func TestInlineComment2(t *testing.T) {
	src := `<?php
	/*/*/`

	expected := []freefloating.String{
		{
			Value:      "<?php",
			StringType: freefloating.TokenType,
			Position:   position.NewPosition(1, 1, 1, 5),
		},
		{
			Value:      "\n\t",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 2, 6, 7),
		},
		{
			Value:      "/*/*/",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(2, 2, 8, 12),
		},
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}

	lexer.Lex(lv)

	actual := lexer.FreeFloating

	assert.DeepEqual(t, expected, actual)
}

func TestEmptyInlineComment(t *testing.T) {
	src := `<?php
	/**/ `

	expected := []freefloating.String{
		{
			Value:      "<?php",
			StringType: freefloating.TokenType,
			Position:   position.NewPosition(1, 1, 1, 5),
		},
		{
			Value:      "\n\t",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 2, 6, 7),
		},
		{
			Value:      "/**/",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(2, 2, 8, 11),
		},
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 12, 12),
		},
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}

	lexer.Lex(lv)

	actual := lexer.FreeFloating

	assert.DeepEqual(t, expected, actual)
}

func TestEmptyInlineComment2(t *testing.T) {
	src := `<?php
	/***/`

	expected := []freefloating.String{
		{
			Value:      "<?php",
			StringType: freefloating.TokenType,
			Position:   position.NewPosition(1, 1, 1, 5),
		},
		{
			Value:      "\n\t",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 2, 6, 7),
		},
		{
			Value:      "/***/",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(2, 2, 8, 12),
		},
	}

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}

	lexer.Lex(lv)

	actual := lv.Tkn.FreeFloating

	assert.DeepEqual(t, expected, actual)
}

func TestMethodCallTokens(t *testing.T) {
	src := `<?php
	$a -> bar ( '' ) ;`

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}

	expected := []freefloating.String{
		{
			Value:      "<?php",
			StringType: freefloating.TokenType,
			Position:   position.NewPosition(1, 1, 1, 5),
		},
		{
			Value:      "\n\t",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 2, 6, 7),
		},
	}
	lexer.Lex(lv)
	actual := lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)

	expected = []freefloating.String{
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 10, 10),
		},
	}
	lexer.Lex(lv)
	actual = lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)

	expected = []freefloating.String{
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 13, 13),
		},
	}
	lexer.Lex(lv)
	actual = lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)

	expected = []freefloating.String{
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 17, 17),
		},
	}
	lexer.Lex(lv)
	actual = lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)

	expected = []freefloating.String{
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 19, 19),
		},
	}
	lexer.Lex(lv)
	actual = lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)

	expected = []freefloating.String{
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 22, 22),
		},
	}
	lexer.Lex(lv)
	actual = lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)

	expected = []freefloating.String{
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 24, 24),
		},
	}
	lexer.Lex(lv)
	actual = lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)
}

func TestYieldFromTokens(t *testing.T) {
	src := `<?php
	yield from $a`

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lexer.WithFreeFloating = true
	lv := &lval{}

	expected := []freefloating.String{
		{
			Value:      "<?php",
			StringType: freefloating.TokenType,
			Position:   position.NewPosition(1, 1, 1, 5),
		},
		{
			Value:      "\n\t",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 2, 6, 7),
		},
	}
	lexer.Lex(lv)
	actual := lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)

	expected = []freefloating.String{
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 18, 18),
		},
	}
	lexer.Lex(lv)
	actual = lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)
}

func TestIgnoreControllCharacters(t *testing.T) {
	src := "<?php \004 echo $b;"

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lv := &lval{}

	expected := "echo"
	lexer.Lex(lv)
	actual := lv.Tkn.Value
	assert.DeepEqual(t, expected, actual)

	expected = "$b"
	lexer.Lex(lv)
	actual = lv.Tkn.Value
	assert.DeepEqual(t, expected, actual)
}

func TestIgnoreControllCharactersAtStringVarOffset(t *testing.T) {
	src := "<?php \"$a[test\004]\";"

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lv := &lval{}

	expected := "\""
	lexer.Lex(lv)
	actual := lv.Tkn.Value
	assert.DeepEqual(t, expected, actual)

	expected = "$a"
	lexer.Lex(lv)
	actual = lv.Tkn.Value
	assert.DeepEqual(t, expected, actual)

	expected = "["
	lexer.Lex(lv)
	actual = lv.Tkn.Value
	assert.DeepEqual(t, expected, actual)

	expected = "test"
	lexer.Lex(lv)
	actual = lv.Tkn.Value
	assert.DeepEqual(t, expected, actual)

	expected = "]"
	lexer.Lex(lv)
	actual = lv.Tkn.Value
	assert.DeepEqual(t, expected, actual)
}

func TestBomInMiddleOfFile(t *testing.T) {
	src := "<?php \xEF\xBB\xBF $a;"

	lexer := scanner.NewLexer(bytes.NewBufferString(src), "test.php")
	lv := &lval{}

	lexer.Lex(lv)

	assert.Assert(t, len(lexer.Errors) > 0)
	assert.Assert(t, lexer.Errors[0].String() == "unicode (UTF-8) BOM in middle of file at line 1")
}
