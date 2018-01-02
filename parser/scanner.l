%{
// Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// blame: jnml, labs.nic.cz

package parser

import (
    "fmt"
    "bytes"
    "github.com/z7zmey/php-parser/token"
)

const (
	INITIAL = iota
	PHP
	STRING
	STRING_VAR
	STRING_VAR_INDEX
	STRING_VAR_NAME
	PROPERTY
	HEREDOC_END
	NOWDOC
	HEREDOC
	BACKQUOTE
)

var heredocLabel []byte

func (l *lexer) Lex(lval *yySymType) int { // Lex(lval *yySymType)
  c := l.Enter()

%}

%s PHP STRING STRING_VAR STRING_VAR_INDEX STRING_VAR_NAME PROPERTY HEREDOC_END NOWDOC HEREDOC BACKQUOTE

%yyb last == '\n' || last = '\0'
%yyt l.getCurrentState()
%yyc c
%yyn c = l.Next()
%yym l.Mark()
%optioncase-insensitive

LNUM            [0-9]+
DNUM            ([0-9]*"."[0-9]+)|([0-9]+"."[0-9]*)
HNUM            0x[0-9a-fA-F]+
BNUM            0b[01]+
EXPONENT_DNUM   (({LNUM}|{DNUM})[eE][+-]?{LNUM})
VAR_NAME        [a-zA-Z_\x7f-\xff][a-zA-Z0-9_\x7f-\xff]*
OPERATORS       [;:,.\[\]()|\/\^&\+-*=%!~$<>?@]
NEW_LINE        (\r|\n|\r\n)

%%
    c = l.Rule0()

<INITIAL>[ \t\n\r]+                             lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil)));
<INITIAL>.
    tb := []byte{}

    for {
        if c == -1 {
            tb = l.TokenBytes(nil);
            break;
        }

        if '?' == rune(c) {
            tb = l.TokenBytes(nil);
            if (len(tb) < 2 || tb[len(tb)-1] != '<') {
                c = l.Next()
                continue;
            }

            tb = l.ungetN(1)
            break;
        }

        c = l.Next()
    }

    lval.token = token.NewToken(l.handleNewLine(tb)); 
    return T_INLINE_HTML

<INITIAL>\<\?php([ \t]|{NEW_LINE})              l.begin(PHP);lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil)));// return T_OPEN_TAG;
<INITIAL>\<\?                                   l.begin(PHP);lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil)));// return T_OPEN_TAG;
<INITIAL>\<\?=                                  l.begin(PHP);lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ECHO;

<PHP>[ \t\n\r]+                                 lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil)));// return T_WHITESPACE
<PHP>\?\>{NEW_LINE}?                            l.begin(INITIAL);lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return rune2Class(';');

<PHP>{DNUM}|{EXPONENT_DNUM}                     lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_DNUMBER
<PHP>{BNUM}
    tb := l.TokenBytes(nil)
    i:=2
    BNUMFOR:for {
        if i > len(tb)-1 {
            break BNUMFOR;
        }
        switch tb[i] {
            case '0': i++;
            default: break BNUMFOR;
        }
    }
    if len(tb) - i < 64 {
        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_LNUMBER
    } else {
        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_DNUMBER
    }
<PHP>{LNUM}
    if len(l.TokenBytes(nil)) < 20 {
        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_LNUMBER
    } else {
        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_DNUMBER
    }
<PHP>{HNUM}
    tb := l.TokenBytes(nil)
    i:=2
    HNUMFOR:for {
        if i > len(tb)-1 {
            break HNUMFOR;
        }
        switch tb[i] {
            case '0': i++;
            default: break HNUMFOR;
        }
    }
    length := len(tb) - i
    if length < 16 || (length == 16 && tb[i] <= '7') {
        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_LNUMBER
    } else {
        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_DNUMBER
    }

<PHP>abstract                                   lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ABSTRACT
<PHP>array                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ARRAY
<PHP>as                                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_AS
<PHP>break                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_BREAK
<PHP>callable                                   lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_CALLABLE
<PHP>case                                       lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_CASE
<PHP>catch                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_CATCH
<PHP>class                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_CLASS
<PHP>clone                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_CLONE
<PHP>const                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_CONST;
<PHP>continue                                   lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_CONTINUE;
<PHP>declare                                    lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_DECLARE;
<PHP>default                                    lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_DEFAULT;
<PHP>do                                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_DO;
<PHP>echo                                       lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ECHO;
<PHP>else                                       lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ELSE;
<PHP>elseif                                     lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ELSEIF;
<PHP>empty                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_EMPTY;
<PHP>enddeclare                                 lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ENDDECLARE
<PHP>endfor                                     lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ENDFOR
<PHP>endforeach                                 lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ENDFOREACH
<PHP>endif                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ENDIF
<PHP>endswitch                                  lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ENDSWITCH
<PHP>endwhile                                   lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ENDWHILE
<PHP>eval                                       lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_EVAL
<PHP>exit|die                                   lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_EXIT
<PHP>extends                                    lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_EXTENDS
<PHP>final                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_FINAL
<PHP>finally                                    lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_FINALLY
<PHP>for                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_FOR
<PHP>foreach                                    lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_FOREACH
<PHP>function|cfunction                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_FUNCTION
<PHP>global                                     lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_GLOBAL
<PHP>goto                                       lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_GOTO
<PHP>if                                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_IF
<PHP>isset                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ISSET
<PHP>implements                                 lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_IMPLEMENTS
<PHP>instanceof                                 lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_INSTANCEOF
<PHP>insteadof                                  lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_INSTEADOF
<PHP>interface                                  lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_INTERFACE
<PHP>list                                       lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_LIST
<PHP>namespace                                  lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_NAMESPACE
<PHP>private                                    lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_PRIVATE
<PHP>public                                     lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_PUBLIC
<PHP>print                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_PRINT
<PHP>protected                                  lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_PROTECTED
<PHP>return                                     lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_RETURN
<PHP>static                                     lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_STATIC
<PHP>switch                                     lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_SWITCH
<PHP>throw                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_THROW
<PHP>trait                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_TRAIT
<PHP>try                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_TRY
<PHP>unset                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_UNSET
<PHP>use                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_USE
<PHP>var                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_VAR
<PHP>while                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_WHILE
<PHP>yield[ \t\n\r]+from[^a-zA-Z0-9_\x80-\xff]  lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_YIELD_FROM
<PHP>yield                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_YIELD
<PHP>include                                    lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_INCLUDE
<PHP>include_once                               lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_INCLUDE_ONCE
<PHP>require                                    lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_REQUIRE
<PHP>require_once                               lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_REQUIRE_ONCE
<PHP>__CLASS__                                  lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_CLASS_C
<PHP>__DIR__                                    lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_DIR
<PHP>__FILE__                                   lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_FILE
<PHP>__FUNCTION__                               lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_FUNC_C
<PHP>__LINE__                                   lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_LINE
<PHP>__NAMESPACE__                              lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_NS_C
<PHP>__METHOD__                                 lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_METHOD_C
<PHP>__TRAIT__                                  lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_TRAIT_C
<PHP>__halt_compiler                            lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_HALT_COMPILER
<PHP>\([ \t]*array[ \t]*\)                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ARRAY_CAST
<PHP>\([ \t]*(bool|boolean)[ \t]*\)             lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_BOOL_CAST
<PHP>\([ \t]*(real|double|float)[ \t]*\)        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_DOUBLE_CAST
<PHP>\([ \t]*(int|integer)[ \t]*\)              lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_INT_CAST
<PHP>\([ \t]*object[ \t]*\)                     lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_OBJECT_CAST
<PHP>\([ \t]*string[ \t]*\)                     lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_STRING_CAST
<PHP>\([ \t]*unset[ \t]*\)                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_UNSET_CAST
<PHP>new                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_NEW
<PHP>and                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_LOGICAL_AND
<PHP>or                                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_LOGICAL_OR
<PHP>xor                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_LOGICAL_XOR
<PHP>\\                                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_NS_SEPARATOR
<PHP>\.\.\.                                     lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ELLIPSIS;
<PHP>::                                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_PAAMAYIM_NEKUDOTAYIM; // T_DOUBLE_COLON
<PHP>&&                                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_BOOLEAN_AND
<PHP>\|\|                                       lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_BOOLEAN_OR
<PHP>&=                                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_AND_EQUAL
<PHP>\|=                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_OR_EQUAL
<PHP>\.=                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_CONCAT_EQUAL;
<PHP>\*=                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_MUL_EQUAL
<PHP>\*\*=                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_POW_EQUAL
<PHP>[/]=                                       lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_DIV_EQUAL;
<PHP>\+=                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_PLUS_EQUAL
<PHP>-=                                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_MINUS_EQUAL
<PHP>\^=                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_XOR_EQUAL
<PHP>%=                                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_MOD_EQUAL
<PHP>--                                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_DEC;
<PHP>\+\+                                       lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_INC
<PHP>=>                                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_DOUBLE_ARROW;
<PHP>\<=\>                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_SPACESHIP
<PHP>\!=|\<\>                                   lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_IS_NOT_EQUAL
<PHP>\!==                                       lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_IS_NOT_IDENTICAL
<PHP>==                                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_IS_EQUAL
<PHP>===                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_IS_IDENTICAL
<PHP>\<\<=                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_SL_EQUAL
<PHP>\>\>=                                      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_SR_EQUAL
<PHP>\>=                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_IS_GREATER_OR_EQUAL
<PHP>\<=                                        lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_IS_SMALLER_OR_EQUAL
<PHP>\*\*                                       lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_POW
<PHP>\<\<                                       lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_SL
<PHP>\>\>                                       lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_SR
<PHP>\?\?                                       lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_COALESCE
<PHP>(#|[/][/]).*{NEW_LINE}                     lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil)));// return T_COMMENT; // TODO: handle ?>
<PHP>([/][*])|([/][*][*])
    tb := l.TokenBytes(nil)
    is_doc_comment := false
    if len(tb) > 2 {
        is_doc_comment = true
        l.phpDocComment = ""
    }

    for {
        if c == -1 {
            break; // TODO: Unterminated comment starting line %d
        }

        p := c
        c = l.Next()

        if rune(p) == '*' && rune(c) == '/' {
            c = l.Next()
            break;
        }
    }

    lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil)))
    if is_doc_comment {
        l.phpDocComment = string(l.TokenBytes(nil))
        // return T_DOC_COMMENT
    } else {
        // return T_COMMENT
    }

<PHP>{OPERATORS}                                lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return rune2Class(rune(l.TokenBytes(nil)[0]))

<PHP>\{                                         l.pushState(PHP); lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return rune2Class(rune(l.TokenBytes(nil)[0]))
<PHP>\}                                         l.popState();     lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return rune2Class(rune(l.TokenBytes(nil)[0])); l.phpDocComment = ""
<PHP>\${VAR_NAME}                               lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_VARIABLE
<PHP>{VAR_NAME}                                 lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_STRING

<PHP>->                                         l.begin(PROPERTY);lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_OBJECT_OPERATOR;
<PROPERTY>[ \t\n\r]+                            lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_WHITESPACE;
<PROPERTY>->                                    lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_OBJECT_OPERATOR;
<PROPERTY>{VAR_NAME}                            l.begin(PHP);lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_STRING;
<PROPERTY>.                                     l.ungetN(1);l.begin(PHP)

<PHP>[\']([^\\\']*([\\].)*)*[\']             lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_CONSTANT_ENCAPSED_STRING;

<PHP>`                                          l.begin(BACKQUOTE); lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return rune2Class(rune(l.TokenBytes(nil)[0]))
<BACKQUOTE>`                                    l.begin(PHP); lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return rune2Class(rune(l.TokenBytes(nil)[0]))

<PHP>[b]?\<\<\<[ \t]*({VAR_NAME}|([']{VAR_NAME}['])|(["]{VAR_NAME}["])){NEW_LINE}
    tb := l.TokenBytes(nil)
    binPrefix := 0
    if tb[0] == 'b' {
        binPrefix = 1
    }
    
    lblFirst := 3 + binPrefix
    lblLast := len(tb)-2
    if tb[lblLast] == '\r' {
        lblLast--
    }

    for {
        if tb[lblFirst] == ' ' || tb[lblFirst] == '\t' {
            lblFirst++
            continue
        }

        break
    }

    switch tb[lblFirst] {
        case '\'' : 
            lblFirst++
            lblLast--
            l.begin(NOWDOC)
        case '"' : 
            lblFirst++
            lblLast--
            l.begin(HEREDOC)
        default:
            l.begin(HEREDOC)
    }

    heredocLabel = make([]byte, lblLast - lblFirst + 1)
    copy(heredocLabel, tb[lblFirst:lblLast+1])
    
    ungetCnt := len(heredocLabel)
    searchLabelAhead := []byte{}
    for i := 0; i < len(heredocLabel); i++ {
        if c == -1 {
            break;
        }
        searchLabelAhead = append(searchLabelAhead, byte(rune(c)))
        c = l.Next()
    }

    if bytes.Equal(heredocLabel, searchLabelAhead) && ';' == rune(c) {
        ungetCnt++
        c = l.Next()
        if '\n' == rune(c) || '\r' == rune(c) {
            l.begin(HEREDOC_END)
        }
    }

    l.ungetN(ungetCnt)

    lval.token = token.NewToken(l.handleNewLine(tb)); 
    return T_START_HEREDOC

<NOWDOC>.
    searchLabel := []byte{}
    tb := []byte{}
    
    for {
        if c == -1 {
            break;
        }

        if '\n' == rune(c) || '\r' == rune(c) {
            if bytes.Equal(append(heredocLabel, ';'), searchLabel) {
                l.begin(HEREDOC_END)
                tb = l.ungetN(len(heredocLabel)+1)
                break;
            }
            
            searchLabel = []byte{}
        } else {
            searchLabel = append(searchLabel, byte(rune(c)))
        }

        c = l.Next()
    }

    lval.token = token.NewToken(l.handleNewLine(tb)); 
    return T_ENCAPSED_AND_WHITESPACE

<HEREDOC_END>{VAR_NAME}\;                       l.begin(PHP);lval.token = token.NewToken(l.handleNewLine(l.ungetN(1))); return T_END_HEREDOC

<PHP>[b]?[\"]
    binPrefix := l.TokenBytes(nil)[0] == 'b'

    beginString := func() int {
        cnt := 1; if (binPrefix) {cnt = 2}
        
        l.ungetN(len(l.TokenBytes(nil))-cnt)
        tokenBytes := l.TokenBytes(nil)[:cnt]
        l.pushState(STRING)

        lval.token = token.NewToken(l.handleNewLine(tokenBytes)); return rune2Class('"')
    }

    F:for {
        if c == -1 {
            break;
        }

        switch c {
            case '"' : 
                c = l.Next();
                lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_CONSTANT_ENCAPSED_STRING
                break F;
            
            case '$':
                c = l.Next();
                if rune(c) == '{' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c >= '\u007f' && c <= 'ÿ' {
                    return beginString()
                    break F;
                } 
                l.ungetN(0)
            
            case '{':
                c = l.Next();
                if rune(c) == '$' {
                    return beginString()
                    break F;
                } 
                l.ungetN(0)

            case '\\':
                c = l.Next();
        }

        c = l.Next()
    }

<STRING>\"                                  l.popState(); lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return rune2Class(rune(l.TokenBytes(nil)[0]))
<STRING,HEREDOC,BACKQUOTE>\{\$              lval.token = token.NewToken(l.handleNewLine(l.ungetN(1))); l.pushState(PHP); return T_CURLY_OPEN
<STRING,HEREDOC,BACKQUOTE>\$\{              l.pushState(STRING_VAR_NAME);lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_DOLLAR_OPEN_CURLY_BRACES
<STRING,HEREDOC,BACKQUOTE>\$                l.ungetN(1);l.pushState(STRING_VAR)
<STRING>.|[ \t\n\r]
    F1:for {
        if c == -1 {
            break;
        }

        switch c {
            case '"' : 
                lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil)));
                return T_ENCAPSED_AND_WHITESPACE
                break F1;
            
            case '$':
                c = l.Next();
                if rune(c) == '{' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c >= '\u007f' && c <= 'ÿ' {
                    l.ungetN(1)
                    tb := l.TokenBytes(nil)
                    lval.token = token.NewToken(l.handleNewLine(tb[:len(tb)-1]));
                    return T_ENCAPSED_AND_WHITESPACE
                    break F1;
                } 
                l.ungetN(0)
            
            case '{':
                c = l.Next();
                if rune(c) == '$' {
                    l.ungetN(1)
                    tb := l.TokenBytes(nil)
                    lval.token = token.NewToken(l.handleNewLine(tb[:len(tb)-1])); 
                    return T_ENCAPSED_AND_WHITESPACE
                    break F1;
                } 
                l.ungetN(0)

            case '\\':
                c = l.Next();
        }

        c = l.Next()
    }

<BACKQUOTE>.|[ \t\n\r]
    F2:for {
        if c == -1 {
            break;
        }

        switch c {
            case '`' :
                lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); 
                return T_ENCAPSED_AND_WHITESPACE
                break F2;
            
            case '$':
                c = l.Next();
                if rune(c) == '{' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c >= '\u007f' && c <= 'ÿ' {
                    l.ungetN(1)
                    tb := l.TokenBytes(nil)
                    lval.token = token.NewToken(l.handleNewLine(tb[:len(tb)-1])); 
                    return T_ENCAPSED_AND_WHITESPACE
                    break F2;
                } 
                l.ungetN(0)
            
            case '{':
                c = l.Next();
                if rune(c) == '$' {
                    l.ungetN(1)
                    tb := l.TokenBytes(nil)
                    lval.token = token.NewToken(l.handleNewLine(tb[:len(tb)-1])); 
                    return T_ENCAPSED_AND_WHITESPACE
                    break F2;
                } 
                l.ungetN(0)

            case '\\':
                c = l.Next();
        }

        c = l.Next()
    }

<HEREDOC>.|[ \t\n\r]
    searchLabel := []byte{}
    tb := []byte{}
    
    HEREDOCFOR:for {
        if c == -1 {
            break;
        }

        switch c {
            case '\n': fallthrough
            case '\r':
                if bytes.Equal(append(heredocLabel, ';'), searchLabel) { // TODO handle ';' as optional
                    l.begin(HEREDOC_END)
                    tb = l.ungetN(len(heredocLabel)+1)
                    break HEREDOCFOR;
                }
                
                searchLabel = []byte{}
            
            case '$':
                c = l.Next();
                if rune(c) == '{' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c >= '\u007f' && c <= 'ÿ' {
                    tb = l.ungetN(1)
                    break HEREDOCFOR;
                } 
                l.ungetN(0)
                searchLabel = []byte{}
            
            case '{':
                c = l.Next();
                if rune(c) == '$' {
                    tb = l.ungetN(1)
                    break HEREDOCFOR;
                } 
                l.ungetN(0)
                searchLabel = []byte{}

            case '\\':
                c = l.Next();
                searchLabel = []byte{}
            
            default:
                searchLabel = append(searchLabel, byte(rune(c)))
        }

        c = l.Next()
    }
    
    lval.token = token.NewToken(l.handleNewLine(tb)); 
    return T_ENCAPSED_AND_WHITESPACE

<STRING_VAR>\${VAR_NAME}                    lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_VARIABLE
<STRING_VAR>->{VAR_NAME}                    lval.token = token.NewToken(l.handleNewLine(l.ungetN(len(l.TokenBytes(nil))-2))); return T_OBJECT_OPERATOR
<STRING_VAR>{VAR_NAME}                      l.popState();lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_STRING
<STRING_VAR>\[                              l.pushState(STRING_VAR_INDEX);lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return rune2Class(rune(l.TokenBytes(nil)[0]))
<STRING_VAR>.|[ \t\n\r]                     l.ungetN(1);l.popState()

<STRING_VAR_INDEX>{LNUM}|{HNUM}|{BNUM}      lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_NUM_STRING
<STRING_VAR_INDEX>\${VAR_NAME}              lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_VARIABLE
<STRING_VAR_INDEX>{VAR_NAME}                lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_STRING
<STRING_VAR_INDEX>\]                        l.popState(); l.popState();lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return rune2Class(rune(l.TokenBytes(nil)[0]))
<STRING_VAR_INDEX>[ \n\r\t\\'#]             l.popState(); l.popState();lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return T_ENCAPSED_AND_WHITESPACE
<STRING_VAR_INDEX>{OPERATORS}               lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return rune2Class(rune(l.TokenBytes(nil)[0]))
<STRING_VAR_INDEX>.                         lval.token = token.NewToken(l.handleNewLine(l.TokenBytes(nil))); return rune2Class(rune(l.TokenBytes(nil)[0]))

<STRING_VAR_NAME>{VAR_NAME}[\[\}]           l.popState();l.pushState(PHP);lval.token = token.NewToken(l.handleNewLine(l.ungetN(1))); return T_STRING_VARNAME
<STRING_VAR_NAME>.                          l.ungetN(1);l.popState();l.pushState(PHP)

%%
    if c, ok := l.Abort(); ok { return int(c) }
    goto yyAction
}