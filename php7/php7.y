%{
package php7

import (
    "fmt"
    "strings"
    "strconv"

    "github.com/z7zmey/php-parser/token"
    "github.com/z7zmey/php-parser/node"
    "github.com/z7zmey/php-parser/node/scalar"
    "github.com/z7zmey/php-parser/node/name"
    "github.com/z7zmey/php-parser/node/stmt"
    "github.com/z7zmey/php-parser/node/expr"
    "github.com/z7zmey/php-parser/node/expr/assign_op"
    "github.com/z7zmey/php-parser/node/expr/binary_op"
    "github.com/z7zmey/php-parser/node/expr/cast"
)

%}

%union{
    node node.Node
    token token.Token
    boolWithToken boolWithToken
    list []node.Node
    foreachVariable foreachVariable
    nodesWithEndToken *nodesWithEndToken
    str string
}

%type <token> $unk
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

%left T_INCLUDE T_INCLUDE_ONCE T_EVAL T_REQUIRE T_REQUIRE_ONCE
%left ','
%left T_LOGICAL_OR
%left T_LOGICAL_XOR
%left T_LOGICAL_AND
%right T_PRINT
%right T_YIELD
%right T_DOUBLE_ARROW
%right T_YIELD_FROM
%left '=' T_PLUS_EQUAL T_MINUS_EQUAL T_MUL_EQUAL T_DIV_EQUAL T_CONCAT_EQUAL T_MOD_EQUAL T_AND_EQUAL T_OR_EQUAL T_XOR_EQUAL T_SL_EQUAL T_SR_EQUAL T_POW_EQUAL
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

%type <boolWithToken> is_reference is_variadic returns_ref

%type <token> reserved_non_modifiers
%type <token> semi_reserved
%type <token> identifier

%type <node> top_statement name statement function_declaration_statement
%type <node> class_declaration_statement trait_declaration_statement
%type <node> interface_declaration_statement
%type <node> group_use_declaration inline_use_declaration
%type <node> mixed_group_use_declaration use_declaration unprefixed_use_declaration
%type <node> const_decl inner_statement
%type <node> expr optional_expr while_statement for_statement 
%type <node> foreach_statement declare_statement finally_statement unset_variable variable
%type <node> extends_from parameter optional_type argument expr_without_variable global_var
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

%type <node> member_modifier
%type <node> use_type
%type <foreachVariable> foreach_variable

%type <nodesWithEndToken> method_body switch_case_list trait_adaptations argument_list ctor_arguments

%type <list> encaps_list backticks_expr namespace_name catch_name_list catch_list class_const_list
%type <list> const_list echo_expr_list for_exprs non_empty_for_exprs global_var_list
%type <list> unprefixed_use_declarations inline_use_declarations property_list static_var_list
%type <list> case_list trait_adaptation_list unset_variables
%type <list> use_declarations lexical_var_list lexical_vars isset_variables non_empty_array_pair_list
%type <list> array_pair_list non_empty_argument_list top_statement_list
%type <list> inner_statement_list parameter_list non_empty_parameter_list class_statement_list
%type <list> interface_extends_list implements_list method_modifiers variable_modifiers
%type <list> non_empty_member_modifiers name_list class_modifiers

%type <str> backup_doc_comment

%%

/////////////////////////////////////////////////////////////////////////

start:
    top_statement_list
        {
            rootnode = stmt.NewStmtList($1)
            positions.AddPosition(rootnode, positionBuilder.NewNodeListPosition($1))
        }
;

reserved_non_modifiers:
      T_INCLUDE {$$=$1} | T_INCLUDE_ONCE {$$=$1} | T_EVAL {$$=$1} | T_REQUIRE {$$=$1} | T_REQUIRE_ONCE {$$=$1} | T_LOGICAL_OR {$$=$1} | T_LOGICAL_XOR {$$=$1} | T_LOGICAL_AND {$$=$1} 
    | T_INSTANCEOF {$$=$1} | T_NEW {$$=$1} | T_CLONE {$$=$1} | T_EXIT {$$=$1} | T_IF {$$=$1} | T_ELSEIF {$$=$1} | T_ELSE {$$=$1} | T_ENDIF {$$=$1} | T_ECHO {$$=$1} | T_DO {$$=$1} | T_WHILE {$$=$1} | T_ENDWHILE {$$=$1} 
    | T_FOR {$$=$1} | T_ENDFOR {$$=$1} | T_FOREACH {$$=$1} | T_ENDFOREACH {$$=$1} | T_DECLARE {$$=$1} | T_ENDDECLARE {$$=$1} | T_AS {$$=$1} | T_TRY {$$=$1} | T_CATCH {$$=$1} | T_FINALLY {$$=$1} 
    | T_THROW {$$=$1} | T_USE {$$=$1} | T_INSTEADOF {$$=$1} | T_GLOBAL {$$=$1} | T_VAR {$$=$1} | T_UNSET {$$=$1} | T_ISSET {$$=$1} | T_EMPTY {$$=$1} | T_CONTINUE {$$=$1} | T_GOTO {$$=$1} 
    | T_FUNCTION {$$=$1} | T_CONST {$$=$1} | T_RETURN {$$=$1} | T_PRINT {$$=$1} | T_YIELD {$$=$1} | T_LIST {$$=$1} | T_SWITCH {$$=$1} | T_ENDSWITCH {$$=$1} | T_CASE {$$=$1} | T_DEFAULT {$$=$1} | T_BREAK {$$=$1} 
    | T_ARRAY {$$=$1} | T_CALLABLE {$$=$1} | T_EXTENDS {$$=$1} | T_IMPLEMENTS {$$=$1} | T_NAMESPACE {$$=$1} | T_TRAIT {$$=$1} | T_INTERFACE {$$=$1} | T_CLASS {$$=$1} 
    | T_CLASS_C {$$=$1} | T_TRAIT_C {$$=$1} | T_FUNC_C {$$=$1} | T_METHOD_C {$$=$1} | T_LINE {$$=$1} | T_FILE {$$=$1} | T_DIR {$$=$1} | T_NS_C {$$=$1} 
;

semi_reserved:
        reserved_non_modifiers {$$=$1}
    | T_STATIC {$$=$1} | T_ABSTRACT {$$=$1} | T_FINAL {$$=$1} | T_PRIVATE {$$=$1} | T_PROTECTED {$$=$1} | T_PUBLIC {$$=$1}
;

identifier:
        T_STRING                                        { fmt.Println("1"); $$ = $1 }
    |   semi_reserved                                   { fmt.Println("2"); $$ = $1 }
;

top_statement_list:
        top_statement_list top_statement                { fmt.Println("3"); $$ = append($1, $2) }
    |   /* empty */                                     { fmt.Println("4"); $$ = []node.Node{} }
;

namespace_name:
    T_STRING
        {
            namePart := name.NewNamePart($1.Value)
            positions.AddPosition(namePart, positionBuilder.NewTokenPosition($1))
            fmt.Println("5"); $$ = []node.Node{namePart}
            comments.AddComments(namePart, $1.Comments())
        }
    |   namespace_name T_NS_SEPARATOR T_STRING
        {
            namePart := name.NewNamePart($3.Value)
            positions.AddPosition(namePart, positionBuilder.NewTokenPosition($3))
            fmt.Println("6"); $$ = append($1, namePart)
            comments.AddComments(namePart, $3.Comments())
        }
;

name:
    namespace_name
        {
            fmt.Println("7"); $$ = name.NewName($1)
            positions.AddPosition($$, positionBuilder.NewNodeListPosition($1))
            comments.AddComments($$, ListGetFirstNodeComments($1))
        }
    | T_NAMESPACE T_NS_SEPARATOR namespace_name
        {
            fmt.Println("8"); $$ = name.NewRelative($3)
            positions.AddPosition($$, positionBuilder.NewTokenNodeListPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    | T_NS_SEPARATOR namespace_name
        {
            fmt.Println("9"); $$ = name.NewFullyQualified($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodeListPosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
;

top_statement:
        statement                                       { fmt.Println("10"); $$ = $1; }
    |   function_declaration_statement                  { fmt.Println("11"); $$ = $1; }
    |   class_declaration_statement                     { fmt.Println("12"); $$ = $1; }
    |   trait_declaration_statement                     { fmt.Println("13"); $$ = $1; }
    |   interface_declaration_statement                 { fmt.Println("14"); $$ = $1; }
    |   T_HALT_COMPILER '(' ')' ';'                     { fmt.Println("15"); $$ = stmt.NewHaltCompiler() }
    |   T_NAMESPACE namespace_name ';'
        {
            name := name.NewName($2)
            positions.AddPosition(name, positionBuilder.NewNodeListPosition($2))
            fmt.Println("16"); $$ = stmt.NewNamespace(name, nil)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))

            comments.AddComments(name, ListGetFirstNodeComments($2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_NAMESPACE namespace_name '{' top_statement_list '}'
        {
            name := name.NewName($2)
            positions.AddPosition(name, positionBuilder.NewNodeListPosition($2))
            fmt.Println("17"); $$ = stmt.NewNamespace(name, $4)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $5))

            comments.AddComments(name, ListGetFirstNodeComments($2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_NAMESPACE '{' top_statement_list '}'
        {
            fmt.Println("18"); $$ = stmt.NewNamespace(nil, $3)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
            comments.AddComments($$, $1.Comments())
        }
    |   T_USE mixed_group_use_declaration ';'           { fmt.Println("19"); $$ = $2 }
    |   T_USE use_type group_use_declaration ';'        { fmt.Println("20"); $$ = $3.(*stmt.GroupUse).SetUseType($2) }
    |   T_USE use_declarations ';'
        {
            fmt.Println("21"); $$ = stmt.NewUseList(nil, $2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    |   T_USE use_type use_declarations ';'             { fmt.Println("22"); $$ = stmt.NewUseList($2, $3) }
    |   T_CONST const_list ';'
        {
            fmt.Println("23"); $$ = stmt.NewConstList($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
;

use_type:
    T_FUNCTION
        {
            fmt.Println("24"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_CONST
        {
            fmt.Println("25"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
;

group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                name := name.NewName($1)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($1))
                fmt.Println("26"); $$ = stmt.NewGroupUse(nil, name, $4)
                positions.AddPosition($$, positionBuilder.NewNodeListTokenPosition($1, $6))

                comments.AddComments(name, ListGetFirstNodeComments($1))
                comments.AddComments($$, ListGetFirstNodeComments($1))
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                name := name.NewName($2)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($2))
                fmt.Println("27"); $$ = stmt.NewGroupUse(nil, name, $5)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $7))

                comments.AddComments(name, ListGetFirstNodeComments($2))
                comments.AddComments($$, $1.Comments())
            }
;

mixed_group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                name := name.NewName($1)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($1))
                fmt.Println("28"); $$ = stmt.NewGroupUse(nil, name, $4)
                positions.AddPosition($$, positionBuilder.NewNodeListTokenPosition($1, $6))

                comments.AddComments(name, ListGetFirstNodeComments($1))
                comments.AddComments($$, ListGetFirstNodeComments($1))
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                name := name.NewName($2)
                positions.AddPosition(name, positionBuilder.NewNodeListPosition($2))
                fmt.Println("29"); $$ = stmt.NewGroupUse(nil, name, $5)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $7))

                comments.AddComments(name, ListGetFirstNodeComments($2))
                comments.AddComments($$, $1.Comments())
            }
;

possible_comma:
        /* empty */
    |   ','
;

inline_use_declarations:
        inline_use_declarations ',' inline_use_declaration
                                                        { fmt.Println("30"); $$ = append($1, $3) }
    |   inline_use_declaration                          { fmt.Println("31"); $$ = []node.Node{$1} }
;

unprefixed_use_declarations:
        unprefixed_use_declarations ',' unprefixed_use_declaration
                                                        { fmt.Println("32"); $$ = append($1, $3) }
    |   unprefixed_use_declaration                      { fmt.Println("33"); $$ = []node.Node{$1} }
;

use_declarations:
        use_declarations ',' use_declaration            { fmt.Println("34"); $$ = append($1, $3) }
    |   use_declaration                                 { fmt.Println("35"); $$ = []node.Node{$1} }
;

inline_use_declaration:
        unprefixed_use_declaration                      { fmt.Println("36"); $$ = $1; }
    |   use_type unprefixed_use_declaration             { fmt.Println("37"); $$ = $2.(*stmt.Use).SetUseType($1) }
;

unprefixed_use_declaration:
    namespace_name
        {
            name := name.NewName($1)
            positions.AddPosition(name, positionBuilder.NewNodeListPosition($1))
            fmt.Println("38"); $$ = stmt.NewUse(nil, name, nil)
            positions.AddPosition($$, positionBuilder.NewNodeListPosition($1))

            comments.AddComments(name, ListGetFirstNodeComments($1))
            comments.AddComments($$, ListGetFirstNodeComments($1))
        }
    |   namespace_name T_AS T_STRING
        {
            name := name.NewName($1)
            positions.AddPosition(name, positionBuilder.NewNodeListPosition($1))
            alias := node.NewIdentifier($3.Value)
            positions.AddPosition(alias, positionBuilder.NewTokenPosition($3))
            fmt.Println("39"); $$ = stmt.NewUse(nil, name, alias)
            positions.AddPosition($$, positionBuilder.NewNodeListTokenPosition($1, $3))

            comments.AddComments(name, ListGetFirstNodeComments($1))
            comments.AddComments(alias, $3.Comments())
            comments.AddComments($$, ListGetFirstNodeComments($1))
        }
;

use_declaration:
        unprefixed_use_declaration                      { fmt.Println("40"); $$ = $1; }
    |   T_NS_SEPARATOR unprefixed_use_declaration       { fmt.Println("41"); $$ = $2; }
;

const_list:
        const_list ',' const_decl                       { fmt.Println("42"); $$ = append($1, $3) }
    |   const_decl                                      { fmt.Println("43"); $$ = []node.Node{$1} }
;

inner_statement_list:
        inner_statement_list inner_statement            { fmt.Println("44"); $$ = append($1, $2) }
    |   /* empty */                                     { fmt.Println("45"); $$ = []node.Node{} }
;

inner_statement:
    statement                                           { fmt.Println("46"); $$ = $1; }
    |   function_declaration_statement                  { fmt.Println("47"); $$ = $1; }
    |   class_declaration_statement                     { fmt.Println("48"); $$ = $1; }
    |   trait_declaration_statement                     { fmt.Println("49"); $$ = $1; }
    |   interface_declaration_statement                 { fmt.Println("50"); $$ = $1; }
    |   T_HALT_COMPILER '(' ')' ';'
        {
            fmt.Println("51"); $$ = stmt.NewHaltCompiler()
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
            comments.AddComments($$, $1.Comments())
        }

statement:
    '{' inner_statement_list '}'
        {
            fmt.Println("52"); $$ = stmt.NewStmtList($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    |   if_stmt                                         { fmt.Println("53"); $$ = $1; }
    |   alt_if_stmt                                     { fmt.Println("54"); $$ = $1; }
    |   T_WHILE '(' expr ')' while_statement
        {
            fmt.Println("55"); $$ = stmt.NewWhile($3, $5)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $5))
            comments.AddComments($$, $1.Comments())
        }
    |   T_DO statement T_WHILE '(' expr ')' ';'
        {
            fmt.Println("56"); $$ = stmt.NewDo($2, $5)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $7))
            comments.AddComments($$, $1.Comments())
        }
    |   T_FOR '(' for_exprs ';' for_exprs ';' for_exprs ')' for_statement
        {
            fmt.Println("57"); $$ = stmt.NewFor($3, $5, $7, $9)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $9))
            comments.AddComments($$, $1.Comments())
        }
    |   T_SWITCH '(' expr ')' switch_case_list
        {
            fmt.Println("58"); $$ = stmt.NewSwitch($3, $5.nodes)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $5.endToken))
            comments.AddComments($$, $1.Comments())
        }
    |   T_BREAK optional_expr ';'
        {
            fmt.Println("59"); $$ = stmt.NewBreak($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    |   T_CONTINUE optional_expr ';'
        {
            fmt.Println("60"); $$ = stmt.NewContinue($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    |   T_RETURN optional_expr ';'
        {
            fmt.Println("61"); $$ = stmt.NewReturn($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    |   T_GLOBAL global_var_list ';'
        {
            fmt.Println("62"); $$ = stmt.NewGlobal($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    |   T_STATIC static_var_list ';'
        {
            fmt.Println("63"); $$ = stmt.NewStatic($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    |   T_ECHO echo_expr_list ';'
        {
            fmt.Println("64"); $$ = stmt.NewEcho($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    |   T_INLINE_HTML
        {
            fmt.Println("65"); $$ = stmt.NewInlineHtml($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   expr ';'
        {
            fmt.Println("66"); $$ = stmt.NewExpression($1)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $2))
            comments.AddComments($$, comments[$1])
        }
    |   T_UNSET '(' unset_variables possible_comma ')' ';' 
        {
            fmt.Println("67"); $$ = stmt.NewUnset($3)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $6))
            comments.AddComments($$, $1.Comments())
        }
    |   T_FOREACH '(' expr T_AS foreach_variable ')' foreach_statement
        {
            fmt.Println("68"); $$ = stmt.NewForeach($3, nil, $5.node, $7, $5.byRef)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $7))
            comments.AddComments($$, $1.Comments())
        }
    |   T_FOREACH '(' expr T_AS variable T_DOUBLE_ARROW foreach_variable ')' foreach_statement
        {
            fmt.Println("69"); $$ = stmt.NewForeach($3, $5, $7.node, $9, $7.byRef)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $9))
            comments.AddComments($$, $1.Comments())
        }
    |   T_DECLARE '(' const_list ')' declare_statement
        {
            fmt.Println("70"); $$ = stmt.NewDeclare($3, $5)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $5))
            comments.AddComments($$, $1.Comments())
        }
    |   ';'
        {
            fmt.Println("71"); $$ = stmt.NewNop()
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_TRY '{' inner_statement_list '}' catch_list finally_statement
            {
                if $6 == nil {
                    fmt.Println("72"); $$ = stmt.NewTry($3, $5, $6)
                    positions.AddPosition($$, positionBuilder.NewTokenNodeListPosition($1, $5))
                } else {
                    fmt.Println("73"); $$ = stmt.NewTry($3, $5, $6)
                    positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $6))
                }

                comments.AddComments($$, $1.Comments())
            }
    |   T_THROW expr ';'
        {
            fmt.Println("74"); $$ = stmt.NewThrow($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    |   T_GOTO T_STRING ';'
        {
            label := node.NewIdentifier($2.Value)
            positions.AddPosition(label, positionBuilder.NewTokenPosition($2))
            fmt.Println("75"); $$ = stmt.NewGoto(label)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))

            comments.AddComments(label, $2.Comments())
            comments.AddComments($$, $1.Comments())
        }
    |   T_STRING ':'
        {
            label := node.NewIdentifier($1.Value)
            positions.AddPosition(label, positionBuilder.NewTokenPosition($1))
            fmt.Println("76"); $$ = stmt.NewLabel(label)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $2))

            comments.AddComments(label, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }

catch_list:
        /* empty */                                     { fmt.Println("77"); $$ = []node.Node{} }
    |   catch_list T_CATCH '(' catch_name_list T_VARIABLE ')' '{' inner_statement_list '}'
        {
            identifier := node.NewIdentifier($5.Value)
            positions.AddPosition(identifier, positionBuilder.NewTokenPosition($5))
            variable := expr.NewVariable(identifier)
            positions.AddPosition(variable, positionBuilder.NewTokenPosition($5))
            catch := stmt.NewCatch($4, variable, $8)
            positions.AddPosition(catch, positionBuilder.NewTokensPosition($2, $9))
            fmt.Println("78"); $$ = append($1, catch)

            comments.AddComments(identifier, $5.Comments())
            comments.AddComments(variable, $5.Comments())
            comments.AddComments(catch, $2.Comments())
        }
;
catch_name_list:
        name                                            { fmt.Println("79"); $$ = []node.Node{$1} }
    |   catch_name_list '|' name                        { fmt.Println("80"); $$ = append($1, $3) }
;

finally_statement:
        /* empty */                                     { fmt.Println("81"); $$ = nil }
    |   T_FINALLY '{' inner_statement_list '}'
        {
            fmt.Println("82"); $$ = stmt.NewFinally($3)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
            comments.AddComments($$, $1.Comments())
        }
;

unset_variables:
        unset_variable                                  { fmt.Println("83"); $$ = []node.Node{$1} }
    |   unset_variables ',' unset_variable              { fmt.Println("84"); $$ = append($1, $3) }
;

unset_variable:
    variable                                            { fmt.Println("85"); $$ = $1 }
;

function_declaration_statement:
    T_FUNCTION returns_ref T_STRING backup_doc_comment '(' parameter_list ')' return_type '{' inner_statement_list '}'
        {
            name := node.NewIdentifier($3.Value)
            positions.AddPosition(name, positionBuilder.NewTokenPosition($3))
            fmt.Println("86"); $$ = stmt.NewFunction(name, $2.value, $6, $8, $10, $4)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $11))

            comments.AddComments(name, $3.Comments())
            comments.AddComments($$, $1.Comments())
        }
;

is_reference:
        /* empty */                                     { fmt.Println("87"); $$ = boolWithToken{false, nil} }
    |   '&'                                             { fmt.Println("88"); $$ = boolWithToken{true, &$1} }
;

is_variadic:
        /* empty */                                     { fmt.Println("89"); $$ = boolWithToken{false, nil} }
    |   T_ELLIPSIS                                      { fmt.Println("90"); $$ = boolWithToken{true, &$1} }
;

class_declaration_statement:
    class_modifiers T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
        {
            name := node.NewIdentifier($3.Value)
            positions.AddPosition(name, positionBuilder.NewTokenPosition($3))
            fmt.Println("91"); $$ = stmt.NewClass(name, $1, nil, $4, $5, $8, $6)
            positions.AddPosition($$, positionBuilder.NewOptionalListTokensPosition($1, $2, $9))
            
            comments.AddComments(name, $3.Comments())
            comments.AddComments($$, ListGetFirstNodeComments($1))
        }
    |   T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
        {
            name := node.NewIdentifier($2.Value)
            positions.AddPosition(name, positionBuilder.NewTokenPosition($2))
            fmt.Println("92"); $$ = stmt.NewClass(name, nil, nil, $3, $4, $7, $5)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $8))

            comments.AddComments(name, $2.Comments())
            comments.AddComments($$, $1.Comments())
        }
;

class_modifiers:
        class_modifier                                  { fmt.Println("93"); $$ = []node.Node{$1} }
    |   class_modifiers class_modifier                  { fmt.Println("94"); $$ = append($1, $2) }
;

class_modifier:
    T_ABSTRACT
        {
            fmt.Println("95"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_FINAL
        {
            fmt.Println("96"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
;

trait_declaration_statement:
    T_TRAIT T_STRING backup_doc_comment '{' class_statement_list '}'
        {
            name := node.NewIdentifier($2.Value)
            positions.AddPosition(name, positionBuilder.NewTokenPosition($2))
            fmt.Println("97"); $$ = stmt.NewTrait(name, $5, $3)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $6))

            comments.AddComments(name, $2.Comments())
            comments.AddComments($$, $1.Comments())
        }
;

interface_declaration_statement:
    T_INTERFACE T_STRING interface_extends_list backup_doc_comment '{' class_statement_list '}'
        {
            name := node.NewIdentifier($2.Value)
            positions.AddPosition(name, positionBuilder.NewTokenPosition($2))
            fmt.Println("98"); $$ = stmt.NewInterface(name, $3, $6, $4)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $7))
            
            comments.AddComments(name, $2.Comments())
            comments.AddComments($$, $1.Comments())
        }
;

extends_from:
        /* empty */                                     { fmt.Println("99"); $$ = nil }
    |   T_EXTENDS name                                  { fmt.Println("100"); $$ = $2; }
;

interface_extends_list:
        /* empty */                                     { fmt.Println("101"); $$ = nil }
    |   T_EXTENDS name_list                             { fmt.Println("102"); $$ = $2 }
;

implements_list:
        /* empty */                                     { fmt.Println("103"); $$ = nil }
    |   T_IMPLEMENTS name_list                          { fmt.Println("104"); $$ = $2 }
;

foreach_variable:
        variable                                        { fmt.Println("105"); $$ = foreachVariable{$1, false} }
    |   '&' variable                                    { fmt.Println("106"); $$ = foreachVariable{$2, true} }
    |   T_LIST '(' array_pair_list ')'
        {
            list := expr.NewList($3)
            positions.AddPosition(list, positionBuilder.NewTokensPosition($1, $4))
            fmt.Println("107"); $$ = foreachVariable{list, false}
            comments.AddComments(list, $1.Comments())
        }
    |   '[' array_pair_list ']'
        {
            list := expr.NewShortList($2)
            positions.AddPosition(list, positionBuilder.NewTokensPosition($1, $3))
            fmt.Println("108"); $$ = foreachVariable{list, false}
            comments.AddComments(list, $1.Comments())
        }
;

for_statement:
        statement                                       { fmt.Println("109"); $$ = $1; }
    |    ':' inner_statement_list T_ENDFOR ';'
        {
            fmt.Println("110"); $$ = stmt.NewStmtList($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
            comments.AddComments($$, $1.Comments())
        }
;

foreach_statement:
        statement                                       { fmt.Println("111"); $$ = $1; }
    |   ':' inner_statement_list T_ENDFOREACH ';'
        {
            fmt.Println("112"); $$ = stmt.NewStmtList($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
            comments.AddComments($$, $1.Comments())
        }
;

declare_statement:
        statement                                       { fmt.Println("113"); $$ = $1; }
    |   ':' inner_statement_list T_ENDDECLARE ';'
        {
            fmt.Println("114"); $$ = stmt.NewStmtList($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
            comments.AddComments($$, $1.Comments())
        }
;

switch_case_list:
        '{' case_list '}'                               { fmt.Println("115"); $$ = &nodesWithEndToken{$2, $3} }
    |   '{' ';' case_list '}'                           { fmt.Println("116"); $$ = &nodesWithEndToken{$3, $4} }
    |   ':' case_list T_ENDSWITCH ';'                   { fmt.Println("117"); $$ = &nodesWithEndToken{$2, $4} }
    |   ':' ';' case_list T_ENDSWITCH ';'               { fmt.Println("118"); $$ = &nodesWithEndToken{$3, $5} }
;

case_list:
        /* empty */                                     { fmt.Println("119"); $$ = []node.Node{} }
    |   case_list T_CASE expr case_separator inner_statement_list
            {
                _case := stmt.NewCase($3, $5)
                positions.AddPosition(_case, positionBuilder.NewTokenNodeListPosition($2, $5))
                fmt.Println("120"); $$ = append($1, _case)
                comments.AddComments(_case, $2.Comments())
            }
    |   case_list T_DEFAULT case_separator inner_statement_list
            {
                _default := stmt.NewDefault($4)
                positions.AddPosition(_default, positionBuilder.NewTokenNodeListPosition($2, $4))
                fmt.Println("121"); $$ = append($1, _default)
                comments.AddComments(_default, $2.Comments())
            }
;

case_separator:
        ':'
    |   ';'
;

while_statement:
        statement                                       { fmt.Println("122"); $$ = $1; }
    |   ':' inner_statement_list T_ENDWHILE ';'
        {
            fmt.Println("123"); $$ = stmt.NewStmtList($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
            comments.AddComments($$, $1.Comments())
        }
;

if_stmt_without_else:
    T_IF '(' expr ')' statement
        {
            fmt.Println("124"); $$ = stmt.NewIf($3, $5, nil, nil)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $5))
            comments.AddComments($$, $1.Comments())
        }
    |   if_stmt_without_else T_ELSEIF '(' expr ')' statement
        { 
            _elseIf := stmt.NewElseIf($4, $6)
            positions.AddPosition(_elseIf, positionBuilder.NewTokenNodePosition($2, $6))
            fmt.Println("125"); $$ = $1.(*stmt.If).AddElseIf(_elseIf)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $6))

            comments.AddComments(_elseIf, $2.Comments())
        }
;

if_stmt:
        if_stmt_without_else %prec T_NOELSE             { fmt.Println("126"); $$ = $1; }
    |   if_stmt_without_else T_ELSE statement
        {
            _else := stmt.NewElse($3)
            positions.AddPosition(_else, positionBuilder.NewTokenNodePosition($2, $3))
            fmt.Println("127"); $$ = $1.(*stmt.If).SetElse(_else)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))

            comments.AddComments($$, $2.Comments())
        }
;

alt_if_stmt_without_else:
    T_IF '(' expr ')' ':' inner_statement_list
        { 
            stmts := stmt.NewStmtList($6)
            positions.AddPosition(stmts, positionBuilder.NewNodeListPosition($6))
            fmt.Println("128"); $$ = stmt.NewAltIf($3, stmts, nil, nil)
            positions.AddPosition($$, positionBuilder.NewTokenNodeListPosition($1, $6))

            comments.AddComments(stmts, $5.Comments())
            comments.AddComments($$, $1.Comments())
        }
    |   alt_if_stmt_without_else T_ELSEIF '(' expr ')' ':' inner_statement_list
        {
            stmts := stmt.NewStmtList($7)
            positions.AddPosition(stmts, positionBuilder.NewNodeListPosition($7))
            _elseIf := stmt.NewAltElseIf($4, stmts)
            positions.AddPosition(_elseIf, positionBuilder.NewTokenNodeListPosition($2, $7))
            fmt.Println("129"); $$ = $1.(*stmt.AltIf).AddElseIf(_elseIf)

            comments.AddComments(stmts, $6.Comments())
            comments.AddComments(_elseIf, $2.Comments())
        }
;

alt_if_stmt:
    alt_if_stmt_without_else T_ENDIF ';'
        {
            fmt.Println("130"); $$ = $1
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $3))
        }
    |   alt_if_stmt_without_else T_ELSE ':' inner_statement_list T_ENDIF ';'
        {
            stmts := stmt.NewStmtList($4)
            positions.AddPosition(stmts, positionBuilder.NewNodeListPosition($4))
            _else := stmt.NewAltElse(stmts)
            positions.AddPosition(_else, positionBuilder.NewTokenNodeListPosition($2, $4))
            fmt.Println("131"); $$ = $1.(*stmt.AltIf).SetElse(_else)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $6))

            comments.AddComments(stmts, $3.Comments())
            comments.AddComments(_else, $2.Comments())
        }
;

parameter_list:
        non_empty_parameter_list                        { fmt.Println("132"); $$ = $1; }
    |   /* empty */                                     { fmt.Println("133"); $$ = nil }
;

non_empty_parameter_list:
        parameter                                       { fmt.Println("134"); $$ = []node.Node{$1} }
    |   non_empty_parameter_list ',' parameter          { fmt.Println("135"); $$ = append($1, $3) }
;

parameter:
    optional_type is_reference is_variadic T_VARIABLE
        {
            identifier := node.NewIdentifier($4.Value)
            positions.AddPosition(identifier, positionBuilder.NewTokenPosition($4))
            variable := expr.NewVariable(identifier)
            positions.AddPosition(variable, positionBuilder.NewTokenPosition($4))

            comments.AddComments($$, $4.Comments())
            comments.AddComments($$, $4.Comments())
            
            if $1 != nil {
                fmt.Println("136"); $$ = node.NewParameter($1, variable, nil, $2.value, $3.value)
                positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $4))
                comments.AddComments($$, comments[$1])
            } else if $2.value == true {
                fmt.Println("137"); $$ = node.NewParameter($1, variable, nil, $2.value, $3.value)
                positions.AddPosition($$, positionBuilder.NewTokensPosition(*$2.token, $4))
                comments.AddComments($$, $2.token.Comments())
            } else if $3.value == true {
                fmt.Println("138"); $$ = node.NewParameter($1, variable, nil, $2.value, $3.value)
                positions.AddPosition($$, positionBuilder.NewTokensPosition(*$3.token, $4))
                comments.AddComments($$, $3.token.Comments())
            } else {
                fmt.Println("139"); $$ = node.NewParameter($1, variable, nil, $2.value, $3.value)
                positions.AddPosition($$, positionBuilder.NewTokenPosition($4))
                comments.AddComments($$, $4.Comments())
            }
        }
    |   optional_type is_reference is_variadic T_VARIABLE '=' expr
        {
            identifier := node.NewIdentifier($4.Value)
            positions.AddPosition(identifier, positionBuilder.NewTokenPosition($4))
            variable := expr.NewVariable(identifier)
            positions.AddPosition(variable, positionBuilder.NewTokenPosition($4))

            comments.AddComments($$, $4.Comments())
            comments.AddComments($$, $4.Comments())

            if $1 != nil {
                fmt.Println("140"); $$ = node.NewParameter($1, variable, $6, $2.value, $3.value)
                positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $6))
                comments.AddComments($$, comments[$1])
            } else if $2.value == true {
                fmt.Println("141"); $$ = node.NewParameter($1, variable, $6, $2.value, $3.value)
                positions.AddPosition($$, positionBuilder.NewTokenNodePosition(*$2.token, $6))
                comments.AddComments($$, $2.token.Comments())
            } else if $3.value == true {
                fmt.Println("142"); $$ = node.NewParameter($1, variable, $6, $2.value, $3.value)
                positions.AddPosition($$, positionBuilder.NewTokenNodePosition(*$3.token, $6))
                comments.AddComments($$, $3.token.Comments())
            } else {
                fmt.Println("143"); $$ = node.NewParameter($1, variable, $6, $2.value, $3.value)
                positions.AddPosition($$, positionBuilder.NewTokenNodePosition($4, $6))
                comments.AddComments($$, $4.Comments())
            }
        }
;

optional_type:
        /* empty */                                     { fmt.Println("144"); $$ = nil }
    |   type_expr                                       { fmt.Println("145"); $$ = $1; }
;

type_expr:
    type                                                { fmt.Println("146"); $$ = $1; }
    |   '?' type
        {
            fmt.Println("147"); $$ = node.NewNullable($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
;

type:
    T_ARRAY
        {
            fmt.Println("148"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_CALLABLE
        {
            fmt.Println("149"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   name                                            { fmt.Println("150"); $$ = $1; }
;

return_type:
        /* empty */                                     { fmt.Println("151"); $$ = nil }
    |   ':' type_expr                                   { fmt.Println("152"); $$ = $2; }
;

argument_list:
        '(' ')'                                         { fmt.Println("153"); $$ = &nodesWithEndToken{[]node.Node{}, $2} }
    |   '(' non_empty_argument_list possible_comma ')'  { fmt.Println("154"); $$ = &nodesWithEndToken{$2, $4} }
;

non_empty_argument_list:
        argument                                        { fmt.Println("155"); $$ = []node.Node{$1} }
    |   non_empty_argument_list ',' argument            { fmt.Println("156"); $$ = append($1, $3) }
;

argument:
    expr
        {
            fmt.Println("157"); $$ = node.NewArgument($1, false, false)
            positions.AddPosition($$, positionBuilder.NewNodePosition($1))
            comments.AddComments($$, comments[$1])
        }
    |   T_ELLIPSIS expr
        {
            fmt.Println("158"); $$ = node.NewArgument($2, true, false)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
;

global_var_list:
        global_var_list ',' global_var                  { fmt.Println("159"); $$ = append($1, $3); }
    |   global_var                                      { fmt.Println("160"); $$ = []node.Node{$1} }
;

global_var:
    simple_variable                                     { fmt.Println("161"); $$ = $1 }
;

static_var_list:
        static_var_list ',' static_var                  { fmt.Println("162"); $$ = append($1, $3) }
    |   static_var                                      { fmt.Println("163"); $$ = []node.Node{$1} }
;

static_var:
        T_VARIABLE
            {
                identifier := node.NewIdentifier($1.Value)
                positions.AddPosition(identifier, positionBuilder.NewTokenPosition($1))
                variable := expr.NewVariable(identifier)
                positions.AddPosition(variable, positionBuilder.NewTokenPosition($1))
                fmt.Println("164"); $$ = stmt.NewStaticVar(variable, nil)
                positions.AddPosition($$, positionBuilder.NewTokenPosition($1))

                comments.AddComments(identifier, $1.Comments())
                comments.AddComments(variable, $1.Comments())
                comments.AddComments($$, $1.Comments())
            }
    |   T_VARIABLE '=' expr
        {
            identifier := node.NewIdentifier($1.Value)
            positions.AddPosition(identifier, positionBuilder.NewTokenPosition($1))
            variable := expr.NewVariable(identifier)
            positions.AddPosition(variable, positionBuilder.NewTokenPosition($1))
            fmt.Println("165"); $$ = stmt.NewStaticVar(variable, $3)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $3))

            comments.AddComments(identifier, $1.Comments())
            comments.AddComments(variable, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
;

class_statement_list:
        class_statement_list class_statement            { fmt.Println("166"); $$ = append($1, $2) }
    |   /* empty */                                     { fmt.Println("167"); $$ = []node.Node{} }
;

class_statement:
    variable_modifiers property_list ';'
        {
            fmt.Println("168"); $$ = stmt.NewPropertyList($1, $2)
            positions.AddPosition($$, positionBuilder.NewNodeListTokenPosition($1, $3))
            comments.AddComments($$, ListGetFirstNodeComments($1))
        }
    |   method_modifiers T_CONST class_const_list ';'
        {
            fmt.Println("169"); $$ = stmt.NewClassConstList($1, $3)
            positions.AddPosition($$, positionBuilder.NewOptionalListTokensPosition($1, $2, $4))
            comments.AddComments($$, ListGetFirstNodeComments($1))
        }
    |   T_USE name_list trait_adaptations
        {
            fmt.Println("170"); $$ = stmt.NewTraitUse($2, $3.nodes)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3.endToken))
            comments.AddComments($$, $1.Comments())
        }
    |   method_modifiers T_FUNCTION returns_ref identifier backup_doc_comment '(' parameter_list ')' return_type method_body
            {
                name := node.NewIdentifier($4.Value)
                positions.AddPosition(name, positionBuilder.NewTokenPosition($4))
                fmt.Println("171"); $$ = stmt.NewClassMethod(name, $1, $3.value, $7, $9, $10.nodes, $5)
                positions.AddPosition($$, positionBuilder.NewOptionalListTokensPosition($1, $2, $10.endToken))
                
                comments.AddComments(name, $4.Comments())
                comments.AddComments($$, ListGetFirstNodeComments($1))
            }
;

name_list:
        name                                            { fmt.Println("172"); $$ = []node.Node{$1} }
    |   name_list ',' name                              { fmt.Println("173"); $$ = append($1, $3) }
;

trait_adaptations:
        ';'                                             { fmt.Println("174"); $$ = &nodesWithEndToken{nil, $1} }
    |   '{' '}'                                         { fmt.Println("175"); $$ = &nodesWithEndToken{nil, $2} }
    |   '{' trait_adaptation_list '}'                   { fmt.Println("176"); $$ = &nodesWithEndToken{$2, $3} }
;

trait_adaptation_list:
        trait_adaptation                                { fmt.Println("177"); $$ = []node.Node{$1} }
    |   trait_adaptation_list trait_adaptation          { fmt.Println("178"); $$ = append($1, $2) }
;

trait_adaptation:
        trait_precedence ';'                            { fmt.Println("179"); $$ = $1; }
    |   trait_alias ';'                                 { fmt.Println("180"); $$ = $1; }
;

trait_precedence:
    absolute_trait_method_reference T_INSTEADOF name_list
        {
            fmt.Println("181"); $$ = stmt.NewTraitUsePrecedence($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodeNodeListPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
;

trait_alias:
    trait_method_reference T_AS T_STRING
        {
            alias := node.NewIdentifier($3.Value)
            positions.AddPosition(alias, positionBuilder.NewTokenPosition($3))
            fmt.Println("182"); $$ = stmt.NewTraitUseAlias($1, nil, alias)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $3))
            
            comments.AddComments(alias, $3.Comments())
            comments.AddComments($$, comments[$1])
        }
    |   trait_method_reference T_AS reserved_non_modifiers
        {
            alias := node.NewIdentifier($3.Value)
            positions.AddPosition(alias, positionBuilder.NewTokenPosition($3))
            fmt.Println("183"); $$ = stmt.NewTraitUseAlias($1, nil, alias)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $3))
            
            comments.AddComments(alias, $3.Comments())
            comments.AddComments($$, comments[$1])
        }
    |   trait_method_reference T_AS member_modifier identifier
        {
            alias := node.NewIdentifier($4.Value)
            positions.AddPosition(alias, positionBuilder.NewTokenPosition($4))
            fmt.Println("184"); $$ = stmt.NewTraitUseAlias($1, $3, alias)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $4))
            
            comments.AddComments(alias, $4.Comments())
            comments.AddComments($$, comments[$1])
        }
    |   trait_method_reference T_AS member_modifier
        {
            fmt.Println("185"); $$ = stmt.NewTraitUseAlias($1, $3, nil)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
;

trait_method_reference:
    identifier
        {
            name := node.NewIdentifier($1.Value)
            positions.AddPosition(name, positionBuilder.NewTokenPosition($1))
            fmt.Println("186"); $$ = stmt.NewTraitMethodRef(nil, name)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            
            comments.AddComments(name, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
    |   absolute_trait_method_reference                 { fmt.Println("187"); $$ = $1; }
;

absolute_trait_method_reference:
    name T_PAAMAYIM_NEKUDOTAYIM identifier
        {
            target := node.NewIdentifier($3.Value)
            positions.AddPosition(target, positionBuilder.NewTokenPosition($3))
            fmt.Println("188"); $$ = stmt.NewTraitMethodRef($1, target)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $3))
            
            comments.AddComments(target, $3.Comments())
            comments.AddComments($$, comments[$1])
        }
;

method_body:
        ';' /* abstract method */                       { fmt.Println("189"); $$ = &nodesWithEndToken{nil, $1} }
    |   '{' inner_statement_list '}'                    { fmt.Println("190"); $$ = &nodesWithEndToken{$2, $3} }
;

variable_modifiers:
        non_empty_member_modifiers                      { fmt.Println("191"); $$ = $1; }
    |   T_VAR
        {
            modifier := node.NewIdentifier($1.Value)
            positions.AddPosition(modifier, positionBuilder.NewTokenPosition($1))
            fmt.Println("192"); $$ = []node.Node{modifier}
            comments.AddComments(modifier, $1.Comments())
        }
;

method_modifiers:
        /* empty */                                     { fmt.Println("193"); $$ = nil }
    |   non_empty_member_modifiers                      { fmt.Println("194"); $$ = $1 }
;

non_empty_member_modifiers:
        member_modifier	                                { fmt.Println("195"); $$ = []node.Node{$1} }
    |   non_empty_member_modifiers member_modifier      { fmt.Println("196"); $$ = append($1, $2) }
;

member_modifier:
    T_PUBLIC
        {
            fmt.Println("197"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_PROTECTED
        {
            fmt.Println("198"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_PRIVATE
        {
            fmt.Println("199"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_STATIC
        {
            fmt.Println("200"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_ABSTRACT
        {
            fmt.Println("201"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_FINAL
        {
            fmt.Println("202"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
;

property_list:
        property_list ',' property                      { fmt.Println("203"); $$ = append($1, $3) }
    |   property                                        { fmt.Println("204"); $$ = []node.Node{$1} }
;

property:
    T_VARIABLE backup_doc_comment
        {
            identifier := node.NewIdentifier($1.Value)
            positions.AddPosition(identifier, positionBuilder.NewTokenPosition($1))
            variable := expr.NewVariable(identifier)
            positions.AddPosition(variable, positionBuilder.NewTokenPosition($1))
            fmt.Println("205"); $$ = stmt.NewProperty(variable, nil, $2)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))

            comments.AddComments(identifier, $1.Comments())
            comments.AddComments(variable, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
    |   T_VARIABLE '=' expr backup_doc_comment
        {
            identifier := node.NewIdentifier($1.Value)
            positions.AddPosition(identifier, positionBuilder.NewTokenPosition($1))
            variable := expr.NewVariable(identifier)
            positions.AddPosition(variable, positionBuilder.NewTokenPosition($1))
            fmt.Println("206"); $$ = stmt.NewProperty(variable, $3, $4)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $3))

            comments.AddComments(identifier, $1.Comments())
            comments.AddComments(variable, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
;

class_const_list:
        class_const_list ',' class_const_decl           { fmt.Println("207"); $$ = append($1, $3) }
    |   class_const_decl                                { fmt.Println("208"); $$ = []node.Node{$1} }
;

class_const_decl:
    identifier '=' expr backup_doc_comment
        {
            name := node.NewIdentifier($1.Value)
            positions.AddPosition(name, positionBuilder.NewTokenPosition($1))
            fmt.Println("209"); $$ = stmt.NewConstant(name, $3, $4)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $3))

            comments.AddComments(name, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
;

const_decl:
    T_STRING '=' expr backup_doc_comment
        {
            name := node.NewIdentifier($1.Value)
            positions.AddPosition(name, positionBuilder.NewTokenPosition($1))
            fmt.Println("210"); $$ = stmt.NewConstant(name, $3, $4)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $3))

            comments.AddComments(name, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
;

echo_expr_list:
        echo_expr_list ',' echo_expr                    { fmt.Println("211"); $$ = append($1, $3) }
    |   echo_expr                                       { fmt.Println("212"); $$ = []node.Node{$1} }
;

echo_expr:
    expr                                                { fmt.Println("213"); $$ = $1 }
;

for_exprs:
        /* empty */                                     { fmt.Println("214"); $$ = nil; }
    |   non_empty_for_exprs                             { fmt.Println("215"); $$ = $1; }
;
non_empty_for_exprs:
        non_empty_for_exprs ',' expr                    { fmt.Println("216"); $$ = append($1, $3) }
    |   expr                                            { fmt.Println("217"); $$ = []node.Node{$1} }
;

anonymous_class:
    T_CLASS ctor_arguments extends_from implements_list backup_doc_comment '{' class_statement_list '}'
        {
            if $2 != nil {
                fmt.Println("218"); $$ = stmt.NewClass(nil, nil, $2.nodes, $3, $4, $7, $5)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $8))
            } else {
                fmt.Println("219"); $$ = stmt.NewClass(nil, nil, nil, $3, $4, $7, $5)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $8))
            }

            comments.AddComments($$, $1.Comments())
        }
;

new_expr:
    T_NEW class_name_reference ctor_arguments
        {
            if $3 != nil {
                fmt.Println("220"); $$ = expr.NewNew($2, $3.nodes)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3.endToken))
            } else {
                fmt.Println("221"); $$ = expr.NewNew($2, nil)
                positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            }

            comments.AddComments($$, $1.Comments())
        }
    |   T_NEW anonymous_class                           { fmt.Println("222"); $$ = expr.NewNew($2, nil) }
;

expr_without_variable:
    T_LIST '(' array_pair_list ')' '=' expr
        {
            list := expr.NewList($3)
            positions.AddPosition(list, positionBuilder.NewTokensPosition($1, $4))
            fmt.Println("223"); $$ = assign_op.NewAssign(list, $6)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $6))

            comments.AddComments(list, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
    |   '[' array_pair_list ']' '=' expr
        {
            shortList := expr.NewShortList($2)
            positions.AddPosition(shortList, positionBuilder.NewTokensPosition($1, $3))
            fmt.Println("224"); $$ = assign_op.NewAssign(shortList, $5)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $5))

            comments.AddComments(shortList, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
    |   variable '=' expr
        {
            fmt.Println("225"); $$ = assign_op.NewAssign($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   variable '=' '&' expr
        {
            fmt.Println("226"); $$ = assign_op.NewAssignRef($1, $4)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $4))
            comments.AddComments($$, comments[$1])
        }
    |   T_CLONE expr
        {
            fmt.Println("227"); $$ = expr.NewClone($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   variable T_PLUS_EQUAL expr
        {
            fmt.Println("228"); $$ = assign_op.NewPlus($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   variable T_MINUS_EQUAL expr
        {
            fmt.Println("229"); $$ = assign_op.NewMinus($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   variable T_MUL_EQUAL expr
        {
            fmt.Println("230"); $$ = assign_op.NewMul($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   variable T_POW_EQUAL expr
        {
            fmt.Println("231"); $$ = assign_op.NewPow($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   variable T_DIV_EQUAL expr
        {
            fmt.Println("232"); $$ = assign_op.NewDiv($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   variable T_CONCAT_EQUAL expr
        {
            fmt.Println("233"); $$ = assign_op.NewConcat($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   variable T_MOD_EQUAL expr
        {
            fmt.Println("234"); $$ = assign_op.NewMod($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   variable T_AND_EQUAL expr
        {
            fmt.Println("235"); $$ = assign_op.NewBitwiseAnd($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   variable T_OR_EQUAL expr
        {
            fmt.Println("236"); $$ = assign_op.NewBitwiseOr($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   variable T_XOR_EQUAL expr
        {
            fmt.Println("237"); $$ = assign_op.NewBitwiseXor($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   variable T_SL_EQUAL expr
        {
            fmt.Println("238"); $$ = assign_op.NewShiftLeft($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   variable T_SR_EQUAL expr
        {
            fmt.Println("239"); $$ = assign_op.NewShiftRight($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   variable T_INC
        {
            fmt.Println("240"); $$ = expr.NewPostInc($1)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $2))
            comments.AddComments($$, comments[$1])
        }
    |   T_INC variable
        {
            fmt.Println("241"); $$ = expr.NewPreInc($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   variable T_DEC
        {
            fmt.Println("242"); $$ = expr.NewPostDec($1)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $2))
            comments.AddComments($$, comments[$1])
        }
    |   T_DEC variable
        {
            fmt.Println("243"); $$ = expr.NewPreDec($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   expr T_BOOLEAN_OR expr
        {
            fmt.Println("244"); $$ = binary_op.NewBooleanOr($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_BOOLEAN_AND expr
        {
            fmt.Println("245"); $$ = binary_op.NewBooleanAnd($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_LOGICAL_OR expr
        {
            fmt.Println("246"); $$ = binary_op.NewLogicalOr($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_LOGICAL_AND expr
        {
            fmt.Println("247"); $$ = binary_op.NewLogicalAnd($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_LOGICAL_XOR expr
        {
            fmt.Println("248"); $$ = binary_op.NewLogicalXor($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr '|' expr
        {
            fmt.Println("249"); $$ = binary_op.NewBitwiseOr($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr '&' expr
        {
            fmt.Println("250"); $$ = binary_op.NewBitwiseAnd($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr '^' expr
        {
            fmt.Println("251"); $$ = binary_op.NewBitwiseXor($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr '.' expr
        {
            fmt.Println("252"); $$ = binary_op.NewConcat($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr '+' expr
        {
            fmt.Println("253"); $$ = binary_op.NewPlus($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr '-' expr
        {
            fmt.Println("254"); $$ = binary_op.NewMinus($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr '*' expr
        {
            fmt.Println("255"); $$ = binary_op.NewMul($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_POW expr
        {
            fmt.Println("256"); $$ = binary_op.NewPow($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr '/' expr
        {
            fmt.Println("257"); $$ = binary_op.NewDiv($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr '%' expr
        {
            fmt.Println("258"); $$ = binary_op.NewMod($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_SL expr
        {
            fmt.Println("259"); $$ = binary_op.NewShiftLeft($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_SR expr
        {
            fmt.Println("260"); $$ = binary_op.NewShiftRight($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   '+' expr %prec T_INC
        {
            fmt.Println("261"); $$ = expr.NewUnaryPlus($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   '-' expr %prec T_INC
        {
            fmt.Println("262"); $$ = expr.NewUnaryMinus($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   '!' expr
        {
            fmt.Println("263"); $$ = expr.NewBooleanNot($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   '~' expr
        {
            fmt.Println("264"); $$ = expr.NewBitwiseNot($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   expr T_IS_IDENTICAL expr
        {
            fmt.Println("265"); $$ = binary_op.NewIdentical($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_IS_NOT_IDENTICAL expr
        {
            fmt.Println("266"); $$ = binary_op.NewNotIdentical($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_IS_EQUAL expr
        {
            fmt.Println("267"); $$ = binary_op.NewEqual($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_IS_NOT_EQUAL expr
        {
            fmt.Println("268"); $$ = binary_op.NewNotEqual($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr '<' expr
        {
            fmt.Println("269"); $$ = binary_op.NewSmaller($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_IS_SMALLER_OR_EQUAL expr
        {
            fmt.Println("270"); $$ = binary_op.NewSmallerOrEqual($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr '>' expr
        {
            fmt.Println("271"); $$ = binary_op.NewGreater($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_IS_GREATER_OR_EQUAL expr
        {
            fmt.Println("272"); $$ = binary_op.NewGreaterOrEqual($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_SPACESHIP expr
        {
            fmt.Println("273"); $$ = binary_op.NewSpaceship($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_INSTANCEOF class_name_reference
        {
            fmt.Println("274"); $$ = expr.NewInstanceOf($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   '(' expr ')'                                    { fmt.Println("275"); $$ = $2; }
    |   new_expr                                        { fmt.Println("276"); $$ = $1; }
    |   expr '?' expr ':' expr
        {
            fmt.Println("277"); $$ = expr.NewTernary($1, $3, $5)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $5))
            comments.AddComments($$, comments[$1])
        }
    |   expr '?' ':' expr
        {
            fmt.Println("278"); $$ = expr.NewTernary($1, nil, $4)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $4))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_COALESCE expr
        {
            fmt.Println("279"); $$ = binary_op.NewCoalesce($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   internal_functions_in_yacc                      { fmt.Println("280"); $$ = $1}
    |   T_INT_CAST expr
        {
            fmt.Println("281"); $$ = cast.NewCastInt($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_DOUBLE_CAST expr
        {
            fmt.Println("282"); $$ = cast.NewCastDouble($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_STRING_CAST expr
        {
            fmt.Println("283"); $$ = cast.NewCastString($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_ARRAY_CAST expr
        {
            fmt.Println("284"); $$ = cast.NewCastArray($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_OBJECT_CAST expr
        {
            fmt.Println("285"); $$ = cast.NewCastObject($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_BOOL_CAST expr
        {
            fmt.Println("286"); $$ = cast.NewCastBool($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_UNSET_CAST expr
        {
            fmt.Println("287"); $$ = cast.NewCastUnset($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_EXIT exit_expr
        {
            fmt.Println("288"); $$ = expr.NewExit($2, strings.EqualFold($1.Value, "die"))
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   '@' expr
        {
            fmt.Println("289"); $$ = expr.NewErrorSuppress($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   scalar                                          { fmt.Println("290"); $$ = $1; }
    |   '`' backticks_expr '`'
        {
            fmt.Println("291"); $$ = expr.NewShellExec($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    |   T_PRINT expr
        {
            fmt.Println("292"); $$ = expr.NewPrint($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_YIELD
        {
            fmt.Println("293"); $$ = expr.NewYield(nil, nil)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_YIELD expr
        {
            fmt.Println("294"); $$ = expr.NewYield(nil, $2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_YIELD expr T_DOUBLE_ARROW expr
        {
            fmt.Println("295"); $$ = expr.NewYield($2, $4)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $4))
            comments.AddComments($$, $1.Comments())
        }
    |   T_YIELD_FROM expr
        {
            fmt.Println("296"); $$ = expr.NewYieldFrom($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                fmt.Println("297"); $$ = expr.NewClosure($5, $7, $8, $10, false, $2.value, $3)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $11))
                
                comments.AddComments($$, $1.Comments())
            }
    |   T_STATIC T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                fmt.Println("298"); $$ = expr.NewClosure($6, $8, $9, $11, true, $3.value, $4)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $12))
                
                comments.AddComments($$, $1.Comments())
            }
;

backup_doc_comment:
	/* empty */ { fmt.Println("299"); $$ = yylex.(*lexer).PhpDocComment; yylex.(*lexer).PhpDocComment = "" }
;

returns_ref:
        /* empty */                                     { fmt.Println("300"); $$ = boolWithToken{false, nil} }
    |   '&'                                             { fmt.Println("301"); $$ = boolWithToken{true, &$1} }
;

lexical_vars:
        /* empty */                                     { fmt.Println("302"); $$ = []node.Node{} }
    |   T_USE '(' lexical_var_list ')'                  { fmt.Println("303"); $$ = $3; }
;

lexical_var_list:
        lexical_var_list ',' lexical_var                { fmt.Println("304"); $$ = append($1, $3) }
    |   lexical_var                                     { fmt.Println("305"); $$ = []node.Node{$1} }
;

lexical_var:
    T_VARIABLE
        {
            identifier := node.NewIdentifier($1.Value)
            positions.AddPosition(identifier, positionBuilder.NewTokenPosition($1))
            variable := expr.NewVariable(identifier)
            positions.AddPosition(variable, positionBuilder.NewTokenPosition($1))
            fmt.Println("306"); $$ = expr.NewClosureUse(variable, false)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))

            comments.AddComments(identifier, $1.Comments())
            comments.AddComments(variable, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
    |   '&' T_VARIABLE
        {
            identifier := node.NewIdentifier($2.Value)
            positions.AddPosition(identifier, positionBuilder.NewTokenPosition($2))
            variable := expr.NewVariable(identifier)
            positions.AddPosition(variable, positionBuilder.NewTokenPosition($2))
            fmt.Println("307"); $$ = expr.NewClosureUse(variable, true)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $2))

            comments.AddComments(identifier, $2.Comments())
            comments.AddComments(variable, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
;

function_call:
    name argument_list
        {
            fmt.Println("308"); $$ = expr.NewFunctionCall($1, $2.nodes)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $2.endToken))
            comments.AddComments($$, comments[$1])
        }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
        {
            fmt.Println("309"); $$ = expr.NewStaticCall($1, $3, $4.nodes)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $4.endToken))
            comments.AddComments($$, comments[$1])
        }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
        {
            fmt.Println("310"); $$ = expr.NewStaticCall($1, $3, $4.nodes)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $4.endToken))
            comments.AddComments($$, comments[$1])
        }
    |   callable_expr argument_list
        {
            fmt.Println("311"); $$ = expr.NewFunctionCall($1, $2.nodes)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $2.endToken))
            comments.AddComments($$, comments[$1])
        }
;

class_name:
    T_STATIC
        {
            fmt.Println("312"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   name                                            { fmt.Println("313"); $$ = $1; }
;

class_name_reference:
        class_name                                      { fmt.Println("314"); $$ = $1; }
    |   new_variable                                    { fmt.Println("315"); $$ = $1; }
;

exit_expr:
        /* empty */                                     { fmt.Println("316"); $$ = nil }
    |   '(' optional_expr ')'                           { fmt.Println("317"); $$ = $2; }
;

backticks_expr:
        /* empty */                                     { fmt.Println("318"); $$ = []node.Node{} }
    |   T_ENCAPSED_AND_WHITESPACE                       { fmt.Println("319"); $$ = []node.Node{scalar.NewEncapsedStringPart($1.Value)} }
    |   encaps_list                                     { fmt.Println("320"); $$ = $1; }
;

ctor_arguments:
        /* empty */	                                    { fmt.Println("321"); $$ = nil }
    |   argument_list                                   { fmt.Println("322"); $$ = $1 }
;

dereferencable_scalar:
    T_ARRAY '(' array_pair_list ')'
        {
            fmt.Println("323"); $$ = expr.NewArray($3)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
            comments.AddComments($$, $1.Comments())
        }
    |   '[' array_pair_list ']'
        {
            fmt.Println("324"); $$ = expr.NewShortArray($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    |   T_CONSTANT_ENCAPSED_STRING
        {
            fmt.Println("325"); $$ = scalar.NewString($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
;

scalar:
    T_LNUMBER
        {
            fmt.Println("326"); $$ = scalar.NewLnumber($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_DNUMBER
        {
            fmt.Println("327"); $$ = scalar.NewDnumber($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_LINE
        {
            fmt.Println("328"); $$ = scalar.NewMagicConstant($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_FILE
        {
            fmt.Println("329"); $$ = scalar.NewMagicConstant($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_DIR
        {
            fmt.Println("330"); $$ = scalar.NewMagicConstant($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_TRAIT_C
        {
            fmt.Println("331"); $$ = scalar.NewMagicConstant($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_METHOD_C
        {
            fmt.Println("332"); $$ = scalar.NewMagicConstant($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_FUNC_C
        {
            fmt.Println("333"); $$ = scalar.NewMagicConstant($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_NS_C
        {
            fmt.Println("334"); $$ = scalar.NewMagicConstant($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_CLASS_C
        {
            fmt.Println("335"); $$ = scalar.NewMagicConstant($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC 
        {
            fmt.Println("336"); $$ = scalar.NewString($2.Value)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))/* TODO: mark as Heredoc*/
            comments.AddComments($$, $1.Comments())
        }
    |   T_START_HEREDOC T_END_HEREDOC
        {
            fmt.Println("337"); $$ = scalar.NewEncapsed(nil)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   '"' encaps_list '"'
        {
            fmt.Println("338"); $$ = scalar.NewEncapsed($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    |   T_START_HEREDOC encaps_list T_END_HEREDOC
        {
            fmt.Println("339"); $$ = scalar.NewEncapsed($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    |   dereferencable_scalar                           { fmt.Println("340"); $$ = $1; }
    |   constant                                        { fmt.Println("341"); $$ = $1; }
;

constant:
    name
        {
            fmt.Println("342"); $$ = expr.NewConstFetch($1)
            positions.AddPosition($$, positionBuilder.NewNodePosition($1))
            comments.AddComments($$, comments[$1])
        }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM identifier
        {
            target := node.NewIdentifier($3.Value)
            positions.AddPosition(target, positionBuilder.NewTokenPosition($3))
            fmt.Println("343"); $$ = expr.NewClassConstFetch($1, target)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $3))

            comments.AddComments(target, $3.Comments())
            comments.AddComments($$, comments[$1])
        }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM identifier
        {
            target := node.NewIdentifier($3.Value)
            positions.AddPosition(target, positionBuilder.NewTokenPosition($3))
            fmt.Println("344"); $$ = expr.NewClassConstFetch($1, target)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $3))

            comments.AddComments(target, $3.Comments())
            comments.AddComments($$, comments[$1])
        }
;

expr:
        variable                                        { fmt.Println("345"); $$ = $1; }
    |   expr_without_variable                           { fmt.Println("346"); $$ = $1; }
;

optional_expr:
        /* empty */                                     { fmt.Println("347"); $$ = nil }
    |   expr                                            { fmt.Println("348"); $$ = $1; }
;

variable_class_name:
    dereferencable                                      { fmt.Println("349"); $$ = $1; }
;

dereferencable:
        variable                                        { fmt.Println("350"); $$ = $1; }
    |   '(' expr ')'                                    { fmt.Println("351"); $$ = $2; }
    |   dereferencable_scalar                           { fmt.Println("352"); $$ = $1; }
;

callable_expr:
        callable_variable                               { fmt.Println("353"); $$ = $1; }
    |   '(' expr ')'                                    { fmt.Println("354"); $$ = $2; }
    |   dereferencable_scalar                           { fmt.Println("355"); $$ = $1; }
;

callable_variable:
    simple_variable                                     { fmt.Println("356"); $$ = $1; }
    |   dereferencable '[' optional_expr ']'
        {
            fmt.Println("357"); $$ = expr.NewArrayDimFetch($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $4))
            comments.AddComments($$, comments[$1])
        }
    |   constant '[' optional_expr ']'
        {
            fmt.Println("358"); $$ = expr.NewArrayDimFetch($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $4))
            comments.AddComments($$, comments[$1])
        }
    |   dereferencable '{' expr '}'
        {
            fmt.Println("359"); $$ = expr.NewArrayDimFetch($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $4))
            comments.AddComments($$, comments[$1])
        }
    |   dereferencable T_OBJECT_OPERATOR property_name argument_list
        {
            fmt.Println("360"); $$ = expr.NewMethodCall($1, $3, $4.nodes)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $4.endToken))
            comments.AddComments($$, comments[$1])
        }
    |   function_call                                   { fmt.Println("361"); $$ = $1; }
;

variable:
        callable_variable                               { fmt.Println("362"); $$ = $1; }
    |   static_member                                   { fmt.Println("363"); $$ = $1; }
    |   dereferencable T_OBJECT_OPERATOR property_name
        {
            fmt.Println("364"); $$ = expr.NewPropertyFetch($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
;

simple_variable:
    T_VARIABLE
        {
            name := node.NewIdentifier($1.Value)
            positions.AddPosition(name, positionBuilder.NewTokenPosition($1))
            fmt.Println("365"); $$ = expr.NewVariable(name)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            
            comments.AddComments(name, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
    |   '$' '{' expr '}'
        {
            fmt.Println("366"); $$ = expr.NewVariable($3)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
            comments.AddComments($$, $1.Comments())
        }
    |   '$' simple_variable
        {
            fmt.Println("367"); $$ = expr.NewVariable($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
;

static_member:
    class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
        {
            fmt.Println("368"); $$ = expr.NewStaticPropertyFetch($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
        {
            fmt.Println("369"); $$ = expr.NewStaticPropertyFetch($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
;

new_variable:
        simple_variable                                 { fmt.Println("370"); $$ = $1 }
    |   new_variable '[' optional_expr ']'
        {
            fmt.Println("371"); $$ = expr.NewArrayDimFetch($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $4))
            comments.AddComments($$, comments[$1])
        }
    |   new_variable '{' expr '}'
        {
            fmt.Println("372"); $$ = expr.NewArrayDimFetch($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $4))
            comments.AddComments($$, comments[$1])
        }
    |   new_variable T_OBJECT_OPERATOR property_name
        {
            fmt.Println("373"); $$ = expr.NewPropertyFetch($1, $3)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                fmt.Println("374"); $$ = expr.NewStaticPropertyFetch($1, $3)
                positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
                comments.AddComments($$, comments[$1])
            }
    |   new_variable T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                fmt.Println("375"); $$ = expr.NewStaticPropertyFetch($1, $3)
                positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
                comments.AddComments($$, comments[$1])
            }
;

member_name:
    identifier
        {
            fmt.Println("376"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   '{' expr '}'                                    { fmt.Println("377"); $$ = $2; }
    |   simple_variable                                 { fmt.Println("378"); $$ = $1 }
;

property_name:
    T_STRING
        {
            fmt.Println("379"); $$ = node.NewIdentifier($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   '{' expr '}'                                    { fmt.Println("380"); $$ = $2; }
    |   simple_variable                                 { fmt.Println("381"); $$ = $1 }
;

array_pair_list:
    non_empty_array_pair_list
        {
            if ($1[len($1)-1] == nil) {
                fmt.Println("382"); $$ = $1[:len($1)-1]
            } else {
                fmt.Println("383"); $$ = $1
            }
        }
;

possible_array_pair:
        /* empty */                                     { fmt.Println("384"); $$ = nil }
    |   array_pair                                      { fmt.Println("385"); $$ = $1; }
;

non_empty_array_pair_list:
        non_empty_array_pair_list ',' possible_array_pair
                                                        { fmt.Println("386"); $$ = append($1, $3) }
    |   possible_array_pair                             { fmt.Println("387"); $$ = []node.Node{$1} }
;

array_pair:
    expr T_DOUBLE_ARROW expr
        {
            fmt.Println("388"); $$ = expr.NewArrayItem($1, $3, false)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $3))
            comments.AddComments($$, comments[$1])
        }
    |   expr
        {
            fmt.Println("389"); $$ = expr.NewArrayItem(nil, $1, false)
            positions.AddPosition($$, positionBuilder.NewNodePosition($1))
            comments.AddComments($$, comments[$1])
        }
    |   expr T_DOUBLE_ARROW '&' variable
        {
            fmt.Println("390"); $$ = expr.NewArrayItem($1, $4, true)
            positions.AddPosition($$, positionBuilder.NewNodesPosition($1, $4))
            comments.AddComments($$, comments[$1])
        }
    |   '&' variable
        {
            fmt.Println("391"); $$ = expr.NewArrayItem(nil, $2, true)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   expr T_DOUBLE_ARROW T_LIST '(' array_pair_list ')'
            {
                // TODO: Cannot use list() as standalone expression
                list := expr.NewList($5)
                positions.AddPosition(list, positionBuilder.NewTokensPosition($3, $6))
                fmt.Println("392"); $$ = expr.NewArrayItem($1, list, false)
                positions.AddPosition($$, positionBuilder.NewNodeTokenPosition($1, $6))

                comments.AddComments(list, $3.Comments())
                comments.AddComments($$, comments[$1])
            }
    |   T_LIST '(' array_pair_list ')'
            {
                // TODO: Cannot use list() as standalone expression
                list := expr.NewList($3)
                positions.AddPosition(list, positionBuilder.NewTokensPosition($1, $4))
                fmt.Println("393"); $$ = expr.NewArrayItem(nil, list, false)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
                
                comments.AddComments(list, $1.Comments())
                comments.AddComments($$, $1.Comments())
            }
;

encaps_list:
        encaps_list encaps_var                          { fmt.Println("394"); $$ = append($1, $2) }
    |   encaps_list T_ENCAPSED_AND_WHITESPACE
        {
            encapsed := scalar.NewEncapsedStringPart($2.Value)
            positions.AddPosition(encapsed, positionBuilder.NewTokenPosition($2))
            fmt.Println("395"); $$ = append($1, encapsed)
            comments.AddComments(encapsed, $2.Comments())
        }
    |   encaps_var                                      { fmt.Println("396"); $$ = []node.Node{$1} }
    |   T_ENCAPSED_AND_WHITESPACE encaps_var
        {
            encapsed := scalar.NewEncapsedStringPart($1.Value)
            positions.AddPosition(encapsed, positionBuilder.NewTokenPosition($1))
            fmt.Println("397"); $$ = []node.Node{encapsed, $2}
            comments.AddComments(encapsed, $1.Comments())
        }
;

encaps_var:
    T_VARIABLE
        {
            name := node.NewIdentifier($1.Value)
            positions.AddPosition(name, positionBuilder.NewTokenPosition($1))
            fmt.Println("398"); $$ = expr.NewVariable(name)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))

            comments.AddComments(name, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
    |   T_VARIABLE '[' encaps_var_offset ']'
        {
            identifier := node.NewIdentifier($1.Value)
            positions.AddPosition(identifier, positionBuilder.NewTokenPosition($1))
            variable := expr.NewVariable(identifier)
            positions.AddPosition(variable, positionBuilder.NewTokenPosition($1))
            fmt.Println("399"); $$ = expr.NewArrayDimFetch(variable, $3)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))

            comments.AddComments(identifier, $1.Comments())
            comments.AddComments(variable, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
    |   T_VARIABLE T_OBJECT_OPERATOR T_STRING
        {
            identifier := node.NewIdentifier($1.Value)
            positions.AddPosition(identifier, positionBuilder.NewTokenPosition($1))
            variable := expr.NewVariable(identifier)
            positions.AddPosition(variable, positionBuilder.NewTokenPosition($1))
            fetch := node.NewIdentifier($3.Value)
            positions.AddPosition(fetch, positionBuilder.NewTokenPosition($3))
            fmt.Println("400"); $$ = expr.NewPropertyFetch(variable, fetch)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            
            comments.AddComments(identifier, $1.Comments())
            comments.AddComments(variable, $1.Comments())
            comments.AddComments(fetch, $3.Comments())
            comments.AddComments($$, $1.Comments())
        }
    |   T_DOLLAR_OPEN_CURLY_BRACES expr '}'
        {
            fmt.Println("401"); $$ = expr.NewVariable($2)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))
            comments.AddComments($$, $1.Comments())
        }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}'
        {
            name := node.NewIdentifier($2.Value)
            positions.AddPosition(name, positionBuilder.NewTokenPosition($2))
            fmt.Println("402"); $$ = expr.NewVariable(name)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $3))

            comments.AddComments(name, $2.Comments())
            comments.AddComments($$, $1.Comments())
        }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
        {
            identifier := node.NewIdentifier($2.Value)
            positions.AddPosition(identifier, positionBuilder.NewTokenPosition($2))
            variable := expr.NewVariable(identifier)
            positions.AddPosition(variable, positionBuilder.NewTokenPosition($2))
            fmt.Println("403"); $$ = expr.NewArrayDimFetch(variable, $4)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $6))


            comments.AddComments(identifier, $2.Comments())
            comments.AddComments(variable, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
    |   T_CURLY_OPEN variable '}'                       { fmt.Println("404"); $$ = $2; }
;
encaps_var_offset:
    T_STRING
        {
            fmt.Println("405"); $$ = scalar.NewString($1.Value)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            comments.AddComments($$, $1.Comments())
        }
    |   T_NUM_STRING
        {
            // TODO: add option to handle 64 bit integer
            if _, err := strconv.Atoi($1.Value); err == nil {
                fmt.Println("406"); $$ = scalar.NewLnumber($1.Value)
                positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            } else {
                fmt.Println("407"); $$ = scalar.NewString($1.Value)
                positions.AddPosition($$, positionBuilder.NewTokenPosition($1))
            }
            comments.AddComments($$, $1.Comments())
        }
    |   '-' T_NUM_STRING
        {
            // TODO: add option to handle 64 bit integer
            if _, err := strconv.Atoi($2.Value); err == nil {
                lnumber := scalar.NewLnumber($2.Value)
                positions.AddPosition(lnumber, positionBuilder.NewTokensPosition($1, $2))
                fmt.Println("408"); $$ = expr.NewUnaryMinus(lnumber)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $2))
                
                comments.AddComments(lnumber, $1.Comments())
            } else {
                $2.Value = "-"+$2.Value
                fmt.Println("409"); $$ = scalar.NewString($2.Value)
                positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $2))
            }

            comments.AddComments($$, $1.Comments())
        }
    |   T_VARIABLE
        {
            identifier := node.NewIdentifier($1.Value)
            positions.AddPosition(identifier, positionBuilder.NewTokenPosition($1))
            fmt.Println("410"); $$ = expr.NewVariable(identifier)
            positions.AddPosition($$, positionBuilder.NewTokenPosition($1))

            comments.AddComments(identifier, $1.Comments())
            comments.AddComments($$, $1.Comments())
        }
;

internal_functions_in_yacc:
    T_ISSET '(' isset_variables possible_comma ')'
        {
            fmt.Println("411"); $$ = expr.NewIsset($3)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $5))
            comments.AddComments($$, $1.Comments())
        }
    |   T_EMPTY '(' expr ')'
        {
            fmt.Println("412"); $$ = expr.NewEmpty($3)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
            comments.AddComments($$, $1.Comments())
        }
    |   T_INCLUDE expr
        {
            fmt.Println("413"); $$ = expr.NewInclude($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_INCLUDE_ONCE expr
        {
            fmt.Println("414"); $$ = expr.NewIncludeOnce($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_EVAL '(' expr ')'
        {
            fmt.Println("415"); $$ = expr.NewEval($3)
            positions.AddPosition($$, positionBuilder.NewTokensPosition($1, $4))
            comments.AddComments($$, $1.Comments())
        }
    |   T_REQUIRE expr
        {
            fmt.Println("416"); $$ = expr.NewRequire($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
    |   T_REQUIRE_ONCE expr
        {
            fmt.Println("417"); $$ = expr.NewRequireOnce($2)
            positions.AddPosition($$, positionBuilder.NewTokenNodePosition($1, $2))
            comments.AddComments($$, $1.Comments())
        }
;

isset_variables:
        isset_variable                                  { fmt.Println("418"); $$ = []node.Node{$1} }
    |   isset_variables ',' isset_variable              { fmt.Println("419"); $$ = append($1, $3) }
;

isset_variable:
    expr                                                { fmt.Println("420"); $$ = $1 }
;

/////////////////////////////////////////////////////////////////////////

%%