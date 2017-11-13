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
)

type lexer struct {
    *lex.Lexer
}

func begin(cond int) {
    sc = cond
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

%s PHP STRING STRING_VAR STRING_VAR_INDEX

%yyb last == '\n' || last = '\0'
%yyt sc
%yyc c
%yyn c = l.Next()
%yym l.Mark()

D           [0-9]+
NC          ([^\\\$\"\{])
ENSCAPED    ([\\].)
DOLLAR      ([\$]+{ENSCAPED})|([\$]+[^a-zA-Z_\x7f-\xff\$\"\{])
CURVE       ([\{]+{ENSCAPED})|([\{]+[^\{\$\"])
ALLOWED     ({NC}|{ENSCAPED}|{DOLLAR}|{CURVE})
STR         {ALLOWED}*
STR_END     [\{\$]?[\"]
VAR_NAME    [a-zA-Z_\x7f-\xff][a-zA-Z0-9_\x7f-\xff]*

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
    F:for {
        c := l.Next()
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
                    if binPrefix {
                        l.ungetN(len(l.TokenBytes(nil))-2)
                        fmt.Println("b\""); begin(STRING)
                    } else {
                        l.ungetN(len(l.TokenBytes(nil))-1)
                        fmt.Println("\""); begin(STRING)
                    }
                    break F;
                } else {
                    l.ungetN(0)
                }
            case '{':
                c = l.Next();
                if rune(c) == '$' {
                    if binPrefix {
                        l.ungetN(len(l.TokenBytes(nil))-2)
                        fmt.Println("b\""); begin(STRING)
                    } else {
                        l.ungetN(len(l.TokenBytes(nil))-1)
                        fmt.Println("\""); begin(STRING)
                    }
                    break F;
                } else {
                    l.ungetN(0)
                }
            case '\\':
                c = l.Next();
        }
    }

<PHP>.                                      fmt.Printf("other: %q\n", l.TokenBytes(nil))

<STRING>\"                                  fmt.Println("\""); begin(PHP)
<STRING>\{\$                                fmt.Printf("T_CURLY_OPEN: %q\n", l.ungetN(1));
<STRING>\$\{                                fmt.Printf("T_DOLLAR_OPEN_CURLY_BRACES: %q\n", l.TokenBytes(nil))
<STRING>\$                                  l.ungetN(1);begin(STRING_VAR)

<STRING>{STR}{STR_END}                      fmt.Printf("T_ENCAPSED_AND_WHITESPACE1: %q\n", l.ungetN(1));
<STRING>{STR}[\{]+[\$]                      fmt.Printf("T_ENCAPSED_AND_WHITESPACE2: %q\n", l.ungetN(2));
<STRING>{STR}[\$]+[\{]                      fmt.Printf("T_ENCAPSED_AND_WHITESPACE3: %q\n", l.ungetN(2));
<STRING>{STR}[^\{][\$]+[a-zA-Z_\x7f-\xff]   fmt.Printf("T_ENCAPSED_AND_WHITESPACE4: %q\n", l.ungetN(2));

<STRING_VAR>\${VAR_NAME}                    fmt.Printf("T_VARIABLE: %q\n", l.TokenBytes(nil));
<STRING_VAR>->{VAR_NAME}                    fmt.Printf("T_OBJECT_OPERATOR: %q\n", l.ungetN(len(l.TokenBytes(nil))-2));
<STRING_VAR>{VAR_NAME}                      fmt.Printf("T_STRING: %q\n", l.TokenBytes(nil));begin(STRING)
<STRING_VAR>\[                              fmt.Println("["); begin(STRING_VAR_INDEX)
<STRING_VAR>.                               l.ungetN(1);begin(STRING)

<STRING_VAR_INDEX>{D}                       fmt.Printf("T_NUM_STRING: %q\n", l.TokenBytes(nil));
<STRING_VAR_INDEX>\${VAR_NAME}              fmt.Printf("T_VARIABLE: %q\n", l.TokenBytes(nil));
<STRING_VAR_INDEX>{VAR_NAME}                fmt.Printf("T_STRING: %q\n", l.TokenBytes(nil));
<STRING_VAR_INDEX>\]                        fmt.Println("\"]\""); begin(STRING)
<STRING_VAR_INDEX>.                         fmt.Printf("%q\n", l.TokenBytes(nil));

<PHP>\${VAR_NAME}                           fmt.Println("T_VARIABLE")

%%
    if c, ok := l.Abort(); ok { return int(c) }
    goto yyAction
}

func main() {
    l := newLexer(os.Stdin, os.Stdout, "file.name")
    l.Lex();
}