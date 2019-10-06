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
    token.HiddenTokens = token.HiddenTokens[:0]
    eof := lex.pe
    var tok TokenType
    lex.prepareToken(token, tok)

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
                lex.prepareToken(token, tok)
                tok = T_INLINE_HTML;
                fbreak;
            };
            '<?' => {
                lex.addHiddenToken(token, T_OPEN_TAG, lex.ts, lex.te)
                fnext php;
            };
            '<?php'i ( [ \t] | newline ) => {
                lex.ungetCnt(lex.te - lex.ts - 5)
                lex.addHiddenToken(token, T_OPEN_TAG, lex.ts, lex.ts+5)
                fnext php;
            };
            '<?='i => {
                tok = T_ECHO;
                lex.prepareToken(token, tok);
                fnext php;
                fbreak;
            };
        *|;

        php := |*
            whitespace_line*                   => {lex.addHiddenToken(token, T_WHITESPACE, lex.ts, lex.te)};
            '?>' newline?                      => {tok = TokenType(int(';')); lex.prepareToken(token, tok); fnext main; fbreak;};
            ';' whitespace_line* '?>' newline? => {tok = TokenType(int(';')); lex.prepareToken(token, tok); fnext main; fbreak;};

            (dnum | exponent_dnum)          => {tok = T_DNUMBER; lex.prepareToken(token, tok); fbreak;};
            bnum => {
                firstNum := 2
                for i := lex.ts + 2; i < lex.te; i++ {
                    if lex.data[i] == '0' {
                        firstNum++
                    }
                }

                if lex.te - lex.ts - firstNum < 64 {
                    tok = T_LNUMBER; lex.prepareToken(token, tok); fbreak;
                }
                tok = T_DNUMBER; lex.prepareToken(token, tok); fbreak;
            };
            lnum => {
                if lex.te - lex.ts < 20 {
                    tok = T_LNUMBER; lex.prepareToken(token, tok); fbreak;
                }
                tok = T_DNUMBER; lex.prepareToken(token, tok); fbreak;
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
                    tok = T_LNUMBER; lex.prepareToken(token, tok); fbreak;
                } 
                tok = T_DNUMBER; lex.prepareToken(token, tok); fbreak;
            };

            'abstract'i                       => {tok = T_ABSTRACT; lex.prepareToken(token, tok); fbreak;};
            'array'i                          => {tok = T_ARRAY; lex.prepareToken(token, tok); fbreak;};
            'as'i                             => {tok = T_AS; lex.prepareToken(token, tok); fbreak;};
            'break'i                          => {tok = T_BREAK; lex.prepareToken(token, tok); fbreak;};
            'callable'i                       => {tok = T_CALLABLE; lex.prepareToken(token, tok); fbreak;};
            'case'i                           => {tok = T_CASE; lex.prepareToken(token, tok); fbreak;};
            'catch'i                          => {tok = T_CATCH; lex.prepareToken(token, tok); fbreak;};
            'class'i                          => {tok = T_CLASS; lex.prepareToken(token, tok); fbreak;};
            'clone'i                          => {tok = T_CLONE; lex.prepareToken(token, tok); fbreak;};
            'const'i                          => {tok = T_CONST; lex.prepareToken(token, tok); fbreak;};
            'continue'i                       => {tok = T_CONTINUE; lex.prepareToken(token, tok); fbreak;};
            'declare'i                        => {tok = T_DECLARE; lex.prepareToken(token, tok); fbreak;};
            'default'i                        => {tok = T_DEFAULT; lex.prepareToken(token, tok); fbreak;};
            'do'i                             => {tok = T_DO; lex.prepareToken(token, tok); fbreak;};
            'echo'i                           => {tok = T_ECHO; lex.prepareToken(token, tok); fbreak;};
            'else'i                           => {tok = T_ELSE; lex.prepareToken(token, tok); fbreak;};
            'elseif'i                         => {tok = T_ELSEIF; lex.prepareToken(token, tok); fbreak;};
            'empty'i                          => {tok = T_EMPTY; lex.prepareToken(token, tok); fbreak;};
            'enddeclare'i                     => {tok = T_ENDDECLARE; lex.prepareToken(token, tok); fbreak;};
            'endfor'i                         => {tok = T_ENDFOR; lex.prepareToken(token, tok); fbreak;};
            'endforeach'i                     => {tok = T_ENDFOREACH; lex.prepareToken(token, tok); fbreak;};
            'endif'i                          => {tok = T_ENDIF; lex.prepareToken(token, tok); fbreak;};
            'endswitch'i                      => {tok = T_ENDSWITCH; lex.prepareToken(token, tok); fbreak;};
            'endwhile'i                       => {tok = T_ENDWHILE; lex.prepareToken(token, tok); fbreak;};
            'eval'i                           => {tok = T_EVAL; lex.prepareToken(token, tok); fbreak;};
            'exit'i | 'die'i                  => {tok = T_EXIT; lex.prepareToken(token, tok); fbreak;};
            'extends'i                        => {tok = T_EXTENDS; lex.prepareToken(token, tok); fbreak;};
            'final'i                          => {tok = T_FINAL; lex.prepareToken(token, tok); fbreak;};
            'finally'i                        => {tok = T_FINALLY; lex.prepareToken(token, tok); fbreak;};
            'for'i                            => {tok = T_FOR; lex.prepareToken(token, tok); fbreak;};
            'foreach'i                        => {tok = T_FOREACH; lex.prepareToken(token, tok); fbreak;};
            'function'i | 'cfunction'i        => {tok = T_FUNCTION; lex.prepareToken(token, tok); fbreak;};
            'global'i                         => {tok = T_GLOBAL; lex.prepareToken(token, tok); fbreak;};
            'goto'i                           => {tok = T_GOTO; lex.prepareToken(token, tok); fbreak;};
            'if'i                             => {tok = T_IF; lex.prepareToken(token, tok); fbreak;};
            'isset'i                          => {tok = T_ISSET; lex.prepareToken(token, tok); fbreak;};
            'implements'i                     => {tok = T_IMPLEMENTS; lex.prepareToken(token, tok); fbreak;};
            'instanceof'i                     => {tok = T_INSTANCEOF; lex.prepareToken(token, tok); fbreak;};
            'insteadof'i                      => {tok = T_INSTEADOF; lex.prepareToken(token, tok); fbreak;};
            'interface'i                      => {tok = T_INTERFACE; lex.prepareToken(token, tok); fbreak;};
            'list'i                           => {tok = T_LIST; lex.prepareToken(token, tok); fbreak;};
            'namespace'i                      => {tok = T_NAMESPACE; lex.prepareToken(token, tok); fbreak;};
            'private'i                        => {tok = T_PRIVATE; lex.prepareToken(token, tok); fbreak;};
            'public'i                         => {tok = T_PUBLIC; lex.prepareToken(token, tok); fbreak;};
            'print'i                          => {tok = T_PRINT; lex.prepareToken(token, tok); fbreak;};
            'protected'i                      => {tok = T_PROTECTED; lex.prepareToken(token, tok); fbreak;};
            'return'i                         => {tok = T_RETURN; lex.prepareToken(token, tok); fbreak;};
            'static'i                         => {tok = T_STATIC; lex.prepareToken(token, tok); fbreak;};
            'switch'i                         => {tok = T_SWITCH; lex.prepareToken(token, tok); fbreak;};
            'throw'i                          => {tok = T_THROW; lex.prepareToken(token, tok); fbreak;};
            'trait'i                          => {tok = T_TRAIT; lex.prepareToken(token, tok); fbreak;};
            'try'i                            => {tok = T_TRY; lex.prepareToken(token, tok); fbreak;};
            'unset'i                          => {tok = T_UNSET; lex.prepareToken(token, tok); fbreak;};
            'use'i                            => {tok = T_USE; lex.prepareToken(token, tok); fbreak;};
            'var'i                            => {tok = T_VAR; lex.prepareToken(token, tok); fbreak;};
            'while'i                          => {tok = T_WHILE; lex.prepareToken(token, tok); fbreak;};
            'yield'i whitespace_line+ 'from'i => {tok = T_YIELD_FROM; lex.prepareToken(token, tok); fbreak;};
            'yield'i                          => {tok = T_YIELD; lex.prepareToken(token, tok); fbreak;};
            'include'i                        => {tok = T_INCLUDE; lex.prepareToken(token, tok); fbreak;};
            'include_once'i                   => {tok = T_INCLUDE_ONCE; lex.prepareToken(token, tok); fbreak;};
            'require'i                        => {tok = T_REQUIRE; lex.prepareToken(token, tok); fbreak;};
            'require_once'i                   => {tok = T_REQUIRE_ONCE; lex.prepareToken(token, tok); fbreak;};
            '__CLASS__'i                      => {tok = T_CLASS_C; lex.prepareToken(token, tok); fbreak;};
            '__DIR__'i                        => {tok = T_DIR; lex.prepareToken(token, tok); fbreak;};
            '__FILE__'i                       => {tok = T_FILE; lex.prepareToken(token, tok); fbreak;};
            '__FUNCTION__'i                   => {tok = T_FUNC_C; lex.prepareToken(token, tok); fbreak;};
            '__LINE__'i                       => {tok = T_LINE; lex.prepareToken(token, tok); fbreak;};
            '__NAMESPACE__'i                  => {tok = T_NS_C; lex.prepareToken(token, tok); fbreak;};
            '__METHOD__'i                     => {tok = T_METHOD_C; lex.prepareToken(token, tok); fbreak;};
            '__TRAIT__'i                      => {tok = T_TRAIT_C; lex.prepareToken(token, tok); fbreak;};
            '__halt_compiler'i                => {tok = T_HALT_COMPILER; lex.prepareToken(token, tok); fnext halt_compiller_open_parenthesis; fbreak;};
            'new'i                            => {tok = T_NEW; lex.prepareToken(token, tok); fbreak;};
            'and'i                            => {tok = T_LOGICAL_AND; lex.prepareToken(token, tok); fbreak;};
            'or'i                             => {tok = T_LOGICAL_OR; lex.prepareToken(token, tok); fbreak;};
            'xor'i                            => {tok = T_LOGICAL_XOR; lex.prepareToken(token, tok); fbreak;};
            '\\'                              => {tok = T_NS_SEPARATOR; lex.prepareToken(token, tok); fbreak;};
            '...'                             => {tok = T_ELLIPSIS; lex.prepareToken(token, tok); fbreak;};
            '::'                              => {tok = T_PAAMAYIM_NEKUDOTAYIM; lex.prepareToken(token, tok); fbreak;};
            '&&'                              => {tok = T_BOOLEAN_AND; lex.prepareToken(token, tok); fbreak;};
            '||'                              => {tok = T_BOOLEAN_OR; lex.prepareToken(token, tok); fbreak;};
            '&='                              => {tok = T_AND_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '|='                              => {tok = T_OR_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '.='                              => {tok = T_CONCAT_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '*='                              => {tok = T_MUL_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '**='                             => {tok = T_POW_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '/='                              => {tok = T_DIV_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '+='                              => {tok = T_PLUS_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '-='                              => {tok = T_MINUS_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '^='                              => {tok = T_XOR_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '%='                              => {tok = T_MOD_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '--'                              => {tok = T_DEC; lex.prepareToken(token, tok); fbreak;};
            '++'                              => {tok = T_INC; lex.prepareToken(token, tok); fbreak;};
            '=>'                              => {tok = T_DOUBLE_ARROW; lex.prepareToken(token, tok); fbreak;};
            '<=>'                             => {tok = T_SPACESHIP; lex.prepareToken(token, tok); fbreak;};
            '!=' | '<>'                       => {tok = T_IS_NOT_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '!=='                             => {tok = T_IS_NOT_IDENTICAL; lex.prepareToken(token, tok); fbreak;};
            '=='                              => {tok = T_IS_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '==='                             => {tok = T_IS_IDENTICAL; lex.prepareToken(token, tok); fbreak;};
            '<<='                             => {tok = T_SL_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '>>='                             => {tok = T_SR_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '>='                              => {tok = T_IS_GREATER_OR_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '<='                              => {tok = T_IS_SMALLER_OR_EQUAL; lex.prepareToken(token, tok); fbreak;};
            '**'                              => {tok = T_POW; lex.prepareToken(token, tok); fbreak;};
            '<<'                              => {tok = T_SL; lex.prepareToken(token, tok); fbreak;};
            '>>'                              => {tok = T_SR; lex.prepareToken(token, tok); fbreak;};
            '??'                              => {tok = T_COALESCE; lex.prepareToken(token, tok); fbreak;};

            '(' whitespace* 'array'i whitespace* ')'                     => {tok = T_ARRAY_CAST; lex.prepareToken(token, tok); fbreak;};
            '(' whitespace* ('bool'i|'boolean'i) whitespace* ')'         => {tok = T_BOOL_CAST; lex.prepareToken(token, tok); fbreak;};
            '(' whitespace* ('real'i|'double'i|'float'i) whitespace* ')' => {tok = T_DOUBLE_CAST; lex.prepareToken(token, tok); fbreak;};
            '(' whitespace* ('int'i|'integer'i) whitespace* ')'          => {tok = T_INT_CAST; lex.prepareToken(token, tok); fbreak;};
            '(' whitespace* 'object'i whitespace* ')'                    => {tok = T_OBJECT_CAST; lex.prepareToken(token, tok); fbreak;};
            '(' whitespace* ('string'i|'binary'i) whitespace* ')'        => {tok = T_STRING_CAST; lex.prepareToken(token, tok); fbreak;};
            '(' whitespace* 'unset'i whitespace* ')'                     => {tok = T_UNSET_CAST; lex.prepareToken(token, tok); fbreak;};

            ('#' | '//') any_line* when is_not_comment_end => {
                lex.ungetStr("?>")
                lex.addHiddenToken(token, T_COMMENT, lex.ts, lex.te)
            };
            '/*' any_line* :>> '*/' {
                tokenType := T_COMMENT

                if lex.te - lex.ts > 4 && string(lex.data[lex.ts:lex.ts+3]) == "/**" {
                    tokenType = T_DOC_COMMENT
                }
                lex.addHiddenToken(token, tokenType, lex.ts, lex.te)
            };

            operators => {
                // rune, _ := utf8.DecodeRune(lex.data[lex.ts:lex.te]);
                // tok = TokenType(Rune2Class(rune));
                tok = TokenType(int(lex.data[lex.ts]));
                lex.prepareToken(token, tok);
                fbreak;
            };

            "{"          => { tok = TokenType(int('{')); lex.prepareToken(token, tok); lex.call(ftargs, fentry(php)); goto _out; };
            "}"          => { tok = TokenType(int('}')); lex.prepareToken(token, tok); lex.ret(1); goto _out;};
            "$" varname  => { tok = T_VARIABLE; lex.prepareToken(token, tok); fbreak; };
            varname      => { tok = T_STRING; lex.prepareToken(token, tok);   fbreak; };

            "->"         => { tok = T_OBJECT_OPERATOR; lex.prepareToken(token, tok); fnext property; fbreak; };

            constant_string => {
                tok = T_CONSTANT_ENCAPSED_STRING;
                lex.prepareToken(token, tok);
                fbreak;
            };

            "b"i? "<<<" [ \t]* ( heredoc_label | ("'" heredoc_label "'") | ('"' heredoc_label '"') ) newline  => {
                lex.heredocLabel = lex.data[lblStart:lblEnd]
                tok = T_START_HEREDOC;
                lex.prepareToken(token, tok);

                if lex.isHeredocEnd(lex.p+1) {
                    fnext heredoc_end;
                } else if lex.data[lblStart-1] == '\'' {
                    fnext nowdoc;
                } else {
                    fnext heredoc;
                }
                fbreak;
            };
            "`" => {tok = TokenType(int('`')); lex.prepareToken(token, tok); fnext backqote; fbreak;};
            '"' => {tok = TokenType(int('"')); lex.prepareToken(token, tok); fnext template_string; fbreak;};

            any_line => {
                c := lex.data[lex.p]
                lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c));
            };
        *|;

        property := |*
            whitespace_line* => {lex.addHiddenToken(token, T_WHITESPACE, lex.ts, lex.te)};
            "->"             => {tok = T_OBJECT_OPERATOR; lex.prepareToken(token, tok); fbreak;};
            varname          => {tok = T_STRING; lex.prepareToken(token, tok); fnext php; fbreak;};
            any              => {lex.ungetCnt(1); fgoto php;};
        *|;

        nowdoc := |*
            any_line* when is_not_heredoc_end => {
                tok = T_ENCAPSED_AND_WHITESPACE;
                lex.prepareToken(token, tok);
                fnext heredoc_end;
                fbreak;
            };
        *|;
        
        heredoc := |*
            "{$" => {lex.ungetCnt(1); tok = T_CURLY_OPEN; lex.prepareToken(token, tok); lex.call(ftargs, fentry(php)); goto _out;};
            "${" => {tok = T_DOLLAR_OPEN_CURLY_BRACES; lex.prepareToken(token, tok); lex.call(ftargs, fentry(string_var_name)); goto _out;};
            "$"  => {lex.ungetCnt(1); fcall string_var;};
            any_line* when is_not_heredoc_end_or_var => {
                tok = T_ENCAPSED_AND_WHITESPACE;
                lex.prepareToken(token, tok);

                if lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
                    fnext heredoc_end;
                }
                fbreak;
            };
        *|;
        
        backqote := |*
            "{$" => {lex.ungetCnt(1); tok = T_CURLY_OPEN; lex.prepareToken(token, tok); lex.call(ftargs, fentry(php)); goto _out;};
            "${" => {tok = T_DOLLAR_OPEN_CURLY_BRACES; lex.prepareToken(token, tok); lex.call(ftargs, fentry(string_var_name)); goto _out;};
            "$"  => {lex.ungetCnt(1); fcall string_var;};
            '`'  => {tok = TokenType(int('`')); lex.prepareToken(token, tok); fnext php; fbreak;};
            any_line* when is_not_backqoute_end_or_var => {
                tok = T_ENCAPSED_AND_WHITESPACE;
                lex.prepareToken(token, tok);
                fbreak;
            };
        *|;
        
        template_string := |*
            "{$" => {lex.ungetCnt(1); tok = T_CURLY_OPEN; lex.prepareToken(token, tok); lex.call(ftargs, fentry(php)); goto _out;};
            "${" => {tok = T_DOLLAR_OPEN_CURLY_BRACES; lex.prepareToken(token, tok); lex.call(ftargs, fentry(string_var_name)); goto _out;};
            "$"  => {lex.ungetCnt(1); fcall string_var;};
            '"'  => {tok = TokenType(int('"')); lex.prepareToken(token, tok); fnext php; fbreak;};
            any_line* when is_not_string_end_or_var => {
                tok = T_ENCAPSED_AND_WHITESPACE;
                lex.prepareToken(token, tok);
                fbreak;
            };
        *|;

        heredoc_end := |*
            varname -- ";" => {
                tok = T_END_HEREDOC;
                lex.prepareToken(token, tok);
                fnext php;
                fbreak;
            };
            varname => {
                tok = T_END_HEREDOC;
                lex.prepareToken(token, tok);
                fnext php;
                fbreak;
            };
        *|;
        
        string_var := |*
            '$' varname        => {tok = T_VARIABLE; lex.prepareToken(token, tok); fbreak;};
            '->' varname_first => {lex.ungetCnt(1); tok = T_OBJECT_OPERATOR; lex.prepareToken(token, tok); fbreak;};
            varname            => {tok = T_STRING; lex.prepareToken(token, tok); fbreak;};
            '['                => {tok = TokenType(int('[')); lex.prepareToken(token, tok); lex.call(ftargs, fentry(string_var_index)); goto _out;};
            any                => {lex.ungetCnt(1); fret;};
        *|;
        
        string_var_index := |*
            lnum | hnum | bnum       => {tok = T_NUM_STRING; lex.prepareToken(token, tok); fbreak;};
            '$' varname              => {tok = T_VARIABLE; lex.prepareToken(token, tok); fbreak;};
            varname                  => {tok = T_STRING; lex.prepareToken(token, tok); fbreak;};
            whitespace_line | [\\'#] => {tok = T_ENCAPSED_AND_WHITESPACE; lex.prepareToken(token, tok); lex.ret(2); goto _out;};
            operators > (svi, 1)     => {lex.prepareToken(token, tok); tok = TokenType(int(lex.data[lex.ts])); fbreak;};
            ']'       > (svi, 2)     => {tok = TokenType(int(']')); lex.prepareToken(token, tok); lex.ret(2); goto _out;};
            any_line => {
                c := lex.data[lex.p]
                lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c));
            };
        *|;

        string_var_name := |*
            varname ("[" | "}") => {lex.ungetCnt(1); tok = T_STRING_VARNAME; lex.prepareToken(token, tok); fnext php; fbreak;};
            any                 => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_open_parenthesis := |*
            whitespace_line* => {lex.addHiddenToken(token, T_WHITESPACE, lex.ts, lex.te)};
            "("              => {tok = TokenType(int('(')); lex.prepareToken(token, tok); fnext halt_compiller_close_parenthesis; fbreak;};
            any              => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_close_parenthesis := |*
            whitespace_line* => {lex.addHiddenToken(token, T_WHITESPACE, lex.ts, lex.te)};
            ")"              => {tok = TokenType(int(')')); lex.prepareToken(token, tok); fnext halt_compiller_close_semicolon; fbreak;};
            any              => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_close_semicolon := |*
            whitespace_line* => {lex.addHiddenToken(token, T_WHITESPACE, lex.ts, lex.te)};
            ";"              => {tok = TokenType(int(';')); lex.prepareToken(token, tok); fnext halt_compiller_end; fbreak;};
            any              => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_end := |*
            any_line* => { lex.addHiddenToken(token, T_COMMENT, lex.ts, lex.te); };
        *|;

        write exec;
    }%%

    lval.Token(token)

    return int(tok);
}