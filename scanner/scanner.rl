package scanner

import (
    "fmt"
)

%%{ 
    machine lexer;
    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;
}%%

func NewLexer(data []byte) *Lexer {
    lex := &Lexer{ 
        data: data,
        pe: len(data),
        stack: make([]int, 0),

        TokenPool: &TokenPool{},
        NewLines: NewLines{make([]int, 0, 128)},
    }
    %% write init;
    return lex
}

func (lex *Lexer) Lex(lval Lval) int {
    token := lex.TokenPool.Get()
    token.SkippedTokens = token.SkippedTokens[:0]
    eof := lex.pe
    var tok TokenID

    lblStart := 0
    lblEnd   := 0

    _, _ = lblStart, lblEnd

    %%{ 
        action heredoc_lbl_start {lblStart = lex.p}
        action heredoc_lbl_end   {lblEnd = lex.p}

        action is_not_heredoc_end { lex.isNotHeredocEnd(lex.p) }
        action is_not_comment_end { lex.isNotPhpCloseToken() && lex.isNotNewLine()  }
        action is_not_heredoc_end_or_var { lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() }
        action is_not_string_end_or_var { lex.isNotStringEnd('"') && lex.isNotStringVar() }
        action is_not_backqoute_end_or_var { lex.isNotStringEnd('`') && lex.isNotStringVar() }

        newline = ('\r\n' >(nl, 1) | '\r' >(nl, 0) | '\n' >(nl, 0)) %{lex.NewLines.Append(lex.p);};
        any_line = any | newline;
        whitespace = [\t\v\f ];
        whitespace_line = [\t\v\f ] | newline;

        lnum = [0-9]+;
        dnum = ( [0-9]* "." [0-9]+ ) | ( [0-9]+ "." [0-9]* );
        hnum = '0x' [0-9a-fA-F]+;
        bnum = '0b' [01]+;

        exponent_dnum = (lnum | dnum) ('e'|'E') ('+'|'-')? lnum;
        varname       = /[a-zA-Z_\x7f-\xff][a-zA-Z0-9_\x7f-\xff]*/;
        varname_first = /[a-zA-Z_\x7f-\xff]/;
        heredoc_label = varname >heredoc_lbl_start %heredoc_lbl_end;
        operators     = ';'|':'|','|'.'|'['|']'|'('|')'|'|'|'/'|'^'|'&'|'+'|'-'|'*'|'='|'%'|'!'|'~'|'$'|'<'|'>'|'?'|'@';
        
        prepush { lex.growCallStack(); }

        constant_string =
            start: (
                "'"        -> qoute
                | "b"i? '"' -> double_qoute
            ),
            qoute: (
                (any - [\\'\r\n]) -> qoute
                | "\r" @{if lex.p+1 != eof && lex.data[lex.p+1] != '\n' {lex.NewLines.Append(lex.p)}} -> qoute
                | "\n" @{lex.NewLines.Append(lex.p)} -> qoute
                | "\\"                               -> qoute_any
                | "'"                                -> final
            ),
            qoute_any: (
                any_line -> qoute
            ),
            double_qoute: (
                (any - [\\"${\r\n])  -> double_qoute
                | "\r" @{if lex.p+1 != eof && lex.data[lex.p+1] != '\n' {lex.NewLines.Append(lex.p)}} -> double_qoute
                | "\n" @{lex.NewLines.Append(lex.p)} -> double_qoute
                | "\\"                               -> double_qoute_any
                | '"'                                -> final
                | '$'                                -> double_qoute_nonvarname
                | '{'                                -> double_qoute_nondollar
            ),
            double_qoute_any: (
                any_line -> double_qoute
            ),
            double_qoute_nondollar: (
                '"'    -> final
                | "\\"               -> double_qoute_any
                | [^$\\"] -> double_qoute
            ),
            double_qoute_nonvarname: (
                '"'                      -> final
                | "\\"                   -> double_qoute_any
                | /[^"\\{a-zA-Z_\x7f-\xff]/ -> double_qoute
            );

        main := |*
            any_line+ -- '<?' => {
                lex.ungetStr("<")
                lval.Token(lex.prepareToken(token))
                tok = T_INLINE_HTML;
                fbreak;
            };
            '<?' => {
                lex.addSkippedToken(token, SkippedTokenTypeToken, lex.ts, lex.te)
                fnext php;
            };
            '<?php'i ( [ \t] | newline ) => {
                lex.ungetCnt(lex.te - lex.ts - 5)
                lex.addSkippedToken(token, SkippedTokenTypeToken, lex.ts, lex.ts+5)
                fnext php;
            };
            '<?='i => {
                lval.Token(lex.prepareToken(token));
                tok = T_ECHO;
                fnext php;
                fbreak;
            };
        *|;

        php := |*
            whitespace_line*                   => {lex.addSkippedToken(token, SkippedTokenTypeWhitespace, lex.ts, lex.te)};
            '?>' newline?                      => {lval.Token(lex.prepareToken(token)); tok = TokenID(int(';')); fnext main; fbreak;};
            ';' whitespace_line* '?>' newline? => {lval.Token(lex.prepareToken(token)); tok = TokenID(int(';')); fnext main; fbreak;};

            (dnum | exponent_dnum)          => {lval.Token(lex.prepareToken(token)); tok = T_DNUMBER; fbreak;};
            bnum => {
                firstNum := 2
                for i := lex.ts + 2; i < lex.te; i++ {
                    if lex.data[i] == '0' {
                        firstNum++
                    }
                }

                if lex.te - lex.ts - firstNum < 64 {
                    lval.Token(lex.prepareToken(token)); tok = T_LNUMBER; fbreak;
                }
                lval.Token(lex.prepareToken(token)); tok = T_DNUMBER; fbreak;
            };
            lnum => {
                if lex.te - lex.ts < 20 {
                    lval.Token(lex.prepareToken(token)); tok = T_LNUMBER; fbreak;
                }
                lval.Token(lex.prepareToken(token)); tok = T_DNUMBER; fbreak;
            };
            hnum => {
                firstNum := lex.ts + 2
                for i := lex.ts + 2; i < lex.te; i++ {
                    if lex.data[i] == '0' {
                        firstNum++
                    }
                }

                length := lex.te - firstNum
                if length < 16 || (length == 16 && lex.data[firstNum] <= '7') {
                    lval.Token(lex.prepareToken(token)); tok = T_LNUMBER; fbreak;
                } 
                lval.Token(lex.prepareToken(token)); tok = T_DNUMBER; fbreak;
            };

            'abstract'i                       => {lval.Token(lex.prepareToken(token)); tok = T_ABSTRACT; fbreak;};
            'array'i                          => {lval.Token(lex.prepareToken(token)); tok = T_ARRAY; fbreak;};
            'as'i                             => {lval.Token(lex.prepareToken(token)); tok = T_AS; fbreak;};
            'break'i                          => {lval.Token(lex.prepareToken(token)); tok = T_BREAK; fbreak;};
            'callable'i                       => {lval.Token(lex.prepareToken(token)); tok = T_CALLABLE; fbreak;};
            'case'i                           => {lval.Token(lex.prepareToken(token)); tok = T_CASE; fbreak;};
            'catch'i                          => {lval.Token(lex.prepareToken(token)); tok = T_CATCH; fbreak;};
            'class'i                          => {lval.Token(lex.prepareToken(token)); tok = T_CLASS; fbreak;};
            'clone'i                          => {lval.Token(lex.prepareToken(token)); tok = T_CLONE; fbreak;};
            'const'i                          => {lval.Token(lex.prepareToken(token)); tok = T_CONST; fbreak;};
            'continue'i                       => {lval.Token(lex.prepareToken(token)); tok = T_CONTINUE; fbreak;};
            'declare'i                        => {lval.Token(lex.prepareToken(token)); tok = T_DECLARE; fbreak;};
            'default'i                        => {lval.Token(lex.prepareToken(token)); tok = T_DEFAULT; fbreak;};
            'do'i                             => {lval.Token(lex.prepareToken(token)); tok = T_DO; fbreak;};
            'echo'i                           => {lval.Token(lex.prepareToken(token)); tok = T_ECHO; fbreak;};
            'else'i                           => {lval.Token(lex.prepareToken(token)); tok = T_ELSE; fbreak;};
            'elseif'i                         => {lval.Token(lex.prepareToken(token)); tok = T_ELSEIF; fbreak;};
            'empty'i                          => {lval.Token(lex.prepareToken(token)); tok = T_EMPTY; fbreak;};
            'enddeclare'i                     => {lval.Token(lex.prepareToken(token)); tok = T_ENDDECLARE; fbreak;};
            'endfor'i                         => {lval.Token(lex.prepareToken(token)); tok = T_ENDFOR; fbreak;};
            'endforeach'i                     => {lval.Token(lex.prepareToken(token)); tok = T_ENDFOREACH; fbreak;};
            'endif'i                          => {lval.Token(lex.prepareToken(token)); tok = T_ENDIF; fbreak;};
            'endswitch'i                      => {lval.Token(lex.prepareToken(token)); tok = T_ENDSWITCH; fbreak;};
            'endwhile'i                       => {lval.Token(lex.prepareToken(token)); tok = T_ENDWHILE; fbreak;};
            'eval'i                           => {lval.Token(lex.prepareToken(token)); tok = T_EVAL; fbreak;};
            'exit'i | 'die'i                  => {lval.Token(lex.prepareToken(token)); tok = T_EXIT; fbreak;};
            'extends'i                        => {lval.Token(lex.prepareToken(token)); tok = T_EXTENDS; fbreak;};
            'final'i                          => {lval.Token(lex.prepareToken(token)); tok = T_FINAL; fbreak;};
            'finally'i                        => {lval.Token(lex.prepareToken(token)); tok = T_FINALLY; fbreak;};
            'for'i                            => {lval.Token(lex.prepareToken(token)); tok = T_FOR; fbreak;};
            'foreach'i                        => {lval.Token(lex.prepareToken(token)); tok = T_FOREACH; fbreak;};
            'function'i | 'cfunction'i        => {lval.Token(lex.prepareToken(token)); tok = T_FUNCTION; fbreak;};
            'global'i                         => {lval.Token(lex.prepareToken(token)); tok = T_GLOBAL; fbreak;};
            'goto'i                           => {lval.Token(lex.prepareToken(token)); tok = T_GOTO; fbreak;};
            'if'i                             => {lval.Token(lex.prepareToken(token)); tok = T_IF; fbreak;};
            'isset'i                          => {lval.Token(lex.prepareToken(token)); tok = T_ISSET; fbreak;};
            'implements'i                     => {lval.Token(lex.prepareToken(token)); tok = T_IMPLEMENTS; fbreak;};
            'instanceof'i                     => {lval.Token(lex.prepareToken(token)); tok = T_INSTANCEOF; fbreak;};
            'insteadof'i                      => {lval.Token(lex.prepareToken(token)); tok = T_INSTEADOF; fbreak;};
            'interface'i                      => {lval.Token(lex.prepareToken(token)); tok = T_INTERFACE; fbreak;};
            'list'i                           => {lval.Token(lex.prepareToken(token)); tok = T_LIST; fbreak;};
            'namespace'i                      => {lval.Token(lex.prepareToken(token)); tok = T_NAMESPACE; fbreak;};
            'private'i                        => {lval.Token(lex.prepareToken(token)); tok = T_PRIVATE; fbreak;};
            'public'i                         => {lval.Token(lex.prepareToken(token)); tok = T_PUBLIC; fbreak;};
            'print'i                          => {lval.Token(lex.prepareToken(token)); tok = T_PRINT; fbreak;};
            'protected'i                      => {lval.Token(lex.prepareToken(token)); tok = T_PROTECTED; fbreak;};
            'return'i                         => {lval.Token(lex.prepareToken(token)); tok = T_RETURN; fbreak;};
            'static'i                         => {lval.Token(lex.prepareToken(token)); tok = T_STATIC; fbreak;};
            'switch'i                         => {lval.Token(lex.prepareToken(token)); tok = T_SWITCH; fbreak;};
            'throw'i                          => {lval.Token(lex.prepareToken(token)); tok = T_THROW; fbreak;};
            'trait'i                          => {lval.Token(lex.prepareToken(token)); tok = T_TRAIT; fbreak;};
            'try'i                            => {lval.Token(lex.prepareToken(token)); tok = T_TRY; fbreak;};
            'unset'i                          => {lval.Token(lex.prepareToken(token)); tok = T_UNSET; fbreak;};
            'use'i                            => {lval.Token(lex.prepareToken(token)); tok = T_USE; fbreak;};
            'var'i                            => {lval.Token(lex.prepareToken(token)); tok = T_VAR; fbreak;};
            'while'i                          => {lval.Token(lex.prepareToken(token)); tok = T_WHILE; fbreak;};
            'yield'i whitespace_line+ 'from'i => {lval.Token(lex.prepareToken(token)); tok = T_YIELD_FROM; fbreak;};
            'yield'i                          => {lval.Token(lex.prepareToken(token)); tok = T_YIELD; fbreak;};
            'include'i                        => {lval.Token(lex.prepareToken(token)); tok = T_INCLUDE; fbreak;};
            'include_once'i                   => {lval.Token(lex.prepareToken(token)); tok = T_INCLUDE_ONCE; fbreak;};
            'require'i                        => {lval.Token(lex.prepareToken(token)); tok = T_REQUIRE; fbreak;};
            'require_once'i                   => {lval.Token(lex.prepareToken(token)); tok = T_REQUIRE_ONCE; fbreak;};
            '__CLASS__'i                      => {lval.Token(lex.prepareToken(token)); tok = T_CLASS_C; fbreak;};
            '__DIR__'i                        => {lval.Token(lex.prepareToken(token)); tok = T_DIR; fbreak;};
            '__FILE__'i                       => {lval.Token(lex.prepareToken(token)); tok = T_FILE; fbreak;};
            '__FUNCTION__'i                   => {lval.Token(lex.prepareToken(token)); tok = T_FUNC_C; fbreak;};
            '__LINE__'i                       => {lval.Token(lex.prepareToken(token)); tok = T_LINE; fbreak;};
            '__NAMESPACE__'i                  => {lval.Token(lex.prepareToken(token)); tok = T_NS_C; fbreak;};
            '__METHOD__'i                     => {lval.Token(lex.prepareToken(token)); tok = T_METHOD_C; fbreak;};
            '__TRAIT__'i                      => {lval.Token(lex.prepareToken(token)); tok = T_TRAIT_C; fbreak;};
            '__halt_compiler'i                => {lval.Token(lex.prepareToken(token)); tok = T_HALT_COMPILER; fnext halt_compiller_open_parenthesis; fbreak;};
            'new'i                            => {lval.Token(lex.prepareToken(token)); tok = T_NEW; fbreak;};
            'and'i                            => {lval.Token(lex.prepareToken(token)); tok = T_LOGICAL_AND; fbreak;};
            'or'i                             => {lval.Token(lex.prepareToken(token)); tok = T_LOGICAL_OR; fbreak;};
            'xor'i                            => {lval.Token(lex.prepareToken(token)); tok = T_LOGICAL_XOR; fbreak;};
            '\\'                              => {lval.Token(lex.prepareToken(token)); tok = T_NS_SEPARATOR; fbreak;};
            '...'                             => {lval.Token(lex.prepareToken(token)); tok = T_ELLIPSIS; fbreak;};
            '::'                              => {lval.Token(lex.prepareToken(token)); tok = T_PAAMAYIM_NEKUDOTAYIM; fbreak;};
            '&&'                              => {lval.Token(lex.prepareToken(token)); tok = T_BOOLEAN_AND; fbreak;};
            '||'                              => {lval.Token(lex.prepareToken(token)); tok = T_BOOLEAN_OR; fbreak;};
            '&='                              => {lval.Token(lex.prepareToken(token)); tok = T_AND_EQUAL; fbreak;};
            '|='                              => {lval.Token(lex.prepareToken(token)); tok = T_OR_EQUAL; fbreak;};
            '.='                              => {lval.Token(lex.prepareToken(token)); tok = T_CONCAT_EQUAL; fbreak;};
            '*='                              => {lval.Token(lex.prepareToken(token)); tok = T_MUL_EQUAL; fbreak;};
            '**='                             => {lval.Token(lex.prepareToken(token)); tok = T_POW_EQUAL; fbreak;};
            '/='                              => {lval.Token(lex.prepareToken(token)); tok = T_DIV_EQUAL; fbreak;};
            '+='                              => {lval.Token(lex.prepareToken(token)); tok = T_PLUS_EQUAL; fbreak;};
            '-='                              => {lval.Token(lex.prepareToken(token)); tok = T_MINUS_EQUAL; fbreak;};
            '^='                              => {lval.Token(lex.prepareToken(token)); tok = T_XOR_EQUAL; fbreak;};
            '%='                              => {lval.Token(lex.prepareToken(token)); tok = T_MOD_EQUAL; fbreak;};
            '--'                              => {lval.Token(lex.prepareToken(token)); tok = T_DEC; fbreak;};
            '++'                              => {lval.Token(lex.prepareToken(token)); tok = T_INC; fbreak;};
            '=>'                              => {lval.Token(lex.prepareToken(token)); tok = T_DOUBLE_ARROW; fbreak;};
            '<=>'                             => {lval.Token(lex.prepareToken(token)); tok = T_SPACESHIP; fbreak;};
            '!=' | '<>'                       => {lval.Token(lex.prepareToken(token)); tok = T_IS_NOT_EQUAL; fbreak;};
            '!=='                             => {lval.Token(lex.prepareToken(token)); tok = T_IS_NOT_IDENTICAL; fbreak;};
            '=='                              => {lval.Token(lex.prepareToken(token)); tok = T_IS_EQUAL; fbreak;};
            '==='                             => {lval.Token(lex.prepareToken(token)); tok = T_IS_IDENTICAL; fbreak;};
            '<<='                             => {lval.Token(lex.prepareToken(token)); tok = T_SL_EQUAL; fbreak;};
            '>>='                             => {lval.Token(lex.prepareToken(token)); tok = T_SR_EQUAL; fbreak;};
            '>='                              => {lval.Token(lex.prepareToken(token)); tok = T_IS_GREATER_OR_EQUAL; fbreak;};
            '<='                              => {lval.Token(lex.prepareToken(token)); tok = T_IS_SMALLER_OR_EQUAL; fbreak;};
            '**'                              => {lval.Token(lex.prepareToken(token)); tok = T_POW; fbreak;};
            '<<'                              => {lval.Token(lex.prepareToken(token)); tok = T_SL; fbreak;};
            '>>'                              => {lval.Token(lex.prepareToken(token)); tok = T_SR; fbreak;};
            '??'                              => {lval.Token(lex.prepareToken(token)); tok = T_COALESCE; fbreak;};

            '(' whitespace* 'array'i whitespace* ')'                     => {lval.Token(lex.prepareToken(token)); tok = T_ARRAY_CAST; fbreak;};
            '(' whitespace* ('bool'i|'boolean'i) whitespace* ')'         => {lval.Token(lex.prepareToken(token)); tok = T_BOOL_CAST; fbreak;};
            '(' whitespace* ('real'i|'double'i|'float'i) whitespace* ')' => {lval.Token(lex.prepareToken(token)); tok = T_DOUBLE_CAST; fbreak;};
            '(' whitespace* ('int'i|'integer'i) whitespace* ')'          => {lval.Token(lex.prepareToken(token)); tok = T_INT_CAST; fbreak;};
            '(' whitespace* 'object'i whitespace* ')'                    => {lval.Token(lex.prepareToken(token)); tok = T_OBJECT_CAST; fbreak;};
            '(' whitespace* ('string'i|'binary'i) whitespace* ')'        => {lval.Token(lex.prepareToken(token)); tok = T_STRING_CAST; fbreak;};
            '(' whitespace* 'unset'i whitespace* ')'                     => {lval.Token(lex.prepareToken(token)); tok = T_UNSET_CAST; fbreak;};

            ('#' | '//') any_line* when is_not_comment_end => {
                lex.ungetStr("?>")
                lex.addSkippedToken(token, SkippedTokenTypeComment, lex.ts, lex.te)
            };
            '/*' any_line* :>> '*/' {
                isDocComment := false;
                if lex.te - lex.ts > 4 && string(lex.data[lex.ts:lex.ts+3]) == "/**" {
                    isDocComment = true;
                }
                lex.addSkippedToken(token, SkippedTokenTypeComment, lex.ts, lex.te)

                if isDocComment {
                    lex.PhpDocComment = string(lex.data[lex.ts:lex.te])
                }
            };

            operators => {
                // rune, _ := utf8.DecodeRune(lex.data[lex.ts:lex.te]);
                // tok = TokenID(Rune2Class(rune));
                lval.Token(lex.prepareToken(token));
                tok = TokenID(int(lex.data[lex.ts]));
                fbreak;
            };

            "{"          => { lval.Token(lex.prepareToken(token)); tok = TokenID(int('{')); lex.call(ftargs, fentry(php)); goto _out; };
            "}"          => { lval.Token(lex.prepareToken(token)); tok = TokenID(int('}')); lex.ret(1); lex.PhpDocComment = ""; goto _out;};
            "$" varname  => { lval.Token(lex.prepareToken(token)); tok = T_VARIABLE; fbreak; };
            varname      => { lval.Token(lex.prepareToken(token)); tok = T_STRING;   fbreak; };

            "->"         => { lval.Token(lex.prepareToken(token)); tok = T_OBJECT_OPERATOR; fnext property; fbreak; };

            constant_string => {
                lval.Token(lex.prepareToken(token));
                tok = T_CONSTANT_ENCAPSED_STRING;
                fbreak;
            };

            "b"i? "<<<" [ \t]* ( heredoc_label | ("'" heredoc_label "'") | ('"' heredoc_label '"') ) newline  => {
                lex.heredocLabel = lex.data[lblStart:lblEnd]
                lval.Token(lex.prepareToken(token));
                tok = T_START_HEREDOC;

                if lex.isHeredocEnd(lex.p+1) {
                    fnext heredoc_end;
                } else if lex.data[lblStart-1] == '\'' {
                    fnext nowdoc;
                } else {
                    fnext heredoc;
                }
                fbreak;
            };
            "`" => {lval.Token(lex.prepareToken(token)); tok = TokenID(int('`')); fnext backqote; fbreak;};
            '"' => {lval.Token(lex.prepareToken(token)); tok = TokenID(int('"')); fnext template_string; fbreak;};

            any_line => {
                c := lex.data[lex.p]
                lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c));
            };
        *|;

        property := |*
            whitespace_line* => {lex.addSkippedToken(token, SkippedTokenTypeWhitespace, lex.ts, lex.te)};
            "->"             => {lval.Token(lex.prepareToken(token)); tok = T_OBJECT_OPERATOR; fbreak;};
            varname          => {lval.Token(lex.prepareToken(token)); tok = T_STRING; fnext php; fbreak;};
            any              => {lex.ungetCnt(1); fgoto php;};
        *|;

        nowdoc := |*
            any_line* when is_not_heredoc_end => {
                lval.Token(lex.prepareToken(token));
                tok = T_ENCAPSED_AND_WHITESPACE;
                fnext heredoc_end;
                fbreak;
            };
        *|;
        
        heredoc := |*
            "{$" => {lex.ungetCnt(1); lval.Token(lex.prepareToken(token)); tok = T_CURLY_OPEN; lex.call(ftargs, fentry(php)); goto _out;};
            "${" => {lval.Token(lex.prepareToken(token)); tok = T_DOLLAR_OPEN_CURLY_BRACES; lex.call(ftargs, fentry(string_var_name)); goto _out;};
            "$"  => {lex.ungetCnt(1); fcall string_var;};
            any_line* when is_not_heredoc_end_or_var => {
                lval.Token(lex.prepareToken(token));
                tok = T_ENCAPSED_AND_WHITESPACE;

                if lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
                    fnext heredoc_end;
                }
                fbreak;
            };
        *|;
        
        backqote := |*
            "{$" => {lex.ungetCnt(1); lval.Token(lex.prepareToken(token)); tok = T_CURLY_OPEN; lex.call(ftargs, fentry(php)); goto _out;};
            "${" => {lval.Token(lex.prepareToken(token)); tok = T_DOLLAR_OPEN_CURLY_BRACES; lex.call(ftargs, fentry(string_var_name)); goto _out;};
            "$"  => {lex.ungetCnt(1); fcall string_var;};
            '`'  => {lval.Token(lex.prepareToken(token)); tok = TokenID(int('`')); fnext php; fbreak;};
            any_line* when is_not_backqoute_end_or_var => {
                lval.Token(lex.prepareToken(token));
                tok = T_ENCAPSED_AND_WHITESPACE;
                fbreak;
            };
        *|;
        
        template_string := |*
            "{$" => {lex.ungetCnt(1); lval.Token(lex.prepareToken(token)); tok = T_CURLY_OPEN; lex.call(ftargs, fentry(php)); goto _out;};
            "${" => {lval.Token(lex.prepareToken(token)); tok = T_DOLLAR_OPEN_CURLY_BRACES; lex.call(ftargs, fentry(string_var_name)); goto _out;};
            "$"  => {lex.ungetCnt(1); fcall string_var;};
            '"'  => {lval.Token(lex.prepareToken(token)); tok = TokenID(int('"')); fnext php; fbreak;};
            any_line* when is_not_string_end_or_var => {
                lval.Token(lex.prepareToken(token));
                tok = T_ENCAPSED_AND_WHITESPACE;
                fbreak;
            };
        *|;

        heredoc_end := |*
            varname -- ";" => {
                lval.Token(lex.prepareToken(token));
                tok = T_END_HEREDOC;
                fnext php;
                fbreak;
            };
            varname => {
                lval.Token(lex.prepareToken(token));
                tok = T_END_HEREDOC;
                fnext php;
                fbreak;
            };
        *|;
        
        string_var := |*
            '$' varname        => {lval.Token(lex.prepareToken(token)); tok = T_VARIABLE; fbreak;};
            '->' varname_first => {lex.ungetCnt(1); lval.Token(lex.prepareToken(token)); tok = T_OBJECT_OPERATOR; fbreak;};
            varname            => {lval.Token(lex.prepareToken(token)); tok = T_STRING; fbreak;};
            '['                => {lval.Token(lex.prepareToken(token)); tok = TokenID(int('[')); lex.call(ftargs, fentry(string_var_index)); goto _out;};
            any                => {lex.ungetCnt(1); fret;};
        *|;
        
        string_var_index := |*
            lnum | hnum | bnum       => {lval.Token(lex.prepareToken(token)); tok = T_NUM_STRING; fbreak;};
            '$' varname              => {lval.Token(lex.prepareToken(token)); tok = T_VARIABLE; fbreak;};
            varname                  => {lval.Token(lex.prepareToken(token)); tok = T_STRING; fbreak;};
            whitespace_line | [\\'#] => {lval.Token(lex.prepareToken(token)); tok = T_ENCAPSED_AND_WHITESPACE; lex.ret(2); goto _out;};
            operators > (svi, 1)     => {lval.Token(lex.prepareToken(token)); tok = TokenID(int(lex.data[lex.ts])); fbreak;};
            ']'       > (svi, 2)     => {lval.Token(lex.prepareToken(token)); tok = TokenID(int(']')); lex.ret(2); goto _out;};
            any_line => {
                c := lex.data[lex.p]
                lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c));
            };
        *|;

        string_var_name := |*
            varname ("[" | "}") => {lex.ungetCnt(1); lval.Token(lex.prepareToken(token)); tok = T_STRING_VARNAME; fnext php; fbreak;};
            any                 => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_open_parenthesis := |*
            whitespace_line* => {lex.addSkippedToken(token, SkippedTokenTypeWhitespace, lex.ts, lex.te)};
            "("              => {lval.Token(lex.prepareToken(token)); tok = TokenID(int('(')); fnext halt_compiller_close_parenthesis; fbreak;};
            any              => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_close_parenthesis := |*
            whitespace_line* => {lex.addSkippedToken(token, SkippedTokenTypeWhitespace, lex.ts, lex.te)};
            ")"              => {lval.Token(lex.prepareToken(token)); tok = TokenID(int(')')); fnext halt_compiller_close_semicolon; fbreak;};
            any              => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_close_semicolon := |*
            whitespace_line* => {lex.addSkippedToken(token, SkippedTokenTypeWhitespace, lex.ts, lex.te)};
            ";"              => {lval.Token(lex.prepareToken(token)); tok = TokenID(int(';')); fnext halt_compiller_end; fbreak;};
            any              => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_end := |*
            any_line* => { lex.addSkippedToken(token, SkippedTokenTypeToken, lex.ts, lex.te); };
        *|;

        write exec;
    }%%

    return int(tok);
}