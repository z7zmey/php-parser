%{
package php5

import (
    "strings"
    "strconv"

    "github.com/z7zmey/php-parser/scanner"
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
    token *scanner.Token
    boolWithToken boolWithToken
    list []node.Node
    foreachVariable foreachVariable
    nodesWithEndToken *nodesWithEndToken
    simpleIndirectReference simpleIndirectReference
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

%left T_INCLUDE T_INCLUDE_ONCE T_EVAL T_REQUIRE T_REQUIRE_ONCE
%left ','
%left T_LOGICAL_OR
%left T_LOGICAL_XOR
%left T_LOGICAL_AND
%right T_PRINT
%right T_YIELD
%left '=' T_PLUS_EQUAL T_MINUS_EQUAL T_MUL_EQUAL T_DIV_EQUAL T_CONCAT_EQUAL T_MOD_EQUAL T_AND_EQUAL T_OR_EQUAL T_XOR_EQUAL T_SL_EQUAL T_SR_EQUAL T_POW_EQUAL
%left '?' ':'
%left T_BOOLEAN_OR
%left T_BOOLEAN_AND 
%left '|'
%left '^'
%left '&'
%nonassoc T_IS_EQUAL T_IS_NOT_EQUAL T_IS_IDENTICAL T_IS_NOT_IDENTICAL
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
%left T_ELSEIF
%left T_ELSE 
%left T_ENDIF 
%right T_STATIC T_ABSTRACT T_FINAL T_PRIVATE T_PROTECTED T_PUBLIC

%type <token> function interface_entry

%type <node> top_statement use_declaration use_function_declaration use_const_declaration common_scalar
%type <node> static_class_constant compound_variable reference_variable class_name variable_class_name
%type <node> dim_offset expr expr_without_variable r_variable w_variable rw_variable variable base_variable_with_function_calls
%type <node> base_variable array_function_dereference function_call inner_statement statement unticked_statement
%type <node> statement global_var static_scalar scalar class_constant static_class_name_scalar class_name_scalar
%type <node> encaps_var encaps_var encaps_var_offset general_constant isset_variable internal_functions_in_yacc assignment_list_element
%type <node> variable_name variable_without_objects dynamic_class_name_reference new_expr class_name_reference static_member
%type <node> function_call fully_qualified_class_name combined_scalar combined_scalar_offset general_constant parenthesis_expr
%type <node> exit_expr yield_expr function_declaration_statement class_declaration_statement constant_declaration
%type <node> else_single new_else_single unset_variable declare_statement
%type <node> finally_statement additional_catch unticked_function_declaration_statement unticked_class_declaration_statement
%type <node> optional_class_type parameter class_entry_type extends_from class_statement class_constant_declaration
%type <node> trait_use_statement function_call_parameter trait_adaptation_statement trait_precedence trait_alias
%type <node> trait_method_reference_fully_qualified trait_method_reference trait_modifiers member_modifier method
%type <node> static_scalar_value static_operation

%type <list> top_statement_list namespace_name use_declarations use_function_declarations use_const_declarations
%type <list> inner_statement_list global_var_list static_var_list encaps_list isset_variables non_empty_array_pair_list
%type <list> array_pair_list assignment_list lexical_var_list lexical_vars elseif_list new_elseif_list non_empty_for_expr
%type <list> for_expr case_list echo_expr_list unset_variables declare_list catch_statement additional_catches
%type <list> non_empty_additional_catches parameter_list non_empty_parameter_list class_statement_list implements_list
%type <list> class_statement_list variable_modifiers method_modifiers class_variable_declaration interface_extends_list
%type <list> interface_list non_empty_function_call_parameter_list trait_list trait_adaptation_list non_empty_trait_adaptation_list
%type <list> trait_reference_list non_empty_member_modifiers backticks_expr static_array_pair_list non_empty_static_array_pair_list

%type <list> chaining_dereference chaining_instance_call chaining_method_or_property instance_call variable_property
%type <list> method_or_not array_method_dereference object_property object_dim_list dynamic_class_name_variable_property
%type <list> dynamic_class_name_variable_properties variable_properties

%type <simpleIndirectReference> simple_indirect_reference
%type <foreachVariable> foreach_variable foreach_optional_arg
%type <nodesWithEndToken> ctor_arguments function_call_parameter_list switch_case_list method_body trait_adaptations
%type <boolWithToken> is_reference is_variadic
%type <altSyntaxNode> while_statement for_statement foreach_statement

%%

start:
        top_statement_list
            {
                yylex.(*Parser).rootNode = stmt.NewStmtList($1)
                yylex.(*Parser).positions.AddPosition(yylex.(*Parser).rootNode, yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
            }
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

top_statement:
        error
            {
                // error
                $$ = nil
            }
    |   statement
            { $$ = $1 }
    |   function_declaration_statement
            { $$ = $1 }
    |   class_declaration_statement
            { $$ = $1 }
    |   T_HALT_COMPILER '(' ')' ';'
            {
                $$ = stmt.NewHaltCompiler()
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
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
    |   T_USE use_declarations ';'
            {
                $$ = stmt.NewUseList(nil, $2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_USE T_FUNCTION use_function_declarations ';'
            {
                useType := node.NewIdentifier($2.Value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                yylex.(*Parser).comments.AddComments($$, $2.Comments())

                $$ = stmt.NewUseList(useType, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_USE T_CONST use_const_declarations ';'
            {
                useType := node.NewIdentifier($2.Value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                yylex.(*Parser).comments.AddComments($$, $2.Comments())

                $$ = stmt.NewUseList(useType, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   constant_declaration ';'
            { $$ = $1 }
;

use_declarations:
        use_declarations ',' use_declaration
            { $$ = append($1, $3) }
    |   use_declaration
            { $$ = []node.Node{$1} }
;

use_declaration:
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
    |   T_NS_SEPARATOR namespace_name
            {
                name := name.NewName($2)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$ = stmt.NewUse(nil, name, nil)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListPosition($2))

                yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($2))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($2))
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                name := name.NewName($2)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                alias := node.NewIdentifier($4.Value)
                yylex.(*Parser).positions.AddPosition(alias, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                $$ = stmt.NewUse(nil, name, alias)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($2, $4))

                yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($2))
                yylex.(*Parser).comments.AddComments(alias, $4.Comments())
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($2))
            }
;

use_function_declarations:
        use_function_declarations ',' use_function_declaration
            { $$ = append($1, $3) }
    |   use_function_declaration
            { $$ = []node.Node{$1} }
;

use_function_declaration:
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
    |   T_NS_SEPARATOR namespace_name
            {
                name := name.NewName($2)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$ = stmt.NewUse(nil, name, nil)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListPosition($2))

                yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($2))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($2))
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                name := name.NewName($2)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                alias := node.NewIdentifier($4.Value)
                yylex.(*Parser).positions.AddPosition(alias, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                $$ = stmt.NewUse(nil, name, alias)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($2, $4))

                yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($2))
                yylex.(*Parser).comments.AddComments(alias, $4.Comments())
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($2))
            }
;

use_const_declarations:
        use_const_declarations ',' use_const_declaration
            { $$ = append($1, $3) }
    |   use_const_declaration
            { $$ = []node.Node{$1} }
;

use_const_declaration:
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
    |   T_NS_SEPARATOR namespace_name
            {
                name := name.NewName($2)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$ = stmt.NewUse(nil, name, nil)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListPosition($2))

                yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($2))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($2))
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                name := name.NewName($2)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                alias := node.NewIdentifier($4.Value)
                yylex.(*Parser).positions.AddPosition(alias, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                $$ = stmt.NewUse(nil, name, alias)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($2, $4))

                yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($2))
                yylex.(*Parser).comments.AddComments(alias, $4.Comments())
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($2))
            }
;

constant_declaration:
        constant_declaration ',' T_STRING '=' static_scalar
            {
                name := node.NewIdentifier($3.Value)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                yylex.(*Parser).comments.AddComments(name, $3.Comments())

                constant := stmt.NewConstant(name, $5, "")
                yylex.(*Parser).positions.AddPosition(constant, yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $5))
                yylex.(*Parser).comments.AddComments(constant, $3.Comments())

                constList := $1.(*stmt.ConstList)
                constList.Consts = append(constList.Consts, constant)

                $$ = $1
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeNodeListPosition($1, constList.Consts))
            }
    |   T_CONST T_STRING '=' static_scalar
            {
                name := node.NewIdentifier($2.Value)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                yylex.(*Parser).comments.AddComments(name, $2.Comments())

                constant := stmt.NewConstant(name, $4, "")
                yylex.(*Parser).positions.AddPosition(constant, yylex.(*Parser).positionBuilder.NewTokenNodePosition($2, $4))
                yylex.(*Parser).comments.AddComments(constant, $2.Comments())

                constList := []node.Node{constant}

                $$ = stmt.NewConstList(constList)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, constList))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
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
            { $$ = []node.Node{} }
;


inner_statement:
        error
            {
                // error
                $$ = nil
            }
    |   statement
            { $$ = $1 }
    |   function_declaration_statement
            { $$ = $1 }
    |   class_declaration_statement
            { $$ = $1 }
    |   T_HALT_COMPILER '(' ')' ';'
            {
                $$ = stmt.NewHaltCompiler()
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;


statement:
        unticked_statement
            { $$ = $1 }
    |   T_STRING ':'
            {
                label := node.NewIdentifier($1.Value)
                yylex.(*Parser).positions.AddPosition(label, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$ = stmt.NewLabel(label)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))

                yylex.(*Parser).comments.AddComments(label, $1.Comments())
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

unticked_statement:
        '{' inner_statement_list '}'
            {
                $$ = stmt.NewStmtList($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_IF parenthesis_expr statement elseif_list else_single
            {
                $$ = stmt.NewIf($2, $3, $4, $5)
                
                if $5 != nil {
                    yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $5))
                } else if len($4) > 0 {
                    yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $4))
                } else {
                    yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))
                }

                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_IF parenthesis_expr ':' inner_statement_list new_elseif_list new_else_single T_ENDIF ';'
            {
                stmts := stmt.NewStmtList($4)
                yylex.(*Parser).positions.AddPosition(stmts, yylex.(*Parser).positionBuilder.NewNodeListPosition($4))

                $$ = stmt.NewAltIf($2, stmts, $5, $6)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $8))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_WHILE parenthesis_expr while_statement
            {
                if ($3.isAlt) {
                    $$ = stmt.NewAltWhile($2, $3.node)
                } else {
                    $$ = stmt.NewWhile($2, $3.node)
                }
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3.node))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_DO statement T_WHILE parenthesis_expr ';'
            {
                $$ = stmt.NewDo($2, $4)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $5))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_FOR '(' for_expr ';' for_expr ';' for_expr ')' for_statement
            {
                if ($9.isAlt) {
                    $$ = stmt.NewAltFor($3, $5, $7, $9.node)
                } else {
                    $$ = stmt.NewFor($3, $5, $7, $9.node)
                }
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $9.node))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_SWITCH parenthesis_expr switch_case_list
            {
                if ($3.endToken.Value == ";") {
                    $$ = stmt.NewAltSwitch($2, $3.nodes)
                } else {
                    $$ = stmt.NewSwitch($2, $3.nodes)
                }
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3.endToken))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_BREAK ';'
            {
                $$ = stmt.NewBreak(nil)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_BREAK expr ';'
            {
                $$ = stmt.NewBreak($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_CONTINUE ';'
            {
                $$ = stmt.NewContinue(nil)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_CONTINUE expr ';'
            {
                $$ = stmt.NewContinue($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_RETURN ';'
            {
                $$ = stmt.NewReturn(nil)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_RETURN expr_without_variable ';'
            {
                $$ = stmt.NewReturn($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_RETURN variable ';'
            {
                $$ = stmt.NewReturn($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   yield_expr ';'
            {
                $$ = stmt.NewExpression($1)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
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
    |   T_UNSET '(' unset_variables ')' ';'
            {
                $$ = stmt.NewUnset($3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $5))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_FOREACH '(' variable T_AS foreach_variable foreach_optional_arg ')' foreach_statement
            {
                if $6.node == nil {
                    if ($8.isAlt) {
                        $$ = stmt.NewAltForeach($3, nil, $5.node, $8.node, $5.byRef)
                    } else {
                        $$ = stmt.NewForeach($3, nil, $5.node, $8.node, $5.byRef)
                    }
                } else {
                    if ($8.isAlt) {
                        $$ = stmt.NewAltForeach($3, $5.node, $6.node, $8.node, $6.byRef)
                    } else {
                        $$ = stmt.NewForeach($3, $5.node, $6.node, $8.node, $6.byRef)
                    }
                }
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $8.node))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_FOREACH '(' expr_without_variable T_AS foreach_variable foreach_optional_arg ')' foreach_statement
            {
                if $6.node == nil {
                    if ($8.isAlt) {
                        $$ = stmt.NewAltForeach($3, nil, $5.node, $8.node, $5.byRef)
                    } else {
                        $$ = stmt.NewForeach($3, nil, $5.node, $8.node, $5.byRef)
                    }
                } else {
                    if ($8.isAlt) {
                        $$ = stmt.NewAltForeach($3, $5.node, $6.node, $8.node, $6.byRef)
                    } else {
                        $$ = stmt.NewForeach($3, $5.node, $6.node, $8.node, $6.byRef)
                    }
                }
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $8.node))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_DECLARE '(' declare_list ')' declare_statement
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
    |   T_TRY '{' inner_statement_list '}' catch_statement finally_statement
            {
                $$ = stmt.NewTry($3, $5, $6)

                if $6 == nil {
                    yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $5))
                } else {
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
;

catch_statement:
        /* empty */
            { $$ = []node.Node{} }
    |   T_CATCH '(' fully_qualified_class_name T_VARIABLE ')' '{' inner_statement_list '}' additional_catches
            {
                identifier := node.NewIdentifier(strings.TrimLeft($4.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                yylex.(*Parser).comments.AddComments(identifier, $4.Comments())
                
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                yylex.(*Parser).comments.AddComments(variable, $4.Comments())
                
                catch := stmt.NewCatch([]node.Node{$3}, variable, $7)
                yylex.(*Parser).positions.AddPosition(catch, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $8))
                yylex.(*Parser).comments.AddComments(catch, $1.Comments())

                $$ = append([]node.Node{catch}, $9...)
            }

finally_statement:
        /* empty */
            { $$ = nil }
    |   T_FINALLY '{' inner_statement_list '}'
            {
                $$ = stmt.NewFinally($3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

additional_catches:
        non_empty_additional_catches
            { $$ = $1 }
    |   /* empty */
            { $$ = []node.Node{} }
;

non_empty_additional_catches:
        additional_catch
            { $$ = []node.Node{$1} }
    |   non_empty_additional_catches additional_catch
            { $$ = append($1, $2) }
;

additional_catch:
        T_CATCH '(' fully_qualified_class_name T_VARIABLE ')' '{' inner_statement_list '}'
            {
                identifier := node.NewIdentifier(strings.TrimLeft($4.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                yylex.(*Parser).comments.AddComments(identifier, $4.Comments())
                
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                yylex.(*Parser).comments.AddComments(variable, $4.Comments())
                
                $$ = stmt.NewCatch([]node.Node{$3}, variable, $7)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $8))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

unset_variables:
        unset_variable
            { $$ = []node.Node{$1} }
    |   unset_variables ',' unset_variable
            { $$ = append($1, $3) }
;

unset_variable:
        variable
            { $$ = $1 }
;

function_declaration_statement:
        unticked_function_declaration_statement
            { $$ = $1 }
;

class_declaration_statement:
        unticked_class_declaration_statement
            { $$ = $1 }
;

is_reference:
        /* empty */
            { $$ = boolWithToken{false, nil} }
    |   '&'
            { $$ = boolWithToken{true, $1} }
;

is_variadic:
        /* empty */
            { $$ = boolWithToken{false, nil} }
    |   T_ELLIPSIS
            { $$ = boolWithToken{true, $1} }
;

unticked_function_declaration_statement:
        function is_reference T_STRING '(' parameter_list ')' '{' inner_statement_list '}'
            {
                name := node.NewIdentifier($3.Value)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                yylex.(*Parser).comments.AddComments(name, $3.Comments())

                $$ = stmt.NewFunction(name, $2.value, $5, nil, $8, "")
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $9))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

unticked_class_declaration_statement:
        class_entry_type T_STRING extends_from implements_list '{' class_statement_list '}'
            {
                switch n := $1.(type) {
                    case *stmt.Class :
                        name := node.NewIdentifier($2.Value)
                        yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                        n.ClassName = name
                        n.Stmts = $6
                        n.Extends = $3
                        n.Implements = $4

                    case *stmt.Trait :
                        // TODO: is it possible that trait extend or implement
                        name := node.NewIdentifier($2.Value)
                        yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                        n.TraitName = name
                        n.Stmts = $6
                }

                $$ = $1
            }
    |   interface_entry T_STRING interface_extends_list '{' class_statement_list '}'
            {
                name := node.NewIdentifier($2.Value)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                yylex.(*Parser).comments.AddComments(name, $2.Comments())
                
                $$ = stmt.NewInterface(name, $3, $5, "")
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $6))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;


class_entry_type:
        T_CLASS
            {
                $$ = stmt.NewClass(nil, nil, nil, nil, nil, nil, "")
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_ABSTRACT T_CLASS
            {
                classModifier := node.NewIdentifier($1.Value)
                yylex.(*Parser).positions.AddPosition(classModifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(classModifier, $1.Comments())

                $$ = stmt.NewClass(nil, []node.Node{classModifier}, nil, nil, nil, nil, "")
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_TRAIT
            {
                $$ = stmt.NewTrait(nil, nil, "")
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_FINAL T_CLASS
            {
                classModifier := node.NewIdentifier($1.Value)
                yylex.(*Parser).positions.AddPosition(classModifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(classModifier, $1.Comments())

                $$ = stmt.NewClass(nil, []node.Node{classModifier}, nil, nil, nil, nil, "")
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

extends_from:
        /* empty */
            { $$ = nil }
    |   T_EXTENDS fully_qualified_class_name
            { $$ = $2 }
;

interface_entry:
        T_INTERFACE
            { $$ = $1 }
;

interface_extends_list:
        /* empty */
            { $$ = nil }
    |   T_EXTENDS interface_list
            { $$ = $2 }
;

implements_list:
        /* empty */
            { $$ = nil }
    |   T_IMPLEMENTS interface_list
            { $$ = $2 }
;

interface_list:
        fully_qualified_class_name
            { $$ = []node.Node{$1} }
    |   interface_list ',' fully_qualified_class_name
            { $$ = append($1, $3) }
;

foreach_optional_arg:
        /* empty */
            { $$ = foreachVariable{nil, false} }
    |   T_DOUBLE_ARROW foreach_variable
            { $$ = $2 }
;

foreach_variable:
        variable
            { $$ = foreachVariable{$1, false} }
    |   '&' variable
            { $$ = foreachVariable{$2, true} }
    |   T_LIST '(' assignment_list ')'
            {
                list := expr.NewList($3)
                yylex.(*Parser).positions.AddPosition(list, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
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
        statement
            { $$ = $1; }
    |   ':' inner_statement_list T_ENDDECLARE ';'
            {
                $$ = stmt.NewStmtList($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;


declare_list:
        T_STRING '=' static_scalar
            {
                name := node.NewIdentifier($1.Value)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(name, $1.Comments())

                constant := stmt.NewConstant(name, $3, "")
                yylex.(*Parser).positions.AddPosition(constant, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))
                yylex.(*Parser).comments.AddComments(constant, $1.Comments())

                $$ = []node.Node{constant}
            }
    |   declare_list ',' T_STRING '=' static_scalar
            {
                name := node.NewIdentifier($3.Value)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                yylex.(*Parser).comments.AddComments(name, $3.Comments())

                constant := stmt.NewConstant(name, $5, "")
                yylex.(*Parser).positions.AddPosition(constant, yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $5))
                yylex.(*Parser).comments.AddComments(constant, $3.Comments())

                $$ = append($1, constant)
            }
;


switch_case_list:
        '{' case_list '}'
            { $$ = &nodesWithEndToken{$2, $3} }
    |   '{' ';' case_list '}'
            { $$ = &nodesWithEndToken{$3, $4} }
    |   ':' case_list T_ENDSWITCH ';'
            { $$ = &nodesWithEndToken{$2, $4} }
    |   ':' ';' case_list T_ENDSWITCH ';'
            { $$ = &nodesWithEndToken{$3, $5} }
;


case_list:
        /* empty */
            { $$ = []node.Node{} }
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



elseif_list:
        /* empty */
            { $$ = nil }
    |   elseif_list T_ELSEIF parenthesis_expr statement
            {
                _elseIf := stmt.NewElseIf($3, $4)
                yylex.(*Parser).positions.AddPosition(_elseIf, yylex.(*Parser).positionBuilder.NewTokenNodePosition($2, $4))
                yylex.(*Parser).comments.AddComments(_elseIf, $2.Comments())

                $$ = append($1, _elseIf)
            }
;


new_elseif_list:
        /* empty */
            { $$ = nil }
    |   new_elseif_list T_ELSEIF parenthesis_expr ':' inner_statement_list
            {
                stmts := stmt.NewStmtList($5)
                yylex.(*Parser).positions.AddPosition(stmts, yylex.(*Parser).positionBuilder.NewNodeListPosition($5))

                _elseIf := stmt.NewAltElseIf($3, stmts)
                yylex.(*Parser).positions.AddPosition(_elseIf, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($2, $5))
                yylex.(*Parser).comments.AddComments(_elseIf, $2.Comments())

                $$ = append($1, _elseIf)
            }
;


else_single:
        /* empty */
            { $$ = nil }
    |   T_ELSE statement
            {
                $$ = stmt.NewElse($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;


new_else_single:
        /* empty */
            { $$ = nil }
    |   T_ELSE ':' inner_statement_list
            {
                stmts := stmt.NewStmtList($3)
                yylex.(*Parser).positions.AddPosition(stmts, yylex.(*Parser).positionBuilder.NewNodeListPosition($3))

                $$ = stmt.NewAltElse(stmts)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;


parameter_list:
        non_empty_parameter_list
            { $$ = $1; }
    |   /* empty */
            { $$ = nil }
;

non_empty_parameter_list:
        parameter
            { $$ = []node.Node{$1} }
    |   non_empty_parameter_list ',' parameter
            { $$ = append($1, $3) }
;

parameter:
        optional_class_type is_reference is_variadic T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeft($4.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                yylex.(*Parser).comments.AddComments($$, $4.Comments())

                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                yylex.(*Parser).comments.AddComments($$, $4.Comments())
                
                $$ = node.NewParameter($1, variable, nil, $2.value, $3.value)
                
                if $1 != nil {
                    yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                    yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
                } else if $2.value == true {
                    yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($2.token, $4))
                    yylex.(*Parser).comments.AddComments($$, $2.token.Comments())
                } else if $3.value == true {
                    yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($3.token, $4))
                    yylex.(*Parser).comments.AddComments($$, $3.token.Comments())
                } else {
                    yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                    yylex.(*Parser).comments.AddComments($$, $4.Comments())
                }
            }
    |   optional_class_type is_reference is_variadic T_VARIABLE '=' static_scalar
            {
                identifier := node.NewIdentifier(strings.TrimLeft($4.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                yylex.(*Parser).comments.AddComments(identifier, $4.Comments())

                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                yylex.(*Parser).comments.AddComments(variable, $4.Comments())

                $$ = node.NewParameter($1, variable, $6, $2.value, $3.value)

                if $1 != nil {
                    yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $6))
                    yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
                } else if $2.value == true {
                    yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($2.token, $6))
                    yylex.(*Parser).comments.AddComments($$, $2.token.Comments())
                } else if $3.value == true {
                    yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($3.token, $6))
                    yylex.(*Parser).comments.AddComments($$, $3.token.Comments())
                } else {
                    yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($4, $6))
                    yylex.(*Parser).comments.AddComments($$, $4.Comments())
                }
            }
;


optional_class_type:
        /* empty */
            { $$ = nil }
    |   T_ARRAY
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
    |   fully_qualified_class_name
            { $$ = $1 }
;


function_call_parameter_list:
        '(' ')'
            { $$ = &nodesWithEndToken{[]node.Node{}, $2} }
    |   '(' non_empty_function_call_parameter_list ')'
            { $$ = &nodesWithEndToken{$2, $3} }
    |   '(' yield_expr ')'
            {
                arg := node.NewArgument($2, false, false)
                yylex.(*Parser).positions.AddPosition(arg, yylex.(*Parser).positionBuilder.NewNodePosition($2))
                yylex.(*Parser).comments.AddComments(arg, yylex.(*Parser).comments[$2])

                $$ = &nodesWithEndToken{[]node.Node{arg}, $3}
            }
;


non_empty_function_call_parameter_list:
        function_call_parameter
            { $$ = []node.Node{$1} }
    |   non_empty_function_call_parameter_list ',' function_call_parameter
            { $$ = append($1, $3) }
;

function_call_parameter:
        expr_without_variable
            {
                $$ = node.NewArgument($1, false, false)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodePosition($1))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   variable
            {
                $$ = node.NewArgument($1, false, false)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodePosition($1))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   '&' w_variable
            {
                $$ = node.NewArgument($2, false, true)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodePosition($2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_ELLIPSIS expr
            {
                $$ = node.NewArgument($2, true, false)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

global_var_list:
        global_var_list ',' global_var
            { $$ = append($1, $3) }
    |   global_var
            { $$ = []node.Node{$1} }
;


global_var:
        T_VARIABLE
            {
                name := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$ = expr.NewVariable(name)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                
                yylex.(*Parser).comments.AddComments(name, $1.Comments())
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   '$' r_variable
            {
                $$ = expr.NewVariable($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   '$' '{' expr '}'
            {
                $$ = expr.NewVariable($3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;


static_var_list:
        static_var_list ',' T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeft($3.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                
                staticVar := stmt.NewStaticVar(variable, nil)
                yylex.(*Parser).positions.AddPosition(staticVar, yylex.(*Parser).positionBuilder.NewTokenPosition($3))

                $$ = append($1, staticVar)

                yylex.(*Parser).comments.AddComments(identifier, $3.Comments())
                yylex.(*Parser).comments.AddComments(variable, $3.Comments())
                yylex.(*Parser).comments.AddComments(staticVar, $3.Comments())
            }
    |   static_var_list ',' T_VARIABLE '=' static_scalar
            {
                identifier := node.NewIdentifier(strings.TrimLeft($3.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                
                staticVar := stmt.NewStaticVar(variable, $5)
                yylex.(*Parser).positions.AddPosition(staticVar, yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $5))

                $$ = append($1, staticVar)

                yylex.(*Parser).comments.AddComments(identifier, $3.Comments())
                yylex.(*Parser).comments.AddComments(variable, $3.Comments())
                yylex.(*Parser).comments.AddComments(staticVar, $3.Comments())
            }
    |   T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                
                staticVar := stmt.NewStaticVar(variable, nil)
                yylex.(*Parser).positions.AddPosition(staticVar, yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                $$ = []node.Node{staticVar}

                yylex.(*Parser).comments.AddComments(identifier, $1.Comments())
                yylex.(*Parser).comments.AddComments(variable, $1.Comments())
                yylex.(*Parser).comments.AddComments(staticVar, $1.Comments())
            }
    |   T_VARIABLE '=' static_scalar
            {
                identifier := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                
                staticVar := stmt.NewStaticVar(variable, $3)
                yylex.(*Parser).positions.AddPosition(staticVar, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))

                $$ = []node.Node{staticVar}

                yylex.(*Parser).comments.AddComments(identifier, $1.Comments())
                yylex.(*Parser).comments.AddComments(variable, $1.Comments())
                yylex.(*Parser).comments.AddComments(staticVar, $1.Comments())
            }

;


class_statement_list:
        class_statement_list class_statement
            { $$ = append($1, $2) }
    |   /* empty */
            { $$ = []node.Node{} }
;


class_statement:
        variable_modifiers class_variable_declaration ';'
            {
                $$ = stmt.NewPropertyList($1, $2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($1))
            }
    |   class_constant_declaration ';'
            { $$ = $1 }
    |   trait_use_statement
            { $$ = $1 }
    |   method_modifiers function is_reference T_STRING '(' parameter_list ')' method_body
            {
                name := node.NewIdentifier($4.Value)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                yylex.(*Parser).comments.AddComments(name, $4.Comments())
                
                $$ = stmt.NewClassMethod(name, $1, $3.value, $6, nil, $8.nodes, "")
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewOptionalListTokensPosition($1, $2, $8.endToken))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($1))
            }
;

trait_use_statement:
        T_USE trait_list trait_adaptations
            {
                $$ = stmt.NewTraitUse($2, $3.nodes)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3.endToken))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

trait_list:
        fully_qualified_class_name
            { $$ = []node.Node{$1} }
    |   trait_list ',' fully_qualified_class_name
            { $$ = append($1, $3) }
;

trait_adaptations:
        ';'
            { $$ = &nodesWithEndToken{nil, $1} }
    |   '{' trait_adaptation_list '}'
            { $$ = &nodesWithEndToken{$2, $3} }
;

trait_adaptation_list:
        /* empty */
            { $$ = nil }
    |   non_empty_trait_adaptation_list
            { $$ = $1 }
;

non_empty_trait_adaptation_list:
        trait_adaptation_statement
            { $$ = []node.Node{$1} }
    |   non_empty_trait_adaptation_list trait_adaptation_statement
            { $$ = append($1, $2) }
;

trait_adaptation_statement:
        trait_precedence ';'
            { $$ = $1 }
    |   trait_alias ';'
            { $$ = $1 }
;

trait_precedence:
        trait_method_reference_fully_qualified T_INSTEADOF trait_reference_list
            {
                $$ = stmt.NewTraitUsePrecedence($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeNodeListPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
;

trait_reference_list:
        fully_qualified_class_name
            { $$ = []node.Node{$1} }
    |   trait_reference_list ',' fully_qualified_class_name
            { $$ = append($1, $3) }
;

trait_method_reference:
        T_STRING
            {
                name := node.NewIdentifier($1.Value)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(name, $1.Comments())
                
                $$ = stmt.NewTraitMethodRef(nil, name)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   trait_method_reference_fully_qualified
            { $$ = $1 }
;

trait_method_reference_fully_qualified:
        fully_qualified_class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                target := node.NewIdentifier($3.Value)
                yylex.(*Parser).positions.AddPosition(target, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                yylex.(*Parser).comments.AddComments(target, $3.Comments())
                
                $$ = stmt.NewTraitMethodRef($1, target)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
;

trait_alias:
        trait_method_reference T_AS trait_modifiers T_STRING
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

trait_modifiers:
        /* empty */
            { $$ = nil }
    |   member_modifier
            { $$ = $1 }
;

method_body:
        ';' /* abstract method */
            { $$ = &nodesWithEndToken{nil, $1} }
    |   '{' inner_statement_list '}'
            { $$ = &nodesWithEndToken{$2, $3} }
;

variable_modifiers:
        non_empty_member_modifiers
            { $$ = $1; }
    |   T_VAR
            {
                modifier := node.NewIdentifier($1.Value)
                yylex.(*Parser).positions.AddPosition(modifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(modifier, $1.Comments())
                
                $$ = []node.Node{modifier}
            }
;

method_modifiers:
        /* empty */
            { $$ = nil }
    |   non_empty_member_modifiers
            { $$ = $1 }
;

non_empty_member_modifiers:
        member_modifier
            { $$ = []node.Node{$1} }
    |   non_empty_member_modifiers member_modifier
            { $$ = append($1, $2) }
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

class_variable_declaration:
        class_variable_declaration ',' T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeft($3.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                yylex.(*Parser).comments.AddComments(identifier, $3.Comments())
                
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                yylex.(*Parser).comments.AddComments(variable, $3.Comments())
                
                property := stmt.NewProperty(variable, nil, "")
                yylex.(*Parser).positions.AddPosition(property, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                yylex.(*Parser).comments.AddComments(property, $3.Comments())

                $$ = append($1, property)
            }
    |   class_variable_declaration ',' T_VARIABLE '=' static_scalar
            {
                identifier := node.NewIdentifier(strings.TrimLeft($3.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                yylex.(*Parser).comments.AddComments(identifier, $3.Comments())
                
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                yylex.(*Parser).comments.AddComments(variable, $3.Comments())
                
                property := stmt.NewProperty(variable, $5, "")
                yylex.(*Parser).positions.AddPosition(property, yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $5))
                yylex.(*Parser).comments.AddComments(property, $3.Comments())

                $$ = append($1, property)
            }
    |   T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(identifier, $1.Comments())
                
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(variable, $1.Comments())
                
                property := stmt.NewProperty(variable, nil, "")
                yylex.(*Parser).positions.AddPosition(property, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(property, $1.Comments())

                $$ = []node.Node{property}
            }
    |   T_VARIABLE '=' static_scalar
            {
                identifier := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(identifier, $1.Comments())
                
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(variable, $1.Comments())
                
                property := stmt.NewProperty(variable, $3, "")
                yylex.(*Parser).positions.AddPosition(property, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))
                yylex.(*Parser).comments.AddComments(property, $1.Comments())

                $$ = []node.Node{property}
            }
;

class_constant_declaration:
        class_constant_declaration ',' T_STRING '=' static_scalar
            {
                name := node.NewIdentifier($3.Value)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                yylex.(*Parser).comments.AddComments(name, $3.Comments())

                constant := stmt.NewConstant(name, $5, "")
                yylex.(*Parser).positions.AddPosition(constant, yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $5))
                yylex.(*Parser).comments.AddComments(constant, $3.Comments())

                $1.(*stmt.ClassConstList).Consts = append($1.(*stmt.ClassConstList).Consts, constant)
                yylex.(*Parser).positions.AddPosition($1, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $5))

                $$ = $1
            }
    |   T_CONST T_STRING '=' static_scalar
            {
                name := node.NewIdentifier($2.Value)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                yylex.(*Parser).comments.AddComments(name, $2.Comments())

                constant := stmt.NewConstant(name, $4, "")
                yylex.(*Parser).positions.AddPosition(constant, yylex.(*Parser).positionBuilder.NewTokenNodePosition($2, $4))
                yylex.(*Parser).comments.AddComments(constant, $2.Comments())

                $$ = stmt.NewClassConstList(nil, []node.Node{constant})
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

echo_expr_list:
        echo_expr_list ',' expr
            { $$ = append($1, $3) }
    |   expr
            { $$ = []node.Node{$1} }
;


for_expr:
        /* empty */
            { $$ = nil }
    |   non_empty_for_expr
            { $$ = $1 }
;

non_empty_for_expr:
        non_empty_for_expr ',' expr
            { $$ = append($1, $3) }
    |   expr
            { $$ = []node.Node{$1} }
;

chaining_method_or_property:
        chaining_method_or_property variable_property
            { $$ = append($1, $2...) }
    |   variable_property
            { $$ = $1 }
;

chaining_dereference:
        chaining_dereference '[' dim_offset ']'
            {
                fetch := expr.NewArrayDimFetch(nil, $3)
                yylex.(*Parser).positions.AddPosition(fetch, yylex.(*Parser).positionBuilder.NewNodePosition($3))

                $$ = append($1, fetch)
            }
    |   '[' dim_offset ']'
            {
                fetch := expr.NewArrayDimFetch(nil, $2)
                yylex.(*Parser).positions.AddPosition(fetch, yylex.(*Parser).positionBuilder.NewNodePosition($2))
                
                $$ = []node.Node{fetch}
            }
;

chaining_instance_call:
        chaining_dereference chaining_method_or_property
            { $$ = append($1, $2...) }
    |   chaining_dereference
            { $$ = $1 }
    |   chaining_method_or_property
            { $$ = $1 }
;

instance_call:
        /* empty */
            { $$ = nil }
    |   chaining_instance_call
            { $$ = $1 }
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
;

expr_without_variable:
        T_LIST '(' assignment_list ')' '=' expr
            {
                list := expr.NewList($3)
                yylex.(*Parser).positions.AddPosition(list, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                $$ = assign.NewAssign(list, $6)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $6))

                yylex.(*Parser).comments.AddComments(list, $1.Comments())
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   variable '=' expr
            {
                $$ = assign.NewAssign($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   variable '=' '&' variable
            {
                $$ = assign.NewReference($1, $4)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   variable '=' '&' T_NEW class_name_reference ctor_arguments
            {
                _new := expr.NewNew($5, nil)
                yylex.(*Parser).positions.AddPosition(_new, yylex.(*Parser).positionBuilder.NewTokenNodePosition($4, $5))

                if $6 != nil {
                    _new = expr.NewNew($5, $6.nodes)
                    yylex.(*Parser).positions.AddPosition(_new, yylex.(*Parser).positionBuilder.NewTokensPosition($4, $6.endToken))
                }
                yylex.(*Parser).comments.AddComments(_new, yylex.(*Parser).comments[$1])

                $$ = assign.NewReference($1, _new)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, _new))
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
    |   rw_variable T_INC
            {
                $$ = expr.NewPostInc($1)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   T_INC rw_variable
            {
                $$ = expr.NewPreInc($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   rw_variable T_DEC
            {
                $$ = expr.NewPostDec($1)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   T_DEC rw_variable
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
    |   expr T_INSTANCEOF class_name_reference
            {
                $$ = expr.NewInstanceOf($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   parenthesis_expr
            { $$ = $1 }
    |   new_expr
            { $$ = $1 }
    |   '(' new_expr ')' instance_call
            {
                $$ = $2

                for _, n := range($4) {
                    switch nn := n.(type) {
                        case *expr.ArrayDimFetch:
                            nn.Variable = $$
                            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            yylex.(*Parser).comments.AddComments(nn, $1.Comments())
                            $$ = nn
                        
                        case *expr.PropertyFetch:
                            nn.Variable = $$
                            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            yylex.(*Parser).comments.AddComments(nn, $1.Comments())
                            $$ = nn
                        
                        case *expr.MethodCall:
                            nn.Variable = $$
                            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            yylex.(*Parser).comments.AddComments(nn, $1.Comments())
                            $$ = nn
                    }
                }
            }
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
    |   internal_functions_in_yacc
            { $$ = $1 }
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
    |   scalar
            { $$ = $1 }
    |   combined_scalar_offset
            { $$ = $1 }
    |   combined_scalar
            { $$ = $1 }
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
    |   function is_reference '(' parameter_list ')' lexical_vars '{' inner_statement_list '}'
            {
                $$ = expr.NewClosure($4, $6, nil, $8, false, $2.value, "")
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $9))
                
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_STATIC function is_reference '(' parameter_list ')' lexical_vars '{' inner_statement_list '}'
            {
                $$ = expr.NewClosure($5, $7, nil, $9, true, $3.value, "")
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $10))
                
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

yield_expr:
        T_YIELD expr_without_variable
            {
                $$ = expr.NewYield(nil, $2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_YIELD variable
            {
                $$ = expr.NewYield(nil, $2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_YIELD expr T_DOUBLE_ARROW expr_without_variable
            {
                $$ = expr.NewYield($2, $4)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_YIELD expr T_DOUBLE_ARROW variable
            {
                $$ = expr.NewYield($2, $4)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

combined_scalar_offset:
        combined_scalar '[' dim_offset ']'
            {
                $$ = expr.NewArrayDimFetch($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   combined_scalar_offset '[' dim_offset ']'
            {
                $$ = expr.NewArrayDimFetch($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   T_CONSTANT_ENCAPSED_STRING '[' dim_offset ']'
            {
                str := scalar.NewString($1.Value)
                yylex.(*Parser).positions.AddPosition(str, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(str, $1.Comments())

                $$ = expr.NewArrayDimFetch(str, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition(str, $4))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[str])
            }
    |   general_constant '[' dim_offset ']'
            {
                $$ = expr.NewArrayDimFetch($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
;

combined_scalar:
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
;

function:
    T_FUNCTION
        { $$ = $1 }
;

lexical_vars:
        /* empty */
            { $$ = []node.Node{} }
    |   T_USE '(' lexical_var_list ')'
            { $$ = $3; }
;

lexical_var_list:
        lexical_var_list ',' T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeft($3.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                yylex.(*Parser).comments.AddComments(identifier, $3.Comments())
                
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                yylex.(*Parser).comments.AddComments(variable, $3.Comments())
                
                use := expr.NewClosureUse(variable, false)
                yylex.(*Parser).positions.AddPosition(use, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                yylex.(*Parser).comments.AddComments(use, $3.Comments())
                
                $$ = append($1, use)
            }
    |   lexical_var_list ',' '&' T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeft($4.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                yylex.(*Parser).comments.AddComments(identifier, $4.Comments())
                
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                yylex.(*Parser).comments.AddComments(variable, $3.Comments())

                use := expr.NewClosureUse(variable, true)
                yylex.(*Parser).positions.AddPosition(use, yylex.(*Parser).positionBuilder.NewTokensPosition($3, $4))
                yylex.(*Parser).comments.AddComments(use, $3.Comments())

                $$ = append($1, use)
            }
    |   T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeft($1.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(identifier, $1.Comments())
                
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(variable, $1.Comments())
                
                use := expr.NewClosureUse(variable, false)
                yylex.(*Parser).positions.AddPosition(use, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(use, $1.Comments())
                
                $$ = []node.Node{use}
            }
    |   '&' T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeft($2.Value, "$"))
                yylex.(*Parser).positions.AddPosition(identifier, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                yylex.(*Parser).comments.AddComments(identifier, $2.Comments())
                
                variable := expr.NewVariable(identifier)
                yylex.(*Parser).positions.AddPosition(variable, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                yylex.(*Parser).comments.AddComments(variable, $1.Comments())

                use := expr.NewClosureUse(variable, true)
                yylex.(*Parser).positions.AddPosition(use, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))
                yylex.(*Parser).comments.AddComments(use, $1.Comments())

                $$ = []node.Node{use}
            }
;

function_call:
        namespace_name function_call_parameter_list
            {
                name := name.NewName($1)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($1))

                $$ = expr.NewFunctionCall(name, $2.nodes)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition(name, $2.endToken))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[name])
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name function_call_parameter_list
            {
                funcName := name.NewRelative($3)
                yylex.(*Parser).positions.AddPosition(funcName, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))
                yylex.(*Parser).comments.AddComments(funcName, $1.Comments())

                $$ = expr.NewFunctionCall(funcName, $4.nodes)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition(funcName, $4.endToken))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[funcName])
            }
    |   T_NS_SEPARATOR namespace_name function_call_parameter_list
            {
                funcName := name.NewFullyQualified($2)
                yylex.(*Parser).positions.AddPosition(funcName, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))
                yylex.(*Parser).comments.AddComments(funcName, $1.Comments())

                $$ = expr.NewFunctionCall(funcName, $3.nodes)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition(funcName, $3.endToken))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[funcName])
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM variable_name function_call_parameter_list
            {
                $$ = expr.NewStaticCall($1, $3, $4.nodes)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4.endToken))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects function_call_parameter_list
            {
                $$ = expr.NewStaticCall($1, $3, $4.nodes)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4.endToken))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_name function_call_parameter_list
            {
                $$ = expr.NewStaticCall($1, $3, $4.nodes)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4.endToken))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects function_call_parameter_list
            {
                $$ = expr.NewStaticCall($1, $3, $4.nodes)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4.endToken))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   variable_without_objects function_call_parameter_list
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
    |   namespace_name 
            {
                $$ = name.NewName($1)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($1))
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = name.NewRelative($3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = name.NewFullyQualified($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

fully_qualified_class_name:
        namespace_name
            {
                $$ = name.NewName($1)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).listGetFirstNodeComments($1))
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = name.NewRelative($3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = name.NewFullyQualified($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

class_name_reference:
        class_name
            { $$ = $1 }
    |   dynamic_class_name_reference
            { $$ = $1 }
;

dynamic_class_name_reference:
        base_variable T_OBJECT_OPERATOR object_property dynamic_class_name_variable_properties
            {
                $$ = $1

                for _, n := range($3) {
                    switch nn := n.(type) {
                        case *expr.ArrayDimFetch:
                            nn.Variable = $$
                            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            yylex.(*Parser).comments.AddComments(nn, yylex.(*Parser).comments[$1])
                            $$ = nn
                        
                        case *expr.PropertyFetch:
                            nn.Variable = $$
                            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            yylex.(*Parser).comments.AddComments(nn, yylex.(*Parser).comments[$1])
                            $$ = nn
                    }
                }

                for _, n := range($4) {
                    switch nn := n.(type) {
                        case *expr.ArrayDimFetch:
                            nn.Variable = $$
                            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            yylex.(*Parser).comments.AddComments(nn, yylex.(*Parser).comments[$1])
                            $$ = nn
                        
                        case *expr.PropertyFetch:
                            nn.Variable = $$
                            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            yylex.(*Parser).comments.AddComments(nn, yylex.(*Parser).comments[$1])
                            $$ = nn
                    }
                }
            }
    |   base_variable 
            { $$ = $1 }
;


dynamic_class_name_variable_properties:
        dynamic_class_name_variable_properties dynamic_class_name_variable_property
            { $$ = append($1, $2...) }
    |   /* empty */
            { $$ = []node.Node{} }
;


dynamic_class_name_variable_property:
        T_OBJECT_OPERATOR object_property
            { $$ = $2 }
;

exit_expr:
        /* empty */
            { $$ = nil }
    |   '(' ')'
            { $$ = nil }
    |   parenthesis_expr
            { $$ = $1 }
;

backticks_expr:
        /* empty */
            { $$ = []node.Node{} }
    |   T_ENCAPSED_AND_WHITESPACE
            { $$ = []node.Node{scalar.NewEncapsedStringPart($1.Value)} }
    |   encaps_list
            { $$ = $1; }
;

ctor_arguments:
        /* empty */
            { $$ = nil }
    |   function_call_parameter_list
            { $$ = $1 }
;

common_scalar:
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
    |   T_CONSTANT_ENCAPSED_STRING
            {
                $$ = scalar.NewString($1.Value)
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
;

static_class_constant:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                target := node.NewIdentifier($3.Value)
                yylex.(*Parser).positions.AddPosition(target, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$ = expr.NewClassConstFetch($1, target)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))

                yylex.(*Parser).comments.AddComments(target, $3.Comments())
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
;

static_scalar:
        static_scalar_value
            { $$ = $1 }
;

static_scalar_value:
        common_scalar
            { $$ = $1 }
    |   static_class_name_scalar
            { $$ = $1 }
    |   namespace_name
            {
                name := name.NewName($1)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($1))

                $$ = expr.NewConstFetch(name)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodePosition(name))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[name])
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                name := name.NewRelative($3)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))
                yylex.(*Parser).comments.AddComments(name, $1.Comments())

                $$ = expr.NewConstFetch(name)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[name])
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := name.NewFullyQualified($2)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))
                yylex.(*Parser).comments.AddComments(name, $1.Comments())

                $$ = expr.NewConstFetch(name)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[name])
            }
    |   T_ARRAY '(' static_array_pair_list ')'
            {
                $$ = expr.NewArray($3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   '[' static_array_pair_list ']'
            {
                $$ = expr.NewShortArray($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   static_class_constant
            { $$ = $1 }
    |   T_CLASS_C
            {
                $$ = scalar.NewMagicConstant($1.Value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   static_operation
            { $$ = $1 }
;

static_operation:
        static_scalar_value '[' static_scalar_value ']'
            {
                $$ = expr.NewArrayDimFetch($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value '+' static_scalar_value
            {
                $$ = binary.NewPlus($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value '-' static_scalar_value
            {
                $$ = binary.NewMinus($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value '*' static_scalar_value
            {
                $$ = binary.NewMul($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value T_POW static_scalar_value
            {
                $$ = binary.NewPow($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value '/' static_scalar_value
            {
                $$ = binary.NewDiv($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value '%' static_scalar_value
            {
                $$ = binary.NewMod($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   '!' static_scalar_value
            {
                $$ = expr.NewBooleanNot($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   '~' static_scalar_value
            {
                $$ = expr.NewBitwiseNot($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   static_scalar_value '|' static_scalar_value
            {
                $$ = binary.NewBitwiseOr($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value '&' static_scalar_value
            {
                $$ = binary.NewBitwiseAnd($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value '^' static_scalar_value
            {
                $$ = binary.NewBitwiseXor($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value T_SL static_scalar_value
            {
                $$ = binary.NewShiftLeft($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value T_SR static_scalar_value
            {
                $$ = binary.NewShiftRight($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value '.' static_scalar_value
            {
                $$ = binary.NewConcat($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value T_LOGICAL_XOR static_scalar_value
            {
                $$ = binary.NewLogicalXor($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value T_LOGICAL_AND static_scalar_value
            {
                $$ = binary.NewLogicalAnd($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value T_LOGICAL_OR static_scalar_value
            {
                $$ = binary.NewLogicalOr($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value T_BOOLEAN_AND static_scalar_value
            {
                $$ = binary.NewBooleanAnd($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value T_BOOLEAN_OR static_scalar_value
            {
                $$ = binary.NewBooleanOr($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value T_IS_IDENTICAL static_scalar_value
            {
                $$ = binary.NewIdentical($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value T_IS_NOT_IDENTICAL static_scalar_value
            {
                $$ = binary.NewNotIdentical($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value T_IS_EQUAL static_scalar_value
            {
                $$ = binary.NewEqual($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value T_IS_NOT_EQUAL static_scalar_value
            {
                $$ = binary.NewNotEqual($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value '<' static_scalar_value
            {
                $$ = binary.NewSmaller($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value '>' static_scalar_value
            {
                $$ = binary.NewGreater($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value T_IS_SMALLER_OR_EQUAL static_scalar_value
            {
                $$ = binary.NewSmallerOrEqual($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value T_IS_GREATER_OR_EQUAL static_scalar_value
            {
                $$ = binary.NewGreaterOrEqual($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value '?' ':' static_scalar_value
            {
                $$ = expr.NewTernary($1, nil, $4)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   static_scalar_value '?' static_scalar_value ':' static_scalar_value
            {
                $$ = expr.NewTernary($1, $3, $5)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $5))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   '+' static_scalar_value
            {
                $$ = expr.NewUnaryPlus($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   '-' static_scalar_value
            {
                $$ = expr.NewUnaryMinus($2)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   '(' static_scalar_value ')'
            { $$ = $2 }
;

general_constant:
        class_constant
            { $$ = $1 }
    |   namespace_name
            {
                name := name.NewName($1)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                yylex.(*Parser).comments.AddComments(name, yylex.(*Parser).listGetFirstNodeComments($1))

                $$ = expr.NewConstFetch(name)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodePosition(name))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[name])
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                name := name.NewRelative($3)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))
                yylex.(*Parser).comments.AddComments(name, $1.Comments())

                $$ = expr.NewConstFetch(name)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodePosition(name))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[name])
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := name.NewFullyQualified($2)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))
                yylex.(*Parser).comments.AddComments(name, $1.Comments())

                $$ = expr.NewConstFetch(name)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodePosition(name))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[name])
            }
;

scalar:
        T_STRING_VARNAME
            {
                name := node.NewIdentifier($1.Value)
                yylex.(*Parser).positions.AddPosition(name, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$ = expr.NewVariable(name)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                yylex.(*Parser).comments.AddComments(name, $1.Comments())
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   general_constant
            { $$ = $1 }
    |   class_name_scalar
            { $$ = $1 }
    |   common_scalar
            { $$ = $1 }
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
    |   T_CLASS_C
            {
                $$ = scalar.NewMagicConstant($1.Value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
;

static_array_pair_list:
        /* empty */
            { $$ = nil }
    |   non_empty_static_array_pair_list possible_comma
            { $$ = $1 }
;

possible_comma:
        /* empty */
    |   ','
;

non_empty_static_array_pair_list:
        non_empty_static_array_pair_list ',' static_scalar_value T_DOUBLE_ARROW static_scalar_value
            {
                arrayItem := expr.NewArrayItem($3, $5, false)
                yylex.(*Parser).positions.AddPosition(arrayItem, yylex.(*Parser).positionBuilder.NewNodesPosition($3, $5))
                yylex.(*Parser).comments.AddComments(arrayItem, yylex.(*Parser).comments[$3])

                $$ = append($1, arrayItem)
            }
    |   non_empty_static_array_pair_list ',' static_scalar_value
            {
                arrayItem := expr.NewArrayItem(nil, $3, false)
                yylex.(*Parser).positions.AddPosition(arrayItem, yylex.(*Parser).positionBuilder.NewNodePosition($3))
                yylex.(*Parser).comments.AddComments(arrayItem, yylex.(*Parser).comments[$3])

                $$ = append($1, arrayItem)
            }
    |   static_scalar_value T_DOUBLE_ARROW static_scalar_value
            {
                arrayItem := expr.NewArrayItem($1, $3, false)
                yylex.(*Parser).positions.AddPosition(arrayItem, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments(arrayItem, yylex.(*Parser).comments[$1])

                $$ = []node.Node{arrayItem}
            }
    |   static_scalar_value
            {
                arrayItem := expr.NewArrayItem(nil, $1, false)
                yylex.(*Parser).positions.AddPosition(arrayItem, yylex.(*Parser).positionBuilder.NewNodePosition($1))
                yylex.(*Parser).comments.AddComments(arrayItem, yylex.(*Parser).comments[$1])

                $$ = []node.Node{arrayItem}
            }
;

expr:
        r_variable
            { $$ = $1 }
    |   expr_without_variable
            { $$ = $1 }
;

parenthesis_expr:
        '(' expr ')'
            { $$ = $2 }
    |   '(' yield_expr ')'
            { $$ = $2 }
;


r_variable:
        variable
            { $$ = $1 }
;


w_variable:
        variable
            { $$ = $1 }
;

rw_variable:
        variable
            { $$ = $1 }
;

variable:
        base_variable_with_function_calls T_OBJECT_OPERATOR object_property method_or_not variable_properties
            {
                $$ = $1

                if $4 != nil {
                    $4[0].(*expr.MethodCall).Method = $3[len($3)-1].(*expr.PropertyFetch).Property
                    $3 = append($3[:len($3)-1], $4...)
                }

                for _, n := range($3) {
                    switch nn := n.(type) {
                        case *expr.ArrayDimFetch:
                            nn.Variable = $$
                            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            yylex.(*Parser).comments.AddComments(nn, yylex.(*Parser).comments[$1])
                            $$ = nn
                        
                        case *expr.PropertyFetch:
                            nn.Variable = $$
                            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            yylex.(*Parser).comments.AddComments(nn, yylex.(*Parser).comments[$1])
                            $$ = nn
                        
                        case *expr.MethodCall:
                            nn.Variable = $$
                            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            yylex.(*Parser).comments.AddComments(nn, yylex.(*Parser).comments[$1])
                            $$ = nn
                    }
                }

                for _, n := range($5) {
                    switch nn := n.(type) {
                        case *expr.ArrayDimFetch:
                            nn.Variable = $$
                            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            yylex.(*Parser).comments.AddComments(nn, yylex.(*Parser).comments[$1])
                            $$ = nn
                        
                        case *expr.PropertyFetch:
                            nn.Variable = $$
                            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            yylex.(*Parser).comments.AddComments(nn, yylex.(*Parser).comments[$1])
                            $$ = nn
                        
                        case *expr.MethodCall:
                            nn.Variable = $$
                            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            yylex.(*Parser).comments.AddComments(nn, yylex.(*Parser).comments[$1])
                            $$ = nn
                    }
                }
            }
    |   base_variable_with_function_calls
            { $$ = $1 }
;

variable_properties:
        variable_properties variable_property
            { $$ = append($1, $2...) }
    |   /* empty */
            { $$ = []node.Node{} }
;


variable_property:
        T_OBJECT_OPERATOR object_property method_or_not
            {
                if $3 != nil {
                    $3[0].(*expr.MethodCall).Method = $2[len($2)-1].(*expr.PropertyFetch).Property
                    $2 = append($2[:len($2)-1], $3...)
                }

                $$ = $2
            }
;

array_method_dereference:
        array_method_dereference '[' dim_offset ']'
            {
                fetch := expr.NewArrayDimFetch(nil, $3)
                yylex.(*Parser).positions.AddPosition(fetch, yylex.(*Parser).positionBuilder.NewNodePosition($3))

                $$ = append($1, fetch)
            }
    |   method '[' dim_offset ']'
            {
                fetch := expr.NewArrayDimFetch(nil, $3)
                yylex.(*Parser).positions.AddPosition(fetch, yylex.(*Parser).positionBuilder.NewNodePosition($3))

                $$ = []node.Node{$1, fetch}
            }
;

method:
        function_call_parameter_list
            {
                $$ = expr.NewMethodCall(nil, nil, $1.nodes)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($1.nodes, $1.endToken))
            }
;

method_or_not:
        method
            { $$ = []node.Node{$1} }
    |   array_method_dereference
            { $$ = $1 }
    |   /* empty */
            { $$ = nil }
;

variable_without_objects:
        reference_variable
            { $$ = $1 }
    |   simple_indirect_reference reference_variable
            {
                $1.last.SetVarName($2)

                for _, n := range($1.all) {
                    yylex.(*Parser).positions[n] =  yylex.(*Parser).positionBuilder.NewNodesPosition(n, $2)
                }

                $$ = $1.all[0]
            }
;

static_member:
        class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects
            {
                $$ = expr.NewStaticPropertyFetch($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects
            {
                $$ = expr.NewStaticPropertyFetch($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }

;

variable_class_name:
        reference_variable
            { $$ = $1 }
;

array_function_dereference:
        array_function_dereference '[' dim_offset ']'
            {
                $$ = expr.NewArrayDimFetch($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   function_call '[' dim_offset ']'
            {
                $$ = expr.NewArrayDimFetch($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
;

base_variable_with_function_calls:
        base_variable               { $$ = $1 }
    |   array_function_dereference  { $$ = $1 }
    |   function_call               { $$ = $1 }
;


base_variable:
        reference_variable
            { $$ = $1 }
    |   simple_indirect_reference reference_variable
            {
                $1.last.SetVarName($2)

                for _, n := range($1.all) {
                    yylex.(*Parser).positions[n] =  yylex.(*Parser).positionBuilder.NewNodesPosition(n, $2)
                }

                $$ = $1.all[0]
            }
    |   static_member
            { $$ = $1 }
;

reference_variable:
        reference_variable '[' dim_offset ']'
            {
                $$ = expr.NewArrayDimFetch($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   reference_variable '{' expr '}'
            {
                $$ = expr.NewArrayDimFetch($1, $3)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   compound_variable
            { $$ = $1 }
;


compound_variable:
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
;

dim_offset:
        /* empty */
            { $$ = nil }
    |   expr
            { $$ = $1 }
;


object_property:
        object_dim_list 
            { $$ = $1 }
    |   variable_without_objects
            {
                fetch := expr.NewPropertyFetch(nil, $1)
                yylex.(*Parser).positions.AddPosition(fetch, yylex.(*Parser).positionBuilder.NewNodePosition($1))

                $$ = []node.Node{fetch}
            }
;

object_dim_list:
        object_dim_list '[' dim_offset ']'
            {
                fetch := expr.NewArrayDimFetch(nil, $3)
                yylex.(*Parser).positions.AddPosition(fetch, yylex.(*Parser).positionBuilder.NewNodePosition($3))

                $$ = append($1, fetch)
            }
    |   object_dim_list '{' expr '}'
            {
                fetch := expr.NewArrayDimFetch(nil, $3)
                yylex.(*Parser).positions.AddPosition(fetch, yylex.(*Parser).positionBuilder.NewNodePosition($3))

                $$ = append($1, fetch)
            }
    |   variable_name
            {
                fetch := expr.NewPropertyFetch(nil, $1)
                yylex.(*Parser).positions.AddPosition(fetch, yylex.(*Parser).positionBuilder.NewNodePosition($1))

                $$ = []node.Node{fetch}
            }
;

variable_name:
        T_STRING
            {
                $$ = node.NewIdentifier($1.Value)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments($$, $1.Comments())
            }
    |   '{' expr '}'
            { $$ = $2 }
;

simple_indirect_reference:
        '$'
            {
                n := expr.NewVariable(nil)
                yylex.(*Parser).positions.AddPosition(n, yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                yylex.(*Parser).comments.AddComments(n, $1.Comments())

                $$ = simpleIndirectReference{[]*expr.Variable{n}, n}
            }
    |   simple_indirect_reference '$'
            {
                n := expr.NewVariable(nil)
                yylex.(*Parser).positions.AddPosition(n, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                yylex.(*Parser).comments.AddComments(n, $2.Comments())

                $1.last.SetVarName(n)

                $1.all = append($1.all, n)
                $1.last = n
                $$ = $1
            }
;

assignment_list:
        assignment_list ',' assignment_list_element
            { $$ = append($1, $3) }
    |   assignment_list_element
            {
                if $1 == nil {
                    $$ = []node.Node{}
                } else {
                    $$ = []node.Node{$1}
                }
            }
;


assignment_list_element:
        variable
            {
                $$ = expr.NewArrayItem(nil, $1, false)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodePosition($1))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   T_LIST '(' assignment_list ')'
            {
                item := expr.NewList($3)
                yylex.(*Parser).positions.AddPosition(item, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                yylex.(*Parser).comments.AddComments(item, $1.Comments())

                $$ = expr.NewArrayItem(nil, item, false)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodePosition(item))
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[item])
            }
    |   /* empty */
            { $$ = nil }
;


array_pair_list:
        /* empty */
            { $$ = []node.Node{} }
    |   non_empty_array_pair_list possible_comma
            { $$ = $1 }
;

non_empty_array_pair_list:
        non_empty_array_pair_list ',' expr T_DOUBLE_ARROW expr
            {
                arrayItem := expr.NewArrayItem($3, $5, false)
                yylex.(*Parser).positions.AddPosition(arrayItem, yylex.(*Parser).positionBuilder.NewNodesPosition($3, $5))
                yylex.(*Parser).comments.AddComments(arrayItem, yylex.(*Parser).comments[$3])

                $$ = append($1, arrayItem)
            }
    |   non_empty_array_pair_list ',' expr
            {
                arrayItem := expr.NewArrayItem(nil, $3, false)
                yylex.(*Parser).positions.AddPosition(arrayItem, yylex.(*Parser).positionBuilder.NewNodePosition($3))
                yylex.(*Parser).comments.AddComments(arrayItem, yylex.(*Parser).comments[$3])

                $$ = append($1, arrayItem)
            }
    |   expr T_DOUBLE_ARROW expr
            {
                arrayItem := expr.NewArrayItem($1, $3, false)
                yylex.(*Parser).positions.AddPosition(arrayItem, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                yylex.(*Parser).comments.AddComments(arrayItem, yylex.(*Parser).comments[$1])

                $$ = []node.Node{arrayItem}
            }
    |   expr
            {
                arrayItem := expr.NewArrayItem(nil, $1, false)
                yylex.(*Parser).positions.AddPosition(arrayItem, yylex.(*Parser).positionBuilder.NewNodePosition($1))
                yylex.(*Parser).comments.AddComments(arrayItem, yylex.(*Parser).comments[$1])

                $$ = []node.Node{arrayItem}
            }
    |   non_empty_array_pair_list ',' expr T_DOUBLE_ARROW '&' w_variable
            {
                arrayItem := expr.NewArrayItem($3, $6, true)
                yylex.(*Parser).positions.AddPosition(arrayItem, yylex.(*Parser).positionBuilder.NewNodesPosition($3, $6))
                yylex.(*Parser).comments.AddComments(arrayItem, yylex.(*Parser).comments[$3])

                $$ = append($1, arrayItem)
            }
    |   non_empty_array_pair_list ',' '&' w_variable
            {
                arrayItem := expr.NewArrayItem(nil, $4, true)
                yylex.(*Parser).positions.AddPosition(arrayItem, yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $4))
                yylex.(*Parser).comments.AddComments(arrayItem, $3.Comments())

                $$ = append($1, arrayItem)
            }
    |   expr T_DOUBLE_ARROW '&' w_variable
            {
                arrayItem := expr.NewArrayItem($1, $4, true)
                yylex.(*Parser).positions.AddPosition(arrayItem, yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4))
                yylex.(*Parser).comments.AddComments(arrayItem, yylex.(*Parser).comments[$1])

                $$ = []node.Node{arrayItem}
            }
    |   '&' w_variable
            {
                arrayItem := expr.NewArrayItem(nil, $2, true)
                yylex.(*Parser).positions.AddPosition(arrayItem, yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                yylex.(*Parser).comments.AddComments(arrayItem, $1.Comments())

                $$ = []node.Node{arrayItem}
            }
;

encaps_list:
        encaps_list encaps_var
            { $$ = append($1, $2) }
    |   encaps_list T_ENCAPSED_AND_WHITESPACE
            {
                encapsed := scalar.NewEncapsedStringPart($2.Value)
                yylex.(*Parser).positions.AddPosition(encapsed, yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                $$ = append($1, encapsed)
                yylex.(*Parser).comments.AddComments(encapsed, $2.Comments())
            }
    |   encaps_var
            { $$ = []node.Node{$1} }
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
                $$ = $2
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
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
    |   T_CURLY_OPEN variable '}'
            { $$ = $2; }
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
    T_ISSET '(' isset_variables ')'
        {
            $$ = expr.NewIsset($3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_EMPTY '(' variable ')'
        {
            $$ = expr.NewEmpty($3)
            yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
            yylex.(*Parser).comments.AddComments($$, $1.Comments())
        }
    |   T_EMPTY '(' expr_without_variable ')'
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
        isset_variable
            { $$ = []node.Node{$1} }
    |   isset_variables ',' isset_variable
            { $$ = append($1, $3) }
;

isset_variable:
        variable
            { $$ = $1 }
    |   expr_without_variable
            { $$ = $1 }
;

class_constant:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                target := node.NewIdentifier($3.Value)
                yylex.(*Parser).positions.AddPosition(target, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$ = expr.NewClassConstFetch($1, target)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))

                yylex.(*Parser).comments.AddComments(target, $3.Comments())
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                target := node.NewIdentifier($3.Value)
                yylex.(*Parser).positions.AddPosition(target, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$ = expr.NewClassConstFetch($1, target)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))

                yylex.(*Parser).comments.AddComments(target, $3.Comments())
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
;

static_class_name_scalar:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_CLASS
            {
                target := node.NewIdentifier($3.Value)
                yylex.(*Parser).positions.AddPosition(target, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$ = expr.NewClassConstFetch($1, target)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))

                yylex.(*Parser).comments.AddComments(target, $3.Comments())
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
;

class_name_scalar:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_CLASS
            {
                target := node.NewIdentifier($3.Value)
                yylex.(*Parser).positions.AddPosition(target, yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$ = expr.NewClassConstFetch($1, target)
                yylex.(*Parser).positions.AddPosition($$, yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))

                yylex.(*Parser).comments.AddComments(target, $3.Comments())
                yylex.(*Parser).comments.AddComments($$, yylex.(*Parser).comments[$1])
            }
;

%%

type foreachVariable struct {
	node  node.Node
	byRef bool
}

type nodesWithEndToken struct {
	nodes    []node.Node
	endToken *scanner.Token
}

type boolWithToken struct {
	value bool
	token *scanner.Token
}

type simpleIndirectReference struct {
	all  []*expr.Variable
	last *expr.Variable
}

type altSyntaxNode struct {
	node  node.Node
	isAlt bool
}
