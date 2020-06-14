package scanner

import (
    "fmt"
    "strconv"
    "strings"

    "github.com/z7zmey/php-parser/freefloating"
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
    lex.FreeFloating = nil
    eof := lex.pe
    var tok TokenID

    token := lex.TokenPool.Get()
    token.FreeFloating = lex.FreeFloating
    token.Value = string(lex.data[0:0])

    lblStart := 0
    lblEnd   := 0

    _, _ = lblStart, lblEnd

    %%{ 
        action heredoc_lbl_start {lblStart = lex.p}
        action heredoc_lbl_end   {lblEnd = lex.p}

        action constant_string_new_line   {
            if lex.data[lex.p] == '\n' {
                lex.NewLines.Append(lex.p)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.NewLines.Append(lex.p)
            }
        }

        action is_not_heredoc_end { lex.isNotHeredocEnd(lex.p) }
        action is_not_comment_end { lex.isNotPhpCloseToken() && lex.isNotNewLine()  }
        action is_not_heredoc_end_or_var { lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() }
        action is_not_string_end_or_var { lex.isNotStringEnd('"') && lex.isNotStringVar() }
        action is_not_backqoute_end_or_var { lex.isNotStringEnd('`') && lex.isNotStringVar() }

        newline = ('\r\n' >(nl, 1) | '\r' >(nl, 0) | '\n' >(nl, 0)) %{lex.NewLines.Append(lex.p);};
        any_line = any | newline;
        whitespace = [\t\v\f ];
        whitespace_line = [\t\v\f ] | newline;

        lnum = [0-9]+('_'[0-9]+)*;
        dnum = (lnum?"." lnum)|(lnum"."lnum?);
        hnum = '0x'[0-9a-fA-F]+('_'[0-9a-fA-F]+)*;
        bnum = '0b'[01]+('_'[01]+)*;

        exponent_dnum = (lnum | dnum) ('e'|'E') ('+'|'-')? lnum;
        varname_first = [a-zA-Z_] | (0x0080..0x00FF);
        varname_second = varname_first | [0-9];
        varname       = varname_first (varname_second)*;
        heredoc_label = varname >heredoc_lbl_start %heredoc_lbl_end;
        operators     = ';'|':'|','|'.'|'['|']'|'('|')'|'|'|'/'|'^'|'&'|'+'|'-'|'*'|'='|'%'|'!'|'~'|'$'|'<'|'>'|'?'|'@';
        
        prepush { lex.growCallStack(); }

        constant_string =
            start: (
                "'"         -> qoute
                | "b"i? '"' -> double_qoute
            ),

            # single qoute string

            qoute: (
                (any - [\\'\r\n])                -> qoute
                | "\r" @constant_string_new_line -> qoute
                | "\n" @constant_string_new_line -> qoute
                | "\\"                           -> qoute_any
                | "'"                            -> final
            ),
            qoute_any: (
                (any - [\r\n])                   -> qoute
                | "\r" @constant_string_new_line -> qoute
                | "\n" @constant_string_new_line -> qoute
            ),

            # double qoute string

            double_qoute: (
                (any - [\\"${\r\n])                -> double_qoute
                | "\r" @constant_string_new_line   -> double_qoute
                | "\n" @constant_string_new_line   -> double_qoute
                | "\\"                             -> double_qoute_any
                | '"'                              -> final
                | '$'                              -> double_qoute_nonvarname
                | '{'                              -> double_qoute_nondollar
            ),
            double_qoute_any: (
                (any - [\r\n])                     -> double_qoute
                | "\r" @constant_string_new_line   -> double_qoute
                | "\n" @constant_string_new_line   -> double_qoute
            ),
            double_qoute_nondollar: (
                (any - [\\$"\r\n])                 -> double_qoute
                | "\r" @constant_string_new_line   -> double_qoute
                | "\n" @constant_string_new_line   -> double_qoute
                | "\\"                             -> double_qoute_any
                | '"'                              -> final
            ),
            double_qoute_nonvarname: (
                (any - [\\{"\r\n] - varname_first) -> double_qoute
                | "\r" @constant_string_new_line   -> double_qoute
                | "\n" @constant_string_new_line   -> double_qoute
                | "\\"                             -> double_qoute_any
                | '"'                              -> final
            );

        main := |*
            "#!" any* :>> newline => {
                lex.addFreeFloating(freefloating.CommentType, lex.ts, lex.te)
            };
            any => {
                fnext html;
                lex.ungetCnt(1)
            };
        *|;

        html := |*
            any_line+ -- '<?' => {
                lex.ungetStr("<")
                lex.setTokenPosition(token)
                tok = T_INLINE_HTML;
                fbreak;
            };
            '<?' => {
                lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.te)
                fnext php;
            };
            '<?php'i ( [ \t] | newline ) => {
                lex.ungetCnt(lex.te - lex.ts - 5)
                lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.ts+5)
                fnext php;
            };
            '<?='i => {
                lex.setTokenPosition(token);
                tok = T_ECHO;
                fnext php;
                fbreak;
            };
        *|;

        php := |*
            whitespace_line*                   => {lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)};
            '?>' newline?                      => {lex.setTokenPosition(token); tok = TokenID(int(';')); fnext html; fbreak;};
            ';' whitespace_line* '?>' newline? => {lex.setTokenPosition(token); tok = TokenID(int(';')); fnext html; fbreak;};

            (dnum | exponent_dnum)          => {lex.setTokenPosition(token); tok = T_DNUMBER; fbreak;};
            bnum => {
                s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
                _, err := strconv.ParseInt(s, 2, 0)

                if err == nil {
                    lex.setTokenPosition(token); tok = T_LNUMBER; fbreak;
                } 
                
                lex.setTokenPosition(token); tok = T_DNUMBER; fbreak;
            };
            lnum => {
                base := 10
                if lex.data[lex.ts] == '0' {
                    base = 8
                }

                s := strings.Replace(string(lex.data[lex.ts:lex.te]), "_", "", -1)
                _, err := strconv.ParseInt(s, base, 0)

                if err == nil {
                    lex.setTokenPosition(token); tok = T_LNUMBER; fbreak;
                } 
                
                lex.setTokenPosition(token); tok = T_DNUMBER; fbreak;
            };
            hnum => {
                s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
                _, err := strconv.ParseInt(s, 16, 0)

                if err == nil {
                    lex.setTokenPosition(token); tok = T_LNUMBER; fbreak;
                } 
                
                lex.setTokenPosition(token); tok = T_DNUMBER; fbreak;
            };

            'abstract'i                       => {lex.setTokenPosition(token); tok = T_ABSTRACT; fbreak;};
            'array'i                          => {lex.setTokenPosition(token); tok = T_ARRAY; fbreak;};
            'as'i                             => {lex.setTokenPosition(token); tok = T_AS; fbreak;};
            'break'i                          => {lex.setTokenPosition(token); tok = T_BREAK; fbreak;};
            'callable'i                       => {lex.setTokenPosition(token); tok = T_CALLABLE; fbreak;};
            'case'i                           => {lex.setTokenPosition(token); tok = T_CASE; fbreak;};
            'catch'i                          => {lex.setTokenPosition(token); tok = T_CATCH; fbreak;};
            'class'i                          => {lex.setTokenPosition(token); tok = T_CLASS; fbreak;};
            'clone'i                          => {lex.setTokenPosition(token); tok = T_CLONE; fbreak;};
            'const'i                          => {lex.setTokenPosition(token); tok = T_CONST; fbreak;};
            'continue'i                       => {lex.setTokenPosition(token); tok = T_CONTINUE; fbreak;};
            'declare'i                        => {lex.setTokenPosition(token); tok = T_DECLARE; fbreak;};
            'default'i                        => {lex.setTokenPosition(token); tok = T_DEFAULT; fbreak;};
            'do'i                             => {lex.setTokenPosition(token); tok = T_DO; fbreak;};
            'echo'i                           => {lex.setTokenPosition(token); tok = T_ECHO; fbreak;};
            'else'i                           => {lex.setTokenPosition(token); tok = T_ELSE; fbreak;};
            'elseif'i                         => {lex.setTokenPosition(token); tok = T_ELSEIF; fbreak;};
            'empty'i                          => {lex.setTokenPosition(token); tok = T_EMPTY; fbreak;};
            'enddeclare'i                     => {lex.setTokenPosition(token); tok = T_ENDDECLARE; fbreak;};
            'endfor'i                         => {lex.setTokenPosition(token); tok = T_ENDFOR; fbreak;};
            'endforeach'i                     => {lex.setTokenPosition(token); tok = T_ENDFOREACH; fbreak;};
            'endif'i                          => {lex.setTokenPosition(token); tok = T_ENDIF; fbreak;};
            'endswitch'i                      => {lex.setTokenPosition(token); tok = T_ENDSWITCH; fbreak;};
            'endwhile'i                       => {lex.setTokenPosition(token); tok = T_ENDWHILE; fbreak;};
            'eval'i                           => {lex.setTokenPosition(token); tok = T_EVAL; fbreak;};
            'exit'i | 'die'i                  => {lex.setTokenPosition(token); tok = T_EXIT; fbreak;};
            'extends'i                        => {lex.setTokenPosition(token); tok = T_EXTENDS; fbreak;};
            'final'i                          => {lex.setTokenPosition(token); tok = T_FINAL; fbreak;};
            'finally'i                        => {lex.setTokenPosition(token); tok = T_FINALLY; fbreak;};
            'for'i                            => {lex.setTokenPosition(token); tok = T_FOR; fbreak;};
            'foreach'i                        => {lex.setTokenPosition(token); tok = T_FOREACH; fbreak;};
            'function'i | 'cfunction'i        => {lex.setTokenPosition(token); tok = T_FUNCTION; fbreak;};
            'fn'i                             => {lex.setTokenPosition(token); tok = T_FN; fbreak;};
            'global'i                         => {lex.setTokenPosition(token); tok = T_GLOBAL; fbreak;};
            'goto'i                           => {lex.setTokenPosition(token); tok = T_GOTO; fbreak;};
            'if'i                             => {lex.setTokenPosition(token); tok = T_IF; fbreak;};
            'isset'i                          => {lex.setTokenPosition(token); tok = T_ISSET; fbreak;};
            'implements'i                     => {lex.setTokenPosition(token); tok = T_IMPLEMENTS; fbreak;};
            'instanceof'i                     => {lex.setTokenPosition(token); tok = T_INSTANCEOF; fbreak;};
            'insteadof'i                      => {lex.setTokenPosition(token); tok = T_INSTEADOF; fbreak;};
            'interface'i                      => {lex.setTokenPosition(token); tok = T_INTERFACE; fbreak;};
            'list'i                           => {lex.setTokenPosition(token); tok = T_LIST; fbreak;};
            'namespace'i                      => {lex.setTokenPosition(token); tok = T_NAMESPACE; fbreak;};
            'private'i                        => {lex.setTokenPosition(token); tok = T_PRIVATE; fbreak;};
            'public'i                         => {lex.setTokenPosition(token); tok = T_PUBLIC; fbreak;};
            'print'i                          => {lex.setTokenPosition(token); tok = T_PRINT; fbreak;};
            'protected'i                      => {lex.setTokenPosition(token); tok = T_PROTECTED; fbreak;};
            'return'i                         => {lex.setTokenPosition(token); tok = T_RETURN; fbreak;};
            'static'i                         => {lex.setTokenPosition(token); tok = T_STATIC; fbreak;};
            'switch'i                         => {lex.setTokenPosition(token); tok = T_SWITCH; fbreak;};
            'throw'i                          => {lex.setTokenPosition(token); tok = T_THROW; fbreak;};
            'trait'i                          => {lex.setTokenPosition(token); tok = T_TRAIT; fbreak;};
            'try'i                            => {lex.setTokenPosition(token); tok = T_TRY; fbreak;};
            'unset'i                          => {lex.setTokenPosition(token); tok = T_UNSET; fbreak;};
            'use'i                            => {lex.setTokenPosition(token); tok = T_USE; fbreak;};
            'var'i                            => {lex.setTokenPosition(token); tok = T_VAR; fbreak;};
            'while'i                          => {lex.setTokenPosition(token); tok = T_WHILE; fbreak;};
            'yield'i whitespace_line* 'from'i => {lex.setTokenPosition(token); tok = T_YIELD_FROM; fbreak;};
            'yield'i                          => {lex.setTokenPosition(token); tok = T_YIELD; fbreak;};
            'include'i                        => {lex.setTokenPosition(token); tok = T_INCLUDE; fbreak;};
            'include_once'i                   => {lex.setTokenPosition(token); tok = T_INCLUDE_ONCE; fbreak;};
            'require'i                        => {lex.setTokenPosition(token); tok = T_REQUIRE; fbreak;};
            'require_once'i                   => {lex.setTokenPosition(token); tok = T_REQUIRE_ONCE; fbreak;};
            '__CLASS__'i                      => {lex.setTokenPosition(token); tok = T_CLASS_C; fbreak;};
            '__DIR__'i                        => {lex.setTokenPosition(token); tok = T_DIR; fbreak;};
            '__FILE__'i                       => {lex.setTokenPosition(token); tok = T_FILE; fbreak;};
            '__FUNCTION__'i                   => {lex.setTokenPosition(token); tok = T_FUNC_C; fbreak;};
            '__LINE__'i                       => {lex.setTokenPosition(token); tok = T_LINE; fbreak;};
            '__NAMESPACE__'i                  => {lex.setTokenPosition(token); tok = T_NS_C; fbreak;};
            '__METHOD__'i                     => {lex.setTokenPosition(token); tok = T_METHOD_C; fbreak;};
            '__TRAIT__'i                      => {lex.setTokenPosition(token); tok = T_TRAIT_C; fbreak;};
            '__halt_compiler'i                => {lex.setTokenPosition(token); tok = T_HALT_COMPILER; fnext halt_compiller_open_parenthesis; fbreak;};
            'new'i                            => {lex.setTokenPosition(token); tok = T_NEW; fbreak;};
            'and'i                            => {lex.setTokenPosition(token); tok = T_LOGICAL_AND; fbreak;};
            'or'i                             => {lex.setTokenPosition(token); tok = T_LOGICAL_OR; fbreak;};
            'xor'i                            => {lex.setTokenPosition(token); tok = T_LOGICAL_XOR; fbreak;};
            '\\'                              => {lex.setTokenPosition(token); tok = T_NS_SEPARATOR; fbreak;};
            '...'                             => {lex.setTokenPosition(token); tok = T_ELLIPSIS; fbreak;};
            '::'                              => {lex.setTokenPosition(token); tok = T_PAAMAYIM_NEKUDOTAYIM; fbreak;};
            '&&'                              => {lex.setTokenPosition(token); tok = T_BOOLEAN_AND; fbreak;};
            '||'                              => {lex.setTokenPosition(token); tok = T_BOOLEAN_OR; fbreak;};
            '&='                              => {lex.setTokenPosition(token); tok = T_AND_EQUAL; fbreak;};
            '|='                              => {lex.setTokenPosition(token); tok = T_OR_EQUAL; fbreak;};
            '.='                              => {lex.setTokenPosition(token); tok = T_CONCAT_EQUAL; fbreak;};
            '*='                              => {lex.setTokenPosition(token); tok = T_MUL_EQUAL; fbreak;};
            '**='                             => {lex.setTokenPosition(token); tok = T_POW_EQUAL; fbreak;};
            '/='                              => {lex.setTokenPosition(token); tok = T_DIV_EQUAL; fbreak;};
            '+='                              => {lex.setTokenPosition(token); tok = T_PLUS_EQUAL; fbreak;};
            '-='                              => {lex.setTokenPosition(token); tok = T_MINUS_EQUAL; fbreak;};
            '^='                              => {lex.setTokenPosition(token); tok = T_XOR_EQUAL; fbreak;};
            '%='                              => {lex.setTokenPosition(token); tok = T_MOD_EQUAL; fbreak;};
            '--'                              => {lex.setTokenPosition(token); tok = T_DEC; fbreak;};
            '++'                              => {lex.setTokenPosition(token); tok = T_INC; fbreak;};
            '=>'                              => {lex.setTokenPosition(token); tok = T_DOUBLE_ARROW; fbreak;};
            '<=>'                             => {lex.setTokenPosition(token); tok = T_SPACESHIP; fbreak;};
            '!=' | '<>'                       => {lex.setTokenPosition(token); tok = T_IS_NOT_EQUAL; fbreak;};
            '!=='                             => {lex.setTokenPosition(token); tok = T_IS_NOT_IDENTICAL; fbreak;};
            '=='                              => {lex.setTokenPosition(token); tok = T_IS_EQUAL; fbreak;};
            '==='                             => {lex.setTokenPosition(token); tok = T_IS_IDENTICAL; fbreak;};
            '<<='                             => {lex.setTokenPosition(token); tok = T_SL_EQUAL; fbreak;};
            '>>='                             => {lex.setTokenPosition(token); tok = T_SR_EQUAL; fbreak;};
            '>='                              => {lex.setTokenPosition(token); tok = T_IS_GREATER_OR_EQUAL; fbreak;};
            '<='                              => {lex.setTokenPosition(token); tok = T_IS_SMALLER_OR_EQUAL; fbreak;};
            '**'                              => {lex.setTokenPosition(token); tok = T_POW; fbreak;};
            '<<'                              => {lex.setTokenPosition(token); tok = T_SL; fbreak;};
            '>>'                              => {lex.setTokenPosition(token); tok = T_SR; fbreak;};
            '??'                              => {lex.setTokenPosition(token); tok = T_COALESCE; fbreak;};
            '??='                             => {lex.setTokenPosition(token); tok = T_COALESCE_EQUAL; fbreak;};

            '(' whitespace* 'array'i whitespace* ')'                     => {lex.setTokenPosition(token); tok = T_ARRAY_CAST; fbreak;};
            '(' whitespace* ('bool'i|'boolean'i) whitespace* ')'         => {lex.setTokenPosition(token); tok = T_BOOL_CAST; fbreak;};
            '(' whitespace* ('real'i|'double'i|'float'i) whitespace* ')' => {lex.setTokenPosition(token); tok = T_DOUBLE_CAST; fbreak;};
            '(' whitespace* ('int'i|'integer'i) whitespace* ')'          => {lex.setTokenPosition(token); tok = T_INT_CAST; fbreak;};
            '(' whitespace* 'object'i whitespace* ')'                    => {lex.setTokenPosition(token); tok = T_OBJECT_CAST; fbreak;};
            '(' whitespace* ('string'i|'binary'i) whitespace* ')'        => {lex.setTokenPosition(token); tok = T_STRING_CAST; fbreak;};
            '(' whitespace* 'unset'i whitespace* ')'                     => {lex.setTokenPosition(token); tok = T_UNSET_CAST; fbreak;};

            ('#' | '//') any_line* when is_not_comment_end => {
                lex.ungetStr("?>")
                lex.addFreeFloating(freefloating.CommentType, lex.ts, lex.te)
            };
            '/*' any_line* :>> '*/' {
                isDocComment := false;
                if lex.te - lex.ts > 4 && string(lex.data[lex.ts:lex.ts+3]) == "/**" {
                    isDocComment = true;
                }
                lex.addFreeFloating(freefloating.CommentType, lex.ts, lex.te)

                if isDocComment {
                    lex.PhpDocComment = string(lex.data[lex.ts:lex.te])
                }
            };

            operators => {
                // rune, _ := utf8.DecodeRune(lex.data[lex.ts:lex.te]);
                // tok = TokenID(Rune2Class(rune));
                lex.setTokenPosition(token);
                tok = TokenID(int(lex.data[lex.ts]));
                fbreak;
            };

            "{"          => { lex.setTokenPosition(token); tok = TokenID(int('{')); lex.call(ftargs, fentry(php)); goto _out; };
            "}"          => { lex.setTokenPosition(token); tok = TokenID(int('}')); lex.ret(1); lex.PhpDocComment = ""; goto _out;};
            "$" varname  => { lex.setTokenPosition(token); tok = T_VARIABLE; fbreak; };
            varname      => { lex.setTokenPosition(token); tok = T_STRING;   fbreak; };

            "->"         => { lex.setTokenPosition(token); tok = T_OBJECT_OPERATOR; fnext property; fbreak; };

            constant_string => {
                lex.setTokenPosition(token);
                tok = T_CONSTANT_ENCAPSED_STRING;
                fbreak;
            };

            "b"i? "<<<" [ \t]* ( heredoc_label | ("'" heredoc_label "'") | ('"' heredoc_label '"') ) newline  => {
                lex.heredocLabel = lex.data[lblStart:lblEnd]
                lex.setTokenPosition(token);
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
            "`" => {lex.setTokenPosition(token); tok = TokenID(int('`')); fnext backqote; fbreak;};
            '"' => {lex.setTokenPosition(token); tok = TokenID(int('"')); fnext template_string; fbreak;};

            any_line => {
                c := lex.data[lex.p]
                lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c));
            };
        *|;

        property := |*
            whitespace_line* => {lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)};
            "->"             => {lex.setTokenPosition(token); tok = T_OBJECT_OPERATOR; fbreak;};
            varname          => {lex.setTokenPosition(token); tok = T_STRING; fnext php; fbreak;};
            any              => {lex.ungetCnt(1); fgoto php;};
        *|;

        nowdoc := |*
            any_line* when is_not_heredoc_end => {
                lex.setTokenPosition(token);
                tok = T_ENCAPSED_AND_WHITESPACE;
                fnext heredoc_end;
                fbreak;
            };
        *|;
        
        heredoc := |*
            "{$" => {lex.ungetCnt(1); lex.setTokenPosition(token); tok = T_CURLY_OPEN; lex.call(ftargs, fentry(php)); goto _out;};
            "${" => {lex.setTokenPosition(token); tok = T_DOLLAR_OPEN_CURLY_BRACES; lex.call(ftargs, fentry(string_var_name)); goto _out;};
            "$"  => {lex.ungetCnt(1); fcall string_var;};
            any_line* when is_not_heredoc_end_or_var => {
                lex.setTokenPosition(token);
                tok = T_ENCAPSED_AND_WHITESPACE;

                if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
                    fnext heredoc_end;
                }
                fbreak;
            };
        *|;
        
        backqote := |*
            "{$"              => {lex.ungetCnt(1); lex.setTokenPosition(token); tok = T_CURLY_OPEN; lex.call(ftargs, fentry(php)); goto _out;};
            "${"              => {lex.setTokenPosition(token); tok = T_DOLLAR_OPEN_CURLY_BRACES; lex.call(ftargs, fentry(string_var_name)); goto _out;};
            "$" varname_first => {lex.ungetCnt(2); fcall string_var;};
            '`'               => {lex.setTokenPosition(token); tok = TokenID(int('`')); fnext php; fbreak;};
            any_line* when is_not_backqoute_end_or_var => {
                lex.setTokenPosition(token);
                tok = T_ENCAPSED_AND_WHITESPACE;
                fbreak;
            };
        *|;
        
        template_string := |*
            "{$"               => {lex.ungetCnt(1); lex.setTokenPosition(token); tok = T_CURLY_OPEN; lex.call(ftargs, fentry(php)); goto _out;};
            "${"               => {lex.setTokenPosition(token); tok = T_DOLLAR_OPEN_CURLY_BRACES; lex.call(ftargs, fentry(string_var_name)); goto _out;};
            "$" varname_first  => {lex.ungetCnt(2); fcall string_var;};
            '"'                => {lex.setTokenPosition(token); tok = TokenID(int('"')); fnext php; fbreak;};
            any_line* when is_not_string_end_or_var => {
                lex.setTokenPosition(token);
                tok = T_ENCAPSED_AND_WHITESPACE;
                fbreak;
            };
        *|;

        heredoc_end := |*
            varname -- ";" => {
                lex.setTokenPosition(token);
                tok = T_END_HEREDOC;
                fnext php;
                fbreak;
            };
            varname => {
                lex.setTokenPosition(token);
                tok = T_END_HEREDOC;
                fnext php;
                fbreak;
            };
        *|;
        
        string_var := |*
            '$' varname        => {lex.setTokenPosition(token); tok = T_VARIABLE; fbreak;};
            '->' varname_first => {lex.ungetCnt(1); lex.setTokenPosition(token); tok = T_OBJECT_OPERATOR; fbreak;};
            varname            => {lex.setTokenPosition(token); tok = T_STRING; fbreak;};
            '['                => {lex.setTokenPosition(token); tok = TokenID(int('[')); lex.call(ftargs, fentry(string_var_index)); goto _out;};
            any                => {lex.ungetCnt(1); fret;};
        *|;
        
        string_var_index := |*
            lnum | hnum | bnum       => {lex.setTokenPosition(token); tok = T_NUM_STRING; fbreak;};
            '$' varname              => {lex.setTokenPosition(token); tok = T_VARIABLE; fbreak;};
            varname                  => {lex.setTokenPosition(token); tok = T_STRING; fbreak;};
            whitespace_line | [\\'#] => {lex.setTokenPosition(token); tok = T_ENCAPSED_AND_WHITESPACE; lex.ret(2); goto _out;};
            operators > (svi, 1)     => {lex.setTokenPosition(token); tok = TokenID(int(lex.data[lex.ts])); fbreak;};
            ']'       > (svi, 2)     => {lex.setTokenPosition(token); tok = TokenID(int(']')); lex.ret(2); goto _out;};
            any_line => {
                c := lex.data[lex.p]
                lex.Error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c));
            };
        *|;

        string_var_name := |*
            varname ("[" | "}") => {lex.ungetCnt(1); lex.setTokenPosition(token); tok = T_STRING_VARNAME; fnext php; fbreak;};
            any                 => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_open_parenthesis := |*
            whitespace_line* => {lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)};
            "("              => {lex.setTokenPosition(token); tok = TokenID(int('(')); fnext halt_compiller_close_parenthesis; fbreak;};
            any              => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_close_parenthesis := |*
            whitespace_line* => {lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)};
            ")"              => {lex.setTokenPosition(token); tok = TokenID(int(')')); fnext halt_compiller_close_semicolon; fbreak;};
            any              => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_close_semicolon := |*
            whitespace_line* => {lex.addFreeFloating(freefloating.WhiteSpaceType, lex.ts, lex.te)};
            ";"              => {lex.setTokenPosition(token); tok = TokenID(int(';')); fnext halt_compiller_end; fbreak;};
            any              => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_end := |*
            any_line* => { lex.addFreeFloating(freefloating.TokenType, lex.ts, lex.te); };
        *|;

        write exec;
    }%%

    token.FreeFloating = lex.FreeFloating
	token.Value = string(lex.data[lex.ts:lex.te])

    lval.Token(token)

    return int(tok);
}