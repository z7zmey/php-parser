%{
// Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// blame: jnml, labs.nic.cz

package main

import (
    "bufio"
    "fmt"
    "os"
    "bytes"
)

var (
    src         = bufio.NewReader(os.Stdin)
    buf         []byte
    current     byte
    isPhp       bool
    sc          int
)

const (
    INITIAL = iota
    PHP
)

func skipHtml() {
    for !isPhp{
        t_open_tag := []byte("<?")
        if p, err := src.Peek(2); err == nil && bytes.Equal(p, t_open_tag) {
            isPhp = true
            break
        }

        _, err := src.ReadByte()
        if err != nil {
            break
        }
    }
}

func getc() byte {
    // skipHtml()

    if current != 0 {
        buf = append(buf, current)
    }
    current = 0
    if b, err := src.ReadByte(); err == nil {
        current = b
    }
    return current
}

func begin(cond int) {
	sc = cond
}

func main() { // This left brace is closed by *1
    c := getc() // init
%}

%s PHP

%yyt sc
%yyc c
%yyn c = getc()
%yyb last == '\n' || last = '\0'

D   [0-9]+

%%
    buf = buf[:0]   // The code before the first rule executed before every scan cycle (rule #0 / state 0 action)


<*>\0                                               return
<INITIAL>[ \t\n\r]+                                 
<INITIAL>.                                          
<INITIAL>\<\?|\<\?php                               fmt.Println("T_OPEN_TAG");begin(PHP)
<INITIAL>\<\?=                                      fmt.Println("T_OPEN_TAG_WITH_ECHO");begin(PHP)

<PHP>.                                      fmt.Printf("%q\n", buf)
<PHP>[ \t\n\r]+                             fmt.Println("T_WHITESPACE")
    // <PHP>\<\?|\<\?php                           fmt.Println("T_OPEN_TAG");
    // <PHP>\<\?=                                  fmt.Println("T_OPEN_TAG_WITH_ECHO");
<PHP>\?\>                                   fmt.Println("T_CLOSE_TAG");begin(INITIAL)
    // <PHP>abstract                               fmt.Println("T_ABSTRACT")
    // <PHP>&=                                     fmt.Println("T_AND_EQUAL")
    // <PHP>\(array\)                              fmt.Println("T_ARRAY_CAST")
    // <PHP>\(bool\)|\(boolean\)                   fmt.Println("T_BOOL_CAST")
    // <PHP>\(real\)|\(double\)|\(float\)          fmt.Println("T_DOUBLE_CAST")
    // <PHP>\(int\)|\(integer\)                    fmt.Println("T_INT_CAST")
    // <PHP>\(object\)                             fmt.Println("T_OBJECT_CAST")
    // <PHP>\(string\)                             fmt.Println("T_STRING_CAST")
    // <PHP>\(unset\)                              fmt.Println("T_UNSET_CAST")
    // <PHP>array                                  fmt.Println("T_ARRAY")
    // <PHP>as                                     fmt.Println("T_AS")
    // <PHP>&&                                     fmt.Println("T_BOOLEAN_AND")
    // <PHP>\|\|                                   fmt.Println("T_BOOLEAN_OR")
    // <PHP>break                                  fmt.Println("T_BREAK")
    // <PHP>callable                               fmt.Println("T_CALLABLE")
    // <PHP>case                                   fmt.Println("T_CASE")
    // <PHP>catch                                  fmt.Println("T_CATCH")
    // <PHP>class                                  fmt.Println("T_CLASS")
    // <PHP>__CLASS__                              fmt.Println("T_CLASS_C")
    // <PHP>__DIR__                                fmt.Println("T_DIR")
    // <PHP>clone                                  fmt.Println("T_CLONE")
    // <PHP>[/][/][^\n]+                           fmt.Println("T_COMMENT"); // TODO: multiline comment
    // <PHP>\.=                                    fmt.Println("T_CONCAT_EQUAL");
    // <PHP>[/]=                                   fmt.Println("T_DIV_EQUAL");
    // <PHP>const                                  fmt.Println("T_CONST");
    // <PHP>continue                               fmt.Println("T_CONTINUE");
    // <PHP>--                                     fmt.Println("T_DEC");
    // <PHP>declare                                fmt.Println("T_DECLARE");
    // <PHP>default                                fmt.Println("T_DEFAULT");
    // <PHP>do                                     fmt.Println("T_DO");
    // <PHP>{D}\.{D}?|\.{D}                        fmt.Println("T_DNUMBER");
    // <PHP>=>                                     fmt.Println("T_DOUBLE_ARROW");
    // <PHP>::                                     fmt.Println("T_DOUBLE_COLON");
    // <PHP>echo                                   fmt.Println("T_ECHO");
    // <PHP>\.\.\.                                 fmt.Println("T_ELLIPSIS");
    // <PHP>else                                   fmt.Println("T_ELSE");
    // <PHP>elseif                                 fmt.Println("T_ELSEIF");
    // <PHP>empty                                  fmt.Println("T_EMPTY");
    // <PHP>endfor                                 fmt.Println("T_ENDFOR")
    // <PHP>endforeach                             fmt.Println("T_ENDFOREACH")
    // <PHP>endif                                  fmt.Println("T_ENDIF")
    // <PHP>endswitch                              fmt.Println("T_ENDSWITCH")
    // <PHP>endwhile                               fmt.Println("T_ENDWHILE")
    // <PHP>exit|die                               fmt.Println("T_EXIT")
    // <PHP>extends                                fmt.Println("T_EXTENDS")
    // <PHP>__FILE__                               fmt.Println("T_FILE")
    // <PHP>final                                  fmt.Println("T_FINAL")
    // <PHP>finally                                fmt.Println("T_FINALLY")
    // <PHP>for                                    fmt.Println("T_FOR")
    // <PHP>foreach                                fmt.Println("T_FOREACH")
    // <PHP>function|cfunction                     fmt.Println("T_FUNCTION")
    // <PHP>__FUNCTION__                           fmt.Println("T_FUNC_C")
    // <PHP>global                                 fmt.Println("T_GLOBAL")
    // <PHP>goto                                   fmt.Println("T_GOTO")
    // <PHP>if                                     fmt.Println("T_IF")
    // <PHP>implements                             fmt.Println("T_IMPLEMENTS")
    // <PHP>\+\+                                   fmt.Println("T_INC")
    // <PHP>instanceof                             fmt.Println("T_INSTANCEOF")
    // <PHP>insteadof                              fmt.Println("T_INSTEADOF")
    // <PHP>interface                              fmt.Println("T_INTERFACE")
    // <PHP>==                                     fmt.Println("T_IS_EQUAL")
    // <PHP>\>=                                    fmt.Println("T_IS_GREATER_OR_EQUAL")
    // <PHP>===                                    fmt.Println("T_IS_IDENTICAL")
    // <PHP>\!=|\<\>                               fmt.Println("T_IS_NOT_EQUAL")
    // <PHP>\!==                                   fmt.Println("T_IS_NOT_IDENTICAL")
    // <PHP>\<=                                    fmt.Println("T_IS_SMALLER_OR_EQUAL")
    // <PHP>\<=\>                                  fmt.Println("T_SPACESHIP")
    // <PHP>__LINE__                               fmt.Println("T_LINE")
    // <PHP>{D}                                    fmt.Println("T_LNUMBER") // TODO: parse 0x1ac
    // <PHP>and                                    fmt.Println("T_LOGICAL_AND")
    // <PHP>or                                     fmt.Println("T_LOGICAL_OR")
    // <PHP>xor                                    fmt.Println("T_LOGICAL_XOR")
    // <PHP>__METHOD__                             fmt.Println("T_METHOD_C")
    // <PHP>-=                                     fmt.Println("T_MINUS_EQUAL")
    // <PHP>%=                                     fmt.Println("T_MOD_EQUAL")
    // <PHP>\*=                                    fmt.Println("T_MUL_EQUAL")
    // <PHP>namespace                              fmt.Println("T_NAMESPACE")
    // <PHP>__NAMESPACE__                          fmt.Println("T_NS_C")
    // <PHP>\\                                     fmt.Println("T_NS_SEPARATOR")
    // <PHP>new                                    fmt.Println("T_NEW")
    // <PHP>->                                     fmt.Println("T_OBJECT_OPERATOR")
    // <PHP>\|=                                    fmt.Println("T_OR_EQUAL")
    // <PHP>\+=                                    fmt.Println("T_PLUS_EQUAL")
    // <PHP>\*\*                                   fmt.Println("T_POW")
    // <PHP>\*\*=                                  fmt.Println("T_POW_EQUAL")
    // <PHP>private                                fmt.Println("T_PRIVATE")
    // <PHP>public                                 fmt.Println("T_PUBLIC")
    // <PHP>protected                              fmt.Println("T_PROTECTED")
    // <PHP>return                                 fmt.Println("T_RETURN")
    // <PHP>\<\<                                   fmt.Println("T_SL")
    // <PHP>\<\<=                                  fmt.Println("T_SL_EQUAL")
    // <PHP>\>\>                                   fmt.Println("T_SR")
    // <PHP>\>\>=                                  fmt.Println("T_SR_EQUAL")
    // <PHP>static                                 fmt.Println("T_STATIC")
    // <PHP>switch                                 fmt.Println("T_SWITCH")
    // <PHP>throw                                  fmt.Println("T_THROW")
    // <PHP>trait                                  fmt.Println("T_TRAIT")
    // <PHP>__TRAIT__                              fmt.Println("T_TRAIT_C")
    // <PHP>try                                    fmt.Println("T_TRY")
    // <PHP>use                                    fmt.Println("T_USE")
    // <PHP>var                                    fmt.Println("T_VAR")
    // <PHP>while                                  fmt.Println("T_WHILE")
    // <PHP>\^=                                    fmt.Println("T_XOR_EQUAL")
    // <PHP>yield                                  fmt.Println("T_YIELD")
    // <PHP>yield\nfrom                            fmt.Println("T_YIELD_FROM")
    // <PHP>'[^']*'                                fmt.Println("T_CONSTANT_ENCAPSED_STRING")

    // <PHP>\{\$                                   fmt.Println("T_CURLY_OPEN");src.UnreadByte();c='$'
    // <PHP>\$[A-Za-z][A-Za-z0-9]*                 fmt.Println("T_VARIABLE") // TODO allow -_ and other

    // <PHP>[/][*][^*]*[*]+([^*/][^*]*[*]+)*[/]    fmt.Println("T_COMMENT");
    // <PHP>[/][*]                                 panic("Unterminated comment")

%%
// \{\$                            fmt.Println("T_CURLY_OPEN");
// \$\{                            fmt.Println("T_DOLLAR_OPEN_CURLY_BRACES");

// T_DOC_COMMENT	/** */	PHPDoc-комментарии
// T_ENCAPSED_AND_WHITESPACE	" $a"	константная часть строки с переменными
// T_ENDDECLARE	enddeclare	declare, альтернативный синтаксис
// T_END_HEREDOC	 	синтаксис heredoc
// T_START_HEREDOC	<<<	синтаксис heredoc
// T_EVAL	eval()	eval()
// T_HALT_COMPILER	__halt_compiler()	__halt_compiler (доступно с PHP 5.1.0)
// T_INCLUDE	include()	include
// T_INCLUDE_ONCE	include_once()	include_once
// T_INLINE_HTML	 	текст вне PHP
// T_ISSET	isset()	isset()
// T_LIST	list()	list()
// T_NUM_STRING	"$a[0]"	цифровой индекс массива внутри строки
// T_PRINT	print()	print
// T_REQUIRE	require()	require
// T_REQUIRE_ONCE	require_once()	require_once
// T_STRING	parent, self и т.п..	идентификаторы, например, ключевые слова вроде parent и self, сюда подходят также имена функций, классов и некоторые другие. Смотрите также T_CONSTANT_ENCAPSED_STRING
// T_STRING_VARNAME	"${a	переменные внутри строки
// T_UNSET	unset()	unset()
// T_VARIABLE	$foo	переменные

// T_BAD_CHARACTER	 	все, что ниже ASCII 32 исключая \t (0x09), \n (0x0a) и \r (0x0d)
// T_CONSTANT_ENCAPSED_STRING	"foo" или 'bar'	строковой синтаксис

    // The golex generated scanner enters top of the user code section when
    // lexem recongition fails. In this example it should never happen.
    fmt.Println("UNDEFENED LEXEM")
    goto yystate0
    // panic("scanner internal error")

} // *1 this right brace
