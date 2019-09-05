%{
package php7

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
%type <node> extends_from
%type <node> implements_list
%type <node> interface_extends_list
%type <node> lexical_vars

%type <node> member_modifier
%type <node> use_type
%type <node> foreach_variable

%%

/////////////////////////////////////////////////////////////////////////

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

name:
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
    | T_NAMESPACE T_NS_SEPARATOR namespace_name
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
    | T_NS_SEPARATOR namespace_name
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
    |   trait_declaration_statement
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   interface_declaration_statement
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
    |   T_USE mixed_group_use_declaration ';'
            {
                $$ = $2
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupUseDeclarationList, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE use_type group_use_declaration ';'
            {
                yylex.(*Parser).SavePosition($2, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))
                yylex.(*Parser).Children($3, ast.NodeGroupUseType, $2)
                
                $$ = $3

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupUseDeclarationList, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$4})

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
    |   T_USE use_type use_declarations ';'
            {
                children := yylex.(*Parser).List.Pop()
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUseList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupUseType, $2)
                yylex.(*Parser).Children($$, ast.NodeGroupUses, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupUseDeclarationList, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONST const_list ';'
            {
                children := yylex.(*Parser).List.Pop()
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtConstList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupConsts, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_type:
        T_FUNCTION
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONST
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

group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                childrenUseDeclarations := yylex.(*Parser).List.Pop()
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // Create GroupUse Node
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtGroupUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenNameParts, []*scanner.Token{$6}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupPrefix, nameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupUseList, childrenUseDeclarations...)

                // save tokens
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSlash, $3.HiddenTokens)
                if $5 != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $5.HiddenTokens)
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, []scanner.Token{*$5})
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $6.HiddenTokens)
                } else {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $6.HiddenTokens)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                childrenUseDeclarations := yylex.(*Parser).List.Pop()
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // Create GroupUse Node
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtGroupUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $7}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupPrefix, nameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupUseList, childrenUseDeclarations...)

                // save tokens
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupUseType, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSlash, $4.HiddenTokens)
                if $6 != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $6.HiddenTokens)
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, []scanner.Token{*$6})
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $7.HiddenTokens)
                } else {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $7.HiddenTokens)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

mixed_group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                childrenUseDeclarations := yylex.(*Parser).List.Pop()
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // Create GroupUse Node
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtGroupUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenNameParts, []*scanner.Token{$6}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupPrefix, nameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupUseList, childrenUseDeclarations...)

                // save tokens
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupEnd, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSlash, $3.HiddenTokens)
                if $5 != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $5.HiddenTokens)
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, []scanner.Token{*$5})
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $6.HiddenTokens)
                } else {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $6.HiddenTokens)
                }
                
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                childrenUseDeclarations := yylex.(*Parser).List.Pop()
                childrenNameParts := yylex.(*Parser).List.Pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNameName,
                })
                yylex.(*Parser).SavePosition(nameNodeID, yylex.(*Parser).NewPosition(childrenNameParts, nil, nil))
                yylex.(*Parser).Children(nameNodeID, ast.NodeGroupParts, childrenNameParts...)

                // Create GroupUse Node
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtGroupUse,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $7}, nil))
                yylex.(*Parser).Children($$, ast.NodeGroupPrefix, nameNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupUseList, childrenUseDeclarations...)

                // save tokens
                yylex.(*Parser).MoveStartTokens(childrenNameParts[0], nameNodeID)
                yylex.(*Parser).AppendTokens(nameNodeID, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupUse, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupUse, []scanner.Token{*$1})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSlash, $4.HiddenTokens)
                if $6 != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $6.HiddenTokens)
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, []scanner.Token{*$6})
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $7.HiddenTokens)
                } else {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $7.HiddenTokens)
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

inline_use_declarations:
        inline_use_declarations ',' inline_use_declaration
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   inline_use_declaration
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unprefixed_use_declarations:
        unprefixed_use_declarations ',' unprefixed_use_declaration
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   unprefixed_use_declaration
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

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

inline_use_declaration:
        unprefixed_use_declaration
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   use_type unprefixed_use_declaration
            {
                yylex.(*Parser).Children($2, ast.NodeGroupUseType, $1)
                
                $$ = $2

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unprefixed_use_declaration:
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
;

use_declaration:
        unprefixed_use_declaration
            {
                $$ = $1

                // save tokens
                yylex.(*Parser).Ast.Foreach($1, func(e graph.Edge, n graph.Node) bool {
                    if n.Type != stxtree.NodeTypeNode {
                        return false
                    }

                    yylex.(*Parser).MoveStartTokens(e.To, $$)

                    return true
                })

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR unprefixed_use_declaration
            {
                $$ = $2;

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSlash, []scanner.Token{*$1})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

const_list:
        const_list ',' const_decl
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   const_decl
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

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
    |   trait_declaration_statement
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   interface_declaration_statement
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

statement:
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
    |   if_stmt
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   alt_if_stmt
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_WHILE '(' expr ')' while_statement
            {
                $$ = $5

                yylex.(*Parser).Children($$, ast.NodeGroupCond, $3)
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$5}))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupWhile, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DO statement T_WHILE '(' expr ')' ';'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtDo,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $7}, nil))
                
                yylex.(*Parser).Children($$, ast.NodeGroupStmt, $2)
                yylex.(*Parser).Children($$, ast.NodeGroupCond, $5)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupWhile, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $7.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$7})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOR '(' for_exprs ';' for_exprs ';' for_exprs ')' for_statement
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
    |   T_SWITCH '(' expr ')' switch_case_list
            {
                $$ = $5

                yylex.(*Parser).Children($$, ast.NodeGroupCond, $3)
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$5}))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSwitch, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BREAK optional_expr ';'
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
    |   T_CONTINUE optional_expr ';'
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
    |   T_RETURN optional_expr ';'
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
    |   T_UNSET '(' unset_variables possible_comma ')' ';' 
            {
                children := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtUnset,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $6}, nil))
                
                yylex.(*Parser).Children($$, ast.NodeGroupVars, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupUnset, $2.HiddenTokens)
                if $4 != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupVarList, $4.HiddenTokens)
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupVarList, []scanner.Token{*$4})
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupVarList, $5.HiddenTokens)
                } else {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupVarList, $5.HiddenTokens)
                }
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCloseParenthesisToken, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$6})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOREACH '(' expr T_AS foreach_variable ')' foreach_statement
            {
                yylex.(*Parser).Children($7, ast.NodeGroupExpr, $3)
                yylex.(*Parser).Children($7, ast.NodeGroupVar, $5)

                $$ = $7
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$7}))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupForeach, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $6.HiddenTokens)


                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOREACH '(' expr T_AS variable T_DOUBLE_ARROW foreach_variable ')' foreach_statement
            {
                yylex.(*Parser).Children($9, ast.NodeGroupExpr, $3)
                yylex.(*Parser).Children($9, ast.NodeGroupKey, $5)
                yylex.(*Parser).Children($9, ast.NodeGroupVar, $7)

                $$ = $9
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$9}))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupForeach, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupKey, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $8.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DECLARE '(' const_list ')' declare_statement
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
    |   T_TRY '{' inner_statement_list '}' catch_list finally_statement
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

catch_list:
        /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   catch_list T_CATCH '(' catch_name_list T_VARIABLE ')' '{' inner_statement_list '}'
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$5}, nil))
                
                varNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprVariable,
                })
                yylex.(*Parser).SavePosition(varNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$5}, nil))
                
                catchNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtCatch,
                })
                yylex.(*Parser).SavePosition(catchNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2, $9}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children(catchNodeID, ast.NodeGroupVar, varNodeID)
                yylex.(*Parser).Children(catchNodeID, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children(catchNodeID, ast.NodeGroupTypes, yylex.(*Parser).List.Pop()...)

                yylex.(*Parser).List.Add(catchNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(varNodeID, ast.TokenGroupStart, $5.HiddenTokens)
                yylex.(*Parser).AppendTokens(catchNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(catchNodeID, ast.TokenGroupCatch, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(catchNodeID, ast.TokenGroupVar, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens(catchNodeID, ast.TokenGroupCond, $7.HiddenTokens)
                yylex.(*Parser).AppendTokens(catchNodeID, ast.TokenGroupStmts, $9.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;
catch_name_list:
        name
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   catch_name_list '|' name
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

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
        T_FUNCTION returns_ref T_STRING backup_doc_comment '(' parameter_list ')' return_type '{' inner_statement_list '}'
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
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $11}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupFunctionName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupReturnType, $8)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children($$, ast.NodeGroupParams, yylex.(*Parser).List.Pop()...)


                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                if $2 != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupFunction, $2.HiddenTokens)
                } 
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $5.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupParamList, $7.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupReturnType, $9.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $11.HiddenTokens)

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

class_declaration_statement:
    class_modifiers T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
            {
                childrenStmts := yylex.(*Parser).List.Pop()
                childrenModifiers := yylex.(*Parser).List.Pop()

                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtClass,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenModifiers, []*scanner.Token{$9}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupImplements, $5)
                yylex.(*Parser).Children($$, ast.NodeGroupExtends, $4)
                yylex.(*Parser).Children($$, ast.NodeGroupClassName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupModifiers, childrenModifiers...)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, childrenStmts...)

                // save tokens
                yylex.(*Parser).MoveStartTokens(childrenModifiers[0], $$)
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupModifierList, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $7.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $9.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtClass,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $8}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupImplements, $4)
                yylex.(*Parser).Children($$, ast.NodeGroupExtends, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupClassName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $8.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_modifiers:
        class_modifier
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_modifiers class_modifier
            {
                yylex.(*Parser).List.Add($2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_modifier:
        T_ABSTRACT
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

trait_declaration_statement:
        T_TRAIT T_STRING backup_doc_comment '{' class_statement_list '}'
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtTrait,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $6}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupTraitName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $6.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

interface_declaration_statement:
        T_INTERFACE T_STRING interface_extends_list backup_doc_comment '{' class_statement_list '}'
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtInterface,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $7}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExtends, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupInterfaceName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $5.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $7.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

extends_from:
        /* empty */
            {
                $$ = 0

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXTENDS name
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

interface_extends_list:
        /* empty */
            {
                $$ = 0

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXTENDS name_list
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
    |   T_IMPLEMENTS name_list
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
    |   T_LIST '(' array_pair_list ')'
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
    |   '[' array_pair_list ']'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprShortList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupItems, yylex.(*Parser).List.Pop()...)

                // save tokensc
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupArrayPairList, $3.HiddenTokens)

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

if_stmt_without_else:
        T_IF '(' expr ')' statement
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtIf,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$5}))

                yylex.(*Parser).Children($$, ast.NodeGroupCond, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupStmt, $5)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupIf, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   if_stmt_without_else T_ELSEIF '(' expr ')' statement
            {
                elseIfNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtElseIf,
                })
                yylex.(*Parser).SavePosition(elseIfNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, []graph.NodeID{$6}))

                yylex.(*Parser).Children(elseIfNodeID, ast.NodeGroupCond, $4)
                yylex.(*Parser).Children(elseIfNodeID, ast.NodeGroupStmt, $6)
                yylex.(*Parser).Children($1, ast.NodeGroupElseIf, elseIfNodeID)

                $$ = $1
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $6}, nil, nil))

                // save tokens
                yylex.(*Parser).AppendTokens(elseIfNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(elseIfNodeID, ast.TokenGroupElseIf, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(elseIfNodeID, ast.TokenGroupExpr, $5.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

if_stmt:
        if_stmt_without_else %prec T_NOELSE
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   if_stmt_without_else T_ELSE statement
            {
                elseNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtElse,
                })
                yylex.(*Parser).SavePosition(elseNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, []graph.NodeID{$3}))
                yylex.(*Parser).SavePosition($1, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children(elseNodeID, ast.NodeGroupStmt, $3)
                yylex.(*Parser).Children($1, ast.NodeGroupElse, elseNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens(elseNodeID, ast.TokenGroupStart, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

alt_if_stmt_without_else:
        T_IF '(' expr ')' ':' inner_statement_list
            {
                children := yylex.(*Parser).List.Pop()

                stmtListNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStmtList,
                })
                yylex.(*Parser).SavePosition(stmtListNodeID, yylex.(*Parser).NewPosition(children, nil, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtAltIf,
                    Flag: ast.NodeFlagAltSyntax,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children))

                yylex.(*Parser).Children(stmtListNodeID, ast.NodeGroupStmts, children...)
                yylex.(*Parser).Children($$, ast.NodeGroupCond, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupStmt, stmtListNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupIf, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupCond, $5.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   alt_if_stmt_without_else T_ELSEIF '(' expr ')' ':' inner_statement_list
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
                yylex.(*Parser).Children(AltElseIfNodeID, ast.NodeGroupCond, $4)
                yylex.(*Parser).Children(AltElseIfNodeID, ast.NodeGroupStmt, stmtListNodeID)
                yylex.(*Parser).Children($1, ast.NodeGroupElseIf, AltElseIfNodeID)

                $$ = $1

                // save tokens
                yylex.(*Parser).AppendTokens(AltElseIfNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(AltElseIfNodeID, ast.TokenGroupElseIf, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(AltElseIfNodeID, ast.TokenGroupExpr, $5.HiddenTokens)
                yylex.(*Parser).AppendTokens(AltElseIfNodeID, ast.TokenGroupCond, $6.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

alt_if_stmt:
        alt_if_stmt_without_else T_ENDIF ';'
            {
                $$ = $1
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$3}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupAltEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   alt_if_stmt_without_else T_ELSE ':' inner_statement_list T_ENDIF ';'
            {
                children := yylex.(*Parser).List.Pop()

                stmtListNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtStmtList,
                })
                yylex.(*Parser).SavePosition(stmtListNodeID, yylex.(*Parser).NewPosition(children, nil, nil))

                AltElseNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtAltElse,
                })
                yylex.(*Parser).SavePosition(AltElseNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, children))

                yylex.(*Parser).Children(stmtListNodeID, ast.NodeGroupStmts, children...)
                yylex.(*Parser).Children(AltElseNodeID, ast.NodeGroupStmt, stmtListNodeID)
                yylex.(*Parser).Children($1, ast.NodeGroupElse, AltElseNodeID)

                $$ = $1
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$6}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens(AltElseNodeID, ast.TokenGroupStart, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(AltElseNodeID, ast.TokenGroupElse, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $5.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupAltEnd, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$6})

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
        optional_type is_reference is_variadic T_VARIABLE
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
    |   optional_type is_reference is_variadic T_VARIABLE '=' expr
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

optional_type:
        /* empty */
            {
                $$ = 0

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   type_expr
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

type_expr:
        type
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '?' type
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeNullable,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

type:
        T_ARRAY
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
    |   name
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

return_type:
        /* empty */
            {
                $$ = 0

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' type_expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtReturnType,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

argument_list:
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
    |   '(' non_empty_argument_list possible_comma ')'
            {
                children := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeArgumentList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupArguments, children...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                if $3 != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupArgumentList, $3.HiddenTokens)
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupArgumentList, []scanner.Token{*$3})
                }
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupArgumentList, $4.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_argument_list:
        argument
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_argument_list ',' argument
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

argument:
        expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeArgument,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition( []graph.NodeID{$1}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $1)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)

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
;

global_var:
        simple_variable
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_var_list:
        static_var_list ',' static_var
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_var
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_var:
        T_VARIABLE
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
                    Type: ast.NodeTypeStmtStaticVar,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupVar, varNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '=' expr
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
                    Type: ast.NodeTypeStmtStaticVar,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$3}))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupVar, varNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

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
        variable_modifiers property_list ';'
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
    |   method_modifiers T_CONST class_const_list ';'
            {
                childrenConstants := yylex.(*Parser).List.Pop()
                childrenModifiers := yylex.(*Parser).List.Pop()

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtClassConstList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(childrenModifiers, []*scanner.Token{$2, $4}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupModifiers, childrenModifiers...)
                yylex.(*Parser).Children($$, ast.NodeGroupConsts, childrenConstants...)

                // save tokens
                if len(childrenModifiers) > 0 {
                    yylex.(*Parser).MoveStartTokens(childrenModifiers[0], $$)
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupModifierList, $2.HiddenTokens)
                } else {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $2.HiddenTokens)
                }
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupConstList, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupSemiColon, []scanner.Token{*$4})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE name_list trait_adaptations
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
    |   method_modifiers T_FUNCTION returns_ref identifier backup_doc_comment '(' parameter_list ')' return_type method_body
            {
                childrenParams := yylex.(*Parser).List.Pop()
                childrenModifiers := yylex.(*Parser).List.Pop()

                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil))

                var posID graph.NodeID
                if len(childrenModifiers) == 0 {
                    posID = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, []graph.NodeID{$10})
                } else {
                    posID = yylex.(*Parser).NewPosition(childrenModifiers, nil, []graph.NodeID{$10})
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
                yylex.(*Parser).Children($$, ast.NodeGroupReturnType, $9)
                yylex.(*Parser).Children($$, ast.NodeGroupStmt, $10)
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
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupParameterList, $8.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

name_list:
        name
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   name_list ',' name
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
    |   '{' '}'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtTraitAdaptationList,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil))

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupAdaptationList, $2.HiddenTokens)

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
        trait_adaptation
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_adaptation_list trait_adaptation
            {
                yylex.(*Parser).List.Add($2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptation:
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
        absolute_trait_method_reference T_INSTEADOF name_list
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

trait_alias:
        trait_method_reference T_AS T_STRING
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtTraitUseAlias,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupRef, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupAlias, identifierNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupRef, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_method_reference T_AS reserved_non_modifiers
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil))

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtTraitUseAlias,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupRef, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupAlias, identifierNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupRef, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(identifierNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_method_reference T_AS member_modifier identifier
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

trait_method_reference:
        identifier
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
    |   absolute_trait_method_reference
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

absolute_trait_method_reference:
        name T_PAAMAYIM_NEKUDOTAYIM identifier
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

property_list:
        property_list ',' property
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   property
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

property:
        T_VARIABLE backup_doc_comment
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
                    Type: ast.NodeTypeStmtProperty,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupVar, varNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '=' expr backup_doc_comment
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
                    Type: ast.NodeTypeStmtProperty,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$3}))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupVar, varNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_const_list:
        class_const_list ',' class_const_decl
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_const_decl
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_const_decl:
        identifier '=' expr backup_doc_comment
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtConstant,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$3}))

                yylex.(*Parser).Children($$, ast.NodeGroupConstantName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

const_decl:
        T_STRING '=' expr backup_doc_comment
            {
                identifierNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeIdentifier,
                })
                yylex.(*Parser).SavePosition(identifierNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtConstant,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$3}))

                yylex.(*Parser).Children($$, ast.NodeGroupConstantName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $3)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

echo_expr_list:
        echo_expr_list ',' echo_expr
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   echo_expr
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

echo_expr:
        expr
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

for_exprs:
        /* empty */
            {
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_for_exprs
            {
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_for_exprs:
        non_empty_for_exprs ',' expr
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

anonymous_class:
        T_CLASS ctor_arguments extends_from implements_list backup_doc_comment '{' class_statement_list '}'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeStmtClass,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $8}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExtends, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupImplements, $4)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children($$, ast.NodeGroupArgumentList, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupName, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $8.HiddenTokens)

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
    |   T_NEW anonymous_class
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprNew,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

expr_without_variable:
        T_LIST '(' array_pair_list ')' '=' expr
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
    |   '[' array_pair_list ']' '=' expr
            {
                listNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprShortList,
                })
                yylex.(*Parser).SavePosition(listNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeAssignAssign,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$5}))

                yylex.(*Parser).Children(listNodeID, ast.NodeGroupItems, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children($$, ast.NodeGroupVar, listNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $5)

                // save tokens
                yylex.(*Parser).AppendTokens(listNodeID, ast.TokenGroupArrayPairList, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $4.HiddenTokens)

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
    |   variable '=' '&' expr
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
    |   variable T_INC
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
    |   T_INC variable
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
    |   variable T_DEC
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
    |   T_DEC variable
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
    |   expr T_SPACESHIP expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinarySpaceship,
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
    |   '(' expr ')'
            {
                $$ = $2;

                // save tokens
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, []scanner.Token{*$1})
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_expr
            {
                $$ = $1

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
    |   expr T_COALESCE expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeBinaryCoalesce,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupLeft, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupRight, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

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
    |   T_YIELD expr
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
    |   T_YIELD expr T_DOUBLE_ARROW expr
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
    |   T_YIELD_FROM expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprYieldFrom,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                var flag ast.NodeFlag
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                }

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprClosure,
                    Flag: flag,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $11}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClosureUse, $7)
                yylex.(*Parser).Children($$, ast.NodeGroupReturnType, $8)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children($$, ast.NodeGroupParams, yylex.(*Parser).List.Pop()...)
                
                // // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                if $2 == nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupFunction, $4.HiddenTokens)
                } else {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupFunction, $2.HiddenTokens)
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupAmpersand, $4.HiddenTokens)
                }
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupParameterList, $6.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupReturnType, $9.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $11.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                flag := ast.NodeFlagStatic
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                }

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprClosure,
                    Flag: flag,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $12}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClosureUse, $8)
                yylex.(*Parser).Children($$, ast.NodeGroupReturnType, $9)
                yylex.(*Parser).Children($$, ast.NodeGroupStmts, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children($$, ast.NodeGroupParams, yylex.(*Parser).List.Pop()...)
                
                // // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStatic, $2.HiddenTokens)
                if $3 == nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupFunction, $5.HiddenTokens)
                } else {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupFunction, $3.HiddenTokens)
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupAmpersand, $5.HiddenTokens)
                }
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupParameterList, $7.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupReturnType, $10.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStmts, $12.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

backup_doc_comment:;

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
        lexical_var_list ',' lexical_var
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   lexical_var
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

lexical_var:
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

                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprReference,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil))

                yylex.(*Parser).Children(varNodeID, ast.NodeGroupVarName, identifierNodeID)
                yylex.(*Parser).Children($$, ast.NodeGroupVar, varNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens(varNodeID, ast.TokenGroupStart, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

function_call:
        name argument_list
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
    |   class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
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
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
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
    |   callable_expr argument_list
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
    |   name
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_name_reference:
        class_name
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_variable
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

exit_expr:
        /* empty */
            {
                $$ = 0

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' optional_expr ')'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprExit,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupExpr, $2)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExit, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExit, []scanner.Token{*$1})
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
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
    |   argument_list
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

dereferencable_scalar:
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
;

scalar:
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
    |   dereferencable_scalar
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   constant
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

constant:
        name
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprConstFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition( []graph.NodeID{$1}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupConstant, $1)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM identifier
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
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM identifier
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

expr:
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

optional_expr:
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

variable_class_name:
        dereferencable
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

dereferencable:
        variable
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' expr ')'
            {
                $$ = $2;

                // save tokens
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, []scanner.Token{*$1})
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   dereferencable_scalar
            {
                $$ = $1;

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

callable_expr:
        callable_variable
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' expr ')'
            {
                $$ = $2;

                // save tokens
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, []scanner.Token{*$1})
                yylex.(*Parser).PrependTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupEnd, []scanner.Token{*$3})

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   dereferencable_scalar
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

callable_variable:
        simple_variable
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   dereferencable '[' optional_expr ']'
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
    |   constant '[' optional_expr ']'
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
    |   dereferencable '{' expr '}'
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
    |   dereferencable T_OBJECT_OPERATOR property_name argument_list
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprMethodCall,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $4}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupMethod, $3)
                yylex.(*Parser).Children($$, ast.NodeGroupArgumentList, $4)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   function_call
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

variable:
        callable_variable
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_member
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   dereferencable T_OBJECT_OPERATOR property_name
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprPropertyFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVar, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupProperty, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

simple_variable:
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
    |   '$' simple_variable
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
;

static_member:
        class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
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
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
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

new_variable:
        simple_variable
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_variable '[' optional_expr ']'
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
    |   new_variable '{' expr '}'
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
    |   new_variable T_OBJECT_OPERATOR property_name
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprPropertyFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupProperty, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprStaticPropertyFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupProperty, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_variable T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprStaticPropertyFetch,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupClass, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupProperty, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVar, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

member_name:
        identifier
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
    |   simple_variable
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

property_name:
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
    |   simple_variable
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

array_pair_list:
        non_empty_array_pair_list
            {
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

possible_array_pair:
        /* empty */
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   array_pair
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_array_pair_list:
        non_empty_array_pair_list ',' possible_array_pair
            {
                prevNodeID := yylex.(*Parser).List.Last()
                yylex.(*Parser).List.Add($3)

                // save tokens
                yylex.(*Parser).AppendTokens(prevNodeID, ast.TokenGroupEnd, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   possible_array_pair
            {
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

array_pair:
        expr T_DOUBLE_ARROW expr
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $3}, nil, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupKey, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupVal, $3)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
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
    |   expr T_DOUBLE_ARROW '&' variable
            {
                refNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprReference,
                })
                yylex.(*Parser).SavePosition(refNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, []graph.NodeID{$4}))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1, $4}, nil, nil))

                yylex.(*Parser).Children(refNodeID, ast.NodeGroupVar, $4)
                yylex.(*Parser).Children($$, ast.NodeGroupKey, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupVal, refNodeID)

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(refNodeID, ast.TokenGroupStart, $3.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' variable
            {
                refNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprReference,
                })
                yylex.(*Parser).SavePosition(refNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, []graph.NodeID{$2}))

                yylex.(*Parser).Children(refNodeID, ast.NodeGroupVar, $2)
                yylex.(*Parser).Children($$, ast.NodeGroupVal, refNodeID)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_DOUBLE_ARROW T_LIST '(' array_pair_list ')'
            {
                listNodeID :=  yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprList,
                })
                yylex.(*Parser).SavePosition(listNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3, $6}, nil))
                
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprArrayItem,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition([]graph.NodeID{$1}, []*scanner.Token{$6}, nil))

                yylex.(*Parser).Children(listNodeID, ast.NodeGroupItems, yylex.(*Parser).List.Pop()...)
                yylex.(*Parser).Children($$, ast.NodeGroupKey, $1)
                yylex.(*Parser).Children($$, ast.NodeGroupVal, listNodeID)

                // TODO: Cannot use list() as standalone expression

                // save tokens
                yylex.(*Parser).MoveStartTokens($1, $$)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupExpr, $2.HiddenTokens)
                yylex.(*Parser).AppendTokens(listNodeID, ast.TokenGroupStart, $3.HiddenTokens)
                yylex.(*Parser).AppendTokens(listNodeID, ast.TokenGroupList, $4.HiddenTokens)
                yylex.(*Parser).AppendTokens(listNodeID, ast.TokenGroupArrayPairList, $6.HiddenTokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LIST '(' array_pair_list ')'
            {
                listNodeID :=  yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
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
    |   '-' T_NUM_STRING
            {
                tknValue := yylex.(*Parser).Ast.FileData[$2.StartPos:$2.EndPos]
                if _, err := strconv.Atoi(string(tknValue)); err == nil {
                    lnumberNodeID := yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                        Type: ast.NodeTypeScalarLnumber,
                    })
                    yylex.(*Parser).SavePosition(lnumberNodeID, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil))
                    
                    $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                        Type: ast.NodeTypeExprUnaryMinus,
                    })
                    yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil))

                    yylex.(*Parser).Children($$, ast.NodeGroupExpr, lnumberNodeID)
                } else {
                    $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                        Type: ast.NodeTypeScalarString,
                    })
                    yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil))
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
        T_ISSET '(' isset_variables possible_comma ')'
            {
                $$ = yylex.(*Parser).Ast.NewNode(ast.SimpleNode{
                    Type: ast.NodeTypeExprIsset,
                })
                yylex.(*Parser).SavePosition($$, yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $5}, nil))

                yylex.(*Parser).Children($$, ast.NodeGroupVars, yylex.(*Parser).List.Pop()...)

                // save tokens
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupStart, $1.HiddenTokens)
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupIsset, $2.HiddenTokens)
                if $4 != nil {
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupVarList, $4.HiddenTokens)
                    yylex.(*Parser).AppendTokens($$, ast.TokenGroupVarList, []scanner.Token{*$4})
                }
                yylex.(*Parser).AppendTokens($$, ast.TokenGroupVarList, $5.HiddenTokens)

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
;

isset_variable:
        expr
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

/////////////////////////////////////////////////////////////////////////

%%
