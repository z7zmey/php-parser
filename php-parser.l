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

func (l *lexer) Lex() int { // Lex(lval *yySymType)
  c := l.Enter()

%}

%s PHP STRING

%yyb last == '\n' || last = '\0'
%yyt sc
%yyc c
%yyn c = l.Next()
%yym l.Mark()

D           [0-9]+
NC          ([^\\\$\"\{])
NCH         [^a-zA-Z_\x7f-\xff]
ENSCAPED    ([\\].)
DOLLAR_E    ([\$]{ENSCAPED})
DOLLAR_N    ([\$][^a-zA-Z_\x7f-\xff\\\$\"\{])
CURVE_E     ([\{]{ENSCAPED})
CURVE       ([\{][^\$\"])
ALLOWED     ({NC}|{ENSCAPED}|{DOLLAR_E}|{DOLLAR_N}|{CURVE_E}|{CURVE})
STR_END     ([\{\$]?[\"])?
STR         {ALLOWED}+{ALLOWED}*
TPL         {STR}{STR_END}
VAR         [a-zA-Z_\x7f-\xff][a-zA-Z0-9_\x7f-\xff]*

%%
    c = l.Rule0()
    // ([\$]{NCH})*

<INITIAL>[ \t\n\r]+
<INITIAL>.
<INITIAL>\<\?|\<\?php                       fmt.Println("T_OPEN_TAG");begin(PHP)
<INITIAL>\<\?=                              fmt.Println("T_OPEN_TAG_WITH_ECHO");begin(PHP)

<PHP>[ \t\n\r]+                             fmt.Println("T_WHITESPACE")
<PHP>\?\>                                   fmt.Println("T_CLOSE_TAG");begin(INITIAL)

<PHP>[\"]{STR}*[\{\$]?[\"]                  fmt.Printf("T_CONSTANT_ENCAPSED_STRING: %s\n", l.TokenBytes(nil));
<PHP>[\']([^\\\']*([\\][\'])*)*[\']         fmt.Printf("T_CONSTANT_ENCAPSED_STRING: %s\n", l.TokenBytes(nil));
<PHP>[\"]                                   fmt.Println("\"");begin(STRING)
<PHP>.                                      fmt.Printf("other: %q\n", l.TokenBytes(nil))

<STRING>\"                                  fmt.Println("\""); begin(PHP)
<STRING>\{                                  fmt.Printf("T_CURLY_OPEN: %q\n", l.TokenBytes(nil));
<STRING>\$\{                                fmt.Printf("T_DOLLAR_OPEN_CURLY_BRACES: %q\n", l.TokenBytes(nil))
<STRING>\${VAR}?                            fmt.Printf("T_VARIABLE: %q\n", l.TokenBytes(nil))
<STRING>{TPL}                               fmt.Printf("T_ENCAPSED_AND_WHITESPACE: %q\n", l.unget('"'));

<PHP>\${VAR}         fmt.Println("T_VARIABLE")

%%
    if c, ok := l.Abort(); ok { return int(c) }
    goto yyAction
}

func main() {
    l := newLexer(os.Stdin, os.Stdout, "file.name")
    l.Lex();
}