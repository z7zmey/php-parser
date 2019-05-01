%{
package php7

import (
    "bytes"
    "strconv"

    "github.com/z7zmey/php-parser/ast"
    "github.com/z7zmey/php-parser/ast/linear"
    "github.com/z7zmey/php-parser/scanner"
)

%}

%union{
    node  linear.NodeID
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
                children := yylex.(*Parser).list.pop()
                nodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeRoot,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(children),
                })
                yylex.(*Parser).ast.Children(0, nodeID, ast.EdgeTypeStmts, children...)

                yylex.(*Parser).ast.RootNode = nodeID

                // save comments
                // yylex.(*Parser).setFreeFloating(yylex.(*Parser).rootNode, freefloating.End, yylex.(*Parser).currentToken.FreeFloating)

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
                    yylex.(*Parser).list.add($2)
                }

                // if inlineHtmlNode, ok := $2.(*stmt.InlineHtml); ok && len($1) > 0 {
                //     prevNode := lastNode($1)
                //     yylex.(*Parser).splitSemiColonAndPhpCloseTag(inlineHtmlNode, prevNode)
                // }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                yylex.(*Parser).list.push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

namespace_name:
        T_STRING
            {
                nodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeNameNamePart,
                    Pos: yylex.(*Parser).ast.NewTokenPosition($1),
                })

                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add(nodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating(namePart, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_NS_SEPARATOR T_STRING
            {
                nodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeNameNamePart,
                    Pos: yylex.(*Parser).ast.NewTokenPosition($3),
                })
                yylex.(*Parser).list.add(nodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(namePart, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

name:
        namespace_name
            {
                children := yylex.(*Parser).list.pop()
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeNameName,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(children),
                })
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeParts, children...)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    | T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                children := yylex.(*Parser).list.pop()
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeNameRelative,
                    Pos:  yylex.(*Parser).ast.NewTokenNodeListPosition($1, children),
                })
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeParts, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Namespace, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    | T_NS_SEPARATOR namespace_name
            {
                children := yylex.(*Parser).list.pop()
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeNameFullyQualified,
                    Pos:  yylex.(*Parser).ast.NewTokenNodeListPosition($1, children),
                })
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeParts, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtHaltCompiler,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.HaltCompiller, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.OpenParenthesisToken, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.CloseParenthesisToken, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE namespace_name ';'
            {
                children := yylex.(*Parser).list.pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeNameName,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(children),
                })
                yylex.(*Parser).ast.Children(0, nameNodeID, ast.EdgeTypeParts, children...)

                // Create Namespace Node
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtNamespace,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeNamespaceName, nameNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).MoveFreeFloating($2[0], name)
                // yylex.(*Parser).setFreeFloating(name, freefloating.End, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE namespace_name '{' top_statement_list '}'
            {
                childrenStmts := yylex.(*Parser).list.pop()
                childrenNameParts := yylex.(*Parser).list.pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeNameName,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(childrenNameParts),
                })
                yylex.(*Parser).ast.Children(0, nameNodeID, ast.EdgeTypeParts, childrenNameParts...)

                // Create Namespace Node
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtNamespace,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $5),
                })
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeNamespaceName, nameNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeStmts, childrenStmts...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).MoveFreeFloating($2[0], name)
                // yylex.(*Parser).setFreeFloating(name, freefloating.End, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $5.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE '{' top_statement_list '}'
            {
                children := yylex.(*Parser).list.pop()
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtNamespace,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeStmts, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Namespace, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE mixed_group_use_declaration ';'
            {
                node := yylex.(*Parser).ast.Nodes.Get($2)
                node.Pos = yylex.(*Parser).ast.NewTokensPosition($1, $3)
                yylex.(*Parser).ast.Nodes.Save($2, node)

                $$ = $2

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.UseDeclarationList, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE use_type group_use_declaration ';'
            {
                node := yylex.(*Parser).ast.Nodes.Get($2)
                node.Pos = yylex.(*Parser).ast.NewTokensPosition($1, $4)
                yylex.(*Parser).ast.Nodes.Save($2, node)

                yylex.(*Parser).ast.Children(0, $3, ast.EdgeTypeUseType, $2)
                
                $$ = $3

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.UseDeclarationList, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE use_declarations ';'
            {
                children := yylex.(*Parser).list.pop()
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtUseList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeUses, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.UseDeclarationList, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE use_type use_declarations ';'
            {
                children := yylex.(*Parser).list.pop()
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtUseList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeUseType, $2)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeUses, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.UseDeclarationList, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONST const_list ';'
            {
                children := yylex.(*Parser).list.pop()
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtConstList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeConsts, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_type:
        T_FUNCTION
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONST
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                childrenUseDeclarations := yylex.(*Parser).list.pop()
                childrenNameParts := yylex.(*Parser).list.pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeNameName,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(childrenNameParts),
                })
                yylex.(*Parser).ast.Children(0, nameNodeID, ast.EdgeTypeParts, childrenNameParts...)

                // Create GroupUse Node
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtGroupUse,
                    Pos:  yylex.(*Parser).ast.NewNodeListTokenPosition(childrenNameParts, $6),
                })
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypePrefix, nameNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeUseList, childrenUseDeclarations...)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1[0], name)
                // yylex.(*Parser).setFreeFloating(name, freefloating.End, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Slash, $3.FreeFloating)
                // if $5 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, append($5.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($5), $6.FreeFloating...)...))
                // } else {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $6.FreeFloating)
                // }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                childrenUseDeclarations := yylex.(*Parser).list.pop()
                childrenNameParts := yylex.(*Parser).list.pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeNameName,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(childrenNameParts),
                })
                yylex.(*Parser).ast.Children(0, nameNodeID, ast.EdgeTypeParts, childrenNameParts...)

                // Create GroupUse Node
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtGroupUse,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $7),
                })
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypePrefix, nameNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeUseList, childrenUseDeclarations...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.UseType, $1.FreeFloating)
                // yylex.(*Parser).MoveFreeFloating($2[0], name)
                // yylex.(*Parser).setFreeFloating(name, freefloating.End, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Slash, $4.FreeFloating)
                // if $6 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, append($6.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($6), $7.FreeFloating...)...))
                // } else {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $7.FreeFloating)
                // }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

mixed_group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                childrenUseDeclarations := yylex.(*Parser).list.pop()
                childrenNameParts := yylex.(*Parser).list.pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeNameName,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(childrenNameParts),
                })
                yylex.(*Parser).ast.Children(0, nameNodeID, ast.EdgeTypeParts, childrenNameParts...)

                // Create GroupUse Node
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtGroupUse,
                    Pos:  yylex.(*Parser).ast.NewNodeListTokenPosition(childrenNameParts, $6),
                })
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypePrefix, nameNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeUseList, childrenUseDeclarations...)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1[0], name)
                // yylex.(*Parser).setFreeFloating(name, freefloating.End, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Slash, $3.FreeFloating)
                // if $5 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, append($5.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($5), $6.FreeFloating...)...))
                // } else {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $6.FreeFloating)
                // }
                
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                childrenUseDeclarations := yylex.(*Parser).list.pop()
                childrenNameParts := yylex.(*Parser).list.pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeNameName,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(childrenNameParts),
                })
                yylex.(*Parser).ast.Children(0, nameNodeID, ast.EdgeTypeParts, childrenNameParts...)

                // Create GroupUse Node
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtGroupUse,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $7),
                })
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypePrefix, nameNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeUseList, childrenUseDeclarations...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Use, append($1.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($1)...))
                // yylex.(*Parser).MoveFreeFloating($2[0], name)
                // yylex.(*Parser).setFreeFloating(name, freefloating.End, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Slash, $4.FreeFloating)
                // if $6 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, append($6.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($6), $7.FreeFloating...)...))
                // } else {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $7.FreeFloating)
                // }

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
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   inline_use_declaration
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unprefixed_use_declarations:
        unprefixed_use_declarations ',' unprefixed_use_declaration
            {
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   unprefixed_use_declaration
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_declarations:
        use_declarations ',' use_declaration
            {
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   use_declaration
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

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
                yylex.(*Parser).ast.Children(0, $2, ast.EdgeTypeUseType, $1)
                
                $$ = $2

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unprefixed_use_declaration:
        namespace_name
            {
                childrenNameParts := yylex.(*Parser).list.pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeNameName,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(childrenNameParts),
                })
                yylex.(*Parser).ast.Children(0, nameNodeID, ast.EdgeTypeParts, childrenNameParts...)

                // Create Use Node
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtUse,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(childrenNameParts),
                })
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeUse, nameNodeID)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1[0], name)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_AS T_STRING
            {
                childrenNameParts := yylex.(*Parser).list.pop()

                // Create Name Node
                nameNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeNameName,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(childrenNameParts),
                })
                yylex.(*Parser).ast.Children(0, nameNodeID, ast.EdgeTypeParts, childrenNameParts...)

                // create Alias Node
                aliasNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($3),
                })

                // Create Use Node
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtUse,
                    Pos:  yylex.(*Parser).ast.NewNodeListTokenPosition(childrenNameParts, $3),
                })
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeUse, nameNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeAlias, aliasNodeID)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1[0], name)
                // yylex.(*Parser).setFreeFloating(name, freefloating.End, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(alias, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_declaration:
        unprefixed_use_declaration
            {
                $$ = $1

                // save coments
                // yylex.(*Parser).MoveFreeFloating($1.(*stmt.Use).Use, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR unprefixed_use_declaration
            {
                $$ = $2;

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Slash, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

const_list:
        const_list ',' const_decl
            {
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   const_decl
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

inner_statement_list:
        inner_statement_list inner_statement
            {
                // if inlineHtmlNode, ok := $2.(*stmt.InlineHtml); ok && len($1) > 0 {
                //     prevNode := lastNode($1)
                //     yylex.(*Parser).splitSemiColonAndPhpCloseTag(inlineHtmlNode, prevNode)
                // }

                if $2 != 0 {
                    yylex.(*Parser).list.add($2)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                yylex.(*Parser).list.push()

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtHaltCompiler,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.HaltCompiller, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.OpenParenthesisToken, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.CloseParenthesisToken, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }

statement:
        '{' inner_statement_list '}'
            {
                children := yylex.(*Parser).list.pop()

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })
                
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeStmts, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)

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
                node := yylex.(*Parser).ast.Nodes.Get($5)
                node.Pos = yylex.(*Parser).ast.NewTokenNodePosition($1, $5)
                yylex.(*Parser).ast.Nodes.Save($5, node)

                yylex.(*Parser).ast.Children(0, $5, ast.EdgeTypeCond, $3)

                $$ = $5

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.While, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DO statement T_WHILE '(' expr ')' ';'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtDo,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $7),
                })
                
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeStmt, $2)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeCond, $5)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.While, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $6.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $7.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($7))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOR '(' for_exprs ';' for_exprs ';' for_exprs ')' for_statement
            {
                node := yylex.(*Parser).ast.Nodes.Get($9)
                node.Pos = yylex.(*Parser).ast.NewTokenNodePosition($1, $9)
                yylex.(*Parser).ast.Nodes.Save($9, node)

                prevNodeID := yylex.(*Parser).ast.Children(0, $9, ast.EdgeTypeLoop, yylex.(*Parser).list.pop()...)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $9, ast.EdgeTypeCond, yylex.(*Parser).list.pop()...)
                yylex.(*Parser).ast.Children(prevNodeID, $9, ast.EdgeTypeInit, yylex.(*Parser).list.pop()...)

                $$ = $9

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.For, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.InitExpr, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.CondExpr, $6.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.IncExpr, $8.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_SWITCH '(' expr ')' switch_case_list
            {
                node := yylex.(*Parser).ast.Nodes.Get($5)
                node.Pos = yylex.(*Parser).ast.NewTokenNodePosition($1, $5)
                yylex.(*Parser).ast.Nodes.Save($5, node)

                yylex.(*Parser).ast.Children(0, $5, ast.EdgeTypeCond, $3)

                $$ = $5

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Switch, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BREAK optional_expr ';'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtBreak,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })
                
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONTINUE optional_expr ';'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtContinue,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })
                
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_RETURN optional_expr ';'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtReturn,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })
                
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_GLOBAL global_var_list ';'
            {
                children := yylex.(*Parser).list.pop()

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtGlobal,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })
                
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVars, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.VarList, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC static_var_list ';'
            {
                children := yylex.(*Parser).list.pop()

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtStatic,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })
                
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVars, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.VarList, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ECHO echo_expr_list ';'
            {
                children := yylex.(*Parser).list.pop()

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtEcho,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })
                
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExprs, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Echo, yylex.(*Parser).GetFreeFloatingToken($1))
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INLINE_HTML
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtInlineHtml,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr ';'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtExpression,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVars, $1)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_UNSET '(' unset_variables possible_comma ')' ';' 
            {
                children := yylex.(*Parser).list.pop()

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtUnset,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $6),
                })
                
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVars, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Unset, $2.FreeFloating)
                // if $4 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.VarList, append($4.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($4), $5.FreeFloating...)...))
                // } else {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.VarList, $5.FreeFloating)
                // }
                // yylex.(*Parser).setFreeFloating($$, freefloating.CloseParenthesisToken, $6.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($6))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOREACH '(' expr T_AS foreach_variable ')' foreach_statement
            {
                node := yylex.(*Parser).ast.Nodes.Get($7)
                node.Pos = yylex.(*Parser).ast.NewTokenNodePosition($1, $7)
                yylex.(*Parser).ast.Nodes.Save($7, node)

                prevNodeID := yylex.(*Parser).ast.Children(0, $7, ast.EdgeTypeExpr, $3)
                yylex.(*Parser).ast.Children(prevNodeID, $7, ast.EdgeTypeVar, $5)

                $$ = $7

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Foreach, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $6.FreeFloating)


                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOREACH '(' expr T_AS variable T_DOUBLE_ARROW foreach_variable ')' foreach_statement
            {
                node := yylex.(*Parser).ast.Nodes.Get($9)
                node.Pos = yylex.(*Parser).ast.NewTokenNodePosition($1, $9)
                yylex.(*Parser).ast.Nodes.Save($9, node)

                prevNodeID := yylex.(*Parser).ast.Children(0, $9, ast.EdgeTypeExpr, $3)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $9, ast.EdgeTypeKey, $5)
                yylex.(*Parser).ast.Children(prevNodeID, $9, ast.EdgeTypeVar, $7)

                $$ = $9

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Foreach, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Key, $6.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $8.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DECLARE '(' const_list ')' declare_statement
            {
                children := yylex.(*Parser).list.pop()

                node := yylex.(*Parser).ast.Nodes.Get($5)
                node.Pos = yylex.(*Parser).ast.NewTokenNodePosition($1, $5)
                yylex.(*Parser).ast.Nodes.Save($5, node)

                yylex.(*Parser).ast.Children(0, $5, ast.EdgeTypeConsts, children...)

                $$ = $5


                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Declare, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.ConstList, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ';'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtNop,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRY '{' inner_statement_list '}' catch_list finally_statement
            {
                childrenCatches := yylex.(*Parser).list.pop()
                childrenStmts := yylex.(*Parser).list.pop()

                var posID linear.PositionID
                if $6 == 0 {
                    posID = yylex.(*Parser).ast.NewTokenNodeListPosition($1, childrenCatches)
                } else {
                    posID = yylex.(*Parser).ast.NewTokenNodePosition($1, $6)
                }

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtTry,
                    Pos:  posID,
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeConsts, $6)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeCatches, childrenCatches...)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeStmts, childrenStmts...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Try, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_THROW expr ';'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtThrow,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_GOTO T_STRING ';'
            {
                LableNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($2),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtGoto,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLabel, LableNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating(label, freefloating.Start, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Label, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STRING ':'
            {
                LableNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtLabel,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLabelName, LableNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Label, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }

catch_list:
        /* empty */
            {
                yylex.(*Parser).list.push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   catch_list T_CATCH '(' catch_name_list T_VARIABLE ')' '{' inner_statement_list '}'
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($5),
                })
                
                varNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($5),
                })
                
                catchNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtCatch,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($2, $9),
                })

                yylex.(*Parser).ast.Children(0, varNodeID, ast.EdgeTypeVarName, identifierNodeID)
                prevNodeID := yylex.(*Parser).ast.Children(0, catchNodeID, ast.EdgeTypeVar, varNodeID)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, catchNodeID, ast.EdgeTypeStmts, yylex.(*Parser).list.pop()...)
                yylex.(*Parser).ast.Children(prevNodeID, catchNodeID, ast.EdgeTypeTypes, yylex.(*Parser).list.pop()...)

                yylex.(*Parser).list.add(catchNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating(catch, freefloating.Start, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(catch, freefloating.Catch, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating(variable, freefloating.Start, $5.FreeFloating)
                // yylex.(*Parser).addDollarToken(variable)
                // yylex.(*Parser).setFreeFloating(catch, freefloating.Var, $6.FreeFloating)
                // yylex.(*Parser).setFreeFloating(catch, freefloating.Cond, $7.FreeFloating)
                // yylex.(*Parser).setFreeFloating(catch, freefloating.Stmts, $9.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;
catch_name_list:
        name
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   catch_name_list '|' name
            {
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtFinally,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeStmts, yylex.(*Parser).list.pop()...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Finally, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unset_variables:
        unset_variable
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   unset_variables ',' unset_variable
            {
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

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
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($3),
                })

                var flag ast.NodeFlag
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                }
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtFunction,
                    Flag: flag,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $11),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeFunctionName, identifierNodeID)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeReturnType, $8)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeStmts, yylex.(*Parser).list.pop()...)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeParams, yylex.(*Parser).list.pop()...)


                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // if $2 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Function, $2.FreeFloating)
                //     yylex.(*Parser).setFreeFloating(name, freefloating.Start, $3.FreeFloating)
                // } else {
                //     yylex.(*Parser).setFreeFloating(name, freefloating.Start, $3.FreeFloating)
                // }
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $5.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.ParamList, $7.FreeFloating)
                // if $8 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Params, (*$8.GetFreeFloating())[freefloating.Colon]); delete((*$8.GetFreeFloating()), freefloating.Colon)
                // }
                // yylex.(*Parser).setFreeFloating($$, freefloating.ReturnType, $9.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $11.FreeFloating)

                // normalize
                // if $8 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Params, (*$$.GetFreeFloating())[freefloating.ReturnType]); delete((*$$.GetFreeFloating()), freefloating.ReturnType)
                // }

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
                childrenStmts := yylex.(*Parser).list.pop()
                childrenModifiers := yylex.(*Parser).list.pop()

                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($3),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtClass,
                    Pos:  yylex.(*Parser).ast.NewNodeListTokenPosition(childrenModifiers, $9),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeImplements, $5)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExtends, $4)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeClassName, identifierNodeID)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeModifiers, childrenModifiers...)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeStmts, childrenStmts...)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1[0], $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.ModifierList, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(name, freefloating.Start, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $7.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $9.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($2),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtClass,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $8),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeImplements, $4)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExtends, $3)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeClassName, identifierNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeStmts, yylex.(*Parser).list.pop()...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating(name, freefloating.Start, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $6.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $8.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_modifiers:
        class_modifier
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_modifiers class_modifier
            {
                yylex.(*Parser).list.add($2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_modifier:
        T_ABSTRACT
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINAL
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_declaration_statement:
        T_TRAIT T_STRING backup_doc_comment '{' class_statement_list '}'
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($2),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtTrait,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $6),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeTraitName, identifierNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeStmts, yylex.(*Parser).list.pop()...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating(name, freefloating.Start, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $6.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

interface_declaration_statement:
        T_INTERFACE T_STRING interface_extends_list backup_doc_comment '{' class_statement_list '}'
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($2),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtInterface,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $7),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExtends, $3)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeInterfaceName, identifierNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeStmts, yylex.(*Parser).list.pop()...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating(name, freefloating.Start, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $5.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $7.FreeFloating)

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtClassExtends,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeClassName, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

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
                children := yylex.(*Parser).list.pop()

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtInterfaceExtends,
                    Pos:  yylex.(*Parser).ast.NewTokenNodeListPosition($1, children),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeInterfaceNames, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

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
                children := yylex.(*Parser).list.pop()

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtClassImplements,
                    Pos:  yylex.(*Parser).ast.NewTokenNodeListPosition($1, children),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeInterfaceNames, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprReference,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LIST '(' array_pair_list ')'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeItems, yylex.(*Parser).list.pop()...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.List, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.ArrayPairList, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' array_pair_list ']'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprShortList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeItems, yylex.(*Parser).list.pop()...)

                // save commentsc
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.ArrayPairList, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

for_statement:
        statement
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtFor,
                    Pos:  yylex.(*Parser).ast.NewNodePosition($1),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeStmt, $1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDFOR ';'
            {
                children := yylex.(*Parser).list.pop()

                stmtListNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(children),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtAltFor,
                    Flag: ast.NodeFlagAltSyntax,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, stmtListNodeID, ast.EdgeTypeStmts, children...)
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeStmt, stmtListNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

foreach_statement:
        statement
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtForeach,
                    Pos:  yylex.(*Parser).ast.NewNodePosition($1),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeStmt, $1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDFOREACH ';'
            {
                children := yylex.(*Parser).list.pop()

                stmtListNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(children),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtAltForeach,
                    Flag: ast.NodeFlagAltSyntax,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, stmtListNodeID, ast.EdgeTypeStmts, children...)
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeStmt, stmtListNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

declare_statement:
        statement
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtDeclare,
                    Pos:  yylex.(*Parser).ast.NewNodePosition($1),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeStmt, $1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDDECLARE ';'
            {
                children := yylex.(*Parser).list.pop()

                stmtListNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(children),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtDeclare,
                    Flag: ast.NodeFlagAltSyntax,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, stmtListNodeID, ast.EdgeTypeStmts, children...)
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeStmt, stmtListNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

switch_case_list:
        '{' case_list '}'
            {
                children := yylex.(*Parser).list.pop()

                caseListNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtCaseList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtSwitch,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, caseListNodeID, ast.EdgeTypeCases, children...)
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeCaseList, caseListNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating(caseList, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating(caseList, freefloating.CaseListEnd, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' ';' case_list '}'
            {
                children := yylex.(*Parser).list.pop()

                caseListNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtCaseList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtSwitch,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, caseListNodeID, ast.EdgeTypeCases, children...)
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeCaseList, caseListNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating(caseList, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating(caseList, freefloating.CaseListStart, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                // yylex.(*Parser).setFreeFloating(caseList, freefloating.CaseListEnd, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' case_list T_ENDSWITCH ';'
            {
                children := yylex.(*Parser).list.pop()

                caseListNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtCaseList,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(children),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtAltSwitch,
                    Flag: ast.NodeFlagAltSyntax,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, caseListNodeID, ast.EdgeTypeCases, children...)
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeCaseList, caseListNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating(caseList, freefloating.CaseListEnd, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' ';' case_list T_ENDSWITCH ';'
            {
                children := yylex.(*Parser).list.pop()

                caseListNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtCaseList,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(children),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtAltSwitch,
                    Flag: ast.NodeFlagAltSyntax,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $5),
                })

                yylex.(*Parser).ast.Children(0, caseListNodeID, ast.EdgeTypeCases, children...)
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeCaseList, caseListNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating(caseList, freefloating.CaseListStart, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                // yylex.(*Parser).setFreeFloating(caseList, freefloating.CaseListEnd, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $5.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($5))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

case_list:
        /* empty */
            {
                yylex.(*Parser).list.push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   case_list T_CASE expr case_separator inner_statement_list
            {
                children := yylex.(*Parser).list.pop()

                caseNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtCase,
                    Pos:  yylex.(*Parser).ast.NewTokenNodeListPosition($2, children),
                })

                yylex.(*Parser).ast.Children(0, caseNodeID, ast.EdgeTypeStmts, children...)

                yylex.(*Parser).list.add(caseNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating(_case, freefloating.Start, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(_case, freefloating.Expr, append($4.FreeFloating))
                // yylex.(*Parser).setFreeFloating(_case, freefloating.CaseSeparator, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   case_list T_DEFAULT case_separator inner_statement_list
            {
                children := yylex.(*Parser).list.pop()

                defaultNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtDefault,
                    Pos:  yylex.(*Parser).ast.NewTokenNodeListPosition($2, children),
                })

                yylex.(*Parser).ast.Children(0, defaultNodeID, ast.EdgeTypeStmts, children...)

                yylex.(*Parser).list.add(defaultNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating(_default, freefloating.Start, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(_default, freefloating.Default, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating(_default, freefloating.CaseSeparator, yylex.(*Parser).GetFreeFloatingToken($3))

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtWhile,
                    Pos:  yylex.(*Parser).ast.NewNodePosition($1),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeStmt, $1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDWHILE ';'
            {
                children := yylex.(*Parser).list.pop()

                stmtListNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(children),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtAltWhile,
                    Flag: ast.NodeFlagAltSyntax,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, stmtListNodeID, ast.EdgeTypeStmts, children...)
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeStmt, stmtListNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

if_stmt_without_else:
        T_IF '(' expr ')' statement
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtIf,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $5),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeCond, $3)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeStmt, $5)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.If, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   if_stmt_without_else T_ELSEIF '(' expr ')' statement
            {
                elseIfNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtElseIf,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($2, $6),
                })

                node := yylex.(*Parser).ast.Nodes.Get($1)
                node.Pos = yylex.(*Parser).ast.NewNodesPosition($1, $6)
                yylex.(*Parser).ast.Nodes.Save($1, node)

                prevNodeID := yylex.(*Parser).ast.Children(0, elseIfNodeID, ast.EdgeTypeCond, $4)
                yylex.(*Parser).ast.Children(prevNodeID, elseIfNodeID, ast.EdgeTypeStmt, $6)
                yylex.(*Parser).ast.Children(0, $1, ast.EdgeTypeElseIf, elseIfNodeID)

                $$ = $1

                // save comments
                // yylex.(*Parser).setFreeFloating(_elseIf, freefloating.Start, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(_elseIf, freefloating.ElseIf, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating(_elseIf, freefloating.Expr, $5.FreeFloating)

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
                elseNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtElse,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($2, $3),
                })

                node := yylex.(*Parser).ast.Nodes.Get($1)
                node.Pos = yylex.(*Parser).ast.NewNodesPosition($1, $3)
                yylex.(*Parser).ast.Nodes.Save($1, node)

                yylex.(*Parser).ast.Children(0, elseNodeID, ast.EdgeTypeStmt, $3)
                yylex.(*Parser).ast.Children(0, $1, ast.EdgeTypeElse, elseNodeID)


                // save comments
                // yylex.(*Parser).setFreeFloating(_else, freefloating.Start, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

alt_if_stmt_without_else:
        T_IF '(' expr ')' ':' inner_statement_list
            {
                children := yylex.(*Parser).list.pop()

                stmtListNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(children),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtAltIf,
                    Flag: ast.NodeFlagAltSyntax,
                    Pos:  yylex.(*Parser).ast.NewTokenNodeListPosition($1, children),
                })

                yylex.(*Parser).ast.Children(0, stmtListNodeID, ast.EdgeTypeStmts, children...)
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeCond, $3)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeStmt, stmtListNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.If, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $5.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   alt_if_stmt_without_else T_ELSEIF '(' expr ')' ':' inner_statement_list
            {
                children := yylex.(*Parser).list.pop()

                stmtListNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(children),
                })

                AltElseIfNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtAltElseIf,
                    Pos:  yylex.(*Parser).ast.NewTokenNodeListPosition($2, children),
                })

                yylex.(*Parser).ast.Children(0, stmtListNodeID, ast.EdgeTypeStmts, children...)
                prevNodeID := yylex.(*Parser).ast.Children(0, AltElseIfNodeID, ast.EdgeTypeCond, $4)
                yylex.(*Parser).ast.Children(prevNodeID, AltElseIfNodeID, ast.EdgeTypeStmt, stmtListNodeID)
                yylex.(*Parser).ast.Children(0, $1, ast.EdgeTypeElseIf, AltElseIfNodeID)

                $$ = $1

                // save comments
                // yylex.(*Parser).setFreeFloating(_elseIf, freefloating.Start, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(_elseIf, freefloating.ElseIf, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating(_elseIf, freefloating.Expr, $5.FreeFloating)
                // yylex.(*Parser).setFreeFloating(_elseIf, freefloating.Cond, $6.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

alt_if_stmt:
        alt_if_stmt_without_else T_ENDIF ';'
            {
                node := yylex.(*Parser).ast.Nodes.Get($1)
                node.Pos = yylex.(*Parser).ast.NewNodeTokenPosition($1, $3)
                yylex.(*Parser).ast.Nodes.Save($1, node)

                $$ = $1

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   alt_if_stmt_without_else T_ELSE ':' inner_statement_list T_ENDIF ';'
            {
                children := yylex.(*Parser).list.pop()

                stmtListNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Pos:  yylex.(*Parser).ast.NewNodeListPosition(children),
                })

                AltElseNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtAltElse,
                    Pos:  yylex.(*Parser).ast.NewTokenNodeListPosition($2, children),
                })

                node := yylex.(*Parser).ast.Nodes.Get($1)
                node.Pos = yylex.(*Parser).ast.NewNodeTokenPosition($1, $6)
                yylex.(*Parser).ast.Nodes.Save($1, node)

                yylex.(*Parser).ast.Children(0, stmtListNodeID, ast.EdgeTypeStmts, children...)
                yylex.(*Parser).ast.Children(0, AltElseNodeID, ast.EdgeTypeStmt, stmtListNodeID)
                yylex.(*Parser).ast.Children(0, $1, ast.EdgeTypeElse, AltElseNodeID)

                $$ = $1

                // save comments
                // yylex.(*Parser).setFreeFloating(_else, freefloating.Start, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(_else, freefloating.Else, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $5.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.AltEnd, $6.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($6))

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
                yylex.(*Parser).list.push()
                
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_parameter_list:
        parameter
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_parameter_list ',' parameter
            {
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

parameter:
        optional_type is_reference is_variadic T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($4),
                })

                varNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($4),
                })

                var posID linear.PositionID
                if $1 != 0 {
                    posID = yylex.(*Parser).ast.NewNodeTokenPosition($1, $4)
                } else if $2 != nil {
                    posID = yylex.(*Parser).ast.NewTokensPosition($2, $4)
                } else if $3 != nil {
                    posID = yylex.(*Parser).ast.NewTokensPosition($3, $4)
                } else {
                    posID = yylex.(*Parser).ast.NewTokenPosition($4)
                }

                var flag ast.NodeFlag
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                }
                if $3 != nil {
                    flag = flag | ast.NodeFlagVariadic
                }

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeParameter,
                    Flag: flag,
                    Pos:  posID,
                })

                yylex.(*Parser).ast.Children(0, varNodeID, ast.EdgeTypeVarName, identifierNodeID)
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVarType, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeVar, varNodeID)

                // // save comments
                // if $1 != nil {
                //     yylex.(*Parser).MoveFreeFloating($1, $$)
                // }
                // if $2 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.OptionalType, $2.FreeFloating)
                // }
                // if $3 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Ampersand, $3.FreeFloating)
                // }
                // yylex.(*Parser).setFreeFloating($$, freefloating.Variadic, $4.FreeFloating)
                // yylex.(*Parser).addDollarToken(variable)

                // // normalize
                // if $3 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Ampersand, (*$$.GetFreeFloating())[freefloating.Variadic]); delete((*$$.GetFreeFloating()), freefloating.Variadic)
                // }
                // if $2 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.OptionalType, (*$$.GetFreeFloating())[freefloating.Ampersand]); delete((*$$.GetFreeFloating()), freefloating.Ampersand)
                // }
                // if $1 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Start, (*$$.GetFreeFloating())[freefloating.OptionalType]); delete((*$$.GetFreeFloating()), freefloating.OptionalType)
                // }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   optional_type is_reference is_variadic T_VARIABLE '=' expr
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($4),
                })

                varNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($4),
                })

                var posID linear.PositionID
                if $1 != 0 {
                    posID = yylex.(*Parser).ast.NewNodesPosition($1, $6)
                } else if $2 != nil {
                    posID = yylex.(*Parser).ast.NewTokenNodePosition($2, $6)
                } else if $3 != nil {
                    posID = yylex.(*Parser).ast.NewTokenNodePosition($3, $6)
                } else {
                    posID = yylex.(*Parser).ast.NewTokenNodePosition($4, $6)
                }

                var flag ast.NodeFlag
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                }
                if $3 != nil {
                    flag = flag | ast.NodeFlagVariadic
                }

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeParameter,
                    Flag: flag,
                    Pos:  posID,
                })

                yylex.(*Parser).ast.Children(0, varNodeID, ast.EdgeTypeVarName, identifierNodeID)
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVarType, $1)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeVar, varNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeDefaultValue, $6)

                // // save comments
                // if $1 != nil {
                //     yylex.(*Parser).MoveFreeFloating($1, $$)
                // }
                // if $2 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.OptionalType, $2.FreeFloating)
                // }
                // if $3 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Ampersand, $3.FreeFloating)
                // }
                // yylex.(*Parser).setFreeFloating($$, freefloating.Variadic, $4.FreeFloating)
                // yylex.(*Parser).addDollarToken(variable)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $5.FreeFloating)

                // // normalize
                // if $3 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Ampersand, (*$$.GetFreeFloating())[freefloating.Variadic]); delete((*$$.GetFreeFloating()), freefloating.Variadic)
                // }
                // if $2 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.OptionalType, (*$$.GetFreeFloating())[freefloating.Ampersand]); delete((*$$.GetFreeFloating()), freefloating.Ampersand)
                // }
                // if $1 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Start, (*$$.GetFreeFloating())[freefloating.OptionalType]); delete((*$$.GetFreeFloating()), freefloating.OptionalType)
                // }

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeNullable,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

type:
        T_ARRAY
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CALLABLE
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

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
                $$ = $2;

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Colon, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

argument_list:
        '(' ')'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeArgumentList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $2),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.ArgumentList, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' non_empty_argument_list possible_comma ')'
            {
                children := yylex.(*Parser).list.pop()

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeArgumentList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeArguments, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // if $3 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.ArgumentList, append($3.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($3), $4.FreeFloating...)...))
                // } else {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.ArgumentList, $4.FreeFloating)
                // }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_argument_list:
        argument
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_argument_list ',' argument
            {
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

argument:
        expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeArgument,
                    Pos:  yylex.(*Parser).ast.NewNodePosition($1),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $1)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ELLIPSIS expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeArgument,
                    Flag: ast.NodeFlagVariadic,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

global_var_list:
        global_var_list ',' global_var
            {
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   global_var
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

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
                yylex.(*Parser).list.add($3)

                // save comments

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_var
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_var:
        T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                varNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtStaticVar,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                yylex.(*Parser).ast.Children(0, varNodeID, ast.EdgeTypeVarName, identifierNodeID)
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, varNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).addDollarToken(variable)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '=' expr
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                varNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtStaticVar,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, varNodeID, ast.EdgeTypeVarName, identifierNodeID)
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, varNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).addDollarToken(variable)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_statement_list:
        class_statement_list class_statement
            {
                yylex.(*Parser).list.add($2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                yylex.(*Parser).list.push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_statement:
        variable_modifiers property_list ';'
            {
                childrenProperties := yylex.(*Parser).list.pop()
                childrenModifiers := yylex.(*Parser).list.pop()

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtPropertyList,
                    Pos:  yylex.(*Parser).ast.NewNodeListTokenPosition(childrenModifiers, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeModifiers, childrenModifiers...)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeProperties, childrenProperties...)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1[0], $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.PropertyList, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   method_modifiers T_CONST class_const_list ';'
            {
                childrenConstants := yylex.(*Parser).list.pop()
                childrenModifiers := yylex.(*Parser).list.pop()

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtClassConstList,
                    Pos:  yylex.(*Parser).ast.NewOptionalListTokensPosition(childrenModifiers, $2, $4),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeModifiers, childrenModifiers...)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeConsts, childrenConstants...)

                // save comments
                // if len($1) > 0 {
                //     yylex.(*Parser).MoveFreeFloating($1[0], $$)
                //     yylex.(*Parser).setFreeFloating($$, freefloating.ModifierList, $2.FreeFloating)
                // } else {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Start, $2.FreeFloating)
                // }
                // yylex.(*Parser).setFreeFloating($$, freefloating.ConstList, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE name_list trait_adaptations
            {
                childrenTraits := yylex.(*Parser).list.pop()

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtTraitUse,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeTraitAdaptationList, $3)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeTraits, childrenTraits...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   method_modifiers T_FUNCTION returns_ref identifier backup_doc_comment '(' parameter_list ')' return_type method_body
            {
                childrenParams := yylex.(*Parser).list.pop()
                childrenModifiers := yylex.(*Parser).list.pop()

                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($4),
                })

                var posID linear.PositionID
                if len(childrenModifiers) == 0 {
                    posID = yylex.(*Parser).ast.NewTokenNodePosition($2, $10)
                } else {
                    posID = yylex.(*Parser).ast.NewNodeListNodePosition(childrenModifiers, $10)
                }

                var flag ast.NodeFlag
                if $3 != nil {
                    flag = flag | ast.NodeFlagRef
                }

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtClassMethod,
                    Flag: flag,
                    Pos:  posID,
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeMethodName, identifierNodeID)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeReturnType, $9)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeStmt, $10)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeModifiers, childrenModifiers...)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeParams, childrenParams...)

                // save comments
                // if len($1) > 0 {
                //     yylex.(*Parser).MoveFreeFloating($1[0], $$)
                //     yylex.(*Parser).setFreeFloating($$, freefloating.ModifierList, $2.FreeFloating)
                // } else {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Start, $2.FreeFloating)
                // }
                // if $3 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Function, $4.FreeFloating)
                // } else {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Function, $3.FreeFloating)
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Ampersand, $4.FreeFloating)
                // }
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $6.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.ParameterList, $8.FreeFloating)
                // if $9 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Params, (*$9.GetFreeFloating())[freefloating.Colon]); delete((*$9.GetFreeFloating()), freefloating.Colon)
                // }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

name_list:
        name
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   name_list ',' name
            {
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptations:
        ';'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtNop,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($1))


                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' '}'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtTraitAdaptationList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $2),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.AdaptationList, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' trait_adaptation_list '}'
            {
                children := yylex.(*Parser).list.pop()

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtTraitAdaptationList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeAdaptations, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.AdaptationList, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptation_list:
        trait_adaptation
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_adaptation_list trait_adaptation
            {
                yylex.(*Parser).list.add($2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptation:
        trait_precedence ';'
            {
                $$ = $1;

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.NameList, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_alias ';'
            {
                $$ = $1;

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Alias, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_precedence:
        absolute_trait_method_reference T_INSTEADOF name_list
            {
                children := yylex.(*Parser).list.pop()

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtTraitUsePrecedence,
                    Pos:  yylex.(*Parser).ast.NewNodeNodeListPosition($1, children),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeRef, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeInsteadof, children...)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Ref, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_alias:
        trait_method_reference T_AS T_STRING
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($3),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtTraitUseAlias,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeRef, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeAlias, identifierNodeID)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Ref, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(alias, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_method_reference T_AS reserved_non_modifiers
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($3),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtTraitUseAlias,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeRef, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeAlias, identifierNodeID)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Ref, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(alias, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_method_reference T_AS member_modifier identifier
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($4),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtTraitUseAlias,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $4),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeRef, $1)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeModifier, $3)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeAlias, identifierNodeID)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Ref, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(alias, freefloating.Start, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_method_reference T_AS member_modifier
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtTraitUseAlias,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeRef, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeModifier, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Ref, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_method_reference:
        identifier
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtTraitMethodRef,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeMethod, identifierNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

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
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewNodePosition($1),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtTraitMethodRef,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeTrait, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeMethod, identifierNodeID)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(target, freefloating.Start, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

method_body:
        ';' /* abstract method */ 
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtNop,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.SemiColon, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' inner_statement_list '}'
            {
                children := yylex.(*Parser).list.pop()

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })
                
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeStmts, children...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $3.FreeFloating)

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
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add(identifierNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating(modifier, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

method_modifiers:
        /* empty */
            {
                yylex.(*Parser).list.push()

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
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_member_modifiers member_modifier
            {
                yylex.(*Parser).list.add($2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

member_modifier:
        T_PUBLIC
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PROTECTED
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PRIVATE
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ABSTRACT
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINAL
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

property_list:
        property_list ',' property
            {
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   property
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

property:
        T_VARIABLE backup_doc_comment
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                varNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtProperty,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                yylex.(*Parser).ast.Children(0, varNodeID, ast.EdgeTypeVarName, identifierNodeID)
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, varNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).addDollarToken(variable)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '=' expr backup_doc_comment
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                varNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtProperty,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, varNodeID, ast.EdgeTypeVarName, identifierNodeID)
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, varNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).addDollarToken(variable)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_const_list:
        class_const_list ',' class_const_decl
            {
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_const_decl
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_const_decl:
        identifier '=' expr backup_doc_comment
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtConstant,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeConstantName, identifierNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

const_decl:
        T_STRING '=' expr backup_doc_comment
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtConstant,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeConstantName, identifierNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

echo_expr_list:
        echo_expr_list ',' echo_expr
            {
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   echo_expr
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

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
                yylex.(*Parser).list.push()

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
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

anonymous_class:
        T_CLASS ctor_arguments extends_from implements_list backup_doc_comment '{' class_statement_list '}'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeStmtClass,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $8),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExtends, $3)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeImplements, $4)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeStmts, yylex.(*Parser).list.pop()...)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeArgumentList, yylex.(*Parser).list.pop()...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $6.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $8.FreeFloating)

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

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprNew,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, lastNodeID),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeClass, $2)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeArgumentList, $3)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NEW anonymous_class
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprNew,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeClass, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

expr_without_variable:
        T_LIST '(' array_pair_list ')' '=' expr
            {
                listNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignAssign,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $6),
                })

                yylex.(*Parser).ast.Children(0, listNodeID, ast.EdgeTypeItems, yylex.(*Parser).list.pop()...)
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, listNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $6)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating(listNode, freefloating.List, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(listNode, freefloating.ArrayPairList, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $5.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' array_pair_list ']' '=' expr
            {
                listNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprShortList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignAssign,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $5),
                })

                yylex.(*Parser).ast.Children(0, listNodeID, ast.EdgeTypeItems, yylex.(*Parser).list.pop()...)
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, listNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $5)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating(shortList, freefloating.ArrayPairList, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignAssign,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' '&' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignReference,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $4),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $4)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Equal, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLONE expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprClone,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_PLUS_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignPlus,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MINUS_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignMinus,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MUL_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignMul,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_POW_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignPow,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_DIV_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignDiv,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_CONCAT_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignConcat,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MOD_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignMod,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_AND_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignBitwiseAnd,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_OR_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignBitwiseOr,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_XOR_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignBitwiseXor,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_SL_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignShiftLeft,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_SR_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeAssignShiftRight,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_INC
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprPostInc,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INC variable
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprPreInc,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_DEC
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprPostDec,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DEC variable
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprPreDec,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_BOOLEAN_OR expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryBooleanOr,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_BOOLEAN_AND expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryBooleanAnd,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_OR expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryLogicalOr,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_AND expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryLogicalAnd,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_XOR expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryLogicalXor,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '|' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryBitwiseOr,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '&' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryBitwiseAnd,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '^' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryBitwiseXor,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '.' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryConcat,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '+' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryPlus,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '-' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryMinus,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '*' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryMul,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_POW expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryPow,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '/' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryDiv,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '%' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryMod,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryShiftLeft,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SR expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryShiftRight,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '+' expr %prec T_INC
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprUnaryPlus,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '-' expr %prec T_INC
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprUnaryMinus,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '!' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprBooleanNot,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '~' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprBitwiseNot,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_IDENTICAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryIdentical,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_NOT_IDENTICAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryNotIdentical,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryEqual,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_NOT_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryNotEqual,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Equal, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '<' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinarySmaller,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_SMALLER_OR_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinarySmallerOrEqual,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '>' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryGreater,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_GREATER_OR_EQUAL expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryGreaterOrEqual,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SPACESHIP expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinarySpaceship,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_INSTANCEOF class_name_reference
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprInstanceOf,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeClass, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' expr ')'
            {
                $$ = $2;

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, append($1.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($1), (*$$.GetFreeFloating())[freefloating.Start]...)...))
                // yylex.(*Parser).setFreeFloating($$, freefloating.End, append((*$$.GetFreeFloating())[freefloating.End], append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_expr
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '?' expr ':' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprTernary,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $5),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeCond, $1)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeIfTrue, $3)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeIfFalse, $5)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.True, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '?' ':' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprTernary,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $4),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeCond, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeIfFalse, $4)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cond, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.True, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_COALESCE expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeBinaryCoalesce,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeLeft, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeRight, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, Coalesceeefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   internal_functions_in_yacc
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INT_CAST expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeCastInt,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOUBLE_CAST expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeCastDouble,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STRING_CAST expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeCastString,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ARRAY_CAST expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeCastArray,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_OBJECT_CAST expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeCastObject,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BOOL_CAST expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeCastBool,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_UNSET_CAST expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeCastUnset,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXIT exit_expr
            {
                $$ = $2

                var flag ast.NodeFlag
                if bytes.EqualFold($1.Value, []byte("die")) {
                    flag = ast.NodeFlagAltSyntax
                }

                if $$ == 0 {
                    $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                        Type: ast.NodeTypeExprExit,
                        Flag: flag,
                        Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                    })
                } else {
                    node := yylex.(*Parser).ast.Nodes.Get($$)
                    node.Pos = yylex.(*Parser).ast.NewTokenNodePosition($1, $2)
                    node.Flag = flag
                    yylex.(*Parser).ast.Nodes.Save($$, node)
                }

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '@' expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprErrorSuppress,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   scalar
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '`' backticks_expr '`'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprShellExec,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeParts, yylex.(*Parser).list.pop()...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PRINT expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprPrint,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprYield,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprYield,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVal, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD expr T_DOUBLE_ARROW expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprYield,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $4),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeKey, $2)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeVal, $4)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD_FROM expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprYieldFrom,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                var flag ast.NodeFlag
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                }

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprClosure,
                    Flag: flag,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $11),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeClosureUse, $7)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeReturnType, $8)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeStmts, yylex.(*Parser).list.pop()...)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeParams, yylex.(*Parser).list.pop()...)
                
                // // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // if $2 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Function, $4.FreeFloating)
                // } else {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Function, $2.FreeFloating)
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Ampersand, $4.FreeFloating)
                // }
                // yylex.(*Parser).setFreeFloating($$, freefloating.ParameterList, $6.FreeFloating)
                // if $8 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.LexicalVars, (*$8.GetFreeFloating())[freefloating.Colon]); delete((*$8.GetFreeFloating()), freefloating.Colon)
                // }
                // yylex.(*Parser).setFreeFloating($$, freefloating.ReturnType, $9.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $11.FreeFloating)

                // // normalize
                // if $8 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.LexicalVars, (*$$.GetFreeFloating())[freefloating.ReturnType]); delete((*$$.GetFreeFloating()), freefloating.ReturnType)
                // }
                // if $7 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Params, (*$$.GetFreeFloating())[freefloating.LexicalVarList]); delete((*$$.GetFreeFloating()), freefloating.LexicalVarList)
                // }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                flag := ast.NodeFlagStatic
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                }

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprClosure,
                    Flag: flag,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $12),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeClosureUse, $8)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeReturnType, $9)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeStmts, yylex.(*Parser).list.pop()...)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeParams, yylex.(*Parser).list.pop()...)
                
                // // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Static, $2.FreeFloating)
                // if $3 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Function, $5.FreeFloating)
                // } else {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Function, $3.FreeFloating)
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Ampersand, $5.FreeFloating)
                // }
                // yylex.(*Parser).setFreeFloating($$, freefloating.ParameterList, $7.FreeFloating)
                // if $9 != nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.LexicalVars, (*$9.GetFreeFloating())[freefloating.Colon]); delete((*$9.GetFreeFloating()), freefloating.Colon)
                // }
                // yylex.(*Parser).setFreeFloating($$, freefloating.ReturnType, $10.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Stmts, $12.FreeFloating)

                // // normalize
                // if $9 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.LexicalVars, (*$$.GetFreeFloating())[freefloating.ReturnType]); delete((*$$.GetFreeFloating()), freefloating.ReturnType)
                // }
                // if $8 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.Params, (*$$.GetFreeFloating())[freefloating.LexicalVarList]); delete((*$$.GetFreeFloating()), freefloating.LexicalVarList)
                // }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

backup_doc_comment:
        /* empty */
            {
                // $$ = yylex.(*Parser).Lexer.GetPhpDocComment()
                yylex.(*Parser).Lexer.SetPhpDocComment("")

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
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
                $$ = 0

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE '(' lexical_var_list ')'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprClosureUse,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeStmts, yylex.(*Parser).list.pop()...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Use, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.LexicalVarList, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

lexical_var_list:
        lexical_var_list ',' lexical_var
            {
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   lexical_var
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

lexical_var:
    T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVarName, identifierNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).addDollarToken($$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($2),
                })

                varNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($2),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprReference,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, varNodeID, ast.EdgeTypeVarName, identifierNodeID)
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, varNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating(variable, freefloating.Start, $2.FreeFloating)
                // yylex.(*Parser).addDollarToken(variable)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

function_call:
        name argument_list
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprFunctionCall,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $2),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeFunction, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeArgumentList, $2)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprStaticCall,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $4),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeClass, $1)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeCall, $3)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeArgumentList, $4)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprStaticCall,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $4),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeClass, $1)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeCall, $3)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeArgumentList, $4)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   callable_expr argument_list
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprFunctionCall,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $2),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeFunction, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeArgumentList, $2)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_name:
        T_STATIC
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprExit,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Exit, append($1.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($1)...))
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

backticks_expr:
        /* empty */
            {
                yylex.(*Parser).list.push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ENCAPSED_AND_WHITESPACE
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add(
                    yylex.(*Parser).ast.Nodes.Create(linear.Node{
                        Type: ast.NodeTypeScalarEncapsedStringPart,
                        Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                    }),
                )

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprArray,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeItems, yylex.(*Parser).list.pop()...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Array, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.ArrayPairList, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' array_pair_list ']'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprShortArray,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeItems, yylex.(*Parser).list.pop()...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.ArrayPairList, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONSTANT_ENCAPSED_STRING
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarString,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

scalar:
        T_LNUMBER
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarLnumber,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DNUMBER
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarDnumber,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LINE
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FILE
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DIR
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRAIT_C
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_METHOD_C
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FUNC_C
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_C
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLASS_C
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC
            {
                stringPartNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarEncapsedStringPart,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($2),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarHeredoc,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeParts, stringPartNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC T_END_HEREDOC
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarHeredoc,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $2),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '"' encaps_list '"'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarEncapsed,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeParts, yylex.(*Parser).list.pop()...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC encaps_list T_END_HEREDOC
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarHeredoc,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeParts, yylex.(*Parser).list.pop()...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprConstFetch,
                    Pos:  yylex.(*Parser).ast.NewNodePosition($1),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeConstant, $1)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM identifier
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($3),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprClassConstFetch,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeClass, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeConstantName, identifierNodeID)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(target, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM identifier
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($3),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprClassConstFetch,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeClass, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeConstantName, identifierNodeID)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(target, freefloating.Start, $3.FreeFloating)

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

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, append($1.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($1), (*$$.GetFreeFloating())[freefloating.Start]...)...))
                // yylex.(*Parser).setFreeFloating($$, freefloating.End, append((*$$.GetFreeFloating())[freefloating.End], append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...)...))

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

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, append($1.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($1), (*$$.GetFreeFloating())[freefloating.Start]...)...))
                // yylex.(*Parser).setFreeFloating($$, freefloating.End, append((*$$.GetFreeFloating())[freefloating.End], append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...)...))

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprArrayDimFetch,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $4),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeDim, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   constant '[' optional_expr ']'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprArrayDimFetch,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $4),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeDim, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   dereferencable '{' expr '}'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprArrayDimFetch,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $4),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeDim, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   dereferencable T_OBJECT_OPERATOR property_name argument_list
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprMethodCall,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $4),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                prevNodeID = yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeMethod, $3)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeArgumentList, $4)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprPropertyFetch,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeProperty, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

simple_variable:
        T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVarName, identifierNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).addDollarToken($$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' '{' expr '}'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVarName, $3)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Dollar, yylex.(*Parser).GetFreeFloatingToken($1))
                // yylex.(*Parser).setFreeFloating($3, freefloating.Start, append($2.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($2), (*$3.GetFreeFloating())[freefloating.Start]...)...))
                // yylex.(*Parser).setFreeFloating($3, freefloating.End, append((*$3.GetFreeFloating())[freefloating.End], append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' simple_variable
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVarName, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Dollar, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_member:
        class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprStaticPropertyFetch,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeClass, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeProperty, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprStaticPropertyFetch,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeClass, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeProperty, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Name, $2.FreeFloating)

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprArrayDimFetch,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $4),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeDim, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_variable '{' expr '}'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprArrayDimFetch,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $4),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeDim, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_variable T_OBJECT_OPERATOR property_name
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprPropertyFetch,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeClass, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeProperty, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprStaticPropertyFetch,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeClass, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeProperty, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_variable T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprStaticPropertyFetch,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeClass, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeProperty, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

member_name:
        identifier
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' expr '}'
            {
                $$ = $2;

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, append($1.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($1), (*$$.GetFreeFloating())[freefloating.Start]...)...))
                // yylex.(*Parser).setFreeFloating($$, freefloating.End, append((*$$.GetFreeFloating())[freefloating.End], append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...)...))

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' expr '}'
            {
                $$ = $2;
                
                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, append($1.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($1), (*$$.GetFreeFloating())[freefloating.Start]...)...))
                // yylex.(*Parser).setFreeFloating($$, freefloating.End, append((*$$.GetFreeFloating())[freefloating.End], append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...)...))

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
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
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
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   possible_array_pair
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

array_pair:
        expr T_DOUBLE_ARROW expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprArrayItem,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $3),
                })

                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeKey, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeVal, $3)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprArrayItem,
                    Pos:  yylex.(*Parser).ast.NewNodePosition($1),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVal, $1)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_DOUBLE_ARROW '&' variable
            {
                refNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprReference,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($3, $4),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprArrayItem,
                    Pos:  yylex.(*Parser).ast.NewNodesPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, refNodeID, ast.EdgeTypeVar, $4)
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeKey, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeVal, refNodeID)

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(reference, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' variable
            {
                refNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprReference,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprArrayItem,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, refNodeID, ast.EdgeTypeVar, $2)
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVal, refNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_DOUBLE_ARROW T_LIST '(' array_pair_list ')'
            {
                listNodeID :=  yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($3, $6),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprArrayItem,
                    Pos:  yylex.(*Parser).ast.NewNodeTokenPosition($1, $6),
                })

                yylex.(*Parser).ast.Children(0, listNodeID, ast.EdgeTypeItems, yylex.(*Parser).list.pop()...)
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeKey, $1)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeVal, listNodeID)

                // TODO: Cannot use list() as standalone expression

                // save comments
                // yylex.(*Parser).MoveFreeFloating($1, $$)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(listNode, freefloating.Start, $3.FreeFloating)
                // yylex.(*Parser).setFreeFloating(listNode, freefloating.List, $4.FreeFloating)
                // yylex.(*Parser).setFreeFloating(listNode, freefloating.ArrayPairList, $6.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LIST '(' array_pair_list ')'
            {
                listNodeID :=  yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprList,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprArrayItem,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, listNodeID, ast.EdgeTypeItems, yylex.(*Parser).list.pop()...)
                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVal, listNodeID)

                // TODO: Cannot use list() as standalone expression
                
                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating(listNode, freefloating.List, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(listNode, freefloating.ArrayPairList, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_list:
        encaps_list encaps_var
            {
                yylex.(*Parser).list.add($2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   encaps_list T_ENCAPSED_AND_WHITESPACE
            {
                yylex.(*Parser).list.add(
                    yylex.(*Parser).ast.Nodes.Create(linear.Node{
                        Type: ast.NodeTypeScalarEncapsedStringPart,
                        Pos:  yylex.(*Parser).ast.NewTokenPosition($2),
                    }),
                )

                // save comments
                // yylex.(*Parser).setFreeFloating(encapsed, freefloating.Start, $2.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   encaps_var
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ENCAPSED_AND_WHITESPACE encaps_var
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add(
                    yylex.(*Parser).ast.Nodes.Create(linear.Node{
                        Type: ast.NodeTypeScalarEncapsedStringPart,
                        Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                    }),
                )
                yylex.(*Parser).list.add($2)

                // save comments
                // yylex.(*Parser).setFreeFloating(encapsed, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_var:
        T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVarName, identifierNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).addDollarToken($$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '[' encaps_var_offset ']'
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                varNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprArrayDimFetch,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, varNodeID, ast.EdgeTypeVarName, identifierNodeID)
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, varNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeDim, $3)

                // save comments
                // yylex.(*Parser).addDollarToken(variable)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($2.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($2)...))
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($4.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE T_OBJECT_OPERATOR T_STRING
            {
                varNameNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                varNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                propertyNameNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($3),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprPropertyFetch,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, varNodeID, ast.EdgeTypeVarName, varNameNodeID)
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, varNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeProperty, propertyNameNodeID)

                // save comments
                // yylex.(*Parser).addDollarToken(variable)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating(fetch, freefloating.Start, $3.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES expr '}'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVarName, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, yylex.(*Parser).GetFreeFloatingToken($1))
                // yylex.(*Parser).setFreeFloating($$, freefloating.End, append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}'
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($2),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $3),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVarName, identifierNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, yylex.(*Parser).GetFreeFloatingToken($1))
                // yylex.(*Parser).setFreeFloating($$, freefloating.End, append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($2),
                })

                varNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($2),
                })
                
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprArrayDimFetch,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $6),
                })

                yylex.(*Parser).ast.Children(0, varNodeID, ast.EdgeTypeVarName, identifierNodeID)
                prevNodeID := yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVar, varNodeID)
                yylex.(*Parser).ast.Children(prevNodeID, $$, ast.EdgeTypeDim, $4)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, yylex.(*Parser).GetFreeFloatingToken($1))
                // yylex.(*Parser).setFreeFloating($$, freefloating.Var, append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...))
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, append($5.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($5)...))
                // yylex.(*Parser).setFreeFloating($$, freefloating.End, append($6.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($6)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CURLY_OPEN variable '}'
            {
                $$ = $2;

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, yylex.(*Parser).GetFreeFloatingToken($1))
                // yylex.(*Parser).setFreeFloating($$, freefloating.End, append($3.FreeFloating, yylex.(*Parser).GetFreeFloatingToken($3)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_var_offset:
        T_STRING
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeScalarString,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NUM_STRING
            {
                // TODO: add option to handle 64 bit integer
                if _, err := strconv.Atoi(string($1.Value)); err == nil {
                    $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                        Type: ast.NodeTypeScalarLnumber,
                        Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                    })
                } else {
                    $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                        Type: ast.NodeTypeScalarString,
                        Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                    })
                }

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '-' T_NUM_STRING
            {
                if _, err := strconv.Atoi(string($2.Value)); err == nil {
                    lnumberNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                        Type: ast.NodeTypeScalarLnumber,
                        Pos:  yylex.(*Parser).ast.NewTokenPosition($2),
                    })
                    
                    $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                        Type: ast.NodeTypeExprUnaryMinus,
                        Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $2),
                    })

                    yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, lnumberNodeID)
                } else {
                    $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                        Type: ast.NodeTypeScalarString,
                        Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $2),
                    })
                }

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                identifierNodeID := yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeIdentifier,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprVariable,
                    Pos:  yylex.(*Parser).ast.NewTokenPosition($1),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVarName, identifierNodeID)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).addDollarToken($$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

internal_functions_in_yacc:
        T_ISSET '(' isset_variables possible_comma ')'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprIsset,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $5),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeVars, yylex.(*Parser).list.pop()...)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Isset, $2.FreeFloating)
                // if $4 == nil {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.VarList, $5.FreeFloating)
                // } else {
                //     yylex.(*Parser).setFreeFloating($$, freefloating.VarList, append($4.FreeFloating, append(yylex.(*Parser).GetFreeFloatingToken($4), $5.FreeFloating...)...))
                // }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EMPTY '(' expr ')'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprEmpty,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Empty, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INCLUDE expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprInclude,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INCLUDE_ONCE expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprIncludeOnce,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EVAL '(' expr ')'
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprEval,
                    Pos:  yylex.(*Parser).ast.NewTokensPosition($1, $4),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $3)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Eval, $2.FreeFloating)
                // yylex.(*Parser).setFreeFloating($$, freefloating.Expr, $4.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_REQUIRE expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprRequire,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_REQUIRE_ONCE expr
            {
                $$ = yylex.(*Parser).ast.Nodes.Create(linear.Node{
                    Type: ast.NodeTypeExprRequireOnce,
                    Pos:  yylex.(*Parser).ast.NewTokenNodePosition($1, $2),
                })

                yylex.(*Parser).ast.Children(0, $$, ast.EdgeTypeExpr, $2)

                // save comments
                // yylex.(*Parser).setFreeFloating($$, freefloating.Start, $1.FreeFloating)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

isset_variables:
        isset_variable
            {
                yylex.(*Parser).list.push()
                yylex.(*Parser).list.add($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   isset_variables ',' isset_variable
            {
                yylex.(*Parser).list.add($3)

                // save comments
                // yylex.(*Parser).setFreeFloating(lastNode($1), freefloating.End, $2.FreeFloating)

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
