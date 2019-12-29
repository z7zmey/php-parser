package scanner

import (
	"testing"

	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/position"
	"gotest.tools/assert"
)

type lval struct {
	Tkn *Token
}

func (lv *lval) Token(t *Token) {
	lv.Tkn = t
}

func TestTokens(t *testing.T) {
	src := `inline html - 
		<? ?>
		<?= ?>
		<?php

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
		T_INLINE_HTML.String(),
		TokenID(int(';')).String(),
		T_INLINE_HTML.String(),
		T_ECHO.String(),
		TokenID(int(';')).String(),
		T_INLINE_HTML.String(),

		T_ABSTRACT.String(),
		T_ARRAY.String(),
		T_AS.String(),
		T_BREAK.String(),
		T_CALLABLE.String(),
		T_CASE.String(),
		T_CATCH.String(),
		T_CLASS.String(),
		T_CLONE.String(),
		T_CONST.String(),
		T_CONTINUE.String(),
		T_DECLARE.String(),
		T_DEFAULT.String(),
		T_DO.String(),
		T_ECHO.String(),
		T_ELSE.String(),
		T_ELSEIF.String(),
		T_EMPTY.String(),
		T_ENDDECLARE.String(),
		T_ENDFOR.String(),
		T_ENDFOREACH.String(),
		T_ENDIF.String(),
		T_ENDSWITCH.String(),
		T_ENDWHILE.String(),
		T_EVAL.String(),
		T_EXIT.String(),
		T_EXTENDS.String(),
		T_FINAL.String(),
		T_FINALLY.String(),
		T_FOR.String(),
		T_FOREACH.String(),
		T_FUNCTION.String(),
		T_FUNCTION.String(),
		T_GLOBAL.String(),
		T_GOTO.String(),
		T_IF.String(),
		T_ISSET.String(),
		T_IMPLEMENTS.String(),
		T_INSTANCEOF.String(),
		T_INSTEADOF.String(),
		T_INTERFACE.String(),
		T_LIST.String(),
		T_NAMESPACE.String(),
		T_PRIVATE.String(),
		T_PUBLIC.String(),
		T_PRINT.String(),
		T_PROTECTED.String(),
		T_RETURN.String(),
		T_STATIC.String(),
		T_SWITCH.String(),
		T_THROW.String(),
		T_TRAIT.String(),
		T_TRY.String(),
		T_UNSET.String(),
		T_USE.String(),
		T_VAR.String(),
		T_WHILE.String(),
		T_YIELD_FROM.String(),
		T_YIELD.String(),
		T_INCLUDE.String(),
		T_INCLUDE_ONCE.String(),
		T_REQUIRE.String(),
		T_REQUIRE_ONCE.String(),

		T_CLASS_C.String(),
		T_DIR.String(),
		T_FILE.String(),
		T_FUNC_C.String(),
		T_LINE.String(),
		T_NS_C.String(),
		T_METHOD_C.String(),
		T_TRAIT_C.String(),
		T_HALT_COMPILER.String(),

		T_NEW.String(),
		T_LOGICAL_AND.String(),
		T_LOGICAL_OR.String(),
		T_LOGICAL_XOR.String(),

		T_NS_SEPARATOR.String(),
		T_ELLIPSIS.String(),
		T_PAAMAYIM_NEKUDOTAYIM.String(),
		T_BOOLEAN_AND.String(),
		T_BOOLEAN_OR.String(),
		T_AND_EQUAL.String(),
		T_OR_EQUAL.String(),
		T_CONCAT_EQUAL.String(),
		T_MUL_EQUAL.String(),
		T_POW_EQUAL.String(),
		T_DIV_EQUAL.String(),
		T_PLUS_EQUAL.String(),
		T_MINUS_EQUAL.String(),
		T_XOR_EQUAL.String(),
		T_MOD_EQUAL.String(),
		T_DEC.String(),
		T_INC.String(),
		T_DOUBLE_ARROW.String(),
		T_SPACESHIP.String(),
		T_IS_NOT_EQUAL.String(),
		T_IS_NOT_EQUAL.String(),
		T_IS_NOT_IDENTICAL.String(),
		T_IS_EQUAL.String(),
		T_IS_IDENTICAL.String(),
		T_SL_EQUAL.String(),
		T_SR_EQUAL.String(),
		T_IS_GREATER_OR_EQUAL.String(),
		T_IS_SMALLER_OR_EQUAL.String(),
		T_POW.String(),
		T_SL.String(),
		T_SR.String(),
		T_COALESCE.String(),

		TokenID(int(';')).String(),
		TokenID(int(':')).String(),
		TokenID(int(',')).String(),
		TokenID(int('.')).String(),
		TokenID(int('[')).String(),
		TokenID(int(']')).String(),
		TokenID(int('(')).String(),
		TokenID(int(')')).String(),
		TokenID(int('|')).String(),
		TokenID(int('/')).String(),
		TokenID(int('^')).String(),
		TokenID(int('&')).String(),
		TokenID(int('+')).String(),
		TokenID(int('-')).String(),
		TokenID(int('*')).String(),
		TokenID(int('=')).String(),
		TokenID(int('%')).String(),
		TokenID(int('!')).String(),
		TokenID(int('~')).String(),
		TokenID(int('$')).String(),
		TokenID(int('<')).String(),
		TokenID(int('>')).String(),
		TokenID(int('?')).String(),
		TokenID(int('@')).String(),
		TokenID(int('{')).String(),
		TokenID(int('}')).String(),

		T_VARIABLE.String(),
		T_STRING.String(),

		T_OBJECT_OPERATOR.String(),
		T_OBJECT_OPERATOR.String(),
		T_STRING.String(),

		T_ARRAY_CAST.String(),
		T_BOOL_CAST.String(),
		T_BOOL_CAST.String(),
		T_DOUBLE_CAST.String(),
		T_DOUBLE_CAST.String(),
		T_DOUBLE_CAST.String(),
		T_INT_CAST.String(),
		T_INT_CAST.String(),
		T_OBJECT_CAST.String(),
		T_STRING_CAST.String(),
		T_STRING_CAST.String(),
		T_UNSET_CAST.String(),
	}

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token == 0 {
			break
		}

		actual = append(actual, TokenID(token).String())
	}

	assert.DeepEqual(t, expected, actual)
}

func TestShebang(t *testing.T) {
	src := `#!/usr/bin/env php
<?php
0.1
`

	expected := []string{
		"#!/usr/bin/env php\n",
		"<?php",
		"\n",
	}

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	token := lexer.Lex(lv)
	assert.Equal(t, token, int(T_DNUMBER))

	for _, tt := range lv.Tkn.FreeFloating {
		actual = append(actual, tt.Value)
	}

	assert.DeepEqual(t, expected, actual)
}

func TestShebangHtml(t *testing.T) {
	src := `#!/usr/bin/env php
<br/><?php
0.1
`

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}

	token := lexer.Lex(lv)
	assert.Equal(t, token, int(T_INLINE_HTML))
	assert.Equal(t, lv.Tkn.FreeFloating[0].Value, "#!/usr/bin/env php\n")

	token = lexer.Lex(lv)
	assert.Equal(t, token, int(T_DNUMBER))
}

func TestNumberTokens(t *testing.T) {
	src := `<?php
		0.1
		.1
		1e10
		.1e10

		0b01111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111
		0b10111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111

		0x0_7FFF_FFFF_FFFF_FFFF
		0x8111_1111_1111_1111

		92233_72036_85477_5807
		0_77777_77777_77777_77777_7

		92233_72036_85477_5808
		0_77777_77777_77777_77777_70
	`

	expected := []string{
		T_DNUMBER.String(),
		T_DNUMBER.String(),
		T_DNUMBER.String(),
		T_DNUMBER.String(),

		T_LNUMBER.String(),
		T_DNUMBER.String(),

		T_LNUMBER.String(),
		T_DNUMBER.String(),

		T_LNUMBER.String(),
		T_LNUMBER.String(),

		T_DNUMBER.String(),
		T_DNUMBER.String(),
	}

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token == 0 {
			break
		}

		actual = append(actual, TokenID(token).String())
	}

	assert.DeepEqual(t, expected, actual)
}

func TestConstantStrings(t *testing.T) {
	src := `<?
		'str'
		'\''
		'\\'

		b"str"
		"\""
		"\\"

		"\$var"
		"$4"
		"$"
		"$\\"

		"{"
		"{a"
		"\{$"
	`

	expected := []string{
		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),

		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),

		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),

		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),
	}

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token == 0 {
			break
		}

		actual = append(actual, TokenID(token).String())
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
		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),
		T_CONSTANT_ENCAPSED_STRING.String(),
	}

	lexer := NewLexer([]byte(src))
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token == 0 {
			break
		}

		actual = append(actual, TokenID(token).String())
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
		TokenID(int('"')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		TokenID(int('"')).String(),

		TokenID(int('"')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		T_CURLY_OPEN.String(),
		T_VARIABLE.String(),
		TokenID(int('}')).String(),
		TokenID(int('"')).String(),

		TokenID(int('"')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_CURLY_OPEN.String(),
		T_VARIABLE.String(),
		TokenID(int('}')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_DOLLAR_OPEN_CURLY_BRACES.String(),
		T_STRING_VARNAME.String(),
		TokenID(int('}')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		TokenID(int('"')).String(),

		TokenID(int('"')).String(),
		T_CURLY_OPEN.String(),
		T_VARIABLE.String(),
		TokenID(int('}')).String(),
		TokenID(int('"')).String(),

		TokenID(int('"')).String(),
		T_VARIABLE.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		TokenID(int('"')).String(),

		TokenID(int('"')).String(),
		T_VARIABLE.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		TokenID(int('"')).String(),

		TokenID(int('"')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		TokenID(int('"')).String(),

		TokenID(int('"')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		TokenID(int('"')).String(),
	}

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token == 0 {
			break
		}

		actual = append(actual, TokenID(token).String())
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
		TokenID(int('`')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		TokenID(int('`')).String(),

		TokenID(int('`')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		T_CURLY_OPEN.String(),
		T_VARIABLE.String(),
		TokenID(int('}')).String(),
		TokenID(int('`')).String(),

		TokenID(int('`')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_CURLY_OPEN.String(),
		T_VARIABLE.String(),
		TokenID(int('}')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_DOLLAR_OPEN_CURLY_BRACES.String(),
		T_STRING_VARNAME.String(),
		TokenID(int('}')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		TokenID(int('`')).String(),

		TokenID(int('`')).String(),
		T_CURLY_OPEN.String(),
		T_VARIABLE.String(),
		TokenID(int('}')).String(),
		TokenID(int('`')).String(),

		TokenID(int('`')).String(),
		T_VARIABLE.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		TokenID(int('`')).String(),

		TokenID(int('`')).String(),
		T_VARIABLE.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		TokenID(int('`')).String(),

		TokenID(int('`')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		TokenID(int('`')).String(),

		TokenID(int('`')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		TokenID(int('`')).String(),
	}

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token == 0 {
			break
		}

		actual = append(actual, TokenID(token).String())
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
		T_START_HEREDOC.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_END_HEREDOC.String(),
		TokenID(int(';')).String(),

		T_START_HEREDOC.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_END_HEREDOC.String(),
		TokenID(int(';')).String(),

		T_START_HEREDOC.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		T_OBJECT_OPERATOR.String(),
		T_STRING.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		TokenID(int('[')).String(),
		T_NUM_STRING.String(),
		TokenID(int(']')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		TokenID(int('[')).String(),
		T_NUM_STRING.String(),
		TokenID(int(']')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		TokenID(int('[')).String(),
		T_NUM_STRING.String(),
		TokenID(int(']')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		TokenID(int('[')).String(),
		T_STRING.String(),
		TokenID(int(']')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		TokenID(int('[')).String(),
		T_VARIABLE.String(),
		TokenID(int(']')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_CURLY_OPEN.String(),
		T_VARIABLE.String(),
		TokenID(int('}')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_DOLLAR_OPEN_CURLY_BRACES.String(),
		T_STRING_VARNAME.String(),
		TokenID(int('}')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_END_HEREDOC.String(),
		TokenID(int(';')).String(),
	}

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token == 0 {
			break
		}

		actual = append(actual, TokenID(token).String())
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
		T_START_HEREDOC.String(),
		T_VARIABLE.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_END_HEREDOC.String(),
		TokenID(int(';')).String(),

		T_START_HEREDOC.String(),
		T_VARIABLE.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_END_HEREDOC.String(),
		TokenID(int(';')).String(),

		T_START_HEREDOC.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_END_HEREDOC.String(),
		TokenID(int(';')).String(),

		T_START_HEREDOC.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_END_HEREDOC.String(),
		TokenID(int(';')).String(),

		T_START_HEREDOC.String(),
		T_VARIABLE.String(),
		T_VARIABLE.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_END_HEREDOC.String(),
	}

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token == 0 {
			break
		}

		actual = append(actual, TokenID(token).String())
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

		T_START_HEREDOC.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_CURLY_OPEN.String(),
		T_VARIABLE.String(),
		TokenID(int('[')).String(),
		T_CONSTANT_ENCAPSED_STRING.String(),
		TokenID(int(']')).String(),
		TokenID(int('}')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_END_HEREDOC.String(),
		TokenID(int(';')).String(),
	}

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token == 0 {
			break
		}

		actual = append(actual, TokenID(token).String())
	}

	assert.DeepEqual(t, expected, actual)
}

func TestHereDocTokens73(t *testing.T) {
	src := `<?php
	<<<"CAT"
		text
	CAT, $b`

	expected := []string{

		T_START_HEREDOC.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_END_HEREDOC.String(),
		TokenID(int(',')).String(),
		T_VARIABLE.String(),
	}

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token == 0 {
			break
		}

		actual = append(actual, TokenID(token).String())
	}

	assert.DeepEqual(t, expected, actual)
}

func TestHereDocTokensBefore73(t *testing.T) {
	src := `<?php
	<<<"CAT"
	CAT
CAT;`

	expected := []string{

		T_START_HEREDOC.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_END_HEREDOC.String(),
		TokenID(int(';')).String(),
	}

	lexer := NewLexer([]byte(src))
	lexer.PHPVersion = "7.2"
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token == 0 {
			break
		}

		actual = append(actual, TokenID(token).String())
	}

	assert.DeepEqual(t, expected, actual)
}

func TestInlineHtmlNopTokens(t *testing.T) {
	src := `<?php
		$a; ?> test <?php
		$a ?> test
	`

	expected := []string{
		T_VARIABLE.String(),
		TokenID(int(';')).String(),
		T_INLINE_HTML.String(),

		T_VARIABLE.String(),
		TokenID(int(';')).String(),
		T_INLINE_HTML.String(),
	}

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}
	actual := []string{}

	for {
		token := lexer.Lex(lv)
		if token == 0 {
			break
		}

		actual = append(actual, TokenID(token).String())
	}

	assert.DeepEqual(t, expected, actual)
}

func TestStringTokensAfterVariable(t *testing.T) {
	src := `<?php "test \"$var\""`

	expected := []string{
		TokenID(int('"')).String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		T_VARIABLE.String(),
		T_ENCAPSED_AND_WHITESPACE.String(),
		TokenID(int('"')).String(),
	}

	expectedTokens := []string{
		"\"",
		"test \\\"",
		"$var",
		"\\\"",
		"\"",
	}

	lexer := NewLexer([]byte(src))
	lv := &lval{}
	actual := []string{}
	actualTokens := []string{}

	for {
		token := lexer.Lex(lv)
		if token == 0 {
			break
		}

		actualTokens = append(actualTokens, lv.Tkn.Value)
		actual = append(actual, TokenID(token).String())
	}

	assert.DeepEqual(t, expected, actual)
	assert.DeepEqual(t, expectedTokens, actualTokens)
}

func TestSlashAfterVariable(t *testing.T) {
	src := `<?php $foo/3`

	expected := []string{
		T_VARIABLE.String(),
		TokenID(int('/')).String(),
		T_LNUMBER.String(),
	}

	expectedTokens := []string{
		"$foo",
		"/",
		"3",
	}

	lexer := NewLexer([]byte(src))
	lv := &lval{}
	actual := []string{}
	actualTokens := []string{}

	for {
		token := lexer.Lex(lv)
		if token == 0 {
			break
		}

		actualTokens = append(actualTokens, lv.Tkn.Value)
		actual = append(actual, TokenID(token).String())
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
			Position:   position.NewPosition(1, 1, 0, 5),
		},
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 1, 5, 6),
		},
		{
			Value:      "//test",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(1, 1, 6, 12),
		},
	}

	lexer := NewLexer([]byte(src))
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
			Position:   position.NewPosition(1, 1, 0, 5),
		},
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 1, 5, 6),
		},
		{
			Value:      "//test\n",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(1, 1, 6, 13),
		},
	}

	lexer := NewLexer([]byte(src))
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
			Position:   position.NewPosition(1, 1, 0, 5),
		},
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 1, 5, 6),
		},
		{
			Value:      "//test\r",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(1, 1, 6, 13),
		},
	}

	lexer := NewLexer([]byte(src))
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
			Position:   position.NewPosition(1, 1, 0, 5),
		},
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 1, 5, 6),
		},
		{
			Value:      "#test\r\n",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(1, 1, 6, 13),
		},
	}

	lexer := NewLexer([]byte(src))
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
			Position:   position.NewPosition(1, 1, 0, 5),
		},
		{
			Value:      "\n\t",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 2, 5, 7),
		},
		{
			Value:      "//test",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(2, 2, 7, 13),
		},
	}

	lexer := NewLexer([]byte(src))
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
			Position:   position.NewPosition(1, 1, 0, 5),
		},
		{
			Value:      "\n\t",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 2, 5, 7),
		},
		{
			Value:      "/*test*/",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(2, 2, 7, 15),
		},
	}

	lexer := NewLexer([]byte(src))
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
			Position:   position.NewPosition(1, 1, 0, 5),
		},
		{
			Value:      "\n\t",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 2, 5, 7),
		},
		{
			Value:      "/*/*/",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(2, 2, 7, 12),
		},
	}

	lexer := NewLexer([]byte(src))
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
			Position:   position.NewPosition(1, 1, 0, 5),
		},
		{
			Value:      "\n\t",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 2, 5, 7),
		},
		{
			Value:      "/**/",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(2, 2, 7, 11),
		},
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 11, 12),
		},
	}

	lexer := NewLexer([]byte(src))
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
			Position:   position.NewPosition(1, 1, 0, 5),
		},
		{
			Value:      "\n\t",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 2, 5, 7),
		},
		{
			Value:      "/***/",
			StringType: freefloating.CommentType,
			Position:   position.NewPosition(2, 2, 7, 12),
		},
	}

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}

	lexer.Lex(lv)

	actual := lv.Tkn.FreeFloating

	assert.DeepEqual(t, expected, actual)
}

func TestMethodCallTokens(t *testing.T) {
	src := `<?php
	$a -> bar ( '' ) ;`

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}

	expected := []freefloating.String{
		{
			Value:      "<?php",
			StringType: freefloating.TokenType,
			Position:   position.NewPosition(1, 1, 0, 5),
		},
		{
			Value:      "\n\t",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 2, 5, 7),
		},
	}
	lexer.Lex(lv)
	actual := lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)

	expected = []freefloating.String{
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 9, 10),
		},
	}
	lexer.Lex(lv)
	actual = lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)

	expected = []freefloating.String{
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 12, 13),
		},
	}
	lexer.Lex(lv)
	actual = lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)

	expected = []freefloating.String{
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 16, 17),
		},
	}
	lexer.Lex(lv)
	actual = lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)

	expected = []freefloating.String{
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 18, 19),
		},
	}
	lexer.Lex(lv)
	actual = lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)

	expected = []freefloating.String{
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 21, 22),
		},
	}
	lexer.Lex(lv)
	actual = lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)

	expected = []freefloating.String{
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 23, 24),
		},
	}
	lexer.Lex(lv)
	actual = lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)
}

func TestYieldFromTokens(t *testing.T) {
	src := `<?php
	yield from $a`

	lexer := NewLexer([]byte(src))
	lexer.WithFreeFloating = true
	lv := &lval{}

	expected := []freefloating.String{
		{
			Value:      "<?php",
			StringType: freefloating.TokenType,
			Position:   position.NewPosition(1, 1, 0, 5),
		},
		{
			Value:      "\n\t",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(1, 2, 5, 7),
		},
	}
	lexer.Lex(lv)
	actual := lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)

	expected = []freefloating.String{
		{
			Value:      " ",
			StringType: freefloating.WhiteSpaceType,
			Position:   position.NewPosition(2, 2, 17, 18),
		},
	}
	lexer.Lex(lv)
	actual = lv.Tkn.FreeFloating
	assert.DeepEqual(t, expected, actual)
}

func TestVarNameByteChars(t *testing.T) {
	src := "<?php $\x80 $\xff"

	lexer := NewLexer([]byte(src))
	lv := &lval{}

	lexer.Lex(lv)
	assert.Equal(t, "$\x80", lv.Tkn.Value)

	lexer.Lex(lv)
	assert.Equal(t, "$\xff", lv.Tkn.Value)
}

func TestStringVarNameByteChars(t *testing.T) {
	src := "<?php \"$\x80 $\xff\""

	lexer := NewLexer([]byte(src))
	lv := &lval{}

	lexer.Lex(lv)
	assert.Equal(t, "\"", lv.Tkn.Value)

	lexer.Lex(lv)
	assert.Equal(t, "$\x80", lv.Tkn.Value)

	lexer.Lex(lv)
	assert.Equal(t, " ", lv.Tkn.Value)

	lexer.Lex(lv)
	assert.Equal(t, "$\xff", lv.Tkn.Value)

	lexer.Lex(lv)
	assert.Equal(t, "\"", lv.Tkn.Value)
}

func TestIgnoreControllCharacters(t *testing.T) {
	src := "<?php \004 echo $b;"

	lexer := NewLexer([]byte(src))
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

	lexer := NewLexer([]byte(src))
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
