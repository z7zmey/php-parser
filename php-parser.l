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
)


type lexer struct {
    *lex.Lexer
}

var stateStack = []int{PHP}

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

%s PHP STRING STRING_VAR STRING_VAR_INDEX STRING_VAR_NAME

%yyb last == '\n' || last = '\0'
%yyt sc
%yyc c
%yyn c = l.Next()
%yym l.Mark()

D           ([0-9]+)
VAR_NAME    [a-zA-Z_\x7f-\xff][a-zA-Z0-9_\x7f-\xff]*
OPERATORS   [;:,.\[\]()|\/\^&\+-*=%!~$<>?@]

%%
    c = l.Rule0()
    // ([\$]{NCH})*

<INITIAL>[ \t\n\r]+
<INITIAL>.
<INITIAL>\<\?|\<\?php                       fmt.Println("T_OPEN_TAG");begin(PHP)
<INITIAL>\<\?=                              fmt.Println("T_OPEN_TAG_WITH_ECHO");begin(PHP)

<PHP>[ \t\n\r]+                             fmt.Println("T_WHITESPACE")
<PHP>\?\>                                   fmt.Println("T_CLOSE_TAG");begin(INITIAL)

<PHP>[\']([^\\\']*([\\][\'])*)*[\']         fmt.Printf("T_CONSTANT_ENCAPSED_STRING: %s\n", l.TokenBytes(nil));
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

<STRING_VAR_INDEX>{D}                       fmt.Printf("T_NUM_STRING: %q\n", l.TokenBytes(nil));
<STRING_VAR_INDEX>\${VAR_NAME}              fmt.Printf("T_VARIABLE: %q\n", l.TokenBytes(nil));
<STRING_VAR_INDEX>{VAR_NAME}                fmt.Printf("T_STRING: %q\n", l.TokenBytes(nil));
<STRING_VAR_INDEX>\]                        fmt.Println("\"]\""); begin(STRING)
<STRING_VAR_INDEX>[ \n\r\t\\'#]             fmt.Printf("T_ENCAPSED_AND_WHITESPACE: %q\n", l.ungetN(1));begin(STRING)
<STRING_VAR_INDEX>{OPERATORS}               fmt.Printf("%q\n", l.TokenBytes(nil));
<STRING_VAR_INDEX>.                         fmt.Printf("%q\n", l.TokenBytes(nil));

<STRING_VAR_NAME>{VAR_NAME}[\[\}]           fmt.Printf("T_STRING_VARNAME: %q\n", l.ungetN(1));popState();pushState(PHP)
<STRING_VAR_NAME>.                          l.ungetN(1);popState();pushState(PHP)

<PHP>\}                                     fmt.Println("}"); popState();
<PHP>\${VAR_NAME}                           fmt.Printf("T_VARIABLE: %q\n", l.TokenBytes(nil))
<PHP>.                                      fmt.Printf("other: %q\n", l.TokenBytes(nil))

%%
    if c, ok := l.Abort(); ok { return int(c) }
    goto yyAction
}

func main() {
    l := newLexer(os.Stdin, os.Stdout, "file.name")
    l.Lex();
}