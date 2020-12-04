%{
package php7

import (
    "strconv"

    "github.com/z7zmey/php-parser/internal/position"
    "github.com/z7zmey/php-parser/pkg/ast"
    "github.com/z7zmey/php-parser/pkg/token"
)

%}

%union{
    node             ast.Vertex
    token            *token.Token
    list             []ast.Vertex
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
%type <node> const_decl inner_statement for_exprs non_empty_for_exprs
%type <node> expr optional_expr parameter_list non_empty_parameter_list
%type <node> declare_statement finally_statement unset_variable variable
%type <node> parameter optional_type argument expr_without_variable global_var_list global_var
%type <node> static_var_list static_var class_statement trait_adaptation trait_precedence trait_alias
%type <node> absolute_trait_method_reference trait_method_reference property echo_expr
%type <node> new_expr anonymous_class class_name class_name_reference simple_variable
%type <node> internal_functions_in_yacc non_empty_array_pair_list array_pair_list
%type <node> exit_expr scalar lexical_var function_call member_name property_name
%type <node> variable_class_name dereferencable_scalar constant dereferencable
%type <node> callable_expr callable_variable static_member new_variable
%type <node> encaps_var encaps_var_offset echo_expr_list catch_name_list name_list
%type <node> if_stmt const_list non_empty_argument_list property_list
%type <node> alt_if_stmt lexical_var_list isset_variables class_const_list
%type <node> if_stmt_without_else unprefixed_use_declarations inline_use_declarations use_declarations
%type <node> class_const_decl namespace_name
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
%type <node> unset_variables
%type <node> extends_from
%type <node> implements_list
%type <node> interface_extends_list
%type <node> lexical_vars

%type <node> member_modifier
%type <node> use_type
%type <node> foreach_variable


%type <list> encaps_list backticks_expr catch_list
%type <list> case_list trait_adaptation_list
%type <list> top_statement_list
%type <list> inner_statement_list class_statement_list
%type <list> method_modifiers variable_modifiers
%type <list> non_empty_member_modifiers class_modifiers

%%

/////////////////////////////////////////////////////////////////////////

start:
        top_statement_list
            {
                yylex.(*Parser).rootNode = &ast.Root{
                    Node: ast.Node{
                        Position: position.NewNodeListPosition($1),
                    },
                    Stmts:  $1,
                    EndTkn: yylex.(*Parser).currentToken,
                }
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
                $$ = &ast.ParserSeparatedList{
                    Items: []ast.Vertex{
                        &ast.NameNamePart{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($1),
                            },
                            StringTkn: $1,
                            Value:     $1.Value,
                        },
                    },
                }
            }
    |   namespace_name T_NS_SEPARATOR T_STRING
            {
                part := &ast.NameNamePart{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($3),
                    },
                    StringTkn:      $3,
                    Value:          $3.Value,
                }

                $1.(*ast.ParserSeparatedList).SeparatorTkns = append($1.(*ast.ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ast.ParserSeparatedList).Items = append($1.(*ast.ParserSeparatedList).Items, part)

                $$ = $1
            }
;

name:
        namespace_name
            {
                $$ = &ast.NameName{
                    Node:  ast.Node{
                        Position: position.NewNodeListPosition($1.(*ast.ParserSeparatedList).Items),
                    },
                    Parts:         $1.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns: $1.(*ast.ParserSeparatedList).SeparatorTkns,
                }
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameRelative{
                    Node:  ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $3.(*ast.ParserSeparatedList).Items),
                    },
                    NsTkn:          $1,
                    NsSeparatorTkn: $2,
                    Parts:          $3.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:  $3.(*ast.ParserSeparatedList).SeparatorTkns,
                }
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameFullyQualified{
                    Node:  ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $2.(*ast.ParserSeparatedList).Items),
                    },
                    NsSeparatorTkn: $1,
                    Parts:          $2.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:  $2.(*ast.ParserSeparatedList).SeparatorTkns,
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
                            Position: position.NewNodeListPosition($2.(*ast.ParserSeparatedList).Items),
                        },
                        Parts:         $2.(*ast.ParserSeparatedList).Items,
                        SeparatorTkns: $2.(*ast.ParserSeparatedList).SeparatorTkns,
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
                            Position: position.NewNodeListPosition($2.(*ast.ParserSeparatedList).Items),
                        },
                        Parts:         $2.(*ast.ParserSeparatedList).Items,
                        SeparatorTkns: $2.(*ast.ParserSeparatedList).SeparatorTkns,
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
                    UseDeclarations: $2.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:   $2.(*ast.ParserSeparatedList).SeparatorTkns,
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
                    UseDeclarations: $3.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:   $3.(*ast.ParserSeparatedList).SeparatorTkns,
                    SemiColonTkn:    $4,
                }
            }
    |   T_CONST const_list ';'
            {
                $$ = &ast.StmtConstList{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    ConstTkn:      $1,
                    Consts:        $2.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns: $2.(*ast.ParserSeparatedList).SeparatorTkns,
                    SemiColonTkn:  $3,
                }
            }
;

use_type:
        T_FUNCTION
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_CONST
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
;

group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                $4.(*ast.ParserSeparatedList).SeparatorTkns = append($4.(*ast.ParserSeparatedList).SeparatorTkns, $5)

                $$ = &ast.StmtGroupUse{
                    Node: ast.Node{
                        Position: position.NewNodeListTokenPosition($1.(*ast.ParserSeparatedList).Items, $6),
                    },
                    Prefix: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($1.(*ast.ParserSeparatedList).Items),
                        },
                        Parts:         $1.(*ast.ParserSeparatedList).Items,
                        SeparatorTkns: $1.(*ast.ParserSeparatedList).SeparatorTkns,
                    },
                    NsSeparatorTkn:       $2,
                    OpenCurlyBracketTkn:  $3,
                    UseDeclarations:      $4.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:        $4.(*ast.ParserSeparatedList).SeparatorTkns,
                    CloseCurlyBracketTkn: $6,
                }
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                $5.(*ast.ParserSeparatedList).SeparatorTkns = append($5.(*ast.ParserSeparatedList).SeparatorTkns, $6)

                $$ = &ast.StmtGroupUse{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $7),
                    },
                    LeadingNsSeparatorTkn: $1,
                    Prefix: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2.(*ast.ParserSeparatedList).Items),
                        },
                        Parts:         $2.(*ast.ParserSeparatedList).Items,
                        SeparatorTkns: $2.(*ast.ParserSeparatedList).SeparatorTkns,
                    },
                    NsSeparatorTkn:       $3,
                    OpenCurlyBracketTkn:  $4,
                    UseDeclarations:      $5.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:        $5.(*ast.ParserSeparatedList).SeparatorTkns,
                    CloseCurlyBracketTkn: $7,
                }
            }
;

mixed_group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                $4.(*ast.ParserSeparatedList).SeparatorTkns = append($4.(*ast.ParserSeparatedList).SeparatorTkns, $5)

                $$ = &ast.StmtGroupUse{
                    Node: ast.Node{
                        Position: position.NewNodeListTokenPosition($1.(*ast.ParserSeparatedList).Items, $6),
                    },
                    Prefix: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($1.(*ast.ParserSeparatedList).Items),
                        },
                        Parts:         $1.(*ast.ParserSeparatedList).Items,
                        SeparatorTkns: $1.(*ast.ParserSeparatedList).SeparatorTkns,
                    },
                    NsSeparatorTkn:       $2,
                    OpenCurlyBracketTkn:  $3,
                    UseDeclarations:      $4.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:        $4.(*ast.ParserSeparatedList).SeparatorTkns,
                    CloseCurlyBracketTkn: $6,
                }
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                $5.(*ast.ParserSeparatedList).SeparatorTkns = append($5.(*ast.ParserSeparatedList).SeparatorTkns, $6)

                $$ = &ast.StmtGroupUse{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $7),
                    },
                    LeadingNsSeparatorTkn: $1,
                    Prefix: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($2.(*ast.ParserSeparatedList).Items),
                        },
                        Parts:         $2.(*ast.ParserSeparatedList).Items,
                        SeparatorTkns: $2.(*ast.ParserSeparatedList).SeparatorTkns,
                    },
                    NsSeparatorTkn:       $3,
                    OpenCurlyBracketTkn:  $4,
                    UseDeclarations:      $5.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:        $5.(*ast.ParserSeparatedList).SeparatorTkns,
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
                $1.(*ast.ParserSeparatedList).SeparatorTkns = append($1.(*ast.ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ast.ParserSeparatedList).Items = append($1.(*ast.ParserSeparatedList).Items, $3)

                $$ = $1
            }
    |   inline_use_declaration
            {
                $$ = &ast.ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
;

unprefixed_use_declarations:
        unprefixed_use_declarations ',' unprefixed_use_declaration
            {
                $1.(*ast.ParserSeparatedList).SeparatorTkns = append($1.(*ast.ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ast.ParserSeparatedList).Items = append($1.(*ast.ParserSeparatedList).Items, $3)

                $$ = $1
            }
    |   unprefixed_use_declaration
            {
                $$ = &ast.ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
;

use_declarations:
        use_declarations ',' use_declaration
            {
                $1.(*ast.ParserSeparatedList).SeparatorTkns = append($1.(*ast.ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ast.ParserSeparatedList).Items = append($1.(*ast.ParserSeparatedList).Items, $3)

                $$ = $1
            }
    |   use_declaration
            {
                $$ = &ast.ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
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
                        Position: position.NewNodeListPosition($1.(*ast.ParserSeparatedList).Items),
                    },
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($1.(*ast.ParserSeparatedList).Items),
                        },
                        Parts:         $1.(*ast.ParserSeparatedList).Items,
                        SeparatorTkns: $1.(*ast.ParserSeparatedList).SeparatorTkns,
                    },
                }
            }
    |   namespace_name T_AS T_STRING
            {
                $$ = &ast.StmtUseDeclaration{
                    Node: ast.Node{
                        Position: position.NewNodeListTokenPosition($1.(*ast.ParserSeparatedList).Items, $3),
                    },
                    Use: &ast.NameName{
                        Node: ast.Node{
                            Position: position.NewNodeListPosition($1.(*ast.ParserSeparatedList).Items),
                        },
                        Parts:         $1.(*ast.ParserSeparatedList).Items,
                        SeparatorTkns: $1.(*ast.ParserSeparatedList).SeparatorTkns,
                    },
                    AsTkn: $2,
                    Alias: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        IdentifierTkn: $3,
                        Value:         $3.Value,
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
                $1.(*ast.ParserSeparatedList).SeparatorTkns = append($1.(*ast.ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ast.ParserSeparatedList).Items = append($1.(*ast.ParserSeparatedList).Items, $3)

                $$ = $1
            }
    |   const_decl
            {
                $$ = &ast.ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
;

inner_statement_list:
        inner_statement_list inner_statement
            {
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
                $9.(*ast.StmtFor).ForTkn = $1
                $9.(*ast.StmtFor).OpenParenthesisTkn = $2
                $9.(*ast.StmtFor).Init = $3.(*ast.ParserSeparatedList).Items
                $9.(*ast.StmtFor).InitSeparatorTkns = $3.(*ast.ParserSeparatedList).SeparatorTkns
                $9.(*ast.StmtFor).InitSemiColonTkn = $4
                $9.(*ast.StmtFor).Cond = $5.(*ast.ParserSeparatedList).Items
                $9.(*ast.StmtFor).CondSeparatorTkns = $5.(*ast.ParserSeparatedList).SeparatorTkns
                $9.(*ast.StmtFor).CondSemiColonTkn = $6
                $9.(*ast.StmtFor).Loop = $7.(*ast.ParserSeparatedList).Items
                $9.(*ast.StmtFor).LoopSeparatorTkns = $7.(*ast.ParserSeparatedList).SeparatorTkns
                $9.(*ast.StmtFor).CloseParenthesisTkn = $8
                $9.(*ast.StmtFor).Node.Position = position.NewTokenNodePosition($1, $9)

                $$ = $9
            }
    |   T_SWITCH '(' expr ')' switch_case_list
            {
                $5.(*ast.StmtSwitch).SwitchTkn = $1
                $5.(*ast.StmtSwitch).OpenParenthesisTkn = $2
                $5.(*ast.StmtSwitch).Cond = $3
                $5.(*ast.StmtSwitch).CloseParenthesisTkn = $4
                $5.(*ast.StmtSwitch).Node.Position = position.NewTokenNodePosition($1, $5)

                $$ = $5
            }
    |   T_BREAK optional_expr ';'
            {
                $$ = &ast.StmtBreak{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    BreakTkn:     $1,
                    Expr:         $2,
                    SemiColonTkn: $3,
                }
            }
    |   T_CONTINUE optional_expr ';'
            {
                $$ = &ast.StmtContinue{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    ContinueTkn:  $1,
                    Expr:         $2,
                    SemiColonTkn: $3,
                }
            }
    |   T_RETURN optional_expr ';'
            {
                $$ = &ast.StmtReturn{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    ReturnTkn:    $1,
                    Expr:         $2,
                    SemiColonTkn: $3,
                }
            }
    |   T_GLOBAL global_var_list ';'
            {
                $2.(*ast.StmtGlobal).GlobalTkn = $1
                $2.(*ast.StmtGlobal).SemiColonTkn = $3
                $2.(*ast.StmtGlobal).Node.Position = position.NewTokensPosition($1, $3)

                $$ = $2
            }
    |   T_STATIC static_var_list ';'
            {
                $2.(*ast.StmtStatic).StaticTkn = $1
                $2.(*ast.StmtStatic).SemiColonTkn = $3
                $2.(*ast.StmtStatic).Node.Position = position.NewTokensPosition($1, $3)

                $$ = $2
            }
    |   T_ECHO echo_expr_list ';'
            {
                $2.(*ast.StmtEcho).EchoTkn = $1
                $2.(*ast.StmtEcho).SemiColonTkn = $3
                $2.(*ast.StmtEcho).Node.Position = position.NewTokensPosition($1, $3)

                $$ = $2
            }
    |   T_INLINE_HTML
            {
                $$ = &ast.StmtInlineHtml{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    InlineHtmlTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   expr ';'
            {
                $$ = &ast.StmtExpression{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $2),
                    },
                    Expr:         $1,
                    SemiColonTkn: $2,
                }
            }
    |   T_UNSET '(' unset_variables possible_comma ')' ';'
            {
                $3.(*ast.StmtUnset).UnsetTkn = $1
                $3.(*ast.StmtUnset).OpenParenthesisTkn = $2
                $3.(*ast.StmtUnset).SeparatorTkns = append($3.(*ast.StmtUnset).SeparatorTkns, $4)
                $3.(*ast.StmtUnset).CloseParenthesisTkn = $5
                $3.(*ast.StmtUnset).SemiColonTkn = $6
                $3.(*ast.StmtUnset).Node.Position = position.NewTokensPosition($1, $6)

                $$ = $3
            }
    |   T_FOREACH '(' expr T_AS foreach_variable ')' foreach_statement
            {
                $7.(*ast.StmtForeach).ForeachTkn = $1
                $7.(*ast.StmtForeach).OpenParenthesisTkn = $2
                $7.(*ast.StmtForeach).Expr = $3
                $7.(*ast.StmtForeach).AsTkn = $4
                $7.(*ast.StmtForeach).Var = $5
                $7.(*ast.StmtForeach).CloseParenthesisTkn = $6
                $7.(*ast.StmtForeach).Node.Position = position.NewTokenNodePosition($1, $7)

                $$ = $7
            }
    |   T_FOREACH '(' expr T_AS variable T_DOUBLE_ARROW foreach_variable ')' foreach_statement
            {
                $9.(*ast.StmtForeach).ForeachTkn = $1
                $9.(*ast.StmtForeach).OpenParenthesisTkn = $2
                $9.(*ast.StmtForeach).Expr = $3
                $9.(*ast.StmtForeach).AsTkn = $4
                $9.(*ast.StmtForeach).Key = $5
                $9.(*ast.StmtForeach).DoubleArrowTkn = $6
                $9.(*ast.StmtForeach).Var = $7
                $9.(*ast.StmtForeach).CloseParenthesisTkn = $8
                $9.(*ast.StmtForeach).Node.Position = position.NewTokenNodePosition($1, $9)

                $$ = $9
            }
    |   T_DECLARE '(' const_list ')' declare_statement
            {
                $5.(*ast.StmtDeclare).DeclareTkn = $1
                $5.(*ast.StmtDeclare).OpenParenthesisTkn = $2
                $5.(*ast.StmtDeclare).Consts = $3.(*ast.ParserSeparatedList).Items
                $5.(*ast.StmtDeclare).SeparatorTkns = $3.(*ast.ParserSeparatedList).SeparatorTkns
                $5.(*ast.StmtDeclare).CloseParenthesisTkn = $4
                $5.(*ast.StmtDeclare).Node.Position = position.NewTokenNodePosition($1, $5)

                $$ = $5
            }
    |   ';'
            {
                $$ = &ast.StmtNop{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    SemiColonTkn: $1,
                }
            }
    |   T_TRY '{' inner_statement_list '}' catch_list finally_statement
            {
                $$ = &ast.StmtTry{
                    TryTkn:            $1,
                    OpenCurlyBracket:  $2,
                    Stmts:             $3,
                    CloseCurlyBracket: $4,
                    Catches:           $5,
                    Finally:           $6,
                }

                if $6 == nil {
                    $$.GetNode().Position = position.NewTokenNodeListPosition($1, $5)
                } else {
                    $$.GetNode().Position = position.NewTokenNodePosition($1, $6)
                }
            }
    |   T_THROW expr ';'
            {
                $$ = &ast.StmtThrow{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    ThrowTkn:     $1,
                    Expr:         $2,
                    SemiColonTkn: $3,
                }
            }
    |   T_GOTO T_STRING ';'
            {
                $$ = &ast.StmtGoto{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    GotoTkn: $1,
                    Label: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($2),
                        },
                        IdentifierTkn: $2,
                        Value:         $2.Value,
                    },
                    SemiColonTkn: $3,
                }
            }
    |   T_STRING ':'
            {
                $$ = &ast.StmtLabel{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $2),
                    },
                    LabelName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                    ColonTkn: $2,
                }
            }

catch_list:
        /* empty */
            {
                $$ = []ast.Vertex{}
            }
    |   catch_list T_CATCH '(' catch_name_list T_VARIABLE ')' '{' inner_statement_list '}'
            {
                catch := $4.(*ast.StmtCatch)
                catch.CatchTkn = $2
                catch.OpenParenthesisTkn = $3
                catch.Var = &ast.ExprVariable{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($5),
                    },
                    VarName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($5),
                        },
                        IdentifierTkn: $5,
                        Value:         $5.Value,
                    },
                }
                catch.CloseParenthesisTkn = $6
                catch.OpenCurlyBracketTkn = $7
                catch.Stmts = $8
                catch.CloseCurlyBracketTkn = $9
                catch.GetNode().Position = position.NewTokensPosition($2, $9)

                $$ = append($1, catch)
            }
;
catch_name_list:
        name
            {
                $$ = &ast.StmtCatch{
                    Types: []ast.Vertex{$1},
                }
            }
    |   catch_name_list '|' name
            {
                $1.(*ast.StmtCatch).SeparatorTkns = append($1.(*ast.StmtCatch).SeparatorTkns, $2)
                $1.(*ast.StmtCatch).Types = append($1.(*ast.StmtCatch).Types, $3)

                $$ = $1
            }
;

finally_statement:
        /* empty */
            {
                $$ = nil
            }
    |   T_FINALLY '{' inner_statement_list '}'
            {
                $$ = &ast.StmtFinally{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    FinallyTkn:           $1,
                    OpenCurlyBracketTkn:  $2,
                    Stmts:                $3,
                    CloseCurlyBracketTkn: $4,
                }
            }
;

unset_variables:
        unset_variable
            {
                $$ = &ast.StmtUnset{
                    Vars: []ast.Vertex{$1},
                }
            }
    |   unset_variables ',' unset_variable
            {
                $1.(*ast.StmtUnset).Vars = append($1.(*ast.StmtUnset).Vars, $3)
                $1.(*ast.StmtUnset).SeparatorTkns = append($1.(*ast.StmtUnset).SeparatorTkns, $2)

                $$ = $1
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
                $$ = &ast.StmtFunction{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $11),
                    },
                    FunctionTkn:  $1,
                    AmpersandTkn: $2,
                    FunctionName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                    OpenParenthesisTkn:   $5,
                    Params:               $6.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:        $6.(*ast.ParserSeparatedList).SeparatorTkns,
                    CloseParenthesisTkn:  $7,
                    ColonTkn:             $8.(*ast.ReturnType).ColonTkn,
                    ReturnType:           $8.(*ast.ReturnType).Type,
                    OpenCurlyBracketTkn:  $9,
                    Stmts:                $10,
                    CloseCurlyBracketTkn: $11,
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
                $$ = &ast.StmtClass{
                    Node: ast.Node{
                        Position: position.NewOptionalListTokensPosition($1, $2, $9),
                    },
                    Modifiers: $1,
                    ClassTkn:  $2,
                    ClassName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                    Extends:           $4,
                    Implements:        $5,
                    OpenCurlyBracket:  $7,
                    Stmts:             $8,
                    CloseCurlyBracket: $9,
                }
            }
    |   T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
            {
                $$ = &ast.StmtClass{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $8),
                    },
                    ClassTkn: $1,
                    ClassName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($2),
                        },
                        IdentifierTkn: $2,
                        Value:         $2.Value,
                    },
                    Extends:           $3,
                    Implements:        $4,
                    OpenCurlyBracket:  $6,
                    Stmts:             $7,
                    CloseCurlyBracket: $8,
                }
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
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_FINAL
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
;

trait_declaration_statement:
        T_TRAIT T_STRING backup_doc_comment '{' class_statement_list '}'
            {
                $$ = &ast.StmtTrait{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $6),
                    },
                    TraitTkn: $1,
                    TraitName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($2),
                        },
                        IdentifierTkn: $2,
                        Value:         $2.Value,
                    },
                    OpenCurlyBracket:  $4,
                    Stmts:             $5,
                    CloseCurlyBracket: $6,
                }
            }
;

interface_declaration_statement:
        T_INTERFACE T_STRING interface_extends_list backup_doc_comment '{' class_statement_list '}'
            {
                $$ = &ast.StmtInterface{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $7),
                    },
                    InterfaceTkn: $1,
                    InterfaceName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($2),
                        },
                        IdentifierTkn: $2,
                        Value:         $2.Value,
                    },
                    Extends:              $3,
                    OpenCurlyBracketTkn:  $5,
                    Stmts:                $6,
                    CloseCurlyBracketTkn: $7,
                }
            }
;

extends_from:
        /* empty */
            {
                $$ = nil
            }
    |   T_EXTENDS name
            {
                $$ = &ast.StmtClassExtends{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    ExtendTkn: $1,
                    ClassName: $2,
                }
            }
;

interface_extends_list:
        /* empty */
            {
                $$ = nil
            }
    |   T_EXTENDS name_list
            {
                $$ = &ast.StmtInterfaceExtends{
                    Node: ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $2.(*ast.ParserSeparatedList).Items),
                    },
                    ExtendsTkn:     $1,
                    InterfaceNames: $2.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:  $2.(*ast.ParserSeparatedList).SeparatorTkns,
                };
            }
;

implements_list:
        /* empty */
            {
                $$ = nil
            }
    |   T_IMPLEMENTS name_list
            {
                $$ = &ast.StmtClassImplements{
                    Node: ast.Node{
                        Position: position.NewTokenNodeListPosition($1, $2.(*ast.ParserSeparatedList).Items),
                    },
                    ImplementsTkn:  $1,
                    InterfaceNames: $2.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:  $2.(*ast.ParserSeparatedList).SeparatorTkns,
                };
            }
;

foreach_variable:
        variable
            {
                $$ = $1
            }
    |   '&' variable
            {
                $$ = &ast.ExprReference{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    AmpersandTkn: $1,
                    Var:          $2,
                }
            }
    |   T_LIST '(' array_pair_list ')'
            {
                $$ = &ast.ExprList{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    ListTkn:         $1,
                    OpenBracketTkn:  $2,
                    Items:           $3.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:   $3.(*ast.ParserSeparatedList).SeparatorTkns,
                    CloseBracketTkn: $4,
                }
            }
    |   '[' array_pair_list ']'
            {
                $$ = &ast.ExprList{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenBracketTkn:  $1,
                    Items:           $2.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:   $2.(*ast.ParserSeparatedList).SeparatorTkns,
                    CloseBracketTkn: $3,
                }
            }
;

for_statement:
        statement
            {
                $$ = &ast.StmtFor{
                    Node: ast.Node{
                        Position: position.NewNodePosition($1),
                    },
                    Stmt: $1,
                }
            }
    |   ':' inner_statement_list T_ENDFOR ';'
            {
                $$ = &ast.StmtFor{
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
                    EndForTkn:    $3,
                    SemiColonTkn: $4,
                }
            }
;

foreach_statement:
        statement
            {
                $$ = &ast.StmtForeach{
                    Node: ast.Node{
                        Position: position.NewNodePosition($1),
                    },
                    Stmt: $1,
                }
            }
    |   ':' inner_statement_list T_ENDFOREACH ';'
            {
                $$ = &ast.StmtForeach{
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
                    EndForeachTkn: $3,
                    SemiColonTkn:  $4,
                }
            }
;

declare_statement:
        statement
            {
                $$ = &ast.StmtDeclare{
                    Node: ast.Node{
                        Position: position.NewNodePosition($1),
                    },
                    Stmt: $1,
                }
            }
    |   ':' inner_statement_list T_ENDDECLARE ';'
            {
                $$ = &ast.StmtDeclare{
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
                    EndDeclareTkn: $3,
                    SemiColonTkn:  $4,
                }
            }
;

switch_case_list:
        '{' case_list '}'
            {
                $$ = &ast.StmtSwitch{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenCurlyBracketTkn:  $1,
                    CaseList:             $2,
                    CloseCurlyBracketTkn: $3,
                }
            }
    |   '{' ';' case_list '}'
            {
                $$ = &ast.StmtSwitch{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    OpenCurlyBracketTkn:  $1,
                    CaseSeparatorTkn:     $2,
                    CaseList:             $3,
                    CloseCurlyBracketTkn: $4,
                }
            }
    |   ':' case_list T_ENDSWITCH ';'
            {
                $$ = &ast.StmtSwitch{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    Alt:          true,
                    ColonTkn:     $1,
                    CaseList:     $2,
                    EndSwitchTkn: $3,
                    SemiColonTkn: $4,
                }
            }
    |   ':' ';' case_list T_ENDSWITCH ';'
            {
                $$ = &ast.StmtSwitch{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $5),
                    },
                    Alt:              true,
                    ColonTkn:         $1,
                    CaseSeparatorTkn: $2,
                    CaseList:         $3,
                    EndSwitchTkn:     $4,
                    SemiColonTkn:     $5,
                }
            }
;

case_list:
        /* empty */
            {
                $$ = nil
            }
    |   case_list T_CASE expr case_separator inner_statement_list
            {
                $$ = append($1, &ast.StmtCase{
                    Node: ast.Node{
                        Position: position.NewTokenNodeListPosition($2, $5),
                    },
                    CaseTkn:          $2,
                    Cond:             $3,
                    CaseSeparatorTkn: $4,
                    Stmts:            $5,
                })
            }
    |   case_list T_DEFAULT case_separator inner_statement_list
            {
                $$ = append($1, &ast.StmtDefault{
                    Node: ast.Node{
                        Position: position.NewTokenNodeListPosition($2, $4),
                    },
                    DefaultTkn:       $2,
                    CaseSeparatorTkn: $3,
                    Stmts:            $4,
                })
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
                $$ = &ast.ParserSeparatedList{}
            }
;

non_empty_parameter_list:
        parameter
            {
                $$ = &ast.ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
    |   non_empty_parameter_list ',' parameter
            {
                $1.(*ast.ParserSeparatedList).SeparatorTkns = append($1.(*ast.ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ast.ParserSeparatedList).Items = append($1.(*ast.ParserSeparatedList).Items, $3)

                $$ = $1
            }
;

parameter:
        optional_type is_reference is_variadic T_VARIABLE
            {
                pos := position.NewTokenPosition($4)
                if $1 != nil {
                    pos = position.NewNodeTokenPosition($1, $4)
                } else if $2 != nil {
                    pos = position.NewTokensPosition($2, $4)
                } else if $3 != nil {
                    pos = position.NewTokensPosition($3, $4)
                }

                $$ = &ast.Parameter{
                    Node: ast.Node{
                        Position: pos,
                    },
                    Type:         $1,
                    AmpersandTkn: $2,
                    VariadicTkn:  $3,
                    Var: &ast.ExprVariable{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($4),
                        },
                        VarName: &ast.Identifier{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($4),
                            },
                            IdentifierTkn: $4,
                            Value:         $4.Value,
                        },
                    },
                }
            }
    |   optional_type is_reference is_variadic T_VARIABLE '=' expr
            {
                pos := position.NewTokenNodePosition($4, $6)
                if $1 != nil {
                    pos = position.NewNodesPosition($1, $6)
                } else if $2 != nil {
                    pos = position.NewTokenNodePosition($2, $6)
                } else if $3 != nil {
                    pos = position.NewTokenNodePosition($3, $6)
                }

                $$ = &ast.Parameter{
                    Node: ast.Node{
                        Position: pos,
                    },
                    Type:         $1,
                    AmpersandTkn: $2,
                    VariadicTkn:  $3,
                    Var: &ast.ExprVariable{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($4),
                        },
                        VarName: &ast.Identifier{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($4),
                            },
                            IdentifierTkn: $4,
                            Value:         $4.Value,
                        },
                    },
                    EqualTkn:     $5,
                    DefaultValue: $6,
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
                $$ = &ast.Nullable{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    QuestionTkn: $1,
                    Expr:        $2,
                }
            }
;

type:
        T_ARRAY
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_CALLABLE
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   name
            {
                $$ = $1
            }
;

return_type:
        /* empty */
            {
                $$ = &ast.ReturnType{}
            }
    |   ':' type_expr
            {
                $$ = &ast.ReturnType{
                    ColonTkn: $1,
                    Type:     $2,
                }
            }
;

argument_list:
        '(' ')'
            {
                $$ = &ast.ArgumentList{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $2),
                    },
                    OpenParenthesisTkn: $1,
                    CloseParenthesisTkn: $2,
                }
            }
    |   '(' non_empty_argument_list possible_comma ')'
            {
                argumentList := $2.(*ast.ArgumentList)
                argumentList.Position = position.NewTokensPosition($1, $4)
                argumentList.OpenParenthesisTkn = $1
                argumentList.SeparatorTkns = append(argumentList.SeparatorTkns, $3)
                argumentList.CloseParenthesisTkn = $4

                $$ = argumentList
            }
;

non_empty_argument_list:
        argument
            {
                $$ = &ast.ArgumentList{
                    Arguments: []ast.Vertex{$1},
                }
            }
    |   non_empty_argument_list ',' argument
            {
                $1.(*ast.ArgumentList).SeparatorTkns = append($1.(*ast.ArgumentList).SeparatorTkns, $2)
                $1.(*ast.ArgumentList).Arguments = append($1.(*ast.ArgumentList).Arguments, $3)

                $$ = $1
            }
;

argument:
        expr
            {
                $$ = &ast.Argument{
                    Node: ast.Node{
                        Position: position.NewNodePosition($1),
                    },
                    Expr:        $1,
                }
            }
    |   T_ELLIPSIS expr
            {
                $$ = &ast.Argument{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    VariadicTkn: $1,
                    Expr:        $2,
                }
            }
;

global_var_list:
        global_var_list ',' global_var
            {
                $1.(*ast.StmtGlobal).Vars = append($1.(*ast.StmtGlobal).Vars, $3)
                $1.(*ast.StmtGlobal).SeparatorTkns = append($1.(*ast.StmtGlobal).SeparatorTkns, $2)

                $$ = $1
            }
    |   global_var
            {
                $$ = &ast.StmtGlobal{
                    Vars: []ast.Vertex{$1},
                }
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
                $1.(*ast.StmtStatic).Vars = append($1.(*ast.StmtStatic).Vars, $3)
                $1.(*ast.StmtStatic).SeparatorTkns = append($1.(*ast.StmtStatic).SeparatorTkns, $2)

                $$ = $1
            }
    |   static_var
            {
                $$ = &ast.StmtStatic{
                    Vars: []ast.Vertex{$1},
                }
            }
;

static_var:
        T_VARIABLE
            {

                $$ = &ast.StmtStaticVar{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    Var: &ast.ExprVariable{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        VarName: &ast.Identifier{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($1),
                            },
                            IdentifierTkn: $1,
                            Value:         $1.Value,
                        },
                    },
                }
            }
    |   T_VARIABLE '=' expr
            {
                $$ = &ast.StmtStaticVar{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $3),
                    },
                    Var: &ast.ExprVariable{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        VarName: &ast.Identifier{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($1),
                            },
                            IdentifierTkn: $1,
                            Value:         $1.Value,
                        },
                    },
                    EqualTkn: $2,
                    Expr:     $3,
                }
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
                $$ = &ast.StmtPropertyList{
                    Node: ast.Node{
                        Position: position.NewNodeListTokenPosition($1, $4),
                    },
                    Modifiers:     $1,
                    Type:          $2,
                    Properties:    $3.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns: $3.(*ast.ParserSeparatedList).SeparatorTkns,
                    SemiColonTkn:  $4,
                }
            }
    |   method_modifiers T_CONST class_const_list ';'
            {
                $$ = &ast.StmtClassConstList{
                    Node: ast.Node{
                        Position: position.NewOptionalListTokensPosition($1, $2, $4),
                    },
                    Modifiers:     $1,
                    ConstTkn:      $2,
                    Consts:        $3.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns: $3.(*ast.ParserSeparatedList).SeparatorTkns,
                    SemiColonTkn:  $4,
                }
            }
    |   T_USE name_list trait_adaptations
            {
                $$ = &ast.StmtTraitUse{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $3),
                    },
                    UseTkn:        $1,
                    Traits:        $2.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns: $2.(*ast.ParserSeparatedList).SeparatorTkns,
                    Adaptations:   $3,
                }
            }
    |   method_modifiers T_FUNCTION returns_ref identifier backup_doc_comment '(' parameter_list ')' return_type method_body
            {
                pos := position.NewTokenNodePosition($2, $10)
                if $1 != nil {
                    $$.GetNode().Position = position.NewNodeListNodePosition($1, $10)
                }

                $$ = &ast.StmtClassMethod{
                    Node: ast.Node{
                        Position: pos,
                    },
                    Modifiers:    $1,
                    FunctionTkn:  $2,
                    AmpersandTkn: $3,
                    MethodName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($4),
                        },
                        IdentifierTkn: $4,
                        Value:         $4.Value,
                    },
                    OpenParenthesisTkn:  $6,
                    Params:              $7.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:       $7.(*ast.ParserSeparatedList).SeparatorTkns,
                    CloseParenthesisTkn: $8,
                    ColonTkn:            $9.(*ast.ReturnType).ColonTkn,
                    ReturnType:          $9.(*ast.ReturnType).Type,
                    Stmt:                $10,
                }
            }
;

name_list:
        name
            {
                $$ = &ast.ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
    |   name_list ',' name
            {
                $1.(*ast.ParserSeparatedList).SeparatorTkns = append($1.(*ast.ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ast.ParserSeparatedList).Items = append($1.(*ast.ParserSeparatedList).Items, $3)

                $$ = $1
            }
;

trait_adaptations:
        ';'
            {
                $$ = &ast.StmtNop{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    SemiColonTkn: $1,
                }
            }
    |   '{' '}'
            {
                $$ = &ast.StmtTraitAdaptationList{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $2),
                    },
                    OpenParenthesisTkn:  $1,
                    CloseParenthesisTkn: $2,
                }
            }
    |   '{' trait_adaptation_list '}'
            {
                $$ = &ast.StmtTraitAdaptationList{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenParenthesisTkn:  $1,
                    Adaptations:         $2,
                    CloseParenthesisTkn: $3,
                }
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
                $1.(*ast.StmtTraitUsePrecedence).SemiColonTkn = $2

                $$ = $1;
            }
    |   trait_alias ';'
            {
                $1.(*ast.StmtTraitUseAlias).SemiColonTkn = $2

                $$ = $1;
            }
;

trait_precedence:
        absolute_trait_method_reference T_INSTEADOF name_list
            {
                $$ = &ast.StmtTraitUsePrecedence{
                    Node: ast.Node{
                        Position: position.NewNodeNodeListPosition($1, $3.(*ast.ParserSeparatedList).Items),
                    },
                    Ref:           $1,
                    InsteadofTkn:  $2,
                    Insteadof:     $3.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns: $3.(*ast.ParserSeparatedList).SeparatorTkns,
                }
            }
;

trait_alias:
        trait_method_reference T_AS T_STRING
            {
                $$ = &ast.StmtTraitUseAlias{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $3),
                    },
                    Ref:   $1,
                    AsTkn: $2,
                    Alias: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
    |   trait_method_reference T_AS reserved_non_modifiers
            {
                $$ = &ast.StmtTraitUseAlias{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $3),
                    },
                    Ref:   $1,
                    AsTkn: $2,
                    Alias: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
    |   trait_method_reference T_AS member_modifier identifier
            {
                $$ = &ast.StmtTraitUseAlias{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $4),
                    },
                    Ref:      $1,
                    AsTkn:    $2,
                    Modifier: $3,
                    Alias: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($4),
                        },
                        IdentifierTkn: $4,
                        Value:         $4.Value,
                    },
                }
            }
    |   trait_method_reference T_AS member_modifier
            {
                $$ = &ast.StmtTraitUseAlias{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Ref:      $1,
                    AsTkn:    $2,
                    Modifier: $3,
                }
            }
;

trait_method_reference:
        identifier
            {
                $$ = &ast.StmtTraitMethodRef{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    Method: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                }
            }
    |   absolute_trait_method_reference
            {
                $$ = $1
            }
;

absolute_trait_method_reference:
        name T_PAAMAYIM_NEKUDOTAYIM identifier
            {
                $$ = &ast.StmtTraitMethodRef{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $3),
                    },
                    Trait:          $1,
                    DoubleColonTkn: $2,
                    Method: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
;

method_body:
        ';' /* abstract method */
            {
                $$ = &ast.StmtNop{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    SemiColonTkn: $1,
                }
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
                $$ = []ast.Vertex{
                    &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                }
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
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_PROTECTED
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_PRIVATE
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_STATIC
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_ABSTRACT
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_FINAL
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
;

property_list:
        property_list ',' property
            {
                $1.(*ast.ParserSeparatedList).SeparatorTkns = append($1.(*ast.ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ast.ParserSeparatedList).Items = append($1.(*ast.ParserSeparatedList).Items, $3)

                $$ = $1
            }
    |   property
            {
                $$ = &ast.ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
;

property:
        T_VARIABLE backup_doc_comment
            {
                $$ = &ast.StmtProperty{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    Var: &ast.ExprVariable{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        VarName: &ast.Identifier{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($1),
                            },
                            IdentifierTkn: $1,
                            Value:         $1.Value,
                        },
                    },
                    Expr: nil,
                }
            }
    |   T_VARIABLE '=' expr backup_doc_comment
            {
                $$ = &ast.StmtProperty{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $3),
                    },
                    Var: &ast.ExprVariable{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        VarName: &ast.Identifier{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($1),
                            },
                            IdentifierTkn: $1,
                            Value:         $1.Value,
                        },
                    },
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
;

class_const_list:
        class_const_list ',' class_const_decl
            {
                $1.(*ast.ParserSeparatedList).SeparatorTkns = append($1.(*ast.ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ast.ParserSeparatedList).Items = append($1.(*ast.ParserSeparatedList).Items, $3)

                $$ = $1
            }
    |   class_const_decl
            {
                $$ = &ast.ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
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
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                    EqualTkn: $2,
                    Expr:     $3,
                }
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
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
;

echo_expr_list:
        echo_expr_list ',' echo_expr
            {
                $1.(*ast.StmtEcho).Exprs = append($1.(*ast.StmtEcho).Exprs, $3)
                $1.(*ast.StmtEcho).SeparatorTkns = append($1.(*ast.StmtEcho).SeparatorTkns, $2)

                $$ = $1
            }
    |   echo_expr
            {
                $$ = &ast.StmtEcho{
                    Exprs: []ast.Vertex{$1},
                }
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
                $$ = &ast.ParserSeparatedList{}
            }
    |   non_empty_for_exprs
            {
                $$ = $1
            }
;

non_empty_for_exprs:
        non_empty_for_exprs ',' expr
            {
                $1.(*ast.ParserSeparatedList).SeparatorTkns = append($1.(*ast.ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ast.ParserSeparatedList).Items = append($1.(*ast.ParserSeparatedList).Items, $3)

                $$ = $1
            }
    |   expr
            {
                $$ = &ast.ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
;

anonymous_class:
        T_CLASS ctor_arguments extends_from implements_list backup_doc_comment '{' class_statement_list '}'
            {
                $$ = &ast.StmtClass{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $8),
                    },
                    ClassTkn:            $1,
                    OpenParenthesisTkn:  $2.(*ast.ArgumentList).OpenParenthesisTkn,
                    Arguments:           $2.(*ast.ArgumentList).Arguments,
                    SeparatorTkns:       $2.(*ast.ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $2.(*ast.ArgumentList).OpenParenthesisTkn,
                    Extends:             $3,
                    Implements:          $4,
                    OpenCurlyBracket:    $6,
                    Stmts:               $7,
                    CloseCurlyBracket:   $8,
                }
            }
;

new_expr:
        T_NEW class_name_reference ctor_arguments
            {
                if $3 != nil {
                    $$ = &ast.ExprNew{
                        Node: ast.Node{
                            Position: position.NewTokenNodePosition($1, $3),
                        },
                        NewTkn:              $1,
                        Class:               $2,
                        OpenParenthesisTkn:  $3.(*ast.ArgumentList).OpenParenthesisTkn,
                        Arguments:           $3.(*ast.ArgumentList).Arguments,
                        SeparatorTkns:       $3.(*ast.ArgumentList).SeparatorTkns,
                        CloseParenthesisTkn: $3.(*ast.ArgumentList).OpenParenthesisTkn,
                    }
                } else {
                    $$ = &ast.ExprNew{
                        Node: ast.Node{
                            Position: position.NewTokenNodePosition($1, $2),
                        },
                        NewTkn: $1,
                        Class:  $2,
                    }
                }
            }
    |   T_NEW anonymous_class
            {
                $$ = &ast.ExprNew{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    NewTkn: $1,
                    Class:  $2,
                }
            }
;

expr_without_variable:
        T_LIST '(' array_pair_list ')' '=' expr
            {
                $$ = &ast.ExprAssign{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $6),
                    },
                    Var: &ast.ExprList{
                        Node: ast.Node{
                            Position: position.NewTokensPosition($1, $4),
                        },
                        ListTkn:         $1,
                        OpenBracketTkn:  $2,
                        Items:           $3.(*ast.ParserSeparatedList).Items,
                        SeparatorTkns:   $3.(*ast.ParserSeparatedList).SeparatorTkns,
                        CloseBracketTkn: $4,
                    },
                    EqualTkn: $5,
                    Expr:     $6,
                }
            }
    |   '[' array_pair_list ']' '=' expr
            {
                $$ = &ast.ExprAssign{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $5),
                    },
                    Var: &ast.ExprList{
                        Node: ast.Node{
                            Position: position.NewTokensPosition($1, $3),
                        },
                        OpenBracketTkn:  $1,
                        Items:           $2.(*ast.ParserSeparatedList).Items,
                        SeparatorTkns:   $2.(*ast.ParserSeparatedList).SeparatorTkns,
                        CloseBracketTkn: $3,
                    },
                    EqualTkn: $4,
                    Expr:     $5,
                }
            }
    |   variable '=' expr
            {
                $$ = &ast.ExprAssign{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable '=' '&' expr
            {
                $$ = &ast.ExprAssignReference{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $4),
                    },
                    Var:          $1,
                    EqualTkn:     $2,
                    AmpersandTkn: $3,
                    Expr:         $4,
                }
            }
    |   T_CLONE expr
            {
                $$ = &ast.ExprClone{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    CloneTkn: $1,
                    Expr:     $2,
                }
            }
    |   variable T_PLUS_EQUAL expr
            {
                $$ = &ast.ExprAssignPlus{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_MINUS_EQUAL expr
            {
                $$ = &ast.ExprAssignMinus{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_MUL_EQUAL expr
            {
                $$ = &ast.ExprAssignMul{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_POW_EQUAL expr
            {
                $$ = &ast.ExprAssignPow{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_DIV_EQUAL expr
            {
                $$ = &ast.ExprAssignDiv{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_CONCAT_EQUAL expr
            {
                $$ = &ast.ExprAssignConcat{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_MOD_EQUAL expr
            {
                $$ = &ast.ExprAssignMod{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_AND_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseAnd{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_OR_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseOr{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_XOR_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseXor{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_SL_EQUAL expr
            {
                $$ = &ast.ExprAssignShiftLeft{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_SR_EQUAL expr
            {
                $$ = &ast.ExprAssignShiftRight{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_COALESCE_EQUAL expr
            {
                $$ = &ast.ExprAssignCoalesce{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:      $1,
                    EqualTkn: $2,
                    Expr:     $3,
                }
            }
    |   variable T_INC
            {
                $$ = &ast.ExprPostInc{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $2),
                    },
                    Var:    $1,
                    IncTkn: $2,
                }
            }
    |   T_INC variable
            {
                $$ = &ast.ExprPreInc{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    IncTkn: $1,
                    Var:    $2,
                }
            }
    |   variable T_DEC
            {
                $$ = &ast.ExprPostDec{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $2),
                    },
                    Var:    $1,
                    DecTkn: $2,
                }
            }
    |   T_DEC variable
            {
                $$ = &ast.ExprPreDec{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    DecTkn: $1,
                    Var:    $2,
                }
            }
    |   expr T_BOOLEAN_OR expr
            {
                $$ = &ast.ExprBinaryBooleanOr{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_BOOLEAN_AND expr
            {
                $$ = &ast.ExprBinaryBooleanAnd{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_LOGICAL_OR expr
            {
                $$ = &ast.ExprBinaryLogicalOr{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_LOGICAL_AND expr
            {
                $$ = &ast.ExprBinaryLogicalAnd{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_LOGICAL_XOR expr
            {
                $$ = &ast.ExprBinaryLogicalXor{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '|' expr
            {
                $$ = &ast.ExprBinaryBitwiseOr{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '&' expr
            {
                $$ = &ast.ExprBinaryBitwiseAnd{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '^' expr
            {
                $$ = &ast.ExprBinaryBitwiseXor{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '.' expr
            {
                $$ = &ast.ExprBinaryConcat{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '+' expr
            {
                $$ = &ast.ExprBinaryPlus{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '-' expr
            {
                $$ = &ast.ExprBinaryMinus{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '*' expr
            {
                $$ = &ast.ExprBinaryMul{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_POW expr
            {
                $$ = &ast.ExprBinaryPow{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '/' expr
            {
                $$ = &ast.ExprBinaryDiv{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '%' expr
            {
                $$ = &ast.ExprBinaryMod{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_SL expr
            {
                $$ = &ast.ExprBinaryShiftLeft{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_SR expr
            {
                $$ = &ast.ExprBinaryShiftRight{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   '+' expr %prec T_INC
            {
                $$ = &ast.ExprUnaryPlus{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    PlusTkn: $1,
                    Expr:    $2,
                }
            }
    |   '-' expr %prec T_INC
            {
                $$ = &ast.ExprUnaryMinus{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    MinusTkn: $1,
                    Expr:     $2,
                }
            }
    |   '!' expr
            {
                $$ = &ast.ExprBooleanNot{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    ExclamationTkn: $1,
                    Expr:           $2,
                }
            }
    |   '~' expr
            {
                $$ = &ast.ExprBitwiseNot{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    TildaTkn: $1,
                    Expr:     $2,
                }
            }
    |   expr T_IS_IDENTICAL expr
            {
                $$ = &ast.ExprBinaryIdentical{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_IS_NOT_IDENTICAL expr
            {
                $$ = &ast.ExprBinaryNotIdentical{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_IS_EQUAL expr
            {
                $$ = &ast.ExprBinaryEqual{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_IS_NOT_EQUAL expr
            {
                $$ = &ast.ExprBinaryNotEqual{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '<' expr
            {
                $$ = &ast.ExprBinarySmaller{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_IS_SMALLER_OR_EQUAL expr
            {
                $$ = &ast.ExprBinarySmallerOrEqual{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr '>' expr
            {
                $$ = &ast.ExprBinaryGreater{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_IS_GREATER_OR_EQUAL expr
            {
                $$ = &ast.ExprBinaryGreaterOrEqual{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_SPACESHIP expr
            {
                $$ = &ast.ExprBinarySpaceship{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   expr T_INSTANCEOF class_name_reference
            {
                $$ = &ast.ExprInstanceOf{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Expr:          $1,
                    InstanceOfTkn: $2,
                    Class:         $3,
                }
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
            }
    |   new_expr
            {
                $$ = $1
            }
    |   expr '?' expr ':' expr
            {
                $$ = &ast.ExprTernary{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $5),
                    },
                    Condition:   $1,
                    QuestionTkn: $2,
                    IfTrue:      $3,
                    ColonTkn:    $4,
                    IfFalse:     $5,
                }
            }
    |   expr '?' ':' expr
            {
                $$ = &ast.ExprTernary{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $4),
                    },
                    Condition:   $1,
                    QuestionTkn: $2,
                    ColonTkn:    $3,
                    IfFalse:     $4,
                }
            }
    |   expr T_COALESCE expr
            {
                $$ = &ast.ExprBinaryCoalesce{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Left:  $1,
                    OpTkn: $2,
                    Right: $3,
                }
            }
    |   internal_functions_in_yacc
            {
                $$ = $1
            }
    |   T_INT_CAST expr
            {
                $$ = &ast.ExprCastInt{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    CastTkn: $1,
                    Expr:    $2,
                }
            }
    |   T_DOUBLE_CAST expr
            {
                $$ = &ast.ExprCastDouble{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    CastTkn: $1,
                    Expr:    $2,
                }
            }
    |   T_STRING_CAST expr
            {
                $$ = &ast.ExprCastString{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    CastTkn: $1,
                    Expr:    $2,
                }
            }
    |   T_ARRAY_CAST expr
            {
                $$ = &ast.ExprCastArray{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    CastTkn: $1,
                    Expr:    $2,
                }
            }
    |   T_OBJECT_CAST expr
            {
                $$ = &ast.ExprCastObject{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    CastTkn: $1,
                    Expr:    $2,
                }
            }
    |   T_BOOL_CAST expr
            {
                $$ = &ast.ExprCastBool{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    CastTkn: $1,
                    Expr:    $2,
                }
            }
    |   T_UNSET_CAST expr
            {
                $$ = &ast.ExprCastUnset{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    CastTkn: $1,
                    Expr:    $2,
                }
            }
    |   T_EXIT exit_expr
            {
                exit := &ast.ExprExit{
                    DieTkn: $1,
                }

                if $2 == nil {
                    exit.Node.Position = position.NewTokenPosition($1)
                } else {
                    exit.Node.Position       = position.NewTokenNodePosition($1, $2)
                    exit.OpenParenthesisTkn  = $2.(*ast.ParserBrackets).OpenBracketTkn
                    exit.Expr                = $2.(*ast.ParserBrackets).Child
                    exit.CloseParenthesisTkn = $2.(*ast.ParserBrackets).CloseBracketTkn
                }

                $$ = exit
            }
    |   '@' expr
            {
                $$ = &ast.ExprErrorSuppress{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    AtTkn: $1,
                    Expr:  $2,
                }
            }
    |   scalar
            {
                $$ = $1
            }
    |   '`' backticks_expr '`'
            {
                $$ = &ast.ExprShellExec{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenBacktickTkn:  $1,
                    Parts:            $2,
                    CloseBacktickTkn: $3,
                }
            }
    |   T_PRINT expr
            {
                $$ = &ast.ExprPrint{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    PrintTkn: $1,
                    Expr:     $2,
                }
            }
    |   T_YIELD
            {
                $$ = &ast.ExprYield{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    YieldTkn: $1,
                }
            }
    |   T_YIELD expr
            {
                $$ = &ast.ExprYield{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    YieldTkn: $1,
                    Value:    $2,
                }
            }
    |   T_YIELD expr T_DOUBLE_ARROW expr
            {
                $$ = &ast.ExprYield{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $4),
                    },
                    YieldTkn:       $1,
                    Key:            $2,
                    DoubleArrowTkn: $3,
                    Value:          $4,
                }
            }
    |   T_YIELD_FROM expr
            {
                $$ = &ast.ExprYieldFrom{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    YieldFromTkn: $1,
                    Expr:         $2,
                }
            }
    |   inline_function
            {
                $$ = $1;
            }
    |   T_STATIC inline_function
            {
                switch n := $2.(type) {
                case *ast.ExprClosure :
                    n.Position = position.NewTokenNodePosition($1, $2)
                    n.StaticTkn = $1;
                case *ast.ExprArrowFunction :
                    n.Position = position.NewTokenNodePosition($1, $2)
                    n.StaticTkn = $1;
                };

                $$ = $2
            }
;

inline_function:
        T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                $$ = &ast.ExprClosure{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $11),
                    },
                    FunctionTkn:          $1,
                    AmpersandTkn:         $2,
                    OpenParenthesisTkn:   $4,
                    Params:               $5.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:        $5.(*ast.ParserSeparatedList).SeparatorTkns,
                    CloseParenthesisTkn:  $6,
                    ClosureUse:           $7,
                    ColonTkn:             $8.(*ast.ReturnType).ColonTkn,
                    ReturnType:           $8.(*ast.ReturnType).Type,
                    OpenCurlyBracketTkn:  $9,
                    Stmts:                $10,
                    CloseCurlyBracketTkn: $11,
                }
            }
    |   T_FN returns_ref '(' parameter_list ')' return_type backup_doc_comment T_DOUBLE_ARROW expr
            {
                $$ = &ast.ExprArrowFunction{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $9),
                    },
                    FnTkn:               $1,
                    AmpersandTkn:        $2,
                    OpenParenthesisTkn:  $3,
                    Params:              $4.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:       $4.(*ast.ParserSeparatedList).SeparatorTkns,
                    CloseParenthesisTkn: $5,
                    ColonTkn:            $6.(*ast.ReturnType).ColonTkn,
                    ReturnType:          $6.(*ast.ReturnType).Type,
                    DoubleArrowTkn:      $8,
                    Expr:                $9,
                }
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
                $$ = &ast.ExprClosureUse{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    UseTkn:              $1,
                    OpenParenthesisTkn:  $2,
                    Uses:                $3.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:       $3.(*ast.ParserSeparatedList).SeparatorTkns,
                    CloseParenthesisTkn: $4,
                }
            }
;

lexical_var_list:
        lexical_var_list ',' lexical_var
            {
                $1.(*ast.ParserSeparatedList).SeparatorTkns = append($1.(*ast.ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ast.ParserSeparatedList).Items = append($1.(*ast.ParserSeparatedList).Items, $3)

                $$ = $1
            }
    |   lexical_var
            {
                $$ = &ast.ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
;

lexical_var:
    T_VARIABLE
            {
                $$ = &ast.ExprVariable{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    VarName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                }
            }
    |   '&' T_VARIABLE
            {
                $$ = &ast.ExprReference{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $2),
                    },
                    AmpersandTkn: $1,
                    Var: &ast.ExprVariable{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($2),
                        },
                        VarName: &ast.Identifier{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($2),
                            },
                            IdentifierTkn: $2,
                            Value:         $2.Value,
                        },
                    },
                }
            }
;

function_call:
        name argument_list
            {
                $$ = &ast.ExprFunctionCall{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $2),
                    },
                    Function:            $1,
                    OpenParenthesisTkn:  $2.(*ast.ArgumentList).OpenParenthesisTkn,
                    Arguments:           $2.(*ast.ArgumentList).Arguments,
                    SeparatorTkns:       $2.(*ast.ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $2.(*ast.ArgumentList).OpenParenthesisTkn,
                }
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
            {
                $$ = &ast.ExprStaticCall{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $4),
                    },
                    Class:               $1,
                    DoubleColonTkn:      $2,
                    Call:                $3,
                    OpenParenthesisTkn:  $4.(*ast.ArgumentList).OpenParenthesisTkn,
                    Arguments:           $4.(*ast.ArgumentList).Arguments,
                    SeparatorTkns:       $4.(*ast.ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $4.(*ast.ArgumentList).OpenParenthesisTkn,
                }
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
            {
                $$ = &ast.ExprStaticCall{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $4),
                    },
                    Class:               $1,
                    DoubleColonTkn:      $2,
                    Call:                $3,
                    OpenParenthesisTkn:  $4.(*ast.ArgumentList).OpenParenthesisTkn,
                    Arguments:           $4.(*ast.ArgumentList).Arguments,
                    SeparatorTkns:       $4.(*ast.ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $4.(*ast.ArgumentList).OpenParenthesisTkn,
                }
            }
    |   callable_expr argument_list
            {
                $$ = &ast.ExprFunctionCall{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $2),
                    },
                    Function:            $1,
                    OpenParenthesisTkn:  $2.(*ast.ArgumentList).OpenParenthesisTkn,
                    Arguments:           $2.(*ast.ArgumentList).Arguments,
                    SeparatorTkns:       $2.(*ast.ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $2.(*ast.ArgumentList).OpenParenthesisTkn,
                }
            }
;

class_name:
        T_STATIC
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
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
            }
;

backticks_expr:
        /* empty */
            {
                $$ = []ast.Vertex{}
            }
    |   T_ENCAPSED_AND_WHITESPACE
            {
                $$ = []ast.Vertex{
                    &ast.ScalarEncapsedStringPart{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        EncapsedStrTkn: $1,
                        Value:          $1.Value,
                    },
                }
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
                $$ = &ast.ExprArray{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    ArrayTkn:        $1,
                    OpenBracketTkn:  $2,
                    Items:           $3.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:   $3.(*ast.ParserSeparatedList).SeparatorTkns,
                    CloseBracketTkn: $4,
                }
            }
    |   '[' array_pair_list ']'
            {
                $$ = &ast.ExprArray{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenBracketTkn:  $1,
                    Items:           $2.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:   $2.(*ast.ParserSeparatedList).SeparatorTkns,
                    CloseBracketTkn: $3,
                }
            }
    |   T_CONSTANT_ENCAPSED_STRING
            {
                $$ = &ast.ScalarString{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    StringTkn: $1,
                    Value:     $1.Value,
                }
            }
;

scalar:
        T_LNUMBER
            {
                $$ = &ast.ScalarLnumber{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    NumberTkn: $1,
                    Value:     $1.Value,
                }
            }
    |   T_DNUMBER
            {
                $$ = &ast.ScalarDnumber{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    NumberTkn: $1,
                    Value:     $1.Value,
                }
            }
    |   T_LINE
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_FILE
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_DIR
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_TRAIT_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_METHOD_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_FUNC_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_NS_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_CLASS_C
            {
                $$ = &ast.ScalarMagicConstant{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    MagicConstTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC
            {
                $$ = &ast.ScalarHeredoc{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenHeredocTkn: $1,
                    Parts: []ast.Vertex{
                        &ast.ScalarEncapsedStringPart{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($2),
                            },
                            EncapsedStrTkn: $2,
                            Value:          $2.Value,
                        },
                    },
                    CloseHeredocTkn: $3,
                }
            }
    |   T_START_HEREDOC T_END_HEREDOC
            {
                $$ = &ast.ScalarHeredoc{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $2),
                    },
                    OpenHeredocTkn:  $1,
                    CloseHeredocTkn: $2,
                }
            }
    |   '"' encaps_list '"'
            {
                $$ = &ast.ScalarEncapsed{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenQoteTkn:  $1,
                    Parts:        $2,
                    CloseQoteTkn: $1,
                }
            }
    |   T_START_HEREDOC encaps_list T_END_HEREDOC
            {
                $$ = &ast.ScalarHeredoc{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenHeredocTkn:  $1,
                    Parts:           $2,
                    CloseHeredocTkn: $3,
                }
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
                $$ = &ast.ExprConstFetch{
                    Node: ast.Node{
                        Position: position.NewNodePosition($1),
                    },
                    Const: $1,
                }
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM identifier
            {
                $$ = &ast.ExprClassConstFetch{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $3),
                    },
                    Class:          $1,
                    DoubleColonTkn: $2,
                    ConstantName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM identifier
            {
                $$ = &ast.ExprClassConstFetch{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $3),
                    },
                    Class:          $1,
                    DoubleColonTkn: $2,
                    ConstantName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
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
                $$ = &ast.ExprArrayDimFetch{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $4),
                    },
                    Var:             $1,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
    |   constant '[' optional_expr ']'
            {
                $$ = &ast.ExprArrayDimFetch{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $4),
                    },
                    Var:             $1,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
    |   dereferencable '{' expr '}'
            {
                $$ = &ast.ExprArrayDimFetch{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $4),
                    },
                    Var:             $1,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
    |   dereferencable T_OBJECT_OPERATOR property_name argument_list
            {
                $$ = &ast.ExprMethodCall{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $4),
                    },
                    Var:                 $1,
                    ObjectOperatorTkn:   $2,
                    Method:              $3,
                    OpenParenthesisTkn:  $4.(*ast.ArgumentList).OpenParenthesisTkn,
                    Arguments:           $4.(*ast.ArgumentList).Arguments,
                    SeparatorTkns:       $4.(*ast.ArgumentList).SeparatorTkns,
                    CloseParenthesisTkn: $4.(*ast.ArgumentList).CloseParenthesisTkn,
                }
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
                $$ = &ast.ExprPropertyFetch{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:               $1,
                    ObjectOperatorTkn: $2,
                    Property:          $3,
                }
            }
;

simple_variable:
        T_VARIABLE
            {
                $$ = &ast.ExprVariable{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    VarName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                }
            }
    |   '$' '{' expr '}'
            {
                $$ = &ast.ExprVariable{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    DollarTkn: $1,
                    VarName: &ast.ParserBrackets{
                        Node: ast.Node{
                            Position: position.NewTokensPosition($2, $4),
                        },
                        OpenBracketTkn:  $2,
                        Child:           $3,
                        CloseBracketTkn: $4,
                    },
                }
            }
    |   '$' simple_variable
            {
                $$ = &ast.ExprVariable{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    DollarTkn: $1,
                    VarName:   $2,
                }
            }
;

static_member:
        class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = &ast.ExprStaticPropertyFetch{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Class:          $1,
                    DoubleColonTkn: $2,
                    Property:       $3,
                }
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = &ast.ExprStaticPropertyFetch{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Class:          $1,
                    DoubleColonTkn: $2,
                    Property:       $3,
                }
            }
;

new_variable:
        simple_variable
            {
                $$ = $1
            }
    |   new_variable '[' optional_expr ']'
            {
                $$ = &ast.ExprArrayDimFetch{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $4),
                    },
                    Var:             $1,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
    |   new_variable '{' expr '}'
            {
                $$ = &ast.ExprArrayDimFetch{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $4),
                    },
                    Var:             $1,
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
    |   new_variable T_OBJECT_OPERATOR property_name
            {
                $$ = &ast.ExprPropertyFetch{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Var:               $1,
                    ObjectOperatorTkn: $2,
                    Property:          $3,
                }
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = &ast.ExprStaticPropertyFetch{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Class:          $1,
                    DoubleColonTkn: $2,
                    Property:       $3,
                }
            }
    |   new_variable T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = &ast.ExprStaticPropertyFetch{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Class:          $1,
                    DoubleColonTkn: $2,
                    Property:       $3,
                }
            }
;

member_name:
        identifier
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   '{' expr '}'
            {
                $$ = &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenBracketTkn:  $1,
                    Child:           $2,
                    CloseBracketTkn: $3,
                }
            }
    |   simple_variable
            {
                $$ = $1
            }
;

property_name:
        T_STRING
            {
                $$ = &ast.Identifier{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    IdentifierTkn: $1,
                    Value:         $1.Value,
                }
            }
    |   '{' expr '}'
            {
                $$ = &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenBracketTkn:  $1,
                    Child:           $2,
                    CloseBracketTkn: $3,
                }
            }
    |   simple_variable
            {
                $$ = $1
            }
;

array_pair_list:
        non_empty_array_pair_list
            {
                pairList := $1.(*ast.ParserSeparatedList)
                fistPair := pairList.Items[0].(*ast.ExprArrayItem)

                if fistPair.Key == nil && fistPair.Val == nil && len(pairList.Items) == 1 {
                    pairList.Items = nil
                }

                $$ = $1
            }
;

possible_array_pair:
        /* empty */
            {
                $$ = &ast.ExprArrayItem{}
            }
    |   array_pair
            {
                $$ = $1
            }
;

non_empty_array_pair_list:
        non_empty_array_pair_list ',' possible_array_pair
            {
                $1.(*ast.ParserSeparatedList).SeparatorTkns = append($1.(*ast.ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ast.ParserSeparatedList).Items = append($1.(*ast.ParserSeparatedList).Items, $3)

                $$ = $1
            }
    |   possible_array_pair
            {
                $$ = &ast.ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
;

array_pair:
        expr T_DOUBLE_ARROW expr
            {
                $$ = &ast.ExprArrayItem{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $3),
                    },
                    Key:            $1,
                    DoubleArrowTkn: $2,
                    Val:            $3,
                }
            }
    |   expr
            {
                $$ = &ast.ExprArrayItem{
                    Node: ast.Node{
                        Position: position.NewNodePosition($1),
                    },
                    Val: $1,
                }
            }
    |   expr T_DOUBLE_ARROW '&' variable
            {
                $$ = &ast.ExprArrayItem{
                    Node: ast.Node{
                        Position: position.NewNodesPosition($1, $4),
                    },
                    Key:            $1,
                    DoubleArrowTkn: $2,
                    Val: &ast.ExprReference{
                        Node: ast.Node{
                            Position: position.NewTokenNodePosition($3, $4),
                        },
                        AmpersandTkn: $3,
                        Var:          $4,
                    },
                }
            }
    |   '&' variable
            {
                $$ = &ast.ExprArrayItem{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    Val: &ast.ExprReference{
                        Node: ast.Node{
                            Position: position.NewTokenNodePosition($1, $2),
                        },
                        AmpersandTkn: $1,
                        Var:          $2,
                    },
                }
            }
    |   T_ELLIPSIS expr
            {
                $$ = &ast.ExprArrayItem{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    EllipsisTkn: $1,
                    Val:         $2,
                }
            }
    |   expr T_DOUBLE_ARROW T_LIST '(' array_pair_list ')'
            {
                $$ = &ast.ExprArrayItem{
                    Node: ast.Node{
                        Position: position.NewNodeTokenPosition($1, $6),
                    },
                    Key:            $1,
                    DoubleArrowTkn: $2,
                    Val: &ast.ExprList{
                        Node: ast.Node{
                            Position: position.NewTokensPosition($3, $6),
                        },
                        ListTkn:         $3,
                        OpenBracketTkn:  $4,
                        Items:           $5.(*ast.ParserSeparatedList).Items,
                        SeparatorTkns:   $5.(*ast.ParserSeparatedList).SeparatorTkns,
                        CloseBracketTkn: $6,
                    },
                }
            }
    |   T_LIST '(' array_pair_list ')'
            {
                $$ = &ast.ExprArrayItem{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    Val: &ast.ExprList{
                        Node: ast.Node{
                            Position: position.NewTokensPosition($1, $4),
                        },
                        ListTkn:         $1,
                        OpenBracketTkn:  $2,
                        Items:           $3.(*ast.ParserSeparatedList).Items,
                        SeparatorTkns:   $3.(*ast.ParserSeparatedList).SeparatorTkns,
                        CloseBracketTkn: $4,
                    },
                }
            }
;

encaps_list:
        encaps_list encaps_var
            {
                $$ = append($1, $2)
            }
    |   encaps_list T_ENCAPSED_AND_WHITESPACE
            {
                $$ = append(
                    $1,
                    &ast.ScalarEncapsedStringPart{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($2),
                        },
                        EncapsedStrTkn: $2,
                        Value:          $2.Value,
                    },
                )
            }
    |   encaps_var
            {
                $$ = []ast.Vertex{$1}
            }
    |   T_ENCAPSED_AND_WHITESPACE encaps_var
            {
                $$ = []ast.Vertex{
                    &ast.ScalarEncapsedStringPart{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        EncapsedStrTkn: $1,
                        Value:          $1.Value,
                    },
                    $2,
                }
            }
;

encaps_var:
        T_VARIABLE
            {
                $$ = &ast.ExprVariable{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    VarName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                }
            }
    |   T_VARIABLE '[' encaps_var_offset ']'
            {
                $$ = &ast.ExprArrayDimFetch{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    Var: &ast.ExprVariable{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        VarName: &ast.Identifier{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($1),
                            },
                            IdentifierTkn: $1,
                            Value:         $1.Value,
                        },
                    },
                    OpenBracketTkn:  $2,
                    Dim:             $3,
                    CloseBracketTkn: $4,
                }
            }
    |   T_VARIABLE T_OBJECT_OPERATOR T_STRING
            {
                $$ = &ast.ExprPropertyFetch{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    Var: &ast.ExprVariable{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        VarName: &ast.Identifier{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($1),
                            },
                            IdentifierTkn: $1,
                            Value:         $1.Value,
                        },
                    },
                    ObjectOperatorTkn: $2,
                    Property: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($3),
                        },
                        IdentifierTkn: $3,
                        Value:         $3.Value,
                    },
                }
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES expr '}'
            {
                $$ = &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenBracketTkn: $1,
                    Child: &ast.ExprVariable{
                        Node: ast.Node{
                            Position: position.NewNodePosition($2),
                        },
                        VarName: $2,
                    },
                    CloseBracketTkn: $3,
                }
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}'
            {
                $$ = &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenBracketTkn: $1,
                    Child: &ast.ExprVariable{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($2),
                        },
                        VarName: &ast.Identifier{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($2),
                            },
                            IdentifierTkn: $2,
                            Value:         $2.Value,
                        },
                    },
                    CloseBracketTkn: $3,
                }
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
            {
                $$ = &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $6),
                    },
                    OpenBracketTkn: $1,
                    Child: &ast.ExprArrayDimFetch{
                        Node: ast.Node{
                            Position: position.NewTokensPosition($2, $5),
                        },
                        Var: &ast.ExprVariable{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($2),
                            },
                            VarName: &ast.Identifier{
                                Node: ast.Node{
                                    Position: position.NewTokenPosition($2),
                                },
                                IdentifierTkn: $2,
                                Value:         $2.Value,
                            },
                        },
                        OpenBracketTkn:       $3,
                        Dim:                  $4,
                        CloseBracketTkn:      $5,
                    },
                    CloseBracketTkn: $6,
                }
            }
    |   T_CURLY_OPEN variable '}'
            {
                $$ = &ast.ParserBrackets{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $3),
                    },
                    OpenBracketTkn:  $1,
                    Child:           $2,
                    CloseBracketTkn: $3,
                }
            }
;

encaps_var_offset:
        T_STRING
            {
                $$ = &ast.ScalarString{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    StringTkn: $1,
                    Value:     $1.Value,
                }
            }
    |   T_NUM_STRING
            {
                // TODO: add option to handle 64 bit integer
                if _, err := strconv.Atoi(string($1.Value)); err == nil {
                    $$ = &ast.ScalarLnumber{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        NumberTkn: $1,
                        Value:     $1.Value,
                    }
                } else {
                    $$ = &ast.ScalarString{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        StringTkn: $1,
                        Value:     $1.Value,
                    }
                }
            }
    |   '-' T_NUM_STRING
            {
                _, err := strconv.Atoi(string($2.Value));
                isInt := err == nil

                if isInt {
                    $$ = &ast.ExprUnaryMinus{
                        Node: ast.Node{
                            Position: position.NewTokensPosition($1, $2),
                        },
                        MinusTkn: $1,
                        Expr: &ast.ScalarLnumber{
                            Node: ast.Node{
                                Position: position.NewTokenPosition($2),
                            },
                            NumberTkn: $2,
                            Value:     $2.Value,
                        },
                    }
                    $$.GetNode().Position = position.NewTokensPosition($1, $2)
                } else {
                    $$ = &ast.ScalarString{
                        Node: ast.Node{
                            Position: position.NewTokensPosition($1, $2),
                        },
                        MinusTkn:  $1,
                        StringTkn: $2,
                        Value:     append([]byte("-"), $2.Value...),
                    }
                }
            }
    |   T_VARIABLE
            {
                $$ = &ast.ExprVariable{
                    Node: ast.Node{
                        Position: position.NewTokenPosition($1),
                    },
                    VarName: &ast.Identifier{
                        Node: ast.Node{
                            Position: position.NewTokenPosition($1),
                        },
                        IdentifierTkn: $1,
                        Value:         $1.Value,
                    },
                }
            }
;

internal_functions_in_yacc:
        T_ISSET '(' isset_variables possible_comma ')'
            {
                if $4 != nil {
                    $3.(*ast.ParserSeparatedList).SeparatorTkns = append($3.(*ast.ParserSeparatedList).SeparatorTkns, $4)
                }

                $$ = &ast.ExprIsset{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $5),
                    },
                    IssetTkn:            $1,
                    OpenParenthesisTkn:  $2,
                    Vars:                $3.(*ast.ParserSeparatedList).Items,
                    SeparatorTkns:       $3.(*ast.ParserSeparatedList).SeparatorTkns,
                    CloseParenthesisTkn: $5,
                }
            }
    |   T_EMPTY '(' expr ')'
            {
                $$ = &ast.ExprEmpty{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    EmptyTkn:            $1,
                    OpenParenthesisTkn:  $2,
                    Expr:                $3,
                    CloseParenthesisTkn: $4,
                }
            }
    |   T_INCLUDE expr
            {
                $$ = &ast.ExprInclude{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    IncludeTkn: $1,
                    Expr:       $2,
                }
            }
    |   T_INCLUDE_ONCE expr
            {
                $$ = &ast.ExprIncludeOnce{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    IncludeTkn: $1,
                    Expr:       $2,
                }
            }
    |   T_EVAL '(' expr ')'
            {
                $$ = &ast.ExprEval{
                    Node: ast.Node{
                        Position: position.NewTokensPosition($1, $4),
                    },
                    EvalTkn:             $1,
                    OpenParenthesisTkn:  $2,
                    Expr:                $3,
                    CloseParenthesisTkn: $4,
                }
            }
    |   T_REQUIRE expr
            {
                $$ = &ast.ExprRequire{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    RequireTkn: $1,
                    Expr:       $2,
                }
            }
    |   T_REQUIRE_ONCE expr
            {
                $$ = &ast.ExprRequireOnce{
                    Node: ast.Node{
                        Position: position.NewTokenNodePosition($1, $2),
                    },
                    RequireOnceTkn: $1,
                    Expr:           $2,
                }
            }
;

isset_variables:
        isset_variable
            {
                $$ = &ast.ParserSeparatedList{
                    Items: []ast.Vertex{$1},
                }
            }
    |   isset_variables ',' isset_variable
            {
                $1.(*ast.ParserSeparatedList).SeparatorTkns = append($1.(*ast.ParserSeparatedList).SeparatorTkns, $2)
                $1.(*ast.ParserSeparatedList).Items = append($1.(*ast.ParserSeparatedList).Items, $3)

                $$ = $1
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
