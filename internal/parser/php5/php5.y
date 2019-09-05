%{
package php5

import (
    "bytes"
    "strconv"

    "github.com/z7zmey/php-parser/pkg/ast"
    "github.com/z7zmey/php-parser/internal/graph"
    "github.com/z7zmey/php-parser/internal/stxtree"
    "github.com/z7zmey/php-parser/internal/scanner"
)

%}

%union{
    node  graph.NodeID
    token *scanner.Token
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
%type <node> namespace_name

%type <node> extends_from
%type <node> implements_list
%type <node> interface_extends_list
%type <node> lexical_vars

%type <token> is_reference is_variadic

%%

start:
        top_statement_list
            {
                children := yylex.(*Parser).List.Pop()
                nodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeRoot,
                })
                yylex.(*Parser).SavePosition(nodeID, yylex.(*Parser).NewPosition(children, nil, nil))
                yylex.(*Parser).Children(nodeID, ast.NodeGroupStmts, children...)

                yylex.(*Parser).Ast.RootNode = nodeID

                // save tokens
                yylex.(*Parser).AppendTokens(nodeID, ast.TokenGroupEnd, yylex.(*Parser).CurrentToken.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

top_statement_list:
        top_statement_list top_statement
            {
                if $2 != 0 {
                    yylex.(*Parser).List.Add($2)
                }

                // TODO
                // if inlineHtmlNode, ok := $2.(*stmt.InlineHtml); ok && len($1) > 0 {
                //     prevNode := lastNode($1)
                //     yylex.(*Parser).splitSemiColonAndPhpCloseTag(inlineHtmlNode, prevNode)
                // }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

namespace_name:
        T_STRING
            {
                nodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameNamePart,
                })
                yylex.(*Parser).SavePosition(nodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(nodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(nodeID, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_NS_SEPARATOR T_STRING
            {
                prevNodeID := yylex.(*Parser).List.Last()
                nodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameNamePart,
                })
                yylex.(*Parser).SavePosition(nodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))
                yylex.(*Parser).List.Add(nodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(nodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

top_statement:
        error
            {
                // error
                $$ = 0

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
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtHaltCompiler,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupHaltCompiller, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupOpenParenthesisToken, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCloseParenthesisToken, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE namespace_name ';'
            {
                children := yylex.(*Parser).List.Pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(children, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, children...)

                // Create Namespace Node
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtNamespace,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupNamespaceName, nameNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens(children[0], nameNodeID)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE namespace_name '{' top_statement_list '}'
            {
                childrenStmts := yylex.(*Parser).List.Pop()
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // Create Namespace Node
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtNamespace,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $5}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupNamespaceName, nameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, childrenStmts...)

                // save tokens
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $5.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE '{' top_statement_list '}'
            {
                children := yylex.(*Parser).List.Pop()
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtNamespace,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupNamespace, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE use_declarations ';'
            {
                children := yylex.(*Parser).List.Pop()
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUseList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUses, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupUseDeclarationList, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE T_FUNCTION use_function_declarations ';'
            {
                children := yylex.(*Parser).List.Pop()

                // use type
                useTypeNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(useTypeNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))

                // use list
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUseList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUseType, useTypeNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupUses, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupUseDeclarationList, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE T_CONST use_const_declarations ';'
            {
                children := yylex.(*Parser).List.Pop()

                // use type
                useTypeNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(useTypeNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))

                // use list
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUseList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUseType, useTypeNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupUses, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupUseDeclarationList, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   constant_declaration ';'
            {
                $$ = $1

                // save position
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$2}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$2})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_declarations:
        use_declarations ',' use_declaration
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   use_declaration
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_declaration:
        namespace_name
            {
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // Create Use Node
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUse, nameNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_AS T_STRING
            {
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // create Alias Node
                
                aliasNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(aliasNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                // Create Use Node
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenNameParts, []*scanner.Token{$3}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUse, nameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupAlias, aliasNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(aliasNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // Create Use Node
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUse, nameNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSlash, []scanner.Token{*$1})
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // create Alias Node
                
                aliasNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(aliasNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))

                // Create Use Node
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenNameParts, []*scanner.Token{$4}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUse, nameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupAlias, aliasNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSlash, []scanner.Token{*$1})
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(aliasNodeID, ast.TokenGroupStart, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_function_declarations:
        use_function_declarations ',' use_function_declaration
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   use_function_declaration
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_function_declaration:
        namespace_name
            {
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // Create Use Node
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUse, nameNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_AS T_STRING
            {
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // create Alias Node
                
                aliasNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(aliasNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                // Create Use Node
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenNameParts, []*scanner.Token{$3}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUse, nameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupAlias, aliasNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(aliasNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // Create Use Node
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUse, nameNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSlash, []scanner.Token{*$1})
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // create Alias Node
                
                aliasNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(aliasNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))

                // Create Use Node
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenNameParts, []*scanner.Token{$4}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUse, nameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupAlias, aliasNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSlash, []scanner.Token{*$1})
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(aliasNodeID, ast.TokenGroupStart, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_const_declarations:
        use_const_declarations ',' use_const_declaration
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   use_const_declaration
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_const_declaration:
        namespace_name
            {
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // Create Use Node
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUse, nameNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_AS T_STRING
            {
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // create Alias Node
                
                aliasNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(aliasNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                // Create Use Node
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenNameParts, []*scanner.Token{$3}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUse, nameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupAlias, aliasNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(aliasNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // Create Use Node
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUse, nameNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSlash, []scanner.Token{*$1})
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_AS T_STRING
            {
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // create Alias Node
                
                aliasNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(aliasNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))

                // Create Use Node
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenNameParts, []*scanner.Token{$4}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUse, nameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupAlias, aliasNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSlash, []scanner.Token{*$1})
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(aliasNodeID, ast.TokenGroupStart, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

constant_declaration:
        constant_declaration ',' T_STRING '=' static_scalar
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))
                
                constantNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtConstant,
                })
                yylex.(*Parser).SavePosition(constantNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, []graph.NodeID{$5}))

                yylex.(*Parser).Children(constantNodeID, ast.NodeGroupConstantName, identifierNodeID)
                yylex.(*Parser).Children(constantNodeID, ast.NodeGroupExpr, $5)

                $$ = $1

                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $5}, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupConsts, constantNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(constantNodeID, ast.TokenGroupStart, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(constantNodeID, ast.TokenGroupName, $4.HiddenTokens)
                // yylex.(*Parser).AppendTokens($1, ast.TokenGroupEnd, $2.HiddenTokens) // TODO: attach to $1 last const

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONST T_STRING '=' static_scalar
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))
                
                constantNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtConstant,
                })
                yylex.(*Parser).SavePosition(constantNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, []graph.NodeID{$4}))

                yylex.(*Parser).Children(constantNodeID, ast.NodeGroupConstantName, identifierNodeID)
                yylex.(*Parser).Children(constantNodeID, ast.NodeGroupExpr, $4)

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtConstList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$4}))
                yylex.(*Parser).Children($$, ast.NodeGroupConsts, constantNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(constantNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(constantNodeID, ast.TokenGroupName, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

inner_statement_list:
        inner_statement_list inner_statement
            {
                // TODO
                // if inlineHtmlNode, ok := $2.(*stmt.InlineHtml); ok && len($1) > 0 {
                //     prevNode := lastNode($1)
                //     yylex.(*Parser).splitSemiColonAndPhpCloseTag(inlineHtmlNode, prevNode)
                // }

                if $2 != 0 {
                    yylex.(*Parser).List.Add($2)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


inner_statement:
        error
            {
                // error
                $$ = 0

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
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtHaltCompiler,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupHaltCompiller, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupOpenParenthesisToken, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCloseParenthesisToken, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$4})

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
                LableNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(LableNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtLabel,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLabelName, LableNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupLabel, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unticked_statement:
        '{' inner_statement_list '}'
            {
                children := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStmtList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))
                
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_IF parenthesis_expr statement elseif_list else_single
            {
                children := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtIf,
                })

                if $5 != 0 {
                    yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$5}))
                } else if len(children) > 0 {
                    yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))
                } else {
                    yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$3}))
                }

                yylex.(*Parser).Children($$, ast.NodeGroupCond, $2)
                yylex.(*Parser).Children($$, ast.NodeGroupStmt, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupElseIf, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                // TODO handle parenthesis_expr tokens

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_IF parenthesis_expr ':' inner_statement_list new_elseif_list new_else_single T_ENDIF ';'
            {
                childrenElseIf := yylex.(*Parser).List.Pop()
                childrenStmt := yylex.(*Parser).List.Pop()

                stmtListNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStmtList,
                })
                yylex.(*Parser).SavePosition(stmtListNodeID, yylex.(*Parser).NewPosition(childrenStmt, nil, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtAltIf,
                    Flag: ast.NodeFlagAltSyntax,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $8}, nil))

                yylex.(*Parser).Children(stmtListNodeID, ast.NodeGroupStmts, childrenStmt...)
                yylex.(*Parser).Children($$, ast.NodeGroupCond, $2)
                yylex.(*Parser).Children($$, ast.NodeGroupStmt, stmtListNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupElseIf, childrenElseIf...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $3.HiddenTokens)
                // TODO handle parenthesis_expr tokens

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_WHILE parenthesis_expr while_statement
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtWhile,
                })

                yylex.(*Parser).Children($$, ast.NodeGroupCond, $2)
                yylex.(*Parser).Children($$, ast.NodeGroupStmt, $3)
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$3}))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                // TODO handle parenthesis_expr tokens

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DO statement T_WHILE parenthesis_expr ';'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtDo,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $5}, nil))
                
                yylex.(*Parser).Children($$, ast.NodeGroupStmt, $2)
                yylex.(*Parser).Children($$, ast.NodeGroupCond, $4)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $5.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$5})
                // TODO handle parenthesis_expr tokens

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOR '(' for_expr ';' for_expr ';' for_expr ')' for_statement
            {
                yylex.(*Parser).Children($9, ast.NodeGroupLoop, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children($9, ast.NodeGroupCond, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children($9, ast.NodeGroupInit, yylex.(*Parser).List.Pop()...)

                $$ = $9
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$9}))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupFor, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupInitExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCondExpr, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupIncExpr, $8.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_SWITCH parenthesis_expr switch_case_list
            {
                $$ = $3

                yylex.(*Parser).Children($$, ast.NodeGroupCond, $2)
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$3}))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                // TODO handle parenthesis_expr tokens

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BREAK ';'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtBreak,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$2})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BREAK expr ';'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtBreak,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))
                
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONTINUE ';'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtContinue,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$2})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONTINUE expr ';'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtContinue,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))
                
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_RETURN ';'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtReturn,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$2})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_RETURN expr_without_variable ';'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtReturn,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))
                
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_RETURN variable ';'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtReturn,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))
                
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   yield_expr ';'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtExpression,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$2}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $1)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$2})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_GLOBAL global_var_list ';'
            {
                children := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtGlobal,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))
                
                yylex.(*Parser).Children($$, ast.NodeGroupVars, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVarList, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC static_var_list ';'
            {
                children := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStatic,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))
                
                yylex.(*Parser).Children($$, ast.NodeGroupVars, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVarList, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ECHO echo_expr_list ';'
            {
                children := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtEcho,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))
                
                yylex.(*Parser).Children($$, ast.NodeGroupExprs, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEcho, []scanner.Token{*$1})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INLINE_HTML
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtInlineHtml,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr ';'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtExpression,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$2}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $1)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$2})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_UNSET '(' unset_variables ')' ';'
            {
                children := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUnset,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $5}, nil))
                
                yylex.(*Parser).Children($$, ast.NodeGroupVars, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupUnset, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVarList, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCloseParenthesisToken, $5.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$5})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOREACH '(' variable T_AS foreach_variable foreach_optional_arg ')' foreach_statement
            {
                yylex.(*Parser).Children($8, ast.NodeGroupExpr, $3)
                if $6 == 0 {
                    yylex.(*Parser).Children($8, ast.NodeGroupVar, $5)
                } else {
                    yylex.(*Parser).Children($8, ast.NodeGroupKey, $5)
                    yylex.(*Parser).Children($8, ast.NodeGroupVar, $6)
                }

                $$ = $8
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$8}))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupForeach, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $7.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOREACH '(' expr_without_variable T_AS foreach_variable foreach_optional_arg ')' foreach_statement
            {
                yylex.(*Parser).Children($8, ast.NodeGroupExpr, $3)
                if $6 == 0 {
                    yylex.(*Parser).Children($8, ast.NodeGroupVar, $5)
                } else {
                    yylex.(*Parser).Children($8, ast.NodeGroupKey, $5)
                    yylex.(*Parser).Children($8, ast.NodeGroupVar, $6)
                }

                $$ = $8
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$8}))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupForeach, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $7.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DECLARE '(' declare_list ')' declare_statement
            {
                children := yylex.(*Parser).List.Pop()
                yylex.(*Parser).Children($5, ast.NodeGroupConsts, children...)

                $$ = $5
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$5}))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupDeclare, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupConstList, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ';'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtNop,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$1})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRY '{' inner_statement_list '}' catch_statement finally_statement
            {
                childrenCatches := yylex.(*Parser).List.Pop()
                childrenStmts := yylex.(*Parser).List.Pop()

                var posID graph.NodeID
                if $6 == 0 {
                    posID = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, childrenCatches)
                } else {
                    posID = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$6})
                }

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtTry,
                })
                yylex.(*Parser).SavePosition($$, posID)

                yylex.(*Parser).Children($$, ast.NodeGroupConsts, $6)
                yylex.(*Parser).Children($$, ast.NodeGroupCatches, childrenCatches...)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, childrenStmts...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupTry, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_THROW expr ';'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtThrow,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_GOTO T_STRING ';'
            {
                LableNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(LableNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtGoto,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLabel, LableNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(LableNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupLabel, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

catch_statement:
        /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CATCH '(' fully_qualified_class_name T_VARIABLE ')' '{' inner_statement_list '}' additional_catches
            {
                additionalCatches := yylex.(*Parser).List.Pop()

                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))
                
                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))
                
                catchNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtCatch,
                })
                yylex.(*Parser).SavePosition(catchNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $8}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children(catchNodeID, ast.NodeGroupVar, varNodeID)
                yylex.(*Parser).Children(catchNodeID, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children(catchNodeID, ast.NodeGroupTypes, $3)

                yylex.(*Parser).List.Add(catchNodeID)
                for _, v := range(additionalCatches) {
                    yylex.(*Parser).List.Add(v)
                }

                // save tokens
                yylex.(*Parser).AppendTokens(varNodeID, ast.TokenGroupStart, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens(catchNodeID, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(catchNodeID, ast.TokenGroupCatch, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(catchNodeID, ast.TokenGroupVar, $5.HiddenTokens)
                yylex.(*Parser).AppendTokens(catchNodeID, ast.TokenGroupCond, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens(catchNodeID, ast.TokenGroupStmts, $8.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

finally_statement:
        /* empty */
            {
                $$ = 0

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINALLY '{' inner_statement_list '}'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtFinally,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupFinally, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

additional_catches:
        non_empty_additional_catches
            {
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_additional_catches:
        additional_catch
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_additional_catches additional_catch
            {
                yylex.(*Parser).List.Add($2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

additional_catch:
        T_CATCH '(' fully_qualified_class_name T_VARIABLE ')' '{' inner_statement_list '}'
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))
                
                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtCatch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $8}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupVar, varNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children($$, ast.NodeGroupTypes, $3)

                // save tokens
                yylex.(*Parser).AppendTokens(varNodeID, ast.TokenGroupStart, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCatch, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $5.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $8.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unset_variables:
        unset_variable
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   unset_variables ',' unset_variable
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

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
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                var flag ast.NodeFlag
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                }
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtFunction,
                    Flag: flag,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $9}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupFunctionName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children($$, ast.NodeGroupParams, yylex.(*Parser).List.Pop()...)


                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                if $2 != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupFunction, $2.HiddenTokens)
                } 
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupParamList, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupParams, $7.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $9.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unticked_class_declaration_statement:
        class_entry_type T_STRING extends_from implements_list '{' class_statement_list '}'
            {
                children := yylex.(*Parser).List.Pop()

                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))

                $$ = $1

                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$7}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClassName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupExtends, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupImplements, $4)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, children...)

                // save tokens
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $5.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $7.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   interface_entry T_STRING interface_extends_list '{' class_statement_list '}'
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtInterface,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $6}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExtends, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupInterfaceName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $6.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


class_entry_type:
        T_CLASS
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtClass,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ABSTRACT T_CLASS
            {
                modifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(modifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtClass,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{modifierNodeID}, []*scanner.Token{$1, $2}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupModifiers, modifierNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupModifierList, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRAIT
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtTrait,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINAL T_CLASS
            {
                modifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(modifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtClass,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{modifierNodeID}, []*scanner.Token{$1, $2}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupModifiers, modifierNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupModifierList, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

extends_from:
        /* empty */
            {
                $$ = 0

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXTENDS fully_qualified_class_name
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtClassExtends,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupClassName, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

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
                $$ = 0

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXTENDS interface_list
            {
                children := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtInterfaceExtends,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))

                yylex.(*Parser).Children($$, ast.NodeGroupInterfaceNames, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

implements_list:
        /* empty */
            {
                $$ = 0

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_IMPLEMENTS interface_list
            {
                children := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtClassImplements,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))

                yylex.(*Parser).Children($$, ast.NodeGroupInterfaceNames, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

interface_list:
        fully_qualified_class_name
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   interface_list ',' fully_qualified_class_name
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

foreach_optional_arg:
        /* empty */
            {
                $$ = 0 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOUBLE_ARROW foreach_variable
            {
                $$ = $2

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupKey, $1.HiddenTokens)

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
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprReference,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LIST '(' assignment_list ')'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupItems, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupList, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupArrayPairList, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

for_statement:
        statement
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtFor,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition( []graph.NodeID{$1}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupStmt, $1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDFOR ';'
            {
                children := yylex.(*Parser).List.Pop()

                stmtListNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStmtList,
                })
                yylex.(*Parser).SavePosition(stmtListNodeID, yylex.(*Parser).NewPosition(children, nil, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtAltFor,
                    Flag: ast.NodeFlagAltSyntax,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children(stmtListNodeID, ast.NodeGroupStmts, children...)
                yylex.(*Parser).Children($$, ast.NodeGroupStmt, stmtListNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupAltEnd, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

foreach_statement:
        statement
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtForeach,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition( []graph.NodeID{$1}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupStmt, $1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDFOREACH ';'
            {
                children := yylex.(*Parser).List.Pop()

                stmtListNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStmtList,
                })
                yylex.(*Parser).SavePosition(stmtListNodeID, yylex.(*Parser).NewPosition(children, nil, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtAltForeach,
                    Flag: ast.NodeFlagAltSyntax,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children(stmtListNodeID, ast.NodeGroupStmts, children...)
                yylex.(*Parser).Children($$, ast.NodeGroupStmt, stmtListNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupAltEnd, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


declare_statement:
        statement
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtDeclare,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition( []graph.NodeID{$1}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupStmt, $1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDDECLARE ';'
            {
                children := yylex.(*Parser).List.Pop()

                stmtListNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStmtList,
                })
                yylex.(*Parser).SavePosition(stmtListNodeID, yylex.(*Parser).NewPosition(children, nil, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtDeclare,
                    Flag: ast.NodeFlagAltSyntax,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children(stmtListNodeID, ast.NodeGroupStmts, children...)
                yylex.(*Parser).Children($$, ast.NodeGroupStmt, stmtListNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupAltEnd, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


declare_list:
        T_STRING '=' static_scalar
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))
                
                constNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtConstant,
                })
                yylex.(*Parser).SavePosition(constNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$3}))

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(constNodeID)

                yylex.(*Parser).Children(constNodeID, ast.NodeGroupConstantName, identifierNodeID)
                yylex.(*Parser).Children(constNodeID, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).AppendTokens(constNodeID, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(constNodeID, ast.TokenGroupName, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   declare_list ',' T_STRING '=' static_scalar
            {
                prevNodeID := yylex.(*Parser).List.Last()
                
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))
                
                constNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtConstant,
                })
                yylex.(*Parser).SavePosition(constNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, []graph.NodeID{$5}))
                
                yylex.(*Parser).List.Add(constNodeID)

                yylex.(*Parser).Children(constNodeID, ast.NodeGroupConstantName, identifierNodeID)
                yylex.(*Parser).Children(constNodeID, ast.NodeGroupExpr, $5)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(constNodeID, ast.TokenGroupStart, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(constNodeID, ast.TokenGroupName, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


switch_case_list:
        '{' case_list '}'
            {
                children := yylex.(*Parser).List.Pop()

                caseListNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtCaseList,
                })
                yylex.(*Parser).SavePosition(caseListNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtSwitch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children(caseListNodeID, ast.NodeGroupCases, children...)
                yylex.(*Parser).Children($$, ast.NodeGroupCaseList, caseListNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(caseListNodeID, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(caseListNodeID, ast.TokenGroupCaseListEnd, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' ';' case_list '}'
            {
                children := yylex.(*Parser).List.Pop()

                caseListNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtCaseList,
                })
                yylex.(*Parser).SavePosition(caseListNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtSwitch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children(caseListNodeID, ast.NodeGroupCases, children...)
                yylex.(*Parser).Children($$, ast.NodeGroupCaseList, caseListNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(caseListNodeID, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(caseListNodeID, ast.TokenGroupCaseListStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(caseListNodeID, ast.TokenGroupCaseListStart, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens(caseListNodeID, ast.TokenGroupCaseListEnd, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' case_list T_ENDSWITCH ';'
            {
                children := yylex.(*Parser).List.Pop()

                caseListNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtCaseList,
                })
                yylex.(*Parser).SavePosition(caseListNodeID, yylex.(*Parser).NewPosition(children, nil, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtAltSwitch,
                    Flag: ast.NodeFlagAltSyntax,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children(caseListNodeID, ast.NodeGroupCases, children...)
                yylex.(*Parser).Children($$, ast.NodeGroupCaseList, caseListNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(caseListNodeID, ast.TokenGroupCaseListEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupAltEnd, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' ';' case_list T_ENDSWITCH ';'
            {
                children := yylex.(*Parser).List.Pop()

                caseListNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtCaseList,
                })
                yylex.(*Parser).SavePosition(caseListNodeID, yylex.(*Parser).NewPosition(children, nil, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtAltSwitch,
                    Flag: ast.NodeFlagAltSyntax,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $5}, nil))

                yylex.(*Parser).Children(caseListNodeID, ast.NodeGroupCases, children...)
                yylex.(*Parser).Children($$, ast.NodeGroupCaseList, caseListNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(caseListNodeID, ast.TokenGroupCaseListStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(caseListNodeID, ast.TokenGroupCaseListStart, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens(caseListNodeID, ast.TokenGroupCaseListEnd, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupAltEnd, $5.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$5})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


case_list:
        /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   case_list T_CASE expr case_separator inner_statement_list
            {
                children := yylex.(*Parser).List.Pop()

                caseNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtCase,
                })
                yylex.(*Parser).SavePosition(caseNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, children))

                yylex.(*Parser).Children(caseNodeID, ast.NodeGroupStmts, children...)

                yylex.(*Parser).List.Add(caseNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(caseNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(caseNodeID, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens(caseNodeID, ast.TokenGroupCaseSeparator, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   case_list T_DEFAULT case_separator inner_statement_list
            {
                children := yylex.(*Parser).List.Pop()

                defaultNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtDefault,
                })
                yylex.(*Parser).SavePosition(defaultNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, children))

                yylex.(*Parser).Children(defaultNodeID, ast.NodeGroupStmts, children...)

                yylex.(*Parser).List.Add(defaultNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(defaultNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(defaultNodeID, ast.TokenGroupDefault, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(defaultNodeID, ast.TokenGroupCaseSeparator, []scanner.Token{*$3})

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
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtWhile,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition( []graph.NodeID{$1}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupStmt, $1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDWHILE ';'
            {
                children := yylex.(*Parser).List.Pop()

                stmtListNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStmtList,
                })
                yylex.(*Parser).SavePosition(stmtListNodeID, yylex.(*Parser).NewPosition(children, nil, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtAltWhile,
                    Flag: ast.NodeFlagAltSyntax,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children(stmtListNodeID, ast.NodeGroupStmts, children...)
                yylex.(*Parser).Children($$, ast.NodeGroupStmt, stmtListNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupAltEnd, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

elseif_list:
        /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   elseif_list T_ELSEIF parenthesis_expr statement
            {
                elseIfNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtElseIf,
                })
                yylex.(*Parser).SavePosition(elseIfNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, []graph.NodeID{$4}))

                yylex.(*Parser).Children(elseIfNodeID, ast.NodeGroupCond, $3)
                yylex.(*Parser).Children(elseIfNodeID, ast.NodeGroupStmt, $4)

                // save tokens
                yylex.(*Parser).AppendTokens(elseIfNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                // TODO handle parenthesis_expr tokens

                yylex.(*Parser).List.Add(elseIfNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


new_elseif_list:
        /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_elseif_list T_ELSEIF parenthesis_expr ':' inner_statement_list
            {
                children := yylex.(*Parser).List.Pop()

                stmtListNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStmtList,
                })
                yylex.(*Parser).SavePosition(stmtListNodeID, yylex.(*Parser).NewPosition(children, nil, nil))

                AltElseIfNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtAltElseIf,
                })
                yylex.(*Parser).SavePosition(AltElseIfNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, children))

                yylex.(*Parser).Children(stmtListNodeID, ast.NodeGroupStmts, children...)
                yylex.(*Parser).Children(AltElseIfNodeID, ast.NodeGroupCond, $3)
                yylex.(*Parser).Children(AltElseIfNodeID, ast.NodeGroupStmt, stmtListNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(AltElseIfNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(AltElseIfNodeID, ast.TokenGroupCond, $4.HiddenTokens)
                // TODO handle parenthesis_expr tokens

                yylex.(*Parser).List.Add(AltElseIfNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


else_single:
        /* empty */
            {
                $$ = 0 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ELSE statement
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtElse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupStmt, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


new_else_single:
        /* empty */
            {
                $$ = 0 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ELSE ':' inner_statement_list
            {
                children := yylex.(*Parser).List.Pop()

                stmtListNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStmtList,
                })
                yylex.(*Parser).SavePosition(stmtListNodeID, yylex.(*Parser).NewPosition(children, nil, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtAltElse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))

                yylex.(*Parser).Children(stmtListNodeID, ast.NodeGroupStmts, children...)
                yylex.(*Parser).Children($$, ast.NodeGroupStmt, stmtListNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupElse, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


parameter_list:
        non_empty_parameter_list
            {
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_parameter_list:
        parameter
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_parameter_list ',' parameter
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

parameter:
        optional_class_type is_reference is_variadic T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))

                var posID graph.NodeID
                if $1 != 0 {
                    posID = yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$4}, nil)
                } else if $2 != nil {
                    posID = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2, $4}, nil)
                } else if $3 != nil {
                    posID = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3, $4}, nil)
                } else {
                    posID = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil)
                }

                var flag ast.NodeFlag
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                }
                if $3 != nil {
                    flag = flag | ast.NodeFlagVariadic
                }

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeParameter,
                    Flag: flag,
                })
                yylex.(*Parser).SavePosition($$, posID)

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupVarType, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupVar, varNodeID)

                // save tokens

                tmp := [4][]scanner.Token{}
                if $2 != nil {
                    tmp[1] = $2.HiddenTokens
                }
                if $3 != nil {
                    tmp[2] = $3.HiddenTokens
                }
                tmp[3] = $4.HiddenTokens

                if $3 == nil {
                    tmp[2] = tmp[3]
                    tmp[3] = nil
                }
                if $2 == nil {
                    tmp[1] = tmp[2]
                    tmp[2] = nil
                }
                if $1 == 0 {
                    tmp[0] = tmp[1]
                    tmp[1] = nil
                }

                if $1 != 0 {
                    yylex.(*Parser).MoveStartTokens($1, $$)
                }
                if tmp[0] != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, tmp[0])
                }
                if tmp[1] != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupOptionalType, tmp[1])
                }
                if tmp[2] != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupAmpersand, tmp[2])
                }
                if tmp[3] != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupVariadic, tmp[3])
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   optional_class_type is_reference is_variadic T_VARIABLE '=' static_scalar
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))

                var posID graph.NodeID
                if $1 != 0 {
                    posID = yylex.(*Parser).NewPosition([]graph.NodeID{$1, $6}, nil, nil)
                } else if $2 != nil {
                    posID = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, []graph.NodeID{$6})
                } else if $3 != nil {
                    posID = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, []graph.NodeID{$6})
                } else {
                    posID = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, []graph.NodeID{$6})
                }

                var flag ast.NodeFlag
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                }
                if $3 != nil {
                    flag = flag | ast.NodeFlagVariadic
                }

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeParameter,
                    Flag: flag,
                })
                yylex.(*Parser).SavePosition($$, posID)

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupVarType, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupVar, varNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupDefaultValue, $6)

                // save tokens

                tmp := [4][]scanner.Token{}
                if $2 != nil {
                    tmp[1] = $2.HiddenTokens
                }
                if $3 != nil {
                    tmp[2] = $3.HiddenTokens
                }
                tmp[3] = $4.HiddenTokens

                if $3 == nil {
                    tmp[2] = tmp[3]
                    tmp[3] = nil
                }
                if $2 == nil {
                    tmp[1] = tmp[2]
                    tmp[2] = nil
                }
                if $1 == 0 {
                    tmp[0] = tmp[1]
                    tmp[1] = nil
                }

                if $1 != 0 {
                    yylex.(*Parser).MoveStartTokens($1, $$)
                }
                if tmp[0] != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, tmp[0])
                }
                if tmp[1] != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupOptionalType, tmp[1])
                }
                if tmp[2] != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupAmpersand, tmp[2])
                }
                if tmp[3] != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupVariadic, tmp[3])
                }
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $5.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


optional_class_type:
        /* empty */
            {
                $$ = 0 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ARRAY
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CALLABLE
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

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
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeArgumentList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupArgumentList, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' non_empty_function_call_parameter_list ')'
            {
                children := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeArgumentList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupArguments, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupArgumentList, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' yield_expr ')'
            {
                argumentNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeArgument,
                })
                yylex.(*Parser).SavePosition(argumentNodeID, yylex.(*Parser).NewPosition([]graph.NodeID{$2}, nil, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeArgumentList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupArguments, argumentNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens($2, argumentNodeID)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupArgumentList, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


non_empty_function_call_parameter_list:
        function_call_parameter
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_function_call_parameter_list ',' function_call_parameter
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

function_call_parameter:
        expr_without_variable
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeArgument,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $1)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeArgument,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $1)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' w_variable
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeArgument,
                    Flag: ast.NodeFlagRef,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ELLIPSIS expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeArgument,
                    Flag: ast.NodeFlagVariadic,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

global_var_list:
        global_var_list ',' global_var
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   global_var
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }


global_var:
        T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVarName, identifierNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' r_variable
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupVarName, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupDollar, []scanner.Token{*$1})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' '{' expr '}'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVarName, $3)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupDollar, []scanner.Token{*$1})
                yylex.(*Parser).PrependTokens($3, ast.TokenGroupStart, []scanner.Token{*$2})
                yylex.(*Parser).PrependTokens($3, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($3, ast.TokenGroupEnd, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($3, ast.TokenGroupEnd, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


static_var_list:
        static_var_list ',' T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                staticVarNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStaticVar,
                })
                yylex.(*Parser).SavePosition(staticVarNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children(staticVarNodeID, ast.NodeGroupVar, varNodeID)

                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add(staticVarNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(staticVarNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_var_list ',' T_VARIABLE '=' static_scalar
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                staticVarNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStaticVar,
                })
                yylex.(*Parser).SavePosition(staticVarNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, []graph.NodeID{$5}))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children(staticVarNodeID, ast.NodeGroupVar, varNodeID)
                yylex.(*Parser).Children(staticVarNodeID, ast.NodeGroupExpr, $5)

                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add(staticVarNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(staticVarNodeID, ast.TokenGroupStart, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(staticVarNodeID, ast.TokenGroupVar, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                staticVarNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStaticVar,
                })
                yylex.(*Parser).SavePosition(staticVarNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children(staticVarNodeID, ast.NodeGroupVar, varNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(staticVarNodeID, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(staticVarNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '=' static_scalar
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                staticVarNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStaticVar,
                })
                yylex.(*Parser).SavePosition(staticVarNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$3}))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children(staticVarNodeID, ast.NodeGroupVar, varNodeID)
                yylex.(*Parser).Children(staticVarNodeID, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).AppendTokens(staticVarNodeID, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(staticVarNodeID, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(staticVarNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


class_statement_list:
        class_statement_list class_statement
            {
                yylex.(*Parser).List.Add($2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


class_statement:
        variable_modifiers class_variable_declaration ';'
            {
                childrenProperties := yylex.(*Parser).List.Pop()
                childrenModifiers := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtPropertyList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenModifiers, []*scanner.Token{$3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupModifiers, childrenModifiers...)
                yylex.(*Parser).Children($$, ast.NodeGroupProperties, childrenProperties...)

                // save tokens
                yylex.(*Parser).MoveStartTokens(childrenModifiers[0], $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupPropertyList, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_constant_declaration ';'
            {
                $$ = $1

                // save position
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$2}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupConstList, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$2})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_use_statement
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   method_modifiers function is_reference T_STRING '(' parameter_list ')' method_body
            {
                childrenParams := yylex.(*Parser).List.Pop()
                childrenModifiers := yylex.(*Parser).List.Pop()

                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))

                var posID graph.NodeID
                if len(childrenModifiers) == 0 {
                    posID = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, []graph.NodeID{$8})
                } else {
                    posID = yylex.(*Parser).NewPosition(childrenModifiers, nil, []graph.NodeID{$8})
                }

                var flag ast.NodeFlag
                if $3 != nil {
                    flag = flag | ast.NodeFlagRef
                }

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtClassMethod,
                    Flag: flag,
                })
                yylex.(*Parser).SavePosition($$, posID)

                yylex.(*Parser).Children($$, ast.NodeGroupMethodName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupStmt, $8)
                yylex.(*Parser).Children($$, ast.NodeGroupModifiers, childrenModifiers...)
                yylex.(*Parser).Children($$, ast.NodeGroupParams, childrenParams...)

                // save tokens
                if len(childrenModifiers) > 0 {
                    yylex.(*Parser).MoveStartTokens(childrenModifiers[0], $$)
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupModifierList, $2.HiddenTokens)
                } else {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $2.HiddenTokens)
                }
                if $3 == nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupFunction, $4.HiddenTokens)
                } else {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupFunction, $3.HiddenTokens)
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupAmpersand, $4.HiddenTokens)
                }
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $5.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupParameterList, $7.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_use_statement:
        T_USE trait_list trait_adaptations
            {
                childrenTraits := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtTraitUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$3}))

                yylex.(*Parser).Children($$, ast.NodeGroupTraitAdaptationList, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupTraits, childrenTraits...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_list:
        fully_qualified_class_name
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_list ',' fully_qualified_class_name
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptations:
        ';'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtNop,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$1})


                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' trait_adaptation_list '}'
            {
                children := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtTraitAdaptationList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupAdaptations, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupAdaptationList, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptation_list:
        /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_trait_adaptation_list
            {
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_trait_adaptation_list:
        trait_adaptation_statement
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_trait_adaptation_list trait_adaptation_statement
            {
                yylex.(*Parser).List.Add($2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptation_statement:
        trait_precedence ';'
            {
                $$ = $1;

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupNameList, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$2})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_alias ';'
            {
                $$ = $1;

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupAlias, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$2})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_precedence:
        trait_method_reference_fully_qualified T_INSTEADOF trait_reference_list
            {
                children := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtTraitUsePrecedence,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, nil, children))

                yylex.(*Parser).Children($$, ast.NodeGroupRef, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupInsteadof, children...)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupRef, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_reference_list:
        fully_qualified_class_name
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_reference_list ',' fully_qualified_class_name
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_method_reference:
        T_STRING
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtTraitMethodRef,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupMethod, identifierNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

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
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtTraitMethodRef,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupTrait, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupMethod, identifierNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_alias:
        trait_method_reference T_AS trait_modifiers T_STRING
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtTraitUseAlias,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupRef, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupModifier, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupAlias, identifierNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupRef, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_method_reference T_AS member_modifier
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtTraitUseAlias,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupRef, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupModifier, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupRef, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_modifiers:
        /* empty */
            {
                $$ = 0 

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
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtNop,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$1})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' inner_statement_list '}'
            {
                children := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStmtList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))
                
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

variable_modifiers:
        non_empty_member_modifiers
            {
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VAR
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(identifierNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

method_modifiers:
        /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_member_modifiers
            {
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_member_modifiers:
        member_modifier
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_member_modifiers member_modifier
            {
                yylex.(*Parser).List.Add($2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

member_modifier:
        T_PUBLIC
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PROTECTED
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PRIVATE
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ABSTRACT
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINAL
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_variable_declaration:
        class_variable_declaration ',' T_VARIABLE
            {
                prevNodeID := yylex.(*Parser).List.Last()

                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))
                
                propertyNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtProperty,
                })
                yylex.(*Parser).SavePosition(propertyNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children(propertyNodeID, ast.NodeGroupVar, varNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(propertyNodeID, ast.TokenGroupStart, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(propertyNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_variable_declaration ',' T_VARIABLE '=' static_scalar
            {
                prevNodeID := yylex.(*Parser).List.Last()

                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))
                
                propertyNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtProperty,
                })
                yylex.(*Parser).SavePosition(propertyNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, []graph.NodeID{$5}))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children(propertyNodeID, ast.NodeGroupVar, varNodeID)
                yylex.(*Parser).Children(propertyNodeID, ast.NodeGroupExpr, $5)

                // save tokens
                yylex.(*Parser).AppendTokens(propertyNodeID, ast.TokenGroupStart, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(propertyNodeID, ast.TokenGroupVar, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).List.Add(propertyNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))
                
                propertyNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtProperty,
                })
                yylex.(*Parser).SavePosition(propertyNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children(propertyNodeID, ast.NodeGroupVar, varNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(propertyNodeID, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(propertyNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '=' static_scalar
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))
                
                propertyNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtProperty,
                })
                yylex.(*Parser).SavePosition(propertyNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$3}))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children(propertyNodeID, ast.NodeGroupVar, varNodeID)
                yylex.(*Parser).Children(propertyNodeID, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).AppendTokens(propertyNodeID, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(propertyNodeID, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(propertyNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_constant_declaration:
        class_constant_declaration ',' T_STRING '=' static_scalar
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))
                
                constantNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtConstant,
                })
                yylex.(*Parser).SavePosition(constantNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, []graph.NodeID{$5}))

                $$ = $1
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $5}, nil, nil))

                yylex.(*Parser).Children(constantNodeID, ast.NodeGroupConstantName, identifierNodeID)
                yylex.(*Parser).Children(constantNodeID, ast.NodeGroupExpr, $5)
                yylex.(*Parser).Children($1, ast.NodeGroupConsts, constantNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(constantNodeID, ast.TokenGroupStart, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(constantNodeID, ast.TokenGroupName, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONST T_STRING '=' static_scalar
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))
                
                constantNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtConstant,
                })
                yylex.(*Parser).SavePosition(constantNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, []graph.NodeID{$4}))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtClassConstList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$4}))

                yylex.(*Parser).Children(constantNodeID, ast.NodeGroupConstantName, identifierNodeID)
                yylex.(*Parser).Children(constantNodeID, ast.NodeGroupExpr, $4)
                yylex.(*Parser).Children($$, ast.NodeGroupConsts, constantNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(constantNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(constantNodeID, ast.TokenGroupName, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

echo_expr_list:
        echo_expr_list ',' expr
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }


for_expr:
        /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_for_expr
            {
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_for_expr:
        non_empty_for_expr ',' expr
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

chaining_method_or_property:
        chaining_method_or_property variable_property
            {
                children2 := yylex.(*Parser).List.Pop()
                children1 := yylex.(*Parser).List.Pop()
                
                yylex.(*Parser).List.Push()

                for _, v := range(children1) {
                    yylex.(*Parser).List.Add(v)
                }
                for _, v := range(children2) {
                    yylex.(*Parser).List.Add(v)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_property
            {
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

chaining_dereference:
        chaining_dereference '[' dim_offset ']'
            {
                fetchNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition(fetchNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2, $4}, nil))

                yylex.(*Parser).Children(fetchNodeID, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupExpr, []scanner.Token{*$4})

                yylex.(*Parser).List.Add(fetchNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' dim_offset ']'
            {
                fetchNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition(fetchNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children(fetchNodeID, ast.NodeGroupDim, $2)

                // save tokens
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupVar, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupVar, []scanner.Token{*$1})
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupExpr, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupExpr, []scanner.Token{*$3})

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(fetchNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

chaining_instance_call:
        chaining_dereference chaining_method_or_property
            {
                children2 := yylex.(*Parser).List.Pop()
                children1 := yylex.(*Parser).List.Pop()
                
                yylex.(*Parser).List.Push()

                for _, v := range(children1) {
                    yylex.(*Parser).List.Add(v)
                }
                for _, v := range(children2) {
                    yylex.(*Parser).List.Add(v)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   chaining_dereference
            {
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   chaining_method_or_property
            {
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

instance_call:
        /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   chaining_instance_call
            {
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

new_expr:
        T_NEW class_name_reference ctor_arguments
            {

                lastNodeID := $3
                if lastNodeID == 0 {
                    lastNodeID = $2
                }

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprNew,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{lastNodeID}))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $2)
                yylex.(*Parser).Children($$, ast.NodeGroupArgumentList, $3)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

expr_without_variable:
        T_LIST '(' assignment_list ')' '=' expr
            {
                listNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprList,
                })
                yylex.(*Parser).SavePosition(listNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignAssign,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$6}))

                yylex.(*Parser).Children(listNodeID, ast.NodeGroupItems, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children($$, ast.NodeGroupVar, listNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $6)

                // save tokens
                yylex.(*Parser).AppendTokens(listNodeID, ast.TokenGroupList, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(listNodeID, ast.TokenGroupArrayPairList, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $5.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignAssign,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' '&' variable
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignReference,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $4}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $4)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEqual, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' '&' T_NEW class_name_reference ctor_arguments
            {

                lastNodeID := $6
                if lastNodeID == 0 {
                    lastNodeID = $5
                }

                newNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprNew,
                })
                yylex.(*Parser).SavePosition(newNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, []graph.NodeID{lastNodeID}))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignReference,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, newNodeID}, nil, nil))

                yylex.(*Parser).Children(newNodeID, ast.NodeGroupClass, $5)
                yylex.(*Parser).Children(newNodeID, ast.NodeGroupArgumentList, $6)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, newNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEqual, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(newNodeID, ast.TokenGroupStart, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLONE expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprClone,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_PLUS_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignPlus,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MINUS_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignMinus,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MUL_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignMul,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_POW_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignPow,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_DIV_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignDiv,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_CONCAT_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignConcat,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MOD_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignMod,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_AND_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignBitwiseAnd,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_OR_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignBitwiseOr,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_XOR_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignBitwiseXor,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_SL_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignShiftLeft,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_SR_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignShiftRight,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   rw_variable T_INC
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprPostInc,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$2}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INC rw_variable
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprPreInc,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   rw_variable T_DEC
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprPostDec,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$2}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DEC rw_variable
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprPreDec,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_BOOLEAN_OR expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryBooleanOr,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_BOOLEAN_AND expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryBooleanAnd,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_OR expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryLogicalOr,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_AND expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryLogicalAnd,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_XOR expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryLogicalXor,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '|' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryBitwiseOr,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '&' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryBitwiseAnd,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '^' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryBitwiseXor,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '.' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryConcat,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '+' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryPlus,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '-' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryMinus,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '*' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryMul,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_POW expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryPow,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '/' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryDiv,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '%' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryMod,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryShiftLeft,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SR expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryShiftRight,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '+' expr %prec T_INC
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprUnaryPlus,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '-' expr %prec T_INC
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprUnaryMinus,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '!' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprBooleanNot,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '~' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprBitwiseNot,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_IDENTICAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryIdentical,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_NOT_IDENTICAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryNotIdentical,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryEqual,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_NOT_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryNotEqual,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEqual, []scanner.Token{*$2})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '<' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinarySmaller,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_SMALLER_OR_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinarySmallerOrEqual,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '>' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryGreater,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_GREATER_OR_EQUAL expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryGreaterOrEqual,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_INSTANCEOF class_name_reference
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprInstanceOf,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupClass, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   parenthesis_expr
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)

                // TODO: handle parenthesis tokens
                //yylex.(*Parser).setFreeFloating($1, freefloating.Start, append((*$1.GetFreeFloating())[freefloating.OpenParenthesisToken], (*$1.GetFreeFloating())[freefloating.Start]...)); delete((*$1.GetFreeFloating()), freefloating.OpenParenthesisToken)
                //yylex.(*Parser).setFreeFloating($1, freefloating.End, append((*$1.GetFreeFloating())[freefloating.End], (*$1.GetFreeFloating())[freefloating.CloseParenthesisToken]...)); delete((*$1.GetFreeFloating()), freefloating.CloseParenthesisToken)
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
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, []scanner.Token{*$1})
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, []scanner.Token{*$3})

                children := yylex.(*Parser).List.Pop()
                for _, n := range(children) {
                    yylex.(*Parser).Children(n, ast.NodeGroupVar, $$)
                    yylex.(*Parser).MoveStartTokens($$, n)
                    $$ = n

                    // save position
                    yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$$, n}, nil, nil))
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '?' expr ':' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprTernary,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $5}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupCond, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupIfTrue, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupIfFalse, $5)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupTrue, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '?' ':' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprTernary,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $4}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupCond, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupIfFalse, $4)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupTrue, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   internal_functions_in_yacc
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INT_CAST expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeCastInt,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCast, []scanner.Token{*$1})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOUBLE_CAST expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeCastDouble,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCast, []scanner.Token{*$1})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STRING_CAST expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeCastString,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCast, []scanner.Token{*$1})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ARRAY_CAST expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeCastArray,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCast, []scanner.Token{*$1})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_OBJECT_CAST expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeCastObject,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCast, []scanner.Token{*$1})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BOOL_CAST expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeCastBool,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCast, []scanner.Token{*$1})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_UNSET_CAST expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeCastUnset,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCast, []scanner.Token{*$1})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXIT exit_expr
            {
                $$ = $2

                var flag ast.NodeFlag
                exitTknValue := yylex.(*Parser).Ast.FileData[$1.StartPos:$1.EndPos]
                if bytes.EqualFold(exitTknValue, []byte("die")) {
                    flag = ast.NodeFlagAltSyntax
                }

                if $$ == 0 {
                    $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                        Type: ast.NodeTypeExprExit,
                        Flag: flag,
                    })
                    yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))
                } else {
                    n := yylex.(*Parser).Ast.Graph.GetNode($$)
                    yylex.(*Parser).Ast.Nodes[n.ID].Flag = flag
                    yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))
                }

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '@' expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprErrorSuppress,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

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
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprShellExec,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupParts, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PRINT expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprPrint,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprYield,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   function is_reference '(' parameter_list ')' lexical_vars '{' inner_statement_list '}'
            {
                var flag ast.NodeFlag
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                }

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprClosure,
                    Flag: flag,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $9}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClosureUse, $6)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children($$, ast.NodeGroupParams, yylex.(*Parser).List.Pop()...)
                
                // // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                if $2 == nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupFunction, $3.HiddenTokens)
                } else {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupFunction, $2.HiddenTokens)
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupAmpersand, $3.HiddenTokens)
                }
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupParameterList, $5.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupReturnType, $7.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $9.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC function is_reference '(' parameter_list ')' lexical_vars '{' inner_statement_list '}'
            {
                flag := ast.NodeFlagStatic
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                }

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprClosure,
                    Flag: flag,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $10}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClosureUse, $7)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children($$, ast.NodeGroupParams, yylex.(*Parser).List.Pop()...)
                
                // // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStatic, $2.HiddenTokens)
                if $3 == nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupFunction, $4.HiddenTokens)
                } else {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupFunction, $3.HiddenTokens)
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupAmpersand, $4.HiddenTokens)
                }
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupParameterList, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupReturnType, $8.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $10.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

yield_expr:
        T_YIELD expr_without_variable
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprYield,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupVal, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD variable
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprYield,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupVal, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD expr T_DOUBLE_ARROW expr_without_variable
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprYield,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$4}))

                yylex.(*Parser).Children($$, ast.NodeGroupKey, $2)
                yylex.(*Parser).Children($$, ast.NodeGroupVal, $4)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD expr T_DOUBLE_ARROW variable
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprYield,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$4}))

                yylex.(*Parser).Children($$, ast.NodeGroupKey, $2)
                yylex.(*Parser).Children($$, ast.NodeGroupVal, $4)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

combined_scalar_offset:
        combined_scalar '[' dim_offset ']'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   combined_scalar_offset '[' dim_offset ']'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONSTANT_ENCAPSED_STRING '[' dim_offset ']'
            {
                stringNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarString,
                })
                yylex.(*Parser).SavePosition(stringNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, stringNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   general_constant '[' dim_offset ']'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

combined_scalar:
        T_ARRAY '(' array_pair_list ')'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArray,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupItems, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupArray, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupArrayPairList, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' array_pair_list ']'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprShortArray,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupItems, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupArrayPairList, $3.HiddenTokens)

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
                $$ = 0

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE '(' lexical_var_list ')'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprClosureUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupUse, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupLexicalVarList, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

lexical_var_list:
        lexical_var_list ',' T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(varNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).List.Add(varNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   lexical_var_list ',' '&' T_VARIABLE
            {
                prevNodeID := yylex.(*Parser).List.Last()

                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))

                refNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprReference,
                })
                yylex.(*Parser).SavePosition(refNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3, $4}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children(refNodeID, ast.NodeGroupVar, varNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(varNodeID, ast.TokenGroupStart, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(varNodeID, ast.TokenGroupStart, $4.HiddenTokens)

                yylex.(*Parser).List.Add(refNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(varNodeID, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(varNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))

                refNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprReference,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children(refNodeID, ast.NodeGroupVar, varNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(varNodeID, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(varNodeID, ast.TokenGroupStart, $2.HiddenTokens)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(varNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

function_call:
        namespace_name function_call_parameter_list
            {
                childrenNameParts := yylex.(*Parser).List.Pop()

                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprFunctionCall,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{nameNodeID, $2}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupFunction, nameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupArgumentList, $2)

                // save tokens
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name function_call_parameter_list
            {
                children := yylex.(*Parser).List.Pop()
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameRelative,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, children...)

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprFunctionCall,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{nameNodeID, $4}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupFunction, nameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupArgumentList, $4)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupNamespace, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name function_call_parameter_list
            {
                children := yylex.(*Parser).List.Pop()
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameFullyQualified,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, children...)

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprFunctionCall,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{nameNodeID, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupFunction, nameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupArgumentList, $3)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM variable_name function_call_parameter_list
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprStaticCall,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $4}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupCall, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupArgumentList, $4)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects function_call_parameter_list
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprStaticCall,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $4}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupCall, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupArgumentList, $4)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_name function_call_parameter_list
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprStaticCall,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $4}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupCall, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupArgumentList, $4)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects function_call_parameter_list
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprStaticCall,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $4}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupCall, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupArgumentList, $4)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_without_objects function_call_parameter_list
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprFunctionCall,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $2}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupFunction, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupArgumentList, $2)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_name:
        T_STATIC
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name 
            {
                children := yylex.(*Parser).List.Pop()
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(children, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupParts, children...)

                // save tokens
                yylex.(*Parser).MoveStartTokens(children[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                children := yylex.(*Parser).List.Pop()
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameRelative,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))
                yylex.(*Parser).Children($$, ast.NodeGroupParts, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupNamespace, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                children := yylex.(*Parser).List.Pop()
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameFullyQualified,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))
                yylex.(*Parser).Children($$, ast.NodeGroupParts, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

fully_qualified_class_name:
        namespace_name
            {
                children := yylex.(*Parser).List.Pop()
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(children, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupParts, children...)

                // save tokens
                yylex.(*Parser).MoveStartTokens(children[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                children := yylex.(*Parser).List.Pop()
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameRelative,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))
                yylex.(*Parser).Children($$, ast.NodeGroupParts, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupNamespace, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                children := yylex.(*Parser).List.Pop()
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameFullyQualified,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))
                yylex.(*Parser).Children($$, ast.NodeGroupParts, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

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
                children2 := yylex.(*Parser).List.Pop()
                children1 := yylex.(*Parser).List.Pop()

                $$ = $1

                // save tokens
                yylex.(*Parser).AppendTokens(children1[0], ast.TokenGroupStart, $2.HiddenTokens)

                for _, n := range(children1) {
                    yylex.(*Parser).Children(n, ast.NodeGroupVar, $$)
                    yylex.(*Parser).MoveStartTokens($$, n)

                    // save position
                    yylex.(*Parser).SavePosition(n, yylex.(*Parser).NewPosition([]graph.NodeID{$$, n}, nil, nil))

                    $$ = n
                }

                for _, n := range(children2) {
                    yylex.(*Parser).Children(n, ast.NodeGroupVar, $$)
                    yylex.(*Parser).MoveStartTokens($$, n)

                    // save position
                    yylex.(*Parser).SavePosition(n, yylex.(*Parser).NewPosition([]graph.NodeID{$$, n}, nil, nil))

                    $$ = n
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
                children2 := yylex.(*Parser).List.Pop()
                children1 := yylex.(*Parser).List.Pop()
                
                yylex.(*Parser).List.Push()

                for _, v := range(children1) {
                    yylex.(*Parser).List.Add(v)
                }
                for _, v := range(children2) {
                    yylex.(*Parser).List.Add(v)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


dynamic_class_name_variable_property:
        T_OBJECT_OPERATOR object_property
            {
                children := yylex.(*Parser).List.Pop()
                
                // save tokens
                yylex.(*Parser).AppendTokens(children[0], ast.TokenGroupVar, $1.HiddenTokens)

                yylex.(*Parser).List.Push()
                for _, v := range(children) {
                    yylex.(*Parser).List.Add(v)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

exit_expr:
        /* empty */
            {
                $$ = 0

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' ')'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprExit,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExit, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExit, []scanner.Token{*$1})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, []scanner.Token{*$2})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   parenthesis_expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprExit,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{yylex.(*Parser).CurrentToken}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)

                // save comments
                // TODO handle parenthesis tokens
                // yylex.(*Parser).setFreeFloating($$, freefloating.Exit, (*$1.GetFreeFloating())[freefloating.OpenParenthesisToken]); delete((*$1.GetFreeFloating()), freefloating.OpenParenthesisToken)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, (*$1.GetFreeFloating())[freefloating.CloseParenthesisToken]); delete((*$1.GetFreeFloating()), freefloating.CloseParenthesisToken)
            }
;

backticks_expr:
        /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ENCAPSED_AND_WHITESPACE
            {
                yylex.(*Parser).List.Push()
                nodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarEncapsedStringPart,
                })
                yylex.(*Parser).SavePosition(nodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))
                yylex.(*Parser).List.Add(nodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   encaps_list
            {
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

ctor_arguments:
        /* empty */
            {
                $$ = 0

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
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarLnumber,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DNUMBER
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarDnumber,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONSTANT_ENCAPSED_STRING
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarString,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LINE
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarMagicConstant,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FILE
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarMagicConstant,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DIR
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarMagicConstant,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRAIT_C
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarMagicConstant,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_METHOD_C
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarMagicConstant,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FUNC_C
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarMagicConstant,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_C
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarMagicConstant,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC
            {
                stringPartNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarEncapsedStringPart,
                })
                yylex.(*Parser).SavePosition(stringPartNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarHeredoc,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupParts, stringPartNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC T_END_HEREDOC
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarHeredoc,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_class_constant:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprClassConstFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupConstantName, identifierNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $3.HiddenTokens)

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
                children := yylex.(*Parser).List.Pop()
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(children, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, children...)

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprConstFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{nameNodeID}, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupConstant, nameNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens(children[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                children := yylex.(*Parser).List.Pop()
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameRelative,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, children...)

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprConstFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{nameNodeID}, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupConstant, nameNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupNamespace, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                children := yylex.(*Parser).List.Pop()
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameFullyQualified,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, children...)

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprConstFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{nameNodeID}, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupConstant, nameNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ARRAY '(' static_array_pair_list ')'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArray,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupItems, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupArray, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupArrayPairList, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' static_array_pair_list ']'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprShortArray,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupItems, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupArrayPairList, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_class_constant
            {
                $$ = $1 

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLASS_C
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarMagicConstant,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

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
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '+' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryPlus,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '-' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryMinus,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '*' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryMul,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_POW static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryPow,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '/' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryDiv,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '%' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryMod,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '!' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprBooleanNot,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '~' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprBitwiseNot,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '|' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryBitwiseOr,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '&' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryBitwiseAnd,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '^' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryBitwiseXor,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_SL static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryShiftLeft,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_SR static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryShiftRight,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '.' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryConcat,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_LOGICAL_XOR static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryLogicalXor,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_LOGICAL_AND static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryLogicalAnd,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_LOGICAL_OR static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryLogicalOr,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_BOOLEAN_AND static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryBooleanAnd,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_BOOLEAN_OR static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryBooleanOr,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_IDENTICAL static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryIdentical,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_NOT_IDENTICAL static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryNotIdentical,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_EQUAL static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryEqual,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_NOT_EQUAL static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryNotEqual,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEqual, []scanner.Token{*$2})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '<' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinarySmaller,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '>' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryGreater,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_SMALLER_OR_EQUAL static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinarySmallerOrEqual,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_IS_GREATER_OR_EQUAL static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryGreaterOrEqual,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '?' ':' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprTernary,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $4}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupCond, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupIfFalse, $4)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupTrue, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value '?' static_scalar_value ':' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprTernary,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $5}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupCond, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupIfTrue, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupIfFalse, $5)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupTrue, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '+' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprUnaryPlus,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '-' static_scalar_value
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprUnaryMinus,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' static_scalar_value ')'
            {
                $$ = $2;

                // save tokens
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, []scanner.Token{*$1})
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, []scanner.Token{*$3})

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
                children := yylex.(*Parser).List.Pop()
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(children, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, children...)

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprConstFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{nameNodeID}, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupConstant, nameNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens(children[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                children := yylex.(*Parser).List.Pop()
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameRelative,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, children...)

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprConstFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{nameNodeID}, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupConstant, nameNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupNamespace, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name
            {
                children := yylex.(*Parser).List.Pop()
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameFullyQualified,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, children...)

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprConstFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{nameNodeID}, nil, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupConstant, nameNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

scalar:
        T_STRING_VARNAME
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVarName, identifierNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, []scanner.Token{*$1})

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
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarEncapsed,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupParts, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC encaps_list T_END_HEREDOC
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarHeredoc,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupParts, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLASS_C
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarMagicConstant,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_array_pair_list:
        /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_static_array_pair_list possible_comma
            {
                prevNodeID := yylex.(*Parser).List.Last()

                if $2 != nil {
                    yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                    yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, []scanner.Token{*$2})
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
                prevNodeID := yylex.(*Parser).List.Last()

                itemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition(itemNodeID, yylex.(*Parser).NewPosition([]graph.NodeID{$3, $5}, nil, nil))

                yylex.(*Parser).Children(itemNodeID, ast.NodeGroupKey, $3)
                yylex.(*Parser).Children(itemNodeID, ast.NodeGroupVal, $5)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).MoveStartTokens($3, itemNodeID)
                yylex.(*Parser).AppendTokens(itemNodeID, ast.TokenGroupExpr, $4.HiddenTokens)

                yylex.(*Parser).List.Add(itemNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_static_array_pair_list ',' static_scalar_value
            {
                prevNodeID := yylex.(*Parser).List.Last()

                itemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition(itemNodeID, yylex.(*Parser).NewPosition( []graph.NodeID{$3}, nil, nil))

                yylex.(*Parser).Children(itemNodeID, ast.NodeGroupVal, $3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).MoveStartTokens($3, itemNodeID)

                yylex.(*Parser).List.Add(itemNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value T_DOUBLE_ARROW static_scalar_value
            {
                itemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition(itemNodeID, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children(itemNodeID, ast.NodeGroupKey, $1)
                yylex.(*Parser).Children(itemNodeID, ast.NodeGroupVal, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, itemNodeID)
                yylex.(*Parser).AppendTokens(itemNodeID, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(itemNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_scalar_value
            {
                itemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition(itemNodeID, yylex.(*Parser).NewPosition( []graph.NodeID{$1}, nil, nil))

                yylex.(*Parser).Children(itemNodeID, ast.NodeGroupVal, $1)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, itemNodeID)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(itemNodeID)

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
                $$ = $2;

                // save tokens
                // TODO: handle parenthesis tokens
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, []scanner.Token{*$1})
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' yield_expr ')'
            {
                $$ = $2;

                // save tokens
                // TODO: handle parenthesis tokens
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, []scanner.Token{*$1})
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, []scanner.Token{*$3})

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
                children5 := yylex.(*Parser).List.Pop()
                children4 := yylex.(*Parser).List.Pop()
                children3 := yylex.(*Parser).List.Pop()

                $$ = $1

                if len(children4) > 0 {

                    node := children3[len(children3)-1]

                    yylex.(*Parser).Ast.Foreach(node, func(e graph.Edge, n graph.Node) bool {
                        if n.Type != stxtree.NodeTypeNode {
                            return false
                        }

                        yylex.(*Parser).Children(children4[0], ast.NodeGroupMethod, e.To)

                        return true
                    })

                    children3 = append(children3[:len(children3)-1], children4...)
                }

                // save tokens
                yylex.(*Parser).AppendTokens(children3[0], ast.TokenGroupVar, $2.HiddenTokens)

                for _, n := range(children3) {
                    yylex.(*Parser).Children(n, ast.NodeGroupVar, $$)
                    yylex.(*Parser).MoveStartTokens($$, n)

                    // save position
                    yylex.(*Parser).SavePosition(n, yylex.(*Parser).NewPosition([]graph.NodeID{$$, n}, nil, nil))

                    $$ = n
                }

                for _, n := range(children5) {
                    yylex.(*Parser).Children(n, ast.NodeGroupVar, $$)
                    yylex.(*Parser).MoveStartTokens($$, n)

                    // save position
                    yylex.(*Parser).SavePosition(n, yylex.(*Parser).NewPosition([]graph.NodeID{$$, n}, nil, nil))

                    $$ = n
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
                children2 := yylex.(*Parser).List.Pop()
                children1 := yylex.(*Parser).List.Pop()
                
                yylex.(*Parser).List.Push()

                for _, v := range(children1) {
                    yylex.(*Parser).List.Add(v)
                }
                for _, v := range(children2) {
                    yylex.(*Parser).List.Add(v)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


variable_property:
        T_OBJECT_OPERATOR object_property method_or_not
            {
                children3 := yylex.(*Parser).List.Pop()
                children2 := yylex.(*Parser).List.Pop()

                if len(children3) > 0 {

                    node := children2[len(children2)-1]

                    yylex.(*Parser).Ast.Foreach(node, func(e graph.Edge, n graph.Node) bool {
                        if n.Type != stxtree.NodeTypeNode {
                            return false
                        }

                        yylex.(*Parser).Children(children3[0], ast.NodeGroupMethod, e.To)

                        return true
                    })

                    children2 = append(children2[:len(children2)-1], children3...)
                }

                yylex.(*Parser).List.Push()

                for _, v := range(children2) {
                    yylex.(*Parser).List.Add(v)
                }

                // save tokens
                yylex.(*Parser).AppendTokens(children2[0], ast.TokenGroupVar, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

array_method_dereference:
        array_method_dereference '[' dim_offset ']'
            {
                fetchNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition(fetchNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2, $4}, nil))

                yylex.(*Parser).Children(fetchNodeID, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupExpr, []scanner.Token{*$4})

                yylex.(*Parser).List.Add(fetchNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   method '[' dim_offset ']'
            {
                fetchNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition(fetchNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2, $4}, nil))

                yylex.(*Parser).Children(fetchNodeID, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens(fetchNodeID, ast.TokenGroupExpr, []scanner.Token{*$4})

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)
                yylex.(*Parser).List.Add(fetchNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

method:
        function_call_parameter_list
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprMethodCall,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupArgumentList, $1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

method_or_not:
        method
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   array_method_dereference
            {
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                yylex.(*Parser).List.Push()

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
                children := yylex.(*Parser).List.Pop()
                last := children[len(children)-1]

                yylex.(*Parser).Children(last, ast.NodeGroupVarName, $2)

                for _, n := range(children) {
                    yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{n, $2}, nil, nil))
                }

                $$ = children[0]

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_member:
        class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprStaticPropertyFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupProperty, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM variable_without_objects
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprStaticPropertyFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupProperty, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $2.HiddenTokens)

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
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   function_call '[' dim_offset ']'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, []scanner.Token{*$4})

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
                children := yylex.(*Parser).List.Pop()
                last := children[len(children)-1]

                yylex.(*Parser).Children(last, ast.NodeGroupVarName, $2)

                for _, n := range(children) {
                    yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{n, $2}, nil, nil))
                }

                $$ = children[0]

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
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   reference_variable '{' expr '}'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, []scanner.Token{*$4})

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
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVarName, identifierNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' '{' expr '}'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVarName, $3)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupDollar, []scanner.Token{*$1})
                yylex.(*Parser).PrependTokens($3, ast.TokenGroupStart, []scanner.Token{*$2})
                yylex.(*Parser).PrependTokens($3, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($3, ast.TokenGroupEnd, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($3, ast.TokenGroupEnd, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

dim_offset:
        /* empty */
            {
                $$ = 0 

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
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_without_objects
            {
                propertyFetchNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprStaticPropertyFetch,
                })
                yylex.(*Parser).SavePosition(propertyFetchNodeID, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, nil, nil))

                yylex.(*Parser).Children(propertyFetchNodeID, ast.NodeGroupProperty, $1)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(propertyFetchNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

object_dim_list:
        object_dim_list '[' dim_offset ']'
            {
                ArrayDimFetchNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition(ArrayDimFetchNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2, $4}, nil))

                yylex.(*Parser).Children(ArrayDimFetchNodeID, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).AppendTokens(ArrayDimFetchNodeID, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(ArrayDimFetchNodeID, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens(ArrayDimFetchNodeID, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens(ArrayDimFetchNodeID, ast.TokenGroupExpr, []scanner.Token{*$4})

                yylex.(*Parser).List.Add(ArrayDimFetchNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   object_dim_list '{' expr '}'
            {
                ArrayDimFetchNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition(ArrayDimFetchNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2, $4}, nil))

                yylex.(*Parser).Children(ArrayDimFetchNodeID, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).AppendTokens(ArrayDimFetchNodeID, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(ArrayDimFetchNodeID, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens(ArrayDimFetchNodeID, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens(ArrayDimFetchNodeID, ast.TokenGroupExpr, []scanner.Token{*$4})

                yylex.(*Parser).List.Add(ArrayDimFetchNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_name
            {
                propertyFetchNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprStaticPropertyFetch,
                })
                yylex.(*Parser).SavePosition(propertyFetchNodeID, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, nil, nil))

                yylex.(*Parser).Children(propertyFetchNodeID, ast.NodeGroupProperty, $1)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(propertyFetchNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

variable_name:
        T_STRING
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' expr '}'
            {
                $$ = $2;
                
                // save tokens
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, []scanner.Token{*$1})
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

simple_indirect_reference:
        '$'
            {
                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens(varNodeID, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(varNodeID, ast.TokenGroupDollar, []scanner.Token{*$1})

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(varNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   simple_indirect_reference '$'
            {
                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))

                yylex.(*Parser).Children(yylex.(*Parser).List.Last(), ast.NodeGroupVarName, varNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(varNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(varNodeID, ast.TokenGroupDollar, []scanner.Token{*$2})

                yylex.(*Parser).List.Add(varNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

assignment_list:
        assignment_list ',' assignment_list_element
            {
                if yylex.(*Parser).List.Len() == 0 {
                    ArrayItemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                        Type: ast.NodeTypeExprArrayItem,
                    })
                    yylex.(*Parser).List.Add(ArrayItemNodeID)
                }

                yylex.(*Parser).AppendTokens(yylex.(*Parser).List.Last(), ast.TokenGroupEnd, $2.HiddenTokens)

                if $3 == 0 {
                    ArrayItemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                        Type: ast.NodeTypeExprArrayItem,
                    })
                    yylex.(*Parser).List.Add(ArrayItemNodeID)
                } else {
                    yylex.(*Parser).List.Add($3)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   assignment_list_element
            {
                yylex.(*Parser).List.Push()
                if $1 != 0 {
                    yylex.(*Parser).List.Add($1)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


assignment_list_element:
        variable
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition( []graph.NodeID{$1}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVal, $1)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LIST '(' assignment_list ')'
            {
                listNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprList,
                })
                yylex.(*Parser).SavePosition(listNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children(listNodeID, ast.NodeGroupItems, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children($$, ast.NodeGroupVal, listNodeID)

                // TODO: Cannot use list() as standalone expression
                
                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(listNodeID, ast.TokenGroupList, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(listNodeID, ast.TokenGroupArrayPairList, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                $$ = 0

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;


array_pair_list:
        /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_array_pair_list possible_comma
            {
                if $2 != nil {
                    // save tokens
                    yylex.(*Parser).AppendTokens(yylex.(*Parser).List.Last(), ast.TokenGroupEnd, $2.HiddenTokens)

                    // seve node
                    ArrayItemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                        Type: ast.NodeTypeExprArrayItem,
                    })
                    yylex.(*Parser).List.Add(ArrayItemNodeID)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_array_pair_list:
        non_empty_array_pair_list ',' expr T_DOUBLE_ARROW expr
            {
                prevNodeID := yylex.(*Parser).List.Last()

                ArrayItemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition(ArrayItemNodeID, yylex.(*Parser).NewPosition([]graph.NodeID{$3, $5}, nil, nil))

                yylex.(*Parser).Children(ArrayItemNodeID, ast.NodeGroupKey, $3)
                yylex.(*Parser).Children(ArrayItemNodeID, ast.NodeGroupVal, $5)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).MoveStartTokens($3, ArrayItemNodeID)
                yylex.(*Parser).AppendTokens(ArrayItemNodeID, ast.TokenGroupExpr, $4.HiddenTokens)

                yylex.(*Parser).List.Add(ArrayItemNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_array_pair_list ',' expr
            {
                prevNodeID := yylex.(*Parser).List.Last()

                ArrayItemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition(ArrayItemNodeID, yylex.(*Parser).NewPosition( []graph.NodeID{$3}, nil, nil))

                yylex.(*Parser).Children(ArrayItemNodeID, ast.NodeGroupVal, $3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).MoveStartTokens($3, ArrayItemNodeID)

                yylex.(*Parser).List.Add(ArrayItemNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_DOUBLE_ARROW expr
            {
                ArrayItemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition(ArrayItemNodeID, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children(ArrayItemNodeID, ast.NodeGroupKey, $1)
                yylex.(*Parser).Children(ArrayItemNodeID, ast.NodeGroupVal, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, ArrayItemNodeID)
                yylex.(*Parser).AppendTokens(ArrayItemNodeID, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(ArrayItemNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                ArrayItemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition(ArrayItemNodeID, yylex.(*Parser).NewPosition( []graph.NodeID{$1}, nil, nil))

                yylex.(*Parser).Children(ArrayItemNodeID, ast.NodeGroupVal, $1)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, ArrayItemNodeID)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(ArrayItemNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_array_pair_list ',' expr T_DOUBLE_ARROW '&' w_variable
            {
                prevNodeID := yylex.(*Parser).List.Last()

                refNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprReference,
                })
                yylex.(*Parser).SavePosition(refNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$5}, []graph.NodeID{$6}))
                
                ArrayItemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition(ArrayItemNodeID, yylex.(*Parser).NewPosition([]graph.NodeID{$3, $6}, nil, nil))

                yylex.(*Parser).Children(refNodeID, ast.NodeGroupVar, $6)
                yylex.(*Parser).Children(ArrayItemNodeID, ast.NodeGroupKey, $3)
                yylex.(*Parser).Children(ArrayItemNodeID, ast.NodeGroupVal, refNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).MoveStartTokens($3, ArrayItemNodeID)
                yylex.(*Parser).AppendTokens(ArrayItemNodeID, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens(refNodeID, ast.TokenGroupStart, $5.HiddenTokens)

                yylex.(*Parser).List.Add(ArrayItemNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_array_pair_list ',' '&' w_variable
            {
                prevNodeID := yylex.(*Parser).List.Last()

                refNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprReference,
                })
                yylex.(*Parser).SavePosition(refNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, []graph.NodeID{$4}))
                
                ArrayItemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition(ArrayItemNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, []graph.NodeID{$4}))

                yylex.(*Parser).Children(refNodeID, ast.NodeGroupVar, $4)
                yylex.(*Parser).Children(ArrayItemNodeID, ast.NodeGroupVal, refNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(ArrayItemNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).List.Add(ArrayItemNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_DOUBLE_ARROW '&' w_variable
            {
                refNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprReference,
                })
                yylex.(*Parser).SavePosition(refNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, []graph.NodeID{$4}))
                
                ArrayItemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition(ArrayItemNodeID, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $4}, nil, nil))

                yylex.(*Parser).Children(refNodeID, ast.NodeGroupVar, $4)
                yylex.(*Parser).Children(ArrayItemNodeID, ast.NodeGroupKey, $1)
                yylex.(*Parser).Children(ArrayItemNodeID, ast.NodeGroupVal, refNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, ArrayItemNodeID)
                yylex.(*Parser).AppendTokens(ArrayItemNodeID, ast.TokenGroupExpr, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(refNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(ArrayItemNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' w_variable
            {
                refNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprReference,
                })
                yylex.(*Parser).SavePosition(refNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))
                
                ArrayItemNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition(ArrayItemNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children(refNodeID, ast.NodeGroupVar, $2)
                yylex.(*Parser).Children(ArrayItemNodeID, ast.NodeGroupVal, refNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(ArrayItemNodeID, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(ArrayItemNodeID)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_list:
        encaps_list encaps_var
            {
                yylex.(*Parser).List.Add($2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   encaps_list T_ENCAPSED_AND_WHITESPACE
            {
                encapsNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarEncapsedStringPart,
                })
                yylex.(*Parser).SavePosition(encapsNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))

                yylex.(*Parser).List.Add(encapsNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(encapsNodeID, ast.TokenGroupStart, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   encaps_var
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ENCAPSED_AND_WHITESPACE encaps_var
            {
                yylex.(*Parser).List.Push()
                
                encapsNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarEncapsedStringPart,
                })
                yylex.(*Parser).SavePosition(encapsNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                yylex.(*Parser).List.Add(encapsNodeID)
                yylex.(*Parser).List.Add($2)

                // save tokens
                yylex.(*Parser).AppendTokens(encapsNodeID, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_var:
        T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVarName, identifierNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '[' encaps_var_offset ']'
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupVar, varNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupDim, $3)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, []scanner.Token{*$2})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE T_OBJECT_OPERATOR T_STRING
            {
                varNameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(varNameNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                propertyNameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(propertyNameNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprPropertyFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, varNameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupVar, varNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupProperty, propertyNameNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(propertyNameNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES expr '}'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVarName, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, []scanner.Token{*$1})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}'
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVarName, identifierNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, []scanner.Token{*$1})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))

                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayDimFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $6}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupVar, varNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupDim, $4)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, []scanner.Token{*$1})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, []scanner.Token{*$3})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $5.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, []scanner.Token{*$5})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, []scanner.Token{*$6})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CURLY_OPEN variable '}'
            {
                $$ = $2;

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, []scanner.Token{*$1})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_var_offset:
        T_STRING
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeScalarString,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NUM_STRING
            {
                // TODO: add option to handle 64 bit integer
                tknValue := yylex.(*Parser).Ast.FileData[$1.StartPos:$1.EndPos]
                if _, err := strconv.Atoi(string(tknValue)); err == nil {
                    $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                        Type: ast.NodeTypeScalarLnumber,
                    })
                    yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))
                } else {
                    $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                        Type: ast.NodeTypeScalarString,
                    })
                    yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))
                }

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVarName, identifierNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

internal_functions_in_yacc:
        T_ISSET '(' isset_variables ')'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprIsset,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVars, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupIsset, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVarList, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EMPTY '(' variable ')'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprEmpty,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEmpty, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EMPTY '(' expr ')'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprEmpty,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEmpty, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INCLUDE expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprInclude,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INCLUDE_ONCE expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprIncludeOnce,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EVAL '(' expr ')'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprEval,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEval, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_REQUIRE expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprRequire,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_REQUIRE_ONCE expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprRequireOnce,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

isset_variables:
        isset_variable
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   isset_variables ',' isset_variable
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }

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
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprClassConstFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupConstantName, identifierNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM T_STRING
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprClassConstFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupConstantName, identifierNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_class_name_scalar:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_CLASS
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprClassConstFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupConstantName, identifierNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_name_scalar:
        class_name T_PAAMAYIM_NEKUDOTAYIM T_CLASS
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprClassConstFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupConstantName, identifierNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

%%
