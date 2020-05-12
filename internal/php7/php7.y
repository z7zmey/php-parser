%{
package php7

import (
    "bytes"
    "strconv"

    "github.com/z7zmey/php-parser/pkg/token"
    "github.com/z7zmey/php-parser/internal/scanner"
    "github.com/z7zmey/php-parser/pkg/ast"
)

%}

%union{
    node             ast.Vertex
    token            *scanner.Token
    list             []ast.Vertex
    str              string

    ClassExtends     *ast.StmtClassExtends
    ClassImplements  *ast.StmtClassImplements
    InterfaceExtends *ast.StmtInterfaceExtends
    ClosureUse       *ast.ExprClosureUse
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

%type <str> backup_doc_comment

%%

/////////////////////////////////////////////////////////////////////////

start:
        top_statement_list
            {
                yylex.(*Parser).rootNode = &ast.Root{ast.Node{}, $1}

                // save position
                yylex.(*Parser).rootNode.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($1)

                yylex.(*Parser).setFreeFloating(yylex.(*Parser).rootNode, token.End, yylex.(*Parser).currentToken.Tokens)

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

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                $$ = []ast.Vertex{}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

namespace_name:
        T_STRING
            {
                namePart := &ast.NameNamePart{ast.Node{}, $1.Value}
                $$ = []ast.Vertex{namePart}

                // save position
                namePart.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(namePart, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_NS_SEPARATOR T_STRING
            {
                namePart := &ast.NameNamePart{ast.Node{}, $3.Value}
                $$ = append($1, namePart)

                // save position
                namePart.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(namePart, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

name:
        namespace_name
            {
                $$ = &ast.NameName{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    | T_NAMESPACE T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameRelative{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Namespace, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    | T_NS_SEPARATOR namespace_name
            {
                $$ = &ast.NameFullyQualified{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                $$ = &ast.StmtHaltCompiler{ast.Node{}}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.HaltCompiller, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.OpenParenthesisToken, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.CloseParenthesisToken, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE namespace_name ';'
            {
                name := &ast.NameName{ast.Node{}, $2}
                $$ = &ast.StmtNamespace{ast.Node{}, name, nil}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).MoveFreeFloating($2[0], name)
                yylex.(*Parser).setFreeFloating(name, token.End, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE namespace_name '{' top_statement_list '}'
            {
                name := &ast.NameName{ast.Node{}, $2}
                $$ = &ast.StmtNamespace{ast.Node{}, name, $4}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).MoveFreeFloating($2[0], name)
                yylex.(*Parser).setFreeFloating(name, token.End, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $5.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NAMESPACE '{' top_statement_list '}'
            {
                $$ = &ast.StmtNamespace{ast.Node{}, nil, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Namespace, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE mixed_group_use_declaration ';'
            {
                $$ = $2

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.UseDeclarationList, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE use_type group_use_declaration ';'
            {
                $3.(*ast.StmtGroupUse).UseType = $2
                $$ = $3

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.UseDeclarationList, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE use_declarations ';'
            {
                $$ = &ast.StmtUseList{ast.Node{}, nil, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.UseDeclarationList, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE use_type use_declarations ';'
            {
                $$ = &ast.StmtUseList{ast.Node{}, $2, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.UseDeclarationList, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONST const_list ';'
            {
                $$ = &ast.StmtConstList{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_type:
        T_FUNCTION
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONST
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                name := &ast.NameName{ast.Node{}, $1}
                $$ = &ast.StmtGroupUse{ast.Node{}, nil, name, $4}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($1, $6)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], name)
                yylex.(*Parser).setFreeFloating(name, token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Slash, $3.Tokens)
                if $5 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.Stmts, append($5.Tokens, append(yylex.(*Parser).GetFreeFloatingToken($5), $6.Tokens...)...))
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Stmts, $6.Tokens)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
            {
                name := &ast.NameName{ast.Node{}, $2}
                $$ = &ast.StmtGroupUse{ast.Node{}, nil, name, $5}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $7)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.UseType, $1.Tokens)
                yylex.(*Parser).MoveFreeFloating($2[0], name)
                yylex.(*Parser).setFreeFloating(name, token.End, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Slash, $4.Tokens)
                if $6 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.Stmts, append($6.Tokens, append(yylex.(*Parser).GetFreeFloatingToken($6), $7.Tokens...)...))
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Stmts, $7.Tokens)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

mixed_group_use_declaration:
        namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                name := &ast.NameName{ast.Node{}, $1}
                $$ = &ast.StmtGroupUse{ast.Node{}, nil, name, $4}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($1, $6)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], name)
                yylex.(*Parser).setFreeFloating(name, token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Slash, $3.Tokens)
                if $5 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.Stmts, append($5.Tokens, append(yylex.(*Parser).GetFreeFloatingToken($5), $6.Tokens...)...))
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Stmts, $6.Tokens)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
            {
                name := &ast.NameName{ast.Node{}, $2}
                $$ = &ast.StmtGroupUse{ast.Node{}, nil, name, $5}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $7)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Use, append($1.Tokens, yylex.(*Parser).GetFreeFloatingToken($1)...))
                yylex.(*Parser).MoveFreeFloating($2[0], name)
                yylex.(*Parser).setFreeFloating(name, token.End, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Slash, $4.Tokens)
                if $6 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.Stmts, append($6.Tokens, append(yylex.(*Parser).GetFreeFloatingToken($6), $7.Tokens...)...))
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Stmts, $7.Tokens)
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
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   inline_use_declaration
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unprefixed_use_declarations:
        unprefixed_use_declarations ',' unprefixed_use_declaration
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   unprefixed_use_declaration
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_declarations:
        use_declarations ',' use_declaration
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   use_declaration
            {
                $$ = []ast.Vertex{$1}

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
                $2.(*ast.StmtUse).UseType = $1
                $$ = $2

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unprefixed_use_declaration:
        namespace_name
            {
                name := &ast.NameName{ast.Node{}, $1}
                $$ = &ast.StmtUse{ast.Node{}, nil, name, nil}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], name)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   namespace_name T_AS T_STRING
            {
                name := &ast.NameName{ast.Node{}, $1}
                alias := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtUse{ast.Node{}, nil, name, alias}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($1)
                alias.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($3)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], name)
                yylex.(*Parser).setFreeFloating(name, token.End, $2.Tokens)
                yylex.(*Parser).setFreeFloating(alias, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

use_declaration:
        unprefixed_use_declaration
            {
                $$ = $1

                // save coments
                yylex.(*Parser).MoveFreeFloating($1.(*ast.StmtUse).Use, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_SEPARATOR unprefixed_use_declaration
            {
                $$ = $2;

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Slash, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Slash, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

const_list:
        const_list ',' const_decl
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   const_decl
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
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

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   /* empty */
            {
                $$ = []ast.Vertex{}

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
                $$ = &ast.StmtHaltCompiler{ast.Node{}, }

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.HaltCompiller, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.OpenParenthesisToken, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.CloseParenthesisToken, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }

statement:
        '{' inner_statement_list '}'
            {
                $$ = &ast.StmtStmtList{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)

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
                switch n := $5.(type) {
                case *ast.StmtWhile :
                    n.Cond = $3
                case *ast.StmtAltWhile :
                    n.Cond = $3
                }

                $$ = $5

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.While, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DO statement T_WHILE '(' expr ')' ';'
            {
                $$ = &ast.StmtDo{ast.Node{}, $2, $5}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $7)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.While, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $6.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Cond, $7.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($7))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
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
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $9)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.For, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.InitExpr, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.CondExpr, $6.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.IncExpr, $8.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_SWITCH '(' expr ')' switch_case_list
            {
                switch n := $5.(type) {
                case *ast.StmtSwitch:
                    n.Cond = $3
                case *ast.StmtAltSwitch:
                    n.Cond = $3
                default:
                    panic("unexpected node type")
                }

                $$ = $5

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Switch, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BREAK optional_expr ';'
            {
                $$ = &ast.StmtBreak{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONTINUE optional_expr ';'
            {
                $$ = &ast.StmtContinue{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_RETURN optional_expr ';'
            {
                $$ = &ast.StmtReturn{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_GLOBAL global_var_list ';'
            {
                $$ = &ast.StmtGlobal{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.VarList, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC static_var_list ';'
            {
                $$ = &ast.StmtStatic{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.VarList, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ECHO echo_expr_list ';'
            {
                $$ = &ast.StmtEcho{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Echo, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INLINE_HTML
            {
                $$ = &ast.StmtInlineHtml{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr ';'
            {
                $$ = &ast.StmtExpression{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_UNSET '(' unset_variables possible_comma ')' ';' 
            {
                $$ = &ast.StmtUnset{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $6)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Unset, $2.Tokens)
                if $4 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.VarList, append($4.Tokens, append(yylex.(*Parser).GetFreeFloatingToken($4), $5.Tokens...)...))
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.VarList, $5.Tokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.CloseParenthesisToken, $6.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($6))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
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
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $7)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Foreach, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $6.Tokens)


                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
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
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $9)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Foreach, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Key, $6.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $8.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DECLARE '(' const_list ')' declare_statement
            {
                $$ = $5
                $$.(*ast.StmtDeclare).Consts = $3

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Declare, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ConstList, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ';'
            {
                $$ = &ast.StmtNop{ast.Node{}, }

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRY '{' inner_statement_list '}' catch_list finally_statement
            {
                if $6 == nil {
                    $$ = &ast.StmtTry{ast.Node{}, $3, $5, $6}
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $5)
                } else {
                    $$ = &ast.StmtTry{ast.Node{}, $3, $5, $6}
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $6)
                }

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Try, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_THROW expr ';'
            {
                $$ = &ast.StmtThrow{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_GOTO T_STRING ';'
            {
                label := &ast.Identifier{ast.Node{}, $2.Value}
                $$ = &ast.StmtGoto{ast.Node{}, label}

                // save position
                label.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(label, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Label, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STRING ':'
            {
                label := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.StmtLabel{ast.Node{}, label}

                // save position
                label.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Label, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }

catch_list:
        /* empty */
            {
                $$ = []ast.Vertex{}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   catch_list T_CATCH '(' catch_name_list T_VARIABLE ')' '{' inner_statement_list '}'
            {
                identifier := &ast.Identifier{ast.Node{}, bytes.TrimLeftFunc($5.Value, isDollar)}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                catch := &ast.StmtCatch{ast.Node{}, $4, variable, $8}
                $$ = append($1, catch)

                // save position
                identifier.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($5)
                variable.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($5)
                catch.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($2, $9)

                // save comments
                yylex.(*Parser).setFreeFloating(catch, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating(catch, token.Catch, $3.Tokens)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $5.Tokens)
                yylex.(*Parser).setFreeFloating(catch, token.Var, $6.Tokens)
                yylex.(*Parser).setFreeFloating(catch, token.Cond, $7.Tokens)
                yylex.(*Parser).setFreeFloating(catch, token.Stmts, $9.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;
catch_name_list:
        name
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   catch_name_list '|' name
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

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
                $$ = &ast.StmtFinally{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Finally, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

unset_variables:
        unset_variable
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   unset_variables ',' unset_variable
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

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
                name := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtFunction{ast.Node{}, $2 != nil, name, $6, $8, $10}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($3)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $11)


                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                if $2 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $2.Tokens)
                    yylex.(*Parser).setFreeFloating(name, token.Start, $3.Tokens)
                } else {
                    yylex.(*Parser).setFreeFloating(name, token.Start, $3.Tokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.Name, $5.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ParamList, $7.Tokens)
                if $8 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.Params, $8.GetNode().Tokens[token.Colon]); delete($8.GetNode().Tokens, token.Colon)
                }
                yylex.(*Parser).setFreeFloating($$, token.ReturnType, $9.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $11.Tokens)

                // normalize
                if $8 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Params, $$.GetNode().Tokens[token.ReturnType]); delete($$.GetNode().Tokens, token.ReturnType)
                }

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
                name := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtClass{ast.Node{}, name, $1, nil, $4, $5, $8}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($3)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewOptionalListTokensPosition($1, $2, $9)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)
                yylex.(*Parser).setFreeFloating($$, token.ModifierList, $2.Tokens)
                yylex.(*Parser).setFreeFloating(name, token.Start, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $7.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $9.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
            {
                name := &ast.Identifier{ast.Node{}, $2.Value}
                $$ = &ast.StmtClass{ast.Node{}, name, nil, nil, $3, $4, $7}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $8)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(name, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $6.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $8.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_modifiers:
        class_modifier
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_modifiers class_modifier
            {
                $$ = append($1, $2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_modifier:
        T_ABSTRACT
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINAL
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_declaration_statement:
        T_TRAIT T_STRING backup_doc_comment '{' class_statement_list '}'
            {
                name := &ast.Identifier{ast.Node{}, $2.Value}
                $$ = &ast.StmtTrait{ast.Node{}, name, $5}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $6)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(name, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $6.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

interface_declaration_statement:
        T_INTERFACE T_STRING interface_extends_list backup_doc_comment '{' class_statement_list '}'
            {
                name := &ast.Identifier{ast.Node{}, $2.Value}
                $$ = &ast.StmtInterface{ast.Node{}, name, $3, $6}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $7)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(name, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $5.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $7.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

extends_from:
        /* empty */
            {
                $$ = nil

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXTENDS name
            {
                $$ = &ast.StmtClassExtends{ast.Node{}, $2};

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

interface_extends_list:
        /* empty */
            {
                $$ = nil

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXTENDS name_list
            {
                $$ = &ast.StmtInterfaceExtends{ast.Node{}, $2};

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

implements_list:
        /* empty */
            {
                $$ = nil

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_IMPLEMENTS name_list
            {
                $$ = &ast.StmtClassImplements{ast.Node{}, $2};

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                $$ = &ast.ExprReference{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LIST '(' array_pair_list ')'
            {
                $$ = &ast.ExprList{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.List, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' array_pair_list ']'
            {
                $$ = &ast.ExprShortList{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save commentsc
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

for_statement:
        statement
            {
                $$ = &ast.StmtFor{ast.Node{}, nil, nil, nil, $1}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodePosition($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDFOR ';'
            {
                stmtList := &ast.StmtStmtList{ast.Node{}, $2}
                $$ = &ast.StmtAltFor{ast.Node{}, nil, nil, nil, stmtList}

                // save position
                stmtList.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

foreach_statement:
        statement
            {
                $$ = &ast.StmtForeach{ast.Node{}, nil, nil, nil, $1}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodePosition($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDFOREACH ';'
            {
                stmtList := &ast.StmtStmtList{ast.Node{}, $2}
                $$ = &ast.StmtAltForeach{ast.Node{}, nil, nil, nil, stmtList}

                // save position
                stmtList.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

declare_statement:
        statement
            {
                $$ = &ast.StmtDeclare{ast.Node{}, false, nil, $1}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodePosition($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDDECLARE ';'
            {
                stmtList := &ast.StmtStmtList{ast.Node{}, $2}
                $$ = &ast.StmtDeclare{ast.Node{}, true, nil, stmtList}

                // save position
                stmtList.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

switch_case_list:
        '{' case_list '}'
            {
                caseList := &ast.StmtCaseList{ast.Node{}, $2}
                $$ = &ast.StmtSwitch{ast.Node{}, nil, caseList}

                // save position
                caseList.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(caseList, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(caseList, token.CaseListEnd, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' ';' case_list '}'
            {
                caseList := &ast.StmtCaseList{ast.Node{}, $3}
                $$ = &ast.StmtSwitch{ast.Node{}, nil, caseList}

                // save position
                caseList.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating(caseList, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(caseList, token.CaseListStart, append($2.Tokens, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating(caseList, token.CaseListEnd, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' case_list T_ENDSWITCH ';'
            {
                caseList := &ast.StmtCaseList{ast.Node{}, $2}
                $$ = &ast.StmtAltSwitch{ast.Node{}, nil, caseList}

                // save position
                caseList.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.Tokens)
                yylex.(*Parser).setFreeFloating(caseList, token.CaseListEnd, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' ';' case_list T_ENDSWITCH ';'
            {

                caseList := &ast.StmtCaseList{ast.Node{}, $3}
                $$ = &ast.StmtAltSwitch{ast.Node{}, nil, caseList}

                // save position
                caseList.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($3)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.Tokens)
                yylex.(*Parser).setFreeFloating(caseList, token.CaseListStart, append($2.Tokens, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating(caseList, token.CaseListEnd, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $5.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($5))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

case_list:
        /* empty */
            {
                $$ = []ast.Vertex{}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   case_list T_CASE expr case_separator inner_statement_list
            {
                _case := &ast.StmtCase{ast.Node{}, $3, $5}
                $$ = append($1, _case)

                // save position
                _case.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($2, $5)

                // save comments
                yylex.(*Parser).setFreeFloating(_case, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating(_case, token.Expr, append($4.Tokens))
                yylex.(*Parser).setFreeFloating(_case, token.CaseSeparator, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   case_list T_DEFAULT case_separator inner_statement_list
            {
                _default := &ast.StmtDefault{ast.Node{}, $4}
                $$ = append($1, _default)

                // save position
                _default.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($2, $4)

                // save comments
                yylex.(*Parser).setFreeFloating(_default, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating(_default, token.Default, $3.Tokens)
                yylex.(*Parser).setFreeFloating(_default, token.CaseSeparator, yylex.(*Parser).GetFreeFloatingToken($3))

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
                $$ = &ast.StmtWhile{ast.Node{}, nil, $1}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodePosition($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' inner_statement_list T_ENDWHILE ';'
            {
                stmtList := &ast.StmtStmtList{ast.Node{}, $2}
                $$ = &ast.StmtAltWhile{ast.Node{}, nil, stmtList}

                // save position
                stmtList.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Cond, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

if_stmt_without_else:
        T_IF '(' expr ')' statement
            {
                $$ = &ast.StmtIf{ast.Node{}, $3, $5, nil, nil}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.If, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   if_stmt_without_else T_ELSEIF '(' expr ')' statement
            {
                _elseIf := &ast.StmtElseIf{ast.Node{}, $4, $6}
                $1.(*ast.StmtIf).ElseIf = append($1.(*ast.StmtIf).ElseIf, _elseIf)

                $$ = $1

                // save position
                _elseIf.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($2, $6)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $6)

                // save comments
                yylex.(*Parser).setFreeFloating(_elseIf, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating(_elseIf, token.ElseIf, $3.Tokens)
                yylex.(*Parser).setFreeFloating(_elseIf, token.Expr, $5.Tokens)

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
                _else := &ast.StmtElse{ast.Node{}, $3}
                $1.(*ast.StmtIf).Else = _else

                $$ = $1

                // save position
                _else.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($2, $3)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(_else, token.Start, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

alt_if_stmt_without_else:
        T_IF '(' expr ')' ':' inner_statement_list
            {
                stmts := &ast.StmtStmtList{ast.Node{}, $6}
                $$ = &ast.StmtAltIf{ast.Node{}, $3, stmts, nil, nil}

                // save position
                stmts.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($6)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($1, $6)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.If, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Cond, $5.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   alt_if_stmt_without_else T_ELSEIF '(' expr ')' ':' inner_statement_list
            {
                stmts := &ast.StmtStmtList{ast.Node{}, $7}
                _elseIf := &ast.StmtAltElseIf{ast.Node{}, $4, stmts}
                $1.(*ast.StmtAltIf).ElseIf = append($1.(*ast.StmtAltIf).ElseIf, _elseIf)

                $$ = $1

                // save position
                stmts.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($7)
                _elseIf.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($2, $7)

                // save comments
                yylex.(*Parser).setFreeFloating(_elseIf, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating(_elseIf, token.ElseIf, $3.Tokens)
                yylex.(*Parser).setFreeFloating(_elseIf, token.Expr, $5.Tokens)
                yylex.(*Parser).setFreeFloating(_elseIf, token.Cond, $6.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

alt_if_stmt:
        alt_if_stmt_without_else T_ENDIF ';'
            {
                $$ = $1

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($3))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   alt_if_stmt_without_else T_ELSE ':' inner_statement_list T_ENDIF ';'
            {
                stmts := &ast.StmtStmtList{ast.Node{}, $4}
                _else := &ast.StmtAltElse{ast.Node{}, stmts}
                $1.(*ast.StmtAltIf).Else = _else

                $$ = $1

                // save position
                stmts.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListPosition($4)
                _else.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodeListPosition($2, $4)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $6)

                // save comments
                yylex.(*Parser).setFreeFloating(_else, token.Start, $2.Tokens)
                yylex.(*Parser).setFreeFloating(_else, token.Else, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $5.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AltEnd, $6.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($6))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

parameter_list:
        non_empty_parameter_list
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

non_empty_parameter_list:
        parameter
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_parameter_list ',' parameter
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

parameter:
        optional_type is_reference is_variadic T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, bytes.TrimLeftFunc($4.Value, isDollar)}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.Parameter{ast.Node{}, $2 != nil, $3 != nil, $1, variable, nil}

                // save position
                identifier.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($4)
                variable.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($4)
                if $1 != nil {
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4)
                } else if $2 != nil {
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($2, $4)
                } else if $3 != nil {
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($3, $4)
                } else {
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($4)
                }

                // save comments
                if $1 != nil {
                    yylex.(*Parser).MoveFreeFloating($1, $$)
                }
                if $2 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.OptionalType, $2.Tokens)
                }
                if $3 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $3.Tokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.Variadic, $4.Tokens)

                // normalize
                if $3 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $$.GetNode().Tokens[token.Variadic]); delete($$.GetNode().Tokens, token.Variadic)
                }
                if $2 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.OptionalType, $$.GetNode().Tokens[token.Ampersand]); delete($$.GetNode().Tokens, token.Ampersand)
                }
                if $1 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Start, $$.GetNode().Tokens[token.OptionalType]); delete($$.GetNode().Tokens, token.OptionalType)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   optional_type is_reference is_variadic T_VARIABLE '=' expr
            {
                identifier := &ast.Identifier{ast.Node{}, bytes.TrimLeftFunc($4.Value, isDollar)}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.Parameter{ast.Node{}, $2 != nil, $3 != nil, $1, variable, $6}

                // save position
                identifier.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($4)
                variable.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($4)
                if $1 != nil {
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $6)
                } else if $2 != nil {
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($2, $6)
                } else if $3 != nil {
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $6)
                } else {
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($4, $6)
                }

                // save comments
                if $1 != nil {
                    yylex.(*Parser).MoveFreeFloating($1, $$)
                }
                if $2 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.OptionalType, $2.Tokens)
                }
                if $3 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $3.Tokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.Variadic, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $5.Tokens)

                // normalize
                if $3 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $$.GetNode().Tokens[token.Variadic]); delete($$.GetNode().Tokens, token.Variadic)
                }
                if $2 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.OptionalType, $$.GetNode().Tokens[token.Ampersand]); delete($$.GetNode().Tokens, token.Ampersand)
                }
                if $1 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Start, $$.GetNode().Tokens[token.OptionalType]); delete($$.GetNode().Tokens, token.OptionalType)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

optional_type:
        /* empty */
            {
                $$ = nil

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
                $$ = &ast.Nullable{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

type:
        T_ARRAY
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CALLABLE
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                $$ = nil

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   ':' type_expr
            {
                $$ = $2;

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Colon, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

argument_list:
        '(' ')'
            {
                $$ = &ast.ArgumentList{ast.Node{}, nil}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ArgumentList, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' non_empty_argument_list possible_comma ')'
            {
                $$ = &ast.ArgumentList{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                if $3 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.ArgumentList, append($3.Tokens, append(yylex.(*Parser).GetFreeFloatingToken($3), $4.Tokens...)...))
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.ArgumentList, $4.Tokens)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_argument_list:
        argument
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_argument_list ',' argument
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

argument:
        expr
            {
                $$ = &ast.Argument{ast.Node{}, false, false, $1}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodePosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ELLIPSIS expr
            {
                $$ = &ast.Argument{ast.Node{}, true, false, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

global_var_list:
        global_var_list ',' global_var
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   global_var
            {
                $$ = []ast.Vertex{$1}

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
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   static_var
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_var:
        T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, bytes.TrimLeftFunc($1.Value, isDollar)}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.StmtStaticVar{ast.Node{}, variable, nil}

                // save position
                identifier.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                variable.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '=' expr
            {
                identifier := &ast.Identifier{ast.Node{}, bytes.TrimLeftFunc($1.Value, isDollar)}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.StmtStaticVar{ast.Node{}, variable, $3}

                // save position
                identifier.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                variable.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

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
                $$ = []ast.Vertex{}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_statement:
        variable_modifiers optional_type property_list ';'
            {
                $$ = &ast.StmtPropertyList{ast.Node{}, $1, $2, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1[0], $$)
                yylex.(*Parser).setFreeFloating($$, token.PropertyList, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   method_modifiers T_CONST class_const_list ';'
            {
                $$ = &ast.StmtClassConstList{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewOptionalListTokensPosition($1, $2, $4)

                // save comments
                if len($1) > 0 {
                    yylex.(*Parser).MoveFreeFloating($1[0], $$)
                    yylex.(*Parser).setFreeFloating($$, token.ModifierList, $2.Tokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Start, $2.Tokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.ConstList, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($4))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE name_list trait_adaptations
            {
                $$ = &ast.StmtTraitUse{ast.Node{}, $2, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   method_modifiers T_FUNCTION returns_ref identifier backup_doc_comment '(' parameter_list ')' return_type method_body
            {
                name := &ast.Identifier{ast.Node{}, $4.Value}
                $$ = &ast.StmtClassMethod{ast.Node{}, $3 != nil, name, $1, $7, $9, $10}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($4)
                if $1 == nil {
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($2, $10)
                } else {
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeListNodePosition($1, $10)
                }

                // save comments
                if len($1) > 0 {
                    yylex.(*Parser).MoveFreeFloating($1[0], $$)
                    yylex.(*Parser).setFreeFloating($$, token.ModifierList, $2.Tokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Start, $2.Tokens)
                }
                if $3 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $4.Tokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $3.Tokens)
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $4.Tokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.Name, $6.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ParameterList, $8.Tokens)
                if $9 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.Params, $9.GetNode().Tokens[token.Colon]); delete($9.GetNode().Tokens, token.Colon)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

name_list:
        name
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   name_list ',' name
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptations:
        ';'
            {
                $$ = &ast.StmtNop{ast.Node{}, }

                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($1))


                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' '}'
            {
                $$ = &ast.StmtTraitAdaptationList{ast.Node{}, nil}

                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AdaptationList, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' trait_adaptation_list '}'
            {
                $$ = &ast.StmtTraitAdaptationList{ast.Node{}, $2}

                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.AdaptationList, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptation_list:
        trait_adaptation
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_adaptation_list trait_adaptation
            {
                $$ = append($1, $2)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_adaptation:
        trait_precedence ';'
            {
                $$ = $1;

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.NameList, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_alias ';'
            {
                $$ = $1;

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Alias, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_precedence:
        absolute_trait_method_reference T_INSTEADOF name_list
            {
                $$ = &ast.StmtTraitUsePrecedence{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeNodeListPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_alias:
        trait_method_reference T_AS T_STRING
            {
                alias := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtTraitUseAlias{ast.Node{}, $1, nil, alias}

                // save position
                alias.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($3)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.Tokens)
                yylex.(*Parser).setFreeFloating(alias, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_method_reference T_AS reserved_non_modifiers
            {
                alias := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtTraitUseAlias{ast.Node{}, $1, nil, alias}

                // save position
                alias.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($3)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.Tokens)
                yylex.(*Parser).setFreeFloating(alias, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_method_reference T_AS member_modifier identifier
            {
                alias := &ast.Identifier{ast.Node{}, $4.Value}
                $$ = &ast.StmtTraitUseAlias{ast.Node{}, $1, $3, alias}

                // save position
                alias.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($4)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.Tokens)
                yylex.(*Parser).setFreeFloating(alias, token.Start, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   trait_method_reference T_AS member_modifier
            {
                $$ = &ast.StmtTraitUseAlias{ast.Node{}, $1, $3, nil}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Ref, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

trait_method_reference:
        identifier
            {
                name := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.StmtTraitMethodRef{ast.Node{}, nil, name}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                target := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.StmtTraitMethodRef{ast.Node{}, $1, target}

                // save position
                target.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($3)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)
                yylex.(*Parser).setFreeFloating(target, token.Start, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

method_body:
        ';' /* abstract method */ 
            {
                $$ = &ast.StmtNop{ast.Node{}, }

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.SemiColon, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' inner_statement_list '}'
            {
                $$ = &ast.StmtStmtList{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

variable_modifiers:
        non_empty_member_modifiers
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VAR
            {
                modifier := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = []ast.Vertex{modifier}

                // save position
                modifier.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(modifier, token.Start, $1.Tokens)

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
                $$ = []ast.Vertex{$1}

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
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PROTECTED
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PRIVATE
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STATIC
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ABSTRACT
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FINAL
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

property_list:
        property_list ',' property
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   property
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

property:
        T_VARIABLE backup_doc_comment
            {
                identifier := &ast.Identifier{ast.Node{}, bytes.TrimLeftFunc($1.Value, isDollar)}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.StmtProperty{ast.Node{}, variable, nil}

                // save position
                identifier.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                variable.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '=' expr backup_doc_comment
            {
                identifier := &ast.Identifier{ast.Node{}, bytes.TrimLeftFunc($1.Value, isDollar)}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.StmtProperty{ast.Node{}, variable, $3}

                // save position
                identifier.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                variable.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_const_list:
        class_const_list ',' class_const_decl
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_const_decl
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_const_decl:
        identifier '=' expr backup_doc_comment
            {
                name := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.StmtConstant{ast.Node{}, name, $3}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

const_decl:
        T_STRING '=' expr backup_doc_comment
            {
                name := &ast.Identifier{ast.Node{}, $1.Value}
                $$ = &ast.StmtConstant{ast.Node{}, name, $3}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

echo_expr_list:
        echo_expr_list ',' echo_expr
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   echo_expr
            {
                $$ = []ast.Vertex{$1}

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
                $$ = nil;

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   non_empty_for_exprs
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

non_empty_for_exprs:
        non_empty_for_exprs ',' expr
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
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
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $8)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Name, $6.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $8.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

new_expr:
        T_NEW class_name_reference ctor_arguments
            {
                if $3 != nil {
                    $$ = &ast.ExprNew{ast.Node{}, $2, $3.(*ast.ArgumentList)}
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $3)
                } else {
                    $$ = &ast.ExprNew{ast.Node{}, $2, nil}
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)
                }

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NEW anonymous_class
            {
                $$ = &ast.ExprNew{ast.Node{}, $2, nil}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

expr_without_variable:
        T_LIST '(' array_pair_list ')' '=' expr
            {
                listNode := &ast.ExprList{ast.Node{}, $3}
                $$ = &ast.ExprAssign{ast.Node{}, listNode, $6}

                // save position
                listNode.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $6)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(listNode, token.List, $2.Tokens)
                yylex.(*Parser).setFreeFloating(listNode, token.ArrayPairList, $4.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $5.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' array_pair_list ']' '=' expr
            {
                shortList := &ast.ExprShortList{ast.Node{}, $2}
                $$ = &ast.ExprAssign{ast.Node{}, shortList, $5}

                // save position
                shortList.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(shortList, token.ArrayPairList, $3.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Var, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' expr
            {
                $$ = &ast.ExprAssign{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable '=' '&' expr
            {
                $$ = &ast.ExprAssignReference{ast.Node{}, $1, $4}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Equal, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLONE expr
            {
                $$ = &ast.ExprClone{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_PLUS_EQUAL expr
            {
                $$ = &ast.ExprAssignPlus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MINUS_EQUAL expr
            {
                $$ = &ast.ExprAssignMinus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MUL_EQUAL expr
            {
                $$ = &ast.ExprAssignMul{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_POW_EQUAL expr
            {
                $$ = &ast.ExprAssignPow{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_DIV_EQUAL expr
            {
                $$ = &ast.ExprAssignDiv{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_CONCAT_EQUAL expr
            {
                $$ = &ast.ExprAssignConcat{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_MOD_EQUAL expr
            {
                $$ = &ast.ExprAssignMod{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_AND_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_OR_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_XOR_EQUAL expr
            {
                $$ = &ast.ExprAssignBitwiseXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_SL_EQUAL expr
            {
                $$ = &ast.ExprAssignShiftLeft{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_SR_EQUAL expr
            {
                $$ = &ast.ExprAssignShiftRight{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_COALESCE_EQUAL expr
            {
                $$ = &ast.ExprAssignCoalesce{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_INC
            {
                $$ = &ast.ExprPostInc{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INC variable
            {
                $$ = &ast.ExprPreInc{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable T_DEC
            {
                $$ = &ast.ExprPostDec{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DEC variable
            {
                $$ = &ast.ExprPreDec{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_BOOLEAN_OR expr
            {
                $$ = &ast.ExprBinaryBooleanOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_BOOLEAN_AND expr
            {
                $$ = &ast.ExprBinaryBooleanAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_OR expr
            {
                $$ = &ast.ExprBinaryLogicalOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_AND expr
            {
                $$ = &ast.ExprBinaryLogicalAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_LOGICAL_XOR expr
            {
                $$ = &ast.ExprBinaryLogicalXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '|' expr
            {
                $$ = &ast.ExprBinaryBitwiseOr{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '&' expr
            {
                $$ = &ast.ExprBinaryBitwiseAnd{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '^' expr
            {
                $$ = &ast.ExprBinaryBitwiseXor{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '.' expr
            {
                $$ = &ast.ExprBinaryConcat{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '+' expr
            {
                $$ = &ast.ExprBinaryPlus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '-' expr
            {
                $$ = &ast.ExprBinaryMinus{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '*' expr
            {
                $$ = &ast.ExprBinaryMul{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_POW expr
            {
                $$ = &ast.ExprBinaryPow{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '/' expr
            {
                $$ = &ast.ExprBinaryDiv{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '%' expr
            {
                $$ = &ast.ExprBinaryMod{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SL expr
            {
                $$ = &ast.ExprBinaryShiftLeft{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SR expr
            {
                $$ = &ast.ExprBinaryShiftRight{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '+' expr %prec T_INC
            {
                $$ = &ast.ExprUnaryPlus{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '-' expr %prec T_INC
            {
                $$ = &ast.ExprUnaryMinus{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '!' expr
            {
                $$ = &ast.ExprBooleanNot{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '~' expr
            {
                $$ = &ast.ExprBitwiseNot{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_IDENTICAL expr
            {
                $$ = &ast.ExprBinaryIdentical{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_NOT_IDENTICAL expr
            {
                $$ = &ast.ExprBinaryNotIdentical{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_EQUAL expr
            {
                $$ = &ast.ExprBinaryEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_NOT_EQUAL expr
            {
                $$ = &ast.ExprBinaryNotEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Equal, yylex.(*Parser).GetFreeFloatingToken($2))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '<' expr
            {
                $$ = &ast.ExprBinarySmaller{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_SMALLER_OR_EQUAL expr
            {
                $$ = &ast.ExprBinarySmallerOrEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '>' expr
            {
                $$ = &ast.ExprBinaryGreater{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_IS_GREATER_OR_EQUAL expr
            {
                $$ = &ast.ExprBinaryGreaterOrEqual{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_SPACESHIP expr
            {
                $$ = &ast.ExprBinarySpaceship{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_INSTANCEOF class_name_reference
            {
                $$ = &ast.ExprInstanceOf{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' expr ')'
            {
                $$ = $2;

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, append($1.Tokens, append(yylex.(*Parser).GetFreeFloatingToken($1), $$.GetNode().Tokens[token.Start]...)...))
                yylex.(*Parser).setFreeFloating($$, token.End, append($$.GetNode().Tokens[token.End], append($3.Tokens, yylex.(*Parser).GetFreeFloatingToken($3)...)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_expr
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '?' expr ':' expr
            {
                $$ = &ast.ExprTernary{ast.Node{}, $1, $3, $5}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $5)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Cond, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.True, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr '?' ':' expr
            {
                $$ = &ast.ExprTernary{ast.Node{}, $1, nil, $4}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Cond, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.True, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_COALESCE expr
            {
                $$ = &ast.ExprBinaryCoalesce{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   internal_functions_in_yacc
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INT_CAST expr
            {
                $$ = &ast.ExprCastInt{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOUBLE_CAST expr
            {
                $$ = &ast.ExprCastDouble{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_STRING_CAST expr
            {
                $$ = &ast.ExprCastString{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ARRAY_CAST expr
            {
                $$ = &ast.ExprCastArray{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_OBJECT_CAST expr
            {
                $$ = &ast.ExprCastObject{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_BOOL_CAST expr
            {
                $$ = &ast.ExprCastBool{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_UNSET_CAST expr
            {
                $$ = &ast.ExprCastUnset{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Cast, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EXIT exit_expr
            {
                var e *ast.ExprExit;
                if $2 != nil {
                    e = $2.(*ast.ExprExit)
                } else {
                    e = &ast.ExprExit{ast.Node{}, false, nil}
                }

                $$ = e

                if (bytes.EqualFold($1.Value, []byte("die"))) {
                    e.Die = true
                }

                // save position
                if $2 == nil {
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                } else {
                    $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)
                }

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '@' expr
            {
                $$ = &ast.ExprErrorSuppress{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   scalar
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '`' backticks_expr '`'
            {
                $$ = &ast.ExprShellExec{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_PRINT expr
            {
                $$ = &ast.ExprPrint{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD
            {
                $$ = &ast.ExprYield{ast.Node{}, nil, nil}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD expr
            {
                $$ = &ast.ExprYield{ast.Node{}, nil, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD expr T_DOUBLE_ARROW expr
            {
                $$ = &ast.ExprYield{ast.Node{}, $2, $4}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_YIELD_FROM expr
            {
                $$ = &ast.ExprYieldFrom{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   inline_function
            {
                $$ = $1;

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
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
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Static, $$.GetNode().Tokens[token.Start]); delete($$.GetNode().Tokens, token.Start)
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens);

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

inline_function:
        T_FUNCTION returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type '{' inner_statement_list '}'
            {
                $$ = &ast.ExprClosure{ast.Node{}, $2 != nil, false, $5, $7, $8, $10}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $11)
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                if $2 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $4.Tokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $2.Tokens)
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $4.Tokens)
                }
                yylex.(*Parser).setFreeFloating($$, token.ParameterList, $6.Tokens)
                if $8 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.LexicalVars, $8.GetNode().Tokens[token.Colon]); delete($8.GetNode().Tokens, token.Colon)
                }
                yylex.(*Parser).setFreeFloating($$, token.ReturnType, $9.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Stmts, $11.Tokens)

                // normalize
                if $8 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.LexicalVars, $$.GetNode().Tokens[token.ReturnType]); delete($$.GetNode().Tokens, token.ReturnType)
                }
                if $7 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Params, $$.GetNode().Tokens[token.LexicalVarList]); delete($$.GetNode().Tokens, token.LexicalVarList)
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FN returns_ref '(' parameter_list ')' return_type backup_doc_comment T_DOUBLE_ARROW expr
            {
                $$ = &ast.ExprArrowFunction{ast.Node{}, $2 != nil, false, $4, $6, $9}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $9)
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                if $2 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $3.Tokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.Function, $2.Tokens)
                    yylex.(*Parser).setFreeFloating($$, token.Ampersand, $3.Tokens)
                };
                yylex.(*Parser).setFreeFloating($$, token.ParameterList, $5.Tokens)
                if $6 != nil {
                    yylex.(*Parser).setFreeFloating($$, token.Params, $6.GetNode().Tokens[token.Colon]); delete($6.GetNode().Tokens, token.Colon)
                };
                yylex.(*Parser).setFreeFloating($$, token.ReturnType, $8.Tokens)

                // normalize
                if $6 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.Params, $$.GetNode().Tokens[token.ReturnType]); delete($$.GetNode().Tokens, token.ReturnType)
                };

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

backup_doc_comment:
        /* empty */
            {
                $$ = yylex.(*Parser).Lexer.GetPhpDocComment()
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
                $$ = nil

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_USE '(' lexical_var_list ')'
            {
                $$ = &ast.ExprClosureUse{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Use, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.LexicalVarList, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

lexical_var_list:
        lexical_var_list ',' lexical_var
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   lexical_var
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

lexical_var:
    T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, bytes.TrimLeftFunc($1.Value, isDollar)}
                $$ = &ast.ExprVariable{ast.Node{}, identifier}

                // save position
                identifier.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, bytes.TrimLeftFunc($2.Value, isDollar)}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.ExprReference{ast.Node{}, variable}

                // save position
                identifier.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($2)
                variable.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(variable, token.Start, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

function_call:
        name argument_list
            {
                $$ = &ast.ExprFunctionCall{ast.Node{}, $1, $2.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
            {
                $$ = &ast.ExprStaticCall{ast.Node{}, $1, $3, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
            {
                $$ = &ast.ExprStaticCall{ast.Node{}, $1, $3, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   callable_expr argument_list
            {
                $$ = &ast.ExprFunctionCall{ast.Node{}, $1, $2.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $2)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

class_name:
        T_STATIC
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                $$ = nil

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '(' optional_expr ')'
            {
                $$ = &ast.ExprExit{ast.Node{}, false, $2};

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Exit, append($1.Tokens, yylex.(*Parser).GetFreeFloatingToken($1)...))
                yylex.(*Parser).setFreeFloating($$, token.Expr, append($3.Tokens, yylex.(*Parser).GetFreeFloatingToken($3)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

backticks_expr:
        /* empty */
            {
                $$ = []ast.Vertex{}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ENCAPSED_AND_WHITESPACE
            {
                part := &ast.ScalarEncapsedStringPart{ast.Node{}, $1.Value}
                $$ = []ast.Vertex{part}

                // save position
                part.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   encaps_list
            {
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

ctor_arguments:
        /* empty */
            {
                $$ = nil

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
                $$ = &ast.ExprArray{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Array, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '[' array_pair_list ']'
            {
                $$ = &ast.ExprShortArray{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.ArrayPairList, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CONSTANT_ENCAPSED_STRING
            {
                $$ = &ast.ScalarString{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

scalar:
        T_LNUMBER
            {
                $$ = &ast.ScalarLnumber{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DNUMBER
            {
                $$ = &ast.ScalarDnumber{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LINE
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FILE
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DIR
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_TRAIT_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_METHOD_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_FUNC_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_NS_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CLASS_C
            {
                $$ = &ast.ScalarMagicConstant{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC
            {
                encapsed := &ast.ScalarEncapsedStringPart{ast.Node{}, $2.Value}
                $$ = &ast.ScalarHeredoc{ast.Node{}, $1.Value, []ast.Vertex{encapsed}}

                // save position
                encapsed.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC T_END_HEREDOC
            {
                $$ = &ast.ScalarHeredoc{ast.Node{}, $1.Value, nil}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '"' encaps_list '"'
            {
                $$ = &ast.ScalarEncapsed{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_START_HEREDOC encaps_list T_END_HEREDOC
            {
                $$ = &ast.ScalarHeredoc{ast.Node{}, $1.Value, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

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
                $$ = &ast.ExprConstFetch{ast.Node{}, $1}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodePosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM identifier
            {
                target := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.ExprClassConstFetch{ast.Node{}, $1, target}

                // save position
                target.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($3)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)
                yylex.(*Parser).setFreeFloating(target, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM identifier
            {
                target := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.ExprClassConstFetch{ast.Node{}, $1, target}

                // save position
                target.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($3)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)
                yylex.(*Parser).setFreeFloating(target, token.Start, $3.Tokens)

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
                $$ = nil

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
                yylex.(*Parser).setFreeFloating($$, token.Start, append($1.Tokens, append(yylex.(*Parser).GetFreeFloatingToken($1), $$.GetNode().Tokens[token.Start]...)...))
                yylex.(*Parser).setFreeFloating($$, token.End, append($$.GetNode().Tokens[token.End], append($3.Tokens, yylex.(*Parser).GetFreeFloatingToken($3)...)...))

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
                yylex.(*Parser).setFreeFloating($$, token.Start, append($1.Tokens, append(yylex.(*Parser).GetFreeFloatingToken($1), $$.GetNode().Tokens[token.Start]...)...))
                yylex.(*Parser).setFreeFloating($$, token.End, append($$.GetNode().Tokens[token.End], append($3.Tokens, yylex.(*Parser).GetFreeFloatingToken($3)...)...))

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
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, append($2.Tokens, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, token.Expr, append($4.Tokens, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   constant '[' optional_expr ']'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, append($2.Tokens, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, token.Expr, append($4.Tokens, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   dereferencable '{' expr '}'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, append($2.Tokens, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, token.Expr, append($4.Tokens, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   dereferencable T_OBJECT_OPERATOR property_name argument_list
            {
                $$ = &ast.ExprMethodCall{ast.Node{}, $1, $3, $4.(*ast.ArgumentList)}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

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
                $$ = &ast.ExprPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

simple_variable:
        T_VARIABLE
            {
                name := &ast.Identifier{ast.Node{}, bytes.TrimLeftFunc($1.Value, isDollar)}
                $$ = &ast.ExprVariable{ast.Node{}, name}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' '{' expr '}'
            {
                $$ = &ast.ExprVariable{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Dollar, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).setFreeFloating($3, token.Start, append($2.Tokens, append(yylex.(*Parser).GetFreeFloatingToken($2), $3.GetNode().Tokens[token.Start]...)...))
                yylex.(*Parser).setFreeFloating($3, token.End, append($3.GetNode().Tokens[token.End], append($4.Tokens, yylex.(*Parser).GetFreeFloatingToken($4)...)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '$' simple_variable
            {
                $$ = &ast.ExprVariable{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Dollar, yylex.(*Parser).GetFreeFloatingToken($1))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

static_member:
        class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = &ast.ExprStaticPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   variable_class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = &ast.ExprStaticPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Name, $2.Tokens)

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
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, append($2.Tokens, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, token.Expr, append($4.Tokens, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_variable '{' expr '}'
            {
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, append($2.Tokens, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, token.Expr, append($4.Tokens, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_variable T_OBJECT_OPERATOR property_name
            {
                $$ = &ast.ExprPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = &ast.ExprStaticPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   new_variable T_PAAMAYIM_NEKUDOTAYIM simple_variable
            {
                $$ = &ast.ExprStaticPropertyFetch{ast.Node{}, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

member_name:
        identifier
            {
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' expr '}'
            {
                $$ = $2;

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, append($1.Tokens, append(yylex.(*Parser).GetFreeFloatingToken($1), $$.GetNode().Tokens[token.Start]...)...))
                yylex.(*Parser).setFreeFloating($$, token.End, append($$.GetNode().Tokens[token.End], append($3.Tokens, yylex.(*Parser).GetFreeFloatingToken($3)...)...))

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
                $$ = &ast.Identifier{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '{' expr '}'
            {
                $$ = $2;
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, append($1.Tokens, append(yylex.(*Parser).GetFreeFloatingToken($1), $$.GetNode().Tokens[token.Start]...)...))
                yylex.(*Parser).setFreeFloating($$, token.End, append($$.GetNode().Tokens[token.End], append($3.Tokens, yylex.(*Parser).GetFreeFloatingToken($3)...)...))

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
                $$ = $1

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

possible_array_pair:
        /* empty */
            {
                $$ = &ast.ExprArrayItem{ast.Node{}, false, nil, nil}

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
                if len($1) == 0 {
                    $1 = []ast.Vertex{&ast.ExprArrayItem{ast.Node{}, false, nil, nil}}
                }

                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   possible_array_pair
            {
                if $1.(*ast.ExprArrayItem).Key == nil && $1.(*ast.ExprArrayItem).Val == nil {
                    $$ = []ast.Vertex{}
                } else {
                    $$ = []ast.Vertex{$1}
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

array_pair:
        expr T_DOUBLE_ARROW expr
            {
                $$ = &ast.ExprArrayItem{ast.Node{}, false, $1, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $3)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr
            {
                $$ = &ast.ExprArrayItem{ast.Node{}, false, nil, $1}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodePosition($1)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_DOUBLE_ARROW '&' variable
            {
                reference := &ast.ExprReference{ast.Node{}, $4}
                $$ = &ast.ExprArrayItem{ast.Node{}, false, $1, reference}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodesPosition($1, $4)
                reference.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($3, $4)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)
                yylex.(*Parser).setFreeFloating(reference, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   '&' variable
            {
                reference := &ast.ExprReference{ast.Node{}, $2}
                $$ = &ast.ExprArrayItem{ast.Node{}, false, nil, reference}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)
                reference.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ELLIPSIS expr
            {
                $$ = &ast.ExprArrayItem{ast.Node{}, true, nil, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   expr T_DOUBLE_ARROW T_LIST '(' array_pair_list ')'
            {
                // TODO: Cannot use list() as standalone expression
                listNode := &ast.ExprList{ast.Node{}, $5}
                $$ = &ast.ExprArrayItem{ast.Node{}, false, $1, listNode}

                // save position
                listNode.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($3, $6)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewNodeTokenPosition($1, $6)

                // save comments
                yylex.(*Parser).MoveFreeFloating($1, $$)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $2.Tokens)
                yylex.(*Parser).setFreeFloating(listNode, token.Start, $3.Tokens)
                yylex.(*Parser).setFreeFloating(listNode, token.List, $4.Tokens)
                yylex.(*Parser).setFreeFloating(listNode, token.ArrayPairList, $6.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_LIST '(' array_pair_list ')'
            {
                // TODO: Cannot use list() as standalone expression
                listNode := &ast.ExprList{ast.Node{}, $3}
                $$ = &ast.ExprArrayItem{ast.Node{}, false, nil, listNode}

                // save position
                listNode.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)
                
                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating(listNode, token.List, $2.Tokens)
                yylex.(*Parser).setFreeFloating(listNode, token.ArrayPairList, $4.Tokens)

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
                encapsed := &ast.ScalarEncapsedStringPart{ast.Node{}, $2.Value}
                $$ = append($1, encapsed)

                // save position
                encapsed.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($2)

                // save comments
                yylex.(*Parser).setFreeFloating(encapsed, token.Start, $2.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   encaps_var
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_ENCAPSED_AND_WHITESPACE encaps_var
            {
                encapsed := &ast.ScalarEncapsedStringPart{ast.Node{}, $1.Value}
                $$ = []ast.Vertex{encapsed, $2}

                // save position
                encapsed.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating(encapsed, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_var:
        T_VARIABLE
            {
                name := &ast.Identifier{ast.Node{}, bytes.TrimLeftFunc($1.Value, isDollar)}
                $$ = &ast.ExprVariable{ast.Node{}, name}

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE '[' encaps_var_offset ']'
            {
                identifier := &ast.Identifier{ast.Node{}, bytes.TrimLeftFunc($1.Value, isDollar)}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, variable, $3}

                // save position
                identifier.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                variable.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Var, append($2.Tokens, yylex.(*Parser).GetFreeFloatingToken($2)...))
                yylex.(*Parser).setFreeFloating($$, token.Expr, append($4.Tokens, yylex.(*Parser).GetFreeFloatingToken($4)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE T_OBJECT_OPERATOR T_STRING
            {
                identifier := &ast.Identifier{ast.Node{}, bytes.TrimLeftFunc($1.Value, isDollar)}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                fetch := &ast.Identifier{ast.Node{}, $3.Value}
                $$ = &ast.ExprPropertyFetch{ast.Node{}, variable, fetch}

                // save position
                identifier.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                variable.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                fetch.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($3)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Var, $2.Tokens)
                yylex.(*Parser).setFreeFloating(fetch, token.Start, $3.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES expr '}'
            {
                variable := &ast.ExprVariable{ast.Node{}, $2}

                $$ = variable

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).setFreeFloating($$, token.End, append($3.Tokens, yylex.(*Parser).GetFreeFloatingToken($3)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}'
            {
                name := &ast.Identifier{ast.Node{}, $2.Value}
                variable := &ast.ExprVariable{ast.Node{}, name}

                $$ = variable

                // save position
                name.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).setFreeFloating($$, token.End, append($3.Tokens, yylex.(*Parser).GetFreeFloatingToken($3)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
            {
                identifier := &ast.Identifier{ast.Node{}, $2.Value}
                variable := &ast.ExprVariable{ast.Node{}, identifier}
                $$ = &ast.ExprArrayDimFetch{ast.Node{}, variable, $4}

                // save position
                identifier.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($2)
                variable.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($2)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $6)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).setFreeFloating($$, token.Var, append($3.Tokens, yylex.(*Parser).GetFreeFloatingToken($3)...))
                yylex.(*Parser).setFreeFloating($$, token.Expr, append($5.Tokens, yylex.(*Parser).GetFreeFloatingToken($5)...))
                yylex.(*Parser).setFreeFloating($$, token.End, append($6.Tokens, yylex.(*Parser).GetFreeFloatingToken($6)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_CURLY_OPEN variable '}'
            {
                $$ = $2;

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, yylex.(*Parser).GetFreeFloatingToken($1))
                yylex.(*Parser).setFreeFloating($$, token.End, append($3.Tokens, yylex.(*Parser).GetFreeFloatingToken($3)...))

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

encaps_var_offset:
        T_STRING
            {
                $$ = &ast.ScalarString{ast.Node{}, $1.Value}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
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
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
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
                    lnumber.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2)
                }
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_VARIABLE
            {
                identifier := &ast.Identifier{ast.Node{}, bytes.TrimLeftFunc($1.Value, isDollar)}
                $$ = &ast.ExprVariable{ast.Node{}, identifier}

                // save position
                identifier.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenPosition($1)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

internal_functions_in_yacc:
        T_ISSET '(' isset_variables possible_comma ')'
            {
                $$ = &ast.ExprIsset{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $5)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Isset, $2.Tokens)
                if $4 == nil {
                    yylex.(*Parser).setFreeFloating($$, token.VarList, $5.Tokens)
                } else {
                    yylex.(*Parser).setFreeFloating($$, token.VarList, append($4.Tokens, append(yylex.(*Parser).GetFreeFloatingToken($4), $5.Tokens...)...))
                }

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EMPTY '(' expr ')'
            {
                $$ = &ast.ExprEmpty{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Empty, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INCLUDE expr
            {
                $$ = &ast.ExprInclude{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_INCLUDE_ONCE expr
            {
                $$ = &ast.ExprIncludeOnce{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_EVAL '(' expr ')'
            {
                $$ = &ast.ExprEval{ast.Node{}, $3}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokensPosition($1, $4)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Eval, $2.Tokens)
                yylex.(*Parser).setFreeFloating($$, token.Expr, $4.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_REQUIRE expr
            {
                $$ = &ast.ExprRequire{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   T_REQUIRE_ONCE expr
            {
                $$ = &ast.ExprRequireOnce{ast.Node{}, $2}

                // save position
                $$.GetNode().Position = yylex.(*Parser).positionBuilder.NewTokenNodePosition($1, $2)

                // save comments
                yylex.(*Parser).setFreeFloating($$, token.Start, $1.Tokens)

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
;

isset_variables:
        isset_variable
            {
                $$ = []ast.Vertex{$1}

                yylex.(*Parser).returnTokenToPool(yyDollar, &yyVAL)
            }
    |   isset_variables ',' isset_variable
            {
                $$ = append($1, $3)

                // save comments
                yylex.(*Parser).setFreeFloating(lastNode($1), token.End, $2.Tokens)

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
