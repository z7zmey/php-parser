%{
package php7

import (
    "strings"
    "strconv"

    "github.com/z7zmey/php-parser/token"
    "github.com/z7zmey/php-parser/node"
    "github.com/z7zmey/php-parser/node/scalar"
    "github.com/z7zmey/php-parser/node/name"
    "github.com/z7zmey/php-parser/node/stmt"
    "github.com/z7zmey/php-parser/node/expr"
    "github.com/z7zmey/php-parser/node/expr/assign"
    "github.com/z7zmey/php-parser/node/expr/binary"
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
    altSyntaxNode altSyntaxNode
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
%type <node> expr optional_expr
%type <node> declare_statement finally_statement unset_variable variable
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

%type <altSyntaxNode> while_statement for_statement foreach_statement

%%

/////////////////////////////////////////////////////////////////////////

start:
    top_statement_list
        {
            yylex.(*Parser).rootNode = stmt.NewStmtList($1)
            yylex.(*Parser).positions.AddPosition(yylex.(*Parser).rootNode, yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
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
        T_STRING                                        { $$ = $1 }
    |   semi_reserved                                   { $$ = $1 }
;

top_statement_list:
    top_statement_list top_statement
        {
            if $2 != nil {
                $$ = append($1, $2)
            }
        }
    |   /* empty */                                     { $$ = []node.Node{} }
;

namespace_name:
    T_STRING
        {
            namePart := name.NewNamePart($1.Value)
            yylex.(*Parser).positions.AddPosition(namePart, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = []node.Node{namePart}
            yylex.(*Parser).comments.AddComments(namePart, $1.Comments())
        }
    |   namespace_name T_NS_SEPARATOR T_STRING
        {
            namePart := name.NewNamePart($3.Value)
            yylex.(*Parser).positions.AddPosition(namePart, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
            $$ = append($1, namePart)
            yylex.(*Parser).comments.AddComments(namePart, $3.Comments())
        }
;

name:
    namespace_name
        {
            $$ = name.NewName($1)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($1))
        }
    | T_NAMESPACE T_NS_SEPARATOR namespace_name
        {
            $$ = name.NewRelative($3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    | T_NS_SEPARATOR namespace_name
        {
            $$ = name.NewFullyQualified($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

top_statement:
    error
        {
            // error
            $$ = nil
        }
    |   statement                                       { $$ = $1; }
    |   function_declaration_statement                  { $$ = $1; }
    |   class_declaration_statement                     { $$ = $1; }
    |   trait_declaration_statement                     { $$ = $1; }
    |   interface_declaration_statement                 { $$ = $1; }
    |   T_HALT_COMPILER '(' ')' ';'                     { $$ = stmt.NewHaltCompiler() }
    |   T_NAMESPACE namespace_name ';'
        {
            name := name.NewName($2)
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
            $$ = stmt.NewNamespace(name, nil)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

            yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_NAMESPACE namespace_name '{' top_statement_list '}'
        {
            name := name.NewName($2)
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
            $$ = stmt.NewNamespace(name, $4)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $5))

            yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_NAMESPACE '{' top_statement_list '}'
        {
            $$ = stmt.NewNamespace(nil, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_USE mixed_group_use_declaration ';'           { $$ = $2 }
    |   T_USE use_type group_use_declaration ';'        { $$ = $3.(*stmt.GroupUse).SetUseType($2) }
    |   T_USE use_declarations ';'
        {
            $$ = stmt.NewUseList(nil, $2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_USE use_type use_declarations ';'             { $$ = stmt.NewUseList($2, $3) }
    |   T_CONST const_list ';'
        {
            $$ = stmt.NewConstList($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

use_type:
    T_FUNCTION
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_CONST
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                name := name.NewName($1)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                $$ = stmt.NewGroupUse(nil, name, $4)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($1, $6))

                yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($1))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($1))
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                name := name.NewName($2)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$ = stmt.NewGroupUse(nil, name, $5)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $7))

                yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

mixed_group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                name := name.NewName($1)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                $$ = stmt.NewGroupUse(nil, name, $4)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($1, $6))

                yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($1))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($1))
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                name := name.NewName($2)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$ = stmt.NewGroupUse(nil, name, $5)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $7))

                yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

possible_comma:
        /* empty */
    |   ','
;

inline_use_declarations:
        inline_use_declarations ',' inline_use_declaration
                                                        { $$ = append($1, $3) }
    |   inline_use_declaration                          { $$ = []node.Node{$1} }
;

unprefixed_use_declarations:
        unprefixed_use_declarations ',' unprefixed_use_declaration
                                                        { $$ = append($1, $3) }
    |   unprefixed_use_declaration                      { $$ = []node.Node{$1} }
;

use_declarations:
        use_declarations ',' use_declaration            { $$ = append($1, $3) }
    |   use_declaration                                 { $$ = []node.Node{$1} }
;

inline_use_declaration:
        unprefixed_use_declaration                      { $$ = $1; }
    |   use_type unprefixed_use_declaration             { $$ = $2.(*stmt.Use).SetUseType($1) }
;

unprefixed_use_declaration:
    namespace_name
        {
            name := name.NewName($1)
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
            $$ = stmt.NewUse(nil, name, nil)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListPosition($1))

            yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($1))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($1))
        }
    |   namespace_name T_AS T_STRING
        {
            name := name.NewName($1)
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
            alias := node.NewIdentifier($3.Value)
            yylex.(*Parser).positions.AddPosition(alias, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
            $$ = stmt.NewUse(nil, name, alias)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($1, $3))

            yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($1))
            yylex.(*Parser).comments.AddComments(alias, $3.Comments())
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($1))
        }
;

use_declaration:
        unprefixed_use_declaration                      { $$ = $1; }
    |   T_NS_SEPARATOR unprefixed_use_declaration       { $$ = $2; }
;

const_list:
        const_list ',' const_decl                       { $$ = append($1, $3) }
    |   const_decl                                      { $$ = []node.Node{$1} }
;

inner_statement_list:
    inner_statement_list inner_statement
        {
            if $2 != nil {
                $$ = append($1, $2)
            }
        }
    |   /* empty */                                     { $$ = []node.Node{} }
;

inner_statement:
    error
        {
            // error
            $$ = nil
        }
    |   statement                                       { $$ = $1; }
    |   function_declaration_statement                  { $$ = $1; }
    |   class_declaration_statement                     { $$ = $1; }
    |   trait_declaration_statement                     { $$ = $1; }
    |   interface_declaration_statement                 { $$ = $1; }
    |   T_HALT_COMPILER '(' ')' ';'
        {
            $$ = stmt.NewHaltCompiler()
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }

statement:
    '{' inner_statement_list '}'
        {
            $$ = stmt.NewStmtList($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   if_stmt                                         { $$ = $1; }
    |   alt_if_stmt                                     { $$ = $1; }
    |   T_WHILE '(' expr ')' while_statement
        {
            if ($5.isAlt) {
                $$ = stmt.NewAltWhile($3, $5.node)
            } else {
                $$ = stmt.NewWhile($3, $5.node)
            }
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $5.node))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_DO statement T_WHILE '(' expr ')' ';'
        {
            $$ = stmt.NewDo($2, $5)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $7))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_FOR '(' for_exprs ';' for_exprs ';' for_exprs ')' for_statement
        {
            if ($9.isAlt) {
                $$ = stmt.NewAltFor($3, $5, $7, $9.node)
            } else {
                $$ = stmt.NewFor($3, $5, $7, $9.node)
            }
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $9.node))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_SWITCH '(' expr ')' switch_case_list
        {
            if ($5.endToken.Value == ";") {
                $$ = stmt.NewAltSwitch($3, $5.nodes)
            } else {
                $$ = stmt.NewSwitch($3, $5.nodes)
            }
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $5.endToken))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_BREAK optional_expr ';'
        {
            $$ = stmt.NewBreak($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_CONTINUE optional_expr ';'
        {
            $$ = stmt.NewContinue($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_RETURN optional_expr ';'
        {
            $$ = stmt.NewReturn($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_GLOBAL global_var_list ';'
        {
            $$ = stmt.NewGlobal($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_STATIC static_var_list ';'
        {
            $$ = stmt.NewStatic($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_ECHO echo_expr_list ';'
        {
            $$ = stmt.NewEcho($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_INLINE_HTML
        {
            $$ = stmt.NewInlineHtml($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   expr ';'
        {
            $$ = stmt.NewExpression($1)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   T_UNSET '(' unset_variables possible_comma ')' ';' 
        {
            $$ = stmt.NewUnset($3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $6))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_FOREACH '(' expr T_AS foreach_variable ')' foreach_statement
        {
            if ($7.isAlt) {
                $$ = stmt.NewAltForeach($3, nil, $5.node, $7.node, $5.byRef)
            } else {
                $$ = stmt.NewForeach($3, nil, $5.node, $7.node, $5.byRef)
            }
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $7.node))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_FOREACH '(' expr T_AS variable T_DOUBLE_ARROW foreach_variable ')' foreach_statement
        {
            if ($9.isAlt) {
                $$ = stmt.NewAltForeach($3, $5, $7.node, $9.node, $7.byRef)
            } else {
                $$ = stmt.NewForeach($3, $5, $7.node, $9.node, $7.byRef)
            }
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $9.node))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_DECLARE '(' const_list ')' declare_statement
        {
            $$ = stmt.NewDeclare($3, $5)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $5))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   ';'
        {
            $$ = stmt.NewNop()
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_TRY '{' inner_statement_list '}' catch_list finally_statement
            {
                if $6 == nil {
                    $$ = stmt.NewTry($3, $5, $6)
                    yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $5))
                } else {
                    $$ = stmt.NewTry($3, $5, $6)
                    yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $6))
                }

                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_THROW expr ';'
        {
            $$ = stmt.NewThrow($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_GOTO T_STRING ';'
        {
            label := node.NewIdentifier($2.Value)
            yylex.(*Parser).positions.AddPosition(label, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
            $$ = stmt.NewGoto(label)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

            yylex.(*Parser).comments.AddComments(label, $2.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_STRING ':'
        {
            label := node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition(label, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = stmt.NewLabel(label)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))

            yylex.(*Parser).comments.AddComments(label, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }

catch_list:
        /* empty */                                     { $$ = []node.Node{} }
    |   catch_list T_CATCH '(' catch_name_list T_VARIABLE ')' '{' inner_statement_list '}'
        {
            identifier := node.NewIdentifier(strings.TrimLeft($5.Value, "$"))
            yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($5))
            variable := expr.NewVariable(identifier)
            yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($5))
            catch := stmt.NewCatch($4, variable, $8)
            yylex.(*Parser).positions.AddPosition(catch, yylex.(*Parser).positionBuilder.NewTokensPosition($2, $9))
            $$ = append($1, catch)

            yylex.(*Parser).comments.AddComments(identifier, $5.Comments())
            yylex.(*Parser).comments.AddComments(variable, $5.Comments())
            yylex.(*Parser).comments.AddComments(catch, $2.Comments())
        }
;
catch_name_list:
        name                                            { $$ = []node.Node{$1} }
    |   catch_name_list '|' name                        { $$ = append($1, $3) }
;

finally_statement:
        /* empty */                                     { $$ = nil }
    |   T_FINALLY '{' inner_statement_list '}'
        {
            $$ = stmt.NewFinally($3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

unset_variables:
        unset_variable                                  { $$ = []node.Node{$1} }
    |   unset_variables ',' unset_variable              { $$ = append($1, $3) }
;

unset_variable:
    variable                                            { $$ = $1 }
;

function_declaration_statement:
    T_FUNCTION returns_ref T_STRING backup_doc_comment '(' parameter_list ')' return_type '{' inner_statement_list '}'
        {
            name := node.NewIdentifier($3.Value)
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
            $$ = stmt.NewFunction(name, $2.value, $6, $8, $10, $4)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $11))

            yylex.(*Parser).comments.AddComments(name, $3.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

is_reference:
        /* empty */                                     { $$ = boolWithToken{false, nil} }
    |   '&'                                             { $$ = boolWithToken{true, &$1} }
;

is_variadic:
        /* empty */                                     { $$ = boolWithToken{false, nil} }
    |   T_ELLIPSIS                                      { $$ = boolWithToken{true, &$1} }
;

class_declaration_statement:
    class_modifiers T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
        {
            name := node.NewIdentifier($3.Value)
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
            $$ = stmt.NewClass(name, $1, nil, $4, $5, $8, $6)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewOptionalListTokensPosition($1, $2, $9))
            
            yylex.(*Parser).comments.AddComments(name, $3.Comments())
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($1))
        }
    |   T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
        {
            name := node.NewIdentifier($2.Value)
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
            $$ = stmt.NewClass(name, nil, nil, $3, $4, $7, $5)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $8))

            yylex.(*Parser).comments.AddComments(name, $2.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

class_modifiers:
        class_modifier                                  { $$ = []node.Node{$1} }
    |   class_modifiers class_modifier                  { $$ = append($1, $2) }
;

class_modifier:
    T_ABSTRACT
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_FINAL
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

trait_declaration_statement:
    T_TRAIT T_STRING backup_doc_comment '{' class_statement_list '}'
        {
            name := node.NewIdentifier($2.Value)
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
            $$ = stmt.NewTrait(name, $5, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $6))

            yylex.(*Parser).comments.AddComments(name, $2.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

interface_declaration_statement:
    T_INTERFACE T_STRING interface_extends_list backup_doc_comment '{' class_statement_list '}'
        {
            name := node.NewIdentifier($2.Value)
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
            $$ = stmt.NewInterface(name, $3, $6, $4)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $7))
            
            yylex.(*Parser).comments.AddComments(name, $2.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

extends_from:
        /* empty */                                     { $$ = nil }
    |   T_EXTENDS name                                  { $$ = $2; }
;

interface_extends_list:
        /* empty */                                     { $$ = nil }
    |   T_EXTENDS name_list                             { $$ = $2 }
;

implements_list:
        /* empty */                                     { $$ = nil }
    |   T_IMPLEMENTS name_list                          { $$ = $2 }
;

foreach_variable:
        variable                                        { $$ = foreachVariable{$1, false} }
    |   '&' variable                                    { $$ = foreachVariable{$2, true} }
    |   T_LIST '(' array_pair_list ')'
        {
            list := expr.NewList($3)
            yylex.(*Parser).positions.AddPosition(list, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            $$ = foreachVariable{list, false}
            yylex.(*Parser).comments.AddComments(list, $1.Comments())
        }
    |   '[' array_pair_list ']'
        {
            list := expr.NewShortList($2)
            yylex.(*Parser).positions.AddPosition(list, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            $$ = foreachVariable{list, false}
            yylex.(*Parser).comments.AddComments(list, $1.Comments())
        }
;

for_statement:
        statement
            { $$ = altSyntaxNode{$1, false} }
    |   ':' inner_statement_list T_ENDFOR ';'
            {
                $$ = altSyntaxNode{stmt.NewStmtList($2), true}
                yylex.(*Parser).positions.AddPosition($$.node, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            }
;

foreach_statement:
        statement
            { $$ = altSyntaxNode{$1, false} }
    |   ':' inner_statement_list T_ENDFOREACH ';'
            {
                $$ = altSyntaxNode{stmt.NewStmtList($2), true}
                yylex.(*Parser).positions.AddPosition($$.node, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            }
;

declare_statement:
        statement                                       { $$ = $1; }
    |   ':' inner_statement_list T_ENDDECLARE ';'
        {
            $$ = stmt.NewStmtList($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

switch_case_list:
        '{' case_list '}'                               { $$ = &nodesWithEndToken{$2, $3} }
    |   '{' ';' case_list '}'                           { $$ = &nodesWithEndToken{$3, $4} }
    |   ':' case_list T_ENDSWITCH ';'                   { $$ = &nodesWithEndToken{$2, $4} }
    |   ':' ';' case_list T_ENDSWITCH ';'               { $$ = &nodesWithEndToken{$3, $5} }
;

case_list:
        /* empty */                                     { $$ = []node.Node{} }
    |   case_list T_CASE expr case_separator inner_statement_list
            {
                _case := stmt.NewCase($3, $5)
                yylex.(*Parser).positions.AddPosition(_case, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($2, $5))
                $$ = append($1, _case)
                yylex.(*Parser).comments.AddComments(_case, $2.Comments())
            }
    |   case_list T_DEFAULT case_separator inner_statement_list
            {
                _default := stmt.NewDefault($4)
                yylex.(*Parser).positions.AddPosition(_default, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($2, $4))
                $$ = append($1, _default)
                yylex.(*Parser).comments.AddComments(_default, $2.Comments())
            }
;

case_separator:
        ':'
    |   ';'
;

while_statement:
        statement
            { $$ = altSyntaxNode{$1, false} }
    |   ':' inner_statement_list T_ENDWHILE ';'
            {
                $$ = altSyntaxNode{stmt.NewStmtList($2), true}
                yylex.(*Parser).positions.AddPosition($$.node, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            }
;

if_stmt_without_else:
    T_IF '(' expr ')' statement
        {
            $$ = stmt.NewIf($3, $5, nil, nil)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $5))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   if_stmt_without_else T_ELSEIF '(' expr ')' statement
        { 
            _elseIf := stmt.NewElseIf($4, $6)
            yylex.(*Parser).positions.AddPosition(_elseIf, yylex.(*Parser).positionBuilder.NewTokenNodePosition($2, $6))
            $$ = $1.(*stmt.If).AddElseIf(_elseIf)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $6))

            yylex.(*Parser).comments.AddComments(_elseIf, $2.Comments())
        }
;

if_stmt:
        if_stmt_without_else %prec T_NOELSE             { $$ = $1; }
    |   if_stmt_without_else T_ELSE statement
        {
            _else := stmt.NewElse($3)
            yylex.(*Parser).positions.AddPosition(_else, yylex.(*Parser).positionBuilder.NewTokenNodePosition($2, $3))
            $$ = $1.(*stmt.If).SetElse(_else)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

            yylex.(*Parser).comments.AddComments($$, $2.Comments())
        }
;

alt_if_stmt_without_else:
    T_IF '(' expr ')' ':' inner_statement_list
        { 
            stmts := stmt.NewStmtList($6)
            yylex.(*Parser).positions.AddPosition(stmts, yylex.(*Parser).positionBuilder.NewNodeListPosition($6))
            $$ = stmt.NewAltIf($3, stmts, nil, nil)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $6))

            yylex.(*Parser).comments.AddComments(stmts, $5.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   alt_if_stmt_without_else T_ELSEIF '(' expr ')' ':' inner_statement_list
        {
            stmts := stmt.NewStmtList($7)
            yylex.(*Parser).positions.AddPosition(stmts, yylex.(*Parser).positionBuilder.NewNodeListPosition($7))
            _elseIf := stmt.NewAltElseIf($4, stmts)
            yylex.(*Parser).positions.AddPosition(_elseIf, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($2, $7))
            $$ = $1.(*stmt.AltIf).AddElseIf(_elseIf)

            yylex.(*Parser).comments.AddComments(stmts, $6.Comments())
            yylex.(*Parser).comments.AddComments(_elseIf, $2.Comments())
        }
;

alt_if_stmt:
    alt_if_stmt_without_else T_ENDIF ';'
        {
            $$ = $1
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))
        }
    |   alt_if_stmt_without_else T_ELSE ':' inner_statement_list T_ENDIF ';'
        {
            stmts := stmt.NewStmtList($4)
            yylex.(*Parser).positions.AddPosition(stmts, yylex.(*Parser).positionBuilder.NewNodeListPosition($4))
            _else := stmt.NewAltElse(stmts)
            yylex.(*Parser).positions.AddPosition(_else, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($2, $4))
            $$ = $1.(*stmt.AltIf).SetElse(_else)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $6))

            yylex.(*Parser).comments.AddComments(stmts, $3.Comments())
            yylex.(*Parser).comments.AddComments(_else, $2.Comments())
        }
;

parameter_list:
        non_empty_parameter_list                        { $$ = $1; }
    |   /* empty */                                     { $$ = nil }
;

non_empty_parameter_list:
        parameter                                       { $$ = []node.Node{$1} }
    |   non_empty_parameter_list ',' parameter          { $$ = append($1, $3) }
;

parameter:
    optional_type is_reference is_variadic T_VARIABLE
        {
            identifier := node.NewIdentifier(strings.TrimLeft($4.Value, "$"))
            yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
            variable := expr.NewVariable(identifier)
            yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($4))

            yylex.(*Parser).comments.AddComments($$, $4.Comments())
            yylex.(*Parser).comments.AddComments($$, $4.Comments())
            
            if $1 != nil {
                $$ = node.NewParameter($1, variable, nil, $2.value, $3.value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            } else if $2.value == true {
                $$ = node.NewParameter($1, variable, nil, $2.value, $3.value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition(*$2.token, $4))
                yylex.(*Parser).comments.AddComments($$, $2.token.Comments())
            } else if $3.value == true {
                $$ = node.NewParameter($1, variable, nil, $2.value, $3.value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition(*$3.token, $4))
                yylex.(*Parser).comments.AddComments($$, $3.token.Comments())
            } else {
                $$ = node.NewParameter($1, variable, nil, $2.value, $3.value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                yylex.(*Parser).comments.AddComments($$, $4.Comments())
            }
        }
    |   optional_type is_reference is_variadic T_VARIABLE '=' expr
        {
            identifier := node.NewIdentifier(strings.TrimLeft($4.Value, "$"))
            yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
            variable := expr.NewVariable(identifier)
            yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($4))

            yylex.(*Parser).comments.AddComments($$, $4.Comments())
            yylex.(*Parser).comments.AddComments($$, $4.Comments())

            if $1 != nil {
                $$ = node.NewParameter($1, variable, $6, $2.value, $3.value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $6))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            } else if $2.value == true {
                $$ = node.NewParameter($1, variable, $6, $2.value, $3.value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition(*$2.token, $6))
                yylex.(*Parser).comments.AddComments($$, $2.token.Comments())
            } else if $3.value == true {
                $$ = node.NewParameter($1, variable, $6, $2.value, $3.value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition(*$3.token, $6))
                yylex.(*Parser).comments.AddComments($$, $3.token.Comments())
            } else {
                $$ = node.NewParameter($1, variable, $6, $2.value, $3.value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($4, $6))
                yylex.(*Parser).comments.AddComments($$, $4.Comments())
            }
        }
;

optional_type:
        /* empty */                                     { $$ = nil }
    |   type_expr                                       { $$ = $1; }
;

type_expr:
    type                                                { $$ = $1; }
    |   '?' type
        {
            $$ = node.NewNullable($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

type:
    T_ARRAY
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_CALLABLE
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   name                                            { $$ = $1; }
;

return_type:
        /* empty */                                     { $$ = nil }
    |   ':' type_expr                                   { $$ = $2; }
;

argument_list:
        '(' ')'                                         { $$ = &nodesWithEndToken{[]node.Node{}, $2} }
    |   '(' non_empty_argument_list possible_comma ')'  { $$ = &nodesWithEndToken{$2, $4} }
;

non_empty_argument_list:
        argument                                        { $$ = []node.Node{$1} }
    |   non_empty_argument_list ',' argument            { $$ = append($1, $3) }
;

argument:
    expr
        {
            $$ = node.NewArgument($1, false, false)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodePosition($1))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   T_ELLIPSIS expr
        {
            $$ = node.NewArgument($2, true, false)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

global_var_list:
        global_var_list ',' global_var                  { $$ = append($1, $3); }
    |   global_var                                      { $$ = []node.Node{$1} }
;

global_var:
    simple_variable                                     { $$ = $1 }
;

static_var_list:
        static_var_list ',' static_var                  { $$ = append($1, $3) }
    |   static_var                                      { $$ = []node.Node{$1} }
;

static_var:
        T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$ = stmt.NewStaticVar(variable, nil)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                yylex.(*Parser).comments.AddComments(identifier, $1.Comments())
                yylex.(*Parser).comments.AddComments(variable, $1.Comments())
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_VARIABLE '=' expr
        {
            identifier := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
            yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            variable := expr.NewVariable(identifier)
            yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = stmt.NewStaticVar(variable, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))

            yylex.(*Parser).comments.AddComments(identifier, $1.Comments())
            yylex.(*Parser).comments.AddComments(variable, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

class_statement_list:
        class_statement_list class_statement            { $$ = append($1, $2) }
    |   /* empty */                                     { $$ = []node.Node{} }
;

class_statement:
    variable_modifiers property_list ';'
        {
            $$ = stmt.NewPropertyList($1, $2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($1))
        }
    |   method_modifiers T_CONST class_const_list ';'
        {
            $$ = stmt.NewClassConstList($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewOptionalListTokensPosition($1, $2, $4))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($1))
        }
    |   T_USE name_list trait_adaptations
        {
            $$ = stmt.NewTraitUse($2, $3.nodes)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3.endToken))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   method_modifiers T_FUNCTION returns_ref identifier backup_doc_comment '(' parameter_list ')' return_type method_body
            {
                name := node.NewIdentifier($4.Value)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                $$ = stmt.NewClassMethod(name, $1, $3.value, $7, $9, $10.nodes, $5)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewOptionalListTokensPosition($1, $2, $10.endToken))
                
                yylex.(*Parser).comments.AddComments(name, $4.Comments())
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($1))
            }
;

name_list:
        name                                            { $$ = []node.Node{$1} }
    |   name_list ',' name                              { $$ = append($1, $3) }
;

trait_adaptations:
        ';'                                             { $$ = &nodesWithEndToken{nil, $1} }
    |   '{' '}'                                         { $$ = &nodesWithEndToken{nil, $2} }
    |   '{' trait_adaptation_list '}'                   { $$ = &nodesWithEndToken{$2, $3} }
;

trait_adaptation_list:
        trait_adaptation                                { $$ = []node.Node{$1} }
    |   trait_adaptation_list trait_adaptation          { $$ = append($1, $2) }
;

trait_adaptation:
        trait_precedence ';'                            { $$ = $1; }
    |   trait_alias ';'                                 { $$ = $1; }
;

trait_precedence:
    absolute_trait_method_reference T_INSTEADOF name_list
        {
            $$ = stmt.NewTraitUsePrecedence($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeNodeListPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
;

trait_alias:
    trait_method_reference T_AS T_STRING
        {
            alias := node.NewIdentifier($3.Value)
            yylex.(*Parser).positions.AddPosition(alias, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
            $$ = stmt.NewTraitUseAlias($1, nil, alias)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))
            
            yylex.(*Parser).comments.AddComments(alias, $3.Comments())
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   trait_method_reference T_AS reserved_non_modifiers
        {
            alias := node.NewIdentifier($3.Value)
            yylex.(*Parser).positions.AddPosition(alias, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
            $$ = stmt.NewTraitUseAlias($1, nil, alias)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))
            
            yylex.(*Parser).comments.AddComments(alias, $3.Comments())
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   trait_method_reference T_AS member_modifier identifier
        {
            alias := node.NewIdentifier($4.Value)
            yylex.(*Parser).positions.AddPosition(alias, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
            $$ = stmt.NewTraitUseAlias($1, $3, alias)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
            
            yylex.(*Parser).comments.AddComments(alias, $4.Comments())
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   trait_method_reference T_AS member_modifier
        {
            $$ = stmt.NewTraitUseAlias($1, $3, nil)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
;

trait_method_reference:
    identifier
        {
            name := node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = stmt.NewTraitMethodRef(nil, name)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            
            yylex.(*Parser).comments.AddComments(name, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   absolute_trait_method_reference                 { $$ = $1; }
;

absolute_trait_method_reference:
    name T_PAAMAYIM_NEKUDOTAYIM identifier
        {
            target := node.NewIdentifier($3.Value)
            yylex.(*Parser).positions.AddPosition(target, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
            $$ = stmt.NewTraitMethodRef($1, target)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))
            
            yylex.(*Parser).comments.AddComments(target, $3.Comments())
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
;

method_body:
        ';' /* abstract method */                       { $$ = &nodesWithEndToken{nil, $1} }
    |   '{' inner_statement_list '}'                    { $$ = &nodesWithEndToken{$2, $3} }
;

variable_modifiers:
        non_empty_member_modifiers                      { $$ = $1; }
    |   T_VAR
        {
            modifier := node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition(modifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = []node.Node{modifier}
            yylex.(*Parser).comments.AddComments(modifier, $1.Comments())
        }
;

method_modifiers:
        /* empty */                                     { $$ = nil }
    |   non_empty_member_modifiers                      { $$ = $1 }
;

non_empty_member_modifiers:
        member_modifier	                                { $$ = []node.Node{$1} }
    |   non_empty_member_modifiers member_modifier      { $$ = append($1, $2) }
;

member_modifier:
    T_PUBLIC
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_PROTECTED
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_PRIVATE
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_STATIC
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_ABSTRACT
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_FINAL
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

property_list:
        property_list ',' property                      { $$ = append($1, $3) }
    |   property                                        { $$ = []node.Node{$1} }
;

property:
    T_VARIABLE backup_doc_comment
        {
            identifier := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
            yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            variable := expr.NewVariable(identifier)
            yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = stmt.NewProperty(variable, nil, $2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))

            yylex.(*Parser).comments.AddComments(identifier, $1.Comments())
            yylex.(*Parser).comments.AddComments(variable, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_VARIABLE '=' expr backup_doc_comment
        {
            identifier := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
            yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            variable := expr.NewVariable(identifier)
            yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = stmt.NewProperty(variable, $3, $4)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))

            yylex.(*Parser).comments.AddComments(identifier, $1.Comments())
            yylex.(*Parser).comments.AddComments(variable, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

class_const_list:
        class_const_list ',' class_const_decl           { $$ = append($1, $3) }
    |   class_const_decl                                { $$ = []node.Node{$1} }
;

class_const_decl:
    identifier '=' expr backup_doc_comment
        {
            name := node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = stmt.NewConstant(name, $3, $4)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))

            yylex.(*Parser).comments.AddComments(name, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

const_decl:
    T_STRING '=' expr backup_doc_comment
        {
            name := node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = stmt.NewConstant(name, $3, $4)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))

            yylex.(*Parser).comments.AddComments(name, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

echo_expr_list:
        echo_expr_list ',' echo_expr                    { $$ = append($1, $3) }
    |   echo_expr                                       { $$ = []node.Node{$1} }
;

echo_expr:
    expr                                                { $$ = $1 }
;

for_exprs:
        /* empty */                                     { $$ = nil; }
    |   non_empty_for_exprs                             { $$ = $1; }
;
non_empty_for_exprs:
        non_empty_for_exprs ',' expr                    { $$ = append($1, $3) }
    |   expr                                            { $$ = []node.Node{$1} }
;

anonymous_class:
    T_CLASS ctor_arguments extends_from implements_list backup_doc_comment '{' class_statement_list '}'
        {
            if $2 != nil {
                $$ = stmt.NewClass(nil, nil, $2.nodes, $3, $4, $7, $5)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $8))
            } else {
                $$ = stmt.NewClass(nil, nil, nil, $3, $4, $7, $5)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $8))
            }

            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

new_expr:
    T_NEW class_name_reference ctor_arguments
        {
            if $3 != nil {
                $$ = expr.NewNew($2, $3.nodes)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3.endToken))
            } else {
                $$ = expr.NewNew($2, nil)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            }

            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_NEW anonymous_class                           { $$ = expr.NewNew($2, nil) }
;

expr_without_variable:
    T_LIST '(' array_pair_list ')' '=' expr
        {
            list := expr.NewList($3)
            yylex.(*Parser).positions.AddPosition(list, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            $$ = assign.NewAssign(list, $6)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $6))

            yylex.(*Parser).comments.AddComments(list, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   '[' array_pair_list ']' '=' expr
        {
            shortList := expr.NewShortList($2)
            yylex.(*Parser).positions.AddPosition(shortList, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            $$ = assign.NewAssign(shortList, $5)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $5))

            yylex.(*Parser).comments.AddComments(shortList, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   variable '=' expr
        {
            $$ = assign.NewAssign($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable '=' '&' expr
        {
            $$ = assign.NewReference($1, $4)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   T_CLONE expr
        {
            $$ = expr.NewClone($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   variable T_PLUS_EQUAL expr
        {
            $$ = assign.NewPlus($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable T_MINUS_EQUAL expr
        {
            $$ = assign.NewMinus($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable T_MUL_EQUAL expr
        {
            $$ = assign.NewMul($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable T_POW_EQUAL expr
        {
            $$ = assign.NewPow($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable T_DIV_EQUAL expr
        {
            $$ = assign.NewDiv($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable T_CONCAT_EQUAL expr
        {
            $$ = assign.NewConcat($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable T_MOD_EQUAL expr
        {
            $$ = assign.NewMod($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable T_AND_EQUAL expr
        {
            $$ = assign.NewBitwiseAnd($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable T_OR_EQUAL expr
        {
            $$ = assign.NewBitwiseOr($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable T_XOR_EQUAL expr
        {
            $$ = assign.NewBitwiseXor($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable T_SL_EQUAL expr
        {
            $$ = assign.NewShiftLeft($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable T_SR_EQUAL expr
        {
            $$ = assign.NewShiftRight($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable T_INC
        {
            $$ = expr.NewPostInc($1)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   T_INC variable
        {
            $$ = expr.NewPreInc($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   variable T_DEC
        {
            $$ = expr.NewPostDec($1)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   T_DEC variable
        {
            $$ = expr.NewPreDec($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   expr T_BOOLEAN_OR expr
        {
            $$ = binary.NewBooleanOr($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_BOOLEAN_AND expr
        {
            $$ = binary.NewBooleanAnd($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_LOGICAL_OR expr
        {
            $$ = binary.NewLogicalOr($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_LOGICAL_AND expr
        {
            $$ = binary.NewLogicalAnd($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_LOGICAL_XOR expr
        {
            $$ = binary.NewLogicalXor($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr '|' expr
        {
            $$ = binary.NewBitwiseOr($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr '&' expr
        {
            $$ = binary.NewBitwiseAnd($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr '^' expr
        {
            $$ = binary.NewBitwiseXor($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr '.' expr
        {
            $$ = binary.NewConcat($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr '+' expr
        {
            $$ = binary.NewPlus($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr '-' expr
        {
            $$ = binary.NewMinus($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr '*' expr
        {
            $$ = binary.NewMul($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_POW expr
        {
            $$ = binary.NewPow($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr '/' expr
        {
            $$ = binary.NewDiv($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr '%' expr
        {
            $$ = binary.NewMod($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_SL expr
        {
            $$ = binary.NewShiftLeft($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_SR expr
        {
            $$ = binary.NewShiftRight($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   '+' expr %prec T_INC
        {
            $$ = expr.NewUnaryPlus($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   '-' expr %prec T_INC
        {
            $$ = expr.NewUnaryMinus($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   '!' expr
        {
            $$ = expr.NewBooleanNot($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   '~' expr
        {
            $$ = expr.NewBitwiseNot($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   expr T_IS_IDENTICAL expr
        {
            $$ = binary.NewIdentical($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_IS_NOT_IDENTICAL expr
        {
            $$ = binary.NewNotIdentical($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_IS_EQUAL expr
        {
            $$ = binary.NewEqual($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_IS_NOT_EQUAL expr
        {
            $$ = binary.NewNotEqual($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr '<' expr
        {
            $$ = binary.NewSmaller($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_IS_SMALLER_OR_EQUAL expr
        {
            $$ = binary.NewSmallerOrEqual($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr '>' expr
        {
            $$ = binary.NewGreater($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_IS_GREATER_OR_EQUAL expr
        {
            $$ = binary.NewGreaterOrEqual($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_SPACESHIP expr
        {
            $$ = binary.NewSpaceship($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_INSTANCEOF class_name_reference
        {
            $$ = expr.NewInstanceOf($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   '(' expr ')'                                    { $$ = $2; }
    |   new_expr                                        { $$ = $1; }
    |   expr '?' expr ':' expr
        {
            $$ = expr.NewTernary($1, $3, $5)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $5))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr '?' ':' expr
        {
            $$ = expr.NewTernary($1, nil, $4)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_COALESCE expr
        {
            $$ = binary.NewCoalesce($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   internal_functions_in_yacc                      { $$ = $1}
    |   T_INT_CAST expr
        {
            $$ = cast.NewInt($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_DOUBLE_CAST expr
        {
            $$ = cast.NewDouble($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_STRING_CAST expr
        {
            $$ = cast.NewString($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_ARRAY_CAST expr
        {
            $$ = cast.NewArray($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_OBJECT_CAST expr
        {
            $$ = cast.NewObject($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_BOOL_CAST expr
        {
            $$ = cast.NewBool($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_UNSET_CAST expr
        {
            $$ = cast.NewUnset($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_EXIT exit_expr
        {
            if (strings.EqualFold($1.Value, "die")) {
                $$ = expr.NewDie($2)
            } else {
                $$ = expr.NewExit($2)
            }
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   '@' expr
        {
            $$ = expr.NewErrorSuppress($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   scalar                                          { $$ = $1; }
    |   '`' backticks_expr '`'
        {
            $$ = expr.NewShellExec($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_PRINT expr
        {
            $$ = expr.NewPrint($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_YIELD
        {
            $$ = expr.NewYield(nil, nil)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_YIELD expr
        {
            $$ = expr.NewYield(nil, $2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_YIELD expr T_DOUBLE_ARROW expr
        {
            $$ = expr.NewYield($2, $4)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_YIELD_FROM expr
        {
            $$ = expr.NewYieldFrom($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                $$ = expr.NewClosure($5, $7, $8, $10, false, $2.value, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $11))
                
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_STATIC T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                $$ = expr.NewClosure($6, $8, $9, $11, true, $3.value, $4)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $12))
                
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

backup_doc_comment:
	/* empty */ { $$ = yylex.(*Parser).PhpDocComment; yylex.(*Parser).PhpDocComment = "" }
;

returns_ref:
        /* empty */                                     { $$ = boolWithToken{false, nil} }
    |   '&'                                             { $$ = boolWithToken{true, &$1} }
;

lexical_vars:
        /* empty */                                     { $$ = []node.Node{} }
    |   T_USE '(' lexical_var_list ')'                  { $$ = $3; }
;

lexical_var_list:
        lexical_var_list ',' lexical_var                { $$ = append($1, $3) }
    |   lexical_var                                     { $$ = []node.Node{$1} }
;

lexical_var:
    T_VARIABLE
        {
            identifier := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
            yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            variable := expr.NewVariable(identifier)
            yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = expr.NewClosureUse(variable, false)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))

            yylex.(*Parser).comments.AddComments(identifier, $1.Comments())
            yylex.(*Parser).comments.AddComments(variable, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   '&' T_VARIABLE
        {
            identifier := node.NewIdentifier(strings.TrimLeft($2.Value, "$"))
            yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
            variable := expr.NewVariable(identifier)
            yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
            $$ = expr.NewClosureUse(variable, true)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))

            yylex.(*Parser).comments.AddComments(identifier, $2.Comments())
            yylex.(*Parser).comments.AddComments(variable, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

function_call:
    name argument_list
        {
            $$ = expr.NewFunctionCall($1, $2.nodes)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2.endToken))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
        {
            $$ = expr.NewStaticCall($1, $3, $4.nodes)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4.endToken))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
        {
            $$ = expr.NewStaticCall($1, $3, $4.nodes)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4.endToken))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   callable_expr argument_list
        {
            $$ = expr.NewFunctionCall($1, $2.nodes)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2.endToken))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
;

class_name:
    T_STATIC
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   name                                            { $$ = $1; }
;

class_name_reference:
        class_name                                      { $$ = $1; }
    |   new_variable                                    { $$ = $1; }
;

exit_expr:
        /* empty */                                     { $$ = nil }
    |   '(' optional_expr ')'                           { $$ = $2; }
;

backticks_expr:
        /* empty */                                     { $$ = []node.Node{} }
    |   T_ENCAPSED_AND_WHITESPACE                       { $$ = []node.Node{scalar.NewEncapsedStringPart($1.Value)} }
    |   encaps_list                                     { $$ = $1; }
;

ctor_arguments:
        /* empty */	                                    { $$ = nil }
    |   argument_list                                   { $$ = $1 }
;

dereferencable_scalar:
    T_ARRAY '(' array_pair_list ')'
        {
            $$ = expr.NewArray($3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   '[' array_pair_list ']'
        {
            $$ = expr.NewShortArray($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_CONSTANT_ENCAPSED_STRING
        {
            $$ = scalar.NewString($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

scalar:
    T_LNUMBER
        {
            $$ = scalar.NewLnumber($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_DNUMBER
        {
            $$ = scalar.NewDnumber($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_LINE
        {
            $$ = scalar.NewMagicConstant($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_FILE
        {
            $$ = scalar.NewMagicConstant($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_DIR
        {
            $$ = scalar.NewMagicConstant($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_TRAIT_C
        {
            $$ = scalar.NewMagicConstant($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_METHOD_C
        {
            $$ = scalar.NewMagicConstant($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_FUNC_C
        {
            $$ = scalar.NewMagicConstant($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_NS_C
        {
            $$ = scalar.NewMagicConstant($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_CLASS_C
        {
            $$ = scalar.NewMagicConstant($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC 
        {
            encapsed := scalar.NewEncapsedStringPart($2.Value)
            yylex.(*Parser).positions.AddPosition(encapsed, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
            yylex.(*Parser).comments.AddComments(encapsed, $2.Comments())

            $$ = scalar.NewHeredoc($1.Value, []node.Node{encapsed})
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_START_HEREDOC T_END_HEREDOC
        {
            $$ = scalar.NewHeredoc($1.Value, nil)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   '"' encaps_list '"'
        {
            $$ = scalar.NewEncapsed($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_START_HEREDOC encaps_list T_END_HEREDOC
        {
            $$ = scalar.NewHeredoc($1.Value, $2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   dereferencable_scalar                           { $$ = $1; }
    |   constant                                        { $$ = $1; }
;

constant:
    name
        {
            $$ = expr.NewConstFetch($1)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodePosition($1))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM identifier
        {
            target := node.NewIdentifier($3.Value)
            yylex.(*Parser).positions.AddPosition(target, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
            $$ = expr.NewClassConstFetch($1, target)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))

            yylex.(*Parser).comments.AddComments(target, $3.Comments())
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM identifier
        {
            target := node.NewIdentifier($3.Value)
            yylex.(*Parser).positions.AddPosition(target, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
            $$ = expr.NewClassConstFetch($1, target)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))

            yylex.(*Parser).comments.AddComments(target, $3.Comments())
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
;

expr:
        variable                                        { $$ = $1; }
    |   expr_without_variable                           { $$ = $1; }
;

optional_expr:
        /* empty */                                     { $$ = nil }
    |   expr                                            { $$ = $1; }
;

variable_class_name:
    dereferencable                                      { $$ = $1; }
;

dereferencable:
        variable                                        { $$ = $1; }
    |   '(' expr ')'                                    { $$ = $2; }
    |   dereferencable_scalar                           { $$ = $1; }
;

callable_expr:
        callable_variable                               { $$ = $1; }
    |   '(' expr ')'                                    { $$ = $2; }
    |   dereferencable_scalar                           { $$ = $1; }
;

callable_variable:
    simple_variable                                     { $$ = $1; }
    |   dereferencable '[' optional_expr ']'
        {
            $$ = expr.NewArrayDimFetch($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   constant '[' optional_expr ']'
        {
            $$ = expr.NewArrayDimFetch($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   dereferencable '{' expr '}'
        {
            $$ = expr.NewArrayDimFetch($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   dereferencable T_OBJECT_OPERATOR property_name argument_list
        {
            $$ = expr.NewMethodCall($1, $3, $4.nodes)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4.endToken))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   function_call                                   { $$ = $1; }
;

variable:
        callable_variable                               { $$ = $1; }
    |   static_member                                   { $$ = $1; }
    |   dereferencable T_OBJECT_OPERATOR property_name
        {
            $$ = expr.NewPropertyFetch($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
;

simple_variable:
    T_VARIABLE
        {
            name := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = expr.NewVariable(name)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            
            yylex.(*Parser).comments.AddComments(name, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   '$' '{' expr '}'
        {
            $$ = expr.NewVariable($3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   '$' simple_variable
        {
            $$ = expr.NewVariable($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

static_member:
    class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
        {
            $$ = expr.NewStaticPropertyFetch($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
        {
            $$ = expr.NewStaticPropertyFetch($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
;

new_variable:
        simple_variable                                 { $$ = $1 }
    |   new_variable '[' optional_expr ']'
        {
            $$ = expr.NewArrayDimFetch($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   new_variable '{' expr '}'
        {
            $$ = expr.NewArrayDimFetch($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   new_variable T_OBJECT_OPERATOR property_name
        {
            $$ = expr.NewPropertyFetch($1, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = expr.NewStaticPropertyFetch($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   new_variable T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = expr.NewStaticPropertyFetch($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
;

member_name:
    identifier
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   '{' expr '}'                                    { $$ = $2; }
    |   simple_variable                                 { $$ = $1 }
;

property_name:
    T_STRING
        {
            $$ = node.NewIdentifier($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   '{' expr '}'                                    { $$ = $2; }
    |   simple_variable                                 { $$ = $1 }
;

array_pair_list:
    non_empty_array_pair_list
        {
            if ($1[len($1)-1] == nil) {
                $$ = $1[:len($1)-1]
            } else {
                $$ = $1
            }
        }
;

possible_array_pair:
        /* empty */                                     { $$ = nil }
    |   array_pair                                      { $$ = $1; }
;

non_empty_array_pair_list:
        non_empty_array_pair_list ',' possible_array_pair
                                                        { $$ = append($1, $3) }
    |   possible_array_pair                             { $$ = []node.Node{$1} }
;

array_pair:
    expr T_DOUBLE_ARROW expr
        {
            $$ = expr.NewArrayItem($1, $3, false)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr
        {
            $$ = expr.NewArrayItem(nil, $1, false)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodePosition($1))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   expr T_DOUBLE_ARROW '&' variable
        {
            $$ = expr.NewArrayItem($1, $4, true)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
        }
    |   '&' variable
        {
            $$ = expr.NewArrayItem(nil, $2, true)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   expr T_DOUBLE_ARROW T_LIST '(' array_pair_list ')'
            {
                // TODO: Cannot use list() as standalone expression
                list := expr.NewList($5)
                yylex.(*Parser).positions.AddPosition(list, yylex.(*Parser).positionBuilder.NewTokensPosition($3, $6))
                $$ = expr.NewArrayItem($1, list, false)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $6))

                yylex.(*Parser).comments.AddComments(list, $3.Comments())
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   T_LIST '(' array_pair_list ')'
            {
                // TODO: Cannot use list() as standalone expression
                list := expr.NewList($3)
                yylex.(*Parser).positions.AddPosition(list, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                $$ = expr.NewArrayItem(nil, list, false)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                
                yylex.(*Parser).comments.AddComments(list, $1.Comments())
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

encaps_list:
        encaps_list encaps_var                          { $$ = append($1, $2) }
    |   encaps_list T_ENCAPSED_AND_WHITESPACE
        {
            encapsed := scalar.NewEncapsedStringPart($2.Value)
            yylex.(*Parser).positions.AddPosition(encapsed, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
            $$ = append($1, encapsed)
            yylex.(*Parser).comments.AddComments(encapsed, $2.Comments())
        }
    |   encaps_var                                      { $$ = []node.Node{$1} }
    |   T_ENCAPSED_AND_WHITESPACE encaps_var
        {
            encapsed := scalar.NewEncapsedStringPart($1.Value)
            yylex.(*Parser).positions.AddPosition(encapsed, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = []node.Node{encapsed, $2}
            yylex.(*Parser).comments.AddComments(encapsed, $1.Comments())
        }
;

encaps_var:
    T_VARIABLE
        {
            name := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = expr.NewVariable(name)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))

            yylex.(*Parser).comments.AddComments(name, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_VARIABLE '[' encaps_var_offset ']'
        {
            identifier := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
            yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            variable := expr.NewVariable(identifier)
            yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = expr.NewArrayDimFetch(variable, $3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

            yylex.(*Parser).comments.AddComments(identifier, $1.Comments())
            yylex.(*Parser).comments.AddComments(variable, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_VARIABLE T_OBJECT_OPERATOR T_STRING
        {
            identifier := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
            yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            variable := expr.NewVariable(identifier)
            yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            fetch := node.NewIdentifier($3.Value)
            yylex.(*Parser).positions.AddPosition(fetch, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
            $$ = expr.NewPropertyFetch(variable, fetch)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            
            yylex.(*Parser).comments.AddComments(identifier, $1.Comments())
            yylex.(*Parser).comments.AddComments(variable, $1.Comments())
            yylex.(*Parser).comments.AddComments(fetch, $3.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_DOLLAR_OPEN_CURLY_BRACES expr '}'
        {
            $$ = expr.NewVariable($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}'
        {
            name := node.NewIdentifier($2.Value)
            yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
            $$ = expr.NewVariable(name)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

            yylex.(*Parser).comments.AddComments(name, $2.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
        {
            identifier := node.NewIdentifier($2.Value)
            yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
            variable := expr.NewVariable(identifier)
            yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
            $$ = expr.NewArrayDimFetch(variable, $4)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $6))


            yylex.(*Parser).comments.AddComments(identifier, $2.Comments())
            yylex.(*Parser).comments.AddComments(variable, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_CURLY_OPEN variable '}'                       { $$ = $2; }
;
encaps_var_offset:
    T_STRING
        {
            $$ = scalar.NewString($1.Value)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_NUM_STRING
        {
            // TODO: add option to handle 64 bit integer
            if _, err := strconv.Atoi($1.Value); err == nil {
                $$ = scalar.NewLnumber($1.Value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            } else {
                $$ = scalar.NewString($1.Value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            }
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   '-' T_NUM_STRING
        {
            // TODO: add option to handle 64 bit integer
            if _, err := strconv.Atoi($2.Value); err == nil {
                lnumber := scalar.NewLnumber($2.Value)
                yylex.(*Parser).positions.AddPosition(lnumber, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))
                $$ = expr.NewUnaryMinus(lnumber)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))
                
                yylex.(*Parser).comments.AddComments(lnumber, $1.Comments())
            } else {
                $2.Value = "-"+$2.Value
                $$ = scalar.NewString($2.Value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))
            }

            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_VARIABLE
        {
            identifier := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
            yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
            $$ = expr.NewVariable(identifier)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))

            yylex.(*Parser).comments.AddComments(identifier, $1.Comments())
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

internal_functions_in_yacc:
    T_ISSET '(' isset_variables possible_comma ')'
        {
            $$ = expr.NewIsset($3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $5))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_EMPTY '(' expr ')'
        {
            $$ = expr.NewEmpty($3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_INCLUDE expr
        {
            $$ = expr.NewInclude($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_INCLUDE_ONCE expr
        {
            $$ = expr.NewIncludeOnce($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_EVAL '(' expr ')'
        {
            $$ = expr.NewEval($3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_REQUIRE expr
        {
            $$ = expr.NewRequire($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_REQUIRE_ONCE expr
        {
            $$ = expr.NewRequireOnce($2)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
;

isset_variables:
        isset_variable                                  { $$ = []node.Node{$1} }
    |   isset_variables ',' isset_variable              { $$ = append($1, $3) }
;

isset_variable:
    expr                                                { $$ = $1 }
;

/////////////////////////////////////////////////////////////////////////

%%

type foreachVariable struct {
	node  node.Node
	byRef bool
}

type nodesWithEndToken struct {
	nodes    []node.Node
	endToken token.Token
}

type boolWithToken struct {
	value bool
	token *token.Token
}

type altSyntaxNode struct {
	node  node.Node
	isAlt bool
}