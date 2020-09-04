%{
package php7

import (
    "bytes"
    "strconv"

    "github.com/z7zmey/php-parser/internal/position"
    "github.com/z7zmey/php-parser/pkg/ast"
    "github.com/z7zmey/php-parser/pkg/token"
)

%}

%union{
    node             ast.Vertex
    token            *token.Token
    tkn              *token.Token
    list             []ast.Vertex

    ClassExtends     *ast.StmtClassExtends
    ClassImplements  *ast.StmtClassImplements
    InterfaceExtends *ast.StmtInterfaceExtends
    ClosureUse       *ast.ExprClosureUse
}

%token <token> T_INCLUDE
%token <token> T_INCLUDE_ONCE
%token <token> T_EXIT
%token <token> T_IF
%token <token> T_LNUMBER
%token <token> T_DNUMBER
%token <token> T_STRING
%token <token> T_STRING_VARNAME
%token <token> T_VARIABLE
%token <token> T_NUM_STRING
%token <token> T_INLINE_HTML
%token <token> T_CHARACTER
%token <token> T_BAD_CHARACTER
%token <token> T_ENCAPSED_AND_WHITESPACE
%token <token> T_CONSTANT_ENCAPSED_STRING
%token <token> T_ECHO
%token <token> T_DO
%token <token> T_WHILE
%token <token> T_ENDWHILE
%token <token> T_FOR
%token <token> T_ENDFOR
%token <token> T_FOREACH
%token <token> T_ENDFOREACH
%token <token> T_DECLARE
%token <token> T_ENDDECLARE
%token <token> T_AS
%token <token> T_SWITCH
%token <token> T_ENDSWITCH
%token <token> T_CASE
%token <token> T_DEFAULT
%token <token> T_BREAK
%token <token> T_CONTINUE
%token <token> T_GOTO
%token <token> T_FUNCTION
%token <token> T_FN
%token <token> T_CONST
%token <token> T_RETURN
%token <token> T_TRY
%token <token> T_CATCH
%token <token> T_FINALLY
%token <token> T_THROW
%token <token> T_USE
%token <token> T_INSTEADOF
%token <token> T_GLOBAL
%token <token> T_VAR
%token <token> T_UNSET
%token <token> T_ISSET
%token <token> T_EMPTY
%token <token> T_HALT_COMPILER
%token <token> T_CLASS
%token <token> T_TRAIT
%token <token> T_INTERFACE
%token <token> T_EXTENDS
%token <token> T_IMPLEMENTS
%token <token> T_OBJECT_OPERATOR
%token <token> T_DOUBLE_ARROW
%token <token> T_LIST
%token <token> T_ARRAY
%token <token> T_CALLABLE
%token <token> T_CLASS_C
%token <token> T_TRAIT_C
%token <token> T_METHOD_C
%token <token> T_FUNC_C
%token <token> T_LINE
%token <token> T_FILE
%token <token> T_COMMENT
%token <token> T_DOC_COMMENT
%token <token> T_OPEN_TAG
%token <token> T_OPEN_TAG_WITH_ECHO
%token <token> T_CLOSE_TAG
%token <token> T_WHITESPACE
%token <token> T_START_HEREDOC
%token <token> T_END_HEREDOC
%token <token> T_DOLLAR_OPEN_CURLY_BRACES
%token <token> T_CURLY_OPEN
%token <token> T_PAAMAYIM_NEKUDOTAYIM
%token <token> T_NAMESPACE
%token <token> T_NS_C
%token <token> T_DIR
%token <token> T_NS_SEPARATOR
%token <token> T_ELLIPSIS
%token <token> T_EVAL
%token <token> T_REQUIRE
%token <token> T_REQUIRE_ONCE
%token <token> T_LOGICAL_OR
%token <token> T_LOGICAL_XOR
%token <token> T_LOGICAL_AND
%token <token> T_INSTANCEOF
%token <token> T_NEW
%token <token> T_CLONE
%token <token> T_ELSEIF
%token <token> T_ELSE
%token <token> T_ENDIF
%token <token> T_PRINT
%token <token> T_YIELD
%token <token> T_STATIC
%token <token> T_ABSTRACT
%token <token> T_FINAL
%token <token> T_PRIVATE
%token <token> T_PROTECTED
%token <token> T_PUBLIC
%token <token> T_INC
%token <token> T_DEC
%token <token> T_YIELD_FROM
%token <token> T_INT_CAST
%token <token> T_DOUBLE_CAST
%token <token> T_STRING_CAST
%token <token> T_ARRAY_CAST
%token <token> T_OBJECT_CAST
%token <token> T_BOOL_CAST
%token <token> T_UNSET_CAST
%token <token> T_COALESCE
%token <token> T_SPACESHIP
%token <token> T_NOELSE
%token <token> T_PLUS_EQUAL
%token <token> T_MINUS_EQUAL
%token <token> T_MUL_EQUAL
%token <token> T_POW_EQUAL
%token <token> T_DIV_EQUAL
%token <token> T_CONCAT_EQUAL
%token <token> T_MOD_EQUAL
%token <token> T_AND_EQUAL
%token <token> T_OR_EQUAL
%token <token> T_XOR_EQUAL
%token <token> T_SL_EQUAL
%token <token> T_SR_EQUAL
%token <token> T_COALESCE_EQUAL
%token <token> T_BOOLEAN_OR
%token <token> T_BOOLEAN_AND
%token <token> T_POW
%token <token> T_SL
%token <token> T_SR
%token <token> T_IS_IDENTICAL
%token <token> T_IS_NOT_IDENTICAL
%token <token> T_IS_EQUAL
%token <token> T_IS_NOT_EQUAL
%token <token> T_IS_SMALLER_OR_EQUAL
%token <token> T_IS_GREATER_OR_EQUAL
%token <token> '"'
%token <token> '`'
%token <token> '{'
%token <token> '}'
%token <token> ';'
%token <token> ':'
%token <token> '('
%token <token> ')'
%token <token> '['
%token <token> ']'
%token <token> '?'
%token <token> '&'
%token <token> '-'
%token <token> '+'
%token <token> '!'
%token <token> '~'
%token <token> '@'
%token <token> '$'
%token <token> ','
%token <token> '|'
%token <token> '='
%token <token> '^'
%token <token> '*'
%token <token> '/'
%token <token> '%'
%token <token> '<'
%token <token> '>'
%token <token> '.'

%left T_INCLUDE T_INCLUDE_ONCE T_EVAL T_REQUIRE T_REQUIRE_ONCE
%left ','
%left T_LOGICAL_OR
%left T_LOGICAL_XOR
%left T_LOGICAL_AND
%right T_PRINT
%right T_YIELD
%right T_DOUBLE_ARROW
%right T_YIELD_FROM
%left '=' T_PLUS_EQUAL T_MINUS_EQUAL T_MUL_EQUAL T_DIV_EQUAL T_CONCAT_EQUAL T_MOD_EQUAL T_AND_EQUAL T_OR_EQUAL T_XOR_EQUAL T_SL_EQUAL T_SR_EQUAL T_POW_EQUAL T_COALESCE_EQUAL
%left '?' ':'
%right T_COALESCE
%left T_BOOLEAN_OR
%left T_BOOLEAN_AND
%left '|'
%left '^'
%left '&'
%nonassoc T_IS_EQUAL T_IS_NOT_EQUAL T_IS_IDENTICAL T_IS_NOT_IDENTICAL T_SPACESHIP
%nonassoc '<' T_IS_SMALLER_OR_EQUAL '>' T_IS_GREATER_OR_EQUAL
%left T_SL T_SR
%left '+' '-' '.'
%left '*' '/' '%'
%right '!'
%nonassoc T_INSTANCEOF
%right '~' T_INC T_DEC T_INT_CAST T_DOUBLE_CAST T_STRING_CAST T_ARRAY_CAST T_OBJECT_CAST T_BOOL_CAST T_UNSET_CAST '@'
%right T_POW
%right '['
%nonassoc T_NEW T_CLONE
%left T_NOELSE
%left T_ELSEIF
%left T_ELSE
%left T_ENDIF
%right T_STATIC T_ABSTRACT T_FINAL T_PRIVATE T_PROTECTED T_PUBLIC

%type <token> is_reference is_variadic returns_ref

%type <token> reserved_non_modifiers
%type <token> semi_reserved
%type <token> identifier
%type <token> possible_comma
%type <token> case_separator

%type <node> top_statement name statement function_declaration_statement
%type <node> class_declaration_statement trait_declaration_statement
%type <node> interface_declaration_statement
%type <node> group_use_declaration inline_use_declaration
%type <node> mixed_group_use_declaration use_declaration unprefixed_use_declaration
%type <node> const_decl inner_statement
%type <node> expr optional_expr
%type <node> declare_statement finally_statement unset_variable variable
%type <node> parameter optional_type argument expr_without_variable global_var
%type <node> static_var class_statement trait_adaptation trait_precedence trait_alias
%type <node> absolute_trait_method_reference trait_method_reference property echo_expr
%type <node> new_expr anonymous_class class_name class_name_reference simple_variable
%type <node> internal_functions_in_yacc
%type <node> exit_expr scalar lexical_var function_call member_name property_name
%type <node> variable_class_name dereferencable_scalar constant dereferencable
%type <node> callable_expr callable_variable static_member new_variable
%type <node> encaps_var encaps_var_offset
%type <node> if_stmt
%type <node> alt_if_stmt
%type <node> if_stmt_without_else
%type <node> class_const_decl
%type <node> alt_if_stmt_without_else
%type <node> array_pair possible_array_pair
%type <node> isset_variable type return_type type_expr
%type <node> class_modifier
%type <node> argument_list ctor_arguments
%type <node> trait_adaptations
%type <node> switch_case_list
%type <node> method_body
%type <node> foreach_statement for_statement while_statement
%type <node> inline_function
%type <ClassExtends> extends_from
%type <ClassImplements> implements_list
%type <InterfaceExtends> interface_extends_list
%type <ClosureUse> lexical_vars

%type <node> member_modifier
%type <node> use_type
%type <node> foreach_variable


%type <list> encaps_list backticks_expr namespace_name catch_name_list catch_list class_const_list
%type <list> const_list echo_expr_list for_exprs non_empty_for_exprs global_var_list
%type <list> unprefixed_use_declarations inline_use_declarations property_list static_var_list
%type <list> case_list trait_adaptation_list unset_variables
%type <list> use_declarations lexical_var_list isset_variables non_empty_array_pair_list
%type <list> array_pair_list non_empty_argument_list top_statement_list
%type <list> inner_statement_list parameter_list non_empty_parameter_list class_statement_list
%type <list> method_modifiers variable_modifiers
%type <list> non_empty_member_modifiers name_list class_modifiers

%%

/////////////////////////////////////////////////////////////////////////

start:
        top_statement_list
            {
                yylex.(*Parser).rootNode = &ast.Root{ast.Node{}, $1}

                // save position
                yylex.(*Parser).rootNode.GetNode().Position = position.NewNodeListPosition($1)

                yylex.(*Parser).setFreeFloating(yylex.(*Parser).rootNode, token.End, yylex.(*Parser).currentToken.SkippedTokens)
            }
;

reserved_non_modifiers:
      T_INCLUDE {$$=$1} | T_INCLUDE_ONCE {$$=$1} | T_EVAL {$$=$1} | T_REQUIRE {$$=$1} | T_REQUIRE_ONCE {$$=$1} | T_LOGICAL_OR {$$=$1} | T_LOGICAL_XOR {$$=$1} | T_LOGICAL_AND {$$=$1} 
    | T_INSTANCEOF {$$=$1} | T_NEW {$$=$1} | T_CLONE {$$=$1} | T_EXIT {$$=$1} | T_IF {$$=$1} | T_ELSEIF {$$=$1} | T_ELSE {$$=$1} | T_ENDIF {$$=$1} | T_ECHO {$$=$1} | T_DO {$$=$1} | T_WHILE {$$=$1} | T_ENDWHILE {$$=$1} 
    | T_FOR {$$=$1} | T_ENDFOR {$$=$1} | T_FOREACH {$$=$1} | T_ENDFOREACH {$$=$1} | T_DECLARE {$$=$1} | T_ENDDECLARE {$$=$1} | T_AS {$$=$1} | T_TRY {$$=$1} | T_CATCH {$$=$1} | T_FINALLY {$$=$1} 
    | T_THROW {$$=$1} | T_USE {$$=$1} | T_INSTEADOF {$$=$1} | T_GLOBAL {$$=$1} | T_VAR {$$=$1} | T_UNSET {$$=$1} | T_ISSET {$$=$1} | T_EMPTY {$$=$1} | T_CONTINUE {$$=$1} | T_GOTO {$$=$1} 
    | T_FUNCTION {$$=$1} | T_CONST {$$=$1} | T_RETURN {$$=$1} | T_PRINT {$$=$1} | T_YIELD {$$=$1} | T_LIST {$$=$1} | T_SWITCH {$$=$1} | T_ENDSWITCH {$$=$1} | T_CASE {$$=$1} | T_DEFAULT {$$=$1} | T_BREAK {$$=$1} 
    | T_ARRAY {$$=$1} | T_CALLABLE {$$=$1} | T_EXTENDS {$$=$1} | T_IMPLEMENTS {$$=$1} | T_NAMESPACE {$$=$1} | T_TRAIT {$$=$1} | T_INTERFACE {$$=$1} | T_CLASS {$$=$1} 
    | T_CLASS_C {$$=$1} | T_TRAIT_C {$$=$1} | T_FUNC_C {$$=$1} | T_METHOD_C {$$=$1} | T_LINE {$$=$1} | T_FILE {$$=$1} | T_DIR {$$=$1} | T_NS_C {$$=$1} | T_FN {$$=$1}
;

semi_reserved:
        reserved_non_modifiers
            {
                $$ = $1
            }
    |   T_STATIC {$$=$1} | T_ABSTRACT {$$=$1} | T_FINAL {$$=$1} | T_PRIVATE {$$=$1} | T_PROTECTED {$$=$1} | T_PUBLIC {$$=$1}
;

identifier:
        T_STRING
            {
                $$ = $1
            }
    |   semi_reserved
            {
                $$ = $1
            }
;

top_statement_list:
        top_statement_list top_statement
            {
                if inlineHtmlNode, ok := $2.(*ast.StmtInlineHtml); ok && len($1) > 0 {
                    prevNode := lastNode($1)
                    yylex.(*Parser).splitSemiColonAndPhpCloseTag(inlineHtmlNode, prevNode)
                }

                if $2 != nil {
                    $$ = append($1, $2)
                }
            }
    |   /* empty */
            {
                $$ = []ast.Vertex{}
            }
;

namespace_name:
        T_STRING
            {
                $$ = []ast.Vertex{
                    &ast.NameNamePart{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        StringTkn: $1,
                        Value:     $1.Value,
                    },
                }
            }
    |   namespace_name T_NS_SEPARATOR T_STRING
            {
                $$ = append($1, &ast.NameNamePart{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($2, $3),
                    },
                    NsSeparatorTkn: $2,
                    StringTkn:      $3,
                    Value:          $3.Value,
                })
            }
;

name:
        namespace_name
            {
                $$ = &ast.NameName{
                    Node:  ast.Node{
                        Position: position.NewNodeListPosition($1),
                    },
                    Parts: $1,
                }
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameRelative{
                    Node:  ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $3),
                    },
                    NsTkn:          $1,
                    NsSeparatorTkn: $2,
                    Parts:          $3,
                }
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameFullyQualified{
                    Node:  ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $2),
                    },
                    NsSeparatorTkn: $1,
                    Parts:          $2,
                }
            }
;

top_statement:
        error
            {
                // error
                $$ = nil
            }
    |   statement
            {
                $$ = $1
            }
    |   function_declaration_statement
            {
                $$ = $1
            }
    |   class_declaration_statement
            {
                $$ = $1
            }
    |   trait_declaration_statement
            {
                $$ = $1
            }
    |   interface_declaration_statement
            {
                $$ = $1
            }
    |   T_HALT_COMPILER '(' ')' ';'
            {
                $$ = &ast.StmtHaltCompiler{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    HaltCompilerTkn:     $1,
                    OpenParenthesisTkn:  $2,
                    CloseParenthesisTkn: $3,
                    SemiColonTkn:        $4,
                }
            }
    |   T_NAMESPACE namespace_name ';'
            {
                $$ = &ast.StmtNamespace{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    NsTkn: $1,
                    Name: &ast.NameName{
                        Node:  ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Parts: $2,
                    },
                    SemiColonTkn: $3,
                }
            }
    |   T_NAMESPACE namespace_name '{' top_statement_list '}'
            {
                $$ = &ast.StmtNamespace{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $5),
                    },
                    NsTkn: $1,
                    Name: &ast.NameName{
                        Node:  ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Parts: $2,
                    },
                    OpenCurlyBracket:  $3,
                    Stmts:             $4,
                    CloseCurlyBracket: $5,
                }
            }
    |   T_NAMESPACE '{' top_statement_list '}'
            {
                $$ = &ast.StmtNamespace{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    NsTkn:             $1,
                    OpenCurlyBracket:  $2,
                    Stmts:             $3,
                    CloseCurlyBracket: $4,
                }
            }
    |   T_USE mixed_group_use_declaration ';'
            {
                use := $2.(*ast.StmtGroupUse)

                use.Node.Position = position.NewTokensPosition($1, $3)
                use.UseTkn        = $1
                use.SemiColonTkn  = $3

                $$ = $2
            }
    |   T_USE use_type group_use_declaration ';'
            {
                use := $3.(*ast.StmtGroupUse)

                use.Node.Position = position.NewTokensPosition($1, $4)
                use.UseTkn        = $1
                use.Type          = $2
                use.SemiColonTkn  = $4

                $$ = $3
            }
    |   T_USE use_declarations ';'
            {
                $$ = &ast.StmtUse{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    UseTkn:          $1,
                    UseDeclarations: $2,
                    SemiColonTkn:    $3,
                }
            }
    |   T_USE use_type use_declarations ';'
            {
                $$ = &ast.StmtUse{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    UseTkn:          $1,
                    Type:            $2,
                    UseDeclarations: $3,
                    SemiColonTkn:    $4,
                }
            }
    |   T_CONST const_list ';'
            {
                $$ = &ast.StmtConstList{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    ConstTkn:     $1,
                    Consts:       $2,
                    SemiColonTkn: $3,
                }
            }
;

use_type:
        T_FUNCTION
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_CONST
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                if len($4) > 0 {
                    $4[len($4)-1].(*ast.StmtUseDeclaration).CommaTkn = $5
                }

                $$ = &ast.StmtGroupUse{
                    Node: ast.Node{
                        Position: position.NewNodeListTokenPosition($1, $6),
                    },
                    Prefix: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($1),
                        },
                        Parts: $1,
                    },
                    NsSeparatorTkn:       $2,
                    OpenCurlyBracketTkn:  $3,
                    UseDeclarations:      $4,
                    CloseCurlyBracketTkn: $6,
                }
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                $5[len($5)-1].(*ast.StmtUseDeclaration).CommaTkn = $6

                $$ = &ast.StmtGroupUse{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $7),
                    },
                    LeadingNsSeparatorTkn: $1,
                    Prefix: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Parts: $2,
                    },
                    NsSeparatorTkn:       $3,
                    OpenCurlyBracketTkn:  $4,
                    UseDeclarations:      $5,
                    CloseCurlyBracketTkn: $7,
                }
            }
;

mixed_group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                $4[len($4)-1].(*ast.StmtUseDeclaration).CommaTkn = $5

                $$ = &ast.StmtGroupUse{
                    Node: ast.Node{
                        Position: position.NewNodeListTokenPosition($1, $6),
                    },
                    Prefix: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($1),
                        },
                        Parts: $1,
                    },
                    NsSeparatorTkn:       $2,
                    OpenCurlyBracketTkn:  $3,
                    UseDeclarations:      $4,
                    CloseCurlyBracketTkn: $6,
                }
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                $5[len($5)-1].(*ast.StmtUseDeclaration).CommaTkn = $6

                $$ = &ast.StmtGroupUse{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $7),
                    },
                    LeadingNsSeparatorTkn: $1,
                    Prefix: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Parts: $2,
                    },
                    NsSeparatorTkn:       $3,
                    OpenCurlyBracketTkn:  $4,
                    UseDeclarations:      $5,
                    CloseCurlyBracketTkn: $7,
                }
            }
;

possible_comma:
        /* empty */
            {
                $$ = nil
            }
    |   ','
            {
                $$ = $1
            }
;

inline_use_declarations:
        inline_use_declarations ',' inline_use_declaration
            {
                $1[len($1)-1].(*ast.StmtUseDeclaration).CommaTkn = $2

                $$ = append($1, $3)
            }
    |   inline_use_declaration
            {
                $$ = []ast.Vertex{$1}
            }
;

unprefixed_use_declarations:
        unprefixed_use_declarations ',' unprefixed_use_declaration
            {
                $1[len($1)-1].(*ast.StmtUseDeclaration).CommaTkn = $2

                $$ = append($1, $3)
            }
    |   unprefixed_use_declaration
            {
                $$ = []ast.Vertex{$1}
            }
;

use_declarations:
        use_declarations ',' use_declaration
            {
                $1[len($1)-1].(*ast.StmtUseDeclaration).CommaTkn = $2

                $$ = append($1, $3)
            }
    |   use_declaration
            {
                $$ = []ast.Vertex{$1}
            }
;

inline_use_declaration:
        unprefixed_use_declaration
            {
                $$ = $1
            }
    |   use_type unprefixed_use_declaration
            {
                decl := $2.(*ast.StmtUseDeclaration)
                decl.Type = $1
                decl.Node.Position = position.NewNodesPosition($1, $2)

                $$ = $2
            }
;

unprefixed_use_declaration:
        namespace_name
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewNodeListPosition($1),
                    },
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($1),
                        },
                        Parts: $1,
                    },
                }
            }
    |   namespace_name T_AS T_STRING
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewNodeListTokenPosition($1, $3),
                    },
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($1),
                        },
                        Parts: $1,
                    },
                    AsTkn: $2,
                    Alias: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        Value: $3.Value,
                    },
                }
            }
;

use_declaration:
        unprefixed_use_declaration
            {
                $$ = $1
            }
    |   T_NS_SEPARATOR unprefixed_use_declaration
            {
                decl := $2.(*ast.StmtUseDeclaration)
                decl.NsSeparatorTkn = $1
                decl.Node.Position = position.NewTokenNodePosition($1, $2)

                $$ = $2
            }
;

const_list:
        const_list ',' const_decl
            {
                lastNode($1).(*ast.StmtConstant).CommaTkn = $2

                $$ = append($1, $3)
            }
    |   const_decl
            {
                $$ = []ast.Vertex{$1}
            }
;

inner_statement_list:
        inner_statement_list inner_statement
            {
                if inlineHtmlNode, ok := $2.(*ast.StmtInlineHtml); ok && len($1) > 0 {
                    prevNode := lastNode($1)
                    yylex.(*Parser).splitSemiColonAndPhpCloseTag(inlineHtmlNode, prevNode)
                }

                if $2 != nil {
                    $$ = append($1, $2)
                }
            }
    |   /* empty */
            {
                $$ = []ast.Vertex{}
            }
;

inner_statement:
        error
            {
                // error
                $$ = nil
            }
    |   statement
            {
                $$ = $1
            }
    |   function_declaration_statement
            {
                $$ = $1
            }
    |   class_declaration_statement
            {
                $$ = $1
            }
    |   trait_declaration_statement
            {
                $$ = $1
            }
    |   interface_declaration_statement
            {
                $$ = $1
            }
    |   T_HALT_COMPILER '(' ')' ';'
            {
                $$ = &ast.StmtHaltCompiler{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    HaltCompilerTkn:     $1,
                    OpenParenthesisTkn:  $2,
                    CloseParenthesisTkn: $3,
                    SemiColonTkn:        $4,
                }
            }

statement:
        '{' inner_statement_list '}'
            {
                $$ = &ast.StmtStmtList{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenCurlyBracket:  $1,
                    Stmts:             $2,
                    CloseCurlyBracket: $3,
                }
            }
    |   if_stmt
            {
                $$ = $1
            }
    |   alt_if_stmt
            {
                $$ = $1
            }
    |   T_WHILE '(' expr ')' while_statement
            {
                $5.(*ast.StmtWhile).WhileTkn = $1
                $5.(*ast.StmtWhile).OpenParenthesisTkn = $2
                $5.(*ast.StmtWhile).Cond = $3
                $5.(*ast.StmtWhile).CloseParenthesisTkn = $4
                $5.(*ast.StmtWhile).Node.Position = position.NewTokenNodePosition($1, $5)

                $$ = $5
            }
    |   T_DO statement T_WHILE '(' expr ')' ';'
            {
                $$ = &ast.StmtDo{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $7),
                    },
                    DoTkn:               $1,
                    Stmt:                $2,
                    WhileTkn:            $3,
                    OpenParenthesisTkn:  $4,
                    Cond:                $5,
                    CloseParenthesisTkn: $6,
                    SemiColonTkn:        $7,
                }
            }
    |   T_FOR '(' for_exprs ';' for_exprs ';' for_exprs ')' for_statement
            {
                switch n := $9.(type) {
                case *ast.StmtFor :
                    n.Init = $3
                    n.Cond = $5
                    n.Loop = $7
                case *ast.StmtAltFor :
                    n.Init = $3
                    n.Cond = $5
                    n.Loop = $7
                }

                $$ = $9

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $9)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.For, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.InitExpr, $4.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.CondExpr, $6.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.IncExpr, $8.SkippedTokens)
            }
    |   T_SWITCH '(' expr ')' switch_case_list
            {
                exprBrackets := &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($2, $4),
                    },
                    OpenBracketTkn:  $2,
                    Child:           $3,
                    CloseBracketTkn: $4,
                }

                switch n := $5.(type) {
                case *ast.StmtSwitch:
                    n.Cond = exprBrackets
                case *ast.StmtAltSwitch:
                    n.Cond = exprBrackets
                default:
                    panic("unexpected node type")
                }

                $$ = $5

                // save position
                exprBrackets.GetNode().Position = position.NewTokensPosition($2, $4)
                $$.GetNode().Position = position.NewTokenNodePosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(exprBrackets, token.Start, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(exprBrackets, token.End, $4.SkippedTokens)
            }
    |   T_BREAK optional_expr ';'
            {
                $$ = &ast.StmtBreak{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.SkippedTokens)
            }
    |   T_CONTINUE optional_expr ';'
            {
                $$ = &ast.StmtContinue{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.SkippedTokens)
            }
    |   T_RETURN optional_expr ';'
            {
                $$ = &ast.StmtReturn{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.SkippedTokens)
            }
    |   T_GLOBAL global_var_list ';'
            {
                $$ = &ast.StmtGlobal{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.VarList, $3.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.SkippedTokens)
            }
    |   T_STATIC static_var_list ';'
            {
                $$ = &ast.StmtStatic{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.VarList, $3.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.SkippedTokens)
            }
    |   T_ECHO echo_expr_list ';'
            {
                $$ = &ast.StmtEcho{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Echo, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.SkippedTokens)
            }
    |   T_INLINE_HTML
            {
                $$ = &ast.StmtInlineHtml{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   expr ';'
            {
                $$ = &ast.StmtExpression{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.SkippedTokens)
            }
    |   T_UNSET '(' unset_variables possible_comma ')' ';'
            {
                $$ = &ast.StmtUnset{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $6)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Unset, $2.SkippedTokens)
                if $4 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.VarList, append($4.SkippedTokens, $5.SkippedTokens...))
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.VarList, $5.SkippedTokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.CloseParenthesisToken, $6.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $6.SkippedTokens)
            }
    |   T_FOREACH '(' expr T_AS foreach_variable ')' foreach_statement
            {
                switch n := $7.(type) {
                case *ast.StmtForeach :
                    n.Expr = $3
                    n.Var = $5
                case *ast.StmtAltForeach :
                    n.Expr = $3
                    n.Var = $5
                }

                $$ = $7

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $7)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Foreach, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $4.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $6.SkippedTokens)
            }
    |   T_FOREACH '(' expr T_AS variable T_DOUBLE_ARROW foreach_variable ')' foreach_statement
            {
                switch n := $9.(type) {
                case *ast.StmtForeach :
                    n.Expr = $3
                    n.Key = $5
                    n.Var = $7
                case *ast.StmtAltForeach :
                    n.Expr = $3
                    n.Key = $5
                    n.Var = $7
                }

                $$ = $9

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $9)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Foreach, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $4.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Key, $6.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $8.SkippedTokens)
            }
    |   T_DECLARE '(' const_list ')' declare_statement
            {
                $$ = $5
                $$.(*ast.StmtDeclare).Consts = $3

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Declare, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ConstList, $4.SkippedTokens)
            }
    |   ';'
            {
                $$ = &ast.StmtNop{ast.Node{}, }

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $1.SkippedTokens)
            }
    |   T_TRY '{' inner_statement_list '}' catch_list finally_statement
            {
                if $6 == nil {
                    $$ = &ast.StmtTry{ast.Node{}, $3, $5, $6}
                    $$.GetNode().Position = position.NewTokenNodeListPosition($1, $5)
                } else {
                    $$ = &ast.StmtTry{ast.Node{}, $3, $5, $6}
                    $$.GetNode().Position = position.NewTokenNodePosition($1, $6)
                }

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Try, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $4.SkippedTokens)
            }
    |   T_THROW expr ';'
            {
                $$ = &ast.StmtThrow{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.SkippedTokens)
            }
    |   T_GOTO T_STRING ';'
            {
                label := &ast.Identifier{ast.Node{}, $2.Value}
                $$ = &ast.StmtGoto{ast.Node{}, label}

                // save position
                label.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating(label, token.Start, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Label, $3.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $3.SkippedTokens)
            }
    |   T_STRING ':'
            {
                label := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.StmtLabel{ast.Node{}, label}

                // save position
                label.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Label, $2.SkippedTokens)
            }

catch_list:
        /* empty */
            {
                $$ = []ast.Vertex{}
            }
    |   catch_list T_CATCH '(' catch_name_list T_VARIABLE ')' '{' inner_statement_list '}'
            {
                identifier := &ast.Identifier{ast.Node{}, $5.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                catch := &ast.StmtCatch{ast.Node{}, $4, variable, $8}
                $$ = append($1, catch)

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($5)
                variable.GetNode().Position = position.NewTokenPosition($5)
                catch.GetNode().Position = position.NewTokensPosition($2, $9)

                // save comments
                yylex.(*Parser).setFreeFloating(catch, token.Start, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(catch, token.Catch, $3.SkippedTokens)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $5.SkippedTokens)
                yylex.(*Parser).setFreeFloating(catch, token.Var, $6.SkippedTokens)
                yylex.(*Parser).setFreeFloating(catch, token.Cond, $7.SkippedTokens)
                yylex.(*Parser).setFreeFloating(catch, token.Stmts, $9.SkippedTokens)
            }
;
catch_name_list:
        name
            {
                $$ = []ast.Vertex{$1}
            }
    |   catch_name_list '|' name
            {
                switch n := lastNode($1).(type) {
                    case *ast.NameName: n.ListSeparatorTkn = $2
                    case *ast.NameFullyQualified: n.ListSeparatorTkn = $2
                    case *ast.NameRelative: n.ListSeparatorTkn = $2
                }
                $$ = append($1, $3)
            }
;

finally_statement:
        /* empty */
            {
                $$ = nil
            }
    |   T_FINALLY '{' inner_statement_list '}'
            {
                $$ = &ast.StmtFinally{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Finally, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $4.SkippedTokens)
            }
;

unset_variables:
        unset_variable
            {
                $$ = []ast.Vertex{$1}
            }
    |   unset_variables ',' unset_variable
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
;

unset_variable:
        variable
            {
                $$ = $1
            }
;

function_declaration_statement:
        T_FUNCTION returns_ref T_STRING backup_doc_comment '(' parameter_list ')' return_type '{' inner_statement_list '}'
            {
                name := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtFunction{ast.Node{}, $2 != nil, name, $6, $8, $10}

                // save position
                name.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewTokensPosition($1, $11)


                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                if $2 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $2.SkippedTokens)
                    yylex.(*Parser).setFreeFloating(name, token.Start, $3.SkippedTokens)
                } else {
                    yylex.(*Parser).setFreeFloating(name, token.Start, $3.SkippedTokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.Name, $5.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ParamList, $7.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ReturnType, $9.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $11.SkippedTokens)

                // normalize
                if $8 == nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Params, $$.GetNode().Tokens[token.ReturnType]); delete($$.GetNode().Tokens, token.ReturnType)
                }
            }
;

is_reference:
        /* empty */
            {
                $$ = nil
            }
    |   '&'
            {
                $$ = $1
            }
;

is_variadic:
        /* empty */
            {
                $$ = nil
            }
    |   T_ELLIPSIS
            {
                $$ = $1
            }
;

class_declaration_statement:
    class_modifiers T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
            {
                name := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtClass{ast.Node{}, name, $1, nil, $4, $5, $8}

                // save position
                name.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewOptionalListTokensPosition($1, $2, $9)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)
                yylex.(*Parser).setFreeFloating($$, token.ModifierList, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(name, token.Start, $3.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $7.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $9.SkippedTokens)
            }
    |   T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
            {
                name := &ast.Identifier{ast.Node{}, $2.Value}
                $$ = &ast.StmtClass{ast.Node{}, name, nil, nil, $3, $4, $7}

                // save position
                name.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $8)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating(name, token.Start, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $6.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $8.SkippedTokens)
            }
;

class_modifiers:
        class_modifier
            {
                $$ = []ast.Vertex{$1}
            }
    |   class_modifiers class_modifier
            {
                $$ = append($1, $2)
            }
;

class_modifier:
        T_ABSTRACT
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_FINAL
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

trait_declaration_statement:
        T_TRAIT T_STRING backup_doc_comment '{' class_statement_list '}'
            {
                name := &ast.Identifier{ast.Node{}, $2.Value}
                $$ = &ast.StmtTrait{ast.Node{}, name, $5}

                // save position
                name.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $6)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating(name, token.Start, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $4.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $6.SkippedTokens)
            }
;

interface_declaration_statement:
        T_INTERFACE T_STRING interface_extends_list backup_doc_comment '{' class_statement_list '}'
            {
                name := &ast.Identifier{ast.Node{}, $2.Value}
                $$ = &ast.StmtInterface{ast.Node{}, name, $3, $6}

                // save position
                name.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $7)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating(name, token.Start, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $5.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $7.SkippedTokens)
            }
;

extends_from:
        /* empty */
            {
                $$ = nil
            }
    |   T_EXTENDS name
            {
                $$ = &ast.StmtClassExtends{ast.Node{}, $2};

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

interface_extends_list:
        /* empty */
            {
                $$ = nil
            }
    |   T_EXTENDS name_list
            {
                $$ = &ast.StmtInterfaceExtends{ast.Node{}, $2};

                // save position
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

implements_list:
        /* empty */
            {
                $$ = nil
            }
    |   T_IMPLEMENTS name_list
            {
                $$ = &ast.StmtClassImplements{ast.Node{}, $2};

                // save position
                $$.GetNode().Position = position.NewTokenNodeListPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

foreach_variable:
        variable
            {
                $$ = $1
            }
    |   '&' variable
            {
                $$ = &ast.ExprReference{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_LIST '(' array_pair_list ')'
            {
                $$ = &ast.ExprList{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.List, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $4.SkippedTokens)
            }
    |   '[' array_pair_list ']'
            {
                $$ = &ast.ExprShortList{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save commentsc
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $3.SkippedTokens)
            }
;

for_statement:
        statement
            {
                $$ = &ast.StmtFor{ast.Node{}, nil, nil, nil, $1}

                // save position
                $$.GetNode().Position = position.NewNodePosition($1)
            }
    |   ':' inner_statement_list T_ENDFOR ';'
            {
                stmtList := &ast.StmtStmtList{
                    Node: ast.Node{
                        Position: position.NewNodeListPosition($2),
                    },
                    Stmts: $2,
                }
                $$ = &ast.StmtAltFor{ast.Node{}, nil, nil, nil, stmtList}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $4.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $4.SkippedTokens)
            }
;

foreach_statement:
        statement
            {
                $$ = &ast.StmtForeach{ast.Node{}, nil, nil, nil, $1}

                // save position
                $$.GetNode().Position = position.NewNodePosition($1)
            }
    |   ':' inner_statement_list T_ENDFOREACH ';'
            {
                stmtList := &ast.StmtStmtList{
                    Node: ast.Node{
                        Position: position.NewNodeListPosition($2),
                    },
                    Stmts: $2,
                }
                $$ = &ast.StmtAltForeach{ast.Node{}, nil, nil, nil, stmtList}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $4.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $4.SkippedTokens)
            }
;

declare_statement:
        statement
            {
                $$ = &ast.StmtDeclare{ast.Node{}, false, nil, $1}

                // save position
                $$.GetNode().Position = position.NewNodePosition($1)
            }
    |   ':' inner_statement_list T_ENDDECLARE ';'
            {
                stmtList := &ast.StmtStmtList{
                    Node: ast.Node{
                        Position: position.NewNodeListPosition($2),
                    },
                    Stmts: $2,
                }
                $$ = &ast.StmtDeclare{ast.Node{}, true, nil, stmtList}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $4.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $4.SkippedTokens)
            }
;

switch_case_list:
        '{' case_list '}'
            {
                caseList := &ast.StmtCaseList{ast.Node{}, $2}
                $$ = &ast.StmtSwitch{ast.Node{}, nil, caseList}

                // save position
                caseList.GetNode().Position = position.NewTokensPosition($1, $3)
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(caseList, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating(caseList, token.CaseListEnd, $3.SkippedTokens)
            }
    |   '{' ';' case_list '}'
            {
                caseList := &ast.StmtCaseList{ast.Node{}, $3}
                $$ = &ast.StmtSwitch{ast.Node{}, nil, caseList}

                // save position
                caseList.GetNode().Position = position.NewTokensPosition($1, $4)
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating(caseList, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(caseList, token.CaseListStart, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(caseList, token.CaseListEnd, $4.SkippedTokens)
            }
    |   ':' case_list T_ENDSWITCH ';'
            {
                caseList := &ast.StmtCaseList{ast.Node{}, $2}
                $$ = &ast.StmtAltSwitch{ast.Node{}, nil, caseList}

                // save position
                caseList.GetNode().Position = position.NewNodeListPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating(caseList, token.CaseListEnd, $3.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $4.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $4.SkippedTokens)
            }
    |   ':' ';' case_list T_ENDSWITCH ';'
            {

                caseList := &ast.StmtCaseList{ast.Node{}, $3}
                $$ = &ast.StmtAltSwitch{ast.Node{}, nil, caseList}

                // save position
                caseList.GetNode().Position = position.NewNodeListPosition($3)
                $$.GetNode().Position = position.NewTokensPosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(caseList, token.CaseListStart, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(caseList, token.CaseListEnd, $4.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $5.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $5.SkippedTokens)
            }
;

case_list:
        /* empty */
            {
                $$ = []ast.Vertex{}
            }
    |   case_list T_CASE expr case_separator inner_statement_list
            {
                _case := &ast.StmtCase{ast.Node{}, $3, $5}
                $$ = append($1, _case)

                // save position
                _case.GetNode().Position = position.NewTokenNodeListPosition($2, $5)

                // save comments
                yylex.(*Parser).setFreeFloating(_case, token.Start, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(_case, token.Expr, append($4.SkippedTokens))
                yylex.(*Parser).setToken(_case, token.CaseSeparator, $4.SkippedTokens)
            }
    |   case_list T_DEFAULT case_separator inner_statement_list
            {
                _default := &ast.StmtDefault{ast.Node{}, $4}
                $$ = append($1, _default)

                // save position
                _default.GetNode().Position = position.NewTokenNodeListPosition($2, $4)

                // save comments
                yylex.(*Parser).setFreeFloating(_default, token.Start, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(_default, token.Default, $3.SkippedTokens)
                yylex.(*Parser).setToken(_default, token.CaseSeparator, $3.SkippedTokens)
            }
;

case_separator:
        ':'
            {
                $$ = $1
            }
    |   ';'
            {
                $$ = $1
            }
;

while_statement:
        statement
            {
                $$ = &ast.StmtWhile{
                    Node: ast.Node{
                        Position: position.NewNodePosition($1),
                    },
                    Stmt: $1,
                }
            }
    |   ':' inner_statement_list T_ENDWHILE ';'
            {
                $$ = &ast.StmtWhile{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    Alt:      true,
                    ColonTkn: $1,
                    Stmt: &ast.StmtStmtList{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2),
                        },
                        Stmts: $2,
                    },
                    EndWhileTkn:  $3,
                    SemiColonTkn: $4,
                }
            }
;

if_stmt_without_else:
        T_IF '(' expr ')' statement
            {
                $$ = &ast.StmtIf{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $5),
                    },
                    IfTkn:               $1,
                    OpenParenthesisTkn:  $2,
                    Cond:                $3,
                    CloseParenthesisTkn: $4,
                    Stmt:                $5,
                }
            }
    |   if_stmt_without_else T_ELSEIF '(' expr ')' statement
            {
                $1.(*ast.StmtIf).ElseIf = append($1.(*ast.StmtIf).ElseIf, &ast.StmtElseIf{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($2, $6),
                    },
                    ElseIfTkn:           $2,
                    OpenParenthesisTkn:  $3,
                    Cond:                $4,
                    CloseParenthesisTkn: $5,
                    Stmt:                $6,
                })

                $1.(*ast.StmtIf).Node.Position = position.NewNodesPosition($1, $6)

                $$ = $1
            }
;

if_stmt:
        if_stmt_without_else %prec T_NOELSE
            {
                $$ = $1
            }
    |   if_stmt_without_else T_ELSE statement
            {
                $1.(*ast.StmtIf).Else = &ast.StmtElse{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($2, $3),
                    },
                    ElseTkn: $2,
                    Stmt:    $3,
                }

                $1.(*ast.StmtIf).Node.Position = position.NewNodesPosition($1, $3)

                $$ = $1
            }
;

alt_if_stmt_without_else:
        T_IF '(' expr ')' ':' inner_statement_list
            {
                $$ = &ast.StmtIf{
                    Node: ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $6),
                    },
                    Alt:                 true,
                    IfTkn:               $1,
                    OpenParenthesisTkn:  $2,
                    Cond:                $3,
                    CloseParenthesisTkn: $4,
                    ColonTkn:            $5,
                    Stmt: &ast.StmtStmtList{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($6),
                        },
                        Stmts: $6,
                    },
                }
            }
    |   alt_if_stmt_without_else T_ELSEIF '(' expr ')' ':' inner_statement_list
            {
                $1.(*ast.StmtIf).ElseIf = append($1.(*ast.StmtIf).ElseIf, &ast.StmtElseIf{
                    Node: ast.Node{
                        Position: position.NewTokenNodeListPosition($2, $7),
                    },
                    Alt:                 true,
                    ElseIfTkn:           $2,
                    OpenParenthesisTkn:  $3,
                    Cond:                $4,
                    CloseParenthesisTkn: $5,
                    ColonTkn:            $6,
                    Stmt: &ast.StmtStmtList{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($7),
                        },
                        Stmts: $7,
                    },
                })

                $$ = $1
            }
;

alt_if_stmt:
        alt_if_stmt_without_else T_ENDIF ';'
            {
                $1.(*ast.StmtIf).EndIfTkn = $2
                $1.(*ast.StmtIf).SemiColonTkn = $3
                $1.(*ast.StmtIf).Node.Position = position.NewNodeTokenPosition($1, $3)

                $$ = $1
            }
    |   alt_if_stmt_without_else T_ELSE ':' inner_statement_list T_ENDIF ';'
            {
                $1.(*ast.StmtIf).Else = &ast.StmtElse{
                    Node: ast.Node{
                        Position: position.NewTokenNodeListPosition($2, $4),
                    },
                    Alt:      true,
                    ElseTkn:  $2,
                    ColonTkn: $3,
                    Stmt: &ast.StmtStmtList{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($4),
                        },
                        Stmts: $4,
                    },
                }
                $1.(*ast.StmtIf).EndIfTkn = $5
                $1.(*ast.StmtIf).SemiColonTkn = $6
                $1.(*ast.StmtIf).Node.Position = position.NewNodeTokenPosition($1, $6)

                $$ = $1
            }
;

parameter_list:
        non_empty_parameter_list
            {
                $$ = $1
            }
    |   /* empty */
            {
                $$ = nil
            }
;

non_empty_parameter_list:
        parameter
            {
                $$ = []ast.Vertex{$1}
            }
    |   non_empty_parameter_list ',' parameter
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
;

parameter:
        optional_type is_reference is_variadic T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $4.Value}
                identifier.GetNode().Position = position.NewTokenPosition($4)

                var variable ast.Vertex
                variable = &ast.ExprVariable{ast.Node{}, identifier}
                variable.GetNode().Position = position.NewTokenPosition($4)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $4.SkippedTokens)

                if $3 != nil {
                    variable = &ast.Variadic{ast.Node{}, variable}
                    variable.GetNode().Position = position.NewTokensPosition($3, $4)
                    yylex.(*Parser).setFreeFloating(variable, token.Start, $3.SkippedTokens)
                }

                if $2 != nil {
                    variable = &ast.Reference{ast.Node{}, variable}
                    variable.GetNode().Position = position.NewTokensPosition($2, $4)
                    yylex.(*Parser).setFreeFloating(variable, token.Start, $2.SkippedTokens)
                }

                $$ = &ast.Parameter{ast.Node{}, $1, variable, nil}

                if $1 != nil {
                    $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)
                } else if $2 != nil {
                    $$.GetNode().Position = position.NewTokensPosition($2, $4)
                } else if $3 != nil {
                    $$.GetNode().Position = position.NewTokensPosition($3, $4)
                } else {
                    $$.GetNode().Position = position.NewTokenPosition($4)
                }
            }
    |   optional_type is_reference is_variadic T_VARIABLE '=' expr
            {
                identifier := &ast.Identifier{ast.Node{}, $4.Value}
                identifier.GetNode().Position = position.NewTokenPosition($4)

                var variable ast.Vertex
                variable = &ast.ExprVariable{ast.Node{}, identifier}
                variable.GetNode().Position = position.NewTokenPosition($4)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $4.SkippedTokens)
                yylex.(*Parser).setFreeFloating(variable, token.End, $5.SkippedTokens)

                if $3 != nil {
                    variable = &ast.Variadic{ast.Node{}, variable}
                    variable.GetNode().Position = position.NewTokensPosition($3, $4)
                    yylex.(*Parser).setFreeFloating(variable, token.Start, $3.SkippedTokens)
                }

                if $2 != nil {
                    variable = &ast.Reference{ast.Node{}, variable}
                    variable.GetNode().Position = position.NewTokensPosition($2, $4)
                    yylex.(*Parser).setFreeFloating(variable, token.Start, $2.SkippedTokens)
                }

                $$ = &ast.Parameter{ast.Node{}, $1, variable, $6}

                if $1 != nil {
                    $$.GetNode().Position = position.NewNodesPosition($1, $6)
                } else if $2 != nil {
                    $$.GetNode().Position = position.NewTokenNodePosition($2, $6)
                } else if $3 != nil {
                    $$.GetNode().Position = position.NewTokenNodePosition($3, $6)
                } else {
                    $$.GetNode().Position = position.NewTokenNodePosition($4, $6)
                }
            }
;

optional_type:
        /* empty */
            {
                $$ = nil
            }
    |   type_expr
            {
                $$ = $1
            }
;

type_expr:
        type
            {
                $$ = $1
            }
    |   '?' type
            {
                $$ = &ast.Nullable{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

type:
        T_ARRAY
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_CALLABLE
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   name
            {
                $$ = $1
            }
;

return_type:
        /* empty */
            {
                $$ = nil
            }
    |   ':' type_expr
            {
                $$ = $2;

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, append($1.SkippedTokens, $$.GetNode().Tokens[token.Start]...))
            }
;

argument_list:
        '(' ')'
            {
                $$ = &ast.ArgumentList{ast.Node{}, nil}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $2.SkippedTokens)
            }
    |   '(' non_empty_argument_list possible_comma ')'
            {
                $$ = &ast.ArgumentList{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, $1.SkippedTokens)
                if $3 != nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.End, append($3.SkippedTokens, $4.SkippedTokens...))
                } else {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.End, $4.SkippedTokens)
                }
            }
;

non_empty_argument_list:
        argument
            {
                $$ = []ast.Vertex{$1}
            }
    |   non_empty_argument_list ',' argument
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
;

argument:
        expr
            {
                $$ = &ast.Argument{ast.Node{}, false, false, $1}

                // save position
                $$.GetNode().Position = position.NewNodePosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
            }
    |   T_ELLIPSIS expr
            {
                $$ = &ast.Argument{ast.Node{}, true, false, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

global_var_list:
        global_var_list ',' global_var
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
    |   global_var
            {
                $$ = []ast.Vertex{$1}
            }
;

global_var:
        simple_variable
            {
                $$ = $1
            }
;

static_var_list:
        static_var_list ',' static_var
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
    |   static_var
            {
                $$ = []ast.Vertex{$1}
            }
;

static_var:
        T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.StmtStaticVar{ast.Node{}, variable, nil}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                variable.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_VARIABLE '=' expr
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.StmtStaticVar{ast.Node{}, variable, $3}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                variable.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
;

class_statement_list:
        class_statement_list class_statement
            {
                $$ = append($1, $2)
            }
    |   /* empty */
            {
                $$ = []ast.Vertex{}
            }
;

class_statement:
        variable_modifiers optional_type property_list ';'
            {
                $$ = &ast.StmtPropertyList{ast.Node{}, $1, $2, $3}

                // save position
                $$.GetNode().Position = position.NewNodeListTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)
                yylex.(*Parser).setFreeFloating($$, token.PropertyList, $4.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $4.SkippedTokens)
            }
    |   method_modifiers T_CONST class_const_list ';'
            {
                $$ = &ast.StmtClassConstList{
                    Node: ast.Node{
                        Position: position.NewOptionalListTokensPosition($1, $2, $4),
                    },
                    Modifiers:    $1,
                    ConstTkn:     $2,
                    Consts:       $3,
                    SemiColonTkn: $4,
                }
            }
    |   T_USE name_list trait_adaptations
            {
                $$ = &ast.StmtTraitUse{ast.Node{}, $2, $3}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   method_modifiers T_FUNCTION returns_ref identifier backup_doc_comment '(' parameter_list ')' return_type method_body
            {
                name := &ast.Identifier{ast.Node{}, $4.Value}
                $$ = &ast.StmtClassMethod{ast.Node{}, $3 != nil, name, $1, $7, $9, $10}

                // save position
                name.GetNode().Position = position.NewTokenPosition($4)
                if $1 == nil {
                    $$.GetNode().Position = position.NewTokenNodePosition($2, $10)
                } else {
                    $$.GetNode().Position = position.NewNodeListNodePosition($1, $10)
                }

                // save comments
                if len($1) > 0 {
                    yylex.(*Parser).MoveFreeFloating($1[0], $$)
                    yylex.(*Parser).setFreeFloating($$, token.ModifierList, $2.SkippedTokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Start, $2.SkippedTokens)
                }
                if $3 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $4.SkippedTokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $3.SkippedTokens)
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $4.SkippedTokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.Name, $6.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ParameterList, $8.SkippedTokens)
            }
;

name_list:
        name
            {
                $$ = []ast.Vertex{$1}
            }
    |   name_list ',' name
            {
                switch n := lastNode($1).(type) {
                    case *ast.NameName: n.ListSeparatorTkn = $2
                    case *ast.NameFullyQualified: n.ListSeparatorTkn = $2
                    case *ast.NameRelative: n.ListSeparatorTkn = $2
                }
                $$ = append($1, $3)
            }
;

trait_adaptations:
        ';'
            {
                $$ = &ast.StmtNop{ast.Node{}, }

                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $1.SkippedTokens)
            }
    |   '{' '}'
            {
                $$ = &ast.StmtTraitAdaptationList{ast.Node{}, nil}

                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.AdaptationList, $2.SkippedTokens)
            }
    |   '{' trait_adaptation_list '}'
            {
                $$ = &ast.StmtTraitAdaptationList{ast.Node{}, $2}

                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.AdaptationList, $3.SkippedTokens)
            }
;

trait_adaptation_list:
        trait_adaptation
            {
                $$ = []ast.Vertex{$1}
            }
    |   trait_adaptation_list trait_adaptation
            {
                $$ = append($1, $2)
            }
;

trait_adaptation:
        trait_precedence ';'
            {
                $$ = $1;

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.NameList, $2.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.SkippedTokens)
            }
    |   trait_alias ';'
            {
                $$ = $1;

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Alias, $2.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $2.SkippedTokens)
            }
;

trait_precedence:
        absolute_trait_method_reference T_INSTEADOF name_list
            {
                $$ = &ast.StmtTraitUsePrecedence{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeNodeListPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.SkippedTokens)
            }
;

trait_alias:
        trait_method_reference T_AS T_STRING
            {
                alias := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtTraitUseAlias{ast.Node{}, $1, nil, alias}

                // save position
                alias.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(alias, token.Start, $3.SkippedTokens)
            }
    |   trait_method_reference T_AS reserved_non_modifiers
            {
                alias := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtTraitUseAlias{ast.Node{}, $1, nil, alias}

                // save position
                alias.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(alias, token.Start, $3.SkippedTokens)
            }
    |   trait_method_reference T_AS member_modifier identifier
            {
                alias := &ast.Identifier{ast.Node{}, $4.Value}
                $$ = &ast.StmtTraitUseAlias{ast.Node{}, $1, $3, alias}

                // save position
                alias.GetNode().Position = position.NewTokenPosition($4)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(alias, token.Start, $4.SkippedTokens)
            }
    |   trait_method_reference T_AS member_modifier
            {
                $$ = &ast.StmtTraitUseAlias{ast.Node{}, $1, $3, nil}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.SkippedTokens)
            }
;

trait_method_reference:
        identifier
            {
                name := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.StmtTraitMethodRef{ast.Node{}, nil, name}

                // save position
                name.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   absolute_trait_method_reference
            {
                $$ = $1
            }
;

absolute_trait_method_reference:
        name T_PAAMAYIM_NEKUDOTAYIM identifier
            {
                target := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtTraitMethodRef{ast.Node{}, $1, target}

                // save position
                target.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(target, token.Start, $2.SkippedTokens)
            }
;

method_body:
        ';' /* abstract method */
            {
                $$ = &ast.StmtNop{ast.Node{}, }

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.SemiColon, $1.SkippedTokens)
            }
    |   '{' inner_statement_list '}'
            {
                $$ = &ast.StmtStmtList{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenCurlyBracket:  $1,
                    Stmts:             $2,
                    CloseCurlyBracket: $3,
                }
            }
;

variable_modifiers:
        non_empty_member_modifiers
            {
                $$ = $1
            }
    |   T_VAR
            {
                modifier := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = []ast.Vertex{modifier}

                // save position
                modifier.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(modifier, token.Start, $1.SkippedTokens)
            }
;

method_modifiers:
        /* empty */
            {
                $$ = nil
            }
    |   non_empty_member_modifiers
            {
                $$ = $1
            }
;

non_empty_member_modifiers:
        member_modifier
            {
                $$ = []ast.Vertex{$1}
            }
    |   non_empty_member_modifiers member_modifier
            {
                $$ = append($1, $2)
            }
;

member_modifier:
        T_PUBLIC
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_PROTECTED
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_PRIVATE
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_STATIC
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_ABSTRACT
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_FINAL
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

property_list:
        property_list ',' property
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
    |   property
            {
                $$ = []ast.Vertex{$1}
            }
;

property:
        T_VARIABLE backup_doc_comment
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.StmtProperty{ast.Node{}, variable, nil}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                variable.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_VARIABLE '=' expr backup_doc_comment
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.StmtProperty{ast.Node{}, variable, $3}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                variable.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
;

class_const_list:
        class_const_list ',' class_const_decl
            {
                lastNode($1).(*ast.StmtConstant).CommaTkn = $2

                $$ = append($1, $3)
            }
    |   class_const_decl
            {
                $$ = []ast.Vertex{$1}
            }
;

class_const_decl:
        identifier '=' expr backup_doc_comment
            {
                $$ = &ast.StmtConstant{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $3),
                    },
                    Name: &ast.Identifier{
                        Node:  ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        Value: $1.Value,
                    },
                    EqualTkn: $2,
                    Expr:     $3,
                }

                yylex.(*Parser).setFreeFloating($$.(*ast.StmtConstant).Name, token.Start, $1.SkippedTokens)
            }
;

const_decl:
        T_STRING '=' expr backup_doc_comment
            {
                $$ = &ast.StmtConstant{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $3),
                    },
                    Name: &ast.Identifier{
                        Node:  ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        Value: $1.Value,
                    },
                    EqualTkn: $2,
                    Expr:     $3,
                }

                yylex.(*Parser).setFreeFloating($$.(*ast.StmtConstant).Name, token.Start, $1.SkippedTokens)
            }
;

echo_expr_list:
        echo_expr_list ',' echo_expr
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
    |   echo_expr
            {
                $$ = []ast.Vertex{$1}
            }
;

echo_expr:
        expr
            {
                $$ = $1
            }
;

for_exprs:
        /* empty */
            {
                $$ = nil;
            }
    |   non_empty_for_exprs
            {
                $$ = $1
            }
;

non_empty_for_exprs:
        non_empty_for_exprs ',' expr
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
    |   expr
            {
                $$ = []ast.Vertex{$1}
            }
;

anonymous_class:
        T_CLASS ctor_arguments extends_from implements_list backup_doc_comment '{' class_statement_list '}'
            {
                if $2 != nil {
                    $$ = &ast.StmtClass{ast.Node{}, nil, nil, $2.(*ast.ArgumentList), $3, $4, $7}
                } else {
                    $$ = &ast.StmtClass{ast.Node{}, nil, nil, nil, $3, $4, $7}
                }

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $8)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $6.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $8.SkippedTokens)
            }
;

new_expr:
        T_NEW class_name_reference ctor_arguments
            {
                if $3 != nil {
                    $$ = &ast.ExprNew{ast.Node{}, $2, $3.(*ast.ArgumentList)}
                    $$.GetNode().Position = position.NewTokenNodePosition($1, $3)
                } else {
                    $$ = &ast.ExprNew{ast.Node{}, $2, nil}
                    $$.GetNode().Position = position.NewTokenNodePosition($1, $2)
                }

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_NEW anonymous_class
            {
                $$ = &ast.ExprNew{ast.Node{}, $2, nil}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

expr_without_variable:
        T_LIST '(' array_pair_list ')' '=' expr
            {
                listNode := &ast.ExprList{ast.Node{}, $3}
                $$ = &ast.ExprAssign{ast.Node{}, listNode, $6}

                // save position
                listNode.GetNode().Position = position.NewTokensPosition($1, $4)
                $$.GetNode().Position = position.NewTokenNodePosition($1, $6)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating(listNode, token.List, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(listNode, token.ArrayPairList, $4.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $5.SkippedTokens)
            }
    |   '[' array_pair_list ']' '=' expr
            {
                shortList := &ast.ExprShortList{ast.Node{}, $2}
                $$ = &ast.ExprAssign{ast.Node{}, shortList, $5}

                // save position
                shortList.GetNode().Position = position.NewTokensPosition($1, $3)
                $$.GetNode().Position = position.NewTokenNodePosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating(shortList, token.ArrayPairList, $3.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $4.SkippedTokens)
            }
    |   variable '=' expr
            {
                $$ = &ast.ExprAssign{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable '=' '&' expr
            {
                $$ = &ast.ExprAssignReference{ast.Node{}, $1, $4}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Equal, $3.SkippedTokens)
            }
    |   T_CLONE expr
            {
                $$ = &ast.ExprClone{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   variable T_PLUS_EQUAL expr
            {
                $$ = &ast.ExprAssignPlus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_MINUS_EQUAL expr
            {
                $$ = &ast.ExprAssignMinus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_MUL_EQUAL expr
            {
                $$ = &ast.ExprAssignMul{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_POW_EQUAL expr
            {
                $$ = &ast.ExprAssignPow{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_DIV_EQUAL expr
            {
                $$ = &ast.ExprAssignDiv{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_CONCAT_EQUAL expr
            {
                $$ = &ast.ExprAssignConcat{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_MOD_EQUAL expr
            {
                $$ = &ast.ExprAssignMod{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_AND_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_OR_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_XOR_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_SL_EQUAL expr
            {
                $$ = &ast.ExprAssignShiftLeft{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_SR_EQUAL expr
            {
                $$ = &ast.ExprAssignShiftRight{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_COALESCE_EQUAL expr
            {
                $$ = &ast.ExprAssignCoalesce{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   variable T_INC
            {
                $$ = &ast.ExprPostInc{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   T_INC variable
            {
                $$ = &ast.ExprPreInc{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   variable T_DEC
            {
                $$ = &ast.ExprPostDec{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   T_DEC variable
            {
                $$ = &ast.ExprPreDec{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   expr T_BOOLEAN_OR expr
            {
                $$ = &ast.ExprBinaryBooleanOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_BOOLEAN_AND expr
            {
                $$ = &ast.ExprBinaryBooleanAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_LOGICAL_OR expr
            {
                $$ = &ast.ExprBinaryLogicalOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_LOGICAL_AND expr
            {
                $$ = &ast.ExprBinaryLogicalAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_LOGICAL_XOR expr
            {
                $$ = &ast.ExprBinaryLogicalXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '|' expr
            {
                $$ = &ast.ExprBinaryBitwiseOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '&' expr
            {
                $$ = &ast.ExprBinaryBitwiseAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '^' expr
            {
                $$ = &ast.ExprBinaryBitwiseXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '.' expr
            {
                $$ = &ast.ExprBinaryConcat{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '+' expr
            {
                $$ = &ast.ExprBinaryPlus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '-' expr
            {
                $$ = &ast.ExprBinaryMinus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '*' expr
            {
                $$ = &ast.ExprBinaryMul{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_POW expr
            {
                $$ = &ast.ExprBinaryPow{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '/' expr
            {
                $$ = &ast.ExprBinaryDiv{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '%' expr
            {
                $$ = &ast.ExprBinaryMod{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_SL expr
            {
                $$ = &ast.ExprBinaryShiftLeft{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_SR expr
            {
                $$ = &ast.ExprBinaryShiftRight{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   '+' expr %prec T_INC
            {
                $$ = &ast.ExprUnaryPlus{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '-' expr %prec T_INC
            {
                $$ = &ast.ExprUnaryMinus{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '!' expr
            {
                $$ = &ast.ExprBooleanNot{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '~' expr
            {
                $$ = &ast.ExprBitwiseNot{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   expr T_IS_IDENTICAL expr
            {
                $$ = &ast.ExprBinaryIdentical{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_IS_NOT_IDENTICAL expr
            {
                $$ = &ast.ExprBinaryNotIdentical{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_IS_EQUAL expr
            {
                $$ = &ast.ExprBinaryEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_IS_NOT_EQUAL expr
            {
                $$ = &ast.ExprBinaryNotEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Equal, $2.SkippedTokens)
            }
    |   expr '<' expr
            {
                $$ = &ast.ExprBinarySmaller{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_IS_SMALLER_OR_EQUAL expr
            {
                $$ = &ast.ExprBinarySmallerOrEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr '>' expr
            {
                $$ = &ast.ExprBinaryGreater{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_IS_GREATER_OR_EQUAL expr
            {
                $$ = &ast.ExprBinaryGreaterOrEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_SPACESHIP expr
            {
                $$ = &ast.ExprBinarySpaceship{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr T_INSTANCEOF class_name_reference
            {
                $$ = &ast.ExprInstanceOf{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   '(' expr ')'
            {
                $$ = &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenBracketTkn:  $1,
                    Child:           $2,
                    CloseBracketTkn: $3,
                }

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.SkippedTokens)
            }
    |   new_expr
            {
                $$ = $1
            }
    |   expr '?' expr ':' expr
            {
                $$ = &ast.ExprTernary{ast.Node{}, $1, $3, $5}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $5)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Cond, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.True, $4.SkippedTokens)
            }
    |   expr '?' ':' expr
            {
                $$ = &ast.ExprTernary{ast.Node{}, $1, nil, $4}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Cond, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.True, $3.SkippedTokens)
            }
    |   expr T_COALESCE expr
            {
                $$ = &ast.ExprBinaryCoalesce{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   internal_functions_in_yacc
            {
                $$ = $1
            }
    |   T_INT_CAST expr
            {
                $$ = &ast.ExprCastInt{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.SkippedTokens)
            }
    |   T_DOUBLE_CAST expr
            {
                $$ = &ast.ExprCastDouble{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.SkippedTokens)
            }
    |   T_STRING_CAST expr
            {
                $$ = &ast.ExprCastString{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.SkippedTokens)
            }
    |   T_ARRAY_CAST expr
            {
                $$ = &ast.ExprCastArray{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.SkippedTokens)
            }
    |   T_OBJECT_CAST expr
            {
                $$ = &ast.ExprCastObject{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.SkippedTokens)
            }
    |   T_BOOL_CAST expr
            {
                $$ = &ast.ExprCastBool{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.SkippedTokens)
            }
    |   T_UNSET_CAST expr
            {
                $$ = &ast.ExprCastUnset{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setToken($$, token.Cast, $1.SkippedTokens)
            }
    |   T_EXIT exit_expr
            {
                $$ = &ast.ExprExit{ast.Node{}, false, $2}

                if (bytes.EqualFold($1.Value, []byte("die"))) {
                    $$.(*ast.ExprExit).Die = true
                }

                // save position
                if $2 == nil {
                    $$.GetNode().Position = position.NewTokenPosition($1)
                } else {
                    $$.GetNode().Position = position.NewTokenNodePosition($1, $2)
                }

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '@' expr
            {
                $$ = &ast.ExprErrorSuppress{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   scalar
            {
                $$ = $1
            }
    |   '`' backticks_expr '`'
            {
                $$ = &ast.ExprShellExec{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_PRINT expr
            {
                $$ = &ast.ExprPrint{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_YIELD
            {
                $$ = &ast.ExprYield{ast.Node{}, nil, nil}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_YIELD expr
            {
                $$ = &ast.ExprYield{ast.Node{}, nil, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_YIELD expr T_DOUBLE_ARROW expr
            {
                $$ = &ast.ExprYield{ast.Node{}, $2, $4}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.SkippedTokens)
            }
    |   T_YIELD_FROM expr
            {
                $$ = &ast.ExprYieldFrom{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   inline_function
            {
                $$ = $1;
            }
    |   T_STATIC inline_function
            {
                $$ = $2;

                switch n := $$.(type) {
                case *ast.ExprClosure :
                    n.Static = true;
                case *ast.ExprArrowFunction :
                    n.Static = true;
                };

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Static, $$.GetNode().Tokens[token.Start]); delete($$.GetNode().Tokens, token.Start)
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens);
            }
;

inline_function:
        T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                $$ = &ast.ExprClosure{ast.Node{}, $2 != nil, false, $5, $7, $8, $10}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $11)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                if $2 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $4.SkippedTokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $2.SkippedTokens)
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $4.SkippedTokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.ParameterList, $6.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ReturnType, $9.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $11.SkippedTokens)

                // normalize
                if $8 == nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.LexicalVars, $$.GetNode().Tokens[token.ReturnType]); delete($$.GetNode().Tokens, token.ReturnType)
                }
                if $7 == nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Params, $$.GetNode().Tokens[token.LexicalVarList]); delete($$.GetNode().Tokens, token.LexicalVarList)
                }
            }
    |   T_FN returns_ref '(' parameter_list ')' return_type backup_doc_comment T_DOUBLE_ARROW expr
            {
                $$ = &ast.ExprArrowFunction{ast.Node{}, $2 != nil, false, $4, $6, $9}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $9)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                if $2 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $3.SkippedTokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $2.SkippedTokens)
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $3.SkippedTokens)
                };
                yylex.(*Parser).setFreeFloating($$, token.ParameterList, $5.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ReturnType, $8.SkippedTokens)

                // normalize
                if $6 == nil {
                    yylex.(*Parser).setFreeFloatingTokens($$, token.Params, $$.GetNode().Tokens[token.ReturnType]); delete($$.GetNode().Tokens, token.ReturnType)
                };
            }
;

backup_doc_comment:
        /* empty */
;

returns_ref:
        /* empty */
            {
                $$ = nil
            }
    |   '&'
            {
                $$ = $1
            }
;

lexical_vars:
        /* empty */
            {
                $$ = nil
            }
    |   T_USE '(' lexical_var_list ')'
            {
                $$ = &ast.ExprClosureUse{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Use, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.LexicalVarList, $4.SkippedTokens)
            }
;

lexical_var_list:
        lexical_var_list ',' lexical_var
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
    |   lexical_var
            {
                $$ = []ast.Vertex{$1}
            }
;

lexical_var:
    T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.ExprVariable{ast.Node{}, identifier}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '&' T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $2.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.ExprReference{ast.Node{}, variable}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($2)
                variable.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $2.SkippedTokens)
            }
;

function_call:
        name argument_list
            {
                $$ = &ast.ExprFunctionCall{ast.Node{}, $1, $2.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
            {
                $$ = &ast.ExprStaticCall{ast.Node{}, $1, $3, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
            {
                $$ = &ast.ExprStaticCall{ast.Node{}, $1, $3, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
    |   callable_expr argument_list
            {
                $$ = &ast.ExprFunctionCall{ast.Node{}, $1, $2.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
            }
;

class_name:
        T_STATIC
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   name
            {
                $$ = $1
            }
;

class_name_reference:
        class_name
            {
                $$ = $1
            }
    |   new_variable
            {
                $$ = $1
            }
;

exit_expr:
        /* empty */
            {
                $$ = nil
            }
    |   '(' optional_expr ')'
            {
                $$ = &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenBracketTkn:  $1,
                    Child:           $2,
                    CloseBracketTkn: $3,
                }

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.SkippedTokens)
            }
;

backticks_expr:
        /* empty */
            {
                $$ = []ast.Vertex{}
            }
    |   T_ENCAPSED_AND_WHITESPACE
            {
                part := &ast.ScalarEncapsedStringPart{ast.Node{}, $1.Value}
                $$ = []ast.Vertex{part}

                // save position
                part.GetNode().Position = position.NewTokenPosition($1)
            }
    |   encaps_list
            {
                $$ = $1
            }
;

ctor_arguments:
        /* empty */
            {
                $$ = nil
            }
    |   argument_list
            {
                $$ = $1
            }
;

dereferencable_scalar:
    T_ARRAY '(' array_pair_list ')'
            {
                $$ = &ast.ExprArray{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Array, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $4.SkippedTokens)
            }
    |   '[' array_pair_list ']'
            {
                $$ = &ast.ExprShortArray{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $3.SkippedTokens)
            }
    |   T_CONSTANT_ENCAPSED_STRING
            {
                $$ = &ast.ScalarString{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

scalar:
        T_LNUMBER
            {
                $$ = &ast.ScalarLnumber{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_DNUMBER
            {
                $$ = &ast.ScalarDnumber{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_LINE
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_FILE
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_DIR
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_TRAIT_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_METHOD_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_FUNC_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_NS_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_CLASS_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC
            {
                encapsed := &ast.ScalarEncapsedStringPart{ast.Node{}, $2.Value}
                $$ = &ast.ScalarHeredoc{ast.Node{}, $1.Value, []ast.Vertex{encapsed}}

                // save position
                encapsed.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_START_HEREDOC T_END_HEREDOC
            {
                $$ = &ast.ScalarHeredoc{ast.Node{}, $1.Value, nil}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '"' encaps_list '"'
            {
                $$ = &ast.ScalarEncapsed{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_START_HEREDOC encaps_list T_END_HEREDOC
            {
                $$ = &ast.ScalarHeredoc{ast.Node{}, $1.Value, $2}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   dereferencable_scalar
            {
                $$ = $1
            }
    |   constant
            {
                $$ = $1
            }
;

constant:
        name
            {
                $$ = &ast.ExprConstFetch{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = position.NewNodePosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM identifier
            {
                target := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.ExprClassConstFetch{ast.Node{}, $1, target}

                // save position
                target.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(target, token.Start, $3.SkippedTokens)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM identifier
            {
                target := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.ExprClassConstFetch{ast.Node{}, $1, target}

                // save position
                target.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(target, token.Start, $3.SkippedTokens)
            }
;

expr:
        variable
            {
                $$ = $1
            }
    |   expr_without_variable
            {
                $$ = $1
            }
;

optional_expr:
        /* empty */
            {
                $$ = nil
            }
    |   expr
            {
                $$ = $1
            }
;

variable_class_name:
        dereferencable
            {
                $$ = $1
            }
;

dereferencable:
        variable
            {
                $$ = $1
            }
    |   '(' expr ')'
            {
                $$ = &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenBracketTkn:  $1,
                    Child:           $2,
                    CloseBracketTkn: $3,
                }

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.SkippedTokens)
            }
    |   dereferencable_scalar
            {
                $$ = $1;
            }
;

callable_expr:
        callable_variable
            {
                $$ = $1
            }
    |   '(' expr ')'
            {
                $$ = &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenBracketTkn:  $1,
                    Child:           $2,
                    CloseBracketTkn: $3,
                }

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.SkippedTokens)
            }
    |   dereferencable_scalar
            {
                $$ = $1
            }
;

callable_variable:
        simple_variable
            {
                $$ = $1
            }
    |   dereferencable '[' optional_expr ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
    |   constant '[' optional_expr ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
    |   dereferencable '{' expr '}'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
    |   dereferencable T_OBJECT_OPERATOR property_name argument_list
            {
                $$ = &ast.ExprMethodCall{ast.Node{}, $1, $3, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   function_call
            {
                $$ = $1
            }
;

variable:
        callable_variable
            {
                $$ = $1
            }
    |   static_member
            {
                $$ = $1
            }
    |   dereferencable T_OBJECT_OPERATOR property_name
            {
                $$ = &ast.ExprPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
;

simple_variable:
        T_VARIABLE
            {
                name := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.ExprVariable{ast.Node{}, name}

                // save position
                name.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '$' '{' expr '}'
            {
                $$ = &ast.ExprVariable{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($3, token.Start, append($2.SkippedTokens, $3.GetNode().Tokens[token.Start]...))
                yylex.(*Parser).setFreeFloatingTokens($3, token.End, append($3.GetNode().Tokens[token.End], $4.SkippedTokens...))
            }
    |   '$' simple_variable
            {
                $$ = &ast.ExprVariable{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

static_member:
        class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = &ast.ExprStaticPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = &ast.ExprStaticPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.SkippedTokens)
            }
;

new_variable:
        simple_variable
            {
                $$ = $1
            }
    |   new_variable '[' optional_expr ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
    |   new_variable '{' expr '}'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
    |   new_variable T_OBJECT_OPERATOR property_name
            {
                $$ = &ast.ExprPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = &ast.ExprStaticPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
    |   new_variable T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = &ast.ExprStaticPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
            }
;

member_name:
        identifier
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '{' expr '}'
            {
                $$ = $2;

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, append($1.SkippedTokens, $$.GetNode().Tokens[token.Start]...))
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, append($$.GetNode().Tokens[token.End], $3.SkippedTokens...))
            }
    |   simple_variable
            {
                $$ = $1
            }
;

property_name:
        T_STRING
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '{' expr '}'
            {
                $$ = $2;

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Start, append($1.SkippedTokens, $$.GetNode().Tokens[token.Start]...))
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, append($$.GetNode().Tokens[token.End], $3.SkippedTokens...))
            }
    |   simple_variable
            {
                $$ = $1
            }
;

array_pair_list:
        non_empty_array_pair_list
            {
                $$ = $1
            }
;

possible_array_pair:
        /* empty */
            {
                $$ = &ast.ExprArrayItem{ast.Node{}, false, nil, nil}
            }
    |   array_pair
            {
                $$ = $1
            }
;

non_empty_array_pair_list:
        non_empty_array_pair_list ',' possible_array_pair
            {
                if len($1) == 0 {
                    $1 = []ast.Vertex{&ast.ExprArrayItem{ast.Node{}, false, nil, nil}}
                }

                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
    |   possible_array_pair
            {
                if $1.(*ast.ExprArrayItem).Key == nil && $1.(*ast.ExprArrayItem).Val == nil {
                    $$ = []ast.Vertex{}
                } else {
                    $$ = []ast.Vertex{$1}
                }
            }
;

array_pair:
        expr T_DOUBLE_ARROW expr
            {
                $$ = &ast.ExprArrayItem{ast.Node{}, false, $1, $3}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
            }
    |   expr
            {
                $$ = &ast.ExprArrayItem{ast.Node{}, false, nil, $1}

                // save position
                $$.GetNode().Position = position.NewNodePosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
            }
    |   expr T_DOUBLE_ARROW '&' variable
            {
                reference := &ast.ExprReference{ast.Node{}, $4}
                $$ = &ast.ExprArrayItem{ast.Node{}, false, $1, reference}

                // save position
                $$.GetNode().Position = position.NewNodesPosition($1, $4)
                reference.GetNode().Position = position.NewTokenNodePosition($3, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(reference, token.Start, $3.SkippedTokens)
            }
    |   '&' variable
            {
                reference := &ast.ExprReference{ast.Node{}, $2}
                $$ = &ast.ExprArrayItem{ast.Node{}, false, nil, reference}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)
                reference.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_ELLIPSIS expr
            {
                $$ = &ast.ExprArrayItem{ast.Node{}, true, nil, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   expr T_DOUBLE_ARROW T_LIST '(' array_pair_list ')'
            {
                // TODO: Cannot use list() as standalone expression
                listNode := &ast.ExprList{ast.Node{}, $5}
                $$ = &ast.ExprArrayItem{ast.Node{}, false, $1, listNode}

                // save position
                listNode.GetNode().Position = position.NewTokensPosition($3, $6)
                $$.GetNode().Position = position.NewNodeTokenPosition($1, $6)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(listNode, token.Start, $3.SkippedTokens)
                yylex.(*Parser).setFreeFloating(listNode, token.List, $4.SkippedTokens)
                yylex.(*Parser).setFreeFloating(listNode, token.ArrayPairList, $6.SkippedTokens)
            }
    |   T_LIST '(' array_pair_list ')'
            {
                // TODO: Cannot use list() as standalone expression
                listNode := &ast.ExprList{ast.Node{}, $3}
                $$ = &ast.ExprArrayItem{ast.Node{}, false, nil, listNode}

                // save position
                listNode.GetNode().Position = position.NewTokensPosition($1, $4)
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating(listNode, token.List, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(listNode, token.ArrayPairList, $4.SkippedTokens)
            }
;

encaps_list:
        encaps_list encaps_var
            {
                $$ = append($1, $2)
            }
    |   encaps_list T_ENCAPSED_AND_WHITESPACE
            {
                encapsed := &ast.ScalarEncapsedStringPart{ast.Node{}, $2.Value}
                $$ = append($1, encapsed)

                // save position
                encapsed.GetNode().Position = position.NewTokenPosition($2)

                // save comments
                yylex.(*Parser).setFreeFloating(encapsed, token.Start, $2.SkippedTokens)
            }
    |   encaps_var
            {
                $$ = []ast.Vertex{$1}
            }
    |   T_ENCAPSED_AND_WHITESPACE encaps_var
            {
                encapsed := &ast.ScalarEncapsedStringPart{ast.Node{}, $1.Value}
                $$ = []ast.Vertex{encapsed, $2}

                // save position
                encapsed.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(encapsed, token.Start, $1.SkippedTokens)
            }
;

encaps_var:
        T_VARIABLE
            {
                name := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.ExprVariable{ast.Node{}, name}

                // save position
                name.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_VARIABLE '[' encaps_var_offset ']'
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, variable, $3}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                variable.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $4.SkippedTokens)
            }
    |   T_VARIABLE T_OBJECT_OPERATOR T_STRING
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                fetch := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.ExprPropertyFetch{ast.Node{}, variable, fetch}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                variable.GetNode().Position = position.NewTokenPosition($1)
                fetch.GetNode().Position = position.NewTokenPosition($3)
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloating(fetch, token.Start, $3.SkippedTokens)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES expr '}'
            {
                variable := &ast.ExprVariable{ast.Node{}, $2}

                $$ = variable

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setToken($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.SkippedTokens)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}'
            {
                name := &ast.Identifier{ast.Node{}, $2.Value}
                variable := &ast.ExprVariable{ast.Node{}, name}

                $$ = variable

                // save position
                name.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setToken($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.SkippedTokens)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
            {
                identifier := &ast.Identifier{ast.Node{}, $2.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, variable, $4}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($2)
                variable.GetNode().Position = position.NewTokenPosition($2)
                $$.GetNode().Position = position.NewTokensPosition($1, $6)

                // save comments
                yylex.(*Parser).setToken(variable, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Var, $3.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.Expr, $5.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $6.SkippedTokens)
            }
    |   T_CURLY_OPEN variable '}'
            {
                $$ = $2;

                // save comments
                yylex.(*Parser).setToken($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens($$, token.End, $3.SkippedTokens)
            }
;

encaps_var_offset:
        T_STRING
            {
                $$ = &ast.ScalarString{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_NUM_STRING
            {
                // TODO: add option to handle 64 bit integer
                if _, err := strconv.Atoi(string($1.Value)); err == nil {
                    $$ = &ast.ScalarLnumber{ast.Node{}, $1.Value}
                } else {
                    $$ = &ast.ScalarString{ast.Node{}, $1.Value}
                }

                // save position
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   '-' T_NUM_STRING
            {
                var lnumber *ast.ScalarLnumber
                // TODO: add option to handle 64 bit integer
                _, err := strconv.Atoi(string($2.Value));
                isInt := err == nil

                if isInt {
                    lnumber = &ast.ScalarLnumber{ast.Node{}, $2.Value}
                    $$ = &ast.ExprUnaryMinus{ast.Node{}, lnumber}
                } else {
                    $2.Value = append([]byte("-"), $2.Value...)
                    $$ = &ast.ScalarString{ast.Node{}, $2.Value}
                }

                // save position
                if isInt {
                    lnumber.GetNode().Position = position.NewTokensPosition($1, $2)
                }
                $$.GetNode().Position = position.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.ExprVariable{ast.Node{}, identifier}

                // save position
                identifier.GetNode().Position = position.NewTokenPosition($1)
                $$.GetNode().Position = position.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

internal_functions_in_yacc:
        T_ISSET '(' isset_variables possible_comma ')'
            {
                $$ = &ast.ExprIsset{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloating($$, token.Isset, $2.SkippedTokens)
                if $4 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.VarList, $5.SkippedTokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.VarList, append($4.SkippedTokens, $5.SkippedTokens...))
                }
            }
    |   T_EMPTY '(' expr ')'
            {
                exprBrackets := &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($2, $4),
                    },
                    OpenBracketTkn:  $2,
                    Child:           $3,
                    CloseBracketTkn: $4,
                }
                $$ = &ast.ExprEmpty{ast.Node{}, exprBrackets}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(exprBrackets, token.Start, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(exprBrackets, token.End, $4.SkippedTokens)
            }
    |   T_INCLUDE expr
            {
                $$ = &ast.ExprInclude{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_INCLUDE_ONCE expr
            {
                $$ = &ast.ExprIncludeOnce{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_EVAL '(' expr ')'
            {
                exprBrackets := &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($2, $4),
                    },
                    OpenBracketTkn:  $2,
                    Child:           $3,
                    CloseBracketTkn: $4,
                }
                $$ = &ast.ExprEval{ast.Node{}, exprBrackets}

                // save position
                $$.GetNode().Position = position.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(exprBrackets, token.Start, $2.SkippedTokens)
                yylex.(*Parser).setFreeFloatingTokens(exprBrackets, token.End, $4.SkippedTokens)
            }
    |   T_REQUIRE expr
            {
                $$ = &ast.ExprRequire{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
    |   T_REQUIRE_ONCE expr
            {
                $$ = &ast.ExprRequireOnce{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = position.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.SkippedTokens)
            }
;

isset_variables:
        isset_variable
            {
                $$ = []ast.Vertex{$1}
            }
    |   isset_variables ',' isset_variable
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.SkippedTokens)
            }
;

isset_variable:
        expr
            {
                $$ = $1
            }
;

/////////////////////////////////////////////////////////////////////////

%%
