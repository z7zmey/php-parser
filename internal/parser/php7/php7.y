%{
package php7

import (
    "bytes"
    "strconv"

    "github.com/z7zmey/php-parser/pkg/ast"
    "github.com/z7zmey/php-parser/internal/scanner"
)

%}

%union{
    node  ast.Node
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

%type <node> name
%type <node> top_statement statement inner_statement
%type <node> function_declaration_statement
%type <node> return_type type_expr type
%type <node> parameter optional_type
%type <node> class_declaration_statement trait_declaration_statement interface_declaration_statement
%type <node> extends_from interface_extends_list implements_list
%type <node> class_modifier class_statement member_modifier property class_const_decl method_body
%type <node> trait_adaptations trait_adaptation trait_precedence trait_alias trait_method_reference absolute_trait_method_reference
%type <node> group_use_declaration inline_use_declaration use_type
%type <node> mixed_group_use_declaration use_declaration unprefixed_use_declaration
%type <node> const_decl
%type <node> expr expr_without_variable optional_expr exit_expr lexical_vars lexical_var
%type <node> internal_functions_in_yacc isset_variable
%type <node> new_expr class_name_reference new_variable ctor_arguments anonymous_class
%type <node> scalar encaps_var encaps_var_offset
%type <node> variable callable_variable simple_variable
%type <node> static_member class_name variable_class_name
%type <node> dereferencable dereferencable_scalar constant
%type <node> property_name argument_list argument
%type <node> function_call callable_expr member_name
%type <node> array_pair possible_array_pair
%type <node> if_stmt alt_if_stmt
%type <node> global_var static_var echo_expr unset_variable foreach_variable
%type <node> finally_statement

%%

/////////////////////////////////////////////////////////////////////////

start:
        top_statement_list
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, yylex.(*Parser).CurrentToken.HiddenTokens )
                
                root := ast.Node{
                    Type: ast.NodeTypeRoot,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes( root ),
                )

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
                if $2.Type != 0 {
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
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                namePartNode := ast.Node{
                    Type: ast.NodeTypeNameNamePart,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(namePartNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_NS_SEPARATOR T_STRING
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$2} )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $3.HiddenTokens )
                
                namePartNode := ast.Node{
                    Type: ast.NodeTypeNameNamePart,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).List.Add(namePartNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

name:
        namespace_name
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupParts),
                )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeNameName,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    | T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupParts),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupNamespace, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeNameRelative,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    | T_NS_SEPARATOR namespace_name
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupParts),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeNameFullyQualified,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

top_statement:
        error
            {
                // error
                $$ = ast.Node{};

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
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupHaltCompiller, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupOpenParenthesisToken, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCloseParenthesisToken, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$4} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtHaltCompiler,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE namespace_name ';'
            {
                // namespace name 

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupParts),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $3.HiddenTokens )
                
                nameNode := ast.Node{
                    Type: ast.NodeTypeNameName,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupNamespaceName,
                }

                // namespace

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(nameNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtNamespace,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE namespace_name '{' top_statement_list '}'
            {
                childrenStmts := yylex.(*Parser).List.Pop(ast.NodeGroupStmts)
                childrenNameParts := yylex.(*Parser).List.Pop(ast.NodeGroupParts)

                // namespace name 

                children := yylex.(*Parser).Ast.AddNodes(
                    childrenNameParts,
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $3.HiddenTokens )
                
                nameNode := ast.Node{
                    Type: ast.NodeTypeNameName,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupNamespaceName,
                }

                // namespace

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(nameNode),
                    childrenStmts,
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $5.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtNamespace,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $5}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE '{' top_statement_list '}'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupNamespace, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $4.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtNamespace,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE mixed_group_use_declaration ';'
            {
                $$ = $2

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupUseDeclarationList, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$3} )

                $$.Position = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil)
                $$.Tokens = yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() )

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE use_type group_use_declaration ';'
            {
                $2.Group = ast.NodeGroupUseType

                $$ = $3

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupUseDeclarationList, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$4} )

                $$.Position = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil)
                $$.Tokens = yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() )
                $$.Children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                    yylex.(*Parser).List.Pop(ast.NodeGroupPrefix, ast.NodeGroupUseList),
                )

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE use_declarations ';'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupUses),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupUseDeclarationList, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtUseList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE use_type use_declarations ';'
            {
                $2.Group = ast.NodeGroupUseType

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                    yylex.(*Parser).List.Pop(ast.NodeGroupUses),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupUseDeclarationList, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$4} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtUseList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONST const_list ';'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupConsts),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtConstList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_type:
        T_FUNCTION
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }
            }
    |   T_CONST
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }
            }
;

group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                childrenUseDeclarations := yylex.(*Parser).List.Pop(ast.NodeGroupUseList)
                childrenNameParts       := yylex.(*Parser).List.Pop(ast.NodeGroupParts)

                // name 

                children := yylex.(*Parser).Ast.AddNodes(
                    childrenNameParts,
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $2.HiddenTokens )
                
                nameNode := ast.Node{
                    Type: ast.NodeTypeNameName,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupPrefix,
                }

                // GroupUse
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtGroupUse,
                    Position: yylex.(*Parser).NewPosition(childrenNameParts, []*scanner.Token{$6}, nil),
                    Children: children,
                }

                // push children

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(childrenUseDeclarations...)
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(nameNode)

                // push tokens

                yylex.(*Parser).PushTokens( ast.TokenGroupSlash, $3.HiddenTokens )
                if $5 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $5.HiddenTokens )
                    yylex.(*Parser).PushTokens( ast.TokenGroupStmts, []scanner.Token{*$5} )
                } 
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $6.HiddenTokens )
                
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                childrenUseDeclarations := yylex.(*Parser).List.Pop(ast.NodeGroupUseList)
                childrenNameParts       := yylex.(*Parser).List.Pop(ast.NodeGroupParts)

                // name 

                children := yylex.(*Parser).Ast.AddNodes(
                    childrenNameParts,
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $3.HiddenTokens )
                
                nameNode := ast.Node{
                    Type: ast.NodeTypeNameName,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupPrefix,
                }

                // GroupUse
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtGroupUse,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $7}, nil),
                    Children: children,
                }

                // push children

                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(childrenUseDeclarations...)
                yylex.(*Parser).List.Push()
                yylex.(*Parser).List.Add(nameNode)

                // push tokens

                yylex.(*Parser).PushTokens( ast.TokenGroupUse, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupUse, []scanner.Token{*$1} )
                yylex.(*Parser).PushTokens( ast.TokenGroupSlash, $4.HiddenTokens )
                if $6 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $6.HiddenTokens )
                    yylex.(*Parser).PushTokens( ast.TokenGroupStmts, []scanner.Token{*$6} )
                } 
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $7.HiddenTokens )
                
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

mixed_group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                childrenUseDeclarations := yylex.(*Parser).List.Pop(ast.NodeGroupUseList)
                childrenNameParts       := yylex.(*Parser).List.Pop(ast.NodeGroupParts)

                // name 

                children := yylex.(*Parser).Ast.AddNodes(
                    childrenNameParts,
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $2.HiddenTokens )
                
                nameNode := ast.Node{
                    Type: ast.NodeTypeNameName,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupPrefix,
                }

                // GroupUse

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(nameNode),
                    childrenUseDeclarations,
                )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtGroupUse,
                    Position: yylex.(*Parser).NewPosition(childrenNameParts, []*scanner.Token{$6}, nil),
                    Children: children,
                }

                // push tokens

                yylex.(*Parser).PushTokens( ast.TokenGroupSlash, $3.HiddenTokens )
                if $5 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $5.HiddenTokens )
                    yylex.(*Parser).PushTokens( ast.TokenGroupStmts, []scanner.Token{*$5} )
                } 
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $6.HiddenTokens )
                
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                childrenUseDeclarations := yylex.(*Parser).List.Pop(ast.NodeGroupUseList)
                childrenNameParts       := yylex.(*Parser).List.Pop(ast.NodeGroupParts)

                // name 

                children := yylex.(*Parser).Ast.AddNodes(
                    childrenNameParts,
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $3.HiddenTokens )
                
                nameNode := ast.Node{
                    Type: ast.NodeTypeNameName,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupPrefix,
                }

                // GroupUse

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(nameNode),
                    childrenUseDeclarations,
                )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtGroupUse,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $7}, nil),
                    Children: children,
                }

                // push tokens

                yylex.(*Parser).PushTokens( ast.TokenGroupUse, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupUse, []scanner.Token{*$1} )
                yylex.(*Parser).PushTokens( ast.TokenGroupSlash, $4.HiddenTokens )
                if $6 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $6.HiddenTokens )
                    yylex.(*Parser).PushTokens( ast.TokenGroupStmts, []scanner.Token{*$6} )
                } 
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $7.HiddenTokens )
                
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
        inline_use_declarations ',' { yylex.(*Parser).PrependToken($2) } inline_use_declaration
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   inline_use_declaration
            {
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unprefixed_use_declarations:
        unprefixed_use_declarations ',' { yylex.(*Parser).PrependToken($2) } unprefixed_use_declaration
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupAlias, ast.NodeGroupUse),
                )

                useNode := ast.Node{
                    Type: ast.NodeTypeStmtUse,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).List.Add(useNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   unprefixed_use_declaration
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupAlias, ast.NodeGroupUse),
                )

                useNode := ast.Node{
                    Type: ast.NodeTypeStmtUse,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).List.Push(useNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_declarations:
        use_declarations ',' { yylex.(*Parser).PrependToken($2) } use_declaration
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   use_declaration
            {
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

inline_use_declaration:
        unprefixed_use_declaration
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupAlias, ast.NodeGroupUse),
                )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtUse,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   use_type unprefixed_use_declaration
            {
                $1.Group = ast.NodeGroupUseType

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1),
                    yylex.(*Parser).List.Pop(ast.NodeGroupAlias, ast.NodeGroupUse),
                )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtUse,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unprefixed_use_declaration:
        namespace_name
            {
                // Name

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupParts),
                )
                
                nameNode := ast.Node{
                    Type: ast.NodeTypeNameName,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupUse,
                }

                yylex.(*Parser).List.Push(nameNode)

                // Alias

                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_AS T_STRING
            {
                // Name

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupParts),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $2.HiddenTokens )
                
                nameNode := ast.Node{
                    Type: ast.NodeTypeNameName,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupUse,
                }

                yylex.(*Parser).List.Push(nameNode)

                // Alias

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $3.HiddenTokens )
                
                aliasNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupAlias,
                }

                yylex.(*Parser).List.Push(aliasNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_declaration:
        unprefixed_use_declaration
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupAlias, ast.NodeGroupUse),
                )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtUse,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR unprefixed_use_declaration
            {
                // Use

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupAlias, ast.NodeGroupUse),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSlash, []scanner.Token{*$1} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtUse,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

const_list:
        const_list ',' { yylex.(*Parser).PrependToken($2) } const_decl
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   const_decl
            {
                yylex.(*Parser).List.Push($1)

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

                if $2.Type != 0 {
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
                $$ = ast.Node{};

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
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupHaltCompiller, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupOpenParenthesisToken, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCloseParenthesisToken, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$4} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtHaltCompiler,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }

statement:
        '{' inner_statement_list '}'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $3.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                $$ = yylex.(*Parser).List.Pop(ast.NodeGroupNil)[0]
                statementNode := yylex.(*Parser).List.Pop(ast.NodeGroupStmt)[0]

                $3.Group = ast.NodeGroupCond
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($3, statementNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupWhile, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )

                $$.Children = children
                $$.Position = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, yylex.(*Parser).Nodes($$))
                $$.Tokens = yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() )

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DO statement T_WHILE '(' expr ')' ';'
            {
                $2.Group = ast.NodeGroupStmt
                $5.Group = ast.NodeGroupCond
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2, $5),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupWhile, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $6.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCond, $7.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$7} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtDo,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $7}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOR '(' for_exprs ';' for_exprs ';' for_exprs ')' for_statement
            {
                $$ = yylex.(*Parser).List.Pop(ast.NodeGroupNil)[0]
                statementNode := yylex.(*Parser).List.Pop(ast.NodeGroupStmt)[0]

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupLoop, ast.NodeGroupCond, ast.NodeGroupInit),
                    yylex.(*Parser).Nodes(statementNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupFor, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupInitExpr, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCondExpr, $6.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupIncExpr, $8.HiddenTokens )

                $$.Children = children
                $$.Position = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, yylex.(*Parser).Nodes($$))
                $$.Tokens = yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() )

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_SWITCH '(' expr ')' switch_case_list
            {
                $$ = yylex.(*Parser).List.Pop(ast.NodeGroupNil)[0]
                caseListNode := yylex.(*Parser).List.Pop(ast.NodeGroupCaseList)[0]

                $3.Group = ast.NodeGroupCond
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($3, caseListNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSwitch, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )

                $$.Children = children
                $$.Position = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, yylex.(*Parser).Nodes($$))
                $$.Tokens = yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() )

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BREAK optional_expr ';'
            {
                var children []ast.Node

                if $2.Type != ast.NodeTypeNil {
                    $2.Group = ast.NodeGroupExpr
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes($2),
                    )
                };

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtBreak,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONTINUE optional_expr ';'
            {
                var children []ast.Node

                if $2.Type != ast.NodeTypeNil {
                    $2.Group = ast.NodeGroupExpr
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes($2),
                    )
                };

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtContinue,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_RETURN optional_expr ';'
            {
                var children []ast.Node

                if $2.Type != ast.NodeTypeNil {
                    $2.Group = ast.NodeGroupExpr
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes($2),
                    )
                };

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtReturn,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_GLOBAL global_var_list ';'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupVars),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVarList, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtGlobal,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC static_var_list ';'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupVars),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVarList, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtStatic,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ECHO echo_expr_list ';'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupExprs),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEcho, []scanner.Token{*$1} )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtEcho,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INLINE_HTML
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtInlineHtml,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr ';'
            {
                $1.Group = ast.NodeGroupExpr
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$2} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtExpression,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$2}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_UNSET '(' unset_variables possible_comma ')' ';' 
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupVars),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupUnset, $2.HiddenTokens )
                if $4 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupVarList, $4.HiddenTokens )
                    yylex.(*Parser).PushTokens( ast.TokenGroupVarList, []scanner.Token{*$4} )
                };
                yylex.(*Parser).PushTokens( ast.TokenGroupVarList, $5.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCloseParenthesisToken, $6.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$6} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtUnset,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $6}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOREACH '(' expr T_AS foreach_variable ')' foreach_statement
            {
                $$ = yylex.(*Parser).List.Pop(ast.NodeGroupNil)[0]
                statementNode := yylex.(*Parser).List.Pop(ast.NodeGroupStmt)[0]

                $3.Group = ast.NodeGroupExpr
                $5.Group = ast.NodeGroupVar
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($3, $5, statementNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupForeach, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $6.HiddenTokens )

                $$.Children = children
                $$.Position = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, yylex.(*Parser).Nodes($$))
                $$.Tokens = yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() )

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FOREACH '(' expr T_AS variable T_DOUBLE_ARROW foreach_variable ')' foreach_statement
            {
                $$ = yylex.(*Parser).List.Pop(ast.NodeGroupNil)[0]
                statementNode := yylex.(*Parser).List.Pop(ast.NodeGroupStmt)[0]

                $3.Group = ast.NodeGroupExpr
                $5.Group = ast.NodeGroupKey
                $7.Group = ast.NodeGroupVar
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($3, $5, $7, statementNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupForeach, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupKey, $6.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $8.HiddenTokens )

                $$.Children = children
                $$.Position = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, yylex.(*Parser).Nodes($$))
                $$.Tokens = yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() )

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DECLARE '(' const_list ')' declare_statement
            {
                $$ = yylex.(*Parser).List.Pop(ast.NodeGroupNil)[0]
                statementNode := yylex.(*Parser).List.Pop(ast.NodeGroupStmt)[0]

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupConsts),
                    yylex.(*Parser).Nodes(statementNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupDeclare, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupConstList, $4.HiddenTokens )

                $$.Children = children
                $$.Position = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, yylex.(*Parser).Nodes($$))
                $$.Tokens = yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() )

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ';'
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$1} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtNop,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRY '{' inner_statement_list '}' catch_list finally_statement
            {
                var nodes []ast.Node

                if $6.Type != ast.NodeTypeNil {
                    $6.Group = ast.NodeGroupFinally
                    nodes = yylex.(*Parser).Nodes($6);
                };

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupCatches, ast.NodeGroupStmts), nodes,
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupTry, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $4.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtTry,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_THROW expr ';'
            {
                $2.Group = ast.NodeGroupExpr
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtThrow,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_GOTO T_STRING ';'
            {
                // Identifier

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupLabel,
                };

                // Goto

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupLabel, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtGoto,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STRING ':'
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Group: ast.NodeGroupLabelName,
                };

                // Label

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupLabel, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtLabel,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$5}, nil),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $5.HiddenTokens )
                
                variableNode := ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVar,
                };

                // catch

                statements := yylex.(*Parser).List.Pop(ast.NodeGroupStmts)
                types := yylex.(*Parser).List.Pop(ast.NodeGroupTypes)

                children = yylex.(*Parser).Ast.AddNodes(
                    types, yylex.(*Parser).Nodes(variableNode), statements,
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCatch, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $6.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCond, $7.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $9.HiddenTokens )

                catchNode := ast.Node{
                    Type: ast.NodeTypeStmtCatch,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2, $9}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).List.Add(catchNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

catch_name_list:
        name
            {
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   catch_name_list '|' { yylex.(*Parser).PrependToken($2) } name
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

finally_statement:
        /* empty */
            {
                $$ = ast.Node{};

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINALLY '{' inner_statement_list '}'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupFinally, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $4.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtFinally,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2, $4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unset_variables:
        unset_variable
            {
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   unset_variables ',' { yylex.(*Parser).PrependToken($2) } unset_variable
            {
                yylex.(*Parser).List.Add($4)

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
                // Identifier

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $3.HiddenTokens )

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil),
                    Group: ast.NodeGroupFunctionName,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                // Function

                var children []ast.Node
                if $8.Type == ast.NodeTypeNil {
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes(identifierNode),
                        yylex.(*Parser).List.Pop(ast.NodeGroupStmts, ast.NodeGroupParams),
                    );
                } else {
                    $8.Group = ast.NodeGroupReturnType
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes(identifierNode, $8),
                        yylex.(*Parser).List.Pop(ast.NodeGroupStmts, ast.NodeGroupParams),
                    );
                };

                var flag ast.NodeFlag
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                };

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                if $2 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupFunction, $2.HiddenTokens )
                };
                yylex.(*Parser).PushTokens( ast.TokenGroupName, $5.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupParamList, $7.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupReturnType, $9.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $11.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtFunction,
                    Flag: flag,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $11}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                // Identifier

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $3.HiddenTokens )

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil),
                    Group: ast.NodeGroupClassName,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                // Class

                childrenStmts := yylex.(*Parser).List.Pop(ast.NodeGroupStmts)
                childrenModifiers := yylex.(*Parser).List.Pop(ast.NodeGroupModifiers)

                nodes := make([]ast.Node, 0, 3)
                nodes = append(nodes, identifierNode)

                if $4.Type != ast.NodeTypeNil {
                    $4.Group = ast.NodeGroupExtends
                    nodes = append(nodes, $4)
                };

                if $5.Type != ast.NodeTypeNil {
                    $5.Group = ast.NodeGroupImplements
                    nodes = append(nodes, $5)
                };

                children := yylex.(*Parser).Ast.AddNodes(
                    childrenModifiers, nodes, childrenStmts,
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupModifierList, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupName, $7.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $9.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtClass,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$9}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
            {
                // Identifier

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil),
                    Group: ast.NodeGroupClassName,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                // Class

                nodes := make([]ast.Node, 0, 3)
                nodes = append(nodes, identifierNode)

                if $3.Type != ast.NodeTypeNil {
                    $3.Group = ast.NodeGroupExtends
                    nodes = append(nodes, $3)
                };

                if $4.Type != ast.NodeTypeNil {
                    $4.Group = ast.NodeGroupImplements
                    nodes = append(nodes, $4)
                };

                children := yylex.(*Parser).Ast.AddNodes(
                    nodes, 
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupName, $6.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $8.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtClass,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $8}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_modifiers:
        class_modifier
            {
                yylex.(*Parser).List.Push($1)

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
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINAL
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_declaration_statement:
        T_TRAIT T_STRING backup_doc_comment '{' class_statement_list '}'
            {
                // Identifier

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil),
                    Group: ast.NodeGroupTraitName,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                // Trait

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupName, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $6.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtTrait,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $6}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

interface_declaration_statement:
        T_INTERFACE T_STRING interface_extends_list backup_doc_comment '{' class_statement_list '}'
            {
                // Identifier

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil),
                    Group: ast.NodeGroupClassName,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                // Interface

                nodes := make([]ast.Node, 0, 3)
                nodes = append(nodes, identifierNode)

                if $3.Type != ast.NodeTypeNil {
                    $3.Group = ast.NodeGroupExtends
                    nodes = append(nodes, $3)
                };

                children := yylex.(*Parser).Ast.AddNodes(
                    nodes, 
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupName, $5.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $7.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtInterface,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $7}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

extends_from:
        /* empty */
            {
                $$ = ast.Node{};

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXTENDS name
            {
                $2.Group = ast.NodeGroupClassName
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtClassExtends,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

interface_extends_list:
        /* empty */
            {
                $$ = ast.Node{};

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXTENDS name_list
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupInterfaceNames),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtInterfaceExtends,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

implements_list:
        /* empty */
            {
                $$ = ast.Node{};

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_IMPLEMENTS name_list
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupInterfaceNames),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtClassImplements,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                $2.Group = ast.NodeGroupVar
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprReference,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LIST '(' array_pair_list ')'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupItems),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupList, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupArrayPairList, $4.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' array_pair_list ']'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupItems),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupArrayPairList, $3.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprShortList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

for_statement:
        statement
            {
                forNode := ast.Node{
                    Type: ast.NodeTypeStmtFor,
                    Position: yylex.(*Parser).NewPosition(yylex.(*Parser).Nodes($1), nil, nil),
                };

                yylex.(*Parser).List.Push($1)
                yylex.(*Parser).List.Push(forNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDFOR ';'
            {
                // Statement
                
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                );
                
                statementNode := ast.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                };

                // While

                yylex.(*Parser).PushTokens( ast.TokenGroupCond, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupAltEnd, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$4} )
                
                forNode := ast.Node{
                    Type: ast.NodeTypeStmtAltFor,
                    Flag: ast.NodeFlagAltSyntax,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$1, $4}, nil),
                };

                yylex.(*Parser).List.Push(statementNode)
                yylex.(*Parser).List.Push(forNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

foreach_statement:
        statement
            {
                foreachNode := ast.Node{
                    Type: ast.NodeTypeStmtForeach,
                    Position: yylex.(*Parser).NewPosition(yylex.(*Parser).Nodes($1), nil, nil),
                };

                yylex.(*Parser).List.Push($1)
                yylex.(*Parser).List.Push(foreachNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDFOREACH ';'
            {
                // Statement
                
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                );
                
                statementNode := ast.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                };

                // While

                yylex.(*Parser).PushTokens( ast.TokenGroupCond, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupAltEnd, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$4} )
                
                foreachNode := ast.Node{
                    Type: ast.NodeTypeStmtAltForeach,
                    Flag: ast.NodeFlagAltSyntax,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$1, $4}, nil),
                };

                yylex.(*Parser).List.Push(statementNode)
                yylex.(*Parser).List.Push(foreachNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

declare_statement:
        statement
            {
                declareNode := ast.Node{
                    Type: ast.NodeTypeStmtDeclare,
                    Position: yylex.(*Parser).NewPosition(yylex.(*Parser).Nodes($1), nil, nil),
                };

                yylex.(*Parser).List.Push($1)
                yylex.(*Parser).List.Push(declareNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDDECLARE ';'
            {
                // Statement
                
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                );
                
                statementNode := ast.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                };

                // While

                yylex.(*Parser).PushTokens( ast.TokenGroupCond, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupAltEnd, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$4} )
                
                declareNode := ast.Node{
                    Type: ast.NodeTypeStmtDeclare,
                    Flag: ast.NodeFlagAltSyntax,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$1, $4}, nil),
                };

                yylex.(*Parser).List.Push(statementNode)
                yylex.(*Parser).List.Push(declareNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

switch_case_list:
        '{' case_list '}'
            {
                // CaseList
                
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupCases),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCaseListEnd, $3.HiddenTokens )
                
                caseListNode := ast.Node{
                    Type: ast.NodeTypeStmtCaseList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                // Switch
                
                switchNode := ast.Node{
                    Type: ast.NodeTypeStmtSwitch,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$1, $3}, nil),
                };

                yylex.(*Parser).List.Push(caseListNode)
                yylex.(*Parser).List.Push(switchNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' ';' case_list '}'
            {
                // CaseList
                
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupCases),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCaseListStart, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCaseListStart, []scanner.Token{*$2} )
                yylex.(*Parser).PushTokens( ast.TokenGroupCaseListEnd, $4.HiddenTokens )
                
                caseListNode := ast.Node{
                    Type: ast.NodeTypeStmtCaseList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                // Switch
                
                switchNode := ast.Node{
                    Type: ast.NodeTypeStmtSwitch,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$1, $4}, nil),
                };

                yylex.(*Parser).List.Push(caseListNode)
                yylex.(*Parser).List.Push(switchNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' case_list T_ENDSWITCH ';'
            {
                // CaseList
                
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupCases),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupCaseListEnd, $3.HiddenTokens )
                
                caseListNode := ast.Node{
                    Type: ast.NodeTypeStmtCaseList,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                // Switch

                yylex.(*Parser).PushTokens( ast.TokenGroupCond, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupAltEnd, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$4} )
                
                switchNode := ast.Node{
                    Type: ast.NodeTypeStmtSwitch,
                    Flag: ast.NodeFlagAltSyntax,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$1, $4}, nil),
                };

                yylex.(*Parser).List.Push(caseListNode)
                yylex.(*Parser).List.Push(switchNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' ';' case_list T_ENDSWITCH ';'
            {
                // CaseList
                
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupCases),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupCaseListStart, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCaseListStart, []scanner.Token{*$2} )
                yylex.(*Parser).PushTokens( ast.TokenGroupCaseListEnd, $4.HiddenTokens )
                
                caseListNode := ast.Node{
                    Type: ast.NodeTypeStmtCaseList,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                // Switch

                yylex.(*Parser).PushTokens( ast.TokenGroupCond, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupAltEnd, $5.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$5} )
                
                switchNode := ast.Node{
                    Type: ast.NodeTypeStmtSwitch,
                    Flag: ast.NodeFlagAltSyntax,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$1, $5}, nil),
                };

                yylex.(*Parser).List.Push(caseListNode)
                yylex.(*Parser).List.Push(switchNode)

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
                $3.Group = ast.NodeGroupExpr
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($3),
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCaseSeparator, []scanner.Token{*$4} )
                
                caseNode := ast.Node{
                    Type: ast.NodeTypeStmtCase,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).List.Add(caseNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   case_list T_DEFAULT case_separator inner_statement_list
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupDefault, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCaseSeparator, []scanner.Token{*$3} )
                
                caseNode := ast.Node{
                    Type: ast.NodeTypeStmtDefault,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).List.Add(caseNode)

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
                whileNode := ast.Node{
                    Type: ast.NodeTypeStmtWhile,
                    Position: yylex.(*Parser).NewPosition(yylex.(*Parser).Nodes($1), nil, nil),
                };

                yylex.(*Parser).List.Push($1)
                yylex.(*Parser).List.Push(whileNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDWHILE ';'
            {
                // Statement
                
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                );
                
                statementNode := ast.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                };

                // While

                yylex.(*Parser).PushTokens( ast.TokenGroupCond, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupAltEnd, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$4} )
                
                whileNode := ast.Node{
                    Type: ast.NodeTypeStmtWhile,
                    Flag: ast.NodeFlagAltSyntax,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$1, $4}, nil),
                };

                yylex.(*Parser).List.Push(statementNode)
                yylex.(*Parser).List.Push(whileNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

if_stmt_without_else:
        T_IF '(' expr ')' statement
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupIf, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )
                
                ifNode := ast.Node{
                    Type: ast.NodeTypeStmtIf,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).List.Push(ifNode)
                yylex.(*Parser).List.Push($3)
                yylex.(*Parser).List.Push($5)
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   if_stmt_without_else T_ELSEIF '(' expr ')' statement
            {
                $4.Group = ast.NodeGroupCond
                $6.Group = ast.NodeGroupStmt

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($4, $6),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupElseIf, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $5.HiddenTokens )
                
                elseIfNode := ast.Node{
                    Type: ast.NodeTypeStmtElseIf,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).List.Add(elseIfNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

if_stmt:
        if_stmt_without_else %prec T_NOELSE
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupElseIf, ast.NodeGroupStmt, ast.NodeGroupCond),
                );

                $$ = yylex.(*Parser).List.Pop(ast.NodeGroupNil)[0]

                $$.Position = yylex.(*Parser).NewPosition(yylex.(*Parser).Nodes($$), nil, children)
                $$.Children = children

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   if_stmt_without_else T_ELSE statement
            {
                // Else
                
                $3.Group = ast.NodeGroupStmt

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($3),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )
                
                elseNode := ast.Node{
                    Type: ast.NodeTypeStmtElse,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupElse,
                };

                // If

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupElseIf, ast.NodeGroupStmt, ast.NodeGroupCond),
                    yylex.(*Parser).Nodes(elseNode),
                );

                $$ = yylex.(*Parser).List.Pop(ast.NodeGroupNil)[0]

                $$.Position = yylex.(*Parser).NewPosition(yylex.(*Parser).Nodes($$), nil, children)
                $$.Children = children

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

alt_if_stmt_without_else:
        T_IF '(' expr ')' ':' inner_statement_list
            {
                // Statement

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                )

                stmtNode := ast.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Group: ast.NodeGroupStmt,
                };

                // Alt if

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupIf, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCond, $5.HiddenTokens )
                
                ifNode := ast.Node{
                    Type: ast.NodeTypeStmtAltIf,
                    Flag: ast.NodeFlagAltSyntax,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).List.Push(ifNode)
                yylex.(*Parser).List.Push($3)
                yylex.(*Parser).List.Push(stmtNode)
                yylex.(*Parser).List.Push()

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   alt_if_stmt_without_else T_ELSEIF '(' expr ')' ':' inner_statement_list
            {
                // Statement
                
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                )

                stmtNode := ast.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Group: ast.NodeGroupStmt,
                };

                // Alt else if

                $4.Group = ast.NodeGroupCond

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($4, stmtNode),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupElseIf, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $5.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCond, $6.HiddenTokens )
                
                elseIfNode := ast.Node{
                    Type: ast.NodeTypeStmtAltElseIf,
                    Flag: ast.NodeFlagAltSyntax,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, children),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).List.Add(elseIfNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

alt_if_stmt:
        alt_if_stmt_without_else T_ENDIF ';'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupElseIf, ast.NodeGroupStmt, ast.NodeGroupCond),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupAltEnd, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$3} )

                $$ = yylex.(*Parser).List.Pop(ast.NodeGroupNil)[0]

                $$.Position = yylex.(*Parser).NewPosition(yylex.(*Parser).Nodes($$), []*scanner.Token{$3}, nil)
                $$.Children = children

                // TODO: rewrite to prevent memory allocations
                $$.Tokens = append( 
                    $$.Tokens, 
                    yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() )...,
                )

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   alt_if_stmt_without_else T_ELSE ':' inner_statement_list T_ENDIF ';'
            {
                // Statement
                
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                )

                stmtNode := ast.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Group: ast.NodeGroupStmt,
                };

                // Alt else

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes( stmtNode ),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupElse, $3.HiddenTokens )
                
                elseNode := ast.Node{
                    Type: ast.NodeTypeStmtAltElse,
                    Flag: ast.NodeFlagAltSyntax,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, children),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupElse,
                };

                // If

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupElseIf, ast.NodeGroupStmt, ast.NodeGroupCond),
                    yylex.(*Parser).Nodes(elseNode),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $5.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupAltEnd, $6.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$6} )

                $$ = yylex.(*Parser).List.Pop(ast.NodeGroupNil)[0]

                $$.Position = yylex.(*Parser).NewPosition(yylex.(*Parser).Nodes($$), []*scanner.Token{$6}, nil)
                $$.Children = children

                // TODO: rewrite to prevent memory allocations
                $$.Tokens = append( 
                    $$.Tokens, 
                    yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() )...,
                )

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
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_parameter_list ',' { yylex.(*Parser).PrependToken($2) } parameter
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

parameter:
        optional_type is_reference is_variadic T_VARIABLE
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                var children []ast.Node
                if $1.Type != ast.NodeTypeNil {
                    $1.Group = ast.NodeGroupVarType
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes($1, identifierNode),
                    );
                } else {
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes(identifierNode),
                    );
                };

                var f bool
                if $1.Type != ast.NodeTypeNil {
                    f = true
                };

                if f && $2 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupOptionalType, $2.HiddenTokens )
                    f = true
                } else if $2 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )
                    f = true
                };

                if f && $3 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupAmpersand, $3.HiddenTokens )
                    f = true
                } else if $3 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupStart, $3.HiddenTokens )
                    f = true
                };

                if f {
                    yylex.(*Parser).PushTokens( ast.TokenGroupVariadic, $4.HiddenTokens )
                    f = true
                } else {
                    yylex.(*Parser).PushTokens( ast.TokenGroupStart, $4.HiddenTokens )
                    f = true
                };

                var position ast.Position
                if $1.Type != ast.NodeTypeNil {
                    position = yylex.(*Parser).NewPosition(children, []*scanner.Token{$4}, nil)
                } else if $2 != nil {
                    position = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2, $4}, nil)
                } else if $3 != nil {
                    position = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3, $4}, nil)
                } else {
                    position = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil)
                };

                var flag ast.NodeFlag
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                }
                if $3 != nil {
                    flag = flag | ast.NodeFlagVariadic
                };

                $$ = ast.Node{
                    Type: ast.NodeTypeParameter,
                    Flag: flag,
                    Position: position,
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   optional_type is_reference is_variadic T_VARIABLE '=' expr
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                $6.Group = ast.NodeGroupDefaultValue
                var children []ast.Node
                if $1.Type != ast.NodeTypeNil {
                    $1.Group = ast.NodeGroupVarType
                        children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes($1, identifierNode, $6),
                    )
                } else {
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes(identifierNode, $6),
                    )
                };

                var f bool
                if $1.Type != ast.NodeTypeNil {
                    f = true
                };

                if f && $2 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupOptionalType, $2.HiddenTokens )
                    f = true
                } else if $2 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )
                    f = true
                };

                if f && $3 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupAmpersand, $3.HiddenTokens )
                    f = true
                } else if $3 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupStart, $3.HiddenTokens )
                    f = true
                };

                if f {
                    yylex.(*Parser).PushTokens( ast.TokenGroupVariadic, $4.HiddenTokens )
                    f = true
                } else {
                    yylex.(*Parser).PushTokens( ast.TokenGroupStart, $4.HiddenTokens )
                    f = true
                };
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $5.HiddenTokens )

                var position ast.Position
                if $1.Type != ast.NodeTypeNil {
                    position = yylex.(*Parser).NewPosition(children, nil, nil)
                } else if $2 != nil {
                    position = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, children)
                } else if $3 != nil {
                    position = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, children)
                } else {
                    position = yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, children)
                };

                var flag ast.NodeFlag
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                };
                if $3 != nil {
                    flag = flag | ast.NodeFlagVariadic
                };

                $$ = ast.Node{
                    Type: ast.NodeTypeParameter,
                    Flag: flag,
                    Position: position,
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

optional_type:
        /* empty */
            {
                $$ = ast.Node{};

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
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeNullable,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

type:
        T_ARRAY
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CALLABLE
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                $$ = ast.Node{};

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' type_expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtReturnType,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

argument_list:
        '(' ')'
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupArgumentList, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeArgumentList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' non_empty_argument_list possible_comma ')'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupArguments),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                if $3 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupArgumentList, $3.HiddenTokens )
                    yylex.(*Parser).PushTokens( ast.TokenGroupArgumentList, []scanner.Token{*$3} )
                }
                yylex.(*Parser).PushTokens( ast.TokenGroupArgumentList, $4.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeArgumentList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_argument_list:
        argument
            {
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_argument_list ',' { yylex.(*Parser).PrependToken($2) } argument
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

argument:
        expr
            {
                $1.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1),
                )

                $$ = ast.Node{
                    Type: ast.NodeTypeArgument,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ELLIPSIS expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeArgument,
                            Flag: ast.NodeFlagVariadic,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

global_var_list:
        global_var_list ',' { yylex.(*Parser).PrependToken($2) } global_var
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   global_var
            {
                yylex.(*Parser).List.Push($1)

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
        static_var_list ',' { yylex.(*Parser).PrependToken($2) } static_var
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_var
            {
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_var:
        T_VARIABLE
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                );
                
                variableNode := ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Group: ast.NodeGroupVar,
                };

                // Static var

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(variableNode),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtStaticVar,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '=' expr
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                );
                
                variableNode := ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Group: ast.NodeGroupVar,
                };

                // Property

                $3.Group = ast.NodeGroupExpr
                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(variableNode, $3),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtStaticVar,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupProperties, ast.NodeGroupModifiers),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupPropertyList, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$3} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtPropertyList,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   method_modifiers T_CONST class_const_list ';'
            {
                childrenConstants := yylex.(*Parser).List.Pop(ast.NodeGroupConsts)
                childrenModifiers := yylex.(*Parser).List.Pop(ast.NodeGroupModifiers)

                children := yylex.(*Parser).Ast.AddNodes(
                    childrenModifiers, childrenConstants,
                );

                if len(childrenModifiers) > 0 {
                    yylex.(*Parser).PushTokens( ast.TokenGroupModifierList, $2.HiddenTokens )
                } else {
                    yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )
                };
                yylex.(*Parser).PushTokens( ast.TokenGroupConstList, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$4} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtClassConstList,
                    Position: yylex.(*Parser).NewPosition(childrenModifiers, []*scanner.Token{$2, $4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE name_list trait_adaptations
            {
                $3.Group = ast.NodeGroupTraitAdaptationList
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupTraits),
                    yylex.(*Parser).Nodes($3),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtTraitUse,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   method_modifiers T_FUNCTION returns_ref identifier backup_doc_comment '(' parameter_list ')' return_type method_body
            {
                // Identifier

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $4.HiddenTokens )

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil),
                    Group: ast.NodeGroupMethodName,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                // Class method

                childrenParams := yylex.(*Parser).List.Pop(ast.NodeGroupParams)
                childrenModifiers := yylex.(*Parser).List.Pop(ast.NodeGroupModifiers)

                nodes := make([]ast.Node, 0, 2)
                if $9.Type != ast.NodeTypeNil {
                    nodes = append(nodes, $9)
                    $9.Group = ast.NodeGroupReturnType
                };
                
                $10.Group = ast.NodeGroupStmt
                nodes = append(nodes, $10)
                
                children := yylex.(*Parser).Ast.AddNodes(
                    childrenModifiers, yylex.(*Parser).Nodes(identifierNode), childrenParams, nodes,
                );

                if len(childrenModifiers) > 0 {
                    yylex.(*Parser).PushTokens( ast.TokenGroupModifierList, $2.HiddenTokens )
                } else {
                    yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )
                };
                if $3 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupFunction, $3.HiddenTokens )
                };
                yylex.(*Parser).PushTokens( ast.TokenGroupName, $6.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupParameterList, $8.HiddenTokens )

                var flag ast.NodeFlag
                if $3 != nil {
                    flag = flag | ast.NodeFlagRef
                };
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtClassMethod,
                    Flag: flag,
                    Position: yylex.(*Parser).NewPosition(childrenModifiers, []*scanner.Token{$2}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

name_list:
        name
            {
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   name_list ',' { yylex.(*Parser).PrependToken($2) } name
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptations:
        ';'
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$1} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtNop,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' '}'
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupAdaptationList, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtTraitAdaptationList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' trait_adaptation_list '}'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupAdaptations),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupAdaptationList, $3.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtTraitAdaptationList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptation_list:
        trait_adaptation
            {
                yylex.(*Parser).List.Push($1)

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
                $1.Group = ast.NodeGroupStmt
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$2} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeWrapper,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$2}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_alias ';'
            {
                $1.Group = ast.NodeGroupStmt
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$2} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeWrapper,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$2}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_precedence:
        absolute_trait_method_reference T_INSTEADOF name_list
            {
                $1.Group = ast.NodeGroupRef
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1),
                    yylex.(*Parser).List.Pop(ast.NodeGroupInsteadof),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupRef, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtTraitUsePrecedence,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_alias:
        trait_method_reference T_AS T_STRING
            {
                // Identifier

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $3.HiddenTokens )

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil),
                    Group: ast.NodeGroupAlias,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                // Alias

                $1.Group = ast.NodeGroupRef
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, identifierNode),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupRef, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtTraitUseAlias,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_method_reference T_AS reserved_non_modifiers
            {
                // Identifier

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $3.HiddenTokens )

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil),
                    Group: ast.NodeGroupAlias,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                // Alias

                $1.Group = ast.NodeGroupRef
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, identifierNode),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupRef, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtTraitUseAlias,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_method_reference T_AS member_modifier identifier
            {
                // Identifier

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $4.HiddenTokens )

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$4}, nil),
                    Group: ast.NodeGroupAlias,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                // Alias

                $1.Group = ast.NodeGroupRef
                $3.Group = ast.NodeGroupModifier
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3, identifierNode),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupRef, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtTraitUseAlias,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_method_reference T_AS member_modifier
            {

                $1.Group = ast.NodeGroupRef
                $3.Group = ast.NodeGroupModifier
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupRef, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtTraitUseAlias,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_method_reference:
        identifier
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Group: ast.NodeGroupMethod,
                };

                // Trait reference

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtTraitMethodRef,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                // Identifier

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $3.HiddenTokens )

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil),
                    Group: ast.NodeGroupMethod,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                // Trait reference

                $1.Group = ast.NodeGroupTrait
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, identifierNode),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupName, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtTraitMethodRef,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

method_body:
        ';' /* abstract method */ 
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupSemiColon, []scanner.Token{*$1} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtNop,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' inner_statement_list '}'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $3.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtStmtList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).List.Push(identifierNode)

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
                yylex.(*Parser).List.Push($1)

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
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PROTECTED
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PRIVATE
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ABSTRACT
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINAL
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

property_list:
        property_list ',' { yylex.(*Parser).PrependToken($2) } property
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   property
            {
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

property:
        T_VARIABLE backup_doc_comment
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                );
                
                variableNode := ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Group: ast.NodeGroupVar,
                };

                // Property

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(variableNode),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtProperty,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '=' expr backup_doc_comment
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                );
                
                variableNode := ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Group: ast.NodeGroupVar,
                };

                // Property

                $3.Group = ast.NodeGroupExpr
                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(variableNode, $3),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtProperty,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_const_list:
        class_const_list ',' { yylex.(*Parser).PrependToken($2) } class_const_decl
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_const_decl
            {
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_const_decl:
        identifier '=' expr backup_doc_comment
            {
                // identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupConstantName,
                };

                // const

                $3.Group = ast.NodeGroupExpr
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupName, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtConstant,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

const_decl:
        T_STRING '=' expr backup_doc_comment
            {
                // identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupConstantName,
                };

                // const

                $3.Group = ast.NodeGroupExpr
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupName, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeStmtConstant,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

echo_expr_list:
        echo_expr_list ',' { yylex.(*Parser).PrependToken($2) } echo_expr
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   echo_expr
            {
                yylex.(*Parser).List.Push($1)

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
        non_empty_for_exprs ',' { yylex.(*Parser).PrependToken($2) } expr
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

anonymous_class:
        T_CLASS ctor_arguments extends_from implements_list backup_doc_comment '{' class_statement_list '}'
            {
                nodes := make([]ast.Node, 0, 3)

                if $2.Type != ast.NodeTypeNil {
                    $2.Group = ast.NodeGroupArgumentList
                    nodes = append(nodes, $2)
                };

                if $3.Type != ast.NodeTypeNil {
                    $3.Group = ast.NodeGroupExtends
                    nodes = append(nodes, $3)
                };

                if $4.Type != ast.NodeTypeNil {
                    $4.Group = ast.NodeGroupImplements
                    nodes = append(nodes, $4)
                };

                children := yylex.(*Parser).Ast.AddNodes(
                    nodes, yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupName, $6.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $8.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeStmtClass,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $8}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

new_expr:
        T_NEW class_name_reference ctor_arguments
            {
                nodes := make([]ast.Node, 0, 2)

                $2.Group = ast.NodeGroupClass
                nodes = append(nodes, $2)

                if $3.Type != ast.NodeTypeNil {
                    $3.Group = ast.NodeGroupArgumentList
                    nodes = append(nodes, $3)
                };

                children := yylex.(*Parser).Ast.AddNodes(
                    nodes,
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprNew,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NEW anonymous_class
            {
                $2.Group = ast.NodeGroupClass

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                );

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprNew,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

expr_without_variable:
        T_LIST '(' array_pair_list ')' '=' expr
            {
                // Array pair list

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupItems),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupList, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupArrayPairList, $4.HiddenTokens )
                
                listNode := ast.Node{
                    Type: ast.NodeTypeExprList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVar,
                };

                // Assign

                $6.Group = ast.NodeGroupExpr

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(listNode, $6),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $5.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignAssign,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' array_pair_list ']' '=' expr
            {
                // Array pair list

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupItems),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupArrayPairList, $3.HiddenTokens )
                
                listNode := ast.Node{
                    Type: ast.NodeTypeExprShortList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVar,
                };

                // Assign

                $5.Group = ast.NodeGroupExpr

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(listNode, $5),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $4.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignAssign,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' expr
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignAssign,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' '&' expr
            {
                $1.Group = ast.NodeGroupVar
                $4.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $4),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEqual, $3.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignReference,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLONE expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprClone,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_PLUS_EQUAL expr
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignPlus,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MINUS_EQUAL expr
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignMinus,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MUL_EQUAL expr
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignMul,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_POW_EQUAL expr
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignPow,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_DIV_EQUAL expr
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignDiv,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_CONCAT_EQUAL expr
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignConcat,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MOD_EQUAL expr
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignMod,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };
                
                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_AND_EQUAL expr
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignBitwiseAnd,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_OR_EQUAL expr
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignBitwiseOr,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_XOR_EQUAL expr
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignBitwiseXor,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_SL_EQUAL expr
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignShiftLeft,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_SR_EQUAL expr
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeAssignShiftRight,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_INC
            {
                $1.Group = ast.NodeGroupVar

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprPostInc,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$2}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INC variable
            {
                $2.Group = ast.NodeGroupVar

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprPreInc,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_DEC
            {
                $1.Group = ast.NodeGroupVar

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprPostDec,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$2}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DEC variable
            {
                $2.Group = ast.NodeGroupVar

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprPreDec,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_BOOLEAN_OR expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryBooleanOr,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_BOOLEAN_AND expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryBooleanAnd,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_OR expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryLogicalOr,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_AND expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryLogicalAnd,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_XOR expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryLogicalXor,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '|' expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryBitwiseOr,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '&' expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryBitwiseAnd,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '^' expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryBitwiseXor,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '.' expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryConcat,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '+' expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryPlus,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '-' expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryMinus,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '*' expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryMul,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_POW expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryPow,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '/' expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryDiv,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '%' expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryMod,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SL expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryShiftLeft,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SR expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryShiftRight,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '+' expr %prec T_INC
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprUnaryPlus,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '-' expr %prec T_INC
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprUnaryMinus,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '!' expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprBooleanNot,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '~' expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprBitwiseNot,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_IDENTICAL expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryIdentical,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_NOT_IDENTICAL expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryNotIdentical,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_EQUAL expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryEqual,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_NOT_EQUAL expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEqual, []scanner.Token{*$2} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryNotEqual,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '<' expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinarySmaller,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_SMALLER_OR_EQUAL expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinarySmallerOrEqual,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '>' expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryGreater,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_GREATER_OR_EQUAL expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryGreaterOrEqual,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SPACESHIP expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinarySpaceship,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_INSTANCEOF class_name_reference
            {
                $1.Group = ast.NodeGroupExpr
                $3.Group = ast.NodeGroupClass

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprInstanceOf,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' expr ')'
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$1} )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeWrapper,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_expr
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '?' expr ':' expr
            {
                $1.Group = ast.NodeGroupCond
                $3.Group = ast.NodeGroupIfTrue
                $5.Group = ast.NodeGroupIfFalse

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3, $5),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupCond, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupTrue, $4.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprTernary,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '?' ':' expr
            {
                $1.Group = ast.NodeGroupCond
                $4.Group = ast.NodeGroupIfFalse

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $4),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupCond, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupTrue, $3.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprTernary,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_COALESCE expr
            {
                $1.Group = ast.NodeGroupLeft
                $3.Group = ast.NodeGroupRight

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeBinaryCoalesce,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   internal_functions_in_yacc
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INT_CAST expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCast, []scanner.Token{*$1} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeCastInt,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOUBLE_CAST expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCast, []scanner.Token{*$1} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeCastDouble,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STRING_CAST expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCast, []scanner.Token{*$1} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeCastString,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ARRAY_CAST expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCast, []scanner.Token{*$1} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeCastArray,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_OBJECT_CAST expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCast, []scanner.Token{*$1} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeCastObject,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BOOL_CAST expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCast, []scanner.Token{*$1} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeCastBool,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_UNSET_CAST expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupCast, []scanner.Token{*$1} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeCastUnset,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXIT exit_expr
            {
                var children []ast.Node

                if $2.Type != ast.NodeTypeNil {
                    $2.Group = ast.NodeGroupExpr
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes($2),
                    )
                };

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                var flag ast.NodeFlag
                exitTknValue := yylex.(*Parser).FileData[$1.StartPos:$1.EndPos]
                if bytes.EqualFold(exitTknValue, []byte("die")) {
                    flag = ast.NodeFlagAltSyntax
                };

                $$ = ast.Node{
                    Type: ast.NodeTypeExprExit,
                    Flag: flag,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '@' expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprErrorSuppress,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   scalar
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '`' backticks_expr '`'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupParts),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprShellExec,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PRINT expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprPrint,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprYield,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD expr
            {
                $2.Group = ast.NodeGroupVal

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprYield,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD expr T_DOUBLE_ARROW expr
            {
                $2.Group = ast.NodeGroupKey
                $4.Group = ast.NodeGroupVal

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2, $4),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $3.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprYield,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD_FROM expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$1} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprYieldFrom,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                stmts := yylex.(*Parser).List.Pop(ast.NodeGroupStmts)
                params := yylex.(*Parser).List.Pop(ast.NodeGroupParams)
                nodes := make([]ast.Node, 0, 2)

                if $7.Type != ast.NodeTypeNil {
                    $7.Group = ast.NodeGroupClosureUse
                    nodes = append(nodes, $7)
                };
                
                if $8.Type != ast.NodeTypeNil {
                    $8.Group = ast.NodeGroupReturnType
                    nodes = append(nodes, $8)
                };

                children := yylex.(*Parser).Ast.AddNodes(
                    params, nodes, stmts,
                );

                var flag ast.NodeFlag
                if $2 != nil {
                    flag = flag | ast.NodeFlagRef
                };

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                if $2 == nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupFunction, $4.HiddenTokens )
                } else {
                    yylex.(*Parser).PushTokens( ast.TokenGroupFunction, $2.HiddenTokens )
                    yylex.(*Parser).PushTokens( ast.TokenGroupAmpersand, $4.HiddenTokens )
                };
                yylex.(*Parser).PushTokens( ast.TokenGroupParameterList, $6.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupReturnType, $9.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $11.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprClosure,
                    Flag: flag,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $11}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                stmts := yylex.(*Parser).List.Pop(ast.NodeGroupStmts)
                params := yylex.(*Parser).List.Pop(ast.NodeGroupParams)
                nodes := make([]ast.Node, 0, 2)

                if $8.Type != ast.NodeTypeNil {
                    $8.Group = ast.NodeGroupClosureUse
                    nodes = append(nodes, $8)
                };
                
                if $9.Type != ast.NodeTypeNil {
                    $9.Group = ast.NodeGroupReturnType
                    nodes = append(nodes, $9)
                };

                children := yylex.(*Parser).Ast.AddNodes(
                    params, nodes, stmts,
                );

                flag := ast.NodeFlagStatic
                if $3 != nil {
                    flag = flag | ast.NodeFlagRef
                };

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStatic, $2.HiddenTokens )
                if $3 == nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupFunction, $5.HiddenTokens )
                } else {
                    yylex.(*Parser).PushTokens( ast.TokenGroupFunction, $3.HiddenTokens )
                    yylex.(*Parser).PushTokens( ast.TokenGroupAmpersand, $5.HiddenTokens )
                };
                yylex.(*Parser).PushTokens( ast.TokenGroupParameterList, $7.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupReturnType, $10.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStmts, $12.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprClosure,
                    Flag: flag,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $12}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                $$ = ast.Node{};

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE '(' lexical_var_list ')'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupStmts),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupUse, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupLexicalVarList, $4.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprClosureUse,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

lexical_var_list:
        lexical_var_list ',' { yylex.(*Parser).PrependToken($2) } lexical_var
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   lexical_var
            {
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

lexical_var:
    T_VARIABLE
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' T_VARIABLE
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )

                variableNode := ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVar,
                };

                // Reference 

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(variableNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprReference,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVal,
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

function_call:
        name argument_list
            {
                $1.Group = ast.NodeGroupFunction
                $2.Group = ast.NodeGroupArgumentList
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $2),
                )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprFunctionCall,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
            {
                $1.Group = ast.NodeGroupClass
                $3.Group = ast.NodeGroupCall
                $4.Group = ast.NodeGroupArgumentList
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3, $4),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupName, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprStaticCall,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
            {
                $1.Group = ast.NodeGroupClass
                $3.Group = ast.NodeGroupCall
                $4.Group = ast.NodeGroupArgumentList
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3, $4),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupName, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprStaticCall,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   callable_expr argument_list
            {
                $1.Group = ast.NodeGroupFunction
                $2.Group = ast.NodeGroupArgumentList
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $2),
                )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprFunctionCall,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_name:
        T_STATIC
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                $$ = ast.Node{};

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' optional_expr ')'
            {
                var children []ast.Node

                if $2.Type != ast.NodeTypeNil {
                    $2.Group = ast.NodeGroupExpr
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes($2),
                    )
                };

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$1} )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeWrapper,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                stringPartNode := ast.Node{
                    Type: ast.NodeTypeScalarEncapsedStringPart,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Group: ast.NodeGroupParts,
                };

                yylex.(*Parser).List.Push(stringPartNode)

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
                $$ = ast.Node{};

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
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupItems),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupArray, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupArrayPairList, $4.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprArray,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' array_pair_list ']'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupItems),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupArrayPairList, $3.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprShortArray,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONSTANT_ENCAPSED_STRING
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarString,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

scalar:
        T_LNUMBER
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarLnumber,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DNUMBER
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarDnumber,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LINE
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FILE
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DIR
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRAIT_C
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_METHOD_C
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FUNC_C
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_C
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLASS_C
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarMagicConstant,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC
            {
                // String part

                stringPartNode := ast.Node{
                    Type: ast.NodeTypeScalarEncapsedStringPart,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil),
                    Group: ast.NodeGroupParts,
                };

                // HEREDOC

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(stringPartNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarHeredoc,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC T_END_HEREDOC
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarHeredoc,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '"' encaps_list '"'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupParts),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarEncapsed,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC encaps_list T_END_HEREDOC
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupParts),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarHeredoc,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                $1.Group = ast.NodeGroupConstant
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1),
                )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprConstFetch,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM identifier
            {
                // Identifier

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $3.HiddenTokens )

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupConstantName,
                };

                // Constant fetch

                $1.Group = ast.NodeGroupClass
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, identifierNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupName, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprClassConstFetch,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM identifier
            {
                // Identifier

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $3.HiddenTokens )

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupConstantName,
                };

                // Constant fetch

                $1.Group = ast.NodeGroupClass
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, identifierNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupName, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprClassConstFetch,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                $$ = ast.Node{};

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
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$1} )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeWrapper,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$1} )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeWrapper,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupDim

                var children []ast.Node
                if $3.Type == ast.NodeTypeNil {
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes($1),
                    )
                } else {
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes($1, $3),
                    )
                };

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, []scanner.Token{*$2} )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, []scanner.Token{*$4} )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprArrayDimFetch,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   constant '[' optional_expr ']'
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupDim

                var children []ast.Node
                if $3.Type == ast.NodeTypeNil {
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes($1),
                    )
                } else {
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes($1, $3),
                    )
                };

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, []scanner.Token{*$2} )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, []scanner.Token{*$4} )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprArrayDimFetch,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   dereferencable '{' expr '}'
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupDim
                
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )
                
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, []scanner.Token{*$2} )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, []scanner.Token{*$4} )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprArrayDimFetch,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   dereferencable T_OBJECT_OPERATOR property_name argument_list
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupMethod
                $4.Group = ast.NodeGroupArgumentList

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3, $4),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprMethodCall,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupProperty

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprPropertyFetch,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

simple_variable:
        T_VARIABLE
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' '{' expr '}'
            {
                $3.Group = ast.NodeGroupVarName
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$1} )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$2} )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, []scanner.Token{*$4} )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' simple_variable
            {
                $2.Group = ast.NodeGroupVarName
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$1} )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_member:
        class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $1.Group = ast.NodeGroupClass
                $3.Group = ast.NodeGroupProperty
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupName, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprStaticPropertyFetch,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $1.Group = ast.NodeGroupClass
                $3.Group = ast.NodeGroupProperty
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupName, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprStaticPropertyFetch,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupDim

                var children []ast.Node
                if $3.Type == ast.NodeTypeNil {
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes($1),
                    )
                } else {
                    children = yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes($1, $3),
                    )
                };

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, []scanner.Token{*$2} )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, []scanner.Token{*$4} )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprArrayDimFetch,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_variable '{' expr '}'
            {
                $1.Group = ast.NodeGroupVar
                $3.Group = ast.NodeGroupDim
                
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )
                
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, []scanner.Token{*$2} )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, []scanner.Token{*$4} )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprArrayDimFetch,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_variable T_OBJECT_OPERATOR property_name
            {
                $1.Group = ast.NodeGroupClass
                $3.Group = ast.NodeGroupProperty

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprPropertyFetch,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $1.Group = ast.NodeGroupClass
                $3.Group = ast.NodeGroupProperty
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupName, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprStaticPropertyFetch,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_variable T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $1.Group = ast.NodeGroupClass
                $3.Group = ast.NodeGroupProperty
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupName, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprStaticPropertyFetch,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

member_name:
        identifier
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' expr '}'
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$1} )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeWrapper,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' expr '}'
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$1} )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, []scanner.Token{*$3} )

                $$ = ast.Node{
                    Type: ast.NodeTypeWrapper,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                items := yylex.(*Parser).List.Pop(ast.NodeGroupItems)

                if len(items) == 1 && items[0].Position.PS == 0 && items[0].Position.PE == 0 {
                    yylex.(*Parser).List.Push()
                } else {
                    yylex.(*Parser).List.Push(items...)
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

possible_array_pair:
        /* empty */
            {
                $$ = ast.Node{
                    Type: ast.NodeTypeExprArrayItem,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   array_pair
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_array_pair_list:
        non_empty_array_pair_list ',' { yylex.(*Parser).PrependToken($2) } possible_array_pair
            {
                yylex.(*Parser).List.Add($4)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   possible_array_pair
            {
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

array_pair:
        expr T_DOUBLE_ARROW expr
            {
                $1.Group = ast.NodeGroupKey
                $3.Group = ast.NodeGroupVal

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprArrayItem,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                $1.Group = ast.NodeGroupVal

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1),
                )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprArrayItem,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_DOUBLE_ARROW '&' variable
            {
                // Reference

                $4.Group = ast.NodeGroupVar

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($4),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $3.HiddenTokens )
                
                refNode := ast.Node{
                    Type: ast.NodeTypeExprReference,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVal,
                };

                // Array item

                $1.Group = ast.NodeGroupKey

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, refNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprArrayItem,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' variable
            {
                $2.Group = ast.NodeGroupVar

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )
                
                refNode := ast.Node{
                    Type: ast.NodeTypeExprReference,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVal,
                };

                // Array item

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(refNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprArrayItem,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_DOUBLE_ARROW T_LIST '(' array_pair_list ')'
            {
                // List

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupItems),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupList, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupArrayPairList, $6.HiddenTokens )

                listNode := ast.Node{
                    Type: ast.NodeTypeExprList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3, $6}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVal,
                };

                // Array item

                $1.Group = ast.NodeGroupKey

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($1, listNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $2.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprArrayItem,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$6}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LIST '(' array_pair_list ')'
            {
                // List

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupItems),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupList, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupArrayPairList, $4.HiddenTokens )

                listNode := ast.Node{
                    Type: ast.NodeTypeExprList,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVal,
                };

                // Array item
                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(listNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprArrayItem,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$1, $4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

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
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )
                
                encapsNode := ast.Node{
                    Type: ast.NodeTypeScalarEncapsedStringPart,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).List.Add(encapsNode)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   encaps_var
            {
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ENCAPSED_AND_WHITESPACE encaps_var
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                encapsNode := ast.Node{
                    Type: ast.NodeTypeScalarEncapsedStringPart,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).List.Push(encapsNode, $2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_var:
        T_VARIABLE
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '[' encaps_var_offset ']'
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                variableNode := ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVar,
                };

                // Fetch

                $3.Group = ast.NodeGroupDim
                
                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(variableNode, $3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, []scanner.Token{*$2} )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, []scanner.Token{*$4} )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprArrayDimFetch,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE T_OBJECT_OPERATOR T_STRING
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                variableNode := ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVar,
                };

                // Property

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $3.HiddenTokens )

                propertyNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$3}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupProperty,
                };

                // Fetch

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(variableNode, propertyNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $2.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprPropertyFetch,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES expr '}'
            {
                // Variable

                $2.Group = ast.NodeGroupVarName

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                variableNode := ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Group: ast.NodeGroupExpr,
                };

                // Wrapper

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(variableNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$1} )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, []scanner.Token{*$3} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeWrapper,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}'
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )

                variableNode := ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupExpr,
                };

                // Wrapper

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(variableNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$1} )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, []scanner.Token{*$3} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeWrapper,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $2.HiddenTokens )

                variableNode := ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVar,
                };

                // Fetch

                $4.Group = ast.NodeGroupDim
                
                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(variableNode, $4),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, []scanner.Token{*$3} )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $5.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, []scanner.Token{*$5} )

                fetchNode := ast.Node{
                    Type: ast.NodeTypeExprArrayDimFetch,
                    Position: yylex.(*Parser).NewPosition(children, []*scanner.Token{$5}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupExpr,
                };

                // Wrapper

                children = yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(fetchNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$1} )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $6.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, []scanner.Token{*$6} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeWrapper,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $6}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CURLY_OPEN variable '}'
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupStart, []scanner.Token{*$1} )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, $3.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEnd, []scanner.Token{*$3} )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeWrapper,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $3}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_var_offset:
        T_STRING
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeScalarString,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NUM_STRING
            {
                nodeType := ast.NodeTypeScalarString
                tknValue := yylex.(*Parser).FileData[$1.StartPos:$1.EndPos]
                if _, err := strconv.Atoi(string(tknValue)); err == nil {
                    nodeType = ast.NodeTypeScalarLnumber
                };

                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: nodeType,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '-' T_NUM_STRING
            {
                yylex.(*Parser).PushTokens( ast.TokenGroupVar, $1.HiddenTokens )

                tknValue := yylex.(*Parser).FileData[$2.StartPos:$2.EndPos]
                if _, err := strconv.Atoi(string(tknValue)); err == nil {
                    lnumberNode := ast.Node{
                        Type: ast.NodeTypeScalarLnumber,
                        Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$2}, nil),
                        Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                        Group: ast.NodeGroupExpr,
                    };
                    
                    children := yylex.(*Parser).Ast.AddNodes(
                        yylex.(*Parser).Nodes(lnumberNode),
                    )

                    $$ = ast.Node{
                        Type: ast.NodeTypeExprUnaryMinus,
                        Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil),
                        Children: children,
                        Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    };
                } else {
                    $$ = ast.Node{
                        Type: ast.NodeTypeScalarString,
                        Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $2}, nil),
                        Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    };
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                // Identifier

                identifierNode := ast.Node{
                    Type: ast.NodeTypeIdentifier,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, nil),
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                    Group: ast.NodeGroupVarName,
                };

                // Variable

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes(identifierNode),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )

                $$ = ast.Node{
                    Type: ast.NodeTypeExprVariable,
                    Position: yylex.(*Parser).NewPosition(children, nil, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

internal_functions_in_yacc:
        T_ISSET '(' isset_variables possible_comma ')'
            {
                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).List.Pop(ast.NodeGroupVars),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupIsset, $2.HiddenTokens )
                if $4 != nil {
                    yylex.(*Parser).PushTokens( ast.TokenGroupVarList, $4.HiddenTokens )
                    yylex.(*Parser).PushTokens( ast.TokenGroupVarList, []scanner.Token{*$4} )
                };
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprIsset,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $5}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EMPTY '(' expr ')'
            {
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEmpty, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprEmpty,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INCLUDE expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprInclude,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INCLUDE_ONCE expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprIncludeOnce,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EVAL '(' expr ')'
            {
                $3.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($3),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupEval, $2.HiddenTokens )
                yylex.(*Parser).PushTokens( ast.TokenGroupExpr, $4.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprEval,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1, $4}, nil),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_REQUIRE expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprRequire,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_REQUIRE_ONCE expr
            {
                $2.Group = ast.NodeGroupExpr

                children := yylex.(*Parser).Ast.AddNodes(
                    yylex.(*Parser).Nodes($2),
                )

                yylex.(*Parser).PushTokens( ast.TokenGroupStart, $1.HiddenTokens )
                
                $$ = ast.Node{
                    Type: ast.NodeTypeExprRequireOnce,
                    Position: yylex.(*Parser).NewPosition(nil, []*scanner.Token{$1}, children),
                    Children: children,
                    Tokens: yylex.(*Parser).Ast.AddTokens( yylex.(*Parser).PopTokens() ),
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

isset_variables:
        isset_variable
            {
                yylex.(*Parser).List.Push($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   isset_variables ',' { yylex.(*Parser).PrependToken($2) } isset_variable
            {
                yylex.(*Parser).List.Add($4)

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
