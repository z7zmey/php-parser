%{
package php5

import (
    "strings"
    "strconv"

    "github.com/z7zmey/php-parser/freefloating"
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
    node                    node.Node
    token                   *scanner.Token
    list                    []node.Node
    simpleIndirectReference simpleIndirectReference

    ClassExtends            *stmt.ClassExtends
    ClassImplements         *stmt.ClassImplements
    InterfaceExtends        *stmt.InterfaceExtends
    ClosureUse              *expr.ClosureUse
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
%type <token> possible_comma
%type <token> case_separator

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
%type <node> optional_class_type parameter class_entry_type class_statement class_constant_declaration
%type <node> trait_use_statement function_call_parameter trait_adaptation_statement trait_precedence trait_alias
%type <node> trait_method_reference_fully_qualified trait_method_reference trait_modifiers member_modifier method
%type <node> static_scalar_value static_operation
%type <node> ctor_arguments function_call_parameter_list
%type <node> trait_adaptations
%type <node> switch_case_list
%type <node> method_body
%type <node> foreach_statement for_statement while_statement
%type <node> foreach_variable foreach_optional_arg

%type <ClassExtends> extends_from
%type <ClassImplements> implements_list
%type <InterfaceExtends> interface_extends_list
%type <ClosureUse> lexical_vars

%type <list> top_statement_list namespace_name use_declarations use_function_declarations use_const_declarations
%type <list> inner_statement_list global_var_list static_var_list encaps_list isset_variables non_empty_array_pair_list
%type <list> array_pair_list assignment_list lexical_var_list elseif_list new_elseif_list non_empty_for_expr
%type <list> for_expr case_list echo_expr_list unset_variables declare_list catch_statement additional_catches
%type <list> non_empty_additional_catches parameter_list non_empty_parameter_list class_statement_list
%type <list> class_statement_list variable_modifiers method_modifiers class_variable_declaration
%type <list> interface_list non_empty_function_call_parameter_list trait_list trait_adaptation_list non_empty_trait_adaptation_list
%type <list> trait_reference_list non_empty_member_modifiers backticks_expr static_array_pair_list non_empty_static_array_pair_list

%type <list> chaining_dereference chaining_instance_call chaining_method_or_property instance_call variable_property
%type <list> method_or_not array_method_dereference object_property object_dim_list dynamic_class_name_variable_property
%type <list> dynamic_class_name_variable_properties variable_properties

%type <simpleIndirectReference> simple_indirect_reference
%type <token> is_reference is_variadic

%%

start:
        top_statement_list
            {
                yylex.(*Parser).rootNode = node.NewRoot($1)
                yylex.(*Parser).rootNode.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))

                yylex.(*Parser).setFreeFloating(yylex.(*Parser).rootNode, freefloating.End, yylex.(*Parser).currentToken.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

top_statement_list:
        top_statement_list top_statement
            {
                if inlineHtmlNode, ok := $2.(*stmt.InlineHtml); ok && len($1) > 0 {
                    prevNode := lastNode($1)
                    yylex.(*Parser).splitSemiColonAndPhpCloseTag(inlineHtmlNode, prevNode)
                }

                if $2 != nil {
                    $$ = append($1, $2)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                $$ = []node.Node{}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

namespace_name:
        T_STRING
            {
                namePart := name.NewNamePart($1.Value)
                $$ = []node.Node{namePart}

                // save position
                namePart.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating(namePart, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_NS_SEPARATOR T_STRING
            {
                namePart := name.NewNamePart($3.Value)
                $$ = append($1, namePart)

                // save position
                namePart.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(namePart, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

top_statement:
        error
            {
                // error
                $$ = nil

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   statement
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   function_declaration_statement
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_declaration_statement
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_HALT_COMPILER '(' ')' ';'
            {
                $$ = stmt.NewHaltCompiler()

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.HaltCompiller, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.OpenParenthesisToken, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.CloseParenthesisToken, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE namespace_name ';'
            {
                name := name.NewName($2)
                $$ = stmt.NewNamespace(name, nil)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).MoveFreeFloating($2[0], name)
                yylex.(*Parser).setFreeFloating(name, freefloating.End, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE namespace_name '{' top_statement_list '}'
            {
                name := name.NewName($2)
                $$ = stmt.NewNamespace(name, $4)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $5))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).MoveFreeFloating($2[0], name)
                yylex.(*Parser).setFreeFloating(name, freefloating.End, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $5.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE '{' top_statement_list '}'
            {
                $$ = stmt.NewNamespace(nil, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Namespace, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE use_declarations ';'
            {
                $$ = stmt.NewUseList(nil, $2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.UseDeclarationList, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE T_FUNCTION use_function_declarations ';'
            {
                useType := node.NewIdentifier($2.Value)
                $$ = stmt.NewUseList(useType, $3)

                // save position
                useType.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(useType, freefloating.Start, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.UseDeclarationList, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE T_CONST use_const_declarations ';'
            {
                useType := node.NewIdentifier($2.Value)
                $$ = stmt.NewUseList(useType, $3)

                // save position
                useType.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(useType, freefloating.Start, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.UseDeclarationList, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   constant_declaration ';'
            {
                $$ = $1

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_declarations:
        use_declarations ',' use_declaration
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   use_declaration
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_declaration:
        namespace_name
            {
                name := name.NewName($1)
                $$ = stmt.NewUse(nil, name, nil)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_AS T_STRING
            {
                name := name.NewName($1)
                alias := node.NewIdentifier($3.Value)
                $$ = stmt.NewUse(nil, name, alias)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                alias.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)
                yylex.(*Parser).setFreeFloating(name, freefloating.End, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(alias, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := name.NewName($2)
                $$ = stmt.NewUse(nil, name, nil)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Slash, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).MoveFreeFloating($2[0], name)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                name := name.NewName($2)
                alias := node.NewIdentifier($4.Value)
                $$ = stmt.NewUse(nil, name, alias)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                alias.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($2, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Slash, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).MoveFreeFloating($2[0], name)
                yylex.(*Parser).setFreeFloating(name, freefloating.End, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating(alias, freefloating.Start, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_function_declarations:
        use_function_declarations ',' use_function_declaration
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   use_function_declaration
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_function_declaration:
        namespace_name
            {
                name := name.NewName($1)
                $$ = stmt.NewUse(nil, name, nil)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_AS T_STRING
            {
                name := name.NewName($1)
                alias := node.NewIdentifier($3.Value)
                $$ = stmt.NewUse(nil, name, alias)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                alias.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)
                yylex.(*Parser).setFreeFloating(name, freefloating.End, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(alias, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := name.NewName($2)
                $$ = stmt.NewUse(nil, name, nil)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Slash, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).MoveFreeFloating($2[0], name)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                name := name.NewName($2)
                alias := node.NewIdentifier($4.Value)
                $$ = stmt.NewUse(nil, name, alias)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                alias.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($2, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Slash, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).MoveFreeFloating($2[0], name)
                yylex.(*Parser).setFreeFloating(name, freefloating.End, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating(alias, freefloating.Start, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_const_declarations:
        use_const_declarations ',' use_const_declaration
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   use_const_declaration
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_const_declaration:
        namespace_name
            {
                name := name.NewName($1)
                $$ = stmt.NewUse(nil, name, nil)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_AS T_STRING
            {
                name := name.NewName($1)
                alias := node.NewIdentifier($3.Value)
                $$ = stmt.NewUse(nil, name, alias)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                alias.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)
                yylex.(*Parser).setFreeFloating(name, freefloating.End, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(alias, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := name.NewName($2)
                $$ = stmt.NewUse(nil, name, nil)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Slash, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).MoveFreeFloating($2[0], name)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                name := name.NewName($2)
                alias := node.NewIdentifier($4.Value)
                $$ = stmt.NewUse(nil, name, alias)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                alias.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($2, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Slash, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).MoveFreeFloating($2[0], name)
                yylex.(*Parser).setFreeFloating(name, freefloating.End, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating(alias, freefloating.Start, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

constant_declaration:
        constant_declaration ',' T_STRING '=' static_scalar
            {
                name := node.NewIdentifier($3.Value)
                constant := stmt.NewConstant(name, $5, "")
                constList := $1.(*stmt.ConstList)
                lastConst := lastNode(constList.Consts)
                constList.Consts = append(constList.Consts, constant)
                $$ = $1

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                constant.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $5))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeNodeListPosition($1, constList.Consts))

                // save comments
                yylex.(*Parser).setFreeFloating(lastConst, freefloating.End, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(constant, freefloating.Start, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating(constant, freefloating.Name, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONST T_STRING '=' static_scalar
            {
                name := node.NewIdentifier($2.Value)
                constant := stmt.NewConstant(name, $4, "")
                constList := []node.Node{constant}
                $$ = stmt.NewConstList(constList)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                constant.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($2, $4))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, constList))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(constant, freefloating.Start, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(constant, freefloating.Name, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

inner_statement_list:
        inner_statement_list inner_statement
            {
                if inlineHtmlNode, ok := $2.(*stmt.InlineHtml); ok && len($1) > 0 {
                    prevNode := lastNode($1)
                    yylex.(*Parser).splitSemiColonAndPhpCloseTag(inlineHtmlNode, prevNode)
                }

                if $2 != nil {
                    $$ = append($1, $2)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                $$ = []node.Node{} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


inner_statement:
        error
            {
                // error
                $$ = nil

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   statement
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   function_declaration_statement
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_declaration_statement
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_HALT_COMPILER '(' ')' ';'
            {
                $$ = stmt.NewHaltCompiler()

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.HaltCompiller, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.OpenParenthesisToken, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.CloseParenthesisToken, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


statement:
        unticked_statement
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STRING ':'
            {
                label := node.NewIdentifier($1.Value)
                $$ = stmt.NewLabel(label)

                // save position
                label.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Label, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unticked_statement:
        '{' inner_statement_list '}'
            {
                $$ = stmt.NewStmtList($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_IF parenthesis_expr statement elseif_list else_single
            {
                $$ = stmt.NewIf($2, $3, $4, $5)

                // save position
                if $5 != nil {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $5))
                } else if len($4) > 0 {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $4))
                } else {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))
                }

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                if len((*$2.GetFreeFloating())[freefloating.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($$, freefloating.If, (*$2.GetFreeFloating())[freefloating.OpenParenthesisToken][:len((*$2.GetFreeFloating())[freefloating.OpenParenthesisToken])-1]); delete((*$2.GetFreeFloating()), freefloating.OpenParenthesisToken)
                }
                if len((*$2.GetFreeFloating())[freefloating.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Expr, (*$2.GetFreeFloating())[freefloating.CloseParenthesisToken][:len((*$2.GetFreeFloating())[freefloating.CloseParenthesisToken])-1]); delete((*$2.GetFreeFloating()), freefloating.CloseParenthesisToken)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_IF parenthesis_expr ':' inner_statement_list new_elseif_list new_else_single T_ENDIF ';'
            {
                stmts := stmt.NewStmtList($4)
                $$ = stmt.NewAltIf($2, stmts, $5, $6)

                // save position
                stmts.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($4))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $8))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                if len((*$2.GetFreeFloating())[freefloating.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($$, freefloating.If, (*$2.GetFreeFloating())[freefloating.OpenParenthesisToken][:len((*$2.GetFreeFloating())[freefloating.OpenParenthesisToken])-1]); delete((*$2.GetFreeFloating()), freefloating.OpenParenthesisToken)
                }
                if len((*$2.GetFreeFloating())[freefloating.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Expr, (*$2.GetFreeFloating())[freefloating.CloseParenthesisToken][:len((*$2.GetFreeFloating())[freefloating.CloseParenthesisToken])-1]); delete((*$2.GetFreeFloating()), freefloating.CloseParenthesisToken)
                }
                yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $7.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $8.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($8))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_WHILE parenthesis_expr while_statement
            {
                switch n := $3.(type) {
                case *stmt.While :
                    n.Cond = $2
                case *stmt.AltWhile :
                    n.Cond = $2
                }

                $$ = $3

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                if len((*$2.GetFreeFloating())[freefloating.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($$, freefloating.While, (*$2.GetFreeFloating())[freefloating.OpenParenthesisToken][:len((*$2.GetFreeFloating())[freefloating.OpenParenthesisToken])-1]); delete((*$2.GetFreeFloating()), freefloating.OpenParenthesisToken)
                }
                if len((*$2.GetFreeFloating())[freefloating.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Expr, (*$2.GetFreeFloating())[freefloating.CloseParenthesisToken][:len((*$2.GetFreeFloating())[freefloating.CloseParenthesisToken])-1]); delete((*$2.GetFreeFloating()), freefloating.CloseParenthesisToken)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DO statement T_WHILE parenthesis_expr ';'
            {
                $$ = stmt.NewDo($2, $4)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $5))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)
                if len((*$4.GetFreeFloating())[freefloating.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($$, freefloating.While, (*$4.GetFreeFloating())[freefloating.OpenParenthesisToken][:len((*$4.GetFreeFloating())[freefloating.OpenParenthesisToken])-1]); delete((*$4.GetFreeFloating()), freefloating.OpenParenthesisToken)
                }
                if len((*$4.GetFreeFloating())[freefloating.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Expr, (*$4.GetFreeFloating())[freefloating.CloseParenthesisToken][:len((*$4.GetFreeFloating())[freefloating.CloseParenthesisToken])-1]); delete((*$4.GetFreeFloating()), freefloating.CloseParenthesisToken)
                }
                yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $5.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($5))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOR '(' for_expr ';' for_expr ';' for_expr ')' for_statement
            {
                switch n := $9.(type) {
                case *stmt.For :
                    n.Init = $3
                    n.Cond = $5
                    n.Loop = $7
                case *stmt.AltFor :
                    n.Init = $3
                    n.Cond = $5
                    n.Loop = $7
                }

                $$ = $9

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $9))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.For, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.InitExpr, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.CondExpr, $6.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.IncExpr, $8.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_SWITCH parenthesis_expr switch_case_list
            {
                switch n := $3.(type) {
                case *stmt.Switch:
                    n.Cond = $2
                case *stmt.AltSwitch:
                    n.Cond = $2
                default:
                    panic("unexpected node type")
                }

                $$ = $3

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                if len((*$2.GetFreeFloating())[freefloating.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Switch, (*$2.GetFreeFloating())[freefloating.OpenParenthesisToken][:len((*$2.GetFreeFloating())[freefloating.OpenParenthesisToken])-1]); delete((*$2.GetFreeFloating()), freefloating.OpenParenthesisToken)
                }
                if len((*$2.GetFreeFloating())[freefloating.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Expr, (*$2.GetFreeFloating())[freefloating.CloseParenthesisToken][:len((*$2.GetFreeFloating())[freefloating.CloseParenthesisToken])-1]); delete((*$2.GetFreeFloating()), freefloating.CloseParenthesisToken)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BREAK ';'
            {
                $$ = stmt.NewBreak(nil)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BREAK expr ';'
            {
                $$ = stmt.NewBreak($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONTINUE ';'
            {
                $$ = stmt.NewContinue(nil)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONTINUE expr ';'
            {
                $$ = stmt.NewContinue($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_RETURN ';'
            {
                $$ = stmt.NewReturn(nil)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_RETURN expr_without_variable ';'
            {
                $$ = stmt.NewReturn($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_RETURN variable ';'
            {
                $$ = stmt.NewReturn($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   yield_expr ';'
            {
                $$ = stmt.NewExpression($1)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_GLOBAL global_var_list ';'
            {
                $$ = stmt.NewGlobal($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.VarList, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC static_var_list ';'
            {
                $$ = stmt.NewStatic($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.VarList, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ECHO echo_expr_list ';'
            {
                $$ = stmt.NewEcho($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Echo, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INLINE_HTML
            {
                $$ = stmt.NewInlineHtml($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr ';'
            {
                $$ = stmt.NewExpression($1)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_UNSET '(' unset_variables ')' ';'
            {
                $$ = stmt.NewUnset($3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $5))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Unset, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.VarList, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.CloseParenthesisToken, $5.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($5))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOREACH '(' variable T_AS foreach_variable foreach_optional_arg ')' foreach_statement
            {
                if $6 == nil {
                    switch n := $8.(type) {
                    case *stmt.Foreach :
                        n.Expr = $3
                        n.Variable = $5
                    case *stmt.AltForeach :
                        n.Expr = $3
                        n.Variable = $5
                    }
                } else {
                    switch n := $8.(type) {
                    case *stmt.Foreach :
                        n.Expr = $3
                        n.Key = $5
                        n.Variable = $6
                    case *stmt.AltForeach :
                        n.Expr = $3
                        n.Key = $5
                        n.Variable = $6
                    }
                }
                
                $$ = $8

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $8))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Foreach, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $4.FreeFloating)
                if $6 != nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Key, (*$6.GetFreeFloating())[freefloating.Key]); delete((*$6.GetFreeFloating()), freefloating.Key)
                }
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $7.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOREACH '(' expr_without_variable T_AS foreach_variable foreach_optional_arg ')' foreach_statement
            {
                if $6 == nil {
                    switch n := $8.(type) {
                    case *stmt.Foreach :
                        n.Expr = $3
                        n.Variable = $5
                    case *stmt.AltForeach :
                        n.Expr = $3
                        n.Variable = $5
                    }
                } else {
                    switch n := $8.(type) {
                    case *stmt.Foreach :
                        n.Expr = $3
                        n.Key = $5
                        n.Variable = $6
                    case *stmt.AltForeach :
                        n.Expr = $3
                        n.Key = $5
                        n.Variable = $6
                    }
                }
                
                // save position
                $$ = $8

                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $8))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Foreach, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $4.FreeFloating)
                if $6 != nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Key, (*$6.GetFreeFloating())[freefloating.Key]); delete((*$6.GetFreeFloating()), freefloating.Key)
                }
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $7.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DECLARE '(' declare_list ')' declare_statement
            {
                $$ = $5
                $$.(*stmt.Declare).Consts = $3

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $5))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Declare, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.ConstList, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ';'
            {
                $$ = stmt.NewNop()

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRY '{' inner_statement_list '}' catch_statement finally_statement
            {
                $$ = stmt.NewTry($3, $5, $6)

                // save position
                if $6 == nil {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $5))
                } else {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $6))
                }

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Try, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_THROW expr ';'
            {
                $$ = stmt.NewThrow($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_GOTO T_STRING ';'
            {
                label := node.NewIdentifier($2.Value)
                $$ = stmt.NewGoto(label)

                // save position
                label.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(label, freefloating.Start, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Label, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

catch_statement:
        /* empty */
            {
                $$ = []node.Node{} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CATCH '(' fully_qualified_class_name T_VARIABLE ')' '{' inner_statement_list '}' additional_catches
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($4.Value, isDollar))
                variable := expr.NewVariable(identifier)
                catchNode := stmt.NewCatch([]node.Node{$3}, variable, $7)
                $$ = append([]node.Node{catchNode}, $9...)

                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                catchNode.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $8))

                // save comments
                yylex.(*Parser).setFreeFloating(catchNode, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(catchNode, freefloating.Catch, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(variable, freefloating.Start, $4.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)
                yylex.(*Parser).setFreeFloating(catchNode, freefloating.Var, $5.FreeFloating)
                yylex.(*Parser).setFreeFloating(catchNode, freefloating.Cond, $6.FreeFloating)
                yylex.(*Parser).setFreeFloating(catchNode, freefloating.Stmts, $8.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

finally_statement:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINALLY '{' inner_statement_list '}'
            {
                $$ = stmt.NewFinally($3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Finally, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

additional_catches:
        non_empty_additional_catches
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                $$ = []node.Node{} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_additional_catches:
        additional_catch
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_additional_catches additional_catch
            {
                $$ = append($1, $2) 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

additional_catch:
        T_CATCH '(' fully_qualified_class_name T_VARIABLE ')' '{' inner_statement_list '}'
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($4.Value, isDollar))
                variable := expr.NewVariable(identifier)
                $$ = stmt.NewCatch([]node.Node{$3}, variable, $7)

                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $8))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Catch, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(variable, freefloating.Start, $4.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $5.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $6.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $8.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unset_variables:
        unset_variable
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   unset_variables ',' unset_variable
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unset_variable:
        variable
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

function_declaration_statement:
        unticked_function_declaration_statement
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_declaration_statement:
        unticked_class_declaration_statement
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
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

unticked_function_declaration_statement:
        function is_reference T_STRING '(' parameter_list ')' '{' inner_statement_list '}'
            {
                name := node.NewIdentifier($3.Value)
                $$ = stmt.NewFunction(name, $2 != nil, $5, nil, $8, "")

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $9))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                if $2 != nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Function, $2.FreeFloating)
                    yylex.(*Parser).setFreeFloating(name, freefloating.Start, $3.FreeFloating)
                } else {
                    yylex.(*Parser).setFreeFloating(name, freefloating.Start, $3.FreeFloating)
                }
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.ParamList, $6.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Params, $7.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $9.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unticked_class_declaration_statement:
        class_entry_type T_STRING extends_from implements_list '{' class_statement_list '}'
            {
                name := node.NewIdentifier($2.Value)
                switch n := $1.(type) {
                    case *stmt.Class :
                        n.ClassName = name
                        n.Stmts = $6
                        n.Extends = $3
                        n.Implements = $4

                    case *stmt.Trait :
                        // TODO: is it possible that trait extend or implement
                        n.TraitName = name
                        n.Stmts = $6
                }
                $$ = $1

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $7))

                // save comments
                yylex.(*Parser).setFreeFloating(name, freefloating.Start, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $5.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $7.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   interface_entry T_STRING interface_extends_list '{' class_statement_list '}'
            {
                name := node.NewIdentifier($2.Value)
                $$ = stmt.NewInterface(name, $3, $5, "")

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $6))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(name, freefloating.Start, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $6.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


class_entry_type:
        T_CLASS
            {
                $$ = stmt.NewClass(nil, nil, nil, nil, nil, nil, "")

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ABSTRACT T_CLASS
            {
                classModifier := node.NewIdentifier($1.Value)
                $$ = stmt.NewClass(nil, []node.Node{classModifier}, nil, nil, nil, nil, "")

                // save position
                classModifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.ModifierList, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRAIT
            {
                $$ = stmt.NewTrait(nil, nil, "")

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINAL T_CLASS
            {
                classModifier := node.NewIdentifier($1.Value)
                $$ = stmt.NewClass(nil, []node.Node{classModifier}, nil, nil, nil, nil, "")

                // save position
                classModifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.ModifierList, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

extends_from:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXTENDS fully_qualified_class_name
            {
                $$ = stmt.NewClassExtends($2);

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

interface_entry:
        T_INTERFACE
            {
                $$ = $1
            }
;

interface_extends_list:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXTENDS interface_list
            {
                $$ = stmt.NewInterfaceExtends($2);

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

implements_list:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_IMPLEMENTS interface_list
            {
                $$ = stmt.NewClassImplements($2);

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

interface_list:
        fully_qualified_class_name
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   interface_list ',' fully_qualified_class_name
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

foreach_optional_arg:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOUBLE_ARROW foreach_variable
            {
                $$ = $2

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Key, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

foreach_variable:
        variable
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' variable
            {
                $$ = expr.NewReference($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LIST '(' assignment_list ')'
            {
                $$ = expr.NewList($3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.List, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.ArrayPairList, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

for_statement:
        statement
            {
                $$ = stmt.NewFor(nil, nil, nil, $1)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDFOR ';'
            {
                stmtList := stmt.NewStmtList($2)
                $$ = stmt.NewAltFor(nil, nil, nil, stmtList)

                // save position
                stmtList.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

foreach_statement:
        statement
            {
                $$ = stmt.NewForeach(nil, nil, nil, $1)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDFOREACH ';'
            {
                stmtList := stmt.NewStmtList($2)
                $$ = stmt.NewAltForeach(nil, nil, nil, stmtList)

                // save position
                stmtList.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


declare_statement:
        statement
            {
                $$ = stmt.NewDeclare(nil, $1, false)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDDECLARE ';'
            {
                stmtList := stmt.NewStmtList($2)
                $$ = stmt.NewDeclare(nil, stmtList, true)

                // save position
                stmtList.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


declare_list:
        T_STRING '=' static_scalar
            {
                name := node.NewIdentifier($1.Value)
                constant := stmt.NewConstant(name, $3, "")
                $$ = []node.Node{constant}

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                constant.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating(constant, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(constant, freefloating.Name, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   declare_list ',' T_STRING '=' static_scalar
            {
                name := node.NewIdentifier($3.Value)
                constant := stmt.NewConstant(name, $5, "")
                $$ = append($1, constant)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                constant.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $5))

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(constant, freefloating.Start, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating(constant, freefloating.Name, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


switch_case_list:
        '{' case_list '}'
            {
                caseList := stmt.NewCaseList($2)
                $$ = stmt.NewSwitch(nil, caseList)

                // save position
                caseList.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating(caseList, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(caseList, freefloating.CaseListEnd, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' ';' case_list '}'
            {
                caseList := stmt.NewCaseList($3)
                $$ = stmt.NewSwitch(nil, caseList)

                // save position
                caseList.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating(caseList, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(caseList, freefloating.CaseListStart, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating(caseList, freefloating.CaseListEnd, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' case_list T_ENDSWITCH ';'
            {
                caseList := stmt.NewCaseList($2)
                $$ = stmt.NewAltSwitch(nil, caseList)

                // save position
                caseList.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(caseList, freefloating.CaseListEnd, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' ';' case_list T_ENDSWITCH ';'
            {

                caseList := stmt.NewCaseList($3)
                $$ = stmt.NewAltSwitch(nil, caseList)

                // save position
                caseList.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $5))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(caseList, freefloating.CaseListStart, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating(caseList, freefloating.CaseListEnd, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $5.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($5))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


case_list:
        /* empty */
            {
                $$ = []node.Node{} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   case_list T_CASE expr case_separator inner_statement_list
            {
                _case := stmt.NewCase($3, $5)
                $$ = append($1, _case)

                // save position
                _case.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($2, $5))

                // save comments
                yylex.(*Parser).setFreeFloating(_case, freefloating.Start, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(_case, freefloating.Expr, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating(_case, freefloating.CaseSeparator, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   case_list T_DEFAULT case_separator inner_statement_list
            {
                _default := stmt.NewDefault($4)
                $$ = append($1, _default)

                // save position
                _default.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($2, $4))

                // save comments
                yylex.(*Parser).setFreeFloating(_default, freefloating.Start, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(_default, freefloating.Default, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating(_default, freefloating.CaseSeparator, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
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
                $$ = stmt.NewWhile(nil, $1)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDWHILE ';'
            {
                stmtList := stmt.NewStmtList($2)
                $$ = stmt.NewAltWhile(nil, stmtList)

                // save position
                stmtList.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;



elseif_list:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   elseif_list T_ELSEIF parenthesis_expr statement
            {
                _elseIf := stmt.NewElseIf($3, $4)
                $$ = append($1, _elseIf)

                // save position
                _elseIf.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($2, $4))

                // save comments
                yylex.(*Parser).setFreeFloating(_elseIf, freefloating.Start, $2.FreeFloating)
                if len((*$3.GetFreeFloating())[freefloating.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating(_elseIf, freefloating.ElseIf, (*$3.GetFreeFloating())[freefloating.OpenParenthesisToken][:len((*$3.GetFreeFloating())[freefloating.OpenParenthesisToken])-1]); delete((*$3.GetFreeFloating()), freefloating.OpenParenthesisToken)
                }
                if len((*$3.GetFreeFloating())[freefloating.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating(_elseIf, freefloating.Expr, (*$3.GetFreeFloating())[freefloating.CloseParenthesisToken][:len((*$3.GetFreeFloating())[freefloating.CloseParenthesisToken])-1]); delete((*$3.GetFreeFloating()), freefloating.CloseParenthesisToken)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


new_elseif_list:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_elseif_list T_ELSEIF parenthesis_expr ':' inner_statement_list
            {
                stmts := stmt.NewStmtList($5)
                _elseIf := stmt.NewAltElseIf($3, stmts)
                $$ = append($1, _elseIf)

                // save position
                stmts.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($5))
                _elseIf.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($2, $5))

                // save comments
                yylex.(*Parser).setFreeFloating(_elseIf, freefloating.Start, $2.FreeFloating)
                if len((*$3.GetFreeFloating())[freefloating.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating(_elseIf, freefloating.ElseIf, (*$3.GetFreeFloating())[freefloating.OpenParenthesisToken][:len((*$3.GetFreeFloating())[freefloating.OpenParenthesisToken])-1]); delete((*$3.GetFreeFloating()), freefloating.OpenParenthesisToken)
                }
                if len((*$3.GetFreeFloating())[freefloating.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating(_elseIf, freefloating.Expr, (*$3.GetFreeFloating())[freefloating.CloseParenthesisToken][:len((*$3.GetFreeFloating())[freefloating.CloseParenthesisToken])-1]); delete((*$3.GetFreeFloating()), freefloating.CloseParenthesisToken)
                }
                yylex.(*Parser).setFreeFloating(_elseIf, freefloating.Cond, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


else_single:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ELSE statement
            {
                $$ = stmt.NewElse($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


new_else_single:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ELSE ':' inner_statement_list
            {
                stmts := stmt.NewStmtList($3)
                $$ = stmt.NewAltElse(stmts)

                // save position
                stmts.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Else, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


parameter_list:
        non_empty_parameter_list
            {
                $$ = $1; 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_parameter_list:
        parameter
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_parameter_list ',' parameter
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

parameter:
        optional_class_type is_reference is_variadic T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($4.Value, isDollar))
                variable := expr.NewVariable(identifier)
                $$ = node.NewParameter($1, variable, nil, $2 != nil, $3 != nil)

                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                if $1 != nil {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                } else if $2 != nil {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($2, $4))
                } else if $3 != nil {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($3, $4))
                } else {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                }

                // save comments
                if $1 != nil {
                    yylex.(*Parser).MoveFreeFloating($1, $$)
                }
                if $2 != nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.OptionalType, $2.FreeFloating)
                }
                if $3 != nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Ampersand, $3.FreeFloating)
                }
                yylex.(*Parser).setFreeFloating($$, freefloating.Variadic, $4.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)

                // normalize
                if $3 == nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Ampersand, (*$$.GetFreeFloating())[freefloating.Variadic]); delete((*$$.GetFreeFloating()), freefloating.Variadic)
                }
                if $2 == nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.OptionalType, (*$$.GetFreeFloating())[freefloating.Ampersand]); delete((*$$.GetFreeFloating()), freefloating.Ampersand)
                }
                if $1 == nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Start, (*$$.GetFreeFloating())[freefloating.OptionalType]); delete((*$$.GetFreeFloating()), freefloating.OptionalType)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   optional_class_type is_reference is_variadic T_VARIABLE '=' static_scalar
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($4.Value, isDollar))
                variable := expr.NewVariable(identifier)
                $$ = node.NewParameter($1, variable, $6, $2 != nil, $3 != nil)

                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                if $1 != nil {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $6))
                } else if $2 != nil {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($2, $6))
                } else if $3 != nil {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $6))
                } else {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($4, $6))
                }

                // save comments
                if $1 != nil {
                    yylex.(*Parser).MoveFreeFloating($1, $$)
                }
                if $2 != nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.OptionalType, $2.FreeFloating)
                }
                if $3 != nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Ampersand, $3.FreeFloating)
                }
                yylex.(*Parser).setFreeFloating($$, freefloating.Variadic, $4.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $5.FreeFloating)

                // normalize
                if $3 == nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Ampersand, (*$$.GetFreeFloating())[freefloating.Variadic]); delete((*$$.GetFreeFloating()), freefloating.Variadic)
                }
                if $2 == nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.OptionalType, (*$$.GetFreeFloating())[freefloating.Ampersand]); delete((*$$.GetFreeFloating()), freefloating.Ampersand)
                }
                if $1 == nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Start, (*$$.GetFreeFloating())[freefloating.OptionalType]); delete((*$$.GetFreeFloating()), freefloating.OptionalType)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


optional_class_type:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ARRAY
            {
                $$ = node.NewIdentifier($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CALLABLE
            {
                $$ = node.NewIdentifier($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   fully_qualified_class_name
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


function_call_parameter_list:
        '(' ')'
            {
                $$ = node.NewArgumentList(nil)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.ArgumentList, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' non_empty_function_call_parameter_list ')'
            {
                $$ = node.NewArgumentList($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.ArgumentList, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' yield_expr ')'
            {
                arg := node.NewArgument($2, false, false)
                $$ = node.NewArgumentList([]node.Node{arg})

                // save position
                arg.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.ArgumentList, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


non_empty_function_call_parameter_list:
        function_call_parameter
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_function_call_parameter_list ',' function_call_parameter
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

function_call_parameter:
        expr_without_variable
            {
                $$ = node.NewArgument($1, false, false)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($1))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable
            {
                $$ = node.NewArgument($1, false, false)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($1))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' w_variable
            {
                $$ = node.NewArgument($2, false, true)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ELLIPSIS expr
            {
                $$ = node.NewArgument($2, true, false)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

global_var_list:
        global_var_list ',' global_var
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   global_var
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


global_var:
        T_VARIABLE
            {
                name := node.NewIdentifier(strings.TrimLeftFunc($1.Value, isDollar))
                $$ = expr.NewVariable(name)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).addDollarToken($$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' r_variable
            {
                $$ = expr.NewVariable($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Dollar, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' '{' expr '}'
            {
                $$ = expr.NewVariable($3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Dollar, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).setFreeFloating($3, freefloating.Start, append($2.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($2), (*$3.GetFreeFloating())[freefloating.Start]...)...))
                yylex.(*Parser).setFreeFloating($3, freefloating.End, append((*$3.GetFreeFloating())[freefloating.End], append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


static_var_list:
        static_var_list ',' T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($3.Value, isDollar))
                variable := expr.NewVariable(identifier)
                staticVar := stmt.NewStaticVar(variable, nil)
                $$ = append($1, staticVar)
                
                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                staticVar.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(staticVar, freefloating.Start, $3.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_var_list ',' T_VARIABLE '=' static_scalar
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($3.Value, isDollar))
                variable := expr.NewVariable(identifier)
                staticVar := stmt.NewStaticVar(variable, $5)
                $$ = append($1, staticVar)
                
                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                staticVar.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $5))

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(staticVar, freefloating.Start, $3.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)
                yylex.(*Parser).setFreeFloating(staticVar, freefloating.Var, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($1.Value, isDollar))
                variable := expr.NewVariable(identifier)
                staticVar := stmt.NewStaticVar(variable, nil)
                $$ = []node.Node{staticVar}
                
                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                staticVar.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating(staticVar, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '=' static_scalar
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($1.Value, isDollar))
                variable := expr.NewVariable(identifier)
                staticVar := stmt.NewStaticVar(variable, $3)
                $$ = []node.Node{staticVar}
                
                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                staticVar.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating(staticVar, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)
                yylex.(*Parser).setFreeFloating(staticVar, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


class_statement_list:
        class_statement_list class_statement
            {
                $$ = append($1, $2) 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                $$ = []node.Node{} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


class_statement:
        variable_modifiers class_variable_declaration ';'
            {
                $$ = stmt.NewPropertyList($1, nil, $2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.PropertyList, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_constant_declaration ';'
            {
                $$ = $1

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.ConstList, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_use_statement
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   method_modifiers function is_reference T_STRING '(' parameter_list ')' method_body
            {
                name := node.NewIdentifier($4.Value)
                $$ = stmt.NewClassMethod(name, $1, $3 != nil, $6, nil, $8, "")

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                if $1 == nil {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($2, $8))
                } else {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListNodePosition($1, $8))
                }

                // save comments
                if len($1) > 0 {
                    yylex.(*Parser).MoveFreeFloating($1[0], $$)
                    yylex.(*Parser).setFreeFloating($$, freefloating.ModifierList, $2.FreeFloating)
                } else {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Start, $2.FreeFloating)
                }
                if $3 == nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Function, $4.FreeFloating)
                } else {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Function, $3.FreeFloating)
                    yylex.(*Parser).setFreeFloating($$, freefloating.Ampersand, $4.FreeFloating)
                }
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $5.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.ParameterList, $7.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_use_statement:
        T_USE trait_list trait_adaptations
            {
                $$ = stmt.NewTraitUse($2, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_list:
        fully_qualified_class_name
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_list ',' fully_qualified_class_name
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptations:
        ';'
            {
                $$ = stmt.NewNop()

                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' trait_adaptation_list '}'
            {
                $$ = stmt.NewTraitAdaptationList($2)

                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.AdaptationList, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptation_list:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_trait_adaptation_list
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_trait_adaptation_list:
        trait_adaptation_statement
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_trait_adaptation_list trait_adaptation_statement
            {
                $$ = append($1, $2) 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptation_statement:
        trait_precedence ';'
            {
                $$ = $1;

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.NameList, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_alias ';'
            {
                $$ = $1;

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Alias, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_precedence:
        trait_method_reference_fully_qualified T_INSTEADOF trait_reference_list
            {
                $$ = stmt.NewTraitUsePrecedence($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeNodeListPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Ref, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_reference_list:
        fully_qualified_class_name
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_reference_list ',' fully_qualified_class_name
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_method_reference:
        T_STRING
            {
                name := node.NewIdentifier($1.Value)
                $$ = stmt.NewTraitMethodRef(nil, name)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_method_reference_fully_qualified
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_method_reference_fully_qualified:
        fully_qualified_class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                target := node.NewIdentifier($3.Value)
                $$ = stmt.NewTraitMethodRef($1, target)

                // save position
                target.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(target, freefloating.Start, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_alias:
        trait_method_reference T_AS trait_modifiers T_STRING
            {
                alias := node.NewIdentifier($4.Value)
                $$ = stmt.NewTraitUseAlias($1, $3, alias)

                // save position
                alias.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Ref, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(alias, freefloating.Start, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_method_reference T_AS member_modifier
            {
                $$ = stmt.NewTraitUseAlias($1, $3, nil)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Ref, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_modifiers:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   member_modifier
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

method_body:
        ';' /* abstract method */ 
            {
                $$ = stmt.NewNop()

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' inner_statement_list '}'
            {
                $$ = stmt.NewStmtList($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

variable_modifiers:
        non_empty_member_modifiers
            {
                $$ = $1; 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VAR
            {
                modifier := node.NewIdentifier($1.Value)
                $$ = []node.Node{modifier}

                // save position
                modifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating(modifier, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

method_modifiers:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_member_modifiers
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_member_modifiers:
        member_modifier
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_member_modifiers member_modifier
            {
                $$ = append($1, $2) 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

member_modifier:
        T_PUBLIC
            {
                $$ = node.NewIdentifier($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PROTECTED
            {
                $$ = node.NewIdentifier($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PRIVATE
            {
                $$ = node.NewIdentifier($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC
            {
                $$ = node.NewIdentifier($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ABSTRACT
            {
                $$ = node.NewIdentifier($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINAL
            {
                $$ = node.NewIdentifier($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_variable_declaration:
        class_variable_declaration ',' T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($3.Value, isDollar))
                variable := expr.NewVariable(identifier)
                property := stmt.NewProperty(variable, nil, "")
                $$ = append($1, property)
                
                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                property.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(property, freefloating.Start, $3.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_variable_declaration ',' T_VARIABLE '=' static_scalar
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($3.Value, isDollar))
                variable := expr.NewVariable(identifier)
                property := stmt.NewProperty(variable, $5, "")
                $$ = append($1, property)
                
                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                property.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $5))
                
                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(property, freefloating.Start, $3.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)
                yylex.(*Parser).setFreeFloating(property, freefloating.Var, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($1.Value, isDollar))
                variable := expr.NewVariable(identifier)
                property := stmt.NewProperty(variable, nil, "")
                $$ = []node.Node{property}
                
                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                property.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                
                // save comments
                yylex.(*Parser).setFreeFloating(property, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '=' static_scalar
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($1.Value, isDollar))
                variable := expr.NewVariable(identifier)
                property := stmt.NewProperty(variable, $3, "")
                $$ = []node.Node{property}
                
                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                property.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))
                
                // save comments
                yylex.(*Parser).setFreeFloating(property, freefloating.Start, $2.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)
                yylex.(*Parser).setFreeFloating(property, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_constant_declaration:
        class_constant_declaration ',' T_STRING '=' static_scalar
            {
                name := node.NewIdentifier($3.Value)
                constant := stmt.NewConstant(name, $5, "")
                constList := $1.(*stmt.ClassConstList)
                lastConst := lastNode(constList.Consts)
                constList.Consts = append(constList.Consts, constant)
                $$ = $1

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                constant.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $5))
                $1.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $5))

                // save comments
                yylex.(*Parser).setFreeFloating(lastConst, freefloating.End, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(constant, freefloating.Start, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating(constant, freefloating.Name, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONST T_STRING '=' static_scalar
            {
                name := node.NewIdentifier($2.Value)
                constant := stmt.NewConstant(name, $4, "")
                $$ = stmt.NewClassConstList(nil, []node.Node{constant})

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                constant.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($2, $4))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(constant, freefloating.Start, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(constant, freefloating.Name, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

echo_expr_list:
        echo_expr_list ',' expr
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


for_expr:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_for_expr
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_for_expr:
        non_empty_for_expr ',' expr
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

chaining_method_or_property:
        chaining_method_or_property variable_property
            {
                $$ = append($1, $2...) 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_property
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

chaining_dereference:
        chaining_dereference '[' dim_offset ']'
            {
                fetch := expr.NewArrayDimFetch(nil, $3)
                $$ = append($1, fetch)
                
                // save position
                fetch.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($3))

                // save comments
                yylex.(*Parser).setFreeFloating(fetch, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating(fetch, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' dim_offset ']'
            {
                fetch := expr.NewArrayDimFetch(nil, $2)
                $$ = []node.Node{fetch}
                
                // save position
                fetch.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($2))
                
                // save comments
                yylex.(*Parser).setFreeFloating(fetch, freefloating.Var, append($1.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($1)...))
                yylex.(*Parser).setFreeFloating(fetch, freefloating.Expr, append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

chaining_instance_call:
        chaining_dereference chaining_method_or_property
            {
                $$ = append($1, $2...) 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   chaining_dereference
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   chaining_method_or_property
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

instance_call:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   chaining_instance_call
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

new_expr:
        T_NEW class_name_reference ctor_arguments
            {

                if $3 != nil {
                    $$ = expr.NewNew($2, $3.(*node.ArgumentList))
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3))
                } else {
                    $$ = expr.NewNew($2, nil)
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                }

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

expr_without_variable:
        T_LIST '(' assignment_list ')' '=' expr
            {
                listNode := expr.NewList($3)
                $$ = assign.NewAssign(listNode, $6)

                // save position
                listNode.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $6))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(listNode, freefloating.List, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(listNode, freefloating.ArrayPairList, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $5.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' expr
            {
                $$ = assign.NewAssign($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' '&' variable
            {
                $$ = assign.NewReference($1, $4)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Equal, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' '&' T_NEW class_name_reference ctor_arguments
            {
                var _new *expr.New

                if $6 != nil {
                    _new = expr.NewNew($5, $6.(*node.ArgumentList))
                } else {
                    _new = expr.NewNew($5, nil)
                }
                $$ = assign.NewReference($1, _new)

                // save position
                if $6 != nil {
                    _new.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($4, $6))
                } else {
                    _new.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($4, $5))
                }
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, _new))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Equal, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating(_new, freefloating.Start, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLONE expr
            {
                $$ = expr.NewClone($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_PLUS_EQUAL expr
            {
                $$ = assign.NewPlus($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MINUS_EQUAL expr
            {
                $$ = assign.NewMinus($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MUL_EQUAL expr
            {
                $$ = assign.NewMul($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_POW_EQUAL expr
            {
                $$ = assign.NewPow($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_DIV_EQUAL expr
            {
                $$ = assign.NewDiv($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_CONCAT_EQUAL expr
            {
                $$ = assign.NewConcat($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MOD_EQUAL expr
            {
                $$ = assign.NewMod($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_AND_EQUAL expr
            {
                $$ = assign.NewBitwiseAnd($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_OR_EQUAL expr
            {
                $$ = assign.NewBitwiseOr($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_XOR_EQUAL expr
            {
                $$ = assign.NewBitwiseXor($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_SL_EQUAL expr
            {
                $$ = assign.NewShiftLeft($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_SR_EQUAL expr
            {
                $$ = assign.NewShiftRight($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   rw_variable T_INC
            {
                $$ = expr.NewPostInc($1)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INC rw_variable
            {
                $$ = expr.NewPreInc($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   rw_variable T_DEC
            {
                $$ = expr.NewPostDec($1)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DEC rw_variable
            {
                $$ = expr.NewPreDec($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_BOOLEAN_OR expr
            {
                $$ = binary.NewBooleanOr($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_BOOLEAN_AND expr
            {
                $$ = binary.NewBooleanAnd($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_OR expr
            {
                $$ = binary.NewLogicalOr($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_AND expr
            {
                $$ = binary.NewLogicalAnd($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_XOR expr
            {
                $$ = binary.NewLogicalXor($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '|' expr
            {
                $$ = binary.NewBitwiseOr($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '&' expr
            {
                $$ = binary.NewBitwiseAnd($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '^' expr
            {
                $$ = binary.NewBitwiseXor($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '.' expr
            {
                $$ = binary.NewConcat($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '+' expr
            {
                $$ = binary.NewPlus($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '-' expr
            {
                $$ = binary.NewMinus($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '*' expr
            {
                $$ = binary.NewMul($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_POW expr
            {
                $$ = binary.NewPow($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '/' expr
            {
                $$ = binary.NewDiv($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '%' expr
            {
                $$ = binary.NewMod($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SL expr
            {
                $$ = binary.NewShiftLeft($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SR expr
            {
                $$ = binary.NewShiftRight($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '+' expr %prec T_INC
            {
                $$ = expr.NewUnaryPlus($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '-' expr %prec T_INC
            {
                $$ = expr.NewUnaryMinus($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '!' expr
            {
                $$ = expr.NewBooleanNot($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '~' expr
            {
                $$ = expr.NewBitwiseNot($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_IDENTICAL expr
            {
                $$ = binary.NewIdentical($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_NOT_IDENTICAL expr
            {
                $$ = binary.NewNotIdentical($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_EQUAL expr
            {
                $$ = binary.NewEqual($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_NOT_EQUAL expr
            {
                $$ = binary.NewNotEqual($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Equal, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '<' expr
            {
                $$ = binary.NewSmaller($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_SMALLER_OR_EQUAL expr
            {
                $$ = binary.NewSmallerOrEqual($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '>' expr
            {
                $$ = binary.NewGreater($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_GREATER_OR_EQUAL expr
            {
                $$ = binary.NewGreaterOrEqual($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_INSTANCEOF class_name_reference
            {
                $$ = expr.NewInstanceOf($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   parenthesis_expr
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)

                yylex.(*Parser).setFreeFloating($1, freefloating.Start, append((*$1.GetFreeFloating())[freefloating.OpenParenthesisToken], (*$1.GetFreeFloating())[freefloating.Start]...)); delete((*$1.GetFreeFloating()), freefloating.OpenParenthesisToken)
                yylex.(*Parser).setFreeFloating($1, freefloating.End, append((*$1.GetFreeFloating())[freefloating.End], (*$1.GetFreeFloating())[freefloating.CloseParenthesisToken]...)); delete((*$1.GetFreeFloating()), freefloating.CloseParenthesisToken)
            }
    |   new_expr
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' new_expr ')' instance_call
            {
                $$ = $2

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, append($1.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($1), (*$$.GetFreeFloating())[freefloating.Start]...)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.End, append((*$$.GetFreeFloating())[freefloating.End], append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...)...))

                for _, n := range($4) {
                    switch nn := n.(type) {
                        case *expr.ArrayDimFetch:
                            nn.Variable = $$
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Variable, $$)
                        
                        case *expr.PropertyFetch:
                            nn.Variable = $$
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Variable, $$)
                        
                        case *expr.MethodCall:
                            nn.Variable = $$
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Variable, $$)
                    }

                    // save position
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($$, n))
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '?' expr ':' expr
            {
                $$ = expr.NewTernary($1, $3, $5)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $5))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.True, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '?' ':' expr
            {
                $$ = expr.NewTernary($1, nil, $4)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.True, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   internal_functions_in_yacc
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INT_CAST expr
            {
                $$ = cast.NewInt($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOUBLE_CAST expr
            {
                $$ = cast.NewDouble($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STRING_CAST expr
            {
                $$ = cast.NewString($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ARRAY_CAST expr
            {
                $$ = cast.NewArray($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_OBJECT_CAST expr
            {
                $$ = cast.NewObject($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BOOL_CAST expr
            {
                $$ = cast.NewBool($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_UNSET_CAST expr
            {
                $$ = cast.NewUnset($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXIT exit_expr
            {
                e := $2.(*expr.Exit)
                $$ = $2

                if (strings.EqualFold($1.Value, "die")) {
                    e.Die = true
                }

                // save position
                if $2.GetPosition() == nil {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                } else {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                }

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '@' expr
            {
                $$ = expr.NewErrorSuppress($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   scalar
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   combined_scalar_offset
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   combined_scalar
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '`' backticks_expr '`'
            {
                $$ = expr.NewShellExec($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PRINT expr
            {
                $$ = expr.NewPrint($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD
            {
                $$ = expr.NewYield(nil, nil)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   function is_reference '(' parameter_list ')' lexical_vars '{' inner_statement_list '}'
            {
                $$ = expr.NewClosure($4, $6, nil, $8, false, $2 != nil, "")

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $9))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                if $2 == nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Function, $3.FreeFloating)
                } else {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Function, $2.FreeFloating)
                    yylex.(*Parser).setFreeFloating($$, freefloating.Ampersand, $3.FreeFloating)
                }
                yylex.(*Parser).setFreeFloating($$, freefloating.ParameterList, $5.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.LexicalVars, $7.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $9.FreeFloating)

                // normalize
                if $6 == nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Params, (*$$.GetFreeFloating())[freefloating.LexicalVars]); delete((*$$.GetFreeFloating()), freefloating.LexicalVars)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC function is_reference '(' parameter_list ')' lexical_vars '{' inner_statement_list '}'
            {
                $$ = expr.NewClosure($5, $7, nil, $9, true, $3 != nil, "")

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $10))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Static, $2.FreeFloating)
                if $3 == nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Function, $4.FreeFloating)
                } else {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Function, $3.FreeFloating)
                    yylex.(*Parser).setFreeFloating($$, freefloating.Ampersand, $4.FreeFloating)
                }
                yylex.(*Parser).setFreeFloating($$, freefloating.ParameterList, $6.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.LexicalVars, $8.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $10.FreeFloating)

                // normalize
                if $7 == nil {
                    yylex.(*Parser).setFreeFloating($$, freefloating.Params, (*$$.GetFreeFloating())[freefloating.LexicalVars]); delete((*$$.GetFreeFloating()), freefloating.LexicalVars)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

yield_expr:
        T_YIELD expr_without_variable
            {
                $$ = expr.NewYield(nil, $2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD variable
            {
                $$ = expr.NewYield(nil, $2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD expr T_DOUBLE_ARROW expr_without_variable
            {
                $$ = expr.NewYield($2, $4)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $4))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD expr T_DOUBLE_ARROW variable
            {
                $$ = expr.NewYield($2, $4)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $4))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

combined_scalar_offset:
        combined_scalar '[' dim_offset ']'
            {
                $$ = expr.NewArrayDimFetch($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   combined_scalar_offset '[' dim_offset ']'
            {
                $$ = expr.NewArrayDimFetch($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONSTANT_ENCAPSED_STRING '[' dim_offset ']'
            {
                str := scalar.NewString($1.Value)
                $$ = expr.NewArrayDimFetch(str, $3)

                // save position
                str.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition(str, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   general_constant '[' dim_offset ']'
            {
                $$ = expr.NewArrayDimFetch($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

combined_scalar:
        T_ARRAY '(' array_pair_list ')'
            {
                $$ = expr.NewArray($3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Array, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.ArrayPairList, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' array_pair_list ']'
            {
                $$ = expr.NewShortArray($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.ArrayPairList, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

function:
    T_FUNCTION
            {
                $$ = $1
            }
;

lexical_vars:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE '(' lexical_var_list ')'
            {
                $$ = expr.NewClosureUse($3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Use, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.LexicalVarList, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

lexical_var_list:
        lexical_var_list ',' T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($3.Value, isDollar))
                variable := expr.NewVariable(identifier)
                $$ = append($1, variable)

                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(variable, freefloating.Start, $3.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   lexical_var_list ',' '&' T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($4.Value, isDollar))
                variable := expr.NewVariable(identifier)
                reference := expr.NewReference(variable)
                $$ = append($1, reference)

                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($4))
                reference.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($3, $4))

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(reference, freefloating.Start, $3.FreeFloating)
                yylex.(*Parser).setFreeFloating(variable, freefloating.Start, $4.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($1.Value, isDollar))
                variable := expr.NewVariable(identifier)
                $$ = []node.Node{variable}

                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating(variable, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($2.Value, isDollar))
                variable := expr.NewVariable(identifier)
                reference := expr.NewReference(variable)
                $$ = []node.Node{reference}
                
                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                reference.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating(reference, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(variable, freefloating.Start, $2.FreeFloating)
                yylex.(*Parser).addDollarToken(variable)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

function_call:
        namespace_name function_call_parameter_list
            {
                name := name.NewName($1)
                $$ = expr.NewFunctionCall(name, $2.(*node.ArgumentList))
                
                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition(name, $2))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name function_call_parameter_list
            {
                funcName := name.NewRelative($3)
                $$ = expr.NewFunctionCall(funcName, $4.(*node.ArgumentList))
                
                // save position
                funcName.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition(funcName, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(funcName, freefloating.Namespace, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name function_call_parameter_list
            {
                funcName := name.NewFullyQualified($2)
                $$ = expr.NewFunctionCall(funcName, $3.(*node.ArgumentList))
                
                // save position
                funcName.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition(funcName, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM variable_name function_call_parameter_list
            {
                $$ = expr.NewStaticCall($1, $3, $4.(*node.ArgumentList))
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects function_call_parameter_list
            {
                $$ = expr.NewStaticCall($1, $3, $4.(*node.ArgumentList))
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_name function_call_parameter_list
            {
                $$ = expr.NewStaticCall($1, $3, $4.(*node.ArgumentList))
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects function_call_parameter_list
            {
                $$ = expr.NewStaticCall($1, $3, $4.(*node.ArgumentList))
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_without_objects function_call_parameter_list
            {
                $$ = expr.NewFunctionCall($1, $2.(*node.ArgumentList))
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $2))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_name:
        T_STATIC
            {
                $$ = node.NewIdentifier($1.Value)
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name 
            {
                $$ = name.NewName($1)
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = name.NewRelative($3)
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Namespace, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = name.NewFullyQualified($2)
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

fully_qualified_class_name:
        namespace_name
            {
                $$ = name.NewName($1)
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = name.NewRelative($3)
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Namespace, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                $$ = name.NewFullyQualified($2)
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_name_reference:
        class_name
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   dynamic_class_name_reference
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

dynamic_class_name_reference:
        base_variable T_OBJECT_OPERATOR object_property dynamic_class_name_variable_properties
            {
                $$ = $1

                // save comments
                yylex.(*Parser).setFreeFloating($3[0], freefloating.Var, $2.FreeFloating)

                for _, n := range($3) {
                    switch nn := n.(type) {
                        case *expr.ArrayDimFetch:
                            nn.Variable = $$
                            $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Variable, $$)
                        
                        case *expr.PropertyFetch:
                            nn.Variable = $$
                            $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Variable, $$)
                    }
                }

                for _, n := range($4) {
                    switch nn := n.(type) {
                        case *expr.ArrayDimFetch:
                            nn.Variable = $$
                            $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Variable, $$)
                        
                        case *expr.PropertyFetch:
                            nn.Variable = $$
                            $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Variable, $$)
                    }
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   base_variable 
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


dynamic_class_name_variable_properties:
        dynamic_class_name_variable_properties dynamic_class_name_variable_property
            {
                $$ = append($1, $2...) 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                $$ = []node.Node{} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


dynamic_class_name_variable_property:
        T_OBJECT_OPERATOR object_property
            {
                $$ = $2
                
                // save comments
                yylex.(*Parser).setFreeFloating($2[0], freefloating.Var, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

exit_expr:
        /* empty */
            {
                $$ = expr.NewExit(nil);

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' ')'
            {
                $$ = expr.NewExit(nil);

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Exit, append($1.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($1)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   parenthesis_expr
            {
                $$ = expr.NewExit($1);

                // save position
                if yylex.(*Parser).currentToken.Value == ")" {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition(yylex.(*Parser).currentToken))
                } else {
                    $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($1))
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Exit, (*$1.GetFreeFloating())[freefloating.OpenParenthesisToken]); delete((*$1.GetFreeFloating()), freefloating.OpenParenthesisToken)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, (*$1.GetFreeFloating())[freefloating.CloseParenthesisToken]); delete((*$1.GetFreeFloating()), freefloating.CloseParenthesisToken)
            }
;

backticks_expr:
        /* empty */
            {
                $$ = []node.Node{} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ENCAPSED_AND_WHITESPACE
            {
                part := scalar.NewEncapsedStringPart($1.Value)
                $$ = []node.Node{part}

                // save position
                part.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   encaps_list
            {
                $$ = $1; 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

ctor_arguments:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   function_call_parameter_list
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

common_scalar:
        T_LNUMBER
            {
                $$ = scalar.NewLnumber($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DNUMBER
            {
                $$ = scalar.NewDnumber($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONSTANT_ENCAPSED_STRING
            {
                $$ = scalar.NewString($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LINE
            {
                $$ = scalar.NewMagicConstant($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FILE
            {
                $$ = scalar.NewMagicConstant($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DIR
            {
                $$ = scalar.NewMagicConstant($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRAIT_C
            {
                $$ = scalar.NewMagicConstant($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_METHOD_C
            {
                $$ = scalar.NewMagicConstant($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FUNC_C
            {
                $$ = scalar.NewMagicConstant($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_C
            {
                $$ = scalar.NewMagicConstant($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC
            {
                encapsed := scalar.NewEncapsedStringPart($2.Value)
                $$ = scalar.NewHeredoc($1.Value, []node.Node{encapsed})

                // save position
                encapsed.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC T_END_HEREDOC
            {
                $$ = scalar.NewHeredoc($1.Value, nil)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_class_constant:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                target := node.NewIdentifier($3.Value)
                $$ = expr.NewClassConstFetch($1, target)
                
                // save position
                target.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(target, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_scalar:
        static_scalar_value
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_scalar_value:
        common_scalar
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_class_name_scalar
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name
            {
                name := name.NewName($1)
                $$ = expr.NewConstFetch(name)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition(name))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                name := name.NewRelative($3)
                $$ = expr.NewConstFetch(name)
                
                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Namespace, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := name.NewFullyQualified($2)
                $$ = expr.NewConstFetch(name)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ARRAY '(' static_array_pair_list ')'
            {
                $$ = expr.NewArray($3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Array, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.ArrayPairList, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' static_array_pair_list ']'
            {
                $$ = expr.NewShortArray($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.ArrayPairList, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_class_constant
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLASS_C
            {
                $$ = scalar.NewMagicConstant($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_operation
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_operation:
        static_scalar_value '[' static_scalar_value ']'
            {
                $$ = expr.NewArrayDimFetch($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '+' static_scalar_value
            {
                $$ = binary.NewPlus($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '-' static_scalar_value
            {
                $$ = binary.NewMinus($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '*' static_scalar_value
            {
                $$ = binary.NewMul($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_POW static_scalar_value
            {
                $$ = binary.NewPow($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '/' static_scalar_value
            {
                $$ = binary.NewDiv($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '%' static_scalar_value
            {
                $$ = binary.NewMod($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '!' static_scalar_value
            {
                $$ = expr.NewBooleanNot($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '~' static_scalar_value
            {
                $$ = expr.NewBitwiseNot($2)
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '|' static_scalar_value
            {
                $$ = binary.NewBitwiseOr($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '&' static_scalar_value
            {
                $$ = binary.NewBitwiseAnd($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '^' static_scalar_value
            {
                $$ = binary.NewBitwiseXor($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_SL static_scalar_value
            {
                $$ = binary.NewShiftLeft($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_SR static_scalar_value
            {
                $$ = binary.NewShiftRight($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '.' static_scalar_value
            {
                $$ = binary.NewConcat($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_LOGICAL_XOR static_scalar_value
            {
                $$ = binary.NewLogicalXor($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_LOGICAL_AND static_scalar_value
            {
                $$ = binary.NewLogicalAnd($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_LOGICAL_OR static_scalar_value
            {
                $$ = binary.NewLogicalOr($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_BOOLEAN_AND static_scalar_value
            {
                $$ = binary.NewBooleanAnd($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_BOOLEAN_OR static_scalar_value
            {
                $$ = binary.NewBooleanOr($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_IDENTICAL static_scalar_value
            {
                $$ = binary.NewIdentical($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_NOT_IDENTICAL static_scalar_value
            {
                $$ = binary.NewNotIdentical($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_EQUAL static_scalar_value
            {
                $$ = binary.NewEqual($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_NOT_EQUAL static_scalar_value
            {
                $$ = binary.NewNotEqual($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Equal, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '<' static_scalar_value
            {
                $$ = binary.NewSmaller($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '>' static_scalar_value
            {
                $$ = binary.NewGreater($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_SMALLER_OR_EQUAL static_scalar_value
            {
                $$ = binary.NewSmallerOrEqual($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_GREATER_OR_EQUAL static_scalar_value
            {
                $$ = binary.NewGreaterOrEqual($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '?' ':' static_scalar_value
            {
                $$ = expr.NewTernary($1, nil, $4)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.True, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '?' static_scalar_value ':' static_scalar_value
            {
                $$ = expr.NewTernary($1, $3, $5)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $5))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.True, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '+' static_scalar_value
            {
                $$ = expr.NewUnaryPlus($2)
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '-' static_scalar_value
            {
                $$ = expr.NewUnaryMinus($2)
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' static_scalar_value ')'
            {
                $$ = $2

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, append($1.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($1), (*$$.GetFreeFloating())[freefloating.Start]...)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.End, append((*$$.GetFreeFloating())[freefloating.End], append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

general_constant:
        class_constant
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name
            {
                name := name.NewName($1)
                $$ = expr.NewConstFetch(name)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewNodeListPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition(name))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                name := name.NewRelative($3)
                $$ = expr.NewConstFetch(name)
                
                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition(name))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(name, freefloating.Namespace, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                name := name.NewFullyQualified($2)
                $$ = expr.NewConstFetch(name)
                
                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition(name))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

scalar:
        T_STRING_VARNAME
            {
                name := node.NewIdentifier($1.Value)
                $$ = expr.NewVariable(name)
                
                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   general_constant
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name_scalar
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   common_scalar
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '"' encaps_list '"'
            {
                $$ = scalar.NewEncapsed($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC encaps_list T_END_HEREDOC
            {
                 $$ = scalar.NewHeredoc($1.Value, $2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLASS_C
            {
                $$ = scalar.NewMagicConstant($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_array_pair_list:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_static_array_pair_list possible_comma
            {
                $$ = $1

                // save comments
                if $2 != nil {
                    yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
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

non_empty_static_array_pair_list:
        non_empty_static_array_pair_list ',' static_scalar_value T_DOUBLE_ARROW static_scalar_value
            {
                arrayItem := expr.NewArrayItem($3, $5, false)
                $$ = append($1, arrayItem)
                
                // save position
                arrayItem.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($3, $5))
                
                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, freefloating.Expr, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_static_array_pair_list ',' static_scalar_value
            {
                arrayItem := expr.NewArrayItem(nil, $3, false)
                $$ = append($1, arrayItem)
                
                // save position
                arrayItem.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($3))
                
                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_DOUBLE_ARROW static_scalar_value
            {
                arrayItem := expr.NewArrayItem($1, $3, false)
                $$ = []node.Node{arrayItem}

                // save position
                arrayItem.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value
            {
                arrayItem := expr.NewArrayItem(nil, $1, false)
                $$ = []node.Node{arrayItem}
                
                // save position
                arrayItem.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($1))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, arrayItem)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

expr:
        r_variable
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr_without_variable
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

parenthesis_expr:
        '(' expr ')'
            {
                $$ = $2

                // save comments
                if len((*$2.GetFreeFloating())[freefloating.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($2, freefloating.Start, append((*$2.GetFreeFloating())[freefloating.OpenParenthesisToken], (*$2.GetFreeFloating())[freefloating.Start]...))
                }
                if len((*$2.GetFreeFloating())[freefloating.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($2, freefloating.End, append((*$2.GetFreeFloating())[freefloating.End], (*$2.GetFreeFloating())[freefloating.CloseParenthesisToken]...))
                }
                yylex.(*Parser).setFreeFloating($2, freefloating.OpenParenthesisToken, append($1.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($1)...))
                yylex.(*Parser).setFreeFloating($2, freefloating.CloseParenthesisToken, append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' yield_expr ')'
            {
                $$ = $2

                // save comments
                if len((*$2.GetFreeFloating())[freefloating.OpenParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($2, freefloating.Start, append((*$2.GetFreeFloating())[freefloating.OpenParenthesisToken], (*$2.GetFreeFloating())[freefloating.Start]...))
                }
                if len((*$2.GetFreeFloating())[freefloating.CloseParenthesisToken]) > 0 {
                    yylex.(*Parser).setFreeFloating($2, freefloating.End, append((*$2.GetFreeFloating())[freefloating.End], (*$2.GetFreeFloating())[freefloating.CloseParenthesisToken]...))
                }
                yylex.(*Parser).setFreeFloating($2, freefloating.OpenParenthesisToken, append($1.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($1)...))
                yylex.(*Parser).setFreeFloating($2, freefloating.CloseParenthesisToken, append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


r_variable:
        variable
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


w_variable:
        variable
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

rw_variable:
        variable
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

variable:
        base_variable_with_function_calls T_OBJECT_OPERATOR object_property method_or_not variable_properties
            {
                $$ = $1

                if $4 != nil {
                    $4[0].(*expr.MethodCall).Method = $3[len($3)-1].(*expr.PropertyFetch).Property
                    $3 = append($3[:len($3)-1], $4...)
                }

                // save comments
                yylex.(*Parser).setFreeFloating($3[0], freefloating.Var, $2.FreeFloating)

                for _, n := range($3) {
                    switch nn := n.(type) {
                        case *expr.ArrayDimFetch:
                            nn.Variable = $$
                            nn.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Variable, $$)
                        
                        case *expr.PropertyFetch:
                            nn.Variable = $$
                            nn.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Variable, $$)
                        
                        case *expr.MethodCall:
                            nn.Variable = $$
                            nn.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Variable, $$)
                    }
                }

                for _, n := range($5) {
                    switch nn := n.(type) {
                        case *expr.ArrayDimFetch:
                            nn.Variable = $$
                            nn.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Variable, $$)
                        
                        case *expr.PropertyFetch:
                            nn.Variable = $$
                            nn.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Variable, $$)
                        
                        case *expr.MethodCall:
                            nn.Variable = $$
                            nn.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($$, nn))
                            $$ = nn
                            yylex.(*Parser).MoveFreeFloating(nn.Variable, $$)
                    }
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   base_variable_with_function_calls
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

variable_properties:
        variable_properties variable_property
            {
                $$ = append($1, $2...) 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                $$ = []node.Node{} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


variable_property:
        T_OBJECT_OPERATOR object_property method_or_not
            {
                if $3 != nil {
                    $3[0].(*expr.MethodCall).Method = $2[len($2)-1].(*expr.PropertyFetch).Property
                    $2 = append($2[:len($2)-1], $3...)
                }

                $$ = $2

                // save comments
                yylex.(*Parser).setFreeFloating($2[0], freefloating.Var, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

array_method_dereference:
        array_method_dereference '[' dim_offset ']'
            {
                fetch := expr.NewArrayDimFetch(nil, $3)
                $$ = append($1, fetch)

                // save position
                fetch.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($3))

                // save comments
                yylex.(*Parser).setFreeFloating(fetch, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating(fetch, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   method '[' dim_offset ']'
            {
                fetch := expr.NewArrayDimFetch(nil, $3)
                $$ = []node.Node{$1, fetch}

                // save position
                fetch.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($3))

                // save comments
                yylex.(*Parser).setFreeFloating(fetch, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating(fetch, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

method:
        function_call_parameter_list
            {
                $$ = expr.NewMethodCall(nil, nil, $1.(*node.ArgumentList))

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

method_or_not:
        method
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   array_method_dereference
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

variable_without_objects:
        reference_variable
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   simple_indirect_reference reference_variable
            {
                $1.last.SetVarName($2)

                for _, n := range($1.all) {
                    n.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition(n, $2))
                }

                $$ = $1.all[0]

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_member:
        class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects
            {
                $$ = expr.NewStaticPropertyFetch($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects
            {
                $$ = expr.NewStaticPropertyFetch($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

variable_class_name:
        reference_variable
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

array_function_dereference:
        array_function_dereference '[' dim_offset ']'
            {
                $$ = expr.NewArrayDimFetch($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   function_call '[' dim_offset ']'
            {
                $$ = expr.NewArrayDimFetch($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

base_variable_with_function_calls:
        base_variable
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   array_function_dereference
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   function_call
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


base_variable:
        reference_variable
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   simple_indirect_reference reference_variable
            {
                $1.last.SetVarName($2)

                for _, n := range($1.all) {
                    n.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition(n, $2))
                }

                $$ = $1.all[0]

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_member
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

reference_variable:
        reference_variable '[' dim_offset ']'
            {
                $$ = expr.NewArrayDimFetch($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   reference_variable '{' expr '}'
            {
                $$ = expr.NewArrayDimFetch($1, $3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   compound_variable
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


compound_variable:
        T_VARIABLE
            {
                name := node.NewIdentifier(strings.TrimLeftFunc($1.Value, isDollar))
                $$ = expr.NewVariable(name)
                
                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).addDollarToken($$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' '{' expr '}'
            {
                $$ = expr.NewVariable($3)
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Dollar, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).setFreeFloating($3, freefloating.Start, append($2.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($2), (*$3.GetFreeFloating())[freefloating.Start]...)...))
                yylex.(*Parser).setFreeFloating($3, freefloating.End, append((*$3.GetFreeFloating())[freefloating.End], append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

dim_offset:
        /* empty */
            {
                $$ = nil 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


object_property:
        object_dim_list 
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_without_objects
            {
                fetch := expr.NewPropertyFetch(nil, $1)
                $$ = []node.Node{fetch}

                // save position
                fetch.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

object_dim_list:
        object_dim_list '[' dim_offset ']'
            {
                fetch := expr.NewArrayDimFetch(nil, $3)
                $$ = append($1, fetch)
                
                // save position
                fetch.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($3))

                // save comments
                yylex.(*Parser).setFreeFloating(fetch, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating(fetch, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   object_dim_list '{' expr '}'
            {
                fetch := expr.NewArrayDimFetch(nil, $3)
                $$ = append($1, fetch)
                
                // save position
                fetch.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($3))

                // save comments
                yylex.(*Parser).setFreeFloating(fetch, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating(fetch, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_name
            {
                fetch := expr.NewPropertyFetch(nil, $1)
                $$ = []node.Node{fetch}
                
                // save position
                fetch.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

variable_name:
        T_STRING
            {
                $$ = node.NewIdentifier($1.Value)
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' expr '}'
            {
                $$ = $2
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, append($1.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($1), (*$$.GetFreeFloating())[freefloating.Start]...)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.End, append((*$$.GetFreeFloating())[freefloating.End], append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

simple_indirect_reference:
        '$'
            {
                n := expr.NewVariable(nil)
                $$ = simpleIndirectReference{[]*expr.Variable{n}, n}
                
                // save position
                n.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                
                // save comments
                yylex.(*Parser).setFreeFloating(n, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(n, freefloating.Dollar, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   simple_indirect_reference '$'
            {
                n := expr.NewVariable(nil)

                $1.last.SetVarName(n)
                $1.all = append($1.all, n)
                $1.last = n
                $$ = $1
                
                // save position
                n.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                
                // save comments
                yylex.(*Parser).setFreeFloating(n, freefloating.Start, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(n, freefloating.Dollar, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

assignment_list:
        assignment_list ',' assignment_list_element
            {
                if len($1) == 0 {
                    $1 = []node.Node{expr.NewArrayItem(nil, nil, false)}
                }

                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   assignment_list_element
            {
                if $1.(*expr.ArrayItem).Key == nil && $1.(*expr.ArrayItem).Val == nil {
                    $$ = []node.Node{}
                } else {
                    $$ = []node.Node{$1}
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


assignment_list_element:
        variable
            {
                $$ = expr.NewArrayItem(nil, $1, false)
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($1))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LIST '(' assignment_list ')'
            {
                listNode := expr.NewList($3)
                $$ = expr.NewArrayItem(nil, listNode, false)
                
                // save position
                listNode.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition(listNode))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating(listNode, freefloating.List, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(listNode, freefloating.ArrayPairList, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                $$ = expr.NewArrayItem(nil, nil, false) 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


array_pair_list:
        /* empty */
            {
                $$ = []node.Node{} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_array_pair_list possible_comma
            {
                $$ = $1

                if $2 != nil {
                    $$ = append($1, expr.NewArrayItem(nil, nil, false))
                }

                // save comments
                if $2 != nil {
                    yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_array_pair_list:
        non_empty_array_pair_list ',' expr T_DOUBLE_ARROW expr
            {
                arrayItem := expr.NewArrayItem($3, $5, false)
                $$ = append($1, arrayItem)

                // save position
                arrayItem.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($3, $5))
                
                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, freefloating.Expr, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_array_pair_list ',' expr
            {
                arrayItem := expr.NewArrayItem(nil, $3, false)
                $$ = append($1, arrayItem)
                
                // save position
                arrayItem.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($3))
                
                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_DOUBLE_ARROW expr
            {
                arrayItem := expr.NewArrayItem($1, $3, false)
                $$ = []node.Node{arrayItem}

                // save position
                arrayItem.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3))
                
                // save comments
                yylex.(*Parser).MoveFreeFloating($1, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                arrayItem := expr.NewArrayItem(nil, $1, false)
                $$ = []node.Node{arrayItem}

                // save position
                arrayItem.SetPosition(yylex.(*Parser).positionBuilder.NewNodePosition($1))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, arrayItem)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_array_pair_list ',' expr T_DOUBLE_ARROW '&' w_variable
            {
                reference := expr.NewReference($6)
                arrayItem := expr.NewArrayItem($3, reference, false)
                $$ = append($1, arrayItem)
                
                // save position
                reference.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($5, $6))
                arrayItem.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($3, $6))
                
                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                yylex.(*Parser).MoveFreeFloating($3, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, freefloating.Expr, $4.FreeFloating)
                yylex.(*Parser).setFreeFloating(reference, freefloating.Start, $5.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_array_pair_list ',' '&' w_variable
            {
                reference := expr.NewReference($4)
                arrayItem := expr.NewArrayItem(nil, reference, false)
                $$ = append($1, arrayItem)
                
                // save position
                reference.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $4))
                arrayItem.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $4))
                
                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(arrayItem, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_DOUBLE_ARROW '&' w_variable
            {
                reference := expr.NewReference($4)
                arrayItem := expr.NewArrayItem($1, reference, false)
                $$ = []node.Node{arrayItem}
                
                // save position
                reference.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $4))
                arrayItem.SetPosition(yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, arrayItem)
                yylex.(*Parser).setFreeFloating(arrayItem, freefloating.Expr, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(reference, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' w_variable
            {
                reference := expr.NewReference($2)
                arrayItem := expr.NewArrayItem(nil, reference, false)
                $$ = []node.Node{arrayItem}
                
                // save position
                reference.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                arrayItem.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))
                
                // save comments
                yylex.(*Parser).setFreeFloating(arrayItem, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_list:
        encaps_list encaps_var
            {
                $$ = append($1, $2) 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   encaps_list T_ENCAPSED_AND_WHITESPACE
            {
                encapsed := scalar.NewEncapsedStringPart($2.Value)
                $$ = append($1, encapsed)

                // save position
                encapsed.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))

                // save comments
                yylex.(*Parser).setFreeFloating(encapsed, freefloating.Start, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   encaps_var
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ENCAPSED_AND_WHITESPACE encaps_var
            {
                encapsed := scalar.NewEncapsedStringPart($1.Value)
                $$ = []node.Node{encapsed, $2}

                // save position
                encapsed.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating(encapsed, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_var:
        T_VARIABLE
            {
                name := node.NewIdentifier(strings.TrimLeftFunc($1.Value, isDollar))
                $$ = expr.NewVariable(name)

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).addDollarToken($$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '[' encaps_var_offset ']'
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($1.Value, isDollar))
                variable := expr.NewVariable(identifier)
                $$ = expr.NewArrayDimFetch(variable, $3)

                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).addDollarToken(variable)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE T_OBJECT_OPERATOR T_STRING
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($1.Value, isDollar))
                variable := expr.NewVariable(identifier)
                fetch := node.NewIdentifier($3.Value)
                $$ = expr.NewPropertyFetch(variable, fetch)

                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                fetch.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).addDollarToken(variable)
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(fetch, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES expr '}'
            {
                variable := expr.NewVariable($2)

                $$ = variable

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).setFreeFloating($$, freefloating.End, append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}'
            {
                name := node.NewIdentifier($2.Value)
                variable := expr.NewVariable(name)

                $$ = variable

                // save position
                name.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).setFreeFloating($$, freefloating.End, append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
            {
                identifier := node.NewIdentifier($2.Value)
                variable := expr.NewVariable(identifier)
                $$ = expr.NewArrayDimFetch(variable, $4)

                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                variable.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($2))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $6))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($5.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($5)...))
                yylex.(*Parser).setFreeFloating($$, freefloating.End, append($6.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($6)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CURLY_OPEN variable '}'
            {
                $$ = $2;

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).setFreeFloating($$, freefloating.End, append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_var_offset:
        T_STRING
            {
                $$ = scalar.NewString($1.Value)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NUM_STRING
            {
                // TODO: add option to handle 64 bit integer
                if _, err := strconv.Atoi($1.Value); err == nil {
                    $$ = scalar.NewLnumber($1.Value)
                } else {
                    $$ = scalar.NewString($1.Value)
                }

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                identifier := node.NewIdentifier(strings.TrimLeftFunc($1.Value, isDollar))
                $$ = expr.NewVariable(identifier)

                // save position
                identifier.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($1))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).addDollarToken($$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

internal_functions_in_yacc:
        T_ISSET '(' isset_variables ')'
            {
                $$ = expr.NewIsset($3)
                
                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Isset, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.VarList, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EMPTY '(' variable ')'
            {
                $$ = expr.NewEmpty($3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Empty, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EMPTY '(' expr ')'
            {
                $$ = expr.NewEmpty($3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Empty, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INCLUDE expr
            {
                $$ = expr.NewInclude($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INCLUDE_ONCE expr
            {
                $$ = expr.NewIncludeOnce($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EVAL '(' expr ')'
            {
                $$ = expr.NewEval($3)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Eval, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_REQUIRE expr
            {
                $$ = expr.NewRequire($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_REQUIRE_ONCE expr
            {
                $$ = expr.NewRequireOnce($2)

                // save position
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2))

                // save comments
                yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

isset_variables:
        isset_variable
            {
                $$ = []node.Node{$1} 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   isset_variables ',' isset_variable
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

isset_variable:
        variable
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr_without_variable
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_constant:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                target := node.NewIdentifier($3.Value)
                $$ = expr.NewClassConstFetch($1, target)

                // save position
                target.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(target, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                target := node.NewIdentifier($3.Value)
                $$ = expr.NewClassConstFetch($1, target)

                // save position
                target.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(target, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_class_name_scalar:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_CLASS
            {
                target := node.NewIdentifier($3.Value)
                $$ = expr.NewClassConstFetch($1, target)

                // save position
                target.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(target, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_name_scalar:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_CLASS
            {
                target := node.NewIdentifier($3.Value)
                $$ = expr.NewClassConstFetch($1, target)

                // save position
                target.SetPosition(yylex.(*Parser).positionBuilder.NewTokenPosition($3))
                $$.SetPosition(yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3))

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)
                yylex.(*Parser).setFreeFloating(target, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

%%

type simpleIndirectReference struct {
	all  []*expr.Variable
	last *expr.Variable
}
