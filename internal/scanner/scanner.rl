package scanner

import (
    "fmt"
    "strconv"
    "strings"

    "github.com/z7zmey/php-parser/pkg/token"
)

%%{ 
    machine lexer;
    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;
}%%

func initLexer(lex *Lexer)  {
    %% write init;
}

func (lex *Lexer) Lex() *token.Token {
    eof := lex.pe
    var tok token.ID

    tkn := lex.tokenPool.Get()

    lblStart := 0
    lblEnd   := 0

    _, _ = lblStart, lblEnd

    %%{ 
        action heredoc_lbl_start {lblStart = lex.p}
        action heredoc_lbl_end   {lblEnd = lex.p}

        action new_line   {
            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        }

        action is_not_heredoc_end { lex.isNotHeredocEnd(lex.p) }
        action is_not_comment_end { lex.isNotPhpCloseToken() && lex.isNotNewLine()  }
        action is_not_heredoc_end_or_var { lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() }
        action is_not_string_end_or_var { lex.isNotStringEnd('"') && lex.isNotStringVar() }
        action is_not_backqoute_end_or_var { lex.isNotStringEnd('`') && lex.isNotStringVar() }

        newline = ('\r\n' >(nl, 1) | '\r' >(nl, 0) | '\n' >(nl, 0)) $new_line %{};
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
                (any - [\\'\r\n]) -> qoute
                | "\r" @new_line  -> qoute
                | "\n" @new_line  -> qoute
                | "\\"            -> qoute_any
                | "'"             -> final
            ),
            qoute_any: (
                (any - [\r\n])   -> qoute
                | "\r" @new_line -> qoute
                | "\n" @new_line -> qoute
            ),

            # double qoute string

            double_qoute: (
                (any - [\\"${\r\n]) -> double_qoute
                | "\r" @new_line    -> double_qoute
                | "\n" @new_line    -> double_qoute
                | "\\"              -> double_qoute_any
                | '"'               -> final
                | '$'               -> double_qoute_nonvarname
                | '{'               -> double_qoute_nondollar
            ),
            double_qoute_any: (
                (any - [\r\n])     -> double_qoute
                | "\r" @new_line   -> double_qoute
                | "\n" @new_line   -> double_qoute
            ),
            double_qoute_nondollar: (
                (any - [\\$"\r\n]) -> double_qoute
                | "\r" @new_line   -> double_qoute
                | "\n" @new_line   -> double_qoute
                | "\\"             -> double_qoute_any
                | '"'              -> final
            ),
            double_qoute_nonvarname: (
                (any - [\\${"\r\n] - varname_first) -> double_qoute
                | "\r" @new_line                    -> double_qoute
                | "\n" @new_line                    -> double_qoute
                | "\\"                              -> double_qoute_any
                | '$'                               -> double_qoute_nonvarname
                | '"'                               -> final
            );

        main := |*
            "#!" any* :>> newline => {
                lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
            };
            any => {
                fnext html;
                lex.ungetCnt(1)
            };
        *|;

        html := |*
            any_line+ -- '<?' => {
                lex.ungetStr("<")
                lex.setTokenPosition(tkn)
                tok = token.T_INLINE_HTML;
                fbreak;
            };
            '<?' => {
                lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.te)
                fnext php;
            };
            '<?php'i ( [ \t] | newline ) => {
                lex.ungetCnt(lex.te - lex.ts - 5)
                lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.ts+5)
                fnext php;
            };
            '<?='i => {
                lex.setTokenPosition(tkn);
                tok = token.T_ECHO;
                fnext php;
                fbreak;
            };
        *|;

        php := |*
            whitespace_line*                   => {lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)};
            '?>' newline?                      => {lex.setTokenPosition(tkn); tok = token.ID(int(';')); fnext html; fbreak;};
            ';' whitespace_line* '?>' newline? => {lex.setTokenPosition(tkn); tok = token.ID(int(';')); fnext html; fbreak;};

            (dnum | exponent_dnum)          => {lex.setTokenPosition(tkn); tok = token.T_DNUMBER; fbreak;};
            bnum => {
                s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
                _, err := strconv.ParseInt(s, 2, 0)

                if err == nil {
                    lex.setTokenPosition(tkn); tok = token.T_LNUMBER; fbreak;
                } 
                
                lex.setTokenPosition(tkn); tok = token.T_DNUMBER; fbreak;
            };
            lnum => {
                base := 10
                if lex.data[lex.ts] == '0' {
                    base = 8
                }

                s := strings.Replace(string(lex.data[lex.ts:lex.te]), "_", "", -1)
                _, err := strconv.ParseInt(s, base, 0)

                if err == nil {
                    lex.setTokenPosition(tkn); tok = token.T_LNUMBER; fbreak;
                } 
                
                lex.setTokenPosition(tkn); tok = token.T_DNUMBER; fbreak;
            };
            hnum => {
                s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
                _, err := strconv.ParseInt(s, 16, 0)

                if err == nil {
                    lex.setTokenPosition(tkn); tok = token.T_LNUMBER; fbreak;
                } 
                
                lex.setTokenPosition(tkn); tok = token.T_DNUMBER; fbreak;
            };

            'abstract'i                       => {lex.setTokenPosition(tkn); tok = token.T_ABSTRACT; fbreak;};
            'array'i                          => {lex.setTokenPosition(tkn); tok = token.T_ARRAY; fbreak;};
            'as'i                             => {lex.setTokenPosition(tkn); tok = token.T_AS; fbreak;};
            'break'i                          => {lex.setTokenPosition(tkn); tok = token.T_BREAK; fbreak;};
            'callable'i                       => {lex.setTokenPosition(tkn); tok = token.T_CALLABLE; fbreak;};
            'case'i                           => {lex.setTokenPosition(tkn); tok = token.T_CASE; fbreak;};
            'catch'i                          => {lex.setTokenPosition(tkn); tok = token.T_CATCH; fbreak;};
            'class'i                          => {lex.setTokenPosition(tkn); tok = token.T_CLASS; fbreak;};
            'clone'i                          => {lex.setTokenPosition(tkn); tok = token.T_CLONE; fbreak;};
            'const'i                          => {lex.setTokenPosition(tkn); tok = token.T_CONST; fbreak;};
            'continue'i                       => {lex.setTokenPosition(tkn); tok = token.T_CONTINUE; fbreak;};
            'declare'i                        => {lex.setTokenPosition(tkn); tok = token.T_DECLARE; fbreak;};
            'default'i                        => {lex.setTokenPosition(tkn); tok = token.T_DEFAULT; fbreak;};
            'do'i                             => {lex.setTokenPosition(tkn); tok = token.T_DO; fbreak;};
            'echo'i                           => {lex.setTokenPosition(tkn); tok = token.T_ECHO; fbreak;};
            'else'i                           => {lex.setTokenPosition(tkn); tok = token.T_ELSE; fbreak;};
            'elseif'i                         => {lex.setTokenPosition(tkn); tok = token.T_ELSEIF; fbreak;};
            'empty'i                          => {lex.setTokenPosition(tkn); tok = token.T_EMPTY; fbreak;};
            'enddeclare'i                     => {lex.setTokenPosition(tkn); tok = token.T_ENDDECLARE; fbreak;};
            'endfor'i                         => {lex.setTokenPosition(tkn); tok = token.T_ENDFOR; fbreak;};
            'endforeach'i                     => {lex.setTokenPosition(tkn); tok = token.T_ENDFOREACH; fbreak;};
            'endif'i                          => {lex.setTokenPosition(tkn); tok = token.T_ENDIF; fbreak;};
            'endswitch'i                      => {lex.setTokenPosition(tkn); tok = token.T_ENDSWITCH; fbreak;};
            'endwhile'i                       => {lex.setTokenPosition(tkn); tok = token.T_ENDWHILE; fbreak;};
            'eval'i                           => {lex.setTokenPosition(tkn); tok = token.T_EVAL; fbreak;};
            'exit'i | 'die'i                  => {lex.setTokenPosition(tkn); tok = token.T_EXIT; fbreak;};
            'extends'i                        => {lex.setTokenPosition(tkn); tok = token.T_EXTENDS; fbreak;};
            'final'i                          => {lex.setTokenPosition(tkn); tok = token.T_FINAL; fbreak;};
            'finally'i                        => {lex.setTokenPosition(tkn); tok = token.T_FINALLY; fbreak;};
            'for'i                            => {lex.setTokenPosition(tkn); tok = token.T_FOR; fbreak;};
            'foreach'i                        => {lex.setTokenPosition(tkn); tok = token.T_FOREACH; fbreak;};
            'function'i | 'cfunction'i        => {lex.setTokenPosition(tkn); tok = token.T_FUNCTION; fbreak;};
            'fn'i                             => {lex.setTokenPosition(tkn); tok = token.T_FN; fbreak;};
            'global'i                         => {lex.setTokenPosition(tkn); tok = token.T_GLOBAL; fbreak;};
            'goto'i                           => {lex.setTokenPosition(tkn); tok = token.T_GOTO; fbreak;};
            'if'i                             => {lex.setTokenPosition(tkn); tok = token.T_IF; fbreak;};
            'isset'i                          => {lex.setTokenPosition(tkn); tok = token.T_ISSET; fbreak;};
            'implements'i                     => {lex.setTokenPosition(tkn); tok = token.T_IMPLEMENTS; fbreak;};
            'instanceof'i                     => {lex.setTokenPosition(tkn); tok = token.T_INSTANCEOF; fbreak;};
            'insteadof'i                      => {lex.setTokenPosition(tkn); tok = token.T_INSTEADOF; fbreak;};
            'interface'i                      => {lex.setTokenPosition(tkn); tok = token.T_INTERFACE; fbreak;};
            'list'i                           => {lex.setTokenPosition(tkn); tok = token.T_LIST; fbreak;};
            'namespace'i                      => {lex.setTokenPosition(tkn); tok = token.T_NAMESPACE; fbreak;};
            'private'i                        => {lex.setTokenPosition(tkn); tok = token.T_PRIVATE; fbreak;};
            'public'i                         => {lex.setTokenPosition(tkn); tok = token.T_PUBLIC; fbreak;};
            'print'i                          => {lex.setTokenPosition(tkn); tok = token.T_PRINT; fbreak;};
            'protected'i                      => {lex.setTokenPosition(tkn); tok = token.T_PROTECTED; fbreak;};
            'return'i                         => {lex.setTokenPosition(tkn); tok = token.T_RETURN; fbreak;};
            'static'i                         => {lex.setTokenPosition(tkn); tok = token.T_STATIC; fbreak;};
            'switch'i                         => {lex.setTokenPosition(tkn); tok = token.T_SWITCH; fbreak;};
            'throw'i                          => {lex.setTokenPosition(tkn); tok = token.T_THROW; fbreak;};
            'trait'i                          => {lex.setTokenPosition(tkn); tok = token.T_TRAIT; fbreak;};
            'try'i                            => {lex.setTokenPosition(tkn); tok = token.T_TRY; fbreak;};
            'unset'i                          => {lex.setTokenPosition(tkn); tok = token.T_UNSET; fbreak;};
            'use'i                            => {lex.setTokenPosition(tkn); tok = token.T_USE; fbreak;};
            'var'i                            => {lex.setTokenPosition(tkn); tok = token.T_VAR; fbreak;};
            'while'i                          => {lex.setTokenPosition(tkn); tok = token.T_WHILE; fbreak;};
            'yield'i whitespace_line+ 'from'i => {lex.setTokenPosition(tkn); tok = token.T_YIELD_FROM; fbreak;};
            'yield'i                          => {lex.setTokenPosition(tkn); tok = token.T_YIELD; fbreak;};
            'include'i                        => {lex.setTokenPosition(tkn); tok = token.T_INCLUDE; fbreak;};
            'include_once'i                   => {lex.setTokenPosition(tkn); tok = token.T_INCLUDE_ONCE; fbreak;};
            'require'i                        => {lex.setTokenPosition(tkn); tok = token.T_REQUIRE; fbreak;};
            'require_once'i                   => {lex.setTokenPosition(tkn); tok = token.T_REQUIRE_ONCE; fbreak;};
            '__CLASS__'i                      => {lex.setTokenPosition(tkn); tok = token.T_CLASS_C; fbreak;};
            '__DIR__'i                        => {lex.setTokenPosition(tkn); tok = token.T_DIR; fbreak;};
            '__FILE__'i                       => {lex.setTokenPosition(tkn); tok = token.T_FILE; fbreak;};
            '__FUNCTION__'i                   => {lex.setTokenPosition(tkn); tok = token.T_FUNC_C; fbreak;};
            '__LINE__'i                       => {lex.setTokenPosition(tkn); tok = token.T_LINE; fbreak;};
            '__NAMESPACE__'i                  => {lex.setTokenPosition(tkn); tok = token.T_NS_C; fbreak;};
            '__METHOD__'i                     => {lex.setTokenPosition(tkn); tok = token.T_METHOD_C; fbreak;};
            '__TRAIT__'i                      => {lex.setTokenPosition(tkn); tok = token.T_TRAIT_C; fbreak;};
            '__halt_compiler'i                => {lex.setTokenPosition(tkn); tok = token.T_HALT_COMPILER; fnext halt_compiller_open_parenthesis; fbreak;};
            'new'i                            => {lex.setTokenPosition(tkn); tok = token.T_NEW; fbreak;};
            'and'i                            => {lex.setTokenPosition(tkn); tok = token.T_LOGICAL_AND; fbreak;};
            'or'i                             => {lex.setTokenPosition(tkn); tok = token.T_LOGICAL_OR; fbreak;};
            'xor'i                            => {lex.setTokenPosition(tkn); tok = token.T_LOGICAL_XOR; fbreak;};
            '\\'                              => {lex.setTokenPosition(tkn); tok = token.T_NS_SEPARATOR; fbreak;};
            '...'                             => {lex.setTokenPosition(tkn); tok = token.T_ELLIPSIS; fbreak;};
            '::'                              => {lex.setTokenPosition(tkn); tok = token.T_PAAMAYIM_NEKUDOTAYIM; fbreak;};
            '&&'                              => {lex.setTokenPosition(tkn); tok = token.T_BOOLEAN_AND; fbreak;};
            '||'                              => {lex.setTokenPosition(tkn); tok = token.T_BOOLEAN_OR; fbreak;};
            '&='                              => {lex.setTokenPosition(tkn); tok = token.T_AND_EQUAL; fbreak;};
            '|='                              => {lex.setTokenPosition(tkn); tok = token.T_OR_EQUAL; fbreak;};
            '.='                              => {lex.setTokenPosition(tkn); tok = token.T_CONCAT_EQUAL; fbreak;};
            '*='                              => {lex.setTokenPosition(tkn); tok = token.T_MUL_EQUAL; fbreak;};
            '**='                             => {lex.setTokenPosition(tkn); tok = token.T_POW_EQUAL; fbreak;};
            '/='                              => {lex.setTokenPosition(tkn); tok = token.T_DIV_EQUAL; fbreak;};
            '+='                              => {lex.setTokenPosition(tkn); tok = token.T_PLUS_EQUAL; fbreak;};
            '-='                              => {lex.setTokenPosition(tkn); tok = token.T_MINUS_EQUAL; fbreak;};
            '^='                              => {lex.setTokenPosition(tkn); tok = token.T_XOR_EQUAL; fbreak;};
            '%='                              => {lex.setTokenPosition(tkn); tok = token.T_MOD_EQUAL; fbreak;};
            '--'                              => {lex.setTokenPosition(tkn); tok = token.T_DEC; fbreak;};
            '++'                              => {lex.setTokenPosition(tkn); tok = token.T_INC; fbreak;};
            '=>'                              => {lex.setTokenPosition(tkn); tok = token.T_DOUBLE_ARROW; fbreak;};
            '<=>'                             => {lex.setTokenPosition(tkn); tok = token.T_SPACESHIP; fbreak;};
            '!=' | '<>'                       => {lex.setTokenPosition(tkn); tok = token.T_IS_NOT_EQUAL; fbreak;};
            '!=='                             => {lex.setTokenPosition(tkn); tok = token.T_IS_NOT_IDENTICAL; fbreak;};
            '=='                              => {lex.setTokenPosition(tkn); tok = token.T_IS_EQUAL; fbreak;};
            '==='                             => {lex.setTokenPosition(tkn); tok = token.T_IS_IDENTICAL; fbreak;};
            '<<='                             => {lex.setTokenPosition(tkn); tok = token.T_SL_EQUAL; fbreak;};
            '>>='                             => {lex.setTokenPosition(tkn); tok = token.T_SR_EQUAL; fbreak;};
            '>='                              => {lex.setTokenPosition(tkn); tok = token.T_IS_GREATER_OR_EQUAL; fbreak;};
            '<='                              => {lex.setTokenPosition(tkn); tok = token.T_IS_SMALLER_OR_EQUAL; fbreak;};
            '**'                              => {lex.setTokenPosition(tkn); tok = token.T_POW; fbreak;};
            '<<'                              => {lex.setTokenPosition(tkn); tok = token.T_SL; fbreak;};
            '>>'                              => {lex.setTokenPosition(tkn); tok = token.T_SR; fbreak;};
            '??'                              => {lex.setTokenPosition(tkn); tok = token.T_COALESCE; fbreak;};
            '??='                             => {lex.setTokenPosition(tkn); tok = token.T_COALESCE_EQUAL; fbreak;};

            '(' whitespace* 'array'i whitespace* ')'                     => {lex.setTokenPosition(tkn); tok = token.T_ARRAY_CAST; fbreak;};
            '(' whitespace* ('bool'i|'boolean'i) whitespace* ')'         => {lex.setTokenPosition(tkn); tok = token.T_BOOL_CAST; fbreak;};
            '(' whitespace* ('real'i|'double'i|'float'i) whitespace* ')' => {lex.setTokenPosition(tkn); tok = token.T_DOUBLE_CAST; fbreak;};
            '(' whitespace* ('int'i|'integer'i) whitespace* ')'          => {lex.setTokenPosition(tkn); tok = token.T_INT_CAST; fbreak;};
            '(' whitespace* 'object'i whitespace* ')'                    => {lex.setTokenPosition(tkn); tok = token.T_OBJECT_CAST; fbreak;};
            '(' whitespace* ('string'i|'binary'i) whitespace* ')'        => {lex.setTokenPosition(tkn); tok = token.T_STRING_CAST; fbreak;};
            '(' whitespace* 'unset'i whitespace* ')'                     => {lex.setTokenPosition(tkn); tok = token.T_UNSET_CAST; fbreak;};

            ('#' | '//') any_line* when is_not_comment_end => {
                lex.ungetStr("?>")
                lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
            };
            '/*' any_line* :>> '*/' {
                isDocComment := false;
                if lex.te - lex.ts > 4 && string(lex.data[lex.ts:lex.ts+3]) == "/**" {
                    isDocComment = true;
                }

                if isDocComment {
                    lex.addFreeFloatingToken(tkn, token.T_DOC_COMMENT, lex.ts, lex.te)
                } else {
                    lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
                }
            };

            operators => {
                lex.setTokenPosition(tkn);
                tok = token.ID(int(lex.data[lex.ts]));
                fbreak;
            };

            "{"          => { lex.setTokenPosition(tkn); tok = token.ID(int('{')); lex.call(ftargs, fentry(php)); goto _out; };
            "}"          => { lex.setTokenPosition(tkn); tok = token.ID(int('}')); lex.ret(1); goto _out;};
            "$" varname  => { lex.setTokenPosition(tkn); tok = token.T_VARIABLE; fbreak; };
            varname      => { lex.setTokenPosition(tkn); tok = token.T_STRING;   fbreak; };

            "->"         => { lex.setTokenPosition(tkn); tok = token.T_OBJECT_OPERATOR; fnext property; fbreak; };

            constant_string => {
                lex.setTokenPosition(tkn);
                tok = token.T_CONSTANT_ENCAPSED_STRING;
                fbreak;
            };

            "b"i? "<<<" [ \t]* ( heredoc_label | ("'" heredoc_label "'") | ('"' heredoc_label '"') ) newline  => {
                lex.heredocLabel = lex.data[lblStart:lblEnd]
                lex.setTokenPosition(tkn);
                tok = token.T_START_HEREDOC;

                if lex.isHeredocEnd(lex.p+1) {
                    fnext heredoc_end;
                } else if lex.data[lblStart-1] == '\'' {
                    fnext nowdoc;
                } else {
                    fnext heredoc;
                }
                fbreak;
            };
            "`" => {lex.setTokenPosition(tkn); tok = token.ID(int('`')); fnext backqote; fbreak;};
            '"' => {lex.setTokenPosition(tkn); tok = token.ID(int('"')); fnext template_string; fbreak;};

            any_line => {
                c := lex.data[lex.p]
                lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c));
            };
        *|;

        property := |*
            whitespace_line* => {lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)};
            "->"             => {lex.setTokenPosition(tkn); tok = token.T_OBJECT_OPERATOR; fbreak;};
            varname          => {lex.setTokenPosition(tkn); tok = token.T_STRING; fnext php; fbreak;};
            any              => {lex.ungetCnt(1); fgoto php;};
        *|;

        nowdoc := |*
            any_line* when is_not_heredoc_end => {
                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;
                fnext heredoc_end;
                fbreak;
            };
        *|;
        
        heredoc := |*
            "{$" => {lex.ungetCnt(1); lex.setTokenPosition(tkn); tok = token.T_CURLY_OPEN; lex.call(ftargs, fentry(php)); goto _out;};
            "${" => {lex.setTokenPosition(tkn); tok = token.T_DOLLAR_OPEN_CURLY_BRACES; lex.call(ftargs, fentry(string_var_name)); goto _out;};
            "$"  => {lex.ungetCnt(1); fcall string_var;};
            any_line* when is_not_heredoc_end_or_var => {
                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;

                if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
                    fnext heredoc_end;
                }
                fbreak;
            };
        *|;
        
        backqote := |*
            "{$"              => {lex.ungetCnt(1); lex.setTokenPosition(tkn); tok = token.T_CURLY_OPEN; lex.call(ftargs, fentry(php)); goto _out;};
            "${"              => {lex.setTokenPosition(tkn); tok = token.T_DOLLAR_OPEN_CURLY_BRACES; lex.call(ftargs, fentry(string_var_name)); goto _out;};
            "$" varname_first => {lex.ungetCnt(2); fcall string_var;};
            '`'               => {lex.setTokenPosition(tkn); tok = token.ID(int('`')); fnext php; fbreak;};
            any_line* when is_not_backqoute_end_or_var => {
                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;
                fbreak;
            };
        *|;
        
        template_string := |*
            "{$"               => {lex.ungetCnt(1); lex.setTokenPosition(tkn); tok = token.T_CURLY_OPEN; lex.call(ftargs, fentry(php)); goto _out;};
            "${"               => {lex.setTokenPosition(tkn); tok = token.T_DOLLAR_OPEN_CURLY_BRACES; lex.call(ftargs, fentry(string_var_name)); goto _out;};
            "$" varname_first  => {lex.ungetCnt(2); fcall string_var;};
            '"'                => {lex.setTokenPosition(tkn); tok = token.ID(int('"')); fnext php; fbreak;};
            any_line* when is_not_string_end_or_var => {
                lex.setTokenPosition(tkn);
                tok = token.T_ENCAPSED_AND_WHITESPACE;
                fbreak;
            };
        *|;

        heredoc_end := |*
            varname -- ";" => {
                lex.setTokenPosition(tkn);
                tok = token.T_END_HEREDOC;
                fnext php;
                fbreak;
            };
            varname => {
                lex.setTokenPosition(tkn);
                tok = token.T_END_HEREDOC;
                fnext php;
                fbreak;
            };
        *|;
        
        string_var := |*
            '$' varname        => {lex.setTokenPosition(tkn); tok = token.T_VARIABLE; fbreak;};
            '->' varname_first => {lex.ungetCnt(1); lex.setTokenPosition(tkn); tok = token.T_OBJECT_OPERATOR; fbreak;};
            varname            => {lex.setTokenPosition(tkn); tok = token.T_STRING; fbreak;};
            '['                => {lex.setTokenPosition(tkn); tok = token.ID(int('[')); lex.call(ftargs, fentry(string_var_index)); goto _out;};
            any                => {lex.ungetCnt(1); fret;};
        *|;
        
        string_var_index := |*
            lnum | hnum | bnum       => {lex.setTokenPosition(tkn); tok = token.T_NUM_STRING; fbreak;};
            '$' varname              => {lex.setTokenPosition(tkn); tok = token.T_VARIABLE; fbreak;};
            varname                  => {lex.setTokenPosition(tkn); tok = token.T_STRING; fbreak;};
            whitespace_line | [\\'#] => {lex.setTokenPosition(tkn); tok = token.T_ENCAPSED_AND_WHITESPACE; lex.ret(2); goto _out;};
            operators > (svi, 1)     => {lex.setTokenPosition(tkn); tok = token.ID(int(lex.data[lex.ts])); fbreak;};
            ']'       > (svi, 2)     => {lex.setTokenPosition(tkn); tok = token.ID(int(']')); lex.ret(2); goto _out;};
            any_line => {
                c := lex.data[lex.p]
                lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c));
            };
        *|;

        string_var_name := |*
            varname ("[" | "}") => {lex.ungetCnt(1); lex.setTokenPosition(tkn); tok = token.T_STRING_VARNAME; fnext php; fbreak;};
            any                 => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_open_parenthesis := |*
            whitespace_line* => {lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)};
            "("              => {lex.setTokenPosition(tkn); tok = token.ID(int('(')); fnext halt_compiller_close_parenthesis; fbreak;};
            any              => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_close_parenthesis := |*
            whitespace_line* => {lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)};
            ")"              => {lex.setTokenPosition(tkn); tok = token.ID(int(')')); fnext halt_compiller_close_semicolon; fbreak;};
            any              => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_close_semicolon := |*
            whitespace_line* => {lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)};
            ";"              => {lex.setTokenPosition(tkn); tok = token.ID(int(';')); fnext halt_compiller_end; fbreak;};
            any              => {lex.ungetCnt(1); fnext php;};
        *|;

        halt_compiller_end := |*
            any_line* => { lex.addFreeFloatingToken(tkn, token.T_HALT_COMPILER, lex.ts, lex.te); };
        *|;

        write exec;
    }%%

    tkn.Value = lex.data[lex.ts:lex.te]
    tkn.ID = token.ID(tok)

    return tkn
}