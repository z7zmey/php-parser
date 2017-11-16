%{
// Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// blame: jnml, labs.nic.cz

package main

import (
    "bufio"
    "go/token"
    "io"
    "unicode"
    "fmt"
    "os"
    "bytes"

    "github.com/cznic/golex/lex"
)

// Allocate Character classes anywhere in [0x80, 0xFF].
const (
    classUnicodeLeter = iota + 0x80
    classUnicodeDigit
    classOther
)

var sc int

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
)


type lexer struct {
    *lex.Lexer
}

var stateStack = []int{PHP}
var heredocLabel []byte

func pushState(state int) {
    sc = state
    stateStack = append(stateStack, state)
    fmt.Printf("PUSH STATE; CURRENT STATE: %d\n", state)
}

func popState() {
    len := len(stateStack)
    if len <= 1 {
        return
    }

    sc = stateStack[len - 2]
    stateStack = stateStack[:len-1]
    
    fmt.Printf("POP STATE; CURRENT STATE: %d\n", sc)
}

func begin(state int) {
    sc = state
}

func rune2Class(r rune) int {
    if r >= 0 && r < 0x80 { // Keep ASCII as it is.
        return int(r)
    }
    if unicode.IsLetter(r) {
        return classUnicodeLeter
    }
    if unicode.IsDigit(r) {
        return classUnicodeDigit
    }
    return classOther
}

func newLexer(src io.Reader, dst io.Writer, fName string) *lexer {
    file := token.NewFileSet().AddFile(fName, -1, 1<<31-1)
    lx, err := lex.New(file, bufio.NewReader(src), lex.RuneClass(rune2Class))
    if (err != nil) { panic(err) }
    return &lexer{lx}
}

type yySymType struct {}

func (l *lexer) unget(r rune) []byte{
    l.Unget(l.Lookahead())
    
    chars := l.Token();
    lastChar := chars[len(chars)-1];

    if lastChar.Rune != r {
        return l.TokenBytes(nil)
    }

    l.Unget(lastChar);

    buf := l.TokenBytes(nil)
    buf = buf[:len(buf)-1]

    return buf
}

func (l *lexer) ungetN(n int) []byte{
    l.Unget(l.Lookahead())
    
    chars := l.Token();
    
    for i := 1; i <= n; i++ {
        char := chars[len(chars)-i];
        l.Unget(char);
    }

    buf := l.TokenBytes(nil)
    buf = buf[:len(buf)-n]

    return buf
}

func (l *lexer) Lex() int { // Lex(lval *yySymType)
  c := l.Enter()

%}

%s PHP STRING STRING_VAR STRING_VAR_INDEX STRING_VAR_NAME PROPERTY HEREDOC_END NOWDOC

%yyb last == '\n' || last = '\0'
%yyt sc
%yyc c
%yyn c = l.Next()
%yym l.Mark()

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
    // ([\$]{NCH})*

<INITIAL>[ \t\n\r]+
<INITIAL>.
<INITIAL>\<\?|\<\?php                           fmt.Println("T_OPEN_TAG");begin(PHP)
<INITIAL>\<\?=                                  fmt.Println("T_OPEN_TAG_WITH_ECHO");begin(PHP)

<PHP>[ \t\n\r]+                                 fmt.Println("T_WHITESPACE")
<PHP>\?\>                                       fmt.Println("T_CLOSE_TAG");begin(INITIAL)

<PHP>abstract                                   fmt.Println("T_ABSTRACT")
<PHP>array                                      fmt.Println("T_ARRAY")
<PHP>as                                         fmt.Println("T_AS")
<PHP>break                                      fmt.Println("T_BREAK")
<PHP>callable                                   fmt.Println("T_CALLABLE")
<PHP>case                                       fmt.Println("T_CASE")
<PHP>catch                                      fmt.Println("T_CATCH")
<PHP>class                                      fmt.Println("T_CLASS")
<PHP>clone                                      fmt.Println("T_CLONE")
<PHP>const                                      fmt.Println("T_CONST");
<PHP>continue                                   fmt.Println("T_CONTINUE");
<PHP>declare                                    fmt.Println("T_DECLARE");
<PHP>default                                    fmt.Println("T_DEFAULT");
<PHP>do                                         fmt.Println("T_DO");
<PHP>echo                                       fmt.Println("T_ECHO");
<PHP>else                                       fmt.Println("T_ELSE");
<PHP>elseif                                     fmt.Println("T_ELSEIF");
<PHP>empty                                      fmt.Println("T_EMPTY");
<PHP>endfor                                     fmt.Println("T_ENDFOR")
<PHP>endforeach                                 fmt.Println("T_ENDFOREACH")
<PHP>endif                                      fmt.Println("T_ENDIF")
<PHP>endswitch                                  fmt.Println("T_ENDSWITCH")
<PHP>endwhile                                   fmt.Println("T_ENDWHILE")
<PHP>exit|die                                   fmt.Println("T_EXIT")
<PHP>extends                                    fmt.Println("T_EXTENDS")
<PHP>final                                      fmt.Println("T_FINAL")
<PHP>finally                                    fmt.Println("T_FINALLY")
<PHP>for                                        fmt.Println("T_FOR")
<PHP>foreach                                    fmt.Println("T_FOREACH")
<PHP>function|cfunction                         fmt.Println("T_FUNCTION")
<PHP>global                                     fmt.Println("T_GLOBAL")
<PHP>goto                                       fmt.Println("T_GOTO")
<PHP>if                                         fmt.Println("T_IF")
<PHP>implements                                 fmt.Println("T_IMPLEMENTS")
<PHP>instanceof                                 fmt.Println("T_INSTANCEOF")
<PHP>insteadof                                  fmt.Println("T_INSTEADOF")
<PHP>interface                                  fmt.Println("T_INTERFACE")
<PHP>namespace                                  fmt.Println("T_NAMESPACE")
<PHP>private                                    fmt.Println("T_PRIVATE")
<PHP>public                                     fmt.Println("T_PUBLIC")
<PHP>protected                                  fmt.Println("T_PROTECTED")
<PHP>return                                     fmt.Println("T_RETURN")
<PHP>static                                     fmt.Println("T_STATIC")
<PHP>switch                                     fmt.Println("T_SWITCH")
<PHP>throw                                      fmt.Println("T_THROW")
<PHP>trait                                      fmt.Println("T_TRAIT")
<PHP>try                                        fmt.Println("T_TRY")
<PHP>use                                        fmt.Println("T_USE")
<PHP>var                                        fmt.Println("T_VAR")
<PHP>while                                      fmt.Println("T_WHILE")
<PHP>yield[ \t\n\r]+from[^a-zA-Z0-9_\x80-\xff]  fmt.Println("T_YIELD_FROM")
<PHP>yield                                      fmt.Println("T_YIELD")
<PHP>__CLASS__                                  fmt.Println("T_CLASS_C")
<PHP>__DIR__                                    fmt.Println("T_DIR")
<PHP>__FILE__                                   fmt.Println("T_FILE")
<PHP>__FUNCTION__                               fmt.Println("T_FUNC_C")
<PHP>__LINE__                                   fmt.Println("T_LINE")
<PHP>__NAMESPACE__                              fmt.Println("T_NS_C")
<PHP>__METHOD__                                 fmt.Println("T_METHOD_C")
<PHP>__TRAIT__                                  fmt.Println("T_TRAIT_C")
<PHP>\([ \t]*array[ \t]*\)                      fmt.Println("T_ARRAY_CAST")
<PHP>\([ \t]*(bool|boolean)[ \t]*\)             fmt.Println("T_BOOL_CAST")
<PHP>\([ \t]*(real|double|float)[ \t]*\)        fmt.Println("T_DOUBLE_CAST")
<PHP>\([ \t]*(int|integer)[ \t]*\)              fmt.Println("T_INT_CAST")
<PHP>\([ \t]*object[ \t]*\)                     fmt.Println("T_OBJECT_CAST")
<PHP>\([ \t]*string[ \t]*\)                     fmt.Println("T_STRING_CAST")
<PHP>\([ \t]*unset[ \t]*\)                      fmt.Println("T_UNSET_CAST")
<PHP>new                                        fmt.Println("T_NEW")
<PHP>and                                        fmt.Println("T_LOGICAL_AND")
<PHP>or                                         fmt.Println("T_LOGICAL_OR")
<PHP>xor                                        fmt.Println("T_LOGICAL_XOR")
<PHP>\\                                         fmt.Println("T_NS_SEPARATOR")
<PHP>\.\.\.                                     fmt.Println("T_ELLIPSIS");
<PHP>::                                         fmt.Println("T_PAAMAYIM_NEKUDOTAYIM"); // T_DOUBLE_COLON
<PHP>&&                                         fmt.Println("T_BOOLEAN_AND")
<PHP>\|\|                                       fmt.Println("T_BOOLEAN_OR")
<PHP>&=                                         fmt.Println("T_AND_EQUAL")
<PHP>\|=                                        fmt.Println("T_OR_EQUAL")
<PHP>\.=                                        fmt.Println("T_CONCAT_EQUAL");
<PHP>\*=                                        fmt.Println("T_MUL_EQUAL")
<PHP>\*\*=                                      fmt.Println("T_POW_EQUAL")
<PHP>[/]=                                       fmt.Println("T_DIV_EQUAL");
<PHP>\+=                                        fmt.Println("T_PLUS_EQUAL")
<PHP>-=                                         fmt.Println("T_MINUS_EQUAL")
<PHP>\^=                                        fmt.Println("T_XOR_EQUAL")
<PHP>%=                                         fmt.Println("T_MOD_EQUAL")
<PHP>--                                         fmt.Println("T_DEC");
<PHP>\+\+                                       fmt.Println("T_INC")
<PHP>=>                                         fmt.Println("T_DOUBLE_ARROW");
<PHP>\<=\>                                      fmt.Println("T_SPACESHIP")
<PHP>\!=|\<\>                                   fmt.Println("T_IS_NOT_EQUAL")
<PHP>\!==                                       fmt.Println("T_IS_NOT_IDENTICAL")
<PHP>==                                         fmt.Println("T_IS_EQUAL")
<PHP>===                                        fmt.Println("T_IS_IDENTICAL")
<PHP>\<\<=                                      fmt.Println("T_SL_EQUAL")
<PHP>\>\>=                                      fmt.Println("T_SR_EQUAL")
<PHP>\>=                                        fmt.Println("T_IS_GREATER_OR_EQUAL")
<PHP>\<=                                        fmt.Println("T_IS_SMALLER_OR_EQUAL")
<PHP>\*\*                                       fmt.Println("T_POW")
<PHP>\<\<                                       fmt.Println("T_SL")
<PHP>\>\>                                       fmt.Println("T_SR")
<PHP>(#|[/][/]){NEW_LINE}                       fmt.Println("T_COMMENT"); // TODO: handle \r\n and allow ?>
<PHP>'[^']*(\\')*'                              fmt.Println("T_CONSTANT_ENCAPSED_STRING")
<PHP>{OPERATORS}                                fmt.Printf("%s\n", l.TokenBytes(nil));

<PHP>\}                                         fmt.Println("}"); popState();
<PHP>\${VAR_NAME}                               fmt.Printf("T_VARIABLE: %q\n", l.TokenBytes(nil))

<PHP>->                                         fmt.Println("T_OBJECT_OPERATOR");begin(PROPERTY)
<PROPERTY>[ \t\n\r]+                            fmt.Println("T_WHITESPACE");
<PROPERTY>->                                    fmt.Println("T_OBJECT_OPERATOR");
<PROPERTY>{VAR_NAME}                            fmt.Println("T_STRING");begin(PHP)
<PROPERTY>.                                     l.ungetN(1);begin(PHP)

<PHP>[\']([^\\\']*([\\][\'])*)*[\']             fmt.Printf("T_CONSTANT_ENCAPSED_STRING: %s\n", l.TokenBytes(nil));

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
            begin(NOWDOC)
        case '"' : 
            lblFirst++
            lblLast--
            fmt.Println("HEREDOC");
            //begin(HEREDOC)
        default:
            //begin(HEREDOC)
    }

    heredocLabel = tb[lblFirst:lblLast+1]
    
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
            begin(HEREDOC_END)
        }
    }

    l.ungetN(ungetCnt)

    fmt.Printf("T_START_HEREDOC: %q\n", tb);

<NOWDOC>.
    searchLabel := []byte{}
    tb := []byte{}
    
    for {
        if c == -1 {
            break;
        }

        if '\n' == rune(c) || '\r' == rune(c) {
            if bytes.Equal(append(heredocLabel, ';'), searchLabel) {
                begin(HEREDOC_END)
                tb = l.ungetN(len(heredocLabel)+1)
                break;
            }
            
            searchLabel = []byte{}
        } else {
            searchLabel = append(searchLabel, byte(rune(c)))
        }

        c = l.Next()
    }
    
    fmt.Printf("T_ENCAPSED_AND_WHITESPACE: %q\n", tb);

<HEREDOC_END>{VAR_NAME}\;                       fmt.Printf("T_END_HEREDOC: %q\n", l.ungetN(1));begin(PHP)
<HEREDOC_END>.                       fmt.Printf("ERROR HEREDOC: %q\n", l.ungetN(1));

<PHP>[b]?[\"]
    binPrefix := l.TokenBytes(nil)[0] == 'b'

    beginString := func() {
        cnt := 1; if (binPrefix) {cnt = 2}
        
        l.ungetN(len(l.TokenBytes(nil))-cnt)
        tokenBytes := l.TokenBytes(nil)[:cnt]
        fmt.Println(string(tokenBytes)) // TODO: RETURN TOKEN
        pushState(STRING)
    }

    F:for {
        if c == -1 {
            break;
        }

        switch c {
            case '"' : 
                c = l.Next();
                fmt.Printf("T_CONSTANT_ENCAPSED_STRING: %s\n", l.TokenBytes(nil));
                break F;
            
            case '$':
                c = l.Next();
                if rune(c) == '{' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c >= '\u007f' && c <= 'ÿ' {
                    beginString()
                    break F;
                } 
                l.ungetN(0)
            
            case '{':
                c = l.Next();
                if rune(c) == '$' {
                    beginString()
                    break F;
                } 
                l.ungetN(0)

            case '\\':
                c = l.Next();
        }

        c = l.Next()
    }

<STRING>\"                                  fmt.Println("\""); popState()
<STRING>\{\$                                fmt.Printf("T_CURLY_OPEN: %q\n", l.ungetN(1));pushState(PHP)
<STRING>\$\{                                fmt.Printf("T_DOLLAR_OPEN_CURLY_BRACES: %q\n", l.TokenBytes(nil));pushState(STRING_VAR_NAME)
<STRING>\$                                  l.ungetN(1);begin(STRING_VAR)
<STRING>.
    F1:for {
        if c == -1 {
            break;
        }

        switch c {
            case '"' : 
                fmt.Printf("T_ENCAPSED_AND_WHITESPACE: %s\n", l.TokenBytes(nil));
                break F1;
            
            case '$':
                c = l.Next();
                if rune(c) == '{' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c >= '\u007f' && c <= 'ÿ' {
                    l.ungetN(1)
                    tb := l.TokenBytes(nil)
                    fmt.Printf("T_ENCAPSED_AND_WHITESPACE: %s\n", tb[:len(tb)-1]);
                    break F1;
                } 
                l.ungetN(0)
            
            case '{':
                c = l.Next();
                if rune(c) == '$' {
                    l.ungetN(1)
                    tb := l.TokenBytes(nil)
                    fmt.Printf("T_ENCAPSED_AND_WHITESPACE: %s\n", tb[:len(tb)-1]);
                    break F1;
                } 
                l.ungetN(0)

            case '\\':
                c = l.Next();
        }

        c = l.Next()
    }

<STRING_VAR>\${VAR_NAME}                    fmt.Printf("T_VARIABLE: %q\n", l.TokenBytes(nil));
<STRING_VAR>->{VAR_NAME}                    fmt.Printf("T_OBJECT_OPERATOR: %q\n", l.ungetN(len(l.TokenBytes(nil))-2));
<STRING_VAR>{VAR_NAME}                      fmt.Printf("T_STRING: %q\n", l.TokenBytes(nil));begin(STRING)
<STRING_VAR>\[                              fmt.Println("["); begin(STRING_VAR_INDEX)
<STRING_VAR>.                               l.ungetN(1);begin(STRING)

<STRING_VAR_INDEX>{LNUM}                    fmt.Printf("T_NUM_STRING: %q\n", l.TokenBytes(nil));
<STRING_VAR_INDEX>\${VAR_NAME}              fmt.Printf("T_VARIABLE: %q\n", l.TokenBytes(nil));
<STRING_VAR_INDEX>{VAR_NAME}                fmt.Printf("T_STRING: %q\n", l.TokenBytes(nil));
<STRING_VAR_INDEX>\]                        fmt.Println("\"]\""); begin(STRING)
<STRING_VAR_INDEX>[ \n\r\t\\'#]             fmt.Printf("T_ENCAPSED_AND_WHITESPACE: %q\n", l.ungetN(1));begin(STRING)
<STRING_VAR_INDEX>{OPERATORS}               fmt.Printf("%q\n", l.TokenBytes(nil));
<STRING_VAR_INDEX>.                         fmt.Printf("%q\n", l.TokenBytes(nil));

<STRING_VAR_NAME>{VAR_NAME}[\[\}]           fmt.Printf("T_STRING_VARNAME: %q\n", l.ungetN(1));popState();pushState(PHP)
<STRING_VAR_NAME>.                          l.ungetN(1);popState();pushState(PHP)

<PHP>.                                      fmt.Printf("other: %q\n", l.TokenBytes(nil))

%%
    if c, ok := l.Abort(); ok { return int(c) }
    goto yyAction
}

func main() {
    l := newLexer(os.Stdin, os.Stdout, "file.name")
    l.Lex();
}